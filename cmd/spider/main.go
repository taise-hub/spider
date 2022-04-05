package main

import (
	"flag"
	"spider"
)

func main() {
	spider := spider.New()
	flag.StringVar(&spider.Option.Target, "u", "", "target URL")
	flag.UintVar(&spider.Option.Depth, "d", 0, "depth of crowling")
	flag.Parse()
	spider.Start()
}
