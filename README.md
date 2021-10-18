# logrgorm2

:smile: logrgorm2 is a logr logging driver for gorm v2

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/vchitai/logrgorm2)
[![License](https://img.shields.io/badge/license-MIT-%2397ca00.svg)](https://github.com/vchitai/logrgorm2/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/vchitai/logrgorm2.svg)](https://github.com/vchitai/logrgorm2/releases)
[![Made by Vchitai](https://img.shields.io/badge/made%20by-Vchitai-blue.svg?style=flat)](https://vchitai.github.io/)

[![GolangCI](https://golangci.com/badges/github.com/vchitai/logrgorm2.svg)](https://golangci.com/r/github.com/vchitai/logrgorm2)
[![codecov](https://codecov.io/gh/vchitai/logrgorm2/branch/main/graph/badge.svg?token=6QWOopYRPD)](https://codecov.io/gh/vchitai/logrgorm2)
[![Go Report Card](https://goreportcard.com/badge/github.com/vchitai/logrgorm2)](https://goreportcard.com/report/github.com/vchitai/logrgorm2)
[![CodeFactor](https://www.codefactor.io/repository/github/vchitai/logrgorm2/badge)](https://www.codefactor.io/repository/github/vchitai/logrgorm2)

If you're using gorm v1, you can use https://github.com/vchitai/logrgorm instead.

## Usage

```go
import "github.com/vchitai/logrgorm2"

stdr.SetVerbosity(1)
logger := stdr.NewWithOptions(stdlog.New(os.Stderr, "", stdlog.LstdFlags), stdr.Options{LogCaller: stdr.All})
logger = logger.WithName("logr")
db, err = gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{Logger: logger})
```

## Install

### Using go

```console
$ go get -u github.com/vchitai/logrgorm2
```

[comment]: <> (## Stargazers over time)

[comment]: <> ([![Stargazers over time]&#40;https://starchart.cc/vchitai/logrgorm2.svg&#41;]&#40;https://starchart.cc/vchitai/logrgorm2&#41;)
