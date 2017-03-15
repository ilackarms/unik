package photon

import (
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
)

func (p *PhotonProvider) DeleteInstance(id string, force bool) error {
	instance, err := p.GetInstance(id)
	if err != nil {
		return errors.New("retrieving instance "+id, err)
	}
	if instance.State == types.InstanceState_Running {
		if !force {
			return errors.New("instance "+instance.Id+"is still running. try again with --force or power off instance first", err)
		} else {
			p.StopInstance(instance.Id)
		}
	}

	task, err := p.client.VMs.Delete(instance.Id)
	if err != nil {
		return errors.New("Delete vm", err)
	}

	task, err = p.waitForTaskSuccess(task)
	if err != nil {
		return errors.New("Delete vm", err)
	}
	return p.state.RemoveInstance(instance)
}
