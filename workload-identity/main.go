package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"time"
)

func main() {
	fmt.Println("Starting...")
	ctx := context.Background()

	for {

	}
	cred, err := azidentity.NewWorkloadIdentityCredential(nil)
	if err != nil {
		panic(err)
	}

	kvClient, err := azsecrets.NewClient("https://dedicated-kv-dev.vault.azure.net/", cred, nil)
	if err != nil {
		panic(err)
	}

	resp, err := kvClient.GetSecret(ctx, "cloudflare-zoneid", "", nil)
	if err != nil {
		panic(err)
	}

	cloudflareZoneID := *resp.Value

	fmt.Println(cloudflareZoneID)

	time.Sleep(24 * time.Hour)
}
