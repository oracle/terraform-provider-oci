// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
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
	// DbManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

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
func (client DbManagementClient) addManagedDatabaseToManagedDatabaseGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups/{managedDatabaseGroupId}/actions/addManagedDatabase", binaryReqBody, extraHeaders)
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

// ChangeDatabaseParameters Changes database parameter values. There are two kinds of database
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
func (client DbManagementClient) changeDatabaseParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/changeDatabaseParameters", binaryReqBody, extraHeaders)
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

// ChangeDbManagementPrivateEndpointCompartment Moves the Database Management private endpoint and its dependent resources to the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeDbManagementPrivateEndpointCompartment.go.html to see an example of how to use ChangeDbManagementPrivateEndpointCompartment API.
func (client DbManagementClient) ChangeDbManagementPrivateEndpointCompartment(ctx context.Context, request ChangeDbManagementPrivateEndpointCompartmentRequest) (response ChangeDbManagementPrivateEndpointCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDbManagementPrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDbManagementPrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDbManagementPrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDbManagementPrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDbManagementPrivateEndpointCompartmentResponse")
	}
	return
}

// changeDbManagementPrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeDbManagementPrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbManagementPrivateEndpoints/{dbManagementPrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDbManagementPrivateEndpointCompartmentResponse
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
func (client DbManagementClient) changeJobCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/actions/changeCompartment", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) changeManagedDatabaseGroupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups/{managedDatabaseGroupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
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

// CreateDbManagementPrivateEndpoint Creates a new Database Management private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateDbManagementPrivateEndpoint.go.html to see an example of how to use CreateDbManagementPrivateEndpoint API.
func (client DbManagementClient) CreateDbManagementPrivateEndpoint(ctx context.Context, request CreateDbManagementPrivateEndpointRequest) (response CreateDbManagementPrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDbManagementPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDbManagementPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDbManagementPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbManagementPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbManagementPrivateEndpointResponse")
	}
	return
}

// createDbManagementPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createDbManagementPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbManagementPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDbManagementPrivateEndpointResponse
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
func (client DbManagementClient) createJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) createManagedDatabaseGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups", binaryReqBody, extraHeaders)
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

// DeleteDbManagementPrivateEndpoint Deletes a specific Database Management private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteDbManagementPrivateEndpoint.go.html to see an example of how to use DeleteDbManagementPrivateEndpoint API.
func (client DbManagementClient) DeleteDbManagementPrivateEndpoint(ctx context.Context, request DeleteDbManagementPrivateEndpointRequest) (response DeleteDbManagementPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbManagementPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDbManagementPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDbManagementPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbManagementPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbManagementPrivateEndpointResponse")
	}
	return
}

// deleteDbManagementPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteDbManagementPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbManagementPrivateEndpoints/{dbManagementPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDbManagementPrivateEndpointResponse
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
func (client DbManagementClient) deleteJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/jobs/{jobId}", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) deleteManagedDatabaseGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedDatabaseGroups/{managedDatabaseGroupId}", binaryReqBody, extraHeaders)
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

// GetAwrDbReport Gets the AWR report for the specific database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetAwrDbReport.go.html to see an example of how to use GetAwrDbReport API.
func (client DbManagementClient) GetAwrDbReport(ctx context.Context, request GetAwrDbReportRequest) (response GetAwrDbReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAwrDbReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrDbReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrDbReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrDbReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrDbReportResponse")
	}
	return
}

// getAwrDbReport implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getAwrDbReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrDbReportResponse
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

// GetAwrDbSqlReport Gets the SQL health check report for one SQL of the specific database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetAwrDbSqlReport.go.html to see an example of how to use GetAwrDbSqlReport API.
func (client DbManagementClient) GetAwrDbSqlReport(ctx context.Context, request GetAwrDbSqlReportRequest) (response GetAwrDbSqlReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAwrDbSqlReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrDbSqlReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrDbSqlReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrDbSqlReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrDbSqlReportResponse")
	}
	return
}

