//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DeviceGroupsListPager provides operations for iterating over paged responses.
type DeviceGroupsListPager struct {
	client    *DeviceGroupsClient
	current   DeviceGroupsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DeviceGroupsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DeviceGroupsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DeviceGroupsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceGroupList.NextLink == nil || len(*p.current.DeviceGroupList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DeviceGroupsListResponse page.
func (p *DeviceGroupsListPager) PageResponse() DeviceGroupsListResponse {
	return p.current
}

// DevicesListPager provides operations for iterating over paged responses.
type DevicesListPager struct {
	client    *DevicesClient
	current   DevicesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DevicesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DevicesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DevicesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceList.NextLink == nil || len(*p.current.DeviceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DevicesListResponse page.
func (p *DevicesListPager) PageResponse() DevicesListResponse {
	return p.current
}

// IotAlertsListPager provides operations for iterating over paged responses.
type IotAlertsListPager struct {
	client    *IotAlertsClient
	current   IotAlertsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotAlertsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotAlertsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotAlertsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AlertListModel.NextLink == nil || len(*p.current.AlertListModel.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotAlertsListResponse page.
func (p *IotAlertsListPager) PageResponse() IotAlertsListResponse {
	return p.current
}

// IotDeviceVulnerabilityListPager provides operations for iterating over paged responses.
type IotDeviceVulnerabilityListPager struct {
	client    *IotDeviceVulnerabilityClient
	current   IotDeviceVulnerabilityListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotDeviceVulnerabilityListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotDeviceVulnerabilityListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotDeviceVulnerabilityListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeviceVulnerabilityListModel.NextLink == nil || len(*p.current.DeviceVulnerabilityListModel.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotDeviceVulnerabilityListResponse page.
func (p *IotDeviceVulnerabilityListPager) PageResponse() IotDeviceVulnerabilityListResponse {
	return p.current
}

// IotRecommendationsListPager provides operations for iterating over paged responses.
type IotRecommendationsListPager struct {
	client    *IotRecommendationsClient
	current   IotRecommendationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotRecommendationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotRecommendationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotRecommendationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecommendationListModel.NextLink == nil || len(*p.current.RecommendationListModel.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotRecommendationsListResponse page.
func (p *IotRecommendationsListPager) PageResponse() IotRecommendationsListResponse {
	return p.current
}

// LocationsListPager provides operations for iterating over paged responses.
type LocationsListPager struct {
	client    *LocationsClient
	current   LocationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LocationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LocationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LocationList.NextLink == nil || len(*p.current.LocationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LocationsListResponse page.
func (p *LocationsListPager) PageResponse() LocationsListResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}
