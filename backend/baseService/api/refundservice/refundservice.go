// Code generated by Kitex v0.12.3. DO NOT EDIT.

package refundservice

import (
	api "DouTok-example/backend/baseService/kitex_gen/api"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"RefundCreate": kitex.NewMethodInfo(
		refundCreateHandler,
		newRefundCreateArgs,
		newRefundCreateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RefundQuery": kitex.NewMethodInfo(
		refundQueryHandler,
		newRefundQueryArgs,
		newRefundQueryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RefundUpdate": kitex.NewMethodInfo(
		refundUpdateHandler,
		newRefundUpdateArgs,
		newRefundUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	refundServiceServiceInfo                = NewServiceInfo()
	refundServiceServiceInfoForClient       = NewServiceInfoForClient()
	refundServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return refundServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return refundServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return refundServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RefundService"
	handlerType := (*api.RefundService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.3",
		Extra:           extra,
	}
	return svcInfo
}

func refundCreateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.RefundCreateRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.RefundService).RefundCreate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RefundCreateArgs:
		success, err := handler.(api.RefundService).RefundCreate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefundCreateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRefundCreateArgs() interface{} {
	return &RefundCreateArgs{}
}

func newRefundCreateResult() interface{} {
	return &RefundCreateResult{}
}

type RefundCreateArgs struct {
	Req *api.RefundCreateRequest
}

func (p *RefundCreateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.RefundCreateRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefundCreateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefundCreateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefundCreateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefundCreateArgs) Unmarshal(in []byte) error {
	msg := new(api.RefundCreateRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefundCreateArgs_Req_DEFAULT *api.RefundCreateRequest

func (p *RefundCreateArgs) GetReq() *api.RefundCreateRequest {
	if !p.IsSetReq() {
		return RefundCreateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefundCreateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefundCreateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefundCreateResult struct {
	Success *api.RefundCreateResponse
}

var RefundCreateResult_Success_DEFAULT *api.RefundCreateResponse

func (p *RefundCreateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.RefundCreateResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefundCreateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefundCreateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefundCreateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefundCreateResult) Unmarshal(in []byte) error {
	msg := new(api.RefundCreateResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefundCreateResult) GetSuccess() *api.RefundCreateResponse {
	if !p.IsSetSuccess() {
		return RefundCreateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefundCreateResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.RefundCreateResponse)
}

func (p *RefundCreateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefundCreateResult) GetResult() interface{} {
	return p.Success
}

func refundQueryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.RefundQueryRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.RefundService).RefundQuery(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RefundQueryArgs:
		success, err := handler.(api.RefundService).RefundQuery(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefundQueryResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRefundQueryArgs() interface{} {
	return &RefundQueryArgs{}
}

func newRefundQueryResult() interface{} {
	return &RefundQueryResult{}
}

type RefundQueryArgs struct {
	Req *api.RefundQueryRequest
}

func (p *RefundQueryArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.RefundQueryRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefundQueryArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefundQueryArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefundQueryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefundQueryArgs) Unmarshal(in []byte) error {
	msg := new(api.RefundQueryRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefundQueryArgs_Req_DEFAULT *api.RefundQueryRequest

func (p *RefundQueryArgs) GetReq() *api.RefundQueryRequest {
	if !p.IsSetReq() {
		return RefundQueryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefundQueryArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefundQueryArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefundQueryResult struct {
	Success *api.RefundQueryResponse
}

var RefundQueryResult_Success_DEFAULT *api.RefundQueryResponse

func (p *RefundQueryResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.RefundQueryResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefundQueryResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefundQueryResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefundQueryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefundQueryResult) Unmarshal(in []byte) error {
	msg := new(api.RefundQueryResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefundQueryResult) GetSuccess() *api.RefundQueryResponse {
	if !p.IsSetSuccess() {
		return RefundQueryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefundQueryResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.RefundQueryResponse)
}

func (p *RefundQueryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefundQueryResult) GetResult() interface{} {
	return p.Success
}

func refundUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.RefundUpdateRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.RefundService).RefundUpdate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RefundUpdateArgs:
		success, err := handler.(api.RefundService).RefundUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefundUpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRefundUpdateArgs() interface{} {
	return &RefundUpdateArgs{}
}

func newRefundUpdateResult() interface{} {
	return &RefundUpdateResult{}
}

type RefundUpdateArgs struct {
	Req *api.RefundUpdateRequest
}

func (p *RefundUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.RefundUpdateRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefundUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefundUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefundUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefundUpdateArgs) Unmarshal(in []byte) error {
	msg := new(api.RefundUpdateRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefundUpdateArgs_Req_DEFAULT *api.RefundUpdateRequest

func (p *RefundUpdateArgs) GetReq() *api.RefundUpdateRequest {
	if !p.IsSetReq() {
		return RefundUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefundUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefundUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefundUpdateResult struct {
	Success *api.RefundUpdateResponse
}

var RefundUpdateResult_Success_DEFAULT *api.RefundUpdateResponse

func (p *RefundUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.RefundUpdateResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefundUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefundUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefundUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefundUpdateResult) Unmarshal(in []byte) error {
	msg := new(api.RefundUpdateResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefundUpdateResult) GetSuccess() *api.RefundUpdateResponse {
	if !p.IsSetSuccess() {
		return RefundUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefundUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.RefundUpdateResponse)
}

func (p *RefundUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefundUpdateResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RefundCreate(ctx context.Context, Req *api.RefundCreateRequest) (r *api.RefundCreateResponse, err error) {
	var _args RefundCreateArgs
	_args.Req = Req
	var _result RefundCreateResult
	if err = p.c.Call(ctx, "RefundCreate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefundQuery(ctx context.Context, Req *api.RefundQueryRequest) (r *api.RefundQueryResponse, err error) {
	var _args RefundQueryArgs
	_args.Req = Req
	var _result RefundQueryResult
	if err = p.c.Call(ctx, "RefundQuery", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefundUpdate(ctx context.Context, Req *api.RefundUpdateRequest) (r *api.RefundUpdateResponse, err error) {
	var _args RefundUpdateArgs
	_args.Req = Req
	var _result RefundUpdateResult
	if err = p.c.Call(ctx, "RefundUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
