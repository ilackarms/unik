package ukvm

import (
	"github.com/ilackarms/unik/pkg/types"
)

func (p *UkvmProvider) ListImages() ([]*types.Image, error) {
	images := []*types.Image{}
	for _, image := range p.state.GetImages() {
		images = append(images, image)
	}
	return images, nil
}
