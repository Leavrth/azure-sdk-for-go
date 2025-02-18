//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationGatewayPrivateEndpointConnectionsClient contains the methods for the ApplicationGatewayPrivateEndpointConnections group.
// Don't use this type directly, use NewApplicationGatewayPrivateEndpointConnectionsClient() instead.
type ApplicationGatewayPrivateEndpointConnectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewApplicationGatewayPrivateEndpointConnectionsClient creates a new instance of ApplicationGatewayPrivateEndpointConnectionsClient with the specified values.
func NewApplicationGatewayPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ApplicationGatewayPrivateEndpointConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ApplicationGatewayPrivateEndpointConnectionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginDelete - Deletes the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsBeginDeleteOptions) (ApplicationGatewayPrivateEndpointConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsDeletePollerResponse{}, err
	}
	result := ApplicationGatewayPrivateEndpointConnectionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationGatewayPrivateEndpointConnectionsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &ApplicationGatewayPrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsGetOptions) (ApplicationGatewayPrivateEndpointConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGatewayPrivateEndpointConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (ApplicationGatewayPrivateEndpointConnectionsGetResponse, error) {
	result := ApplicationGatewayPrivateEndpointConnectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGatewayPrivateEndpointConnection); err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all private endpoint connections on an application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) List(resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateEndpointConnectionsListOptions) *ApplicationGatewayPrivateEndpointConnectionsListPager {
	return &ApplicationGatewayPrivateEndpointConnectionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, applicationGatewayName, options)
		},
		advancer: func(ctx context.Context, resp ApplicationGatewayPrivateEndpointConnectionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationGatewayPrivateEndpointConnectionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateEndpointConnectionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listHandleResponse(resp *http.Response) (ApplicationGatewayPrivateEndpointConnectionsListResponse, error) {
	result := ApplicationGatewayPrivateEndpointConnectionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGatewayPrivateEndpointConnectionListResult); err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsBeginUpdateOptions) (ApplicationGatewayPrivateEndpointConnectionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters, options)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsUpdatePollerResponse{}, err
	}
	result := ApplicationGatewayPrivateEndpointConnectionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationGatewayPrivateEndpointConnectionsClient.Update", "azure-async-operation", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsUpdatePollerResponse{}, err
	}
	result.Poller = &ApplicationGatewayPrivateEndpointConnectionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the specified private endpoint connection on application gateway.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) update(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
