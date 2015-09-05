package llogger

import (
	"io"
	"log"
	"os"
)

const (
	DISABLED = iota
	FATAL    = iota
	ERROR    = iota
	WARN     = iota
	INFO     = iota
	DEBUG    = iota
)

type Logger struct {
	level  int
	logger *log.Logger
}

func NewLogger(out io.Writer, level int, prefix string) *Logger {
	if level != DISABLED || level != FATAL || level != ERROR || level != WARN || level != INFO || level != DEBUG {
		log.Fatal("Wrong level provided, user llogger.[DISABLED|FATAL|ERROR|WARN|INFO|DEBUG]")
	}
	return &Logger{level: level, logger: log.New(out, prefix, log.LstdFlags)}
}

func NewStdLogger(level int, prefix string) *Logger {
	return &Logger{level: level, logger: log.New(os.Stderr, prefix, log.LstdFlags)}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.output(FATAL, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.output(ERROR, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(WARN, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.output(INFO, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.output(INFO, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(DEBUG, v...)
}

func (l *Logger) Fatalf(f string, v ...interface{}) {
	l.outputf(FATAL, f, v...)
}

func (l *Logger) Errorf(f string, v ...interface{}) {
	l.outputf(ERROR, f, v...)
}

func (l *Logger) Warnf(f string, v ...interface{}) {
	l.outputf(WARN, f, v...)
}

func (l *Logger) Infof(f string, v ...interface{}) {
	l.outputf(INFO, f, v...)
}

func (l *Logger) Printf(f string, v ...interface{}) {
	l.outputf(INFO, f, v...)
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	l.outputf(DEBUG, f, v...)
}

func (l *Logger) output(level int, v ...interface{}) {
	if l.level < level {
		return
	}
	switch level {
	case DEBUG:
		l.logger.Println(append([]interface{}{"[Debug]"}, v...))
	case INFO:
		l.logger.Println(append([]interface{}{"[Info]"}, v...))
	case WARN:
		l.logger.Println(append([]interface{}{"[Warn]"}, v...))
	case ERROR:
		l.logger.Println(append([]interface{}{"[Error]"}, v...))
	case FATAL:
		l.logger.Println(append([]interface{}{"[Fatal]"}, v...))
	case DISABLED:
		return
	}
}

func (l *Logger) outputf(level int, f string, v ...interface{}) {
	if l.level < level {
		return
	}
	switch level {
	case DEBUG:
		l.logger.Printf("[Debug]"+f, v...)
	case INFO:
		l.logger.Printf("[Info]"+f, v...)
	case WARN:
		l.logger.Printf("[Warn]"+f, v...)
	case ERROR:
		l.logger.Printf("[Error]"+f, v...)
	case FATAL:
		l.logger.Printf("[Fatal]"+f, v...)
	case DISABLED:
	}
}
