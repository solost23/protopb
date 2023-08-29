// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: order_machine/order_machine.proto

package order_machine

import (
	common "github.com/solost23/protopb/gen/go/common"
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

type OrderSource int32

const (
	OrderSource_App          OrderSource = 0
	OrderSource_Web          OrderSource = 1
	OrderSource_SmallProgram OrderSource = 2
)

// Enum value maps for OrderSource.
var (
	OrderSource_name = map[int32]string{
		0: "App",
		1: "Web",
		2: "SmallProgram",
	}
	OrderSource_value = map[string]int32{
		"App":          0,
		"Web":          1,
		"SmallProgram": 2,
	}
)

func (x OrderSource) Enum() *OrderSource {
	p := new(OrderSource)
	*p = x
	return p
}

func (x OrderSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderSource) Descriptor() protoreflect.EnumDescriptor {
	return file_order_machine_order_machine_proto_enumTypes[0].Descriptor()
}

func (OrderSource) Type() protoreflect.EnumType {
	return &file_order_machine_order_machine_proto_enumTypes[0]
}

func (x OrderSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderSource.Descriptor instead.
func (OrderSource) EnumDescriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{0}
}

type OrderPayWay int32

const (
	OrderPayWay_AliPay     OrderPayWay = 0
	OrderPayWay_WeChat     OrderPayWay = 1
	OrderPayWay_OnlineBank OrderPayWay = 2
)

// Enum value maps for OrderPayWay.
var (
	OrderPayWay_name = map[int32]string{
		0: "AliPay",
		1: "WeChat",
		2: "OnlineBank",
	}
	OrderPayWay_value = map[string]int32{
		"AliPay":     0,
		"WeChat":     1,
		"OnlineBank": 2,
	}
)

func (x OrderPayWay) Enum() *OrderPayWay {
	p := new(OrderPayWay)
	*p = x
	return p
}

func (x OrderPayWay) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderPayWay) Descriptor() protoreflect.EnumDescriptor {
	return file_order_machine_order_machine_proto_enumTypes[1].Descriptor()
}

func (OrderPayWay) Type() protoreflect.EnumType {
	return &file_order_machine_order_machine_proto_enumTypes[1]
}

func (x OrderPayWay) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderPayWay.Descriptor instead.
func (OrderPayWay) EnumDescriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{1}
}

type OrderEvent int32

const (
	// 创建订单
	OrderEvent_EventCreate OrderEvent = 0
	// 确定订单
	OrderEvent_EventConfirm OrderEvent = 1
	// 修改订单
	OrderEvent_EventModify OrderEvent = 2
	// 支付订单
	OrderEvent_EventPay OrderEvent = 3
	// 发送订单
	OrderEvent_EventSend OrderEvent = 4
	// 用户收货
	OrderEvent_EventAccept OrderEvent = 5
	// 用户评价
	OrderEvent_EventEvaluation OrderEvent = 6
	// 用户申请退款
	OrderEvent_EventRefundAccept OrderEvent = 7
	// 用户取消退款
	OrderEvent_EventRefundCancel OrderEvent = 8
	// 后台同意退款
	OrderEvent_EventRefundConfirm OrderEvent = 9
	// 后台拒绝退款
	OrderEvent_EventRefundReject OrderEvent = 10
)

// Enum value maps for OrderEvent.
var (
	OrderEvent_name = map[int32]string{
		0:  "EventCreate",
		1:  "EventConfirm",
		2:  "EventModify",
		3:  "EventPay",
		4:  "EventSend",
		5:  "EventAccept",
		6:  "EventEvaluation",
		7:  "EventRefundAccept",
		8:  "EventRefundCancel",
		9:  "EventRefundConfirm",
		10: "EventRefundReject",
	}
	OrderEvent_value = map[string]int32{
		"EventCreate":        0,
		"EventConfirm":       1,
		"EventModify":        2,
		"EventPay":           3,
		"EventSend":          4,
		"EventAccept":        5,
		"EventEvaluation":    6,
		"EventRefundAccept":  7,
		"EventRefundCancel":  8,
		"EventRefundConfirm": 9,
		"EventRefundReject":  10,
	}
)

func (x OrderEvent) Enum() *OrderEvent {
	p := new(OrderEvent)
	*p = x
	return p
}

func (x OrderEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_order_machine_order_machine_proto_enumTypes[2].Descriptor()
}

