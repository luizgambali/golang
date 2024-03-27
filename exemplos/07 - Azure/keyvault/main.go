//reference: https://learn.microsoft.com/pt-br/azure/key-vault/keys/quick-create-go

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
)

func main() {
	keyVaultName := os.Getenv("KEY_VAULT_NAME")
	keyVaultUrl := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)

	// create credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// create azkeys client
	client, err := azkeys.NewClient(keyVaultUrl, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	// create RSA Key
	rsaKeyParams := azkeys.CreateKeyParameters{
		Kty:     to.Ptr(azkeys.JSONWebKeyTypeRSA),
		KeySize: to.Ptr(int32(2048)),
	}
	rsaResp, err := client.CreateKey(context.TODO(), "new-rsa-key", rsaKeyParams, nil)
	if err != nil {
		log.Fatalf("failed to create rsa key: %v", err)
	}
	fmt.Printf("New RSA key ID: %s\n", *rsaResp.Key.KID)

	// create EC Key
	ecKeyParams := azkeys.CreateKeyParameters{
		Kty:   to.Ptr(azkeys.JSONWebKeyTypeEC),
		Curve: to.Ptr(azkeys.JSONWebKeyCurveNameP256),
	}
	ecResp, err := client.CreateKey(context.TODO(), "new-ec-key", ecKeyParams, nil)
	if err != nil {
		log.Fatalf("failed to create ec key: %v", err)
	}
	fmt.Printf("New EC key ID: %s\n", *ecResp.Key.KID)

	// list all vault keys
	fmt.Println("List all vault keys:")
	pager := client.NewListKeysPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, key := range page.Value {
			fmt.Println(*key.KID)
		}
	}

	// update key properties to disable key
	updateParams := azkeys.UpdateKeyParameters{
		KeyAttributes: &azkeys.KeyAttributes{
			Enabled: to.Ptr(false),
		},
	}
	// an empty string version updates the latest version of the key
	version := ""
	updateResp, err := client.UpdateKey(context.TODO(), "new-rsa-key", version, updateParams, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Key %s Enabled attribute set to: %t\n", *updateResp.Key.KID, *updateResp.Attributes.Enabled)

	// delete the created keys
	for _, keyName := range []string{"new-rsa-key", "new-ec-key"} {
		// DeleteKey returns when Key Vault has begun deleting the key. That can take several
		// seconds to complete, so it may be necessary to wait before performing other operations
		// on the deleted key.
		delResp, err := client.DeleteKey(context.TODO(), keyName, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully deleted key %s", *delResp.Key.KID)
	}
}
