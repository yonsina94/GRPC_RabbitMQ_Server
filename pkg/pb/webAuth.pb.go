// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: pkg/pb/webAuth.proto

package pb

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

type StartRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *StartRegistrationRequest) Reset() {
	*x = StartRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRegistrationRequest) ProtoMessage() {}

func (x *StartRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRegistrationRequest.ProtoReflect.Descriptor instead.
func (*StartRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{0}
}

func (x *StartRegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type StartRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey *PublicKey `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *StartRegistrationResponse) Reset() {
	*x = StartRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRegistrationResponse) ProtoMessage() {}

func (x *StartRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRegistrationResponse.ProtoReflect.Descriptor instead.
func (*StartRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{1}
}

func (x *StartRegistrationResponse) GetPublicKey() *PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type FinishRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string               `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Id         string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	RawId      string               `protobuf:"bytes,3,opt,name=rawId,proto3" json:"rawId,omitempty"`
	Type       string               `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Transports []string             `protobuf:"bytes,5,rep,name=transports,proto3" json:"transports,omitempty"`
	Response   *AttestationResponse `protobuf:"bytes,6,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *FinishRegistrationRequest) Reset() {
	*x = FinishRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishRegistrationRequest) ProtoMessage() {}

func (x *FinishRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishRegistrationRequest.ProtoReflect.Descriptor instead.
func (*FinishRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{2}
}

func (x *FinishRegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FinishRegistrationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FinishRegistrationRequest) GetRawId() string {
	if x != nil {
		return x.RawId
	}
	return ""
}

func (x *FinishRegistrationRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FinishRegistrationRequest) GetTransports() []string {
	if x != nil {
		return x.Transports
	}
	return nil
}

func (x *FinishRegistrationRequest) GetResponse() *AttestationResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type AttestationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttestationObject string `protobuf:"bytes,1,opt,name=attestationObject,proto3" json:"attestationObject,omitempty"`
	ClientDataJSON    string `protobuf:"bytes,2,opt,name=clientDataJSON,proto3" json:"clientDataJSON,omitempty"`
}

func (x *AttestationResponse) Reset() {
	*x = AttestationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationResponse) ProtoMessage() {}

func (x *AttestationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationResponse.ProtoReflect.Descriptor instead.
func (*AttestationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{3}
}

func (x *AttestationResponse) GetAttestationObject() string {
	if x != nil {
		return x.AttestationObject
	}
	return ""
}

func (x *AttestationResponse) GetClientDataJSON() string {
	if x != nil {
		return x.ClientDataJSON
	}
	return ""
}

type FinishRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FinishRegistrationResponse) Reset() {
	*x = FinishRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishRegistrationResponse) ProtoMessage() {}

func (x *FinishRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishRegistrationResponse.ProtoReflect.Descriptor instead.
func (*FinishRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{4}
}

func (x *FinishRegistrationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge              string                  `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Rp                     *RelyingParty           `protobuf:"bytes,2,opt,name=rp,proto3" json:"rp,omitempty"`
	User                   *User                   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PubKeyCredParams       []*PubKeyCredParams     `protobuf:"bytes,4,rep,name=pubKeyCredParams,proto3" json:"pubKeyCredParams,omitempty"`
	AuthenticatorSelection *AuthenticatorSelection `protobuf:"bytes,5,opt,name=authenticatorSelection,proto3" json:"authenticatorSelection,omitempty"`
	Timeout                int32                   `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Attestation            string                  `protobuf:"bytes,7,opt,name=attestation,proto3" json:"attestation,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{5}
}

func (x *PublicKey) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *PublicKey) GetRp() *RelyingParty {
	if x != nil {
		return x.Rp
	}
	return nil
}

func (x *PublicKey) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PublicKey) GetPubKeyCredParams() []*PubKeyCredParams {
	if x != nil {
		return x.PubKeyCredParams
	}
	return nil
}

func (x *PublicKey) GetAuthenticatorSelection() *AuthenticatorSelection {
	if x != nil {
		return x.AuthenticatorSelection
	}
	return nil
}

func (x *PublicKey) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *PublicKey) GetAttestation() string {
	if x != nil {
		return x.Attestation
	}
	return ""
}

type RelyingParty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RelyingParty) Reset() {
	*x = RelyingParty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelyingParty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelyingParty) ProtoMessage() {}

func (x *RelyingParty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelyingParty.ProtoReflect.Descriptor instead.
func (*RelyingParty) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{6}
}

func (x *RelyingParty) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RelyingParty) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Id          string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PubKeyCredParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Alg  int32  `protobuf:"varint,2,opt,name=alg,proto3" json:"alg,omitempty"`
}

func (x *PubKeyCredParams) Reset() {
	*x = PubKeyCredParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeyCredParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeyCredParams) ProtoMessage() {}

func (x *PubKeyCredParams) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeyCredParams.ProtoReflect.Descriptor instead.
func (*PubKeyCredParams) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{8}
}

func (x *PubKeyCredParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PubKeyCredParams) GetAlg() int32 {
	if x != nil {
		return x.Alg
	}
	return 0
}

type AuthenticatorSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserVerification        string `protobuf:"bytes,1,opt,name=userVerification,proto3" json:"userVerification,omitempty"`
	AuthenticatorAttachment string `protobuf:"bytes,2,opt,name=authenticatorAttachment,proto3" json:"authenticatorAttachment,omitempty"`
}

func (x *AuthenticatorSelection) Reset() {
	*x = AuthenticatorSelection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_webAuth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatorSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatorSelection) ProtoMessage() {}

func (x *AuthenticatorSelection) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_webAuth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatorSelection.ProtoReflect.Descriptor instead.
func (*AuthenticatorSelection) Descriptor() ([]byte, []int) {
	return file_pkg_pb_webAuth_proto_rawDescGZIP(), []int{9}
}

func (x *AuthenticatorSelection) GetUserVerification() string {
	if x != nil {
		return x.UserVerification
	}
	return ""
}

func (x *AuthenticatorSelection) GetAuthenticatorAttachment() string {
	if x != nil {
		return x.AuthenticatorAttachment
	}
	return ""
}

var File_pkg_pb_webAuth_proto protoreflect.FileDescriptor

var file_pkg_pb_webAuth_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x65, 0x62, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62,
	0x22, 0x30, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x4e, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x22, 0xc6, 0x01, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x77, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x77, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x13, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x53,
	0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x34, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd3,
	0x02, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x02, 0x72, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75,
	0x62, 0x2e, 0x52, 0x65, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x02,
	0x72, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x43, 0x72, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x43, 0x72, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x10, 0x70, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x58,
	0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x16, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x43, 0x72, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6c, 0x67,
	0x22, 0x7e, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x17, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x32, 0xc9, 0x01, 0x0a, 0x08, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x12, 0x5c, 0x0a,
	0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x12, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x64, 0x48, 0x75,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x6e, 0x73, 0x69,
	0x6e, 0x61, 0x39, 0x34, 0x2f, 0x47, 0x52, 0x50, 0x43, 0x5f, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x4d, 0x51, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_webAuth_proto_rawDescOnce sync.Once
	file_pkg_pb_webAuth_proto_rawDescData = file_pkg_pb_webAuth_proto_rawDesc
)

func file_pkg_pb_webAuth_proto_rawDescGZIP() []byte {
	file_pkg_pb_webAuth_proto_rawDescOnce.Do(func() {
		file_pkg_pb_webAuth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_webAuth_proto_rawDescData)
	})
	return file_pkg_pb_webAuth_proto_rawDescData
}

var file_pkg_pb_webAuth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_pb_webAuth_proto_goTypes = []interface{}{
	(*StartRegistrationRequest)(nil),   // 0: NftIdHub.StartRegistrationRequest
	(*StartRegistrationResponse)(nil),  // 1: NftIdHub.StartRegistrationResponse
	(*FinishRegistrationRequest)(nil),  // 2: NftIdHub.FinishRegistrationRequest
	(*AttestationResponse)(nil),        // 3: NftIdHub.AttestationResponse
	(*FinishRegistrationResponse)(nil), // 4: NftIdHub.FinishRegistrationResponse
	(*PublicKey)(nil),                  // 5: NftIdHub.PublicKey
	(*RelyingParty)(nil),               // 6: NftIdHub.RelyingParty
	(*User)(nil),                       // 7: NftIdHub.User
	(*PubKeyCredParams)(nil),           // 8: NftIdHub.PubKeyCredParams
	(*AuthenticatorSelection)(nil),     // 9: NftIdHub.AuthenticatorSelection
}
var file_pkg_pb_webAuth_proto_depIdxs = []int32{
	5, // 0: NftIdHub.StartRegistrationResponse.publicKey:type_name -> NftIdHub.PublicKey
	3, // 1: NftIdHub.FinishRegistrationRequest.response:type_name -> NftIdHub.AttestationResponse
	6, // 2: NftIdHub.PublicKey.rp:type_name -> NftIdHub.RelyingParty
	7, // 3: NftIdHub.PublicKey.user:type_name -> NftIdHub.User
	8, // 4: NftIdHub.PublicKey.pubKeyCredParams:type_name -> NftIdHub.PubKeyCredParams
	9, // 5: NftIdHub.PublicKey.authenticatorSelection:type_name -> NftIdHub.AuthenticatorSelection
	0, // 6: NftIdHub.NftIdHub.StartRegistration:input_type -> NftIdHub.StartRegistrationRequest
	2, // 7: NftIdHub.NftIdHub.FinishRegistration:input_type -> NftIdHub.FinishRegistrationRequest
	1, // 8: NftIdHub.NftIdHub.StartRegistration:output_type -> NftIdHub.StartRegistrationResponse
	4, // 9: NftIdHub.NftIdHub.FinishRegistration:output_type -> NftIdHub.FinishRegistrationResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_pb_webAuth_proto_init() }
func file_pkg_pb_webAuth_proto_init() {
	if File_pkg_pb_webAuth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_webAuth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRegistrationRequest); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRegistrationResponse); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishRegistrationRequest); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationResponse); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishRegistrationResponse); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelyingParty); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeyCredParams); i {
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
		file_pkg_pb_webAuth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatorSelection); i {
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
			RawDescriptor: file_pkg_pb_webAuth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_webAuth_proto_goTypes,
		DependencyIndexes: file_pkg_pb_webAuth_proto_depIdxs,
		MessageInfos:      file_pkg_pb_webAuth_proto_msgTypes,
	}.Build()
	File_pkg_pb_webAuth_proto = out.File
	file_pkg_pb_webAuth_proto_rawDesc = nil
	file_pkg_pb_webAuth_proto_goTypes = nil
	file_pkg_pb_webAuth_proto_depIdxs = nil
}
