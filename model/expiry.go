package model

type ExpiryData struct {
	Domainname    string `json:"domain_name"`
	Expireddate   string `json:"expired_date"`
	Remainingdays int    `json:"remaining_days"`
}

// type DataInventory struct {
// 	Id       int          `json:"id"`
//  Datapool []ExpiryData `json:"data_pool,omitempty"`
// }
