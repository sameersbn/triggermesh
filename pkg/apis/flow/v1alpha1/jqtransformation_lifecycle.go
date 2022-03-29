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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*JQTransformation) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("JQTransformation")
}

// GetConditionSet implements duckv1.KRShaped.
func (t *JQTransformation) GetConditionSet() apis.ConditionSet {
	if t.Spec.Sink.Ref != nil || t.Spec.Sink.URI != nil {
		return eventSenderConditionSet
	}
	return targetConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *JQTransformation) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *JQTransformation) GetStatusManager() *StatusManager {
	return &StatusManager{
		ConditionSet: t.GetConditionSet(),
		TargetStatus: &t.Status,
	}
}

// GetSink implements EventSender.
func (t *JQTransformation) GetSink() *duckv1.Destination {
	return &t.Spec.Sink
}