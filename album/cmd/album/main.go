package main

import (
	"flag"
	"fmt"
	conf2 "gin-demo/album/internal/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var confPath string

func init() {
	flag.StringVar(&confPath, "conf", "/Users/lizhengjun/gin-demo/album/configs/dev.yaml", "配置文件地址")
}

func main() {
	flag.Parse()
	yamlFile,err := ioutil.ReadFile(confPath)

	if err != nil {
		log.Fatal(err)
	}
	conf := conf2.NewBootStrap()

	err = yaml.Unmarshal(yamlFile,conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf.Jwt.Secret)
	fmt.Println(conf.Server.Http.Port)
	fmt.Println(conf.Server.Http.ReadTimeout)
}
