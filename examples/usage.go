package main

import "github.com/jurka/llogger"

var log = llogger.NewStdLogger(llogger.INFO, "prefix")

func main() {
	log.Info("This is info")
	log.Fatal("This is fatal")
	log.Debug("This is debug") // you woun't see this
	log.Error("This is error")
	log.Warn("This is warn")

}
