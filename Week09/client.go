package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var ch = make(chan int)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("server is not starting")
		return
	}

	defer conn.Close()
	go func(conn net.Conn) {
		for {
			var data = make([]byte, 128)
			_, err := conn.Read(data)
			if err != nil {
				log.Printf("err:%v \n", err)
			}
			fmt.Println("Reply:", string(data))
		}
	}(conn)
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err == nil {
			fmt.Printf("client send：%s", input)
		}
		//将从输入中读取的内容写入到连接中
		b := []byte(input)
		conn.Write(b)

		select {
		case <-ch:
			fmt.Println("server error，please reconnecting")
			return
		default:
			//不加default的话，那么<-ch会阻塞for，下一个输入就没法进行
		}
	}
}
