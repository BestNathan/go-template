package healthy

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var hm *HealthyModule

func init() {
	hm = &HealthyModule{
		healthCtlr: HC(),
	}
}

type HealthyModule struct {
	healthCtlr *HealthController
}

func HM() *HealthyModule {
	if hm == nil {
		panic(errors.New("healthy module is not inited"))
	}

	return hm
}

func (h *HealthyModule) Route(r gin.IRouter) {
	h.healthCtlr.Route(r.Group("/health"))
}
