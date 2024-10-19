package iap

import (
	"crypto/rand"
	"encoding/hex"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randomString() string {
	buf := make([]byte, 16)
	if n, err := rand.Read(buf); err != nil || n != len(buf) {
		panic("failed to make random string")
	}
	return hex.EncodeToString(buf)
}

func TestSuccessFrame(t *testing.T) {
	t.Run("Write", func(t *testing.T) {
		id := randomString()
		buf := makeSuccessFrame(id)
		assert.Len(t, buf, 6+len(id))
	})
}

func TestAckFrame(t *testing.T) {
	t.Run("Write", func(t *testing.T) {
		buf := makeAckFrame(0x1337)
		assert.Len(t, buf, 10)
		assert.Equal(t, []byte{0x0, 0x7}, buf[0:2])
	})
}

func TestDataFrame(t *testing.T) {
	t.Run("Write", func(t *testing.T) {
		buf := makeDataFrame([]byte{0x13, 0x37})
		assert.Len(t, buf, 8)
		assert.Equal(t, []byte{0x0, 0x4}, buf[0:2])
		assert.Equal(t, []byte{0x0, 0x0, 0x0, 0x2}, buf[2:6])
		assert.Equal(t, []byte{0x13, 0x37}, buf[6:8])
	})
}

func TestConn(t *testing.T) {
	t.Run("With double Close", func(t *testing.T) {
		r, _ := net.Pipe()
		conn := newConn(r)

		assert.NoError(t, conn.Close())
		assert.NoError(t, conn.Close())
	})
}

func TestRead(t *testing.T) {
	t.Run("Without ACK", func(t *testing.T) {
		r, w := net.Pipe()

		defer r.Close()
		defer w.Close()

		conn := newConn(r)
		defer func() {
			assert.NoError(t, conn.Close())
		}()
		assert.False(t, conn.Connected())

		w.Write(makeSuccessFrame(randomString()))
		w.Write(makeDataFrame([]byte{0x13, 0x37}))

		buf := make([]byte, 2)
		n, err := conn.Read(buf)

		assert.NoError(t, err)
		assert.NotEmpty(t, conn.SessionID())
		assert.True(t, conn.Connected())
		assert.Equal(t, 2, n)
		assert.Equal(t, []byte{0x13, 0x37}, buf)
	})
}

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
