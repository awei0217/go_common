package mysql_binlog

import (
	"context"
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

func Start() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "47.94.195.36",
		Port:     3306,
		User:     "root",
		Password: "root",
	}
	syncer := replication.NewBinlogSyncer(cfg)
	streamer, err := syncer.StartSync(mysql.Position{"mysql_binlog.000001", 0})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		ev, err := streamer.GetEvent(context.Background())
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(ev.RawData))
	}
}
