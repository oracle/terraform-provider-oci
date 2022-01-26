// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//AccessRequestsClient a client for AccessRequests
type AccessRequestsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAccessRequestsClientWithConfigurationProvider Creates a new default AccessRequests client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAccessRequestsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AccessRequestsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newAccessRequestsClientFromBaseClient(baseClient, provider)
}

// NewAccessRequestsClientWithOboToken Creates a new default AccessRequests client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAccessRequestsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AccessRequestsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAccessRequestsClientFromBaseClient(baseClient, configProvider)
}

func newAccessRequestsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AccessRequestsClient, err error) {
	// AccessRequests service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AccessRequestsClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AccessRequestsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("operatoraccesscontrol", "https://operator-access-control.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AccessRequestsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *AccessRequestsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApproveAccessRequest Approves an access request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ApproveAccessRequest.go.html to see an example of how to use ApproveAccessRequest API.
func (client AccessRequestsClient) ApproveAccessRequest(ctx context.Context, request ApproveAccessRequestRequest) (response ApproveAccessRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.approveAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApproveAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApproveAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApproveAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApproveAccessRequestResponse")
	}
	return
}

// approveAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) approveAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/accessRequests/{accessRequestId}/action/approve", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApproveAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAccessRequest Gets details of an access request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/GetAccessRequest.go.html to see an example of how to use GetAccessRequest API.
func (client AccessRequestsClient) GetAccessRequest(ctx context.Context, request GetAccessRequestRequest) (response GetAccessRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAccessRequestResponse")
	}
	return
}

// getAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) getAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/accessRequests/{accessRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAccessRequestHistories Returns a history of all status associated with the accessRequestId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListAccessRequestHistories.go.html to see an example of how to use ListAccessRequestHistories API.
func (client AccessRequestsClient) ListAccessRequestHistories(ctx context.Context, request ListAccessRequestHistoriesRequest) (response ListAccessRequestHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAccessRequestHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAccessRequestHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAccessRequestHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAccessRequestHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAccessRequestHistoriesResponse")
	}
	return
}

// listAccessRequestHistories implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) listAccessRequestHistories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/accessRequests/{accessRequestId}/history", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAccessRequestHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAccessRequests Lists all access requests in the compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListAccessRequests.go.html to see an example of how to use ListAccessRequests API.
func (client AccessRequestsClient) ListAccessRequests(ctx context.Context, request ListAccessRequestsRequest) (response ListAccessRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAccessRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAccessRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAccessRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAccessRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAccessRequestsResponse")
	}
	return
}

// listAccessRequests implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) listAccessRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/accessRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAccessRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RejectAccessRequest Rejects an access request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/RejectAccessRequest.go.html to see an example of how to use RejectAccessRequest API.
func (client AccessRequestsClient) RejectAccessRequest(ctx context.Context, request RejectAccessRequestRequest) (response RejectAccessRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rejectAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RejectAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RejectAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RejectAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RejectAccessRequestResponse")
	}
	return
}

// rejectAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) rejectAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/accessRequests/{accessRequestId}/action/reject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RejectAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReviewAccessRequest Reviews the access request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ReviewAccessRequest.go.html to see an example of how to use ReviewAccessRequest API.
func (client AccessRequestsClient) ReviewAccessRequest(ctx context.Context, request ReviewAccessRequestRequest) (response ReviewAccessRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.reviewAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReviewAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReviewAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReviewAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReviewAccessRequestResponse")
	}
	return
}

// reviewAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) reviewAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/accessRequests/{accessRequestId}/action/review", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReviewAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RevokeAccessRequest Revokes an already approved access request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/RevokeAccessRequest.go.html to see an example of how to use RevokeAccessRequest API.
func (client AccessRequestsClient) RevokeAccessRequest(ctx context.Context, request RevokeAccessRequestRequest) (response RevokeAccessRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.revokeAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RevokeAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RevokeAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RevokeAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RevokeAccessRequestResponse")
	}
	return
}

// revokeAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client AccessRequestsClient) revokeAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/accessRequests/{accessRequestId}/action/revoke", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RevokeAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
