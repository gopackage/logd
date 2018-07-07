package main

import "github.com/gopackage/logd"

func main() {
	server := &logd.Server{}
	server.Start()
}
