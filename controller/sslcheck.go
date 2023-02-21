package controller

import (
	"crypto/tls"
	"fmt"
	"go-ssl-checker/service/manifest"
	"go-ssl-checker/service/notification"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SSLCheck(c *gin.Context) {
	manifests, err := manifest.ParseInventory(false)
	if err != nil {
		fmt.Println("error while parsing inventory on ssl check")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error on controller/sslcheck.go function SSLCheck": err,
		})
		return
	}

	for _, v := range manifests {
		//debug domain name
		fmt.Println("running check on : " + v.Domainname)

		conn, err := net.Dial("tcp", v.Domainname+":"+v.Domainport)
		if err != nil {
			fmt.Println("error while netdial tcp")
			log.Fatal(err)
		}

		client := tls.Client(conn, &tls.Config{
			ServerName: v.Domainname,
		})
		defer client.Close()

		if err := client.Handshake(); err != nil {
			fmt.Println("error while client handshake")
			log.Fatal(err)
		}

		cert := client.ConnectionState().PeerCertificates[0]
		expireddate := cert.NotAfter.Format(time.RFC3339)
		currenttime := time.Now()
		expiredelta := cert.NotAfter.Sub(currenttime)
		deltainteger := int64(expiredelta.Hours() / 24)

		if deltainteger > 300 {
			notification.RemindUpdate(v.Domainname)
		}

		c.JSON(http.StatusOK, gin.H{
			"domain_name":            v.Domainname,
			"ssl_expiry_date":        expireddate,
			"remaining_days_expired": deltainteger,
		})
		client.Close()
	}
}

func SSLList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ssl list is reachable",
	})
}
