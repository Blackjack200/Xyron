// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: primitive_types.proto

package xyron

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

type Vec2F struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Vec2F) Reset() {
	*x = Vec2F{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec2F) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec2F) ProtoMessage() {}

func (x *Vec2F) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec2F.ProtoReflect.Descriptor instead.
func (*Vec2F) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{0}
}

func (x *Vec2F) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec2F) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Vec2I struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Vec2I) Reset() {
	*x = Vec2I{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec2I) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec2I) ProtoMessage() {}

func (x *Vec2I) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec2I.ProtoReflect.Descriptor instead.
func (*Vec2I) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{1}
}

func (x *Vec2I) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec2I) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Vec3F struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vec3F) Reset() {
	*x = Vec3F{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec3F) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec3F) ProtoMessage() {}

func (x *Vec3F) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec3F.ProtoReflect.Descriptor instead.
func (*Vec3F) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{2}
}

func (x *Vec3F) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec3F) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vec3F) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Vec3I struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vec3I) Reset() {
	*x = Vec3I{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec3I) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec3I) ProtoMessage() {}

func (x *Vec3I) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec3I.ProtoReflect.Descriptor instead.
func (*Vec3I) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{3}
}

func (x *Vec3I) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec3I) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vec3I) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Loc3F struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position  *Vec3F `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Direction *Vec3F `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *Loc3F) Reset() {
	*x = Loc3F{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loc3F) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loc3F) ProtoMessage() {}

func (x *Loc3F) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loc3F.ProtoReflect.Descriptor instead.
func (*Loc3F) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{4}
}

func (x *Loc3F) GetPosition() *Vec3F {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Loc3F) GetDirection() *Vec3F {
	if x != nil {
		return x.Direction
	}
	return nil
}

type AxisAlignedBoundingBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *Vec3F `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"`
	Max *Vec3F `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *AxisAlignedBoundingBox) Reset() {
	*x = AxisAlignedBoundingBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AxisAlignedBoundingBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AxisAlignedBoundingBox) ProtoMessage() {}

func (x *AxisAlignedBoundingBox) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AxisAlignedBoundingBox.ProtoReflect.Descriptor instead.
func (*AxisAlignedBoundingBox) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{5}
}

func (x *AxisAlignedBoundingBox) GetMin() *Vec3F {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *AxisAlignedBoundingBox) GetMax() *Vec3F {
	if x != nil {
		return x.Max
	}
	return nil
}

type BlockData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelativePosition *Vec3I `protobuf:"bytes,1,opt,name=relativePosition,proto3" json:"relativePosition,omitempty"`
	Feature          string `protobuf:"bytes,2,opt,name=feature,proto3" json:"feature,omitempty"` //registered in BlockFeatureRegistry
}

func (x *BlockData) Reset() {
	*x = BlockData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockData) ProtoMessage() {}

func (x *BlockData) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockData.ProtoReflect.Descriptor instead.
func (*BlockData) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{6}
}

func (x *BlockData) GetRelativePosition() *Vec3I {
	if x != nil {
		return x.RelativePosition
	}
	return nil
}

func (x *BlockData) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

type BlockFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollisionBoxes []*AxisAlignedBoundingBox `protobuf:"bytes,2,rep,name=collisionBoxes,proto3" json:"collisionBoxes,omitempty"`
	Friction       float32                   `protobuf:"fixed32,3,opt,name=friction,proto3" json:"friction,omitempty"`
	IsSolid        bool                      `protobuf:"varint,4,opt,name=isSolid,proto3" json:"isSolid,omitempty"`
	IsLiquid       bool                      `protobuf:"varint,5,opt,name=isLiquid,proto3" json:"isLiquid,omitempty"`
	IsAir          bool                      `protobuf:"varint,6,opt,name=isAir,proto3" json:"isAir,omitempty"`
	IsSlime        bool                      `protobuf:"varint,7,opt,name=isSlime,proto3" json:"isSlime,omitempty"`
	IsClimbable    bool                      `protobuf:"varint,8,opt,name=isClimbable,proto3" json:"isClimbable,omitempty"`
}

func (x *BlockFeature) Reset() {
	*x = BlockFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockFeature) ProtoMessage() {}

func (x *BlockFeature) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockFeature.ProtoReflect.Descriptor instead.
func (*BlockFeature) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{7}
}

func (x *BlockFeature) GetCollisionBoxes() []*AxisAlignedBoundingBox {
	if x != nil {
		return x.CollisionBoxes
	}
	return nil
}

func (x *BlockFeature) GetFriction() float32 {
	if x != nil {
		return x.Friction
	}
	return 0
}

func (x *BlockFeature) GetIsSolid() bool {
	if x != nil {
		return x.IsSolid
	}
	return false
}

func (x *BlockFeature) GetIsLiquid() bool {
	if x != nil {
		return x.IsLiquid
	}
	return false
}

func (x *BlockFeature) GetIsAir() bool {
	if x != nil {
		return x.IsAir
	}
	return false
}

func (x *BlockFeature) GetIsSlime() bool {
	if x != nil {
		return x.IsSlime
	}
	return false
}

func (x *BlockFeature) GetIsClimbable() bool {
	if x != nil {
		return x.IsClimbable
	}
	return false
}

type BlockFeatureRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features map[string]*BlockFeature `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BlockFeatureRegistry) Reset() {
	*x = BlockFeatureRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockFeatureRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockFeatureRegistry) ProtoMessage() {}

