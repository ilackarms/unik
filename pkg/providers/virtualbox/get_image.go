package virtualbox

import (
	"github.com/ilackarms/unik/pkg/providers/common"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *VirtualboxProvider) GetImage(nameOrIdPrefix string) (*types.Image, error) {
	return common.GetImage(p, nameOrIdPrefix)
}
