package nutanix

import (
	"context"
	"fmt"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/provider"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

type nutanixProvider struct {
	Log logrus.FieldLogger
}

// NewNutanixProvider creates a new vSphere provider.
func NewNutanixProvider(log logrus.FieldLogger) provider.Provider {
	return &nutanixProvider{
		Log: log,
	}
}

// Name returns the name of the provider
func (p *nutanixProvider) Name() models.PlatformType {
	return models.PlatformTypeNutanix
}

func (p *nutanixProvider) IsHostSupported(ctx context.Context, host *models.Host) (bool, error) {
	// during the discovery there is a short time that host didn't return its inventory to the service
	if host.Inventory == "" {
		return false, nil
	}
	hostInventory, err := common.UnmarshalInventory(context.Background(), host.Inventory)
	if err != nil {
		return false, fmt.Errorf("error marshaling host to inventory, error %w", err)
	}
	return hostInventory.SystemVendor.Manufacturer == NutanixManufacturer, nil
}

func (p *nutanixProvider) AreHostsSupported(hosts []*models.Host) (bool, error) {
	for _, h := range hosts {
		supported, err := p.IsHostSupported(context.Background(), h)
		if err != nil {
			return false, fmt.Errorf("error while checking if host is supported, error is: %w", err)
		}
		if !supported {
			return false, nil
		}
	}
	return true, nil
}

func (p *nutanixProvider) IsProviderForPlatform(platform *models.Platform) bool {
	return platform != nil &&
		platform.Type != nil &&
		*platform.Type == p.Name()
}
