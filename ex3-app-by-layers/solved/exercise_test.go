package solved

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController(t *testing.T) {
	ctrl := NewController()

	router := gin.Default()
	router.GET("/", ctrl.ControllerHandler)

	// Mocking
	ctrl.service.ServiceLogics = func(input string) string {
		return "All right!"
	}

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
