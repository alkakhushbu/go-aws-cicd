package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestABC(t *testing.T) {
	// t.Errorf("This test would fail")

	//the below test would pass
	responseCode := http.StatusOK
	respondBody := "Hello World! Moved from gin to http standard library!"

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()

	ping(rec, req)
	require.Equal(t, responseCode, rec.Code)
	require.Equal(t, respondBody, rec.Body.String())
}
