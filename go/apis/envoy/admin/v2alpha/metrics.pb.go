// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2alpha/metrics.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SimpleMetric_Type int32

const (
	SimpleMetric_COUNTER SimpleMetric_Type = 0
	SimpleMetric_GAUGE   SimpleMetric_Type = 1
)

var SimpleMetric_Type_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
}

var SimpleMetric_Type_value = map[string]int32{
	"COUNTER": 0,
	"GAUGE":   1,
}

func (x SimpleMetric_Type) String() string {
	return proto.EnumName(SimpleMetric_Type_name, int32(x))
}

func (SimpleMetric_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_680a736ec6584458, []int{0, 0}
}

// Proto representation of an Envoy Counter or Gauge value.
type SimpleMetric struct {
	// Type of the metric represented.
	Type SimpleMetric_Type `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.admin.v2alpha.SimpleMetric_Type" json:"type,omitempty"`
	// Current metric value.
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	// Name of the metric.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMetric) Reset()         { *m = SimpleMetric{} }
func (m *SimpleMetric) String() string { return proto.CompactTextString(m) }
func (*SimpleMetric) ProtoMessage()    {}
func (*SimpleMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_680a736ec6584458, []int{0}
}
func (m *SimpleMetric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleMetric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMetric.Merge(m, src)
}
func (m *SimpleMetric) XXX_Size() int {
	return m.Size()
}
func (m *SimpleMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMetric.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMetric proto.InternalMessageInfo

func (m *SimpleMetric) GetType() SimpleMetric_Type {
	if m != nil {
		return m.Type
	}
	return SimpleMetric_COUNTER
}

func (m *SimpleMetric) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SimpleMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("envoy.admin.v2alpha.SimpleMetric_Type", SimpleMetric_Type_name, SimpleMetric_Type_value)
	proto.RegisterType((*SimpleMetric)(nil), "envoy.admin.v2alpha.SimpleMetric")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/metrics.proto", fileDescriptor_680a736ec6584458) }

var fileDescriptor_680a736ec6584458 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0xd4,
	0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06,
	0x2b, 0xd1, 0x03, 0x2b, 0xd1, 0x83, 0x2a, 0x51, 0x9a, 0xc2, 0xc8, 0xc5, 0x13, 0x9c, 0x99, 0x5b,
	0x90, 0x93, 0xea, 0x0b, 0x56, 0x2c, 0x64, 0xc5, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa6, 0x87, 0x45, 0x93, 0x1e, 0xb2, 0x06, 0xbd, 0x90, 0xca, 0x82,
	0xd4, 0x20, 0xb0, 0x1e, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x92, 0xe3, 0x62, 0x01, 0xe9, 0x13, 0xe2, 0xe6, 0x62,
	0x77, 0xf6, 0x0f, 0xf5, 0x0b, 0x71, 0x0d, 0x12, 0x60, 0x10, 0xe2, 0xe4, 0x62, 0x75, 0x77, 0x0c,
	0x75, 0x77, 0x15, 0x60, 0x74, 0xb2, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0xb9, 0x14, 0x33, 0xf3, 0x21, 0xee, 0x28, 0x28, 0xca, 0xaf, 0xa8, 0xc4, 0xe6,
	0x24, 0x27, 0x1e, 0x88, 0x6b, 0x8a, 0x03, 0x40, 0x5e, 0x0d, 0x60, 0x4c, 0x62, 0x03, 0xfb, 0xd9,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x90, 0x0a, 0x6c, 0x0a, 0x18, 0x01, 0x00, 0x00,
}

func (m *SimpleMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleMetric) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.Type))
	}
	if m.Value != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.Value))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SimpleMetric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMetrics(uint64(m.Type))
	}
	if m.Value != 0 {
		n += 1 + sovMetrics(uint64(m.Value))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetrics(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SimpleMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: SimpleMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SimpleMetric_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetrics
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
func skipMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetrics
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
				return 0, ErrInvalidLengthMetrics
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMetrics
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetrics
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetrics(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMetrics
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetrics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics   = fmt.Errorf("proto: integer overflow")
)
