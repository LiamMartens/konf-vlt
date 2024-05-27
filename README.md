# konfvlt
Can be used as a provider for [konf](https://github.com/nil-go/konf)
*Does not support watching at this time*

## Usage
```go
package main

import (
	"os"

	konfvlt "github.com/LiamMartens/konf-vlt"
	"github.com/nil-go/konf"
)

func main() {
	var config konf.Config
	vault_config := konfvlt.VaultConfig{
		OrganizationId:  os.Getenv("ORGANIZATION_ID"),
		ProjectId:       os.Getenv("PROJECT_ID"),
		ApplicationName: os.Getenv("APPLICATION_NAME"),
		ClientId:        os.Getenv("CLIENT_ID"),
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
	}
	vlt_provider, err := konfvlt.New(vault_config, konfvlt.ProviderConfig{
		SplitUnderscore: true,
	})
	if err != nil {
		t.Fatalf("failed to authenticate with VLT (%s)", err.Error())
	}
	config.Load(vlt_provider)
	konf.SetDefault(&config)
}
```
