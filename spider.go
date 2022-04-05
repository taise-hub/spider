package spider

import (
	"fmt"
	"io"
	"net/http"
	"spider/parser"
	"strings"
)

type Spider struct {
	client *http.Client
	Option option
}

func New() *Spider {
	return &Spider{
		client: &http.Client{},
		Option: option{},
	}
}

// Crowlを開始する。
func (s *Spider) Start() error {
	dom, err := s.sendRequest(s.Option.Target)
	if err != nil {
		return err
	}
	urls, err := parser.Parse(dom)
	if err != nil {
		return err
	}
	a := s.matchTargetOrigin(urls)
	fmt.Printf("%v\n", a)
	return nil
}

func (s *Spider) sendRequest(url string) ([]byte, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (s *Spider) matchTargetOrigin(paths []string) []string {
	var list []string
	for _, path := range paths {
		if s.checkSameOrigin(path) {
			list = append(list, path)
		}
	}
	return list
}

func isFullPath(path string) bool {
	return strings.HasPrefix(path, "http://") ||
		strings.HasPrefix(path, "https://")
}

func (s *Spider) checkSameOrigin(path string) bool {
	if !isFullPath(path) {
		return true
	}
	return strings.HasPrefix(path, s.Option.Target)
}
