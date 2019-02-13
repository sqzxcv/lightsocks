package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sqzxcv/lightsocks"
	"github.com/sqzxcv/lightsocks/cmd"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	// config.ReadConfig()
	// config.SaveConfig()

	config.Password = "yshAwgIkaoBDJw2lq1yNUsZXu8vNPJP57LfvNmDphbXX23a6O4ZfnCAMBJ7TtL96LYPaeah15z8Gny9tjmgOZJDc8mFJT9huZoisVP+z9SNpiZ36JqFrB70xRhXqxK0LA1lMWnEswxPZAMdyR3zQ1n43EaAPlprMH1ESsKOM5lZnKh32JZcJc2y2p4cQ/ugpBRnuot3xQtRB5DjrRU6PmCEb7RY6WIt0lRczkq97NEuZm2Ndz+V9f7I5K/epvvjBcM4BvAiKPlXhgvDSsW+UKPNQyar8Hvs9MOJ4/RzfptX0Ta7AMkSEdy7guDVepCIYFEpIW8WB3mWRU2Ia0bnjCg=="
	config.ListenAddr = "127.0.0.1:4433"
	config.RemoteAddr = "104.237.141.191:4433"

	// 启动 local 端并监听
	lsLocal, err := lightsocks.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Println("sqzxcv读取密码错误,错入原因如下:")
		log.Fatalln(err)
	}
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("sqzxcv--使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
远程服务地址 remote：
%s
密码 password：
%s
	`, listenAddr, config.RemoteAddr, config.Password))
		log.Printf("sqzxcv--lightsocks-local:%s 启动成功 监听在 %s\n", version, listenAddr.String())
	}))
}
