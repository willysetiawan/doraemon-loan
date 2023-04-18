package httpcli

// import "github.com/go-resty/resty/v2"

// var client = resty.New().SetDebug(true)

// type HttpClient struct {
// 	ReqHeaders map[string]string
// 	URL        string
// 	ReqBody    interface{}
// }

// func Post(httpClient *HttpClient) (res *resty.Response, err error) {
// 	res, err = client.R().SetHeaders(httpClient.ReqHeaders).SetBody(httpClient.ReqBody).Post(httpClient.URL)
// 	return
// }

import (
	"bytes"
	"errors"
	"net"
	"net/http"
	helper "process-loan/lib/validate"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type (
	ResponseInfoSendHttp struct {
		Error      interface{} `json:"error"`
		StatusCode int         `json:"status_code"`
		Status     string      `json:"status"`
		Proto      string      `json:"proto"`
		Time       string      `json:"time"`
		ReceivedAt time.Time   `json:"received_at"`
		Body       interface{} `json:"body"`
	}
	RequestInfoSendHttp struct {
		DNSLookup      string      `json:"dns_lookup"`
		ConnTime       string      `json:"conn_time"`
		TCPConnTime    string      `json:"tcp_conn_time"`
		TLSHandshake   string      `json:"tls_hand_shake"`
		ServerTime     string      `json:"server_time"`
		ResponseTime   string      `json:"response_time"`
		TotalTime      string      `json:"total_time"`
		IsConnReused   bool        `json:"is_conn_reused"`
		IsConnWasIdle  bool        `json:"is_conn_was_idle"`
		ConnIdleTime   string      `json:"conn_idle_time"`
		RequestAttempt int         `json:"request_attempt"`
		RemoteAddr     net.Addr    `json:"remote_addr"`
		Body           interface{} `json:"body"`
	}
	InfoSendHttp struct {
		ResponseInfo ResponseInfoSendHttp `json:"response_info"`
		RequestInfo  RequestInfoSendHttp  `json:"request_info"`
	}
)

func Send(method string, url string, header map[string]string, req interface{}) (res []byte, info InfoSendHttp, err error) {
	var resp *resty.Response
	client := resty.New()
	request := client.R().
		EnableTrace().
		SetHeaders(header)
	switch method {
	case http.MethodGet:
		if !helper.NilInterface(req) {
			v, _ := query.Values(req)
			url = url + "?" + v.Encode()
		}
		resp, err = request.Get(url)

	case http.MethodPost:
		request.SetBody(req)
		resp, err = request.Post(url)
	case http.MethodDelete:
		request.SetBody(req)
		resp, err = request.Delete(url)
	case http.MethodPut:
		request.SetBody(req)
		resp, err = request.Put(url)
	case http.MethodPatch:
		request.SetBody(req)
		resp, err = request.Patch(url)
	default:
		err = errors.New("method not supported")
		return
	}
	if err != nil {
		return
	}
	ti := resp.Request.TraceInfo()
	resLog := ResponseInfoSendHttp{
		Error:      err,
		StatusCode: resp.StatusCode(),
		Status:     resp.Status(),
		Proto:      resp.Proto(),
		Time:       resp.Time().String(),
		ReceivedAt: resp.ReceivedAt().Local(),
		Body:       resp.String(),
	}
	reqLog := RequestInfoSendHttp{
		DNSLookup:      ti.DNSLookup.String(),
		ConnTime:       ti.ConnTime.String(),
		TCPConnTime:    ti.TCPConnTime.String(),
		TLSHandshake:   ti.TLSHandshake.String(),
		ServerTime:     ti.ServerTime.String(),
		ResponseTime:   ti.ResponseTime.String(),
		TotalTime:      ti.TotalTime.String(),
		IsConnReused:   ti.IsConnReused,
		IsConnWasIdle:  ti.IsConnWasIdle,
		ConnIdleTime:   ti.ConnIdleTime.String(),
		RequestAttempt: ti.RequestAttempt,
		RemoteAddr:     ti.RemoteAddr,
		Body:           req,
	}
	info.RequestInfo = reqLog
	info.ResponseInfo = resLog
	res = []byte(resp.String())

	return
}

func SendFormDataString(method string, url string, header map[string]string, req map[string]string) (res []byte, info InfoSendHttp, err error) {
	var resp *resty.Response
	client := resty.New()
	request := client.R().
		EnableTrace().
		SetHeaders(header)
	switch method {
	case http.MethodPost:
		request.SetFormData(req)
		resp, err = request.Post(url)
	default:
		err = errors.New("method not supported")
		return
	}
	if err != nil {
		return
	}
	ti := resp.Request.TraceInfo()
	resLog := ResponseInfoSendHttp{
		Error:      err,
		StatusCode: resp.StatusCode(),
		Status:     resp.Status(),
		Proto:      resp.Proto(),
		Time:       resp.Time().String(),
		ReceivedAt: resp.ReceivedAt().Local(),
		Body:       resp.String(),
	}
	reqLog := RequestInfoSendHttp{
		DNSLookup:      ti.DNSLookup.String(),
		ConnTime:       ti.ConnTime.String(),
		TCPConnTime:    ti.TCPConnTime.String(),
		TLSHandshake:   ti.TLSHandshake.String(),
		ServerTime:     ti.ServerTime.String(),
		ResponseTime:   ti.ResponseTime.String(),
		TotalTime:      ti.TotalTime.String(),
		IsConnReused:   ti.IsConnReused,
		IsConnWasIdle:  ti.IsConnWasIdle,
		ConnIdleTime:   ti.ConnIdleTime.String(),
		RequestAttempt: ti.RequestAttempt,
		RemoteAddr:     ti.RemoteAddr,
		Body:           req,
	}
	info.RequestInfo = reqLog
	info.ResponseInfo = resLog
	res = []byte(resp.String())

	return
}

func SendFileFaceCompare(method string, url string, header map[string]string, req [][]byte) (res []byte, info InfoSendHttp, err error) {
	var resp *resty.Response
	client := resty.New()
	request := client.R().
		EnableTrace().
		SetHeaders(header)
	switch method {
	case http.MethodPost:
		request.SetFileReader("firstImage", "onboarding.jpeg", bytes.NewReader(req[0]))
		request.SetFileReader("secondImage", "selfie_binding.jpeg", bytes.NewReader(req[1]))
		resp, err = request.Post(url)
	default:
		err = errors.New("method not supported")
		return
	}
	if err != nil {
		return
	}
	ti := resp.Request.TraceInfo()
	resLog := ResponseInfoSendHttp{
		Error:      err,
		StatusCode: resp.StatusCode(),
		Status:     resp.Status(),
		Proto:      resp.Proto(),
		Time:       resp.Time().String(),
		ReceivedAt: resp.ReceivedAt().Local(),
		Body:       resp.String(),
	}
	reqLog := RequestInfoSendHttp{
		DNSLookup:      ti.DNSLookup.String(),
		ConnTime:       ti.ConnTime.String(),
		TCPConnTime:    ti.TCPConnTime.String(),
		TLSHandshake:   ti.TLSHandshake.String(),
		ServerTime:     ti.ServerTime.String(),
		ResponseTime:   ti.ResponseTime.String(),
		TotalTime:      ti.TotalTime.String(),
		IsConnReused:   ti.IsConnReused,
		IsConnWasIdle:  ti.IsConnWasIdle,
		ConnIdleTime:   ti.ConnIdleTime.String(),
		RequestAttempt: ti.RequestAttempt,
		RemoteAddr:     ti.RemoteAddr,
		Body:           req,
	}
	info.RequestInfo = reqLog
	info.ResponseInfo = resLog
	res = []byte(resp.String())

	return
}
