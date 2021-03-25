// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v37/common"
	"github.com/oracle/oci-go-sdk/v37/common/auth"
	"net/http"
)

//DbManagementClient a client for DbManagement
type DbManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbManagementClientWithConfigurationProvider Creates a new default DbManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbManagementClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDbManagementClientFromBaseClient(baseClient, provider)
}

// NewDbManagementClientWithOboToken Creates a new default DbManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDbManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbManagementClientFromBaseClient(baseClient, configProvider)
}

func newDbManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbManagementClient, err error) {
	client = DbManagementClient{BaseClient: baseClient}
	client.BasePath = "20201101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemanagement", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddManagedDatabaseToManagedDatabaseGroup Adds a Managed Database to a specific Managed Database Group.
// After the database is added, it will be included in the
// management activities performed on the Managed Database Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/AddManagedDatabaseToManagedDatabaseGroup.go.html to see an example of how to use AddManagedDatabaseToManagedDatabaseGroup API.
func (client DbManagementClient) AddManagedDatabaseToManagedDatabaseGroup(ctx context.Context, request AddManagedDatabaseToManagedDatabaseGroupRequest) (response AddManagedDatabaseToManagedDatabaseGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addManagedDatabaseToManagedDatabaseGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddManagedDatabaseToManagedDatabaseGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddManagedDatabaseToManagedDatabaseGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddManagedDatabaseToManagedDatabaseGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddManagedDatabaseToManagedDatabaseGroupResponse")
	}
	return
}

// addManagedDatabaseToManagedDatabaseGroup implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) addManagedDatabaseToManagedDatabaseGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups/{managedDatabaseGroupId}/actions/addManagedDatabase")
	if err != nil {
		return nil, err
	}

	var response AddManagedDatabaseToManagedDatabaseGroupResponse
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

// ChangeDatabaseParameters Changes database parameters' values. There are two kinds of database
// parameters:
// - Dynamic parameters: They can be changed for the current Oracle
// Database instance. The changes take effect immediately.
// - Static parameters: They cannot be changed for the current instance.
// You must change these parameters and then restart the database before
// changes take effect.
// **Note:** If the instance is started using a text initialization
// parameter file, the parameter changes are applicable only for the
// current instance. You must update them manually to be passed to
// a future instance.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeDatabaseParameters.go.html to see an example of how to use ChangeDatabaseParameters API.
func (client DbManagementClient) ChangeDatabaseParameters(ctx context.Context, request ChangeDatabaseParametersRequest) (response ChangeDatabaseParametersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseParametersResponse")
	}
	return
}

// changeDatabaseParameters implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeDatabaseParameters(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/changeDatabaseParameters")
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseParametersResponse
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

// ChangeJobCompartment Moves a job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeJobCompartment.go.html to see an example of how to use ChangeJobCompartment API.
func (client DbManagementClient) ChangeJobCompartment(ctx context.Context, request ChangeJobCompartmentRequest) (response ChangeJobCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeJobCompartmentResponse")
	}
	return
}

// changeJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeJobCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeJobCompartmentResponse
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

// ChangeManagedDatabaseGroupCompartment Moves a Managed Database Group to a different compartment.
// The destination compartment must not have a Managed Database Group
// with the same name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeManagedDatabaseGroupCompartment.go.html to see an example of how to use ChangeManagedDatabaseGroupCompartment API.
func (client DbManagementClient) ChangeManagedDatabaseGroupCompartment(ctx context.Context, request ChangeManagedDatabaseGroupCompartmentRequest) (response ChangeManagedDatabaseGroupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeManagedDatabaseGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagedDatabaseGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagedDatabaseGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagedDatabaseGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagedDatabaseGroupCompartmentResponse")
	}
	return
}

// changeManagedDatabaseGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeManagedDatabaseGroupCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups/{managedDatabaseGroupId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeManagedDatabaseGroupCompartmentResponse
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

// CreateJob Creates a job to be executed on a Managed Database or Managed Database Group. Only one
// of the parameters, managedDatabaseId or managedDatabaseGroupId should be provided as
// input in CreateJobDetails resource in request body.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateJob.go.html to see an example of how to use CreateJob API.
func (client DbManagementClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobResponse")
	}
	return
}

// createJob implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs")
	if err != nil {
		return nil, err
	}

	var response CreateJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &job{})
	return response, err
}

