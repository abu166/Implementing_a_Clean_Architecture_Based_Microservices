// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: order.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       string                 `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if x != nil {
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
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	TotalPrice    float32                `protobuf:"fixed32,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if x != nil {
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
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateOrderResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	TotalPrice    float32                `protobuf:"fixed32,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetOrderResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetOrderResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *GetOrderResponse) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusRequest) Reset() {
	*x = UpdateOrderStatusRequest{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusRequest) ProtoMessage() {}

func (x *UpdateOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOrderStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateOrderStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusResponse) Reset() {
	*x = UpdateOrderStatusResponse{}
	mi := &file_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResponse) ProtoMessage() {}

func (x *UpdateOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateOrderStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

type ListOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*GetOrderResponse    `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *ListOrdersResponse) GetOrders() []*GetOrderResponse {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{9}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	mi := &file_order_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{10}
}

func (x *CreateProductResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_order_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{11}
}

func (x *GetProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	mi := &file_order_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{12}
}

func (x *GetProductResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GetProductResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *GetProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_order_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductResponse) Reset() {
	*x = UpdateProductResponse{}
	mi := &file_order_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductResponse) ProtoMessage() {}

func (x *UpdateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{14}
}

func (x *UpdateProductResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateProductResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *UpdateProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_order_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{15}
}

func (x *DeleteProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductResponse) Reset() {
	*x = DeleteProductResponse{}
	mi := &file_order_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResponse) ProtoMessage() {}

func (x *DeleteProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{16}
}

func (x *DeleteProductResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      string                 `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"` // Optional filter by category
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`        // Pagination: Page number
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`      // Pagination: Items per page
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	mi := &file_order_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{17}
}

