package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// An OriginClusterIssuer represents the Cloudflare Origin CA as an external cert-manager issuer.
// It is scoped to a single namespace, so it can be used only by resources in the same
// namespace.
type OriginClusterIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Desired state of the OriginClusterIssuer resource
	Spec OriginClusterIssuerSpec `json:"spec,omitempty"`

	// Status of the OriginClusterIssuer. This is set and managed automatically.
	// +optional
	Status OriginClusterIssuerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginClusterIssuerList is a list of OriginClusterIssuers.
type OriginClusterIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata.omitempty"`

	Items []OriginClusterIssuer `json:"items"`
}

// OriginClusterIssuerSpec is the specification of an OriginClusterIssuer. This includes any
// configuration required for the issuer.
type OriginClusterIssuerSpec struct {
	// RequestType is the signature algorithm Cloudflare should use to sign the certificate.
	RequestType RequestType `json:"requestType"`

	// Auth configures how to authenticate with the Cloudflare API.
	Auth OriginClusterIssuerAuthentication `json:"auth"`
}

// OriginClusterIssuerStatus contains status information about an OriginClusterIssuer
type OriginClusterIssuerStatus struct {
	// List of status conditions to indicate the status of an OriginClusterIssuer
	// Known condition types are `Ready`.
	// +optional
	Conditions []OriginClusterIssuerCondition `json:"conditions,omitempty"`
}

// OriginClusterIssuerAuthentication defines how to authenticate with the Cloudflare API.
// Only one of `serviceKeyRef` may be specified.
type OriginClusterIssuerAuthentication struct {
	// ServiceKeyRef authenticates with an API Service Key.
	// +optional
	ServiceKeyRef SecretKeySelector `json:"serviceKeyRef,omitempty"`
}

// SecretKeySelector contains a reference to a secret.
type SecretKeySelector struct {
	// Name of the secret in the OriginClusterIssuer's namespace to select from.
	Name string `json:"name"`
	// Key of the secret to select from. Must be a valid secret key.
	Key string `json:"key"`
	// Namespace where secret is located.
	Namespace string `json:"namespace"`
}

// OriginClusterIssuerCondition contains condition information for the OriginClusterIssuer.
type OriginClusterIssuerCondition struct {
	// Type of the condition, known values are ('Ready')
	Type ConditionType `json:"type"`

	// Status of the condition, one of ('True', 'False', 'Unknown')
	Status ConditionStatus `json:"status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message is a human readable description of the details of the last
	// transition1, complementing reason.
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:validation:Enum=OriginRSA;OriginECC

// RequestType represents the signature algorithm used to sign certificates.
type RequestType string

const (
	// RequestTypeOriginRSA represents an RSA256 signature.
	RequestTypeOriginRSA RequestType = "OriginRSA"

	// RequestTypeOriginECC represents an ECDSA signature.
	RequestTypeOriginECC RequestType = "OriginECC"
)

// +kubebuilder:validation:Enum=Ready

// ConditionType represents an OriginClusterIssuer condition value.
type ConditionType string

const (
	// ConditionReady represents that an OriginClusterIssuer condition is in
	// a ready state and able to issue certificates.
	// If the `status` of this condition is `False`, CertificateRequest
	// controllers should prevent attempts to sign certificates.
	ConditionReady ConditionType = "Ready"
)

// +kubebuilder:validation:Enum=True;False;Unknown

// ConditionStatus represents a condition's status.
type ConditionStatus string

const (
	// ConditionTrue represents the fact that a given condition is true.
	ConditionTrue ConditionStatus = "True"

	// ConditionFalse represents the fact that a given condition is false.
	ConditionFalse ConditionStatus = "False"

	// ConditionUnknown represents the fact that a given condition is unknown.
	ConditionUnknown ConditionStatus = "Unknown"
)
