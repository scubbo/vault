// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: sdk/logical/identity.proto

package logical

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

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier for the entity
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is the human-friendly unique identifier for the entity
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Aliases contains thhe alias mappings for the given entity
	Aliases []*Alias `protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// Metadata represents the custom data tied to this entity
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Disabled is true if the entity is disabled.
	Disabled bool `protobuf:"varint,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// NamespaceID is the identifier of the namespace to which this entity
	// belongs to.
	NamespaceID string `protobuf:"bytes,6,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_logical_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_sdk_logical_identity_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetAliases() []*Alias {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Entity) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Entity) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Entity) GetNamespaceID() string {
	if x != nil {
		return x.NamespaceID
	}
	return ""
}

type Alias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MountType is the backend mount's type to which this identity belongs
	MountType string `protobuf:"bytes,1,opt,name=mount_type,json=mountType,proto3" json:"mount_type,omitempty"`
	// MountAccessor is the identifier of the mount entry to which this
	// identity belongs
	MountAccessor string `protobuf:"bytes,2,opt,name=mount_accessor,json=mountAccessor,proto3" json:"mount_accessor,omitempty"`
	// Name is the identifier of this identity in its authentication source
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata represents the custom data tied to this alias. Fields added
	// to it should have a low rate of change (or no change) because each
	// change incurs a storage write, so quickly-changing fields can have
	// a significant performance impact at scale. See the SDK's
	// "aliasmetadata" package for a helper that eases and standardizes
	// using this safely.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID is the unique identifier for the alias
	ID string `protobuf:"bytes,5,opt,name=ID,proto3" json:"ID,omitempty"`
	// NamespaceID is the identifier of the namespace to which this alias
	// belongs.
	NamespaceID string `protobuf:"bytes,6,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
	// Custom Metadata represents the custom data tied to this alias
	CustomMetadata map[string]string `protobuf:"bytes,7,rep,name=custom_metadata,json=customMetadata,proto3" json:"custom_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Local indicates if the alias only belongs to the cluster where it was
	// created. If true, the alias will be stored in a location that are ignored
	// by the performance replication subsystem.
	Local bool `protobuf:"varint,8,opt,name=local,proto3" json:"local,omitempty"`
}

func (x *Alias) Reset() {
	*x = Alias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_logical_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alias) ProtoMessage() {}

func (x *Alias) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alias.ProtoReflect.Descriptor instead.
func (*Alias) Descriptor() ([]byte, []int) {
	return file_sdk_logical_identity_proto_rawDescGZIP(), []int{1}
}

func (x *Alias) GetMountType() string {
	if x != nil {
		return x.MountType
	}
	return ""
}

func (x *Alias) GetMountAccessor() string {
	if x != nil {
		return x.MountAccessor
	}
	return ""
}

func (x *Alias) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Alias) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Alias) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Alias) GetNamespaceID() string {
	if x != nil {
		return x.NamespaceID
	}
	return ""
}

func (x *Alias) GetCustomMetadata() map[string]string {
	if x != nil {
		return x.CustomMetadata
	}
	return nil
}

func (x *Alias) GetLocal() bool {
	if x != nil {
		return x.Local
	}
	return false
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier for the group
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is the human-friendly unique identifier for the group
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata represents the custom data tied to this group
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// NamespaceID is the identifier of the namespace to which this group
	// belongs to.
	NamespaceID string `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_logical_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_sdk_logical_identity_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Group) GetNamespaceID() string {
	if x != nil {
		return x.NamespaceID
	}
	return ""
}

type MFAMethodID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ID           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	UsesPasscode bool   `protobuf:"varint,3,opt,name=uses_passcode,json=usesPasscode,proto3" json:"uses_passcode,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MFAMethodID) Reset() {
	*x = MFAMethodID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_logical_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MFAMethodID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MFAMethodID) ProtoMessage() {}

func (x *MFAMethodID) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MFAMethodID.ProtoReflect.Descriptor instead.
func (*MFAMethodID) Descriptor() ([]byte, []int) {
	return file_sdk_logical_identity_proto_rawDescGZIP(), []int{3}
}

func (x *MFAMethodID) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MFAMethodID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MFAMethodID) GetUsesPasscode() bool {
	if x != nil {
		return x.UsesPasscode
	}
	return false
}

func (x *MFAMethodID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MFAConstraintAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Any []*MFAMethodID `protobuf:"bytes,1,rep,name=any,proto3" json:"any,omitempty"`
}

func (x *MFAConstraintAny) Reset() {
	*x = MFAConstraintAny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_logical_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MFAConstraintAny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MFAConstraintAny) ProtoMessage() {}

func (x *MFAConstraintAny) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MFAConstraintAny.ProtoReflect.Descriptor instead.
func (*MFAConstraintAny) Descriptor() ([]byte, []int) {
	return file_sdk_logical_identity_proto_rawDescGZIP(), []int{4}
}

func (x *MFAConstraintAny) GetAny() []*MFAMethodID {
	if x != nil {
		return x.Any
	}
	return nil
}

type MFARequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MFARequestID   string                       `protobuf:"bytes,1,opt,name=mfa_request_id,json=mfaRequestId,proto3" json:"mfa_request_id,omitempty"`
	MFAConstraints map[string]*MFAConstraintAny `protobuf:"bytes,2,rep,name=mfa_constraints,json=mfaConstraints,proto3" json:"mfa_constraints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MFARequirement) Reset() {
	*x = MFARequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_logical_identity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MFARequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MFARequirement) ProtoMessage() {}

