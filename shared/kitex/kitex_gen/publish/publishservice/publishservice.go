// Code generated by Kitex v0.6.2. DO NOT EDIT.

package publishservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	publish "tuzi-tiktok/kitex/kitex_gen/publish"
)

func serviceInfo() *kitex.ServiceInfo {
	return publishServiceServiceInfo
}

var publishServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PublishService"
	handlerType := (*publish.PublishService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishVideo":   kitex.NewMethodInfo(publishVideoHandler, newPublishVideoArgs, newPublishVideoResult, false),
		"GetPublishList": kitex.NewMethodInfo(getPublishListHandler, newGetPublishListArgs, newGetPublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "idl",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.PublishRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishService).PublishVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishVideoArgs:
		success, err := handler.(publish.PublishService).PublishVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishVideoResult)
		realResult.Success = success
	}
	return nil
}
func newPublishVideoArgs() interface{} {
	return &PublishVideoArgs{}
}

func newPublishVideoResult() interface{} {
	return &PublishVideoResult{}
}

type PublishVideoArgs struct {
	Req *publish.PublishRequest
}

func (p *PublishVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(publish.PublishRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishVideoArgs) Unmarshal(in []byte) error {
	msg := new(publish.PublishRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishVideoArgs_Req_DEFAULT *publish.PublishRequest

func (p *PublishVideoArgs) GetReq() *publish.PublishRequest {
	if !p.IsSetReq() {
		return PublishVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PublishVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PublishVideoResult struct {
	Success *publish.PublishResponse
}

var PublishVideoResult_Success_DEFAULT *publish.PublishResponse

func (p *PublishVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(publish.PublishResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishVideoResult) Unmarshal(in []byte) error {
	msg := new(publish.PublishResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishVideoResult) GetSuccess() *publish.PublishResponse {
	if !p.IsSetSuccess() {
		return PublishVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.PublishResponse)
}

func (p *PublishVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PublishVideoResult) GetResult() interface{} {
	return p.Success
}

func getPublishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.PublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.PublishService).GetPublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPublishListArgs:
		success, err := handler.(publish.PublishService).GetPublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPublishListResult)
		realResult.Success = success
	}
	return nil
}
func newGetPublishListArgs() interface{} {
	return &GetPublishListArgs{}
}

func newGetPublishListResult() interface{} {
	return &GetPublishListResult{}
}

type GetPublishListArgs struct {
	Req *publish.PublishListRequest
}

func (p *GetPublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(publish.PublishListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPublishListArgs) Unmarshal(in []byte) error {
	msg := new(publish.PublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPublishListArgs_Req_DEFAULT *publish.PublishListRequest

func (p *GetPublishListArgs) GetReq() *publish.PublishListRequest {
	if !p.IsSetReq() {
		return GetPublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPublishListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPublishListResult struct {
	Success *publish.PublishListResponse
}

var GetPublishListResult_Success_DEFAULT *publish.PublishListResponse

func (p *GetPublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(publish.PublishListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPublishListResult) Unmarshal(in []byte) error {
	msg := new(publish.PublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPublishListResult) GetSuccess() *publish.PublishListResponse {
	if !p.IsSetSuccess() {
		return GetPublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.PublishListResponse)
}

func (p *GetPublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPublishListResult) GetResult() interface{} {
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

func (p *kClient) PublishVideo(ctx context.Context, Req *publish.PublishRequest) (r *publish.PublishResponse, err error) {
	var _args PublishVideoArgs
	_args.Req = Req
	var _result PublishVideoResult
	if err = p.c.Call(ctx, "PublishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPublishList(ctx context.Context, Req *publish.PublishListRequest) (r *publish.PublishListResponse, err error) {
	var _args GetPublishListArgs
	_args.Req = Req
	var _result GetPublishListResult
	if err = p.c.Call(ctx, "GetPublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}