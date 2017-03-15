package common

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/djannot/aws-sdk-go/aws"
	"github.com/djannot/aws-sdk-go/aws/session"
	"github.com/djannot/aws-sdk-go/service/s3"
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/config"
	"github.com/ilackarms/unik/pkg/types"
	"github.com/layer-x/layerx-commons/lxhttpclient"
	"io"
	"net/http"
	"os"
)

const (
	unik_hub_region = "us-east-1"
	unik_hub_bucket = "unik-hub"
	unik_image_info = "Unik-Image-Info"
)

func PullImage(config config.HubConfig, imageName string, writer io.Writer) (*types.Image, error) {
	//to trigger modified djannot/aws-sdk
	os.Setenv("S3_AUTH_PROXY_URL", config.URL)

	//search available images, get user for image name
	resp, body, err := lxhttpclient.Get(config.URL, "/images", nil)
	if err != nil {
		return nil, errors.New("performing GET request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("failed GETting image list status %v: %s", resp.StatusCode, string(body)), nil)
	}
	var images []*types.UserImage
	if err := json.Unmarshal(body, &images); err != nil {
		logrus.Fatal(err)
	}
	var user string
	for _, image := range images {
		if image.Name == imageName {
			user = image.Owner
			break
		}
	}
	if user == "" {
		return nil, errors.New("could not find image "+imageName, nil)
	}

	metadata, err := s3Download(imageKey(user, imageName), config.Password, writer)
	if err != nil {
		return nil, errors.New("downloading image", err)
	}
	var image types.Image
	if err := json.Unmarshal([]byte(metadata), &image); err != nil {
		return nil, errors.New("unmarshalling metadata for image", err)
	}
	logrus.Infof("downloaded image %v", image)
	return &image, nil
}

func PushImage(config config.HubConfig, image *types.Image, imagePath string) error {
	//to trigger modified djannot/aws-sdk
	os.Setenv("S3_AUTH_PROXY_URL", config.URL)
	metadata, err := json.Marshal(image)
	if err != nil {
		return errors.New("converting image metadata to json", err)
	}
	//upload image
	reader, err := os.Open(imagePath)
	if err != nil {
		return errors.New("opening file", err)
	}
	defer reader.Close()
	fileInfo, err := reader.Stat()
	if err != nil {
		return errors.New("getting file info", err)
	}
	if err := s3Upload(config, imageKey(config.Username, image.Name), string(metadata), reader, fileInfo.Size()); err != nil {
		return errors.New("uploading image file", err)
	}
	logrus.Infof("Image %v pushed to %s", image, config.URL)
	return nil
}

func RemoteDeleteImage(config config.HubConfig, imageName string) error {
	//to trigger modified djannot/aws-sdk
	os.Setenv("S3_AUTH_PROXY_URL", config.URL)
	if err := s3Delete(config, imageKey(config.Username, imageName)); err != nil {
		return errors.New("deleting image file", err)
	}
	logrus.Infof("Image %v deleted from %s", imageName, config.URL)
	return nil
}

func s3Download(key, password string, writer io.Writer) (string, error) {
	params := &s3.GetObjectInput{
		Bucket:   aws.String(unik_hub_bucket),
		Key:      aws.String(key),
		Password: aws.String(password),
	}
	result, err := s3.New(session.New(&aws.Config{Region: aws.String(unik_hub_region)})).GetObject(params)
	if err != nil {
		return "", errors.New("failed to download from s3", err)
	}
	n, err := io.Copy(writer, result.Body)
	if err != nil {
		return "", errors.New("copying image bytes", err)
	}
	logrus.Infof("downloaded %v bytes", n)
	if result.Metadata[unik_image_info] == nil {
		return "", errors.New(fmt.Sprintf(unik_image_info+" was empty. full metadata: %+v", result.Metadata), nil)
	}
	return *result.Metadata[unik_image_info], nil
}

func s3Upload(config config.HubConfig, key, metadata string, body io.ReadSeeker, length int64) error {
	params := &s3.PutObjectInput{
		Body:   body,
		Bucket: aws.String(unik_hub_bucket),
		Key:    aws.String(key),
		Metadata: map[string]*string{
			"unik-password": aws.String(config.Password),
			"unik-email":    aws.String(config.Username),
			"unik-access":   aws.String("public"),
			unik_image_info: aws.String(metadata),
		},
	}
	result, err := s3.New(session.New(&aws.Config{Region: aws.String(unik_hub_region)})).PutObject(params)
	if err != nil {
		return errors.New("uploading image to s3 backend", err)
	}
	logrus.Infof("uploaded %v bytes: %v", length, result)
	return nil
}

// unik hub has to do it itself to validate user
func s3Delete(config config.HubConfig, key string) error {
	deleteMessage := struct {
		Username string `json:"user"`
		Password string `json:"pass"`
		Key      string `json:"key"`
	}{
		Username: config.Username,
		Password: config.Password,
		Key:      key,
	}
	resp, body, err := lxhttpclient.Post(config.URL, "/delete_image", nil, deleteMessage)
	if err != nil {
		return errors.New("failed to perform delete request", err)
	}
	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("expected status code 204, got %v: %s", resp.StatusCode, string(body)), nil)
	}
	return nil
}

func imageKey(username, imageName string) string {
	return "/" + username + "/" + imageName + "/latest" //TODO: support image versioning
}
