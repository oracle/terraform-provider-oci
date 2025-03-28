// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.oracle.com/iaas/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// StreamClient a client for Stream
type StreamClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewStreamClientWithConfigurationProvider Creates a new default Stream client with the given configuration provider.
// the configuration provider will be used for the default signer
func NewStreamClientWithConfigurationProvider(configProvider common.ConfigurationProvider, endpoint string) (client StreamClient, err error) {
	if enabled := common.CheckForEnabledServices("streaming"); !enabled {
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
	return newStreamClientFromBaseClient(baseClient, provider, endpoint)
}

// NewStreamClientWithOboToken Creates a new default Stream client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
func NewStreamClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string, endpoint string) (client StreamClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newStreamClientFromBaseClient(baseClient, configProvider, endpoint)
}

func newStreamClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider, endpoint string) (client StreamClient, err error) {
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = StreamClient{BaseClient: baseClient}
	client.BasePath = "20180418"
	client.Host = endpoint
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *StreamClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *StreamClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ConsumerCommit Provides a mechanism to manually commit offsets, if not using commit-on-get consumer semantics.
// This commits offsets assicated with the provided cursor, extends the timeout on each of the affected partitions, and returns an updated cursor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ConsumerCommit.go.html to see an example of how to use ConsumerCommit API.
// A default retry strategy applies to this operation ConsumerCommit()
func (client StreamClient) ConsumerCommit(ctx context.Context, request ConsumerCommitRequest) (response ConsumerCommitResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.consumerCommit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConsumerCommitResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConsumerCommitResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConsumerCommitResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConsumerCommitResponse")
	}
	return
}

// consumerCommit implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) consumerCommit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/commit", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConsumerCommitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Group/ConsumerCommit"
		err = common.PostProcessServiceError(err, "Stream", "ConsumerCommit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConsumerHeartbeat Allows long-running processes to extend the timeout on partitions reserved by a consumer instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ConsumerHeartbeat.go.html to see an example of how to use ConsumerHeartbeat API.
// A default retry strategy applies to this operation ConsumerHeartbeat()
func (client StreamClient) ConsumerHeartbeat(ctx context.Context, request ConsumerHeartbeatRequest) (response ConsumerHeartbeatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.consumerHeartbeat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConsumerHeartbeatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConsumerHeartbeatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConsumerHeartbeatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConsumerHeartbeatResponse")
	}
	return
}

// consumerHeartbeat implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) consumerHeartbeat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/heartbeat", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConsumerHeartbeatResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Group/ConsumerHeartbeat"
		err = common.PostProcessServiceError(err, "Stream", "ConsumerHeartbeat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCursor Creates a cursor. Cursors are used to consume a stream, starting from a specific point in the partition and going forward from there.
// You can create a cursor based on an offset, a time, the trim horizon, or the most recent message in the stream. As the oldest message
// inside the retention period boundary, using the trim horizon effectively lets you consume all messages in the stream. A cursor based
// on the most recent message allows consumption of only messages that are added to the stream after you create the cursor. Cursors expire
// five minutes after you receive them from the service.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/CreateCursor.go.html to see an example of how to use CreateCursor API.
// A default retry strategy applies to this operation CreateCursor()
func (client StreamClient) CreateCursor(ctx context.Context, request CreateCursorRequest) (response CreateCursorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createCursor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCursorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCursorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCursorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCursorResponse")
	}
	return
}

// createCursor implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) createCursor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/cursors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCursorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Cursor/CreateCursor"
		err = common.PostProcessServiceError(err, "Stream", "CreateCursor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGroupCursor Creates a group-cursor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/CreateGroupCursor.go.html to see an example of how to use CreateGroupCursor API.
// A default retry strategy applies to this operation CreateGroupCursor()
func (client StreamClient) CreateGroupCursor(ctx context.Context, request CreateGroupCursorRequest) (response CreateGroupCursorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createGroupCursor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGroupCursorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGroupCursorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGroupCursorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGroupCursorResponse")
	}
	return
}

// createGroupCursor implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) createGroupCursor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/groupCursors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGroupCursorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Cursor/CreateGroupCursor"
		err = common.PostProcessServiceError(err, "Stream", "CreateGroupCursor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGroup Returns the current state of a consumer group.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/GetGroup.go.html to see an example of how to use GetGroup API.
// A default retry strategy applies to this operation GetGroup()
func (client StreamClient) GetGroup(ctx context.Context, request GetGroupRequest) (response GetGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGroupResponse")
	}
	return
}

// getGroup implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) getGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams/{streamId}/groups/{groupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Group/GetGroup"
		err = common.PostProcessServiceError(err, "Stream", "GetGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMessages Returns messages from the specified stream using the specified cursor as the starting point for consumption. By default, the number of messages returned is undefined, but the service returns as many as possible.
// To get messages, you must first obtain a cursor using the CreateCursor operation.
// In the response, retrieve the value of the 'opc-next-cursor' header to pass as a parameter to get the next batch of messages in the stream.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/GetMessages.go.html to see an example of how to use GetMessages API.
// A default retry strategy applies to this operation GetMessages()
func (client StreamClient) GetMessages(ctx context.Context, request GetMessagesRequest) (response GetMessagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMessages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMessagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMessagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMessagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMessagesResponse")
	}
	return
}

// getMessages implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) getMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams/{streamId}/messages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMessagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Message/GetMessages"
		err = common.PostProcessServiceError(err, "Stream", "GetMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutMessages Emits messages to a stream. There's no limit to the number of messages in a request, but the total size of a message or request must be 1 MiB or less.
// The service calculates the partition ID from the message key and stores messages that share a key on the same partition.
// If a message does not contain a key or if the key is null, the service generates a message key for you.
// The partition ID cannot be passed as a parameter.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/PutMessages.go.html to see an example of how to use PutMessages API.
func (client StreamClient) PutMessages(ctx context.Context, request PutMessagesRequest) (response PutMessagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putMessages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutMessagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutMessagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutMessagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutMessagesResponse")
	}
	return
}

// putMessages implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) putMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/messages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutMessagesResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Message/PutMessages"
		err = common.PostProcessServiceError(err, "Stream", "PutMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGroup Forcefully changes the current location of a group as a whole; reseting processing location of all consumers to a particular location in the stream.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/UpdateGroup.go.html to see an example of how to use UpdateGroup API.
// A default retry strategy applies to this operation UpdateGroup()
func (client StreamClient) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGroupResponse")
	}
	return
}

// updateGroup implements the OCIOperation interface (enables retrying operations)
func (client StreamClient) updateGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streams/{streamId}/groups/{groupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Group/UpdateGroup"
		err = common.PostProcessServiceError(err, "Stream", "UpdateGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
