package main

import (
	. "code.be.staff.com/staff/StaffGo/public/zap"
	"time"
)

func main() {
	logger := LoadConfiguration("./logconfig.json")
	defer Close()
	log := logger.Sugar()

	for {
		log.Error("error")
		log.Info("info")
		time.Sleep(5 * time.Second)
	}
}
