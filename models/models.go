package models

import (
	"log"
	"time"

	g "github.com/soniah/gosnmp"
)

var (
	Community             = <password>
	CISCO_NMG_2010        = <ip>
	CISCO_NMG_3F_N        = <ip>
	CISCO_NMG_4F_S        = <ip>
	CISCO_NMG_4F_N        = <ip>
	IfAdminStatus         = ".1.3.6.1.2.1.2.2.1.7"             //check out if the port is locked by admin
	IfOperStatus          = ".1.3.6.1.2.1.2.2.1.8"             //check out if the Port is being used
	IfOperSuspendedStatus = ".1.3.6.1.4.1.9.6.1.101.43.1.1.24" //check out if the Port is suspended by switch
	IfReActivate          = ".1.3.6.1.4.1.9.6.1.101.43.1.1.27" //Unsuspended Port

	Snmp     *g.GoSNMP
	Snmp_Set *g.SnmpPDU
)

func init() {
	Snmp = &g.GoSNMP{
		Target:    "", // default nothing
		Port:      161,
		Community: Community,
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	Snmp_Set = &g.SnmpPDU{
		Name:  "", // default nothing, want oi
		Type:  g.Integer,
		Value: 2, // default to locked
	}

	log.SetPrefix("[SNMP_CONNECT] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
