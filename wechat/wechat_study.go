package wechat

import (
	"github.com/qianlnk/gobot"
)

func WeChat() {
	cfg := gobot.Load()
	rebot, err := gobot.NewWecat(cfg)
	if err != nil {
		panic(err)
	}

	rebot.Start()
}
