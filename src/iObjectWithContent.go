package main

type iObjectWithContent interface {
	setContent(value string)
	getContent() string
	startTask()
	endTask()
}
