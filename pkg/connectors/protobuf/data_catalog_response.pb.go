// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_catalog_response.proto

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

type DataStore_DataStoreType int32

const (
	DataStore_UNKNOWN DataStore_DataStoreType = 0
	DataStore_LOCAL   DataStore_DataStoreType = 1
	DataStore_S3      DataStore_DataStoreType = 2
	DataStore_DB2     DataStore_DataStoreType = 3
	DataStore_KAFKA   DataStore_DataStoreType = 4
)

var DataStore_DataStoreType_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOCAL",
	2: "S3",
	3: "DB2",
	4: "KAFKA",
}

var DataStore_DataStoreType_value = map[string]int32{
	"UNKNOWN": 0,
	"LOCAL":   1,
	"S3":      2,
	"DB2":     3,
	"KAFKA":   4,
}

func (x DataStore_DataStoreType) String() string {
	return proto.EnumName(DataStore_DataStoreType_name, int32(x))
}

func (DataStore_DataStoreType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{2, 0}
}

type DataComponentMetadata struct {
	ComponentType string `protobuf:"bytes,1,opt,name=component_type,json=componentType,proto3" json:"component_type,omitempty"`
	//Named terms, that exist in Catalog toxonomy and the values for these terms
	//for columns we will have "SchemaDetails" key, that will include technical schema details for this column
	//TODO: Consider create special field for schema outside of metadata
	NamedMetadata map[string]string `protobuf:"bytes,2,rep,name=named_metadata,json=namedMetadata,proto3" json:"named_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//Tags - can be any free text added to a component (no taxonomy)
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataComponentMetadata) Reset()         { *m = DataComponentMetadata{} }
func (m *DataComponentMetadata) String() string { return proto.CompactTextString(m) }
func (*DataComponentMetadata) ProtoMessage()    {}
func (*DataComponentMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{0}
}

func (m *DataComponentMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataComponentMetadata.Unmarshal(m, b)
}
func (m *DataComponentMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataComponentMetadata.Marshal(b, m, deterministic)
}
func (m *DataComponentMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataComponentMetadata.Merge(m, src)
}
func (m *DataComponentMetadata) XXX_Size() int {
	return xxx_messageInfo_DataComponentMetadata.Size(m)
}
func (m *DataComponentMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DataComponentMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DataComponentMetadata proto.InternalMessageInfo

func (m *DataComponentMetadata) GetComponentType() string {
	if m != nil {
		return m.ComponentType
	}
	return ""
}

func (m *DataComponentMetadata) GetNamedMetadata() map[string]string {
	if m != nil {
		return m.NamedMetadata
	}
	return nil
}

func (m *DataComponentMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DatasetMetadata struct {
	DatasetNamedMetadata map[string]string `protobuf:"bytes,1,rep,name=dataset_named_metadata,json=datasetNamedMetadata,proto3" json:"dataset_named_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//Tags - can be any free text added to a component (no taxonomy)
	DatasetTags []string `protobuf:"bytes,2,rep,name=dataset_tags,json=datasetTags,proto3" json:"dataset_tags,omitempty"`
	//metadata for each component in asset. In tabular data each column is a component, then we will have: column name -> column metadata
	ComponentsMetadata   map[string]*DataComponentMetadata `protobuf:"bytes,3,rep,name=components_metadata,json=componentsMetadata,proto3" json:"components_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DatasetMetadata) Reset()         { *m = DatasetMetadata{} }
func (m *DatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*DatasetMetadata) ProtoMessage()    {}
func (*DatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{1}
}

func (m *DatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetMetadata.Unmarshal(m, b)
}
func (m *DatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetMetadata.Marshal(b, m, deterministic)
}
func (m *DatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetMetadata.Merge(m, src)
}
func (m *DatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_DatasetMetadata.Size(m)
}
func (m *DatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetMetadata proto.InternalMessageInfo

func (m *DatasetMetadata) GetDatasetNamedMetadata() map[string]string {
	if m != nil {
		return m.DatasetNamedMetadata
	}
	return nil
}

func (m *DatasetMetadata) GetDatasetTags() []string {
	if m != nil {
		return m.DatasetTags
	}
	return nil
}

func (m *DatasetMetadata) GetComponentsMetadata() map[string]*DataComponentMetadata {
	if m != nil {
		return m.ComponentsMetadata
	}
	return nil
}

type DataStore struct {
	Type DataStore_DataStoreType `protobuf:"varint,1,opt,name=type,proto3,enum=connectors.DataStore_DataStoreType" json:"type,omitempty"`
	Name string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now
	//should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2
	Db2                  *Db2DataStore   `protobuf:"bytes,3,opt,name=db2,proto3" json:"db2,omitempty"`
	S3                   *S3DataStore    `protobuf:"bytes,4,opt,name=s3,proto3" json:"s3,omitempty"`
	Kafka                *KafkaDataStore `protobuf:"bytes,5,opt,name=kafka,proto3" json:"kafka,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DataStore) Reset()         { *m = DataStore{} }
func (m *DataStore) String() string { return proto.CompactTextString(m) }
func (*DataStore) ProtoMessage()    {}
func (*DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{2}
}

func (m *DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStore.Unmarshal(m, b)
}
func (m *DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStore.Marshal(b, m, deterministic)
}
func (m *DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStore.Merge(m, src)
}
func (m *DataStore) XXX_Size() int {
	return xxx_messageInfo_DataStore.Size(m)
}
func (m *DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_DataStore proto.InternalMessageInfo

func (m *DataStore) GetType() DataStore_DataStoreType {
	if m != nil {
		return m.Type
	}
	return DataStore_UNKNOWN
}

func (m *DataStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataStore) GetDb2() *Db2DataStore {
	if m != nil {
		return m.Db2
	}
	return nil
}

func (m *DataStore) GetS3() *S3DataStore {
	if m != nil {
		return m.S3
	}
	return nil
}

func (m *DataStore) GetKafka() *KafkaDataStore {
	if m != nil {
		return m.Kafka
	}
	return nil
}

type DatasetDetails struct {
	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DataOwner  string     `protobuf:"bytes,2,opt,name=data_owner,json=dataOwner,proto3" json:"data_owner,omitempty"`
	DataStore  *DataStore `protobuf:"bytes,3,opt,name=data_store,json=dataStore,proto3" json:"data_store,omitempty"`
	DataFormat string     `protobuf:"bytes,4,opt,name=data_format,json=dataFormat,proto3" json:"data_format,omitempty"`
	Geo        string     `protobuf:"bytes,5,opt,name=geo,proto3" json:"geo,omitempty"`
	//LocationType locationType = 10;  //publicCloud/privateCloud etc. Should be filled later when we understand better if we have a closed set of values and how they are used.
	Metadata             *DatasetMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DatasetDetails) Reset()         { *m = DatasetDetails{} }
func (m *DatasetDetails) String() string { return proto.CompactTextString(m) }
func (*DatasetDetails) ProtoMessage()    {}
func (*DatasetDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{3}
}

func (m *DatasetDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetDetails.Unmarshal(m, b)
}
func (m *DatasetDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetDetails.Marshal(b, m, deterministic)
}
func (m *DatasetDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetDetails.Merge(m, src)
}
func (m *DatasetDetails) XXX_Size() int {
	return xxx_messageInfo_DatasetDetails.Size(m)
}
func (m *DatasetDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetDetails.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetDetails proto.InternalMessageInfo

func (m *DatasetDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasetDetails) GetDataOwner() string {
	if m != nil {
		return m.DataOwner
	}
	return ""
}

func (m *DatasetDetails) GetDataStore() *DataStore {
	if m != nil {
		return m.DataStore
	}
	return nil
}

func (m *DatasetDetails) GetDataFormat() string {
	if m != nil {
		return m.DataFormat
	}
	return ""
}

func (m *DatasetDetails) GetGeo() string {
	if m != nil {
		return m.Geo
	}
	return ""
}

func (m *DatasetDetails) GetMetadata() *DatasetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CatalogDatasetInfo struct {
	DatasetId            string          `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	Details              *DatasetDetails `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CatalogDatasetInfo) Reset()         { *m = CatalogDatasetInfo{} }
func (m *CatalogDatasetInfo) String() string { return proto.CompactTextString(m) }
func (*CatalogDatasetInfo) ProtoMessage()    {}
func (*CatalogDatasetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{4}
}

func (m *CatalogDatasetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogDatasetInfo.Unmarshal(m, b)
}
func (m *CatalogDatasetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogDatasetInfo.Marshal(b, m, deterministic)
}
func (m *CatalogDatasetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogDatasetInfo.Merge(m, src)
}
func (m *CatalogDatasetInfo) XXX_Size() int {
	return xxx_messageInfo_CatalogDatasetInfo.Size(m)
}
func (m *CatalogDatasetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogDatasetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogDatasetInfo proto.InternalMessageInfo

func (m *CatalogDatasetInfo) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *CatalogDatasetInfo) GetDetails() *DatasetDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type Db2DataStore struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Table                string   `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	Port                 string   `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Ssl                  string   `protobuf:"bytes,5,opt,name=ssl,proto3" json:"ssl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Db2DataStore) Reset()         { *m = Db2DataStore{} }
func (m *Db2DataStore) String() string { return proto.CompactTextString(m) }
func (*Db2DataStore) ProtoMessage()    {}
func (*Db2DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{5}
}

func (m *Db2DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Db2DataStore.Unmarshal(m, b)
}
func (m *Db2DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Db2DataStore.Marshal(b, m, deterministic)
}
func (m *Db2DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Db2DataStore.Merge(m, src)
}
func (m *Db2DataStore) XXX_Size() int {
	return xxx_messageInfo_Db2DataStore.Size(m)
}
func (m *Db2DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_Db2DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_Db2DataStore proto.InternalMessageInfo

func (m *Db2DataStore) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Db2DataStore) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Db2DataStore) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Db2DataStore) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Db2DataStore) GetSsl() string {
	if m != nil {
		return m.Ssl
	}
	return ""
}

type S3DataStore struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	ObjectKey            string   `protobuf:"bytes,3,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S3DataStore) Reset()         { *m = S3DataStore{} }
func (m *S3DataStore) String() string { return proto.CompactTextString(m) }
func (*S3DataStore) ProtoMessage()    {}
func (*S3DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{6}
}

func (m *S3DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3DataStore.Unmarshal(m, b)
}
func (m *S3DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3DataStore.Marshal(b, m, deterministic)
}
func (m *S3DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3DataStore.Merge(m, src)
}
func (m *S3DataStore) XXX_Size() int {
	return xxx_messageInfo_S3DataStore.Size(m)
}
func (m *S3DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_S3DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_S3DataStore proto.InternalMessageInfo

func (m *S3DataStore) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *S3DataStore) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *S3DataStore) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

func (m *S3DataStore) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type KafkaDataStore struct {
	TopicName             string   `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	BootstrapServers      string   `protobuf:"bytes,2,opt,name=bootstrap_servers,json=bootstrapServers,proto3" json:"bootstrap_servers,omitempty"`
	SchemaRegistry        string   `protobuf:"bytes,3,opt,name=schema_registry,json=schemaRegistry,proto3" json:"schema_registry,omitempty"`
	KeyDeserializer       string   `protobuf:"bytes,4,opt,name=key_deserializer,json=keyDeserializer,proto3" json:"key_deserializer,omitempty"`
	ValueDeserializer     string   `protobuf:"bytes,5,opt,name=value_deserializer,json=valueDeserializer,proto3" json:"value_deserializer,omitempty"`
	SecurityProtocol      string   `protobuf:"bytes,6,opt,name=security_protocol,json=securityProtocol,proto3" json:"security_protocol,omitempty"`
	SaslMechanism         string   `protobuf:"bytes,7,opt,name=sasl_mechanism,json=saslMechanism,proto3" json:"sasl_mechanism,omitempty"`
	SslTruststore         string   `protobuf:"bytes,8,opt,name=ssl_truststore,json=sslTruststore,proto3" json:"ssl_truststore,omitempty"`
	SslTruststorePassword string   `protobuf:"bytes,9,opt,name=ssl_truststore_password,json=sslTruststorePassword,proto3" json:"ssl_truststore_password,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *KafkaDataStore) Reset()         { *m = KafkaDataStore{} }
func (m *KafkaDataStore) String() string { return proto.CompactTextString(m) }
func (*KafkaDataStore) ProtoMessage()    {}
func (*KafkaDataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_995adc688b28b07a, []int{7}
}

func (m *KafkaDataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaDataStore.Unmarshal(m, b)
}
func (m *KafkaDataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaDataStore.Marshal(b, m, deterministic)
}
func (m *KafkaDataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaDataStore.Merge(m, src)
}
func (m *KafkaDataStore) XXX_Size() int {
	return xxx_messageInfo_KafkaDataStore.Size(m)
}
func (m *KafkaDataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaDataStore.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaDataStore proto.InternalMessageInfo

func (m *KafkaDataStore) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *KafkaDataStore) GetBootstrapServers() string {
	if m != nil {
		return m.BootstrapServers
	}
	return ""
}

func (m *KafkaDataStore) GetSchemaRegistry() string {
	if m != nil {
		return m.SchemaRegistry
	}
	return ""
}

func (m *KafkaDataStore) GetKeyDeserializer() string {
	if m != nil {
		return m.KeyDeserializer
	}
	return ""
}

func (m *KafkaDataStore) GetValueDeserializer() string {
	if m != nil {
		return m.ValueDeserializer
	}
	return ""
}

func (m *KafkaDataStore) GetSecurityProtocol() string {
	if m != nil {
		return m.SecurityProtocol
	}
	return ""
}

func (m *KafkaDataStore) GetSaslMechanism() string {
	if m != nil {
		return m.SaslMechanism
	}
	return ""
}

func (m *KafkaDataStore) GetSslTruststore() string {
	if m != nil {
		return m.SslTruststore
	}
	return ""
}

func (m *KafkaDataStore) GetSslTruststorePassword() string {
	if m != nil {
		return m.SslTruststorePassword
	}
	return ""
}

func init() {
	proto.RegisterEnum("connectors.DataStore_DataStoreType", DataStore_DataStoreType_name, DataStore_DataStoreType_value)
	proto.RegisterType((*DataComponentMetadata)(nil), "connectors.DataComponentMetadata")
	proto.RegisterMapType((map[string]string)(nil), "connectors.DataComponentMetadata.NamedMetadataEntry")
	proto.RegisterType((*DatasetMetadata)(nil), "connectors.DatasetMetadata")
	proto.RegisterMapType((map[string]*DataComponentMetadata)(nil), "connectors.DatasetMetadata.ComponentsMetadataEntry")
	proto.RegisterMapType((map[string]string)(nil), "connectors.DatasetMetadata.DatasetNamedMetadataEntry")
	proto.RegisterType((*DataStore)(nil), "connectors.DataStore")
	proto.RegisterType((*DatasetDetails)(nil), "connectors.DatasetDetails")
	proto.RegisterType((*CatalogDatasetInfo)(nil), "connectors.CatalogDatasetInfo")
	proto.RegisterType((*Db2DataStore)(nil), "connectors.Db2DataStore")
	proto.RegisterType((*S3DataStore)(nil), "connectors.S3DataStore")
	proto.RegisterType((*KafkaDataStore)(nil), "connectors.KafkaDataStore")
}

func init() {
	proto.RegisterFile("data_catalog_response.proto", fileDescriptor_995adc688b28b07a)
}

var fileDescriptor_995adc688b28b07a = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x6e, 0xdb, 0x46,
	0x10, 0xae, 0x48, 0xff, 0x71, 0x14, 0xc9, 0xf2, 0x36, 0x8e, 0x59, 0x05, 0x45, 0x1d, 0x15, 0x45,
	0xdc, 0x06, 0x15, 0x0a, 0xc9, 0xad, 0x8b, 0x3e, 0xd5, 0xb1, 0x9b, 0x22, 0x50, 0x62, 0x07, 0xb4,
	0x8b, 0x3e, 0xf4, 0x81, 0x58, 0x91, 0x6b, 0x89, 0x15, 0xc9, 0x25, 0x76, 0x56, 0x49, 0xd4, 0x63,
	0xf4, 0x0c, 0x3d, 0x45, 0xef, 0xd3, 0x03, 0xf4, 0x06, 0xc5, 0xfe, 0x90, 0xa2, 0x1c, 0xc5, 0x45,
	0xde, 0x76, 0xbe, 0xf9, 0x66, 0xe6, 0x9b, 0xd9, 0xe1, 0x12, 0x1e, 0xc6, 0x54, 0xd2, 0x30, 0xa2,
	0x92, 0xa6, 0x7c, 0x12, 0x0a, 0x86, 0x05, 0xcf, 0x91, 0xf5, 0x0b, 0xc1, 0x25, 0x27, 0x10, 0xf1,
	0x3c, 0x67, 0x91, 0xe4, 0x02, 0x7b, 0xff, 0x36, 0x60, 0xff, 0x9c, 0x4a, 0x7a, 0xc6, 0xb3, 0x82,
	0xe7, 0x2c, 0x97, 0x2f, 0x99, 0xa4, 0x2a, 0x98, 0x7c, 0x01, 0xed, 0xa8, 0x04, 0x43, 0xb9, 0x28,
	0x98, 0xdf, 0x38, 0x6c, 0x1c, 0x79, 0x41, 0xab, 0x42, 0xaf, 0x17, 0x05, 0x23, 0xbf, 0x41, 0x3b,
	0xa7, 0x19, 0x8b, 0xc3, 0xcc, 0x06, 0xfa, 0xce, 0xa1, 0x7b, 0xd4, 0x1c, 0x1c, 0xf7, 0x97, 0x55,
	0xfa, 0x6b, 0x2b, 0xf4, 0x2f, 0x54, 0x5c, 0x69, 0xfd, 0x94, 0x4b, 0xb1, 0x08, 0x5a, 0x79, 0x1d,
	0x23, 0x04, 0x36, 0x24, 0x9d, 0xa0, 0xef, 0x1e, 0xba, 0x47, 0x5e, 0xa0, 0xcf, 0xdd, 0x1f, 0x81,
	0xbc, 0x1b, 0x48, 0x3a, 0xe0, 0xce, 0xd8, 0xc2, 0x4a, 0x54, 0x47, 0x72, 0x1f, 0x36, 0x5f, 0xd3,
	0x74, 0xce, 0x7c, 0x47, 0x63, 0xc6, 0xf8, 0xc1, 0xf9, 0xbe, 0xd1, 0xfb, 0xdb, 0x85, 0x5d, 0xa5,
	0x08, 0xd9, 0xb2, 0xdb, 0x19, 0x3c, 0x88, 0x0d, 0x14, 0xde, 0x6a, 0xa7, 0xa1, 0xdb, 0xf9, 0xf6,
	0x76, 0x3b, 0xb5, 0xe0, 0xd2, 0x5e, 0xd3, 0xcf, 0xfd, 0x78, 0x8d, 0x8b, 0x3c, 0x82, 0x7b, 0x65,
	0x31, 0xdd, 0x9e, 0xa3, 0xdb, 0x6b, 0x5a, 0xec, 0x9a, 0x4e, 0x90, 0xc4, 0xf0, 0x71, 0x35, 0x67,
	0x5c, 0x8a, 0x71, 0xb5, 0x98, 0xe1, 0x5d, 0x62, 0xaa, 0x39, 0xe3, 0xaa, 0x14, 0x12, 0xbd, 0xe3,
	0xe8, 0xfe, 0x0c, 0x9f, 0xbc, 0x57, 0xfb, 0x87, 0x8c, 0xb4, 0x3b, 0x85, 0x83, 0xf7, 0xd4, 0x5d,
	0x93, 0xe6, 0xa4, 0x9e, 0xa6, 0x39, 0x78, 0xf4, 0xbf, 0x9b, 0x52, 0xbf, 0xbc, 0xbf, 0x1c, 0xf0,
	0x14, 0xe9, 0x4a, 0x72, 0xc1, 0xc8, 0x09, 0x6c, 0x54, 0xab, 0xd9, 0x1e, 0x7c, 0x7e, 0x3b, 0x93,
	0x26, 0x2d, 0x4f, 0x6a, 0x61, 0x03, 0x1d, 0xa0, 0x36, 0x4b, 0xdd, 0xb3, 0xed, 0x44, 0x9f, 0xc9,
	0x57, 0xe0, 0xc6, 0xe3, 0x81, 0xef, 0x6a, 0x55, 0xfe, 0x4a, 0xae, 0xf1, 0xa0, 0x4a, 0x12, 0x28,
	0x12, 0x79, 0x0c, 0x0e, 0x0e, 0xfd, 0x0d, 0x4d, 0x3d, 0xa8, 0x53, 0xaf, 0x86, 0x4b, 0xa6, 0x83,
	0x43, 0xf2, 0x0d, 0x6c, 0xce, 0xe8, 0xcd, 0x8c, 0xfa, 0x9b, 0x9a, 0xdb, 0xad, 0x73, 0x47, 0xca,
	0xb1, 0xa4, 0x1b, 0x62, 0xef, 0x0c, 0x5a, 0x2b, 0x8a, 0x49, 0x13, 0xb6, 0x7f, 0xb9, 0x18, 0x5d,
	0x5c, 0xfe, 0x7a, 0xd1, 0xf9, 0x88, 0x78, 0xb0, 0xf9, 0xe2, 0xf2, 0xec, 0xf4, 0x45, 0xa7, 0x41,
	0xb6, 0xc0, 0xb9, 0x1a, 0x76, 0x1c, 0xb2, 0x0d, 0xee, 0xf9, 0xd3, 0x41, 0xc7, 0x55, 0xbe, 0xd1,
	0xe9, 0xb3, 0xd1, 0x69, 0x67, 0xa3, 0xf7, 0x4f, 0x03, 0xda, 0xf6, 0x6a, 0xcf, 0x99, 0xa4, 0x49,
	0x8a, 0x55, 0xcb, 0x8d, 0x5a, 0xcb, 0x9f, 0x02, 0xe8, 0x97, 0x82, 0xbf, 0xc9, 0x99, 0xb0, 0xc3,
	0xf0, 0x14, 0x72, 0xa9, 0x00, 0x72, 0x6c, 0xdd, 0xa8, 0xb4, 0xd8, 0xc1, 0xec, 0xaf, 0x1d, 0xb2,
	0x89, 0x32, 0x97, 0xf2, 0x19, 0xe8, 0x55, 0x0e, 0x6f, 0xb8, 0xc8, 0xa8, 0xd4, 0x43, 0xf2, 0x02,
	0x9d, 0xe8, 0x99, 0x46, 0xd4, 0x4a, 0x4c, 0x18, 0xd7, 0x13, 0xf1, 0x02, 0x75, 0x24, 0x27, 0xb0,
	0x53, 0xed, 0xf8, 0xb6, 0x2e, 0xf3, 0xf0, 0x8e, 0x1d, 0x0f, 0x2a, 0x72, 0x2f, 0x01, 0x72, 0x66,
	0x5e, 0x39, 0xcb, 0x79, 0x9e, 0xdf, 0xf0, 0xb2, 0x2d, 0xf5, 0x81, 0x25, 0xb1, 0x6d, 0xd8, 0xb3,
	0xc8, 0xf3, 0x98, 0x1c, 0xc3, 0x76, 0x6c, 0x86, 0x62, 0x7b, 0xea, 0xae, 0x29, 0x66, 0xc7, 0x16,
	0x94, 0xd4, 0xde, 0x5b, 0xb8, 0x57, 0xdf, 0x03, 0xd5, 0xc5, 0x5c, 0xa4, 0xe5, 0x62, 0xcf, 0x45,
	0x4a, 0xba, 0xb0, 0xa3, 0x8a, 0x8c, 0x29, 0x96, 0x8b, 0x55, 0xd9, 0xea, 0xdb, 0x91, 0x74, 0x9c,
	0x9a, 0x29, 0x7a, 0x81, 0x31, 0xd4, 0x9d, 0x14, 0x5c, 0x94, 0x33, 0xd2, 0x67, 0x95, 0x17, 0x31,
	0x2d, 0xa7, 0x83, 0x98, 0xf6, 0xde, 0x42, 0xb3, 0xb6, 0x56, 0xaa, 0x0c, 0xcb, 0xe3, 0x82, 0x27,
	0xb9, 0xb4, 0xd5, 0x2b, 0x9b, 0x3c, 0x80, 0xad, 0xf1, 0x3c, 0x9a, 0x31, 0x69, 0x05, 0x58, 0x4b,
	0x4d, 0x84, 0x8f, 0x7f, 0x67, 0x91, 0x0c, 0xd5, 0xc7, 0x68, 0x34, 0x78, 0x06, 0x19, 0xb1, 0x85,
	0x0a, 0x13, 0x6c, 0x92, 0xf0, 0xdc, 0x2a, 0xb1, 0x56, 0xef, 0x4f, 0x17, 0xda, 0xab, 0x5b, 0xaa,
	0x32, 0x49, 0x5e, 0x24, 0x51, 0x58, 0x5b, 0x26, 0x4f, 0x23, 0xea, 0x0d, 0x21, 0x4f, 0x60, 0x6f,
	0xcc, 0xb9, 0x44, 0x29, 0x68, 0x11, 0x22, 0x13, 0xaf, 0x99, 0x40, 0xab, 0xa5, 0x53, 0x39, 0xae,
	0x0c, 0x4e, 0x1e, 0xc3, 0x2e, 0x46, 0x53, 0x96, 0xd1, 0x50, 0xd5, 0x43, 0x29, 0x4a, 0x69, 0x6d,
	0x03, 0x07, 0x16, 0x25, 0x5f, 0x42, 0x67, 0xc6, 0x16, 0x61, 0xcc, 0x90, 0x89, 0x84, 0xa6, 0xc9,
	0x1f, 0x4c, 0x58, 0xa5, 0xbb, 0x33, 0xb6, 0x38, 0xaf, 0xc1, 0xe4, 0x6b, 0x20, 0xfa, 0xb5, 0x58,
	0x25, 0x9b, 0x69, 0xee, 0x69, 0xcf, 0x0a, 0xfd, 0x09, 0xec, 0x21, 0x8b, 0xe6, 0x22, 0x91, 0x8b,
	0x50, 0xff, 0x1e, 0x23, 0x9e, 0xfa, 0x5b, 0x46, 0x6f, 0xe9, 0x78, 0x65, 0x71, 0xf5, 0x4f, 0x44,
	0x8a, 0x69, 0x98, 0xb1, 0x68, 0x4a, 0xf3, 0x04, 0x33, 0xbd, 0xac, 0x5e, 0xd0, 0x52, 0xe8, 0xcb,
	0x12, 0xd4, 0x34, 0x4c, 0x43, 0x29, 0xe6, 0x28, 0xcd, 0xa7, 0xb3, 0x63, 0x69, 0x98, 0x5e, 0x57,
	0x20, 0xf9, 0x0e, 0x0e, 0x56, 0x69, 0x61, 0x41, 0x11, 0xdf, 0x70, 0x11, 0xfb, 0x9e, 0xe6, 0xef,
	0xaf, 0xf0, 0x5f, 0x59, 0xe7, 0xd3, 0x16, 0x34, 0x23, 0x9e, 0xf5, 0x63, 0x2a, 0x33, 0x86, 0xd3,
	0xf1, 0x96, 0x96, 0x3d, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xc9, 0x1f, 0x9f, 0xf4, 0x07,
	0x00, 0x00,
}
