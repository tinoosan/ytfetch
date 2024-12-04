package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func download(url *string) string {
	cmd := exec.Command("yt-dlp", "-x", "--audio-format","wav", *url)
  fmt.Println("Command args: ", cmd.Args)
	if *url == "" {
		fmt.Println("You must provide a url")
		return ""
	}
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Error excecuting video download")
		return ""
	}
	return string(output)
}

func main() {
	url := flag.String("u", "", "Youtube URL to download")
	flag.Parse()
	output := download(url)
	fmt.Printf("Arguments: %s \n", flag.Args())
	fmt.Println("Command output:", string(output))
}
