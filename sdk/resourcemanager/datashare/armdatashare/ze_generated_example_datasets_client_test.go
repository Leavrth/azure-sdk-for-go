//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_Get.json
func ExampleDataSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		"<data-set-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataSetClassification.GetDataSet().ID: %s\n", *res.GetDataSet().ID)
}

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_Create.json
func ExampleDataSetsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		"<data-set-name>",
		&armdatashare.BlobDataSet{
			DataSet: armdatashare.DataSet{
				Kind: armdatashare.DataSetKindBlob.ToPtr(),
			},
			Properties: &armdatashare.BlobProperties{
				ContainerName:      to.StringPtr("<container-name>"),
				FilePath:           to.StringPtr("<file-path>"),
				ResourceGroup:      to.StringPtr("<resource-group>"),
				StorageAccountName: to.StringPtr("<storage-account-name>"),
				SubscriptionID:     to.StringPtr("<subscription-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataSetClassification.GetDataSet().ID: %s\n", *res.GetDataSet().ID)
}

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_Delete.json
func ExampleDataSetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		"<data-set-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_ListByShare.json
func ExampleDataSetsClient_ListByShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetsClient("<subscription-id>", cred, nil)
	pager := client.ListByShare("<resource-group-name>",
		"<account-name>",
		"<share-name>",
		&armdatashare.DataSetsListByShareOptions{SkipToken: nil,
			Filter:  nil,
			Orderby: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DataSetClassification.GetDataSet().ID: %s\n", *v.GetDataSet().ID)
		}
	}
}
