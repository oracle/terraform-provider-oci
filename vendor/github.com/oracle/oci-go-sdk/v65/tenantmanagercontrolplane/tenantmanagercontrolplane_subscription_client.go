// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SubscriptionClient a client for Subscription
type SubscriptionClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSubscriptionClientWithConfigurationProvider Creates a new default Subscription client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSubscriptionClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SubscriptionClient, err error) {
	if enabled := common.CheckForEnabledServices("tenantmanagercontrolplane"); !enabled {
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
	return newSubscriptionClientFromBaseClient(baseClient, provider)
}

// NewSubscriptionClientWithOboToken Creates a new default Subscription client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSubscriptionClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SubscriptionClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSubscriptionClientFromBaseClient(baseClient, configProvider)
}

func newSubscriptionClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SubscriptionClient, err error) {
	// Subscription service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Subscription"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SubscriptionClient{BaseClient: baseClient}
	client.BasePath = "20230401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SubscriptionClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SubscriptionClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SubscriptionClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateSubscriptionMapping Assign the tenancy record identified by the compartment ID to the given subscription ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/CreateSubscriptionMapping.go.html to see an example of how to use CreateSubscriptionMapping API.
func (client SubscriptionClient) CreateSubscriptionMapping(ctx context.Context, request CreateSubscriptionMappingRequest) (response CreateSubscriptionMappingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSubscriptionMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSubscriptionMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSubscriptionMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSubscriptionMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSubscriptionMappingResponse")
	}
	return
}

// createSubscriptionMapping implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) createSubscriptionMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptionMappings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSubscriptionMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SubscriptionMapping/CreateSubscriptionMapping"
		err = common.PostProcessServiceError(err, "Subscription", "CreateSubscriptionMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSubscriptionMapping Delete the subscription mapping details by subscription mapping ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/DeleteSubscriptionMapping.go.html to see an example of how to use DeleteSubscriptionMapping API.
func (client SubscriptionClient) DeleteSubscriptionMapping(ctx context.Context, request DeleteSubscriptionMappingRequest) (response DeleteSubscriptionMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSubscriptionMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSubscriptionMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSubscriptionMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSubscriptionMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSubscriptionMappingResponse")
	}
	return
}

// deleteSubscriptionMapping implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) deleteSubscriptionMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/subscriptionMappings/{subscriptionMappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSubscriptionMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SubscriptionMapping/DeleteSubscriptionMapping"
		err = common.PostProcessServiceError(err, "Subscription", "DeleteSubscriptionMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssignedSubscription Get the assigned subscription details by assigned subscription ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetAssignedSubscription.go.html to see an example of how to use GetAssignedSubscription API.
func (client SubscriptionClient) GetAssignedSubscription(ctx context.Context, request GetAssignedSubscriptionRequest) (response GetAssignedSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssignedSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssignedSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssignedSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssignedSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssignedSubscriptionResponse")
	}
	return
}

// getAssignedSubscription implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) getAssignedSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assignedSubscriptions/{assignedSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssignedSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/AssignedSubscription/GetAssignedSubscription"
		err = common.PostProcessServiceError(err, "Subscription", "GetAssignedSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assignedsubscription{})
	return response, err
}

// GetSubscription Gets the subscription details by subscription ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetSubscription.go.html to see an example of how to use GetSubscription API.
func (client SubscriptionClient) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (response GetSubscriptionResponse, err error) {
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
func (client SubscriptionClient) getSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/Subscription/GetSubscription"
		err = common.PostProcessServiceError(err, "Subscription", "GetSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &subscription{})
	return response, err
}

// GetSubscriptionMapping Get the subscription mapping details by subscription mapping ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetSubscriptionMapping.go.html to see an example of how to use GetSubscriptionMapping API.
func (client SubscriptionClient) GetSubscriptionMapping(ctx context.Context, request GetSubscriptionMappingRequest) (response GetSubscriptionMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSubscriptionMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSubscriptionMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSubscriptionMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSubscriptionMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSubscriptionMappingResponse")
	}
	return
}

