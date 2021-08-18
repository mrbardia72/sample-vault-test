package main

import (
	"testing"

	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/vault"
)

const (
	testVaultToken = "myroot"
)

func TestPutSecret(t *testing.T) {
	cluster := vault.NewTestCluster(t, &vault.CoreConfig{
		DevToken: testVaultToken,
	}, &vault.TestClusterOptions{
		HandlerFunc: vaulthttp.Handler,
	})
	cluster.Start()
	defer cluster.Cleanup()

	core := cluster.Cores[0].Core
	vault.TestWaitActive(t, core)
	client := cluster.Cores[0].Client

	err := putSecret(client, map[string]interface{}{"kazemi": "bardia"}, "secret")
	if err != nil {
		t.Fatal(err)
	}

	data, err := client.Logical().Read("secret/data/secret")
	if err != nil {
		t.Fatal(err)
	}

	if secret, ok := data.Data["kazemi"].(string); ok {
		if secret != "bardia" {
			t.Fatalf("Wrong secret returned: %s", secret)
		}
	} else {
		t.Fatal("Could not get secret")
	}
}
