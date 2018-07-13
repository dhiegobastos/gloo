// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec.proto

/*
Package aws is a generated protocol buffer package.

It is generated from these files:
	spec.proto

It has these top-level messages:
	UpstreamSpec
	FunctionSpec
*/
package aws

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for AWS Lambda Upstreams
type UpstreamSpec struct {
	// The AWS Region in which to run Lambda Functions
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	SecretRef string `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
}

func (m *UpstreamSpec) Reset()                    { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string            { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()               {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{0} }

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() string {
	if m != nil {
		return m.SecretRef
	}
	return ""
}

// Function Spec for Functions on AWS Lambda Upstreams
type FunctionSpec struct {
	// The Name of the Lambda Function as it appears in the AWS Lambda Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// The Qualifier for the Lambda Function. Qualifiers act as a kind of version
	// for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info.
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
}

func (m *FunctionSpec) Reset()                    { *m = FunctionSpec{} }
func (m *FunctionSpec) String() string            { return proto.CompactTextString(m) }
func (*FunctionSpec) ProtoMessage()               {}
func (*FunctionSpec) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{1} }

func (m *FunctionSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *FunctionSpec) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.api.aws.v1.UpstreamSpec")
	proto.RegisterType((*FunctionSpec)(nil), "gloo.api.aws.v1.FunctionSpec")
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.SecretRef != that1.SecretRef {
		return false
	}
	return true
}
func (this *FunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionSpec)
	if !ok {
		that2, ok := that.(FunctionSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.Qualifier != that1.Qualifier {
		return false
	}
	return true
}

func init() { proto.RegisterFile("spec.proto", fileDescriptorSpec) }

var fileDescriptorSpec = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x59, 0x0f, 0x85, 0x1d, 0x56, 0x84, 0x45, 0xa4, 0x88, 0x8a, 0xd4, 0x83, 0x5e, 0x9a,
	0x20, 0xbe, 0x81, 0xa0, 0x47, 0xc1, 0x8a, 0x17, 0x2f, 0x25, 0x0d, 0xb3, 0x71, 0x30, 0xc9, 0xc4,
	0xfc, 0xb1, 0xaf, 0xe4, 0x73, 0xf9, 0x24, 0xd2, 0x74, 0xb7, 0xb7, 0xf9, 0x7e, 0xdf, 0xf0, 0x83,
	0x0f, 0x20, 0x05, 0xd4, 0x22, 0x44, 0xce, 0xdc, 0x9f, 0x18, 0xcb, 0x2c, 0x54, 0x20, 0xa1, 0xb6,
	0x49, 0xfc, 0xdc, 0x9f, 0x9f, 0x1a, 0x36, 0x5c, 0x3b, 0xb9, 0xbb, 0xf6, 0x6f, 0x8b, 0x27, 0xe8,
	0xde, 0x43, 0xca, 0x11, 0x95, 0x7b, 0x0b, 0xa8, 0xfb, 0x33, 0x98, 0x45, 0x34, 0xc4, 0x7e, 0xde,
	0x5c, 0x37, 0x77, 0xed, 0x6a, 0x4c, 0xfd, 0x25, 0x40, 0x42, 0x1d, 0x31, 0xaf, 0x23, 0x0e, 0xf3,
	0xa3, 0xda, 0xb5, 0x7b, 0xb2, 0xc2, 0x61, 0xf1, 0x0a, 0xdd, 0x73, 0xf1, 0x3a, 0x13, 0xfb, 0xaa,
	0xb9, 0x81, 0xe3, 0x61, 0xcc, 0x6b, 0xaf, 0x1c, 0x8e, 0xb6, 0x6e, 0x82, 0x2f, 0xca, 0x61, 0x7f,
	0x01, 0xed, 0x77, 0x51, 0x96, 0x06, 0xc2, 0x38, 0x29, 0x0f, 0xe0, 0x71, 0xf9, 0xfb, 0x77, 0xd5,
	0x7c, 0xdc, 0x1a, 0xca, 0x9f, 0x65, 0x23, 0x34, 0x3b, 0x99, 0xd8, 0xf2, 0x92, 0x58, 0xee, 0x96,
	0xc9, 0xf0, 0x65, 0x64, 0xb0, 0xc5, 0x90, 0x4f, 0x52, 0x6d, 0xd3, 0x66, 0x56, 0xf7, 0x3c, 0xfc,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x56, 0xfc, 0xb5, 0x04, 0x01, 0x00, 0x00,
}
