// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// UsagePlansClient a client for UsagePlans
type UsagePlansClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewUsagePlansClientWithConfigurationProvider Creates a new default UsagePlans client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewUsagePlansClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client UsagePlansClient, err error) {
	if enabled := common.CheckForEnabledServices("apigateway"); !enabled {
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
	return newUsagePlansClientFromBaseClient(baseClient, provider)
}

// NewUsagePlansClientWithOboToken Creates a new default UsagePlans client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewUsagePlansClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client UsagePlansClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newUsagePlansClientFromBaseClient(baseClient, configProvider)
}

func newUsagePlansClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client UsagePlansClient, err error) {
	// UsagePlans service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("UsagePlans"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = UsagePlansClient{BaseClient: baseClient}
	client.BasePath = "20190501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *UsagePlansClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apigateway", "https://apigateway.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *UsagePlansClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *UsagePlansClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeUsagePlanCompartment Changes the usage plan compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ChangeUsagePlanCompartment.go.html to see an example of how to use ChangeUsagePlanCompartment API.
func (client UsagePlansClient) ChangeUsagePlanCompartment(ctx context.Context, request ChangeUsagePlanCompartmentRequest) (response ChangeUsagePlanCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeUsagePlanCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeUsagePlanCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeUsagePlanCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeUsagePlanCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeUsagePlanCompartmentResponse")
	}
	return
}

// changeUsagePlanCompartment implements the OCIOperation interface (enables retrying operations)
func (client UsagePlansClient) changeUsagePlanCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/usagePlans/{usagePlanId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeUsagePlanCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/UsagePlan/ChangeUsagePlanCompartment"
		err = common.PostProcessServiceError(err, "UsagePlans", "ChangeUsagePlanCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUsagePlan Creates a new usage plan.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/CreateUsagePlan.go.html to see an example of how to use CreateUsagePlan API.
// A default retry strategy applies to this operation CreateUsagePlan()
func (client UsagePlansClient) CreateUsagePlan(ctx context.Context, request CreateUsagePlanRequest) (response CreateUsagePlanResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUsagePlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUsagePlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUsagePlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUsagePlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUsagePlanResponse")
	}
	return
}

// createUsagePlan implements the OCIOperation interface (enables retrying operations)
func (client UsagePlansClient) createUsagePlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/usagePlans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUsagePlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "UsagePlans", "CreateUsagePlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUsagePlan Deletes the usage plan with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/DeleteUsagePlan.go.html to see an example of how to use DeleteUsagePlan API.
func (client UsagePlansClient) DeleteUsagePlan(ctx context.Context, request DeleteUsagePlanRequest) (response DeleteUsagePlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUsagePlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUsagePlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUsagePlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUsagePlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUsagePlanResponse")
	}
	return
}

// deleteUsagePlan implements the OCIOperation interface (enables retrying operations)
func (client UsagePlansClient) deleteUsagePlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/usagePlans/{usagePlanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUsagePlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/UsagePlan/DeleteUsagePlan"
		err = common.PostProcessServiceError(err, "UsagePlans", "DeleteUsagePlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUsagePlan Gets a usage plan by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/GetUsagePlan.go.html to see an example of how to use GetUsagePlan API.
// A default retry strategy applies to this operation GetUsagePlan()
func (client UsagePlansClient) GetUsagePlan(ctx context.Context, request GetUsagePlanRequest) (response GetUsagePlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUsagePlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUsagePlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUsagePlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUsagePlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUsagePlanResponse")
	}
	return
}

// getUsagePlan implements the OCIOperation interface (enables retrying operations)
func (client UsagePlansClient) getUsagePlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usagePlans/{usagePlanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUsagePlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/UsagePlan/GetUsagePlan"
		err = common.PostProcessServiceError(err, "UsagePlans", "GetUsagePlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUsagePlans Returns a list of usage plans.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListUsagePlans.go.html to see an example of how to use ListUsagePlans API.
// A default retry strategy applies to this operation ListUsagePlans()
func (client UsagePlansClient) ListUsagePlans(ctx context.Context, request ListUsagePlansRequest) (response ListUsagePlansResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsagePlans, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsagePlansResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsagePlansResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsagePlansResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsagePlansResponse")
	}
	return
}

// listUsagePlans implements the OCIOperation interface (enables retrying operations)
func (client UsagePlansClient) listUsagePlans(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usagePlans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsagePlansResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/UsagePlan/ListUsagePlans"
		err = common.PostProcessServiceError(err, "UsagePlans", "ListUsagePlans", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUsagePlan Updates the usage plan with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/UpdateUsagePlan.go.html to see an example of how to use UpdateUsagePlan API.
func (client UsagePlansClient) UpdateUsagePlan(ctx context.Context, request UpdateUsagePlanRequest) (response UpdateUsagePlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUsagePlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUsagePlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUsagePlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUsagePlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUsagePlanResponse")
	}
	return
}

// updateUsagePlan implements the OCIOperation interface (enables retrying operations)
func (client UsagePlansClient) updateUsagePlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/usagePlans/{usagePlanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUsagePlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/UsagePlan/UpdateUsagePlan"
		err = common.PostProcessServiceError(err, "UsagePlans", "UpdateUsagePlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
