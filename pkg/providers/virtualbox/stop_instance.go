package virtualbox

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/providers/virtualbox/virtualboxclient"
)

func (p *VirtualboxProvider) StopInstance(id string) error {
	instance, err := p.GetInstance(id)
	if err != nil {
		return errors.New("retrieving instance "+id, err)
	}
	if err := virtualboxclient.PowerOffVm(instance.Id); err != nil {
		return errors.New("failed to stop instance "+instance.Id, err)
	}
	return nil
}
