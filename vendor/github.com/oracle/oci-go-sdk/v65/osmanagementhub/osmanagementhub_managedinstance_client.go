// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ManagedInstanceClient a client for ManagedInstance
type ManagedInstanceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagedInstanceClientWithConfigurationProvider Creates a new default ManagedInstance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagedInstanceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagedInstanceClient, err error) {
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
	return newManagedInstanceClientFromBaseClient(baseClient, provider)
}

// NewManagedInstanceClientWithOboToken Creates a new default ManagedInstance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewManagedInstanceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagedInstanceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newManagedInstanceClientFromBaseClient(baseClient, configProvider)
}

func newManagedInstanceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagedInstanceClient, err error) {
	// ManagedInstance service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ManagedInstance"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ManagedInstanceClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagedInstanceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagedInstanceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ManagedInstanceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachProfileToManagedInstance Adds profile to a managed instance. After the profile has been added,
// the instance can be registered as a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachProfileToManagedInstance.go.html to see an example of how to use AttachProfileToManagedInstance API.
// A default retry strategy applies to this operation AttachProfileToManagedInstance()
func (client ManagedInstanceClient) AttachProfileToManagedInstance(ctx context.Context, request AttachProfileToManagedInstanceRequest) (response AttachProfileToManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachProfileToManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachProfileToManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachProfileToManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachProfileToManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachProfileToManagedInstanceResponse")
	}
	return
}

// attachProfileToManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) attachProfileToManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/attachProfile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachProfileToManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/AttachProfileToManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "AttachProfileToManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachSoftwareSourcesToManagedInstance Adds software sources to a managed instance. After the software source has been added,
// then packages from that software source can be installed on the managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachSoftwareSourcesToManagedInstance.go.html to see an example of how to use AttachSoftwareSourcesToManagedInstance API.
// A default retry strategy applies to this operation AttachSoftwareSourcesToManagedInstance()
func (client ManagedInstanceClient) AttachSoftwareSourcesToManagedInstance(ctx context.Context, request AttachSoftwareSourcesToManagedInstanceRequest) (response AttachSoftwareSourcesToManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachSoftwareSourcesToManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachSoftwareSourcesToManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachSoftwareSourcesToManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachSoftwareSourcesToManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachSoftwareSourcesToManagedInstanceResponse")
	}
	return
}

// attachSoftwareSourcesToManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) attachSoftwareSourcesToManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/attachSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachSoftwareSourcesToManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/AttachSoftwareSourcesToManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "AttachSoftwareSourcesToManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedInstance Unregisters the specified managed instance from the service. Once unregistered, the service will no longer manage the instance.
// For Linux instances, yum or DNF repository files will be restored to their state prior to registration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteManagedInstance.go.html to see an example of how to use DeleteManagedInstance API.
// A default retry strategy applies to this operation DeleteManagedInstance()
func (client ManagedInstanceClient) DeleteManagedInstance(ctx context.Context, request DeleteManagedInstanceRequest) (response DeleteManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagedInstanceResponse")
	}
	return
}

// deleteManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) deleteManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedInstances/{managedInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/DeleteManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "DeleteManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachProfileFromManagedInstance Detaches profile from a managed instance. After the profile has been removed,
// the instance cannot be registered as a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DetachProfileFromManagedInstance.go.html to see an example of how to use DetachProfileFromManagedInstance API.
// A default retry strategy applies to this operation DetachProfileFromManagedInstance()
func (client ManagedInstanceClient) DetachProfileFromManagedInstance(ctx context.Context, request DetachProfileFromManagedInstanceRequest) (response DetachProfileFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachProfileFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachProfileFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachProfileFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachProfileFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachProfileFromManagedInstanceResponse")
	}
	return
}

// detachProfileFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) detachProfileFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/detachProfile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachProfileFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/DetachProfileFromManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "DetachProfileFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachSoftwareSourcesFromManagedInstance Removes software sources from a managed instance.
// Packages will no longer be able to be installed from these software sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DetachSoftwareSourcesFromManagedInstance.go.html to see an example of how to use DetachSoftwareSourcesFromManagedInstance API.
// A default retry strategy applies to this operation DetachSoftwareSourcesFromManagedInstance()
func (client ManagedInstanceClient) DetachSoftwareSourcesFromManagedInstance(ctx context.Context, request DetachSoftwareSourcesFromManagedInstanceRequest) (response DetachSoftwareSourcesFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachSoftwareSourcesFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachSoftwareSourcesFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachSoftwareSourcesFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachSoftwareSourcesFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachSoftwareSourcesFromManagedInstanceResponse")
	}
	return
}

// detachSoftwareSourcesFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) detachSoftwareSourcesFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/detachSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachSoftwareSourcesFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/DetachSoftwareSourcesFromManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "DetachSoftwareSourcesFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableModuleStreamOnManagedInstance Disables a module stream on a managed instance.  After the stream is
// disabled, it is no longer possible to install the profiles that are
// contained by the stream.  All installed profiles must be removed prior
// to disabling a module stream.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DisableModuleStreamOnManagedInstance.go.html to see an example of how to use DisableModuleStreamOnManagedInstance API.
// A default retry strategy applies to this operation DisableModuleStreamOnManagedInstance()
func (client ManagedInstanceClient) DisableModuleStreamOnManagedInstance(ctx context.Context, request DisableModuleStreamOnManagedInstanceRequest) (response DisableModuleStreamOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableModuleStreamOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableModuleStreamOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableModuleStreamOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableModuleStreamOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableModuleStreamOnManagedInstanceResponse")
	}
	return
}

// disableModuleStreamOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) disableModuleStreamOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/disableModuleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableModuleStreamOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/DisableModuleStreamOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "DisableModuleStreamOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableModuleStreamOnManagedInstance Enables a module stream on a managed instance.  After the stream is
// enabled, it is possible to install the profiles that are contained
// by the stream.  Enabling a stream that is already enabled will
// succeed.  Attempting to enable a different stream for a module that
// already has a stream enabled results in an error.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/EnableModuleStreamOnManagedInstance.go.html to see an example of how to use EnableModuleStreamOnManagedInstance API.
// A default retry strategy applies to this operation EnableModuleStreamOnManagedInstance()
func (client ManagedInstanceClient) EnableModuleStreamOnManagedInstance(ctx context.Context, request EnableModuleStreamOnManagedInstanceRequest) (response EnableModuleStreamOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableModuleStreamOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableModuleStreamOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableModuleStreamOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableModuleStreamOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableModuleStreamOnManagedInstanceResponse")
	}
	return
}

// enableModuleStreamOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) enableModuleStreamOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/enableModuleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableModuleStreamOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/EnableModuleStreamOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "EnableModuleStreamOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstance Gets information about the specified managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagedInstance.go.html to see an example of how to use GetManagedInstance API.
// A default retry strategy applies to this operation GetManagedInstance()
func (client ManagedInstanceClient) GetManagedInstance(ctx context.Context, request GetManagedInstanceRequest) (response GetManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedInstanceResponse")
	}
	return
}

// getManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) getManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/GetManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "GetManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWindowsUpdate Returns a Windows Update object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetWindowsUpdate.go.html to see an example of how to use GetWindowsUpdate API.
// A default retry strategy applies to this operation GetWindowsUpdate()
func (client ManagedInstanceClient) GetWindowsUpdate(ctx context.Context, request GetWindowsUpdateRequest) (response GetWindowsUpdateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWindowsUpdate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWindowsUpdateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWindowsUpdateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWindowsUpdateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWindowsUpdateResponse")
	}
	return
}

// getWindowsUpdate implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) getWindowsUpdate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/windowsUpdates/{windowsUpdateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWindowsUpdateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/WindowsUpdate/GetWindowsUpdate"
		err = common.PostProcessServiceError(err, "ManagedInstance", "GetWindowsUpdate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAllWindowsUpdatesOnManagedInstancesInCompartment Installs all of the available Windows updates for managed instances in a compartment. This applies only to standalone Windows instances. This will not update instances that belong to a group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallAllWindowsUpdatesOnManagedInstancesInCompartment.go.html to see an example of how to use InstallAllWindowsUpdatesOnManagedInstancesInCompartment API.
// A default retry strategy applies to this operation InstallAllWindowsUpdatesOnManagedInstancesInCompartment()
func (client ManagedInstanceClient) InstallAllWindowsUpdatesOnManagedInstancesInCompartment(ctx context.Context, request InstallAllWindowsUpdatesOnManagedInstancesInCompartmentRequest) (response InstallAllWindowsUpdatesOnManagedInstancesInCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installAllWindowsUpdatesOnManagedInstancesInCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallAllWindowsUpdatesOnManagedInstancesInCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallAllWindowsUpdatesOnManagedInstancesInCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallAllWindowsUpdatesOnManagedInstancesInCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallAllWindowsUpdatesOnManagedInstancesInCompartmentResponse")
	}
	return
}

