package logger

import (
	"os"
	"sync"
)

var singleLogger *logger
var mutex sync.Mutex

type logger struct {
	file *os.File
}

func GetInstance() *logger {
	mutex.Lock()         // only one thread can come in at a time
	defer mutex.Unlock() // unlock after end of function

	if singleLogger == nil {
		singleLogger = &logger{}
		singleLogger.file, _ = os.Create("log.txt")
	}

	return singleLogger
}

func (l *logger) Log(data string) {
	l.file.WriteString(data + "\n")
}
