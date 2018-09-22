package httpsecure

import (
	"github.com/devopsfaith/krakend/config"
	"github.com/unrolled/secure"
)

// Namespace is the key to use to store and access the custom config data
const Namespace = "github_com/devopsfaith/krakend-httpsecure"

// ZeroCfg is the zero value for the Config struct.
// Deprecated: the config getter does not return a ZeroCfg when no config available
var ZeroCfg = secure.Options{}

// ConfigGetter implements the config.ConfigGetter interface. It parses the extra config for the
// package and returns nil if something goes wrong.
func ConfigGetter(e config.ExtraConfig) interface{} {
	v, ok := e[Namespace]
	if !ok {
		return nil
	}
	tmp, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}

	cfg := secure.Options{}

	getStrings(tmp, "allowed_hosts", &cfg.AllowedHosts)
	getStrings(tmp, "host_proxy_headers", &cfg.HostsProxyHeaders)

	getInt64(tmp, "sts_seconds", &cfg.STSSeconds)

	getString(tmp, "custom_frame_options_value", &cfg.CustomFrameOptionsValue)
	getString(tmp, "content_security_policy", &cfg.ContentSecurityPolicy)
	getString(tmp, "public_key", &cfg.PublicKey)
	getString(tmp, "ssl_host", &cfg.SSLHost)
	getString(tmp, "referrer_policy", &cfg.ReferrerPolicy)

	getBool(tmp, "content_type_nosniff", &cfg.ContentTypeNosniff)
	getBool(tmp, "browser_xss_filter", &cfg.BrowserXssFilter)
	getBool(tmp, "is_development", &cfg.IsDevelopment)
	getBool(tmp, "sts_include_subdomains", &cfg.STSIncludeSubdomains)
	getBool(tmp, "frame_deny", &cfg.FrameDeny)
	getBool(tmp, "ssl_redirect", &cfg.SSLRedirect)

	return cfg
}

func getStrings(data map[string]interface{}, key string, v *[]string) {
	if vs, ok := data[key]; ok {
		result := []string{}
		for _, v := range vs.([]interface{}) {
			result = append(result, v.(string))
		}
		*v = result
	}
}

func getString(data map[string]interface{}, key string, v *string) {
	if val, ok := data[key]; ok {
		*v = val.(string)
	}
}

func getBool(data map[string]interface{}, key string, v *bool) {
	if val, ok := data[key]; ok {
		*v = val.(bool)
	}
}

func getInt64(data map[string]interface{}, key string, v *int64) {
	if val, ok := data[key]; ok {
		*v = int64(val.(int))
	}
}
