package mylogging

import (
	"io"
	"os"

	logging "github.com/op/go-logging"
)

// Logger interface to go-logging.Logger struct
type Logger interface {
	Critical(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Notice(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	// with custom format print
	Criticalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

// Level just copy from go-logging package
type Level int

// Log levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// String returns the string representation of a logging level.
func (p Level) String() string {
	return levelNames[p]
}

const (
	pkgLogID      = "flogging"
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel  = INFO
)

var (
	logger *logging.Logger

	defaultOutput *os.File

	modulesLevel map[string]Level // Holds the map of all modules and their respective log level

	// in hyperledger/fabric - flogging package - use lock to write modulesLevel, write logging.SetLevel() etc, read modulesLevel etc.
	// lock sync.RWMutex
	// once sync.Once
)

func init() {
	logger = logging.MustGetLogger(pkgLogID)
	modulesLevel = make(map[string]Level)
	Reset()
	return
}

// Reset to init or reset to default setting
func Reset() {
	defaultOutput = os.Stderr
	modulesLevel[""] = defaultLevel
	InitBackend(SetFormat(defaultFormat), defaultOutput)
	return
}

// SetFormat parser format string to logging.formatter
func SetFormat(formatSpec string) logging.Formatter {
	if formatSpec == "" {
		formatSpec = defaultFormat
	}
	return logging.MustStringFormatter(formatSpec)
}

// InitBackend init backend, store it to logging.defaultBackend
func InitBackend(formatter logging.Formatter, output io.Writer) {
	backend := logging.NewLogBackend(output, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, formatter)
	logging.SetBackend(backendFormatter).SetLevel(getLoggingLevel(modulesLevel[""]), "")
	return
}

// getLoggingLevel convert type
func getLoggingLevel(level Level) logging.Level {
	return logging.Level(level)
}

// SetModuleLevel set(Level, moduleName) to logging
func SetModuleLevel(level Level, moduleName string) {
	modulesLevel[moduleName] = level
	logging.SetLevel(getLoggingLevel(modulesLevel[moduleName]), moduleName)
	return
}

// MustGetLogger use logging.MustGetLogger to get *Logger
// if a module do not SetModuleLevel before, it will get default level, then set level to modulesLevel map.
func MustGetLogger(moduleName string) Logger {
	l := logging.MustGetLogger(moduleName)
	modulesLevel[moduleName] = GetModuleLevel(moduleName)
	return l
}

// GetModuleLevel get logging.Level from logging by moduleName
func GetModuleLevel(moduleName string) Level {
	// logging.GetLevel() returns the logging level for the module, if defined.
	// Otherwise, it returns the default logging level, as set by
	// `flogging/logging.go`.
	level := logging.GetLevel(moduleName)
	return Level(level)
}

// GetArrayModulesLevel returns a array for modulesLevel
func GetArrayModulesLevel() [][]string {
	var modLevelList = [][]string{}
	for k, v := range modulesLevel {
		modLevelList = append(modLevelList, []string{k, levelNames[v]})
	}
	return modLevelList
}

// OverrideModulesLevel to set all modules to the same level
func OverrideModulesLevel(level Level) {
	for k := range modulesLevel {
		SetModuleLevel(level, k)
	}
}
