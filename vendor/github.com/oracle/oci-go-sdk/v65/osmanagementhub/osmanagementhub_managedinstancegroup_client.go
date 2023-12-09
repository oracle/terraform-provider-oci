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

// ManagedInstanceGroupClient a client for ManagedInstanceGroup
type ManagedInstanceGroupClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagedInstanceGroupClientWithConfigurationProvider Creates a new default ManagedInstanceGroup client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagedInstanceGroupClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagedInstanceGroupClient, err error) {
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
	return newManagedInstanceGroupClientFromBaseClient(baseClient, provider)
}

// NewManagedInstanceGroupClientWithOboToken Creates a new default ManagedInstanceGroup client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewManagedInstanceGroupClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagedInstanceGroupClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newManagedInstanceGroupClientFromBaseClient(baseClient, configProvider)
}

func newManagedInstanceGroupClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagedInstanceGroupClient, err error) {
	// ManagedInstanceGroup service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ManagedInstanceGroup"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ManagedInstanceGroupClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagedInstanceGroupClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagedInstanceGroupClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ManagedInstanceGroupClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachManagedInstancesToManagedInstanceGroup Adds managed instances to the specified managed instance group. After the managed
// instances have been added, then operations can be performed on the managed
// instance group which will then apply to all managed instances in the
// group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachManagedInstancesToManagedInstanceGroup.go.html to see an example of how to use AttachManagedInstancesToManagedInstanceGroup API.
// A default retry strategy applies to this operation AttachManagedInstancesToManagedInstanceGroup()
func (client ManagedInstanceGroupClient) AttachManagedInstancesToManagedInstanceGroup(ctx context.Context, request AttachManagedInstancesToManagedInstanceGroupRequest) (response AttachManagedInstancesToManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachManagedInstancesToManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachManagedInstancesToManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachManagedInstancesToManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachManagedInstancesToManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachManagedInstancesToManagedInstanceGroupResponse")
	}
	return
}

// attachManagedInstancesToManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) attachManagedInstancesToManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/attachManagedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachManagedInstancesToManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/AttachManagedInstancesToManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "AttachManagedInstancesToManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachSoftwareSourcesToManagedInstanceGroup Attaches software sources to the specified managed instance group. The software sources must be compatible with the content for the managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachSoftwareSourcesToManagedInstanceGroup.go.html to see an example of how to use AttachSoftwareSourcesToManagedInstanceGroup API.
// A default retry strategy applies to this operation AttachSoftwareSourcesToManagedInstanceGroup()
func (client ManagedInstanceGroupClient) AttachSoftwareSourcesToManagedInstanceGroup(ctx context.Context, request AttachSoftwareSourcesToManagedInstanceGroupRequest) (response AttachSoftwareSourcesToManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachSoftwareSourcesToManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachSoftwareSourcesToManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachSoftwareSourcesToManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachSoftwareSourcesToManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachSoftwareSourcesToManagedInstanceGroupResponse")
	}
	return
}

// attachSoftwareSourcesToManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) attachSoftwareSourcesToManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/attachSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachSoftwareSourcesToManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/AttachSoftwareSourcesToManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "AttachSoftwareSourcesToManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateManagedInstanceGroup Creates a new managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateManagedInstanceGroup.go.html to see an example of how to use CreateManagedInstanceGroup API.
// A default retry strategy applies to this operation CreateManagedInstanceGroup()
func (client ManagedInstanceGroupClient) CreateManagedInstanceGroup(ctx context.Context, request CreateManagedInstanceGroupRequest) (response CreateManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagedInstanceGroupResponse")
	}
	return
}

// createManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) createManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/CreateManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "CreateManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedInstanceGroup Deletes a specified managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteManagedInstanceGroup.go.html to see an example of how to use DeleteManagedInstanceGroup API.
// A default retry strategy applies to this operation DeleteManagedInstanceGroup()
func (client ManagedInstanceGroupClient) DeleteManagedInstanceGroup(ctx context.Context, request DeleteManagedInstanceGroupRequest) (response DeleteManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagedInstanceGroupResponse")
	}
	return
}

// deleteManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) deleteManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedInstanceGroups/{managedInstanceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/DeleteManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "DeleteManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachManagedInstancesFromManagedInstanceGroup Removes a managed instance from the specified managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DetachManagedInstancesFromManagedInstanceGroup.go.html to see an example of how to use DetachManagedInstancesFromManagedInstanceGroup API.
// A default retry strategy applies to this operation DetachManagedInstancesFromManagedInstanceGroup()
func (client ManagedInstanceGroupClient) DetachManagedInstancesFromManagedInstanceGroup(ctx context.Context, request DetachManagedInstancesFromManagedInstanceGroupRequest) (response DetachManagedInstancesFromManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachManagedInstancesFromManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachManagedInstancesFromManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachManagedInstancesFromManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachManagedInstancesFromManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachManagedInstancesFromManagedInstanceGroupResponse")
	}
	return
}

// detachManagedInstancesFromManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) detachManagedInstancesFromManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/detachManagedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachManagedInstancesFromManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/DetachManagedInstancesFromManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "DetachManagedInstancesFromManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachSoftwareSourcesFromManagedInstanceGroup Detaches software sources from a group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DetachSoftwareSourcesFromManagedInstanceGroup.go.html to see an example of how to use DetachSoftwareSourcesFromManagedInstanceGroup API.
// A default retry strategy applies to this operation DetachSoftwareSourcesFromManagedInstanceGroup()
func (client ManagedInstanceGroupClient) DetachSoftwareSourcesFromManagedInstanceGroup(ctx context.Context, request DetachSoftwareSourcesFromManagedInstanceGroupRequest) (response DetachSoftwareSourcesFromManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachSoftwareSourcesFromManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachSoftwareSourcesFromManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachSoftwareSourcesFromManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachSoftwareSourcesFromManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachSoftwareSourcesFromManagedInstanceGroupResponse")
	}
	return
}

// detachSoftwareSourcesFromManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) detachSoftwareSourcesFromManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/detachSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachSoftwareSourcesFromManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/DetachSoftwareSourcesFromManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "DetachSoftwareSourcesFromManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableModuleStreamOnManagedInstanceGroup Disables a module stream on a managed instance group. After the stream is
// disabled, it is no longer possible to install the profiles that are
// contained by the stream. All installed profiles must be removed prior
// to disabling a module stream.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DisableModuleStreamOnManagedInstanceGroup.go.html to see an example of how to use DisableModuleStreamOnManagedInstanceGroup API.
// A default retry strategy applies to this operation DisableModuleStreamOnManagedInstanceGroup()
func (client ManagedInstanceGroupClient) DisableModuleStreamOnManagedInstanceGroup(ctx context.Context, request DisableModuleStreamOnManagedInstanceGroupRequest) (response DisableModuleStreamOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableModuleStreamOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableModuleStreamOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableModuleStreamOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableModuleStreamOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableModuleStreamOnManagedInstanceGroupResponse")
	}
	return
}

// disableModuleStreamOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) disableModuleStreamOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/disableModuleStream", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableModuleStreamOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/DisableModuleStreamOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "DisableModuleStreamOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableModuleStreamOnManagedInstanceGroup Enables a module stream on a managed instance group.  After the stream is
// enabled, it is possible to install the profiles that are contained
// by the stream.  Enabling a stream that is already enabled will
// succeed.  Attempting to enable a different stream for a module that
// already has a stream enabled results in an error.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/EnableModuleStreamOnManagedInstanceGroup.go.html to see an example of how to use EnableModuleStreamOnManagedInstanceGroup API.
// A default retry strategy applies to this operation EnableModuleStreamOnManagedInstanceGroup()
func (client ManagedInstanceGroupClient) EnableModuleStreamOnManagedInstanceGroup(ctx context.Context, request EnableModuleStreamOnManagedInstanceGroupRequest) (response EnableModuleStreamOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableModuleStreamOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableModuleStreamOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableModuleStreamOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableModuleStreamOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableModuleStreamOnManagedInstanceGroupResponse")
	}
	return
}

// enableModuleStreamOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) enableModuleStreamOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/enableModuleStream", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableModuleStreamOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/EnableModuleStreamOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "EnableModuleStreamOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstanceGroup Gets information about the specified managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagedInstanceGroup.go.html to see an example of how to use GetManagedInstanceGroup API.
// A default retry strategy applies to this operation GetManagedInstanceGroup()
func (client ManagedInstanceGroupClient) GetManagedInstanceGroup(ctx context.Context, request GetManagedInstanceGroupRequest) (response GetManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedInstanceGroupResponse")
	}
	return
}

// getManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) getManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/GetManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "GetManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallModuleStreamProfileOnManagedInstanceGroup Installs a profile for an module stream. The stream must be
// enabled before a profile can be installed. If a module stream
// defines multiple profiles, each one can be installed independently.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallModuleStreamProfileOnManagedInstanceGroup.go.html to see an example of how to use InstallModuleStreamProfileOnManagedInstanceGroup API.
// A default retry strategy applies to this operation InstallModuleStreamProfileOnManagedInstanceGroup()
func (client ManagedInstanceGroupClient) InstallModuleStreamProfileOnManagedInstanceGroup(ctx context.Context, request InstallModuleStreamProfileOnManagedInstanceGroupRequest) (response InstallModuleStreamProfileOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installModuleStreamProfileOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallModuleStreamProfileOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallModuleStreamProfileOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallModuleStreamProfileOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallModuleStreamProfileOnManagedInstanceGroupResponse")
	}
	return
}

// installModuleStreamProfileOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) installModuleStreamProfileOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/installStreamProfile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallModuleStreamProfileOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/InstallModuleStreamProfileOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "InstallModuleStreamProfileOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackagesOnManagedInstanceGroup Installs package(s) on each managed instance in a managed instance group. The package must be compatible with the
// instances in the managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallPackagesOnManagedInstanceGroup.go.html to see an example of how to use InstallPackagesOnManagedInstanceGroup API.
// A default retry strategy applies to this operation InstallPackagesOnManagedInstanceGroup()
func (client ManagedInstanceGroupClient) InstallPackagesOnManagedInstanceGroup(ctx context.Context, request InstallPackagesOnManagedInstanceGroupRequest) (response InstallPackagesOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installPackagesOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallPackagesOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallPackagesOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallPackagesOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallPackagesOnManagedInstanceGroupResponse")
	}
	return
}

// installPackagesOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) installPackagesOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/installPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallPackagesOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/InstallPackagesOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "InstallPackagesOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroupAvailableModules Lists available modules that for the specified managed instance group. Filter the list against a variety of
// criteria including but not limited to its name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupAvailableModules.go.html to see an example of how to use ListManagedInstanceGroupAvailableModules API.
// A default retry strategy applies to this operation ListManagedInstanceGroupAvailableModules()
func (client ManagedInstanceGroupClient) ListManagedInstanceGroupAvailableModules(ctx context.Context, request ListManagedInstanceGroupAvailableModulesRequest) (response ListManagedInstanceGroupAvailableModulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceGroupAvailableModules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceGroupAvailableModulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceGroupAvailableModulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceGroupAvailableModulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceGroupAvailableModulesResponse")
	}
	return
}

