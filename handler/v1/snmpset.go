package v1

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"nmg/netflow/models"

	"github.com/gin-gonic/gin"
	g "github.com/soniah/gosnmp"
)

type Data struct {
	Params *Params `json:"params"`
}

type Params struct {
	Device          string `json:"device"`
	Action          string `json:"action"`
	AdminOid        string `json:"adminoid"`
	AdminStatus     int    `json:"adminstatus"`
	SuspendedOid    string `json:"suspendedoid"`
	SuspendedStatus int    `json:"suspendedstatus"`
}

var Snmp_Set = *models.Snmp_Set

func SnmpSet_Status(c *gin.Context) {

	if SetDatas(c.Request.Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Query, need Payload",
		})
		return
	}

	var set_req []g.SnmpPDU

	err := snmp.Connect()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to Connect: Please check Switch Status",
		})
		return
	}
	defer snmp.Conn.Close()

	set_req = append(set_req, Snmp_Set)

	result, err := snmp.Set(set_req)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to Set, Please check the request setting",
		})
		return
	}

	c.JSON(http.StatusOK, result)
	return

}

func SetDatas(payload io.ReadCloser) (err error) {
	//detect parameter
	defer func() {
		if recover() != nil {
			err = errors.New("Failed to Query, need Payload")
		}
	}()

	body, err2 := ioutil.ReadAll(payload)
	if err2 != nil {
		log.Println("Failed to read payload")
	}
	str := string(body)
	data := &Data{}

	err2 = json.Unmarshal([]byte(str), data)
	if err2 != nil {
		log.Println("Failed to decode payload")
	}

	switch data.Params.Device {
	case "nmg":
		snmp.Target = models.CISCO_NMG_2010
	case "3fn":
		snmp.Target = models.CISCO_NMG_3F_N
	case "4fs":
		snmp.Target = models.CISCO_NMG_4F_S
	case "4fn":
		snmp.Target = models.CISCO_NMG_4F_N
	}

	switch data.Params.Action {
	case "Admin":
		Snmp_Set.Name = data.Params.AdminOid
		if data.Params.AdminStatus == 2 { // if locked
			Snmp_Set.Value = 1 // then unlock
		} else {
			Snmp_Set.Value = 2 // or lock
		}
	case "Suspended":
		oid := data.Params.SuspendedOid
		port_number := oid[len(oid)-1 : len(oid)] // get the port number
		Snmp_Set.Name = models.IfReActivate + "." + port_number

		if data.Params.SuspendedStatus == 1 { // if suspended
			Snmp_Set.Value = 1 // then unsuspended
		}
	}
	return err
}
