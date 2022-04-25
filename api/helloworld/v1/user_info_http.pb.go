// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserInfoServiceHTTPServer interface {
	Delete(context.Context, *IdRequest) (*CommReply, error)
	Get(context.Context, *IdRequest) (*UserInfoReply, error)
	List(context.Context, *ListRequest) (*ListUserInfoReply, error)
	Save(context.Context, *UserInfo) (*CommReply, error)
}

func RegisterUserInfoServiceHTTPServer(s *http.Server, srv UserInfoServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/userInfo/save", _UserInfoService_Save0_HTTP_Handler(srv))
	r.POST("/userInfo/delete", _UserInfoService_Delete0_HTTP_Handler(srv))
	r.GET("/userInfo/get/{id}", _UserInfoService_Get0_HTTP_Handler(srv))
	r.POST("/userInfo/list", _UserInfoService_List0_HTTP_Handler(srv))
}

func _UserInfoService_Save0_HTTP_Handler(srv UserInfoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/demo.v1.UserInfoService/Save")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Save(ctx, req.(*UserInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommReply)
		return ctx.Result(200, reply)
	}
}

func _UserInfoService_Delete0_HTTP_Handler(srv UserInfoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/demo.v1.UserInfoService/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommReply)
		return ctx.Result(200, reply)
	}
}

func _UserInfoService_Get0_HTTP_Handler(srv UserInfoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/demo.v1.UserInfoService/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _UserInfoService_List0_HTTP_Handler(srv UserInfoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/demo.v1.UserInfoService/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserInfoReply)
		return ctx.Result(200, reply)
	}
}

type UserInfoServiceHTTPClient interface {
	Delete(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *CommReply, err error)
	Get(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	List(ctx context.Context, req *ListRequest, opts ...http.CallOption) (rsp *ListUserInfoReply, err error)
	Save(ctx context.Context, req *UserInfo, opts ...http.CallOption) (rsp *CommReply, err error)
}

type UserInfoServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserInfoServiceHTTPClient(client *http.Client) UserInfoServiceHTTPClient {
	return &UserInfoServiceHTTPClientImpl{client}
}

func (c *UserInfoServiceHTTPClientImpl) Delete(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*CommReply, error) {
	var out CommReply
	pattern := "/userInfo/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/demo.v1.UserInfoService/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserInfoServiceHTTPClientImpl) Get(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/userInfo/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/demo.v1.UserInfoService/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserInfoServiceHTTPClientImpl) List(ctx context.Context, in *ListRequest, opts ...http.CallOption) (*ListUserInfoReply, error) {
	var out ListUserInfoReply
	pattern := "/userInfo/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/demo.v1.UserInfoService/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserInfoServiceHTTPClientImpl) Save(ctx context.Context, in *UserInfo, opts ...http.CallOption) (*CommReply, error) {
	var out CommReply
	pattern := "/userInfo/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/demo.v1.UserInfoService/Save"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
