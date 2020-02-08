/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

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

package node

import (
	"oneinfra.ereslibre.es/m/internal/pkg/infra"
)

const (
	kubeScheduler = "k8s.gcr.io/kube-scheduler:v1.17.0"
)

type KubeScheduler struct{}

func (kubeScheduler *KubeScheduler) Reconcile(hypervisor *infra.Hypervisor) error {
	return nil
}
