package main

import (
	//"os"
	//"fmt"
	"flag"

    "github.com/thewayma/suricata_checker/g"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()
	g.ParseConfig(*cfg)

	g.InitRedisConnPool()
    /*
	g.InitHbsClient()

	store.InitHistoryBigMap()

	go rpc.Start()
	go cron.SyncStrategies()
	go cron.CleanStale()
    */

	select {}
}
