// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/plgd-dev/cloud/resource-aggregate/pb/resources.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_UNKNOWN            Status = 0
	Status_OK                 Status = 1
	Status_BAD_REQUEST        Status = 2
	Status_UNAUTHORIZED       Status = 3
	Status_FORBIDDEN          Status = 4
	Status_NOT_FOUND          Status = 5
	Status_UNAVAILABLE        Status = 6
	Status_NOT_IMPLEMENTED    Status = 7
	Status_ACCEPTED           Status = 8
	Status_ERROR              Status = 9
	Status_METHOD_NOT_ALLOWED Status = 10
)

var Status_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "OK",
	2:  "BAD_REQUEST",
	3:  "UNAUTHORIZED",
	4:  "FORBIDDEN",
	5:  "NOT_FOUND",
	6:  "UNAVAILABLE",
	7:  "NOT_IMPLEMENTED",
	8:  "ACCEPTED",
	9:  "ERROR",
	10: "METHOD_NOT_ALLOWED",
}

var Status_value = map[string]int32{
	"UNKNOWN":            0,
	"OK":                 1,
	"BAD_REQUEST":        2,
	"UNAUTHORIZED":       3,
	"FORBIDDEN":          4,
	"NOT_FOUND":          5,
	"UNAVAILABLE":        6,
	"NOT_IMPLEMENTED":    7,
	"ACCEPTED":           8,
	"ERROR":              9,
	"METHOD_NOT_ALLOWED": 10,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9784b4964f94935b, []int{0}
}

// https://github.com/openconnectivityfoundation/core/blob/master/schemas/oic.oic-link-schema.json
type Resource struct {
	Id                    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Href                  string                 `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
	ResourceTypes         []string               `protobuf:"bytes,3,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	Interfaces            []string               `protobuf:"bytes,4,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	DeviceId              string                 `protobuf:"bytes,5,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	InstanceId            int64                  `protobuf:"varint,6,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Anchor                string                 `protobuf:"bytes,7,opt,name=anchor,proto3" json:"anchor,omitempty"`
	Policies              *Policies              `protobuf:"bytes,8,opt,name=policies,proto3" json:"policies,omitempty"`
	Title                 string                 `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	SupportedContentTypes []string               `protobuf:"bytes,10,rep,name=supported_content_types,json=supportedContentTypes,proto3" json:"supported_content_types,omitempty"`
	EndpointInformations  []*EndpointInformation `protobuf:"bytes,11,rep,name=endpoint_informations,json=endpointInformations,proto3" json:"endpoint_informations,omitempty"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9784b4964f94935b, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *Resource) GetResourceTypes() []string {
	if m != nil {
		return m.ResourceTypes
	}
	return nil
}

func (m *Resource) GetInterfaces() []string {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *Resource) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Resource) GetInstanceId() int64 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *Resource) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

func (m *Resource) GetPolicies() *Policies {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Resource) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Resource) GetSupportedContentTypes() []string {
	if m != nil {
		return m.SupportedContentTypes
	}
	return nil
}

func (m *Resource) GetEndpointInformations() []*EndpointInformation {
	if m != nil {
		return m.EndpointInformations
	}
	return nil
}

type Policies struct {
	BitFlags int32 `protobuf:"varint,1,opt,name=bit_flags,json=bitFlags,proto3" json:"bit_flags,omitempty"`
}

func (m *Policies) Reset()         { *m = Policies{} }
func (m *Policies) String() string { return proto.CompactTextString(m) }
func (*Policies) ProtoMessage()    {}
func (*Policies) Descriptor() ([]byte, []int) {
	return fileDescriptor_9784b4964f94935b, []int{1}
}
func (m *Policies) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policies.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policies.Merge(m, src)
}
func (m *Policies) XXX_Size() int {
	return m.Size()
}
func (m *Policies) XXX_DiscardUnknown() {
	xxx_messageInfo_Policies.DiscardUnknown(m)
}

var xxx_messageInfo_Policies proto.InternalMessageInfo

