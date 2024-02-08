// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration (WAA) API
//
// API for the Web Application Acceleration service.
// Use this API to manage regional Web App Acceleration policies such as Caching and Compression
// for accelerating HTTP services.
//

package waa

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// WaaClient a client for Waa
type WaaClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewWaaClientWithConfigurationProvider Creates a new default Waa client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewWaaClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client WaaClient, err error) {
	if enabled := common.CheckForEnabledServices("waa"); !enabled {
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
	return newWaaClientFromBaseClient(baseClient, provider)
}

// NewWaaClientWithOboToken Creates a new default Waa client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewWaaClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client WaaClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newWaaClientFromBaseClient(baseClient, configProvider)
}

func newWaaClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client WaaClient, err error) {
	// Waa service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Waa"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = WaaClient{BaseClient: baseClient}
	client.BasePath = "20211230"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *WaaClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("waa", "https://waa.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *WaaClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *WaaClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeWebAppAccelerationCompartment Moves a Web App Acceleration resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/ChangeWebAppAccelerationCompartment.go.html to see an example of how to use ChangeWebAppAccelerationCompartment API.
func (client WaaClient) ChangeWebAppAccelerationCompartment(ctx context.Context, request ChangeWebAppAccelerationCompartmentRequest) (response ChangeWebAppAccelerationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeWebAppAccelerationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeWebAppAccelerationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeWebAppAccelerationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeWebAppAccelerationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeWebAppAccelerationCompartmentResponse")
	}
	return
}

// changeWebAppAccelerationCompartment implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) changeWebAppAccelerationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppAccelerations/{webAppAccelerationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeWebAppAccelerationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/ChangeWebAppAccelerationCompartment"
		err = common.PostProcessServiceError(err, "Waa", "ChangeWebAppAccelerationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeWebAppAccelerationPolicyCompartment Moves a WebAppAccelerationfPolicy resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/ChangeWebAppAccelerationPolicyCompartment.go.html to see an example of how to use ChangeWebAppAccelerationPolicyCompartment API.
func (client WaaClient) ChangeWebAppAccelerationPolicyCompartment(ctx context.Context, request ChangeWebAppAccelerationPolicyCompartmentRequest) (response ChangeWebAppAccelerationPolicyCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeWebAppAccelerationPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeWebAppAccelerationPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeWebAppAccelerationPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeWebAppAccelerationPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeWebAppAccelerationPolicyCompartmentResponse")
	}
	return
}

// changeWebAppAccelerationPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) changeWebAppAccelerationPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppAccelerationPolicies/{webAppAccelerationPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeWebAppAccelerationPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAccelerationPolicy/ChangeWebAppAccelerationPolicyCompartment"
		err = common.PostProcessServiceError(err, "Waa", "ChangeWebAppAccelerationPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateWebAppAcceleration Creates a new WebAppAcceleration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/CreateWebAppAcceleration.go.html to see an example of how to use CreateWebAppAcceleration API.
func (client WaaClient) CreateWebAppAcceleration(ctx context.Context, request CreateWebAppAccelerationRequest) (response CreateWebAppAccelerationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createWebAppAcceleration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateWebAppAccelerationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateWebAppAccelerationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateWebAppAccelerationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateWebAppAccelerationResponse")
	}
	return
}

// createWebAppAcceleration implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) createWebAppAcceleration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppAccelerations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateWebAppAccelerationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/CreateWebAppAcceleration"
		err = common.PostProcessServiceError(err, "Waa", "CreateWebAppAcceleration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &webappacceleration{})
	return response, err
}

// CreateWebAppAccelerationPolicy Creates a new WebAppAccelerationPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/CreateWebAppAccelerationPolicy.go.html to see an example of how to use CreateWebAppAccelerationPolicy API.
func (client WaaClient) CreateWebAppAccelerationPolicy(ctx context.Context, request CreateWebAppAccelerationPolicyRequest) (response CreateWebAppAccelerationPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createWebAppAccelerationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateWebAppAccelerationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateWebAppAccelerationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateWebAppAccelerationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateWebAppAccelerationPolicyResponse")
	}
	return
}

