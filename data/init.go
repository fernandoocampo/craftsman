package data

import (
	"fmt"
	"os"

	"github.com/fernandoocampo/craftsman/util"
	"github.com/sirupsen/logrus"
)

var log *util.LogHandle

func init() {
	var err error
	log, err = util.NewLogger(util.Options{LogLevel: "info", LogFormat: "text", LogFields: logrus.Fields{"pkg": "dao", "srv": "user"}})
	if err != nil {
		fmt.Printf("cant load logger: %v", err)
		os.Exit(1)
	}
}
