// +build !cgo

package ukvm

import "github.com/ilackarms/pkg/errors"

func (p *UkvmProvider) StopInstance(id string) error {

	return errors.New("Stopping ukvm instance is not supported without cgo", nil)
}
