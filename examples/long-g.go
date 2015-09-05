// Example for long running G

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jurka/llogger"
)

func longTask() {
	log := llogger.NewStdLogger(llogger.INFO, fmt.Sprint("longTask", "-", rand.Int31n(89999)+10000, " "))
	for {
		log.Info("This is sample message about progress")
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start tasks:")

	go longTask()
	go longTask()
	go longTask()

	time.Sleep(10 * time.Second)
}
