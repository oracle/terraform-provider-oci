// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SoftwareSourceClient a client for SoftwareSource
type SoftwareSourceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSoftwareSourceClientWithConfigurationProvider Creates a new default SoftwareSource client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSoftwareSourceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SoftwareSourceClient, err error) {
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
	return newSoftwareSourceClientFromBaseClient(baseClient, provider)
}

// NewSoftwareSourceClientWithOboToken Creates a new default SoftwareSource client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSoftwareSourceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SoftwareSourceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSoftwareSourceClientFromBaseClient(baseClient, configProvider)
}

func newSoftwareSourceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SoftwareSourceClient, err error) {
	// SoftwareSource service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("SoftwareSource"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SoftwareSourceClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SoftwareSourceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SoftwareSourceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SoftwareSourceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAvailabilityOfSoftwareSources Updates the availability for a list of specified software sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ChangeAvailabilityOfSoftwareSources.go.html to see an example of how to use ChangeAvailabilityOfSoftwareSources API.
// A default retry strategy applies to this operation ChangeAvailabilityOfSoftwareSources()
func (client SoftwareSourceClient) ChangeAvailabilityOfSoftwareSources(ctx context.Context, request ChangeAvailabilityOfSoftwareSourcesRequest) (response ChangeAvailabilityOfSoftwareSourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAvailabilityOfSoftwareSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAvailabilityOfSoftwareSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAvailabilityOfSoftwareSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAvailabilityOfSoftwareSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAvailabilityOfSoftwareSourcesResponse")
	}
	return
}

// changeAvailabilityOfSoftwareSources implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) changeAvailabilityOfSoftwareSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/actions/changeAvailability", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAvailabilityOfSoftwareSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ChangeAvailabilityOfSoftwareSources"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ChangeAvailabilityOfSoftwareSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEntitlement Registers the necessary entitlement credentials for OS vendor software sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateEntitlement.go.html to see an example of how to use CreateEntitlement API.
// A default retry strategy applies to this operation CreateEntitlement()
func (client SoftwareSourceClient) CreateEntitlement(ctx context.Context, request CreateEntitlementRequest) (response CreateEntitlementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntitlementResponse")
	}
	return
}

// createEntitlement implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) createEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/entitlements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEntitlementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/CreateEntitlement"
		err = common.PostProcessServiceError(err, "SoftwareSource", "CreateEntitlement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSoftwareSource Creates a new versioned or custom software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateSoftwareSource.go.html to see an example of how to use CreateSoftwareSource API.
// A default retry strategy applies to this operation CreateSoftwareSource()
func (client SoftwareSourceClient) CreateSoftwareSource(ctx context.Context, request CreateSoftwareSourceRequest) (response CreateSoftwareSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSoftwareSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSoftwareSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSoftwareSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSoftwareSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSoftwareSourceResponse")
	}
	return
}

// createSoftwareSource implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) createSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/CreateSoftwareSource"
		err = common.PostProcessServiceError(err, "SoftwareSource", "CreateSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &softwaresource{})
	return response, err
}

// DeleteSoftwareSource Deletes the specified software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteSoftwareSource.go.html to see an example of how to use DeleteSoftwareSource API.
// A default retry strategy applies to this operation DeleteSoftwareSource()
func (client SoftwareSourceClient) DeleteSoftwareSource(ctx context.Context, request DeleteSoftwareSourceRequest) (response DeleteSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSoftwareSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSoftwareSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSoftwareSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSoftwareSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSoftwareSourceResponse")
	}
	return
}

// deleteSoftwareSource implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) deleteSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/softwareSources/{softwareSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/DeleteSoftwareSource"
		err = common.PostProcessServiceError(err, "SoftwareSource", "DeleteSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetErratum Gets information about the specified erratum by its advisory name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetErratum.go.html to see an example of how to use GetErratum API.
// A default retry strategy applies to this operation GetErratum()
func (client SoftwareSourceClient) GetErratum(ctx context.Context, request GetErratumRequest) (response GetErratumResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getErratum, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetErratumResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetErratumResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetErratumResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetErratumResponse")
	}
	return
}

