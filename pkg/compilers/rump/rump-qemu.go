package rump

import (
	"fmt"
	"strings"

	"io/ioutil"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/ilackarms/unik/pkg/compilers"
	unikos "github.com/ilackarms/unik/pkg/os"
	"github.com/ilackarms/unik/pkg/types"
)

func CreateImageQemu(kernel string, args string, mntPoints, bakedEnv []string, noCleanup bool) (*types.RawImage, error) {
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
	c = setRumpCmdLine(c, "program.bin", argv, false)

	bootBlk := blk{
		Source:     "dev",
		Path:       "/dev/ld0e",
		FSType:     "blk",
		MountPoint: "/bootpart",
	}
	c.Blk = append(c.Blk, bootBlk)

	res := &types.RawImage{}
	res.RunSpec.Compiler = compilers.Rump

	for i, mntPoint := range mntPoints {
		deviceMapped := fmt.Sprintf("ld%ca", '1'+i)
		blk := blk{
			Source:     "dev",
			Path:       "/dev/" + deviceMapped,
			FSType:     "blk",
			MountPoint: mntPoint,
		}

		c.Blk = append(c.Blk, blk)
		logrus.Debugf("adding mount point to image: %s:%s", mntPoint, deviceMapped)
		res.RunSpec.DeviceMappings = append(res.RunSpec.DeviceMappings,
			types.DeviceMapping{MountPoint: mntPoint, DeviceName: deviceMapped})
	}

	// virtualbox network
	c.Net = &net{
		If:     "vioif0",
		Type:   "inet",
		Method: DHCP,
	}

	cmdline, err := toRumpJson(c)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("writing rump json config: %s", cmdline)

	imgFile, err := compilers.BuildBootableImage(kernel, cmdline, true, noCleanup)
	if err != nil {
		return nil, err
	}

	//copy kernel for qemu
	if err := unikos.CopyFile(kernel, filepath.Join(filepath.Dir(imgFile), "program.bin")); err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(filepath.Join(filepath.Dir(imgFile), "cmdline"), []byte(cmdline), 0644); err != nil {
		return nil, err
	}

	res.LocalImagePath = imgFile
	res.StageSpec.ImageFormat = types.ImageFormat_RAW
	res.RunSpec.DefaultInstanceMemory = 512
	return res, nil

}
