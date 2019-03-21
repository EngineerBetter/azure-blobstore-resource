// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/data_items.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A representation of an image.
type Image struct {
	// Input only. The data representing the image.
	// For Predict calls [image_bytes][] must be set, as other options are not
	// currently supported by prediction API. You can read the contents of an
	// uploaded image by using the [content_uri][] field.
	//
	// Types that are valid to be assigned to Data:
	//	*Image_ImageBytes
	//	*Image_InputConfig
	Data isImage_Data `protobuf_oneof:"data"`
	// Output only. HTTP URI to the thumbnail image.
	ThumbnailUri         string   `protobuf:"bytes,4,opt,name=thumbnail_uri,json=thumbnailUri,proto3" json:"thumbnail_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_16254e250b81d9d6, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

type isImage_Data interface {
	isImage_Data()
}

type Image_ImageBytes struct {
	ImageBytes []byte `protobuf:"bytes,1,opt,name=image_bytes,json=imageBytes,proto3,oneof"`
}

type Image_InputConfig struct {
	InputConfig *InputConfig `protobuf:"bytes,6,opt,name=input_config,json=inputConfig,proto3,oneof"`
}

func (*Image_ImageBytes) isImage_Data() {}

func (*Image_InputConfig) isImage_Data() {}

func (m *Image) GetData() isImage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Image) GetImageBytes() []byte {
	if x, ok := m.GetData().(*Image_ImageBytes); ok {
		return x.ImageBytes
	}
	return nil
}

func (m *Image) GetInputConfig() *InputConfig {
	if x, ok := m.GetData().(*Image_InputConfig); ok {
		return x.InputConfig
	}
	return nil
}

func (m *Image) GetThumbnailUri() string {
	if m != nil {
		return m.ThumbnailUri
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Image) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Image_OneofMarshaler, _Image_OneofUnmarshaler, _Image_OneofSizer, []interface{}{
		(*Image_ImageBytes)(nil),
		(*Image_InputConfig)(nil),
	}
}

func _Image_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Image)
	// data
	switch x := m.Data.(type) {
	case *Image_ImageBytes:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ImageBytes)
	case *Image_InputConfig:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InputConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Image.Data has unexpected type %T", x)
	}
	return nil
}

func _Image_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Image)
	switch tag {
	case 1: // data.image_bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Data = &Image_ImageBytes{x}
		return true, err
	case 6: // data.input_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InputConfig)
		err := b.DecodeMessage(msg)
		m.Data = &Image_InputConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Image_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Image)
	// data
	switch x := m.Data.(type) {
	case *Image_ImageBytes:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ImageBytes)))
		n += len(x.ImageBytes)
	case *Image_InputConfig:
		s := proto.Size(x.InputConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A representation of a text snippet.
type TextSnippet struct {
	// Required. The content of the text snippet as a string. Up to 250000
	// characters long.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The format of the source text. For example, "text/html" or
	// "text/plain". If left blank the format is automatically determined from
	// the type of the uploaded content. The default is "text/html". Up to 25000
	// characters long.
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Output only. HTTP URI where you can download the content.
	ContentUri           string   `protobuf:"bytes,4,opt,name=content_uri,json=contentUri,proto3" json:"content_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSnippet) Reset()         { *m = TextSnippet{} }
func (m *TextSnippet) String() string { return proto.CompactTextString(m) }
func (*TextSnippet) ProtoMessage()    {}
func (*TextSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_16254e250b81d9d6, []int{1}
}
func (m *TextSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSnippet.Unmarshal(m, b)
}
func (m *TextSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSnippet.Marshal(b, m, deterministic)
}
func (dst *TextSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSnippet.Merge(dst, src)
}
func (m *TextSnippet) XXX_Size() int {
	return xxx_messageInfo_TextSnippet.Size(m)
}
func (m *TextSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_TextSnippet proto.InternalMessageInfo

func (m *TextSnippet) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextSnippet) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *TextSnippet) GetContentUri() string {
	if m != nil {
		return m.ContentUri
	}
	return ""
}

// Example data used for training or prediction.
type ExamplePayload struct {
	// Required. Input only. The example data.
	//
	// Types that are valid to be assigned to Payload:
	//	*ExamplePayload_Image
	//	*ExamplePayload_TextSnippet
	Payload              isExamplePayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExamplePayload) Reset()         { *m = ExamplePayload{} }
