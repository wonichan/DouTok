// Code generated by Kitex v0.12.3. DO NOT EDIT.

package fileservice

import (
	api "DouTok-example/backend/baseService/kitex_gen/api"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PreSignGet(ctx context.Context, Req *api.PreSignGetRequest, callOptions ...callopt.Option) (r *api.PreSignGetResponse, err error)
	PreSignPut(ctx context.Context, Req *api.PreSignPutRequest, callOptions ...callopt.Option) (r *api.PreSignPutResponse, err error)
	ReportUploaded(ctx context.Context, Req *api.ReportUploadedRequest, callOptions ...callopt.Option) (r *api.ReportUploadedResponse, err error)
	PreSignSlicingPut(ctx context.Context, Req *api.PreSignSlicingPutRequest, callOptions ...callopt.Option) (r *api.PreSignSlicingPutResponse, err error)
	GetProgressRate4SlicingPut(ctx context.Context, Req *api.GetProgressRate4SlicingPutRequest, callOptions ...callopt.Option) (r *api.GetProgressRate4SlicingPutResponse, err error)
	MergeFileParts(ctx context.Context, Req *api.MergeFilePartsRequest, callOptions ...callopt.Option) (r *api.MergeFilePartsResponse, err error)
	RemoveFile(ctx context.Context, Req *api.RemoveFileRequest, callOptions ...callopt.Option) (r *api.RemoveFileResponse, err error)
	GetFileInfoById(ctx context.Context, Req *api.GetFileInfoByIdRequest, callOptions ...callopt.Option) (r *api.GetFileInfoByIdResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFileServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFileServiceClient struct {
	*kClient
}

func (p *kFileServiceClient) PreSignGet(ctx context.Context, Req *api.PreSignGetRequest, callOptions ...callopt.Option) (r *api.PreSignGetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PreSignGet(ctx, Req)
}

func (p *kFileServiceClient) PreSignPut(ctx context.Context, Req *api.PreSignPutRequest, callOptions ...callopt.Option) (r *api.PreSignPutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PreSignPut(ctx, Req)
}

func (p *kFileServiceClient) ReportUploaded(ctx context.Context, Req *api.ReportUploadedRequest, callOptions ...callopt.Option) (r *api.ReportUploadedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReportUploaded(ctx, Req)
}

func (p *kFileServiceClient) PreSignSlicingPut(ctx context.Context, Req *api.PreSignSlicingPutRequest, callOptions ...callopt.Option) (r *api.PreSignSlicingPutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PreSignSlicingPut(ctx, Req)
}

func (p *kFileServiceClient) GetProgressRate4SlicingPut(ctx context.Context, Req *api.GetProgressRate4SlicingPutRequest, callOptions ...callopt.Option) (r *api.GetProgressRate4SlicingPutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProgressRate4SlicingPut(ctx, Req)
}

func (p *kFileServiceClient) MergeFileParts(ctx context.Context, Req *api.MergeFilePartsRequest, callOptions ...callopt.Option) (r *api.MergeFilePartsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MergeFileParts(ctx, Req)
}

func (p *kFileServiceClient) RemoveFile(ctx context.Context, Req *api.RemoveFileRequest, callOptions ...callopt.Option) (r *api.RemoveFileResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveFile(ctx, Req)
}

func (p *kFileServiceClient) GetFileInfoById(ctx context.Context, Req *api.GetFileInfoByIdRequest, callOptions ...callopt.Option) (r *api.GetFileInfoByIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileInfoById(ctx, Req)
}
