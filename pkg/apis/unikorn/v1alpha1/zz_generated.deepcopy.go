//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022-2024 EscherCloud.

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
	net "net"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationNamedReference) DeepCopyInto(out *ApplicationNamedReference) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(ApplicationReference)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationNamedReference.
func (in *ApplicationNamedReference) DeepCopy() *ApplicationNamedReference {
	if in == nil {
		return nil
	}
	out := new(ApplicationNamedReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationReference) DeepCopyInto(out *ApplicationReference) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ApplicationReferenceKind)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationReference.
func (in *ApplicationReference) DeepCopy() *ApplicationReference {
	if in == nil {
		return nil
	}
	out := new(ApplicationReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplication) DeepCopyInto(out *HelmApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplication.
func (in *HelmApplication) DeepCopy() *HelmApplication {
	if in == nil {
		return nil
	}
	out := new(HelmApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplicationDependency) DeepCopyInto(out *HelmApplicationDependency) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplicationDependency.
func (in *HelmApplicationDependency) DeepCopy() *HelmApplicationDependency {
	if in == nil {
		return nil
	}
	out := new(HelmApplicationDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplicationList) DeepCopyInto(out *HelmApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplicationList.
func (in *HelmApplicationList) DeepCopy() *HelmApplicationList {
	if in == nil {
		return nil
	}
	out := new(HelmApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplicationParameter) DeepCopyInto(out *HelmApplicationParameter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplicationParameter.
func (in *HelmApplicationParameter) DeepCopy() *HelmApplicationParameter {
	if in == nil {
		return nil
	}
	out := new(HelmApplicationParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplicationSpec) DeepCopyInto(out *HelmApplicationSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(string)
		**out = **in
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exported != nil {
		in, out := &in.Exported, &out.Exported
		*out = new(bool)
		**out = **in
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]HelmApplicationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplicationSpec.
func (in *HelmApplicationSpec) DeepCopy() *HelmApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(HelmApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplicationStatus) DeepCopyInto(out *HelmApplicationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplicationStatus.
func (in *HelmApplicationStatus) DeepCopy() *HelmApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(HelmApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApplicationVersion) DeepCopyInto(out *HelmApplicationVersion) {
	*out = *in
	if in.Repo != nil {
		in, out := &in.Repo, &out.Repo
		*out = new(string)
		**out = **in
	}
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Release != nil {
		in, out := &in.Release, &out.Release
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]HelmApplicationParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateNamespace != nil {
		in, out := &in.CreateNamespace, &out.CreateNamespace
		*out = new(bool)
		**out = **in
	}
	if in.ServerSideApply != nil {
		in, out := &in.ServerSideApply, &out.ServerSideApply
		*out = new(bool)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]HelmApplicationDependency, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Recommends != nil {
		in, out := &in.Recommends, &out.Recommends
		*out = make([]HelmApplicationDependency, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApplicationVersion.
func (in *HelmApplicationVersion) DeepCopy() *HelmApplicationVersion {
	if in == nil {
		return nil
	}
	out := new(HelmApplicationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv4Address) DeepCopyInto(out *IPv4Address) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv4Address.
func (in *IPv4Address) DeepCopy() *IPv4Address {
	if in == nil {
		return nil
	}
	out := new(IPv4Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv4Prefix.
func (in *IPv4Prefix) DeepCopy() *IPv4Prefix {
	if in == nil {
		return nil
	}
	out := new(IPv4Prefix)
	in.DeepCopyInto(out)
	return out
}
