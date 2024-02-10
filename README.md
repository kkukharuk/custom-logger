# custom-logger
Custom logging library

## Install

    go get github.com/kkukharuk/custom-logger

## Usage

### Configure

**Type** - [_string_] where are the work logs written. Variants: stdout, stderr, file (default: stdout)

**Format** - [_string_] formating work log. Variants: text, json (default: text)

**Level** - [_string_] the level of work logs. Variants: info, error, warn, debug (default: error)

**LogFile** - [_string_] the path to the work logs file.

**Size** - [_int_] the size of the work log files (in megabytes).

**Backaps** - [_int_] the number of backup files of work logs.

**age** - [_int_] storage time for backups of work log files.

**Compress** - [_bool_] is it necessary to compress archived work log files.

**Prefix** - [_string_] the label of the application/utility will always be displayed in the work logs.

### Example

    package main

    import (
        log "github.com/kkukharuk/custom-logger"
    )

    func main() {
        logConfig := log.Config{
            Type:   "stdout",
            Format: "text",
            Level:  "debug",
            Prefix: "example",
        }
        logger, err := log.Init(logConfig)
        if err != nil {
            panic(err)
        }
        logger.Info("Info message", "test1")
        logger.Error("Error message", "test2")
        logger.Warn("Warning message", "test3")
        logger.Debug("Debug message", "test4")
        logger.Fatal("Fatal message", "test5")
    }

Result:

    2024-02-10T14:05:46.477997 [info] example :: init :: Initialize logger
    2024-02-10T14:05:46.478149 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478163 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478172 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠄⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478180 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠄⠄⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478187 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠁⠄⠄⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478195 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478202 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⣿⣿⣿
    2024-02-10T14:05:46.478209 [info] example :: init :: ⣿⣿⣿⠄⠄⠄⠄⠄⣿⠛⠋⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣿⣿⣿
    2024-02-10T14:05:46.478228 [info] example :: init :: ⣿⣿⣿⠄⠄⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿
    2024-02-10T14:05:46.478237 [info] example :: init :: ⣿⣿⣿⠄⠄⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣴⣿⣿⣿
    2024-02-10T14:05:46.478244 [info] example :: init :: ⣿⣿⣿⠄⠄⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢈⣿⣿⣿
    2024-02-10T14:05:46.478252 [info] example :: init :: ⣿⣿⣿⠄⠄⣴⣶⡄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠰⣿⣿⣿⣿
    2024-02-10T14:05:46.478259 [info] example :: init :: ⣿⣿⣿⣤⣤⣭⣯⣤⣿⣿⣿⣷⣶⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣤⣿⣿⣿⣿
    2024-02-10T14:05:46.478267 [info] example :: init :: ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
    2024-02-10T14:05:46.478274 [info] example :: test1 :: Info message
    2024-02-10T14:05:46.478282 [error] example :: test2 :: Error message
    2024-02-10T14:05:46.478290 [warning] example :: test3 :: Warning message
    2024-02-10T14:05:46.478305 [debug] example :: test4 :: Debug message
    2024-02-10T14:05:46.478314 [fatal] example :: test5 :: Fatal message