// getErratum implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) getErratum(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/errata/{name}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetErratumResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Erratum/GetErratum"
		err = common.PostProcessServiceError(err, "SoftwareSource", "GetErratum", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModuleStream Gets information about the specified module stream in a software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetModuleStream.go.html to see an example of how to use GetModuleStream API.
// A default retry strategy applies to this operation GetModuleStream()
func (client SoftwareSourceClient) GetModuleStream(ctx context.Context, request GetModuleStreamRequest) (response GetModuleStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModuleStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModuleStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModuleStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModuleStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModuleStreamResponse")
	}
	return
}

// getModuleStream implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) getModuleStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/moduleStreams/{moduleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModuleStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ModuleStream/GetModuleStream"
		err = common.PostProcessServiceError(err, "SoftwareSource", "GetModuleStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModuleStreamProfile Gets information about the specified module stream profile in a software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetModuleStreamProfile.go.html to see an example of how to use GetModuleStreamProfile API.
// A default retry strategy applies to this operation GetModuleStreamProfile()
func (client SoftwareSourceClient) GetModuleStreamProfile(ctx context.Context, request GetModuleStreamProfileRequest) (response GetModuleStreamProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModuleStreamProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModuleStreamProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModuleStreamProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModuleStreamProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModuleStreamProfileResponse")
	}
	return
}

// getModuleStreamProfile implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) getModuleStreamProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/moduleStreamProfiles/{profileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModuleStreamProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ModuleStreamProfile/GetModuleStreamProfile"
		err = common.PostProcessServiceError(err, "SoftwareSource", "GetModuleStreamProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPackageGroup Gets information about the specified package group from a software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetPackageGroup.go.html to see an example of how to use GetPackageGroup API.
// A default retry strategy applies to this operation GetPackageGroup()
func (client SoftwareSourceClient) GetPackageGroup(ctx context.Context, request GetPackageGroupRequest) (response GetPackageGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPackageGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPackageGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPackageGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPackageGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPackageGroupResponse")
	}
	return
}

// getPackageGroup implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) getPackageGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/packageGroups/{packageGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPackageGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/PackageGroup/GetPackageGroup"
		err = common.PostProcessServiceError(err, "SoftwareSource", "GetPackageGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwarePackage Gets information about the specified software package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetSoftwarePackage.go.html to see an example of how to use GetSoftwarePackage API.
// A default retry strategy applies to this operation GetSoftwarePackage()
func (client SoftwareSourceClient) GetSoftwarePackage(ctx context.Context, request GetSoftwarePackageRequest) (response GetSoftwarePackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSoftwarePackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSoftwarePackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSoftwarePackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSoftwarePackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSoftwarePackageResponse")
	}
	return
}

// getSoftwarePackage implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) getSoftwarePackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/softwarePackages/{softwarePackageName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSoftwarePackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/GetSoftwarePackage"
		err = common.PostProcessServiceError(err, "SoftwareSource", "GetSoftwarePackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwareSource Gets information about the specified software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetSoftwareSource.go.html to see an example of how to use GetSoftwareSource API.
// A default retry strategy applies to this operation GetSoftwareSource()
func (client SoftwareSourceClient) GetSoftwareSource(ctx context.Context, request GetSoftwareSourceRequest) (response GetSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSoftwareSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSoftwareSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSoftwareSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSoftwareSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSoftwareSourceResponse")
	}
	return
}

// getSoftwareSource implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) getSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/GetSoftwareSource"
		err = common.PostProcessServiceError(err, "SoftwareSource", "GetSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &softwaresource{})
	return response, err
}

// ListEntitlements Lists entitlements in the specified tenancy OCID. Filter the list against a variety of criteria including but
// not limited to its CSI, and vendor name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListEntitlements.go.html to see an example of how to use ListEntitlements API.
// A default retry strategy applies to this operation ListEntitlements()
func (client SoftwareSourceClient) ListEntitlements(ctx context.Context, request ListEntitlementsRequest) (response ListEntitlementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntitlements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntitlementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntitlementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntitlementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntitlementsResponse")
	}
	return
}

// listEntitlements implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listEntitlements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/entitlements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEntitlementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListEntitlements"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListEntitlements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListErrata Lists all of the currently available errata. Filter the list against a variety of criteria including but not
// limited to its name, classification type, advisory severity, and OS family.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListErrata.go.html to see an example of how to use ListErrata API.
// A default retry strategy applies to this operation ListErrata()
func (client SoftwareSourceClient) ListErrata(ctx context.Context, request ListErrataRequest) (response ListErrataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listErrata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListErrataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListErrataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListErrataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListErrataResponse")
	}
	return
}

