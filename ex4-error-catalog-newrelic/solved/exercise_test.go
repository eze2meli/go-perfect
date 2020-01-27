package solved

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService(t *testing.T) {
	ctrl := NewController()

	router := gin.Default()
	router.GET("/", ctrl.ControllerHandler)

	t.Run("Error empty input", func(t *testing.T) {
		performRequest(router, "GET", "/?input=")
		assert.Contains(t, loggedNewRelicErr.Error(), ErrServiceLogics.InputEmpty)
	})

	t.Run("Error invalid input", func(t *testing.T) {
		performRequest(router, "GET", "/?input=NaN")
		assert.Contains(t, loggedNewRelicErr.Error(), ErrServiceLogics.InputInvalid)
	})
}

// aux func
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}