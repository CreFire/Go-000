package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var done = make(chan int)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
			done <- 1
		})
		s := NewServer(":8080",mux)
		go func() {
			err := s.Start()
			if err != nil {
				fmt.Printf(" err:%s \n ", err)
			}
		}()

		select {
		case <-done:
			fmt.Printf(" http done \n ")
			return s.Stop()
		case <-ctx.Done():
			fmt.Printf("signal-> http  done \n ")
			return errors.New("signal-> http  done \n ")
		}
	})
	group.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			fmt.Printf("signal done \n ")
			return errors.New("signal done \n ")
		case <-ctx.Done():
			fmt.Printf("http server -> signal done \n ")
			return errors.New("http server -> signal done \n ")
		}
	})

	// 捕获err
	fmt.Println("开始捕捉err")
	err := group.Wait()
	fmt.Println("=======", err)
}

//http服务
type httpServer struct {
	s       *http.Server
	handler http.Handler
	cxt     context.Context
}

func NewServer(address string, mux http.Handler) *httpServer {
	h := &httpServer{cxt: context.Background()}
	h.s = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	return h
}

func (h *httpServer) Start() error {
	fmt.Println("httpServer start")
	return h.s.ListenAndServe()
}

func (h *httpServer) Stop() error {
	_ = h.s.Shutdown(h.cxt)
	return fmt.Errorf("httpServer结束")
}