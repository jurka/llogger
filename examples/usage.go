package main

import "github.com/jurka/llogger"

var log = llogger.NewStdLogger(llogger.INFO, "prefix")

func main() {
	log.Info("This is INFO")
	log.Print("`Print` is also INFO")
	log.Fatal("This is FATAL")
	log.Debug("This is DEBUG") // you woun't see this
	log.Error("This is ERROR")
	log.Warn("This is WARN")
}
