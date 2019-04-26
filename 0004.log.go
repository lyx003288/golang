package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

const (
	//LogDir log dir
	LogDir           string = "./log/"
	debugLevel       string = "Debug"
	infoLevel        string = "Info"
	errorLevel       string = "Error"
	maxLogSize       int64  = (10 * 1024 * 1024)
	maxLogFileNumber int    = 10
)

var (
	//DebugLog debug log type
	debugLog *LogType
	//InfoLog info log type
	infoLog *LogType
	//ErrorLog error log type
	errorLog *LogType
)

//LogType log type
type LogType struct {
	file   *os.File
	size   int64
	logger *log.Logger
}

func init() {
	if isExist, _ := PathExists(LogDir); !isExist {
		os.Mkdir(LogDir, os.ModePerm)
	}

	initDebugLog()
	initInfoLog()
	initErrorLog()
}

func initDebugLog() {
	debugFile, debugLogger := generateLogger(debugLevel)
	debugLog = &LogType{file: debugFile, size: 0, logger: debugLogger}
}

func initInfoLog() {
	infoFile, infoLogger := generateLogger(infoLevel)
	infoLog = &LogType{file: infoFile, size: 0, logger: infoLogger}
}

func initErrorLog() {
	errorFile, errorLogger := generateLogger(errorLevel)
	errorLog = &LogType{file: errorFile, size: 0, logger: errorLogger}
}

func generateLogger(level string) (*os.File, *log.Logger) {
	logFile, err := os.OpenFile(LogDir+level+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open ", level, " log file error: ", err)
	}

	multiWrite := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWrite, "["+level+"] ", log.LstdFlags|log.Lshortfile)
	return logFile, logger
}

func renameLogFile(logName string) {
	oldestIndex := maxLogFileNumber - 1
	oldLogName := logName + "." + strconv.Itoa(oldestIndex)
	isExist, _ := PathExists(oldLogName)
	if isExist {
		os.Remove(oldLogName)
	}

	for i := maxLogFileNumber - 2; i > 0; i-- {
		oldLogName = logName + "." + strconv.Itoa(i)
		isExist, _ := PathExists(oldLogName)
		if isExist {
			newLogName := logName + "." + strconv.Itoa(i+1)
			os.Rename(oldLogName, newLogName)
		}
	}

	os.Rename(logName, logName+".1")
}

//Debug debug method
func Debug(format string, v ...interface{}) {
	doLog(debugLevel, format, v...)
}

//Info debug method
func Info(format string, v ...interface{}) {
	doLog(infoLevel, format, v...)
}

//Error debug method
func Error(format string, v ...interface{}) {
	doLog(errorLevel, format, v...)
}

func doLog(level string, format string, v ...interface{}) {
	var logType *LogType
	switch level {
	case debugLevel:
		logType = debugLog
	case infoLevel:
		logType = infoLog
	case errorLevel:
		logType = errorLog
	}
	fileInfo, err := logType.file.Stat()
	if err != nil {
		log.Panic("log "+level+" stat error, ", err)
	}

	size := fileInfo.Size()
	if size > maxLogSize {
		logType.file.Close()

		logName := fileInfo.Name()
		renameLogFile(path.Join(LogDir, logName))

		logFile, logger := generateLogger(errorLevel)
		logType.file = logFile
		logType.size = 0
		logType.logger = logger
	}

	logType.logger.Output(3, fmt.Sprintf(format, v...))
}
