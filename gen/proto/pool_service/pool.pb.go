// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: proto/pool_service/pool.proto

package pool_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{0}
}

type StreamPoolsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Since         int32                  `protobuf:"varint,1,opt,name=since,proto3" json:"since,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamPoolsRequest) Reset() {
	*x = StreamPoolsRequest{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamPoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamPoolsRequest) ProtoMessage() {}

func (x *StreamPoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamPoolsRequest.ProtoReflect.Descriptor instead.
func (*StreamPoolsRequest) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{1}
}

func (x *StreamPoolsRequest) GetSince() int32 {
	if x != nil {
		return x.Since
	}
	return 0
}

type GetPoolsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pools         []*Pool                `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPoolsResponse) Reset() {
	*x = GetPoolsResponse{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoolsResponse) ProtoMessage() {}

func (x *GetPoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoolsResponse.ProtoReflect.Descriptor instead.
func (*GetPoolsResponse) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{2}
}

func (x *GetPoolsResponse) GetPools() []*Pool {
	if x != nil {
		return x.Pools
	}
	return nil
}

type GetAssetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserAssets    []string               `protobuf:"bytes,1,rep,name=user_assets,json=userAssets,proto3" json:"user_assets,omitempty"`
	SearchValue   string                 `protobuf:"bytes,2,opt,name=search_value,json=searchValue,proto3" json:"search_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetsRequest) Reset() {
	*x = GetAssetsRequest{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsRequest) ProtoMessage() {}

func (x *GetAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsRequest.ProtoReflect.Descriptor instead.
func (*GetAssetsRequest) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetsRequest) GetUserAssets() []string {
	if x != nil {
		return x.UserAssets
	}
	return nil
}

func (x *GetAssetsRequest) GetSearchValue() string {
	if x != nil {
		return x.SearchValue
	}
	return ""
}

type GetAssetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Assets        []*Asset               `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetsResponse) Reset() {
	*x = GetAssetsResponse{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsResponse) ProtoMessage() {}

func (x *GetAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsResponse.ProtoReflect.Descriptor instead.
func (*GetAssetsResponse) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssetsResponse) GetAssets() []*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

type Pool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Fee           string                 `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Assets        []*Asset               `protobuf:"bytes,4,rep,name=assets,proto3" json:"assets,omitempty"`
	Reserve0      string                 `protobuf:"bytes,5,opt,name=reserve0,proto3" json:"reserve0,omitempty"`
	Reserve1      string                 `protobuf:"bytes,6,opt,name=reserve1,proto3" json:"reserve1,omitempty"`
	Dex           string                 `protobuf:"bytes,7,opt,name=dex,proto3" json:"dex,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pool) Reset() {
	*x = Pool{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

func (x *Pool) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{5}
}

func (x *Pool) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Pool) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Pool) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *Pool) GetAssets() []*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *Pool) GetReserve0() string {
	if x != nil {
		return x.Reserve0
	}
	return ""
}

func (x *Pool) GetReserve1() string {
	if x != nil {
		return x.Reserve1
	}
	return ""
}

func (x *Pool) GetDex() string {
	if x != nil {
		return x.Dex
	}
	return ""
}

func (x *Pool) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Asset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Symbol        string                 `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Image         string                 `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Decimals      uint32                 `protobuf:"varint,6,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Verification  string                 `protobuf:"bytes,7,opt,name=verification,proto3" json:"verification,omitempty"`
	IsScam        bool                   `protobuf:"varint,8,opt,name=is_scam,json=isScam,proto3" json:"is_scam,omitempty"`
	IsWallet      bool                   `protobuf:"varint,9,opt,name=is_wallet,json=isWallet,proto3" json:"is_wallet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Asset) Reset() {
	*x = Asset{}
	mi := &file_proto_pool_service_pool_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_service_pool_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_proto_pool_service_pool_proto_rawDescGZIP(), []int{6}
}

func (x *Asset) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Asset) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Asset) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Asset) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *Asset) GetVerification() string {
	if x != nil {
		return x.Verification
	}
	return ""
}

func (x *Asset) GetIsScam() bool {
	if x != nil {
		return x.IsScam
	}
	return false
}

func (x *Asset) GetIsWallet() bool {
	if x != nil {
		return x.IsWallet
	}
	return false
}

var File_proto_pool_service_pool_proto protoreflect.FileDescriptor

var file_proto_pool_service_pool_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73,
	0x22, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x04, 0x50,
	0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x66, 0x65, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x30, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x30, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x78, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xed, 0x01, 0x0a, 0x05, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x63, 0x61, 0x6d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x63, 0x61, 0x6d, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x32, 0xf8, 0x01, 0x0a, 0x0b, 0x50, 0x6f,
	0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x20,
	0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x2d, 0x70, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x3b, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pool_service_pool_proto_rawDescOnce sync.Once
	file_proto_pool_service_pool_proto_rawDescData = file_proto_pool_service_pool_proto_rawDesc
)

func file_proto_pool_service_pool_proto_rawDescGZIP() []byte {
	file_proto_pool_service_pool_proto_rawDescOnce.Do(func() {
		file_proto_pool_service_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pool_service_pool_proto_rawDescData)
	})
	return file_proto_pool_service_pool_proto_rawDescData
}

var file_proto_pool_service_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_pool_service_pool_proto_goTypes = []any{
	(*Empty)(nil),              // 0: pool_service.Empty
	(*StreamPoolsRequest)(nil), // 1: pool_service.StreamPoolsRequest
	(*GetPoolsResponse)(nil),   // 2: pool_service.GetPoolsResponse
	(*GetAssetsRequest)(nil),   // 3: pool_service.GetAssetsRequest
	(*GetAssetsResponse)(nil),  // 4: pool_service.GetAssetsResponse
	(*Pool)(nil),               // 5: pool_service.Pool
	(*Asset)(nil),              // 6: pool_service.Asset
}
var file_proto_pool_service_pool_proto_depIdxs = []int32{
	5, // 0: pool_service.GetPoolsResponse.pools:type_name -> pool_service.Pool
	6, // 1: pool_service.GetAssetsResponse.assets:type_name -> pool_service.Asset
	6, // 2: pool_service.Pool.assets:type_name -> pool_service.Asset
	0, // 3: pool_service.PoolService.GetPools:input_type -> pool_service.Empty
	1, // 4: pool_service.PoolService.StreamNewPools:input_type -> pool_service.StreamPoolsRequest
	3, // 5: pool_service.PoolService.GetAssets:input_type -> pool_service.GetAssetsRequest
	2, // 6: pool_service.PoolService.GetPools:output_type -> pool_service.GetPoolsResponse
	2, // 7: pool_service.PoolService.StreamNewPools:output_type -> pool_service.GetPoolsResponse
	4, // 8: pool_service.PoolService.GetAssets:output_type -> pool_service.GetAssetsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_pool_service_pool_proto_init() }
func file_proto_pool_service_pool_proto_init() {
	if File_proto_pool_service_pool_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_pool_service_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pool_service_pool_proto_goTypes,
		DependencyIndexes: file_proto_pool_service_pool_proto_depIdxs,
		MessageInfos:      file_proto_pool_service_pool_proto_msgTypes,
	}.Build()
	File_proto_pool_service_pool_proto = out.File
	file_proto_pool_service_pool_proto_rawDesc = nil
	file_proto_pool_service_pool_proto_goTypes = nil
	file_proto_pool_service_pool_proto_depIdxs = nil
}
