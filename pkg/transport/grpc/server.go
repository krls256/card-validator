package grpc

import (
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type ServiceRegistration[SVC any] func(grpc.ServiceRegistrar, SVC)

func NewServer[SVC any](cfg Config, svc SVC, sr ServiceRegistration[SVC]) *Server[SVC] {
	return &Server[SVC]{
		cfg: cfg,
		svc: svc,
		sr:  sr,
	}
}

type Server[SVC any] struct {
	cfg Config
	svc SVC
	sr  ServiceRegistration[SVC]

	srv *grpc.Server
}

func (s *Server[SVC]) RunAsync() {
	go func() {
		if err := s.Run(); err != nil {
			slog.Error(err.Error())
		}
	}()
}

func (s *Server[SVC]) Run() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", s.cfg.DNS())
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	slog.Info("grpc listening", "address", tcpAddr.String())

	s.srv = grpc.NewServer()

	s.sr(s.srv, s.svc)

	if err = s.srv.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (s *Server[SVC]) Shutdown() {
	s.srv.GracefulStop()
}