// CreateManagedDatabaseGroup Creates a Managed Database Group. The group does not contain any
// Managed Databases when it is created, and they must be added later.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateManagedDatabaseGroup.go.html to see an example of how to use CreateManagedDatabaseGroup API.
func (client DbManagementClient) CreateManagedDatabaseGroup(ctx context.Context, request CreateManagedDatabaseGroupRequest) (response CreateManagedDatabaseGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagedDatabaseGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagedDatabaseGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagedDatabaseGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagedDatabaseGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagedDatabaseGroupResponse")
	}
	return
}

// createManagedDatabaseGroup implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createManagedDatabaseGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups")
	if err != nil {
		return nil, err
	}

	var response CreateManagedDatabaseGroupResponse
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

// DeleteJob Deletes the job specified by jobId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteJob.go.html to see an example of how to use DeleteJob API.
func (client DbManagementClient) DeleteJob(ctx context.Context, request DeleteJobRequest) (response DeleteJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobResponse")
	}
	return
}

// deleteJob implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/jobs/{jobId}")
	if err != nil {
		return nil, err
	}

	var response DeleteJobResponse
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

// DeleteManagedDatabaseGroup Deletes the Managed Database Group specified by managedDatabaseGroupId.
// If the group contains Managed Databases, then it cannot be deleted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteManagedDatabaseGroup.go.html to see an example of how to use DeleteManagedDatabaseGroup API.
func (client DbManagementClient) DeleteManagedDatabaseGroup(ctx context.Context, request DeleteManagedDatabaseGroupRequest) (response DeleteManagedDatabaseGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagedDatabaseGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagedDatabaseGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagedDatabaseGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagedDatabaseGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagedDatabaseGroupResponse")
	}
	return
}

// deleteManagedDatabaseGroup implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteManagedDatabaseGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedDatabaseGroups/{managedDatabaseGroupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagedDatabaseGroupResponse
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

// GetDatabaseFleetHealthMetrics Gets the health metrics for a fleet of databases in a compartment or in a Managed Database Group.
// Either the CompartmentId or the ManagedDatabaseGroupId query parameters must be provided to retrieve the health metrics.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDatabaseFleetHealthMetrics.go.html to see an example of how to use GetDatabaseFleetHealthMetrics API.
func (client DbManagementClient) GetDatabaseFleetHealthMetrics(ctx context.Context, request GetDatabaseFleetHealthMetricsRequest) (response GetDatabaseFleetHealthMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseFleetHealthMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseFleetHealthMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseFleetHealthMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseFleetHealthMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseFleetHealthMetricsResponse")
	}
	return
}

// getDatabaseFleetHealthMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getDatabaseFleetHealthMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleetMetrics")
	if err != nil {
		return nil, err
	}

	var response GetDatabaseFleetHealthMetricsResponse
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

// GetDatabaseHomeMetrics Gets a summary of the activity and resource usage metrics like DB Time, CPU, User I/O, Wait, Storage, and Memory for a Managed Database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDatabaseHomeMetrics.go.html to see an example of how to use GetDatabaseHomeMetrics API.
func (client DbManagementClient) GetDatabaseHomeMetrics(ctx context.Context, request GetDatabaseHomeMetricsRequest) (response GetDatabaseHomeMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseHomeMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseHomeMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseHomeMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseHomeMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseHomeMetricsResponse")
	}
	return
}

// getDatabaseHomeMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getDatabaseHomeMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseHomeMetrics")
	if err != nil {
		return nil, err
	}

	var response GetDatabaseHomeMetricsResponse
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

// GetJob Gets the details for the job specified by jobId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetJob.go.html to see an example of how to use GetJob API.
func (client DbManagementClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobResponse")
	}
	return
}

// getJob implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}")
	if err != nil {
		return nil, err
	}

	var response GetJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &job{})
	return response, err
}

// GetJobExecution Gets the details for the job execution specified by jobExecutionId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetJobExecution.go.html to see an example of how to use GetJobExecution API.
func (client DbManagementClient) GetJobExecution(ctx context.Context, request GetJobExecutionRequest) (response GetJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobExecutionResponse")
	}
	return
}

// getJobExecution implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getJobExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobExecutions/{jobExecutionId}")
	if err != nil {
		return nil, err
	}

	var response GetJobExecutionResponse
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

// GetJobRun Gets the details for the job run specified by jobRunId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetJobRun.go.html to see an example of how to use GetJobRun API.
func (client DbManagementClient) GetJobRun(ctx context.Context, request GetJobRunRequest) (response GetJobRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobRunResponse")
	}
	return
}