func (m *Policies) GetBitFlags() int32 {
	if m != nil {
		return m.BitFlags
	}
	return 0
}

type Content struct {
	Data              []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ContentType       string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	CoapContentFormat int32  `protobuf:"varint,3,opt,name=coap_content_format,json=coapContentFormat,proto3" json:"coap_content_format,omitempty"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_9784b4964f94935b, []int{2}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Content.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return m.Size()
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Content) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Content) GetCoapContentFormat() int32 {
	if m != nil {
		return m.CoapContentFormat
	}
	return 0
}

type EndpointInformation struct {
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Priority int64  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (m *EndpointInformation) Reset()         { *m = EndpointInformation{} }
func (m *EndpointInformation) String() string { return proto.CompactTextString(m) }
func (*EndpointInformation) ProtoMessage()    {}
func (*EndpointInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9784b4964f94935b, []int{3}
}
func (m *EndpointInformation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EndpointInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EndpointInformation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndpointInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointInformation.Merge(m, src)
}
func (m *EndpointInformation) XXX_Size() int {
	return m.Size()
}
func (m *EndpointInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointInformation.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointInformation proto.InternalMessageInfo

func (m *EndpointInformation) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *EndpointInformation) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterEnum("ocf.cloud.resourceaggregate.pb.Status", Status_name, Status_value)
	proto.RegisterType((*Resource)(nil), "ocf.cloud.resourceaggregate.pb.Resource")
	proto.RegisterType((*Policies)(nil), "ocf.cloud.resourceaggregate.pb.Policies")
	proto.RegisterType((*Content)(nil), "ocf.cloud.resourceaggregate.pb.Content")
	proto.RegisterType((*EndpointInformation)(nil), "ocf.cloud.resourceaggregate.pb.EndpointInformation")
}

func init() {
	proto.RegisterFile("github.com/plgd-dev/cloud/resource-aggregate/pb/resources.proto", fileDescriptor_9784b4964f94935b)
}

var fileDescriptor_9784b4964f94935b = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xb5, 0xac, 0xd8, 0x96, 0xaf, 0x9d, 0x44, 0xdf, 0xe4, 0xe7, 0x13, 0xdf, 0x07, 0xaa, 0x6b,
	0x28, 0x35, 0x85, 0xc8, 0x90, 0x40, 0x37, 0x5d, 0x14, 0x3b, 0x92, 0x89, 0x89, 0x2d, 0xa5, 0x13,
	0xbb, 0x81, 0x6c, 0x84, 0x2c, 0x8d, 0xed, 0x01, 0x47, 0x23, 0xa4, 0x71, 0x20, 0x6f, 0xd1, 0x97,
	0xe9, 0x3b, 0xb4, 0xbb, 0x2c, 0xbb, 0x2c, 0xc9, 0x8b, 0x14, 0x8d, 0x64, 0x25, 0xd0, 0xd2, 0xd2,
	0x9d, 0xce, 0x39, 0xf7, 0xdc, 0xb9, 0xe7, 0x8e, 0x06, 0xde, 0x2f, 0x28, 0x5f, 0xae, 0x67, 0x86,
	0xcf, 0x6e, 0xba, 0xd1, 0x6a, 0x11, 0x1c, 0x05, 0xe4, 0xb6, 0xeb, 0xaf, 0xd8, 0x3a, 0xe8, 0xc6,
	0x24, 0x61, 0xeb, 0xd8, 0x27, 0x47, 0xde, 0x62, 0x11, 0x93, 0x85, 0xc7, 0x49, 0x37, 0x9a, 0x15,
	0x6c, 0x62, 0x44, 0x31, 0xe3, 0x0c, 0xe9, 0xcc, 0x9f, 0x1b, 0xc2, 0x60, 0x6c, 0xa4, 0xa2, 0xde,
	0x88, 0x66, 0xed, 0xaf, 0x32, 0x28, 0x38, 0x17, 0xd0, 0x0e, 0x94, 0x69, 0xa0, 0x49, 0x2d, 0xa9,
	0x53, 0xc7, 0x65, 0x1a, 0x20, 0x04, 0x5b, 0xcb, 0x98, 0xcc, 0xb5, 0xb2, 0x60, 0xc4, 0x37, 0x7a,
	0x05, 0x3b, 0x9b, 0x46, 0x2e, 0xbf, 0x8b, 0x48, 0xa2, 0xc9, 0x2d, 0xb9, 0x53, 0xc7, 0xdb, 0x1b,
	0x76, 0x92, 0x92, 0x48, 0x07, 0xa0, 0x21, 0x27, 0xf1, 0xdc, 0xf3, 0x49, 0xa2, 0x6d, 0x89, 0x92,
	0x67, 0x0c, 0xfa, 0x1f, 0xea, 0x01, 0xb9, 0xa5, 0x3e, 0x71, 0x69, 0xa0, 0x55, 0x44, 0x7f, 0x25,
	0x23, 0x86, 0x01, 0x7a, 0x01, 0x0d, 0x1a, 0x26, 0xdc, 0x0b, 0x33, 0xb9, 0xda, 0x92, 0x3a, 0x72,
	0xea, 0xce, 0xa8, 0x61, 0x80, 0x0e, 0xa1, 0xea, 0x85, 0xfe, 0x92, 0xc5, 0x5a, 0x4d, 0x58, 0x73,
	0x84, 0x4c, 0x50, 0x22, 0xb6, 0xa2, 0x3e, 0x25, 0x89, 0xa6, 0xb4, 0xa4, 0x4e, 0xe3, 0xb8, 0x63,
	0xfc, 0x7e, 0x01, 0xc6, 0x45, 0x5e, 0x8f, 0x0b, 0x27, 0xda, 0x87, 0x0a, 0xa7, 0x7c, 0x45, 0xb4,
	0xba, 0x68, 0x9e, 0x01, 0xf4, 0x16, 0xfe, 0x4d, 0xd6, 0x51, 0xc4, 0x62, 0x4e, 0x02, 0xd7, 0x67,
	0x21, 0x27, 0x21, 0xcf, 0x37, 0x00, 0x22, 0xde, 0x41, 0x21, 0x9f, 0x66, 0x6a, 0xb6, 0x89, 0x25,
	0x1c, 0x90, 0x30, 0x88, 0x18, 0x0d, 0xb9, 0x4b, 0xc3, 0x39, 0x8b, 0x6f, 0x3c, 0x4e, 0x59, 0x98,
	0x68, 0x8d, 0x96, 0xdc, 0x69, 0x1c, 0x9f, 0xfc, 0x69, 0x40, 0x2b, 0x37, 0x0f, 0x9f, 0xbc, 0x78,
	0x9f, 0xfc, 0x4c, 0x26, 0xed, 0xd7, 0xa0, 0x6c, 0xd2, 0xa4, 0xfb, 0x9d, 0x51, 0xee, 0xce, 0x57,
	0xde, 0x22, 0x11, 0x37, 0x5a, 0xc1, 0xca, 0x8c, 0xf2, 0x41, 0x8a, 0xdb, 0x11, 0xd4, 0xf2, 0x11,
	0xd3, 0x2b, 0x0e, 0x3c, 0xee, 0x89, 0x92, 0x26, 0x16, 0xdf, 0xe8, 0x25, 0x34, 0x9f, 0xe7, 0xcb,
	0xaf, 0xbf, 0xe1, 0x3f, 0xa5, 0x42, 0x06, 0xec, 0xf9, 0xcc, 0x8b, 0x8a, 0x3d, 0x64, 0x53, 0x68,
	0xb2, 0x38, 0xe8, 0x9f, 0x54, 0xca, 0x0f, 0x18, 0x08, 0xa1, 0x3d, 0x86, 0xbd, 0x5f, 0xe4, 0x40,
	0xff, 0x81, 0xb2, 0x49, 0x92, 0xff, 0x76, 0x05, 0x4e, 0xb5, 0x28, 0xa6, 0x2c, 0xa6, 0xfc, 0x4e,
	0x4c, 0x20, 0xe3, 0x02, 0xbf, 0xf9, 0x2c, 0x41, 0xf5, 0x92, 0x7b, 0x7c, 0x9d, 0xa0, 0x06, 0xd4,
	0xa6, 0xf6, 0xb9, 0xed, 0x5c, 0xd9, 0x6a, 0x09, 0x55, 0xa1, 0xec, 0x9c, 0xab, 0x12, 0xda, 0x85,
	0x46, 0xbf, 0x67, 0xba, 0xd8, 0xfa, 0x30, 0xb5, 0x2e, 0x27, 0x6a, 0x19, 0xa9, 0xd0, 0x9c, 0xda,
	0xbd, 0xe9, 0xe4, 0xcc, 0xc1, 0xc3, 0x6b, 0xcb, 0x54, 0x65, 0xb4, 0x0d, 0xf5, 0x81, 0x83, 0xfb,
	0x43, 0xd3, 0xb4, 0x6c, 0x75, 0x2b, 0x85, 0xb6, 0x33, 0x71, 0x07, 0xce, 0xd4, 0x36, 0xd5, 0x4a,
	0xda, 0x60, 0x6a, 0xf7, 0x3e, 0xf6, 0x86, 0xa3, 0x5e, 0x7f, 0x64, 0xa9, 0x55, 0xb4, 0x07, 0xbb,
	0xa9, 0x3e, 0x1c, 0x5f, 0x8c, 0xac, 0xb1, 0x65, 0x4f, 0x2c, 0x53, 0xad, 0xa1, 0x26, 0x28, 0xbd,
	0xd3, 0x53, 0xeb, 0x22, 0x45, 0x0a, 0xaa, 0x43, 0xc5, 0xc2, 0xd8, 0xc1, 0x6a, 0x1d, 0x1d, 0x02,
	0x1a, 0x5b, 0x93, 0x33, 0xc7, 0x74, 0x53, 0x53, 0x6f, 0x34, 0x72, 0xae, 0x2c, 0x53, 0x85, 0xfe,
	0xe8, 0xcb, 0x83, 0x2e, 0xdd, 0x3f, 0xe8, 0xd2, 0xf7, 0x07, 0x5d, 0xfa, 0xf4, 0xa8, 0x97, 0xee,
	0x1f, 0xf5, 0xd2, 0xb7, 0x47, 0xbd, 0x74, 0x7d, 0xfc, 0x97, 0x0f, 0xfd, 0x5d, 0x34, 0x9b, 0x55,
	0xc5, 0x13, 0x3f, 0xf9, 0x11, 0x00, 0x00, 0xff, 0xff, 0x41, 0x11, 0xe4, 0xa8, 0x25, 0x04, 0x00,
	0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EndpointInformations) > 0 {
		for iNdEx := len(m.EndpointInformations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EndpointInformations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResources(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.SupportedContentTypes) > 0 {
		for iNdEx := len(m.SupportedContentTypes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SupportedContentTypes[iNdEx])
			copy(dAtA[i:], m.SupportedContentTypes[iNdEx])
			i = encodeVarintResources(dAtA, i, uint64(len(m.SupportedContentTypes[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintResources(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Policies != nil {
		{
			size, err := m.Policies.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResources(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Anchor) > 0 {
		i -= len(m.Anchor)
		copy(dAtA[i:], m.Anchor)
		i = encodeVarintResources(dAtA, i, uint64(len(m.Anchor)))
		i--
		dAtA[i] = 0x3a
	}
	if m.InstanceId != 0 {
		i = encodeVarintResources(dAtA, i, uint64(m.InstanceId))
		i--
		dAtA[i] = 0x30
	}
	if len(m.DeviceId) > 0 {
		i -= len(m.DeviceId)
		copy(dAtA[i:], m.DeviceId)
		i = encodeVarintResources(dAtA, i, uint64(len(m.DeviceId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Interfaces) > 0 {
		for iNdEx := len(m.Interfaces) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Interfaces[iNdEx])
			copy(dAtA[i:], m.Interfaces[iNdEx])
			i = encodeVarintResources(dAtA, i, uint64(len(m.Interfaces[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ResourceTypes) > 0 {
		for iNdEx := len(m.ResourceTypes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResourceTypes[iNdEx])
			copy(dAtA[i:], m.ResourceTypes[iNdEx])
			i = encodeVarintResources(dAtA, i, uint64(len(m.ResourceTypes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Href) > 0 {
		i -= len(m.Href)
		copy(dAtA[i:], m.Href)
		i = encodeVarintResources(dAtA, i, uint64(len(m.Href)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintResources(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Policies) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policies) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policies) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BitFlags != 0 {
		i = encodeVarintResources(dAtA, i, uint64(m.BitFlags))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Content) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Content) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoapContentFormat != 0 {
		i = encodeVarintResources(dAtA, i, uint64(m.CoapContentFormat))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContentType) > 0 {
		i -= len(m.ContentType)
		copy(dAtA[i:], m.ContentType)
		i = encodeVarintResources(dAtA, i, uint64(len(m.ContentType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintResources(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EndpointInformation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndpointInformation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EndpointInformation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Priority != 0 {
		i = encodeVarintResources(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintResources(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResources(dAtA []byte, offset int, v uint64) int {
	offset -= sovResources(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	l = len(m.Href)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	if len(m.ResourceTypes) > 0 {
		for _, s := range m.ResourceTypes {
			l = len(s)
			n += 1 + l + sovResources(uint64(l))
		}
	}
	if len(m.Interfaces) > 0 {
		for _, s := range m.Interfaces {
			l = len(s)
			n += 1 + l + sovResources(uint64(l))
		}
	}
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	if m.InstanceId != 0 {
		n += 1 + sovResources(uint64(m.InstanceId))
	}
	l = len(m.Anchor)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	if m.Policies != nil {
		l = m.Policies.Size()
		n += 1 + l + sovResources(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	if len(m.SupportedContentTypes) > 0 {
		for _, s := range m.SupportedContentTypes {
			l = len(s)
			n += 1 + l + sovResources(uint64(l))
		}
	}
	if len(m.EndpointInformations) > 0 {
		for _, e := range m.EndpointInformations {
			l = e.Size()
			n += 1 + l + sovResources(uint64(l))
		}
	}
	return n
}

func (m *Policies) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BitFlags != 0 {
		n += 1 + sovResources(uint64(m.BitFlags))
	}
	return n
}

func (m *Content) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	if m.CoapContentFormat != 0 {
		n += 1 + sovResources(uint64(m.CoapContentFormat))
	}
	return n
}

func (m *EndpointInformation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovResources(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovResources(uint64(m.Priority))
	}
	return n
}

func sovResources(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResources(x uint64) (n int) {
	return sovResources(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResources
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Href", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Href = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceTypes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceTypes = append(m.ResourceTypes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interfaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interfaces = append(m.Interfaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceId", wireType)
			}
			m.InstanceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InstanceId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Anchor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Anchor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policies == nil {
				m.Policies = &Policies{}
			}
			if err := m.Policies.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedContentTypes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedContentTypes = append(m.SupportedContentTypes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndpointInformations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndpointInformations = append(m.EndpointInformations, &EndpointInformation{})
			if err := m.EndpointInformations[len(m.EndpointInformations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResources(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Policies) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResources
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Policies: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policies: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitFlags", wireType)
			}
			m.BitFlags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BitFlags |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResources(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Content) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResources
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoapContentFormat", wireType)
			}
			m.CoapContentFormat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoapContentFormat |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResources(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EndpointInformation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResources
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EndpointInformation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndpointInformation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResources
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResources
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResources
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResources(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResources
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResources(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResources
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResources
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResources
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthResources
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResources
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResources
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResources        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResources          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResources = fmt.Errorf("proto: unexpected end of group")
)
