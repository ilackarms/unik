package openstack

import "github.com/ilackarms/pkg/errors"

func (p *OpenstackProvider) DetachVolume(id string) error {
	return errors.New("not yet supportded for openstack", nil)
}
