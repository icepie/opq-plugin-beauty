package main

import (
	"log"
	"sync"
	"time"

	"github.com/icepie/xiaoice-beauty/client"

	"github.com/mcoo/OPQBot"
)

var (
	Bot OPQBot.BotManager
	IB  client.IceBeauty
)

// 启动 bot
func start(qq int64, opqUrl string) {

	Bot = OPQBot.NewBotManager(qq, opqUrl)
	err := Bot.Start()
	if err != nil {
		log.Println(err.Error())
		log.Println("连接失败！！")
		time.Sleep(5 * time.Second)
		log.Println("尝试重新连接")
		start(qq, opqUrl)
	}
	defer Bot.Stop()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// 群消息事件
	err = Bot.AddEvent(OPQBot.EventNameOnGroupMessage, groupMsgHandle)
	if err != nil {
		log.Println(err.Error())
	}
	// 好友消息事件
	err = Bot.AddEvent(OPQBot.EventNameOnFriendMessage, friendMsgHandle)
	if err != nil {
		log.Println(err.Error())
	}
	err = Bot.AddEvent(OPQBot.EventNameOnGroupShut, func(botQQ int64, packet OPQBot.GroupShutPack) {
		log.Println(botQQ, packet)
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = Bot.AddEvent(OPQBot.EventNameOnConnected, func() {
		log.Println("连接成功！！！")
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = Bot.AddEvent(OPQBot.EventNameOnDisconnected, func() {
		log.Println("连接断开！！")
		time.Sleep(5 * time.Second)
		log.Println("尝试重新连接")
		start(qq, opqUrl)
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = Bot.AddEvent(OPQBot.EventNameOnOther, func(botQQ int64, e interface{}) {
		log.Println(e)
	})
	if err != nil {
		log.Println(err.Error())
	}
	wg.Wait()

}

func main() {
	// 创建小冰颜值鉴定代理对象
	var err error
	IB, err = client.New()
	if err != nil {
		log.Println("error: ", err)
	}

	// 启动 bot
	start(0, "http://127.0.0.1:8881")
}