func (x *BlockFeatureRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockFeatureRegistry.ProtoReflect.Descriptor instead.
func (*BlockFeatureRegistry) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{8}
}

func (x *BlockFeatureRegistry) GetFeatures() map[string]*BlockFeature {
	if x != nil {
		return x.Features
	}
	return nil
}

type PotionEffectRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Effects []string `protobuf:"bytes,1,rep,name=effects,proto3" json:"effects,omitempty"`
}

func (x *PotionEffectRegistry) Reset() {
	*x = PotionEffectRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PotionEffectRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PotionEffectRegistry) ProtoMessage() {}

func (x *PotionEffectRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PotionEffectRegistry.ProtoReflect.Descriptor instead.
func (*PotionEffectRegistry) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{9}
}

func (x *PotionEffectRegistry) GetEffects() []string {
	if x != nil {
		return x.Effects
	}
	return nil
}

type ItemData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	VanillaName string `protobuf:"bytes,2,opt,name=vanillaName,proto3" json:"vanillaName,omitempty"`
	Feature     string `protobuf:"bytes,3,opt,name=feature,proto3" json:"feature,omitempty"` //registered in ItemFeatureRegistry
}

func (x *ItemData) Reset() {
	*x = ItemData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemData) ProtoMessage() {}

func (x *ItemData) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemData.ProtoReflect.Descriptor instead.
func (*ItemData) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{10}
}

func (x *ItemData) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ItemData) GetVanillaName() string {
	if x != nil {
		return x.VanillaName
	}
	return ""
}

func (x *ItemData) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

type ItemFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsArmor            bool `protobuf:"varint,1,opt,name=isArmor,proto3" json:"isArmor,omitempty"`
	IsBlockPlanterItem bool `protobuf:"varint,2,opt,name=isBlockPlanterItem,proto3" json:"isBlockPlanterItem,omitempty"`
	IsDamageable       bool `protobuf:"varint,3,opt,name=isDamageable,proto3" json:"isDamageable,omitempty"`
	IsFood             bool `protobuf:"varint,4,opt,name=isFood,proto3" json:"isFood,omitempty"`
	IsThrowable        bool `protobuf:"varint,5,opt,name=isThrowable,proto3" json:"isThrowable,omitempty"`
	IsTool             bool `protobuf:"varint,6,opt,name=isTool,proto3" json:"isTool,omitempty"`
	IsBow              bool `protobuf:"varint,7,opt,name=isBow,proto3" json:"isBow,omitempty"`
	IsCrossBow         bool `protobuf:"varint,8,opt,name=isCrossBow,proto3" json:"isCrossBow,omitempty"`
}

