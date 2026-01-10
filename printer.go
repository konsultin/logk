package logk

import (
	"github.com/go-konsultin/logk/level"
	logkOption "github.com/go-konsultin/logk/option"
)

type Printer interface {
	Print(namespace string, outLevel level.LogLevel, msg string, options *logkOption.Options)
}
