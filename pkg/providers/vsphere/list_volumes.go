package vsphere

import (
	"github.com/ilackarms/unik/pkg/types"
)

func (p *VsphereProvider) ListVolumes() ([]*types.Volume, error) {
	volumes := []*types.Volume{}
	for _, volume := range p.state.GetVolumes() {
		volumes = append(volumes, volume)
	}
	return volumes, nil
}
