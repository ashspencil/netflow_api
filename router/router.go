package router

import (
	v1 "nmg/netflow/handler/v1"

	"github.com/gin-gonic/gin"
)

func FlowRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/snmpwalk/ifoperstatus", v1.Snmpwalk_IfOperStatus)
		apiv1.GET("/snmpwalk/ifadminstatus", v1.Snmpwalk_IfAdminStatus)
		apiv1.GET("/snmpwalk/ifopersuspendedstatus", v1.Snmpwalk_IfOperSuspendedStatus)
		apiv1.GET("/snmpwalk/all", v1.Snmpwalk_All)
		apiv1.PUT("/snmpset/status", v1.SnmpSet_Status)
	}
	return r
}
