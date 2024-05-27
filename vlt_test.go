package konfvlt_test

import (
	"os"
	"testing"

	konfvlt "github.com/LiamMartens/konf-vlt"
	"github.com/nil-go/konf"
)

func TestKonfProvider(t *testing.T) {
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

	if konf.Get[string]("shallowkey") != "value" {
		t.Fatalf("shallow key not loaded: %q", "shallowkey")
	}

	if konf.Get[string]("foo.bar.baz") != "deepvalue" {
		t.Fatalf("deep key not loaded: %q", "foo.bar.baz")
	}
}
