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

// AnnouncementsPreferencesClient a client for AnnouncementsPreferences
type AnnouncementsPreferencesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAnnouncementsPreferencesClientWithConfigurationProvider Creates a new default AnnouncementsPreferences client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAnnouncementsPreferencesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AnnouncementsPreferencesClient, err error) {
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
	return newAnnouncementsPreferencesClientFromBaseClient(baseClient, provider)
}

// NewAnnouncementsPreferencesClientWithOboToken Creates a new default AnnouncementsPreferences client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAnnouncementsPreferencesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AnnouncementsPreferencesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAnnouncementsPreferencesClientFromBaseClient(baseClient, configProvider)
}

func newAnnouncementsPreferencesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AnnouncementsPreferencesClient, err error) {
	// AnnouncementsPreferences service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AnnouncementsPreferences"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AnnouncementsPreferencesClient{BaseClient: baseClient}
	client.BasePath = "20180904"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AnnouncementsPreferencesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("announcements", "https://announcements.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AnnouncementsPreferencesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AnnouncementsPreferencesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAnnouncementsPreference Creates a request that specifies preferences for the tenancy regarding receiving announcements by email.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/CreateAnnouncementsPreference.go.html to see an example of how to use CreateAnnouncementsPreference API.
func (client AnnouncementsPreferencesClient) CreateAnnouncementsPreference(ctx context.Context, request CreateAnnouncementsPreferenceRequest) (response CreateAnnouncementsPreferenceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAnnouncementsPreference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAnnouncementsPreferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAnnouncementsPreferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAnnouncementsPreferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAnnouncementsPreferenceResponse")
	}
	return
}

// createAnnouncementsPreference implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementsPreferencesClient) createAnnouncementsPreference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/announcementsPreferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAnnouncementsPreferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementsPreferencesSummary/CreateAnnouncementsPreference"
		err = common.PostProcessServiceError(err, "AnnouncementsPreferences", "CreateAnnouncementsPreference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnnouncementsPreference Gets the current preferences of the tenancy regarding receiving announcements by email.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/GetAnnouncementsPreference.go.html to see an example of how to use GetAnnouncementsPreference API.
func (client AnnouncementsPreferencesClient) GetAnnouncementsPreference(ctx context.Context, request GetAnnouncementsPreferenceRequest) (response GetAnnouncementsPreferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnnouncementsPreference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnnouncementsPreferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnnouncementsPreferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnnouncementsPreferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnnouncementsPreferenceResponse")
	}
	return
}

// getAnnouncementsPreference implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementsPreferencesClient) getAnnouncementsPreference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcementsPreferences/{preferenceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnnouncementsPreferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementsPreferences/GetAnnouncementsPreference"
		err = common.PostProcessServiceError(err, "AnnouncementsPreferences", "GetAnnouncementsPreference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAnnouncementsPreferences Gets the current preferences of the tenancy regarding receiving announcements by email.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListAnnouncementsPreferences.go.html to see an example of how to use ListAnnouncementsPreferences API.
func (client AnnouncementsPreferencesClient) ListAnnouncementsPreferences(ctx context.Context, request ListAnnouncementsPreferencesRequest) (response ListAnnouncementsPreferencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnnouncementsPreferences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnnouncementsPreferencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnnouncementsPreferencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnnouncementsPreferencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnnouncementsPreferencesResponse")
	}
	return
}

// listAnnouncementsPreferences implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementsPreferencesClient) listAnnouncementsPreferences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcementsPreferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAnnouncementsPreferencesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementsPreferencesSummary/ListAnnouncementsPreferences"
		err = common.PostProcessServiceError(err, "AnnouncementsPreferences", "ListAnnouncementsPreferences", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAnnouncementsPreference Updates the preferences of the tenancy regarding receiving announcements by email.
// This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/UpdateAnnouncementsPreference.go.html to see an example of how to use UpdateAnnouncementsPreference API.
func (client AnnouncementsPreferencesClient) UpdateAnnouncementsPreference(ctx context.Context, request UpdateAnnouncementsPreferenceRequest) (response UpdateAnnouncementsPreferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnnouncementsPreference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnnouncementsPreferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnnouncementsPreferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnnouncementsPreferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnnouncementsPreferenceResponse")
	}
	return
}

// updateAnnouncementsPreference implements the OCIOperation interface (enables retrying operations)
func (client AnnouncementsPreferencesClient) updateAnnouncementsPreference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/announcementsPreferences/{preferenceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAnnouncementsPreferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/announcements/0.0.1/AnnouncementsPreferencesSummary/UpdateAnnouncementsPreference"
		err = common.PostProcessServiceError(err, "AnnouncementsPreferences", "UpdateAnnouncementsPreference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
