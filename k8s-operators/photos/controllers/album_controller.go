/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	photosv1alpha1 "k8s-operator/api/v1alpha1"
)

// AlbumReconciler reconciles a Album object
type AlbumReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=photos.lynq.peddy.ai,resources=albums,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=photos.lynq.peddy.ai,resources=albums/status,verbs=get;update;patch

func (r *AlbumReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("album", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *AlbumReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&photosv1alpha1.Album{}).
		Complete(r)
}
