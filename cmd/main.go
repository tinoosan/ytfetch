package main

import (
	"flag"
	"yfetch/internal/downloader"
)

func main() {
	url := flag.String("u", "", "Youtube URL to download")
	dirPath := flag.String("p", "", "Directory to download file to")
	flag.Parse()
	d := downloader.NewDownloader(url, dirPath)
	d.Download()
}
