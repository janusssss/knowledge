package gomonkey

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func TestHello(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	patch := gomonkey.ApplyFunc(Hello, func(c echo.Context) error {
		return c.String(http.StatusOK, "The force be with you")
	})
	defer patch.Reset()

	_ = Hello(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "The force be with you", rec.Body.String())
}
