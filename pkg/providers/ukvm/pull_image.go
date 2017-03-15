package ukvm

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *UkvmProvider) PullImage(params types.PullImagePararms) error {

	return errors.New("pulling image not supported for ukvm", nil)
}