// createWebAppAccelerationPolicy implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) createWebAppAccelerationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppAccelerationPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateWebAppAccelerationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAccelerationPolicy/CreateWebAppAccelerationPolicy"
		err = common.PostProcessServiceError(err, "Waa", "CreateWebAppAccelerationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWebAppAcceleration Deletes a WebAppAcceleration resource identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/DeleteWebAppAcceleration.go.html to see an example of how to use DeleteWebAppAcceleration API.
func (client WaaClient) DeleteWebAppAcceleration(ctx context.Context, request DeleteWebAppAccelerationRequest) (response DeleteWebAppAccelerationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWebAppAcceleration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWebAppAccelerationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWebAppAccelerationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWebAppAccelerationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWebAppAccelerationResponse")
	}
	return
}

// deleteWebAppAcceleration implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) deleteWebAppAcceleration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/webAppAccelerations/{webAppAccelerationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWebAppAccelerationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/DeleteWebAppAcceleration"
		err = common.PostProcessServiceError(err, "Waa", "DeleteWebAppAcceleration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWebAppAccelerationPolicy Deletes a WebAppAccelerationPolicy resource identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/DeleteWebAppAccelerationPolicy.go.html to see an example of how to use DeleteWebAppAccelerationPolicy API.
func (client WaaClient) DeleteWebAppAccelerationPolicy(ctx context.Context, request DeleteWebAppAccelerationPolicyRequest) (response DeleteWebAppAccelerationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWebAppAccelerationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWebAppAccelerationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWebAppAccelerationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWebAppAccelerationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWebAppAccelerationPolicyResponse")
	}
	return
}

// deleteWebAppAccelerationPolicy implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) deleteWebAppAccelerationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/webAppAccelerationPolicies/{webAppAccelerationPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWebAppAccelerationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAccelerationPolicy/DeleteWebAppAccelerationPolicy"
		err = common.PostProcessServiceError(err, "Waa", "DeleteWebAppAccelerationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWebAppAcceleration Gets a WebAppAcceleration by OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/GetWebAppAcceleration.go.html to see an example of how to use GetWebAppAcceleration API.
func (client WaaClient) GetWebAppAcceleration(ctx context.Context, request GetWebAppAccelerationRequest) (response GetWebAppAccelerationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWebAppAcceleration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWebAppAccelerationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWebAppAccelerationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWebAppAccelerationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWebAppAccelerationResponse")
	}
	return
}

// getWebAppAcceleration implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) getWebAppAcceleration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppAccelerations/{webAppAccelerationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWebAppAccelerationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/GetWebAppAcceleration"
		err = common.PostProcessServiceError(err, "Waa", "GetWebAppAcceleration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &webappacceleration{})
	return response, err
}

// GetWebAppAccelerationPolicy Gets a WebAppAccelerationPolicy with the given OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/GetWebAppAccelerationPolicy.go.html to see an example of how to use GetWebAppAccelerationPolicy API.
func (client WaaClient) GetWebAppAccelerationPolicy(ctx context.Context, request GetWebAppAccelerationPolicyRequest) (response GetWebAppAccelerationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWebAppAccelerationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWebAppAccelerationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWebAppAccelerationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWebAppAccelerationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWebAppAccelerationPolicyResponse")
	}
	return
}

// getWebAppAccelerationPolicy implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) getWebAppAccelerationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppAccelerationPolicies/{webAppAccelerationPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWebAppAccelerationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAccelerationPolicy/GetWebAppAccelerationPolicy"
		err = common.PostProcessServiceError(err, "Waa", "GetWebAppAccelerationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWebAppAccelerationPolicies Gets a list of all WebAppAccelerationPolicies in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/ListWebAppAccelerationPolicies.go.html to see an example of how to use ListWebAppAccelerationPolicies API.
func (client WaaClient) ListWebAppAccelerationPolicies(ctx context.Context, request ListWebAppAccelerationPoliciesRequest) (response ListWebAppAccelerationPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWebAppAccelerationPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWebAppAccelerationPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWebAppAccelerationPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWebAppAccelerationPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWebAppAccelerationPoliciesResponse")
	}
	return
}

