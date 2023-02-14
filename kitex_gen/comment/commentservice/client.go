// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	comment "github.com/ClubWeGo/commentmicro/kitex_gen/comment"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CommentMethod(ctx context.Context, request *comment.CommentReq, callOptions ...callopt.Option) (r *comment.CommentResp, err error)
	CommentListMethod(ctx context.Context, request *comment.CommentListReq, callOptions ...callopt.Option) (r *comment.CommentListResp, err error)
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
	return &kCommentServiceClient{
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

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CommentMethod(ctx context.Context, request *comment.CommentReq, callOptions ...callopt.Option) (r *comment.CommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentMethod(ctx, request)
}

func (p *kCommentServiceClient) CommentListMethod(ctx context.Context, request *comment.CommentListReq, callOptions ...callopt.Option) (r *comment.CommentListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentListMethod(ctx, request)
}