func (OrderEvent) Type() protoreflect.EnumType {
	return &file_order_machine_order_machine_proto_enumTypes[2]
}

func (x OrderEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderEvent.Descriptor instead.
func (OrderEvent) EnumDescriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{2}
}

type OrderStatus int32

const (
	// 默认
	OrderStatus_StatusDefault OrderStatus = 0
	// 已预定
	OrderStatus_StatusReserved OrderStatus = 10
	// 待支付
	OrderStatus_StatusWaitPayment OrderStatus = 20
	// 已支付，待发货
	OrderStatus_StatusAlreadyPaymentWaitSendGoods OrderStatus = 30
	// 待收货
	OrderStatus_StatusWaitAcceptGoods OrderStatus = 40
	// 待评价
	OrderStatus_StatusWaitEvaluation OrderStatus = 50
	// 已评价，订单完成
	OrderStatus_StatusAlreadyEvaluationOrderOver OrderStatus = 60
	// 待退款，审核中
	OrderStatus_StatusWaitRefundChecking OrderStatus = 70
	// 退款成功
	OrderStatus_StatusRefundSuccess OrderStatus = 80
	// 退款失败
	OrderStatus_StatusRefundFailed OrderStatus = 90
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0:  "StatusDefault",
		10: "StatusReserved",
		20: "StatusWaitPayment",
		30: "StatusAlreadyPaymentWaitSendGoods",
		40: "StatusWaitAcceptGoods",
		50: "StatusWaitEvaluation",
		60: "StatusAlreadyEvaluationOrderOver",
		70: "StatusWaitRefundChecking",
		80: "StatusRefundSuccess",
		90: "StatusRefundFailed",
	}
	OrderStatus_value = map[string]int32{
		"StatusDefault":                     0,
		"StatusReserved":                    10,
		"StatusWaitPayment":                 20,
		"StatusAlreadyPaymentWaitSendGoods": 30,
		"StatusWaitAcceptGoods":             40,
		"StatusWaitEvaluation":              50,
		"StatusAlreadyEvaluationOrderOver":  60,
		"StatusWaitRefundChecking":          70,
		"StatusRefundSuccess":               80,
		"StatusRefundFailed":                90,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_order_machine_order_machine_proto_enumTypes[3].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_order_machine_order_machine_proto_enumTypes[3]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{3}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header      *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	CourseIds   []int32               `protobuf:"zigzag32,2,rep,packed,name=courseIds,proto3" json:"courseIds,omitempty"`
	CourseNums  []int32               `protobuf:"zigzag32,3,rep,packed,name=courseNums,proto3" json:"courseNums,omitempty"`
	OrderSource OrderSource           `protobuf:"varint,4,opt,name=orderSource,proto3,enum=order_machine.OrderSource" json:"orderSource,omitempty"`
	OrderPayWay OrderPayWay           `protobuf:"varint,5,opt,name=orderPayWay,proto3,enum=order_machine.OrderPayWay" json:"orderPayWay,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CreateOrderRequest) GetCourseIds() []int32 {
	if x != nil {
		return x.CourseIds
	}
	return nil
}

func (x *CreateOrderRequest) GetCourseNums() []int32 {
	if x != nil {
		return x.CourseNums
	}
	return nil
}

func (x *CreateOrderRequest) GetOrderSource() OrderSource {
	if x != nil {
		return x.OrderSource
	}
	return OrderSource_App
}

func (x *CreateOrderRequest) GetOrderPayWay() OrderPayWay {
	if x != nil {
		return x.OrderPayWay
	}
	return OrderPayWay_AliPay
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	OrderId   int32             `protobuf:"zigzag32,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *CreateOrderResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{2}
}

func (x *ListOrderRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type DetailOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     int32        `protobuf:"zigzag32,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	OrderStatus int32        `protobuf:"zigzag32,2,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	OrderEvents []OrderEvent `protobuf:"varint,3,rep,packed,name=orderEvents,proto3,enum=order_machine.OrderEvent" json:"orderEvents,omitempty"`
}

func (x *DetailOrder) Reset() {
	*x = DetailOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailOrder) ProtoMessage() {}

func (x *DetailOrder) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailOrder.ProtoReflect.Descriptor instead.
func (*DetailOrder) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{3}
}

func (x *DetailOrder) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DetailOrder) GetOrderStatus() int32 {
	if x != nil {
		return x.OrderStatus
	}
	return 0
}

