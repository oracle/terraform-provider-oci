// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//OsManagementClient a client for OsManagement
type OsManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOsManagementClientWithConfigurationProvider Creates a new default OsManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOsManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OsManagementClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newOsManagementClientFromBaseClient(baseClient, configProvider)
}

// NewOsManagementClientWithOboToken Creates a new default OsManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOsManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OsManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newOsManagementClientFromBaseClient(baseClient, configProvider)
}

func newOsManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OsManagementClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *OsManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddPackagesToSoftwareSource Adds a given list of Software Packages to a specific Software Source.
func (client OsManagementClient) AddPackagesToSoftwareSource(ctx context.Context, request AddPackagesToSoftwareSourceRequest) (response AddPackagesToSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) addPackagesToSoftwareSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/{softwareSourceId}/actions/addPackages")
	if err != nil {
		return nil, err
	}

	var response AddPackagesToSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachChildSoftwareSourceToManagedInstance Adds a child software source to a managed instance. After the software
// source has been added, then packages from that software source can be
// installed on the managed instance.
func (client OsManagementClient) AttachChildSoftwareSourceToManagedInstance(ctx context.Context, request AttachChildSoftwareSourceToManagedInstanceRequest) (response AttachChildSoftwareSourceToManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) attachChildSoftwareSourceToManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/attachChildSoftwareSource")
	if err != nil {
		return nil, err
	}

	var response AttachChildSoftwareSourceToManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachManagedInstanceToManagedInstanceGroup Adds a Managed Instance to a Managed Instance Group. After the Managed
// Instance has been added, then operations can be performed on the Managed
// Instance Group which will then apply to all Managed Instances in the
// group.
func (client OsManagementClient) AttachManagedInstanceToManagedInstanceGroup(ctx context.Context, request AttachManagedInstanceToManagedInstanceGroupRequest) (response AttachManagedInstanceToManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) attachManagedInstanceToManagedInstanceGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/attachManagedInstance")
	if err != nil {
		return nil, err
	}

	var response AttachManagedInstanceToManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachParentSoftwareSourceToManagedInstance Adds a parent software source to a managed instance. After the software
// source has been added, then packages from that software source can be
// installed on the managed instance. Software sources that have this
// software source as a parent will be able to be added to this managed instance.
func (client OsManagementClient) AttachParentSoftwareSourceToManagedInstance(ctx context.Context, request AttachParentSoftwareSourceToManagedInstanceRequest) (response AttachParentSoftwareSourceToManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) attachParentSoftwareSourceToManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/attachParentSoftwareSource")
	if err != nil {
		return nil, err
	}

	var response AttachParentSoftwareSourceToManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeManagedInstanceGroupCompartment Moves a resource into a different compartment. When provided, If-Match