// installAllWindowsUpdatesOnManagedInstancesInCompartment implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) installAllWindowsUpdatesOnManagedInstancesInCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/actions/installWindowsUpdates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallAllWindowsUpdatesOnManagedInstancesInCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/InstallAllWindowsUpdatesOnManagedInstancesInCompartment"
		err = common.PostProcessServiceError(err, "ManagedInstance", "InstallAllWindowsUpdatesOnManagedInstancesInCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallModuleStreamProfileOnManagedInstance Installs a profile for an module stream.  The stream must be
// enabled before a profile can be installed.  If a module stream
// defines multiple profiles, each one can be installed independently.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallModuleStreamProfileOnManagedInstance.go.html to see an example of how to use InstallModuleStreamProfileOnManagedInstance API.
// A default retry strategy applies to this operation InstallModuleStreamProfileOnManagedInstance()
func (client ManagedInstanceClient) InstallModuleStreamProfileOnManagedInstance(ctx context.Context, request InstallModuleStreamProfileOnManagedInstanceRequest) (response InstallModuleStreamProfileOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installModuleStreamProfileOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallModuleStreamProfileOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallModuleStreamProfileOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallModuleStreamProfileOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallModuleStreamProfileOnManagedInstanceResponse")
	}
	return
}

// installModuleStreamProfileOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) installModuleStreamProfileOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/installStreamProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallModuleStreamProfileOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/InstallModuleStreamProfileOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "InstallModuleStreamProfileOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackagesOnManagedInstance Installs packages on a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallPackagesOnManagedInstance.go.html to see an example of how to use InstallPackagesOnManagedInstance API.
// A default retry strategy applies to this operation InstallPackagesOnManagedInstance()
func (client ManagedInstanceClient) InstallPackagesOnManagedInstance(ctx context.Context, request InstallPackagesOnManagedInstanceRequest) (response InstallPackagesOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installPackagesOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallPackagesOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallPackagesOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallPackagesOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallPackagesOnManagedInstanceResponse")
	}
	return
}

// installPackagesOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) installPackagesOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/installPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallPackagesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/InstallPackagesOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "InstallPackagesOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallWindowsUpdatesOnManagedInstance Installs Windows updates on the specified managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallWindowsUpdatesOnManagedInstance.go.html to see an example of how to use InstallWindowsUpdatesOnManagedInstance API.
// A default retry strategy applies to this operation InstallWindowsUpdatesOnManagedInstance()
func (client ManagedInstanceClient) InstallWindowsUpdatesOnManagedInstance(ctx context.Context, request InstallWindowsUpdatesOnManagedInstanceRequest) (response InstallWindowsUpdatesOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installWindowsUpdatesOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallWindowsUpdatesOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallWindowsUpdatesOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallWindowsUpdatesOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallWindowsUpdatesOnManagedInstanceResponse")
	}
	return
}

// installWindowsUpdatesOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) installWindowsUpdatesOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/installWindowsUpdates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallWindowsUpdatesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/InstallWindowsUpdatesOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "InstallWindowsUpdatesOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceAvailablePackages Returns a list of packages that are available for installation on a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceAvailablePackages.go.html to see an example of how to use ListManagedInstanceAvailablePackages API.
// A default retry strategy applies to this operation ListManagedInstanceAvailablePackages()
func (client ManagedInstanceClient) ListManagedInstanceAvailablePackages(ctx context.Context, request ListManagedInstanceAvailablePackagesRequest) (response ListManagedInstanceAvailablePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceAvailablePackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceAvailablePackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceAvailablePackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceAvailablePackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceAvailablePackagesResponse")
	}
	return
}

// listManagedInstanceAvailablePackages implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceAvailablePackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/availablePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceAvailablePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceAvailablePackages"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceAvailablePackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceAvailableSoftwareSources Returns a list of software sources that can be attached to the specified managed instance. Any software sources already attached to the instance are not included in the list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceAvailableSoftwareSources.go.html to see an example of how to use ListManagedInstanceAvailableSoftwareSources API.
// A default retry strategy applies to this operation ListManagedInstanceAvailableSoftwareSources()
func (client ManagedInstanceClient) ListManagedInstanceAvailableSoftwareSources(ctx context.Context, request ListManagedInstanceAvailableSoftwareSourcesRequest) (response ListManagedInstanceAvailableSoftwareSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceAvailableSoftwareSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceAvailableSoftwareSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceAvailableSoftwareSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceAvailableSoftwareSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceAvailableSoftwareSourcesResponse")
	}
	return
}

