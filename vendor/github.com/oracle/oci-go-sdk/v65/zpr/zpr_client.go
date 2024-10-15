// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Zero Trust Packet Routing Control Plane API
//
// Use the Zero Trust Packet Routing Control Plane API to manage ZPR configuration and policy. See the Zero Trust Packet Routing (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/home.htm) documentation for more information.
//

package zpr

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ZprClient a client for Zpr
type ZprClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewZprClientWithConfigurationProvider Creates a new default Zpr client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewZprClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ZprClient, err error) {
	if enabled := common.CheckForEnabledServices("zpr"); !enabled {
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
	return newZprClientFromBaseClient(baseClient, provider)
}

// NewZprClientWithOboToken Creates a new default Zpr client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewZprClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ZprClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newZprClientFromBaseClient(baseClient, configProvider)
}

func newZprClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ZprClient, err error) {
	// Zpr service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Zpr"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ZprClient{BaseClient: baseClient}
	client.BasePath = "20240301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ZprClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("zpr", "https://zpr.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ZprClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ZprClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateConfiguration Initiates the process to onboard ZPR
// in a root compartment (the root compartment is the tenancy). It creates an object of ZPR configuration as part of onboarding.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/CreateConfiguration.go.html to see an example of how to use CreateConfiguration API.
// A default retry strategy applies to this operation CreateConfiguration()
func (client ZprClient) CreateConfiguration(ctx context.Context, request CreateConfigurationRequest) (response CreateConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConfigurationResponse")
	}
	return
}

// createConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) createConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/Configuration/CreateConfiguration"
		err = common.PostProcessServiceError(err, "Zpr", "CreateConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateZprPolicy Creates a ZprPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/CreateZprPolicy.go.html to see an example of how to use CreateZprPolicy API.
// A default retry strategy applies to this operation CreateZprPolicy()
func (client ZprClient) CreateZprPolicy(ctx context.Context, request CreateZprPolicyRequest) (response CreateZprPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createZprPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateZprPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateZprPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateZprPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateZprPolicyResponse")
	}
	return
}

// createZprPolicy implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) createZprPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/zprPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateZprPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/ZprPolicy/CreateZprPolicy"
		err = common.PostProcessServiceError(err, "Zpr", "CreateZprPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteZprPolicy Deletes a ZprPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/DeleteZprPolicy.go.html to see an example of how to use DeleteZprPolicy API.
// A default retry strategy applies to this operation DeleteZprPolicy()
func (client ZprClient) DeleteZprPolicy(ctx context.Context, request DeleteZprPolicyRequest) (response DeleteZprPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteZprPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteZprPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteZprPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteZprPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteZprPolicyResponse")
	}
	return
}

// deleteZprPolicy implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) deleteZprPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/zprPolicies/{zprPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteZprPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/ZprPolicy/DeleteZprPolicy"
		err = common.PostProcessServiceError(err, "Zpr", "DeleteZprPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfiguration Retrieves the ZPR configuration details for the root compartment (the root compartment is the tenancy).
// Returns ZPR configuration for root compartment (the root compartment is the tenancy).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/GetConfiguration.go.html to see an example of how to use GetConfiguration API.
// A default retry strategy applies to this operation GetConfiguration()
func (client ZprClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationResponse")
	}
	return
}

// getConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) getConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/Configuration/GetConfiguration"
		err = common.PostProcessServiceError(err, "Zpr", "GetConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetZprConfigurationWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/GetZprConfigurationWorkRequest.go.html to see an example of how to use GetZprConfigurationWorkRequest API.
// A default retry strategy applies to this operation GetZprConfigurationWorkRequest()
func (client ZprClient) GetZprConfigurationWorkRequest(ctx context.Context, request GetZprConfigurationWorkRequestRequest) (response GetZprConfigurationWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getZprConfigurationWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetZprConfigurationWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetZprConfigurationWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetZprConfigurationWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetZprConfigurationWorkRequestResponse")
	}
	return
}

// getZprConfigurationWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) getZprConfigurationWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprConfigurationWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetZprConfigurationWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequest/GetZprConfigurationWorkRequest"
		err = common.PostProcessServiceError(err, "Zpr", "GetZprConfigurationWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetZprPolicy Gets information about a ZprPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/GetZprPolicy.go.html to see an example of how to use GetZprPolicy API.
// A default retry strategy applies to this operation GetZprPolicy()
func (client ZprClient) GetZprPolicy(ctx context.Context, request GetZprPolicyRequest) (response GetZprPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getZprPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetZprPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetZprPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetZprPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetZprPolicyResponse")
	}
	return
}

// getZprPolicy implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) getZprPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprPolicies/{zprPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetZprPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/ZprPolicy/GetZprPolicy"
		err = common.PostProcessServiceError(err, "Zpr", "GetZprPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetZprPolicyWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/GetZprPolicyWorkRequest.go.html to see an example of how to use GetZprPolicyWorkRequest API.
// A default retry strategy applies to this operation GetZprPolicyWorkRequest()
func (client ZprClient) GetZprPolicyWorkRequest(ctx context.Context, request GetZprPolicyWorkRequestRequest) (response GetZprPolicyWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getZprPolicyWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetZprPolicyWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetZprPolicyWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetZprPolicyWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetZprPolicyWorkRequestResponse")
	}
	return
}

// getZprPolicyWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) getZprPolicyWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprPolicyWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetZprPolicyWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequest/GetZprPolicyWorkRequest"
		err = common.PostProcessServiceError(err, "Zpr", "GetZprPolicyWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprConfigurationWorkRequestErrors Lists the errors for a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprConfigurationWorkRequestErrors.go.html to see an example of how to use ListZprConfigurationWorkRequestErrors API.
// A default retry strategy applies to this operation ListZprConfigurationWorkRequestErrors()
func (client ZprClient) ListZprConfigurationWorkRequestErrors(ctx context.Context, request ListZprConfigurationWorkRequestErrorsRequest) (response ListZprConfigurationWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprConfigurationWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprConfigurationWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprConfigurationWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprConfigurationWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprConfigurationWorkRequestErrorsResponse")
	}
	return
}

// listZprConfigurationWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprConfigurationWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprConfigurationWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprConfigurationWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequestError/ListZprConfigurationWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprConfigurationWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprConfigurationWorkRequestLogs Lists the logs for a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprConfigurationWorkRequestLogs.go.html to see an example of how to use ListZprConfigurationWorkRequestLogs API.
// A default retry strategy applies to this operation ListZprConfigurationWorkRequestLogs()
func (client ZprClient) ListZprConfigurationWorkRequestLogs(ctx context.Context, request ListZprConfigurationWorkRequestLogsRequest) (response ListZprConfigurationWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprConfigurationWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprConfigurationWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprConfigurationWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprConfigurationWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprConfigurationWorkRequestLogsResponse")
	}
	return
}

// listZprConfigurationWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprConfigurationWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprConfigurationWorkRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprConfigurationWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequestLogEntry/ListZprConfigurationWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprConfigurationWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprConfigurationWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprConfigurationWorkRequests.go.html to see an example of how to use ListZprConfigurationWorkRequests API.
// A default retry strategy applies to this operation ListZprConfigurationWorkRequests()
func (client ZprClient) ListZprConfigurationWorkRequests(ctx context.Context, request ListZprConfigurationWorkRequestsRequest) (response ListZprConfigurationWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprConfigurationWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprConfigurationWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprConfigurationWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprConfigurationWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprConfigurationWorkRequestsResponse")
	}
	return
}

// listZprConfigurationWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprConfigurationWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprConfigurationWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprConfigurationWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequest/ListZprConfigurationWorkRequests"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprConfigurationWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprPolicies Gets a list of ZprPolicies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicies.go.html to see an example of how to use ListZprPolicies API.
// A default retry strategy applies to this operation ListZprPolicies()
func (client ZprClient) ListZprPolicies(ctx context.Context, request ListZprPoliciesRequest) (response ListZprPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprPoliciesResponse")
	}
	return
}

// listZprPolicies implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/ZprPolicyCollection/ListZprPolicies"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprPolicyWorkRequestErrors Lists the errors for a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicyWorkRequestErrors.go.html to see an example of how to use ListZprPolicyWorkRequestErrors API.
// A default retry strategy applies to this operation ListZprPolicyWorkRequestErrors()
func (client ZprClient) ListZprPolicyWorkRequestErrors(ctx context.Context, request ListZprPolicyWorkRequestErrorsRequest) (response ListZprPolicyWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprPolicyWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprPolicyWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprPolicyWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprPolicyWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprPolicyWorkRequestErrorsResponse")
	}
	return
}

// listZprPolicyWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprPolicyWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprPolicyWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprPolicyWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequestError/ListZprPolicyWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprPolicyWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprPolicyWorkRequestLogs Lists the logs for a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicyWorkRequestLogs.go.html to see an example of how to use ListZprPolicyWorkRequestLogs API.
// A default retry strategy applies to this operation ListZprPolicyWorkRequestLogs()
func (client ZprClient) ListZprPolicyWorkRequestLogs(ctx context.Context, request ListZprPolicyWorkRequestLogsRequest) (response ListZprPolicyWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprPolicyWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprPolicyWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprPolicyWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprPolicyWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprPolicyWorkRequestLogsResponse")
	}
	return
}

// listZprPolicyWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprPolicyWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprPolicyWorkRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprPolicyWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequestLogEntry/ListZprPolicyWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprPolicyWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZprPolicyWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicyWorkRequests.go.html to see an example of how to use ListZprPolicyWorkRequests API.
// A default retry strategy applies to this operation ListZprPolicyWorkRequests()
func (client ZprClient) ListZprPolicyWorkRequests(ctx context.Context, request ListZprPolicyWorkRequestsRequest) (response ListZprPolicyWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZprPolicyWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZprPolicyWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZprPolicyWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZprPolicyWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZprPolicyWorkRequestsResponse")
	}
	return
}

// listZprPolicyWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) listZprPolicyWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zprPolicyWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZprPolicyWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/WorkRequest/ListZprPolicyWorkRequests"
		err = common.PostProcessServiceError(err, "Zpr", "ListZprPolicyWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateZprPolicy Updates a specific ZprPolicy. If updating on statements, the entire list of policy statements is required, which will replace the existing policy statements associated with the policy ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/UpdateZprPolicy.go.html to see an example of how to use UpdateZprPolicy API.
// A default retry strategy applies to this operation UpdateZprPolicy()
func (client ZprClient) UpdateZprPolicy(ctx context.Context, request UpdateZprPolicyRequest) (response UpdateZprPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateZprPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateZprPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateZprPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateZprPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateZprPolicyResponse")
	}
	return
}

// updateZprPolicy implements the OCIOperation interface (enables retrying operations)
func (client ZprClient) updateZprPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/zprPolicies/{zprPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateZprPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/zero-trust-packet-routing/20240301/ZprPolicy/UpdateZprPolicy"
		err = common.PostProcessServiceError(err, "Zpr", "UpdateZprPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