func (x *MFARequirement) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_identity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MFARequirement.ProtoReflect.Descriptor instead.
func (*MFARequirement) Descriptor() ([]byte, []int) {
	return file_sdk_logical_identity_proto_rawDescGZIP(), []int{5}
}

func (x *MFARequirement) GetMFARequestID() string {
	if x != nil {
		return x.MFARequestID
	}
	return ""
}

func (x *MFARequirement) GetMFAConstraints() map[string]*MFAConstraintAny {
	if x != nil {
		return x.MFAConstraints
	}
	return nil
}

var File_sdk_logical_identity_proto protoreflect.FileDescriptor

var file_sdk_logical_identity_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x22, 0x8d, 0x02, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x39,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb1, 0x03, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc5, 0x01, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x61, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x6a, 0x0a, 0x0b, 0x4d, 0x46, 0x41, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x73, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a,
	0x10, 0x4d, 0x46, 0x41, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x41, 0x6e,
	0x79, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x4d, 0x46, 0x41, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x49, 0x44, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x22, 0xea, 0x01, 0x0a, 0x0e, 0x4d, 0x46,
	0x41, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x6d, 0x66, 0x61, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x66, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x54, 0x0a, 0x0f, 0x6d, 0x66, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x4d, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x66, 0x61, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x6d, 0x66, 0x61, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x5c, 0x0a, 0x13, 0x4d, 0x66, 0x61, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x4d, 0x46, 0x41, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_logical_identity_proto_rawDescOnce sync.Once
	file_sdk_logical_identity_proto_rawDescData = file_sdk_logical_identity_proto_rawDesc
)

func file_sdk_logical_identity_proto_rawDescGZIP() []byte {
	file_sdk_logical_identity_proto_rawDescOnce.Do(func() {
		file_sdk_logical_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_logical_identity_proto_rawDescData)
	})
	return file_sdk_logical_identity_proto_rawDescData
}

var file_sdk_logical_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_sdk_logical_identity_proto_goTypes = []interface{}{
	(*Entity)(nil),           // 0: logical.Entity
	(*Alias)(nil),            // 1: logical.Alias
	(*Group)(nil),            // 2: logical.Group
	(*MFAMethodID)(nil),      // 3: logical.MFAMethodID
	(*MFAConstraintAny)(nil), // 4: logical.MFAConstraintAny
	(*MFARequirement)(nil),   // 5: logical.MFARequirement
	nil,                      // 6: logical.Entity.MetadataEntry
	nil,                      // 7: logical.Alias.MetadataEntry
	nil,                      // 8: logical.Alias.CustomMetadataEntry
	nil,                      // 9: logical.Group.MetadataEntry
	nil,                      // 10: logical.MFARequirement.MFAConstraintsEntry
}
var file_sdk_logical_identity_proto_depIDxs = []int32{
	1,  // 0: logical.Entity.aliases:type_name -> logical.Alias
	6,  // 1: logical.Entity.metadata:type_name -> logical.Entity.MetadataEntry
	7,  // 2: logical.Alias.metadata:type_name -> logical.Alias.MetadataEntry
	8,  // 3: logical.Alias.custom_metadata:type_name -> logical.Alias.CustomMetadataEntry
	9,  // 4: logical.Group.metadata:type_name -> logical.Group.MetadataEntry
	3,  // 5: logical.MFAConstraintAny.any:type_name -> logical.MFAMethodID
	10, // 6: logical.MFARequirement.mfa_constraints:type_name -> logical.MFARequirement.MFAConstraintsEntry
	4,  // 7: logical.MFARequirement.MFAConstraintsEntry.value:type_name -> logical.MFAConstraintAny
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_sdk_logical_identity_proto_init() }
func file_sdk_logical_identity_proto_init() {
	if File_sdk_logical_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_logical_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_sdk_logical_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alias); i {
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
		file_sdk_logical_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_sdk_logical_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MFAMethodID); i {
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
		file_sdk_logical_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MFAConstraintAny); i {
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
		file_sdk_logical_identity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MFARequirement); i {
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
			RawDescriptor: file_sdk_logical_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_logical_identity_proto_goTypes,
		DependencyIndexes: file_sdk_logical_identity_proto_depIDxs,
		MessageInfos:      file_sdk_logical_identity_proto_msgTypes,
	}.Build()
	File_sdk_logical_identity_proto = out.File
	file_sdk_logical_identity_proto_rawDesc = nil
	file_sdk_logical_identity_proto_goTypes = nil
	file_sdk_logical_identity_proto_depIDxs = nil
}