// listManagedInstanceGroupAvailableModules implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) listManagedInstanceGroupAvailableModules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}/availableModules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupAvailableModulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ListManagedInstanceGroupAvailableModules"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ListManagedInstanceGroupAvailableModules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroupAvailablePackages Lists available packages on the specified managed instances group. Filter the list against a variety
// of criteria including but not limited to the package name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupAvailablePackages.go.html to see an example of how to use ListManagedInstanceGroupAvailablePackages API.
// A default retry strategy applies to this operation ListManagedInstanceGroupAvailablePackages()
func (client ManagedInstanceGroupClient) ListManagedInstanceGroupAvailablePackages(ctx context.Context, request ListManagedInstanceGroupAvailablePackagesRequest) (response ListManagedInstanceGroupAvailablePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceGroupAvailablePackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceGroupAvailablePackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceGroupAvailablePackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceGroupAvailablePackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceGroupAvailablePackagesResponse")
	}
	return
}

// listManagedInstanceGroupAvailablePackages implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) listManagedInstanceGroupAvailablePackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}/availablePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupAvailablePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ListManagedInstanceGroupAvailablePackages"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ListManagedInstanceGroupAvailablePackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroupAvailableSoftwareSources Lists available software sources for a specified managed instance group. Filter the list against a variety of
// criteria including but not limited to its name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupAvailableSoftwareSources.go.html to see an example of how to use ListManagedInstanceGroupAvailableSoftwareSources API.
// A default retry strategy applies to this operation ListManagedInstanceGroupAvailableSoftwareSources()
func (client ManagedInstanceGroupClient) ListManagedInstanceGroupAvailableSoftwareSources(ctx context.Context, request ListManagedInstanceGroupAvailableSoftwareSourcesRequest) (response ListManagedInstanceGroupAvailableSoftwareSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceGroupAvailableSoftwareSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceGroupAvailableSoftwareSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceGroupAvailableSoftwareSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceGroupAvailableSoftwareSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceGroupAvailableSoftwareSourcesResponse")
	}
	return
}

// listManagedInstanceGroupAvailableSoftwareSources implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) listManagedInstanceGroupAvailableSoftwareSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}/availableSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupAvailableSoftwareSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ListManagedInstanceGroupAvailableSoftwareSources"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ListManagedInstanceGroupAvailableSoftwareSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroupInstalledPackages Lists installed packages on the specified managed instances group. Filter the list against a variety
// of criteria including but not limited to the package name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupInstalledPackages.go.html to see an example of how to use ListManagedInstanceGroupInstalledPackages API.
// A default retry strategy applies to this operation ListManagedInstanceGroupInstalledPackages()
func (client ManagedInstanceGroupClient) ListManagedInstanceGroupInstalledPackages(ctx context.Context, request ListManagedInstanceGroupInstalledPackagesRequest) (response ListManagedInstanceGroupInstalledPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceGroupInstalledPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceGroupInstalledPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceGroupInstalledPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceGroupInstalledPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceGroupInstalledPackagesResponse")
	}
	return
}

// listManagedInstanceGroupInstalledPackages implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) listManagedInstanceGroupInstalledPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}/installedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupInstalledPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ListManagedInstanceGroupInstalledPackages"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ListManagedInstanceGroupInstalledPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroupModules Retrieve a list of module streams, along with a summary of their
// status, from a managed instance group.  Filters may be applied to select
// a subset of module streams based on the filter criteria.
// The 'moduleName' attribute filters against the name of a module.
// It accepts strings of the format "<module>".  If this attribute
// is defined, only streams that belong to the specified module are
// included in the result set.  If it is not defined, the request is
// not subject to this filter.
// The "status" attribute filters against the state of a module stream.
// Valid values are "ENABLED", "DISABLED", and "ACTIVE".  If the
// attribute is set to "ENABLED", only module streams that are enabled
// are included in the result set.  If the attribute is set to "DISABLED",
// only module streams that are not enabled are included in the result
// set.  If the attribute is set to "ACTIVE", only module streams that
// are active are included in the result set.  If the attribute is not
// defined, the request is not subject to this filter.
// When sorting by the display name, the result set is sorted first
// by the module name and then by the stream name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupModules.go.html to see an example of how to use ListManagedInstanceGroupModules API.
// A default retry strategy applies to this operation ListManagedInstanceGroupModules()
func (client ManagedInstanceGroupClient) ListManagedInstanceGroupModules(ctx context.Context, request ListManagedInstanceGroupModulesRequest) (response ListManagedInstanceGroupModulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceGroupModules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceGroupModulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceGroupModulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceGroupModulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceGroupModulesResponse")
	}
	return
}

