package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var (
	linesum int
	mutex   *sync.Mutex = new(sync.Mutex)
)

var (
	// the dir where souce file stored
	rootPath, _ = os.Getwd()
	// exclude these sub dirs
	nodirs [5]string = [...]string{"/bitbucket.org", "/github.com", "/goplayer", "/uniqush", "/code.google.com"}
	// the suffix name you care
	suffixname string = ".go"
)

func main() {

	canAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8002")

	conn, err := net.ListenUDP("udp", canAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, 13)
		_, _, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//strData := string(data)
		fmt.Println("Received:", data)

		//upper := strings.ToUpper(strData)
		//_, err = conn.WriteToUDP([]byte(upper), rAddr)
		//if err != nil {
		//	fmt.Println(err)
		//	continue
		//}

		//fmt.Println("Send:", upper)

		//fmt.Printf("%x", -4700)
		////cmd := exec.Command("ffmpeg",
		////	"-f", "avfoundation",
		////	"-i", "1",
		////	"-vcodec", "libx264",
		////	"-preset", "ultrafast",
		////	"-acodec", "libfaac",
		////	"-f", "flv", "rtmp://video.nissanchina.cn/mec/456")
		//
		//cmd.Stdin = os.Stdin
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//
		//cmd.Run()
		//
		//argsLen := len(os.Args)
		//fmt.Println("argsLen:", argsLen)
		//if argsLen == 2 {
		//	rootPath = os.Args[1]
		//} else if argsLen == 3 {
		//	rootPath = os.Args[1]
		//	suffixname = os.Args[2]
		//}
		//// sync chan using for waiting
		//done := make(chan bool)
		//go codeLineSum(rootPath, done)
		//<-done
		//
		//fmt.Println("total line:", linesum)
	}

}
