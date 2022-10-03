package main

import (
	"fmt"
	"sync"
	"time"
)

type objectWithMutex struct {
	mu      sync.Mutex
	content string
}

func (object *objectWithMutex) setContent(value string) {
	object.mu.Lock()
	defer object.mu.Unlock()
	object.content = value
	fmt.Printf("%d | CONTENT SET · %s\n", time.Now().Nanosecond(), value)

}

func (object *objectWithMutex) getContent() string {
	object.mu.Lock()
	defer object.mu.Unlock()
	value := object.content
	fmt.Printf("%d | CONTENT GET %s\n", time.Now().Nanosecond(), value)
	return value
}
