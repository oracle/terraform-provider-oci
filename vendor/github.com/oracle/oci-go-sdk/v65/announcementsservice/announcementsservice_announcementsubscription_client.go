// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// AnnouncementSubscriptionClient a client for AnnouncementSubscription
type AnnouncementSubscriptionClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAnnouncementSubscriptionClientWithConfigurationProvider Creates a new default AnnouncementSubscription client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAnnouncementSubscriptionClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AnnouncementSubscriptionClient, err error) {
	if enabled := common.CheckForEnabledServices("announcementsservice"); !enabled {
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
	return newAnnouncementSubscriptionClientFromBaseClient(baseClient, provider)
}

// NewAnnouncementSubscriptionClientWithOboToken Creates a new default AnnouncementSubscription client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAnnouncementSubscriptionClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AnnouncementSubscriptionClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAnnouncementSubscriptionClientFromBaseClient(baseClient, configProvider)
}

func newAnnouncementSubscriptionClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AnnouncementSubscriptionClient, err error) {
	// AnnouncementSubscription service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AnnouncementSubscription"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AnnouncementSubscriptionClient{BaseClient: baseClient}
	client.BasePath = "20180904"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AnnouncementSubscriptionClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("announcements", "https://announcements.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AnnouncementSubscriptionClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AnnouncementSubscriptionClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAnnouncementSubscriptionCompartment Moves the specified announcement subscription from one compartment to another compartment. When provided, If-Match is checked against ETag values of the resource.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ChangeAnnouncementSubscriptionCompartment.go.html to see an example of how to use ChangeAnnouncementSubscriptionCompartment API.
func (client AnnouncementSubscriptionClient) ChangeAnnouncementSubscriptionCompartment(ctx context.Context, request ChangeAnnouncementSubscriptionCompartmentRequest) (response ChangeAnnouncementSubscriptionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeAnnouncementSubscriptionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAnnouncementSubscriptionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAnnouncementSubscriptionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAnnouncementSubscriptionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAnnouncementSubscriptionCompartmentResponse")
	}
	return
}

// changeAnnouncementSubscriptionCompartment implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) changeAnnouncementSubscriptionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/announcementSubscriptions/{announcementSubscriptionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAnnouncementSubscriptionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/ChangeAnnouncementSubscriptionCompartment"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "ChangeAnnouncementSubscriptionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAnnouncementSubscription Creates a new announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/CreateAnnouncementSubscription.go.html to see an example of how to use CreateAnnouncementSubscription API.
func (client AnnouncementSubscriptionClient) CreateAnnouncementSubscription(ctx context.Context, request CreateAnnouncementSubscriptionRequest) (response CreateAnnouncementSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAnnouncementSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAnnouncementSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAnnouncementSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAnnouncementSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAnnouncementSubscriptionResponse")
	}
	return
}

// createAnnouncementSubscription implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) createAnnouncementSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/announcementSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAnnouncementSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/CreateAnnouncementSubscription"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "CreateAnnouncementSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFilterGroup Creates a new filter group in the specified announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/CreateFilterGroup.go.html to see an example of how to use CreateFilterGroup API.
func (client AnnouncementSubscriptionClient) CreateFilterGroup(ctx context.Context, request CreateFilterGroupRequest) (response CreateFilterGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFilterGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFilterGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFilterGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFilterGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFilterGroupResponse")
	}
	return
}

// createFilterGroup implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) createFilterGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/announcementSubscriptions/{announcementSubscriptionId}/filterGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFilterGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/CreateFilterGroup"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "CreateFilterGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAnnouncementSubscription Deletes the specified announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/DeleteAnnouncementSubscription.go.html to see an example of how to use DeleteAnnouncementSubscription API.
func (client AnnouncementSubscriptionClient) DeleteAnnouncementSubscription(ctx context.Context, request DeleteAnnouncementSubscriptionRequest) (response DeleteAnnouncementSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAnnouncementSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAnnouncementSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAnnouncementSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAnnouncementSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAnnouncementSubscriptionResponse")
	}
	return
}

// deleteAnnouncementSubscription implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) deleteAnnouncementSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/announcementSubscriptions/{announcementSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAnnouncementSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/DeleteAnnouncementSubscription"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "DeleteAnnouncementSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFilterGroup Deletes a filter group in the specified announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/DeleteFilterGroup.go.html to see an example of how to use DeleteFilterGroup API.
func (client AnnouncementSubscriptionClient) DeleteFilterGroup(ctx context.Context, request DeleteFilterGroupRequest) (response DeleteFilterGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFilterGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFilterGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFilterGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFilterGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFilterGroupResponse")
	}
	return
}

// deleteFilterGroup implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) deleteFilterGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/announcementSubscriptions/{announcementSubscriptionId}/filterGroups/{filterGroupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFilterGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/DeleteFilterGroup"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "DeleteFilterGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnnouncementSubscription Gets the specified announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/GetAnnouncementSubscription.go.html to see an example of how to use GetAnnouncementSubscription API.
func (client AnnouncementSubscriptionClient) GetAnnouncementSubscription(ctx context.Context, request GetAnnouncementSubscriptionRequest) (response GetAnnouncementSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnnouncementSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnnouncementSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnnouncementSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnnouncementSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnnouncementSubscriptionResponse")
	}
	return
}

// getAnnouncementSubscription implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) getAnnouncementSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcementSubscriptions/{announcementSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnnouncementSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/GetAnnouncementSubscription"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "GetAnnouncementSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAnnouncementSubscriptions Gets a list of all announcement subscriptions in the specified compartment.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListAnnouncementSubscriptions.go.html to see an example of how to use ListAnnouncementSubscriptions API.
func (client AnnouncementSubscriptionClient) ListAnnouncementSubscriptions(ctx context.Context, request ListAnnouncementSubscriptionsRequest) (response ListAnnouncementSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnnouncementSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnnouncementSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnnouncementSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnnouncementSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnnouncementSubscriptionsResponse")
	}
	return
}

// listAnnouncementSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) listAnnouncementSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcementSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAnnouncementSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscriptionCollection/ListAnnouncementSubscriptions"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "ListAnnouncementSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAnnouncementSubscription Updates the specified announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/UpdateAnnouncementSubscription.go.html to see an example of how to use UpdateAnnouncementSubscription API.
func (client AnnouncementSubscriptionClient) UpdateAnnouncementSubscription(ctx context.Context, request UpdateAnnouncementSubscriptionRequest) (response UpdateAnnouncementSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnnouncementSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnnouncementSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnnouncementSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnnouncementSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnnouncementSubscriptionResponse")
	}
	return
}

// updateAnnouncementSubscription implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) updateAnnouncementSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/announcementSubscriptions/{announcementSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAnnouncementSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/UpdateAnnouncementSubscription"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "UpdateAnnouncementSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFilterGroup Updates a filter group in the specified announcement subscription.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/UpdateFilterGroup.go.html to see an example of how to use UpdateFilterGroup API.
func (client AnnouncementSubscriptionClient) UpdateFilterGroup(ctx context.Context, request UpdateFilterGroupRequest) (response UpdateFilterGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFilterGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFilterGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFilterGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFilterGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFilterGroupResponse")
	}
	return
}

// updateFilterGroup implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementSubscriptionClient) updateFilterGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/announcementSubscriptions/{announcementSubscriptionId}/filterGroups/{filterGroupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFilterGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementSubscription/UpdateFilterGroup"
		err = common.PostProcessServiceError(err, "AnnouncementSubscription", "UpdateFilterGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
