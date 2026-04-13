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

package credhelper

import (
	"testing"
)

func TestACRCredHelper_Get_NonACR(t *testing.T) {
	helper := NewACRCredentialsHelper()

	tests := []struct {
		name      string
		serverURL string
	}{
		{name: "docker hub", serverURL: "index.docker.io"},
		{name: "gcr", serverURL: "gcr.io"},
		{name: "ecr", serverURL: "123456789.dkr.ecr.us-east-1.amazonaws.com"},
		{name: "ghcr", serverURL: "ghcr.io"},
		{name: "localhost", serverURL: "localhost:5000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := helper.Get(tt.serverURL)
			if err == nil {
				t.Errorf("Get(%q) should return error for non-ACR registry", tt.serverURL)
			}
		})
	}
}

func TestACRCredHelper_Add(t *testing.T) {
	helper := NewACRCredentialsHelper()
	err := helper.Add(nil)
	if err == nil {
		t.Error("Add should return error")
	}
}

func TestACRCredHelper_Delete(t *testing.T) {
	helper := NewACRCredentialsHelper()
	err := helper.Delete("")
	if err == nil {
		t.Error("Delete should return error")
	}
}

func TestACRCredHelper_List(t *testing.T) {
	helper := NewACRCredentialsHelper()
	_, err := helper.List()
	if err == nil {
		t.Error("List should return error")
	}
}

func TestKeychainHelper_Get_NonACR(t *testing.T) {
	helper := NewKeychainHelper()

	tests := []struct {
		name      string
		serverURL string
	}{
		{name: "docker hub", serverURL: "index.docker.io"},
		{name: "gcr", serverURL: "gcr.io"},
		{name: "ecr", serverURL: "123456789.dkr.ecr.us-east-1.amazonaws.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, pass, err := helper.Get(tt.serverURL)
			if err != nil {
				t.Errorf("KeychainHelper.Get(%q) returned error for non-ACR registry: %v", tt.serverURL, err)
			}
			if user != "" || pass != "" {
				t.Errorf("KeychainHelper.Get(%q) = (%q, %q), want empty credentials for non-ACR", tt.serverURL, user, pass)
			}
		})
	}
}
