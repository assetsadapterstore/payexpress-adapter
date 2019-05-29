package payexpress

import (
	"fmt"
	"github.com/blocktree/openwallet/log"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
)

type ClientInterface interface {
	Call(path string, request []interface{}) (*gjson.Result, error)
}

// A Client is a Bitcoin RPC client. It performs RPCs over HTTP using JSON
// request and responses. A Client must be configured with a secret token
// to authenticate with other Cores on the network.
type Client struct {
	BaseURL string
	Debug   bool
	client  *req.Req
}

const (
	ErrNotFound         = 404
	ErrUnknownException = 500
)

type Error struct {
	Code int    `json:"code,omitempty"`
	Err  string `json:"error,omitempty"`
}

//Errorf 生成OWError
func Errorf(code int, format string, a ...interface{}) *Error {
	err := &Error{
		Code: code,
		Err:  fmt.Sprintf(format, a...),
	}
	return err
}

//Error 错误信息
func (err *Error) Error() string {
	return fmt.Sprintf("[%d]%s", err.Code, err.Err)
}

func NewClient(url string, debug bool) *Client {

	url = strings.TrimSuffix(url, "/")
	c := Client{
		BaseURL: url,
		Debug:   debug,
	}

	api := req.New()
	//trans, _ := api.Client().Transport.(*http.Transport)
	//trans.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c.client = api

	return &c
}

// Call calls a remote procedure on another node, specified by the path.
func (c *Client) Call(path, method string, request interface{}) (*gjson.Result, *Error) {

	if c.client == nil {
		return nil, Errorf(ErrUnknownException, "API url is not setup.")
	}

	body := make([]interface{}, 0)

	authHeader := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	body = append(body, authHeader)
	if request != nil {
		body = append(body, req.BodyJSON(&request))
	}

	if c.Debug {
		log.Std.Info("Start Request API...")
	}
	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	r, err := c.client.Do(method, url, body...)

	if c.Debug {
		log.Std.Info("Request API Completed")
	}

	if c.Debug {
		log.Std.Info("%+v", r)
	}

	if err != nil {
		return nil, Errorf(ErrUnknownException, err.Error())
	}

	resp := gjson.ParseBytes(r.Bytes())

	if respErr := isError(r); respErr != nil {
		return nil, respErr
	}

	return &resp, nil
}

//isError 是否报错
func isError(r *req.Resp) *Error {

	if r.Response().StatusCode != http.StatusOK {
		message := gjson.GetBytes(r.Bytes(), "title").String()
		status := gjson.GetBytes(r.Bytes(), "status").Int()
		return Errorf(int(status), message)
	}

	return nil
}
