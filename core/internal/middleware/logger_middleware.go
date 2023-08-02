package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type LoggerMiddleware struct {
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go writelog(r) //记录日志
		next(w, r)
	}
}

func writelog(r *http.Request) {
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	msg := fmt.Sprintf("[Log] [%v] [%v] [%v] [%v] \n", date, r.RemoteAddr, r.Method, r.URL.Path)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println("open log file error:", err)
	}
	_, err = file.WriteString(msg)
	if err != nil {
		fmt.Println("record log error:", err)
	}

}
