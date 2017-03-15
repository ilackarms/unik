package openstack

import (
	"github.com/ilackarms/pkg/errors"
)

func (p *OpenstackProvider) GetInstanceLogs(id string) (string, error) {
	return "", errors.New("not yet supportded for openstack", nil)
}
