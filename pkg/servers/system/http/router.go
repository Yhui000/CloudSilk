package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterRemoteServiceTaskQueueRouter(r)
	RegisterRemoteServiceTaskRouter(r)
	RegisterRemoteServiceRouter(r)
	RegisterInvocationAuthenticationRouter(r)
	RegisterSerialNumberRouter(r)
	RegisterSystemEventTriggerRouter(r)
	RegisterSystemEventRouter(r)
	RegisterSystemParamsConfigRouter(r)
	RegisterTaskQueueExecutionRouter(r)
	RegisterTaskQueueRouter(r)
	RegisterDataMappingRouter(r)
	RegisterPrintServerRouter(r)
}
