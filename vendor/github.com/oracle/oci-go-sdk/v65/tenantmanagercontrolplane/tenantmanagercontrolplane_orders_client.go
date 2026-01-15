// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OrdersClient a client for Orders
type OrdersClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOrdersClientWithConfigurationProvider Creates a new default Orders client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOrdersClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OrdersClient, err error) {
	if enabled := common.CheckForEnabledServices("tenantmanagercontrolplane"); !enabled {
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
	return newOrdersClientFromBaseClient(baseClient, provider)
}

// NewOrdersClientWithOboToken Creates a new default Orders client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOrdersClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OrdersClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOrdersClientFromBaseClient(baseClient, configProvider)
}

func newOrdersClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OrdersClient, err error) {
	// Orders service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Orders"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OrdersClient{BaseClient: baseClient}
	client.BasePath = "20230401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OrdersClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OrdersClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OrdersClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateOrder Triggers an order activation workflow on behalf of the tenant, given by compartment ID in the body.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ActivateOrder.go.html to see an example of how to use ActivateOrder API.
func (client OrdersClient) ActivateOrder(ctx context.Context, request ActivateOrderRequest) (response ActivateOrderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.activateOrder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateOrderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateOrderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateOrderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateOrderResponse")
	}
	return
}

// activateOrder implements the OCIOperation interface (enables retrying operations)
func (client OrdersClient) activateOrder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/orders/{activationToken}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateOrderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/Order/ActivateOrder"
		err = common.PostProcessServiceError(err, "Orders", "ActivateOrder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOrder Returns the order details given by the order ID in the JWT.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetOrder.go.html to see an example of how to use GetOrder API.
func (client OrdersClient) GetOrder(ctx context.Context, request GetOrderRequest) (response GetOrderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOrder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOrderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOrderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOrderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOrderResponse")
	}
	return
}

// getOrder implements the OCIOperation interface (enables retrying operations)
func (client OrdersClient) getOrder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/orders/{activationToken}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOrderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/Order/GetOrder"
		err = common.PostProcessServiceError(err, "Orders", "GetOrder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
