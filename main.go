package main

import (
	"flag"
    "github.com/thewayma/suricata_checker/g"
    "github.com/thewayma/suricata_checker/rpc"
    _"github.com/thewayma/suricata_checker/check"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()

	g.ParseConfig(*cfg)
	g.InitRedisConnPool()
	g.InitHbsClient()

    go rpc.Start()
    /*
	go cron.SyncStrategies()
	go cron.CleanStale()
    */

	select {}
}
