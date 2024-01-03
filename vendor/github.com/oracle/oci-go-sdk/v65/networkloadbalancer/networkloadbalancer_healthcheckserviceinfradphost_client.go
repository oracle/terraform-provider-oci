// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// HealthCheckServiceInfraDpHostClient a client for HealthCheckServiceInfraDpHost
type HealthCheckServiceInfraDpHostClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewHealthCheckServiceInfraDpHostClientWithConfigurationProvider Creates a new default HealthCheckServiceInfraDpHost client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewHealthCheckServiceInfraDpHostClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client HealthCheckServiceInfraDpHostClient, err error) {
	if enabled := common.CheckForEnabledServices("networkloadbalancer"); !enabled {
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
	return newHealthCheckServiceInfraDpHostClientFromBaseClient(baseClient, provider)
}

// NewHealthCheckServiceInfraDpHostClientWithOboToken Creates a new default HealthCheckServiceInfraDpHost client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewHealthCheckServiceInfraDpHostClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client HealthCheckServiceInfraDpHostClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newHealthCheckServiceInfraDpHostClientFromBaseClient(baseClient, configProvider)
}

func newHealthCheckServiceInfraDpHostClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client HealthCheckServiceInfraDpHostClient, err error) {
	// HealthCheckServiceInfraDpHost service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("HealthCheckServiceInfraDpHost"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = HealthCheckServiceInfraDpHostClient{BaseClient: baseClient}
	client.BasePath = "20200501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *HealthCheckServiceInfraDpHostClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("networkloadbalancer", "https://network-load-balancer-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *HealthCheckServiceInfraDpHostClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *HealthCheckServiceInfraDpHostClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DeleteHealthCheckServiceInfraDpHost Delete hcs dp host configuration by identifier.
// A default retry strategy applies to this operation DeleteHealthCheckServiceInfraDpHost()
func (client HealthCheckServiceInfraDpHostClient) DeleteHealthCheckServiceInfraDpHost(ctx context.Context, request DeleteHealthCheckServiceInfraDpHostRequest) (response DeleteHealthCheckServiceInfraDpHostResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteHealthCheckServiceInfraDpHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteHealthCheckServiceInfraDpHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteHealthCheckServiceInfraDpHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteHealthCheckServiceInfraDpHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteHealthCheckServiceInfraDpHostResponse")
	}
	return
}

// deleteHealthCheckServiceInfraDpHost implements the OCIOperation interface (enables retrying operations)
func (client HealthCheckServiceInfraDpHostClient) deleteHealthCheckServiceInfraDpHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/healthCheckServiceInfraDpHost/{availabilityDomain}/{dpHostId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteHealthCheckServiceInfraDpHostResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/DpHost/DeleteHealthCheckServiceInfraDpHost"
		err = common.PostProcessServiceError(err, "HealthCheckServiceInfraDpHost", "DeleteHealthCheckServiceInfraDpHost", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetHealthCheckServiceInfraDpHost Retrieves hcs dp host configuration information by identifier.
// A default retry strategy applies to this operation GetHealthCheckServiceInfraDpHost()
func (client HealthCheckServiceInfraDpHostClient) GetHealthCheckServiceInfraDpHost(ctx context.Context, request GetHealthCheckServiceInfraDpHostRequest) (response GetHealthCheckServiceInfraDpHostResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHealthCheckServiceInfraDpHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHealthCheckServiceInfraDpHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHealthCheckServiceInfraDpHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHealthCheckServiceInfraDpHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHealthCheckServiceInfraDpHostResponse")
	}
	return
}

// getHealthCheckServiceInfraDpHost implements the OCIOperation interface (enables retrying operations)
func (client HealthCheckServiceInfraDpHostClient) getHealthCheckServiceInfraDpHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/healthCheckServiceInfraDpHost/{availabilityDomain}/{dpHostId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHealthCheckServiceInfraDpHostResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/DpHost/GetHealthCheckServiceInfraDpHost"
		err = common.PostProcessServiceError(err, "HealthCheckServiceInfraDpHost", "GetHealthCheckServiceInfraDpHost", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PostHeartBeat Post HCS dp host heart beat
// A default retry strategy applies to this operation PostHeartBeat()
func (client HealthCheckServiceInfraDpHostClient) PostHeartBeat(ctx context.Context, request PostHeartBeatRequest) (response PostHeartBeatResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.postHeartBeat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PostHeartBeatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PostHeartBeatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PostHeartBeatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PostHeartBeatResponse")
	}
	return
}

// postHeartBeat implements the OCIOperation interface (enables retrying operations)
func (client HealthCheckServiceInfraDpHostClient) postHeartBeat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/healthCheckServiceInfraDpHost/heartBeat", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PostHeartBeatResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/DpHost/PostHeartBeat"
		err = common.PostProcessServiceError(err, "HealthCheckServiceInfraDpHost", "PostHeartBeat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateHealthCheckServiceInfraDpHost Update HCS dp host
// A default retry strategy applies to this operation UpdateHealthCheckServiceInfraDpHost()
func (client HealthCheckServiceInfraDpHostClient) UpdateHealthCheckServiceInfraDpHost(ctx context.Context, request UpdateHealthCheckServiceInfraDpHostRequest) (response UpdateHealthCheckServiceInfraDpHostResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateHealthCheckServiceInfraDpHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHealthCheckServiceInfraDpHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHealthCheckServiceInfraDpHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHealthCheckServiceInfraDpHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHealthCheckServiceInfraDpHostResponse")
	}
	return
}

// updateHealthCheckServiceInfraDpHost implements the OCIOperation interface (enables retrying operations)
func (client HealthCheckServiceInfraDpHostClient) updateHealthCheckServiceInfraDpHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/healthCheckServiceInfraDpHost/{availabilityDomain}/{dpHostId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateHealthCheckServiceInfraDpHostResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/DpHost/UpdateHealthCheckServiceInfraDpHost"
		err = common.PostProcessServiceError(err, "HealthCheckServiceInfraDpHost", "UpdateHealthCheckServiceInfraDpHost", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
