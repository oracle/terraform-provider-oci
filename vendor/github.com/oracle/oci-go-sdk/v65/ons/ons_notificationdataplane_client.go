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

// NotificationDataPlaneClient a client for NotificationDataPlane
type NotificationDataPlaneClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewNotificationDataPlaneClientWithConfigurationProvider Creates a new default NotificationDataPlane client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewNotificationDataPlaneClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client NotificationDataPlaneClient, err error) {
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
	return newNotificationDataPlaneClientFromBaseClient(baseClient, provider)
}

// NewNotificationDataPlaneClientWithOboToken Creates a new default NotificationDataPlane client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewNotificationDataPlaneClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client NotificationDataPlaneClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newNotificationDataPlaneClientFromBaseClient(baseClient, configProvider)
}

func newNotificationDataPlaneClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client NotificationDataPlaneClient, err error) {
	// NotificationDataPlane service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("NotificationDataPlane"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = NotificationDataPlaneClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *NotificationDataPlaneClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("notification", "https://notification.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *NotificationDataPlaneClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *NotificationDataPlaneClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeSubscriptionCompartment Moves a subscription into a different compartment within the same tenancy. For information about moving
// resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/ChangeSubscriptionCompartment.go.html to see an example of how to use ChangeSubscriptionCompartment API.
func (client NotificationDataPlaneClient) ChangeSubscriptionCompartment(ctx context.Context, request ChangeSubscriptionCompartmentRequest) (response ChangeSubscriptionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSubscriptionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSubscriptionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSubscriptionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSubscriptionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSubscriptionCompartmentResponse")
	}
	return
}

// changeSubscriptionCompartment implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) changeSubscriptionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptions/{subscriptionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSubscriptionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/ChangeSubscriptionCompartment"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "ChangeSubscriptionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSubscription Creates a subscription for the specified topic and sends a subscription confirmation URL to the endpoint. The subscription remains in "Pending" status until it has been confirmed.
// For information about confirming subscriptions, see
// To confirm a subscription (https://docs.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#confirmSub).
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/CreateSubscription.go.html to see an example of how to use CreateSubscription API.
func (client NotificationDataPlaneClient) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (response CreateSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSubscriptionResponse")
	}
	return
}

// createSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) createSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/CreateSubscription"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "CreateSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSubscription Deletes the specified subscription.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/DeleteSubscription.go.html to see an example of how to use DeleteSubscription API.
func (client NotificationDataPlaneClient) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) (response DeleteSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSubscriptionResponse")
	}
	return
}

// deleteSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) deleteSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/subscriptions/{subscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/DeleteSubscription"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "DeleteSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfirmSubscription Gets the confirmation details for the specified subscription.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/GetConfirmSubscription.go.html to see an example of how to use GetConfirmSubscription API.
func (client NotificationDataPlaneClient) GetConfirmSubscription(ctx context.Context, request GetConfirmSubscriptionRequest) (response GetConfirmSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfirmSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfirmSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfirmSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfirmSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfirmSubscriptionResponse")
	}
	return
}

// getConfirmSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) getConfirmSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{id}/confirmation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfirmSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/GetConfirmSubscription"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "GetConfirmSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSubscription Gets the specified subscription's configuration information.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/GetSubscription.go.html to see an example of how to use GetSubscription API.
func (client NotificationDataPlaneClient) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (response GetSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSubscriptionResponse")
	}
	return
}

// getSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) getSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/GetSubscription"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "GetSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUnsubscription Unsubscribes the subscription from the topic.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/GetUnsubscription.go.html to see an example of how to use GetUnsubscription API.
func (client NotificationDataPlaneClient) GetUnsubscription(ctx context.Context, request GetUnsubscriptionRequest) (response GetUnsubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUnsubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUnsubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUnsubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUnsubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUnsubscriptionResponse")
	}
	return
}

// getUnsubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) getUnsubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{id}/unsubscription", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUnsubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/GetUnsubscription"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "GetUnsubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSubscriptions Lists the subscriptions in the specified compartment or topic.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/ListSubscriptions.go.html to see an example of how to use ListSubscriptions API.
func (client NotificationDataPlaneClient) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (response ListSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSubscriptionsResponse")
	}
	return
}

// listSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) listSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/ListSubscriptions"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "ListSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishMessage Publishes a message to the specified topic.
// The topic endpoint is required for this operation.
// To get the topic endpoint, use GetTopic
// and review the `apiEndpoint` value in the response (NotificationTopic).
// Limits information follows.
// Message size limit per request: 64KB.
// Message delivery rate limit per endpoint: 60 messages per minute for HTTP-based protocols, 10 messages per minute for the `EMAIL` protocol.
// HTTP-based protocols use URL endpoints that begin with "http:" or "https:".
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60 per topic. (This TPM limit represents messages per minute.)
// For more information about publishing messages, see Publishing Messages (https://docs.oracle.com/iaas/Content/Notification/Tasks/publishingmessages.htm).
// For steps to request a limit increase, see Requesting a Service Limit Increase (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm#three).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/PublishMessage.go.html to see an example of how to use PublishMessage API.
func (client NotificationDataPlaneClient) PublishMessage(ctx context.Context, request PublishMessageRequest) (response PublishMessageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.publishMessage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishMessageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishMessageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishMessageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishMessageResponse")
	}
	return
}

// publishMessage implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) publishMessage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/topics/{topicId}/messages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishMessageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/NotificationTopic/PublishMessage"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "PublishMessage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResendSubscriptionConfirmation Resends the confirmation details for the specified subscription.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/ResendSubscriptionConfirmation.go.html to see an example of how to use ResendSubscriptionConfirmation API.
func (client NotificationDataPlaneClient) ResendSubscriptionConfirmation(ctx context.Context, request ResendSubscriptionConfirmationRequest) (response ResendSubscriptionConfirmationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resendSubscriptionConfirmation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResendSubscriptionConfirmationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResendSubscriptionConfirmationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResendSubscriptionConfirmationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResendSubscriptionConfirmationResponse")
	}
	return
}

// resendSubscriptionConfirmation implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) resendSubscriptionConfirmation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptions/{id}/resendConfirmation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResendSubscriptionConfirmationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/ResendSubscriptionConfirmation"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "ResendSubscriptionConfirmation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSubscription Updates the specified subscription's configuration.
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/UpdateSubscription.go.html to see an example of how to use UpdateSubscription API.
func (client NotificationDataPlaneClient) UpdateSubscription(ctx context.Context, request UpdateSubscriptionRequest) (response UpdateSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSubscriptionResponse")
	}
	return
}

// updateSubscription implements the OCIOperation interface (enables retrying operations)
func (client NotificationDataPlaneClient) updateSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/subscriptions/{subscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/Subscription/UpdateSubscription"
		err = common.PostProcessServiceError(err, "NotificationDataPlane", "UpdateSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
