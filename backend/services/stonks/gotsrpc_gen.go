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
	StonksServiceGoTSRPCProxyNewUser      = "NewUser"
	StonksServiceGoTSRPCProxyStartSession = "StartSession"
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
			arg_stonk string
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
		newUserRet := p.service.NewUser(&rw, r, arg_name)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{newUserRet}
			if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
				gotsrpc.ErrorCouldNotReply(w)
				return
			}
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	case StonksServiceGoTSRPCProxyStartSession:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_id string
		)
		args = []interface{}{&arg_id}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		rw := gotsrpc.ResponseWriter{ResponseWriter: w}
		startSessionRet, startSessionRet_1 := p.service.StartSession(&rw, r, arg_id)
		callStats.Execution = time.Since(executionStart)
		if rw.Status() == http.StatusOK {
			rets = []interface{}{startSessionRet, startSessionRet_1}
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
