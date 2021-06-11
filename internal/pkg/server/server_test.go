package server_test

import (
	"errors"
	"go-template/internal/pkg/server"
	"testing"
	"time"
)

type Server struct {
	name      string
	stop      bool
	emitError bool
}

func (s Server) Name() string {
	return s.name
}

func (s *Server) Serve() error {
	for {
		if s.stop {
			if s.emitError {
				return errors.New("emit error")
			}
			return nil
		}
		time.Sleep(time.Second)
	}
}

func (s *Server) Stop() {
	s.stop = true
}

func TestServe(t *testing.T) {

	t.Run("error is nil", func(t *testing.T) {
		srv1 := &Server{"srv1", false, true}
		srv2 := &Server{"srv2", false, false}

		go func() {
			time.Sleep(time.Second)
			srv1.Stop()
			time.Sleep(time.Second)
			srv2.Stop()
		}()

		err := server.Serve(srv1, srv2)
		if err == nil {
			t.Fatalf("error should not be nil")
		}

	})

	t.Run("error is not nil", func(t *testing.T) {
		srv1 := &Server{"srv1", false, true}
		srv2 := &Server{"srv2", false, false}

		go func() {
			time.Sleep(time.Second)
			srv2.Stop()
			time.Sleep(time.Second)
			srv1.Stop()
		}()

		err := server.Serve(srv1, srv2)
		if err == nil {
			t.Fatalf("error should be nil")
		}

	})

}
