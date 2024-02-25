//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginClusterIssuer) DeepCopyInto(out *OriginClusterIssuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginClusterIssuer.
func (in *OriginClusterIssuer) DeepCopy() *OriginClusterIssuer {
	if in == nil {
		return nil
	}
	out := new(OriginClusterIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OriginClusterIssuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginClusterIssuerAuthentication) DeepCopyInto(out *OriginClusterIssuerAuthentication) {
	*out = *in
	out.ServiceKeyRef = in.ServiceKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginClusterIssuerAuthentication.
func (in *OriginClusterIssuerAuthentication) DeepCopy() *OriginClusterIssuerAuthentication {
	if in == nil {
		return nil
	}
	out := new(OriginClusterIssuerAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginClusterIssuerCondition) DeepCopyInto(out *OriginClusterIssuerCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginClusterIssuerCondition.
func (in *OriginClusterIssuerCondition) DeepCopy() *OriginClusterIssuerCondition {
	if in == nil {
		return nil
	}
	out := new(OriginClusterIssuerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginClusterIssuerList) DeepCopyInto(out *OriginClusterIssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OriginClusterIssuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginClusterIssuerList.
func (in *OriginClusterIssuerList) DeepCopy() *OriginClusterIssuerList {
	if in == nil {
		return nil
	}
	out := new(OriginClusterIssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OriginClusterIssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginClusterIssuerSpec) DeepCopyInto(out *OriginClusterIssuerSpec) {
	*out = *in
	out.Auth = in.Auth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginClusterIssuerSpec.
func (in *OriginClusterIssuerSpec) DeepCopy() *OriginClusterIssuerSpec {
	if in == nil {
		return nil
	}
	out := new(OriginClusterIssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginClusterIssuerStatus) DeepCopyInto(out *OriginClusterIssuerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OriginClusterIssuerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginClusterIssuerStatus.
func (in *OriginClusterIssuerStatus) DeepCopy() *OriginClusterIssuerStatus {
	if in == nil {
		return nil
	}
	out := new(OriginClusterIssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}
