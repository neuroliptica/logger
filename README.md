# logger
Simple personal logging wrapper around channels and log package.

## Install

```bash
$ go get github.com/neuroliptica/logger
```

## Usage

### Bind to default global logger

```go
package main

import "github.com/neuroliptica/logger"

func main() {
    // Create logger instance pointer and bind to default channel
    mainLogger := logger.MakeLogger("main-logger").BindToDefault()
    
    // Log message without format
    mainLogger.Log("message from main")
        // hh:mm:ss [main-logger] message from main
    
    // Log message with format
    mainLogger.Logf("append string with name %s %d times", "str", 1)
        // hh:mm:ss [main-logger] append string with name str 1 times
}
```

### Bind to custom logger

```go
package main

import (
    "log"
    "github.com/neuroliptica/logger"
)

func main() {
    // Create custom log messages destination
    custom := make(chan string)
    go func() {
        for msg := range custom {
            log.Println(msg)
        }
    }()
    
    mainLogger := logger.MakeLogger("custom-logger").BindToChannel(custom)
    mainLogger.Log("log message for custom dest")
}
```