func (x *ListProductsRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ListProductsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*GetProductResponse  `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_order_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{18}
}

func (x *ListProductsResponse) GetProducts() []*GetProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\x05order\"W\n" +
	"\tOrderItem\x12\x18\n" +
	"\aproduct\x18\x01 \x01(\tR\aproduct\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\"U\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x05R\x06userId\x12&\n" +
	"\x05items\x18\x02 \x03(\v2\x10.order.OrderItemR\x05items\"^\n" +
	"\x13CreateOrderResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x12\x1f\n" +
	"\vtotal_price\x18\x03 \x01(\x02R\n" +
	"totalPrice\"!\n" +
	"\x0fGetOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"\x9c\x01\n" +
	"\x10GetOrderResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\x12\x1f\n" +
	"\vtotal_price\x18\x04 \x01(\x02R\n" +
	"totalPrice\x12&\n" +
	"\x05items\x18\x05 \x03(\v2\x10.order.OrderItemR\x05items\"B\n" +
	"\x18UpdateOrderStatusRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"5\n" +
	"\x19UpdateOrderStatusResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\x13\n" +
	"\x11ListOrdersRequest\"E\n" +
	"\x12ListOrdersResponse\x12/\n" +
	"\x06orders\x18\x01 \x03(\v2\x17.order.GetOrderResponseR\x06orders\"\x94\x01\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x02 \x01(\tR\bcategory\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x14\n" +
	"\x05stock\x18\x04 \x01(\x05R\x05stock\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\"\xa5\x01\n" +
	"\x15CreateProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x03 \x01(\tR\bcategory\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x02R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"\xa2\x01\n" +
	"\x12GetProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x03 \x01(\tR\bcategory\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x02R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\"\xa4\x01\n" +
	"\x14UpdateProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x03 \x01(\tR\bcategory\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x02R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\"\xa5\x01\n" +
	"\x15UpdateProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x03 \x01(\tR\bcategory\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x02R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\"&\n" +
	"\x14DeleteProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"1\n" +
	"\x15DeleteProductResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"[\n" +
	"\x13ListProductsRequest\x12\x1a\n" +
	"\bcategory\x18\x01 \x01(\tR\bcategory\x12\x12\n" +
	"\x04page\x18\x02 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x05R\x05limit\"M\n" +
	"\x14ListProductsResponse\x125\n" +
	"\bproducts\x18\x01 \x03(\v2\x19.order.GetProductResponseR\bproducts2\xac\x02\n" +
	"\fOrderService\x12D\n" +
	"\vCreateOrder\x12\x19.order.CreateOrderRequest\x1a\x1a.order.CreateOrderResponse\x12;\n" +
	"\bGetOrder\x12\x16.order.GetOrderRequest\x1a\x17.order.GetOrderResponse\x12V\n" +
	"\x11UpdateOrderStatus\x12\x1f.order.UpdateOrderStatusRequest\x1a .order.UpdateOrderStatusResponse\x12A\n" +
	"\n" +
	"ListOrders\x12\x18.order.ListOrdersRequest\x1a\x19.order.ListOrdersResponse2\x80\x03\n" +
	"\x0eProductService\x12J\n" +
	"\rCreateProduct\x12\x1b.order.CreateProductRequest\x1a\x1c.order.CreateProductResponse\x12A\n" +
	"\n" +
	"GetProduct\x12\x18.order.GetProductRequest\x1a\x19.order.GetProductResponse\x12J\n" +
	"\rUpdateProduct\x12\x1b.order.UpdateProductRequest\x1a\x1c.order.UpdateProductResponse\x12J\n" +
	"\rDeleteProduct\x12\x1b.order.DeleteProductRequest\x1a\x1c.order.DeleteProductResponse\x12G\n" +
	"\fListProducts\x12\x1a.order.ListProductsRequest\x1a\x1b.order.ListProductsResponseB\tZ\a./protob\x06proto3"

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData []byte
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)))
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_order_proto_goTypes = []any{
	(*OrderItem)(nil),                 // 0: order.OrderItem
	(*CreateOrderRequest)(nil),        // 1: order.CreateOrderRequest
	(*CreateOrderResponse)(nil),       // 2: order.CreateOrderResponse
	(*GetOrderRequest)(nil),           // 3: order.GetOrderRequest
	(*GetOrderResponse)(nil),          // 4: order.GetOrderResponse
	(*UpdateOrderStatusRequest)(nil),  // 5: order.UpdateOrderStatusRequest
	(*UpdateOrderStatusResponse)(nil), // 6: order.UpdateOrderStatusResponse
	(*ListOrdersRequest)(nil),         // 7: order.ListOrdersRequest
	(*ListOrdersResponse)(nil),        // 8: order.ListOrdersResponse
	(*CreateProductRequest)(nil),      // 9: order.CreateProductRequest
	(*CreateProductResponse)(nil),     // 10: order.CreateProductResponse
	(*GetProductRequest)(nil),         // 11: order.GetProductRequest
	(*GetProductResponse)(nil),        // 12: order.GetProductResponse
	(*UpdateProductRequest)(nil),      // 13: order.UpdateProductRequest
	(*UpdateProductResponse)(nil),     // 14: order.UpdateProductResponse
	(*DeleteProductRequest)(nil),      // 15: order.DeleteProductRequest
	(*DeleteProductResponse)(nil),     // 16: order.DeleteProductResponse
	(*ListProductsRequest)(nil),       // 17: order.ListProductsRequest
	(*ListProductsResponse)(nil),      // 18: order.ListProductsResponse
}
var file_order_proto_depIdxs = []int32{
	0,  // 0: order.CreateOrderRequest.items:type_name -> order.OrderItem
	0,  // 1: order.GetOrderResponse.items:type_name -> order.OrderItem
	4,  // 2: order.ListOrdersResponse.orders:type_name -> order.GetOrderResponse
	12, // 3: order.ListProductsResponse.products:type_name -> order.GetProductResponse
	1,  // 4: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	3,  // 5: order.OrderService.GetOrder:input_type -> order.GetOrderRequest
	5,  // 6: order.OrderService.UpdateOrderStatus:input_type -> order.UpdateOrderStatusRequest
	7,  // 7: order.OrderService.ListOrders:input_type -> order.ListOrdersRequest
	9,  // 8: order.ProductService.CreateProduct:input_type -> order.CreateProductRequest
	11, // 9: order.ProductService.GetProduct:input_type -> order.GetProductRequest
	13, // 10: order.ProductService.UpdateProduct:input_type -> order.UpdateProductRequest
	15, // 11: order.ProductService.DeleteProduct:input_type -> order.DeleteProductRequest
	17, // 12: order.ProductService.ListProducts:input_type -> order.ListProductsRequest
	2,  // 13: order.OrderService.CreateOrder:output_type -> order.CreateOrderResponse
	4,  // 14: order.OrderService.GetOrder:output_type -> order.GetOrderResponse
	6,  // 15: order.OrderService.UpdateOrderStatus:output_type -> order.UpdateOrderStatusResponse
	8,  // 16: order.OrderService.ListOrders:output_type -> order.ListOrdersResponse
	10, // 17: order.ProductService.CreateProduct:output_type -> order.CreateProductResponse
	12, // 18: order.ProductService.GetProduct:output_type -> order.GetProductResponse
	14, // 19: order.ProductService.UpdateProduct:output_type -> order.UpdateProductResponse
	16, // 20: order.ProductService.DeleteProduct:output_type -> order.DeleteProductResponse
	18, // 21: order.ProductService.ListProducts:output_type -> order.ListProductsResponse
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
