package kusto

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EventHubConnectionsClient is the the Azure Kusto management API provides a RESTful set of web services that interact
// with Azure Kusto services to manage your clusters and databases. The API enables you to create, update, and delete
// clusters and databases.
type EventHubConnectionsClient struct {
	BaseClient
}

// NewEventHubConnectionsClient creates an instance of the EventHubConnectionsClient client.
func NewEventHubConnectionsClient(subscriptionID string) EventHubConnectionsClient {
	return NewEventHubConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEventHubConnectionsClientWithBaseURI creates an instance of the EventHubConnectionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewEventHubConnectionsClientWithBaseURI(baseURI string, subscriptionID string) EventHubConnectionsClient {
	return EventHubConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a Event Hub connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// eventHubConnectionName - the name of the event hub connection.
// parameters - the Event Hub connection parameters supplied to the CreateOrUpdate operation.
func (client EventHubConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string, parameters EventHubConnection) (result EventHubConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.EventHubConnectionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.EventHubConnectionProperties.EventHubResourceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.EventHubConnectionProperties.ConsumerGroup", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("kusto.EventHubConnectionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, eventHubConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EventHubConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string, parameters EventHubConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":            autorest.Encode("path", clusterName),
		"databaseName":           autorest.Encode("path", databaseName),
		"eventHubConnectionName": autorest.Encode("path", eventHubConnectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/eventhubconnections/{eventHubConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubConnectionsClient) CreateOrUpdateSender(req *http.Request) (future EventHubConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EventHubConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result EventHubConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the Event Hub connection with the given name.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// eventHubConnectionName - the name of the event hub connection.
func (client EventHubConnectionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string) (result EventHubConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, databaseName, eventHubConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EventHubConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":            autorest.Encode("path", clusterName),
		"databaseName":           autorest.Encode("path", databaseName),
		"eventHubConnectionName": autorest.Encode("path", eventHubConnectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/eventhubconnections/{eventHubConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubConnectionsClient) DeleteSender(req *http.Request) (future EventHubConnectionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EventHubConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EventhubConnectionValidation checks that the Event Hub data connection parameters are valid.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// parameters - the Event Hub connection parameters supplied to the CreateOrUpdate operation.
func (client EventHubConnectionsClient) EventhubConnectionValidation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters EventHubConnectionValidation) (result EventHubConnectionValidationListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubConnectionsClient.EventhubConnectionValidation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.EventHubConnectionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.EventHubConnectionProperties.EventHubResourceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.EventHubConnectionProperties.ConsumerGroup", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("kusto.EventHubConnectionsClient", "EventhubConnectionValidation", err.Error())
	}

	req, err := client.EventhubConnectionValidationPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "EventhubConnectionValidation", nil, "Failure preparing request")
		return
	}

	resp, err := client.EventhubConnectionValidationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "EventhubConnectionValidation", resp, "Failure sending request")
		return
	}

	result, err = client.EventhubConnectionValidationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "EventhubConnectionValidation", resp, "Failure responding to request")
		return
	}

	return
}

// EventhubConnectionValidationPreparer prepares the EventhubConnectionValidation request.
func (client EventHubConnectionsClient) EventhubConnectionValidationPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters EventHubConnectionValidation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/eventhubConnectionValidation", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EventhubConnectionValidationSender sends the EventhubConnectionValidation request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubConnectionsClient) EventhubConnectionValidationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EventhubConnectionValidationResponder handles the response to the EventhubConnectionValidation request. The method always
// closes the http.Response Body.
func (client EventHubConnectionsClient) EventhubConnectionValidationResponder(resp *http.Response) (result EventHubConnectionValidationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns an Event Hub connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// eventHubConnectionName - the name of the event hub connection.
func (client EventHubConnectionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string) (result EventHubConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, databaseName, eventHubConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EventHubConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":            autorest.Encode("path", clusterName),
		"databaseName":           autorest.Encode("path", databaseName),
		"eventHubConnectionName": autorest.Encode("path", eventHubConnectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/eventhubconnections/{eventHubConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EventHubConnectionsClient) GetResponder(resp *http.Response) (result EventHubConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase returns the list of Event Hub connections of the given Kusto database.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
func (client EventHubConnectionsClient) ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result EventHubConnectionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubConnectionsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client EventHubConnectionsClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/eventhubconnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubConnectionsClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client EventHubConnectionsClient) ListByDatabaseResponder(resp *http.Response) (result EventHubConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a Event Hub connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// eventHubConnectionName - the name of the event hub connection.
// parameters - the Event Hub connection parameters supplied to the Update operation.
func (client EventHubConnectionsClient) Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string, parameters EventHubConnectionUpdate) (result EventHubConnectionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubConnectionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, eventHubConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.EventHubConnectionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EventHubConnectionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, eventHubConnectionName string, parameters EventHubConnectionUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":            autorest.Encode("path", clusterName),
		"databaseName":           autorest.Encode("path", databaseName),
		"eventHubConnectionName": autorest.Encode("path", eventHubConnectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-07-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/eventhubconnections/{eventHubConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubConnectionsClient) UpdateSender(req *http.Request) (future EventHubConnectionsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EventHubConnectionsClient) UpdateResponder(resp *http.Response) (result EventHubConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
