// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package stonks

import (
	go_context "context"
	go_net_http "net/http"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	pkg_errors "github.com/pkg/errors"
)

type StonksServiceGoTSRPCClient interface {
	GetStonkInfo(ctx go_context.Context, stonk string) (retGetStonkInfo_0 StonkInfo, retGetStonkInfo_1 *Err, clientErr error)
	NewUser(ctx go_context.Context, name string) (retNewUser_0 *Err, clientErr error)
	StartSession(ctx go_context.Context, id string) (retStartSession_0 []User, retStartSession_1 *Err, clientErr error)
}

type HTTPStonksServiceGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultStonksServiceGoTSRPCClient(url string) *HTTPStonksServiceGoTSRPCClient {
	return NewStonksServiceGoTSRPCClient(url, "/service/stonks")
}

func NewStonksServiceGoTSRPCClient(url string, endpoint string) *HTTPStonksServiceGoTSRPCClient {
	return NewStonksServiceGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewStonksServiceGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPStonksServiceGoTSRPCClient {
	return &HTTPStonksServiceGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPStonksServiceGoTSRPCClient) GetStonkInfo(ctx go_context.Context, stonk string) (retGetStonkInfo_0 StonkInfo, retGetStonkInfo_1 *Err, clientErr error) {
	args := []interface{}{stonk}
	reply := []interface{}{&retGetStonkInfo_0, &retGetStonkInfo_1}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "GetStonkInfo", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call stonks.StonksServiceGoTSRPCProxy GetStonkInfo")
	}
	return
}

func (tsc *HTTPStonksServiceGoTSRPCClient) NewUser(ctx go_context.Context, name string) (retNewUser_0 *Err, clientErr error) {
	args := []interface{}{name}
	reply := []interface{}{&retNewUser_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "NewUser", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call stonks.StonksServiceGoTSRPCProxy NewUser")
	}
	return
}

func (tsc *HTTPStonksServiceGoTSRPCClient) StartSession(ctx go_context.Context, id string) (retStartSession_0 []User, retStartSession_1 *Err, clientErr error) {
	args := []interface{}{id}
	reply := []interface{}{&retStartSession_0, &retStartSession_1}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "StartSession", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call stonks.StonksServiceGoTSRPCProxy StartSession")
	}
	return
}
