package controller

import (
	"crypto/tls"
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/model"
	"go-ssl-checker/service/manifest"
	"go-ssl-checker/service/notification"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SSLCheck(c *gin.Context) {
	manifests, err := manifest.ParseInventory(false)
	if err != nil {
		fmt.Println("error while parsing inventory on ssl check in file controller/sslcheck.go")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error on controller/sslcheck.go function SSLCheck": err,
		})
		return
	}
	report := 0

	for _, v := range manifests {
		//debug domain name
		fmt.Println("running check on : " + v.Domainname)

		conn, err := net.Dial("tcp", v.Domainname+":"+v.Domainport)
		if err != nil {
			fmt.Println("error while netdial tcp in file controller/sslcheck.go")
			log.Fatal(err)
		}

		client := tls.Client(conn, &tls.Config{
			ServerName: v.Domainname,
		})
		defer client.Close()

		if err := client.Handshake(); err != nil {
			fmt.Println("error while client handshake in file controller/sslcheck.go")
			log.Fatal(err)
		}

		cert := client.ConnectionState().PeerCertificates[0]
		expireddate := cert.NotAfter.Format(time.RFC3339)
		currenttime := time.Now()
		expiredelta := cert.NotAfter.Sub(currenttime)
		deltainteger := int(expiredelta.Hours() / 24)

		//ubah env threshold jadi int untuk compare
		intthreshold, err := strconv.Atoi(config.Common.Threshold)
		if err != nil {
			fmt.Println("error while changing env threshold string to int in file controller/sslcheck.go")
			log.Fatal(err)
		}
		if deltainteger < intthreshold {
			notification.RemindUpdate(v.Domainname)
			report = report + 1
		}

		list := &model.Expiry_data{
			Domainname:    v.Domainname,
			Expireddate:   expireddate,
			Remainingdays: deltainteger,
		}

		c.JSON(http.StatusOK, list)

		// membuat json manual
		// c.JSON(http.StatusOK, gin.H{
		// 	"domain_name":            v.Domainname,
		// 	"ssl_expiry_date":        expireddate,
		// 	"remaining_days_expired": deltainteger,
		// })
		client.Close()
	}

	fmt.Printf("we got %v reports of near expired domain \n", report)
}

func SSLList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ssl list is reachable",
	})
}
