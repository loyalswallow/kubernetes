/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	networkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ServiceCIDRApplyConfiguration represents a declarative configuration of the ServiceCIDR type for use
// with apply.
type ServiceCIDRApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ServiceCIDRSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *ServiceCIDRStatusApplyConfiguration `json:"status,omitempty"`
}

// ServiceCIDR constructs a declarative configuration of the ServiceCIDR type for use with
// apply.
func ServiceCIDR(name string) *ServiceCIDRApplyConfiguration {
	b := &ServiceCIDRApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ServiceCIDR")
	b.WithAPIVersion("networking.k8s.io/v1alpha1")
	return b
}

// ExtractServiceCIDR extracts the applied configuration owned by fieldManager from
// serviceCIDR. If no managedFields are found in serviceCIDR for fieldManager, a
// ServiceCIDRApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// serviceCIDR must be a unmodified ServiceCIDR API object that was retrieved from the Kubernetes API.
// ExtractServiceCIDR provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractServiceCIDR(serviceCIDR *networkingv1alpha1.ServiceCIDR, fieldManager string) (*ServiceCIDRApplyConfiguration, error) {
	return extractServiceCIDR(serviceCIDR, fieldManager, "")
}

// ExtractServiceCIDRStatus is the same as ExtractServiceCIDR except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractServiceCIDRStatus(serviceCIDR *networkingv1alpha1.ServiceCIDR, fieldManager string) (*ServiceCIDRApplyConfiguration, error) {
	return extractServiceCIDR(serviceCIDR, fieldManager, "status")
}

func extractServiceCIDR(serviceCIDR *networkingv1alpha1.ServiceCIDR, fieldManager string, subresource string) (*ServiceCIDRApplyConfiguration, error) {
	b := &ServiceCIDRApplyConfiguration{}
	err := managedfields.ExtractInto(serviceCIDR, internal.Parser().Type("io.k8s.api.networking.v1alpha1.ServiceCIDR"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(serviceCIDR.Name)

	b.WithKind("ServiceCIDR")
	b.WithAPIVersion("networking.k8s.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithKind(value string) *ServiceCIDRApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithAPIVersion(value string) *ServiceCIDRApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithName(value string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithGenerateName(value string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithNamespace(value string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithUID(value types.UID) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithResourceVersion(value string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithGeneration(value int64) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ServiceCIDRApplyConfiguration) WithLabels(entries map[string]string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ServiceCIDRApplyConfiguration) WithAnnotations(entries map[string]string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ServiceCIDRApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ServiceCIDRApplyConfiguration) WithFinalizers(values ...string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithTag(value string) *ServiceCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Tag = &value
	return b
}

func (b *ServiceCIDRApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithSpec(value *ServiceCIDRSpecApplyConfiguration) *ServiceCIDRApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ServiceCIDRApplyConfiguration) WithStatus(value *ServiceCIDRStatusApplyConfiguration) *ServiceCIDRApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ServiceCIDRApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