// getJobRun implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getJobRun(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobRuns/{jobRunId}")
	if err != nil {
		return nil, err
	}

	var response GetJobRunResponse
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

// GetManagedDatabase Gets the details for the Managed Database specified by managedDatabaseId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetManagedDatabase.go.html to see an example of how to use GetManagedDatabase API.
func (client DbManagementClient) GetManagedDatabase(ctx context.Context, request GetManagedDatabaseRequest) (response GetManagedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedDatabaseResponse")
	}
	return
}

// getManagedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getManagedDatabase(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}")
	if err != nil {
		return nil, err
	}

	var response GetManagedDatabaseResponse
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

// GetManagedDatabaseGroup Gets the details for the Managed Database Group specified by managedDatabaseGroupId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetManagedDatabaseGroup.go.html to see an example of how to use GetManagedDatabaseGroup API.
func (client DbManagementClient) GetManagedDatabaseGroup(ctx context.Context, request GetManagedDatabaseGroupRequest) (response GetManagedDatabaseGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedDatabaseGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedDatabaseGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedDatabaseGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedDatabaseGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedDatabaseGroupResponse")
	}
	return
}

// getManagedDatabaseGroup implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getManagedDatabaseGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabaseGroups/{managedDatabaseGroupId}")
	if err != nil {
		return nil, err
	}

	var response GetManagedDatabaseGroupResponse
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

// ListDatabaseParameters Gets the list of database parameters for the specified Managed Database. The parameters are listed in alphabetical order, along with their current values.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListDatabaseParameters.go.html to see an example of how to use ListDatabaseParameters API.
func (client DbManagementClient) ListDatabaseParameters(ctx context.Context, request ListDatabaseParametersRequest) (response ListDatabaseParametersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseParametersResponse")
	}
	return
}

// listDatabaseParameters implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listDatabaseParameters(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/databaseParameters")
	if err != nil {
		return nil, err
	}

	var response ListDatabaseParametersResponse
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

// ListJobExecutions Gets the job execution for a specific ID or the list of job executions for a job, Managed Database or Managed Database Group
// in a specific compartment. Only one of the parameters, ID, jobId, managedDatabaseId or managedDatabaseGroupId should be provided.
// If none of these parameters is provided, all the job executions in the compartment are listed. Job executions can also be filtered
// based on the name and status parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListJobExecutions.go.html to see an example of how to use ListJobExecutions API.
func (client DbManagementClient) ListJobExecutions(ctx context.Context, request ListJobExecutionsRequest) (response ListJobExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobExecutionsResponse")
	}
	return
}

// listJobExecutions implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listJobExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobExecutions")
	if err != nil {
		return nil, err
	}

	var response ListJobExecutionsResponse
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

// ListJobRuns Gets the job run for a specific ID or the list of job runs for a job, Managed Database or Managed Database Group
// in a specific compartment. Only one of the parameters, ID, jobId, managedDatabaseId, or managedDatabaseGroupId
// should be provided. If none of these parameters is provided, all the job runs in the compartment are listed.
// Job runs can also be filtered based on name and runStatus parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListJobRuns.go.html to see an example of how to use ListJobRuns API.
func (client DbManagementClient) ListJobRuns(ctx context.Context, request ListJobRunsRequest) (response ListJobRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobRunsResponse")
	}
	return
}

// listJobRuns implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listJobRuns(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobRuns")
	if err != nil {
		return nil, err
	}

	var response ListJobRunsResponse
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

// ListJobs Gets the job for a specific ID or the list of jobs for a Managed Database or Managed Database Group
// in a specific compartment. Only one of the parameters, ID, managedDatabaseId or managedDatabaseGroupId,
// should be provided. If none of these parameters is provided, all the jobs in the compartment are listed.
// Jobs can also be filtered based on the name and lifecycleState parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListJobs.go.html to see an example of how to use ListJobs API.
func (client DbManagementClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobsResponse")
	}
	return
}

// listJobs implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs")
	if err != nil {
		return nil, err
	}

	var response ListJobsResponse
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

// ListManagedDatabaseGroups Gets the Managed Database Group for a specific ID or the list of Managed Database Groups in
// a specific compartment. Managed Database Groups can also be filtered based on the name parameter.
// Only one of the parameters, ID or name should be provided. If none of these parameters is provided,
// all the Managed Database Groups in the compartment are listed.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedDatabaseGroups.go.html to see an example of how to use ListManagedDatabaseGroups API.
func (client DbManagementClient) ListManagedDatabaseGroups(ctx context.Context, request ListManagedDatabaseGroupsRequest) (response ListManagedDatabaseGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedDatabaseGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedDatabaseGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedDatabaseGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedDatabaseGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedDatabaseGroupsResponse")
	}
	return
}

