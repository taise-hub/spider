package parser

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func Parse(data []byte) ([]string, error) {
	var urls []string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urls = append(urls, url)
	})
	return urls, nil
}