func (x *ItemFeature) Reset() {
	*x = ItemFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemFeature) ProtoMessage() {}

func (x *ItemFeature) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemFeature.ProtoReflect.Descriptor instead.
func (*ItemFeature) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{11}
}

func (x *ItemFeature) GetIsArmor() bool {
	if x != nil {
		return x.IsArmor
	}
	return false
}

func (x *ItemFeature) GetIsBlockPlanterItem() bool {
	if x != nil {
		return x.IsBlockPlanterItem
	}
	return false
}

func (x *ItemFeature) GetIsDamageable() bool {
	if x != nil {
		return x.IsDamageable
	}
	return false
}

func (x *ItemFeature) GetIsFood() bool {
	if x != nil {
		return x.IsFood
	}
	return false
}

func (x *ItemFeature) GetIsThrowable() bool {
	if x != nil {
		return x.IsThrowable
	}
	return false
}

func (x *ItemFeature) GetIsTool() bool {
	if x != nil {
		return x.IsTool
	}
	return false
}

func (x *ItemFeature) GetIsBow() bool {
	if x != nil {
		return x.IsBow
	}
	return false
}

func (x *ItemFeature) GetIsCrossBow() bool {
	if x != nil {
		return x.IsCrossBow
	}
	return false
}

type ItemFeatureRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features map[string]*ItemFeatureRegistry `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ItemFeatureRegistry) Reset() {
	*x = ItemFeatureRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primitive_types_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemFeatureRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemFeatureRegistry) ProtoMessage() {}

func (x *ItemFeatureRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_primitive_types_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemFeatureRegistry.ProtoReflect.Descriptor instead.
func (*ItemFeatureRegistry) Descriptor() ([]byte, []int) {
	return file_primitive_types_proto_rawDescGZIP(), []int{12}
}

func (x *ItemFeatureRegistry) GetFeatures() map[string]*ItemFeatureRegistry {
	if x != nil {
		return x.Features
	}
	return nil
}

var File_primitive_types_proto protoreflect.FileDescriptor

var file_primitive_types_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x22, 0x23, 0x0a, 0x05, 0x56, 0x65, 0x63, 0x32, 0x66, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x23, 0x0a, 0x05, 0x56, 0x65, 0x63, 0x32, 0x69, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x31, 0x0a, 0x05, 0x56, 0x65,
	0x63, 0x33, 0x66, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12,
	0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x22, 0x31, 0x0a,
	0x05, 0x56, 0x65, 0x63, 0x33, 0x69, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a,
	0x22, 0x61, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x33, 0x66, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x66, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x66, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x16, 0x41, 0x78, 0x69, 0x73, 0x41, 0x6c, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x20, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x66, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12,
	0x20, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x66, 0x52, 0x03, 0x6d, 0x61,
	0x78, 0x22, 0x61, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a,
	0x0a, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x69, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x78, 0x69, 0x73, 0x41, 0x6c, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x52, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x66, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73,
	0x53, 0x6f, 0x6c, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53,
	0x6f, 0x6c, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x41, 0x69, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x41, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x53, 0x6c, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53, 0x6c, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x6d, 0x62, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x6d, 0x62, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x1a, 0x52, 0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x14, 0x50, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x08, 0x49, 0x74,
	0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x0b, 0x49, 0x74, 0x65,
	0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x44, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x46, 0x6f, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x77, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x77, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x54, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x54, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x6f,
	0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x6f, 0x77, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x42, 0x6f, 0x77, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x42, 0x6f, 0x77, 0x22, 0xb8,
	0x01, 0x0a, 0x13, 0x49, 0x74, 0x65, 0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x46, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x59,
	0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3d, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63,
	0x6b, 0x32, 0x30, 0x30, 0x2e, 0x78, 0x79, 0x72, 0x6f, 0x6e, 0x5a, 0x06, 0x78, 0x79, 0x72, 0x6f,
	0x6e, 0x2f, 0x88, 0x01, 0x01, 0xca, 0x02, 0x0d, 0x70, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x73, 0x5c,
	0x78, 0x79, 0x72, 0x6f, 0x6e, 0xd0, 0x02, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_primitive_types_proto_rawDescOnce sync.Once
	file_primitive_types_proto_rawDescData = file_primitive_types_proto_rawDesc
)

func file_primitive_types_proto_rawDescGZIP() []byte {
	file_primitive_types_proto_rawDescOnce.Do(func() {
		file_primitive_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_primitive_types_proto_rawDescData)
	})
	return file_primitive_types_proto_rawDescData
}

var file_primitive_types_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_primitive_types_proto_goTypes = []interface{}{
	(*Vec2F)(nil),                  // 0: xchange.Vec2f
	(*Vec2I)(nil),                  // 1: xchange.Vec2i
	(*Vec3F)(nil),                  // 2: xchange.Vec3f
	(*Vec3I)(nil),                  // 3: xchange.Vec3i
	(*Loc3F)(nil),                  // 4: xchange.Loc3f
	(*AxisAlignedBoundingBox)(nil), // 5: xchange.AxisAlignedBoundingBox
	(*BlockData)(nil),              // 6: xchange.BlockData
	(*BlockFeature)(nil),           // 7: xchange.BlockFeature
	(*BlockFeatureRegistry)(nil),   // 8: xchange.BlockFeatureRegistry
	(*PotionEffectRegistry)(nil),   // 9: xchange.PotionEffectRegistry
	(*ItemData)(nil),               // 10: xchange.ItemData
	(*ItemFeature)(nil),            // 11: xchange.ItemFeature
	(*ItemFeatureRegistry)(nil),    // 12: xchange.ItemFeatureRegistry
	nil,                            // 13: xchange.BlockFeatureRegistry.FeaturesEntry
	nil,                            // 14: xchange.ItemFeatureRegistry.FeaturesEntry
}
var file_primitive_types_proto_depIdxs = []int32{
	2,  // 0: xchange.Loc3f.position:type_name -> xchange.Vec3f
	2,  // 1: xchange.Loc3f.direction:type_name -> xchange.Vec3f
	2,  // 2: xchange.AxisAlignedBoundingBox.min:type_name -> xchange.Vec3f
	2,  // 3: xchange.AxisAlignedBoundingBox.max:type_name -> xchange.Vec3f
	3,  // 4: xchange.BlockData.relativePosition:type_name -> xchange.Vec3i
	5,  // 5: xchange.BlockFeature.collisionBoxes:type_name -> xchange.AxisAlignedBoundingBox
	13, // 6: xchange.BlockFeatureRegistry.features:type_name -> xchange.BlockFeatureRegistry.FeaturesEntry
	14, // 7: xchange.ItemFeatureRegistry.features:type_name -> xchange.ItemFeatureRegistry.FeaturesEntry
	7,  // 8: xchange.BlockFeatureRegistry.FeaturesEntry.value:type_name -> xchange.BlockFeature
	12, // 9: xchange.ItemFeatureRegistry.FeaturesEntry.value:type_name -> xchange.ItemFeatureRegistry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_primitive_types_proto_init() }
func file_primitive_types_proto_init() {
	if File_primitive_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_primitive_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec2F); i {
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
		file_primitive_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec2I); i {
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
		file_primitive_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec3F); i {
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
		file_primitive_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec3I); i {
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
		file_primitive_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loc3F); i {
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
		file_primitive_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AxisAlignedBoundingBox); i {
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
		file_primitive_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockData); i {
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
		file_primitive_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockFeature); i {
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
		file_primitive_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockFeatureRegistry); i {
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
		file_primitive_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PotionEffectRegistry); i {
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
		file_primitive_types_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemData); i {
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
		file_primitive_types_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemFeature); i {
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
		file_primitive_types_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemFeatureRegistry); i {
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
			RawDescriptor: file_primitive_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_primitive_types_proto_goTypes,
		DependencyIndexes: file_primitive_types_proto_depIdxs,
		MessageInfos:      file_primitive_types_proto_msgTypes,
	}.Build()
	File_primitive_types_proto = out.File
	file_primitive_types_proto_rawDesc = nil
	file_primitive_types_proto_goTypes = nil
	file_primitive_types_proto_depIdxs = nil
}
