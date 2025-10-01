// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MonitoredRegionClient a client for MonitoredRegion
type MonitoredRegionClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMonitoredRegionClientWithConfigurationProvider Creates a new default MonitoredRegion client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMonitoredRegionClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MonitoredRegionClient, err error) {
	if enabled := common.CheckForEnabledServices("resourceanalytics"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newMonitoredRegionClientFromBaseClient(baseClient, provider)
}

// NewMonitoredRegionClientWithOboToken Creates a new default MonitoredRegion client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMonitoredRegionClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MonitoredRegionClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMonitoredRegionClientFromBaseClient(baseClient, configProvider)
}

func newMonitoredRegionClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MonitoredRegionClient, err error) {
	// MonitoredRegion service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MonitoredRegion"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MonitoredRegionClient{BaseClient: baseClient}
	client.BasePath = "20241031"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MonitoredRegionClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("resourceanalytics", "https://resource-analytics.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MonitoredRegionClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *MonitoredRegionClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateMonitoredRegion Creates a MonitoredRegion.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/CreateMonitoredRegion.go.html to see an example of how to use CreateMonitoredRegion API.
// A default retry strategy applies to this operation CreateMonitoredRegion()
func (client MonitoredRegionClient) CreateMonitoredRegion(ctx context.Context, request CreateMonitoredRegionRequest) (response CreateMonitoredRegionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createMonitoredRegion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMonitoredRegionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMonitoredRegionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMonitoredRegionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMonitoredRegionResponse")
	}
	return
}

// createMonitoredRegion implements the OCIOperation interface (enables retrying operations)
func (client MonitoredRegionClient) createMonitoredRegion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredRegions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMonitoredRegionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/MonitoredRegion/CreateMonitoredRegion"
		err = common.PostProcessServiceError(err, "MonitoredRegion", "CreateMonitoredRegion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMonitoredRegion Deletes a MonitoredRegion.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/DeleteMonitoredRegion.go.html to see an example of how to use DeleteMonitoredRegion API.
// A default retry strategy applies to this operation DeleteMonitoredRegion()
func (client MonitoredRegionClient) DeleteMonitoredRegion(ctx context.Context, request DeleteMonitoredRegionRequest) (response DeleteMonitoredRegionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMonitoredRegion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMonitoredRegionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMonitoredRegionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMonitoredRegionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMonitoredRegionResponse")
	}
	return
}

// deleteMonitoredRegion implements the OCIOperation interface (enables retrying operations)
func (client MonitoredRegionClient) deleteMonitoredRegion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/monitoredRegions/{monitoredRegionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMonitoredRegionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/MonitoredRegion/DeleteMonitoredRegion"
		err = common.PostProcessServiceError(err, "MonitoredRegion", "DeleteMonitoredRegion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMonitoredRegion Gets information about a MonitoredRegion.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/GetMonitoredRegion.go.html to see an example of how to use GetMonitoredRegion API.
// A default retry strategy applies to this operation GetMonitoredRegion()
func (client MonitoredRegionClient) GetMonitoredRegion(ctx context.Context, request GetMonitoredRegionRequest) (response GetMonitoredRegionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitoredRegion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitoredRegionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitoredRegionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitoredRegionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitoredRegionResponse")
	}
	return
}

// getMonitoredRegion implements the OCIOperation interface (enables retrying operations)
func (client MonitoredRegionClient) getMonitoredRegion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredRegions/{monitoredRegionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMonitoredRegionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/MonitoredRegion/GetMonitoredRegion"
		err = common.PostProcessServiceError(err, "MonitoredRegion", "GetMonitoredRegion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMonitoredRegions Gets a list of MonitoredRegions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListMonitoredRegions.go.html to see an example of how to use ListMonitoredRegions API.
// A default retry strategy applies to this operation ListMonitoredRegions()
func (client MonitoredRegionClient) ListMonitoredRegions(ctx context.Context, request ListMonitoredRegionsRequest) (response ListMonitoredRegionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMonitoredRegions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMonitoredRegionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMonitoredRegionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMonitoredRegionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMonitoredRegionsResponse")
	}
	return
}

// listMonitoredRegions implements the OCIOperation interface (enables retrying operations)
func (client MonitoredRegionClient) listMonitoredRegions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredRegions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMonitoredRegionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/MonitoredRegionCollection/ListMonitoredRegions"
		err = common.PostProcessServiceError(err, "MonitoredRegion", "ListMonitoredRegions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
