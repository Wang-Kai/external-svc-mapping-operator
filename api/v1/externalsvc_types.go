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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalSvcSpec defines the desired state of ExternalSvc
type ExternalSvcSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ExternalSvcIP   string `json:"external_svc_ip"`
	ExternalSvcPort int    `json:"external_svc_port"`
	InternalSvcPort int    `json:"internal_svc_port,omitempty"`
}

// ExternalSvcStatus defines the observed state of ExternalSvc
type ExternalSvcStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ExternalSvc is the Schema for the externalsvcs API
type ExternalSvc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalSvcSpec   `json:"spec,omitempty"`
	Status ExternalSvcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalSvcList contains a list of ExternalSvc
type ExternalSvcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalSvc `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalSvc{}, &ExternalSvcList{})
}
