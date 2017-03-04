package main

import (
	"log"
	"os"
	"sort"

	"github.com/BurntSushi/toml"
)

type MainCfgDef struct {
	InlcudeDir   []string
	ExcludeDir   []string
	LogFile      string
	IntervalTime int
}

var (
	MainCfg MainCfgDef
)

func init() {

	if _, err := toml.DecodeFile("config.toml", &MainCfg); err != nil {
		log.Fatalln("load config.toml error, ", err)
	}

	sort.Strings(MainCfg.InlcudeDir)
	sort.Strings(MainCfg.ExcludeDir)

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	if len(MainCfg.LogFile) != 0 {
		file, err := os.OpenFile(MainCfg.LogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		}
	}
}

func main() {
	addFileMonitor()

	log.Printf("server starting(total:%d)...\n", MmonitorDirNum)

	fileMonitorLoop()

	log.Println("server stopped")
}
