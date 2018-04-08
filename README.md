# emersyx_log [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url] [![GoDoc][godoc-img]][godoc-url]

This repository has been archived and is not under development anymore. The relevant code has been merged into the main
[emersyx repository][emersyx-repo]. This repository now serves for historical purposes only, as it contains the code for
emersyx version 0.1.

Logging functionality to be reused throughout emersyx components. Under the hood, this package uses the [go log][1]
package, so it does not reimplement anything fancy. The reason for having it as a separate module is to make integration
of the same code base easier into the all emersyx components. The result is a consistent logging format with little
effort.

[build-img]: https://travis-ci.org/emersyx/emersyx_log.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_log
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_log
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_log
[godoc-img]: https://godoc.org/emersyx.net/emersyx_log?status.svg
[godoc-url]: https://godoc.org/emersyx.net/emersyx_log
[emersyx-repo]: https://github.com/emersyx/emersyx
[1]: https://golang.org/pkg/log/
