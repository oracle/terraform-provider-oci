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

// QueueAdminClient a client for QueueAdmin
type QueueAdminClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewQueueAdminClientWithConfigurationProvider Creates a new default QueueAdmin client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewQueueAdminClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client QueueAdminClient, err error) {
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
	return newQueueAdminClientFromBaseClient(baseClient, provider)
}

// NewQueueAdminClientWithOboToken Creates a new default QueueAdmin client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewQueueAdminClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client QueueAdminClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newQueueAdminClientFromBaseClient(baseClient, configProvider)
}

func newQueueAdminClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client QueueAdminClient, err error) {
	// QueueAdmin service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("QueueAdmin"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = QueueAdminClient{BaseClient: baseClient}
	client.BasePath = "20210201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *QueueAdminClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("queue", "https://messaging.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *QueueAdminClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *QueueAdminClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddCapability Adds a capability to the queue such as CONSUMER_GROUPS.
// A default retry strategy applies to this operation AddCapability()
func (client QueueAdminClient) AddCapability(ctx context.Context, request AddCapabilityRequest) (response AddCapabilityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addCapability, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddCapabilityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddCapabilityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddCapabilityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddCapabilityResponse")
	}
	return
}

// addCapability implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) addCapability(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/actions/addCapability", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddCapabilityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/AddCapability"
		err = common.PostProcessServiceError(err, "QueueAdmin", "AddCapability", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeQueueCompartment Moves a queue from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangeQueueCompartment()
func (client QueueAdminClient) ChangeQueueCompartment(ctx context.Context, request ChangeQueueCompartmentRequest) (response ChangeQueueCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeQueueCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeQueueCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeQueueCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeQueueCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeQueueCompartmentResponse")
	}
	return
}

// changeQueueCompartment implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) changeQueueCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeQueueCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/ChangeQueueCompartment"
		err = common.PostProcessServiceError(err, "QueueAdmin", "ChangeQueueCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConsumerGroup Creates a new consumer group.
// A default retry strategy applies to this operation CreateConsumerGroup()
func (client QueueAdminClient) CreateConsumerGroup(ctx context.Context, request CreateConsumerGroupRequest) (response CreateConsumerGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConsumerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConsumerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConsumerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConsumerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConsumerGroupResponse")
	}
	return
}

// createConsumerGroup implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) createConsumerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/consumerGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConsumerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/ConsumerGroup/CreateConsumerGroup"
		err = common.PostProcessServiceError(err, "QueueAdmin", "CreateConsumerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateQueue Creates a new queue.
// A default retry strategy applies to this operation CreateQueue()
func (client QueueAdminClient) CreateQueue(ctx context.Context, request CreateQueueRequest) (response CreateQueueResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createQueue, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateQueueResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateQueueResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateQueueResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateQueueResponse")
	}
	return
}

// createQueue implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) createQueue(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateQueueResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/CreateQueue"
		err = common.PostProcessServiceError(err, "QueueAdmin", "CreateQueue", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConsumerGroup Deletes a consumer group resource by identifier.
// A default retry strategy applies to this operation DeleteConsumerGroup()
func (client QueueAdminClient) DeleteConsumerGroup(ctx context.Context, request DeleteConsumerGroupRequest) (response DeleteConsumerGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConsumerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConsumerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConsumerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConsumerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConsumerGroupResponse")
	}
	return
}

// deleteConsumerGroup implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) deleteConsumerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/consumerGroups/{consumerGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConsumerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/ConsumerGroup/DeleteConsumerGroup"
		err = common.PostProcessServiceError(err, "QueueAdmin", "DeleteConsumerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteQueue Deletes a queue resource by identifier.
// A default retry strategy applies to this operation DeleteQueue()
func (client QueueAdminClient) DeleteQueue(ctx context.Context, request DeleteQueueRequest) (response DeleteQueueResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteQueue, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteQueueResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteQueueResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteQueueResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteQueueResponse")
	}
	return
}

// deleteQueue implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) deleteQueue(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/queues/{queueId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteQueueResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/DeleteQueue"
		err = common.PostProcessServiceError(err, "QueueAdmin", "DeleteQueue", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConsumerGroup Gets a consumer group by identifier.
// A default retry strategy applies to this operation GetConsumerGroup()
func (client QueueAdminClient) GetConsumerGroup(ctx context.Context, request GetConsumerGroupRequest) (response GetConsumerGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConsumerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConsumerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConsumerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConsumerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConsumerGroupResponse")
	}
	return
}

// getConsumerGroup implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) getConsumerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/consumerGroups/{consumerGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConsumerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/ConsumerGroup/GetConsumerGroup"
		err = common.PostProcessServiceError(err, "QueueAdmin", "GetConsumerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQueue Gets a queue by identifier.
// A default retry strategy applies to this operation GetQueue()
func (client QueueAdminClient) GetQueue(ctx context.Context, request GetQueueRequest) (response GetQueueResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQueue, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQueueResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQueueResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQueueResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQueueResponse")
	}
	return
}

