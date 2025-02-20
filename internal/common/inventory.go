package common

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/elastic/go-freelru"
	"github.com/openshift/assisted-service/models"
	"github.com/zeebo/xxh3"
)

var once sync.Once
var InventoryCache *freelru.SyncedLRU[string, models.Inventory]

func hashStringXXH3HASH(s string) uint32 {
	return uint32(xxh3.HashString(s))
}

func MarshalInventory(ctx context.Context, inventory *models.Inventory) (string, error) {

	if data, err := json.Marshal(inventory); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func UnmarshalInventory(ctx context.Context, inventoryStr string) (*models.Inventory, error) {
	once.Do(func() {
		var err error
		InventoryCache, err = freelru.NewSynced[string, models.Inventory](64, hashStringXXH3HASH)
		if err != nil {
			panic(err)
		}
	})

	inventory, ok := InventoryCache.Get(inventoryStr)
	if ok {
		return &inventory, nil
	}

	err := sonic.UnmarshalString(inventoryStr, &inventory)
	if err != nil {
		return nil, err
	}

	InventoryCache.Add(inventoryStr, inventory)

	return &inventory, nil
}
