package rpc

import (
	"log"
	"net"
	"net/rpc"
	"github.com/thewayma/suricata_checker/g"
	"github.com/thewayma/suricata_checker/check"
)

type Judge struct{}

func (this *Judge) Send(items []*g.JudgeItem, resp *g.SimpleRpcResponse) error {
	remain := g.Config().Remain
	now := time.Now().Unix()

	for _, item := range items {
		pk := item.PrimaryKey()
		check.HistoryBigMap[pk[0:2]].PushFrontAndMaintain(pk, item, remain, now)
	}
	return nil
}

func Start() {
	if !g.Config().Rpc.Enabled {
		return
	}
	addr := g.Config().Rpc.Listen
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail: %s", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
	} else {
		log.Println("rpc listening", addr)
	}

	rpc.Register(new(Judge))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept occur error: %s", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