// is checked against ETag values of the resource.
func (client OsManagementClient) ChangeManagedInstanceGroupCompartment(ctx context.Context, request ChangeManagedInstanceGroupCompartmentRequest) (response ChangeManagedInstanceGroupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) changeManagedInstanceGroupCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeManagedInstanceGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeScheduledJobCompartment Moves a resource into a different compartment. When provided, If-Match
// is checked against ETag values of the resource.
func (client OsManagementClient) ChangeScheduledJobCompartment(ctx context.Context, request ChangeScheduledJobCompartmentRequest) (response ChangeScheduledJobCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) changeScheduledJobCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeScheduledJobCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSoftwareSourceCompartment Moves a resource into a different compartment. When provided, If-Match
// is checked against ETag values of the resource.
func (client OsManagementClient) ChangeSoftwareSourceCompartment(ctx context.Context, request ChangeSoftwareSourceCompartmentRequest) (response ChangeSoftwareSourceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) changeSoftwareSourceCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/{softwareSourceId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeSoftwareSourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateManagedInstanceGroup Creates a new Managed Instance Group on the management system.
// This will not contain any managed instances after it is first created,
// and they must be added later.
func (client OsManagementClient) CreateManagedInstanceGroup(ctx context.Context, request CreateManagedInstanceGroupRequest) (response CreateManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) createManagedInstanceGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups")
	if err != nil {
		return nil, err
	}

	var response CreateManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateScheduledJob Creates a new Scheduled Job to perform a specific package operation on
// a set of managed instances or managed instance groups.  Can be created
// as a one-time execution in the future, or as a recurring execution
// that repeats on a defined interval.
func (client OsManagementClient) CreateScheduledJob(ctx context.Context, request CreateScheduledJobRequest) (response CreateScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) createScheduledJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs")
	if err != nil {
		return nil, err
	}

	var response CreateScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSoftwareSource Creates a new custom Software Source on the management system.
// This will not contain any packages after it is first created,
// and they must be added later.
func (client OsManagementClient) CreateSoftwareSource(ctx context.Context, request CreateSoftwareSourceRequest) (response CreateSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) createSoftwareSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources")
	if err != nil {
		return nil, err
	}

	var response CreateSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedInstanceGroup Deletes a Managed Instance Group from the management system
func (client OsManagementClient) DeleteManagedInstanceGroup(ctx context.Context, request DeleteManagedInstanceGroupRequest) (response DeleteManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) deleteManagedInstanceGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedInstanceGroups/{managedInstanceGroupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteScheduledJob Cancels an existing Scheduled Job on the management system
func (client OsManagementClient) DeleteScheduledJob(ctx context.Context, request DeleteScheduledJobRequest) (response DeleteScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) deleteScheduledJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/scheduledJobs/{scheduledJobId}")
	if err != nil {
		return nil, err
	}

	var response DeleteScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSoftwareSource Deletes a custom Software Source on the management system
func (client OsManagementClient) DeleteSoftwareSource(ctx context.Context, request DeleteSoftwareSourceRequest) (response DeleteSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) deleteSoftwareSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/softwareSources/{softwareSourceId}")
	if err != nil {
		return nil, err
	}

	var response DeleteSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachChildSoftwareSourceFromManagedInstance Removes a child software source from a managed instance. Packages will no longer be able to be
// installed from these software sources.
func (client OsManagementClient) DetachChildSoftwareSourceFromManagedInstance(ctx context.Context, request DetachChildSoftwareSourceFromManagedInstanceRequest) (response DetachChildSoftwareSourceFromManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) detachChildSoftwareSourceFromManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/detachChildSoftwareSource")
	if err != nil {
		return nil, err
	}

	var response DetachChildSoftwareSourceFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachManagedInstanceFromManagedInstanceGroup Removes a Managed Instance from a Managed Instance Group.
func (client OsManagementClient) DetachManagedInstanceFromManagedInstanceGroup(ctx context.Context, request DetachManagedInstanceFromManagedInstanceGroupRequest) (response DetachManagedInstanceFromManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) detachManagedInstanceFromManagedInstanceGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstanceGroups/{managedInstanceGroupId}/actions/detachManagedInstance")
	if err != nil {
		return nil, err
	}

	var response DetachManagedInstanceFromManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachParentSoftwareSourceFromManagedInstance Removes a software source from a managed instance. All child software sources will also be removed
// from the managed instance. Packages will no longer be able to be installed from these software sources.
func (client OsManagementClient) DetachParentSoftwareSourceFromManagedInstance(ctx context.Context, request DetachParentSoftwareSourceFromManagedInstanceRequest) (response DetachParentSoftwareSourceFromManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) detachParentSoftwareSourceFromManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/detachParentSoftwareSource")
	if err != nil {
		return nil, err
	}

	var response DetachParentSoftwareSourceFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetErratum Returns a specific erratum.
func (client OsManagementClient) GetErratum(ctx context.Context, request GetErratumRequest) (response GetErratumResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getErratum(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/errata/{erratumId}")
	if err != nil {
		return nil, err
	}

	var response GetErratumResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstance Returns a specific Managed Instance.
func (client OsManagementClient) GetManagedInstance(ctx context.Context, request GetManagedInstanceRequest) (response GetManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}")
	if err != nil {
		return nil, err
	}

	var response GetManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstanceGroup Returns a specific Managed Instance Group.
func (client OsManagementClient) GetManagedInstanceGroup(ctx context.Context, request GetManagedInstanceGroupRequest) (response GetManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getManagedInstanceGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups/{managedInstanceGroupId}")
	if err != nil {
		return nil, err
	}

	var response GetManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledJob Gets the detailed information for the Scheduled Job with the given ID.
func (client OsManagementClient) GetScheduledJob(ctx context.Context, request GetScheduledJobRequest) (response GetScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getScheduledJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs/{scheduledJobId}")
	if err != nil {
		return nil, err
	}

	var response GetScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwarePackage Returns a specific Software Package.
func (client OsManagementClient) GetSoftwarePackage(ctx context.Context, request GetSoftwarePackageRequest) (response GetSoftwarePackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getSoftwarePackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/softwarePackages/{softwarePackageName}")
	if err != nil {
		return nil, err
	}

	var response GetSoftwarePackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwareSource Returns a specific Software Source.
func (client OsManagementClient) GetSoftwareSource(ctx context.Context, request GetSoftwareSourceRequest) (response GetSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getSoftwareSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}")
	if err != nil {
		return nil, err
	}

	var response GetSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWindowsUpdate Returns a Windows Update object.
func (client OsManagementClient) GetWindowsUpdate(ctx context.Context, request GetWindowsUpdateRequest) (response GetWindowsUpdateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getWindowsUpdate(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/updates/{windowsUpdate}")
	if err != nil {
		return nil, err
	}

	var response GetWindowsUpdateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the detailed information for the work request with the given ID.
func (client OsManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAllPackageUpdatesOnManagedInstance Install all of the available package updates for the managed instance.
func (client OsManagementClient) InstallAllPackageUpdatesOnManagedInstance(ctx context.Context, request InstallAllPackageUpdatesOnManagedInstanceRequest) (response InstallAllPackageUpdatesOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) installAllPackageUpdatesOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/updateAll")
	if err != nil {
		return nil, err
	}

	var response InstallAllPackageUpdatesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAllWindowsUpdatesOnManagedInstance Install all of the available Windows updates for the managed instance.
func (client OsManagementClient) InstallAllWindowsUpdatesOnManagedInstance(ctx context.Context, request InstallAllWindowsUpdatesOnManagedInstanceRequest) (response InstallAllWindowsUpdatesOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) installAllWindowsUpdatesOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/updates/installAll")
	if err != nil {
		return nil, err
	}

	var response InstallAllWindowsUpdatesOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackageOnManagedInstance Installs a package on a managed instance.
func (client OsManagementClient) InstallPackageOnManagedInstance(ctx context.Context, request InstallPackageOnManagedInstanceRequest) (response InstallPackageOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) installPackageOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/install")
	if err != nil {
		return nil, err
	}

	var response InstallPackageOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackageUpdateOnManagedInstance Updates a package on a managed instance.
func (client OsManagementClient) InstallPackageUpdateOnManagedInstance(ctx context.Context, request InstallPackageUpdateOnManagedInstanceRequest) (response InstallPackageUpdateOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) installPackageUpdateOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/update")
	if err != nil {
		return nil, err
	}

	var response InstallPackageUpdateOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallWindowsUpdateOnManagedInstance Installs a Windows update on a managed instance.
func (client OsManagementClient) InstallWindowsUpdateOnManagedInstance(ctx context.Context, request InstallWindowsUpdateOnManagedInstanceRequest) (response InstallWindowsUpdateOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) installWindowsUpdateOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/updates/install")
	if err != nil {
		return nil, err
	}

	var response InstallWindowsUpdateOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailablePackagesForManagedInstance Returns a list of packages available for install on the Managed Instance.
func (client OsManagementClient) ListAvailablePackagesForManagedInstance(ctx context.Context, request ListAvailablePackagesForManagedInstanceRequest) (response ListAvailablePackagesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listAvailablePackagesForManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/packages/available")
	if err != nil {
		return nil, err
	}

	var response ListAvailablePackagesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableSoftwareSourcesForManagedInstance Returns a list of available software sources for a Managed Instance.
func (client OsManagementClient) ListAvailableSoftwareSourcesForManagedInstance(ctx context.Context, request ListAvailableSoftwareSourcesForManagedInstanceRequest) (response ListAvailableSoftwareSourcesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listAvailableSoftwareSourcesForManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/availableSoftwareSources")
	if err != nil {
		return nil, err
	}

	var response ListAvailableSoftwareSourcesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableUpdatesForManagedInstance Returns a list of available updates for a Managed Instance.
func (client OsManagementClient) ListAvailableUpdatesForManagedInstance(ctx context.Context, request ListAvailableUpdatesForManagedInstanceRequest) (response ListAvailableUpdatesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listAvailableUpdatesForManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/packages/updates")
	if err != nil {
		return nil, err
	}

	var response ListAvailableUpdatesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableWindowsUpdatesForManagedInstance Returns a list of available Windows updates for a Managed Instance. This is only applicable to Windows instances.
func (client OsManagementClient) ListAvailableWindowsUpdatesForManagedInstance(ctx context.Context, request ListAvailableWindowsUpdatesForManagedInstanceRequest) (response ListAvailableWindowsUpdatesForManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listAvailableWindowsUpdatesForManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/updates/available")
	if err != nil {
		return nil, err
	}

	var response ListAvailableWindowsUpdatesForManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceGroups Returns a list of all Managed Instance Groups.
func (client OsManagementClient) ListManagedInstanceGroups(ctx context.Context, request ListManagedInstanceGroupsRequest) (response ListManagedInstanceGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listManagedInstanceGroups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstanceGroups")
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstances Returns a list of all Managed Instances.
func (client OsManagementClient) ListManagedInstances(ctx context.Context, request ListManagedInstancesRequest) (response ListManagedInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listManagedInstances(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances")
	if err != nil {
		return nil, err
	}

	var response ListManagedInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPackagesInstalledOnManagedInstance Returns a list of installed packages on the Managed Instance.
func (client OsManagementClient) ListPackagesInstalledOnManagedInstance(ctx context.Context, request ListPackagesInstalledOnManagedInstanceRequest) (response ListPackagesInstalledOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listPackagesInstalledOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/packages")
	if err != nil {
		return nil, err
	}

	var response ListPackagesInstalledOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledJobs Returns a list of all of the currently active Scheduled Jobs in the system
func (client OsManagementClient) ListScheduledJobs(ctx context.Context, request ListScheduledJobsRequest) (response ListScheduledJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listScheduledJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs")
	if err != nil {
		return nil, err
	}

	var response ListScheduledJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareSourcePackages Lists Software Packages in a Software Source
func (client OsManagementClient) ListSoftwareSourcePackages(ctx context.Context, request ListSoftwareSourcePackagesRequest) (response ListSoftwareSourcePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listSoftwareSourcePackages(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/{softwareSourceId}/softwarePackages")
	if err != nil {
		return nil, err
	}

	var response ListSoftwareSourcePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareSources Returns a list of all Software Sources.
func (client OsManagementClient) ListSoftwareSources(ctx context.Context, request ListSoftwareSourcesRequest) (response ListSoftwareSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listSoftwareSources(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources")
	if err != nil {
		return nil, err
	}

	var response ListSoftwareSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUpcomingScheduledJobs Returns a list of all of the Scheduled Jobs whose next execution time is at or before the specified time.
func (client OsManagementClient) ListUpcomingScheduledJobs(ctx context.Context, request ListUpcomingScheduledJobsRequest) (response ListUpcomingScheduledJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listUpcomingScheduledJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs/upcomingSchedules")
	if err != nil {
		return nil, err
	}

	var response ListUpcomingScheduledJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWindowsUpdates Returns a list of Windows Updates.
func (client OsManagementClient) ListWindowsUpdates(ctx context.Context, request ListWindowsUpdatesRequest) (response ListWindowsUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWindowsUpdates(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/updates")
	if err != nil {
		return nil, err
	}

	var response ListWindowsUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWindowsUpdatesInstalledOnManagedInstance Returns a list of installed Windows updates for a Managed Instance. This is only applicable to Windows instances.
func (client OsManagementClient) ListWindowsUpdatesInstalledOnManagedInstance(ctx context.Context, request ListWindowsUpdatesInstalledOnManagedInstanceRequest) (response ListWindowsUpdatesInstalledOnManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWindowsUpdatesInstalledOnManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/updates/installed")
	if err != nil {
		return nil, err
	}

	var response ListWindowsUpdatesInstalledOnManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Gets the errors for the work request with the given ID.
func (client OsManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the log entries for the work request with the given ID.
func (client OsManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
func (client OsManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackageFromManagedInstance Removes an installed package from a managed instance.
func (client OsManagementClient) RemovePackageFromManagedInstance(ctx context.Context, request RemovePackageFromManagedInstanceRequest) (response RemovePackageFromManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) removePackageFromManagedInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/packages/remove")
	if err != nil {
		return nil, err
	}

	var response RemovePackageFromManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackagesFromSoftwareSource Removes a given list of Software Packages from a specific Software Source.
func (client OsManagementClient) RemovePackagesFromSoftwareSource(ctx context.Context, request RemovePackagesFromSoftwareSourceRequest) (response RemovePackagesFromSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) removePackagesFromSoftwareSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareSources/{softwareSourceId}/actions/removePackages")
	if err != nil {
		return nil, err
	}

	var response RemovePackagesFromSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunScheduledJobNow This will trigger an already created Scheduled Job to being executing
// immediately instead of waiting for its next regularly scheduled time.
func (client OsManagementClient) RunScheduledJobNow(ctx context.Context, request RunScheduledJobNowRequest) (response RunScheduledJobNowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) runScheduledJobNow(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/runNow")
	if err != nil {
		return nil, err
	}

	var response RunScheduledJobNowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSoftwarePackages Searches all of the available Software Sources and returns any/all Software Packages matching
// the search criteria.
func (client OsManagementClient) SearchSoftwarePackages(ctx context.Context, request SearchSoftwarePackagesRequest) (response SearchSoftwarePackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) searchSoftwarePackages(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareSources/softwarePackages")
	if err != nil {
		return nil, err
	}

	var response SearchSoftwarePackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SkipNextScheduledJobExecution This will force an already created Scheduled Job to skip its
// next regularly scheduled execution
func (client OsManagementClient) SkipNextScheduledJobExecution(ctx context.Context, request SkipNextScheduledJobExecutionRequest) (response SkipNextScheduledJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) skipNextScheduledJobExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/skipNextExecution")
	if err != nil {
		return nil, err
	}

	var response SkipNextScheduledJobExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedInstanceGroup Updates a specific Managed Instance Group.
func (client OsManagementClient) UpdateManagedInstanceGroup(ctx context.Context, request UpdateManagedInstanceGroupRequest) (response UpdateManagedInstanceGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) updateManagedInstanceGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedInstanceGroups/{managedInstanceGroupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagedInstanceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateScheduledJob Updates an existing Scheduled Job on the management system.
func (client OsManagementClient) UpdateScheduledJob(ctx context.Context, request UpdateScheduledJobRequest) (response UpdateScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) updateScheduledJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/scheduledJobs/{scheduledJobId}")
	if err != nil {
		return nil, err
	}

	var response UpdateScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSoftwareSource Updates an existing custom Software Source on the management system.
func (client OsManagementClient) UpdateSoftwareSource(ctx context.Context, request UpdateSoftwareSourceRequest) (response UpdateSoftwareSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OsManagementClient) updateSoftwareSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/softwareSources/{softwareSourceId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSoftwareSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
