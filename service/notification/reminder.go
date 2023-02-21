package notification

import (
	"fmt"
	"go-ssl-checker/config"
)

func RemindUpdate(c string) {
	fmt.Println("sending slack notification for domain : " + c + " cause it hit the threshold of " + config.Common.Threshold + " days")
}
