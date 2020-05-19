package znet

import (
	"fmt"
	"net"
	"zinx_anno/ziface"
)

type Server struct {
	IPVer string
	name  string
	addr  string
	port  int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP:%s, Port:%d is start", s.addr, s.port)
	
	go func() {
		conn, err := net.ResolveTCPAddr(s.IPVer, fmt.Sprintf("%s:%d", s.addr, s.port))
		if err != nil {
			fmt.Printf("resolve tcp addr error:", err)
			return
		}

		listen, err := net.ListenTCP(s.IPVer, conn)
		if err != nil {
			fmt.Printf("listen error:", err)
			return
		}

		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buff err", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buff err", err)
						continue
					}
				}
			}()
		}
	}()
			
}

func (s *Server) Stop() {
	fmt.Printf("[Start] Server Listener at IP:%s, Port:%d is stop", s.addr, s.port)
}

func (s *Server) Server() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		IPVer: "tcp4",
		name:  name,
		addr:  "0.0.0.0",
		port:  7777,
	}
	return s
}


