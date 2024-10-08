package http

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/require"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

func newOneRouteHandler(path string, handler fiber.Handler) *oneRouteHandler {
	return &oneRouteHandler{path: path, handler: handler}
}

type oneRouteHandler struct {
	path    string
	handler fiber.Handler
}

func (h *oneRouteHandler) Register(router fiber.Router) {
	router.Get(h.path, h.handler)
}

func (h *oneRouteHandler) GenerateRequest(t *testing.T, cfg Config) *http.Request {
	r, err := http.NewRequest(http.MethodGet, "http://"+filepath.Join(cfg.DNS(), h.path), nil)
	require.Nil(t, err)

	return r
}

// go test ./pkg/transport/http -run Test_Response_OK
func Test_Response_OK(t *testing.T) {
	h := newOneRouteHandler("test", func(ctx fiber.Ctx) error {
		return OK(ctx, nil)
	})

	server := NewServer(context.Background(), "", logger, correctCfg, []Handler{h})
	server.AsyncRun()

	time.Sleep(time.Millisecond * 100)

	defer server.Shutdown(context.Background())

	req := h.GenerateRequest(t, correctCfg)
	resp, err := http.DefaultClient.Do(req)

	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, http.StatusOK)
}

// go test ./pkg/transport/http -run Test_Response_BadRequest
func Test_Response_BadRequest(t *testing.T) {
	testResponseWriter(t, BadRequest, http.StatusBadRequest)
}

// go test ./pkg/transport/http -run Test_Response_ServerError
func Test_Response_ServerError(t *testing.T) {
	testResponseWriter(t, ServerError, http.StatusInternalServerError)
}

func testResponseWriter(t *testing.T, fn func(ctx fiber.Ctx, err error) error, status int) {
	h := newOneRouteHandler("test", func(ctx fiber.Ctx) error {
		return fn(ctx, nil)
	})

	server := NewServer(context.Background(), "", logger, correctCfg, []Handler{h})
	server.AsyncRun()

	time.Sleep(time.Millisecond * 100)

	defer server.Shutdown(context.Background())

	req := h.GenerateRequest(t, correctCfg)
	resp, err := http.DefaultClient.Do(req)

	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, status)
}
