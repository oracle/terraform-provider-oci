// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// NotificationControlPlaneClient a client for NotificationControlPlane
type NotificationControlPlaneClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewNotificationControlPlaneClientWithConfigurationProvider Creates a new default NotificationControlPlane client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewNotificationControlPlaneClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client NotificationControlPlaneClient, err error) {
	if enabled := common.CheckForEnabledServices("ons"); !enabled {
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
	return newNotificationControlPlaneClientFromBaseClient(baseClient, provider)
}

// NewNotificationControlPlaneClientWithOboToken Creates a new default NotificationControlPlane client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewNotificationControlPlaneClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client NotificationControlPlaneClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newNotificationControlPlaneClientFromBaseClient(baseClient, configProvider)
}

func newNotificationControlPlaneClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client NotificationControlPlaneClient, err error) {
	// NotificationControlPlane service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("NotificationControlPlane"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = NotificationControlPlaneClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *NotificationControlPlaneClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("notification", "https://notification.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *NotificationControlPlaneClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *NotificationControlPlaneClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeTopicCompartment Moves a topic into a different compartment within the same tenancy. For information about moving resources
// between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/ChangeTopicCompartment.go.html to see an example of how to use ChangeTopicCompartment API.
func (client NotificationControlPlaneClient) ChangeTopicCompartment(ctx context.Context, request ChangeTopicCompartmentRequest) (response ChangeTopicCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTopicCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTopicCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTopicCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTopicCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTopicCompartmentResponse")
	}
	return
}

// changeTopicCompartment implements the OCIOperation interface (enables retrying operations)
func (client NotificationControlPlaneClient) changeTopicCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/topics/{topicId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTopicCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/ChangeTopicCompartment"
		err = common.PostProcessServiceError(err, "NotificationControlPlane", "ChangeTopicCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTopic Creates a topic in the specified compartment. For general information about topics, see
// Managing Topics and Subscriptions (https://docs.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the topic to reside.
// For information about access control and compartments, see Overview of the IAM Service (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
// You must specify a display name for the topic.
// All Oracle Cloud Infrastructure resources, including topics, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response. You can also
// retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the resource in the
// Console. For more information, see Resource Identifiers (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/CreateTopic.go.html to see an example of how to use CreateTopic API.
func (client NotificationControlPlaneClient) CreateTopic(ctx context.Context, request CreateTopicRequest) (response CreateTopicResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTopic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTopicResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTopicResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTopicResponse")
	}
	return
}

// createTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationControlPlaneClient) createTopic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/topics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/CreateTopic"
		err = common.PostProcessServiceError(err, "NotificationControlPlane", "CreateTopic", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTopic Deletes the specified topic.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/DeleteTopic.go.html to see an example of how to use DeleteTopic API.
func (client NotificationControlPlaneClient) DeleteTopic(ctx context.Context, request DeleteTopicRequest) (response DeleteTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTopic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTopicResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTopicResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTopicResponse")
	}
	return
}

// deleteTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationControlPlaneClient) deleteTopic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/topics/{topicId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/DeleteTopic"
		err = common.PostProcessServiceError(err, "NotificationControlPlane", "DeleteTopic", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTopic Gets the specified topic's configuration information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/GetTopic.go.html to see an example of how to use GetTopic API.
func (client NotificationControlPlaneClient) GetTopic(ctx context.Context, request GetTopicRequest) (response GetTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTopic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTopicResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTopicResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTopicResponse")
	}
	return
}

// getTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationControlPlaneClient) getTopic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/topics/{topicId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/GetTopic"
		err = common.PostProcessServiceError(err, "NotificationControlPlane", "GetTopic", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTopics Lists topics in the specified compartment.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 120.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/ListTopics.go.html to see an example of how to use ListTopics API.
func (client NotificationControlPlaneClient) ListTopics(ctx context.Context, request ListTopicsRequest) (response ListTopicsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTopics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTopicsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTopicsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTopicsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTopicsResponse")
	}
	return
}

// listTopics implements the OCIOperation interface (enables retrying operations)
func (client NotificationControlPlaneClient) listTopics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/topics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTopicsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/ListTopics"
		err = common.PostProcessServiceError(err, "NotificationControlPlane", "ListTopics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTopic Updates the specified topic's configuration.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/UpdateTopic.go.html to see an example of how to use UpdateTopic API.
func (client NotificationControlPlaneClient) UpdateTopic(ctx context.Context, request UpdateTopicRequest) (response UpdateTopicResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTopic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTopicResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTopicResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTopicResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTopicResponse")
	}
	return
}

// updateTopic implements the OCIOperation interface (enables retrying operations)
func (client NotificationControlPlaneClient) updateTopic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/topics/{topicId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTopicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/UpdateTopic"
		err = common.PostProcessServiceError(err, "NotificationControlPlane", "UpdateTopic", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
