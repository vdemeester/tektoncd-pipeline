// Copyright 2015 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file defines the GRPCLB LoadBalancing protocol.
//
// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/lb/v1/load_balancer.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.27.1
// source: grpc/lb/v1/load_balancer.proto

package grpc_lb_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoadBalanceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to LoadBalanceRequestType:
	//
	//	*LoadBalanceRequest_InitialRequest
	//	*LoadBalanceRequest_ClientStats
	LoadBalanceRequestType isLoadBalanceRequest_LoadBalanceRequestType `protobuf_oneof:"load_balance_request_type"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *LoadBalanceRequest) Reset() {
	*x = LoadBalanceRequest{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalanceRequest) ProtoMessage() {}

func (x *LoadBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalanceRequest.ProtoReflect.Descriptor instead.
func (*LoadBalanceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *LoadBalanceRequest) GetLoadBalanceRequestType() isLoadBalanceRequest_LoadBalanceRequestType {
	if x != nil {
		return x.LoadBalanceRequestType
	}
	return nil
}

func (x *LoadBalanceRequest) GetInitialRequest() *InitialLoadBalanceRequest {
	if x != nil {
		if x, ok := x.LoadBalanceRequestType.(*LoadBalanceRequest_InitialRequest); ok {
			return x.InitialRequest
		}
	}
	return nil
}

func (x *LoadBalanceRequest) GetClientStats() *ClientStats {
	if x != nil {
		if x, ok := x.LoadBalanceRequestType.(*LoadBalanceRequest_ClientStats); ok {
			return x.ClientStats
		}
	}
	return nil
}

type isLoadBalanceRequest_LoadBalanceRequestType interface {
	isLoadBalanceRequest_LoadBalanceRequestType()
}

type LoadBalanceRequest_InitialRequest struct {
	// This message should be sent on the first request to the load balancer.
	InitialRequest *InitialLoadBalanceRequest `protobuf:"bytes,1,opt,name=initial_request,json=initialRequest,proto3,oneof"`
}

type LoadBalanceRequest_ClientStats struct {
	// The client stats should be periodically reported to the load balancer
	// based on the duration defined in the InitialLoadBalanceResponse.
	ClientStats *ClientStats `protobuf:"bytes,2,opt,name=client_stats,json=clientStats,proto3,oneof"`
}

func (*LoadBalanceRequest_InitialRequest) isLoadBalanceRequest_LoadBalanceRequestType() {}

func (*LoadBalanceRequest_ClientStats) isLoadBalanceRequest_LoadBalanceRequestType() {}

type InitialLoadBalanceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the load balanced service (e.g., service.googleapis.com). Its
	// length should be less than 256 bytes.
	// The name might include a port number. How to handle the port number is up
	// to the balancer.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitialLoadBalanceRequest) Reset() {
	*x = InitialLoadBalanceRequest{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitialLoadBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialLoadBalanceRequest) ProtoMessage() {}

func (x *InitialLoadBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialLoadBalanceRequest.ProtoReflect.Descriptor instead.
func (*InitialLoadBalanceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{1}
}

func (x *InitialLoadBalanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Contains the number of calls finished for a particular load balance token.
type ClientStatsPerToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// See Server.load_balance_token.
	LoadBalanceToken string `protobuf:"bytes,1,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`
	// The total number of RPCs that finished associated with the token.
	NumCalls      int64 `protobuf:"varint,2,opt,name=num_calls,json=numCalls,proto3" json:"num_calls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientStatsPerToken) Reset() {
	*x = ClientStatsPerToken{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientStatsPerToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStatsPerToken) ProtoMessage() {}

func (x *ClientStatsPerToken) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStatsPerToken.ProtoReflect.Descriptor instead.
func (*ClientStatsPerToken) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{2}
}

func (x *ClientStatsPerToken) GetLoadBalanceToken() string {
	if x != nil {
		return x.LoadBalanceToken
	}
	return ""
}

func (x *ClientStatsPerToken) GetNumCalls() int64 {
	if x != nil {
		return x.NumCalls
	}
	return 0
}

// Contains client level statistics that are useful to load balancing. Each
// count except the timestamp should be reset to zero after reporting the stats.
type ClientStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The timestamp of generating the report.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The total number of RPCs that started.
	NumCallsStarted int64 `protobuf:"varint,2,opt,name=num_calls_started,json=numCallsStarted,proto3" json:"num_calls_started,omitempty"`
	// The total number of RPCs that finished.
	NumCallsFinished int64 `protobuf:"varint,3,opt,name=num_calls_finished,json=numCallsFinished,proto3" json:"num_calls_finished,omitempty"`
	// The total number of RPCs that failed to reach a server except dropped RPCs.
	NumCallsFinishedWithClientFailedToSend int64 `protobuf:"varint,6,opt,name=num_calls_finished_with_client_failed_to_send,json=numCallsFinishedWithClientFailedToSend,proto3" json:"num_calls_finished_with_client_failed_to_send,omitempty"`
	// The total number of RPCs that finished and are known to have been received
	// by a server.
	NumCallsFinishedKnownReceived int64 `protobuf:"varint,7,opt,name=num_calls_finished_known_received,json=numCallsFinishedKnownReceived,proto3" json:"num_calls_finished_known_received,omitempty"`
	// The list of dropped calls.
	CallsFinishedWithDrop []*ClientStatsPerToken `protobuf:"bytes,8,rep,name=calls_finished_with_drop,json=callsFinishedWithDrop,proto3" json:"calls_finished_with_drop,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ClientStats) Reset() {
	*x = ClientStats{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStats) ProtoMessage() {}

