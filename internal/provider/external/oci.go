package external

import (
	"context"
	"fmt"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/provider"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

const (
	OCIManufacturer string = "OracleCloud.com"
)

type ociExternalProvider struct {
	baseExternalProvider
}

func NewOciExternalProvider(log logrus.FieldLogger) provider.Provider {
	p := &ociExternalProvider{
		baseExternalProvider: baseExternalProvider{
			Log: log,
		},
	}
	p.Provider = p
	return p
}

func (p *ociExternalProvider) IsHostSupported(host *models.Host) (bool, error) {
	return IsOciHost(host)
}

func (p *ociExternalProvider) AreHostsSupported(hosts []*models.Host) (bool, error) {
	for _, h := range hosts {
		supported, err := p.IsHostSupported(h)
		if err != nil {
			return false, fmt.Errorf("error while checking if host is supported, error is: %w", err)
		}
		if !supported {
			return false, nil
		}
	}
	return true, nil
}

func (p *ociExternalProvider) IsProviderForPlatform(platform *models.Platform) bool {
	return common.IsOciExternalIntegrationEnabled(platform)
}

func IsOciHost(host *models.Host) (bool, error) {
	// during the discovery there is a short time that host didn't return its inventory to the service
	if host.Inventory == "" {
		return false, nil
	}
	hostInventory, err := common.UnmarshalInventory(context.Background(), host.Inventory)
	if err != nil {
		return false, fmt.Errorf("error marshaling host to inventory, error %w", err)
	}
	return hostInventory.SystemVendor.Manufacturer == OCIManufacturer, nil
}
