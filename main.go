package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Unknwon/goconfig"
	"github.com/hudl/fargo"

	service "github.com/leeningli/backing-fulfillment/service"
)

var (
	eureka_url string
	port       string
)

func main() {
	intPort, _ := strconv.Atoi(port)

	server := service.NewServer()
	go server.Run(":" + port)
	fmt.Println("12334")
	c := fargo.NewConn(eureka_url + "/eureka")
	i := fargo.Instance{
		HostName:         "lee",
		Port:             intPort,
		App:              "BackingFulfillmentApp",
		IPAddr:           "127.0.0.1",
		VipAddress:       "127.0.0.1",
		SecureVipAddress: "127.0.0.1",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.UP,
	}
	c.RegisterInstance(&i)
	f, _ := c.GetApps()
	for key, theApp := range f {
		fmt.Println("App:", key, "First Host Name:", theApp.Instances[0].HostName)
	}
	app, _ := c.GetApp("BackingFulfillmentApp")
	fmt.Printf("my app info is %v\n", app)
}

func init() {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	eureka_url, err = cfg.GetValue("eureka", "url")
	if err != nil {
		fmt.Println(err)
		eureka_url = "http://127.0.0.1:8080" //docker run -d -it -p 8080:8761 springcloud/eureka-------http://127.0.0.1:8080/eureka/apps/TestEurekaApp
	}
	port, err = cfg.GetValue("main", "port")
	if err != nil {
		fmt.Println(err)
		port = "3001"
	}
}
