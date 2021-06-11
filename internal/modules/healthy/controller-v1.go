package healthy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func HC() *HealthController {
	return &HealthController{}
}

func (hc *HealthController) Route(r gin.IRouter) {
	r.GET("", hc.health)
}

func (hc *HealthController) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": "0", "message": "SUCCESS"})
}
