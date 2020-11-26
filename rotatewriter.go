package utils

import (
	"log"
	"os"
	"strings"
	"time"
)

var stdoutLogger = log.New(os.Stdout, "", log.Ltime|log.Ldate|log.Lshortfile)

//RotateWriter 一个自动滚动文件的写入器
type RotateWriter struct {
	file         *os.File
	filename     string
	filepath     string
	lastCreateOn time.Time
	CreateOn     time.Time
}

//Write 数据写入文件
func (writer *RotateWriter) Write(p []byte) (n int, err error) {
	defer func() {
		if err := recover(); err != nil {
			stdoutLogger.Println(err)
		}
	}()
	writer.checkFile()
	n, err = writer.file.Write(p)
	if err != nil {
		stdoutLogger.Println(err)
	}
	return
}

//checkFile 检查文件 如果最后创建日期的时间是不是当天，创建新文件
//
func (writer *RotateWriter) checkFile() {
	ny, nm, nd := time.Now().Date()
	y, m, d := writer.lastCreateOn.Date()
	if ny == y && nm == m && nd == d && writer.file != nil {
		return
	}
	writer.file = createFile(writer.filepath, writer.filename)
}

//NewRotateWriter 创建一个新的按天滚动的RotateWriter
func NewRotateWriter(path, name string) *RotateWriter {
	//创建文件
	file := createFile(path, name)
	logger := &RotateWriter{
		CreateOn:     time.Now(),
		lastCreateOn: time.Now(),
		file:         file,
		filename:     name,
		filepath:     path,
	}
	return logger
}

//createFile 按照日期创建文件
func createFile(path, name string) *os.File {
	filename := name + `_` + time.Now().Format("2006_01_02") + `.log`
	var dir string
	path = strings.TrimSpace(path)

	if len(path) == 0 {
		dir = `/logs/`
	} else if strings.HasSuffix(path, `/`) {
		dir = `logs/` + path
	} else {
		dir = `/logs/` + path + `/`
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0666)
		if err != nil {
			stdoutLogger.Println(err.Error())
		}
	}
	f, err := os.OpenFile(dir+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		stdoutLogger.Println(err)
		return os.Stdout
	}
	return f
}
