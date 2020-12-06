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
	"strconv"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	opdemov1 "github.com/opdemo/op01/api/v1"
)

// OpDemoReconciler reconciles a OpDemo object
type OpDemoReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=opdemo.opdemo.org,resources=opdemoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=opdemo.opdemo.org,resources=opdemoes/status,verbs=get;update;patch

func (r *OpDemoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	//_ = context.Background()
	ctx := context.Background()
	//	_ = r.Log.WithValues("opdemo", req.NamespacedName)
	reconcileLog := r.Log.WithValues("opdemo", req.NamespacedName)

	// your logic here
	reconcileLog.Info("OP: Reconcile: method called")
	reconcileLog.Info("OP: Reconcile: req.Name= " + req.Name)
	reconcileLog.Info("OP: Reconcile: req.Namespace= " + req.Namespace)
	// get instance of the OpDemo
	opdemo := &opdemov1.OpDemo{}
	err := r.Get(ctx, req.NamespacedName, opdemo)
	if err != nil {
		reconcileLog.Info("OP: Reconcile: error getting OpDemo instance")
		return ctrl.Result{}, err
	}
	reconcileLog.Info("OP: Reconcile: spec Foo = " + opdemo.Spec.Foo)

	podList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(opdemo.Namespace),
	}
	if err = r.List(ctx, podList, listOpts...); err != nil {
		reconcileLog.Info("OP: Reconcile: error listing pods")
		return ctrl.Result{}, err
	}
	reconcileLog.Info("OP: Reconcile: Number of pods = " + strconv.Itoa(len(podList.Items)))
	// for _, podInfo := range (*podList).Items {
	//	reconcileLog.Info("OP: Reconcile: podInfo.Name = " + podInfo.Name)
	//}

	// end custom logic

	return ctrl.Result{}, nil
}

func (r *OpDemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&opdemov1.OpDemo{}).
		Complete(r)
}