// getAwrDbSqlReport implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getAwrDbSqlReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSqlReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrDbSqlReportResponse
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

// GetClusterCacheMetric Gets the metrics related to cluster cache for the Oracle
// Real Application Clusters (Oracle RAC) database specified
// by managedDatabaseId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetClusterCacheMetric.go.html to see an example of how to use GetClusterCacheMetric API.
func (client DbManagementClient) GetClusterCacheMetric(ctx context.Context, request GetClusterCacheMetricRequest) (response GetClusterCacheMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterCacheMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterCacheMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterCacheMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterCacheMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterCacheMetricResponse")
	}
	return
}

// getClusterCacheMetric implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getClusterCacheMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/clusterCacheMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterCacheMetricResponse
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
func (client DbManagementClient) getDatabaseFleetHealthMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleetMetrics", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) getDatabaseHomeMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseHomeMetrics", binaryReqBody, extraHeaders)
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

// GetDbManagementPrivateEndpoint Gets the details of a specific Database Management private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDbManagementPrivateEndpoint.go.html to see an example of how to use GetDbManagementPrivateEndpoint API.
func (client DbManagementClient) GetDbManagementPrivateEndpoint(ctx context.Context, request GetDbManagementPrivateEndpointRequest) (response GetDbManagementPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbManagementPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbManagementPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbManagementPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbManagementPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbManagementPrivateEndpointResponse")
	}
	return
}

// getDbManagementPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getDbManagementPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbManagementPrivateEndpoints/{dbManagementPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbManagementPrivateEndpointResponse
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
func (client DbManagementClient) getJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) getJobExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobExecutions/{jobExecutionId}", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) getJobRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobRuns/{jobRunId}", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) getManagedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) getManagedDatabaseGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabaseGroups/{managedDatabaseGroupId}", binaryReqBody, extraHeaders)
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

// GetPdbMetrics Gets a summary of the resource usage metrics such as CPU, User I/O, and Storage for each
// PDB within a specific CDB. If comparmentId is specified, then the metrics for
// each PDB (within the CDB) in the specified compartment are retrieved.
// If compartmentId is not specified, then the metrics for all the PDBs within the CDB are retrieved.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetPdbMetrics.go.html to see an example of how to use GetPdbMetrics API.
func (client DbManagementClient) GetPdbMetrics(ctx context.Context, request GetPdbMetricsRequest) (response GetPdbMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPdbMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPdbMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPdbMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPdbMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPdbMetricsResponse")
	}
	return
}

// getPdbMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getPdbMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/pdbMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPdbMetricsResponse
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

// GetUser Gets the details of a specific user for the specified managedDatabaseId and userName.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetUser.go.html to see an example of how to use GetUser API.
func (client DbManagementClient) GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserResponse")
	}
	return
}

// getUser implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserResponse
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

// GetWorkRequest Gets the status of the work request with the given Work Request ID
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client DbManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DbManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssociatedDatabases Gets the list of databases using a specific Database Management private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAssociatedDatabases.go.html to see an example of how to use ListAssociatedDatabases API.
func (client DbManagementClient) ListAssociatedDatabases(ctx context.Context, request ListAssociatedDatabasesRequest) (response ListAssociatedDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssociatedDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssociatedDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssociatedDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssociatedDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssociatedDatabasesResponse")
	}
	return
}

// listAssociatedDatabases implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listAssociatedDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbManagementPrivateEndpoints/{dbManagementPrivateEndpointId}/associatedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssociatedDatabasesResponse
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

// ListAwrDbSnapshots Lists AWR snapshots for the specified database in the AWR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAwrDbSnapshots.go.html to see an example of how to use ListAwrDbSnapshots API.
func (client DbManagementClient) ListAwrDbSnapshots(ctx context.Context, request ListAwrDbSnapshotsRequest) (response ListAwrDbSnapshotsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAwrDbSnapshots, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrDbSnapshotsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrDbSnapshotsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrDbSnapshotsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrDbSnapshotsResponse")
	}
	return
}

