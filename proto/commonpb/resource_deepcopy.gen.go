// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package commonpb

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Identifier within kubernetes types, where deepcopy-gen is used.
func (in *Identifier) DeepCopyInto(out *Identifier) {
	p := proto.Clone(in).(*Identifier)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identifier. Required by controller-gen.
func (in *Identifier) DeepCopy() *Identifier {
	if in == nil {
		return nil
	}
	out := new(Identifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Identifier. Required by controller-gen.
func (in *Identifier) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
