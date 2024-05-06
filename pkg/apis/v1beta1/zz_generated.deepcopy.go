//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	apisv1beta1 "sigs.k8s.io/karpenter/pkg/apis/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMI) DeepCopyInto(out *AMI) {
	*out = *in
	if in.Requirements != nil {
		in, out := &in.Requirements, &out.Requirements
		*out = make([]apisv1beta1.NodeSelectorRequirementWithMinValues, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMI.
func (in *AMI) DeepCopy() *AMI {
	if in == nil {
		return nil
	}
	out := new(AMI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMISelectorTerm) DeepCopyInto(out *AMISelectorTerm) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMISelectorTerm.
func (in *AMISelectorTerm) DeepCopy() *AMISelectorTerm {
	if in == nil {
		return nil
	}
	out := new(AMISelectorTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDevice) DeepCopyInto(out *BlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.IOPS != nil {
		in, out := &in.IOPS, &out.IOPS
		*out = new(int64)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDevice.
func (in *BlockDevice) DeepCopy() *BlockDevice {
	if in == nil {
		return nil
	}
	out := new(BlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceMapping) DeepCopyInto(out *BlockDeviceMapping) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.EBS != nil {
		in, out := &in.EBS, &out.EBS
		*out = new(BlockDevice)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMapping.
func (in *BlockDeviceMapping) DeepCopy() *BlockDeviceMapping {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservation) DeepCopyInto(out *CapacityReservation) {
	*out = *in
	if in.Requirements != nil {
		in, out := &in.Requirements, &out.Requirements
		*out = make([]apisv1beta1.NodeSelectorRequirementWithMinValues, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservation.
func (in *CapacityReservation) DeepCopy() *CapacityReservation {
	if in == nil {
		return nil
	}
	out := new(CapacityReservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationSelectorTerm) DeepCopyInto(out *CapacityReservationSelectorTerm) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationSelectorTerm.
func (in *CapacityReservationSelectorTerm) DeepCopy() *CapacityReservationSelectorTerm {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationSelectorTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EC2NodeClass) DeepCopyInto(out *EC2NodeClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EC2NodeClass.
func (in *EC2NodeClass) DeepCopy() *EC2NodeClass {
	if in == nil {
		return nil
	}
	out := new(EC2NodeClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EC2NodeClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EC2NodeClassList) DeepCopyInto(out *EC2NodeClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EC2NodeClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EC2NodeClassList.
func (in *EC2NodeClassList) DeepCopy() *EC2NodeClassList {
	if in == nil {
		return nil
	}
	out := new(EC2NodeClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EC2NodeClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EC2NodeClassSpec) DeepCopyInto(out *EC2NodeClassSpec) {
	*out = *in
	if in.SubnetSelectorTerms != nil {
		in, out := &in.SubnetSelectorTerms, &out.SubnetSelectorTerms
		*out = make([]SubnetSelectorTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupSelectorTerms != nil {
		in, out := &in.SecurityGroupSelectorTerms, &out.SecurityGroupSelectorTerms
		*out = make([]SecurityGroupSelectorTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AssociatePublicIPAddress != nil {
		in, out := &in.AssociatePublicIPAddress, &out.AssociatePublicIPAddress
		*out = new(bool)
		**out = **in
	}
	if in.AMISelectorTerms != nil {
		in, out := &in.AMISelectorTerms, &out.AMISelectorTerms
		*out = make([]AMISelectorTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AMIFamily != nil {
		in, out := &in.AMIFamily, &out.AMIFamily
		*out = new(string)
		**out = **in
	}
	if in.CapacityReservationSelectorTerms != nil {
		in, out := &in.CapacityReservationSelectorTerms, &out.CapacityReservationSelectorTerms
		*out = make([]CapacityReservationSelectorTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.InstanceProfile != nil {
		in, out := &in.InstanceProfile, &out.InstanceProfile
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BlockDeviceMappings != nil {
		in, out := &in.BlockDeviceMappings, &out.BlockDeviceMappings
		*out = make([]*BlockDeviceMapping, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BlockDeviceMapping)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.InstanceStorePolicy != nil {
		in, out := &in.InstanceStorePolicy, &out.InstanceStorePolicy
		*out = new(InstanceStorePolicy)
		**out = **in
	}
	if in.DetailedMonitoring != nil {
		in, out := &in.DetailedMonitoring, &out.DetailedMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = new(MetadataOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EC2NodeClassSpec.
func (in *EC2NodeClassSpec) DeepCopy() *EC2NodeClassSpec {
	if in == nil {
		return nil
	}
	out := new(EC2NodeClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EC2NodeClassStatus) DeepCopyInto(out *EC2NodeClassStatus) {
	*out = *in
	if in.CapacityReservations != nil {
		in, out := &in.CapacityReservations, &out.CapacityReservations
		*out = make([]CapacityReservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]Subnet, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]SecurityGroup, len(*in))
		copy(*out, *in)
	}
	if in.AMIs != nil {
		in, out := &in.AMIs, &out.AMIs
		*out = make([]AMI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EC2NodeClassStatus.
func (in *EC2NodeClassStatus) DeepCopy() *EC2NodeClassStatus {
	if in == nil {
		return nil
	}
	out := new(EC2NodeClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptions) DeepCopyInto(out *MetadataOptions) {
	*out = *in
	if in.HTTPEndpoint != nil {
		in, out := &in.HTTPEndpoint, &out.HTTPEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HTTPProtocolIPv6 != nil {
		in, out := &in.HTTPProtocolIPv6, &out.HTTPProtocolIPv6
		*out = new(string)
		**out = **in
	}
	if in.HTTPPutResponseHopLimit != nil {
		in, out := &in.HTTPPutResponseHopLimit, &out.HTTPPutResponseHopLimit
		*out = new(int64)
		**out = **in
	}
	if in.HTTPTokens != nil {
		in, out := &in.HTTPTokens, &out.HTTPTokens
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptions.
func (in *MetadataOptions) DeepCopy() *MetadataOptions {
	if in == nil {
		return nil
	}
	out := new(MetadataOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupSelectorTerm) DeepCopyInto(out *SecurityGroupSelectorTerm) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupSelectorTerm.
func (in *SecurityGroupSelectorTerm) DeepCopy() *SecurityGroupSelectorTerm {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupSelectorTerm)
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
func (in *SubnetSelectorTerm) DeepCopyInto(out *SubnetSelectorTerm) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSelectorTerm.
func (in *SubnetSelectorTerm) DeepCopy() *SubnetSelectorTerm {
	if in == nil {
		return nil
	}
	out := new(SubnetSelectorTerm)
	in.DeepCopyInto(out)
	return out
}
