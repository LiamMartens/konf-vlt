package konfvlt

import (
	"strings"

	"github.com/LiamMartens/konf-vlt/deepmap"
	"github.com/howellzach/vlt-go"
)

type VaultConfig struct {
	OrganizationId  string
	ProjectId       string
	ApplicationName string
	ClientId        string
	ClientSecret    string
}

type ProviderConfig struct {
	SplitUnderscore bool
}

type VaultSecretsProvider struct {
	ProviderConfig
	vaultClient *vlt.Client
}

func New(vaultConfig VaultConfig, providerConfig ProviderConfig) (*VaultSecretsProvider, error) {
	vault_client, err := vlt.NewClient(vaultConfig.OrganizationId, vaultConfig.ProjectId, vaultConfig.ApplicationName, vaultConfig.ClientId, vaultConfig.ClientSecret)
	if err != nil {
		return nil, err
	}

	provider := new(VaultSecretsProvider)
	provider.ProviderConfig = providerConfig
	provider.vaultClient = &vault_client

	return provider, nil
}

func (vault *VaultSecretsProvider) Load() (map[string]any, error) {
	secrets, err := vault.vaultClient.GetAllSecrets()
	if err != nil {
		return nil, err
	}
	values := make(map[string]any)
	for _, secret := range secrets {
		if vault.SplitUnderscore {
			keys := strings.Split(secret.Name, "_")
			if len(keys) == 0 && keys[0] == "" {
				continue
			}
			deepmap.DeepInsert(values, keys, secret.Version.Value)
		} else {
			values[secret.Name] = secret.Version.Value
		}
	}
	return values, nil
}
