// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//PhoneMessagingClient a client for PhoneMessaging
type PhoneMessagingClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPhoneMessagingClientWithConfigurationProvider Creates a new default PhoneMessaging client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPhoneMessagingClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PhoneMessagingClient, err error) {
	if enabled := common.CheckForEnabledServices("ons"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newPhoneMessagingClientFromBaseClient(baseClient, provider)
}

// NewPhoneMessagingClientWithOboToken Creates a new default PhoneMessaging client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewPhoneMessagingClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PhoneMessagingClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPhoneMessagingClientFromBaseClient(baseClient, configProvider)
}

func newPhoneMessagingClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PhoneMessagingClient, err error) {
	// PhoneMessaging service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("PhoneMessaging"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PhoneMessagingClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PhoneMessagingClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("notification", "https://notification.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PhoneMessagingClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *PhoneMessagingClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// SendMessage Sends a message.
// A default retry strategy applies to this operation SendMessage()
func (client PhoneMessagingClient) SendMessage(ctx context.Context, request SendMessageRequest) (response SendMessageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.sendMessage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SendMessageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SendMessageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SendMessageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SendMessageResponse")
	}
	return
}

// sendMessage implements the OCIOperation interface (enables retrying operations)
func (client PhoneMessagingClient) sendMessage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/phoneNumbers/{phoneNumberId}/actions/sendMessage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SendMessageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/PhoneNumber/SendMessage"
		err = common.PostProcessServiceError(err, "PhoneMessaging", "SendMessage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
