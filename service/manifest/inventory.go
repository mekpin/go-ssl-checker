package manifest

import (
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/model"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func ParseInventory(mask bool) ([]model.Inventory, error) {
	// //debug
	// fmt.Println("running parseinventory")
	f, err := os.ReadFile(config.Manifest.InventoryPath)
	if err != nil {
		log.Info().Str("message", "error while reading manifest file").Send()

		return nil, err
	}

	var inventories []model.Inventory
	// //debug
	// fmt.Println("running yml unmarshal")
	if err = yaml.Unmarshal(f, &inventories); err != nil {
		log.Info().Str("message", "error while unmarshalling inventories").Send()
		return nil, err
	}
	//debug
	fmt.Println("returning inventories")
	return inventories, nil
}
