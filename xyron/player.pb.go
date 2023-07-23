// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: player.proto

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

type GameMode int32

const (
	GameMode_Survival  GameMode = 0
	GameMode_Creative  GameMode = 1
	GameMode_Adventure GameMode = 2
	GameMode_Spectator GameMode = 3
)

// Enum value maps for GameMode.
var (
	GameMode_name = map[int32]string{
		0: "Survival",
		1: "Creative",
		2: "Adventure",
		3: "Spectator",
	}
	GameMode_value = map[string]int32{
		"Survival":  0,
		"Creative":  1,
		"Adventure": 2,
		"Spectator": 3,
	}
)

func (x GameMode) Enum() *GameMode {
	p := new(GameMode)
	*p = x
	return p
}

func (x GameMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameMode) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[0].Descriptor()
}

func (GameMode) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[0]
}

func (x GameMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameMode.Descriptor instead.
func (GameMode) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

type InputMode int32

const (
	InputMode_MouseKeyboard    InputMode = 0
	InputMode_Touch            InputMode = 1
	InputMode_Gamepad          InputMode = 2
	InputMode_MotionController InputMode = 3
)

// Enum value maps for InputMode.
var (
	InputMode_name = map[int32]string{
		0: "MouseKeyboard",
		1: "Touch",
		2: "Gamepad",
		3: "MotionController",
	}
	InputMode_value = map[string]int32{
		"MouseKeyboard":    0,
		"Touch":            1,
		"Gamepad":          2,
		"MotionController": 3,
	}
)

func (x InputMode) Enum() *InputMode {
	p := new(InputMode)
	*p = x
	return p
}

func (x InputMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InputMode) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[1].Descriptor()
}

func (InputMode) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[1]
}

func (x InputMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InputMode.Descriptor instead.
func (InputMode) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

type DeviceOS int32

const (
	DeviceOS_Android       DeviceOS = 0
	DeviceOS_IOS           DeviceOS = 1
	DeviceOS_OSX           DeviceOS = 2
	DeviceOS_AMAZON        DeviceOS = 3
	DeviceOS_GEAR_VR       DeviceOS = 4
	DeviceOS_HOLOLENS      DeviceOS = 5
	DeviceOS_WINDOWS_10    DeviceOS = 6
	DeviceOS_WIN32         DeviceOS = 7
	DeviceOS_DEDICATED     DeviceOS = 8
	DeviceOS_TVOS          DeviceOS = 9
	DeviceOS_PLAYSTATION   DeviceOS = 10
	DeviceOS_NINTENDO      DeviceOS = 11
	DeviceOS_XBOX          DeviceOS = 12
	DeviceOS_WINDOWS_PHONE DeviceOS = 13
)

// Enum value maps for DeviceOS.
var (
	DeviceOS_name = map[int32]string{
		0:  "Android",
		1:  "IOS",
		2:  "OSX",
		3:  "AMAZON",
		4:  "GEAR_VR",
		5:  "HOLOLENS",
		6:  "WINDOWS_10",
		7:  "WIN32",
		8:  "DEDICATED",
		9:  "TVOS",
		10: "PLAYSTATION",
		11: "NINTENDO",
		12: "XBOX",
		13: "WINDOWS_PHONE",
	}
	DeviceOS_value = map[string]int32{
		"Android":       0,
		"IOS":           1,
		"OSX":           2,
		"AMAZON":        3,
		"GEAR_VR":       4,
		"HOLOLENS":      5,
		"WINDOWS_10":    6,
		"WIN32":         7,
		"DEDICATED":     8,
		"TVOS":          9,
		"PLAYSTATION":   10,
		"NINTENDO":      11,
		"XBOX":          12,
		"WINDOWS_PHONE": 13,
	}
)

func (x DeviceOS) Enum() *DeviceOS {
	p := new(DeviceOS)
	*p = x
	return p
}

func (x DeviceOS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceOS) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[2].Descriptor()
}

func (DeviceOS) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[2]
}

