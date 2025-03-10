package gotool

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient 是HTTP客户端的接口定义
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPRequest 封装HTTP请求的参数
type HTTPRequest struct {
	URL         string
	Method      string
	Headers     map[string]string
	QueryParams map[string]string
	Body        interface{}
	Timeout     time.Duration
	Client      HTTPClient
	Context     context.Context
}

// HTTPResponse 封装HTTP响应
type HTTPResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	Error      error
}

// NewHTTPRequest 创建一个新的HTTP请求
func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{
		Method:      http.MethodGet,
		Headers:     make(map[string]string),
		QueryParams: make(map[string]string),
		Timeout:     30 * time.Second,
		Client:      &http.Client{},
		Context:     context.Background(),
	}
}

// SetURL 设置请求URL
func (r *HTTPRequest) SetURL(url string) *HTTPRequest {
	r.URL = url
	return r
}

// SetMethod 设置HTTP方法
func (r *HTTPRequest) SetMethod(method string) *HTTPRequest {
	r.Method = method
	return r
}

// SetHeader 设置单个请求头
func (r *HTTPRequest) SetHeader(key, value string) *HTTPRequest {
	r.Headers[key] = value
	return r
}

// SetHeaders 批量设置请求头
func (r *HTTPRequest) SetHeaders(headers map[string]string) *HTTPRequest {
	for k, v := range headers {
		r.Headers[k] = v
	}
	return r
}

// SetQueryParam 设置单个查询参数
func (r *HTTPRequest) SetQueryParam(key, value string) *HTTPRequest {
	r.QueryParams[key] = value
	return r
}

// SetQueryParams 批量设置查询参数
func (r *HTTPRequest) SetQueryParams(params map[string]string) *HTTPRequest {
	for k, v := range params {
		r.QueryParams[k] = v
	}
	return r
}

// SetBody 设置请求体
func (r *HTTPRequest) SetBody(body interface{}) *HTTPRequest {
	r.Body = body
	return r
}

// SetTimeout 设置请求超时时间
func (r *HTTPRequest) SetTimeout(timeout time.Duration) *HTTPRequest {
	r.Timeout = timeout
	return r
}

// SetClient 设置HTTP客户端
func (r *HTTPRequest) SetClient(client HTTPClient) *HTTPRequest {
	r.Client = client
	return r
}

// SetContext 设置上下文
func (r *HTTPRequest) SetContext(ctx context.Context) *HTTPRequest {
	r.Context = ctx
	return r
}

// buildURL 构建完整的URL，包括查询参数
func (r *HTTPRequest) buildURL() (string, error) {
	parsedURL, err := url.Parse(r.URL)
	if err != nil {
		return "", err
	}

	q := parsedURL.Query()
	for k, v := range r.QueryParams {
		q.Set(k, v)
	}

	parsedURL.RawQuery = q.Encode()
	return parsedURL.String(), nil
}

// prepareBody 准备请求体
func (r *HTTPRequest) prepareBody() (io.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}

	switch body := r.Body.(type) {
	case string:
		return strings.NewReader(body), nil
	case []byte:
		return bytes.NewReader(body), nil
	case io.Reader:
		return body, nil
	default:
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(jsonBody), nil
	}
}

// Send 发送HTTP请求
func (r *HTTPRequest) Send() *HTTPResponse {
	response := &HTTPResponse{}

	// 构建URL
	fullURL, err := r.buildURL()
	if err != nil {
		response.Error = fmt.Errorf("failed to build URL: %w", err)
		return response
	}

	// 准备请求体
	bodyReader, err := r.prepareBody()
	if err != nil {
		response.Error = fmt.Errorf("failed to prepare request body: %w", err)
		return response
	}

	// 创建请求
	req, err := http.NewRequestWithContext(r.Context, r.Method, fullURL, bodyReader)
	if err != nil {
		response.Error = fmt.Errorf("failed to create request: %w", err)
		return response
	}

	// 设置请求头
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// 如果是JSON请求体且未设置Content-Type，则设置为application/json
	if _, ok := r.Body.(string); !ok && r.Body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 创建客户端并设置超时
	client := r.Client
	if httpClient, ok := client.(*http.Client); ok {
		httpClient.Timeout = r.Timeout
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		response.Error = fmt.Errorf("request failed: %w", err)
		return response
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.Error = fmt.Errorf("failed to read response body: %w", err)
		return response
	}

	// 填充响应
	response.StatusCode = resp.StatusCode
	response.Headers = resp.Header
	response.Body = body

	return response
}

// String 将响应体转换为字符串
func (r *HTTPResponse) String() string {
	if r.Error != nil {
		return ""
	}
	return string(r.Body)
}

// JSON 将响应体解析为JSON
func (r *HTTPResponse) JSON(v interface{}) error {
	if r.Error != nil {
		return r.Error
	}
	return json.Unmarshal(r.Body, v)
}

// IsSuccess 检查响应是否成功（状态码2xx）
func (r *HTTPResponse) IsSuccess() bool {
	return r.Error == nil && r.StatusCode >= 200 && r.StatusCode < 300
}

// GET 发送GET请求
func GET(url string) *HTTPResponse {
	return NewHTTPRequest().SetMethod(http.MethodGet).SetURL(url).Send()
}

// POST 发送POST请求
func POST(url string, body interface{}) *HTTPResponse {
	return NewHTTPRequest().SetMethod(http.MethodPost).SetURL(url).SetBody(body).Send()
}

// PUT 发送PUT请求
func PUT(url string, body interface{}) *HTTPResponse {
	return NewHTTPRequest().SetMethod(http.MethodPut).SetURL(url).SetBody(body).Send()
}

// DELETE 发送DELETE请求
func DELETE(url string) *HTTPResponse {
	return NewHTTPRequest().SetMethod(http.MethodDelete).SetURL(url).Send()
}

// PATCH 发送PATCH请求
func PATCH(url string, body interface{}) *HTTPResponse {
	return NewHTTPRequest().SetMethod(http.MethodPatch).SetURL(url).SetBody(body).Send()
}

// GetJSON 发送GET请求并将响应解析为JSON
func GetJSON(url string, v interface{}) error {
	resp := GET(url)
	if resp.Error != nil {
		return resp.Error
	}
	return resp.JSON(v)
}

// PostJSON 发送POST请求并将响应解析为JSON
func PostJSON(url string, body, v interface{}) error {
	resp := POST(url, body)
	if resp.Error != nil {
		return resp.Error
	}
	return resp.JSON(v)
}
