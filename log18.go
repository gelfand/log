//go:build go1.18

package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	// reset is a reset color.
	reset = "\033[0m"
	// red is a red color.
	red = "\033[31m"
	// boldRED is a bold red color.
	boldRed = "\033[1;31m"
	// green is a green color.
	green = "\033[32m"
	// orange is a orange color.
	orange = "\033[33m"
	// cyan is a cyan color.
	cyan = "\033[36m"
)

// logLvl is a type which represent log record type.
type logLvl int8

const (
	lvlDebug logLvl = iota
	lvlInfo
	lvlWarn
	lvlError
	lvlFatal
)

// String implements Stringer interface for the logLvl.
func (lvl logLvl) String() string {
	switch lvl {
	case lvlDebug:
		return "DBUG"
	case lvlInfo:
		return "INFO"
	case lvlWarn:
		return "WARN"
	case lvlError:
		return "EROR"
	case lvlFatal:
		return "FATL"
	default:
		return ""
	}
}

// color is a helper function for coloring log record.
func (lvl logLvl) color() string {
	switch lvl {
	case lvlDebug:
		return cyan
	case lvlInfo:
		return green
	case lvlWarn:
		return orange
	case lvlError:
		return red
	case lvlFatal:
		return boldRed
	default:
		return ""
	}
}

// reset is a helper function for resetting color.
func (lvl logLvl) reset() string {
	return reset
}

// Logger is an interface type which represents a logger.
type Logger interface {
	// Debug is the debug level log method.
	Debug(arg any)
	// Debugf the same as Debug except formats string with provided arguments.
	Debugf(format string, args ...any)
	// Info is the info level log method.
	Info(arg any)
	// Infof the same as Info except formats string with provided arguments.
	Infof(format string, args ...any)
	// Warn is the warn level log method.
	Warn(arg any)
	// Warnf the same as Warn except formats string with provided arguments.
	Warnf(format string, args ...any)
	// Error is the error level log method.
	Error(arg any)
	// Errorf the same as Error except formats string with provided arguments.
	Errorf(format string, args ...any)
	// Fatal is the fatal level log method.
	Fatal(arg any)
	// Fatalf the same as Fatal except formats string with provided arguments.
	Fatalf(format string, args ...any)
}

// Log is the implementation of the Logger interface.
type Log struct {
	// w is the io.Writer to write log record.
	w io.Writer
}

// New creates a new Logger with provided io.Writer.
func New(w io.Writer) Logger {
	return &Log{w}
}

// rootLogger is a global variable being used for logging without creating
// Log object from the caller side.
var rootLogger = New(os.Stderr)

// Debug writes debug record.
func (l *Log) Debug(arg any) {
	fmt.Fprint(l.w, doLog(lvlDebug, arg))
}

// Debugf writes a Debug record.
func (l *Log) Debugf(format string, args ...any) {
	fmt.Fprint(l.w, logf(lvlDebug, format, args))
}

// Info writes info record.
func (l *Log) Info(arg any) {
	fmt.Fprint(l.w, doLog(lvlInfo, arg))
}

// Infof is the same as Info except formats string with provided arguments.
func (l *Log) Infof(format string, args ...any) {
	fmt.Fprint(l.w, logf(lvlInfo, format, args))
}

// Warn writes warning record.
func (l *Log) Warn(arg any) {
	fmt.Fprint(l.w, doLog(lvlWarn, arg))
}

// Warnf is the same as Warn except formats string with provided arguments.
func (l *Log) Warnf(format string, args ...any) {
	fmt.Fprint(l.w, logf(lvlWarn, format, args))
}

// Erorr is the error level log method.
func (l *Log) Error(arg any) {
	fmt.Fprint(l.w, doLog(lvlError, arg))
}

// Erorrf is the same as Error except formats string with provided arguments.
func (l *Log) Errorf(format string, args ...any) {
	fmt.Fprint(l.w, logf(lvlError, format, args))
}

// Fatal is the fatal level log method.
func (l *Log) Fatal(arg any) {
	fmt.Fprint(l.w, doLog(lvlFatal, arg))
	die()
}

// Fatalf is the same as Fatal except formats string with provided arguments.
func (l *Log) Fatalf(format string, args ...any) {
	fmt.Fprint(l.w, logf(lvlFatal, format, args))
	die()
}

// Info writes info record.
func Info(arg any) {
	rootLogger.Info(arg)
}

// Infof is the same as Info except formats string with provided arguments.
func Infof(format string, args ...any) {
	rootLogger.Infof(format, args...)
}

// Warn writes warning record.
func Warn(arg any) {
	rootLogger.Warn(arg)
}

// Warnf is the same as Warn except formats string with provided arguments.
func Warnf(format string, args ...any) {
	rootLogger.Warnf(format, args...)
}

// Error writes error record.
func Error(arg any) {
	rootLogger.Error(arg)
}

// Errorf is the same as Error except formats string with provided arguments.
func Errorf(format string, args ...any) {
	rootLogger.Errorf(format, args...)
}

// Debug writes debug record.
func Debug(arg any) {
	rootLogger.Debug(arg)
}

// Debugf is the same as Debug except formats string with provided arguments.
func Debugf(format string, args ...any) {
	rootLogger.Debugf(format, args...)
}

// Fatal is the fatal level log method.
func Fatal(arg any) {
	rootLogger.Fatal(arg)
}

// Fatalf is the same as Fatal except formats string with provided arguments.
func Fatalf(format string, args ...any) {
	rootLogger.Fatalf(format, args...)
}

// logf is a helper function for formatting log record.
func logf(lvl logLvl, format string, a []any) string {
	return doLog(lvl, fmt.Sprintf(format, a...))
}

// doLog is a helper function for logging.
func doLog(lvl logLvl, arg any) string {
	record := lvl.color() + lvl.String() + lvl.reset()

	// Add time.
	record += "[" + getTime() + "] "
	// Format.
	record += fmt.Sprint(arg)
	record += "\n"

	return record
}

// getTime returns current time in HH:MM:SS.NANO XM
// for example: 00:30:23.607883616 PM.
func getTime() string {
	currentTime := time.Now()

	var hourString string
	var dayPart string

	if hour := currentTime.Hour(); hour >= 12 {
		hour -= 12
		hourString = fmt.Sprintf("%02d", hour)
		dayPart = "PM"
	} else {
		hourString = fmt.Sprintf("%02d", hour)
		dayPart = "AM"
	}

	minute := fmt.Sprintf("%02d", time.Now().Minute())
	second := fmt.Sprintf("%02d", time.Now().Second())
	nanosecondString := fmt.Sprintf("%09d", currentTime.Nanosecond())

	logTime := hourString + ":" + minute + ":" + second + "." + nanosecondString + " " + dayPart

	return logTime
}

// die is a helper function for exiting the program.
func die() {
	os.Exit(1)
}
