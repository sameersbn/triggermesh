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

package httppollersource

import (
	"context"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	rt "knative.dev/pkg/reconciler/testing"

	tmapis "github.com/triggermesh/triggermesh/pkg/apis"
	commonv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/common/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	fakeinjectionclient "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client/fake"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/sources/v1alpha1/httppollersource"
	common "github.com/triggermesh/triggermesh/pkg/reconciler"
	. "github.com/triggermesh/triggermesh/pkg/reconciler/testing"
)

func TestReconcileSource(t *testing.T) {
	adapterCfg := &adapterConfig{
		Image:   "registry/image:tag",
		configs: &source.EmptyVarsGenerator{},
	}

	ctor := reconcilerCtor(adapterCfg)
	src := newEventSource()
	ab := adapterBuilder(adapterCfg)

	TestReconcileAdapter(t, ctor, src, ab)
}

// reconcilerCtor returns a Ctor for a source Reconciler.
func reconcilerCtor(cfg *adapterConfig) Ctor {
	return func(t *testing.T, ctx context.Context, _ *rt.TableRow, ls *Listers) controller.Reconciler {
		r := &Reconciler{
			base:       NewTestDeploymentReconciler(ctx, ls),
			adapterCfg: cfg,
			srcLister:  ls.GetHTTPPollerSourceLister().HTTPPollerSources,
		}

		return reconcilerv1alpha1.NewReconciler(ctx, logging.FromContext(ctx),
			fakeinjectionclient.Get(ctx), ls.GetHTTPPollerSourceLister(),
			controller.GetEventRecorder(ctx), r)
	}
}

// newEventSource returns a test source object with a minimal set of pre-filled attributes.
func newEventSource() *v1alpha1.HTTPPollerSource {
	endpoint, err := apis.ParseURL("https://test")
	if err != nil {
		panic(err)
	}

	skipVerify := false
	username := "username"
	cacert := "cacert-contents"

	src := &v1alpha1.HTTPPollerSource{
		Spec: v1alpha1.HTTPPollerSourceSpec{
			EventType:  "test-type",
			Endpoint:   *endpoint,
			Method:     "GET",
			Interval:   tmapis.Duration(time.Second * 5),
			SkipVerify: &skipVerify,
			Headers: map[string]string{
				"h1": "v1",
			},
			BasicAuthUsername: &username,
			BasicAuthPassword: &commonv1alpha1.ValueFromField{
				ValueFromSecret: &v1.SecretKeySelector{
					Key: "key",
				},
			},
			CACertificate: &cacert,
		},
	}

	Populate(src)

	return src
}

// adapterBuilder returns a slim Reconciler containing only the fields accessed
// by r.BuildAdapter().
func adapterBuilder(cfg *adapterConfig) common.AdapterDeploymentBuilder {
	return &Reconciler{
		adapterCfg: cfg,
	}
}
