package main

import "github.com/gfelixc/maritime/port/ui/server"

func main() {
	err := server.BootGRPCServer()
	if err != nil {
		panic(err)
	}
}