// listManagedDatabaseGroups implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listManagedDatabaseGroups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabaseGroups")
	if err != nil {
		return nil, err
	}

	var response ListManagedDatabaseGroupsResponse
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

// ListManagedDatabases Gets the Managed Database for a specific ID or the list of Managed Databases in a specific compartment.
// Managed Databases can also be filtered based on the name parameter. Only one of the parameters, ID or name
// should be provided. If none of these parameters is provided, all the Managed Databases in the compartment are listed.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedDatabases.go.html to see an example of how to use ListManagedDatabases API.
func (client DbManagementClient) ListManagedDatabases(ctx context.Context, request ListManagedDatabasesRequest) (response ListManagedDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedDatabasesResponse")
	}
	return
}

// listManagedDatabases implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listManagedDatabases(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases")
	if err != nil {
		return nil, err
	}

	var response ListManagedDatabasesResponse
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

// ListTablespaces Gets the list of tablespaces for the specified managedDatabaseId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListTablespaces.go.html to see an example of how to use ListTablespaces API.
func (client DbManagementClient) ListTablespaces(ctx context.Context, request ListTablespacesRequest) (response ListTablespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTablespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTablespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTablespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTablespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTablespacesResponse")
	}
	return
}

// listTablespaces implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listTablespaces(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/tablespaces")
	if err != nil {
		return nil, err
	}

	var response ListTablespacesResponse
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

// RemoveManagedDatabaseFromManagedDatabaseGroup Removes a Managed Database from a Managed Database Group. Any management
// activities that are currently running on this database will continue to
// run to completion. However, any activities scheduled to run in the future
// will not be performed on this database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/RemoveManagedDatabaseFromManagedDatabaseGroup.go.html to see an example of how to use RemoveManagedDatabaseFromManagedDatabaseGroup API.
func (client DbManagementClient) RemoveManagedDatabaseFromManagedDatabaseGroup(ctx context.Context, request RemoveManagedDatabaseFromManagedDatabaseGroupRequest) (response RemoveManagedDatabaseFromManagedDatabaseGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeManagedDatabaseFromManagedDatabaseGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveManagedDatabaseFromManagedDatabaseGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveManagedDatabaseFromManagedDatabaseGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveManagedDatabaseFromManagedDatabaseGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveManagedDatabaseFromManagedDatabaseGroupResponse")
	}
	return
}

// removeManagedDatabaseFromManagedDatabaseGroup implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) removeManagedDatabaseFromManagedDatabaseGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups/{managedDatabaseGroupId}/actions/removeManagedDatabase")
	if err != nil {
		return nil, err
	}

	var response RemoveManagedDatabaseFromManagedDatabaseGroupResponse
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

// ResetDatabaseParameters Resets database parameters' values to their default or startup values.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ResetDatabaseParameters.go.html to see an example of how to use ResetDatabaseParameters API.
func (client DbManagementClient) ResetDatabaseParameters(ctx context.Context, request ResetDatabaseParametersRequest) (response ResetDatabaseParametersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resetDatabaseParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResetDatabaseParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResetDatabaseParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResetDatabaseParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResetDatabaseParametersResponse")
	}
	return
}

// resetDatabaseParameters implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) resetDatabaseParameters(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/resetDatabaseParameters")
	if err != nil {
		return nil, err
	}

	var response ResetDatabaseParametersResponse
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

// UpdateManagedDatabaseGroup Updates the Managed Database Group specified by managedDatabaseGroupId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateManagedDatabaseGroup.go.html to see an example of how to use UpdateManagedDatabaseGroup API.
func (client DbManagementClient) UpdateManagedDatabaseGroup(ctx context.Context, request UpdateManagedDatabaseGroupRequest) (response UpdateManagedDatabaseGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateManagedDatabaseGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagedDatabaseGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagedDatabaseGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagedDatabaseGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagedDatabaseGroupResponse")
	}
	return
}

// updateManagedDatabaseGroup implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateManagedDatabaseGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedDatabaseGroups/{managedDatabaseGroupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagedDatabaseGroupResponse
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
