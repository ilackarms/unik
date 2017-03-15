package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/ilackarms/pkg/errors"
)

func (p *AwsProvider) DeleteImage(id string, force bool) error {
	image, err := p.GetImage(id)
	if err != nil {
		return errors.New("retrieving image", err)
	}
	instances, err := p.ListInstances()
	if err != nil {
		return errors.New("retrieving list of instances", err)
	}
	for _, instance := range instances {
		if instance.ImageId == image.Id {
			if !force {
				return errors.New("instance "+instance.Id+" found which uses image "+image.Id+"; try again with force=true", nil)
			} else {
				err = p.DeleteInstance(instance.Id, true)
				if err != nil {
					return errors.New("failed to delete instance "+instance.Id+" which is using image "+image.Id, err)
				}
			}
		}
	}

	ec2svc := p.newEC2()
	deleteAmiParam := &ec2.DeregisterImageInput{
		ImageId: aws.String(image.Id),
	}
	_, err = ec2svc.DeregisterImage(deleteAmiParam)
	if err != nil {
		return errors.New("failed deleting image "+image.Id, err)
	}

	snap, err := getSnapshotForImage(ec2svc, image.Id)
	if err != nil {
		return err
	}
	deleteSnapshotParam := &ec2.DeleteSnapshotInput{
		SnapshotId: aws.String(*snap.SnapshotId),
	}
	if _, err := ec2svc.DeleteSnapshot(deleteSnapshotParam); err != nil {
		return errors.New("failed deleting snapshot "+*snap.SnapshotId, err)
	}
	deleteVolumeParam := &ec2.DeleteVolumeInput{
		VolumeId: aws.String(*snap.VolumeId),
	}
	if _, err := ec2svc.DeleteVolume(deleteVolumeParam); err != nil {
		return errors.New("failed deleting volumme "+*snap.VolumeId, err)
	}
	return p.state.RemoveImage(image)
}

func getSnapshotForImage(ec2svc *ec2.EC2, imageId string) (*ec2.Snapshot, error) {
	describeSnapshotsOutput, err := ec2svc.DescribeSnapshots(&ec2.DescribeSnapshotsInput{})
	if err != nil {
		return nil, errors.New("getting ec2 snapshot list", err)
	}

	for _, snapshot := range describeSnapshotsOutput.Snapshots {
		for _, tag := range snapshot.Tags {
			if *tag.Key == UNIK_IMAGE_ID && *tag.Value == imageId {
				return snapshot, nil
			}
		}
	}
	return nil, errors.New("snapshot for image "+imageId+" not found", nil)
}