// listManagedInstanceGroupModules implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) listManagedInstanceGroupModules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}/modules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupModulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ListManagedInstanceGroupModules"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ListManagedInstanceGroupModules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroups Lists managed instance groups that match the specified compartment or managed instance group OCID. Filter the
// list against a variety of criteria including but not limited to its name, status, architecture, and OS family.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroups.go.html to see an example of how to use ListManagedInstanceGroups API.
// A default retry strategy applies to this operation ListManagedInstanceGroups()
func (client ManagedInstanceGroupClient) ListManagedInstanceGroups(ctx context.Context, request ListManagedInstanceGroupsRequest) (response ListManagedInstanceGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceGroupsResponse")
	}
	return
}

// listManagedInstanceGroups implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) listManagedInstanceGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ListManagedInstanceGroups"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ListManagedInstanceGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ManageModuleStreamsOnManagedInstanceGroup Perform an operation involving modules, streams, and profiles on a
// managed instance group.  Each operation may enable or disable an arbitrary
// amount of module streams, and install or remove an arbitrary number
// of module stream profiles.  When the operation is complete, the
// state of the modules, streams, and profiles on the managed instance group
// will match the state indicated in the operation.
// Each module stream specified in the list of module streams to enable
// will be in the "ENABLED" state upon completion of the operation.
// If there was already a stream of that module enabled, any work
// required to switch from the current stream to the new stream is
// performed implicitly.
// Each module stream specified in the list of module streams to disable
// will be in the "DISABLED" state upon completion of the operation.
// Any profiles that are installed for the module stream will be removed
// as part of the operation.
// Each module stream profile specified in the list of profiles to install
// will be in the "INSTALLED" state upon completion of the operation,
// indicating that any packages that are part of the profile are installed
// on the managed instance.  If the module stream containing the profile
// is not enabled, it will be enabled as part of the operation. There
// is an exception when attempting to install a stream of a profile when
// another stream of the same module is enabled.  It is an error to attempt
// to install a profile of another module stream, unless enabling the
// new module stream is explicitly included in this operation.
// Each module stream profile specified in the list of profiles to remove
// will be in the "AVAILABLE" state upon completion of the operation.
// The status of packages within the profile after the operation is
// complete is defined by the package manager on the managed instance group.
// Operations that contain one or more elements that are not allowed
// are rejected.
// The result of this request is a work request object. The returned
// work request is the parent of a structure of other work requests.  Taken
// as a whole, this structure indicates the entire set of work to be
// performed to complete the operation.
// This interface can also be used to perform a dry run of the operation
// rather than committing it to a managed instance group.  If a dry run is
// requested, the OS Management Hub service will evaluate the operation
// against the current module, stream, and profile state on the managed
// instance.  It will calculate the impact of the operation on all
// modules, streams, and profiles on the managed instance, including those
// that are implicitly impacted by the operation.
// The work request resulting from a dry run behaves differently than
// a work request resulting from a committable operation.  Dry run
// work requests are always singletons and never have children. The
// impact of the operation is returned using the log and error
// facilities of work requests. The impact of operations that are
// allowed by the OS Management Hub service are communicated as one or
// more work request log entries.  Operations that are not allowed
// by the OS Management Hub service are communicated as one or more
// work request error entries.  Each entry, for either logs or errors,
// contains a structured message containing the results of one
// or more operations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ManageModuleStreamsOnManagedInstanceGroup.go.html to see an example of how to use ManageModuleStreamsOnManagedInstanceGroup API.
// A default retry strategy applies to this operation ManageModuleStreamsOnManagedInstanceGroup()
func (client ManagedInstanceGroupClient) ManageModuleStreamsOnManagedInstanceGroup(ctx context.Context, request ManageModuleStreamsOnManagedInstanceGroupRequest) (response ManageModuleStreamsOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.manageModuleStreamsOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ManageModuleStreamsOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ManageModuleStreamsOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ManageModuleStreamsOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ManageModuleStreamsOnManagedInstanceGroupResponse")
	}
	return
}

// manageModuleStreamsOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) manageModuleStreamsOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/manageModuleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ManageModuleStreamsOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/ManageModuleStreamsOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "ManageModuleStreamsOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveModuleStreamProfileFromManagedInstanceGroup Removes a profile for a module stream that is installed on a managed instance group.
// If a module stream is provided, rather than a fully qualified profile, all
// profiles that have been installed for the module stream will be removed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RemoveModuleStreamProfileFromManagedInstanceGroup.go.html to see an example of how to use RemoveModuleStreamProfileFromManagedInstanceGroup API.
// A default retry strategy applies to this operation RemoveModuleStreamProfileFromManagedInstanceGroup()
func (client ManagedInstanceGroupClient) RemoveModuleStreamProfileFromManagedInstanceGroup(ctx context.Context, request RemoveModuleStreamProfileFromManagedInstanceGroupRequest) (response RemoveModuleStreamProfileFromManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeModuleStreamProfileFromManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveModuleStreamProfileFromManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveModuleStreamProfileFromManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveModuleStreamProfileFromManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveModuleStreamProfileFromManagedInstanceGroupResponse")
	}
	return
}

// removeModuleStreamProfileFromManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) removeModuleStreamProfileFromManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/removeStreamProfile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveModuleStreamProfileFromManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/RemoveModuleStreamProfileFromManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "RemoveModuleStreamProfileFromManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackagesFromManagedInstanceGroup Removes package(s) from each managed instance in a specified managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RemovePackagesFromManagedInstanceGroup.go.html to see an example of how to use RemovePackagesFromManagedInstanceGroup API.
// A default retry strategy applies to this operation RemovePackagesFromManagedInstanceGroup()
func (client ManagedInstanceGroupClient) RemovePackagesFromManagedInstanceGroup(ctx context.Context, request RemovePackagesFromManagedInstanceGroupRequest) (response RemovePackagesFromManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removePackagesFromManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePackagesFromManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePackagesFromManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePackagesFromManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePackagesFromManagedInstanceGroupResponse")
	}
	return
}

// removePackagesFromManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) removePackagesFromManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/removePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePackagesFromManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/RemovePackagesFromManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "RemovePackagesFromManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAllPackagesOnManagedInstanceGroup Updates all packages on each managed instance in the specified managed instance group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateAllPackagesOnManagedInstanceGroup.go.html to see an example of how to use UpdateAllPackagesOnManagedInstanceGroup API.
// A default retry strategy applies to this operation UpdateAllPackagesOnManagedInstanceGroup()
func (client ManagedInstanceGroupClient) UpdateAllPackagesOnManagedInstanceGroup(ctx context.Context, request UpdateAllPackagesOnManagedInstanceGroupRequest) (response UpdateAllPackagesOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAllPackagesOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAllPackagesOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAllPackagesOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAllPackagesOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAllPackagesOnManagedInstanceGroupResponse")
	}
	return
}

// updateAllPackagesOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) updateAllPackagesOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/updateAllPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAllPackagesOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/UpdateAllPackagesOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "UpdateAllPackagesOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedInstanceGroup Updates the specified managed instance group's name, description, and tags.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateManagedInstanceGroup.go.html to see an example of how to use UpdateManagedInstanceGroup API.
// A default retry strategy applies to this operation UpdateManagedInstanceGroup()
func (client ManagedInstanceGroupClient) UpdateManagedInstanceGroup(ctx context.Context, request UpdateManagedInstanceGroupRequest) (response UpdateManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagedInstanceGroupResponse")
	}
	return
}

// updateManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client ManagedInstanceGroupClient) updateManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedInstanceGroups/{managedInstanceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagedInstanceGroup/UpdateManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "ManagedInstanceGroup", "UpdateManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
