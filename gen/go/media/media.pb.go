// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: media/media.proto

package mediav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendMediaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Author        string                 `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Link          string                 `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Content       *MediaFile             `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Poster        *MediaFile             `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
	Duration      float32                `protobuf:"fixed32,6,opt,name=duration,proto3" json:"duration,omitempty"`
	IsActive      bool                   `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMediaRequest) Reset() {
	*x = SendMediaRequest{}
	mi := &file_media_media_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaRequest) ProtoMessage() {}

func (x *SendMediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaRequest.ProtoReflect.Descriptor instead.
func (*SendMediaRequest) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{0}
}

func (x *SendMediaRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *SendMediaRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendMediaRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *SendMediaRequest) GetContent() *MediaFile {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SendMediaRequest) GetPoster() *MediaFile {
	if x != nil {
		return x.Poster
	}
	return nil
}

func (x *SendMediaRequest) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *SendMediaRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type SendMediaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMediaResponse) Reset() {
	*x = SendMediaResponse{}
	mi := &file_media_media_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaResponse) ProtoMessage() {}

func (x *SendMediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaResponse.ProtoReflect.Descriptor instead.
func (*SendMediaResponse) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{1}
}

func (x *SendMediaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MediaFile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Mimetype      string                 `protobuf:"bytes,3,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaFile) Reset() {
	*x = MediaFile{}
	mi := &file_media_media_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaFile) ProtoMessage() {}

func (x *MediaFile) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaFile.ProtoReflect.Descriptor instead.
func (*MediaFile) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{2}
}

func (x *MediaFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *MediaFile) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *MediaFile) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

var File_media_media_proto protoreflect.FileDescriptor

const file_media_media_proto_rawDesc = "" +
	"\n" +
	"\x11media/media.proto\x12\x05media\"\xe1\x01\n" +
	"\x10SendMediaRequest\x12\x16\n" +
	"\x06author\x18\x01 \x01(\tR\x06author\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x12\n" +
	"\x04link\x18\x03 \x01(\tR\x04link\x12*\n" +
	"\acontent\x18\x04 \x01(\v2\x10.media.MediaFileR\acontent\x12(\n" +
	"\x06poster\x18\x05 \x01(\v2\x10.media.MediaFileR\x06poster\x12\x1a\n" +
	"\bduration\x18\x06 \x01(\x02R\bduration\x12\x1b\n" +
	"\tis_active\x18\a \x01(\bR\bisActive\"-\n" +
	"\x11SendMediaResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"]\n" +
	"\tMediaFile\x12\x18\n" +
	"\acontent\x18\x01 \x01(\fR\acontent\x12\x1a\n" +
	"\bfilename\x18\x02 \x01(\tR\bfilename\x12\x1a\n" +
	"\bmimetype\x18\x03 \x01(\tR\bmimetype2G\n" +
	"\x05Media\x12>\n" +
	"\tSendMedia\x12\x17.media.SendMediaRequest\x1a\x18.media.SendMediaResponseBAZ?github.com/co1seam/ember-backend-api-contracts/media/v1;mediav1b\x06proto3"

var (
	file_media_media_proto_rawDescOnce sync.Once
	file_media_media_proto_rawDescData []byte
)

func file_media_media_proto_rawDescGZIP() []byte {
	file_media_media_proto_rawDescOnce.Do(func() {
		file_media_media_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_media_media_proto_rawDesc), len(file_media_media_proto_rawDesc)))
	})
	return file_media_media_proto_rawDescData
}

var file_media_media_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_media_media_proto_goTypes = []any{
	(*SendMediaRequest)(nil),  // 0: media.SendMediaRequest
	(*SendMediaResponse)(nil), // 1: media.SendMediaResponse
	(*MediaFile)(nil),         // 2: media.MediaFile
}
var file_media_media_proto_depIdxs = []int32{
	2, // 0: media.SendMediaRequest.content:type_name -> media.MediaFile
	2, // 1: media.SendMediaRequest.poster:type_name -> media.MediaFile
	0, // 2: media.Media.SendMedia:input_type -> media.SendMediaRequest
	1, // 3: media.Media.SendMedia:output_type -> media.SendMediaResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_media_media_proto_init() }
func file_media_media_proto_init() {
	if File_media_media_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_media_media_proto_rawDesc), len(file_media_media_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_media_media_proto_goTypes,
		DependencyIndexes: file_media_media_proto_depIdxs,
		MessageInfos:      file_media_media_proto_msgTypes,
	}.Build()
	File_media_media_proto = out.File
	file_media_media_proto_goTypes = nil
	file_media_media_proto_depIdxs = nil
}
