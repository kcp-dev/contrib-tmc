//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	tenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"
	conditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/apis/conditions/v1alpha1"

	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceToSync) DeepCopyInto(out *ResourceToSync) {
	*out = *in
	out.GroupResource = in.GroupResource
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceToSync.
func (in *ResourceToSync) DeepCopy() *ResourceToSync {
	if in == nil {
		return nil
	}
	out := new(ResourceToSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncTarget) DeepCopyInto(out *SyncTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncTarget.
func (in *SyncTarget) DeepCopy() *SyncTarget {
	if in == nil {
		return nil
	}
	out := new(SyncTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SyncTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncTargetList) DeepCopyInto(out *SyncTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SyncTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncTargetList.
func (in *SyncTargetList) DeepCopy() *SyncTargetList {
	if in == nil {
		return nil
	}
	out := new(SyncTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SyncTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncTargetSpec) DeepCopyInto(out *SyncTargetSpec) {
	*out = *in
	if in.EvictAfter != nil {
		in, out := &in.EvictAfter, &out.EvictAfter
		*out = (*in).DeepCopy()
	}
	if in.SupportedAPIExports != nil {
		in, out := &in.SupportedAPIExports, &out.SupportedAPIExports
		*out = make([]tenancyv1alpha1.APIExportReference, len(*in))
		copy(*out, *in)
	}
	if in.Cells != nil {
		in, out := &in.Cells, &out.Cells
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncTargetSpec.
func (in *SyncTargetSpec) DeepCopy() *SyncTargetSpec {
	if in == nil {
		return nil
	}
	out := new(SyncTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncTargetStatus) DeepCopyInto(out *SyncTargetStatus) {
	*out = *in
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(v1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncedResources != nil {
		in, out := &in.SyncedResources, &out.SyncedResources
		*out = make([]ResourceToSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastSyncerHeartbeatTime != nil {
		in, out := &in.LastSyncerHeartbeatTime, &out.LastSyncerHeartbeatTime
		*out = (*in).DeepCopy()
	}
	if in.VirtualWorkspaces != nil {
		in, out := &in.VirtualWorkspaces, &out.VirtualWorkspaces
		*out = make([]VirtualWorkspace, len(*in))
		copy(*out, *in)
	}
	if in.TunnelWorkspaces != nil {
		in, out := &in.TunnelWorkspaces, &out.TunnelWorkspaces
		*out = make([]TunnelWorkspace, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncTargetStatus.
func (in *SyncTargetStatus) DeepCopy() *SyncTargetStatus {
	if in == nil {
		return nil
	}
	out := new(SyncTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelWorkspace) DeepCopyInto(out *TunnelWorkspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelWorkspace.
func (in *TunnelWorkspace) DeepCopy() *TunnelWorkspace {
	if in == nil {
		return nil
	}
	out := new(TunnelWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualWorkspace) DeepCopyInto(out *VirtualWorkspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualWorkspace.
func (in *VirtualWorkspace) DeepCopy() *VirtualWorkspace {
	if in == nil {
		return nil
	}
	out := new(VirtualWorkspace)
	in.DeepCopyInto(out)
	return out
}
