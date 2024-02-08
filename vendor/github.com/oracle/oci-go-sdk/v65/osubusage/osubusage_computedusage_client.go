// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Usage Computation
//
// OneSubscription API Common set of Subscription Plan Management (SPM) Usage Computation resources
//

package osubusage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ComputedUsageClient a client for ComputedUsage
type ComputedUsageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewComputedUsageClientWithConfigurationProvider Creates a new default ComputedUsage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComputedUsageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComputedUsageClient, err error) {
	if enabled := common.CheckForEnabledServices("osubusage"); !enabled {
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
	return newComputedUsageClientFromBaseClient(baseClient, provider)
}

// NewComputedUsageClientWithOboToken Creates a new default ComputedUsage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewComputedUsageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ComputedUsageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newComputedUsageClientFromBaseClient(baseClient, configProvider)
}

func newComputedUsageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ComputedUsageClient, err error) {
	// ComputedUsage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ComputedUsage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ComputedUsageClient{BaseClient: baseClient}
	client.BasePath = "oalapp/service/onesubs/proxy/20210501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ComputedUsageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osubusage", "https://csaap-e.oracle.com")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ComputedUsageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ComputedUsageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetComputedUsage This is an API which returns Computed Usage corresponding to the id passed
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubusage/GetComputedUsage.go.html to see an example of how to use GetComputedUsage API.
func (client ComputedUsageClient) GetComputedUsage(ctx context.Context, request GetComputedUsageRequest) (response GetComputedUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getComputedUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetComputedUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetComputedUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetComputedUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetComputedUsageResponse")
	}
	return
}

// getComputedUsage implements the OCIOperation interface (enables retrying operations)
func (client ComputedUsageClient) getComputedUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/computedUsages/{computedUsageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetComputedUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ComputedUsage", "GetComputedUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListComputedUsageAggregateds This is a collection API which returns a list of aggregated computed usage details (there can be multiple Parent Products under a given SubID each of which is represented under Subscription Service Line # in SPM).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubusage/ListComputedUsageAggregateds.go.html to see an example of how to use ListComputedUsageAggregateds API.
func (client ComputedUsageClient) ListComputedUsageAggregateds(ctx context.Context, request ListComputedUsageAggregatedsRequest) (response ListComputedUsageAggregatedsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listComputedUsageAggregateds, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListComputedUsageAggregatedsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListComputedUsageAggregatedsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListComputedUsageAggregatedsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListComputedUsageAggregatedsResponse")
	}
	return
}

// listComputedUsageAggregateds implements the OCIOperation interface (enables retrying operations)
func (client ComputedUsageClient) listComputedUsageAggregateds(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/computedUsages/aggregated", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListComputedUsageAggregatedsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ComputedUsage", "ListComputedUsageAggregateds", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListComputedUsages This is a collection API which returns a list of Computed Usages for given filters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubusage/ListComputedUsages.go.html to see an example of how to use ListComputedUsages API.
func (client ComputedUsageClient) ListComputedUsages(ctx context.Context, request ListComputedUsagesRequest) (response ListComputedUsagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listComputedUsages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListComputedUsagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListComputedUsagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListComputedUsagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListComputedUsagesResponse")
	}
	return
}

// listComputedUsages implements the OCIOperation interface (enables retrying operations)
func (client ComputedUsageClient) listComputedUsages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/computedUsages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListComputedUsagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ComputedUsage", "ListComputedUsages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
