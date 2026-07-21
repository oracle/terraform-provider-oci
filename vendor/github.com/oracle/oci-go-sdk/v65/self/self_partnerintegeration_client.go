// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// PartnerIntegerationClient a client for PartnerIntegeration
type PartnerIntegerationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPartnerIntegerationClientWithConfigurationProvider Creates a new default PartnerIntegeration client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPartnerIntegerationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PartnerIntegerationClient, err error) {
	if enabled := common.CheckForEnabledServices("self"); !enabled {
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
	return newPartnerIntegerationClientFromBaseClient(baseClient, provider)
}

// NewPartnerIntegerationClientWithOboToken Creates a new default PartnerIntegeration client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewPartnerIntegerationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PartnerIntegerationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPartnerIntegerationClientFromBaseClient(baseClient, configProvider)
}

func newPartnerIntegerationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PartnerIntegerationClient, err error) {
	// PartnerIntegeration service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("PartnerIntegeration"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PartnerIntegerationClient{BaseClient: baseClient}
	client.BasePath = "20260129"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PartnerIntegerationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("self", "https://self.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PartnerIntegerationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PartnerIntegerationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateSubscription Activates a subscription identified by the provided subscription ID and updates its status.
// A default retry strategy applies to this operation ActivateSubscription()
func (client PartnerIntegerationClient) ActivateSubscription(ctx context.Context, request ActivateSubscriptionRequest) (response ActivateSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateSubscriptionResponse")
	}
	return
}

// activateSubscription implements the OCIOperation interface (enables retrying operations)
func (client PartnerIntegerationClient) activateSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/partners/subscriptions/{subscriptionId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "partnerIntegeration", "ActivateSubscription")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/self/20260129/PartnerSubscription/ActivateSubscription"
		err = common.PostProcessServiceError(err, "PartnerIntegeration", "ActivateSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPartners Lists marketplace publisher partner info for a compartment.
// A default retry strategy applies to this operation ListPartners()
func (client PartnerIntegerationClient) ListPartners(ctx context.Context, request ListPartnersRequest) (response ListPartnersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPartners, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPartnersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPartnersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPartnersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPartnersResponse")
	}
	return
}

// listPartners implements the OCIOperation interface (enables retrying operations)
func (client PartnerIntegerationClient) listPartners(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/partners", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPartnersResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "partnerIntegeration", "ListPartners")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/self/20260129/Partner/ListPartners"
		err = common.PostProcessServiceError(err, "PartnerIntegeration", "ListPartners", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListingSubscriptions Gets information about a Subscription.
// A default retry strategy applies to this operation ListingSubscriptions()
func (client PartnerIntegerationClient) ListingSubscriptions(ctx context.Context, request ListingSubscriptionsRequest) (response ListingSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listingSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListingSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListingSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListingSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListingSubscriptionsResponse")
	}
	return
}

// listingSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client PartnerIntegerationClient) listingSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/partners/subscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListingSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "partnerIntegeration", "ListingSubscriptions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/self/20260129/ListingSubscriptionsCollection/ListingSubscriptions"
		err = common.PostProcessServiceError(err, "PartnerIntegeration", "ListingSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResolveSubscription This API returns the subscription details by resolving JWT token to corresponding subscription and move its state to Pending Activation state.
// A default retry strategy applies to this operation ResolveSubscription()
func (client PartnerIntegerationClient) ResolveSubscription(ctx context.Context, request ResolveSubscriptionRequest) (response ResolveSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resolveSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResolveSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResolveSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResolveSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResolveSubscriptionResponse")
	}
	return
}

// resolveSubscription implements the OCIOperation interface (enables retrying operations)
func (client PartnerIntegerationClient) resolveSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/partners/subscriptions/actions/resolve", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResolveSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "partnerIntegeration", "ResolveSubscription")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/self/20260129/PartnerSubscription/ResolveSubscription"
		err = common.PostProcessServiceError(err, "PartnerIntegeration", "ResolveSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SubmitSubscriptionUsageBatch Asynchronously submits a UTF-8 CSV usage file for marketplace offers. The file
// must not exceed 50 MB or 10,000 rows and must include required usage columns.
// A default retry strategy applies to this operation SubmitSubscriptionUsageBatch()
func (client PartnerIntegerationClient) SubmitSubscriptionUsageBatch(ctx context.Context, request SubmitSubscriptionUsageBatchRequest) (response SubmitSubscriptionUsageBatchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.submitSubscriptionUsageBatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubmitSubscriptionUsageBatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubmitSubscriptionUsageBatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubmitSubscriptionUsageBatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubmitSubscriptionUsageBatchResponse")
	}
	return
}

// submitSubscriptionUsageBatch implements the OCIOperation interface (enables retrying operations)
func (client PartnerIntegerationClient) submitSubscriptionUsageBatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/partners/actions/batchUsageRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubmitSubscriptionUsageBatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "partnerIntegeration", "SubmitSubscriptionUsageBatch")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/self/20260129/CreateSubscriptionUsageRecordDetails/SubmitSubscriptionUsageBatch"
		err = common.PostProcessServiceError(err, "PartnerIntegeration", "SubmitSubscriptionUsageBatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SubmitSubscriptionUsageRecords Synchronously submits usage records for marketplace offers. Each record must include
// `id`, `marketplaceOfferId`, `amount`, `currencyCode`, `timeUsageStarted`,
// `timeUsageEnded`, and `usageDimensionName`.
// A default retry strategy applies to this operation SubmitSubscriptionUsageRecords()
func (client PartnerIntegerationClient) SubmitSubscriptionUsageRecords(ctx context.Context, request SubmitSubscriptionUsageRecordsRequest) (response SubmitSubscriptionUsageRecordsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.submitSubscriptionUsageRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubmitSubscriptionUsageRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubmitSubscriptionUsageRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubmitSubscriptionUsageRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubmitSubscriptionUsageRecordsResponse")
	}
	return
}

// submitSubscriptionUsageRecords implements the OCIOperation interface (enables retrying operations)
func (client PartnerIntegerationClient) submitSubscriptionUsageRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/partners/actions/submitUsageRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubmitSubscriptionUsageRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "partnerIntegeration", "SubmitSubscriptionUsageRecords")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/self/20260129/CreateSubscriptionUsageRecordDetails/SubmitSubscriptionUsageRecords"
		err = common.PostProcessServiceError(err, "PartnerIntegeration", "SubmitSubscriptionUsageRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
