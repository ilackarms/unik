package types

import "github.com/ilackarms/unik/pkg/config"

type RunInstanceParams struct {
	Name                 string
	ImageId              string
	MntPointsToVolumeIds map[string]string
	Env                  map[string]string
	InstanceMemory       int
	NoCleanup            bool
	DebugMode            bool
}

type StageImageParams struct {
	Name      string
	RawImage  *RawImage
	Force     bool
	NoCleanup bool
}

type CreateVolumeParams struct {
	Name      string
	ImagePath string
	NoCleanup bool
}

type CompileImageParams struct {
	SourcesDir string
	Args       string
	MntPoints  []string
	NoCleanup  bool
	SizeMB     int
}

type PullImagePararms struct {
	Config    config.HubConfig
	ImageName string
	Force     bool
}

type PushImagePararms struct {
	Config    config.HubConfig
	ImageName string
}

type RemoteDeleteImagePararms struct {
	Config    config.HubConfig
	ImageName string
}
