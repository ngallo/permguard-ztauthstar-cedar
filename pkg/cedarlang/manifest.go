// Copyright 2024 Nitro Agility S.r.l.
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
//
// SPDX-License-Identifier: Apache-2.0

package cedarlang

import (
	"errors"

	azztasmanifests "github.com/permguard/permguard-ztauthstar/pkg/ztauthstar/authstarmodels/manifests"
)

// BuildManifest builds the manifest.
func BuildManifest(manifest *azztasmanifests.Manifest, template string) (*azztasmanifests.Manifest, error) {
	if manifest == nil {
		return nil, errors.New("[cedar] manifest is nil")
	}
	return manifest, nil
}

// ValidateManifest validates the manifest.
func ValidateManifest(manifest *azztasmanifests.Manifest) (bool, error) {
	if manifest == nil {
		return false, errors.New("[cedar] manifest is nil")
	}
	return true, nil
}
