package httpsecure

import (
	"encoding/json"
	"fmt"

	"github.com/devopsfaith/krakend/config"
)

func ExampleConfigGetter() {
	cfg := ConfigGetter(config.ExtraConfig{
		Namespace: map[string]interface{}{
			"allowed_hosts":      []interface{}{"host1"},
			"host_proxy_headers": []interface{}{"x-custom-header"},
			"sts_seconds":        10.0,
			"ssl_redirect":       true,
			"ssl_host":           "secure.example.com",
		},
	})
	fmt.Println(cfg)

	// output:
	// {false false false false false false true false false false false       secure.example.com [host1] [x-custom-header] <nil> map[] 10 }
}

func ExampleConfigGetter_fromParsedData() {
	sample := `{
            "allowed_hosts": ["host1"],
            "ssl_proxy_headers": {},
            "sts_seconds": 300,
            "frame_deny": true,
            "sts_include_subdomains": true
        }`
	parsedCfg := map[string]interface{}{}

	if err := json.Unmarshal([]byte(sample), &parsedCfg); err != nil {
		fmt.Println(err)
		return
	}

	cfg := ConfigGetter(config.ExtraConfig{Namespace: parsedCfg})
	fmt.Println(cfg)

	// output:
	// {false false false true false false false false false true false        [host1] [] <nil> map[] 300 }
}
