// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docs

// ReferenceShort contains short help text.
const ReferenceShort = `Overview of mpdev commands`

// ReferenceLong contains expanded help text.
const ReferenceLong = `mpdev contains commands to both configure and construct 
artifacts needed for publishing to the Google Cloud Marketplace.
`

// ReferenceExamples contains examples.
const ReferenceExamples = `
  # get a package
  $ mpdev pkg get https://github.com/GoogleCloudPlatform/marketplace-tools.git/examples/deployment-manager/autogen/singlevm wordpress
  fetching package examples/deployment-manager/autogen/singlevm from \
    https://github.com/GoogleCloudPlatform/marketplace-tools to wordpress

  # set a value in configuration
  $ mpdev cfg set image wordpress-image
	set 1 fields

  # apply mpdev configuration from mpdev.yaml file
  $ mpdev apply -f mpdev.yaml
  ...
  all resources have been created
`

// ApplyShort contains short help text for apply command.
const ApplyShort = `Updates or creates artifacts defined by the configurations in filename or stdin`

// ApplyLong contains expanded help text for apply command
const ApplyLong = `Updates or creates artifacts defined by the configurations in filename or stdin.
See https://github.com/GoogleCloudPlatform/marketplace-tools/blob/master/docs/mpdev-reference.md#mpdev-resources
for list of resources that can be created.
`

// ApplyExamples contains examples for apply command.
const ApplyExamples = `
  # apply the configuration in dm.yaml
  mpdev apply -f dm.yaml

  # apply the yaml configuration passed into stdin
  cat mpdev.yaml | mpdev apply -f

  # apply the configuration in dm.yaml and gce.yaml
  mpdev apply -f dm.yaml,gce.yaml

  # dryrun of configuration in dm.yaml
  mpdev apply -f dm.yaml --dryrun
`

// TestShort contains short help text for test command.
const TestShort = `Tests artifacts defined by the configurations in filename or stdin`

// TestLong contains expanded help text for test command
const TestLong = `Tests artifacts defined by the configurations in filename or stdin.
See https://github.com/GoogleCloudPlatform/marketplace-tools/blob/master/docs/mpdev-reference.md#mpdev-resources
for list of resources that can be tested.
`

// TestExamples contains examples for test command.
// TODO(willtang): add examples to docstring
const TestExamples = ``
