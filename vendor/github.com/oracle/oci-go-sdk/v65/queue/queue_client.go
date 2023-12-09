// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// QueueClient a client for Queue
type QueueClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewQueueClientWithConfigurationProvider Creates a new default Queue client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewQueueClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client QueueClient, err error) {
	if enabled := common.CheckForEnabledServices("queue"); !enabled {
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
	return newQueueClientFromBaseClient(baseClient, provider)
}

// NewQueueClientWithOboToken Creates a new default Queue client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewQueueClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client QueueClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newQueueClientFromBaseClient(baseClient, configProvider)
}

func newQueueClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client QueueClient, err error) {
	// Queue service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Queue"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = QueueClient{BaseClient: baseClient}
	client.BasePath = "20210201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *QueueClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("queue", "https://messaging.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *QueueClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *QueueClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DeleteMessage Deletes the message represented by the receipt from the queue.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to delete messages.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/DeleteMessage.go.html to see an example of how to use DeleteMessage API.
// A default retry strategy applies to this operation DeleteMessage()
func (client QueueClient) DeleteMessage(ctx context.Context, request DeleteMessageRequest) (response DeleteMessageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMessage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMessageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMessageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMessageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMessageResponse")
	}
	return
}

// deleteMessage implements the OCIOperation interface (enables retrying operations)
func (client QueueClient) deleteMessage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/queues/{queueId}/messages/{messageReceipt}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMessageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/GetMessage/DeleteMessage"
		err = common.PostProcessServiceError(err, "Queue", "DeleteMessage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMessages Deletes multiple messages from the queue.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to delete messages.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/DeleteMessages.go.html to see an example of how to use DeleteMessages API.
// A default retry strategy applies to this operation DeleteMessages()
func (client QueueClient) DeleteMessages(ctx context.Context, request DeleteMessagesRequest) (response DeleteMessagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMessages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMessagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMessagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMessagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMessagesResponse")
	}
	return
}

// deleteMessages implements the OCIOperation interface (enables retrying operations)
func (client QueueClient) deleteMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/messages/actions/deleteMessages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMessagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/GetMessage/DeleteMessages"
		err = common.PostProcessServiceError(err, "Queue", "DeleteMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMessages Consumes messages from the queue.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to consume messages.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
// GetMessages accepts optional channelFilter query parameter that can filter source channels of the messages.
// When channelFilter is present, service will return available messages from the channel which ID exactly matched the filter.
// When filter is not specified, messages will be returned from a random non-empty channel within a queue.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/GetMessages.go.html to see an example of how to use GetMessages API.
func (client QueueClient) GetMessages(ctx context.Context, request GetMessagesRequest) (response GetMessagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client QueueClient) getMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queues/{queueId}/messages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMessagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/GetMessage/GetMessages"
		err = common.PostProcessServiceError(err, "Queue", "GetMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStats Gets the statistics for the queue and its dead letter queue.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to get a queue's statistics.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/GetStats.go.html to see an example of how to use GetStats API.
// A default retry strategy applies to this operation GetStats()
func (client QueueClient) GetStats(ctx context.Context, request GetStatsRequest) (response GetStatsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStatsResponse")
	}
	return
}

// getStats implements the OCIOperation interface (enables retrying operations)
func (client QueueClient) getStats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queues/{queueId}/stats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStatsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/QueueStats/GetStats"
		err = common.PostProcessServiceError(err, "Queue", "GetStats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListChannels Gets the list of IDs of non-empty channels.
// It will return an approximate list of IDs of non-empty channels. That information is based on the queue level statistics.
// API supports optional channelFilter parameter which will filter the returned results according to the specified filter.
// List of channel IDs is approximate, because statistics is refreshed once per-second, and that list represents a snapshot of the past information. API is paginated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/ListChannels.go.html to see an example of how to use ListChannels API.
// A default retry strategy applies to this operation ListChannels()
func (client QueueClient) ListChannels(ctx context.Context, request ListChannelsRequest) (response ListChannelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listChannels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListChannelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListChannelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListChannelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListChannelsResponse")
	}
	return
}

// listChannels implements the OCIOperation interface (enables retrying operations)
func (client QueueClient) listChannels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queues/{queueId}/channels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListChannelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/ChannelCollection/ListChannels"
		err = common.PostProcessServiceError(err, "Queue", "ListChannels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutMessages Puts messages into the queue.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to produce messages.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/PutMessages.go.html to see an example of how to use PutMessages API.
func (client QueueClient) PutMessages(ctx context.Context, request PutMessagesRequest) (response PutMessagesResponse, err error) {
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
func (client QueueClient) putMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/messages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutMessagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/PutMessage/PutMessages"
		err = common.PostProcessServiceError(err, "Queue", "PutMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMessage Updates the visibility of the message represented by the receipt.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to update messages.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/UpdateMessage.go.html to see an example of how to use UpdateMessage API.
// A default retry strategy applies to this operation UpdateMessage()
func (client QueueClient) UpdateMessage(ctx context.Context, request UpdateMessageRequest) (response UpdateMessageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMessage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMessageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMessageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMessageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMessageResponse")
	}
	return
}

// updateMessage implements the OCIOperation interface (enables retrying operations)
func (client QueueClient) updateMessage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/queues/{queueId}/messages/{messageReceipt}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMessageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/UpdatedMessage/UpdateMessage"
		err = common.PostProcessServiceError(err, "Queue", "UpdateMessage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMessages Updates multiple messages in the queue.
// You must use the messages endpoint (https://docs.cloud.oracle.com/iaas/Content/queue/messages.htm#messages__messages-endpoint) to update messages.
// The messages endpoint may be different for different queues. Use GetQueue to find the queue's `messagesEndpoint`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/UpdateMessages.go.html to see an example of how to use UpdateMessages API.
// A default retry strategy applies to this operation UpdateMessages()
func (client QueueClient) UpdateMessages(ctx context.Context, request UpdateMessagesRequest) (response UpdateMessagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMessages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMessagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMessagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMessagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMessagesResponse")
	}
	return
}

// updateMessages implements the OCIOperation interface (enables retrying operations)
func (client QueueClient) updateMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/messages/actions/updateMessages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMessagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/GetMessage/UpdateMessages"
		err = common.PostProcessServiceError(err, "Queue", "UpdateMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
