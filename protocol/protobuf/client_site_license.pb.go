// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: steammessages_sitelicenseclient.proto

package protobuf

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

type CMsgClientSiteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteId                 *uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	SiteName               *string `protobuf:"bytes,2,opt,name=site_name,json=siteName" json:"site_name,omitempty"`
	AllowCachedCredentials *bool   `protobuf:"varint,3,opt,name=allow_cached_credentials,json=allowCachedCredentials" json:"allow_cached_credentials,omitempty"`
}

func (x *CMsgClientSiteInfo) Reset() {
	*x = CMsgClientSiteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteInfo) ProtoMessage() {}

func (x *CMsgClientSiteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteInfo.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteInfo) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgClientSiteInfo) GetSiteId() uint64 {
	if x != nil && x.SiteId != nil {
		return *x.SiteId
	}
	return 0
}

func (x *CMsgClientSiteInfo) GetSiteName() string {
	if x != nil && x.SiteName != nil {
		return *x.SiteName
	}
	return ""
}

func (x *CMsgClientSiteInfo) GetAllowCachedCredentials() bool {
	if x != nil && x.AllowCachedCredentials != nil {
		return *x.AllowCachedCredentials
	}
	return false
}

type CMsgClientSiteLicenseCheckout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
}

func (x *CMsgClientSiteLicenseCheckout) Reset() {
	*x = CMsgClientSiteLicenseCheckout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteLicenseCheckout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteLicenseCheckout) ProtoMessage() {}

func (x *CMsgClientSiteLicenseCheckout) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteLicenseCheckout.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteLicenseCheckout) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgClientSiteLicenseCheckout) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type CMsgClientSiteLicenseCheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult *int32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
}

// Default values for CMsgClientSiteLicenseCheckoutResponse fields.
const (
	Default_CMsgClientSiteLicenseCheckoutResponse_Eresult = int32(2)
)

func (x *CMsgClientSiteLicenseCheckoutResponse) Reset() {
	*x = CMsgClientSiteLicenseCheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteLicenseCheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteLicenseCheckoutResponse) ProtoMessage() {}

func (x *CMsgClientSiteLicenseCheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteLicenseCheckoutResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteLicenseCheckoutResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgClientSiteLicenseCheckoutResponse) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgClientSiteLicenseCheckoutResponse_Eresult
}

type CMsgClientSiteLicenseGetAvailableSeats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
}

func (x *CMsgClientSiteLicenseGetAvailableSeats) Reset() {
	*x = CMsgClientSiteLicenseGetAvailableSeats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteLicenseGetAvailableSeats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteLicenseGetAvailableSeats) ProtoMessage() {}

func (x *CMsgClientSiteLicenseGetAvailableSeats) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteLicenseGetAvailableSeats.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteLicenseGetAvailableSeats) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{3}
}

func (x *CMsgClientSiteLicenseGetAvailableSeats) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type CMsgClientSiteLicenseGetAvailableSeatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult *int32  `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	Seats   *uint32 `protobuf:"varint,2,opt,name=seats" json:"seats,omitempty"`
}

// Default values for CMsgClientSiteLicenseGetAvailableSeatsResponse fields.
const (
	Default_CMsgClientSiteLicenseGetAvailableSeatsResponse_Eresult = int32(2)
)

func (x *CMsgClientSiteLicenseGetAvailableSeatsResponse) Reset() {
	*x = CMsgClientSiteLicenseGetAvailableSeatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteLicenseGetAvailableSeatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteLicenseGetAvailableSeatsResponse) ProtoMessage() {}

func (x *CMsgClientSiteLicenseGetAvailableSeatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteLicenseGetAvailableSeatsResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteLicenseGetAvailableSeatsResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{4}
}

func (x *CMsgClientSiteLicenseGetAvailableSeatsResponse) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgClientSiteLicenseGetAvailableSeatsResponse_Eresult
}

func (x *CMsgClientSiteLicenseGetAvailableSeatsResponse) GetSeats() uint32 {
	if x != nil && x.Seats != nil {
		return *x.Seats
	}
	return 0
}

type CMsgClientSiteLicenseGetContentCacheInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CMsgClientSiteLicenseGetContentCacheInfo) Reset() {
	*x = CMsgClientSiteLicenseGetContentCacheInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteLicenseGetContentCacheInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteLicenseGetContentCacheInfo) ProtoMessage() {}

func (x *CMsgClientSiteLicenseGetContentCacheInfo) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteLicenseGetContentCacheInfo.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteLicenseGetContentCacheInfo) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{5}
}

type CMsgClientSiteLicenseGetContentCacheInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseCache    *bool   `protobuf:"varint,1,opt,name=use_cache,json=useCache" json:"use_cache,omitempty"`
	Ipv4Address *uint32 `protobuf:"varint,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	PortNumber  *uint32 `protobuf:"varint,3,opt,name=port_number,json=portNumber" json:"port_number,omitempty"`
	P2PGroup    *uint32 `protobuf:"varint,4,opt,name=p2p_group,json=p2pGroup" json:"p2p_group,omitempty"`
	IpAddress   *string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) Reset() {
	*x = CMsgClientSiteLicenseGetContentCacheInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_sitelicenseclient_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientSiteLicenseGetContentCacheInfoResponse) ProtoMessage() {}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_sitelicenseclient_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientSiteLicenseGetContentCacheInfoResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientSiteLicenseGetContentCacheInfoResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_sitelicenseclient_proto_rawDescGZIP(), []int{6}
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) GetUseCache() bool {
	if x != nil && x.UseCache != nil {
		return *x.UseCache
	}
	return false
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) GetIpv4Address() uint32 {
	if x != nil && x.Ipv4Address != nil {
		return *x.Ipv4Address
	}
	return 0
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) GetPortNumber() uint32 {
	if x != nil && x.PortNumber != nil {
		return *x.PortNumber
	}
	return 0
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) GetP2PGroup() uint32 {
	if x != nil && x.P2PGroup != nil {
		return *x.P2PGroup
	}
	return 0
}

