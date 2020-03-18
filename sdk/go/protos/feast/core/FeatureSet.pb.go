// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/FeatureSet.proto

package core

import (
	fmt "fmt"
	types "github.com/gojek/feast/sdk/go/protos/feast/types"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	v0 "tensorflow_metadata/proto/v0"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FeatureSetStatus int32

const (
	FeatureSetStatus_STATUS_INVALID FeatureSetStatus = 0
	FeatureSetStatus_STATUS_PENDING FeatureSetStatus = 1
	FeatureSetStatus_STATUS_READY   FeatureSetStatus = 2
)

var FeatureSetStatus_name = map[int32]string{
	0: "STATUS_INVALID",
	1: "STATUS_PENDING",
	2: "STATUS_READY",
}

var FeatureSetStatus_value = map[string]int32{
	"STATUS_INVALID": 0,
	"STATUS_PENDING": 1,
	"STATUS_READY":   2,
}

func (x FeatureSetStatus) String() string {
	return proto.EnumName(FeatureSetStatus_name, int32(x))
}

func (FeatureSetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_972fbd278ac19c0c, []int{0}
}

type FeatureSet struct {
	// User-specified specifications of this feature set.
	Spec *FeatureSetSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// System-populated metadata for this feature set.
	Meta                 *FeatureSetMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FeatureSet) Reset()         { *m = FeatureSet{} }
func (m *FeatureSet) String() string { return proto.CompactTextString(m) }
func (*FeatureSet) ProtoMessage()    {}
func (*FeatureSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_972fbd278ac19c0c, []int{0}
}

func (m *FeatureSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSet.Unmarshal(m, b)
}
func (m *FeatureSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSet.Marshal(b, m, deterministic)
}
func (m *FeatureSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSet.Merge(m, src)
}
func (m *FeatureSet) XXX_Size() int {
	return xxx_messageInfo_FeatureSet.Size(m)
}
func (m *FeatureSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSet.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSet proto.InternalMessageInfo

func (m *FeatureSet) GetSpec() *FeatureSetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *FeatureSet) GetMeta() *FeatureSetMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type FeatureSetSpec struct {
	// Name of project that this feature set belongs to.
	Project string `protobuf:"bytes,7,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the feature set. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Feature set version.
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// List of entities contained within this featureSet.
	// This allows the feature to be used during joins between feature sets.
	// If the featureSet is ingested into a store that supports keys, this value
	// will be made a key.
	Entities []*EntitySpec `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	// List of features contained within this featureSet.
	Features []*FeatureSpec `protobuf:"bytes,4,rep,name=features,proto3" json:"features,omitempty"`
	// Features in this feature set will only be retrieved if they are found
	// after [time - max_age]. Missing or older feature values will be returned
	// as nulls and indicated to end user
	MaxAge *duration.Duration `protobuf:"bytes,5,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// Optional. Source on which feature rows can be found.
	// If not set, source will be set to the default value configured in Feast Core.
	Source               *Source  `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureSetSpec) Reset()         { *m = FeatureSetSpec{} }
func (m *FeatureSetSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSetSpec) ProtoMessage()    {}
func (*FeatureSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_972fbd278ac19c0c, []int{1}
}

func (m *FeatureSetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSetSpec.Unmarshal(m, b)
}
func (m *FeatureSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSetSpec.Marshal(b, m, deterministic)
}
func (m *FeatureSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSetSpec.Merge(m, src)
}
func (m *FeatureSetSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSetSpec.Size(m)
}
func (m *FeatureSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSetSpec proto.InternalMessageInfo

func (m *FeatureSetSpec) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *FeatureSetSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSetSpec) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FeatureSetSpec) GetEntities() []*EntitySpec {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *FeatureSetSpec) GetFeatures() []*FeatureSpec {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *FeatureSetSpec) GetMaxAge() *duration.Duration {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *FeatureSetSpec) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type EntitySpec struct {
	// Name of the entity.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value type of the feature.
	ValueType types.ValueType_Enum `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=feast.types.ValueType_Enum" json:"value_type,omitempty"`
	// Types that are valid to be assigned to PresenceConstraints:
	//	*EntitySpec_Presence
	//	*EntitySpec_GroupPresence
	PresenceConstraints isEntitySpec_PresenceConstraints `protobuf_oneof:"presence_constraints"`
	// The shape of the feature which governs the number of values that appear in
	// each example.
	//
	// Types that are valid to be assigned to ShapeType:
	//	*EntitySpec_Shape
	//	*EntitySpec_ValueCount
	ShapeType isEntitySpec_ShapeType `protobuf_oneof:"shape_type"`
	// Domain for the values of the feature.
	//
	// Types that are valid to be assigned to DomainInfo:
	//	*EntitySpec_Domain
	//	*EntitySpec_IntDomain
	//	*EntitySpec_FloatDomain
	//	*EntitySpec_StringDomain
	//	*EntitySpec_BoolDomain
	//	*EntitySpec_StructDomain
	//	*EntitySpec_NaturalLanguageDomain
	//	*EntitySpec_ImageDomain
	//	*EntitySpec_MidDomain
	//	*EntitySpec_UrlDomain
	//	*EntitySpec_TimeDomain
	//	*EntitySpec_TimeOfDayDomain
	DomainInfo           isEntitySpec_DomainInfo `protobuf_oneof:"domain_info"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EntitySpec) Reset()         { *m = EntitySpec{} }
func (m *EntitySpec) String() string { return proto.CompactTextString(m) }
func (*EntitySpec) ProtoMessage()    {}
func (*EntitySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_972fbd278ac19c0c, []int{2}
}

func (m *EntitySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntitySpec.Unmarshal(m, b)
}
func (m *EntitySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntitySpec.Marshal(b, m, deterministic)
}
func (m *EntitySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntitySpec.Merge(m, src)
}
func (m *EntitySpec) XXX_Size() int {
	return xxx_messageInfo_EntitySpec.Size(m)
}
func (m *EntitySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EntitySpec.DiscardUnknown(m)
}

var xxx_messageInfo_EntitySpec proto.InternalMessageInfo

func (m *EntitySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntitySpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_INVALID
}

type isEntitySpec_PresenceConstraints interface {
	isEntitySpec_PresenceConstraints()
}

type EntitySpec_Presence struct {
	Presence *v0.FeaturePresence `protobuf:"bytes,3,opt,name=presence,proto3,oneof"`
}

type EntitySpec_GroupPresence struct {
	GroupPresence *v0.FeaturePresenceWithinGroup `protobuf:"bytes,4,opt,name=group_presence,json=groupPresence,proto3,oneof"`
}

func (*EntitySpec_Presence) isEntitySpec_PresenceConstraints() {}

func (*EntitySpec_GroupPresence) isEntitySpec_PresenceConstraints() {}

func (m *EntitySpec) GetPresenceConstraints() isEntitySpec_PresenceConstraints {
	if m != nil {
		return m.PresenceConstraints
	}
	return nil
}

func (m *EntitySpec) GetPresence() *v0.FeaturePresence {
	if x, ok := m.GetPresenceConstraints().(*EntitySpec_Presence); ok {
		return x.Presence
	}
	return nil
}

func (m *EntitySpec) GetGroupPresence() *v0.FeaturePresenceWithinGroup {
	if x, ok := m.GetPresenceConstraints().(*EntitySpec_GroupPresence); ok {
		return x.GroupPresence
	}
	return nil
}

type isEntitySpec_ShapeType interface {
	isEntitySpec_ShapeType()
}

type EntitySpec_Shape struct {
	Shape *v0.FixedShape `protobuf:"bytes,5,opt,name=shape,proto3,oneof"`
}

type EntitySpec_ValueCount struct {
	ValueCount *v0.ValueCount `protobuf:"bytes,6,opt,name=value_count,json=valueCount,proto3,oneof"`
}

func (*EntitySpec_Shape) isEntitySpec_ShapeType() {}

func (*EntitySpec_ValueCount) isEntitySpec_ShapeType() {}

func (m *EntitySpec) GetShapeType() isEntitySpec_ShapeType {
	if m != nil {
		return m.ShapeType
	}
	return nil
}

func (m *EntitySpec) GetShape() *v0.FixedShape {
	if x, ok := m.GetShapeType().(*EntitySpec_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *EntitySpec) GetValueCount() *v0.ValueCount {
	if x, ok := m.GetShapeType().(*EntitySpec_ValueCount); ok {
		return x.ValueCount
	}
	return nil
}

type isEntitySpec_DomainInfo interface {
	isEntitySpec_DomainInfo()
}

type EntitySpec_Domain struct {
	Domain string `protobuf:"bytes,7,opt,name=domain,proto3,oneof"`
}

type EntitySpec_IntDomain struct {
	IntDomain *v0.IntDomain `protobuf:"bytes,8,opt,name=int_domain,json=intDomain,proto3,oneof"`
}

type EntitySpec_FloatDomain struct {
	FloatDomain *v0.FloatDomain `protobuf:"bytes,9,opt,name=float_domain,json=floatDomain,proto3,oneof"`
}

type EntitySpec_StringDomain struct {
	StringDomain *v0.StringDomain `protobuf:"bytes,10,opt,name=string_domain,json=stringDomain,proto3,oneof"`
}

type EntitySpec_BoolDomain struct {
	BoolDomain *v0.BoolDomain `protobuf:"bytes,11,opt,name=bool_domain,json=boolDomain,proto3,oneof"`
}

type EntitySpec_StructDomain struct {
	StructDomain *v0.StructDomain `protobuf:"bytes,12,opt,name=struct_domain,json=structDomain,proto3,oneof"`
}

type EntitySpec_NaturalLanguageDomain struct {
	NaturalLanguageDomain *v0.NaturalLanguageDomain `protobuf:"bytes,13,opt,name=natural_language_domain,json=naturalLanguageDomain,proto3,oneof"`
}

type EntitySpec_ImageDomain struct {
	ImageDomain *v0.ImageDomain `protobuf:"bytes,14,opt,name=image_domain,json=imageDomain,proto3,oneof"`
}

type EntitySpec_MidDomain struct {
	MidDomain *v0.MIDDomain `protobuf:"bytes,15,opt,name=mid_domain,json=midDomain,proto3,oneof"`
}

type EntitySpec_UrlDomain struct {
	UrlDomain *v0.URLDomain `protobuf:"bytes,16,opt,name=url_domain,json=urlDomain,proto3,oneof"`
}

type EntitySpec_TimeDomain struct {
	TimeDomain *v0.TimeDomain `protobuf:"bytes,17,opt,name=time_domain,json=timeDomain,proto3,oneof"`
}

type EntitySpec_TimeOfDayDomain struct {
	TimeOfDayDomain *v0.TimeOfDayDomain `protobuf:"bytes,18,opt,name=time_of_day_domain,json=timeOfDayDomain,proto3,oneof"`
}

func (*EntitySpec_Domain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_IntDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_FloatDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_StringDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_BoolDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_StructDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_NaturalLanguageDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_ImageDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_MidDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_UrlDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_TimeDomain) isEntitySpec_DomainInfo() {}

func (*EntitySpec_TimeOfDayDomain) isEntitySpec_DomainInfo() {}

func (m *EntitySpec) GetDomainInfo() isEntitySpec_DomainInfo {
	if m != nil {
		return m.DomainInfo
	}
	return nil
}

func (m *EntitySpec) GetDomain() string {
	if x, ok := m.GetDomainInfo().(*EntitySpec_Domain); ok {
		return x.Domain
	}
	return ""
}

func (m *EntitySpec) GetIntDomain() *v0.IntDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_IntDomain); ok {
		return x.IntDomain
	}
	return nil
}

func (m *EntitySpec) GetFloatDomain() *v0.FloatDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_FloatDomain); ok {
		return x.FloatDomain
	}
	return nil
}

func (m *EntitySpec) GetStringDomain() *v0.StringDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_StringDomain); ok {
		return x.StringDomain
	}
	return nil
}

func (m *EntitySpec) GetBoolDomain() *v0.BoolDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_BoolDomain); ok {
		return x.BoolDomain
	}
	return nil
}

func (m *EntitySpec) GetStructDomain() *v0.StructDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_StructDomain); ok {
		return x.StructDomain
	}
	return nil
}

func (m *EntitySpec) GetNaturalLanguageDomain() *v0.NaturalLanguageDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_NaturalLanguageDomain); ok {
		return x.NaturalLanguageDomain
	}
	return nil
}

func (m *EntitySpec) GetImageDomain() *v0.ImageDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_ImageDomain); ok {
		return x.ImageDomain
	}
	return nil
}

func (m *EntitySpec) GetMidDomain() *v0.MIDDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_MidDomain); ok {
		return x.MidDomain
	}
	return nil
}

func (m *EntitySpec) GetUrlDomain() *v0.URLDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_UrlDomain); ok {
		return x.UrlDomain
	}
	return nil
}

func (m *EntitySpec) GetTimeDomain() *v0.TimeDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_TimeDomain); ok {
		return x.TimeDomain
	}
	return nil
}

func (m *EntitySpec) GetTimeOfDayDomain() *v0.TimeOfDayDomain {
	if x, ok := m.GetDomainInfo().(*EntitySpec_TimeOfDayDomain); ok {
		return x.TimeOfDayDomain
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EntitySpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EntitySpec_Presence)(nil),
		(*EntitySpec_GroupPresence)(nil),
		(*EntitySpec_Shape)(nil),
		(*EntitySpec_ValueCount)(nil),
		(*EntitySpec_Domain)(nil),
		(*EntitySpec_IntDomain)(nil),
		(*EntitySpec_FloatDomain)(nil),
		(*EntitySpec_StringDomain)(nil),
		(*EntitySpec_BoolDomain)(nil),
		(*EntitySpec_StructDomain)(nil),
		(*EntitySpec_NaturalLanguageDomain)(nil),
		(*EntitySpec_ImageDomain)(nil),
		(*EntitySpec_MidDomain)(nil),
		(*EntitySpec_UrlDomain)(nil),
		(*EntitySpec_TimeDomain)(nil),
		(*EntitySpec_TimeOfDayDomain)(nil),
	}
}

type FeatureSpec struct {
	// Name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value type of the feature.
	ValueType types.ValueType_Enum `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=feast.types.ValueType_Enum" json:"value_type,omitempty"`
	// Types that are valid to be assigned to PresenceConstraints:
	//	*FeatureSpec_Presence
	//	*FeatureSpec_GroupPresence
	PresenceConstraints isFeatureSpec_PresenceConstraints `protobuf_oneof:"presence_constraints"`
	// The shape of the feature which governs the number of values that appear in
	// each example.
	//
	// Types that are valid to be assigned to ShapeType:
	//	*FeatureSpec_Shape
	//	*FeatureSpec_ValueCount
	ShapeType isFeatureSpec_ShapeType `protobuf_oneof:"shape_type"`
	// Domain for the values of the feature.
	//
	// Types that are valid to be assigned to DomainInfo:
	//	*FeatureSpec_Domain
	//	*FeatureSpec_IntDomain
	//	*FeatureSpec_FloatDomain
	//	*FeatureSpec_StringDomain
	//	*FeatureSpec_BoolDomain
	//	*FeatureSpec_StructDomain
	//	*FeatureSpec_NaturalLanguageDomain
	//	*FeatureSpec_ImageDomain
	//	*FeatureSpec_MidDomain
	//	*FeatureSpec_UrlDomain
	//	*FeatureSpec_TimeDomain
	//	*FeatureSpec_TimeOfDayDomain
	DomainInfo           isFeatureSpec_DomainInfo `protobuf_oneof:"domain_info"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FeatureSpec) Reset()         { *m = FeatureSpec{} }
func (m *FeatureSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSpec) ProtoMessage()    {}
func (*FeatureSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_972fbd278ac19c0c, []int{3}
}

func (m *FeatureSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSpec.Unmarshal(m, b)
}
func (m *FeatureSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSpec.Marshal(b, m, deterministic)
}
func (m *FeatureSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSpec.Merge(m, src)
}
func (m *FeatureSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSpec.Size(m)
}
func (m *FeatureSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSpec proto.InternalMessageInfo

func (m *FeatureSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_INVALID
}

type isFeatureSpec_PresenceConstraints interface {
	isFeatureSpec_PresenceConstraints()
}

type FeatureSpec_Presence struct {
	Presence *v0.FeaturePresence `protobuf:"bytes,3,opt,name=presence,proto3,oneof"`
}

type FeatureSpec_GroupPresence struct {
	GroupPresence *v0.FeaturePresenceWithinGroup `protobuf:"bytes,4,opt,name=group_presence,json=groupPresence,proto3,oneof"`
}

func (*FeatureSpec_Presence) isFeatureSpec_PresenceConstraints() {}

func (*FeatureSpec_GroupPresence) isFeatureSpec_PresenceConstraints() {}

func (m *FeatureSpec) GetPresenceConstraints() isFeatureSpec_PresenceConstraints {
	if m != nil {
		return m.PresenceConstraints
	}
	return nil
}

func (m *FeatureSpec) GetPresence() *v0.FeaturePresence {
	if x, ok := m.GetPresenceConstraints().(*FeatureSpec_Presence); ok {
		return x.Presence
	}
	return nil
}

func (m *FeatureSpec) GetGroupPresence() *v0.FeaturePresenceWithinGroup {
	if x, ok := m.GetPresenceConstraints().(*FeatureSpec_GroupPresence); ok {
		return x.GroupPresence
	}
	return nil
}

type isFeatureSpec_ShapeType interface {
	isFeatureSpec_ShapeType()
}

type FeatureSpec_Shape struct {
	Shape *v0.FixedShape `protobuf:"bytes,5,opt,name=shape,proto3,oneof"`
}

type FeatureSpec_ValueCount struct {
	ValueCount *v0.ValueCount `protobuf:"bytes,6,opt,name=value_count,json=valueCount,proto3,oneof"`
}

func (*FeatureSpec_Shape) isFeatureSpec_ShapeType() {}

func (*FeatureSpec_ValueCount) isFeatureSpec_ShapeType() {}

func (m *FeatureSpec) GetShapeType() isFeatureSpec_ShapeType {
	if m != nil {
		return m.ShapeType
	}
	return nil
}

func (m *FeatureSpec) GetShape() *v0.FixedShape {
	if x, ok := m.GetShapeType().(*FeatureSpec_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *FeatureSpec) GetValueCount() *v0.ValueCount {
	if x, ok := m.GetShapeType().(*FeatureSpec_ValueCount); ok {
		return x.ValueCount
	}
	return nil
}

type isFeatureSpec_DomainInfo interface {
	isFeatureSpec_DomainInfo()
}

type FeatureSpec_Domain struct {
	Domain string `protobuf:"bytes,7,opt,name=domain,proto3,oneof"`
}

type FeatureSpec_IntDomain struct {
	IntDomain *v0.IntDomain `protobuf:"bytes,8,opt,name=int_domain,json=intDomain,proto3,oneof"`
}

type FeatureSpec_FloatDomain struct {
	FloatDomain *v0.FloatDomain `protobuf:"bytes,9,opt,name=float_domain,json=floatDomain,proto3,oneof"`
}

type FeatureSpec_StringDomain struct {
	StringDomain *v0.StringDomain `protobuf:"bytes,10,opt,name=string_domain,json=stringDomain,proto3,oneof"`
}

type FeatureSpec_BoolDomain struct {
	BoolDomain *v0.BoolDomain `protobuf:"bytes,11,opt,name=bool_domain,json=boolDomain,proto3,oneof"`
}

type FeatureSpec_StructDomain struct {
	StructDomain *v0.StructDomain `protobuf:"bytes,12,opt,name=struct_domain,json=structDomain,proto3,oneof"`
}

type FeatureSpec_NaturalLanguageDomain struct {
	NaturalLanguageDomain *v0.NaturalLanguageDomain `protobuf:"bytes,13,opt,name=natural_language_domain,json=naturalLanguageDomain,proto3,oneof"`
}

type FeatureSpec_ImageDomain struct {
	ImageDomain *v0.ImageDomain `protobuf:"bytes,14,opt,name=image_domain,json=imageDomain,proto3,oneof"`
}

type FeatureSpec_MidDomain struct {
	MidDomain *v0.MIDDomain `protobuf:"bytes,15,opt,name=mid_domain,json=midDomain,proto3,oneof"`
}

type FeatureSpec_UrlDomain struct {
	UrlDomain *v0.URLDomain `protobuf:"bytes,16,opt,name=url_domain,json=urlDomain,proto3,oneof"`
}

type FeatureSpec_TimeDomain struct {
	TimeDomain *v0.TimeDomain `protobuf:"bytes,17,opt,name=time_domain,json=timeDomain,proto3,oneof"`
}

type FeatureSpec_TimeOfDayDomain struct {
	TimeOfDayDomain *v0.TimeOfDayDomain `protobuf:"bytes,18,opt,name=time_of_day_domain,json=timeOfDayDomain,proto3,oneof"`
}

func (*FeatureSpec_Domain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_IntDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_FloatDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_StringDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_BoolDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_StructDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_NaturalLanguageDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_ImageDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_MidDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_UrlDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_TimeDomain) isFeatureSpec_DomainInfo() {}

func (*FeatureSpec_TimeOfDayDomain) isFeatureSpec_DomainInfo() {}

func (m *FeatureSpec) GetDomainInfo() isFeatureSpec_DomainInfo {
	if m != nil {
		return m.DomainInfo
	}
	return nil
}

func (m *FeatureSpec) GetDomain() string {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_Domain); ok {
		return x.Domain
	}
	return ""
}

func (m *FeatureSpec) GetIntDomain() *v0.IntDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_IntDomain); ok {
		return x.IntDomain
	}
	return nil
}

func (m *FeatureSpec) GetFloatDomain() *v0.FloatDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_FloatDomain); ok {
		return x.FloatDomain
	}
	return nil
}

func (m *FeatureSpec) GetStringDomain() *v0.StringDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_StringDomain); ok {
		return x.StringDomain
	}
	return nil
}

func (m *FeatureSpec) GetBoolDomain() *v0.BoolDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_BoolDomain); ok {
		return x.BoolDomain
	}
	return nil
}

func (m *FeatureSpec) GetStructDomain() *v0.StructDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_StructDomain); ok {
		return x.StructDomain
	}
	return nil
}

func (m *FeatureSpec) GetNaturalLanguageDomain() *v0.NaturalLanguageDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_NaturalLanguageDomain); ok {
		return x.NaturalLanguageDomain
	}
	return nil
}

func (m *FeatureSpec) GetImageDomain() *v0.ImageDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_ImageDomain); ok {
		return x.ImageDomain
	}
	return nil
}

func (m *FeatureSpec) GetMidDomain() *v0.MIDDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_MidDomain); ok {
		return x.MidDomain
	}
	return nil
}

func (m *FeatureSpec) GetUrlDomain() *v0.URLDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_UrlDomain); ok {
		return x.UrlDomain
	}
	return nil
}

func (m *FeatureSpec) GetTimeDomain() *v0.TimeDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_TimeDomain); ok {
		return x.TimeDomain
	}
	return nil
}

func (m *FeatureSpec) GetTimeOfDayDomain() *v0.TimeOfDayDomain {
	if x, ok := m.GetDomainInfo().(*FeatureSpec_TimeOfDayDomain); ok {
		return x.TimeOfDayDomain
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeatureSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeatureSpec_Presence)(nil),
		(*FeatureSpec_GroupPresence)(nil),
		(*FeatureSpec_Shape)(nil),
		(*FeatureSpec_ValueCount)(nil),
		(*FeatureSpec_Domain)(nil),
		(*FeatureSpec_IntDomain)(nil),
		(*FeatureSpec_FloatDomain)(nil),
		(*FeatureSpec_StringDomain)(nil),
		(*FeatureSpec_BoolDomain)(nil),
		(*FeatureSpec_StructDomain)(nil),
		(*FeatureSpec_NaturalLanguageDomain)(nil),
		(*FeatureSpec_ImageDomain)(nil),
		(*FeatureSpec_MidDomain)(nil),
		(*FeatureSpec_UrlDomain)(nil),
		(*FeatureSpec_TimeDomain)(nil),
		(*FeatureSpec_TimeOfDayDomain)(nil),
	}
}

type FeatureSetMeta struct {
	// Created timestamp of this specific feature set.
	CreatedTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp,omitempty"`
	// Status of the feature set.
	// Used to indicate whether the feature set is ready for consumption or ingestion.
	// Currently supports 2 states:
	// 1) STATUS_PENDING - A feature set is in pending state if Feast has not spun up the jobs
	// necessary to push rows for this feature set to stores subscribing to this feature set.
	// 2) STATUS_READY - Feature set is ready for consumption or ingestion
	Status               FeatureSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=feast.core.FeatureSetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FeatureSetMeta) Reset()         { *m = FeatureSetMeta{} }
