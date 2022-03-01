// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opts/ps_opts_manage_relay.proto

package opts

import (
	fmt "fmt"
	args "github.com/batchcorp/plumber-schemas/build/go/protos/args"
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

type GetRelayOptions struct {
	// @gotags: kong:"help='ID of the relay to get (leave empty to get all)'"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the relay to get (leave empty to get all)'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRelayOptions) Reset()         { *m = GetRelayOptions{} }
func (m *GetRelayOptions) String() string { return proto.CompactTextString(m) }
func (*GetRelayOptions) ProtoMessage()    {}
func (*GetRelayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86dbec1f9d2837a2, []int{0}
}

func (m *GetRelayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRelayOptions.Unmarshal(m, b)
}
func (m *GetRelayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRelayOptions.Marshal(b, m, deterministic)
}
func (m *GetRelayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRelayOptions.Merge(m, src)
}
func (m *GetRelayOptions) XXX_Size() int {
	return xxx_messageInfo_GetRelayOptions.Size(m)
}
func (m *GetRelayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRelayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GetRelayOptions proto.InternalMessageInfo

func (m *GetRelayOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateRelayOptions struct {
	// @gotags: kong:"help='ID of the connection to use',required"
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" kong:"help='ID of the connection to use',required"`
	// @gotags: kong:"help='Secret collection token',required"
	CollectionToken string `protobuf:"bytes,2,opt,name=collection_token,json=collectionToken,proto3" json:"collection_token,omitempty" kong:"help='Secret collection token',required"`
	// @gotags: kong:"help='How many messages to send in a single batch',default=1000"
	BatchSize int32 `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty" kong:"help='How many messages to send in a single batch',default=1000"`
	// @gotags: kong:"help='How many times plumber will try re-sending a batch',default=3"
	BatchMaxRetry int32 `protobuf:"varint,4,opt,name=batch_max_retry,json=batchMaxRetry,proto3" json:"batch_max_retry,omitempty" kong:"help='How many times plumber will try re-sending a batch',default=3"`
	// @gotags: kong:"help='How many workers to launch per relay',default=10"
	NumWorkers int32 `protobuf:"varint,5,opt,name=num_workers,json=numWorkers,proto3" json:"num_workers,omitempty" kong:"help='How many workers to launch per relay',default=10"`
	// @gotags: kong:"help='Alternative collector to relay events to',default='grpc-collector.batch.sh:9000'"
	BatchshGrpcAddress string `protobuf:"bytes,6,opt,name=batchsh_grpc_address,json=batchshGrpcAddress,proto3" json:"batchsh_grpc_address,omitempty" kong:"help='Alternative collector to relay events to',default='grpc-collector.batch.sh:9000'"`
	// @gotags: kong:"help='Whether to use TLS with collector'"
	BatchshGrpcDisableTls bool `protobuf:"varint,7,opt,name=batchsh_grpc_disable_tls,json=batchshGrpcDisableTls,proto3" json:"batchsh_grpc_disable_tls,omitempty" kong:"help='Whether to use TLS with collector'"`
	// @gotags: kong:"help='How long to wait before giving up talking to the gRPC collector',default=5"
	BatchshGrpcTimeoutSeconds int32 `protobuf:"varint,8,opt,name=batchsh_grpc_timeout_seconds,json=batchshGrpcTimeoutSeconds,proto3" json:"batchsh_grpc_timeout_seconds,omitempty" kong:"help='How long to wait before giving up talking to the gRPC collector',default=5"`
	// @gotags: kong:"cmd,help='Apache Kafka'"
	Kafka *args.KafkaRelayArgs `protobuf:"bytes,100,opt,name=kafka,proto3" json:"kafka,omitempty" kong:"cmd,help='Apache Kafka'"`
	// @gotags: kong:"cmd,help='AWS Simple Queue System'"
	AwsSqs *args.AWSSQSRelayArgs `protobuf:"bytes,101,opt,name=aws_sqs,json=awsSqs,proto3" json:"aws_sqs,omitempty" kong:"cmd,help='AWS Simple Queue System'"`
	// @gotags: kong:"cmd,help='MongoDB (CDC)'"
	Mongo *args.MongoReadArgs `protobuf:"bytes,102,opt,name=mongo,proto3" json:"mongo,omitempty" kong:"cmd,help='MongoDB (CDC)'"`
	// @gotags: kong:"cmd,help='NSQ'"
	Nsq *args.NSQReadArgs `protobuf:"bytes,103,opt,name=nsq,proto3" json:"nsq,omitempty" kong:"cmd,help='NSQ'"`
	// @gotags: kong:"cmd,help='MQTT'"
	Rabbit *args.RabbitReadArgs `protobuf:"bytes,104,opt,name=rabbit,proto3" json:"rabbit,omitempty" kong:"cmd,help='MQTT'"`
	// @gotags: kong:"cmd,help='MQTT'"
	Mqtt *args.MQTTReadArgs `protobuf:"bytes,105,opt,name=mqtt,proto3" json:"mqtt,omitempty" kong:"cmd,help='MQTT'"`
	// @gotags: kong:"cmd,help='Azure Service Bus'"
	AzureServiceBus *args.AzureServiceBusReadArgs `protobuf:"bytes,106,opt,name=azure_service_bus,json=azureServiceBus,proto3" json:"azure_service_bus,omitempty" kong:"cmd,help='Azure Service Bus'"`
	// @gotags: kong:"cmd,help='Google Cloud Platform Pub/Sub'"
	GcpPubsub *args.GCPPubSubReadArgs `protobuf:"bytes,107,opt,name=gcp_pubsub,json=gcpPubsub,proto3" json:"gcp_pubsub,omitempty" kong:"cmd,help='Google Cloud Platform Pub/Sub'"`
	// @gotags: kong:"cmd,help='KubeMQ Queue'"
	KubemqQueue *args.KubeMQQueueReadArgs `protobuf:"bytes,108,opt,name=kubemq_queue,json=kubemqQueue,proto3" json:"kubemq_queue,omitempty" kong:"cmd,help='KubeMQ Queue'"`
	// @gotags: kong:"cmd,help='Redis PubSub'"
	RedisPubsub *args.RedisPubSubReadArgs `protobuf:"bytes,109,opt,name=redis_pubsub,json=redisPubsub,proto3" json:"redis_pubsub,omitempty" kong:"cmd,help='Redis PubSub'"`
	// @gotags: kong:"cmd,help='Redis Streams'"
	RedisStreams *args.RedisStreamsReadArgs `protobuf:"bytes,111,opt,name=redis_streams,json=redisStreams,proto3" json:"redis_streams,omitempty" kong:"cmd,help='Redis Streams'"`
	// @gotags: kong:"cmd,help='PostgreSQL (CDC)'"
	Postgres *args.PostgresReadArgs `protobuf:"bytes,112,opt,name=postgres,proto3" json:"postgres,omitempty" kong:"cmd,help='PostgreSQL (CDC)'"`
	// @gotags: kong:"cmd,help='NATS'"
	Nats *args.NatsReadArgs `protobuf:"bytes,113,opt,name=nats,proto3" json:"nats,omitempty" kong:"cmd,help='NATS'"`
	// @gotags: kong:"cmd,help='NATS Streaming'"
	NatsStreaming *args.NatsStreamingReadArgs `protobuf:"bytes,114,opt,name=nats_streaming,json=natsStreaming,proto3" json:"nats_streaming,omitempty" kong:"cmd,help='NATS Streaming'"`
	// @gotags: kong:"cmd,help='NATS JetStream'"
	NatsJetstream        *args.NatsJetstreamReadArgs `protobuf:"bytes,115,opt,name=nats_jetstream,json=natsJetstream,proto3" json:"nats_jetstream,omitempty" kong:"cmd,help='NATS JetStream'"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CreateRelayOptions) Reset()         { *m = CreateRelayOptions{} }
func (m *CreateRelayOptions) String() string { return proto.CompactTextString(m) }
func (*CreateRelayOptions) ProtoMessage()    {}
func (*CreateRelayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86dbec1f9d2837a2, []int{1}
}

func (m *CreateRelayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRelayOptions.Unmarshal(m, b)
}
func (m *CreateRelayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRelayOptions.Marshal(b, m, deterministic)
}
func (m *CreateRelayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRelayOptions.Merge(m, src)
}
func (m *CreateRelayOptions) XXX_Size() int {
	return xxx_messageInfo_CreateRelayOptions.Size(m)
}
func (m *CreateRelayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRelayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRelayOptions proto.InternalMessageInfo

func (m *CreateRelayOptions) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *CreateRelayOptions) GetCollectionToken() string {
	if m != nil {
		return m.CollectionToken
	}
	return ""
}

func (m *CreateRelayOptions) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *CreateRelayOptions) GetBatchMaxRetry() int32 {
	if m != nil {
		return m.BatchMaxRetry
	}
	return 0
}

func (m *CreateRelayOptions) GetNumWorkers() int32 {
	if m != nil {
		return m.NumWorkers
	}
	return 0
}

func (m *CreateRelayOptions) GetBatchshGrpcAddress() string {
	if m != nil {
		return m.BatchshGrpcAddress
	}
	return ""
}

func (m *CreateRelayOptions) GetBatchshGrpcDisableTls() bool {
	if m != nil {
		return m.BatchshGrpcDisableTls
	}
	return false
}

func (m *CreateRelayOptions) GetBatchshGrpcTimeoutSeconds() int32 {
	if m != nil {
		return m.BatchshGrpcTimeoutSeconds
	}
	return 0
}

func (m *CreateRelayOptions) GetKafka() *args.KafkaRelayArgs {
	if m != nil {
		return m.Kafka
	}
	return nil
}

func (m *CreateRelayOptions) GetAwsSqs() *args.AWSSQSRelayArgs {
	if m != nil {
		return m.AwsSqs
	}
	return nil
}

func (m *CreateRelayOptions) GetMongo() *args.MongoReadArgs {
	if m != nil {
		return m.Mongo
	}
	return nil
}

func (m *CreateRelayOptions) GetNsq() *args.NSQReadArgs {
	if m != nil {
		return m.Nsq
	}
	return nil
}

func (m *CreateRelayOptions) GetRabbit() *args.RabbitReadArgs {
	if m != nil {
		return m.Rabbit
	}
	return nil
}

func (m *CreateRelayOptions) GetMqtt() *args.MQTTReadArgs {
	if m != nil {
		return m.Mqtt
	}
	return nil
}

func (m *CreateRelayOptions) GetAzureServiceBus() *args.AzureServiceBusReadArgs {
	if m != nil {
		return m.AzureServiceBus
	}
	return nil
}

func (m *CreateRelayOptions) GetGcpPubsub() *args.GCPPubSubReadArgs {
	if m != nil {
		return m.GcpPubsub
	}
	return nil
}

func (m *CreateRelayOptions) GetKubemqQueue() *args.KubeMQQueueReadArgs {
	if m != nil {
		return m.KubemqQueue
	}
	return nil
}

func (m *CreateRelayOptions) GetRedisPubsub() *args.RedisPubSubReadArgs {
	if m != nil {
		return m.RedisPubsub
	}
	return nil
}

func (m *CreateRelayOptions) GetRedisStreams() *args.RedisStreamsReadArgs {
	if m != nil {
		return m.RedisStreams
	}
	return nil
}

func (m *CreateRelayOptions) GetPostgres() *args.PostgresReadArgs {
	if m != nil {
		return m.Postgres
	}
	return nil
}

func (m *CreateRelayOptions) GetNats() *args.NatsReadArgs {
	if m != nil {
		return m.Nats
	}
	return nil
}

func (m *CreateRelayOptions) GetNatsStreaming() *args.NatsStreamingReadArgs {
	if m != nil {
		return m.NatsStreaming
	}
	return nil
}

func (m *CreateRelayOptions) GetNatsJetstream() *args.NatsJetstreamReadArgs {
	if m != nil {
		return m.NatsJetstream
	}
	return nil
}

type DeleteRelayOptions struct {
	// @gotags: kong:"help='ID of the relay to delete',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the relay to delete',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRelayOptions) Reset()         { *m = DeleteRelayOptions{} }
func (m *DeleteRelayOptions) String() string { return proto.CompactTextString(m) }
func (*DeleteRelayOptions) ProtoMessage()    {}
func (*DeleteRelayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86dbec1f9d2837a2, []int{2}
}

func (m *DeleteRelayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRelayOptions.Unmarshal(m, b)
}
func (m *DeleteRelayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRelayOptions.Marshal(b, m, deterministic)
}
func (m *DeleteRelayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRelayOptions.Merge(m, src)
}
func (m *DeleteRelayOptions) XXX_Size() int {
	return xxx_messageInfo_DeleteRelayOptions.Size(m)
}
func (m *DeleteRelayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRelayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRelayOptions proto.InternalMessageInfo

func (m *DeleteRelayOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StopRelayOptions struct {
	// @gotags: kong:"help='ID of the relay to stop',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the relay to stop',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRelayOptions) Reset()         { *m = StopRelayOptions{} }
func (m *StopRelayOptions) String() string { return proto.CompactTextString(m) }
func (*StopRelayOptions) ProtoMessage()    {}
func (*StopRelayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86dbec1f9d2837a2, []int{3}
}

func (m *StopRelayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRelayOptions.Unmarshal(m, b)
}
func (m *StopRelayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRelayOptions.Marshal(b, m, deterministic)
}
func (m *StopRelayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRelayOptions.Merge(m, src)
}
func (m *StopRelayOptions) XXX_Size() int {
	return xxx_messageInfo_StopRelayOptions.Size(m)
}
func (m *StopRelayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRelayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StopRelayOptions proto.InternalMessageInfo

func (m *StopRelayOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResumeRelayOptions struct {
	// @gotags: kong:"help='ID of the relay to resume',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the relay to resume',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeRelayOptions) Reset()         { *m = ResumeRelayOptions{} }
func (m *ResumeRelayOptions) String() string { return proto.CompactTextString(m) }
func (*ResumeRelayOptions) ProtoMessage()    {}
func (*ResumeRelayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86dbec1f9d2837a2, []int{4}
}

func (m *ResumeRelayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRelayOptions.Unmarshal(m, b)
}
func (m *ResumeRelayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRelayOptions.Marshal(b, m, deterministic)
}
func (m *ResumeRelayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRelayOptions.Merge(m, src)
}
func (m *ResumeRelayOptions) XXX_Size() int {
	return xxx_messageInfo_ResumeRelayOptions.Size(m)
}
func (m *ResumeRelayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRelayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRelayOptions proto.InternalMessageInfo

func (m *ResumeRelayOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRelayOptions)(nil), "protos.opts.GetRelayOptions")
	proto.RegisterType((*CreateRelayOptions)(nil), "protos.opts.CreateRelayOptions")
	proto.RegisterType((*DeleteRelayOptions)(nil), "protos.opts.DeleteRelayOptions")
	proto.RegisterType((*StopRelayOptions)(nil), "protos.opts.StopRelayOptions")
	proto.RegisterType((*ResumeRelayOptions)(nil), "protos.opts.ResumeRelayOptions")
}

func init() { proto.RegisterFile("opts/ps_opts_manage_relay.proto", fileDescriptor_86dbec1f9d2837a2) }

var fileDescriptor_86dbec1f9d2837a2 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xed, 0x6e, 0xdb, 0x36,
	0x18, 0x85, 0xe1, 0xb4, 0x49, 0x13, 0xba, 0x69, 0x3a, 0x62, 0x1f, 0x4c, 0xda, 0x2c, 0x8e, 0x17,
	0x0c, 0xde, 0x80, 0xda, 0xdb, 0x8a, 0x61, 0x28, 0x86, 0x61, 0x48, 0x53, 0x2c, 0xe8, 0x86, 0x74,
	0xb6, 0x64, 0xa0, 0xc0, 0xfe, 0x10, 0x94, 0xf4, 0x56, 0x56, 0x2d, 0x89, 0x12, 0x3f, 0x96, 0x36,
	0x37, 0xb0, 0xdb, 0x1e, 0x48, 0xca, 0xb2, 0x19, 0xbb, 0xf9, 0x25, 0xf8, 0x9c, 0xe7, 0x1c, 0xd1,
	0xa4, 0xf9, 0x1a, 0x9d, 0xf0, 0x4a, 0xc9, 0x51, 0x25, 0xa9, 0x79, 0xd2, 0x82, 0x95, 0x2c, 0x05,
	0x2a, 0x20, 0x67, 0x1f, 0x87, 0x95, 0xe0, 0x8a, 0xe3, 0xae, 0x7d, 0xc8, 0xa1, 0xf1, 0x8f, 0x8e,
	0x98, 0x48, 0x2d, 0x6d, 0x9e, 0x94, 0x5d, 0x4b, 0x2a, 0x6b, 0xe9, 0xc0, 0xa3, 0x33, 0xdf, 0xbb,
	0xd1, 0x02, 0xa8, 0x04, 0xf1, 0x6f, 0x16, 0x03, 0x8d, 0xf4, 0x82, 0x3a, 0xf6, 0xa8, 0x34, 0xae,
	0x68, 0xa5, 0x23, 0xa9, 0xa3, 0xc6, 0x26, 0x9e, 0x3d, 0x67, 0xef, 0xe6, 0xac, 0x71, 0x4e, 0x7c,
	0x47, 0x47, 0x50, 0xd4, 0xb4, 0xd6, 0xa0, 0x61, 0x63, 0xb4, 0xe0, 0x65, 0xca, 0x1b, 0xe7, 0x2b,
	0xdf, 0xa9, 0x95, 0xda, 0x68, 0x94, 0x4c, 0x2d, 0x56, 0x79, 0xba, 0x66, 0xd0, 0xf7, 0xa0, 0xa4,
	0x12, 0xc0, 0x8a, 0x4f, 0x23, 0xce, 0xcf, 0xca, 0xb4, 0x41, 0xbe, 0xf4, 0x11, 0x59, 0x37, 0xfa,
	0x13, 0x4f, 0xaf, 0xb8, 0x54, 0xa9, 0x80, 0xc5, 0xab, 0x0f, 0x3d, 0x53, 0xb0, 0x28, 0xca, 0xd4,
	0xc6, 0x2d, 0x10, 0x90, 0x64, 0xd2, 0xdf, 0xbd, 0xde, 0x06, 0xc0, 0x2d, 0xaa, 0x69, 0xef, 0x9f,
	0xa2, 0x83, 0x4b, 0x50, 0x81, 0x39, 0xdf, 0xbf, 0x2b, 0x95, 0xf1, 0x52, 0xe2, 0x47, 0x68, 0x2b,
	0x4b, 0x48, 0xa7, 0xd7, 0x19, 0xec, 0x05, 0x5b, 0x59, 0xd2, 0xff, 0x0f, 0x21, 0x7c, 0x21, 0x80,
	0x29, 0xf0, 0xb0, 0x6f, 0xd0, 0x7e, 0xcc, 0xcb, 0x12, 0x62, 0xf3, 0x91, 0xb6, 0x89, 0x87, 0x4b,
	0xf1, 0x75, 0x82, 0xbf, 0x43, 0x8f, 0x63, 0x9e, 0xe7, 0x0d, 0xa4, 0xf8, 0x1c, 0x4a, 0xb2, 0x65,
	0xb9, 0x83, 0xa5, 0x3e, 0x35, 0x32, 0x3e, 0x46, 0x28, 0x62, 0x2a, 0x9e, 0x51, 0x99, 0xdd, 0x00,
	0xb9, 0xd7, 0xeb, 0x0c, 0xb6, 0x83, 0x3d, 0xab, 0x84, 0xd9, 0x0d, 0xe0, 0x6f, 0xd1, 0x81, 0xb3,
	0x0b, 0xf6, 0x81, 0x0a, 0x50, 0xe2, 0x23, 0xb9, 0x6f, 0x99, 0x7d, 0x2b, 0x5f, 0xb1, 0x0f, 0x81,
	0x11, 0xf1, 0x09, 0xea, 0x96, 0xba, 0xa0, 0xd7, 0x5c, 0xcc, 0x41, 0x48, 0xb2, 0x6d, 0x19, 0x54,
	0xea, 0xe2, 0xad, 0x53, 0xf0, 0x0f, 0xe8, 0x73, 0x9b, 0x90, 0x33, 0x9a, 0x8a, 0x2a, 0xa6, 0x2c,
	0x49, 0x04, 0x48, 0x49, 0x76, 0xec, 0xb2, 0x70, 0xe3, 0x5d, 0x8a, 0x2a, 0x3e, 0x77, 0x0e, 0xfe,
	0x05, 0x11, 0x2f, 0x91, 0x64, 0x92, 0x45, 0x39, 0x50, 0x95, 0x4b, 0xf2, 0xa0, 0xd7, 0x19, 0xec,
	0x06, 0x5f, 0xac, 0xa4, 0x5e, 0x39, 0x77, 0x9a, 0x4b, 0xfc, 0x3b, 0x7a, 0xea, 0x05, 0x55, 0x56,
	0x00, 0xd7, 0x8a, 0x4a, 0x88, 0x79, 0x99, 0x48, 0xb2, 0x6b, 0x17, 0x77, 0xb8, 0x12, 0x9e, 0x3a,
	0x22, 0x74, 0x00, 0xfe, 0x11, 0x6d, 0xdb, 0x9f, 0x3c, 0x49, 0x7a, 0x9d, 0x41, 0xf7, 0xa7, 0x27,
	0xc3, 0xe6, 0xee, 0x99, 0xe3, 0x1c, 0xfe, 0x65, 0x1c, 0x7b, 0x24, 0xe7, 0x22, 0x95, 0x81, 0x23,
	0xf1, 0xcf, 0xe8, 0x41, 0x73, 0x0d, 0x09, 0xd8, 0xd0, 0x53, 0x2f, 0x74, 0xfe, 0x36, 0x0c, 0x27,
	0xe1, 0x32, 0xb5, 0xc3, 0xae, 0x65, 0x58, 0x9b, 0x5d, 0xd9, 0xb6, 0x37, 0x84, 0xbc, 0xb3, 0xa1,
	0x23, 0x2f, 0x74, 0x65, 0x9c, 0x00, 0x58, 0xe2, 0x5e, 0x64, 0x41, 0xfc, 0x3d, 0xba, 0x57, 0xca,
	0x9a, 0xa4, 0x96, 0x27, 0x1e, 0xff, 0x26, 0x9c, 0xb4, 0xb4, 0x81, 0xf0, 0x73, 0xb4, 0xe3, 0x7e,
	0xb8, 0x64, 0xb6, 0xe1, 0x8b, 0x04, 0xd6, 0x6a, 0x13, 0x0d, 0x8a, 0x9f, 0xa1, 0xfb, 0xe6, 0x6a,
	0x92, 0xcc, 0x46, 0x0e, 0xfd, 0x15, 0x4d, 0xa6, 0xd3, 0x36, 0x60, 0x31, 0x3c, 0x46, 0x9f, 0xad,
	0xcd, 0x18, 0xf2, 0xde, 0x66, 0xcf, 0xfc, 0x2d, 0x30, 0x54, 0xe8, 0xa0, 0x97, 0x5a, 0xb6, 0x35,
	0x07, 0xcc, 0x37, 0xf0, 0x6f, 0x08, 0x2d, 0xe7, 0x11, 0x99, 0xdb, 0xaa, 0xaf, 0xbd, 0xaa, 0xcb,
	0x8b, 0xf1, 0x58, 0x47, 0xa1, 0x8e, 0xda, 0x92, 0xbd, 0x34, 0xae, 0xc6, 0x36, 0x80, 0x2f, 0xd0,
	0xc3, 0xd5, 0xa9, 0x44, 0x72, 0x5b, 0xd0, 0xf3, 0xcf, 0x50, 0x47, 0x70, 0x35, 0x99, 0x18, 0xbf,
	0xad, 0xe8, 0xba, 0x94, 0x15, 0x4d, 0xc9, 0xea, 0xbd, 0x26, 0xc5, 0x86, 0x92, 0xc0, 0x00, 0xb7,
	0xd6, 0xd1, 0x15, 0x8d, 0x68, 0x56, 0xf2, 0x07, 0xda, 0xf7, 0xee, 0x3e, 0xe1, 0xb6, 0xe5, 0x74,
	0xbd, 0x25, 0x74, 0x40, 0x5b, 0xe3, 0x5e, 0xde, 0xa8, 0xf8, 0x05, 0xda, 0x5d, 0x0c, 0x27, 0x52,
	0xd9, 0x8a, 0x63, 0xaf, 0x62, 0xdc, 0x98, 0x6d, 0xbc, 0xc5, 0xcd, 0x61, 0x9a, 0x91, 0x48, 0xea,
	0x0d, 0x87, 0xf9, 0x86, 0xa9, 0x65, 0xc4, 0x62, 0xf8, 0x35, 0x7a, 0xe4, 0x4f, 0x50, 0x22, 0x6c,
	0xb0, 0xbf, 0x16, 0x0c, 0x17, 0x44, 0xdb, 0xb0, 0x5f, 0xae, 0xca, 0x6d, 0x55, 0x3b, 0xaf, 0x89,
	0xfc, 0x44, 0xd5, 0x9f, 0x0b, 0xc2, 0xaf, 0x6a, 0xe5, 0xfe, 0x19, 0xc2, 0xaf, 0x20, 0x87, 0x5b,
	0x83, 0xf0, 0xf6, 0xbc, 0xec, 0xa3, 0xc7, 0xa1, 0xe2, 0xd5, 0x9d, 0xcc, 0x19, 0xc2, 0x01, 0x48,
	0x5d, 0xdc, 0xd9, 0xf4, 0xf2, 0xd7, 0x7f, 0x5e, 0xa4, 0x99, 0x9a, 0xe9, 0x68, 0x18, 0xf3, 0x62,
	0x64, 0xc7, 0x44, 0xcc, 0x45, 0x35, 0xaa, 0x72, 0x5d, 0x44, 0x20, 0x9e, 0xc9, 0x78, 0x06, 0x05,
	0x93, 0xa3, 0x48, 0x67, 0x79, 0x32, 0x4a, 0xf9, 0xc8, 0x7d, 0xa3, 0x91, 0xf9, 0x6b, 0x8e, 0x76,
	0xec, 0x87, 0xe7, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x21, 0xd9, 0xc4, 0xe5, 0xd1, 0x07, 0x00,
	0x00,
}