package core

import (
	"crypto/tls"
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/model"
	"go-ssl-checker/service/notification"
	"log"
	"net"
	"strconv"
	"time"
)

func SSLExpireCheck(manifests []model.Inventory) map[string]model.Expiry_data {
	report := 0
	list := make(map[string]model.Expiry_data)

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
		list["tes"] = model.Expiry_data{
			Domainname:    v.Domainname,
			Expireddate:   expireddate,
			Remainingdays: deltainteger,
		}

	}

	fmt.Printf("we got %v reports of near expired domain \n", report)
	return list

}