// listAwrDbSnapshots implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listAwrDbSnapshots(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrDbSnapshotsResponse
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

// ListAwrDbs Gets the list of databases and their snapshot summary details available in the AWR of the specified Managed Database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAwrDbs.go.html to see an example of how to use ListAwrDbs API.
func (client DbManagementClient) ListAwrDbs(ctx context.Context, request ListAwrDbsRequest) (response ListAwrDbsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAwrDbs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrDbsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrDbsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrDbsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrDbsResponse")
	}
	return
}

// listAwrDbs implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listAwrDbs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrDbsResponse
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

// ListConsumerGroupPrivileges Gets the list of Consumer Group Privileges granted for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListConsumerGroupPrivileges.go.html to see an example of how to use ListConsumerGroupPrivileges API.
func (client DbManagementClient) ListConsumerGroupPrivileges(ctx context.Context, request ListConsumerGroupPrivilegesRequest) (response ListConsumerGroupPrivilegesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConsumerGroupPrivileges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConsumerGroupPrivilegesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConsumerGroupPrivilegesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConsumerGroupPrivilegesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConsumerGroupPrivilegesResponse")
	}
	return
}

// listConsumerGroupPrivileges implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listConsumerGroupPrivileges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/consumerGroupPrivileges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConsumerGroupPrivilegesResponse
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

// ListDataAccessContainers Gets the list of Containers if it does not apply to all containers for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListDataAccessContainers.go.html to see an example of how to use ListDataAccessContainers API.
func (client DbManagementClient) ListDataAccessContainers(ctx context.Context, request ListDataAccessContainersRequest) (response ListDataAccessContainersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAccessContainers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataAccessContainersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataAccessContainersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAccessContainersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAccessContainersResponse")
	}
	return
}

// listDataAccessContainers implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listDataAccessContainers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/dataAccessContainers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataAccessContainersResponse
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
func (client DbManagementClient) listDatabaseParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/databaseParameters", binaryReqBody, extraHeaders)
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

// ListDbManagementPrivateEndpoints Gets a list of Database Management private endpoints.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListDbManagementPrivateEndpoints.go.html to see an example of how to use ListDbManagementPrivateEndpoints API.
func (client DbManagementClient) ListDbManagementPrivateEndpoints(ctx context.Context, request ListDbManagementPrivateEndpointsRequest) (response ListDbManagementPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbManagementPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbManagementPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbManagementPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbManagementPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbManagementPrivateEndpointsResponse")
	}
	return
}

// listDbManagementPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listDbManagementPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbManagementPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbManagementPrivateEndpointsResponse
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

// ListJobExecutions Gets the job execution for a specific ID or the list of job executions for a job, job run, Managed Database or Managed Database Group
// in a specific compartment. Only one of the parameters, ID, jobId, jobRunId, managedDatabaseId or managedDatabaseGroupId should be provided.
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
func (client DbManagementClient) listJobExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobExecutions", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) listJobRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobRuns", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) listJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs", binaryReqBody, extraHeaders)
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
func (client DbManagementClient) listManagedDatabaseGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabaseGroups", binaryReqBody, extraHeaders)
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
// Managed Databases can be filtered based on the name parameter. Only one of the parameters, ID or name
// should be provided. If neither of these parameters is provided, all the Managed Databases in the compartment
// are listed. Managed Databases can also be filtered based on the deployment type and management option.
// If the deployment type is not specified or if it is `ONPREMISE`, then the management option is not
// considered and Managed Databases with `ADVANCED` management option are listed.
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
func (client DbManagementClient) listManagedDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases", binaryReqBody, extraHeaders)
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

// ListObjectPrivileges Gets the list of Object Privileges granted for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListObjectPrivileges.go.html to see an example of how to use ListObjectPrivileges API.
func (client DbManagementClient) ListObjectPrivileges(ctx context.Context, request ListObjectPrivilegesRequest) (response ListObjectPrivilegesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listObjectPrivileges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListObjectPrivilegesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListObjectPrivilegesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListObjectPrivilegesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListObjectPrivilegesResponse")
	}
	return
}

