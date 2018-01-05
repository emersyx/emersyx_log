# emersyx_log [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url]

Logging functionality to be reused throughout emersyx components. Under the hood, this package uses the [go log][1]
package, so it does not reimplement anything fancy. The reason for having it as a separate module is to make integration
of the same code base easier into the all emersyx components. The result is a consistent logging format with little
effort.

[build-img]: https://travis-ci.org/emersyx/emersyx_log.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_log
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_log
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_log
[1]: https://golang.org/pkg/log/
