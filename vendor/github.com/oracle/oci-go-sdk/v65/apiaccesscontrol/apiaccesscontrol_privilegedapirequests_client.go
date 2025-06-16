// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// PrivilegedApiRequestsClient a client for PrivilegedApiRequests
type PrivilegedApiRequestsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPrivilegedApiRequestsClientWithConfigurationProvider Creates a new default PrivilegedApiRequests client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPrivilegedApiRequestsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PrivilegedApiRequestsClient, err error) {
	if enabled := common.CheckForEnabledServices("apiaccesscontrol"); !enabled {
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
	return newPrivilegedApiRequestsClientFromBaseClient(baseClient, provider)
}

// NewPrivilegedApiRequestsClientWithOboToken Creates a new default PrivilegedApiRequests client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewPrivilegedApiRequestsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PrivilegedApiRequestsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPrivilegedApiRequestsClientFromBaseClient(baseClient, configProvider)
}

func newPrivilegedApiRequestsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PrivilegedApiRequestsClient, err error) {
	// PrivilegedApiRequests service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("PrivilegedApiRequests"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PrivilegedApiRequestsClient{BaseClient: baseClient}
	client.BasePath = "20241130"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PrivilegedApiRequestsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apiaccesscontrol", "https://pactl.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PrivilegedApiRequestsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PrivilegedApiRequestsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApprovePrivilegedApiRequest Approves privilegedApi request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ApprovePrivilegedApiRequest.go.html to see an example of how to use ApprovePrivilegedApiRequest API.
// A default retry strategy applies to this operation ApprovePrivilegedApiRequest()
func (client PrivilegedApiRequestsClient) ApprovePrivilegedApiRequest(ctx context.Context, request ApprovePrivilegedApiRequestRequest) (response ApprovePrivilegedApiRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.approvePrivilegedApiRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApprovePrivilegedApiRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApprovePrivilegedApiRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApprovePrivilegedApiRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApprovePrivilegedApiRequestResponse")
	}
	return
}

// approvePrivilegedApiRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) approvePrivilegedApiRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiRequests/{privilegedApiRequestId}/actions/approve", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApprovePrivilegedApiRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "ApprovePrivilegedApiRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ClosePrivilegedApiRequest Closes privilegedApi request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ClosePrivilegedApiRequest.go.html to see an example of how to use ClosePrivilegedApiRequest API.
// A default retry strategy applies to this operation ClosePrivilegedApiRequest()
func (client PrivilegedApiRequestsClient) ClosePrivilegedApiRequest(ctx context.Context, request ClosePrivilegedApiRequestRequest) (response ClosePrivilegedApiRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.closePrivilegedApiRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ClosePrivilegedApiRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ClosePrivilegedApiRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ClosePrivilegedApiRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ClosePrivilegedApiRequestResponse")
	}
	return
}

// closePrivilegedApiRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) closePrivilegedApiRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiRequests/{privilegedApiRequestId}/actions/close", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ClosePrivilegedApiRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "ClosePrivilegedApiRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivilegedApiRequest Creates a PrivilegedApiRequest.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/CreatePrivilegedApiRequest.go.html to see an example of how to use CreatePrivilegedApiRequest API.
// A default retry strategy applies to this operation CreatePrivilegedApiRequest()
func (client PrivilegedApiRequestsClient) CreatePrivilegedApiRequest(ctx context.Context, request CreatePrivilegedApiRequestRequest) (response CreatePrivilegedApiRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPrivilegedApiRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivilegedApiRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivilegedApiRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivilegedApiRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivilegedApiRequestResponse")
	}
	return
}

// createPrivilegedApiRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) createPrivilegedApiRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePrivilegedApiRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "CreatePrivilegedApiRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivilegedApiRequest Gets information about a PrivilegedApiRequest.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/GetPrivilegedApiRequest.go.html to see an example of how to use GetPrivilegedApiRequest API.
// A default retry strategy applies to this operation GetPrivilegedApiRequest()
func (client PrivilegedApiRequestsClient) GetPrivilegedApiRequest(ctx context.Context, request GetPrivilegedApiRequestRequest) (response GetPrivilegedApiRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivilegedApiRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivilegedApiRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivilegedApiRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivilegedApiRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivilegedApiRequestResponse")
	}
	return
}

// getPrivilegedApiRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) getPrivilegedApiRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privilegedApiRequests/{privilegedApiRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivilegedApiRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "GetPrivilegedApiRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivilegedApiRequests Lists all privilegedApi requests in the compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListPrivilegedApiRequests.go.html to see an example of how to use ListPrivilegedApiRequests API.
// A default retry strategy applies to this operation ListPrivilegedApiRequests()
func (client PrivilegedApiRequestsClient) ListPrivilegedApiRequests(ctx context.Context, request ListPrivilegedApiRequestsRequest) (response ListPrivilegedApiRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivilegedApiRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivilegedApiRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivilegedApiRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivilegedApiRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivilegedApiRequestsResponse")
	}
	return
}

// listPrivilegedApiRequests implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) listPrivilegedApiRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privilegedApiRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPrivilegedApiRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "ListPrivilegedApiRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RejectPrivilegedApiRequest Rejects privilegedApi request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/RejectPrivilegedApiRequest.go.html to see an example of how to use RejectPrivilegedApiRequest API.
// A default retry strategy applies to this operation RejectPrivilegedApiRequest()
func (client PrivilegedApiRequestsClient) RejectPrivilegedApiRequest(ctx context.Context, request RejectPrivilegedApiRequestRequest) (response RejectPrivilegedApiRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rejectPrivilegedApiRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RejectPrivilegedApiRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RejectPrivilegedApiRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RejectPrivilegedApiRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RejectPrivilegedApiRequestResponse")
	}
	return
}

// rejectPrivilegedApiRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) rejectPrivilegedApiRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiRequests/{privilegedApiRequestId}/actions/reject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RejectPrivilegedApiRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "RejectPrivilegedApiRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RevokePrivilegedApiRequest Revokes an already approved privilegedApi request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/RevokePrivilegedApiRequest.go.html to see an example of how to use RevokePrivilegedApiRequest API.
// A default retry strategy applies to this operation RevokePrivilegedApiRequest()
func (client PrivilegedApiRequestsClient) RevokePrivilegedApiRequest(ctx context.Context, request RevokePrivilegedApiRequestRequest) (response RevokePrivilegedApiRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.revokePrivilegedApiRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RevokePrivilegedApiRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RevokePrivilegedApiRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RevokePrivilegedApiRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RevokePrivilegedApiRequestResponse")
	}
	return
}

// revokePrivilegedApiRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiRequestsClient) revokePrivilegedApiRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiRequests/{privilegedApiRequestId}/actions/revoke", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RevokePrivilegedApiRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivilegedApiRequests", "RevokePrivilegedApiRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
