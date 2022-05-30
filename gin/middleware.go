package gin

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/config"
	secure "github.com/unrolled/secure"

	httpsecure "github.com/krakendio/krakend-httpsecure/v2"
)

var ErrNoConfig = errors.New("no config present for the httpsecure module")

// Register registers the secure middleware into the gin engine
func Register(cfg config.ExtraConfig, engine *gin.Engine) error {
	opt, ok := httpsecure.ConfigGetter(cfg).(secure.Options)
	if !ok {
		return ErrNoConfig
	}
	engine.Use(secureMw(opt))
	return nil
}

// NewSecureMw creates a secured middleware for the gin engine
func NewSecureMw(cfg config.ExtraConfig) gin.HandlerFunc {
	opt, ok := httpsecure.ConfigGetter(cfg).(secure.Options)
	if !ok {
		return func(c *gin.Context) {}
	}

	return secureMw(opt)
}

// secureMw creates a secured middleware for the gin engine
func secureMw(opt secure.Options) gin.HandlerFunc {
	secureMiddleware := secure.New(opt)

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)

		if err != nil {
			c.Abort()
			return
		}

		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
	}
}
