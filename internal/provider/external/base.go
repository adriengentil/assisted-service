package external

import (
	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/internal/installcfg"
	"github.com/openshift/assisted-service/internal/provider"
	"github.com/openshift/assisted-service/internal/usage"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

// baseExternalProvider provides a default implementation suitable for platforms relying on the external platform.
// Compose it and implement Name() to fullfil the Provider interface.
type baseExternalProvider struct {
	provider.Provider
	Log logrus.FieldLogger
}

func (p *baseExternalProvider) IsHostSupported(_ *models.Host) (bool, error) {
	return true, nil
}

func (p *baseExternalProvider) AreHostsSupported(hosts []*models.Host) (bool, error) {
	return true, nil
}

func (p baseExternalProvider) PreCreateManifestsHook(cluster *common.Cluster, envVars *[]string, workDir string) error {
	return nil
}

func (p baseExternalProvider) PostCreateManifestsHook(cluster *common.Cluster, envVars *[]string, workDir string) error {
	return nil
}

func (p baseExternalProvider) AddPlatformToInstallConfig(cfg *installcfg.InstallerConfigBaremetal, cluster *common.Cluster) error {
	cfg.Platform = installcfg.Platform{
		External: &installcfg.ExternalInstallConfigPlatform{
			PlatformName:           cluster.Cluster.Platform.External.PlatformName,
			CloudControllerManager: installcfg.CloudControllerManager(cluster.Cluster.Platform.External.CloudControllerManager),
		},
	}

	cfg.Networking.MachineNetwork = provider.GetMachineNetworkForUserManagedNetworking(p.Log, cluster)
	if cluster.NetworkType != nil {
		cfg.Networking.NetworkType = swag.StringValue(cluster.NetworkType)
	}

	if common.IsSingleNodeCluster(cluster) {

		if cfg.Networking.NetworkType == "" {
			cfg.Networking.NetworkType = models.ClusterNetworkTypeOVNKubernetes
		}

		bootstrap := common.GetBootstrapHost(cluster)
		if bootstrap != nil {
			cfg.BootstrapInPlace = installcfg.BootstrapInPlace{InstallationDisk: hostutil.GetHostInstallationPath(bootstrap)}
		}
	}

	return nil
}

func (p *baseExternalProvider) CleanPlatformValuesFromDBUpdates(_ map[string]interface{}) error {
	return nil
}

func (p *baseExternalProvider) SetPlatformUsages(
	usages map[string]models.Usage,
	usageApi usage.API) error {
	props := &map[string]interface{}{
		"platform_type": p.Provider.Name()}
	usageApi.Add(usages, usage.PlatformSelectionUsage, props)
	return nil
}
