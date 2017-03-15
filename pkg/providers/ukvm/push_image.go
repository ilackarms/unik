package ukvm

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *UkvmProvider) PushImage(params types.PushImagePararms) error {
	return errors.New("pushing image not supported for ukvm", nil)
}
