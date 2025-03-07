package httpServer

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Server struct {
	app             *fiber.App
	notify          chan error
	shutdownTimeout time.Duration
	addr            string
}

func New(router *fiber.App, opts ...Options) *Server {
	s := &Server{
		app:    router,
		notify: make(chan error, 1),
	}

	for _, o := range opts {
		o(s)
	}

	return s
}

func (s *Server) Start() error {
	err := s.app.Listen(s.addr)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.app.ShutdownWithContext(ctx)
}
