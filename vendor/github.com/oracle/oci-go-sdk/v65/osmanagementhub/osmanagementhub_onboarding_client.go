// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OnboardingClient a client for Onboarding
type OnboardingClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOnboardingClientWithConfigurationProvider Creates a new default Onboarding client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOnboardingClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OnboardingClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagementhub"); !enabled {
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
	return newOnboardingClientFromBaseClient(baseClient, provider)
}

// NewOnboardingClientWithOboToken Creates a new default Onboarding client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOnboardingClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OnboardingClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOnboardingClientFromBaseClient(baseClient, configProvider)
}

func newOnboardingClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OnboardingClient, err error) {
	// Onboarding service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Onboarding"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OnboardingClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OnboardingClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OnboardingClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OnboardingClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachLifecycleStageToProfile Attaches the specified lifecycle stage to a profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachLifecycleStageToProfile.go.html to see an example of how to use AttachLifecycleStageToProfile API.
// A default retry strategy applies to this operation AttachLifecycleStageToProfile()
func (client OnboardingClient) AttachLifecycleStageToProfile(ctx context.Context, request AttachLifecycleStageToProfileRequest) (response AttachLifecycleStageToProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachLifecycleStageToProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachLifecycleStageToProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachLifecycleStageToProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachLifecycleStageToProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachLifecycleStageToProfileResponse")
	}
	return
}

// attachLifecycleStageToProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) attachLifecycleStageToProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles/{profileId}/actions/attachLifecycleStage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachLifecycleStageToProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/AttachLifecycleStageToProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "AttachLifecycleStageToProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachManagedInstanceGroupToProfile Attaches the specified managed instance group to a profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachManagedInstanceGroupToProfile.go.html to see an example of how to use AttachManagedInstanceGroupToProfile API.
// A default retry strategy applies to this operation AttachManagedInstanceGroupToProfile()
func (client OnboardingClient) AttachManagedInstanceGroupToProfile(ctx context.Context, request AttachManagedInstanceGroupToProfileRequest) (response AttachManagedInstanceGroupToProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachManagedInstanceGroupToProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachManagedInstanceGroupToProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachManagedInstanceGroupToProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachManagedInstanceGroupToProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachManagedInstanceGroupToProfileResponse")
	}
	return
}

// attachManagedInstanceGroupToProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) attachManagedInstanceGroupToProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles/{profileId}/actions/attachManagedInstanceGroup", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachManagedInstanceGroupToProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/AttachManagedInstanceGroupToProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "AttachManagedInstanceGroupToProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachManagementStationToProfile Attaches the specified management station to a profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachManagementStationToProfile.go.html to see an example of how to use AttachManagementStationToProfile API.
// A default retry strategy applies to this operation AttachManagementStationToProfile()
func (client OnboardingClient) AttachManagementStationToProfile(ctx context.Context, request AttachManagementStationToProfileRequest) (response AttachManagementStationToProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachManagementStationToProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachManagementStationToProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachManagementStationToProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachManagementStationToProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachManagementStationToProfileResponse")
	}
	return
}

// attachManagementStationToProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) attachManagementStationToProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles/{profileId}/actions/attachManagementStation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachManagementStationToProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/AttachManagementStationToProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "AttachManagementStationToProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachSoftwareSourcesToProfile Attaches the specified software sources to a profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachSoftwareSourcesToProfile.go.html to see an example of how to use AttachSoftwareSourcesToProfile API.
// A default retry strategy applies to this operation AttachSoftwareSourcesToProfile()
func (client OnboardingClient) AttachSoftwareSourcesToProfile(ctx context.Context, request AttachSoftwareSourcesToProfileRequest) (response AttachSoftwareSourcesToProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachSoftwareSourcesToProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachSoftwareSourcesToProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachSoftwareSourcesToProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachSoftwareSourcesToProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachSoftwareSourcesToProfileResponse")
	}
	return
}

// attachSoftwareSourcesToProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) attachSoftwareSourcesToProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles/{profileId}/actions/attachSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachSoftwareSourcesToProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/AttachSoftwareSourcesToProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "AttachSoftwareSourcesToProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeProfileCompartment Moves the profile to a different compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ChangeProfileCompartment.go.html to see an example of how to use ChangeProfileCompartment API.
// A default retry strategy applies to this operation ChangeProfileCompartment()
func (client OnboardingClient) ChangeProfileCompartment(ctx context.Context, request ChangeProfileCompartmentRequest) (response ChangeProfileCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeProfileCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeProfileCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeProfileCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeProfileCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeProfileCompartmentResponse")
	}
	return
}

// changeProfileCompartment implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) changeProfileCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles/{profileId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeProfileCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/ChangeProfileCompartment"
		err = common.PostProcessServiceError(err, "Onboarding", "ChangeProfileCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProfile Creates a registration profile. A profile defines the content applied to the instance when registering it with the service.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateProfile.go.html to see an example of how to use CreateProfile API.
// A default retry strategy applies to this operation CreateProfile()
func (client OnboardingClient) CreateProfile(ctx context.Context, request CreateProfileRequest) (response CreateProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProfileResponse")
	}
	return
}

// createProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) createProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/CreateProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "CreateProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &profile{})
	return response, err
}

// DeleteProfile Deletes a specified registration profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteProfile.go.html to see an example of how to use DeleteProfile API.
// A default retry strategy applies to this operation DeleteProfile()
func (client OnboardingClient) DeleteProfile(ctx context.Context, request DeleteProfileRequest) (response DeleteProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProfileResponse")
	}
	return
}

// deleteProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) deleteProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/profiles/{profileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/DeleteProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "DeleteProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachSoftwareSourcesFromProfile Detaches the specified software sources from a profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DetachSoftwareSourcesFromProfile.go.html to see an example of how to use DetachSoftwareSourcesFromProfile API.
// A default retry strategy applies to this operation DetachSoftwareSourcesFromProfile()
func (client OnboardingClient) DetachSoftwareSourcesFromProfile(ctx context.Context, request DetachSoftwareSourcesFromProfileRequest) (response DetachSoftwareSourcesFromProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachSoftwareSourcesFromProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachSoftwareSourcesFromProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachSoftwareSourcesFromProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachSoftwareSourcesFromProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachSoftwareSourcesFromProfileResponse")
	}
	return
}

// detachSoftwareSourcesFromProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) detachSoftwareSourcesFromProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles/{profileId}/actions/detachSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachSoftwareSourcesFromProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/DetachSoftwareSourcesFromProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "DetachSoftwareSourcesFromProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProfile Returns information about the specified registration profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetProfile.go.html to see an example of how to use GetProfile API.
// A default retry strategy applies to this operation GetProfile()
func (client OnboardingClient) GetProfile(ctx context.Context, request GetProfileRequest) (response GetProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProfileResponse")
	}
	return
}

// getProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) getProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profiles/{profileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/GetProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "GetProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &profile{})
	return response, err
}

// GetProfileVersion Returns information about the version of the specified registration profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetProfileVersion.go.html to see an example of how to use GetProfileVersion API.
// A default retry strategy applies to this operation GetProfileVersion()
func (client OnboardingClient) GetProfileVersion(ctx context.Context, request GetProfileVersionRequest) (response GetProfileVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProfileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProfileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProfileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProfileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProfileVersionResponse")
	}
	return
}

// getProfileVersion implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) getProfileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profiles/{profileId}/version/{profileVersion}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProfileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ProfileVersion/GetProfileVersion"
		err = common.PostProcessServiceError(err, "Onboarding", "GetProfileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfileAvailableSoftwareSources Lists available software sources for a specified profile. Filter the list against a variety of criteria including but not limited to the software source name. The results list only software sources that have not already been added to the profile.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListProfileAvailableSoftwareSources.go.html to see an example of how to use ListProfileAvailableSoftwareSources API.
// A default retry strategy applies to this operation ListProfileAvailableSoftwareSources()
func (client OnboardingClient) ListProfileAvailableSoftwareSources(ctx context.Context, request ListProfileAvailableSoftwareSourcesRequest) (response ListProfileAvailableSoftwareSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfileAvailableSoftwareSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfileAvailableSoftwareSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfileAvailableSoftwareSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfileAvailableSoftwareSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfileAvailableSoftwareSourcesResponse")
	}
	return
}

// listProfileAvailableSoftwareSources implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) listProfileAvailableSoftwareSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profiles/{profileId}/availableSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfileAvailableSoftwareSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/ListProfileAvailableSoftwareSources"
		err = common.PostProcessServiceError(err, "Onboarding", "ListProfileAvailableSoftwareSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfiles Lists registration profiles that match the specified compartment or profile OCID. Filter the list against a
// variety of criteria including but not limited to its name, status, vendor name, and architecture type.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListProfiles.go.html to see an example of how to use ListProfiles API.
// A default retry strategy applies to this operation ListProfiles()
func (client OnboardingClient) ListProfiles(ctx context.Context, request ListProfilesRequest) (response ListProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfilesResponse")
	}
	return
}

// listProfiles implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) listProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/ListProfiles"
		err = common.PostProcessServiceError(err, "Onboarding", "ListProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProfile Updates the specified profile's name, description, and tags.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateProfile.go.html to see an example of how to use UpdateProfile API.
// A default retry strategy applies to this operation UpdateProfile()
func (client OnboardingClient) UpdateProfile(ctx context.Context, request UpdateProfileRequest) (response UpdateProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProfileResponse")
	}
	return
}

// updateProfile implements the OCIOperation interface (enables retrying operations)
func (client OnboardingClient) updateProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/profiles/{profileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Profile/UpdateProfile"
		err = common.PostProcessServiceError(err, "Onboarding", "UpdateProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &profile{})
	return response, err
}