// listWebAppAccelerationPolicies implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) listWebAppAccelerationPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppAccelerationPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWebAppAccelerationPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAccelerationPolicy/ListWebAppAccelerationPolicies"
		err = common.PostProcessServiceError(err, "Waa", "ListWebAppAccelerationPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWebAppAccelerations Gets a list of all WebAppAccelerations in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/ListWebAppAccelerations.go.html to see an example of how to use ListWebAppAccelerations API.
func (client WaaClient) ListWebAppAccelerations(ctx context.Context, request ListWebAppAccelerationsRequest) (response ListWebAppAccelerationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWebAppAccelerations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWebAppAccelerationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWebAppAccelerationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWebAppAccelerationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWebAppAccelerationsResponse")
	}
	return
}

// listWebAppAccelerations implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) listWebAppAccelerations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppAccelerations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWebAppAccelerationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/ListWebAppAccelerations"
		err = common.PostProcessServiceError(err, "Waa", "ListWebAppAccelerations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PurgeWebAppAccelerationCache Clears resources from the cache of the WebAppAcceleration. Each new request for a purged resource will be
// forwarded to the origin server to fetch a new version of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/PurgeWebAppAccelerationCache.go.html to see an example of how to use PurgeWebAppAccelerationCache API.
func (client WaaClient) PurgeWebAppAccelerationCache(ctx context.Context, request PurgeWebAppAccelerationCacheRequest) (response PurgeWebAppAccelerationCacheResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.purgeWebAppAccelerationCache, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PurgeWebAppAccelerationCacheResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PurgeWebAppAccelerationCacheResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PurgeWebAppAccelerationCacheResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PurgeWebAppAccelerationCacheResponse")
	}
	return
}

// purgeWebAppAccelerationCache implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) purgeWebAppAccelerationCache(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppAccelerations/{webAppAccelerationId}/actions/purgeCache", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PurgeWebAppAccelerationCacheResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/PurgeWebAppAccelerationCache"
		err = common.PostProcessServiceError(err, "Waa", "PurgeWebAppAccelerationCache", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWebAppAcceleration Updates the WebAppAcceleration identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/UpdateWebAppAcceleration.go.html to see an example of how to use UpdateWebAppAcceleration API.
func (client WaaClient) UpdateWebAppAcceleration(ctx context.Context, request UpdateWebAppAccelerationRequest) (response UpdateWebAppAccelerationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWebAppAcceleration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWebAppAccelerationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWebAppAccelerationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWebAppAccelerationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWebAppAccelerationResponse")
	}
	return
}

// updateWebAppAcceleration implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) updateWebAppAcceleration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/webAppAccelerations/{webAppAccelerationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWebAppAccelerationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAcceleration/UpdateWebAppAcceleration"
		err = common.PostProcessServiceError(err, "Waa", "UpdateWebAppAcceleration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWebAppAccelerationPolicy Update the WebAppAccelerationPolicy identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/UpdateWebAppAccelerationPolicy.go.html to see an example of how to use UpdateWebAppAccelerationPolicy API.
func (client WaaClient) UpdateWebAppAccelerationPolicy(ctx context.Context, request UpdateWebAppAccelerationPolicyRequest) (response UpdateWebAppAccelerationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWebAppAccelerationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWebAppAccelerationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWebAppAccelerationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWebAppAccelerationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWebAppAccelerationPolicyResponse")
	}
	return
}

// updateWebAppAccelerationPolicy implements the OCIOperation interface (enables retrying operations)
func (client WaaClient) updateWebAppAccelerationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/webAppAccelerationPolicies/{webAppAccelerationPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWebAppAccelerationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waa/20211230/WebAppAccelerationPolicy/UpdateWebAppAccelerationPolicy"
		err = common.PostProcessServiceError(err, "Waa", "UpdateWebAppAccelerationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
