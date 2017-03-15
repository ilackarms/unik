package aws

import (
	"github.com/ilackarms/unik/pkg/providers/common"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *AwsProvider) GetInstance(nameOrIdPrefix string) (*types.Instance, error) {
	return common.GetInstance(p, nameOrIdPrefix)
}
