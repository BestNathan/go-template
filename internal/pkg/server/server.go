package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type IServer interface {
	Name() string
	Serve() error
}

type Quit struct {
	name string
	err  error
}

func Serve(srvs ...IServer) error {
	qCh := make(chan *Quit)
	sCh := make(chan os.Signal)

	for _, srv := range srvs {
		go func(s IServer) {
			if err := s.Serve(); err != nil {
				qCh <- &Quit{s.Name(), err}
				return
			}
		}(srv)
	}

	signal.Notify(sCh, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case q := <-qCh:
			fmt.Printf("%s occurs an error: %s\n", q.name, q.err.Error())
			return q.err
		case s := <-sCh:
			fmt.Printf("receive signal: %s\n", s.String())
			return nil
		}
	}
}
