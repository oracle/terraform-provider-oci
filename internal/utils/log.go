// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"syscall"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/fatih/color"
)

var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
)

//TFProviderLogger interface for logging in the Terraform Provider
type TFProviderLogger interface {
	//LogLevel returns the log level of TFProviderLogger
	LogLevel() int

	//Log logs v with the provided format if the current log level is loglevel
	Log(logLevel int, format string, v ...interface{}) error
}

//NONE no logging messages
const NONE = 0

//INFO minimal logging messages
const INFO = 1

//DEBUG debug logging messages
const DEBUG = 2

//defaultTFProviderLogger the default implementation of the TFProviderLogger
type defaultTFProviderLogger struct {
	currentLoggingLevel int
	debugLogger         *log.Logger
	infoLogger          *log.Logger
	nullLogger          *log.Logger
}

var defaultTFLogger TFProviderLogger
var loggerLock sync.Mutex

//initializes the defaultTFProviderLogger as default Logger
func init() {
	l, _ := NewTFProviderLogger()
	SetTFProviderLogger(l)
}

//SetTFProviderLogger sets the defaultTFLogger to logger
func SetTFProviderLogger(logger TFProviderLogger) {
	loggerLock.Lock()
	defaultTFLogger = logger
	loggerLock.Unlock()
}

// NewTFProviderLogger creates a defaultTFProviderLogger
// The value of the "OCI_TF_LOG" environment variable controls the logging level.
// Default logging level is INFO "i"
func NewTFProviderLogger() (defaultTFProviderLogger, error) {
	logger := defaultTFProviderLogger{}

	logOutput := os.Stderr
	if logPath := os.Getenv(globalvar.EnvOCITFLogFile); logPath != "" {
		logOutput, _ = os.OpenFile(logPath, syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0666)
	}

	logger.currentLoggingLevel = NONE
	logger.debugLogger = log.New(logOutput, "DEBUG ", log.Ldate|log.Lmicroseconds)
	logger.infoLogger = log.New(logOutput, "INFO ", log.Ldate|log.Lmicroseconds)
	logger.nullLogger = log.New(ioutil.Discard, "", log.Ldate|log.Lmicroseconds)

	logLevel := GetEnvSettingWithDefault("TF_LOG", "i")

	switch strings.ToLower(logLevel) {
	case "null":
		logger.currentLoggingLevel = NONE
		break
	case "i", "info", "INFO":
		logger.currentLoggingLevel = INFO
		break
	case "d", "debug", "DEBUG":
		logger.currentLoggingLevel = DEBUG
		break
	default:
		logger.currentLoggingLevel = INFO
	}

	return logger, nil
}

func (l defaultTFProviderLogger) getLoggerForLevel(logLevel int) *log.Logger {
	if logLevel > l.currentLoggingLevel {
		return l.nullLogger
	}

	switch logLevel {
	case NONE:
		return l.nullLogger
	case INFO:
		return l.infoLogger
	case DEBUG:
		return l.debugLogger
	default:
		return l.nullLogger
	}
}

//LogLevel returns the current debug level
func (l defaultTFProviderLogger) LogLevel() int {
	return l.currentLoggingLevel
}

func (l defaultTFProviderLogger) Log(logLevel int, format string, v ...interface{}) error {
	logger := l.getLoggerForLevel(logLevel)
	_ = logger.Output(4, fmt.Sprintf(format, v...))
	return nil
}

// Log logs v
func Log(v ...interface{}) {
	_ = defaultTFLogger.Log(INFO, "%v", v...)
}

// Logln logs v appending a new line at the end
func Logln(v ...interface{}) {
	_ = defaultTFLogger.Log(INFO, "%v\n", v...)
}

// Logf logs v with the provided format
func Logf(format string, v ...interface{}) {
	_ = defaultTFLogger.Log(INFO, format, v...)
}

// Debug logs v if debug mode is set
func Debug(v ...interface{}) {
	m := fmt.Sprint(v...)
	_ = defaultTFLogger.Log(DEBUG, "%s", m)
}

// Debugln logs v appending a new line if debug mode is set
func Debugln(v ...interface{}) {
	m := fmt.Sprint(v...)
	_ = defaultTFLogger.Log(DEBUG, "%s\n", m)
}

// Debugf logs v with the provided format if debug mode is set
func Debugf(format string, v ...interface{}) {
	_ = defaultTFLogger.Log(DEBUG, format, v...)
}
