// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package errorpb

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Options within kubernetes types, where deepcopy-gen is used.
func (in *Options) DeepCopyInto(out *Options) {
	p := proto.Clone(in).(*Options)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options. Required by controller-gen.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Options. Required by controller-gen.
func (in *Options) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Fields within kubernetes types, where deepcopy-gen is used.
func (in *Fields) DeepCopyInto(out *Fields) {
	p := proto.Clone(in).(*Fields)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fields. Required by controller-gen.
func (in *Fields) DeepCopy() *Fields {
	if in == nil {
		return nil
	}
	out := new(Fields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Fields. Required by controller-gen.
func (in *Fields) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
