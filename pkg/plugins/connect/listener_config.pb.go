// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: listener_config.proto

/*
Package connect is a generated protocol buffer package.

It is generated from these files:
	listener_config.proto

It has these top-level messages:
	ListenerConfig
	InboundListenerConfig
	AuthConfig
	OutboundListenerConfig
*/
package connect

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// the listenerConfig must be either an InboundListener or an OutboundListener
type ListenerConfig struct {
	// Types that are valid to be assigned to Config:
	//	*ListenerConfig_Inbound
	//	*ListenerConfig_Outbound
	Config isListenerConfig_Config `protobuf_oneof:"config"`
}

func (m *ListenerConfig) Reset()                    { *m = ListenerConfig{} }
func (m *ListenerConfig) String() string            { return proto.CompactTextString(m) }
func (*ListenerConfig) ProtoMessage()               {}
func (*ListenerConfig) Descriptor() ([]byte, []int) { return fileDescriptorListenerConfig, []int{0} }

type isListenerConfig_Config interface {
	isListenerConfig_Config()
	Equal(interface{}) bool
}

type ListenerConfig_Inbound struct {
	Inbound *InboundListenerConfig `protobuf:"bytes,1,opt,name=inbound,oneof"`
}
type ListenerConfig_Outbound struct {
	Outbound *OutboundListenerConfig `protobuf:"bytes,2,opt,name=outbound,oneof"`
}

func (*ListenerConfig_Inbound) isListenerConfig_Config()  {}
func (*ListenerConfig_Outbound) isListenerConfig_Config() {}

func (m *ListenerConfig) GetConfig() isListenerConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ListenerConfig) GetInbound() *InboundListenerConfig {
	if x, ok := m.GetConfig().(*ListenerConfig_Inbound); ok {
		return x.Inbound
	}
	return nil
}

func (m *ListenerConfig) GetOutbound() *OutboundListenerConfig {
	if x, ok := m.GetConfig().(*ListenerConfig_Outbound); ok {
		return x.Outbound
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListenerConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListenerConfig_OneofMarshaler, _ListenerConfig_OneofUnmarshaler, _ListenerConfig_OneofSizer, []interface{}{
		(*ListenerConfig_Inbound)(nil),
		(*ListenerConfig_Outbound)(nil),
	}
}

func _ListenerConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListenerConfig)
	// config
	switch x := m.Config.(type) {
	case *ListenerConfig_Inbound:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Inbound); err != nil {
			return err
		}
	case *ListenerConfig_Outbound:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Outbound); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ListenerConfig.Config has unexpected type %T", x)
	}
	return nil
}

