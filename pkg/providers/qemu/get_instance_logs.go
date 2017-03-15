package qemu

import "github.com/ilackarms/pkg/errors"

func (p *QemuProvider) GetInstanceLogs(id string) (string, error) {
	return "", errors.New("not supported", nil)
}
