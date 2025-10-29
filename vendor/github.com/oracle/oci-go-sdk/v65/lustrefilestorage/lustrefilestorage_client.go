// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LustreFileStorageClient a client for LustreFileStorage
type LustreFileStorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLustreFileStorageClientWithConfigurationProvider Creates a new default LustreFileStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLustreFileStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LustreFileStorageClient, err error) {
	if enabled := common.CheckForEnabledServices("lustrefilestorage"); !enabled {
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
	return newLustreFileStorageClientFromBaseClient(baseClient, provider)
}

// NewLustreFileStorageClientWithOboToken Creates a new default LustreFileStorage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLustreFileStorageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LustreFileStorageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLustreFileStorageClientFromBaseClient(baseClient, configProvider)
}

func newLustreFileStorageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LustreFileStorageClient, err error) {
	// LustreFileStorage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LustreFileStorage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LustreFileStorageClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LustreFileStorageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LustreFileStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LustreFileStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client LustreFileStorageClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLustreFileSystemCompartment Moves a Lustre file system into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ChangeLustreFileSystemCompartment.go.html to see an example of how to use ChangeLustreFileSystemCompartment API.
// A default retry strategy applies to this operation ChangeLustreFileSystemCompartment()
func (client LustreFileStorageClient) ChangeLustreFileSystemCompartment(ctx context.Context, request ChangeLustreFileSystemCompartmentRequest) (response ChangeLustreFileSystemCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeLustreFileSystemCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLustreFileSystemCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLustreFileSystemCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLustreFileSystemCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLustreFileSystemCompartmentResponse")
	}
	return
}

// changeLustreFileSystemCompartment implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) changeLustreFileSystemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lustreFileSystems/{lustreFileSystemId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeLustreFileSystemCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LustreFileSystem/ChangeLustreFileSystemCompartment"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ChangeLustreFileSystemCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeObjectStorageLinkCompartment Moves an Object Storage link into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ChangeObjectStorageLinkCompartment.go.html to see an example of how to use ChangeObjectStorageLinkCompartment API.
// A default retry strategy applies to this operation ChangeObjectStorageLinkCompartment()
func (client LustreFileStorageClient) ChangeObjectStorageLinkCompartment(ctx context.Context, request ChangeObjectStorageLinkCompartmentRequest) (response ChangeObjectStorageLinkCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeObjectStorageLinkCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeObjectStorageLinkCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeObjectStorageLinkCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeObjectStorageLinkCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeObjectStorageLinkCompartmentResponse")
	}
	return
}

// changeObjectStorageLinkCompartment implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) changeObjectStorageLinkCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/objectStorageLinks/{objectStorageLinkId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeObjectStorageLinkCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/ChangeObjectStorageLinkCompartment"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ChangeObjectStorageLinkCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLustreFileSystem Creates a Lustre file system.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/CreateLustreFileSystem.go.html to see an example of how to use CreateLustreFileSystem API.
// A default retry strategy applies to this operation CreateLustreFileSystem()
func (client LustreFileStorageClient) CreateLustreFileSystem(ctx context.Context, request CreateLustreFileSystemRequest) (response CreateLustreFileSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createLustreFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLustreFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLustreFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLustreFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLustreFileSystemResponse")
	}
	return
}

// createLustreFileSystem implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) createLustreFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lustreFileSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLustreFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LustreFileSystem/CreateLustreFileSystem"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "CreateLustreFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateObjectStorageLink Creates an Object Storage link.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/CreateObjectStorageLink.go.html to see an example of how to use CreateObjectStorageLink API.
// A default retry strategy applies to this operation CreateObjectStorageLink()
func (client LustreFileStorageClient) CreateObjectStorageLink(ctx context.Context, request CreateObjectStorageLinkRequest) (response CreateObjectStorageLinkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createObjectStorageLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateObjectStorageLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateObjectStorageLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateObjectStorageLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateObjectStorageLinkResponse")
	}
	return
}

// createObjectStorageLink implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) createObjectStorageLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/objectStorageLinks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateObjectStorageLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/CreateObjectStorageLink"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "CreateObjectStorageLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLustreFileSystem Deletes a Lustre file system.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/DeleteLustreFileSystem.go.html to see an example of how to use DeleteLustreFileSystem API.
// A default retry strategy applies to this operation DeleteLustreFileSystem()
func (client LustreFileStorageClient) DeleteLustreFileSystem(ctx context.Context, request DeleteLustreFileSystemRequest) (response DeleteLustreFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLustreFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLustreFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLustreFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLustreFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLustreFileSystemResponse")
	}
	return
}

