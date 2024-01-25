package test

import (
	"net/http"
	"net/http/httptest"
)

func (uts *UnitTestSuite) TestHealthCheck() {
	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)

	uts.Equal(http.StatusOK, resp.StatusCode)
}
