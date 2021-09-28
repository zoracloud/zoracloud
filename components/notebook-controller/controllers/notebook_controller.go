/*
Copyright 2021.

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

	// "github.com/zoracloud/components/notebook-controller/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// "sigs.k8s.io/controller-runtime/pkg/log"
	// notebookv1beta1 "github.com/zoracloud/components/notebook-controller/api/v1beta1"
)

// NotebookReconciler reconciles a Notebook object
type NotebookReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=notebook.zoracloud.ai,resources=notebooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=notebook.zoracloud.ai,resources=notebooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=notebook.zoracloud.ai,resources=notebooks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Notebook object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *NotebookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	// log := r.Log.WithValues("notebook", req.NamespacedName)

	// TODO(yanniszark): Can we avoid reconciling Events and Notebook in the same queue?
	event := &corev1.Event{}
	var getEventErr error
	getEventErr = r.Get(ctx, req.NamespacedName, event)

	if getEventErr == nil {
		// involvedNotebook := &notebookv1beta1.Notebook{}
		// nbName, err := n

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NotebookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return nil
}
