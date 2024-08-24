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

// AnnouncementClient a client for Announcement
type AnnouncementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAnnouncementClientWithConfigurationProvider Creates a new default Announcement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAnnouncementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AnnouncementClient, err error) {
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
	return newAnnouncementClientFromBaseClient(baseClient, provider)
}

// NewAnnouncementClientWithOboToken Creates a new default Announcement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAnnouncementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AnnouncementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAnnouncementClientFromBaseClient(baseClient, configProvider)
}

func newAnnouncementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AnnouncementClient, err error) {
	// Announcement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Announcement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AnnouncementClient{BaseClient: baseClient}
	client.BasePath = "20180904"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AnnouncementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("announcements", "https://announcements.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AnnouncementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AnnouncementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetAnnouncement Gets the details of a specific announcement.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/GetAnnouncement.go.html to see an example of how to use GetAnnouncement API.
func (client AnnouncementClient) GetAnnouncement(ctx context.Context, request GetAnnouncementRequest) (response GetAnnouncementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnnouncement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnnouncementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnnouncementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnnouncementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnnouncementResponse")
	}
	return
}

// getAnnouncement implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementClient) getAnnouncement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcements/{announcementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnnouncementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/Announcement/GetAnnouncement"
		err = common.PostProcessServiceError(err, "Announcement", "GetAnnouncement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnnouncementCompartment Gets the compartment details of an announcement.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/GetAnnouncementCompartment.go.html to see an example of how to use GetAnnouncementCompartment API.
func (client AnnouncementClient) GetAnnouncementCompartment(ctx context.Context, request GetAnnouncementCompartmentRequest) (response GetAnnouncementCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnnouncementCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnnouncementCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnnouncementCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnnouncementCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnnouncementCompartmentResponse")
	}
	return
}

// getAnnouncementCompartment implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementClient) getAnnouncementCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcements/{announcementId}/compartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnnouncementCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementCompartment/GetAnnouncementCompartment"
		err = common.PostProcessServiceError(err, "Announcement", "GetAnnouncementCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnnouncementUserStatus Gets information about whether a specific announcement was acknowledged by a user.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/GetAnnouncementUserStatus.go.html to see an example of how to use GetAnnouncementUserStatus API.
func (client AnnouncementClient) GetAnnouncementUserStatus(ctx context.Context, request GetAnnouncementUserStatusRequest) (response GetAnnouncementUserStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnnouncementUserStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnnouncementUserStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnnouncementUserStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnnouncementUserStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnnouncementUserStatusResponse")
	}
	return
}

// getAnnouncementUserStatus implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementClient) getAnnouncementUserStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcements/{announcementId}/userStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnnouncementUserStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementUserStatusDetails/GetAnnouncementUserStatus"
		err = common.PostProcessServiceError(err, "Announcement", "GetAnnouncementUserStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAnnouncements Gets a list of announcements for the current tenancy.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListAnnouncements.go.html to see an example of how to use ListAnnouncements API.
func (client AnnouncementClient) ListAnnouncements(ctx context.Context, request ListAnnouncementsRequest) (response ListAnnouncementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnnouncements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnnouncementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnnouncementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnnouncementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnnouncementsResponse")
	}
	return
}

// listAnnouncements implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementClient) listAnnouncements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAnnouncementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementsCollection/ListAnnouncements"
		err = common.PostProcessServiceError(err, "Announcement", "ListAnnouncements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAnnouncementUserStatus Updates the status of the specified announcement with regard to whether it has been marked as read.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/UpdateAnnouncementUserStatus.go.html to see an example of how to use UpdateAnnouncementUserStatus API.
func (client AnnouncementClient) UpdateAnnouncementUserStatus(ctx context.Context, request UpdateAnnouncementUserStatusRequest) (response UpdateAnnouncementUserStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnnouncementUserStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnnouncementUserStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnnouncementUserStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnnouncementUserStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnnouncementUserStatusResponse")
	}
	return
}

// updateAnnouncementUserStatus implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementClient) updateAnnouncementUserStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/announcements/{announcementId}/userStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAnnouncementUserStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementUserStatusDetails/UpdateAnnouncementUserStatus"
		err = common.PostProcessServiceError(err, "Announcement", "UpdateAnnouncementUserStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
