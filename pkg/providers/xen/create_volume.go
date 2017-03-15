package xen

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *XenProvider) CreateVolume(params types.CreateVolumeParams) (_ *types.Volume, err error) {
	if _, volumeErr := p.GetImage(params.Name); volumeErr == nil {
		return nil, errors.New("volume already exists", nil)
	}

	volumePath := getVolumePath(params.Name)
	if err := os.MkdirAll(filepath.Dir(volumePath), 0755); err != nil {
		return nil, errors.New("creating directory for volume file", err)
	}
	defer func() {
		if err != nil {
			if params.NoCleanup {
				logrus.Warnf("because --no-cleanup flag was provided, not cleaning up failed volume %s at %s", params.Name, volumePath)
			} else {
				os.RemoveAll(filepath.Dir(volumePath))
			}
		}
	}()
	logrus.WithField("raw-image", params.ImagePath).Infof("creating volume from raw image")

	rawImageFile, err := os.Stat(params.ImagePath)
	if err != nil {
		return nil, errors.New("statting raw image file", err)
	}
	sizeMb := rawImageFile.Size() >> 20

	if err := os.Rename(params.ImagePath, volumePath); err != nil {
		return nil, errors.New("copying raw image from "+params.ImagePath+"to "+volumePath, err)
	}

	volume := &types.Volume{
		Id:             params.Name,
		Name:           params.Name,
		SizeMb:         sizeMb,
		Attachment:     "",
		Infrastructure: types.Infrastructure_XEN,
		Created:        time.Now(),
	}

	err = p.state.ModifyVolumes(func(volumes map[string]*types.Volume) error {
		volumes[volume.Id] = volume
		return nil
	})
	if err != nil {
		return nil, errors.New("modifying volume map in state", err)
	}
	return volume, nil

}
