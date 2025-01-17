/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzzer

import (
	"kubeops.dev/installer/apis/installer/v1alpha1"

	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
)

// Funcs returns the fuzzer functions for this api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(s *v1alpha1.Auditor, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.ClusterConnector, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.ExternalDnsOperator, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.KubeUiServer, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.Panopticon, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.Scanner, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.Supervisor, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.PrepareCluster, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again
		},
	}
}
