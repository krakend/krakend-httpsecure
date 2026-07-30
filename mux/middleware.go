package mux

import (
	"net/http"

	"github.com/luraproject/lura/v3/config"
	"github.com/luraproject/lura/v3/router/mux"
	"github.com/unrolled/secure"

	httpsecure "github.com/krakend/krakend-httpsecure/v3"
)

// NewSecureMw creates a secured middleware for the mux engine
func NewSecureMw(cfg config.ExtraConfig) mux.HandlerMiddleware {
	opt, ok := httpsecure.ConfigGetter(cfg).(secure.Options)
	if !ok {
		return identityMiddleware{}
	}

	return secure.New(opt)
}

type identityMiddleware struct{}

func (i identityMiddleware) Handler(h http.Handler) http.Handler {
	return h
}
