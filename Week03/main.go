package main

import (
	"context"
	"errors"
	"fmt"
	"go-project-layout/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.NewServer()
	// http server
	g.Go(func() error {
		fmt.Println("http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			svr.Shutdown(context.TODO())
		}()
		return svr.Start()
	})
}
