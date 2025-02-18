//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

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

// BotConnectionClient contains the methods for the BotConnection group.
// Don't use this type directly, use NewBotConnectionClient() instead.
type BotConnectionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewBotConnectionClient creates a new instance of BotConnectionClient with the specified values.
func NewBotConnectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BotConnectionClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BotConnectionClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Register a new Auth Connection for a Bot Service
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) Create(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionCreateOptions) (BotConnectionCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, connectionName, parameters, options)
	if err != nil {
		return BotConnectionCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotConnectionCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *BotConnectionClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *BotConnectionClient) createHandleResponse(resp *http.Response) (BotConnectionCreateResponse, error) {
	result := BotConnectionCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *BotConnectionClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a Connection Setting registration for a Bot Service
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionDeleteOptions) (BotConnectionDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, connectionName, options)
	if err != nil {
		return BotConnectionDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return BotConnectionDeleteResponse{}, client.deleteHandleError(resp)
	}
	return BotConnectionDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BotConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *BotConnectionClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a Connection Setting registration for a Bot Service
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) Get(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionGetOptions) (BotConnectionGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, connectionName, options)
	if err != nil {
		return BotConnectionGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotConnectionGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BotConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BotConnectionClient) getHandleResponse(resp *http.Response) (BotConnectionGetResponse, error) {
	result := BotConnectionGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BotConnectionClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByBotService - Returns all the Connection Settings registered to a particular BotService resource
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) ListByBotService(resourceGroupName string, resourceName string, options *BotConnectionListByBotServiceOptions) *BotConnectionListByBotServicePager {
	return &BotConnectionListByBotServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBotServiceCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp BotConnectionListByBotServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConnectionSettingResponseList.NextLink)
		},
	}
}

// listByBotServiceCreateRequest creates the ListByBotService request.
func (client *BotConnectionClient) listByBotServiceCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *BotConnectionListByBotServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBotServiceHandleResponse handles the ListByBotService response.
func (client *BotConnectionClient) listByBotServiceHandleResponse(resp *http.Response) (BotConnectionListByBotServiceResponse, error) {
	result := BotConnectionListByBotServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSettingResponseList); err != nil {
		return BotConnectionListByBotServiceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByBotServiceHandleError handles the ListByBotService error response.
func (client *BotConnectionClient) listByBotServiceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListServiceProviders - Lists the available Service Providers for creating Connection Settings
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) ListServiceProviders(ctx context.Context, options *BotConnectionListServiceProvidersOptions) (BotConnectionListServiceProvidersResponse, error) {
	req, err := client.listServiceProvidersCreateRequest(ctx, options)
	if err != nil {
		return BotConnectionListServiceProvidersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionListServiceProvidersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotConnectionListServiceProvidersResponse{}, client.listServiceProvidersHandleError(resp)
	}
	return client.listServiceProvidersHandleResponse(resp)
}

// listServiceProvidersCreateRequest creates the ListServiceProviders request.
func (client *BotConnectionClient) listServiceProvidersCreateRequest(ctx context.Context, options *BotConnectionListServiceProvidersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BotService/listAuthServiceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listServiceProvidersHandleResponse handles the ListServiceProviders response.
func (client *BotConnectionClient) listServiceProvidersHandleResponse(resp *http.Response) (BotConnectionListServiceProvidersResponse, error) {
	result := BotConnectionListServiceProvidersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceProviderResponseList); err != nil {
		return BotConnectionListServiceProvidersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listServiceProvidersHandleError handles the ListServiceProviders error response.
func (client *BotConnectionClient) listServiceProvidersHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListWithSecrets - Get a Connection Setting registration for a Bot Service
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) ListWithSecrets(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionListWithSecretsOptions) (BotConnectionListWithSecretsResponse, error) {
	req, err := client.listWithSecretsCreateRequest(ctx, resourceGroupName, resourceName, connectionName, options)
	if err != nil {
		return BotConnectionListWithSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionListWithSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotConnectionListWithSecretsResponse{}, client.listWithSecretsHandleError(resp)
	}
	return client.listWithSecretsHandleResponse(resp)
}

// listWithSecretsCreateRequest creates the ListWithSecrets request.
func (client *BotConnectionClient) listWithSecretsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionListWithSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}/listWithSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listWithSecretsHandleResponse handles the ListWithSecrets response.
func (client *BotConnectionClient) listWithSecretsHandleResponse(resp *http.Response) (BotConnectionListWithSecretsResponse, error) {
	result := BotConnectionListWithSecretsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionListWithSecretsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listWithSecretsHandleError handles the ListWithSecrets error response.
func (client *BotConnectionClient) listWithSecretsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates a Connection Setting registration for a Bot Service
// If the operation fails it returns the *Error error type.
func (client *BotConnectionClient) Update(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionUpdateOptions) (BotConnectionUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, connectionName, parameters, options)
	if err != nil {
		return BotConnectionUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotConnectionUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *BotConnectionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *BotConnectionClient) updateHandleResponse(resp *http.Response) (BotConnectionUpdateResponse, error) {
	result := BotConnectionUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *BotConnectionClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
