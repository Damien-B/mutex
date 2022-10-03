package main

import (
	"fmt"
	"sync"
	"time"
)

type access interface {
	setContent(value string)
	getContent() string
}

type objectWithRWMutex struct {
	mu      sync.RWMutex
	content string
}

func (object *objectWithRWMutex) setContent(value string) {
	object.mu.Lock()
	defer object.mu.Unlock()
	object.content = value
	fmt.Printf("%d | CONTENT SET · %s\n", time.Now().Nanosecond(), value)

}

func (object *objectWithRWMutex) getContent() string {
	object.mu.RLock()
	defer object.mu.RUnlock()
	value := object.content
	fmt.Printf("%d | CONTENT GET %s\n", time.Now().Nanosecond(), value)
	return value
}
