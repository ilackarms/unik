package aws

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers/common"
)

func (p *AwsProvider) GetInstanceLogs(id string) (string, error) {
	instance, err := p.GetInstance(id)
	if err != nil {
		return "", errors.New("retrieving instance "+id, err)
	}
	return common.GetInstanceLogs(instance)
}
