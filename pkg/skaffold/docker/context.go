/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package docker

import (
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
	"github.com/pkg/errors"
)

func contextTarPaths(dockerfilePath string, context string) ([]string, error) {
	return GetDockerfileDependencies(dockerfilePath, context)
}

func CreateDockerTarContext(w io.Writer, dockerfilePath, context string) error {
	paths, err := contextTarPaths(dockerfilePath, context)
	if err != nil {
		return errors.Wrap(err, "getting relative tar paths")
	}
	if err := util.CreateTar(w, context, paths); err != nil {
		return errors.Wrap(err, "creating tar gz")
	}
	return nil
}

func CreateDockerTarGzContext(w io.Writer, dockerfilePath, context string) error {
	paths, err := contextTarPaths(dockerfilePath, context)
	if err != nil {
		return errors.Wrap(err, "getting relative tar paths")
	}
	if err := util.CreateTarGz(w, context, paths); err != nil {
		return errors.Wrap(err, "creating tar gz")
	}
	return nil
}