// deleteLustreFileSystem implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) deleteLustreFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/lustreFileSystems/{lustreFileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLustreFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LustreFileSystem/DeleteLustreFileSystem"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "DeleteLustreFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteObjectStorageLink Deletes an Object Storage link.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/DeleteObjectStorageLink.go.html to see an example of how to use DeleteObjectStorageLink API.
// A default retry strategy applies to this operation DeleteObjectStorageLink()
func (client LustreFileStorageClient) DeleteObjectStorageLink(ctx context.Context, request DeleteObjectStorageLinkRequest) (response DeleteObjectStorageLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteObjectStorageLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteObjectStorageLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteObjectStorageLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteObjectStorageLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteObjectStorageLinkResponse")
	}
	return
}

// deleteObjectStorageLink implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) deleteObjectStorageLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/objectStorageLinks/{objectStorageLinkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteObjectStorageLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/DeleteObjectStorageLink"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "DeleteObjectStorageLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLustreFileSystem Gets information about a Lustre file system.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/GetLustreFileSystem.go.html to see an example of how to use GetLustreFileSystem API.
// A default retry strategy applies to this operation GetLustreFileSystem()
func (client LustreFileStorageClient) GetLustreFileSystem(ctx context.Context, request GetLustreFileSystemRequest) (response GetLustreFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLustreFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLustreFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLustreFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLustreFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLustreFileSystemResponse")
	}
	return
}

// getLustreFileSystem implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) getLustreFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lustreFileSystems/{lustreFileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLustreFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LustreFileSystem/GetLustreFileSystem"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "GetLustreFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObjectStorageLink Gets information about an Object Storage link.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/GetObjectStorageLink.go.html to see an example of how to use GetObjectStorageLink API.
// A default retry strategy applies to this operation GetObjectStorageLink()
func (client LustreFileStorageClient) GetObjectStorageLink(ctx context.Context, request GetObjectStorageLinkRequest) (response GetObjectStorageLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObjectStorageLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectStorageLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectStorageLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectStorageLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectStorageLinkResponse")
	}
	return
}

// getObjectStorageLink implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) getObjectStorageLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/objectStorageLinks/{objectStorageLinkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetObjectStorageLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/GetObjectStorageLink"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "GetObjectStorageLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSyncJob Gets details of a sync job associated with an Object Storage link when `objectStorageLink` and a unique ID are provided.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/GetSyncJob.go.html to see an example of how to use GetSyncJob API.
// A default retry strategy applies to this operation GetSyncJob()
func (client LustreFileStorageClient) GetSyncJob(ctx context.Context, request GetSyncJobRequest) (response GetSyncJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSyncJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSyncJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSyncJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSyncJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSyncJobResponse")
	}
	return
}

// getSyncJob implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) getSyncJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/objectStorageLinks/{objectStorageLinkId}/syncJobs/{syncJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSyncJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/GetSyncJob"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "GetSyncJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client LustreFileStorageClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLustreFileSystems Gets a list of Lustre file systems.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListLustreFileSystems.go.html to see an example of how to use ListLustreFileSystems API.
// A default retry strategy applies to this operation ListLustreFileSystems()
func (client LustreFileStorageClient) ListLustreFileSystems(ctx context.Context, request ListLustreFileSystemsRequest) (response ListLustreFileSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLustreFileSystems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLustreFileSystemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLustreFileSystemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLustreFileSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLustreFileSystemsResponse")
	}
	return
}

// listLustreFileSystems implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) listLustreFileSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lustreFileSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLustreFileSystemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LustreFileSystemCollection/ListLustreFileSystems"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ListLustreFileSystems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListObjectStorageLinks Gets a list of Object Storage links.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListObjectStorageLinks.go.html to see an example of how to use ListObjectStorageLinks API.
// A default retry strategy applies to this operation ListObjectStorageLinks()
func (client LustreFileStorageClient) ListObjectStorageLinks(ctx context.Context, request ListObjectStorageLinksRequest) (response ListObjectStorageLinksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listObjectStorageLinks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListObjectStorageLinksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListObjectStorageLinksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListObjectStorageLinksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListObjectStorageLinksResponse")
	}
	return
}

// listObjectStorageLinks implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) listObjectStorageLinks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/objectStorageLinks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListObjectStorageLinksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLinkCollection/ListObjectStorageLinks"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ListObjectStorageLinks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSyncJobs Lists all sync jobs associated with the Object Storage link. Contains a unique ID for each sync job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListSyncJobs.go.html to see an example of how to use ListSyncJobs API.
// A default retry strategy applies to this operation ListSyncJobs()
func (client LustreFileStorageClient) ListSyncJobs(ctx context.Context, request ListSyncJobsRequest) (response ListSyncJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSyncJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSyncJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSyncJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSyncJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSyncJobsResponse")
	}
	return
}

