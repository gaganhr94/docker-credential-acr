// Copyright 2026 docker-credential-acr authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-credential-acr-env",
	Short: "Docker credential helper for Azure Container Registry",
	Long: `A Docker credential helper that provides authentication for Azure Container Registry
using environment-based Azure credentials (service principal, managed identity,
workload identity, Azure CLI, etc.).

Configure Docker to use this helper by adding the following to ~/.docker/config.json:

  {
    "credHelpers": {
      "myregistry.azurecr.io": "acr-env"
    }
  }`,
}

// Execute runs the root command.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
