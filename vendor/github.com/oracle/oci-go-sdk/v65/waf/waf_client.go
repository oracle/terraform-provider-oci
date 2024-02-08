// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// WafClient a client for Waf
type WafClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewWafClientWithConfigurationProvider Creates a new default Waf client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewWafClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client WafClient, err error) {
	if enabled := common.CheckForEnabledServices("waf"); !enabled {
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
	return newWafClientFromBaseClient(baseClient, provider)
}

// NewWafClientWithOboToken Creates a new default Waf client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewWafClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client WafClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newWafClientFromBaseClient(baseClient, configProvider)
}

func newWafClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client WafClient, err error) {
	// Waf service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Waf"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = WafClient{BaseClient: baseClient}
	client.BasePath = "20210930"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *WafClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("waf", "https://waf.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *WafClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *WafClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeNetworkAddressListCompartment Moves a NetworkAddressList resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ChangeNetworkAddressListCompartment.go.html to see an example of how to use ChangeNetworkAddressListCompartment API.
// A default retry strategy applies to this operation ChangeNetworkAddressListCompartment()
func (client WafClient) ChangeNetworkAddressListCompartment(ctx context.Context, request ChangeNetworkAddressListCompartmentRequest) (response ChangeNetworkAddressListCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeNetworkAddressListCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeNetworkAddressListCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeNetworkAddressListCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeNetworkAddressListCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeNetworkAddressListCompartmentResponse")
	}
	return
}

// changeNetworkAddressListCompartment implements the OCIOperation interface (enables retrying operations)
func (client WafClient) changeNetworkAddressListCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkAddressLists/{networkAddressListId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeNetworkAddressListCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/NetworkAddressList/ChangeNetworkAddressListCompartment"
		err = common.PostProcessServiceError(err, "Waf", "ChangeNetworkAddressListCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeWebAppFirewallCompartment Moves a Web App Firewall resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ChangeWebAppFirewallCompartment.go.html to see an example of how to use ChangeWebAppFirewallCompartment API.
// A default retry strategy applies to this operation ChangeWebAppFirewallCompartment()
func (client WafClient) ChangeWebAppFirewallCompartment(ctx context.Context, request ChangeWebAppFirewallCompartmentRequest) (response ChangeWebAppFirewallCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeWebAppFirewallCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeWebAppFirewallCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeWebAppFirewallCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeWebAppFirewallCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeWebAppFirewallCompartmentResponse")
	}
	return
}

// changeWebAppFirewallCompartment implements the OCIOperation interface (enables retrying operations)
func (client WafClient) changeWebAppFirewallCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppFirewalls/{webAppFirewallId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeWebAppFirewallCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewall/ChangeWebAppFirewallCompartment"
		err = common.PostProcessServiceError(err, "Waf", "ChangeWebAppFirewallCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeWebAppFirewallPolicyCompartment Moves a WebAppFirewallfPolicy resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ChangeWebAppFirewallPolicyCompartment.go.html to see an example of how to use ChangeWebAppFirewallPolicyCompartment API.
// A default retry strategy applies to this operation ChangeWebAppFirewallPolicyCompartment()
func (client WafClient) ChangeWebAppFirewallPolicyCompartment(ctx context.Context, request ChangeWebAppFirewallPolicyCompartmentRequest) (response ChangeWebAppFirewallPolicyCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeWebAppFirewallPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeWebAppFirewallPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeWebAppFirewallPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeWebAppFirewallPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeWebAppFirewallPolicyCompartmentResponse")
	}
	return
}

// changeWebAppFirewallPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client WafClient) changeWebAppFirewallPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppFirewallPolicies/{webAppFirewallPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeWebAppFirewallPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewallPolicy/ChangeWebAppFirewallPolicyCompartment"
		err = common.PostProcessServiceError(err, "Waf", "ChangeWebAppFirewallPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNetworkAddressList Creates a new NetworkAddressList.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/CreateNetworkAddressList.go.html to see an example of how to use CreateNetworkAddressList API.
// A default retry strategy applies to this operation CreateNetworkAddressList()
func (client WafClient) CreateNetworkAddressList(ctx context.Context, request CreateNetworkAddressListRequest) (response CreateNetworkAddressListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNetworkAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNetworkAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNetworkAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNetworkAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNetworkAddressListResponse")
	}
	return
}

// createNetworkAddressList implements the OCIOperation interface (enables retrying operations)
func (client WafClient) createNetworkAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkAddressLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNetworkAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/NetworkAddressList/CreateNetworkAddressList"
		err = common.PostProcessServiceError(err, "Waf", "CreateNetworkAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &networkaddresslist{})
	return response, err
}

// CreateWebAppFirewall Creates a new WebAppFirewall.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/CreateWebAppFirewall.go.html to see an example of how to use CreateWebAppFirewall API.
// A default retry strategy applies to this operation CreateWebAppFirewall()
func (client WafClient) CreateWebAppFirewall(ctx context.Context, request CreateWebAppFirewallRequest) (response CreateWebAppFirewallResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createWebAppFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateWebAppFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateWebAppFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateWebAppFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateWebAppFirewallResponse")
	}
	return
}

// createWebAppFirewall implements the OCIOperation interface (enables retrying operations)
func (client WafClient) createWebAppFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppFirewalls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateWebAppFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewall/CreateWebAppFirewall"
		err = common.PostProcessServiceError(err, "Waf", "CreateWebAppFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &webappfirewall{})
	return response, err
}

// CreateWebAppFirewallPolicy Creates a new WebAppFirewallPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/CreateWebAppFirewallPolicy.go.html to see an example of how to use CreateWebAppFirewallPolicy API.
// A default retry strategy applies to this operation CreateWebAppFirewallPolicy()
func (client WafClient) CreateWebAppFirewallPolicy(ctx context.Context, request CreateWebAppFirewallPolicyRequest) (response CreateWebAppFirewallPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createWebAppFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateWebAppFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateWebAppFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateWebAppFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateWebAppFirewallPolicyResponse")
	}
	return
}

// createWebAppFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client WafClient) createWebAppFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/webAppFirewallPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateWebAppFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewallPolicy/CreateWebAppFirewallPolicy"
		err = common.PostProcessServiceError(err, "Waf", "CreateWebAppFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNetworkAddressList Deletes a NetworkAddressList resource identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/DeleteNetworkAddressList.go.html to see an example of how to use DeleteNetworkAddressList API.
// A default retry strategy applies to this operation DeleteNetworkAddressList()
func (client WafClient) DeleteNetworkAddressList(ctx context.Context, request DeleteNetworkAddressListRequest) (response DeleteNetworkAddressListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNetworkAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNetworkAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNetworkAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNetworkAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNetworkAddressListResponse")
	}
	return
}

// deleteNetworkAddressList implements the OCIOperation interface (enables retrying operations)
func (client WafClient) deleteNetworkAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkAddressLists/{networkAddressListId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNetworkAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/NetworkAddressList/DeleteNetworkAddressList"
		err = common.PostProcessServiceError(err, "Waf", "DeleteNetworkAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWebAppFirewall Deletes a WebAppFirewall resource identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/DeleteWebAppFirewall.go.html to see an example of how to use DeleteWebAppFirewall API.
// A default retry strategy applies to this operation DeleteWebAppFirewall()
func (client WafClient) DeleteWebAppFirewall(ctx context.Context, request DeleteWebAppFirewallRequest) (response DeleteWebAppFirewallResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWebAppFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWebAppFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWebAppFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWebAppFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWebAppFirewallResponse")
	}
	return
}

// deleteWebAppFirewall implements the OCIOperation interface (enables retrying operations)
func (client WafClient) deleteWebAppFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/webAppFirewalls/{webAppFirewallId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWebAppFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewall/DeleteWebAppFirewall"
		err = common.PostProcessServiceError(err, "Waf", "DeleteWebAppFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWebAppFirewallPolicy Deletes a WebAppFirewallPolicy resource identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/DeleteWebAppFirewallPolicy.go.html to see an example of how to use DeleteWebAppFirewallPolicy API.
// A default retry strategy applies to this operation DeleteWebAppFirewallPolicy()
func (client WafClient) DeleteWebAppFirewallPolicy(ctx context.Context, request DeleteWebAppFirewallPolicyRequest) (response DeleteWebAppFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWebAppFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWebAppFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWebAppFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWebAppFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWebAppFirewallPolicyResponse")
	}
	return
}

// deleteWebAppFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client WafClient) deleteWebAppFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/webAppFirewallPolicies/{webAppFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWebAppFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewallPolicy/DeleteWebAppFirewallPolicy"
		err = common.PostProcessServiceError(err, "Waf", "DeleteWebAppFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNetworkAddressList Gets a NetworkAddressList by OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/GetNetworkAddressList.go.html to see an example of how to use GetNetworkAddressList API.
// A default retry strategy applies to this operation GetNetworkAddressList()
func (client WafClient) GetNetworkAddressList(ctx context.Context, request GetNetworkAddressListRequest) (response GetNetworkAddressListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNetworkAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkAddressListResponse")
	}
	return
}

// getNetworkAddressList implements the OCIOperation interface (enables retrying operations)
func (client WafClient) getNetworkAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkAddressLists/{networkAddressListId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/NetworkAddressList/GetNetworkAddressList"
		err = common.PostProcessServiceError(err, "Waf", "GetNetworkAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &networkaddresslist{})
	return response, err
}

// GetWebAppFirewall Gets a WebAppFirewall by OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/GetWebAppFirewall.go.html to see an example of how to use GetWebAppFirewall API.
// A default retry strategy applies to this operation GetWebAppFirewall()
func (client WafClient) GetWebAppFirewall(ctx context.Context, request GetWebAppFirewallRequest) (response GetWebAppFirewallResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWebAppFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWebAppFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWebAppFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWebAppFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWebAppFirewallResponse")
	}
	return
}

// getWebAppFirewall implements the OCIOperation interface (enables retrying operations)
func (client WafClient) getWebAppFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppFirewalls/{webAppFirewallId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWebAppFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewall/GetWebAppFirewall"
		err = common.PostProcessServiceError(err, "Waf", "GetWebAppFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &webappfirewall{})
	return response, err
}

// GetWebAppFirewallPolicy Gets a WebAppFirewallPolicy with the given OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/GetWebAppFirewallPolicy.go.html to see an example of how to use GetWebAppFirewallPolicy API.
// A default retry strategy applies to this operation GetWebAppFirewallPolicy()
func (client WafClient) GetWebAppFirewallPolicy(ctx context.Context, request GetWebAppFirewallPolicyRequest) (response GetWebAppFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWebAppFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWebAppFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWebAppFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWebAppFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWebAppFirewallPolicyResponse")
	}
	return
}

// getWebAppFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client WafClient) getWebAppFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppFirewallPolicies/{webAppFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWebAppFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewallPolicy/GetWebAppFirewallPolicy"
		err = common.PostProcessServiceError(err, "Waf", "GetWebAppFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the WorkRequest with the given OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client WafClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client WafClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Waf", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkAddressLists Gets a list of all NetworkAddressLists in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListNetworkAddressLists.go.html to see an example of how to use ListNetworkAddressLists API.
// A default retry strategy applies to this operation ListNetworkAddressLists()
func (client WafClient) ListNetworkAddressLists(ctx context.Context, request ListNetworkAddressListsRequest) (response ListNetworkAddressListsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkAddressLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkAddressListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkAddressListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkAddressListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkAddressListsResponse")
	}
	return
}

// listNetworkAddressLists implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listNetworkAddressLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkAddressLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkAddressListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/NetworkAddressList/ListNetworkAddressLists"
		err = common.PostProcessServiceError(err, "Waf", "ListNetworkAddressLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProtectionCapabilities Lists of protection capabilities filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListProtectionCapabilities.go.html to see an example of how to use ListProtectionCapabilities API.
// A default retry strategy applies to this operation ListProtectionCapabilities()
func (client WafClient) ListProtectionCapabilities(ctx context.Context, request ListProtectionCapabilitiesRequest) (response ListProtectionCapabilitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProtectionCapabilities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProtectionCapabilitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProtectionCapabilitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProtectionCapabilitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProtectionCapabilitiesResponse")
	}
	return
}

// listProtectionCapabilities implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listProtectionCapabilities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/protectionCapabilities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProtectionCapabilitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/ProtectionCapability/ListProtectionCapabilities"
		err = common.PostProcessServiceError(err, "Waf", "ListProtectionCapabilities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProtectionCapabilityGroupTags Lists of available group tags filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListProtectionCapabilityGroupTags.go.html to see an example of how to use ListProtectionCapabilityGroupTags API.
// A default retry strategy applies to this operation ListProtectionCapabilityGroupTags()
func (client WafClient) ListProtectionCapabilityGroupTags(ctx context.Context, request ListProtectionCapabilityGroupTagsRequest) (response ListProtectionCapabilityGroupTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProtectionCapabilityGroupTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProtectionCapabilityGroupTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProtectionCapabilityGroupTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProtectionCapabilityGroupTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProtectionCapabilityGroupTagsResponse")
	}
	return
}

// listProtectionCapabilityGroupTags implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listProtectionCapabilityGroupTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/protectionCapabilities/groupTags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProtectionCapabilityGroupTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/ProtectionCapability/ListProtectionCapabilityGroupTags"
		err = common.PostProcessServiceError(err, "Waf", "ListProtectionCapabilityGroupTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWebAppFirewallPolicies Gets a list of all WebAppFirewallPolicies in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWebAppFirewallPolicies.go.html to see an example of how to use ListWebAppFirewallPolicies API.
// A default retry strategy applies to this operation ListWebAppFirewallPolicies()
func (client WafClient) ListWebAppFirewallPolicies(ctx context.Context, request ListWebAppFirewallPoliciesRequest) (response ListWebAppFirewallPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWebAppFirewallPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWebAppFirewallPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWebAppFirewallPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWebAppFirewallPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWebAppFirewallPoliciesResponse")
	}
	return
}

// listWebAppFirewallPolicies implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listWebAppFirewallPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppFirewallPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWebAppFirewallPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewallPolicy/ListWebAppFirewallPolicies"
		err = common.PostProcessServiceError(err, "Waf", "ListWebAppFirewallPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWebAppFirewalls Gets a list of all WebAppFirewalls in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWebAppFirewalls.go.html to see an example of how to use ListWebAppFirewalls API.
// A default retry strategy applies to this operation ListWebAppFirewalls()
func (client WafClient) ListWebAppFirewalls(ctx context.Context, request ListWebAppFirewallsRequest) (response ListWebAppFirewallsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWebAppFirewalls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWebAppFirewallsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWebAppFirewallsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWebAppFirewallsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWebAppFirewallsResponse")
	}
	return
}

// listWebAppFirewalls implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listWebAppFirewalls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/webAppFirewalls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWebAppFirewallsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewall/ListWebAppFirewalls"
		err = common.PostProcessServiceError(err, "Waf", "ListWebAppFirewalls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client WafClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Waf", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client WafClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Waf", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the WorkRequests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client WafClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client WafClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Waf", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNetworkAddressList Update the NetworkAddressList identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/UpdateNetworkAddressList.go.html to see an example of how to use UpdateNetworkAddressList API.
// A default retry strategy applies to this operation UpdateNetworkAddressList()
func (client WafClient) UpdateNetworkAddressList(ctx context.Context, request UpdateNetworkAddressListRequest) (response UpdateNetworkAddressListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNetworkAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNetworkAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNetworkAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNetworkAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNetworkAddressListResponse")
	}
	return
}

// updateNetworkAddressList implements the OCIOperation interface (enables retrying operations)
func (client WafClient) updateNetworkAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkAddressLists/{networkAddressListId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNetworkAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/NetworkAddressList/UpdateNetworkAddressList"
		err = common.PostProcessServiceError(err, "Waf", "UpdateNetworkAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWebAppFirewall Updates the WebAppFirewall identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/UpdateWebAppFirewall.go.html to see an example of how to use UpdateWebAppFirewall API.
// A default retry strategy applies to this operation UpdateWebAppFirewall()
func (client WafClient) UpdateWebAppFirewall(ctx context.Context, request UpdateWebAppFirewallRequest) (response UpdateWebAppFirewallResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWebAppFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWebAppFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWebAppFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWebAppFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWebAppFirewallResponse")
	}
	return
}

// updateWebAppFirewall implements the OCIOperation interface (enables retrying operations)
func (client WafClient) updateWebAppFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/webAppFirewalls/{webAppFirewallId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWebAppFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewall/UpdateWebAppFirewall"
		err = common.PostProcessServiceError(err, "Waf", "UpdateWebAppFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWebAppFirewallPolicy Update the WebAppFirewallPolicy identified by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/UpdateWebAppFirewallPolicy.go.html to see an example of how to use UpdateWebAppFirewallPolicy API.
// A default retry strategy applies to this operation UpdateWebAppFirewallPolicy()
func (client WafClient) UpdateWebAppFirewallPolicy(ctx context.Context, request UpdateWebAppFirewallPolicyRequest) (response UpdateWebAppFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWebAppFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWebAppFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWebAppFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWebAppFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWebAppFirewallPolicyResponse")
	}
	return
}

// updateWebAppFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client WafClient) updateWebAppFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/webAppFirewallPolicies/{webAppFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWebAppFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/waf/20210930/WebAppFirewallPolicy/UpdateWebAppFirewallPolicy"
		err = common.PostProcessServiceError(err, "Waf", "UpdateWebAppFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
