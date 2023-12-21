package utils

import "log"

const (
	InfoColor = "\033[1;36m%s\033[0m"
)

func LogInfo(msg string) {
	log.Printf(InfoColor, msg)
}
