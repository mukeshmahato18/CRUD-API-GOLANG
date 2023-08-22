package main

import (
	"fmt"
	"log"
	"net"

	cache "github.com/mukeshmahato18/cache/Cache"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts

	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}
	log.Printf("server starting on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read error %s\n", err)
			break
		}
		go s.handleCommand(conn, buf[:n])
	}
}

func(s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	rawStr := string(rawCmd)
	parts := strings.Split(rawStr, " ")
}

func(s *Server) handleSetCmd(conn net.Conn, ){}