func (x DeviceOS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceOS.Descriptor instead.
func (DeviceOS) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{2}
}

type PlayerAction int32

const (
	PlayerAction_Jump              PlayerAction = 0
	PlayerAction_Swing             PlayerAction = 1
	PlayerAction_StartSprint       PlayerAction = 2
	PlayerAction_StopSprint        PlayerAction = 3
	PlayerAction_StartSneak        PlayerAction = 4
	PlayerAction_StopSneak         PlayerAction = 5
	PlayerAction_StartSprintFlying PlayerAction = 6
	PlayerAction_StopSprintFlying  PlayerAction = 7
	PlayerAction_StartGliding      PlayerAction = 8
	PlayerAction_StopGliding       PlayerAction = 9
	PlayerAction_StartSwimming     PlayerAction = 10
	PlayerAction_StopSwimming      PlayerAction = 11
)

// Enum value maps for PlayerAction.
var (
	PlayerAction_name = map[int32]string{
		0:  "Jump",
		1:  "Swing",
		2:  "StartSprint",
		3:  "StopSprint",
		4:  "StartSneak",
		5:  "StopSneak",
		6:  "StartSprintFlying",
		7:  "StopSprintFlying",
		8:  "StartGliding",
		9:  "StopGliding",
		10: "StartSwimming",
		11: "StopSwimming",
	}
	PlayerAction_value = map[string]int32{
		"Jump":              0,
		"Swing":             1,
		"StartSprint":       2,
		"StopSprint":        3,
		"StartSneak":        4,
		"StopSneak":         5,
		"StartSprintFlying": 6,
		"StopSprintFlying":  7,
		"StartGliding":      8,
		"StopGliding":       9,
		"StartSwimming":     10,
		"StopSwimming":      11,
	}
)

func (x PlayerAction) Enum() *PlayerAction {
	p := new(PlayerAction)
	*p = x
	return p
}

func (x PlayerAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerAction) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[3].Descriptor()
}

func (PlayerAction) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[3]
}

func (x PlayerAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerAction.Descriptor instead.
func (PlayerAction) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{3}
}

type DamageCause int32

const (
	DamageCause_Contact         DamageCause = 0
	DamageCause_EntityAttack    DamageCause = 1
	DamageCause_Projectile      DamageCause = 2
	DamageCause_Suffocation     DamageCause = 3
	DamageCause_Fall            DamageCause = 4
	DamageCause_Fire            DamageCause = 5
	DamageCause_FireTick        DamageCause = 6
	DamageCause_Lava            DamageCause = 7
	DamageCause_Drowning        DamageCause = 8
	DamageCause_BlockExplosion  DamageCause = 9
	DamageCause_EntityExplosion DamageCause = 10
	DamageCause_Void            DamageCause = 11
	DamageCause_Suicide         DamageCause = 12
	DamageCause_Magic           DamageCause = 13
	DamageCause_Custom          DamageCause = 14
	DamageCause_Starvation      DamageCause = 15
	DamageCause_FallingBlock    DamageCause = 16
)

// Enum value maps for DamageCause.
var (
	DamageCause_name = map[int32]string{
		0:  "Contact",
		1:  "EntityAttack",
		2:  "Projectile",
		3:  "Suffocation",
		4:  "Fall",
		5:  "Fire",
		6:  "FireTick",
		7:  "Lava",
		8:  "Drowning",
		9:  "BlockExplosion",
		10: "EntityExplosion",
		11: "Void",
		12: "Suicide",
		13: "Magic",
		14: "Custom",
		15: "Starvation",
		16: "FallingBlock",
	}
	DamageCause_value = map[string]int32{
		"Contact":         0,
		"EntityAttack":    1,
		"Projectile":      2,
		"Suffocation":     3,
		"Fall":            4,
		"Fire":            5,
		"FireTick":        6,
		"Lava":            7,
		"Drowning":        8,
		"BlockExplosion":  9,
		"EntityExplosion": 10,
		"Void":            11,
		"Suicide":         12,
		"Magic":           13,
		"Custom":          14,
		"Starvation":      15,
		"FallingBlock":    16,
	}
)

