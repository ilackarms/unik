package gcloud

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers/common"
)

func (p *GcloudProvider) GetInstanceLogs(id string) (string, error) {
	instance, err := p.GetInstance(id)
	if err != nil {
		return "", errors.New("retrieving instance "+id, err)
	}
	return common.GetInstanceLogs(instance)
}