func (m *FeatureSetMeta) String() string { return proto.CompactTextString(m) }
func (*FeatureSetMeta) ProtoMessage()    {}
func (*FeatureSetMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_972fbd278ac19c0c, []int{4}
}

func (m *FeatureSetMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSetMeta.Unmarshal(m, b)
}
func (m *FeatureSetMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSetMeta.Marshal(b, m, deterministic)
}
func (m *FeatureSetMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSetMeta.Merge(m, src)
}
func (m *FeatureSetMeta) XXX_Size() int {
	return xxx_messageInfo_FeatureSetMeta.Size(m)
}
func (m *FeatureSetMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSetMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSetMeta proto.InternalMessageInfo

func (m *FeatureSetMeta) GetCreatedTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTimestamp
	}
	return nil
}

func (m *FeatureSetMeta) GetStatus() FeatureSetStatus {
	if m != nil {
		return m.Status
	}
	return FeatureSetStatus_STATUS_INVALID
}

func init() {
	proto.RegisterEnum("feast.core.FeatureSetStatus", FeatureSetStatus_name, FeatureSetStatus_value)
	proto.RegisterType((*FeatureSet)(nil), "feast.core.FeatureSet")
	proto.RegisterType((*FeatureSetSpec)(nil), "feast.core.FeatureSetSpec")
	proto.RegisterType((*EntitySpec)(nil), "feast.core.EntitySpec")
	proto.RegisterType((*FeatureSpec)(nil), "feast.core.FeatureSpec")
	proto.RegisterType((*FeatureSetMeta)(nil), "feast.core.FeatureSetMeta")
}

