// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// A description of the UsageApi API.
//

package usageapi

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//UsageapiClient a client for Usageapi
type UsageapiClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewUsageapiClientWithConfigurationProvider Creates a new default Usageapi client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewUsageapiClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client UsageapiClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newUsageapiClientFromBaseClient(baseClient, configProvider)
}

// NewUsageapiClientWithOboToken Creates a new default Usageapi client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewUsageapiClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client UsageapiClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newUsageapiClientFromBaseClient(baseClient, configProvider)
}

func newUsageapiClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client UsageapiClient, err error) {
	client = UsageapiClient{BaseClient: baseClient}
	client.BasePath = "20200107"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *UsageapiClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("usageapi", "https://usageapi.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *UsageapiClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *UsageapiClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// RequestSummarizedConfigurations Returns the list of config for UI dropdown list
func (client UsageapiClient) RequestSummarizedConfigurations(ctx context.Context, request RequestSummarizedConfigurationsRequest) (response RequestSummarizedConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedConfigurationsResponse")
	}
	return
}

// requestSummarizedConfigurations implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestSummarizedConfigurations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedConfigurationsResponse
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

// RequestSummarizedUsages Returns the usage for the given account
func (client UsageapiClient) RequestSummarizedUsages(ctx context.Context, request RequestSummarizedUsagesRequest) (response RequestSummarizedUsagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedUsages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedUsagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedUsagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedUsagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedUsagesResponse")
	}
	return
}

// requestSummarizedUsages implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestSummarizedUsages(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/usage")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedUsagesResponse
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
