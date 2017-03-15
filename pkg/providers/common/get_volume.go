package common

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers"
	"github.com/ilackarms/unik/pkg/types"
	"strings"
)

func GetVolume(p providers.Provider, nameOrIdPrefix string) (*types.Volume, error) {
	volumes, err := p.ListVolumes()
	if err != nil {
		return nil, errors.New("retrieving volume list", err)
	}
	for _, volume := range volumes {
		if strings.Contains(volume.Id, nameOrIdPrefix) || strings.Contains(volume.Name, nameOrIdPrefix) {
			return volume, nil
		}
	}
	return nil, errors.New("volume with name or id containing '"+nameOrIdPrefix+"' not found", nil)
}
