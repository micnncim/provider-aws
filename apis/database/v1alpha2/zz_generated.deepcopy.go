// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroup) DeepCopyInto(out *DBSubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroup.
func (in *DBSubnetGroup) DeepCopy() *DBSubnetGroup {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBSubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupExternalStatus) DeepCopyInto(out *DBSubnetGroupExternalStatus) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]Subnet, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupExternalStatus.
func (in *DBSubnetGroupExternalStatus) DeepCopy() *DBSubnetGroupExternalStatus {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupList) DeepCopyInto(out *DBSubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBSubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupList.
func (in *DBSubnetGroupList) DeepCopy() *DBSubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBSubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupNameReferencer) DeepCopyInto(out *DBSubnetGroupNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupNameReferencer.
func (in *DBSubnetGroupNameReferencer) DeepCopy() *DBSubnetGroupNameReferencer {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupNameReferencerForRDSInstance) DeepCopyInto(out *DBSubnetGroupNameReferencerForRDSInstance) {
	*out = *in
	out.DBSubnetGroupNameReferencer = in.DBSubnetGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupNameReferencerForRDSInstance.
func (in *DBSubnetGroupNameReferencerForRDSInstance) DeepCopy() *DBSubnetGroupNameReferencerForRDSInstance {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupNameReferencerForRDSInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupParameters) DeepCopyInto(out *DBSubnetGroupParameters) {
	*out = *in
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]*SubnetIDReferencerForDBSubnetGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetIDReferencerForDBSubnetGroup)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupParameters.
func (in *DBSubnetGroupParameters) DeepCopy() *DBSubnetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupSpec) DeepCopyInto(out *DBSubnetGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.DBSubnetGroupParameters.DeepCopyInto(&out.DBSubnetGroupParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupSpec.
func (in *DBSubnetGroupSpec) DeepCopy() *DBSubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupStatus) DeepCopyInto(out *DBSubnetGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.DBSubnetGroupExternalStatus.DeepCopyInto(&out.DBSubnetGroupExternalStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupStatus.
func (in *DBSubnetGroupStatus) DeepCopy() *DBSubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstance) DeepCopyInto(out *RDSInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstance.
func (in *RDSInstance) DeepCopy() *RDSInstance {
	if in == nil {
		return nil
	}
	out := new(RDSInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceClass) DeepCopyInto(out *RDSInstanceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceClass.
func (in *RDSInstanceClass) DeepCopy() *RDSInstanceClass {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceClassList) DeepCopyInto(out *RDSInstanceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDSInstanceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceClassList.
func (in *RDSInstanceClassList) DeepCopy() *RDSInstanceClassList {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceClassSpecTemplate) DeepCopyInto(out *RDSInstanceClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.RDSInstanceParameters.DeepCopyInto(&out.RDSInstanceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceClassSpecTemplate.
func (in *RDSInstanceClassSpecTemplate) DeepCopy() *RDSInstanceClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceList) DeepCopyInto(out *RDSInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDSInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceList.
func (in *RDSInstanceList) DeepCopy() *RDSInstanceList {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceParameters) DeepCopyInto(out *RDSInstanceParameters) {
	*out = *in
	if in.SubnetGroupNameRef != nil {
		in, out := &in.SubnetGroupNameRef, &out.SubnetGroupNameRef
		*out = new(DBSubnetGroupNameReferencerForRDSInstance)
		**out = **in
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]*SecurityGroupIDReferencerForRDSInstance, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityGroupIDReferencerForRDSInstance)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceParameters.
func (in *RDSInstanceParameters) DeepCopy() *RDSInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceSpec) DeepCopyInto(out *RDSInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.RDSInstanceParameters.DeepCopyInto(&out.RDSInstanceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceSpec.
func (in *RDSInstanceSpec) DeepCopy() *RDSInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceStatus) DeepCopyInto(out *RDSInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceStatus.
func (in *RDSInstanceStatus) DeepCopy() *RDSInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupIDReferencerForRDSInstance) DeepCopyInto(out *SecurityGroupIDReferencerForRDSInstance) {
	*out = *in
	out.SecurityGroupIDReferencer = in.SecurityGroupIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupIDReferencerForRDSInstance.
func (in *SecurityGroupIDReferencerForRDSInstance) DeepCopy() *SecurityGroupIDReferencerForRDSInstance {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupIDReferencerForRDSInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetIDReferencerForDBSubnetGroup) DeepCopyInto(out *SubnetIDReferencerForDBSubnetGroup) {
	*out = *in
	out.SubnetIDReferencer = in.SubnetIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetIDReferencerForDBSubnetGroup.
func (in *SubnetIDReferencerForDBSubnetGroup) DeepCopy() *SubnetIDReferencerForDBSubnetGroup {
	if in == nil {
		return nil
	}
	out := new(SubnetIDReferencerForDBSubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}