func _ListenerConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListenerConfig)
	switch tag {
	case 1: // config.inbound
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InboundListenerConfig)
		err := b.DecodeMessage(msg)
		m.Config = &ListenerConfig_Inbound{msg}
		return true, err
	case 2: // config.outbound
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OutboundListenerConfig)
		err := b.DecodeMessage(msg)
		m.Config = &ListenerConfig_Outbound{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ListenerConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListenerConfig)
	// config
	switch x := m.Config.(type) {
	case *ListenerConfig_Inbound:
		s := proto.Size(x.Inbound)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ListenerConfig_Outbound:
		s := proto.Size(x.Outbound)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// configuration for the inbound listener
// this listener does authentication and connects
// clients to the local service
type InboundListenerConfig struct {
	// configuration for tls-based auth filter
	AuthConfig *AuthConfig `protobuf:"bytes,1,opt,name=auth_config,json=authConfig" json:"auth_config,omitempty"`
	// the address of the local upstream being proxied
	// the service being proxied must be reachable by Envoy
	LocalServiceAddress string `protobuf:"bytes,2,opt,name=local_service_address,json=localServiceAddress,proto3" json:"local_service_address,omitempty"`
	// the name of the local consul service being proxied
	LocalServiceName string `protobuf:"bytes,3,opt,name=local_service_name,json=localServiceName,proto3" json:"local_service_name,omitempty"`
}

func (m *InboundListenerConfig) Reset()         { *m = InboundListenerConfig{} }
func (m *InboundListenerConfig) String() string { return proto.CompactTextString(m) }
func (*InboundListenerConfig) ProtoMessage()    {}
func (*InboundListenerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorListenerConfig, []int{1}
}

func (m *InboundListenerConfig) GetAuthConfig() *AuthConfig {
	if m != nil {
		return m.AuthConfig
	}
	return nil
}

func (m *InboundListenerConfig) GetLocalServiceAddress() string {
	if m != nil {
		return m.LocalServiceAddress
	}
	return ""
}

func (m *InboundListenerConfig) GetLocalServiceName() string {
	if m != nil {
		return m.LocalServiceName
	}
	return ""
}

// AuthConfig contains information necessary to
// communicate with the Authentication Server (Consul Agent)
type AuthConfig struct {
	// The name of the service who owns this proxy
	// Target must be delivered by the filter as part of the authorize request payload
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// the hostname of the authorization REST service
	AuthorizeHostname string `protobuf:"bytes,2,opt,name=authorize_hostname,json=authorizeHostname,proto3" json:"authorize_hostname,omitempty"`
	// the port of the authorization REST service
	AuthorizePort uint32 `protobuf:"varint,3,opt,name=authorize_port,json=authorizePort,proto3" json:"authorize_port,omitempty"`
	// the request path for the authorization REST service
	// NOTE: currently ignored by the plugin and filter
	AuthorizePath string `protobuf:"bytes,4,opt,name=authorize_path,json=authorizePath,proto3" json:"authorize_path,omitempty"`
	// Connection Timeout tells the filter to set a timeout for unresponsive connections created to this upstream.
	// If not provided by the user, it will set to a default value
	RequestTimeout *time.Duration `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,stdduration" json:"request_timeout,omitempty"`
}

func (m *AuthConfig) Reset()                    { *m = AuthConfig{} }
func (m *AuthConfig) String() string            { return proto.CompactTextString(m) }
func (*AuthConfig) ProtoMessage()               {}
func (*AuthConfig) Descriptor() ([]byte, []int) { return fileDescriptorListenerConfig, []int{2} }

func (m *AuthConfig) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *AuthConfig) GetAuthorizeHostname() string {
	if m != nil {
		return m.AuthorizeHostname
	}
	return ""
}

func (m *AuthConfig) GetAuthorizePort() uint32 {
	if m != nil {
		return m.AuthorizePort
	}
	return 0
}

func (m *AuthConfig) GetAuthorizePath() string {
	if m != nil {
		return m.AuthorizePath
	}
	return ""
}

func (m *AuthConfig) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

// The configuration for the outbound listeners which serve as "tcp routes"
type OutboundListenerConfig struct {
	// The name of the consul service which is the destination for the listener
	DestinationConsulService string `protobuf:"bytes,1,opt,name=destination_consul_service,json=destinationConsulService,proto3" json:"destination_consul_service,omitempty"`
	// TODO (ilackarms): support destination type in Consul Connect API
	DestinationConsulType string `protobuf:"bytes,2,opt,name=destination_consul_type,json=destinationConsulType,proto3" json:"destination_consul_type,omitempty"`
}

func (m *OutboundListenerConfig) Reset()         { *m = OutboundListenerConfig{} }
func (m *OutboundListenerConfig) String() string { return proto.CompactTextString(m) }
func (*OutboundListenerConfig) ProtoMessage()    {}
func (*OutboundListenerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorListenerConfig, []int{3}
}

func (m *OutboundListenerConfig) GetDestinationConsulService() string {
	if m != nil {
		return m.DestinationConsulService
	}
	return ""
}

func (m *OutboundListenerConfig) GetDestinationConsulType() string {
	if m != nil {
		return m.DestinationConsulType
	}
	return ""
}

