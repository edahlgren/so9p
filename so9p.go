package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

var servermap map[fid]*sfid
var clientfid fid
func main() {

	if os.Args[1] == "s" {

		servermap = make(map[fid]*sfid, 128)
		S := new(So9ps)
		S.Fs.Name = "/"
		rpc.Register(S)
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal(err)
		}
		rpc.Accept(l)
	} else {
		var client so9pc
		var err error
		client.Client, err = rpc.Dial("tcp", "localhost"+":1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		fi, err := client.attach("/", 1)
		if err != nil {
			log.Fatal("attach", err)
		}
		newfid, fi, err := client.walk(1, "etc")
		if err != nil {
			log.Fatal("walk", err)
		}
		fmt.Printf("newfid %v fi %v err %v\n", newfid, fi, err)
	}

}
