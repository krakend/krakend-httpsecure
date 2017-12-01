package mux

import (
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/router/mux"
	"github.com/unrolled/secure"

	"github.com/devopsfaith/krakend-httpsecure"
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