// getQueue implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) getQueue(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queues/{queueId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetQueueResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/GetQueue"
		err = common.PostProcessServiceError(err, "QueueAdmin", "GetQueue", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
// A default retry strategy applies to this operation GetWorkRequest()
func (client QueueAdminClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client QueueAdminClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "QueueAdmin", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConsumerGroups Returns a list of consumer groups.
// A default retry strategy applies to this operation ListConsumerGroups()
func (client QueueAdminClient) ListConsumerGroups(ctx context.Context, request ListConsumerGroupsRequest) (response ListConsumerGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConsumerGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConsumerGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConsumerGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConsumerGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConsumerGroupsResponse")
	}
	return
}

// listConsumerGroups implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) listConsumerGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/consumerGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConsumerGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/ConsumerGroupCollection/ListConsumerGroups"
		err = common.PostProcessServiceError(err, "QueueAdmin", "ListConsumerGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListQueues Returns a list of queues.
// A default retry strategy applies to this operation ListQueues()
func (client QueueAdminClient) ListQueues(ctx context.Context, request ListQueuesRequest) (response ListQueuesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQueues, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQueuesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQueuesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQueuesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQueuesResponse")
	}
	return
}

// listQueues implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) listQueues(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queues", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListQueuesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/QueueCollection/ListQueues"
		err = common.PostProcessServiceError(err, "QueueAdmin", "ListQueues", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client QueueAdminClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client QueueAdminClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/WorkRequestErrorCollection/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "QueueAdmin", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client QueueAdminClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client QueueAdminClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/WorkRequestLogEntryCollection/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "QueueAdmin", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client QueueAdminClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client QueueAdminClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/WorkRequestSummaryCollection/ListWorkRequests"
		err = common.PostProcessServiceError(err, "QueueAdmin", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PurgeQueue Deletes all messages present in the queue or in the specified consumer group, or deletes all the messages in the specific channel at the time of invocation.
// Only one concurrent purge operation is supported for any given queue.
// However multiple concurrent purge operations are supported for different queues.
// Purge request without specification of target channels will clean up all messages in the queue and in the child channels.
// Purge request without specification of consumer group will either clean up all messages in the queue or in the primary consumer group, depending on the presence of the CONSUMER_GROUPS capability on the queue.
// To purge all consumer groups, the special value 'all' can be used.
// A default retry strategy applies to this operation PurgeQueue()
func (client QueueAdminClient) PurgeQueue(ctx context.Context, request PurgeQueueRequest) (response PurgeQueueResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.purgeQueue, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PurgeQueueResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PurgeQueueResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PurgeQueueResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PurgeQueueResponse")
	}
	return
}

// purgeQueue implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) purgeQueue(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/actions/purge", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PurgeQueueResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/PurgeQueue"
		err = common.PostProcessServiceError(err, "QueueAdmin", "PurgeQueue", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveCapability Removes a capability from the queue such as CONSUMER_GROUPS.
// A default retry strategy applies to this operation RemoveCapability()
func (client QueueAdminClient) RemoveCapability(ctx context.Context, request RemoveCapabilityRequest) (response RemoveCapabilityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeCapability, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveCapabilityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveCapabilityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveCapabilityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveCapabilityResponse")
	}
	return
}

// removeCapability implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) removeCapability(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queues/{queueId}/actions/removeCapability", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveCapabilityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/RemoveCapability"
		err = common.PostProcessServiceError(err, "QueueAdmin", "RemoveCapability", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConsumerGroup Updates the specified consumer group.
// A default retry strategy applies to this operation UpdateConsumerGroup()
func (client QueueAdminClient) UpdateConsumerGroup(ctx context.Context, request UpdateConsumerGroupRequest) (response UpdateConsumerGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConsumerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConsumerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConsumerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConsumerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConsumerGroupResponse")
	}
	return
}

// updateConsumerGroup implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) updateConsumerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/consumerGroups/{consumerGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConsumerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/ConsumerGroup/UpdateConsumerGroup"
		err = common.PostProcessServiceError(err, "QueueAdmin", "UpdateConsumerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateQueue Updates the specified queue.
// A default retry strategy applies to this operation UpdateQueue()
func (client QueueAdminClient) UpdateQueue(ctx context.Context, request UpdateQueueRequest) (response UpdateQueueResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateQueue, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateQueueResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateQueueResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateQueueResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateQueueResponse")
	}
	return
}

// updateQueue implements the OCIOperation interface (enables retrying operations)
func (client QueueAdminClient) updateQueue(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/queues/{queueId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateQueueResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/queue/20210201/Queue/UpdateQueue"
		err = common.PostProcessServiceError(err, "QueueAdmin", "UpdateQueue", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
