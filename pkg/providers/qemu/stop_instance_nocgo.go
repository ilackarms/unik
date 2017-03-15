// +build !cgo

package qemu

import "github.com/ilackarms/pkg/errors"

func (p *QemuProvider) StopInstance(id string) error {

	return errors.New("Stopping qemu instance is not supported without cgo", nil)
}
