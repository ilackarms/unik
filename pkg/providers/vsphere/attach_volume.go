package vsphere

import (
	"github.com/Sirupsen/logrus"
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers/common"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *VsphereProvider) AttachVolume(id, instanceId, mntPoint string) error {
	volume, err := p.GetVolume(id)
	if err != nil {
		return errors.New("retrieving volume "+id, err)
	}
	instance, err := p.GetInstance(instanceId)
	if err != nil {
		return errors.New("retrieving instance "+instanceId, err)
	}
	image, err := p.GetImage(instance.ImageId)
	if err != nil {
		return errors.New("retrieving image for instance", err)
	}
	controllerPort, err := common.GetControllerPortForMnt(image, mntPoint)
	if err != nil {
		return errors.New("getting controller port for mnt point", err)
	}
	logrus.Infof("attaching %s to %s on controller port %v", volume.Id, instance.Id, controllerPort)
	if err := p.getClient().AttachDisk(instance.Id, getVolumeDatastorePath(volume.Name), controllerPort, image.RunSpec.StorageDriver); err != nil {
		return errors.New("attaching disk to vm", err)
	}
	if err := p.state.ModifyVolumes(func(volumes map[string]*types.Volume) error {
		volume, ok := volumes[volume.Id]
		if !ok {
			return errors.New("no record of "+volume.Id+" in the state", nil)
		}
		volume.Attachment = instance.Id
		return nil
	}); err != nil {
		return errors.New("modifying volumes in state", err)
	}
	return nil
}
