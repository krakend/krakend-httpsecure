package httpsecure

import (
	"fmt"
	"testing"

	"github.com/devopsfaith/krakend/config"
)

func TestConfigGetter(t *testing.T) {
	cfg := ConfigGetter(config.ExtraConfig{
		Namespace: map[string]interface{}{
			"allowed_hosts":      []interface{}{"host1"},
			"host_proxy_headers": []interface{}{"x-custom-header"},
			"sts_seconds":        10,
			"ssl_redirect":       true,
			"ssl_host":           "secure.example.com",
		},
	})
	fmt.Println(cfg)
}
