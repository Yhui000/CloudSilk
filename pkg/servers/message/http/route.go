package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterMessageTypeRouter(r)
	RegisterMessageTemplateRouter(r)
}
