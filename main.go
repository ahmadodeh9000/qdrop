package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	qrcode "github.com/skip2/go-qrcode"
)

// getting random token

func getInviteToken() string {
	tok := make([]byte, 8)
	if _, err := rand.Read(tok); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(tok)
}

// serving a file
func serveFile(ln net.Listener, path, name string, t string) error {
	mux := http.NewServeMux()
	done := make(chan struct{})
	var once sync.Once

	mux.HandleFunc("/"+t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition",
			mime.FormatMediaType("attachment", map[string]string{"filename": name}))
		http.ServeFile(w, r, path)
		once.Do(func() { close(done) })
	})

	// when da file is downloaded it shuts down
	srv := &http.Server{Handler: mux}
	go func() {
		<-done
		fmt.Println("FILE HAS BEEN DOWNLOADED !")
		srv.Shutdown(context.Background())
	}()

	if err := srv.Serve(ln); err != http.ErrServerClosed {
		return err
	}

	return nil
}

// getting local ip addr

func startListener() (net.Listener, int) {
	ln, err := net.Listen("tcp", ":0")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	port := ln.Addr().(*net.TCPAddr).Port

	return ln, port
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "1.1.1.1:80")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer conn.Close()

	ip := conn.LocalAddr().(*net.UDPAddr).IP.String()

	return ip

}

// file stuff

func getFile() (string, string) {
	if len(os.Args) < 2 {
		fmt.Println("usage: qdrop <path>")
		os.Exit(1)
	}

	fpath := os.Args[1]

	info, err := os.Stat(fpath)

	if err != nil {
		fmt.Println("qdrop: CANNOT READ FILE", err)
		os.Exit(1)
	}

	if info.IsDir() {
		fmt.Println("qdrop: is a dir not a file lol")
		os.Exit(1)
	}

	return fpath, filepath.Base(fpath)
}

// frontend

func getFileSize(path string) string {
	info, err := os.Stat(path)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	bytes := info.Size()

	if bytes > 1000 {

		sizeKB := float64(bytes) / 1000.0
		return fmt.Sprintf("%v KB", sizeKB)
	}

	return fmt.Sprintf("%v B", float64(bytes))

}

func show(name string, path string, link string) {
	// print logo

	// printing other info

	fmt.Printf("FILE : \t %v (%v)\n", name, getFileSize(path))
	fmt.Printf("LINK : \t %v\n", link)

	// printing qr code
	getQrCode(link)
}

// prints qr code to the terminal

func getQrCode(url string) {
	qr, err := qrcode.New(url, qrcode.Low)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(qr.ToSmallString(false))
}

func main() {
	path, name := getFile()
	ip_addr := getLocalIP()
	ln, port := startListener()
	tok := getInviteToken()

	defer ln.Close()

	url := fmt.Sprintf("http://%v:%v/"+tok, ip_addr, port)

	show(name, path, url)
	serveFile(ln, path, name, tok)
}
