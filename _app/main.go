package main

import "github.com/gopackage/logd"

func main() {
	server := &logd.Server{}
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
