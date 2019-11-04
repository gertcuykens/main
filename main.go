package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"

	. "github.com/gertcuykens/main/hello"
	cakebaker "github.com/gertcuykens/module/v5"
)

func main() {
	fmt.Println(Hello())
	srv := rpc.NewServer()
	srv.Register(new(cakebaker.CakeBaker))

	go func(srv *rpc.Server) {
		l, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer l.Close()

		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			go func(conn net.Conn) {
				srv.ServeConn(conn)
				defer conn.Close()
			}(conn)
		}
	}(srv)

	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var reply string
	err = client.Call("CakeBaker.BakeIt", 5, &reply)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer client.Close()
	fmt.Printf("%s\n", reply)
}
