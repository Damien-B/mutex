package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var isTestingRWMutex bool = true
	var sut iObjectWithContent
	if isTestingRWMutex {
		sut = &objectWithRWMutex{content: "0"}
	} else {
		sut = &objectWithMutex{content: "0"}
	}

	go appendObjectContent(sut)
	go appendObjectContent(sut)
	go appendObjectContent(sut)
	go appendObjectContent(sut)
	go appendObjectContent(sut)
	go appendObjectContent(sut)

	<-interrupt
}

func appendObjectContent(object iObjectWithContent) {
	newValue := object.getContent() + "-"
	object.setContent(newValue)
}