// listObjectPrivileges implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listObjectPrivileges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/objectPrivileges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListObjectPrivilegesResponse
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

// ListProxiedForUsers Gets the list of Users for which the current user acts as proxy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListProxiedForUsers.go.html to see an example of how to use ListProxiedForUsers API.
func (client DbManagementClient) ListProxiedForUsers(ctx context.Context, request ListProxiedForUsersRequest) (response ListProxiedForUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProxiedForUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProxiedForUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProxiedForUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProxiedForUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProxiedForUsersResponse")
	}
	return
}

// listProxiedForUsers implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listProxiedForUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/proxiedForUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProxiedForUsersResponse
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

// ListProxyUsers Gets the list of proxy users for the current User.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListProxyUsers.go.html to see an example of how to use ListProxyUsers API.
func (client DbManagementClient) ListProxyUsers(ctx context.Context, request ListProxyUsersRequest) (response ListProxyUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProxyUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProxyUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProxyUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProxyUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProxyUsersResponse")
	}
	return
}

// listProxyUsers implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listProxyUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/proxyUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProxyUsersResponse
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

// ListRoles Gets the list of roles granted for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListRoles.go.html to see an example of how to use ListRoles API.
func (client DbManagementClient) ListRoles(ctx context.Context, request ListRolesRequest) (response ListRolesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRolesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRolesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRolesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRolesResponse")
	}
	return
}

// listRoles implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listRoles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/roles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRolesResponse
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

// ListSystemPrivileges Gets the list of System Privileges granted for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSystemPrivileges.go.html to see an example of how to use ListSystemPrivileges API.
func (client DbManagementClient) ListSystemPrivileges(ctx context.Context, request ListSystemPrivilegesRequest) (response ListSystemPrivilegesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSystemPrivileges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSystemPrivilegesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSystemPrivilegesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSystemPrivilegesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSystemPrivilegesResponse")
	}
	return
}

// listSystemPrivileges implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listSystemPrivileges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users/{userName}/systemPrivileges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSystemPrivilegesResponse
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
func (client DbManagementClient) listTablespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/tablespaces", binaryReqBody, extraHeaders)
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

// ListUsers Gets the list of users for the specified managedDatabaseId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListUsers.go.html to see an example of how to use ListUsers API.
func (client DbManagementClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsersResponse")
	}
	return
}

// listUsers implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsersResponse
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

// ListWorkRequestErrors Returns a paginated list of errors for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client DbManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DbManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client DbManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DbManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests The list of work requests in a specific compartment was retrieved successfully.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client DbManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DbManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
func (client DbManagementClient) removeManagedDatabaseFromManagedDatabaseGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabaseGroups/{managedDatabaseGroupId}/actions/removeManagedDatabase", binaryReqBody, extraHeaders)
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

// ResetDatabaseParameters Resets database parameter values to their default or startup values.
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
func (client DbManagementClient) resetDatabaseParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/resetDatabaseParameters", binaryReqBody, extraHeaders)
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

// SummarizeAwrDbCpuUsages Summarizes the AWR CPU resource limits and metrics for the specified database in AWR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbCpuUsages.go.html to see an example of how to use SummarizeAwrDbCpuUsages API.
func (client DbManagementClient) SummarizeAwrDbCpuUsages(ctx context.Context, request SummarizeAwrDbCpuUsagesRequest) (response SummarizeAwrDbCpuUsagesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbCpuUsages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbCpuUsagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbCpuUsagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbCpuUsagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbCpuUsagesResponse")
	}
	return
}

// summarizeAwrDbCpuUsages implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbCpuUsages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbCpuUsages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbCpuUsagesResponse
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

// SummarizeAwrDbMetrics Summarizes the metric samples for the specified database in the AWR. The metric samples are summarized based on the Time dimension for each metric.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbMetrics.go.html to see an example of how to use SummarizeAwrDbMetrics API.
func (client DbManagementClient) SummarizeAwrDbMetrics(ctx context.Context, request SummarizeAwrDbMetricsRequest) (response SummarizeAwrDbMetricsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbMetricsResponse")
	}
	return
}

// summarizeAwrDbMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbMetricsResponse
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

// SummarizeAwrDbParameterChanges Summarizes the database parameter change history for one database parameter of the specified database in AWR. One change history record contains
// the previous value, the changed value, and the corresponding time range. If the database parameter value was changed multiple times within the time range, then multiple change history records are created for the same parameter.
// Note that this API only returns information on change history details for one database parameter.
// To get a list of all the database parameters whose values were changed during a specified time range, use the following API endpoint:
// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbParameters
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbParameterChanges.go.html to see an example of how to use SummarizeAwrDbParameterChanges API.
func (client DbManagementClient) SummarizeAwrDbParameterChanges(ctx context.Context, request SummarizeAwrDbParameterChangesRequest) (response SummarizeAwrDbParameterChangesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbParameterChanges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbParameterChangesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbParameterChangesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbParameterChangesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbParameterChangesResponse")
	}
	return
}

// summarizeAwrDbParameterChanges implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbParameterChanges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbParameterChanges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbParameterChangesResponse
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

// SummarizeAwrDbParameters Summarizes the database parameter history for the specified database in AWR. This includes the list of database
// parameters, with information on whether the parameter values were modified within the query time range. Note that
// each database parameter is only listed once. Depending on the optional query parameters, the returned summary gets all the database parameters, which include:
// - Each parameter whose value was changed during the time range:  (valueChanged ="Y")
// - Each parameter whose value was unchanged during the time range:  (valueChanged ="N")
// - Each parameter whose value was changed at the system level during the time range: (valueChanged ="Y"  and valueModified = "SYSTEM_MOD")
// - Each parameter whose value was unchanged during the time range, however, the value is not the default value: (valueChanged ="N" and  valueDefault = "FALSE")
// Note that this API does not return information on the number of times each database parameter has been changed within the time range. To get the database parameter value change history for a specific parameter, use the following API endpoint:
// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbParameterChanges
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbParameters.go.html to see an example of how to use SummarizeAwrDbParameters API.
func (client DbManagementClient) SummarizeAwrDbParameters(ctx context.Context, request SummarizeAwrDbParametersRequest) (response SummarizeAwrDbParametersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbParametersResponse")
	}
	return
}

// summarizeAwrDbParameters implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbParameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbParametersResponse
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

// SummarizeAwrDbSnapshotRanges Summarizes the AWR snapshot ranges that contain continuous snapshots, for the specified Managed Database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbSnapshotRanges.go.html to see an example of how to use SummarizeAwrDbSnapshotRanges API.
func (client DbManagementClient) SummarizeAwrDbSnapshotRanges(ctx context.Context, request SummarizeAwrDbSnapshotRangesRequest) (response SummarizeAwrDbSnapshotRangesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbSnapshotRanges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbSnapshotRangesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbSnapshotRangesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbSnapshotRangesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbSnapshotRangesResponse")
	}
	return
}

// summarizeAwrDbSnapshotRanges implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbSnapshotRanges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbSnapshotRangesResponse
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

// SummarizeAwrDbSysstats Summarizes the AWR SYSSTAT sample data for the specified database in AWR. The statistical data is summarized based on the Time dimension for each statistic.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbSysstats.go.html to see an example of how to use SummarizeAwrDbSysstats API.
func (client DbManagementClient) SummarizeAwrDbSysstats(ctx context.Context, request SummarizeAwrDbSysstatsRequest) (response SummarizeAwrDbSysstatsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbSysstats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbSysstatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbSysstatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbSysstatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbSysstatsResponse")
	}
	return
}

// summarizeAwrDbSysstats implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbSysstats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSysstats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbSysstatsResponse
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

