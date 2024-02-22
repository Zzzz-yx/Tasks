package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000") //监听tcp协议的端口
	if err != nil {
		log.Fatal(err) //如果错误会退出
	}

	for {
		//会阻塞，接收客户端的连接
		fmt.Println("accept before")
		conn, err := listener.Accept()
		fmt.Println("accept after", conn.RemoteAddr().String()) //获取远程地址
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	//发送当前时间给客户端
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05")) //格式化时间
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
