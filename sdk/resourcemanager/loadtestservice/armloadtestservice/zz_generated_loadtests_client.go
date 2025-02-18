//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armloadtestservice

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

// LoadTestsClient contains the methods for the LoadTests group.
// Don't use this type directly, use NewLoadTestsClient() instead.
type LoadTestsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLoadTestsClient creates a new instance of LoadTestsClient with the specified values.
func NewLoadTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LoadTestsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &LoadTestsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update LoadTest resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResource LoadTestResource, options *LoadTestsCreateOrUpdateOptions) (LoadTestsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadTestName, loadTestResource, options)
	if err != nil {
		return LoadTestsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadTestsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadTestsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResource LoadTestResource, options *LoadTestsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, loadTestResource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LoadTestsClient) createOrUpdateHandleResponse(resp *http.Response) (LoadTestsCreateOrUpdateResponse, error) {
	result := LoadTestsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResource); err != nil {
		return LoadTestsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LoadTestsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete a LoadTest resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) BeginDelete(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsBeginDeleteOptions) (LoadTestsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, loadTestName, options)
	if err != nil {
		return LoadTestsDeletePollerResponse{}, err
	}
	result := LoadTestsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LoadTestsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return LoadTestsDeletePollerResponse{}, err
	}
	result.Poller = &LoadTestsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a LoadTest resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) deleteOperation(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadTestName, options)
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
func (client *LoadTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LoadTestsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a LoadTest resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) Get(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsGetOptions) (LoadTestsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadTestName, options)
	if err != nil {
		return LoadTestsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadTestsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadTestsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadTestsClient) getHandleResponse(resp *http.Response) (LoadTestsGetResponse, error) {
	result := LoadTestsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResource); err != nil {
		return LoadTestsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LoadTestsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists loadtest resources in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) ListByResourceGroup(resourceGroupName string, options *LoadTestsListByResourceGroupOptions) *LoadTestsListByResourceGroupPager {
	return &LoadTestsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp LoadTestsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LoadTestResourcePageList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LoadTestsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LoadTestsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LoadTestsClient) listByResourceGroupHandleResponse(resp *http.Response) (LoadTestsListByResourceGroupResponse, error) {
	result := LoadTestsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResourcePageList); err != nil {
		return LoadTestsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *LoadTestsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists loadtests resources in a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) ListBySubscription(options *LoadTestsListBySubscriptionOptions) *LoadTestsListBySubscriptionPager {
	return &LoadTestsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LoadTestsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LoadTestResourcePageList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LoadTestsClient) listBySubscriptionCreateRequest(ctx context.Context, options *LoadTestsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LoadTestService/loadTests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LoadTestsClient) listBySubscriptionHandleResponse(resp *http.Response) (LoadTestsListBySubscriptionResponse, error) {
	result := LoadTestsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResourcePageList); err != nil {
		return LoadTestsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *LoadTestsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update a loadtest resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoadTestsClient) Update(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResourcePatchRequestBody LoadTestResourcePatchRequestBody, options *LoadTestsUpdateOptions) (LoadTestsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, loadTestName, loadTestResourcePatchRequestBody, options)
	if err != nil {
		return LoadTestsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadTestsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadTestsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *LoadTestsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResourcePatchRequestBody LoadTestResourcePatchRequestBody, options *LoadTestsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, loadTestResourcePatchRequestBody)
}

// updateHandleResponse handles the Update response.
func (client *LoadTestsClient) updateHandleResponse(resp *http.Response) (LoadTestsUpdateResponse, error) {
	result := LoadTestsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResource); err != nil {
		return LoadTestsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *LoadTestsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
