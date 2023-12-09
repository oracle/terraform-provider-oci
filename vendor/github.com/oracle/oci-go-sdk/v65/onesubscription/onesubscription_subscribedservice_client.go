// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SubscribedServiceClient a client for SubscribedService
type SubscribedServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSubscribedServiceClientWithConfigurationProvider Creates a new default SubscribedService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSubscribedServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SubscribedServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("onesubscription"); !enabled {
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
	return newSubscribedServiceClientFromBaseClient(baseClient, provider)
}

// NewSubscribedServiceClientWithOboToken Creates a new default SubscribedService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSubscribedServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SubscribedServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSubscribedServiceClientFromBaseClient(baseClient, configProvider)
}

func newSubscribedServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SubscribedServiceClient, err error) {
	// SubscribedService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("SubscribedService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SubscribedServiceClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SubscribedServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("onesubscription", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SubscribedServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SubscribedServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetSubscribedService This API returns the subscribed service details corresponding to the id provided
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/GetSubscribedService.go.html to see an example of how to use GetSubscribedService API.
func (client SubscribedServiceClient) GetSubscribedService(ctx context.Context, request GetSubscribedServiceRequest) (response GetSubscribedServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSubscribedService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSubscribedServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSubscribedServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSubscribedServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSubscribedServiceResponse")
	}
	return
}

// getSubscribedService implements the OCIOperation interface (enables retrying operations)
func (client SubscribedServiceClient) getSubscribedService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscribedServices/{subscribedServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSubscribedServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "SubscribedService", "GetSubscribedService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSubscribedServices This list API returns all subscribed services for given Subscription ID
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListSubscribedServices.go.html to see an example of how to use ListSubscribedServices API.
func (client SubscribedServiceClient) ListSubscribedServices(ctx context.Context, request ListSubscribedServicesRequest) (response ListSubscribedServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSubscribedServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSubscribedServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSubscribedServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSubscribedServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSubscribedServicesResponse")
	}
	return
}

// listSubscribedServices implements the OCIOperation interface (enables retrying operations)
func (client SubscribedServiceClient) listSubscribedServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscribedServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSubscribedServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "SubscribedService", "ListSubscribedServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
