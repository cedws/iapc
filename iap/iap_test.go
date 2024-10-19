package iap

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/stretchr/testify/assert"
)

var testData = []byte("hello")

var wsListener net.Listener

func wsUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{proxySubproto},
	})
	if err != nil {
		panic(err)
	}
	defer conn.CloseNow()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn.Write(ctx, websocket.MessageBinary, makeSuccessFrame(randomString()))
	conn.Write(ctx, websocket.MessageBinary, makeDataFrame(testData))

	conn.Close(websocket.StatusNormalClosure, "")
}

func randomString() string {
	buf := make([]byte, 16)
	if n, err := rand.Read(buf); err != nil || n != len(buf) {
		panic("failed to make random string")
	}
	return hex.EncodeToString(buf)
}

func TestMain(m *testing.M) {
	// websocket library rejects non-URL origins, need to override
	proxyOrigin = ""

	mux := http.NewServeMux()
	mux.HandleFunc("/", wsUpgradeHandler)

	var err error

	wsListener, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(wsListener, mux)

	m.Run()
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
	t.Run("E2E Read", func(t *testing.T) {
		conn, err := dial(context.Background(), "ws://"+wsListener.Addr().String())
		assert.NoError(t, err)
		defer func() {
			assert.NoError(t, conn.Close())
		}()

		buf := make([]byte, len(testData))
		_, err = conn.Read(buf)

		assert.NoError(t, err)
		assert.Equal(t, testData, buf)

		assert.True(t, conn.Connected())
		assert.NotEmpty(t, conn.SessionID())
	})

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
		w.Write(makeDataFrame(testData))

		buf := make([]byte, len(testData))
		n, err := conn.Read(buf)

		assert.NoError(t, err)
		assert.Equal(t, testData, buf)
		assert.Equal(t, len(testData), n)

		assert.NotEmpty(t, conn.SessionID())
		assert.True(t, conn.Connected())
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
