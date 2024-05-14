package main

import "URL_shorter/cmd/server"

func main() {
	server.Init()
	server.Run()
	defer server.Close()
}
