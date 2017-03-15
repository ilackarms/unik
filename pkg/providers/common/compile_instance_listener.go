package common

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ilackarms/pkg/errors"
	"github.com/ilackarms/unik/instance-listener/bindata"
	"github.com/ilackarms/unik/pkg/compilers/rump"
	"github.com/ilackarms/unik/pkg/types"
)

func CompileInstanceListener(sourceDir, instanceListenerPrefix, dockerImage string, createImageFunc func(kernel, args string, mntPoints, bakedEnv []string, noCleanup bool) (*types.RawImage, error), enablePersistence bool) (*types.RawImage, error) {
	mainData, err := bindata.Asset("instance-listener/main.go")
	if err != nil {
		return nil, errors.New("reading binary data of instance listener main", err)
	}
	if err := ioutil.WriteFile(filepath.Join(sourceDir, "main.go"), mainData, 0644); err != nil {
		return nil, errors.New("copying contents of instance listener main.go", err)
	}
	godepsData, err := bindata.Asset("instance-listener/Godeps/Godeps.json")
	if err != nil {
		return nil, errors.New("reading binary data of instance listener Godeps", err)
	}
	if err := os.MkdirAll(filepath.Join(sourceDir, "Godeps"), 0755); err != nil {
		return nil, errors.New("creating Godeps dir", err)
	}
	if err := ioutil.WriteFile(filepath.Join(sourceDir, "Godeps", "Godeps.json"), godepsData, 0644); err != nil {
		return nil, errors.New("copying contents of instance listener Godeps.json", err)
	}

	args := "-prefix " + instanceListenerPrefix
	if enablePersistence {
		args = args + " -enablePersistence"
	}

	params := types.CompileImageParams{
		SourcesDir: sourceDir,
		Args:       args,
	}

	if enablePersistence {
		params.MntPoints = []string{"/data"}
	}

	rumpGoCompiler := &rump.RumpGoCompiler{
		RumCompilerBase: rump.RumCompilerBase{
			DockerImage: dockerImage,
			CreateImage: createImageFunc,
		},
		BootstrapType: rump.BootstrapTypeNoStub,
	}
	return rumpGoCompiler.CompileRawImage(params)
}
