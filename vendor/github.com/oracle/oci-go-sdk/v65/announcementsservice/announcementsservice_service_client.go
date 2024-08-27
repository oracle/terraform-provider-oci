// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ServiceClient a client for Service
type ServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewServiceClientWithConfigurationProvider Creates a new default Service client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("announcementsservice"); !enabled {
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
	return newServiceClientFromBaseClient(baseClient, provider)
}

// NewServiceClientWithOboToken Creates a new default Service client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newServiceClientFromBaseClient(baseClient, configProvider)
}

func newServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ServiceClient, err error) {
	// Service service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Service"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ServiceClient{BaseClient: baseClient}
	client.BasePath = "20180904"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("announcements", "https://announcements.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListServices List all active services
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListServices.go.html to see an example of how to use ListServices API.
func (client ServiceClient) ListServices(ctx context.Context, request ListServicesRequest) (response ListServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServicesResponse")
	}
	return
}

// listServices implements the OCIOperation interface (enables retrying operations)
func (client ServiceClient) listServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/services", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/Service/ListServices"
		err = common.PostProcessServiceError(err, "Service", "ListServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