func (m *ExamplePayload) String() string { return proto.CompactTextString(m) }
func (*ExamplePayload) ProtoMessage()    {}
func (*ExamplePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_items_16254e250b81d9d6, []int{2}
}
func (m *ExamplePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamplePayload.Unmarshal(m, b)
}
func (m *ExamplePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamplePayload.Marshal(b, m, deterministic)
}
func (dst *ExamplePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamplePayload.Merge(dst, src)
}
func (m *ExamplePayload) XXX_Size() int {
	return xxx_messageInfo_ExamplePayload.Size(m)
}
func (m *ExamplePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamplePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ExamplePayload proto.InternalMessageInfo

type isExamplePayload_Payload interface {
	isExamplePayload_Payload()
}

type ExamplePayload_Image struct {
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3,oneof"`
}

type ExamplePayload_TextSnippet struct {
	TextSnippet *TextSnippet `protobuf:"bytes,2,opt,name=text_snippet,json=textSnippet,proto3,oneof"`
}

func (*ExamplePayload_Image) isExamplePayload_Payload() {}

func (*ExamplePayload_TextSnippet) isExamplePayload_Payload() {}

func (m *ExamplePayload) GetPayload() isExamplePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ExamplePayload) GetImage() *Image {
	if x, ok := m.GetPayload().(*ExamplePayload_Image); ok {
		return x.Image
	}
	return nil
}

func (m *ExamplePayload) GetTextSnippet() *TextSnippet {
	if x, ok := m.GetPayload().(*ExamplePayload_TextSnippet); ok {
		return x.TextSnippet
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExamplePayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExamplePayload_OneofMarshaler, _ExamplePayload_OneofUnmarshaler, _ExamplePayload_OneofSizer, []interface{}{
		(*ExamplePayload_Image)(nil),
		(*ExamplePayload_TextSnippet)(nil),
	}
}

func _ExamplePayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExamplePayload)
	// payload
	switch x := m.Payload.(type) {
	case *ExamplePayload_Image:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Image); err != nil {
			return err
		}
	case *ExamplePayload_TextSnippet:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextSnippet); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ExamplePayload.Payload has unexpected type %T", x)
	}
	return nil
}

func _ExamplePayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExamplePayload)
	switch tag {
	case 1: // payload.image
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Image)
		err := b.DecodeMessage(msg)
		m.Payload = &ExamplePayload_Image{msg}
		return true, err
	case 2: // payload.text_snippet
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextSnippet)
		err := b.DecodeMessage(msg)
		m.Payload = &ExamplePayload_TextSnippet{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ExamplePayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExamplePayload)
	// payload
	switch x := m.Payload.(type) {
	case *ExamplePayload_Image:
		s := proto.Size(x.Image)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExamplePayload_TextSnippet:
		s := proto.Size(x.TextSnippet)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Image)(nil), "google.cloud.automl.v1beta1.Image")
	proto.RegisterType((*TextSnippet)(nil), "google.cloud.automl.v1beta1.TextSnippet")
	proto.RegisterType((*ExamplePayload)(nil), "google.cloud.automl.v1beta1.ExamplePayload")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/data_items.proto", fileDescriptor_data_items_16254e250b81d9d6)
}

