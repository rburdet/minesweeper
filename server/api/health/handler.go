package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct{}

func (hh *HealthHandler) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}
