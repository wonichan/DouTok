// Code generated by Kitex v0.12.3. DO NOT EDIT.

package fileservice

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
	"PreSignGet": kitex.NewMethodInfo(
		preSignGetHandler,
		newPreSignGetArgs,
		newPreSignGetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"PreSignPut": kitex.NewMethodInfo(
		preSignPutHandler,
		newPreSignPutArgs,
		newPreSignPutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ReportUploaded": kitex.NewMethodInfo(
		reportUploadedHandler,
		newReportUploadedArgs,
		newReportUploadedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"PreSignSlicingPut": kitex.NewMethodInfo(
		preSignSlicingPutHandler,
		newPreSignSlicingPutArgs,
		newPreSignSlicingPutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetProgressRate4SlicingPut": kitex.NewMethodInfo(
		getProgressRate4SlicingPutHandler,
		newGetProgressRate4SlicingPutArgs,
		newGetProgressRate4SlicingPutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MergeFileParts": kitex.NewMethodInfo(
		mergeFilePartsHandler,
		newMergeFilePartsArgs,
		newMergeFilePartsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RemoveFile": kitex.NewMethodInfo(
		removeFileHandler,
		newRemoveFileArgs,
		newRemoveFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetFileInfoById": kitex.NewMethodInfo(
		getFileInfoByIdHandler,
		newGetFileInfoByIdArgs,
		newGetFileInfoByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	fileServiceServiceInfo                = NewServiceInfo()
	fileServiceServiceInfoForClient       = NewServiceInfoForClient()
	fileServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return fileServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return fileServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return fileServiceServiceInfoForClient
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
	serviceName := "FileService"
	handlerType := (*api.FileService)(nil)
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

func preSignGetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.PreSignGetRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).PreSignGet(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PreSignGetArgs:
		success, err := handler.(api.FileService).PreSignGet(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PreSignGetResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPreSignGetArgs() interface{} {
	return &PreSignGetArgs{}
}

func newPreSignGetResult() interface{} {
	return &PreSignGetResult{}
}

type PreSignGetArgs struct {
	Req *api.PreSignGetRequest
}

func (p *PreSignGetArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.PreSignGetRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PreSignGetArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PreSignGetArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PreSignGetArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PreSignGetArgs) Unmarshal(in []byte) error {
	msg := new(api.PreSignGetRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PreSignGetArgs_Req_DEFAULT *api.PreSignGetRequest

func (p *PreSignGetArgs) GetReq() *api.PreSignGetRequest {
	if !p.IsSetReq() {
		return PreSignGetArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PreSignGetArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PreSignGetArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PreSignGetResult struct {
	Success *api.PreSignGetResponse
}

var PreSignGetResult_Success_DEFAULT *api.PreSignGetResponse

func (p *PreSignGetResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.PreSignGetResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PreSignGetResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PreSignGetResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PreSignGetResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PreSignGetResult) Unmarshal(in []byte) error {
	msg := new(api.PreSignGetResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PreSignGetResult) GetSuccess() *api.PreSignGetResponse {
	if !p.IsSetSuccess() {
		return PreSignGetResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PreSignGetResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.PreSignGetResponse)
}

func (p *PreSignGetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PreSignGetResult) GetResult() interface{} {
	return p.Success
}

func preSignPutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.PreSignPutRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).PreSignPut(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PreSignPutArgs:
		success, err := handler.(api.FileService).PreSignPut(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PreSignPutResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPreSignPutArgs() interface{} {
	return &PreSignPutArgs{}
}

func newPreSignPutResult() interface{} {
	return &PreSignPutResult{}
}

type PreSignPutArgs struct {
	Req *api.PreSignPutRequest
}

func (p *PreSignPutArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.PreSignPutRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PreSignPutArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PreSignPutArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PreSignPutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PreSignPutArgs) Unmarshal(in []byte) error {
	msg := new(api.PreSignPutRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PreSignPutArgs_Req_DEFAULT *api.PreSignPutRequest

func (p *PreSignPutArgs) GetReq() *api.PreSignPutRequest {
	if !p.IsSetReq() {
		return PreSignPutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PreSignPutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PreSignPutArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PreSignPutResult struct {
	Success *api.PreSignPutResponse
}

var PreSignPutResult_Success_DEFAULT *api.PreSignPutResponse

func (p *PreSignPutResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.PreSignPutResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PreSignPutResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PreSignPutResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PreSignPutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PreSignPutResult) Unmarshal(in []byte) error {
	msg := new(api.PreSignPutResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PreSignPutResult) GetSuccess() *api.PreSignPutResponse {
	if !p.IsSetSuccess() {
		return PreSignPutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PreSignPutResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.PreSignPutResponse)
}

func (p *PreSignPutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PreSignPutResult) GetResult() interface{} {
	return p.Success
}

func reportUploadedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.ReportUploadedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).ReportUploaded(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ReportUploadedArgs:
		success, err := handler.(api.FileService).ReportUploaded(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ReportUploadedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newReportUploadedArgs() interface{} {
	return &ReportUploadedArgs{}
}

func newReportUploadedResult() interface{} {
	return &ReportUploadedResult{}
}

type ReportUploadedArgs struct {
	Req *api.ReportUploadedRequest
}

func (p *ReportUploadedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.ReportUploadedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ReportUploadedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ReportUploadedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ReportUploadedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ReportUploadedArgs) Unmarshal(in []byte) error {
	msg := new(api.ReportUploadedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ReportUploadedArgs_Req_DEFAULT *api.ReportUploadedRequest

func (p *ReportUploadedArgs) GetReq() *api.ReportUploadedRequest {
	if !p.IsSetReq() {
		return ReportUploadedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReportUploadedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ReportUploadedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ReportUploadedResult struct {
	Success *api.ReportUploadedResponse
}

var ReportUploadedResult_Success_DEFAULT *api.ReportUploadedResponse

func (p *ReportUploadedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.ReportUploadedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ReportUploadedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ReportUploadedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ReportUploadedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ReportUploadedResult) Unmarshal(in []byte) error {
	msg := new(api.ReportUploadedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReportUploadedResult) GetSuccess() *api.ReportUploadedResponse {
	if !p.IsSetSuccess() {
		return ReportUploadedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReportUploadedResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.ReportUploadedResponse)
}

func (p *ReportUploadedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReportUploadedResult) GetResult() interface{} {
	return p.Success
}

func preSignSlicingPutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.PreSignSlicingPutRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).PreSignSlicingPut(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PreSignSlicingPutArgs:
		success, err := handler.(api.FileService).PreSignSlicingPut(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PreSignSlicingPutResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPreSignSlicingPutArgs() interface{} {
	return &PreSignSlicingPutArgs{}
}

func newPreSignSlicingPutResult() interface{} {
	return &PreSignSlicingPutResult{}
}

type PreSignSlicingPutArgs struct {
	Req *api.PreSignSlicingPutRequest
}

func (p *PreSignSlicingPutArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.PreSignSlicingPutRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PreSignSlicingPutArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PreSignSlicingPutArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PreSignSlicingPutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PreSignSlicingPutArgs) Unmarshal(in []byte) error {
	msg := new(api.PreSignSlicingPutRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PreSignSlicingPutArgs_Req_DEFAULT *api.PreSignSlicingPutRequest

func (p *PreSignSlicingPutArgs) GetReq() *api.PreSignSlicingPutRequest {
	if !p.IsSetReq() {
		return PreSignSlicingPutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PreSignSlicingPutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PreSignSlicingPutArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PreSignSlicingPutResult struct {
	Success *api.PreSignSlicingPutResponse
}

var PreSignSlicingPutResult_Success_DEFAULT *api.PreSignSlicingPutResponse

func (p *PreSignSlicingPutResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.PreSignSlicingPutResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PreSignSlicingPutResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PreSignSlicingPutResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PreSignSlicingPutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PreSignSlicingPutResult) Unmarshal(in []byte) error {
	msg := new(api.PreSignSlicingPutResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PreSignSlicingPutResult) GetSuccess() *api.PreSignSlicingPutResponse {
	if !p.IsSetSuccess() {
		return PreSignSlicingPutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PreSignSlicingPutResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.PreSignSlicingPutResponse)
}

func (p *PreSignSlicingPutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PreSignSlicingPutResult) GetResult() interface{} {
	return p.Success
}

func getProgressRate4SlicingPutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.GetProgressRate4SlicingPutRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).GetProgressRate4SlicingPut(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetProgressRate4SlicingPutArgs:
		success, err := handler.(api.FileService).GetProgressRate4SlicingPut(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetProgressRate4SlicingPutResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetProgressRate4SlicingPutArgs() interface{} {
	return &GetProgressRate4SlicingPutArgs{}
}

func newGetProgressRate4SlicingPutResult() interface{} {
	return &GetProgressRate4SlicingPutResult{}
}

type GetProgressRate4SlicingPutArgs struct {
	Req *api.GetProgressRate4SlicingPutRequest
}

func (p *GetProgressRate4SlicingPutArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.GetProgressRate4SlicingPutRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetProgressRate4SlicingPutArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetProgressRate4SlicingPutArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetProgressRate4SlicingPutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetProgressRate4SlicingPutArgs) Unmarshal(in []byte) error {
	msg := new(api.GetProgressRate4SlicingPutRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetProgressRate4SlicingPutArgs_Req_DEFAULT *api.GetProgressRate4SlicingPutRequest

func (p *GetProgressRate4SlicingPutArgs) GetReq() *api.GetProgressRate4SlicingPutRequest {
	if !p.IsSetReq() {
		return GetProgressRate4SlicingPutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetProgressRate4SlicingPutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetProgressRate4SlicingPutArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetProgressRate4SlicingPutResult struct {
	Success *api.GetProgressRate4SlicingPutResponse
}

var GetProgressRate4SlicingPutResult_Success_DEFAULT *api.GetProgressRate4SlicingPutResponse

func (p *GetProgressRate4SlicingPutResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.GetProgressRate4SlicingPutResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetProgressRate4SlicingPutResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetProgressRate4SlicingPutResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetProgressRate4SlicingPutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetProgressRate4SlicingPutResult) Unmarshal(in []byte) error {
	msg := new(api.GetProgressRate4SlicingPutResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProgressRate4SlicingPutResult) GetSuccess() *api.GetProgressRate4SlicingPutResponse {
	if !p.IsSetSuccess() {
		return GetProgressRate4SlicingPutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetProgressRate4SlicingPutResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.GetProgressRate4SlicingPutResponse)
}

func (p *GetProgressRate4SlicingPutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetProgressRate4SlicingPutResult) GetResult() interface{} {
	return p.Success
}

func mergeFilePartsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.MergeFilePartsRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).MergeFileParts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MergeFilePartsArgs:
		success, err := handler.(api.FileService).MergeFileParts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MergeFilePartsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMergeFilePartsArgs() interface{} {
	return &MergeFilePartsArgs{}
}

func newMergeFilePartsResult() interface{} {
	return &MergeFilePartsResult{}
}

type MergeFilePartsArgs struct {
	Req *api.MergeFilePartsRequest
}

func (p *MergeFilePartsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.MergeFilePartsRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MergeFilePartsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MergeFilePartsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MergeFilePartsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MergeFilePartsArgs) Unmarshal(in []byte) error {
	msg := new(api.MergeFilePartsRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MergeFilePartsArgs_Req_DEFAULT *api.MergeFilePartsRequest

func (p *MergeFilePartsArgs) GetReq() *api.MergeFilePartsRequest {
	if !p.IsSetReq() {
		return MergeFilePartsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MergeFilePartsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MergeFilePartsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MergeFilePartsResult struct {
	Success *api.MergeFilePartsResponse
}

var MergeFilePartsResult_Success_DEFAULT *api.MergeFilePartsResponse

func (p *MergeFilePartsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.MergeFilePartsResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MergeFilePartsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MergeFilePartsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MergeFilePartsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MergeFilePartsResult) Unmarshal(in []byte) error {
	msg := new(api.MergeFilePartsResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MergeFilePartsResult) GetSuccess() *api.MergeFilePartsResponse {
	if !p.IsSetSuccess() {
		return MergeFilePartsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MergeFilePartsResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.MergeFilePartsResponse)
}

func (p *MergeFilePartsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MergeFilePartsResult) GetResult() interface{} {
	return p.Success
}

func removeFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.RemoveFileRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).RemoveFile(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RemoveFileArgs:
		success, err := handler.(api.FileService).RemoveFile(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RemoveFileResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRemoveFileArgs() interface{} {
	return &RemoveFileArgs{}
}

func newRemoveFileResult() interface{} {
	return &RemoveFileResult{}
}

type RemoveFileArgs struct {
	Req *api.RemoveFileRequest
}

func (p *RemoveFileArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.RemoveFileRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RemoveFileArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RemoveFileArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RemoveFileArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RemoveFileArgs) Unmarshal(in []byte) error {
	msg := new(api.RemoveFileRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RemoveFileArgs_Req_DEFAULT *api.RemoveFileRequest

func (p *RemoveFileArgs) GetReq() *api.RemoveFileRequest {
	if !p.IsSetReq() {
		return RemoveFileArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RemoveFileArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RemoveFileArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RemoveFileResult struct {
	Success *api.RemoveFileResponse
}

var RemoveFileResult_Success_DEFAULT *api.RemoveFileResponse

func (p *RemoveFileResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.RemoveFileResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RemoveFileResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RemoveFileResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RemoveFileResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RemoveFileResult) Unmarshal(in []byte) error {
	msg := new(api.RemoveFileResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RemoveFileResult) GetSuccess() *api.RemoveFileResponse {
	if !p.IsSetSuccess() {
		return RemoveFileResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RemoveFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.RemoveFileResponse)
}

func (p *RemoveFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RemoveFileResult) GetResult() interface{} {
	return p.Success
}

func getFileInfoByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.GetFileInfoByIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FileService).GetFileInfoById(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetFileInfoByIdArgs:
		success, err := handler.(api.FileService).GetFileInfoById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFileInfoByIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetFileInfoByIdArgs() interface{} {
	return &GetFileInfoByIdArgs{}
}

func newGetFileInfoByIdResult() interface{} {
	return &GetFileInfoByIdResult{}
}

type GetFileInfoByIdArgs struct {
	Req *api.GetFileInfoByIdRequest
}

func (p *GetFileInfoByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.GetFileInfoByIdRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFileInfoByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFileInfoByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFileInfoByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetFileInfoByIdArgs) Unmarshal(in []byte) error {
	msg := new(api.GetFileInfoByIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFileInfoByIdArgs_Req_DEFAULT *api.GetFileInfoByIdRequest

func (p *GetFileInfoByIdArgs) GetReq() *api.GetFileInfoByIdRequest {
	if !p.IsSetReq() {
		return GetFileInfoByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFileInfoByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFileInfoByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFileInfoByIdResult struct {
	Success *api.GetFileInfoByIdResponse
}

var GetFileInfoByIdResult_Success_DEFAULT *api.GetFileInfoByIdResponse

func (p *GetFileInfoByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.GetFileInfoByIdResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFileInfoByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFileInfoByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFileInfoByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetFileInfoByIdResult) Unmarshal(in []byte) error {
	msg := new(api.GetFileInfoByIdResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFileInfoByIdResult) GetSuccess() *api.GetFileInfoByIdResponse {
	if !p.IsSetSuccess() {
		return GetFileInfoByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFileInfoByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.GetFileInfoByIdResponse)
}

func (p *GetFileInfoByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFileInfoByIdResult) GetResult() interface{} {
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

func (p *kClient) PreSignGet(ctx context.Context, Req *api.PreSignGetRequest) (r *api.PreSignGetResponse, err error) {
	var _args PreSignGetArgs
	_args.Req = Req
	var _result PreSignGetResult
	if err = p.c.Call(ctx, "PreSignGet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PreSignPut(ctx context.Context, Req *api.PreSignPutRequest) (r *api.PreSignPutResponse, err error) {
	var _args PreSignPutArgs
	_args.Req = Req
	var _result PreSignPutResult
	if err = p.c.Call(ctx, "PreSignPut", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReportUploaded(ctx context.Context, Req *api.ReportUploadedRequest) (r *api.ReportUploadedResponse, err error) {
	var _args ReportUploadedArgs
	_args.Req = Req
	var _result ReportUploadedResult
	if err = p.c.Call(ctx, "ReportUploaded", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PreSignSlicingPut(ctx context.Context, Req *api.PreSignSlicingPutRequest) (r *api.PreSignSlicingPutResponse, err error) {
	var _args PreSignSlicingPutArgs
	_args.Req = Req
	var _result PreSignSlicingPutResult
	if err = p.c.Call(ctx, "PreSignSlicingPut", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProgressRate4SlicingPut(ctx context.Context, Req *api.GetProgressRate4SlicingPutRequest) (r *api.GetProgressRate4SlicingPutResponse, err error) {
	var _args GetProgressRate4SlicingPutArgs
	_args.Req = Req
	var _result GetProgressRate4SlicingPutResult
	if err = p.c.Call(ctx, "GetProgressRate4SlicingPut", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MergeFileParts(ctx context.Context, Req *api.MergeFilePartsRequest) (r *api.MergeFilePartsResponse, err error) {
	var _args MergeFilePartsArgs
	_args.Req = Req
	var _result MergeFilePartsResult
	if err = p.c.Call(ctx, "MergeFileParts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveFile(ctx context.Context, Req *api.RemoveFileRequest) (r *api.RemoveFileResponse, err error) {
	var _args RemoveFileArgs
	_args.Req = Req
	var _result RemoveFileResult
	if err = p.c.Call(ctx, "RemoveFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFileInfoById(ctx context.Context, Req *api.GetFileInfoByIdRequest) (r *api.GetFileInfoByIdResponse, err error) {
	var _args GetFileInfoByIdArgs
	_args.Req = Req
	var _result GetFileInfoByIdResult
	if err = p.c.Call(ctx, "GetFileInfoById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
