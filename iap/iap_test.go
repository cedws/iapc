package iap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectURL(t *testing.T) {
	url := connectURL(&dialOptions{
		Zone:    "zone",
		Region:  "region",
		Project: "project",
	})

	assert.Contains(t, url, proxyHost)
	assert.Contains(t, url, proxyPath)

	assert.Contains(t, url, "zone=zone")
	assert.Contains(t, url, "region=region")
	assert.Contains(t, url, "project=project")

	assert.NotContains(t, url, "token=")
	assert.NotContains(t, url, "group=")
	assert.NotContains(t, url, "port=")
}