func init() {
	proto.RegisterType((*ListenerConfig)(nil), "gloo.api.connect.v1.ListenerConfig")
	proto.RegisterType((*InboundListenerConfig)(nil), "gloo.api.connect.v1.InboundListenerConfig")
	proto.RegisterType((*AuthConfig)(nil), "gloo.api.connect.v1.AuthConfig")
	proto.RegisterType((*OutboundListenerConfig)(nil), "gloo.api.connect.v1.OutboundListenerConfig")
}
func (this *ListenerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerConfig)
	if !ok {
		that2, ok := that.(ListenerConfig)
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
	if that1.Config == nil {
		if this.Config != nil {
			return false
		}
	} else if this.Config == nil {
		return false
	} else if !this.Config.Equal(that1.Config) {
		return false
	}
	return true
}
func (this *ListenerConfig_Inbound) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerConfig_Inbound)
	if !ok {
		that2, ok := that.(ListenerConfig_Inbound)
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
	if !this.Inbound.Equal(that1.Inbound) {
		return false
	}
	return true
}
func (this *ListenerConfig_Outbound) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerConfig_Outbound)
	if !ok {
		that2, ok := that.(ListenerConfig_Outbound)
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
	if !this.Outbound.Equal(that1.Outbound) {
		return false
	}
	return true
}
func (this *InboundListenerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InboundListenerConfig)
	if !ok {
		that2, ok := that.(InboundListenerConfig)
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
	if !this.AuthConfig.Equal(that1.AuthConfig) {
		return false
	}
	if this.LocalServiceAddress != that1.LocalServiceAddress {
		return false
	}
	if this.LocalServiceName != that1.LocalServiceName {
		return false
	}
	return true
}
func (this *AuthConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthConfig)
	if !ok {
		that2, ok := that.(AuthConfig)
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
	if this.Target != that1.Target {
		return false
	}
	if this.AuthorizeHostname != that1.AuthorizeHostname {
		return false
	}
	if this.AuthorizePort != that1.AuthorizePort {
		return false
	}
	if this.AuthorizePath != that1.AuthorizePath {
		return false
	}
	if this.RequestTimeout != nil && that1.RequestTimeout != nil {
		if *this.RequestTimeout != *that1.RequestTimeout {
			return false
		}
	} else if this.RequestTimeout != nil {
		return false
	} else if that1.RequestTimeout != nil {
		return false
	}
	return true
}
func (this *OutboundListenerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OutboundListenerConfig)
	if !ok {
		that2, ok := that.(OutboundListenerConfig)
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
	if this.DestinationConsulService != that1.DestinationConsulService {
		return false
	}
	if this.DestinationConsulType != that1.DestinationConsulType {
		return false
	}
	return true
}

func init() { proto.RegisterFile("listener_config.proto", fileDescriptorListenerConfig) }

