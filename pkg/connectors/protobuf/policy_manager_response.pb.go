// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy_manager_response.proto

package connectors

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EnforcementAction_EnforcementActionLevel int32

const (
	EnforcementAction_UNKNOWN EnforcementAction_EnforcementActionLevel = 0
	EnforcementAction_DATASET EnforcementAction_EnforcementActionLevel = 1
	EnforcementAction_COLUMN  EnforcementAction_EnforcementActionLevel = 2
	EnforcementAction_ROW     EnforcementAction_EnforcementActionLevel = 3
	EnforcementAction_CELL    EnforcementAction_EnforcementActionLevel = 4
)

var EnforcementAction_EnforcementActionLevel_name = map[int32]string{
	0: "UNKNOWN",
	1: "DATASET",
	2: "COLUMN",
	3: "ROW",
	4: "CELL",
}

var EnforcementAction_EnforcementActionLevel_value = map[string]int32{
	"UNKNOWN": 0,
	"DATASET": 1,
	"COLUMN":  2,
	"ROW":     3,
	"CELL":    4,
}

func (x EnforcementAction_EnforcementActionLevel) String() string {
	return proto.EnumName(EnforcementAction_EnforcementActionLevel_name, int32(x))
}

func (EnforcementAction_EnforcementActionLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{0, 0}
}

type EnforcementAction struct {
	Name                 string                                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Level                EnforcementAction_EnforcementActionLevel `protobuf:"varint,3,opt,name=level,proto3,enum=connectors.EnforcementAction_EnforcementActionLevel" json:"level,omitempty"`
	Args                 map[string]string                        `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *EnforcementAction) Reset()         { *m = EnforcementAction{} }
func (m *EnforcementAction) String() string { return proto.CompactTextString(m) }
func (*EnforcementAction) ProtoMessage()    {}
func (*EnforcementAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{0}
}

func (m *EnforcementAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnforcementAction.Unmarshal(m, b)
}
func (m *EnforcementAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnforcementAction.Marshal(b, m, deterministic)
}
func (m *EnforcementAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnforcementAction.Merge(m, src)
}
func (m *EnforcementAction) XXX_Size() int {
	return xxx_messageInfo_EnforcementAction.Size(m)
}
func (m *EnforcementAction) XXX_DiscardUnknown() {
	xxx_messageInfo_EnforcementAction.DiscardUnknown(m)
}

var xxx_messageInfo_EnforcementAction proto.InternalMessageInfo

func (m *EnforcementAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnforcementAction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EnforcementAction) GetLevel() EnforcementAction_EnforcementActionLevel {
	if m != nil {
		return m.Level
	}
	return EnforcementAction_UNKNOWN
}

func (m *EnforcementAction) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type OperationDecision struct {
	Operation            *AccessOperation     `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	EnforcementActions   []*EnforcementAction `protobuf:"bytes,2,rep,name=enforcement_actions,json=enforcementActions,proto3" json:"enforcement_actions,omitempty"`
	UsedPolicies         []*Policy            `protobuf:"bytes,3,rep,name=used_policies,json=usedPolicies,proto3" json:"used_policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OperationDecision) Reset()         { *m = OperationDecision{} }
func (m *OperationDecision) String() string { return proto.CompactTextString(m) }
func (*OperationDecision) ProtoMessage()    {}
func (*OperationDecision) Descriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{1}
}

func (m *OperationDecision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationDecision.Unmarshal(m, b)
}
func (m *OperationDecision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationDecision.Marshal(b, m, deterministic)
}
func (m *OperationDecision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationDecision.Merge(m, src)
}
func (m *OperationDecision) XXX_Size() int {
	return xxx_messageInfo_OperationDecision.Size(m)
}
func (m *OperationDecision) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationDecision.DiscardUnknown(m)
}

var xxx_messageInfo_OperationDecision proto.InternalMessageInfo

func (m *OperationDecision) GetOperation() *AccessOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *OperationDecision) GetEnforcementActions() []*EnforcementAction {
	if m != nil {
		return m.EnforcementActions
	}
	return nil
}

func (m *OperationDecision) GetUsedPolicies() []*Policy {
	if m != nil {
		return m.UsedPolicies
	}
	return nil
}

type DatasetDecision struct {
	Dataset              *DatasetIdentifier   `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	Decisions            []*OperationDecision `protobuf:"bytes,2,rep,name=decisions,proto3" json:"decisions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DatasetDecision) Reset()         { *m = DatasetDecision{} }
func (m *DatasetDecision) String() string { return proto.CompactTextString(m) }
func (*DatasetDecision) ProtoMessage()    {}
func (*DatasetDecision) Descriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{2}
}

func (m *DatasetDecision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetDecision.Unmarshal(m, b)
}
func (m *DatasetDecision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetDecision.Marshal(b, m, deterministic)
}
func (m *DatasetDecision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetDecision.Merge(m, src)
}
func (m *DatasetDecision) XXX_Size() int {
	return xxx_messageInfo_DatasetDecision.Size(m)
}
func (m *DatasetDecision) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetDecision.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetDecision proto.InternalMessageInfo

func (m *DatasetDecision) GetDataset() *DatasetIdentifier {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func (m *DatasetDecision) GetDecisions() []*OperationDecision {
	if m != nil {
		return m.Decisions
	}
	return nil
}

type Policy struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Hierarchy            []string `protobuf:"bytes,5,rep,name=hierarchy,proto3" json:"hierarchy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{3}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Policy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Policy) GetHierarchy() []string {
	if m != nil {
		return m.Hierarchy
	}
	return nil
}

