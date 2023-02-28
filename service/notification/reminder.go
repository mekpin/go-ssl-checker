package notification

import (
	"fmt"
	"go-ssl-checker/config"
)

func RemindUpdate(alert string) {
	fmt.Println("sending slack notification for domain : " + alert + " cause it hit the threshold of " + config.Common.Threshold + " days")

}
