package gcloud

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *GcloudProvider) CreateVolume(params types.CreateVolumeParams) (*types.Volume, error) {
	return nil, errors.New("not yet implemented", nil)
}
func (p *GcloudProvider) CreateEmptyVolume(name string, size int) (*types.Volume, error) {
	return nil, errors.New("not yet implemented", nil)
}
