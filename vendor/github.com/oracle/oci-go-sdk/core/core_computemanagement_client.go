// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ComputeManagementClient a client for ComputeManagement
type ComputeManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewComputeManagementClientWithConfigurationProvider Creates a new default ComputeManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComputeManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComputeManagementClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = ComputeManagementClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ComputeManagementClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ComputeManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ComputeManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateInstanceConfiguration Creates an instance configuration
func (client ComputeManagementClient) CreateInstanceConfiguration(ctx context.Context, request CreateInstanceConfigurationRequest) (response CreateInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createInstanceConfiguration, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInstanceConfigurationResponse")
	}
	return
}

// createInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) createInstanceConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instanceConfigurations")
	if err != nil {
		return nil, err
	}

	var response CreateInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInstanceConfiguration Deletes an instance configuration.
func (client ComputeManagementClient) DeleteInstanceConfiguration(ctx context.Context, request DeleteInstanceConfigurationRequest) (response DeleteInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInstanceConfiguration, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInstanceConfigurationResponse")
	}
	return
}

// deleteInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) deleteInstanceConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/instanceConfigurations/{instanceConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response DeleteInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInstanceConfiguration Gets the specified instance configuration
func (client ComputeManagementClient) GetInstanceConfiguration(ctx context.Context, request GetInstanceConfigurationRequest) (response GetInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstanceConfiguration, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstanceConfigurationResponse")
	}
	return
}

// getInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) getInstanceConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceConfigurations/{instanceConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response GetInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LaunchInstanceConfiguration Launch an instance from an instance configuration
func (client ComputeManagementClient) LaunchInstanceConfiguration(ctx context.Context, request LaunchInstanceConfigurationRequest) (response LaunchInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.launchInstanceConfiguration, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(LaunchInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LaunchInstanceConfigurationResponse")
	}
	return
}

// launchInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) launchInstanceConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instanceConfigurations/{instanceConfigurationId}/actions/launch")
	if err != nil {
		return nil, err
	}

	var response LaunchInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstanceConfigurations Lists the available instanceConfigurations in the specific compartment.
func (client ComputeManagementClient) ListInstanceConfigurations(ctx context.Context, request ListInstanceConfigurationsRequest) (response ListInstanceConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstanceConfigurations, policy)
	if err != nil {
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstanceConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstanceConfigurationsResponse")
	}
	return
}

// listInstanceConfigurations implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) listInstanceConfigurations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceConfigurations")
	if err != nil {
		return nil, err
	}

	var response ListInstanceConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
