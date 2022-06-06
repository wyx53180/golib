package golog

//package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type logger struct {
	log_level     int
	log_file_name string
	log_file      *os.File
	log_file_num  int
}

var mu sync.Mutex

const (
	DEBUG          int               = iota //0
	INFO                                    //1
	WARNING                                 //2
	ERROR                                   //3
	one_file_size  = 5 * 1024 * 1024        // 5M
	max_file_count = 5
)

func str_to_log_level(str string) int {
	switch strings.ToLower(str) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	default:
		panic(errors.New("unknown log level"))
	}
}

func NewLogger(level, file_name string) *logger {
	log_level := str_to_log_level(level)
	file, err := os.OpenFile(file_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return &logger{
		log_level:     log_level,
		log_file:      file,
		log_file_num:  0,
		log_file_name: file_name,
	}
}
func (l *logger) judge_info_ouput(level int) bool {
	return level >= l.log_level
}

func process_data_message(data interface{}, tag string) string {
	// 获取代码运行信息,文件，行号，方法名
	pc, path, line, ok := runtime.Caller(2)
	if !ok {
		panic(ok)
	}
	func_name := runtime.FuncForPC(pc).Name()
	path = filepath.Base(path)
	// 处理时间
	now := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s [%s][%s, %s, %d]:%v\n", now, tag, path, func_name, line, data)
}

func (l *logger) check_file_size() {
	fileinfo, err := os.Stat(l.log_file_name)
	if err != nil {
		panic(err)
	}
	if fileinfo.Size() > one_file_size {
		// 超过size，重新开文件
		l.log_file_num += 1
		if l.log_file_num > max_file_count {
			// 文件数，重置为1
			l.log_file_num = 1
		}
		file_name := fmt.Sprintf("%s.%d", l.log_file_name, l.log_file_num)
		file, err := os.OpenFile(file_name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		// 关闭之前的日志文件
		l.log_file.Close()
		l.log_file = file
	}
}

func (l *logger) write_to_file(data string) {
	mu.Lock()
	defer mu.Unlock()
	l.check_file_size()
	l.log_file.WriteString(data)
}

func (l *logger) Info(data ...interface{}) {
	if l.judge_info_ouput(INFO) {
		sdata := process_data_message(data, "INFO")
		l.write_to_file(sdata)
	}
}

func (l *logger) Warning(data ...interface{}) {
	if l.judge_info_ouput(WARNING) {
		sdata := process_data_message(data, "WARNING")
		l.write_to_file(sdata)
	}
}
func (l *logger) Debug(data ...interface{}) {
	if l.judge_info_ouput(DEBUG) {
		sdata := process_data_message(data, "DEBUG")
		l.write_to_file(sdata)
	}
}
func (l *logger) Error(data ...interface{}) {
	if l.judge_info_ouput(ERROR) {
		sdata := process_data_message(data, "ERROR")
		l.write_to_file(sdata)
	}
}

func (l *logger) Close() {
	l.log_file.Close()
}

/*
func main() {
	l := NewLogger("debug", "my.log")
	l.Info("info")
	l.Debug("debug")
	l.Error("adf", 123, "eror")
}
*/
