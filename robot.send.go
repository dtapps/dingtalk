package dingtalk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
	"time"
)

type RobotSendResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type RobotSendResult struct {
	Result RobotSendResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newRobotSendResult(result RobotSendResponse, body []byte, http gorequest.Response) *RobotSendResult {
	return &RobotSendResult{Result: result, Body: body, Http: http}
}

// RobotSend 发送消息
// https://open.dingtalk.com/document/group/custom-robot-access
func (c *Client) RobotSend(ctx context.Context, access_token string, notMustParams ...gorequest.Params) (*RobotSendResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "robot/send")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response RobotSendResponse
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/robot/send?access_token=%s", access_token), params, http.MethodPost, &response)
	return newRobotSendResult(response, request.ResponseBody, request), err
}

// RobotSendSign 发送消息签名版
// https://open.dingtalk.com/document/group/custom-robot-access
func (c *Client) RobotSendSign(ctx context.Context, access_token string, secret string, notMustParams ...gorequest.Params) (*RobotSendResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "robot/send")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 时间
	timestamp := time.Now().UnixNano() / 1e6

	// 请求
	var response RobotSendResponse
	request, err := c.request(ctx, fmt.Sprintf("robot/send?access_token=%s&timestamp=%d&sign=%s", access_token, timestamp, c.robotSendSignGetSign(secret, timestamp)), params, http.MethodPost, &response)
	return newRobotSendResult(response, request.ResponseBody, request), err
}

func (c *Client) robotSendSignGetSign(secret string, timestamp int64) string {
	secStr := fmt.Sprintf("%d\n%s", timestamp, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(secStr))
	result := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}
