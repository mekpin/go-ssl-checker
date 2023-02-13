package controller

import (
	"crypto/tls"
	"fmt"
	"go-ssl-checker/config"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SSLCheck(c *gin.Context) {
	//getexpiry := ""
	AllDomainPort := config.Common.Alldomainport
	Domain1 := config.Common.Domain1
	// Domain2 := config.Common.Domain2
	// Domain3 := config.Common.Domain3
	// Domain4 := config.Common.Domain4
	// Domain5 := config.Common.Domain5

	conn, err := net.Dial("tcp", Domain1+":"+AllDomainPort)
	if err != nil {
		log.Fatal(err)
	}

	client := tls.Client(conn, &tls.Config{
		ServerName: Domain1,
	})
	defer client.Close()

	if err := client.Handshake(); err != nil {
		log.Fatal(err)
	}

	cert := client.ConnectionState().PeerCertificates[0]
	expirydomain1 := cert.NotAfter.Format(time.RFC3339)
	// parseexpirydomain1 := time.Parse(layout, time.L)(expirydomain1)
	currenttime := time.Now()
	delta1 := cert.NotAfter.Sub(currenttime)

	fmt.Println(expirydomain1)

	c.JSON(http.StatusOK, gin.H{
		"domain_name":            Domain1,
		"ssl_expiry_date":        expirydomain1,
		"remaining_days_expired": int64(delta1.Hours() / 24),
	})
}

func SSLList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ssl list is reachable",
	})
}
