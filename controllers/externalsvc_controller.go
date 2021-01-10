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
	"fmt"

	"github.com/go-logr/logr"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mappingv1 "github.com/Wang-Kai/external-svc-mapping-operator/api/v1"
)

// ExternalSvcReconciler reconciles a ExternalSvc object
type ExternalSvcReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=mapping.sf.ucloud.cn,resources=externalsvcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mapping.sf.ucloud.cn,resources=externalsvcs/status,verbs=get;update;patch

func (r *ExternalSvcReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("externalsvc", req.NamespacedName)

	log.Info("fetching ExternalSvc resource")
	//  fetch external-svc resource
	externalsvc := &mappingv1.ExternalSvc{}
	if err := r.Client.Get(ctx, req.NamespacedName, externalsvc); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// check if there is external svc exist or not
	rawSvc := &core.Service{}
	err := r.Client.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: fmt.Sprintf("%s-svc", req.Name)}, rawSvc)
	if apierrors.IsNotFound(err) {
		// create ExternalSvc
		r.CreateExternalSvc(ctx, externalsvc)

	}
	if err != nil {
		// handle error
		return ctrl.Result{}, nil
	}

	// externalsvc has exist, update it
	// TODO

	return ctrl.Result{}, nil
}

func (r *ExternalSvcReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mappingv1.ExternalSvc{}).
		Complete(r)
}

func (r *ExternalSvcReconciler) CreateExternalSvc(ctx context.Context, externalSvc *mappingv1.ExternalSvc) error {

	ownerReference := *metav1.NewControllerRef(externalSvc, mappingv1.GroupVersion.WithKind("ExternalSvc"))

	// create svc
	svc := &core.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:            externalSvc.Name,
			Namespace:       externalSvc.Namespace,
			OwnerReferences: []metav1.OwnerReference{ownerReference},
		},
		Spec: core.ServiceSpec{
			Ports: []core.ServicePort{
				{
					Port: int32(externalSvc.Spec.InternalSvcPort),
				},
			},
		},
	}
	if err := r.Client.Create(ctx, svc); err != nil {
		return err
	}

	// create ep
	ep := &core.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:            externalSvc.Name,
			Namespace:       externalSvc.Namespace,
			OwnerReferences: []metav1.OwnerReference{ownerReference},
		},
		Subsets: []core.EndpointSubset{
			{
				Addresses: []core.EndpointAddress{
					{
						IP: externalSvc.Spec.ExternalSvcIP,
					},
				},
				Ports: []core.EndpointPort{
					{
						Port: int32(externalSvc.Spec.InternalSvcPort),
					},
				},
			},
		},
	}
	if err := r.Client.Create(ctx, ep); err != nil {
		return err
	}

	return nil
}
