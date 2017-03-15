package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/ilackarms/pkg/errors"
)

func (p *AwsProvider) DeleteVolume(id string, force bool) error {
	volume, err := p.GetVolume(id)
	if err != nil {
		return errors.New("retrieving volume "+id, err)
	}
	if volume.Attachment != "" {
		if force {
			if err := p.DetachVolume(volume.Id); err != nil {
				return errors.New("detaching volume for deletion", err)
			} else {
				return errors.New("volume "+volume.Id+" is attached to instance."+volume.Attachment+", try again with --force or detach volume first", err)
			}
		}
	}
	param := &ec2.DeleteVolumeInput{
		VolumeId: aws.String(volume.Id),
	}
	_, err = p.newEC2().DeleteVolume(param)
	if err != nil {
		return errors.New("failed to terminate volume "+volume.Id, err)
	}
	return p.state.RemoveVolume(volume)
}
