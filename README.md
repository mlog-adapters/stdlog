# stdlog adapter
[![License][license-img]][license] [![Actions Status][action-img]][action] [![GoDoc][godoc-img]][godoc] [![Go Report Card][goreport-img]][goreport] [![Coverage Status][codecov-img]][codecov]

stdlog provides standard logger adapter to mlog

TODO: add more info

Example:
```
	...
import (
	"log"
	"os"

	"github.com/mlog-adapters/stdlog"
	"go.melnyk.org/mlog"
	"go.melnyk.org/mlog/console"
)
	...
	lb := console.NewLogbook(os.Stdout)     // Create new logbook from console implementation
	lb.SetLevel(mlog.Default, mlog.Verbose) // Set default logging level to Verbose
	logger := lb.Joiner().Join("std")       // Get (join) logger with name "std"
	...
	s := &http.Server{
       Addr:           ":8080",
       Handler:        router,
       ErrorLog:       stdlog.NewStdLog(logger, mlog.Warning), // Get std adaoter to mlog logger
	}
	...
	 restore := stdlog.RedirectStdLog(logger, mlog.Warning) // Redirect standard logger to mlog logger
	 defer restore() // Restore standard logger state

	 log.Print("just standard log output")
	...
```

## Development Status: In active development
All APIs are in active development and not finalized, and breaking changes will be made in the 0.x series.


[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license]: https://github.com/mmelnyk/mpool/blob/master/LICENSE
[action-img]: https://github.com/mlog-adapters/stdlog/workflows/Test/badge.svg
[action]: https://github.com/mlog-adapters/stdlog/actions
[godoc-img]: https://godoc.org/github.com/mlog-adapters/stdlog?status.svg
[godoc]: https://godoc.org/github.com/mlog-adapters/stdlog
[goreport-img]: https://goreportcard.com/badge/github.com/mlog-adapters/stdlog
[goreport]: https://goreportcard.com/report/github.com/mlog-adapters/stdlog
[codecov-img]: https://codecov.io/gh/mlog-adapters/stdlog/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/mlog-adapters/stdlog
