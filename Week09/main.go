package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type rwConn struct {
	readMu    sync.Mutex // for TestHandlerBodyClose
	readBuf   bytes.Buffer
	writeBuf  bytes.Buffer
	closeFunc func() error // called if non-nil
	closec    chan bool    // else, if non-nil, send value to it on close
}

func (c *rwConn) Close() error {
	if c.closeFunc != nil {
		return c.closeFunc()
	}
	select {
	case c.closec <- true:
	default:
	}
	return nil
}

func (c *rwConn) Read(b []byte) (int, error) {
	c.readMu.Lock()
	defer c.readMu.Unlock()
	return c.readBuf.Read(b)
}

func (c *rwConn) Write(b []byte) (int, error) {
	return c.writeBuf.Write(b)
}

func main() {
	ctx := context.Background()
	var output bytes.Buffer
	var input bytes.Buffer
	rw := &rwConn{
		readBuf:  input,
		writeBuf: output,
		closec:   make(chan bool, 1),
	}

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("err:%v \n", err)
		return
	}
	fmt.Printf("start listen \n ")
	for {
		var msg = make(chan []byte, 10)
		conn, err := listen.Accept()
		fmt.Printf("accapt conn:%v \n ", conn.RemoteAddr())
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		go rw.QueryRead(ctx, conn, msg)
		go rw.QueryWrite(ctx, conn, msg)
	}

}

//创建grpcServer
func (rw *rwConn) QueryRead(ctx context.Context, conn net.Conn, msg chan []byte) {

	var data = make([]byte, 128)
	for {
		n, err := conn.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		fmt.Println("n:", n)
		msg <- data[:n]
	}
	fmt.Printf("rw:write:%v", rw.writeBuf.String())
	return
}

func (rw *rwConn) QueryWrite(ctx context.Context, conn net.Conn, msg chan []byte) {
	for {
		data, ok := <-msg
		if !ok {
			break
		}
		fmt.Println("input:", string(data))
		n, err := conn.Write(data)
		fmt.Println("output:", n)
		if err != nil {
			if err != io.EOF {
				fmt.Println("err", err)
			}
		}
		data = data[:0]
		rw.writeBuf.Write(data)
	}
}
