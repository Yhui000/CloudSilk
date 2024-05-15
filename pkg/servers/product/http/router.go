package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterProductInfoRouter(r)
	RegisterProductOrderAttributeRouter(r)
	RegisterProductOrderPackageRouter(r)
	RegisterProductOrderProcessStepRouter(r)
	RegisterProductOrderProcessRouter(r)
	RegisterProductOrderRouter(r)
	RegisterProductOrderLabelRouter(r)
	RegisterProductOrderBomRouter(r)
	RegisterProductProcessRouteRouter(r)
	RegisterProductPackageMatchRuleRouter(r)
	RegisterProductPackageTypeRouter(r)
	RegisterProductPackageRouter(r)
	RegisterProductPackageStackRuleRouter(r)
	RegisterProductReworkTypeRouter(r)
	RegisterProductReworkCauseRouter(r)
	RegisterProductReworkSolutionRouter(r)
	RegisterProductReleaseStrategyRouter(r)
	RegisterProductAttributeValuateRuleRouter(r)
	RegisterProductOrderPriorityRuleRouter(r)
	RegisterProductOrderReleaseRuleRouter(r)
	RegisterProductOrderAttachmentRouter(r)
	RegisterProductPackageRecordRouter(r)
	RegisterProductRhythmRecordRouter(r)
	RegisterProductReworkRecordRouter(r)
	RegisterProductTestRecordRouter(r)
	RegisterProductIssueRecordRouter(r)
	RegisterProductProcessRecordRouter(r)
	RegisterProductWorkRecordRouter(r)
	RegisterProductReleaseRecordRouter(r)
	RegisterProductOrderPalletRouter(r)
}