// listErrata implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listErrata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/errata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListErrataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Erratum/ListErrata"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListErrata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModuleStreamProfiles Lists module stream profiles from the specified software source OCID. Filter the list against a variety of
// criteria including but not limited to its module name, stream name, and (profile) name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListModuleStreamProfiles.go.html to see an example of how to use ListModuleStreamProfiles API.
// A default retry strategy applies to this operation ListModuleStreamProfiles()
func (client SoftwareSourceClient) ListModuleStreamProfiles(ctx context.Context, request ListModuleStreamProfilesRequest) (response ListModuleStreamProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModuleStreamProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModuleStreamProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModuleStreamProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModuleStreamProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModuleStreamProfilesResponse")
	}
	return
}

// listModuleStreamProfiles implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listModuleStreamProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/moduleStreamProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModuleStreamProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListModuleStreamProfiles"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListModuleStreamProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModuleStreams Lists module streams from the specified software source OCID. Filter the list against a variety of
// criteria including but not limited to its module name and (stream) name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListModuleStreams.go.html to see an example of how to use ListModuleStreams API.
// A default retry strategy applies to this operation ListModuleStreams()
func (client SoftwareSourceClient) ListModuleStreams(ctx context.Context, request ListModuleStreamsRequest) (response ListModuleStreamsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModuleStreams, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModuleStreamsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModuleStreamsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModuleStreamsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModuleStreamsResponse")
	}
	return
}

// listModuleStreams implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listModuleStreams(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/moduleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModuleStreamsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListModuleStreams"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListModuleStreams", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPackageGroups Lists package groups that associate with the specified software source OCID. Filter the list against a
// variety of criteria including but not limited to its name, and package group type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListPackageGroups.go.html to see an example of how to use ListPackageGroups API.
// A default retry strategy applies to this operation ListPackageGroups()
func (client SoftwareSourceClient) ListPackageGroups(ctx context.Context, request ListPackageGroupsRequest) (response ListPackageGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPackageGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPackageGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPackageGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPackageGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPackageGroupsResponse")
	}
	return
}

// listPackageGroups implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listPackageGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/packageGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPackageGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListPackageGroups"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListPackageGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwarePackages Lists software packages in the specified software source.  Filter the list against a variety of criteria
// including but not limited to its name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwarePackages.go.html to see an example of how to use ListSoftwarePackages API.
// A default retry strategy applies to this operation ListSoftwarePackages()
func (client SoftwareSourceClient) ListSoftwarePackages(ctx context.Context, request ListSoftwarePackagesRequest) (response ListSoftwarePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSoftwarePackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSoftwarePackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSoftwarePackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSoftwarePackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSoftwarePackagesResponse")
	}
	return
}

// listSoftwarePackages implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listSoftwarePackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/softwarePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSoftwarePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListSoftwarePackages"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListSoftwarePackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareSourceVendors Lists available software source vendors. Filter the list against a variety of criteria including but not limited
// to its name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwareSourceVendors.go.html to see an example of how to use ListSoftwareSourceVendors API.
// A default retry strategy applies to this operation ListSoftwareSourceVendors()
func (client SoftwareSourceClient) ListSoftwareSourceVendors(ctx context.Context, request ListSoftwareSourceVendorsRequest) (response ListSoftwareSourceVendorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSoftwareSourceVendors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSoftwareSourceVendorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSoftwareSourceVendorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSoftwareSourceVendorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSoftwareSourceVendorsResponse")
	}
	return
}

// listSoftwareSourceVendors implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listSoftwareSourceVendors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSourceVendors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSoftwareSourceVendorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListSoftwareSourceVendors"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListSoftwareSourceVendors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareSources Lists software sources that match the specified tenancy or software source OCID. Filter the list against a
// variety of criteria including but not limited to its name, status, architecture, and OS family.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwareSources.go.html to see an example of how to use ListSoftwareSources API.
// A default retry strategy applies to this operation ListSoftwareSources()
func (client SoftwareSourceClient) ListSoftwareSources(ctx context.Context, request ListSoftwareSourcesRequest) (response ListSoftwareSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSoftwareSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSoftwareSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSoftwareSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSoftwareSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSoftwareSourcesResponse")
	}
	return
}