func init() { proto.RegisterFile("feast/core/FeatureSet.proto", fileDescriptor_972fbd278ac19c0c) }

var fileDescriptor_972fbd278ac19c0c = []byte{
	// 938 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x97, 0xdf, 0x6e, 0xe2, 0xc6,
	0x17, 0xc7, 0x03, 0x49, 0x48, 0x38, 0x10, 0xc2, 0x8e, 0x7e, 0xbf, 0x8d, 0xbb, 0x5b, 0xb5, 0x29,
	0xad, 0xd4, 0x74, 0xa5, 0xda, 0x2b, 0xb6, 0x57, 0x7b, 0x17, 0x0a, 0x0d, 0xa8, 0x59, 0x1a, 0x19,
	0x36, 0x55, 0xdb, 0x0b, 0x6b, 0x30, 0x83, 0x33, 0xbb, 0xf6, 0x8c, 0xe5, 0x19, 0xd3, 0xf0, 0x14,
	0x7d, 0x80, 0xf6, 0xa6, 0x6f, 0x5a, 0xcd, 0xd8, 0x63, 0x93, 0x28, 0xd0, 0x3e, 0x00, 0x77, 0x39,
	0x3e, 0xdf, 0xf9, 0xf8, 0xfc, 0xcb, 0xc1, 0x03, 0x2f, 0x17, 0x04, 0x0b, 0xe9, 0xf8, 0x3c, 0x21,
	0xce, 0x0f, 0x04, 0xcb, 0x34, 0x21, 0x13, 0x22, 0xed, 0x38, 0xe1, 0x92, 0x23, 0xd0, 0x4e, 0x5b,
	0x39, 0x5f, 0x9c, 0x65, 0x42, 0xb9, 0x8a, 0x89, 0x70, 0x6e, 0x71, 0x98, 0x92, 0x4c, 0x64, 0x1c,
	0x9a, 0x30, 0xe1, 0x69, 0xe2, 0x1b, 0xc7, 0x67, 0x01, 0xe7, 0x41, 0x48, 0x1c, 0x6d, 0xcd, 0xd2,
	0x85, 0x33, 0x4f, 0x13, 0x2c, 0x29, 0x67, 0xb9, 0xff, 0xf3, 0xc7, 0x7e, 0x49, 0x23, 0x22, 0x24,
	0x8e, 0xe2, 0x5c, 0xf0, 0x8d, 0x24, 0x4c, 0xf0, 0x64, 0x11, 0xf2, 0xdf, 0xbd, 0x88, 0x48, 0x3c,
	0xc7, 0x12, 0x67, 0x6a, 0x67, 0xf9, 0xda, 0x11, 0xfe, 0x1d, 0x89, 0x70, 0x26, 0xed, 0x84, 0x00,
	0x65, 0xf4, 0xc8, 0x86, 0x03, 0x11, 0x13, 0xdf, 0xaa, 0x9c, 0x57, 0x2e, 0x1a, 0xdd, 0x17, 0x76,
	0x99, 0x86, 0x5d, 0xaa, 0x26, 0x31, 0xf1, 0x5d, 0xad, 0x53, 0x7a, 0xc5, 0xb7, 0xaa, 0xdb, 0xf4,
	0xef, 0x88, 0xc4, 0xae, 0xd6, 0x75, 0xfe, 0xae, 0x42, 0xeb, 0x21, 0x08, 0x59, 0x70, 0x14, 0x27,
	0xfc, 0x03, 0xf1, 0xa5, 0x75, 0x74, 0x5e, 0xb9, 0xa8, 0xbb, 0xc6, 0x44, 0x08, 0x0e, 0x18, 0x8e,
	0x88, 0x0e, 0xa6, 0xee, 0xea, 0xbf, 0x95, 0x7a, 0x49, 0x12, 0x41, 0x39, 0xd3, 0xef, 0x3c, 0x74,
	0x8d, 0x89, 0xba, 0x70, 0x4c, 0x98, 0xa4, 0x92, 0x12, 0x61, 0xed, 0x9f, 0xef, 0x5f, 0x34, 0xba,
	0xcf, 0xd7, 0xc3, 0x19, 0x28, 0xdf, 0x4a, 0x87, 0x5e, 0xe8, 0xd0, 0x1b, 0x38, 0x5e, 0x64, 0xd1,
	0x08, 0xeb, 0x40, 0x9f, 0x39, 0x7b, 0x2a, 0x05, 0x7d, 0xc8, 0x08, 0x51, 0x17, 0x8e, 0x22, 0x7c,
	0xef, 0xe1, 0x80, 0x58, 0x87, 0x3a, 0xed, 0x4f, 0xec, 0xac, 0x1f, 0xb6, 0xe9, 0x87, 0xdd, 0xcf,
	0xfb, 0xe5, 0xd6, 0x22, 0x7c, 0x7f, 0x19, 0x10, 0xf4, 0x0a, 0x6a, 0x42, 0x77, 0xd8, 0xaa, 0xe9,
	0x23, 0x68, 0xfd, 0x35, 0x59, 0xef, 0xdd, 0x5c, 0xd1, 0xf9, 0x13, 0x00, 0xca, 0x68, 0x9f, 0xac,
	0xc2, 0x5b, 0x80, 0xa5, 0x1a, 0x24, 0x4f, 0x0d, 0x95, 0x2e, 0x44, 0xab, 0xfb, 0x32, 0x47, 0xea,
	0x39, 0xb3, 0xf5, 0x9c, 0x4d, 0x57, 0xb1, 0x4a, 0x3c, 0x8d, 0xdc, 0xfa, 0xd2, 0xd8, 0x68, 0x00,
	0xc7, 0x71, 0x42, 0x04, 0x61, 0x3e, 0xb1, 0xf6, 0x75, 0x30, 0x5f, 0xdb, 0xe5, 0xb8, 0xd8, 0x66,
	0x5c, 0xec, 0xe5, 0x6b, 0x93, 0xff, 0x4d, 0x2e, 0x1f, 0xee, 0xb9, 0xc5, 0x51, 0xf4, 0x1b, 0xb4,
	0x82, 0x84, 0xa7, 0xb1, 0x57, 0xc0, 0x0e, 0x34, 0xac, 0xfb, 0x1f, 0x61, 0x3f, 0x53, 0x79, 0x47,
	0xd9, 0x95, 0x42, 0x0c, 0xf7, 0xdc, 0x13, 0xcd, 0x32, 0x3e, 0xf4, 0x16, 0x0e, 0xc5, 0x1d, 0x8e,
	0x4d, 0x81, 0x3b, 0x1b, 0x99, 0xf4, 0x9e, 0xcc, 0x27, 0x4a, 0x39, 0xac, 0xb8, 0xd9, 0x11, 0x34,
	0x80, 0x46, 0x56, 0x1b, 0x9f, 0xa7, 0x4c, 0xe6, 0xf5, 0xde, 0x48, 0xd0, 0x75, 0xfa, 0x5e, 0x29,
	0x87, 0x15, 0x37, 0x2b, 0xaa, 0xb6, 0x90, 0x05, 0xb5, 0x39, 0x8f, 0x30, 0x65, 0xd9, 0x54, 0x0e,
	0xab, 0x6e, 0x6e, 0xa3, 0x1e, 0x00, 0x65, 0xd2, 0xcb, 0xbd, 0xc7, 0x9a, 0xff, 0xc5, 0x26, 0xfe,
	0x88, 0xc9, 0xbe, 0x16, 0x0e, 0xab, 0x6e, 0x9d, 0x1a, 0x03, 0x0d, 0xa1, 0xb9, 0x08, 0x39, 0x2e,
	0x28, 0x75, 0x4d, 0xf9, 0x72, 0x63, 0x9e, 0x4a, 0x5b, 0x70, 0x1a, 0x8b, 0xd2, 0x44, 0x3f, 0xc2,
	0x89, 0x90, 0x09, 0x65, 0x81, 0x41, 0x81, 0x46, 0x7d, 0xb5, 0x09, 0x35, 0xd1, 0xe2, 0x82, 0xd5,
	0x14, 0x6b, 0xb6, 0xaa, 0xdd, 0x8c, 0xf3, 0xd0, 0xa0, 0x1a, 0xdb, 0x6b, 0xd7, 0xe3, 0x3c, 0x2c,
	0x40, 0x30, 0x2b, 0xac, 0x3c, 0xa6, 0xd4, 0x2f, 0xd2, 0x6b, 0xfe, 0x6b, 0x4c, 0xa9, 0x2f, 0x1f,
	0xc4, 0x54, 0xd8, 0x28, 0x80, 0x33, 0xa6, 0x26, 0x07, 0x87, 0x5e, 0x88, 0x59, 0x90, 0xe2, 0x80,
	0x18, 0xec, 0x89, 0xc6, 0x7e, 0xbb, 0x09, 0x3b, 0xce, 0x8e, 0x5d, 0xe7, 0xa7, 0x0a, 0xfe, 0xff,
	0xd9, 0x53, 0x0e, 0xd5, 0x13, 0x1a, 0xad, 0xd1, 0x5b, 0xdb, 0x7b, 0x32, 0x8a, 0xd6, 0x99, 0x0d,
	0x5a, 0x9a, 0x6a, 0x42, 0x22, 0x3a, 0x37, 0x9c, 0xd3, 0xed, 0x13, 0xf2, 0x6e, 0xd4, 0x2f, 0x27,
	0x24, 0xa2, 0xf3, 0x92, 0x91, 0x26, 0x45, 0x27, 0xda, 0xdb, 0x19, 0xef, 0xdd, 0xeb, 0x92, 0x91,
	0x26, 0x61, 0xd9, 0x4e, 0xf5, 0xcb, 0x60, 0x20, 0xcf, 0xb6, 0xb7, 0x73, 0x4a, 0xa3, 0x32, 0x1f,
	0x90, 0x85, 0x85, 0x6e, 0x01, 0x69, 0x0c, 0x5f, 0x78, 0x73, 0xbc, 0x32, 0x34, 0xb4, 0x7d, 0x77,
	0x28, 0xda, 0x4f, 0x8b, 0x3e, 0x5e, 0x15, 0xc8, 0x53, 0xf9, 0xf0, 0x51, 0xef, 0x39, 0xfc, 0xcf,
	0x2c, 0x0f, 0xcf, 0xe7, 0x4c, 0xc8, 0x04, 0x53, 0x26, 0x45, 0xaf, 0x09, 0xa0, 0xff, 0x95, 0xf5,
	0x76, 0xeb, 0x9d, 0x40, 0x23, 0x7b, 0xa3, 0x47, 0xd9, 0x82, 0x77, 0xfe, 0x02, 0x68, 0xac, 0xed,
	0xe5, 0xdd, 0x7a, 0xdc, 0xad, 0xc7, 0xdd, 0x7a, 0xdc, 0xad, 0xc7, 0xdd, 0x7a, 0xcc, 0xd6, 0xe3,
	0x1f, 0x95, 0xf5, 0x0f, 0x6c, 0xf5, 0xe5, 0x8d, 0xae, 0xe0, 0x99, 0x9f, 0x10, 0x2c, 0xc9, 0xdc,
	0x2b, 0xee, 0x09, 0xc5, 0x07, 0xfe, 0xe3, 0x2f, 0xd7, 0xa9, 0x51, 0xb8, 0xed, 0xfc, 0x50, 0xf1,
	0x04, 0x7d, 0x07, 0x35, 0x21, 0xb1, 0x4c, 0x45, 0xbe, 0x52, 0x3f, 0xdd, 0x70, 0x3d, 0xd0, 0x1a,
	0x37, 0xd7, 0xbe, 0xba, 0x86, 0xf6, 0x63, 0x1f, 0x42, 0xd0, 0x9a, 0x4c, 0x2f, 0xa7, 0xef, 0x27,
	0xde, 0x68, 0x7c, 0x7b, 0x79, 0x3d, 0xea, 0xb7, 0xf7, 0xd6, 0x9e, 0xdd, 0x0c, 0xc6, 0xfd, 0xd1,
	0xf8, 0xaa, 0x5d, 0x41, 0x6d, 0x68, 0xe6, 0xcf, 0xdc, 0xc1, 0x65, 0xff, 0x97, 0x76, 0xb5, 0x37,
	0x86, 0xb5, 0xab, 0x55, 0xef, 0xb4, 0x24, 0xdf, 0xa8, 0x0c, 0x7e, 0x75, 0x02, 0x2a, 0xef, 0xd2,
	0x99, 0xed, 0xf3, 0xc8, 0x09, 0xf8, 0x07, 0xf2, 0xd1, 0xc9, 0xee, 0x58, 0x62, 0xfe, 0xd1, 0x09,
	0x78, 0x76, 0x05, 0x12, 0x4e, 0x79, 0xef, 0x9a, 0xd5, 0xf4, 0xa3, 0x37, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x2a, 0x9e, 0xd9, 0xce, 0x0d, 0x00, 0x00,
}