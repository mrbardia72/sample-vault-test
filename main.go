package main

import (
	"fmt"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

// Write the secret
func putSecret(client *api.Client, secret map[string]interface{}, secretPath string) error {
	_, err := client.Logical().Write(fmt.Sprintf("secret/data/%s", secretPath), secret)
	return err
}

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialise vault client")
	}
	err = putSecret(client, map[string]interface{}{"kazemi": "bardia"}, "langroud")
	if err != nil {
		logrus.WithError(err).Fatal("Could not write secret")
	}
}
