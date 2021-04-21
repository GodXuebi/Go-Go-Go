package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")
	log.SetLevel(log.TraceLevel)
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	log.Fatal("Bye.")
	// 记完日志后会调用 panic()
	log.Panic("I'm bailing.")

}
