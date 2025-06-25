package url

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func InitUrl(input string) *TargetUrl {
	return &TargetUrl{Input: input}
}

type TargetUrl struct {
	Input string `json:"input" dc:"输入的 URL"`
}

type ParsedUrl struct {
	Host       string `json:"host" dc:"完整域名"`
	MainDomain string `json:"main_domain" dc:"主域名"`
	Path       string `json:"path" dc:"路径"`
	Query      string `json:"query" dc:"查询参数"`
}

// extractMainDomain 提取主域名（比如 example.com）
func (t *TargetUrl) extractMainDomain(host string) (string, error) {
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return "", errors.New("host 格式无效")
	}
	// 获取最后两个部分作为主域名
	return strings.Join(parts[len(parts)-2:], "."), nil
}

// ValidateAndExtractURL 校验 URL 并提取主域名
func (t *TargetUrl) ValidateAndExtractURL() (string, error) {
	// 解析 URL
	parsedURL, err := url.Parse(t.Input)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", errors.New("无效的 URL")
	}

	// 提取主域名
	mainDomain, err := t.extractMainDomain(parsedURL.Hostname())
	if err != nil {
		return "", err
	}
	return mainDomain, nil
}

// Parse 解析为对象
func (t *TargetUrl) Parse() (*ParsedUrl, error) {
	// 解析 URL
	parsedURL, err := url.Parse(t.Input)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, errors.New("无效的 URL")
	}

	// 提取主域名
	mainDomain, err := t.extractMainDomain(parsedURL.Hostname())
	if err != nil {
		return nil, err
	}

	// 拼接参数部分
	params := parsedURL.RawQuery
	return &ParsedUrl{
		Host:       parsedURL.Host,
		MainDomain: mainDomain,
		Path:       parsedURL.Path,
		Query:      params,
	}, nil
}

// GetDomainAndParams 返回域名 + 参数的字符串
func (t *TargetUrl) GetDomainAndParams() (string, error) {
	parsed, err := t.Parse()
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s?%s", parsed.Host, parsed.Query), nil
}
