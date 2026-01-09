# logk - Logger Library

📝 Custom Konsultin leveled logger with namespace and child logger support.

## Installation

```bash
go get github.com/konsultin/logk
```

## Quick Start

```go
import "github.com/konsultin/logk"

// Get default logger
log := logk.Get()

// Log messages at different levels
log.Info("Server started")
log.Error("Connection failed", logkOption.Error(err))
log.Debug("Processing request", logkOption.AddMetadata("id", 123))

// Formatted logging
log.Infof("User %s logged in", username)
log.Errorf("Failed to process: %v", err)

// Create child logger with namespace
userLog := log.NewChild(logkOption.WithNamespace("user.service"))
userLog.Info("User created")
```

## Features

- **Multi-level Logging** - FATAL, ERROR, WARN, INFO, DEBUG, TRACE
- **Namespace Support** - Organize logs by domain/component
- **Child Loggers** - Create scoped loggers inheriting parent config
- **Metadata Attachment** - Add context data to log entries
- **Environment Config** - Configure via LOG_LEVEL and LOG_NAMESPACE

## License

MIT License - see [LICENSE](LICENSE)
