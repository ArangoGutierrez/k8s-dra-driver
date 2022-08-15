//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableDevice) DeepCopyInto(out *AllocatableDevice) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(AllocatableGpu)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(AllocatableMigDevice)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableDevice.
func (in *AllocatableDevice) DeepCopy() *AllocatableDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatableDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableGpu) DeepCopyInto(out *AllocatableGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableGpu.
func (in *AllocatableGpu) DeepCopy() *AllocatableGpu {
	if in == nil {
		return nil
	}
	out := new(AllocatableGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableMigDevice) DeepCopyInto(out *AllocatableMigDevice) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableMigDevice.
func (in *AllocatableMigDevice) DeepCopy() *AllocatableMigDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatableMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedDevice) DeepCopyInto(out *AllocatedDevice) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(AllocatedGpu)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(AllocatedMigDevice)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedDevice.
func (in *AllocatedDevice) DeepCopy() *AllocatedDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatedDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedGpu) DeepCopyInto(out *AllocatedGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedGpu.
func (in *AllocatedGpu) DeepCopy() *AllocatedGpu {
	if in == nil {
		return nil
	}
	out := new(AllocatedGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedMigDevice) DeepCopyInto(out *AllocatedMigDevice) {
	*out = *in
	out.Placement = in.Placement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedMigDevice.
func (in *AllocatedMigDevice) DeepCopy() *AllocatedMigDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatedMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClass) DeepCopyInto(out *DeviceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClass.
func (in *DeviceClass) DeepCopy() *DeviceClass {
	if in == nil {
		return nil
	}
	out := new(DeviceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClassList) DeepCopyInto(out *DeviceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClassList.
func (in *DeviceClassList) DeepCopy() *DeviceClassList {
	if in == nil {
		return nil
	}
	out := new(DeviceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClassSpec) DeepCopyInto(out *DeviceClassSpec) {
	*out = *in
	if in.DeviceSelector != nil {
		in, out := &in.DeviceSelector, &out.DeviceSelector
		*out = make([]DeviceSelector, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClassSpec.
func (in *DeviceClassSpec) DeepCopy() *DeviceClassSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceRequirements) DeepCopyInto(out *DeviceRequirements) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(GpuClaimSpec)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(MigDeviceClaimSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceRequirements.
func (in *DeviceRequirements) DeepCopy() *DeviceRequirements {
	if in == nil {
		return nil
	}
	out := new(DeviceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSelector) DeepCopyInto(out *DeviceSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSelector.
func (in *DeviceSelector) DeepCopy() *DeviceSelector {
	if in == nil {
		return nil
	}
	out := new(DeviceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuClaim) DeepCopyInto(out *GpuClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuClaim.
func (in *GpuClaim) DeepCopy() *GpuClaim {
	if in == nil {
		return nil
	}
	out := new(GpuClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuClaimList) DeepCopyInto(out *GpuClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GpuClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuClaimList.
func (in *GpuClaimList) DeepCopy() *GpuClaimList {
	if in == nil {
		return nil
	}
	out := new(GpuClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuClaimSpec) DeepCopyInto(out *GpuClaimSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuClaimSpec.
func (in *GpuClaimSpec) DeepCopy() *GpuClaimSpec {
	if in == nil {
		return nil
	}
	out := new(GpuClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDeviceClaim) DeepCopyInto(out *MigDeviceClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDeviceClaim.
func (in *MigDeviceClaim) DeepCopy() *MigDeviceClaim {
	if in == nil {
		return nil
	}
	out := new(MigDeviceClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigDeviceClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDeviceClaimList) DeepCopyInto(out *MigDeviceClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigDeviceClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDeviceClaimList.
func (in *MigDeviceClaimList) DeepCopy() *MigDeviceClaimList {
	if in == nil {
		return nil
	}
	out := new(MigDeviceClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigDeviceClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDeviceClaimSpec) DeepCopyInto(out *MigDeviceClaimSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDeviceClaimSpec.
func (in *MigDeviceClaimSpec) DeepCopy() *MigDeviceClaimSpec {
	if in == nil {
		return nil
	}
	out := new(MigDeviceClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDevicePlacement) DeepCopyInto(out *MigDevicePlacement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDevicePlacement.
func (in *MigDevicePlacement) DeepCopy() *MigDevicePlacement {
	if in == nil {
		return nil
	}
	out := new(MigDevicePlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationState) DeepCopyInto(out *NodeAllocationState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationState.
func (in *NodeAllocationState) DeepCopy() *NodeAllocationState {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateList) DeepCopyInto(out *NodeAllocationStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeAllocationState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateList.
func (in *NodeAllocationStateList) DeepCopy() *NodeAllocationStateList {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateSpec) DeepCopyInto(out *NodeAllocationStateSpec) {
	*out = *in
	if in.AllocatableDevices != nil {
		in, out := &in.AllocatableDevices, &out.AllocatableDevices
		*out = make([]AllocatableDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClaimRequirements != nil {
		in, out := &in.ClaimRequirements, &out.ClaimRequirements
		*out = make(map[string]DeviceRequirements, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ClaimAllocations != nil {
		in, out := &in.ClaimAllocations, &out.ClaimAllocations
		*out = make(map[string][]AllocatedDevice, len(*in))
		for key, val := range *in {
			var outVal []AllocatedDevice
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]AllocatedDevice, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateSpec.
func (in *NodeAllocationStateSpec) DeepCopy() *NodeAllocationStateSpec {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateSpec)
	in.DeepCopyInto(out)
	return out
}