// getSubscriptionMapping implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) getSubscriptionMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptionMappings/{subscriptionMappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSubscriptionMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SubscriptionMapping/GetSubscriptionMapping"
		err = common.PostProcessServiceError(err, "Subscription", "GetSubscriptionMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssignedSubscriptionLineItems List line item summaries that a assigned subscription owns.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListAssignedSubscriptionLineItems.go.html to see an example of how to use ListAssignedSubscriptionLineItems API.
// A default retry strategy applies to this operation ListAssignedSubscriptionLineItems()
func (client SubscriptionClient) ListAssignedSubscriptionLineItems(ctx context.Context, request ListAssignedSubscriptionLineItemsRequest) (response ListAssignedSubscriptionLineItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssignedSubscriptionLineItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssignedSubscriptionLineItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssignedSubscriptionLineItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssignedSubscriptionLineItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssignedSubscriptionLineItemsResponse")
	}
	return
}

// listAssignedSubscriptionLineItems implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) listAssignedSubscriptionLineItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assignedSubscriptions/{assignedSubscriptionId}/assignedSubscriptionLineItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssignedSubscriptionLineItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/AssignedSubscriptionLineItemSummary/ListAssignedSubscriptionLineItems"
		err = common.PostProcessServiceError(err, "Subscription", "ListAssignedSubscriptionLineItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssignedSubscriptions Lists subscriptions that are consumed by the compartment. Only the root compartment is allowed.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListAssignedSubscriptions.go.html to see an example of how to use ListAssignedSubscriptions API.
func (client SubscriptionClient) ListAssignedSubscriptions(ctx context.Context, request ListAssignedSubscriptionsRequest) (response ListAssignedSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssignedSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssignedSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssignedSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssignedSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssignedSubscriptionsResponse")
	}
	return
}

// listAssignedSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) listAssignedSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assignedSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssignedSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/AssignedSubscription/ListAssignedSubscriptions"
		err = common.PostProcessServiceError(err, "Subscription", "ListAssignedSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableRegions List the available regions based on subscription ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListAvailableRegions.go.html to see an example of how to use ListAvailableRegions API.
func (client SubscriptionClient) ListAvailableRegions(ctx context.Context, request ListAvailableRegionsRequest) (response ListAvailableRegionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailableRegions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailableRegionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailableRegionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailableRegionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailableRegionsResponse")
	}
	return
}

// listAvailableRegions implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) listAvailableRegions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}/availableRegions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailableRegionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/Subscription/ListAvailableRegions"
		err = common.PostProcessServiceError(err, "Subscription", "ListAvailableRegions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSubscriptionLineItems Lists the line items in a subscription.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSubscriptionLineItems.go.html to see an example of how to use ListSubscriptionLineItems API.
// A default retry strategy applies to this operation ListSubscriptionLineItems()
func (client SubscriptionClient) ListSubscriptionLineItems(ctx context.Context, request ListSubscriptionLineItemsRequest) (response ListSubscriptionLineItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSubscriptionLineItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSubscriptionLineItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSubscriptionLineItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSubscriptionLineItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSubscriptionLineItemsResponse")
	}
	return
}

// listSubscriptionLineItems implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) listSubscriptionLineItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}/subscriptionLineItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSubscriptionLineItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SubscriptionLineItemSummary/ListSubscriptionLineItems"
		err = common.PostProcessServiceError(err, "Subscription", "ListSubscriptionLineItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSubscriptionMappings Lists the subscription mappings for all the subscriptions owned by a given compartmentId. Only the root compartment is allowed.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSubscriptionMappings.go.html to see an example of how to use ListSubscriptionMappings API.
func (client SubscriptionClient) ListSubscriptionMappings(ctx context.Context, request ListSubscriptionMappingsRequest) (response ListSubscriptionMappingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSubscriptionMappings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSubscriptionMappingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSubscriptionMappingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSubscriptionMappingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSubscriptionMappingsResponse")
	}
	return
}

// listSubscriptionMappings implements the OCIOperation interface (enables retrying operations)
func (client SubscriptionClient) listSubscriptionMappings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptionMappings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSubscriptionMappingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SubscriptionMapping/ListSubscriptionMappings"
		err = common.PostProcessServiceError(err, "Subscription", "ListSubscriptionMappings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSubscriptions List the subscriptions that a compartment owns. Only the root compartment is allowed.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSubscriptions.go.html to see an example of how to use ListSubscriptions API.
func (client SubscriptionClient) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (response ListSubscriptionsResponse, err error) {
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
func (client SubscriptionClient) listSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/Subscription/ListSubscriptions"
		err = common.PostProcessServiceError(err, "Subscription", "ListSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
