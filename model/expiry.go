package model

type ExpiryData struct {
	Domainname    string `json:"domain_name"`
	Expireddate   string `json:"expired_date"`
	Remainingdays int    `json:"remaining_days"`
	Servername    string `json:"server_name"`
	Domainport    string `json:"domain_port"`
}

type ReminderData struct {
	Domainname    string `json:"domain_name"`
	Expireddate   string `json:"expired_date"`
	Remainingdays int    `json:"remaining_days"`
	Servername    string `json:"server_name"`
	Domainport    string `json:"domain_port"`
}