func (x *DetailOrder) GetOrderEvents() []OrderEvent {
	if x != nil {
		return x.OrderEvents
	}
	return nil
}

type ListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	Records   []*DetailOrder    `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrderResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *ListOrderResponse) GetRecords() []*DetailOrder {
	if x != nil {
		return x.Records
	}
	return nil
}

type SwitchOrderStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header                *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	OrderId               int32                 `protobuf:"zigzag32,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	OrderEvent            OrderEvent            `protobuf:"varint,3,opt,name=orderEvent,proto3,enum=order_machine.OrderEvent" json:"orderEvent,omitempty"`
	RefundAcceptIntroduce string                `protobuf:"bytes,4,opt,name=refundAcceptIntroduce,proto3" json:"refundAcceptIntroduce,omitempty"`
	Evaluation            string                `protobuf:"bytes,5,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
}

func (x *SwitchOrderStateRequest) Reset() {
	*x = SwitchOrderStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchOrderStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchOrderStateRequest) ProtoMessage() {}

func (x *SwitchOrderStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchOrderStateRequest.ProtoReflect.Descriptor instead.
func (*SwitchOrderStateRequest) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{5}
}

func (x *SwitchOrderStateRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SwitchOrderStateRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SwitchOrderStateRequest) GetOrderEvent() OrderEvent {
	if x != nil {
		return x.OrderEvent
	}
	return OrderEvent_EventCreate
}

func (x *SwitchOrderStateRequest) GetRefundAcceptIntroduce() string {
	if x != nil {
		return x.RefundAcceptIntroduce
	}
	return ""
}

func (x *SwitchOrderStateRequest) GetEvaluation() string {
	if x != nil {
		return x.Evaluation
	}
	return ""
}

type SwitchOrderStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo   *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	OrderId     int32             `protobuf:"zigzag32,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	OrderStatus OrderStatus       `protobuf:"varint,3,opt,name=orderStatus,proto3,enum=order_machine.OrderStatus" json:"orderStatus,omitempty"`
}

func (x *SwitchOrderStateResponse) Reset() {
	*x = SwitchOrderStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_machine_order_machine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchOrderStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchOrderStateResponse) ProtoMessage() {}

func (x *SwitchOrderStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_machine_order_machine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchOrderStateResponse.ProtoReflect.Descriptor instead.
func (*SwitchOrderStateResponse) Descriptor() ([]byte, []int) {
	return file_order_machine_order_machine_proto_rawDescGZIP(), []int{6}
}

func (x *SwitchOrderStateResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *SwitchOrderStateResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SwitchOrderStateResponse) GetOrderStatus() OrderStatus {
	if x != nil {
		return x.OrderStatus
	}
	return OrderStatus_StatusDefault
}

var File_order_machine_order_machine_proto protoreflect.FileDescriptor

var file_order_machine_order_machine_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x11,
	0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x11, 0x52,
	0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x57, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x57, 0x61, 0x79, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x57, 0x61, 0x79, 0x22, 0x60, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a,
	0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0xf3, 0x01, 0x0a, 0x17, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x15, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x18, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x3c, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x31, 0x0a,
	0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x70, 0x70, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x65, 0x62, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x10, 0x02,
	0x2a, 0x35, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x57, 0x61, 0x79, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x50, 0x61, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x42, 0x61, 0x6e, 0x6b, 0x10, 0x02, 0x2a, 0xe0, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x6e, 0x64, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x06, 0x12, 0x15, 0x0a,
	0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x0a, 0x2a, 0x9c, 0x02, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x10,
	0x0a, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x57, 0x61, 0x69, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x14, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x57, 0x61, 0x69, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x10, 0x1e, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x57, 0x61, 0x69, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x10, 0x28, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x57, 0x61, 0x69, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x32, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6c,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x10, 0x3c, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x10, 0x46, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10,
	0x50, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x5a, 0x32, 0xa0, 0x02, 0x0a, 0x13, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x73,
	0x74, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_machine_order_machine_proto_rawDescOnce sync.Once
	file_order_machine_order_machine_proto_rawDescData = file_order_machine_order_machine_proto_rawDesc
)

func file_order_machine_order_machine_proto_rawDescGZIP() []byte {
	file_order_machine_order_machine_proto_rawDescOnce.Do(func() {
		file_order_machine_order_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_machine_order_machine_proto_rawDescData)
	})
	return file_order_machine_order_machine_proto_rawDescData
}