// listSoftwareSources implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) listSoftwareSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSoftwareSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/ListSoftwareSources"
		err = common.PostProcessServiceError(err, "SoftwareSource", "ListSoftwareSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSoftwareSourceModuleStreams Lists modules from a list of software sources. Filter the list against a variety of
// criteria including the module name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SearchSoftwareSourceModuleStreams.go.html to see an example of how to use SearchSoftwareSourceModuleStreams API.
// A default retry strategy applies to this operation SearchSoftwareSourceModuleStreams()
func (client SoftwareSourceClient) SearchSoftwareSourceModuleStreams(ctx context.Context, request SearchSoftwareSourceModuleStreamsRequest) (response SearchSoftwareSourceModuleStreamsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchSoftwareSourceModuleStreams, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSoftwareSourceModuleStreamsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSoftwareSourceModuleStreamsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSoftwareSourceModuleStreamsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSoftwareSourceModuleStreamsResponse")
	}
	return
}

// searchSoftwareSourceModuleStreams implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) searchSoftwareSourceModuleStreams(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSourceModuleStreams/actions/search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSoftwareSourceModuleStreamsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/SearchSoftwareSourceModuleStreams"
		err = common.PostProcessServiceError(err, "SoftwareSource", "SearchSoftwareSourceModuleStreams", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSoftwareSourceModules Lists modules from a list of software sources. Filter the list against a variety of
// criteria including the (module) name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SearchSoftwareSourceModules.go.html to see an example of how to use SearchSoftwareSourceModules API.
// A default retry strategy applies to this operation SearchSoftwareSourceModules()
func (client SoftwareSourceClient) SearchSoftwareSourceModules(ctx context.Context, request SearchSoftwareSourceModulesRequest) (response SearchSoftwareSourceModulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchSoftwareSourceModules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSoftwareSourceModulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSoftwareSourceModulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSoftwareSourceModulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSoftwareSourceModulesResponse")
	}
	return
}

// searchSoftwareSourceModules implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) searchSoftwareSourceModules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSourceModules/actions/search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSoftwareSourceModulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/SearchSoftwareSourceModules"
		err = common.PostProcessServiceError(err, "SoftwareSource", "SearchSoftwareSourceModules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSoftwareSourcePackageGroups Searches the package groups from the specified list of software sources. Filter the list against a variety of criteria
// including but not limited to its name, and group type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SearchSoftwareSourcePackageGroups.go.html to see an example of how to use SearchSoftwareSourcePackageGroups API.
// A default retry strategy applies to this operation SearchSoftwareSourcePackageGroups()
func (client SoftwareSourceClient) SearchSoftwareSourcePackageGroups(ctx context.Context, request SearchSoftwareSourcePackageGroupsRequest) (response SearchSoftwareSourcePackageGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchSoftwareSourcePackageGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSoftwareSourcePackageGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSoftwareSourcePackageGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSoftwareSourcePackageGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSoftwareSourcePackageGroupsResponse")
	}
	return
}

// searchSoftwareSourcePackageGroups implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) searchSoftwareSourcePackageGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSourcePackageGroups/actions/search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSoftwareSourcePackageGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/SearchSoftwareSourcePackageGroups"
		err = common.PostProcessServiceError(err, "SoftwareSource", "SearchSoftwareSourcePackageGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSoftwareSource Updates the specified software source's details, including but not limited to name, description, and tags.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateSoftwareSource.go.html to see an example of how to use UpdateSoftwareSource API.
// A default retry strategy applies to this operation UpdateSoftwareSource()
func (client SoftwareSourceClient) UpdateSoftwareSource(ctx context.Context, request UpdateSoftwareSourceRequest) (response UpdateSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSoftwareSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSoftwareSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSoftwareSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSoftwareSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSoftwareSourceResponse")
	}
	return
}

// updateSoftwareSource implements the OCIOperation interface (enables retrying operations)
func (client SoftwareSourceClient) updateSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/softwareSources/{softwareSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/SoftwareSource/UpdateSoftwareSource"
		err = common.PostProcessServiceError(err, "SoftwareSource", "UpdateSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &softwaresource{})
	return response, err
}
