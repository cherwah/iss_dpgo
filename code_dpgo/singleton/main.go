package main

import (
	"singleton/logger"
	"strconv"
	"time"
)

func main() {

	// starting thread to simulate database log
	go db_log()

	// starting thread to simulate network log
	go net_log()

	// starting thread to simulate io log
	go io_log()

	// starting thread to simulate error log
	go err_log()

	time.Sleep(10 * time.Second)
}

// database log
func db_log() {
	log := logger.GetInstance()

	for i := 1; i <= 50; i++ {
		log.Log("database log_id: " + strconv.Itoa(i))
	}
}

// network log
func net_log() {
	log := logger.GetInstance()

	for i := 1; i <= 50; i++ {
		log.Log("network log_id: " + strconv.Itoa(i))
	}
}

// file-io log
func io_log() {
	log := logger.GetInstance()

	for i := 1; i <= 50; i++ {
		log.Log("io log_id: " + strconv.Itoa(i))
	}
}

// error log
func err_log() {
	log := logger.GetInstance()

	for i := 1; i <= 50; i++ {
		log.Log("error log_id: " + strconv.Itoa(i))
	}
}
