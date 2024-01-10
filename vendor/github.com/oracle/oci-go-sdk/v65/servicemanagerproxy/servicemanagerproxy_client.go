// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ServiceManagerProxyClient a client for ServiceManagerProxy
type ServiceManagerProxyClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewServiceManagerProxyClientWithConfigurationProvider Creates a new default ServiceManagerProxy client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewServiceManagerProxyClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ServiceManagerProxyClient, err error) {
	if enabled := common.CheckForEnabledServices("servicemanagerproxy"); !enabled {
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
	return newServiceManagerProxyClientFromBaseClient(baseClient, provider)
}

// NewServiceManagerProxyClientWithOboToken Creates a new default ServiceManagerProxy client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewServiceManagerProxyClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ServiceManagerProxyClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newServiceManagerProxyClientFromBaseClient(baseClient, configProvider)
}

func newServiceManagerProxyClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ServiceManagerProxyClient, err error) {
	// ServiceManagerProxy service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ServiceManagerProxy"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ServiceManagerProxyClient{BaseClient: baseClient}
	client.BasePath = "20210914"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ServiceManagerProxyClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("servicemanagerproxy", "https://smproxy.{region}.ocs.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ServiceManagerProxyClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ServiceManagerProxyClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetServiceEnvironment Get the detailed information for a specific service environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemanagerproxy/GetServiceEnvironment.go.html to see an example of how to use GetServiceEnvironment API.
func (client ServiceManagerProxyClient) GetServiceEnvironment(ctx context.Context, request GetServiceEnvironmentRequest) (response GetServiceEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceEnvironmentResponse")
	}
	return
}

// getServiceEnvironment implements the OCIOperation interface (enables retrying operations)
func (client ServiceManagerProxyClient) getServiceEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceEnvironments/{serviceEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/smp/20210914/ServiceEnvironment/GetServiceEnvironment"
		err = common.PostProcessServiceError(err, "ServiceManagerProxy", "GetServiceEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceEnvironments List the details of Software as a Service (SaaS) environments provisioned by Service Manager.
// Information includes the service instance endpoints and service definition details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemanagerproxy/ListServiceEnvironments.go.html to see an example of how to use ListServiceEnvironments API.
func (client ServiceManagerProxyClient) ListServiceEnvironments(ctx context.Context, request ListServiceEnvironmentsRequest) (response ListServiceEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceEnvironmentsResponse")
	}
	return
}

// listServiceEnvironments implements the OCIOperation interface (enables retrying operations)
func (client ServiceManagerProxyClient) listServiceEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceEnvironmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/smp/20210914/ServiceEnvironment/ListServiceEnvironments"
		err = common.PostProcessServiceError(err, "ServiceManagerProxy", "ListServiceEnvironments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
