package common

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/types"
	"github.com/layer-x/layerx-commons/lxhttpclient"
)

const UnikLogsPort = 9967

func GetInstanceLogs(instance *types.Instance) (string, error) {
	if instance.IpAddress == "" {
		return "", errors.New("instance has not been assigned a public ip address", nil)
	}
	_, body, err := lxhttpclient.Get(instance.IpAddress+fmt.Sprintf(":%v", UnikLogsPort), "/logs", nil)
	if err != nil {
		return "", errors.New("faiiled to connect to instance at "+instance.IpAddress+" for logs", err)
	}
	logrus.WithFields(logrus.Fields{"response-length": len(body), "instance": instance}).Debugf("received stdout from instance")
	return string(body), nil
}
