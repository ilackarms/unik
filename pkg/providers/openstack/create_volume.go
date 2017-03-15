package openstack

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *OpenstackProvider) CreateVolume(params types.CreateVolumeParams) (_ *types.Volume, err error) {
	return nil, errors.New("not yet supportded for openstack", nil)
}