func (x DamageCause) Enum() *DamageCause {
	p := new(DamageCause)
	*p = x
	return p
}

func (x DamageCause) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DamageCause) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[4].Descriptor()
}

func (DamageCause) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[4]
}

func (x DamageCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DamageCause.Descriptor instead.
func (DamageCause) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{4}
}

type BreakBlockStatus int32

const (
	BreakBlockStatus_StartBreak  BreakBlockStatus = 0
	BreakBlockStatus_AbortBreak  BreakBlockStatus = 1
	BreakBlockStatus_FinishBreak BreakBlockStatus = 2
)

// Enum value maps for BreakBlockStatus.
var (
	BreakBlockStatus_name = map[int32]string{
		0: "StartBreak",
		1: "AbortBreak",
		2: "FinishBreak",
	}
	BreakBlockStatus_value = map[string]int32{
		"StartBreak":  0,
		"AbortBreak":  1,
		"FinishBreak": 2,
	}
)

func (x BreakBlockStatus) Enum() *BreakBlockStatus {
	p := new(BreakBlockStatus)
	*p = x
	return p
}

func (x BreakBlockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BreakBlockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[5].Descriptor()
}

func (BreakBlockStatus) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[5]
}

func (x BreakBlockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BreakBlockStatus.Descriptor instead.
func (BreakBlockStatus) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{5}
}

type ConsumeStatus int32

const (
	ConsumeStatus_Start ConsumeStatus = 0
	ConsumeStatus_Stop  ConsumeStatus = 1
)

// Enum value maps for ConsumeStatus.
var (
	ConsumeStatus_name = map[int32]string{
		0: "Start",
		1: "Stop",
	}
	ConsumeStatus_value = map[string]int32{
		"Start": 0,
		"Stop":  1,
	}
)

func (x ConsumeStatus) Enum() *ConsumeStatus {
	p := new(ConsumeStatus)
	*p = x
	return p
}

func (x ConsumeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsumeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[6].Descriptor()
}

func (ConsumeStatus) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[6]
}

func (x ConsumeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsumeStatus.Descriptor instead.
func (ConsumeStatus) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{6}
}

type EntityPositionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location          *Loc3F                  `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	BoundingBox       *AxisAlignedBoundingBox `protobuf:"bytes,2,opt,name=boundingBox,proto3" json:"boundingBox,omitempty"`
	Below             *BlockData              `protobuf:"bytes,3,opt,name=below,proto3" json:"below,omitempty"`
	IsImmobile        bool                    `protobuf:"varint,4,opt,name=isImmobile,proto3" json:"isImmobile,omitempty"`
	IsOnGround        bool                    `protobuf:"varint,5,opt,name=isOnGround,proto3" json:"isOnGround,omitempty"`
	AllowFlying       bool                    `protobuf:"varint,6,opt,name=allowFlying,proto3" json:"allowFlying,omitempty"`
	IsFlying          bool                    `protobuf:"varint,7,opt,name=isFlying,proto3" json:"isFlying,omitempty"`
	CollidedBlocks    []*BlockData            `protobuf:"bytes,8,rep,name=collidedBlocks,proto3" json:"collidedBlocks,omitempty"`
	IntersectedBlocks []*BlockData            `protobuf:"bytes,9,rep,name=intersectedBlocks,proto3" json:"intersectedBlocks,omitempty"`
}

func (x *EntityPositionData) Reset() {
	*x = EntityPositionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityPositionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityPositionData) ProtoMessage() {}

func (x *EntityPositionData) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityPositionData.ProtoReflect.Descriptor instead.
func (*EntityPositionData) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

func (x *EntityPositionData) GetLocation() *Loc3F {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *EntityPositionData) GetBoundingBox() *AxisAlignedBoundingBox {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *EntityPositionData) GetBelow() *BlockData {
	if x != nil {
		return x.Below
	}
	return nil
}

func (x *EntityPositionData) GetIsImmobile() bool {
	if x != nil {
		return x.IsImmobile
	}
	return false
}

func (x *EntityPositionData) GetIsOnGround() bool {
	if x != nil {
		return x.IsOnGround
	}
	return false
}

func (x *EntityPositionData) GetAllowFlying() bool {
	if x != nil {
		return x.AllowFlying
	}
	return false
}

func (x *EntityPositionData) GetIsFlying() bool {
	if x != nil {
		return x.IsFlying
	}
	return false
}

func (x *EntityPositionData) GetCollidedBlocks() []*BlockData {
	if x != nil {
		return x.CollidedBlocks
	}
	return nil
}

func (x *EntityPositionData) GetIntersectedBlocks() []*BlockData {
	if x != nil {
		return x.IntersectedBlocks
	}
	return nil
}

type AttackData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cause    DamageCause         `protobuf:"varint,1,opt,name=cause,proto3,enum=xchange.DamageCause" json:"cause,omitempty"`
	Attacker *EntityPositionData `protobuf:"bytes,2,opt,name=attacker,proto3" json:"attacker,omitempty"`
	Target   *EntityPositionData `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AttackData) Reset() {
	*x = AttackData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackData) ProtoMessage() {}