var fileDescriptor_data_items_16254e250b81d9d6 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x69, 0xeb, 0xa8, 0x53, 0x38, 0xf8, 0x14, 0xad, 0x48, 0x2b, 0x85, 0x43, 0x0f,
	0x28, 0xd1, 0xc6, 0x0d, 0x4e, 0xcb, 0x84, 0xe8, 0x0e, 0x93, 0xa6, 0xb0, 0x71, 0x40, 0x95, 0x22,
	0x27, 0x35, 0xc6, 0x92, 0xed, 0x67, 0x25, 0x2f, 0xa8, 0xb9, 0xf3, 0x59, 0xf8, 0x30, 0x7c, 0x2a,
	0x64, 0x3b, 0xb0, 0x1e, 0x50, 0x76, 0xeb, 0xfb, 0xbf, 0xdf, 0x7b, 0xef, 0xef, 0x7f, 0x43, 0xde,
	0x0a, 0x00, 0xa1, 0x78, 0x56, 0x2b, 0xe8, 0x76, 0x19, 0xeb, 0x10, 0xb4, 0xca, 0x7e, 0x5c, 0x54,
	0x1c, 0xd9, 0x45, 0xb6, 0x63, 0xc8, 0x4a, 0x89, 0x5c, 0xb7, 0xa9, 0x6d, 0x00, 0x81, 0x2e, 0x02,
	0x9d, 0x7a, 0x3a, 0x0d, 0x74, 0x3a, 0xd0, 0x67, 0x2f, 0x87, 0x55, 0xcc, 0xca, 0x8c, 0x19, 0x03,
	0xc8, 0x50, 0x82, 0x19, 0x46, 0xcf, 0xde, 0x8c, 0x1d, 0x92, 0x10, 0xa8, 0xd5, 0xaf, 0x88, 0x9c,
	0xdc, 0x68, 0x26, 0x38, 0x7d, 0x45, 0x62, 0xe9, 0x7e, 0x94, 0x55, 0x8f, 0xbc, 0x4d, 0xa2, 0x65,
	0xb4, 0x9e, 0x6f, 0x26, 0x05, 0xf1, 0x62, 0xee, 0x34, 0x7a, 0x4b, 0xe6, 0xd2, 0xd8, 0x0e, 0xcb,
	0x1a, 0xcc, 0x37, 0x29, 0x92, 0xe9, 0x32, 0x5a, 0xc7, 0x97, 0xeb, 0x74, 0xc4, 0x64, 0x7a, 0xe3,
	0x06, 0xae, 0x3d, 0xbf, 0x99, 0x14, 0xb1, 0x7c, 0x2c, 0xe9, 0x6b, 0xf2, 0x1c, 0xbf, 0x77, 0xba,
	0x32, 0x4c, 0xaa, 0xb2, 0x6b, 0x64, 0x72, 0xbc, 0x8c, 0xd6, 0xb3, 0x62, 0xfe, 0x4f, 0x7c, 0x68,
	0x64, 0x3e, 0x25, 0xc7, 0x2e, 0x95, 0x15, 0x27, 0xf1, 0x3d, 0xdf, 0xe3, 0x67, 0x23, 0xad, 0xe5,
	0x48, 0x13, 0x72, 0x5a, 0x83, 0x41, 0x6e, 0xd0, 0x3b, 0x9d, 0x15, 0x7f, 0x4b, 0xba, 0x20, 0x33,
	0x2d, 0x35, 0x2f, 0xb1, 0xb7, 0x3c, 0x39, 0xf2, 0xbd, 0x67, 0x4e, 0xb8, 0xef, 0x2d, 0xa7, 0xe7,
	0x24, 0x1e, 0xb8, 0x83, 0x83, 0x64, 0x90, 0x1e, 0x1a, 0xe9, 0xf2, 0x78, 0xf1, 0x71, 0xcf, 0xb4,
	0x55, 0xfc, 0x8e, 0xf5, 0x0a, 0xd8, 0x8e, 0xbe, 0x27, 0x27, 0x3e, 0x03, 0x7f, 0x28, 0xbe, 0x5c,
	0x8d, 0x3f, 0xd7, 0x91, 0x9b, 0x49, 0x11, 0x46, 0x5c, 0x62, 0xc8, 0xf7, 0x58, 0xb6, 0xc1, 0xb6,
	0xf7, 0xf3, 0x54, 0x62, 0x07, 0xcf, 0x74, 0x89, 0xe1, 0x63, 0x99, 0xcf, 0xc8, 0xa9, 0x0d, 0xae,
	0xf2, 0x9f, 0x11, 0x39, 0xaf, 0x41, 0x8f, 0x6d, 0xba, 0x8b, 0xbe, 0x5e, 0x0d, 0x6d, 0x01, 0x8a,
	0x19, 0x91, 0x42, 0x23, 0x32, 0xc1, 0x8d, 0xff, 0xeb, 0xb3, 0xd0, 0x62, 0x56, 0xb6, 0xff, 0xfd,
	0x46, 0x3e, 0x84, 0xf2, 0xf7, 0xd1, 0xe2, 0x93, 0x07, 0xb7, 0xd7, 0x0e, 0xda, 0x5e, 0x75, 0x08,
	0xb7, 0x6a, 0xfb, 0x25, 0x40, 0xd5, 0xd4, 0xef, 0x7a, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0x52, 0x09, 0x2a, 0xd7, 0x02, 0x00, 0x00,
}
