# emersyx_log [![Build Status](https://travis-ci.org/emersyx/emersyx_log.svg?branch=master)](https://travis-ci.org/emersyx/emersyx_log)

Logging functionality to be reused throughout emersyx components. Under the hood, this package uses the [go log][1]
package, so it does not reimplement anything fancy. The reason for having it as a separate module is to make integration
of the same code base easier into the all emersyx components. The result is a consistent logging format with little
effort.

[1]: https://golang.org/pkg/log/
