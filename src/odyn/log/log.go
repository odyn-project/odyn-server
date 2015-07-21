// Copyright 2015 Odyn Authors (see AUTHORS file for project)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


// Logging library for Odyn.
package log

import (
    "log"
    "os"
    "fmt"
)

type OdynLogger struct {
    logFile *os.File
    logger *log.Logger
    errorLogger *log.Logger
    warnLogger *log.Logger
}

var std = OdynLogger{}

// If /var/log/canopy files cannot be opened, then fallback to just logging to STDOUT
func initFallback() error {
    std.infoLogger = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)
    std.errorLogger = log.New(os.Stdout, "ERROR ", log.LstdFlags | log.Lshortfile)
    std.warnLogger = log.New(os.Stdout, "WARN ", log.LstdFlags | log.Lshortfile)

    return nil
}

// Initialize Odyn logger
func Init(logFilename string) error {
    var err error
    std.logFile, err = os.OpenFile(logFilename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666);
    if err != nil {
        fmt.Println("Error opening file " + logFilename + ": ", err)
        fmt.Println("Falling back to STDOUT for logging")
        return initFallback()
    }
    std.infoLogger = log.New(std.logFile, "", log.LstdFlags | log.Lshortfile)

    std.errorLogger = log.New(std.logFile, "ERROR ", log.LstdFlags | log.Lshortfile)
    std.warnLogger = log.New(std.logFile, "WARN ", log.LstdFlags | log.Lshortfile)

    return nil
}

// Close Odyn log file
func Shutdown() {
    std.infoLogger.Output(2, fmt.Sprintln("Goodbye"));
    if (std.logFile != nil) {
        std.logFile.Close()
    }
}

// Log an error
func Error(v ...interface{}) {
    std.errorLogger.Output(2, fmt.Sprintln(v...))
}

func ErrorCalldepth(calldepth int, v ...interface{}) {
    std.errorLogger.Output(calldepth, fmt.Sprintln(v...))
}

// Log a warning
func Warn(v ...interface{}) {
    std.warnLogger.Output(2, fmt.Sprintln(v...))
}

// Log an information statement
func Info(v ...interface{}) {
    std.infoLogger.Output(2, fmt.Sprintln(v...))
}
