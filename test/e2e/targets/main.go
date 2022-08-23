/*
Copyright 2022 TriggerMesh Inc.

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

package targets

import (
	_ "github.com/triggermesh/triggermesh/test/e2e/targets/awskinesis"
	_ "github.com/triggermesh/triggermesh/test/e2e/targets/awss3"
	_ "github.com/triggermesh/triggermesh/test/e2e/targets/awssqs"
	_ "github.com/triggermesh/triggermesh/test/e2e/targets/azureeventhubs"
	_ "github.com/triggermesh/triggermesh/test/e2e/targets/googlecloudstorage"
)
