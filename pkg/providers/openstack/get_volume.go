package openstack

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *OpenstackProvider) GetVolume(nameOrIdPrefix string) (*types.Volume, error) {
	return nil, errors.New("not yet supportded for openstack", nil)
}
