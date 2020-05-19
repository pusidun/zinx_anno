package main

import (
	"zinx_anno/znet"
)

func main() {
	s := znet.NewServer("test")
	s.Server()
}