// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/ttrpc/example/example.proto

package example

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type Method1Request struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar                  string   `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method1Request) Reset()      { *m = Method1Request{} }
func (*Method1Request) ProtoMessage() {}
func (*Method1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b23bbed948ff84, []int{0}
}
func (m *Method1Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Method1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Method1Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Method1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method1Request.Merge(m, src)
}
func (m *Method1Request) XXX_Size() int {
	return m.Size()
}
func (m *Method1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Method1Request.DiscardUnknown(m)
}

var xxx_messageInfo_Method1Request proto.InternalMessageInfo

type Method1Response struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar                  string   `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method1Response) Reset()      { *m = Method1Response{} }
func (*Method1Response) ProtoMessage() {}
func (*Method1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b23bbed948ff84, []int{1}
}
func (m *Method1Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Method1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Method1Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Method1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method1Response.Merge(m, src)
}
func (m *Method1Response) XXX_Size() int {
	return m.Size()
}
func (m *Method1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Method1Response.DiscardUnknown(m)
}

var xxx_messageInfo_Method1Response proto.InternalMessageInfo

type Method2Request struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method2Request) Reset()      { *m = Method2Request{} }
func (*Method2Request) ProtoMessage() {}
func (*Method2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b23bbed948ff84, []int{2}
}
func (m *Method2Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Method2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Method2Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Method2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method2Request.Merge(m, src)
}
func (m *Method2Request) XXX_Size() int {
	return m.Size()
}
func (m *Method2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Method2Request.DiscardUnknown(m)
}

var xxx_messageInfo_Method2Request proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Method1Request)(nil), "ttrpc.example.v1.Method1Request")
	proto.RegisterType((*Method1Response)(nil), "ttrpc.example.v1.Method1Response")
	proto.RegisterType((*Method2Request)(nil), "ttrpc.example.v1.Method2Request")
}

func init() {
	proto.RegisterFile("github.com/containerd/ttrpc/example/example.proto", fileDescriptor_d5b23bbed948ff84)
}

var fileDescriptor_d5b23bbed948ff84 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0xd1, 0x2f, 0x29, 0x29, 0x2a, 0x48, 0xd6, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x85,
	0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x02, 0x60, 0x49, 0x3d, 0x98, 0x60, 0x99, 0xa1,
	0x94, 0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x3e, 0xa9, 0x34, 0x4d, 0x3f, 0x35,
	0xb7, 0xa0, 0xa4, 0x12, 0xa2, 0x5c, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1,
	0x20, 0xa2, 0x4a, 0x26, 0x5c, 0x7c, 0xbe, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0x86, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x69, 0xf9, 0xf9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x20, 0x26, 0x48, 0x24, 0x29, 0xb1, 0x48, 0x82, 0x09, 0x22, 0x92, 0x94, 0x58, 0xa4,
	0x64, 0xca, 0xc5, 0x0f, 0xd7, 0x55, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x4a, 0x94, 0x36, 0x0d, 0x98,
	0x65, 0x46, 0x30, 0xcb, 0xc4, 0xb8, 0xd8, 0x12, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0xa0, 0x1a, 0xa1,
	0x3c, 0xa3, 0x79, 0x8c, 0x5c, 0xec, 0xae, 0x10, 0x8f, 0x09, 0xf9, 0x71, 0xb1, 0x43, 0x2d, 0x13,
	0x52, 0xd0, 0x43, 0xf7, 0xb3, 0x1e, 0xaa, 0xeb, 0xa5, 0x14, 0xf1, 0xa8, 0x80, 0xba, 0xd4, 0x19,
	0x66, 0x9e, 0x11, 0x11, 0xe6, 0x89, 0xe9, 0x41, 0xc2, 0x54, 0x0f, 0x16, 0xa6, 0x7a, 0xae, 0xa0,
	0x30, 0x75, 0x72, 0x3d, 0xf1, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x86, 0x47, 0x72, 0x8c,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x94, 0x36, 0x11,
	0x31, 0x69, 0x0d, 0xa5, 0x93, 0xd8, 0xc0, 0xc6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4a,
	0xc9, 0x2c, 0xdc, 0xff, 0x01, 0x00, 0x00,
}

func (m *Method1Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Method1Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Method1Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Bar) > 0 {
		i -= len(m.Bar)
		copy(dAtA[i:], m.Bar)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Bar)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Foo) > 0 {
		i -= len(m.Foo)
		copy(dAtA[i:], m.Foo)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Foo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Method1Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Method1Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Method1Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Bar) > 0 {
		i -= len(m.Bar)
		copy(dAtA[i:], m.Bar)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Bar)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Foo) > 0 {
		i -= len(m.Foo)
		copy(dAtA[i:], m.Foo)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Foo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Method2Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Method2Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Method2Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExample(dAtA []byte, offset int, v uint64) int {
	offset -= sovExample(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Method1Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Bar)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Method1Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Bar)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Method2Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExample(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExample(x uint64) (n int) {
	return sovExample(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Method1Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Method1Request{`,
		`Foo:` + fmt.Sprintf("%v", this.Foo) + `,`,
		`Bar:` + fmt.Sprintf("%v", this.Bar) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Method1Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Method1Response{`,
		`Foo:` + fmt.Sprintf("%v", this.Foo) + `,`,
		`Bar:` + fmt.Sprintf("%v", this.Bar) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Method2Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Method2Request{`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringExample(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

type ExampleService interface {
	Method1(ctx context.Context, req *Method1Request) (*Method1Response, error)
	Method2(ctx context.Context, req *Method1Request) (*types.Empty, error)
}

func RegisterExampleService(srv *github_com_containerd_ttrpc.Server, svc ExampleService) {
	srv.Register("ttrpc.example.v1.Example", map[string]github_com_containerd_ttrpc.Method{
		"Method1": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req Method1Request
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Method1(ctx, &req)
		},
		"Method2": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req Method1Request
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Method2(ctx, &req)
		},
	})
}

type exampleClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewExampleClient(client *github_com_containerd_ttrpc.Client) ExampleService {
	return &exampleClient{
		client: client,
	}
}

func (c *exampleClient) Method1(ctx context.Context, req *Method1Request) (*Method1Response, error) {
	var resp Method1Response
	if err := c.client.Call(ctx, "ttrpc.example.v1.Example", "Method1", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *exampleClient) Method2(ctx context.Context, req *Method1Request) (*types.Empty, error) {
	var resp types.Empty
	if err := c.client.Call(ctx, "ttrpc.example.v1.Example", "Method2", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (m *Method1Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
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
			return fmt.Errorf("proto: Method1Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Method1Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Foo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Foo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Method1Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
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
			return fmt.Errorf("proto: Method1Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Method1Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Foo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Foo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Method2Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
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
			return fmt.Errorf("proto: Method2Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Method2Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExample(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
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
				return 0, ErrInvalidLengthExample
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExample
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExample
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExample        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExample          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExample = fmt.Errorf("proto: unexpected end of group")
)
