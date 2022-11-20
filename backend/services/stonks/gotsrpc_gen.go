// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package stonks

import (
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	time "time"

	gotsrpc "github.com/foomo/gotsrpc/v2"
)

const (
	StonksServiceGoTSRPCProxyGetStonkInfo = "GetStonkInfo"
	StonksServiceGoTSRPCProxyGetUserInfo  = "GetUserInfo"
	StonksServiceGoTSRPCProxyNewUser      = "NewUser"
	StonksServiceGoTSRPCProxyPlaceOrder   = "PlaceOrder"
	StonksServiceGoTSRPCProxyUpdateOrder  = "UpdateOrder"
)

type StonksServiceGoTSRPCProxy struct {
	EndPoint string
	service  *StonksService
}

func NewDefaultStonksServiceGoTSRPCProxy(service *StonksService) *StonksServiceGoTSRPCProxy {
	return NewStonksServiceGoTSRPCProxy(service, "/service/stonks")
}

func NewStonksServiceGoTSRPCProxy(service *StonksService, endpoint string) *StonksServiceGoTSRPCProxy {
	return &StonksServiceGoTSRPCProxy{
		EndPoint: endpoint,
		service:  service,
	}
}

// ServeHTTP exposes your service
func (p *StonksServiceGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	} else if r.Method != http.MethodPost {
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	defer io.Copy(ioutil.Discard, r.Body) // Drain Request Body

	funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)
	callStats, _ := gotsrpc.GetStatsForRequest(r)
	callStats.Func = funcName
	callStats.Package = "github.com/hackaTUM/GameOfStonks/services/stonks"
	callStats.Service = "StonksService"
	switch funcName {
	case StonksServiceGoTSRPCProxyGetStonkInfo:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_stonk StonkName
		)
		args = []interface{}{&arg_stonk}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		rw := gotsrpc.ResponseWriter{ResponseWriter: w}
		getStonkInfoRet, getStonkInfoRet_1 := p.service.GetStonkInfo(&rw, r, arg_stonk)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{getStonkInfoRet, getStonkInfoRet_1}
			if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
				gotsrpc.ErrorCouldNotReply(w)
				return
			}
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	case StonksServiceGoTSRPCProxyGetUserInfo:
		var (
			args []interface{}
			rets []interface{}
		)
		executionStart := time.Now()
		rw := gotsrpc.ResponseWriter{ResponseWriter: w}
		getUserInfoRet, getUserInfoRet_1, getUserInfoRet_2 := p.service.GetUserInfo(&rw, r)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{getUserInfoRet, getUserInfoRet_1, getUserInfoRet_2}
			if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
				gotsrpc.ErrorCouldNotReply(w)
				return
			}
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	case StonksServiceGoTSRPCProxyNewUser:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_name string
		)
		args = []interface{}{&arg_name}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		rw := gotsrpc.ResponseWriter{ResponseWriter: w}
		newUserRet, newUserRet_1 := p.service.NewUser(&rw, r, arg_name)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{newUserRet, newUserRet_1}
			if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
				gotsrpc.ErrorCouldNotReply(w)
				return
			}
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	case StonksServiceGoTSRPCProxyPlaceOrder:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_cmd PlaceOrderCmd
		)
		args = []interface{}{&arg_cmd}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		rw := gotsrpc.ResponseWriter{ResponseWriter: w}
		placeOrderRet := p.service.PlaceOrder(&rw, r, arg_cmd)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{placeOrderRet}
			if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
				gotsrpc.ErrorCouldNotReply(w)
				return
			}
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	case StonksServiceGoTSRPCProxyUpdateOrder:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_cmd UpdateOrderCmd
		)
		args = []interface{}{&arg_cmd}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		rw := gotsrpc.ResponseWriter{ResponseWriter: w}
		updateOrderRet := p.service.UpdateOrder(&rw, r, arg_cmd)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{updateOrderRet}
			if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
				gotsrpc.ErrorCouldNotReply(w)
				return
			}
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	default:
		gotsrpc.ClearStats(r)
		gotsrpc.ErrorFuncNotFound(w)
	}
}
