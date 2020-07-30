package main

import (
	"flag"

	"pika/server/servers"
)

func main() {

	var addr string
	flag.StringVar(&addr, "addr", "9999", "服务器端口,默认9999")
	flag.Parse()

	addr = ":" + addr
	servers.Run(addr)
}
