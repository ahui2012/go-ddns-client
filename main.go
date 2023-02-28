package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ahui2012/go-ddns-client/config"
	"github.com/ahui2012/go-ddns-client/ddns"
	"github.com/ahui2012/go-ddns-client/pubip"
)

var WorkingPath string

func init() {
	exeFilePath, err := os.Executable()
	if err != nil {
		log.Panicln(err)
	}

	WorkingPath = filepath.Dir(exeFilePath)
	logFilePath := filepath.Join(WorkingPath, "ddns.log")
	fmt.Println("log file path:", logFilePath)
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(f)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("catch a error:", err)
		}
	}()

	fmt.Println("running at", WorkingPath)
	log.Println("started")
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatalln("can load config from config.json:", err)
		return
	}
	interval := time.Duration(cfg.UpdateInterval) * time.Second

	ddns.Init(cfg.Domains)
	for {
		update(cfg)
		time.Sleep(interval)
	}
}

func update(cfg *config.AppConfig) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("catch a error:", err)
		}
	}()

	publicIP := pubip.GetPublicIP(cfg.PubIPUrls)
	if publicIP == "" {
		log.Fatalln("can not get public ip")
		return
	}

	ddns.Update(publicIP)
}
