package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"orm/conf"
	"orm/db"
	"orm/generate"
	"os"
)

var configFile = flag.String("f", "conf.json", "The config file")

func main() {
	flag.Parse()
	load(*configFile, &conf.Conf)
	db.NewDB(conf.Conf.Db.DbName)
	generate.Genertate(conf.Conf.Db.Tables...)
}

func load(configString string, config interface{}) error {
	data, err := readFile(configString)
	if err != nil {
		return err
	}
	json.Unmarshal(data, config)
	return err
}

func readFile(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}
	return fd, nil
}
