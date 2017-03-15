package photon

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *PhotonProvider) ListVolumes() ([]*types.Volume, error) {
	return nil, errors.New("not implemented", nil)
}
