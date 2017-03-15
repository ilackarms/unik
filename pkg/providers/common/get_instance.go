package common

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers"
	"github.com/ilackarms/unik/pkg/types"
	"strings"
)

func GetInstance(p providers.Provider, nameOrIdPrefix string) (*types.Instance, error) {
	instances, err := p.ListInstances()
	if err != nil {
		return nil, errors.New("retrieving instance list", err)
	}
	for _, instance := range instances {
		if strings.Contains(instance.Id, nameOrIdPrefix) || strings.Contains(instance.Name, nameOrIdPrefix) {
			return instance, nil
		}
	}
	return nil, errors.New("instance with name or id containing '"+nameOrIdPrefix+"' not found", nil)
}