var file_order_machine_order_machine_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_order_machine_order_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_machine_order_machine_proto_goTypes = []interface{}{
	(OrderSource)(0),                 // 0: order_machine.OrderSource
	(OrderPayWay)(0),                 // 1: order_machine.OrderPayWay
	(OrderEvent)(0),                  // 2: order_machine.OrderEvent
	(OrderStatus)(0),                 // 3: order_machine.OrderStatus
	(*CreateOrderRequest)(nil),       // 4: order_machine.CreateOrderRequest
	(*CreateOrderResponse)(nil),      // 5: order_machine.CreateOrderResponse
	(*ListOrderRequest)(nil),         // 6: order_machine.ListOrderRequest
	(*DetailOrder)(nil),              // 7: order_machine.DetailOrder
	(*ListOrderResponse)(nil),        // 8: order_machine.ListOrderResponse
	(*SwitchOrderStateRequest)(nil),  // 9: order_machine.SwitchOrderStateRequest
	(*SwitchOrderStateResponse)(nil), // 10: order_machine.SwitchOrderStateResponse
	(*common.RequestHeader)(nil),     // 11: common.RequestHeader
	(*common.ErrorInfo)(nil),         // 12: common.ErrorInfo
}
var file_order_machine_order_machine_proto_depIdxs = []int32{
	11, // 0: order_machine.CreateOrderRequest.header:type_name -> common.RequestHeader
	0,  // 1: order_machine.CreateOrderRequest.orderSource:type_name -> order_machine.OrderSource
	1,  // 2: order_machine.CreateOrderRequest.orderPayWay:type_name -> order_machine.OrderPayWay
	12, // 3: order_machine.CreateOrderResponse.errorInfo:type_name -> common.ErrorInfo
	11, // 4: order_machine.ListOrderRequest.header:type_name -> common.RequestHeader
	2,  // 5: order_machine.DetailOrder.orderEvents:type_name -> order_machine.OrderEvent
	12, // 6: order_machine.ListOrderResponse.errorInfo:type_name -> common.ErrorInfo
	7,  // 7: order_machine.ListOrderResponse.records:type_name -> order_machine.DetailOrder
	11, // 8: order_machine.SwitchOrderStateRequest.header:type_name -> common.RequestHeader
	2,  // 9: order_machine.SwitchOrderStateRequest.orderEvent:type_name -> order_machine.OrderEvent
	12, // 10: order_machine.SwitchOrderStateResponse.errorInfo:type_name -> common.ErrorInfo
	3,  // 11: order_machine.SwitchOrderStateResponse.orderStatus:type_name -> order_machine.OrderStatus
	4,  // 12: order_machine.OrderMachineService.CreateOrder:input_type -> order_machine.CreateOrderRequest
	6,  // 13: order_machine.OrderMachineService.ListOrder:input_type -> order_machine.ListOrderRequest
	9,  // 14: order_machine.OrderMachineService.SwitchOrderState:input_type -> order_machine.SwitchOrderStateRequest
	5,  // 15: order_machine.OrderMachineService.CreateOrder:output_type -> order_machine.CreateOrderResponse
	8,  // 16: order_machine.OrderMachineService.ListOrder:output_type -> order_machine.ListOrderResponse
	10, // 17: order_machine.OrderMachineService.SwitchOrderState:output_type -> order_machine.SwitchOrderStateResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_order_machine_order_machine_proto_init() }
func file_order_machine_order_machine_proto_init() {
	if File_order_machine_order_machine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_machine_order_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_order_machine_order_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_order_machine_order_machine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderRequest); i {
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
		file_order_machine_order_machine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailOrder); i {
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
		file_order_machine_order_machine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderResponse); i {
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
		file_order_machine_order_machine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchOrderStateRequest); i {
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
		file_order_machine_order_machine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchOrderStateResponse); i {
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
			RawDescriptor: file_order_machine_order_machine_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_machine_order_machine_proto_goTypes,
		DependencyIndexes: file_order_machine_order_machine_proto_depIdxs,
		EnumInfos:         file_order_machine_order_machine_proto_enumTypes,
		MessageInfos:      file_order_machine_order_machine_proto_msgTypes,
	}.Build()
	File_order_machine_order_machine_proto = out.File
	file_order_machine_order_machine_proto_rawDesc = nil
	file_order_machine_order_machine_proto_goTypes = nil
	file_order_machine_order_machine_proto_depIdxs = nil
}