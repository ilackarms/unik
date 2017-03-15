package openstack

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *OpenstackProvider) ListVolumes() ([]*types.Volume, error) {
	return nil, errors.New("not yet supportded for openstack", nil)
}
