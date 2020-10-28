package protos

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/dhiguero/grpc-proto-manager/internal/pkg/files"
	"github.com/rs/zerolog/log"
)

// DockerCmdProvider is a proto generator based on issuing docker commands. Future
// implementations will rely on the docker library.
type DockerCmdProvider struct {
}

// NewDockerCmdGenerator uses an external command to launch the docker container with the proto tools.
func NewDockerCmdGenerator() (Generator, error) {
	log.Debug().Msg("Using DockerCmd proto generator")
	return &DockerCmdProvider{}, nil
}

// Generate a set of proto stubs in a given language.
func (dcp *DockerCmdProvider) Generate(rootPath string, targetName string, generatedPath string, language string) error {
	// Based on the documentation available at: https://github.com/namely/docker-protoc

	log.Debug().Str("rootPath", rootPath).Str("targetName", targetName).Str("generatedPath", generatedPath).Str("language", language).Msg("generating protos")

	cmdArgs := []string{
		"run",
		"-v", fmt.Sprintf("%s:/defs", rootPath), // source proto definition. This should be the root so imports work :)
		"namely/protoc-all:1.32_4", // Image, maybe move this as a constant or config value.
		"-l", language,             // Target language
		"-d", targetName, // Directory to take protos from
		"-i", ".", // Include local path
		"-o", "generated", // Path where the resulting code is stored.
		// Extra options from the available arguments
		"--with-gateway",   //Generate grpc-gateway files (experimental)
		"--with-validator", // Generate validations for (go gogo cpp java python)
	}

	cmd := exec.Command("docker", cmdArgs...)
	log.Debug().Interface("cmd", cmd).Msg("docker generation cmd")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unable to generate protos for %s due to %w: %s", targetName, err, string(stdoutStderr))
	}
	log.Debug().Msg("execution finished")
	log.Debug().Str("output", string(stdoutStderr)).Msg("protos successfully generated")
	err = dcp.copyAllSourceFiles(path.Join(rootPath, targetName), generatedPath)
	if err != nil {
		return fmt.Errorf("unable to copy source files: %w", err)
	}
	return dcp.moveGeneratedFiles(path.Join(rootPath, "generated"), generatedPath)
}

// copyAllSourceFiles copies all source files into the generated path so it contains everything that will be uploaded
func (dcp *DockerCmdProvider) copyAllSourceFiles(source string, generatedPath string) error {
	toCopy := make(map[string]string, 0)
	filepath.Walk(source, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			toCopy[currentPath] = info.Name()
		}

		return nil
	})
	for filePath, fileName := range toCopy {
		log.Debug().Str("toCopy", filePath).Str("fileName", fileName).Msg("moving file")
		err := files.CopyFile(filePath, path.Join(generatedPath, fileName))
		if err != nil {
			return err
		}
	}
	return nil
}

// moveGenerateFiles moves the generated files into the temp directory.
func (dcp *DockerCmdProvider) moveGeneratedFiles(rootPath string, generatedPath string) error {
	log.Debug().Str("rootPath", rootPath).Str("generatedPath", generatedPath).Msg("moving generated content")

	// Find the generated files. This is a bit of a hack since the generated structure depends on language specs.
	// So we find all the files
	toCopy := make(map[string]string, 0)
	filepath.Walk(rootPath, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			toCopy[currentPath] = info.Name()
		}

		return nil
	})

	for filePath, fileName := range toCopy {
		log.Debug().Str("toCopy", filePath).Str("fileName", fileName).Msg("moving file")
		err := os.Rename(filePath, path.Join(generatedPath, fileName))
		if err != nil {
			return err
		}
	}

	// Cleanup the temporal generated directory.
	err := os.RemoveAll(rootPath)
	if err != nil {
		return err
	}

	return nil
}
