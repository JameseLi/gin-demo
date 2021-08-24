package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Server struct {
	Env string `yaml:"env"`
}

func NewServer() *Server {
	return &Server{}
}

var filePath string

func init() {
	flag.StringVar(&filePath, "configs", "/Users/lizhengjun/gin-demo/config-gin/configs/test.yaml", "config 文件路径地址")
}

func main() {
	flag.Parse()
	fmt.Printf("filePath: %s\n", filePath)

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//初始化配置文件
	conf := NewServer()
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/config", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"env": conf.Env,
		})
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
