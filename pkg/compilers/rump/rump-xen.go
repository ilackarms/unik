package rump

import (
	"fmt"
	"strings"

	"github.com/ilackarms/unik/pkg/compilers"
	"github.com/ilackarms/unik/pkg/types"
)

func CreateImageXen(kernel, args string, mntPoints, bakedEnv []string, noCleanup bool) (*types.RawImage, error) {
	return createImageXen(kernel, args, mntPoints, bakedEnv, noCleanup, false)
}

func CreateImageXenAddStub(kernel, args string, mntPoints, bakedEnv []string, noCleanup bool) (*types.RawImage, error) {
	return createImageXen(kernel, args, mntPoints, bakedEnv, noCleanup, true)
}

func createImageXen(kernel, args string, mntPoints, bakedEnv []string, noCleanup, addStub bool) (*types.RawImage, error) {
	// create rump config
	var c rumpConfig
	if bakedEnv != nil {
		c.Env = make(map[string]string)
		for i, pair := range bakedEnv {
			c.Env[fmt.Sprintf("env%d", i)] = pair
		}
	}

	argv := []string{}
	if args != "" {
		argv = strings.Split(args, " ")
	}
	c = setRumpCmdLine(c, "program.bin", argv, addStub)

	res := &types.RawImage{}
	volIndex := 0
	// add root -> sda1 mapping
	res.RunSpec.DeviceMappings = append(res.RunSpec.DeviceMappings, types.DeviceMapping{MountPoint: "/", DeviceName: "/dev/sda1"})

	bootBlk := blk{
		Source:     "etfs",
		Path:       "sda1",
		FSType:     "blk",
		MountPoint: "/bootpart",
	}
	c.Blk = append(c.Blk, bootBlk)

	for _, mntPoint := range mntPoints {
		// start from sdb; sda is for root.
		volIndex++
		deviceMapped := fmt.Sprintf("sd%c1", 'a'+volIndex)
		blk := blk{
			Source:     "etfs",
			Path:       deviceMapped,
			FSType:     "blk",
			MountPoint: mntPoint,
		}

		c.Blk = append(c.Blk, blk)
		res.RunSpec.DeviceMappings = append(res.RunSpec.DeviceMappings, types.DeviceMapping{MountPoint: mntPoint, DeviceName: "/dev/" + deviceMapped})
	}

	// aws network
	c.Net = &net{
		If:     "xenif0",
		Cloner: "true",
		Type:   "inet",
		Method: DHCP,
	}

	cmdline, err := toRumpJson(c)
	if err != nil {
		return nil, err
	}
	imgFile, err := compilers.BuildBootableImage(kernel, cmdline, false, noCleanup)

	if err != nil {
		return nil, err
	}

	res.LocalImagePath = imgFile
	res.StageSpec = types.StageSpec{
		ImageFormat:           types.ImageFormat_RAW,
		XenVirtualizationType: types.XenVirtualizationType_Paravirtual,
	}
	res.RunSpec.DefaultInstanceMemory = 1024 // todo: investigate why rump uses so much memory on xen

	return res, nil
}
