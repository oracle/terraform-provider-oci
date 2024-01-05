// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OsManagementClient a client for OsManagement
type OsManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOsManagementClientWithConfigurationProvider Creates a new default OsManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOsManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OsManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagement"); !enabled {
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
	return newOsManagementClientFromBaseClient(baseClient, provider)
}

// NewOsManagementClientWithOboToken Creates a new default OsManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOsManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OsManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOsManagementClientFromBaseClient(baseClient, configProvider)
}

func newOsManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OsManagementClient, err error) {
	// OsManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OsManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OsManagementClient{BaseClient: baseClient}
	client.BasePath = "20190801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OsManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagement", "https://osms.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OsManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OsManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddPackagesToSoftwareSource Adds a given list of Software Packages to a specific Software Source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/AddPackagesToSoftwareSource.go.html to see an example of how to use AddPackagesToSoftwareSource API.
func (client OsManagementClient) AddPackagesToSoftwareSource(ctx context.Context, request AddPackagesToSoftwareSourceRequest) (response AddPackagesToSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addPackagesToSoftwareSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddPackagesToSoftwareSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddPackagesToSoftwareSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddPackagesToSoftwareSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddPackagesToSoftwareSourceResponse")
	}
	return
}

// addPackagesToSoftwareSource implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) addPackagesToSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/{softwareSourceId}/actions/addPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddPackagesToSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/AddPackagesToSoftwareSource"
		err = common.PostProcessServiceError(err, "OsManagement", "AddPackagesToSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachChildSoftwareSourceToManagedInstance Adds a child software source to a managed instance. After the software
// source has been added, then packages from that software source can be
// installed on the managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/AttachChildSoftwareSourceToManagedInstance.go.html to see an example of how to use AttachChildSoftwareSourceToManagedInstance API.
func (client OsManagementClient) AttachChildSoftwareSourceToManagedInstance(ctx context.Context, request AttachChildSoftwareSourceToManagedInstanceRequest) (response AttachChildSoftwareSourceToManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachChildSoftwareSourceToManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachChildSoftwareSourceToManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachChildSoftwareSourceToManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachChildSoftwareSourceToManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachChildSoftwareSourceToManagedInstanceResponse")
	}
	return
}

// attachChildSoftwareSourceToManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) attachChildSoftwareSourceToManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/attachChildSoftwareSource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachChildSoftwareSourceToManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/AttachChildSoftwareSourceToManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "AttachChildSoftwareSourceToManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachManagedInstanceToManagedInstanceGroup Adds a Managed Instance to a Managed Instance Group. After the Managed
// Instance has been added, then operations can be performed on the Managed
// Instance Group which will then apply to all Managed Instances in the
// group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/AttachManagedInstanceToManagedInstanceGroup.go.html to see an example of how to use AttachManagedInstanceToManagedInstanceGroup API.
func (client OsManagementClient) AttachManagedInstanceToManagedInstanceGroup(ctx context.Context, request AttachManagedInstanceToManagedInstanceGroupRequest) (response AttachManagedInstanceToManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachManagedInstanceToManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachManagedInstanceToManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachManagedInstanceToManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachManagedInstanceToManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachManagedInstanceToManagedInstanceGroupResponse")
	}
	return
}

// attachManagedInstanceToManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) attachManagedInstanceToManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/attachManagedInstance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachManagedInstanceToManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/AttachManagedInstanceToManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "AttachManagedInstanceToManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachParentSoftwareSourceToManagedInstance Adds a parent software source to a managed instance. After the software
// source has been added, then packages from that software source can be
// installed on the managed instance. Software sources that have this
// software source as a parent will be able to be added to this managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/AttachParentSoftwareSourceToManagedInstance.go.html to see an example of how to use AttachParentSoftwareSourceToManagedInstance API.
func (client OsManagementClient) AttachParentSoftwareSourceToManagedInstance(ctx context.Context, request AttachParentSoftwareSourceToManagedInstanceRequest) (response AttachParentSoftwareSourceToManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachParentSoftwareSourceToManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachParentSoftwareSourceToManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachParentSoftwareSourceToManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachParentSoftwareSourceToManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachParentSoftwareSourceToManagedInstanceResponse")
	}
	return
}

// attachParentSoftwareSourceToManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) attachParentSoftwareSourceToManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/attachParentSoftwareSource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachParentSoftwareSourceToManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/AttachParentSoftwareSourceToManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "AttachParentSoftwareSourceToManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeManagedInstanceGroupCompartment Moves a resource into a different compartment. When provided, If-Match
// is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ChangeManagedInstanceGroupCompartment.go.html to see an example of how to use ChangeManagedInstanceGroupCompartment API.
func (client OsManagementClient) ChangeManagedInstanceGroupCompartment(ctx context.Context, request ChangeManagedInstanceGroupCompartmentRequest) (response ChangeManagedInstanceGroupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeManagedInstanceGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagedInstanceGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagedInstanceGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagedInstanceGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagedInstanceGroupCompartmentResponse")
	}
	return
}

// changeManagedInstanceGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) changeManagedInstanceGroupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeManagedInstanceGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/ChangeManagedInstanceGroupCompartment"
		err = common.PostProcessServiceError(err, "OsManagement", "ChangeManagedInstanceGroupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeScheduledJobCompartment Moves a resource into a different compartment. When provided, If-Match
// is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ChangeScheduledJobCompartment.go.html to see an example of how to use ChangeScheduledJobCompartment API.
func (client OsManagementClient) ChangeScheduledJobCompartment(ctx context.Context, request ChangeScheduledJobCompartmentRequest) (response ChangeScheduledJobCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeScheduledJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeScheduledJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeScheduledJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeScheduledJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeScheduledJobCompartmentResponse")
	}
	return
}

// changeScheduledJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) changeScheduledJobCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeScheduledJobCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/ChangeScheduledJobCompartment"
		err = common.PostProcessServiceError(err, "OsManagement", "ChangeScheduledJobCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSoftwareSourceCompartment Moves a resource into a different compartment. When provided, If-Match
// is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ChangeSoftwareSourceCompartment.go.html to see an example of how to use ChangeSoftwareSourceCompartment API.
func (client OsManagementClient) ChangeSoftwareSourceCompartment(ctx context.Context, request ChangeSoftwareSourceCompartmentRequest) (response ChangeSoftwareSourceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSoftwareSourceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSoftwareSourceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSoftwareSourceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSoftwareSourceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSoftwareSourceCompartmentResponse")
	}
	return
}

// changeSoftwareSourceCompartment implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) changeSoftwareSourceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/{softwareSourceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSoftwareSourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/ChangeSoftwareSourceCompartment"
		err = common.PostProcessServiceError(err, "OsManagement", "ChangeSoftwareSourceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateManagedInstanceGroup Creates a new Managed Instance Group on the management system.
// This will not contain any managed instances after it is first created,
// and they must be added later.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/CreateManagedInstanceGroup.go.html to see an example of how to use CreateManagedInstanceGroup API.
func (client OsManagementClient) CreateManagedInstanceGroup(ctx context.Context, request CreateManagedInstanceGroupRequest) (response CreateManagedInstanceGroupResponse, err error) {
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
func (client OsManagementClient) createManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/CreateManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "CreateManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateScheduledJob Creates a new Scheduled Job to perform a specific package operation on
// a set of managed instances or managed instance groups.  Can be created
// as a one-time execution in the future, or as a recurring execution
// that repeats on a defined interval.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/CreateScheduledJob.go.html to see an example of how to use CreateScheduledJob API.
func (client OsManagementClient) CreateScheduledJob(ctx context.Context, request CreateScheduledJobRequest) (response CreateScheduledJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduledJobResponse")
	}
	return
}

// createScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) createScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/CreateScheduledJob"
		err = common.PostProcessServiceError(err, "OsManagement", "CreateScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSoftwareSource Creates a new custom Software Source on the management system.
// This will not contain any packages after it is first created,
// and they must be added later.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/CreateSoftwareSource.go.html to see an example of how to use CreateSoftwareSource API.
func (client OsManagementClient) CreateSoftwareSource(ctx context.Context, request CreateSoftwareSourceRequest) (response CreateSoftwareSourceResponse, err error) {
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
func (client OsManagementClient) createSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/CreateSoftwareSource"
		err = common.PostProcessServiceError(err, "OsManagement", "CreateSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedInstanceGroup Deletes a Managed Instance Group from the management system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DeleteManagedInstanceGroup.go.html to see an example of how to use DeleteManagedInstanceGroup API.
func (client OsManagementClient) DeleteManagedInstanceGroup(ctx context.Context, request DeleteManagedInstanceGroupRequest) (response DeleteManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) deleteManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/DeleteManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "DeleteManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteScheduledJob Cancels an existing Scheduled Job on the management system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DeleteScheduledJob.go.html to see an example of how to use DeleteScheduledJob API.
func (client OsManagementClient) DeleteScheduledJob(ctx context.Context, request DeleteScheduledJobRequest) (response DeleteScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduledJobResponse")
	}
	return
}

// deleteScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) deleteScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/scheduledJobs/{scheduledJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/DeleteScheduledJob"
		err = common.PostProcessServiceError(err, "OsManagement", "DeleteScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSoftwareSource Deletes a custom Software Source on the management system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DeleteSoftwareSource.go.html to see an example of how to use DeleteSoftwareSource API.
func (client OsManagementClient) DeleteSoftwareSource(ctx context.Context, request DeleteSoftwareSourceRequest) (response DeleteSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) deleteSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/DeleteSoftwareSource"
		err = common.PostProcessServiceError(err, "OsManagement", "DeleteSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachChildSoftwareSourceFromManagedInstance Removes a child software source from a managed instance. Packages will no longer be able to be
// installed from these software sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DetachChildSoftwareSourceFromManagedInstance.go.html to see an example of how to use DetachChildSoftwareSourceFromManagedInstance API.
func (client OsManagementClient) DetachChildSoftwareSourceFromManagedInstance(ctx context.Context, request DetachChildSoftwareSourceFromManagedInstanceRequest) (response DetachChildSoftwareSourceFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachChildSoftwareSourceFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachChildSoftwareSourceFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachChildSoftwareSourceFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachChildSoftwareSourceFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachChildSoftwareSourceFromManagedInstanceResponse")
	}
	return
}

// detachChildSoftwareSourceFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) detachChildSoftwareSourceFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/detachChildSoftwareSource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachChildSoftwareSourceFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/DetachChildSoftwareSourceFromManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "DetachChildSoftwareSourceFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachManagedInstanceFromManagedInstanceGroup Removes a Managed Instance from a Managed Instance Group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DetachManagedInstanceFromManagedInstanceGroup.go.html to see an example of how to use DetachManagedInstanceFromManagedInstanceGroup API.
func (client OsManagementClient) DetachManagedInstanceFromManagedInstanceGroup(ctx context.Context, request DetachManagedInstanceFromManagedInstanceGroupRequest) (response DetachManagedInstanceFromManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachManagedInstanceFromManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachManagedInstanceFromManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachManagedInstanceFromManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachManagedInstanceFromManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachManagedInstanceFromManagedInstanceGroupResponse")
	}
	return
}

// detachManagedInstanceFromManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) detachManagedInstanceFromManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/detachManagedInstance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachManagedInstanceFromManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/DetachManagedInstanceFromManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "DetachManagedInstanceFromManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachParentSoftwareSourceFromManagedInstance Removes a software source from a managed instance. All child software sources will also be removed
// from the managed instance. Packages will no longer be able to be installed from these software sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DetachParentSoftwareSourceFromManagedInstance.go.html to see an example of how to use DetachParentSoftwareSourceFromManagedInstance API.
func (client OsManagementClient) DetachParentSoftwareSourceFromManagedInstance(ctx context.Context, request DetachParentSoftwareSourceFromManagedInstanceRequest) (response DetachParentSoftwareSourceFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachParentSoftwareSourceFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachParentSoftwareSourceFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachParentSoftwareSourceFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachParentSoftwareSourceFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachParentSoftwareSourceFromManagedInstanceResponse")
	}
	return
}

// detachParentSoftwareSourceFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) detachParentSoftwareSourceFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/detachParentSoftwareSource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachParentSoftwareSourceFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/DetachParentSoftwareSourceFromManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "DetachParentSoftwareSourceFromManagedInstance", apiReferenceLink)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DisableModuleStreamOnManagedInstance.go.html to see an example of how to use DisableModuleStreamOnManagedInstance API.
func (client OsManagementClient) DisableModuleStreamOnManagedInstance(ctx context.Context, request DisableModuleStreamOnManagedInstanceRequest) (response DisableModuleStreamOnManagedInstanceResponse, err error) {
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
func (client OsManagementClient) disableModuleStreamOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/moduleStreams/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableModuleStreamOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamDetails/DisableModuleStreamOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "DisableModuleStreamOnManagedInstance", apiReferenceLink)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/EnableModuleStreamOnManagedInstance.go.html to see an example of how to use EnableModuleStreamOnManagedInstance API.
func (client OsManagementClient) EnableModuleStreamOnManagedInstance(ctx context.Context, request EnableModuleStreamOnManagedInstanceRequest) (response EnableModuleStreamOnManagedInstanceResponse, err error) {
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
func (client OsManagementClient) enableModuleStreamOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/moduleStreams/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableModuleStreamOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamDetails/EnableModuleStreamOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "EnableModuleStreamOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetErratum Returns a specific erratum.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetErratum.go.html to see an example of how to use GetErratum API.
func (client OsManagementClient) GetErratum(ctx context.Context, request GetErratumRequest) (response GetErratumResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getErratum(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/errata/{erratumId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetErratumResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/Erratum/GetErratum"
		err = common.PostProcessServiceError(err, "OsManagement", "GetErratum", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstance Returns a specific Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetManagedInstance.go.html to see an example of how to use GetManagedInstance API.
func (client OsManagementClient) GetManagedInstance(ctx context.Context, request GetManagedInstanceRequest) (response GetManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/GetManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "GetManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstanceGroup Returns a specific Managed Instance Group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetManagedInstanceGroup.go.html to see an example of how to use GetManagedInstanceGroup API.
func (client OsManagementClient) GetManagedInstanceGroup(ctx context.Context, request GetManagedInstanceGroupRequest) (response GetManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/GetManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "GetManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModuleStream Retrieve a detailed description of a module stream from a software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetModuleStream.go.html to see an example of how to use GetModuleStream API.
func (client OsManagementClient) GetModuleStream(ctx context.Context, request GetModuleStreamRequest) (response GetModuleStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getModuleStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/modules/{moduleName}/streams/{streamName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModuleStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStream/GetModuleStream"
		err = common.PostProcessServiceError(err, "OsManagement", "GetModuleStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModuleStreamProfile Retrieve a detailed description of a module stream profile from a software source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetModuleStreamProfile.go.html to see an example of how to use GetModuleStreamProfile API.
func (client OsManagementClient) GetModuleStreamProfile(ctx context.Context, request GetModuleStreamProfileRequest) (response GetModuleStreamProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getModuleStreamProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/modules/{moduleName}/streams/{streamName}/profiles/{profileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModuleStreamProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamProfile/GetModuleStreamProfile"
		err = common.PostProcessServiceError(err, "OsManagement", "GetModuleStreamProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledJob Gets the detailed information for the Scheduled Job with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetScheduledJob.go.html to see an example of how to use GetScheduledJob API.
func (client OsManagementClient) GetScheduledJob(ctx context.Context, request GetScheduledJobRequest) (response GetScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledJobResponse")
	}
	return
}

// getScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) getScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs/{scheduledJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/GetScheduledJob"
		err = common.PostProcessServiceError(err, "OsManagement", "GetScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwarePackage Returns a specific Software Package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetSoftwarePackage.go.html to see an example of how to use GetSoftwarePackage API.
func (client OsManagementClient) GetSoftwarePackage(ctx context.Context, request GetSoftwarePackageRequest) (response GetSoftwarePackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getSoftwarePackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/GetSoftwarePackage"
		err = common.PostProcessServiceError(err, "OsManagement", "GetSoftwarePackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwareSource Returns a specific Software Source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetSoftwareSource.go.html to see an example of how to use GetSoftwareSource API.
func (client OsManagementClient) GetSoftwareSource(ctx context.Context, request GetSoftwareSourceRequest) (response GetSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/GetSoftwareSource"
		err = common.PostProcessServiceError(err, "OsManagement", "GetSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWindowsUpdate Returns a Windows Update object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetWindowsUpdate.go.html to see an example of how to use GetWindowsUpdate API.
func (client OsManagementClient) GetWindowsUpdate(ctx context.Context, request GetWindowsUpdateRequest) (response GetWindowsUpdateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getWindowsUpdate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/updates/{windowsUpdate}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWindowsUpdateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/WindowsUpdate/GetWindowsUpdate"
		err = common.PostProcessServiceError(err, "OsManagement", "GetWindowsUpdate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the detailed information for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client OsManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "OsManagement", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAllPackageUpdatesOnManagedInstance Install all of the available package updates for the managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallAllPackageUpdatesOnManagedInstance.go.html to see an example of how to use InstallAllPackageUpdatesOnManagedInstance API.
func (client OsManagementClient) InstallAllPackageUpdatesOnManagedInstance(ctx context.Context, request InstallAllPackageUpdatesOnManagedInstanceRequest) (response InstallAllPackageUpdatesOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installAllPackageUpdatesOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallAllPackageUpdatesOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallAllPackageUpdatesOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallAllPackageUpdatesOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallAllPackageUpdatesOnManagedInstanceResponse")
	}
	return
}

// installAllPackageUpdatesOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) installAllPackageUpdatesOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/updateAll", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallAllPackageUpdatesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/InstallAllPackageUpdatesOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallAllPackageUpdatesOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAllUpdatesOnManagedInstanceGroup Install all of the available updates for the Managed Instance Group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallAllUpdatesOnManagedInstanceGroup.go.html to see an example of how to use InstallAllUpdatesOnManagedInstanceGroup API.
func (client OsManagementClient) InstallAllUpdatesOnManagedInstanceGroup(ctx context.Context, request InstallAllUpdatesOnManagedInstanceGroupRequest) (response InstallAllUpdatesOnManagedInstanceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installAllUpdatesOnManagedInstanceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallAllUpdatesOnManagedInstanceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallAllUpdatesOnManagedInstanceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallAllUpdatesOnManagedInstanceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallAllUpdatesOnManagedInstanceGroupResponse")
	}
	return
}

// installAllUpdatesOnManagedInstanceGroup implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) installAllUpdatesOnManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/updates/installAll", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallAllUpdatesOnManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/InstallAllUpdatesOnManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallAllUpdatesOnManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAllWindowsUpdatesOnManagedInstance Install all of the available Windows updates for the managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallAllWindowsUpdatesOnManagedInstance.go.html to see an example of how to use InstallAllWindowsUpdatesOnManagedInstance API.
func (client OsManagementClient) InstallAllWindowsUpdatesOnManagedInstance(ctx context.Context, request InstallAllWindowsUpdatesOnManagedInstanceRequest) (response InstallAllWindowsUpdatesOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installAllWindowsUpdatesOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallAllWindowsUpdatesOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallAllWindowsUpdatesOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallAllWindowsUpdatesOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallAllWindowsUpdatesOnManagedInstanceResponse")
	}
	return
}

// installAllWindowsUpdatesOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) installAllWindowsUpdatesOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/updates/installAll", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallAllWindowsUpdatesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/InstallAllWindowsUpdatesOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallAllWindowsUpdatesOnManagedInstance", apiReferenceLink)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallModuleStreamProfileOnManagedInstance.go.html to see an example of how to use InstallModuleStreamProfileOnManagedInstance API.
func (client OsManagementClient) InstallModuleStreamProfileOnManagedInstance(ctx context.Context, request InstallModuleStreamProfileOnManagedInstanceRequest) (response InstallModuleStreamProfileOnManagedInstanceResponse, err error) {
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
func (client OsManagementClient) installModuleStreamProfileOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/streamProfiles/install", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallModuleStreamProfileOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamProfileDetails/InstallModuleStreamProfileOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallModuleStreamProfileOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackageOnManagedInstance Installs a package on a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallPackageOnManagedInstance.go.html to see an example of how to use InstallPackageOnManagedInstance API.
func (client OsManagementClient) InstallPackageOnManagedInstance(ctx context.Context, request InstallPackageOnManagedInstanceRequest) (response InstallPackageOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installPackageOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallPackageOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallPackageOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallPackageOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallPackageOnManagedInstanceResponse")
	}
	return
}

// installPackageOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) installPackageOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/install", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallPackageOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/InstallPackageOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallPackageOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackageUpdateOnManagedInstance Updates a package on a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallPackageUpdateOnManagedInstance.go.html to see an example of how to use InstallPackageUpdateOnManagedInstance API.
func (client OsManagementClient) InstallPackageUpdateOnManagedInstance(ctx context.Context, request InstallPackageUpdateOnManagedInstanceRequest) (response InstallPackageUpdateOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installPackageUpdateOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallPackageUpdateOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallPackageUpdateOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallPackageUpdateOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallPackageUpdateOnManagedInstanceResponse")
	}
	return
}

// installPackageUpdateOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) installPackageUpdateOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/update", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallPackageUpdateOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/InstallPackageUpdateOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallPackageUpdateOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallWindowsUpdateOnManagedInstance Installs a Windows update on a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallWindowsUpdateOnManagedInstance.go.html to see an example of how to use InstallWindowsUpdateOnManagedInstance API.
func (client OsManagementClient) InstallWindowsUpdateOnManagedInstance(ctx context.Context, request InstallWindowsUpdateOnManagedInstanceRequest) (response InstallWindowsUpdateOnManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installWindowsUpdateOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallWindowsUpdateOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallWindowsUpdateOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallWindowsUpdateOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallWindowsUpdateOnManagedInstanceResponse")
	}
	return
}

// installWindowsUpdateOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) installWindowsUpdateOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/updates/install", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallWindowsUpdateOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/InstallWindowsUpdateOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "InstallWindowsUpdateOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailablePackagesForManagedInstance Returns a list of packages available for install on the Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailablePackagesForManagedInstance.go.html to see an example of how to use ListAvailablePackagesForManagedInstance API.
func (client OsManagementClient) ListAvailablePackagesForManagedInstance(ctx context.Context, request ListAvailablePackagesForManagedInstanceRequest) (response ListAvailablePackagesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailablePackagesForManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailablePackagesForManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailablePackagesForManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailablePackagesForManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailablePackagesForManagedInstanceResponse")
	}
	return
}

// listAvailablePackagesForManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listAvailablePackagesForManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/packages/available", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailablePackagesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListAvailablePackagesForManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListAvailablePackagesForManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableSoftwareSourcesForManagedInstance Returns a list of available software sources for a Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailableSoftwareSourcesForManagedInstance.go.html to see an example of how to use ListAvailableSoftwareSourcesForManagedInstance API.
func (client OsManagementClient) ListAvailableSoftwareSourcesForManagedInstance(ctx context.Context, request ListAvailableSoftwareSourcesForManagedInstanceRequest) (response ListAvailableSoftwareSourcesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailableSoftwareSourcesForManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailableSoftwareSourcesForManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailableSoftwareSourcesForManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailableSoftwareSourcesForManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailableSoftwareSourcesForManagedInstanceResponse")
	}
	return
}

// listAvailableSoftwareSourcesForManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listAvailableSoftwareSourcesForManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/availableSoftwareSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailableSoftwareSourcesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListAvailableSoftwareSourcesForManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListAvailableSoftwareSourcesForManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableUpdatesForManagedInstance Returns a list of available updates for a Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailableUpdatesForManagedInstance.go.html to see an example of how to use ListAvailableUpdatesForManagedInstance API.
func (client OsManagementClient) ListAvailableUpdatesForManagedInstance(ctx context.Context, request ListAvailableUpdatesForManagedInstanceRequest) (response ListAvailableUpdatesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailableUpdatesForManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailableUpdatesForManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailableUpdatesForManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailableUpdatesForManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailableUpdatesForManagedInstanceResponse")
	}
	return
}

// listAvailableUpdatesForManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listAvailableUpdatesForManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/packages/updates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailableUpdatesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListAvailableUpdatesForManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListAvailableUpdatesForManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableWindowsUpdatesForManagedInstance Returns a list of available Windows updates for a Managed Instance. This is only applicable to Windows instances.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailableWindowsUpdatesForManagedInstance.go.html to see an example of how to use ListAvailableWindowsUpdatesForManagedInstance API.
func (client OsManagementClient) ListAvailableWindowsUpdatesForManagedInstance(ctx context.Context, request ListAvailableWindowsUpdatesForManagedInstanceRequest) (response ListAvailableWindowsUpdatesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailableWindowsUpdatesForManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailableWindowsUpdatesForManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailableWindowsUpdatesForManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailableWindowsUpdatesForManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailableWindowsUpdatesForManagedInstanceResponse")
	}
	return
}

// listAvailableWindowsUpdatesForManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listAvailableWindowsUpdatesForManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/updates/available", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailableWindowsUpdatesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListAvailableWindowsUpdatesForManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListAvailableWindowsUpdatesForManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListErrata Returns a list of all of the currently available Errata in the system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListErrata.go.html to see an example of how to use ListErrata API.
func (client OsManagementClient) ListErrata(ctx context.Context, request ListErrataRequest) (response ListErrataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listErrata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ErratumSummary/ListErrata"
		err = common.PostProcessServiceError(err, "OsManagement", "ListErrata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceErrata Returns a list of errata relevant to the Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListManagedInstanceErrata.go.html to see an example of how to use ListManagedInstanceErrata API.
func (client OsManagementClient) ListManagedInstanceErrata(ctx context.Context, request ListManagedInstanceErrataRequest) (response ListManagedInstanceErrataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listManagedInstanceErrata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListManagedInstanceErrata"
		err = common.PostProcessServiceError(err, "OsManagement", "ListManagedInstanceErrata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroups Returns a list of all Managed Instance Groups.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListManagedInstanceGroups.go.html to see an example of how to use ListManagedInstanceGroups API.
func (client OsManagementClient) ListManagedInstanceGroups(ctx context.Context, request ListManagedInstanceGroupsRequest) (response ListManagedInstanceGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listManagedInstanceGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroupSummary/ListManagedInstanceGroups"
		err = common.PostProcessServiceError(err, "OsManagement", "ListManagedInstanceGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstances Returns a list of all Managed Instances.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListManagedInstances.go.html to see an example of how to use ListManagedInstances API.
func (client OsManagementClient) ListManagedInstances(ctx context.Context, request ListManagedInstancesRequest) (response ListManagedInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listManagedInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceSummary/ListManagedInstances"
		err = common.PostProcessServiceError(err, "OsManagement", "ListManagedInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModuleStreamProfiles Retrieve a list of module stream profiles from a software source.
// Filters may be applied to select a subset of module stream profiles
// based on the filter criteria.
// The "moduleName", "streamName", and "profileName" attributes combine
// to form a set of filters on the list of module stream profiles.  If
// a "moduleName" is provided, only profiles that belong to that module
// are returned.  If both a "moduleName" and "streamName" are given,
// only profiles belonging to that module stream are returned.  Finally,
// if all three are given then only the particular profile indicated
// by the triple is returned.  It is not valid to supply a "streamName"
// without a "moduleName".  It is also not valid to supply a "profileName"
// without a "streamName".
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreamProfiles.go.html to see an example of how to use ListModuleStreamProfiles API.
func (client OsManagementClient) ListModuleStreamProfiles(ctx context.Context, request ListModuleStreamProfilesRequest) (response ListModuleStreamProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listModuleStreamProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/streamProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModuleStreamProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamProfileSummary/ListModuleStreamProfiles"
		err = common.PostProcessServiceError(err, "OsManagement", "ListModuleStreamProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModuleStreamProfilesOnManagedInstance Retrieve a list of module stream profiles, along with a summary of their
// of their status, from a managed instance.  Filters may be applied to
// select a subset of profiles based on the filter criteria.
// The "moduleName", "streamName", and "profileName" attributes combine
// to form a set of filters on the list of module stream profiles.  If
// a "modulName" is provided, only profiles that belong to that module
// are returned.  If both a "moduleName" and "streamName" are given,
// only profiles belonging to that module stream are returned.  Finally,
// if all three are given then only the particular profile indicated
// by the triple is returned.  It is not valid to supply a "streamName"
// without a "moduleName".  It is also not valid to supply a "profileName"
// without a "streamName".
// The "status" attribute filters against the state of a module stream
// profile.  Valid values are "INSTALLED" and "AVAILABLE".  If the
// attribute is set to "INSTALLED", only module stream profiles that
// are installed are included in the result set.  If the attribute is
// set to "AVAILABLE", only module stream profiles that are not
// installed are included in the result set.  If the attribute is not
// defined, the request is not subject to this filter.
// When sorting by display name, the result set is sorted first by
// module name, then by stream name, and finally by profile name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreamProfilesOnManagedInstance.go.html to see an example of how to use ListModuleStreamProfilesOnManagedInstance API.
func (client OsManagementClient) ListModuleStreamProfilesOnManagedInstance(ctx context.Context, request ListModuleStreamProfilesOnManagedInstanceRequest) (response ListModuleStreamProfilesOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModuleStreamProfilesOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModuleStreamProfilesOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModuleStreamProfilesOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModuleStreamProfilesOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModuleStreamProfilesOnManagedInstanceResponse")
	}
	return
}

// listModuleStreamProfilesOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listModuleStreamProfilesOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/streamProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModuleStreamProfilesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListModuleStreamProfilesOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListModuleStreamProfilesOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModuleStreams Retrieve a list of module streams from a software source.
// Filters may be applied to select a subset of module streams
// based on the filter criteria.
// The 'moduleName' attribute filters against the name of a module.
// It accepts strings of the format "<module>".  If this attribute
// is defined, only streams that belong to the specified module are
// included in the result set.  If it is not defined, the request is
// not subject to this filter.  The 'streamName' attribute filters
// against the name of a stream of a module.  If this attribute is
// defined, only the particular module stream that matches both the
// module and stream names is included in the result set.  It is
// not valid to supply 'streamName' without also supplying a
// 'moduleName'.
// When sorting by display name, the result set is sorted first by
// module name, then by stream name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreams.go.html to see an example of how to use ListModuleStreams API.
func (client OsManagementClient) ListModuleStreams(ctx context.Context, request ListModuleStreamsRequest) (response ListModuleStreamsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listModuleStreams(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamSummary/ListModuleStreams"
		err = common.PostProcessServiceError(err, "OsManagement", "ListModuleStreams", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModuleStreamsOnManagedInstance Retrieve a list of module streams, along with a summary of their
// status, from a managed instance.  Filters may be applied to select
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreamsOnManagedInstance.go.html to see an example of how to use ListModuleStreamsOnManagedInstance API.
func (client OsManagementClient) ListModuleStreamsOnManagedInstance(ctx context.Context, request ListModuleStreamsOnManagedInstanceRequest) (response ListModuleStreamsOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModuleStreamsOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModuleStreamsOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModuleStreamsOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModuleStreamsOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModuleStreamsOnManagedInstanceResponse")
	}
	return
}

// listModuleStreamsOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listModuleStreamsOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/moduleStreams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModuleStreamsOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListModuleStreamsOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListModuleStreamsOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPackagesInstalledOnManagedInstance Returns a list of installed packages on the Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListPackagesInstalledOnManagedInstance.go.html to see an example of how to use ListPackagesInstalledOnManagedInstance API.
func (client OsManagementClient) ListPackagesInstalledOnManagedInstance(ctx context.Context, request ListPackagesInstalledOnManagedInstanceRequest) (response ListPackagesInstalledOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPackagesInstalledOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPackagesInstalledOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPackagesInstalledOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPackagesInstalledOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPackagesInstalledOnManagedInstanceResponse")
	}
	return
}

// listPackagesInstalledOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listPackagesInstalledOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/packages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPackagesInstalledOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListPackagesInstalledOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListPackagesInstalledOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledJobs Returns a list of all of the currently active Scheduled Jobs in the system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListScheduledJobs.go.html to see an example of how to use ListScheduledJobs API.
func (client OsManagementClient) ListScheduledJobs(ctx context.Context, request ListScheduledJobsRequest) (response ListScheduledJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledJobsResponse")
	}
	return
}

// listScheduledJobs implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listScheduledJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/ListScheduledJobs"
		err = common.PostProcessServiceError(err, "OsManagement", "ListScheduledJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareSourcePackages Lists Software Packages in a Software Source
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListSoftwareSourcePackages.go.html to see an example of how to use ListSoftwareSourcePackages API.
func (client OsManagementClient) ListSoftwareSourcePackages(ctx context.Context, request ListSoftwareSourcePackagesRequest) (response ListSoftwareSourcePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSoftwareSourcePackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSoftwareSourcePackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSoftwareSourcePackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSoftwareSourcePackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSoftwareSourcePackagesResponse")
	}
	return
}

// listSoftwareSourcePackages implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listSoftwareSourcePackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/softwarePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSoftwareSourcePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/ListSoftwareSourcePackages"
		err = common.PostProcessServiceError(err, "OsManagement", "ListSoftwareSourcePackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareSources Returns a list of all Software Sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListSoftwareSources.go.html to see an example of how to use ListSoftwareSources API.
func (client OsManagementClient) ListSoftwareSources(ctx context.Context, request ListSoftwareSourcesRequest) (response ListSoftwareSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listSoftwareSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSourceSummary/ListSoftwareSources"
		err = common.PostProcessServiceError(err, "OsManagement", "ListSoftwareSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUpcomingScheduledJobs Returns a list of all of the Scheduled Jobs whose next execution time is at or before the specified time.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListUpcomingScheduledJobs.go.html to see an example of how to use ListUpcomingScheduledJobs API.
func (client OsManagementClient) ListUpcomingScheduledJobs(ctx context.Context, request ListUpcomingScheduledJobsRequest) (response ListUpcomingScheduledJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUpcomingScheduledJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUpcomingScheduledJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUpcomingScheduledJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUpcomingScheduledJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUpcomingScheduledJobsResponse")
	}
	return
}

// listUpcomingScheduledJobs implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listUpcomingScheduledJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs/upcomingSchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUpcomingScheduledJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/ListUpcomingScheduledJobs"
		err = common.PostProcessServiceError(err, "OsManagement", "ListUpcomingScheduledJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWindowsUpdates Returns a list of Windows Updates.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWindowsUpdates.go.html to see an example of how to use ListWindowsUpdates API.
func (client OsManagementClient) ListWindowsUpdates(ctx context.Context, request ListWindowsUpdatesRequest) (response ListWindowsUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWindowsUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/updates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWindowsUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/WindowsUpdateSummary/ListWindowsUpdates"
		err = common.PostProcessServiceError(err, "OsManagement", "ListWindowsUpdates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWindowsUpdatesInstalledOnManagedInstance Returns a list of installed Windows updates for a Managed Instance. This is only applicable to Windows instances.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWindowsUpdatesInstalledOnManagedInstance.go.html to see an example of how to use ListWindowsUpdatesInstalledOnManagedInstance API.
func (client OsManagementClient) ListWindowsUpdatesInstalledOnManagedInstance(ctx context.Context, request ListWindowsUpdatesInstalledOnManagedInstanceRequest) (response ListWindowsUpdatesInstalledOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWindowsUpdatesInstalledOnManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWindowsUpdatesInstalledOnManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWindowsUpdatesInstalledOnManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWindowsUpdatesInstalledOnManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWindowsUpdatesInstalledOnManagedInstanceResponse")
	}
	return
}

// listWindowsUpdatesInstalledOnManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) listWindowsUpdatesInstalledOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/updates/installed", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWindowsUpdatesInstalledOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ListWindowsUpdatesInstalledOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ListWindowsUpdatesInstalledOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Gets the errors for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client OsManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/WorkRequest/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "OsManagement", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the log entries for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client OsManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/WorkRequest/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "OsManagement", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client OsManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "OsManagement", "ListWorkRequests", apiReferenceLink)
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
// The result of this request is a WorkRequest object.  The returned
// WorkRequest is the parent of a structure of other WorkRequests.  Taken
// as a whole, this structure indicates the entire set of work to be
// performed to complete the operation.
// This interface can also be used to perform a dry run of the operation
// rather than committing it to a managed instance.  If a dry run is
// requested, the OS Management Service will evaluate the operation
// against the current module, stream, and profile state on the managed
// instance.  It will calculate the impact of the operation on all
// modules, streams, and profiles on the managed instance, including those
// that are implicitly impacted by the operation.
// The WorkRequest resulting from a dry run behaves differently than
// a WorkRequest resulting from a committable operation.  Dry run
// WorkRequests are always singletons and never have children.  The
// impact of the operation is returned using the log and error
// facilities of WorkRequests.  The impact of operations that are
// allowed by the OS Management Service are communicated as one or
// more work request log entries.  Operations that are not allowed
// by the OS Management Service are communicated as one or more
// work requst error entries.  Each entry, for either logs or errors,
// contains a structured message containing the results of one
// or more operations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ManageModuleStreamsOnManagedInstance.go.html to see an example of how to use ManageModuleStreamsOnManagedInstance API.
func (client OsManagementClient) ManageModuleStreamsOnManagedInstance(ctx context.Context, request ManageModuleStreamsOnManagedInstanceRequest) (response ManageModuleStreamsOnManagedInstanceResponse, err error) {
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
func (client OsManagementClient) manageModuleStreamsOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/moduleStreams/manage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ManageModuleStreamsOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/ManageModuleStreamsOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "ManageModuleStreamsOnManagedInstance", apiReferenceLink)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/RemoveModuleStreamProfileFromManagedInstance.go.html to see an example of how to use RemoveModuleStreamProfileFromManagedInstance API.
func (client OsManagementClient) RemoveModuleStreamProfileFromManagedInstance(ctx context.Context, request RemoveModuleStreamProfileFromManagedInstanceRequest) (response RemoveModuleStreamProfileFromManagedInstanceResponse, err error) {
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
func (client OsManagementClient) removeModuleStreamProfileFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/streamProfiles/remove", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveModuleStreamProfileFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamProfileDetails/RemoveModuleStreamProfileFromManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "RemoveModuleStreamProfileFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackageFromManagedInstance Removes an installed package from a managed instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/RemovePackageFromManagedInstance.go.html to see an example of how to use RemovePackageFromManagedInstance API.
func (client OsManagementClient) RemovePackageFromManagedInstance(ctx context.Context, request RemovePackageFromManagedInstanceRequest) (response RemovePackageFromManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removePackageFromManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePackageFromManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePackageFromManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePackageFromManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePackageFromManagedInstanceResponse")
	}
	return
}

// removePackageFromManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) removePackageFromManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/remove", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePackageFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstance/RemovePackageFromManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "RemovePackageFromManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackagesFromSoftwareSource Removes a given list of Software Packages from a specific Software Source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/RemovePackagesFromSoftwareSource.go.html to see an example of how to use RemovePackagesFromSoftwareSource API.
func (client OsManagementClient) RemovePackagesFromSoftwareSource(ctx context.Context, request RemovePackagesFromSoftwareSourceRequest) (response RemovePackagesFromSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removePackagesFromSoftwareSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePackagesFromSoftwareSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePackagesFromSoftwareSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePackagesFromSoftwareSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePackagesFromSoftwareSourceResponse")
	}
	return
}

// removePackagesFromSoftwareSource implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) removePackagesFromSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/{softwareSourceId}/actions/removePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePackagesFromSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/RemovePackagesFromSoftwareSource"
		err = common.PostProcessServiceError(err, "OsManagement", "RemovePackagesFromSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunScheduledJobNow This will trigger an already created Scheduled Job to being executing
// immediately instead of waiting for its next regularly scheduled time.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/RunScheduledJobNow.go.html to see an example of how to use RunScheduledJobNow API.
func (client OsManagementClient) RunScheduledJobNow(ctx context.Context, request RunScheduledJobNowRequest) (response RunScheduledJobNowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.runScheduledJobNow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunScheduledJobNowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunScheduledJobNowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunScheduledJobNowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunScheduledJobNowResponse")
	}
	return
}

// runScheduledJobNow implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) runScheduledJobNow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/runNow", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunScheduledJobNowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/RunScheduledJobNow"
		err = common.PostProcessServiceError(err, "OsManagement", "RunScheduledJobNow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSoftwarePackages Searches all of the available Software Sources and returns any/all Software Packages matching
// the search criteria.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/SearchSoftwarePackages.go.html to see an example of how to use SearchSoftwarePackages API.
func (client OsManagementClient) SearchSoftwarePackages(ctx context.Context, request SearchSoftwarePackagesRequest) (response SearchSoftwarePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchSoftwarePackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSoftwarePackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSoftwarePackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSoftwarePackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSoftwarePackagesResponse")
	}
	return
}

// searchSoftwarePackages implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) searchSoftwarePackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/softwarePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSoftwarePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/SearchSoftwarePackages"
		err = common.PostProcessServiceError(err, "OsManagement", "SearchSoftwarePackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SkipNextScheduledJobExecution This will force an already created Scheduled Job to skip its
// next regularly scheduled execution
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/SkipNextScheduledJobExecution.go.html to see an example of how to use SkipNextScheduledJobExecution API.
func (client OsManagementClient) SkipNextScheduledJobExecution(ctx context.Context, request SkipNextScheduledJobExecutionRequest) (response SkipNextScheduledJobExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.skipNextScheduledJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SkipNextScheduledJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SkipNextScheduledJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SkipNextScheduledJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SkipNextScheduledJobExecutionResponse")
	}
	return
}

// skipNextScheduledJobExecution implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) skipNextScheduledJobExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/skipNextExecution", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SkipNextScheduledJobExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/SkipNextScheduledJobExecution"
		err = common.PostProcessServiceError(err, "OsManagement", "SkipNextScheduledJobExecution", apiReferenceLink)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/SwitchModuleStreamOnManagedInstance.go.html to see an example of how to use SwitchModuleStreamOnManagedInstance API.
func (client OsManagementClient) SwitchModuleStreamOnManagedInstance(ctx context.Context, request SwitchModuleStreamOnManagedInstanceRequest) (response SwitchModuleStreamOnManagedInstanceResponse, err error) {
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
func (client OsManagementClient) switchModuleStreamOnManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/moduleStreams/switch", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchModuleStreamOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ModuleStreamDetails/SwitchModuleStreamOnManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "SwitchModuleStreamOnManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedInstance Updates a specific Managed Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/UpdateManagedInstance.go.html to see an example of how to use UpdateManagedInstance API.
func (client OsManagementClient) UpdateManagedInstance(ctx context.Context, request UpdateManagedInstanceRequest) (response UpdateManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) updateManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/UpdateManagedInstanceDetails/UpdateManagedInstance"
		err = common.PostProcessServiceError(err, "OsManagement", "UpdateManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedInstanceGroup Updates a specific Managed Instance Group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/UpdateManagedInstanceGroup.go.html to see an example of how to use UpdateManagedInstanceGroup API.
func (client OsManagementClient) UpdateManagedInstanceGroup(ctx context.Context, request UpdateManagedInstanceGroupRequest) (response UpdateManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) updateManagedInstanceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ManagedInstanceGroup/UpdateManagedInstanceGroup"
		err = common.PostProcessServiceError(err, "OsManagement", "UpdateManagedInstanceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateScheduledJob Updates an existing Scheduled Job on the management system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/UpdateScheduledJob.go.html to see an example of how to use UpdateScheduledJob API.
func (client OsManagementClient) UpdateScheduledJob(ctx context.Context, request UpdateScheduledJobRequest) (response UpdateScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduledJobResponse")
	}
	return
}

// updateScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client OsManagementClient) updateScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/scheduledJobs/{scheduledJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/ScheduledJob/UpdateScheduledJob"
		err = common.PostProcessServiceError(err, "OsManagement", "UpdateScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSoftwareSource Updates an existing custom Software Source on the management system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/UpdateSoftwareSource.go.html to see an example of how to use UpdateSoftwareSource API.
func (client OsManagementClient) UpdateSoftwareSource(ctx context.Context, request UpdateSoftwareSourceRequest) (response UpdateSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) updateSoftwareSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/os-management/20190801/SoftwareSource/UpdateSoftwareSource"
		err = common.PostProcessServiceError(err, "OsManagement", "UpdateSoftwareSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
