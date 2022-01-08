package wpprof

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultPrefix = "/debug/pprof"
)

type Option func(*Options)

type Options struct {
	prefix string
	handlers []gin.HandlerFunc
}

func WithPrefix(prefix string) Option {
	return func(opts *Options) {
		opts.prefix = prefix
	}
}

func WithHandlers(handlers ...gin.HandlerFunc) Option {
	return func(opts *Options) {
		opts.handlers = handlers
	}
}

func getPrefix(prefixOptions string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions
	}
	return prefix
}

// Register the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.GrouterGroup. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func Register(g *gin.Engine, opts ...Option) {
	options := Options{}

	for _, opt := range opts {
		opt(&options)
	}

	prefix := getPrefix(options.prefix)

	r := g.Group(prefix,options.handlers...)
	{
		r.GET("/", gin.WrapF(pprof.Index))
		r.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		r.GET("/profile", gin.WrapF(pprof.Profile))
		r.POST("/symbol", gin.WrapF(pprof.Symbol))
		r.GET("/symbol", gin.WrapF(pprof.Symbol))
		r.GET("/trace", gin.WrapF(pprof.Trace))
		r.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		r.GET("/block", gin.WrapH(pprof.Handler("block")))
		r.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		r.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		r.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		r.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
