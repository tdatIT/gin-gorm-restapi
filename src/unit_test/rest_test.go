package unit_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tdatIT/gin-gorm-restapi/src/config"
	"github.com/tdatIT/gin-gorm-restapi/src/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI(t *testing.T) {
	config.InitConnection()
	r := gin.New()
	routers.RegisterRouters(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"product_id\":1,\"name\":\"iPhone14 Promax 256GB VN/A\",\"price\":1299.5,\"create_at\":\"2006-01-03T05:04:05+07:00\"}", w.Body.String())

}