// listSyncJobs implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) listSyncJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/objectStorageLinks/{objectStorageLinkId}/syncJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSyncJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/ListSyncJobs"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ListSyncJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client LustreFileStorageClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client LustreFileStorageClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client LustreFileStorageClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartExportToObject Starts the export of data from the Lustre file system to Object Storage.
// The Lustre file system path and Object Storage object prefix are defined in the Object Storage link resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/StartExportToObject.go.html to see an example of how to use StartExportToObject API.
// A default retry strategy applies to this operation StartExportToObject()
func (client LustreFileStorageClient) StartExportToObject(ctx context.Context, request StartExportToObjectRequest) (response StartExportToObjectResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startExportToObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartExportToObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartExportToObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartExportToObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartExportToObjectResponse")
	}
	return
}

// startExportToObject implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) startExportToObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/objectStorageLinks/{objectStorageLinkId}/actions/startExportToObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartExportToObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/StartExportToObject"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "StartExportToObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartImportFromObject Starts the import of data from Object Storage to the Lustre file system.
// The Lustre file system path and Object Storage object prefix are defined in the Object Storage link resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/StartImportFromObject.go.html to see an example of how to use StartImportFromObject API.
// A default retry strategy applies to this operation StartImportFromObject()
func (client LustreFileStorageClient) StartImportFromObject(ctx context.Context, request StartImportFromObjectRequest) (response StartImportFromObjectResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startImportFromObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartImportFromObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartImportFromObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartImportFromObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartImportFromObjectResponse")
	}
	return
}

// startImportFromObject implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) startImportFromObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/objectStorageLinks/{objectStorageLinkId}/actions/startImportFromObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartImportFromObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/StartImportFromObject"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "StartImportFromObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopExportToObject Stops the export of data from the Lustre file system to Object Storage.
// The Lustre file system path and Object Storage object prefix are defined in the Object Storage link resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/StopExportToObject.go.html to see an example of how to use StopExportToObject API.
// A default retry strategy applies to this operation StopExportToObject()
func (client LustreFileStorageClient) StopExportToObject(ctx context.Context, request StopExportToObjectRequest) (response StopExportToObjectResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopExportToObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopExportToObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopExportToObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopExportToObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopExportToObjectResponse")
	}
	return
}

// stopExportToObject implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) stopExportToObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/objectStorageLinks/{objectStorageLinkId}/actions/stopExportToObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopExportToObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/StopExportToObject"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "StopExportToObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopImportFromObject Stops the import of data from Object Storage to the Lustre file system.
// The Lustre file system path and Object Storage object prefix are defined in the Object Storage link resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/StopImportFromObject.go.html to see an example of how to use StopImportFromObject API.
// A default retry strategy applies to this operation StopImportFromObject()
func (client LustreFileStorageClient) StopImportFromObject(ctx context.Context, request StopImportFromObjectRequest) (response StopImportFromObjectResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopImportFromObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopImportFromObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopImportFromObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopImportFromObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopImportFromObjectResponse")
	}
	return
}

// stopImportFromObject implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) stopImportFromObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/objectStorageLinks/{objectStorageLinkId}/actions/stopImportFromObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopImportFromObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/StopImportFromObject"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "StopImportFromObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLustreFileSystem Updates a Lustre file system.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/UpdateLustreFileSystem.go.html to see an example of how to use UpdateLustreFileSystem API.
// A default retry strategy applies to this operation UpdateLustreFileSystem()
func (client LustreFileStorageClient) UpdateLustreFileSystem(ctx context.Context, request UpdateLustreFileSystemRequest) (response UpdateLustreFileSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLustreFileSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLustreFileSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLustreFileSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLustreFileSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLustreFileSystemResponse")
	}
	return
}

// updateLustreFileSystem implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) updateLustreFileSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/lustreFileSystems/{lustreFileSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLustreFileSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LustreFileSystem/UpdateLustreFileSystem"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "UpdateLustreFileSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateObjectStorageLink Updates an Object Storage link.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/UpdateObjectStorageLink.go.html to see an example of how to use UpdateObjectStorageLink API.
// A default retry strategy applies to this operation UpdateObjectStorageLink()
func (client LustreFileStorageClient) UpdateObjectStorageLink(ctx context.Context, request UpdateObjectStorageLinkRequest) (response UpdateObjectStorageLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateObjectStorageLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateObjectStorageLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateObjectStorageLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateObjectStorageLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateObjectStorageLinkResponse")
	}
	return
}

// updateObjectStorageLink implements the OCIOperation interface (enables retrying operations)
func (client LustreFileStorageClient) updateObjectStorageLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/objectStorageLinks/{objectStorageLinkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateObjectStorageLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ObjectStorageLink/UpdateObjectStorageLink"
		err = common.PostProcessServiceError(err, "LustreFileStorage", "UpdateObjectStorageLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
