package photon

import (
	"github.com/ilackarms/unik/pkg/providers/common"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *PhotonProvider) GetVolume(nameOrIdPrefix string) (*types.Volume, error) {
	return common.GetVolume(p, nameOrIdPrefix)
}