type ComponentVersion struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentVersion) Reset()         { *m = ComponentVersion{} }
func (m *ComponentVersion) String() string { return proto.CompactTextString(m) }
func (*ComponentVersion) ProtoMessage()    {}
func (*ComponentVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{4}
}

func (m *ComponentVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentVersion.Unmarshal(m, b)
}
func (m *ComponentVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentVersion.Marshal(b, m, deterministic)
}
func (m *ComponentVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentVersion.Merge(m, src)
}
func (m *ComponentVersion) XXX_Size() int {
	return xxx_messageInfo_ComponentVersion.Size(m)
}
func (m *ComponentVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentVersion proto.InternalMessageInfo

func (m *ComponentVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ComponentVersion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ComponentVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PoliciesDecisions struct {
	ComponentVersions    []*ComponentVersion  `protobuf:"bytes,1,rep,name=component_versions,json=componentVersions,proto3" json:"component_versions,omitempty"`
	DatasetDecisions     []*DatasetDecision   `protobuf:"bytes,2,rep,name=dataset_decisions,json=datasetDecisions,proto3" json:"dataset_decisions,omitempty"`
	GeneralDecisions     []*OperationDecision `protobuf:"bytes,3,rep,name=general_decisions,json=generalDecisions,proto3" json:"general_decisions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PoliciesDecisions) Reset()         { *m = PoliciesDecisions{} }
func (m *PoliciesDecisions) String() string { return proto.CompactTextString(m) }
func (*PoliciesDecisions) ProtoMessage()    {}
func (*PoliciesDecisions) Descriptor() ([]byte, []int) {
	return fileDescriptor_136c41ae692a5fdb, []int{5}
}

func (m *PoliciesDecisions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoliciesDecisions.Unmarshal(m, b)
}
func (m *PoliciesDecisions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoliciesDecisions.Marshal(b, m, deterministic)
}
func (m *PoliciesDecisions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoliciesDecisions.Merge(m, src)
}
func (m *PoliciesDecisions) XXX_Size() int {
	return xxx_messageInfo_PoliciesDecisions.Size(m)
}
func (m *PoliciesDecisions) XXX_DiscardUnknown() {
	xxx_messageInfo_PoliciesDecisions.DiscardUnknown(m)
}

var xxx_messageInfo_PoliciesDecisions proto.InternalMessageInfo

func (m *PoliciesDecisions) GetComponentVersions() []*ComponentVersion {
	if m != nil {
		return m.ComponentVersions
	}
	return nil
}

func (m *PoliciesDecisions) GetDatasetDecisions() []*DatasetDecision {
	if m != nil {
		return m.DatasetDecisions
	}
	return nil
}

func (m *PoliciesDecisions) GetGeneralDecisions() []*OperationDecision {
	if m != nil {
		return m.GeneralDecisions
	}
	return nil
}

func init() {
	proto.RegisterEnum("connectors.EnforcementAction_EnforcementActionLevel", EnforcementAction_EnforcementActionLevel_name, EnforcementAction_EnforcementActionLevel_value)
	proto.RegisterType((*EnforcementAction)(nil), "connectors.EnforcementAction")
	proto.RegisterMapType((map[string]string)(nil), "connectors.EnforcementAction.ArgsEntry")
	proto.RegisterType((*OperationDecision)(nil), "connectors.OperationDecision")
	proto.RegisterType((*DatasetDecision)(nil), "connectors.DatasetDecision")
	proto.RegisterType((*Policy)(nil), "connectors.Policy")
	proto.RegisterType((*ComponentVersion)(nil), "connectors.ComponentVersion")
	proto.RegisterType((*PoliciesDecisions)(nil), "connectors.PoliciesDecisions")
}

func init() {
	proto.RegisterFile("policy_manager_response.proto", fileDescriptor_136c41ae692a5fdb)
}

var fileDescriptor_136c41ae692a5fdb = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xd9, 0x4e, 0x9b, 0x5f, 0x26, 0xb4, 0x38, 0x0b, 0x42, 0x56, 0xff, 0x48, 0x91, 0x2f,
	0xe4, 0x94, 0x43, 0x41, 0x2a, 0xd0, 0x53, 0x68, 0x23, 0x41, 0x1b, 0x92, 0x62, 0x5a, 0x7a, 0xb4,
	0x96, 0xf5, 0x34, 0xb5, 0x88, 0x77, 0xcd, 0xee, 0xb6, 0x52, 0x6e, 0xdc, 0xf8, 0x72, 0xdc, 0xf9,
	0x26, 0x9c, 0x91, 0xd7, 0x7f, 0xe2, 0x26, 0x51, 0xd5, 0xdb, 0xce, 0xcb, 0x9b, 0xb7, 0x6f, 0xde,
	0x4e, 0x0c, 0xfb, 0xa9, 0x98, 0xc5, 0x6c, 0x1e, 0x26, 0x94, 0xd3, 0x29, 0xca, 0x50, 0xa2, 0x4a,
	0x05, 0x57, 0xd8, 0x4f, 0xa5, 0xd0, 0x82, 0x00, 0x13, 0x9c, 0x23, 0xd3, 0x42, 0xaa, 0x9d, 0xbd,
	0x15, 0xea, 0x8f, 0x5b, 0x54, 0x3a, 0x67, 0xfa, 0xbf, 0x6d, 0xe8, 0x0c, 0xf9, 0xb5, 0x90, 0x0c,
	0x13, 0xe4, 0x7a, 0xc0, 0x74, 0x2c, 0x38, 0x21, 0xd0, 0xe0, 0x34, 0x41, 0xcf, 0xea, 0x5a, 0xbd,
	0x56, 0x60, 0xce, 0x64, 0x1b, 0xec, 0x38, 0xf2, 0x6c, 0x83, 0xd8, 0x71, 0x44, 0x4e, 0x61, 0x63,
	0x86, 0x77, 0x38, 0xf3, 0x9c, 0xae, 0xd5, 0xdb, 0x3e, 0x78, 0xdd, 0x5f, 0xdc, 0xd9, 0x5f, 0x51,
	0x5c, 0x45, 0x46, 0x59, 0x6f, 0x90, 0x4b, 0x90, 0x23, 0x68, 0x50, 0x39, 0x55, 0x5e, 0xa3, 0xeb,
	0xf4, 0xda, 0x07, 0x2f, 0x1f, 0x96, 0x1a, 0xc8, 0xa9, 0x1a, 0x72, 0x2d, 0xe7, 0x81, 0x69, 0xda,
	0x39, 0x84, 0x56, 0x05, 0x11, 0x17, 0x9c, 0xef, 0x38, 0x2f, 0x8c, 0x67, 0x47, 0xf2, 0x1c, 0x36,
	0xee, 0xe8, 0xec, 0x16, 0x0b, 0xeb, 0x79, 0xf1, 0xce, 0x7e, 0x63, 0xf9, 0x9f, 0xe1, 0xc5, 0x7a,
	0x5b, 0xa4, 0x0d, 0xcd, 0xcb, 0xf1, 0xd9, 0x78, 0x72, 0x35, 0x76, 0xff, 0xcb, 0x8a, 0x93, 0xc1,
	0xc5, 0xe0, 0xcb, 0xf0, 0xc2, 0xb5, 0x08, 0xc0, 0xe6, 0xf1, 0x64, 0x74, 0xf9, 0x69, 0xec, 0xda,
	0xa4, 0x09, 0x4e, 0x30, 0xb9, 0x72, 0x1d, 0xf2, 0x3f, 0x34, 0x8e, 0x87, 0xa3, 0x91, 0xdb, 0xf0,
	0xff, 0x58, 0xd0, 0x99, 0xa4, 0x28, 0x69, 0xa6, 0x75, 0x82, 0x2c, 0x56, 0x59, 0x9c, 0x6f, 0xa1,
	0x25, 0x4a, 0xd0, 0x58, 0x6b, 0x1f, 0xec, 0xd6, 0x67, 0x1c, 0x30, 0x86, 0x4a, 0x55, 0x7d, 0xc1,
	0x82, 0x4d, 0xc6, 0xf0, 0x0c, 0x17, 0x1e, 0x43, 0x6a, 0x4c, 0x2a, 0xcf, 0x36, 0x41, 0xed, 0x3f,
	0x18, 0x54, 0x40, 0x70, 0x19, 0x52, 0xe4, 0x10, 0xb6, 0x6e, 0x15, 0x46, 0xa1, 0x59, 0x8a, 0x18,
	0x95, 0xe7, 0x18, 0x25, 0x52, 0x57, 0x3a, 0x37, 0x0b, 0x13, 0x3c, 0xc9, 0x88, 0xe7, 0x05, 0xcf,
	0xff, 0x65, 0xc1, 0xd3, 0x13, 0xaa, 0xa9, 0x42, 0x5d, 0xcd, 0x75, 0x08, 0xcd, 0x28, 0x87, 0x8a,
	0xa9, 0xee, 0x19, 0x2a, 0xd8, 0x1f, 0x23, 0xe4, 0x3a, 0xbe, 0x8e, 0x51, 0x06, 0x25, 0x9b, 0x1c,
	0x41, 0x2b, 0x2a, 0x44, 0xd6, 0xce, 0xb2, 0x12, 0x61, 0xb0, 0xe0, 0xfb, 0x3f, 0x2d, 0xd8, 0xcc,
	0x2d, 0x16, 0x3b, 0x69, 0x55, 0x3b, 0x59, 0xee, 0xad, 0x5d, 0xdb, 0xdb, 0x2e, 0xb4, 0x23, 0x54,
	0x4c, 0xc6, 0xa9, 0x89, 0xdf, 0x31, 0x3f, 0xd5, 0xa1, 0xac, 0x4b, 0xcf, 0x53, 0xf4, 0x1a, 0x79,
	0x57, 0x76, 0x26, 0x7b, 0xd0, 0xba, 0x89, 0x51, 0x52, 0xc9, 0x6e, 0xe6, 0xde, 0x46, 0xd7, 0xe9,
	0xb5, 0x82, 0x05, 0xe0, 0x9f, 0x83, 0x7b, 0x2c, 0x92, 0x54, 0x70, 0xe4, 0xfa, 0x2b, 0x4a, 0xf5,
	0xd8, 0xff, 0x8c, 0x07, 0xcd, 0xbb, 0x9c, 0x5e, 0xf8, 0x28, 0x4b, 0xff, 0xaf, 0x05, 0x9d, 0x32,
	0xeb, 0x72, 0x68, 0x45, 0xce, 0x80, 0xb0, 0xf2, 0x9e, 0xb0, 0xa0, 0x2a, 0xcf, 0x32, 0x81, 0xed,
	0xd5, 0x03, 0x5b, 0x76, 0x13, 0x74, 0xd8, 0x12, 0xa2, 0xc8, 0x07, 0xe8, 0x14, 0xf9, 0x87, 0xcb,
	0xe1, 0xef, 0xae, 0x79, 0xb7, 0x2a, 0x7a, 0x37, 0xba, 0x0f, 0x28, 0x72, 0x0a, 0x9d, 0x29, 0x72,
	0x94, 0x74, 0x56, 0x53, 0x72, 0x1e, 0xf3, 0x8c, 0x6e, 0xd1, 0x57, 0x69, 0xbd, 0xdf, 0x82, 0x36,
	0x13, 0x49, 0x3f, 0xa2, 0x3a, 0x41, 0x75, 0xf3, 0x6d, 0xd3, 0x7c, 0x96, 0x5e, 0xfd, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x72, 0x50, 0xb9, 0xe1, 0x04, 0x00, 0x00,
}
