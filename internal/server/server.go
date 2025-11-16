package server

import (
	"fmt"
	"io"
	"net"

	"github.com/pixambi/pixam-protohackers/internal/config"
)

type Server struct {
	Config   *config.Config
	listener net.Listener
}

func New(cfg *config.Config) *Server {
	return &Server{
		Config: cfg,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	s.listener = ln

	fmt.Println("listening on port 10000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("accept: %w", err)
		}
		fmt.Println("connection from ", conn.RemoteAddr())
		go s.handle(conn)
	}
}

func (s *Server) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		fmt.Println("copy: ", err.Error())
	}
}
