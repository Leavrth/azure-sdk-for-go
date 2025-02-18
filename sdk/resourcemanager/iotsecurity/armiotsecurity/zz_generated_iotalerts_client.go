//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// IotAlertsClient contains the methods for the IotAlerts group.
// Don't use this type directly, use NewIotAlertsClient() instead.
type IotAlertsClient struct {
	ep                  string
	pl                  runtime.Pipeline
	subscriptionID      string
	iotDefenderLocation string
}

// NewIotAlertsClient creates a new instance of IotAlertsClient with the specified values.
func NewIotAlertsClient(subscriptionID string, iotDefenderLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *IotAlertsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IotAlertsClient{subscriptionID: subscriptionID, iotDefenderLocation: iotDefenderLocation, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get IoT alert
// If the operation fails it returns the *ErrorResponse error type.
func (client *IotAlertsClient) Get(ctx context.Context, deviceGroupName string, alertID string, options *IotAlertsGetOptions) (IotAlertsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceGroupName, alertID, options)
	if err != nil {
		return IotAlertsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotAlertsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotAlertsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IotAlertsClient) getCreateRequest(ctx context.Context, deviceGroupName string, alertID string, options *IotAlertsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/alerts/{alertId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if alertID == "" {
		return nil, errors.New("parameter alertID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", url.PathEscape(alertID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotAlertsClient) getHandleResponse(resp *http.Response) (IotAlertsGetResponse, error) {
	result := IotAlertsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertModel); err != nil {
		return IotAlertsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IotAlertsClient) getHandleError(resp *http.Response) error {
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

// List - List IoT alerts
// If the operation fails it returns the *ErrorResponse error type.
func (client *IotAlertsClient) List(deviceGroupName string, options *IotAlertsListOptions) *IotAlertsListPager {
	return &IotAlertsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, deviceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IotAlertsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AlertListModel.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IotAlertsClient) listCreateRequest(ctx context.Context, deviceGroupName string, options *IotAlertsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/alerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IotAlertsClient) listHandleResponse(resp *http.Response) (IotAlertsListResponse, error) {
	result := IotAlertsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertListModel); err != nil {
		return IotAlertsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IotAlertsClient) listHandleError(resp *http.Response) error {
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

// Patch - Update an existing alert
// If the operation fails it returns the *ErrorResponse error type.
func (client *IotAlertsClient) Patch(ctx context.Context, deviceGroupName string, alertID string, alertPatchModel AlertPatchPropertiesModel, options *IotAlertsPatchOptions) (IotAlertsPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, deviceGroupName, alertID, alertPatchModel, options)
	if err != nil {
		return IotAlertsPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotAlertsPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotAlertsPatchResponse{}, client.patchHandleError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *IotAlertsClient) patchCreateRequest(ctx context.Context, deviceGroupName string, alertID string, alertPatchModel AlertPatchPropertiesModel, options *IotAlertsPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/alerts/{alertId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if alertID == "" {
		return nil, errors.New("parameter alertID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", url.PathEscape(alertID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, alertPatchModel)
}

// patchHandleResponse handles the Patch response.
func (client *IotAlertsClient) patchHandleResponse(resp *http.Response) (IotAlertsPatchResponse, error) {
	result := IotAlertsPatchResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertModel); err != nil {
		return IotAlertsPatchResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// patchHandleError handles the Patch error response.
func (client *IotAlertsClient) patchHandleError(resp *http.Response) error {
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