var fileDescriptorListenerConfig = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0x4e, 0xf3, 0x0f, 0xcd, 0x56, 0x0d, 0xb0, 0x25, 0x34, 0xe4, 0xd0, 0x96, 0x9c,
	0x10, 0x05, 0x9b, 0x16, 0x09, 0x71, 0xe0, 0x40, 0x53, 0x84, 0x52, 0x09, 0x01, 0x32, 0x3d, 0x71,
	0x89, 0x36, 0xf6, 0xd6, 0x5e, 0xe1, 0xec, 0x98, 0xdd, 0xd9, 0x48, 0xe5, 0x21, 0xb8, 0xc2, 0x23,
	0x70, 0xe1, 0x05, 0xe0, 0xc2, 0x9b, 0x20, 0x71, 0xe3, 0x2d, 0x50, 0x76, 0x37, 0x6e, 0x9a, 0xfa,
	0x36, 0xf6, 0x7c, 0xbf, 0x99, 0xf1, 0xe7, 0x8f, 0x74, 0x0b, 0xa1, 0x91, 0x4b, 0xae, 0xc6, 0x09,
	0xc8, 0x33, 0x91, 0x85, 0xa5, 0x02, 0x04, 0xba, 0x95, 0x15, 0x00, 0x21, 0x2b, 0x45, 0x98, 0x80,
	0x94, 0x3c, 0xc1, 0x70, 0x76, 0xd0, 0xdf, 0x9e, 0xb1, 0x42, 0xa4, 0x0c, 0x79, 0xb4, 0x28, 0x9c,
	0xba, 0xbf, 0x93, 0x01, 0x64, 0x05, 0x8f, 0xec, 0xd3, 0xc4, 0x9c, 0x45, 0xa9, 0x51, 0x0c, 0x05,
	0x48, 0xdf, 0xbf, 0x95, 0x41, 0x06, 0xb6, 0x8c, 0xe6, 0x95, 0x7b, 0x3b, 0xf8, 0x1e, 0x90, 0xce,
	0x2b, 0xbf, 0xfd, 0xd8, 0x2e, 0xa7, 0x2f, 0xc9, 0x35, 0x21, 0x27, 0x60, 0x64, 0xda, 0x0b, 0xf6,
	0x82, 0x7b, 0x1b, 0x87, 0xf7, 0xc3, 0x9a, 0x43, 0xc2, 0x13, 0xa7, 0xb9, 0x0c, 0x8f, 0xfe, 0x8b,
	0x17, 0x30, 0x3d, 0x21, 0xeb, 0x60, 0xd0, 0x0d, 0x6a, 0xd8, 0x41, 0xfb, 0xb5, 0x83, 0xde, 0x78,
	0xd1, 0x95, 0x49, 0x15, 0x3e, 0x5c, 0x27, 0x2d, 0xe7, 0xcc, 0xe0, 0x67, 0x40, 0xba, 0xb5, 0x9b,
	0xe9, 0x73, 0xb2, 0xc1, 0x0c, 0xe6, 0xde, 0x42, 0x7f, 0xfa, 0x6e, 0xed, 0xc6, 0x23, 0x83, 0xb9,
	0xa3, 0x62, 0xc2, 0xaa, 0x9a, 0x1e, 0x92, 0x6e, 0x01, 0x09, 0x2b, 0xc6, 0x9a, 0xab, 0x99, 0x48,
	0xf8, 0x98, 0xa5, 0xa9, 0xe2, 0x5a, 0xdb, 0xeb, 0xdb, 0xf1, 0x96, 0x6d, 0xbe, 0x73, 0xbd, 0x23,
	0xd7, 0xa2, 0x0f, 0x08, 0xbd, 0xcc, 0x48, 0x36, 0xe5, 0xbd, 0x35, 0x0b, 0xdc, 0x58, 0x06, 0x5e,
	0xb3, 0x29, 0x1f, 0x7c, 0x69, 0x10, 0x72, 0xb1, 0x9c, 0xde, 0x25, 0x2d, 0x64, 0x2a, 0xe3, 0x68,
	0xaf, 0x6d, 0x0f, 0xdb, 0x3f, 0xfe, 0xfe, 0x5a, 0x6b, 0xaa, 0xc6, 0x5e, 0x10, 0xfb, 0x06, 0x7d,
	0x4a, 0xe8, 0xfc, 0x42, 0x50, 0xe2, 0x13, 0x1f, 0xe7, 0xa0, 0xd1, 0xce, 0x6f, 0xac, 0xca, 0x6f,
	0x56, 0xa2, 0x91, 0xd7, 0xd0, 0x47, 0xa4, 0x73, 0x41, 0x96, 0xa0, 0xd0, 0x5e, 0xb5, 0xb9, 0x4c,
	0x6d, 0x56, 0x82, 0xb7, 0xa0, 0x70, 0x85, 0x60, 0x98, 0xf7, 0x9a, 0xab, 0x7b, 0x96, 0x08, 0x86,
	0x39, 0x1d, 0x91, 0xeb, 0x8a, 0x7f, 0x34, 0x5c, 0xe3, 0x18, 0xc5, 0x94, 0x83, 0xc1, 0xde, 0xff,
	0xd6, 0xf7, 0x3b, 0xa1, 0x4b, 0x63, 0xb8, 0x48, 0x63, 0xf8, 0xc2, 0xa7, 0x71, 0xd8, 0xfc, 0xfa,
	0x7b, 0x37, 0x88, 0x3b, 0x9e, 0x3b, 0x75, 0xd8, 0xe0, 0x73, 0x40, 0x6e, 0xd7, 0x07, 0x81, 0x3e,
	0x23, 0xfd, 0x94, 0x6b, 0x14, 0xd2, 0xf2, 0xf3, 0xff, 0xab, 0x4d, 0xe5, 0xb7, 0x73, 0x2e, 0xee,
	0x2d, 0x29, 0x8e, 0xad, 0xc0, 0xdb, 0x4e, 0x9f, 0x90, 0xed, 0x1a, 0x1a, 0xcf, 0x4b, 0xef, 0x62,
	0xdc, 0xbd, 0x82, 0x9e, 0x9e, 0x97, 0x7c, 0x78, 0xf0, 0xed, 0xcf, 0x4e, 0xf0, 0x7e, 0x3f, 0x13,
	0x98, 0x9b, 0x49, 0x98, 0xc0, 0x34, 0xd2, 0x50, 0xc0, 0x43, 0x01, 0xd1, 0x3c, 0x51, 0x51, 0xf9,
	0x21, 0x8b, 0xca, 0xc2, 0x64, 0x42, 0xea, 0xc8, 0x27, 0x6b, 0xd2, 0xb2, 0x1f, 0xfb, 0xf8, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x8e, 0x98, 0x99, 0xcf, 0x03, 0x00, 0x00,
}
