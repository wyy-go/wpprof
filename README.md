# wpprof


![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wpprof?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wpprof)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wpprof)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wpprof/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wpprof)](https://goreportcard.com/report/github.com/wyy-go/wpprof)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wpprof?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wpprof/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wpprof)


gin pprof middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

## Usage

### Start using it

Download and install it:

```bash
go get github.com/wyy-go/wpprof
```

Import it in your code:

```go
import "github.com/wyy-go/wpprof"
```

### Example

```go
package main

import (
	"github.com/wyy-go/wpprof"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  wpprof.Register(router)
  router.Run(":8080")
}
```

### change default path prefix

```go
func main() {
	router := gin.Default()
	// default is "/debug/pprof"
	wpprof.Register(router, "/dev/pprof")
	router.Run(":8080")
}
```

### custom router group

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wpprof"
)

func main() {
	router := gin.Default()
	wpprof.Register(router,
		wpprof.WithPrefix("/admin"),
		wpprof.WithHandlers(func(c *gin.Context) {
			if c.Request.Header.Get("Authorization") != "foobar" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
			c.Next()
		}),
	)
	router.Run(":8080")
}

```

### Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```
