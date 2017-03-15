package common

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers"
	"github.com/ilackarms/unik/pkg/types"
	"strings"
)

func GetImage(p providers.Provider, nameOrIdPrefix string) (*types.Image, error) {
	images, err := p.ListImages()
	if err != nil {
		return nil, errors.New("retrieving image list", err)
	}
	for _, image := range images {
		if strings.Contains(image.Id, nameOrIdPrefix) || strings.Contains(image.Name, nameOrIdPrefix) {
			return image, nil
		}
	}
	return nil, errors.New("image with name or id containing '"+nameOrIdPrefix+"' not found", nil)
}
