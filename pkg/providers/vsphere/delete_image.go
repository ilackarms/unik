package vsphere

import (
	"github.com/Sirupsen/logrus"
	"github.com/ilackarms/pkg/errors"
)

func (p *VsphereProvider) DeleteImage(id string, force bool) error {
	image, err := p.GetImage(id)
	if err != nil {
		return errors.New("retrieving image", err)
	}
	instances, err := p.ListInstances()
	if err != nil {
		return errors.New("retrieving list of instances", err)
	}
	for _, instance := range instances {
		if instance.ImageId == image.Id {
			if !force {
				return errors.New("instance "+instance.Id+" found which uses image "+image.Id+"; try again with force=true", nil)
			} else {
				logrus.Warnf("deleting instance %s which belongs to image %s", instance.Id, image.Id)
				err = p.DeleteInstance(instance.Id, true)
				if err != nil {
					return errors.New("failed to delete instance "+instance.Id+" which is using image "+image.Id, err)
				}
			}
		}
	}

	imageDir := getImageDatastoreDir(image.Name)
	logrus.Infof("deleting image file at %s", imageDir)
	if err := p.getClient().Rmdir(imageDir); err != nil {
		return errors.New("deleting image file at "+imageDir, err)
	}
	return p.state.RemoveImage(image)
}
