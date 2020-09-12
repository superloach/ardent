package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var dir string

func init() {
	flag.StringVar(&dir, "d", "assets/", "Asset directory")
	flag.Parse()
}

func parseConfigFiles() (confs []config, err error) {
	err = filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) != ".yml" {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			var conf config
			if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
				return err
			}

			conf.filepath = strings.Replace(path, ".yml", ".png", 1)
			confs = append(confs, conf)

			return nil
		},
	)

	return
}

func main() {
	confs, err := parseConfigFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, conf := range confs {
		asset, err := conf.toAsset()
		if err != nil {
			log.Fatal(err)
		}

		d, err := asset.MarshalBinary()
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(
			strings.Replace(conf.filepath, ".png", ".asset", 1),
			d,
			0777,
		); err != nil {
			log.Fatal(err)
		}
	}
}
