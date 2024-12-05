package downloader

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Downloader struct {
	Url  *string
	Path *string
}

func NewDownloader(url *string, dirPath *string) *Downloader {
	if dirPath == nil || *dirPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf(homeDir, " not found", err)
		}
    var downloadDir string
    downloadDir = homeDir + "/Downloads"
    dirPath = &downloadDir
    fmt.Printf("Resolved dirPath to %s\n", *dirPath)
	}
	d := &Downloader{Url: url, Path: dirPath}
  fmt.Printf("Downloader: %+v %v\n", *d.Url, *d.Path)
	return d
}

func (d *Downloader) Download() {
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "wav", "-P", *d.Path, *d.Url)
	fmt.Println("Command args: ", cmd.Args)
	if *d.Url == "" {
		fmt.Println("You must provide a Url")
		return
	}
	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error excecuting video download")
		return
	}
  cmd.Start()
  if _, err := io.Copy(os.Stdout, output); err != nil {
    log.Fatal(err)
  }
  cmd.Wait()


}
