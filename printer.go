package logk

import (
	"github.com/konsultin/logk/level"
	logkOption "github.com/konsultin/logk/option"
)

type Printer interface {
	Print(namespace string, outLevel level.LogLevel, msg string, options *logkOption.Options)
}
