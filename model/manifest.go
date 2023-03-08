package model

type Inventory struct {
	Domainname string `yaml:"domain_name" json:"domain_name"`
	Domainport string `yaml:"domain_port" json:"domain_port"`
	Servername string `yaml:"server_name" json:"server_name"`
}

func (m *Inventory) Mask() *Inventory {
	return &Inventory{
		Domainname: m.Domainname,
		Domainport: m.Domainport,
		Servername: m.Servername,
	}
}
