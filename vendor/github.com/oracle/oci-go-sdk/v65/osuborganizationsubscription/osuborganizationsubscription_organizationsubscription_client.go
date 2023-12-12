// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription Gateway API Organization's Subscription
//
// API that returns data for the list of subscription ids returned from Organizations API
//

package osuborganizationsubscription

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OrganizationSubscriptionClient a client for OrganizationSubscription
type OrganizationSubscriptionClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOrganizationSubscriptionClientWithConfigurationProvider Creates a new default OrganizationSubscription client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOrganizationSubscriptionClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OrganizationSubscriptionClient, err error) {
	if enabled := common.CheckForEnabledServices("osuborganizationsubscription"); !enabled {
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
	return newOrganizationSubscriptionClientFromBaseClient(baseClient, provider)
}

// NewOrganizationSubscriptionClientWithOboToken Creates a new default OrganizationSubscription client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOrganizationSubscriptionClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OrganizationSubscriptionClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOrganizationSubscriptionClientFromBaseClient(baseClient, configProvider)
}

func newOrganizationSubscriptionClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OrganizationSubscriptionClient, err error) {
	// OrganizationSubscription service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OrganizationSubscription"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OrganizationSubscriptionClient{BaseClient: baseClient}
	client.BasePath = "oalapp/service/onesubs/proxy/20210501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OrganizationSubscriptionClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osuborganizationsubscription", "https://csaap-e.oracle.com")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OrganizationSubscriptionClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OrganizationSubscriptionClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListOrganizationSubscriptions API that returns data for the list of subscription ids returned from Organizations API
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osuborganizationsubscription/ListOrganizationSubscriptions.go.html to see an example of how to use ListOrganizationSubscriptions API.
func (client OrganizationSubscriptionClient) ListOrganizationSubscriptions(ctx context.Context, request ListOrganizationSubscriptionsRequest) (response ListOrganizationSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOrganizationSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOrganizationSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOrganizationSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOrganizationSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOrganizationSubscriptionsResponse")
	}
	return
}

// listOrganizationSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client OrganizationSubscriptionClient) listOrganizationSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/organizationSubscription", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOrganizationSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OrganizationSubscription", "ListOrganizationSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
