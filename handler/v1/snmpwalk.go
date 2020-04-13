package v1

import (
	"errors"
	"net/http"
	"net/url"
	"nmg/netflow/models"

	"github.com/gin-gonic/gin"
)

var snmp = *models.Snmp

// Request for IfOperStatus
func Snmpwalk_IfOperStatus(c *gin.Context) {

	oper := c.Request.URL.Query()

	if SetDevices(oper) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Query, need Parameter",
		})
		return
	}

	err := snmp.Connect()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to Connect: Please check Switch Status",
		})
		return
	}
	defer snmp.Conn.Close()

	oid := models.IfOperStatus
	result, err := snmp.WalkAll(oid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not Found, Please check the request setting",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": result})
}

// Request for IfAdminStatus
func Snmpwalk_IfAdminStatus(c *gin.Context) {

	admin := c.Request.URL.Query()

	if SetDevices(admin) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Query, need Parameter",
		})
		return
	}

	err := snmp.Connect()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to Connect: Please check Switch Status",
		})
		return
	}
	defer snmp.Conn.Close()

	oid := models.IfAdminStatus
	result, err := snmp.WalkAll(oid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not Found, Please check the request setting",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": result})
}

// Request for IfOperSuspendedStatus
func Snmpwalk_IfOperSuspendedStatus(c *gin.Context) {

	suspend := c.Request.URL.Query()

	if SetDevices(suspend) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Query, need Parameter",
		})
		return
	}

	err := snmp.Connect()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to Connect: Please check Switch Status",
		})
		return
	}
	defer snmp.Conn.Close()

	oid := models.IfOperSuspendedStatus
	result, err := snmp.WalkAll(oid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not Found, Please check the request setting",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": result})
}

// Request for all Status
func Snmpwalk_All(c *gin.Context) {

	q := c.Request.URL.Query()

	if SetDevices(q) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Query, need Parameter",
		})
		return
	}

	err := snmp.Connect()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to Connect: Please check Switch Status",
		})
		return
	}
	defer snmp.Conn.Close()

	oid := models.IfOperSuspendedStatus
	IfOperSuspendedStatus, err := snmp.WalkAll(oid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not Found, Please check the request setting",
		})
		return
	}

	oid = models.IfOperStatus
	IfOperStatus, err := snmp.WalkAll(oid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not Found, Please check the request setting",
		})
		return
	}

	oid = models.IfAdminStatus
	IfAdminStatus, err := snmp.WalkAll(oid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not Found, Please check the request setting",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"suspended": IfOperSuspendedStatus, "oper": IfOperStatus, "admin": IfAdminStatus})
}

func SetDevices(q url.Values) (err error) {
	// detect parameter
	defer func() {
		if recover() != nil {
			err = errors.New("Failed to Query, need Parameter")
		}
	}()

	switch q["device"][0] {
	case "nmg":
		snmp.Target = models.CISCO_NMG_2010
	case "3fn":
		snmp.Target = models.CISCO_NMG_3F_N
	case "4fs":
		snmp.Target = models.CISCO_NMG_4F_S
	case "4fn":
		snmp.Target = models.CISCO_NMG_4F_N
	}

	return err
}