func (x *ClientStats) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStats.ProtoReflect.Descriptor instead.
func (*ClientStats) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{3}
}

func (x *ClientStats) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ClientStats) GetNumCallsStarted() int64 {
	if x != nil {
		return x.NumCallsStarted
	}
	return 0
}

func (x *ClientStats) GetNumCallsFinished() int64 {
	if x != nil {
		return x.NumCallsFinished
	}
	return 0
}

func (x *ClientStats) GetNumCallsFinishedWithClientFailedToSend() int64 {
	if x != nil {
		return x.NumCallsFinishedWithClientFailedToSend
	}
	return 0
}

func (x *ClientStats) GetNumCallsFinishedKnownReceived() int64 {
	if x != nil {
		return x.NumCallsFinishedKnownReceived
	}
	return 0
}

func (x *ClientStats) GetCallsFinishedWithDrop() []*ClientStatsPerToken {
	if x != nil {
		return x.CallsFinishedWithDrop
	}
	return nil
}

type LoadBalanceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to LoadBalanceResponseType:
	//
	//	*LoadBalanceResponse_InitialResponse
	//	*LoadBalanceResponse_ServerList
	//	*LoadBalanceResponse_FallbackResponse
	LoadBalanceResponseType isLoadBalanceResponse_LoadBalanceResponseType `protobuf_oneof:"load_balance_response_type"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *LoadBalanceResponse) Reset() {
	*x = LoadBalanceResponse{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalanceResponse) ProtoMessage() {}

func (x *LoadBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalanceResponse.ProtoReflect.Descriptor instead.
func (*LoadBalanceResponse) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{4}
}

func (x *LoadBalanceResponse) GetLoadBalanceResponseType() isLoadBalanceResponse_LoadBalanceResponseType {
	if x != nil {
		return x.LoadBalanceResponseType
	}
	return nil
}

func (x *LoadBalanceResponse) GetInitialResponse() *InitialLoadBalanceResponse {
	if x != nil {
		if x, ok := x.LoadBalanceResponseType.(*LoadBalanceResponse_InitialResponse); ok {
			return x.InitialResponse
		}
	}
	return nil
}

func (x *LoadBalanceResponse) GetServerList() *ServerList {
	if x != nil {
		if x, ok := x.LoadBalanceResponseType.(*LoadBalanceResponse_ServerList); ok {
			return x.ServerList
		}
	}
	return nil
}

func (x *LoadBalanceResponse) GetFallbackResponse() *FallbackResponse {
	if x != nil {
		if x, ok := x.LoadBalanceResponseType.(*LoadBalanceResponse_FallbackResponse); ok {
			return x.FallbackResponse
		}
	}
	return nil
}

type isLoadBalanceResponse_LoadBalanceResponseType interface {
	isLoadBalanceResponse_LoadBalanceResponseType()
}

type LoadBalanceResponse_InitialResponse struct {
	// This message should be sent on the first response to the client.
	InitialResponse *InitialLoadBalanceResponse `protobuf:"bytes,1,opt,name=initial_response,json=initialResponse,proto3,oneof"`
}

type LoadBalanceResponse_ServerList struct {
	// Contains the list of servers selected by the load balancer. The client
	// should send requests to these servers in the specified order.
	ServerList *ServerList `protobuf:"bytes,2,opt,name=server_list,json=serverList,proto3,oneof"`
}

type LoadBalanceResponse_FallbackResponse struct {
	// If this field is set, then the client should eagerly enter fallback
	// mode (even if there are existing, healthy connections to backends).
	FallbackResponse *FallbackResponse `protobuf:"bytes,3,opt,name=fallback_response,json=fallbackResponse,proto3,oneof"`
}

func (*LoadBalanceResponse_InitialResponse) isLoadBalanceResponse_LoadBalanceResponseType() {}

func (*LoadBalanceResponse_ServerList) isLoadBalanceResponse_LoadBalanceResponseType() {}

func (*LoadBalanceResponse_FallbackResponse) isLoadBalanceResponse_LoadBalanceResponseType() {}

type FallbackResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FallbackResponse) Reset() {
	*x = FallbackResponse{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FallbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FallbackResponse) ProtoMessage() {}

func (x *FallbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FallbackResponse.ProtoReflect.Descriptor instead.
func (*FallbackResponse) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{5}
}

type InitialLoadBalanceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This interval defines how often the client should send the client stats
	// to the load balancer. Stats should only be reported when the duration is
	// positive.
	ClientStatsReportInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=client_stats_report_interval,json=clientStatsReportInterval,proto3" json:"client_stats_report_interval,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *InitialLoadBalanceResponse) Reset() {
	*x = InitialLoadBalanceResponse{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitialLoadBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialLoadBalanceResponse) ProtoMessage() {}

func (x *InitialLoadBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialLoadBalanceResponse.ProtoReflect.Descriptor instead.
func (*InitialLoadBalanceResponse) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{6}
}

func (x *InitialLoadBalanceResponse) GetClientStatsReportInterval() *durationpb.Duration {
	if x != nil {
		return x.ClientStatsReportInterval
	}
	return nil
}

type ServerList struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Contains a list of servers selected by the load balancer. The list will
	// be updated when server resolutions change or as needed to balance load
	// across more servers. The client should consume the server list in order
	// unless instructed otherwise via the client_config.
	Servers       []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerList) Reset() {
	*x = ServerList{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerList) ProtoMessage() {}

func (x *ServerList) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerList.ProtoReflect.Descriptor instead.
func (*ServerList) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{7}
}

func (x *ServerList) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

// Contains server information. When the drop field is not true, use the other
// fields.
type Server struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A resolved address for the server, serialized in network-byte-order. It may
	// either be an IPv4 or IPv6 address.
	IpAddress []byte `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// A resolved port number for the server.
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// An opaque but printable token for load reporting. The client must include
	// the token of the picked server into the initial metadata when it starts a
	// call to that server. The token is used by the server to verify the request
	// and to allow the server to report load to the gRPC LB system. The token is
	// also used in client stats for reporting dropped calls.
	//
	// Its length can be variable but must be less than 50 bytes.
	LoadBalanceToken string `protobuf:"bytes,3,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`
	// Indicates whether this particular request should be dropped by the client.
	// If the request is dropped, there will be a corresponding entry in
	// ClientStats.calls_finished_with_drop.
	Drop          bool `protobuf:"varint,4,opt,name=drop,proto3" json:"drop,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_lb_v1_load_balancer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_grpc_lb_v1_load_balancer_proto_rawDescGZIP(), []int{8}
}

func (x *Server) GetIpAddress() []byte {
	if x != nil {
		return x.IpAddress
	}
	return nil
}

func (x *Server) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Server) GetLoadBalanceToken() string {
	if x != nil {
		return x.LoadBalanceToken
	}
	return ""
}

func (x *Server) GetDrop() bool {
	if x != nil {
		return x.Drop
	}
	return false
}

var File_grpc_lb_v1_load_balancer_proto protoreflect.FileDescriptor

var file_grpc_lb_v1_load_balancer_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01,
	0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x42, 0x1b, 0x0a, 0x19, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x2f, 0x0a, 0x19, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x60, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x50, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x22, 0xb0, 0x03, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a,
	0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x43, 0x61,
	0x6c, 0x6c, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75,
	0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x5d, 0x0a, 0x2d, 0x6e, 0x75, 0x6d, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x26, 0x6e, 0x75, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x54, 0x6f, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x48, 0x0a, 0x21, 0x6e, 0x75, 0x6d, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x1d, 0x6e, 0x75, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x12, 0x58, 0x0a, 0x18, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x15, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x44, 0x72, 0x6f, 0x70, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x90, 0x02, 0x0a, 0x13, 0x4c, 0x6f, 0x61, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f,
	0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x4b, 0x0a, 0x11, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x66, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x0a, 0x1a,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x46, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e,
	0x0a, 0x1a, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x1c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x40,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04,
	0x22, 0x83, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2c,
	0x0a, 0x12, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x72, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x72, 0x6f, 0x70,
	0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x32, 0x62, 0x0a, 0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x57, 0x0a, 0x0d, 0x69, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4c, 0x6f, 0x61,
	0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6c, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x62,
	0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_grpc_lb_v1_load_balancer_proto_rawDescOnce sync.Once
	file_grpc_lb_v1_load_balancer_proto_rawDescData []byte
)

func file_grpc_lb_v1_load_balancer_proto_rawDescGZIP() []byte {
	file_grpc_lb_v1_load_balancer_proto_rawDescOnce.Do(func() {
		file_grpc_lb_v1_load_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_grpc_lb_v1_load_balancer_proto_rawDesc), len(file_grpc_lb_v1_load_balancer_proto_rawDesc)))
	})
	return file_grpc_lb_v1_load_balancer_proto_rawDescData
}

var file_grpc_lb_v1_load_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_lb_v1_load_balancer_proto_goTypes = []any{
	(*LoadBalanceRequest)(nil),         // 0: grpc.lb.v1.LoadBalanceRequest
	(*InitialLoadBalanceRequest)(nil),  // 1: grpc.lb.v1.InitialLoadBalanceRequest
	(*ClientStatsPerToken)(nil),        // 2: grpc.lb.v1.ClientStatsPerToken
	(*ClientStats)(nil),                // 3: grpc.lb.v1.ClientStats
	(*LoadBalanceResponse)(nil),        // 4: grpc.lb.v1.LoadBalanceResponse
	(*FallbackResponse)(nil),           // 5: grpc.lb.v1.FallbackResponse
	(*InitialLoadBalanceResponse)(nil), // 6: grpc.lb.v1.InitialLoadBalanceResponse
	(*ServerList)(nil),                 // 7: grpc.lb.v1.ServerList
	(*Server)(nil),                     // 8: grpc.lb.v1.Server
	(*timestamppb.Timestamp)(nil),      // 9: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),        // 10: google.protobuf.Duration
}
var file_grpc_lb_v1_load_balancer_proto_depIdxs = []int32{
	1,  // 0: grpc.lb.v1.LoadBalanceRequest.initial_request:type_name -> grpc.lb.v1.InitialLoadBalanceRequest
	3,  // 1: grpc.lb.v1.LoadBalanceRequest.client_stats:type_name -> grpc.lb.v1.ClientStats
	9,  // 2: grpc.lb.v1.ClientStats.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 3: grpc.lb.v1.ClientStats.calls_finished_with_drop:type_name -> grpc.lb.v1.ClientStatsPerToken
	6,  // 4: grpc.lb.v1.LoadBalanceResponse.initial_response:type_name -> grpc.lb.v1.InitialLoadBalanceResponse
	7,  // 5: grpc.lb.v1.LoadBalanceResponse.server_list:type_name -> grpc.lb.v1.ServerList
	5,  // 6: grpc.lb.v1.LoadBalanceResponse.fallback_response:type_name -> grpc.lb.v1.FallbackResponse
	10, // 7: grpc.lb.v1.InitialLoadBalanceResponse.client_stats_report_interval:type_name -> google.protobuf.Duration
	8,  // 8: grpc.lb.v1.ServerList.servers:type_name -> grpc.lb.v1.Server
	0,  // 9: grpc.lb.v1.LoadBalancer.BalanceLoad:input_type -> grpc.lb.v1.LoadBalanceRequest
	4,  // 10: grpc.lb.v1.LoadBalancer.BalanceLoad:output_type -> grpc.lb.v1.LoadBalanceResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_grpc_lb_v1_load_balancer_proto_init() }
func file_grpc_lb_v1_load_balancer_proto_init() {
	if File_grpc_lb_v1_load_balancer_proto != nil {
		return
	}
	file_grpc_lb_v1_load_balancer_proto_msgTypes[0].OneofWrappers = []any{
		(*LoadBalanceRequest_InitialRequest)(nil),
		(*LoadBalanceRequest_ClientStats)(nil),
	}
	file_grpc_lb_v1_load_balancer_proto_msgTypes[4].OneofWrappers = []any{
		(*LoadBalanceResponse_InitialResponse)(nil),
		(*LoadBalanceResponse_ServerList)(nil),
		(*LoadBalanceResponse_FallbackResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_grpc_lb_v1_load_balancer_proto_rawDesc), len(file_grpc_lb_v1_load_balancer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_lb_v1_load_balancer_proto_goTypes,
		DependencyIndexes: file_grpc_lb_v1_load_balancer_proto_depIdxs,
		MessageInfos:      file_grpc_lb_v1_load_balancer_proto_msgTypes,
	}.Build()
	File_grpc_lb_v1_load_balancer_proto = out.File
	file_grpc_lb_v1_load_balancer_proto_goTypes = nil
	file_grpc_lb_v1_load_balancer_proto_depIdxs = nil
}
