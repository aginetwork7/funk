// Code generated by protoc-gen-go-sql. DO NOT EDIT.
// versions:
// - protoc-gen-go-enum v0.0.2
// - protoc                 v5.28.2
// source: errorpb/code.proto

package errorpb

import grpc "google.golang.org/grpc"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func CodeValues() []string {
	return []string{"OK", "Canceled", "Unknown", "InvalidArgument", "DeadlineExceeded", "NotFound", "AlreadyExists", "PermissionDenied", "ResourceExhausted", "FailedPrecondition", "Aborted", "OutOfRange", "Unimplemented", "Internal", "Unavailable", "DataLoss", "Unauthenticated", "TooManyRequests"}
}

func CodeDecode(name string) Code {
	return Code(Code_value[name])
}
