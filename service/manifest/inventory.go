package manifest

import (
	"go-ssl-checker/config"
	"go-ssl-checker/model"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseInventory(mask bool) ([]model.Inventory, error) {
	f, err := os.ReadFile(config.Manifest.InventoryPath)
	if err != nil {
		return nil, err
	}

	var inventories []model.Inventory

	if err = yaml.Unmarshal(f, &inventories); err != nil {
		return nil, err
	}

	return inventories, nil

	// list := make(map[string]model.Inventory)

	// for _, v := range inventories {
	// 	if mask {
	// 		list[v.Domainname] = model.Inventory{
	// 			Domainname: v.Domainname,
	// 			Domainport: v.Domainport,
	// 		}
	// 	} else {
	// 		// fmt.Printf("domain_name", ":", "%s", list[v.Domainname])
	// 		// fmt.Printf("domain_port", ":", "%s", list[v.Domainport])
	// 		list[v.Domainname] = model.Inventory{
	// 			Domainname: v.Domainname,
	// 			Domainport: v.Domainport,
	// 		}
	// 	}
	// }

	// log.Debug().Interface("inventory", list).Send()
	// return list, nil
}
