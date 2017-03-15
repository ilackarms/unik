package osv

import (
	"github.com/Sirupsen/logrus"
	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/pkg/compilers"
	"github.com/ilackarms/unik/pkg/types"
	unikutil "github.com/ilackarms/unik/pkg/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type OSvJavaCompiler struct {
	ImageFinisher ImageFinisher
}

// javaProjectConfig defines available inputs
type javaProjectConfig struct {
	MainFile    string `yaml:"main_file"`
	RuntimeArgs string `yaml:"runtime_args"`
	BuildCmd    string `yaml:"build_command"`
}

func (r *OSvJavaCompiler) CompileRawImage(params types.CompileImageParams) (*types.RawImage, error) {
	sourcesDir := params.SourcesDir

	var config javaProjectConfig
	data, err := ioutil.ReadFile(filepath.Join(sourcesDir, "manifest.yaml"))
	if err != nil {
		return nil, errors.New("failed to read manifest.yaml file", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.New("failed to parse yaml manifest.yaml file", err)
	}

	container := unikutil.NewContainer("compilers-osv-java").WithVolume("/dev", "/dev").WithVolume(sourcesDir+"/", "/project_directory")
	var args []string
	if r.ImageFinisher.UseEc2() {
		args = append(args, "-ec2")
	}

	args = append(args, "-main_file", config.MainFile)
	args = append(args, "-args", params.Args)
	if config.BuildCmd != "" {
		args = append(args, "-buildCmd", config.BuildCmd)
	}
	if len(config.RuntimeArgs) > 0 {
		args = append(args, "-runtime", config.RuntimeArgs)
	}

	logrus.WithFields(logrus.Fields{
		"args": args,
	}).Debugf("running compilers-osv-java container")

	if err := container.Run(args...); err != nil {
		return nil, errors.New("failed running compilers-osv-java on "+sourcesDir, err)
	}

	// And finally bootstrap.
	convertParams := FinishParams{
		CompileParams:    params,
		CapstanImagePath: filepath.Join(sourcesDir, "boot.qcow2"),
	}
	return r.ImageFinisher.FinishImage(convertParams)
}

func (r *OSvJavaCompiler) Usage() *compilers.CompilerUsage {
	return &compilers.CompilerUsage{
		PrepareApplication: "Compile your Java application into a fat jar.",
		ConfigurationFiles: map[string]string{
			"/manifest.yaml": "main_file: <relative-path-to-your-fat-jar>",
		},
	}
}