// SummarizeAwrDbTopWaitEvents Summarizes the AWR top wait events.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbTopWaitEvents.go.html to see an example of how to use SummarizeAwrDbTopWaitEvents API.
func (client DbManagementClient) SummarizeAwrDbTopWaitEvents(ctx context.Context, request SummarizeAwrDbTopWaitEventsRequest) (response SummarizeAwrDbTopWaitEventsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbTopWaitEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbTopWaitEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbTopWaitEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbTopWaitEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbTopWaitEventsResponse")
	}
	return
}

// summarizeAwrDbTopWaitEvents implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbTopWaitEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbTopWaitEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbTopWaitEventsResponse
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

// SummarizeAwrDbWaitEventBuckets Summarizes AWR wait event data into value buckets and frequency, for the specified database in the AWR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbWaitEventBuckets.go.html to see an example of how to use SummarizeAwrDbWaitEventBuckets API.
func (client DbManagementClient) SummarizeAwrDbWaitEventBuckets(ctx context.Context, request SummarizeAwrDbWaitEventBucketsRequest) (response SummarizeAwrDbWaitEventBucketsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbWaitEventBuckets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbWaitEventBucketsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbWaitEventBucketsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbWaitEventBucketsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbWaitEventBucketsResponse")
	}
	return
}

// summarizeAwrDbWaitEventBuckets implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbWaitEventBuckets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbWaitEventBuckets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbWaitEventBucketsResponse
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

// SummarizeAwrDbWaitEvents Summarizes the AWR wait event sample data for the specified database in the AWR. The event data is summarized based on the Time dimension for each event.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbWaitEvents.go.html to see an example of how to use SummarizeAwrDbWaitEvents API.
func (client DbManagementClient) SummarizeAwrDbWaitEvents(ctx context.Context, request SummarizeAwrDbWaitEventsRequest) (response SummarizeAwrDbWaitEventsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDbWaitEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDbWaitEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDbWaitEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDbWaitEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDbWaitEventsResponse")
	}
	return
}

// summarizeAwrDbWaitEvents implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeAwrDbWaitEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbWaitEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDbWaitEventsResponse
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

// SummarizeJobExecutionsStatuses Gets the number of job executions grouped by status for a job, Managed Database, or Database Group in a specific compartment. Only one of the parameters, jobId, managedDatabaseId, or managedDatabaseGroupId should be provided.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeJobExecutionsStatuses.go.html to see an example of how to use SummarizeJobExecutionsStatuses API.
func (client DbManagementClient) SummarizeJobExecutionsStatuses(ctx context.Context, request SummarizeJobExecutionsStatusesRequest) (response SummarizeJobExecutionsStatusesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeJobExecutionsStatuses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeJobExecutionsStatusesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeJobExecutionsStatusesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeJobExecutionsStatusesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeJobExecutionsStatusesResponse")
	}
	return
}

// summarizeJobExecutionsStatuses implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeJobExecutionsStatuses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobExecutionsStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeJobExecutionsStatusesResponse
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

// UpdateDbManagementPrivateEndpoint Updates one or more attributes of a specific Database Management private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateDbManagementPrivateEndpoint.go.html to see an example of how to use UpdateDbManagementPrivateEndpoint API.
func (client DbManagementClient) UpdateDbManagementPrivateEndpoint(ctx context.Context, request UpdateDbManagementPrivateEndpointRequest) (response UpdateDbManagementPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbManagementPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbManagementPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbManagementPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbManagementPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbManagementPrivateEndpointResponse")
	}
	return
}

// updateDbManagementPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateDbManagementPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbManagementPrivateEndpoints/{dbManagementPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbManagementPrivateEndpointResponse
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

// UpdateJob Updates the details for the recurring scheduled job specified by jobId. Note that non-recurring (one time) jobs cannot be updated.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateJob.go.html to see an example of how to use UpdateJob API.
func (client DbManagementClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobResponse")
	}
	return
}

// updateJob implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/jobs/{jobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJobResponse
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
func (client DbManagementClient) updateManagedDatabaseGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedDatabaseGroups/{managedDatabaseGroupId}", binaryReqBody, extraHeaders)
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
