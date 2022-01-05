package lg

import "fmt"

type Appender interface {
	append(msg string)
}

type ConsoleAppender struct {
	Appender
	pattern string
}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

func (c *ConsoleAppender) append(msg string) {
	fmt.Println(msg)
}

type RollingFileAppender struct {
	Appender
	lgpath   string
	filename string
	pattern  string
}

func NewRollingFileAppender(lgpath, filename string) *RollingFileAppender {
	return &RollingFileAppender{}
}

func (r *RollingFileAppender) append(msg string) {

}