// listManagedInstanceAvailableSoftwareSources implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceAvailableSoftwareSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/availableSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceAvailableSoftwareSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceAvailableSoftwareSources"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceAvailableSoftwareSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceAvailableWindowsUpdates Returns a list of Windows updates that can be installed on the specified managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceAvailableWindowsUpdates.go.html to see an example of how to use ListManagedInstanceAvailableWindowsUpdates API.
// A default retry strategy applies to this operation ListManagedInstanceAvailableWindowsUpdates()
func (client ManagedInstanceClient) ListManagedInstanceAvailableWindowsUpdates(ctx context.Context, request ListManagedInstanceAvailableWindowsUpdatesRequest) (response ListManagedInstanceAvailableWindowsUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceAvailableWindowsUpdates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceAvailableWindowsUpdatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceAvailableWindowsUpdatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceAvailableWindowsUpdatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceAvailableWindowsUpdatesResponse")
	}
	return
}

// listManagedInstanceAvailableWindowsUpdates implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceAvailableWindowsUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/availableWindowsUpdates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceAvailableWindowsUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceAvailableWindowsUpdates"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceAvailableWindowsUpdates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceErrata Returns a list of applicable errata on the managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceErrata.go.html to see an example of how to use ListManagedInstanceErrata API.
// A default retry strategy applies to this operation ListManagedInstanceErrata()
func (client ManagedInstanceClient) ListManagedInstanceErrata(ctx context.Context, request ListManagedInstanceErrataRequest) (response ListManagedInstanceErrataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceErrata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceErrataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceErrataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceErrataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceErrataResponse")
	}
	return
}

// listManagedInstanceErrata implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceErrata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/errata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceErrataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceErrata"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceErrata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceInstalledPackages Lists the packages that are installed on the managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceInstalledPackages.go.html to see an example of how to use ListManagedInstanceInstalledPackages API.
// A default retry strategy applies to this operation ListManagedInstanceInstalledPackages()
func (client ManagedInstanceClient) ListManagedInstanceInstalledPackages(ctx context.Context, request ListManagedInstanceInstalledPackagesRequest) (response ListManagedInstanceInstalledPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceInstalledPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceInstalledPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceInstalledPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceInstalledPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceInstalledPackagesResponse")
	}
	return
}

// listManagedInstanceInstalledPackages implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceInstalledPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/installedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceInstalledPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceInstalledPackages"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceInstalledPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceInstalledWindowsUpdates Returns a list of Windows updates that have been installed on the specified managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceInstalledWindowsUpdates.go.html to see an example of how to use ListManagedInstanceInstalledWindowsUpdates API.
// A default retry strategy applies to this operation ListManagedInstanceInstalledWindowsUpdates()
func (client ManagedInstanceClient) ListManagedInstanceInstalledWindowsUpdates(ctx context.Context, request ListManagedInstanceInstalledWindowsUpdatesRequest) (response ListManagedInstanceInstalledWindowsUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceInstalledWindowsUpdates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceInstalledWindowsUpdatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceInstalledWindowsUpdatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceInstalledWindowsUpdatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceInstalledWindowsUpdatesResponse")
	}
	return
}

// listManagedInstanceInstalledWindowsUpdates implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceInstalledWindowsUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/installedWindowsUpdates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceInstalledWindowsUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceInstalledWindowsUpdates"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceInstalledWindowsUpdates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceModules Retrieves a list of modules, along with streams of the modules, from a managed instance. Filters may be applied to select a subset of modules based on the filter criteria.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceModules.go.html to see an example of how to use ListManagedInstanceModules API.
// A default retry strategy applies to this operation ListManagedInstanceModules()
func (client ManagedInstanceClient) ListManagedInstanceModules(ctx context.Context, request ListManagedInstanceModulesRequest) (response ListManagedInstanceModulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceModules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceModulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceModulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceModulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceModulesResponse")
	}
	return
}

// listManagedInstanceModules implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceModules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/modules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceModulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceModules"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceModules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceUpdatablePackages Returns a list of updatable packages for a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceUpdatablePackages.go.html to see an example of how to use ListManagedInstanceUpdatablePackages API.
// A default retry strategy applies to this operation ListManagedInstanceUpdatablePackages()
func (client ManagedInstanceClient) ListManagedInstanceUpdatablePackages(ctx context.Context, request ListManagedInstanceUpdatablePackagesRequest) (response ListManagedInstanceUpdatablePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceUpdatablePackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceUpdatablePackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceUpdatablePackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceUpdatablePackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceUpdatablePackagesResponse")
	}
	return
}

