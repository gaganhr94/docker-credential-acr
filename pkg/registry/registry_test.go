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

package registry

import "testing"

func TestIsACRRegistry(t *testing.T) {
	tests := []struct {
		name      string
		serverURL string
		want      bool
	}{
		// Positive cases
		{name: "azurecr.io", serverURL: "myregistry.azurecr.io", want: true},
		{name: "azurecr.cn", serverURL: "myregistry.azurecr.cn", want: true},
		{name: "azurecr.de", serverURL: "myregistry.azurecr.de", want: true},
		{name: "azurecr.us", serverURL: "myregistry.azurecr.us", want: true},
		{name: "mcr", serverURL: "mcr.microsoft.com", want: true},
		{name: "nested subdomain", serverURL: "my.nested.registry.azurecr.io", want: true},

		// Negative cases
		{name: "azurecr.me", serverURL: "myregistry.azurecr.me", want: false},
		{name: "docker hub", serverURL: "index.docker.io", want: false},
		{name: "gcr", serverURL: "gcr.io", want: false},
		{name: "ecr", serverURL: "123456789.dkr.ecr.us-east-1.amazonaws.com", want: false},
		{name: "localhost", serverURL: "localhost", want: false},
		{name: "localhost with port", serverURL: "localhost:5000", want: false},
		{name: "ip address", serverURL: "192.168.1.1", want: false},
		{name: "empty", serverURL: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsACRRegistry(tt.serverURL)
			if got != tt.want {
				t.Errorf("IsACRRegistry(%q) = %v, want %v", tt.serverURL, got, tt.want)
			}
		})
	}
}