func (x *CMsgClientSiteLicenseGetContentCacheInfoResponse) GetIpAddress() string {
	if x != nil && x.IpAddress != nil {
		return *x.IpAddress
	}
	return ""
}

var File_steammessages_sitelicenseclient_proto protoreflect.FileDescriptor

var file_steammessages_sitelicenseclient_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x73, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x35, 0x0a, 0x1d, 0x43, 0x4d, 0x73, 0x67,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x22,
	0x44, 0x0a, 0x25, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x65, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07, 0x65, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0x0a, 0x26, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x47, 0x65,
	0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x47, 0x65,
	0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07, 0x65, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0x2a, 0x0a, 0x28, 0x43, 0x4d,
	0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x30, 0x43, 0x4d, 0x73, 0x67, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70, 0x76, 0x34,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x69, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x32, 0x70, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x32, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x05, 0x48, 0x01, 0x80, 0x01, 0x00,
}

var (
	file_steammessages_sitelicenseclient_proto_rawDescOnce sync.Once
	file_steammessages_sitelicenseclient_proto_rawDescData = file_steammessages_sitelicenseclient_proto_rawDesc
)

func file_steammessages_sitelicenseclient_proto_rawDescGZIP() []byte {
	file_steammessages_sitelicenseclient_proto_rawDescOnce.Do(func() {
		file_steammessages_sitelicenseclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_sitelicenseclient_proto_rawDescData)
	})
	return file_steammessages_sitelicenseclient_proto_rawDescData
}

var file_steammessages_sitelicenseclient_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_steammessages_sitelicenseclient_proto_goTypes = []any{
	(*CMsgClientSiteInfo)(nil),                               // 0: CMsgClientSiteInfo
	(*CMsgClientSiteLicenseCheckout)(nil),                    // 1: CMsgClientSiteLicenseCheckout
	(*CMsgClientSiteLicenseCheckoutResponse)(nil),            // 2: CMsgClientSiteLicenseCheckoutResponse
	(*CMsgClientSiteLicenseGetAvailableSeats)(nil),           // 3: CMsgClientSiteLicenseGetAvailableSeats
	(*CMsgClientSiteLicenseGetAvailableSeatsResponse)(nil),   // 4: CMsgClientSiteLicenseGetAvailableSeatsResponse
	(*CMsgClientSiteLicenseGetContentCacheInfo)(nil),         // 5: CMsgClientSiteLicenseGetContentCacheInfo
	(*CMsgClientSiteLicenseGetContentCacheInfoResponse)(nil), // 6: CMsgClientSiteLicenseGetContentCacheInfoResponse
}
var file_steammessages_sitelicenseclient_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steammessages_sitelicenseclient_proto_init() }
func file_steammessages_sitelicenseclient_proto_init() {
	if File_steammessages_sitelicenseclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_steammessages_sitelicenseclient_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_sitelicenseclient_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteLicenseCheckout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_sitelicenseclient_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteLicenseCheckoutResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_sitelicenseclient_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteLicenseGetAvailableSeats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_sitelicenseclient_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteLicenseGetAvailableSeatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_sitelicenseclient_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteLicenseGetContentCacheInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_sitelicenseclient_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientSiteLicenseGetContentCacheInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_steammessages_sitelicenseclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steammessages_sitelicenseclient_proto_goTypes,
		DependencyIndexes: file_steammessages_sitelicenseclient_proto_depIdxs,
		MessageInfos:      file_steammessages_sitelicenseclient_proto_msgTypes,
	}.Build()
	File_steammessages_sitelicenseclient_proto = out.File
	file_steammessages_sitelicenseclient_proto_rawDesc = nil
	file_steammessages_sitelicenseclient_proto_goTypes = nil
	file_steammessages_sitelicenseclient_proto_depIdxs = nil
}