// listManagedInstanceUpdatablePackages implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstanceUpdatablePackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/updatablePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceUpdatablePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstanceUpdatablePackages"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstanceUpdatablePackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstances Lists managed instances that match the specified compartment or managed instance OCID. Filter the list against a variety of criteria including but not limited to its name, status, architecture, and OS version.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstances.go.html to see an example of how to use ListManagedInstances API.
// A default retry strategy applies to this operation ListManagedInstances()
func (client ManagedInstanceClient) ListManagedInstances(ctx context.Context, request ListManagedInstancesRequest) (response ListManagedInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstancesResponse")
	}
	return
}

// listManagedInstances implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listManagedInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ListManagedInstances"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListManagedInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWindowsUpdates Lists Windows updates that have been reported to the service.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListWindowsUpdates.go.html to see an example of how to use ListWindowsUpdates API.
// A default retry strategy applies to this operation ListWindowsUpdates()
func (client ManagedInstanceClient) ListWindowsUpdates(ctx context.Context, request ListWindowsUpdatesRequest) (response ListWindowsUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWindowsUpdates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWindowsUpdatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWindowsUpdatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWindowsUpdatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWindowsUpdatesResponse")
	}
	return
}

// listWindowsUpdates implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) listWindowsUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/windowsUpdates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWindowsUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/WindowsUpdateCollection/ListWindowsUpdates"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ListWindowsUpdates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ManageModuleStreamsOnManagedInstance Enables or disables module streams and installs or removes module stream profiles. Once complete, the state of the modules, streams, and
// profiles will match the state indicated in the operation. See ManageModuleStreamsOnManagedInstanceDetails
// for more information. You can preform this operation as a dry run. For a dry run, the service evaluates the operation against the
// current module, stream, and profile state on the managed instance, but does not commit the changes. Instead, the service returns work request log or error entries indicating the impact of the operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ManageModuleStreamsOnManagedInstance.go.html to see an example of how to use ManageModuleStreamsOnManagedInstance API.
// A default retry strategy applies to this operation ManageModuleStreamsOnManagedInstance()
func (client ManagedInstanceClient) ManageModuleStreamsOnManagedInstance(ctx context.Context, request ManageModuleStreamsOnManagedInstanceRequest) (response ManageModuleStreamsOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.manageModuleStreamsOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ManageModuleStreamsOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ManageModuleStreamsOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ManageModuleStreamsOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ManageModuleStreamsOnManagedInstanceResponse")
	}
	return
}

// manageModuleStreamsOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) manageModuleStreamsOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/manageModuleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ManageModuleStreamsOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/ManageModuleStreamsOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "ManageModuleStreamsOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshSoftwareOnManagedInstance Refreshes the package or Windows update information on a managed instance with the latest data from the software source. This does not update packages on the instance. It provides the service with the latest package data.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RefreshSoftwareOnManagedInstance.go.html to see an example of how to use RefreshSoftwareOnManagedInstance API.
// A default retry strategy applies to this operation RefreshSoftwareOnManagedInstance()
func (client ManagedInstanceClient) RefreshSoftwareOnManagedInstance(ctx context.Context, request RefreshSoftwareOnManagedInstanceRequest) (response RefreshSoftwareOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshSoftwareOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshSoftwareOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshSoftwareOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshSoftwareOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshSoftwareOnManagedInstanceResponse")
	}
	return
}

// refreshSoftwareOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) refreshSoftwareOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/refreshSoftware", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshSoftwareOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/RefreshSoftwareOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "RefreshSoftwareOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveModuleStreamProfileFromManagedInstance Removes a profile for a module stream that is installed on a managed instance.
// If a module stream is provided, rather than a fully qualified profile, all
// profiles that have been installed for the module stream will be removed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RemoveModuleStreamProfileFromManagedInstance.go.html to see an example of how to use RemoveModuleStreamProfileFromManagedInstance API.
// A default retry strategy applies to this operation RemoveModuleStreamProfileFromManagedInstance()
func (client ManagedInstanceClient) RemoveModuleStreamProfileFromManagedInstance(ctx context.Context, request RemoveModuleStreamProfileFromManagedInstanceRequest) (response RemoveModuleStreamProfileFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeModuleStreamProfileFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveModuleStreamProfileFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveModuleStreamProfileFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveModuleStreamProfileFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveModuleStreamProfileFromManagedInstanceResponse")
	}
	return
}

// removeModuleStreamProfileFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) removeModuleStreamProfileFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/removeStreamProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveModuleStreamProfileFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/RemoveModuleStreamProfileFromManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "RemoveModuleStreamProfileFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackagesFromManagedInstance Removes an installed package from a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RemovePackagesFromManagedInstance.go.html to see an example of how to use RemovePackagesFromManagedInstance API.
// A default retry strategy applies to this operation RemovePackagesFromManagedInstance()
func (client ManagedInstanceClient) RemovePackagesFromManagedInstance(ctx context.Context, request RemovePackagesFromManagedInstanceRequest) (response RemovePackagesFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removePackagesFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePackagesFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePackagesFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePackagesFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePackagesFromManagedInstanceResponse")
	}
	return
}

// removePackagesFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) removePackagesFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/removePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePackagesFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/RemovePackagesFromManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "RemovePackagesFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchModuleStreamOnManagedInstance Enables a new stream for a module that already has a stream enabled.
// If any profiles or packages from the original module are installed,
// switching to a new stream will remove the existing packages and
// install their counterparts in the new stream.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SwitchModuleStreamOnManagedInstance.go.html to see an example of how to use SwitchModuleStreamOnManagedInstance API.
// A default retry strategy applies to this operation SwitchModuleStreamOnManagedInstance()
func (client ManagedInstanceClient) SwitchModuleStreamOnManagedInstance(ctx context.Context, request SwitchModuleStreamOnManagedInstanceRequest) (response SwitchModuleStreamOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.switchModuleStreamOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchModuleStreamOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchModuleStreamOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchModuleStreamOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchModuleStreamOnManagedInstanceResponse")
	}
	return
}

// switchModuleStreamOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) switchModuleStreamOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/switchModuleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchModuleStreamOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/SwitchModuleStreamOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "SwitchModuleStreamOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAllPackagesOnManagedInstancesInCompartment Install all of the available package updates for all of the managed instances in a compartment. This applies only to standalone non-Windows instances. This will not update instances that belong to a group or lifecycle environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateAllPackagesOnManagedInstancesInCompartment.go.html to see an example of how to use UpdateAllPackagesOnManagedInstancesInCompartment API.
// A default retry strategy applies to this operation UpdateAllPackagesOnManagedInstancesInCompartment()
func (client ManagedInstanceClient) UpdateAllPackagesOnManagedInstancesInCompartment(ctx context.Context, request UpdateAllPackagesOnManagedInstancesInCompartmentRequest) (response UpdateAllPackagesOnManagedInstancesInCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAllPackagesOnManagedInstancesInCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAllPackagesOnManagedInstancesInCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAllPackagesOnManagedInstancesInCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAllPackagesOnManagedInstancesInCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAllPackagesOnManagedInstancesInCompartmentResponse")
	}
	return
}

// updateAllPackagesOnManagedInstancesInCompartment implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) updateAllPackagesOnManagedInstancesInCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/actions/updatePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAllPackagesOnManagedInstancesInCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/UpdateAllPackagesOnManagedInstancesInCompartment"
		err = common.PostProcessServiceError(err, "ManagedInstance", "UpdateAllPackagesOnManagedInstancesInCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedInstance Updates the specified managed instance information, such as description, ONS topic, and associated management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateManagedInstance.go.html to see an example of how to use UpdateManagedInstance API.
// A default retry strategy applies to this operation UpdateManagedInstance()
func (client ManagedInstanceClient) UpdateManagedInstance(ctx context.Context, request UpdateManagedInstanceRequest) (response UpdateManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagedInstanceResponse")
	}
	return
}

// updateManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) updateManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedInstances/{managedInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/UpdateManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "UpdateManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePackagesOnManagedInstance Updates a package on a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdatePackagesOnManagedInstance.go.html to see an example of how to use UpdatePackagesOnManagedInstance API.
// A default retry strategy applies to this operation UpdatePackagesOnManagedInstance()
func (client ManagedInstanceClient) UpdatePackagesOnManagedInstance(ctx context.Context, request UpdatePackagesOnManagedInstanceRequest) (response UpdatePackagesOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePackagesOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePackagesOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePackagesOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePackagesOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePackagesOnManagedInstanceResponse")
	}
	return
}

// updatePackagesOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceClient) updatePackagesOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/updatePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePackagesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstance/UpdatePackagesOnManagedInstance"
		err = common.PostProcessServiceError(err, "ManagedInstance", "UpdatePackagesOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
