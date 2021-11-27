package logger

import "log"

// TODO: logファイルに残すようにする
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}
