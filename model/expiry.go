package model

type Expiry_data struct {
	Domainname    string `json:"domain_name"`
	Expireddate   string `json:"expired_date"`
	Remainingdays int    `json:"remaining_days"`
}
