package logd

import (
	"encoding/base64"
	"io"
	"net"

	"github.com/apex/log"
	"github.com/rogpeppe/fastuuid"
)

// Server is the core logd.
type Server struct {
	gen  *fastuuid.Generator
	conn *net.UDPConn
}

// New creates a new server.
func New() *Server {
	return &Server{
		gen: fastuuid.MustNewGenerator(),
	}
}

// Start the server - blocks the thread.
func (s *Server) Start() error {
	addr, err := net.ResolveUDPAddr("udp", ":9044")
	if err != nil {
		return err
	}
	log.WithField("addr", addr).Info("Dialing")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	log.Info("Connected")

	s.conn = conn
	buffer := make([]byte, 512)
	for {
		size, remote, err := conn.ReadFrom(buffer)
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
		// TODO process the data
		log.WithFields(log.Fields{"size": size, "remote": remote, "data": string(buffer[0:size])}).Info("log")
	}
}

// Stop terminates the server (not thread safe).
func (s *Server) Stop() {
	// Not thread safe but we don't anticipate multiple calls to Stop.
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
}

// UUID generates a 24-byte UUID.
func (s *Server) UUID() string {
	uuid := s.gen.Next()
	key := base64.StdEncoding.EncodeToString(uuid[:])

	return key
}
