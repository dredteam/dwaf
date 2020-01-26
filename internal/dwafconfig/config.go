package dwafconfig

import (
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"net/url"
	"time"
)

type ReverseProxyConfiguration struct {
	Server *http.Server
	URL    *url.URL
}

type ProductConfiguration struct {
	LicenseKey string
}

type Configuration struct {
	ReverseProxy ReverseProxyConfiguration
	Product      ProductConfiguration
}

func GetConfiguration() Configuration {
	c := Configuration{}
	c.ReverseProxy = getReverseProxyConfiguration()
	c.Product = getProductConfiguration()
	return c
}

func getReverseProxyConfiguration() ReverseProxyConfiguration {
	type Specification struct {
		ServerAddress           string        `default:":8080" split_words:"true"`
		ServerReadHeaderTimeout time.Duration `default:"4s" split_words:"true"`
		ServerReadTimeout       time.Duration `default:"9s" split_words:"true"`
		ServerWriteTimeout      time.Duration `default:"9s" split_words:"true"`
		ServerIdleTimeout       time.Duration `default:"9s" split_words:"true"`
		ServerMaxHeaderBytes    int           `default:"1048576" split_words:"true"`
		ServerProxyUrl          string        `required:"true" split_words:"true"`
	}
	var spec Specification
	err := envconfig.Process("dwaf", &spec)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:              spec.ServerAddress,
		ReadHeaderTimeout: spec.ServerReadHeaderTimeout,
		ReadTimeout:       spec.ServerReadTimeout,
		WriteTimeout:      spec.ServerWriteTimeout,
		IdleTimeout:       spec.ServerIdleTimeout,
		MaxHeaderBytes:    spec.ServerMaxHeaderBytes,
	}

	return ReverseProxyConfiguration{Server: server, URL: parseUrl(spec.ServerProxyUrl)}
}

func getProductConfiguration() ProductConfiguration {
	return ProductConfiguration{}
}

func parseUrl(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
