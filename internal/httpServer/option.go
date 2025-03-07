package httpServer

import (
	"time"
)

type Options func(*Server)

func OptionSet(host, port string, sTimeout time.Duration) Options {
	return func(s *Server) {
		Addr(host, port)(s)
		ShutdownTimeout(sTimeout)(s)
	}
}

func Addr(host, port string) Options {
	return func(s *Server) {
		s.addr = ":" + port //net.JoinHostPort(host, port)
	}
}

func ShutdownTimeout(sTimeout time.Duration) Options {
	return func(s *Server) {
		s.shutdownTimeout = sTimeout
	}
}
