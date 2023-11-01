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

// ManagedInstanceClient a client for ManagedInstance
type ManagedInstanceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagedInstanceClientWithConfigurationProvider Creates a new default ManagedInstance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagedInstanceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagedInstanceClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagementhub"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
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

// ListManagedInstanceAvailablePackages Returns a list of available packages for a managed instance.
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

// ListManagedInstanceAvailableSoftwareSources Returns a list of available software sources for a managed instance.
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

// ListManagedInstanceModules Retrieve a list of modules, along with streams of the modules,
// from a managed instance.  Filters may be applied to select
// a subset of modules based on the filter criteria.
// The 'name' attribute filters against the name of a module.
// It accepts strings of the format "<string>".
// The 'nameContains' attribute filters against the name of a module
// based on partial match. It accepts strings of the format "<string>".
// If this attribute is defined, only matching modules are included
// in the result set. If it is not defined, the request  is not subject
// to this filter.
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

// ManageModuleStreamsOnManagedInstance Perform an operation involving modules, streams, and profiles on a
// managed instance.  Each operation may enable or disable an arbitrary
// amount of module streams, and install or remove an arbitrary number
// of module stream profiles.  When the operation is complete, the
// state of the modules, streams, and profiles on the managed instance
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
// is not enabled, it will be enabled as part of the operation.  There
// is an exception when attempting to install a stream of a profile when
// another stream of the same module is enabled.  It is an error to attempt
// to install a profile of another module stream, unless enabling the
// new module stream is explicitly included in this operation.
// Each module stream profile specified in the list of profiles to remove
// will be in the "AVAILABLE" state upon completion of the operation.
// The status of packages within the profile after the operation is
// complete is defined by the package manager on the managed instance.
// Operations that contain one or more elements that are not allowed
// are rejected.
// The result of this request is a work request object.  The returned
// work request is the parent of a structure of other WorkRequests.  Taken
// as a whole, this structure indicates the entire set of work to be
// performed to complete the operation.
// This interface can also be used to perform a dry run of the operation
// rather than committing it to a managed instance.  If a dry run is
// requested, the OS Management Hub service will evaluate the operation
// against the current module, stream, and profile state on the managed
// instance.  It will calculate the impact of the operation on all
// modules, streams, and profiles on the managed instance, including those
// that are implicitly impacted by the operation.
// The WorkRequest resulting from a dry run behaves differently than
// a WorkRequest resulting from a committable operation.  Dry run
// WorkRequests are always singletons and never have children.  The
// impact of the operation is returned using the log and error
// facilities of work requests.  The impact of operations that are
// allowed by the OS Management Hub service are communicated as one or
// more work request log entries.  Operations that are not allowed
// by the OS Management Hub service are communicated as one or more
// work request error entries.  Each entry, for either logs or errors,
// contains a structured message containing the results of one
// or more operations.
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

// RefreshSoftwareOnManagedInstance Refresh all installed and updatable software information on a managed instance.
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

// UpdateAllPackagesOnManagedInstancesInCompartment Install all of the available package updates for all of the managed instances in a compartment.
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

// UpdateManagedInstance Updates the managed instance.
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
