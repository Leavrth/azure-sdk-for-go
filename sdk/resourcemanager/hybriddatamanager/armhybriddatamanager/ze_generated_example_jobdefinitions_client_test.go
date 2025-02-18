//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_ListByDataService-GET-example-71.json
func ExampleJobDefinitionsClient_ListByDataService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.ListByDataService("<data-service-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobDefinitionsListByDataServiceOptions{Filter: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("JobDefinition.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_Get-GET-example-81.json
func ExampleJobDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JobDefinition.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_CreateOrUpdate-PUT-example-83.json
func ExampleJobDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		armhybriddatamanager.JobDefinition{
			Properties: &armhybriddatamanager.JobDefinitionProperties{
				DataServiceInput: map[string]interface{}{
					"AzureStorageType": "Blob",
					"BackupChoice":     "UseExistingLatest",
					"ContainerName":    "containerfromtest",
					"DeviceName":       "8600-SHG0997877L71FC",
					"FileNameFilter":   "*",
					"IsDirectoryMode":  false,
					"RootDirectories": []interface{}{
						"\\",
					},
					"VolumeNames": []interface{}{
						"TestAutomation",
					},
				},
				DataSinkID:       to.StringPtr("<data-sink-id>"),
				DataSourceID:     to.StringPtr("<data-source-id>"),
				RunLocation:      armhybriddatamanager.RunLocationWestus.ToPtr(),
				State:            armhybriddatamanager.StateEnabled.ToPtr(),
				UserConfirmation: armhybriddatamanager.UserConfirmationRequired.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JobDefinition.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_Delete-DELETE-example-81.json
func ExampleJobDefinitionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_Run-POST-example-132.json
func ExampleJobDefinitionsClient_BeginRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRun(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		armhybriddatamanager.RunParameters{
			CustomerSecrets: []*armhybriddatamanager.CustomerSecret{},
			DataServiceInput: map[string]interface{}{
				"AzureStorageType": "Blob",
				"BackupChoice":     "UseExistingLatest",
				"ContainerName":    "containerfromtest",
				"DeviceName":       "8600-SHG0997877L71FC",
				"FileNameFilter":   "*",
				"IsDirectoryMode":  false,
				"RootDirectories": []interface{}{
					"\\",
				},
				"VolumeNames": []interface{}{
					"TestAutomation",
				},
			},
			UserConfirmation: armhybriddatamanager.UserConfirmationNotRequired.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_ListByDataManager-GET-example-191.json
func ExampleJobDefinitionsClient_ListByDataManager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.ListByDataManager("<resource-group-name>",
		"<data-manager-name>",
		&armhybriddatamanager.JobDefinitionsListByDataManagerOptions{Filter: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("JobDefinition.ID: %s\n", *v.ID)
		}
	}
}
