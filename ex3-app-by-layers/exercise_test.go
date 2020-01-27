package ex3_app_by_layers

import (
	"github.com/eze2meli/go-perfect/ex3-app-by-layers/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO 3: Mock service call. Avoid using mock frameworks
func TestController(t *testing.T) {
	router := gin.Default()
	router.GET("/", controller.ControllerHandler)

	// Dont modify...
	w := performRequest(router, "GET", "/?input=hello")
	assert.Equal(t, 200, w.Code)
}

// aux func
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