func (x *AttackData) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackData.ProtoReflect.Descriptor instead.
func (*AttackData) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

func (x *AttackData) GetCause() DamageCause {
	if x != nil {
		return x.Cause
	}
	return DamageCause_Contact
}

func (x *AttackData) GetAttacker() *EntityPositionData {
	if x != nil {
		return x.Attacker
	}
	return nil
}

func (x *AttackData) GetTarget() *EntityPositionData {
	if x != nil {
		return x.Target
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os   DeviceOS `protobuf:"varint,1,opt,name=os,proto3,enum=xchange.DeviceOS" json:"os,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetOs() DeviceOS {
	if x != nil {
		return x.Os
	}
	return DeviceOS_Android
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_player_proto protoreflect.FileDescriptor

var file_player_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x15, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9,
	0x03, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x33, 0x66, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x41, 0x78, 0x69, 0x73, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x42, 0x6f, 0x78, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x65, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x62, 0x65, 0x6c, 0x6f, 0x77, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x49, 0x6d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x49, 0x6d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x4f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x64,
	0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x40, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x0a, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x05, 0x63, 0x61, 0x75,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x05,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x33,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0x3f, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x44, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x03, 0x2a, 0x4c, 0x0a, 0x09, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x6f, 0x75, 0x73, 0x65,
	0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x6f,
	0x75, 0x63, 0x68, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x61, 0x64,
	0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x10, 0x03, 0x2a, 0xc0, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x53, 0x58, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4d, 0x41, 0x5a, 0x4f, 0x4e, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x47, 0x45, 0x41, 0x52, 0x5f, 0x56, 0x52, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x48, 0x4f, 0x4c, 0x4f, 0x4c, 0x45, 0x4e, 0x53, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x5f, 0x31, 0x30, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x57,
	0x49, 0x4e, 0x33, 0x32, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x44, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x56, 0x4f, 0x53, 0x10, 0x09, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x4c, 0x41, 0x59, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a,
	0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x44, 0x4f, 0x10, 0x0b, 0x12, 0x08,
	0x0a, 0x04, 0x58, 0x42, 0x4f, 0x58, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x53, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x0d, 0x2a, 0xd8, 0x01, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04,
	0x4a, 0x75, 0x6d, 0x70, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x77, 0x69, 0x6e, 0x67, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6e, 0x65, 0x61, 0x6b,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x10,
	0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x70,
	0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x10, 0x07, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x6c, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x08,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x47, 0x6c, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x10,
	0x09, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x77, 0x69, 0x6d, 0x6d, 0x69,
	0x6e, 0x67, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x77, 0x69, 0x6d,
	0x6d, 0x69, 0x6e, 0x67, 0x10, 0x0b, 0x2a, 0x80, 0x02, 0x0a, 0x0b, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x75, 0x66, 0x66, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x6c, 0x6c, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x72, 0x65, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x69,
	0x72, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x61, 0x76, 0x61,
	0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x72, 0x6f, 0x77, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x08,
	0x12, 0x12, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x73, 0x69,
	0x6f, 0x6e, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x78,
	0x70, 0x6c, 0x6f, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x10, 0x0c,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x67, 0x69, 0x63, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x61, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x10, 0x2a, 0x43, 0x0a, 0x10, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x10, 0x02, 0x2a, 0x24,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x74,
	0x6f, 0x70, 0x10, 0x01, 0x42, 0x3d, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x32, 0x30, 0x30, 0x2e,
	0x78, 0x79, 0x72, 0x6f, 0x6e, 0x5a, 0x06, 0x78, 0x79, 0x72, 0x6f, 0x6e, 0x2f, 0x88, 0x01, 0x01,
	0xca, 0x02, 0x0d, 0x70, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x73, 0x5c, 0x78, 0x79, 0x72, 0x6f, 0x6e,
	0xd0, 0x02, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_proto_rawDescOnce sync.Once
	file_player_proto_rawDescData = file_player_proto_rawDesc
)

func file_player_proto_rawDescGZIP() []byte {
	file_player_proto_rawDescOnce.Do(func() {
		file_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_proto_rawDescData)
	})
	return file_player_proto_rawDescData
}

var file_player_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_player_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_player_proto_goTypes = []interface{}{
	(GameMode)(0),                  // 0: xchange.GameMode
	(InputMode)(0),                 // 1: xchange.InputMode
	(DeviceOS)(0),                  // 2: xchange.DeviceOS
	(PlayerAction)(0),              // 3: xchange.PlayerAction
	(DamageCause)(0),               // 4: xchange.DamageCause
	(BreakBlockStatus)(0),          // 5: xchange.BreakBlockStatus
	(ConsumeStatus)(0),             // 6: xchange.ConsumeStatus
	(*EntityPositionData)(nil),     // 7: xchange.EntityPositionData
	(*AttackData)(nil),             // 8: xchange.AttackData
	(*Player)(nil),                 // 9: xchange.Player
	(*Loc3F)(nil),                  // 10: xchange.Loc3f
	(*AxisAlignedBoundingBox)(nil), // 11: xchange.AxisAlignedBoundingBox
	(*BlockData)(nil),              // 12: xchange.BlockData
}
var file_player_proto_depIdxs = []int32{
	10, // 0: xchange.EntityPositionData.location:type_name -> xchange.Loc3f
	11, // 1: xchange.EntityPositionData.boundingBox:type_name -> xchange.AxisAlignedBoundingBox
	12, // 2: xchange.EntityPositionData.below:type_name -> xchange.BlockData
	12, // 3: xchange.EntityPositionData.collidedBlocks:type_name -> xchange.BlockData
	12, // 4: xchange.EntityPositionData.intersectedBlocks:type_name -> xchange.BlockData
	4,  // 5: xchange.AttackData.cause:type_name -> xchange.DamageCause
	7,  // 6: xchange.AttackData.attacker:type_name -> xchange.EntityPositionData
	7,  // 7: xchange.AttackData.target:type_name -> xchange.EntityPositionData
	2,  // 8: xchange.Player.os:type_name -> xchange.DeviceOS
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_player_proto_init() }
func file_player_proto_init() {
	if File_player_proto != nil {
		return
	}
	file_primitive_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityPositionData); i {
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
		file_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackData); i {
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
		file_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_player_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_proto_goTypes,
		DependencyIndexes: file_player_proto_depIdxs,
		EnumInfos:         file_player_proto_enumTypes,
		MessageInfos:      file_player_proto_msgTypes,
	}.Build()
	File_player_proto = out.File
	file_player_proto_rawDesc = nil
	file_player_proto_goTypes = nil
	file_player_proto_depIdxs = nil
}
