package notification

import (
	"fmt"
	"time"
)

func DailyCheck() {
	jalanjam := time.Now()
	fmt.Println("running ssl daily check on " + jalanjam.String())
}
