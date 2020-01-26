package dwafconfig

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestGetReverseProxyConfiguration(t *testing.T) {

	t.Run("DWAF_SERVER_PROXY_URL=http://example.com/", func(t *testing.T) {
		var got ReverseProxyConfiguration
		var duration time.Duration
		_ = os.Setenv("DWAF_SERVER_PROXY_URL", "http://example.com/")
		got = getReverseProxyConfiguration()

		assert.Equal(t, got.URL.String(), "http://example.com/")
		assert.Equal(t, got.Server.Addr, ":8080")
		assert.Equal(t, got.Server.MaxHeaderBytes, 1048576)

		duration, _ = time.ParseDuration("9s")
		assert.Equal(t, got.Server.ReadTimeout, duration)
		assert.Equal(t, got.Server.WriteTimeout, duration)
		assert.Equal(t, got.Server.IdleTimeout, duration)

		duration, _ = time.ParseDuration("4s")
		assert.Equal(t, got.Server.ReadHeaderTimeout, duration)

		_ = os.Unsetenv("SERVER_PROXY_URL")
	})

	t.Run("DWAF_SERVER_PROXY_URL=https://dredteam.com/with-path", func(t *testing.T) {
		var got ReverseProxyConfiguration
		_ = os.Setenv("DWAF_SERVER_PROXY_URL", "https://dredteam.com/with-path")
		got = getReverseProxyConfiguration()
		assert.Equal(t, got.URL.String(), "https://dredteam.com/with-path")
		_ = os.Unsetenv("SERVER_PROXY_URL")
	})

	t.Run("DWAF_SERVER_ADDRESS=:8000", func(t *testing.T) {
		var got ReverseProxyConfiguration
		_ = os.Setenv("DWAF_SERVER_ADDRESS", ":8000")
		got = getReverseProxyConfiguration()
		assert.Equal(t, got.Server.Addr, ":8000")
		_ = os.Unsetenv("DWAF_SERVER_ADDRESS")
	})
}
