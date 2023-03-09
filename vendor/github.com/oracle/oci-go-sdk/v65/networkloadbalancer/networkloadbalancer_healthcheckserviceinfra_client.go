// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

//HealthCheckServiceInfraClient a client for HealthCheckServiceInfra
type HealthCheckServiceInfraClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewHealthCheckServiceInfraClientWithConfigurationProvider Creates a new default HealthCheckServiceInfra client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewHealthCheckServiceInfraClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client HealthCheckServiceInfraClient, err error) {
	if enabled := common.CheckForEnabledServices("networkloadbalancer"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newHealthCheckServiceInfraClientFromBaseClient(baseClient, provider)
}

// NewHealthCheckServiceInfraClientWithOboToken Creates a new default HealthCheckServiceInfra client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewHealthCheckServiceInfraClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client HealthCheckServiceInfraClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newHealthCheckServiceInfraClientFromBaseClient(baseClient, configProvider)
}

func newHealthCheckServiceInfraClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client HealthCheckServiceInfraClient, err error) {
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = HealthCheckServiceInfraClient{BaseClient: baseClient}
	client.BasePath = "20200501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *HealthCheckServiceInfraClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("networkloadbalancer", "https://network-load-balancer-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *HealthCheckServiceInfraClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *HealthCheckServiceInfraClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// RegisterHealthCheckServiceInfraDpHost Create a HCS dp host
// A default retry strategy applies to this operation RegisterHealthCheckServiceInfraDpHost()
func (client HealthCheckServiceInfraClient) RegisterHealthCheckServiceInfraDpHost(ctx context.Context, request RegisterHealthCheckServiceInfraDpHostRequest) (response RegisterHealthCheckServiceInfraDpHostResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.registerHealthCheckServiceInfraDpHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RegisterHealthCheckServiceInfraDpHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RegisterHealthCheckServiceInfraDpHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RegisterHealthCheckServiceInfraDpHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RegisterHealthCheckServiceInfraDpHostResponse")
	}
	return
}

// registerHealthCheckServiceInfraDpHost implements the OCIOperation interface (enables retrying operations)
func (client HealthCheckServiceInfraClient) registerHealthCheckServiceInfraDpHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/healthCheckServiceInfraDpHost", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RegisterHealthCheckServiceInfraDpHostResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/DpHost/RegisterHealthCheckServiceInfraDpHost"
		err = common.PostProcessServiceError(err, "HealthCheckServiceInfra", "RegisterHealthCheckServiceInfraDpHost", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
