// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DbManagementClient a client for DbManagement
type DbManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbManagementClientWithConfigurationProvider Creates a new default DbManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("databasemanagement"); !enabled {
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
	return newDbManagementClientFromBaseClient(baseClient, provider)
}

// NewDbManagementClientWithOboToken Creates a new default DbManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDbManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbManagementClientFromBaseClient(baseClient, configProvider)
}

func newDbManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbManagementClient, err error) {
	// DbManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbManagement"))
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
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DbManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddDataFiles Adds data files or temp files to the tablespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/AddDataFiles.go.html to see an example of how to use AddDataFiles API.
func (client DbManagementClient) AddDataFiles(ctx context.Context, request AddDataFilesRequest) (response AddDataFilesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addDataFiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDataFilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDataFilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDataFilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDataFilesResponse")
	}
	return
}

// addDataFiles implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) addDataFiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/tablespaces/{tablespaceName}/actions/addDataFiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDataFilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/AddDataFiles"
		err = common.PostProcessServiceError(err, "DbManagement", "AddDataFiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddManagedDatabaseToManagedDatabaseGroup Adds a Managed Database to a specific Managed Database Group.
// After the database is added, it will be included in the
// management activities performed on the Managed Database Group.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/AddManagedDatabaseToManagedDatabaseGroup"
		err = common.PostProcessServiceError(err, "DbManagement", "AddManagedDatabaseToManagedDatabaseGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddmTasks Lists the metadata for each ADDM task who's end snapshot time falls within the provided start and end time. Details include
// the name of the ADDM task, description, user, status and creation date time.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/AddmTasks.go.html to see an example of how to use AddmTasks API.
func (client DbManagementClient) AddmTasks(ctx context.Context, request AddmTasksRequest) (response AddmTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addmTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddmTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddmTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddmTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddmTasksResponse")
	}
	return
}

// addmTasks implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) addmTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/addmTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddmTasksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/AddmTasksCollection/AddmTasks"
		err = common.PostProcessServiceError(err, "DbManagement", "AddmTasks", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ChangeDatabaseParameters"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeDatabaseParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDbManagementPrivateEndpointCompartment Moves the Database Management private endpoint and its dependent resources to the specified compartment.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DbManagementPrivateEndpoint/ChangeDbManagementPrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeDbManagementPrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExternalDbSystemCompartment Moves the external DB system and its related resources (excluding databases) to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeExternalDbSystemCompartment.go.html to see an example of how to use ChangeExternalDbSystemCompartment API.
// A default retry strategy applies to this operation ChangeExternalDbSystemCompartment()
func (client DbManagementClient) ChangeExternalDbSystemCompartment(ctx context.Context, request ChangeExternalDbSystemCompartmentRequest) (response ChangeExternalDbSystemCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExternalDbSystemCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExternalDbSystemCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExternalDbSystemCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExternalDbSystemCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExternalDbSystemCompartmentResponse")
	}
	return
}

// changeExternalDbSystemCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeExternalDbSystemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystems/{externalDbSystemId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalDbSystemCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/ChangeExternalDbSystemCompartment"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeExternalDbSystemCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExternalExadataInfrastructureCompartment Moves the Exadata infrastructure and its related resources (Exadata storage server, Exadata storage server connectors and Exadata storage server grid) to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeExternalExadataInfrastructureCompartment.go.html to see an example of how to use ChangeExternalExadataInfrastructureCompartment API.
// A default retry strategy applies to this operation ChangeExternalExadataInfrastructureCompartment()
func (client DbManagementClient) ChangeExternalExadataInfrastructureCompartment(ctx context.Context, request ChangeExternalExadataInfrastructureCompartmentRequest) (response ChangeExternalExadataInfrastructureCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExternalExadataInfrastructureCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExternalExadataInfrastructureCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExternalExadataInfrastructureCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExternalExadataInfrastructureCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExternalExadataInfrastructureCompartmentResponse")
	}
	return
}

// changeExternalExadataInfrastructureCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeExternalExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataInfrastructures/{externalExadataInfrastructureId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalExadataInfrastructureCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/ChangeExternalExadataInfrastructureCompartment"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeExternalExadataInfrastructureCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeJobCompartment Moves a job.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Job/ChangeJobCompartment"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeJobCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeManagedDatabaseGroupCompartment Moves a Managed Database Group to a different compartment.
// The destination compartment must not have a Managed Database Group
// with the same name.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/ChangeManagedDatabaseGroupCompartment"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeManagedDatabaseGroupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePlanRetention Changes the retention period of unused plans. The period can range
// between 5 and 523 weeks.
// The database purges plans that have not been used for longer than
// the plan retention period.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangePlanRetention.go.html to see an example of how to use ChangePlanRetention API.
func (client DbManagementClient) ChangePlanRetention(ctx context.Context, request ChangePlanRetentionRequest) (response ChangePlanRetentionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changePlanRetention, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePlanRetentionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePlanRetentionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePlanRetentionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePlanRetentionResponse")
	}
	return
}

// changePlanRetention implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changePlanRetention(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/changePlanRetention", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePlanRetentionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ChangePlanRetention"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangePlanRetention", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSpaceBudget Changes the disk space limit for the SQL Management Base. The allowable
// range for this limit is between 1% and 50%.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeSpaceBudget.go.html to see an example of how to use ChangeSpaceBudget API.
func (client DbManagementClient) ChangeSpaceBudget(ctx context.Context, request ChangeSpaceBudgetRequest) (response ChangeSpaceBudgetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeSpaceBudget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSpaceBudgetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSpaceBudgetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSpaceBudgetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSpaceBudgetResponse")
	}
	return
}

// changeSpaceBudget implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeSpaceBudget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/changeSpaceBudget", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSpaceBudgetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ChangeSpaceBudget"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeSpaceBudget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSqlPlanBaselinesAttributes Changes one or more attributes of a single plan or all plans associated with a SQL statement.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ChangeSqlPlanBaselinesAttributes.go.html to see an example of how to use ChangeSqlPlanBaselinesAttributes API.
func (client DbManagementClient) ChangeSqlPlanBaselinesAttributes(ctx context.Context, request ChangeSqlPlanBaselinesAttributesRequest) (response ChangeSqlPlanBaselinesAttributesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeSqlPlanBaselinesAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSqlPlanBaselinesAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSqlPlanBaselinesAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSqlPlanBaselinesAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSqlPlanBaselinesAttributesResponse")
	}
	return
}

// changeSqlPlanBaselinesAttributes implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) changeSqlPlanBaselinesAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/changeSqlPlanBaselinesAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSqlPlanBaselinesAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ChangeSqlPlanBaselinesAttributes"
		err = common.PostProcessServiceError(err, "DbManagement", "ChangeSqlPlanBaselinesAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CheckExternalDbSystemConnectorConnectionStatus Checks the status of the external DB system component connection specified in this connector.
// This operation will refresh the connectionStatus and timeConnectionStatusLastUpdated fields.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CheckExternalDbSystemConnectorConnectionStatus.go.html to see an example of how to use CheckExternalDbSystemConnectorConnectionStatus API.
// A default retry strategy applies to this operation CheckExternalDbSystemConnectorConnectionStatus()
func (client DbManagementClient) CheckExternalDbSystemConnectorConnectionStatus(ctx context.Context, request CheckExternalDbSystemConnectorConnectionStatusRequest) (response CheckExternalDbSystemConnectorConnectionStatusResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.checkExternalDbSystemConnectorConnectionStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CheckExternalDbSystemConnectorConnectionStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CheckExternalDbSystemConnectorConnectionStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CheckExternalDbSystemConnectorConnectionStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CheckExternalDbSystemConnectorConnectionStatusResponse")
	}
	return
}

// checkExternalDbSystemConnectorConnectionStatus implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) checkExternalDbSystemConnectorConnectionStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystemConnectors/{externalDbSystemConnectorId}/actions/checkConnectionStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CheckExternalDbSystemConnectorConnectionStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemConnector/CheckExternalDbSystemConnectorConnectionStatus"
		err = common.PostProcessServiceError(err, "DbManagement", "CheckExternalDbSystemConnectorConnectionStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldbsystemconnector{})
	return response, err
}

// CheckExternalExadataStorageConnector Checks the status of the Exadata storage server connection specified by exadataStorageConnectorId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CheckExternalExadataStorageConnector.go.html to see an example of how to use CheckExternalExadataStorageConnector API.
// A default retry strategy applies to this operation CheckExternalExadataStorageConnector()
func (client DbManagementClient) CheckExternalExadataStorageConnector(ctx context.Context, request CheckExternalExadataStorageConnectorRequest) (response CheckExternalExadataStorageConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.checkExternalExadataStorageConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CheckExternalExadataStorageConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CheckExternalExadataStorageConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CheckExternalExadataStorageConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CheckExternalExadataStorageConnectorResponse")
	}
	return
}

// checkExternalExadataStorageConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) checkExternalExadataStorageConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataStorageConnectors/{externalExadataStorageConnectorId}/actions/checkStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CheckExternalExadataStorageConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageConnector/CheckExternalExadataStorageConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "CheckExternalExadataStorageConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureAutomaticCaptureFilters Configures automatic capture filters to capture only those statements
// that match the filter criteria.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ConfigureAutomaticCaptureFilters.go.html to see an example of how to use ConfigureAutomaticCaptureFilters API.
func (client DbManagementClient) ConfigureAutomaticCaptureFilters(ctx context.Context, request ConfigureAutomaticCaptureFiltersRequest) (response ConfigureAutomaticCaptureFiltersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.configureAutomaticCaptureFilters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureAutomaticCaptureFiltersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureAutomaticCaptureFiltersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureAutomaticCaptureFiltersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureAutomaticCaptureFiltersResponse")
	}
	return
}

// configureAutomaticCaptureFilters implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) configureAutomaticCaptureFilters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/configureAutomaticCaptureFilters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureAutomaticCaptureFiltersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ConfigureAutomaticCaptureFilters"
		err = common.PostProcessServiceError(err, "DbManagement", "ConfigureAutomaticCaptureFilters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureAutomaticSpmEvolveAdvisorTask Configures the Automatic SPM Evolve Advisor task `SYS_AUTO_SPM_EVOLVE_TASK`
// by specifying task parameters. As the task is owned by `SYS`, only `SYS` can
// set task parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ConfigureAutomaticSpmEvolveAdvisorTask.go.html to see an example of how to use ConfigureAutomaticSpmEvolveAdvisorTask API.
func (client DbManagementClient) ConfigureAutomaticSpmEvolveAdvisorTask(ctx context.Context, request ConfigureAutomaticSpmEvolveAdvisorTaskRequest) (response ConfigureAutomaticSpmEvolveAdvisorTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.configureAutomaticSpmEvolveAdvisorTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureAutomaticSpmEvolveAdvisorTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureAutomaticSpmEvolveAdvisorTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureAutomaticSpmEvolveAdvisorTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureAutomaticSpmEvolveAdvisorTaskResponse")
	}
	return
}

// configureAutomaticSpmEvolveAdvisorTask implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) configureAutomaticSpmEvolveAdvisorTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/configureAutomaticSpmEvolveAdvisorTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureAutomaticSpmEvolveAdvisorTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ConfigureAutomaticSpmEvolveAdvisorTask"
		err = common.PostProcessServiceError(err, "DbManagement", "ConfigureAutomaticSpmEvolveAdvisorTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDbManagementPrivateEndpoint Creates a new Database Management private endpoint.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DbManagementPrivateEndpoint/CreateDbManagementPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateDbManagementPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalDbSystem Creates an external DB system and its related resources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateExternalDbSystem.go.html to see an example of how to use CreateExternalDbSystem API.
// A default retry strategy applies to this operation CreateExternalDbSystem()
func (client DbManagementClient) CreateExternalDbSystem(ctx context.Context, request CreateExternalDbSystemRequest) (response CreateExternalDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalDbSystemResponse")
	}
	return
}

// createExternalDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createExternalDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/CreateExternalDbSystem"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateExternalDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalDbSystemConnector Creates a new external connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateExternalDbSystemConnector.go.html to see an example of how to use CreateExternalDbSystemConnector API.
// A default retry strategy applies to this operation CreateExternalDbSystemConnector()
func (client DbManagementClient) CreateExternalDbSystemConnector(ctx context.Context, request CreateExternalDbSystemConnectorRequest) (response CreateExternalDbSystemConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalDbSystemConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalDbSystemConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalDbSystemConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalDbSystemConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalDbSystemConnectorResponse")
	}
	return
}

// createExternalDbSystemConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createExternalDbSystemConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystemConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalDbSystemConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemConnector/CreateExternalDbSystemConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateExternalDbSystemConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldbsystemconnector{})
	return response, err
}

// CreateExternalDbSystemDiscovery Creates an external DB system discovery resource and initiates the discovery process.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateExternalDbSystemDiscovery.go.html to see an example of how to use CreateExternalDbSystemDiscovery API.
// A default retry strategy applies to this operation CreateExternalDbSystemDiscovery()
func (client DbManagementClient) CreateExternalDbSystemDiscovery(ctx context.Context, request CreateExternalDbSystemDiscoveryRequest) (response CreateExternalDbSystemDiscoveryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalDbSystemDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalDbSystemDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalDbSystemDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalDbSystemDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalDbSystemDiscoveryResponse")
	}
	return
}

// createExternalDbSystemDiscovery implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createExternalDbSystemDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystemDiscoveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalDbSystemDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemDiscovery/CreateExternalDbSystemDiscovery"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateExternalDbSystemDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalExadataInfrastructure Creates an OCI resource for the Exadata infrastructure and enables the Monitoring service for the Exadata infrastructure.
// The following resource/subresources are created:
//
//	Infrastructure
//	Storage server connectors
//	Storage servers
//	Storage grids
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateExternalExadataInfrastructure.go.html to see an example of how to use CreateExternalExadataInfrastructure API.
// A default retry strategy applies to this operation CreateExternalExadataInfrastructure()
func (client DbManagementClient) CreateExternalExadataInfrastructure(ctx context.Context, request CreateExternalExadataInfrastructureRequest) (response CreateExternalExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalExadataInfrastructureResponse")
	}
	return
}

// createExternalExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createExternalExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/CreateExternalExadataInfrastructure"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateExternalExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalExadataStorageConnector Creates the Exadata storage server connector after validating the connection information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateExternalExadataStorageConnector.go.html to see an example of how to use CreateExternalExadataStorageConnector API.
// A default retry strategy applies to this operation CreateExternalExadataStorageConnector()
func (client DbManagementClient) CreateExternalExadataStorageConnector(ctx context.Context, request CreateExternalExadataStorageConnectorRequest) (response CreateExternalExadataStorageConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalExadataStorageConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalExadataStorageConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalExadataStorageConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalExadataStorageConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalExadataStorageConnectorResponse")
	}
	return
}

// createExternalExadataStorageConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createExternalExadataStorageConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataStorageConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalExadataStorageConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageConnector/CreateExternalExadataStorageConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateExternalExadataStorageConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJob Creates a job to be executed on a Managed Database or Managed Database Group. Only one
// of the parameters, managedDatabaseId or managedDatabaseGroupId should be provided as
// input in CreateJobDetails resource in request body.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Job/CreateJob"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &job{})
	return response, err
}

// CreateManagedDatabaseGroup Creates a Managed Database Group. The group does not contain any
// Managed Databases when it is created, and they must be added later.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/CreateManagedDatabaseGroup"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateManagedDatabaseGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTablespace Creates a tablespace within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateTablespace.go.html to see an example of how to use CreateTablespace API.
func (client DbManagementClient) CreateTablespace(ctx context.Context, request CreateTablespaceRequest) (response CreateTablespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTablespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTablespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTablespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTablespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTablespaceResponse")
	}
	return
}

// createTablespace implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) createTablespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/tablespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTablespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/CreateTablespace"
		err = common.PostProcessServiceError(err, "DbManagement", "CreateTablespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDbManagementPrivateEndpoint Deletes a specific Database Management private endpoint.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DbManagementPrivateEndpoint/DeleteDbManagementPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteDbManagementPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalDbSystem Deletes the external DB system specified by `externalDbSystemId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteExternalDbSystem.go.html to see an example of how to use DeleteExternalDbSystem API.
func (client DbManagementClient) DeleteExternalDbSystem(ctx context.Context, request DeleteExternalDbSystemRequest) (response DeleteExternalDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalDbSystemResponse")
	}
	return
}

// deleteExternalDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteExternalDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalDbSystems/{externalDbSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/DeleteExternalDbSystem"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteExternalDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalDbSystemConnector Deletes the external connector specified by `externalDbSystemConnectorId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteExternalDbSystemConnector.go.html to see an example of how to use DeleteExternalDbSystemConnector API.
func (client DbManagementClient) DeleteExternalDbSystemConnector(ctx context.Context, request DeleteExternalDbSystemConnectorRequest) (response DeleteExternalDbSystemConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalDbSystemConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalDbSystemConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalDbSystemConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalDbSystemConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalDbSystemConnectorResponse")
	}
	return
}

// deleteExternalDbSystemConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteExternalDbSystemConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalDbSystemConnectors/{externalDbSystemConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalDbSystemConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemConnector/DeleteExternalDbSystemConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteExternalDbSystemConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalDbSystemDiscovery Deletes the external DB system discovery resource specified by `externalDbSystemDiscoveryId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteExternalDbSystemDiscovery.go.html to see an example of how to use DeleteExternalDbSystemDiscovery API.
func (client DbManagementClient) DeleteExternalDbSystemDiscovery(ctx context.Context, request DeleteExternalDbSystemDiscoveryRequest) (response DeleteExternalDbSystemDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalDbSystemDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalDbSystemDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalDbSystemDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalDbSystemDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalDbSystemDiscoveryResponse")
	}
	return
}

// deleteExternalDbSystemDiscovery implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteExternalDbSystemDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalDbSystemDiscoveries/{externalDbSystemDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalDbSystemDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemDiscovery/DeleteExternalDbSystemDiscovery"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteExternalDbSystemDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalExadataInfrastructure Deletes the Exadata infrastructure specified by externalExadataInfrastructureId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteExternalExadataInfrastructure.go.html to see an example of how to use DeleteExternalExadataInfrastructure API.
func (client DbManagementClient) DeleteExternalExadataInfrastructure(ctx context.Context, request DeleteExternalExadataInfrastructureRequest) (response DeleteExternalExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalExadataInfrastructureResponse")
	}
	return
}

// deleteExternalExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteExternalExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalExadataInfrastructures/{externalExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/DeleteExternalExadataInfrastructure"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteExternalExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalExadataStorageConnector Deletes the Exadata storage server connector specified by exadataStorageConnectorId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeleteExternalExadataStorageConnector.go.html to see an example of how to use DeleteExternalExadataStorageConnector API.
func (client DbManagementClient) DeleteExternalExadataStorageConnector(ctx context.Context, request DeleteExternalExadataStorageConnectorRequest) (response DeleteExternalExadataStorageConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalExadataStorageConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalExadataStorageConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalExadataStorageConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalExadataStorageConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalExadataStorageConnectorResponse")
	}
	return
}

// deleteExternalExadataStorageConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deleteExternalExadataStorageConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalExadataStorageConnectors/{externalExadataStorageConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalExadataStorageConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageConnector/DeleteExternalExadataStorageConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteExternalExadataStorageConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJob Deletes the job specified by jobId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Job/DeleteJob"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedDatabaseGroup Deletes the Managed Database Group specified by managedDatabaseGroupId.
// If the group contains Managed Databases, then it cannot be deleted.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/DeleteManagedDatabaseGroup"
		err = common.PostProcessServiceError(err, "DbManagement", "DeleteManagedDatabaseGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePreferredCredential Deletes the preferred credential based on the credentialName.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DeletePreferredCredential.go.html to see an example of how to use DeletePreferredCredential API.
func (client DbManagementClient) DeletePreferredCredential(ctx context.Context, request DeletePreferredCredentialRequest) (response DeletePreferredCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePreferredCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePreferredCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePreferredCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePreferredCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePreferredCredentialResponse")
	}
	return
}

// deletePreferredCredential implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) deletePreferredCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedDatabases/{managedDatabaseId}/preferredCredentials/{credentialName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePreferredCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/PreferredCredential/DeletePreferredCredential"
		err = common.PostProcessServiceError(err, "DbManagement", "DeletePreferredCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAutomaticInitialPlanCapture Disables automatic initial plan capture.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableAutomaticInitialPlanCapture.go.html to see an example of how to use DisableAutomaticInitialPlanCapture API.
func (client DbManagementClient) DisableAutomaticInitialPlanCapture(ctx context.Context, request DisableAutomaticInitialPlanCaptureRequest) (response DisableAutomaticInitialPlanCaptureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableAutomaticInitialPlanCapture, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAutomaticInitialPlanCaptureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAutomaticInitialPlanCaptureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAutomaticInitialPlanCaptureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAutomaticInitialPlanCaptureResponse")
	}
	return
}

// disableAutomaticInitialPlanCapture implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableAutomaticInitialPlanCapture(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/disableAutomaticInitialPlanCapture", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAutomaticInitialPlanCaptureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/DisableAutomaticInitialPlanCapture"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableAutomaticInitialPlanCapture", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAutomaticSpmEvolveAdvisorTask Disables the Automatic SPM Evolve Advisor task.
// One client controls both Automatic SQL Tuning Advisor and Automatic SPM Evolve Advisor.
// Thus, the same task enables or disables both.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableAutomaticSpmEvolveAdvisorTask.go.html to see an example of how to use DisableAutomaticSpmEvolveAdvisorTask API.
func (client DbManagementClient) DisableAutomaticSpmEvolveAdvisorTask(ctx context.Context, request DisableAutomaticSpmEvolveAdvisorTaskRequest) (response DisableAutomaticSpmEvolveAdvisorTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableAutomaticSpmEvolveAdvisorTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAutomaticSpmEvolveAdvisorTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAutomaticSpmEvolveAdvisorTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAutomaticSpmEvolveAdvisorTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAutomaticSpmEvolveAdvisorTaskResponse")
	}
	return
}

// disableAutomaticSpmEvolveAdvisorTask implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableAutomaticSpmEvolveAdvisorTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/disableAutomaticSpmEvolveAdvisorTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAutomaticSpmEvolveAdvisorTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/DisableAutomaticSpmEvolveAdvisorTask"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableAutomaticSpmEvolveAdvisorTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalDbSystemDatabaseManagement Disables Database Management service for all the components of the specified
// external DB system (except databases).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableExternalDbSystemDatabaseManagement.go.html to see an example of how to use DisableExternalDbSystemDatabaseManagement API.
// A default retry strategy applies to this operation DisableExternalDbSystemDatabaseManagement()
func (client DbManagementClient) DisableExternalDbSystemDatabaseManagement(ctx context.Context, request DisableExternalDbSystemDatabaseManagementRequest) (response DisableExternalDbSystemDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalDbSystemDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalDbSystemDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalDbSystemDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalDbSystemDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalDbSystemDatabaseManagementResponse")
	}
	return
}

// disableExternalDbSystemDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableExternalDbSystemDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystems/{externalDbSystemId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalDbSystemDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/DisableExternalDbSystemDatabaseManagement"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableExternalDbSystemDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalDbSystemStackMonitoring Disables Stack Monitoring for all the components of the specified
// external DB system (except databases).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableExternalDbSystemStackMonitoring.go.html to see an example of how to use DisableExternalDbSystemStackMonitoring API.
// A default retry strategy applies to this operation DisableExternalDbSystemStackMonitoring()
func (client DbManagementClient) DisableExternalDbSystemStackMonitoring(ctx context.Context, request DisableExternalDbSystemStackMonitoringRequest) (response DisableExternalDbSystemStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalDbSystemStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalDbSystemStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalDbSystemStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalDbSystemStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalDbSystemStackMonitoringResponse")
	}
	return
}

// disableExternalDbSystemStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableExternalDbSystemStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystems/{externalDbSystemId}/actions/disableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalDbSystemStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/DisableExternalDbSystemStackMonitoring"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableExternalDbSystemStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalExadataInfrastructureManagement Disables Database Management for the Exadata infrastructure specified by externalExadataInfrastructureId.
// It covers the following components:
// - Exadata infrastructure
// - Exadata storage grid
// - Exadata storage server
// Note that Database Management will not be disabled for the DB systems within the Exadata infrastructure and should be disabled explicitly, if required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableExternalExadataInfrastructureManagement.go.html to see an example of how to use DisableExternalExadataInfrastructureManagement API.
// A default retry strategy applies to this operation DisableExternalExadataInfrastructureManagement()
func (client DbManagementClient) DisableExternalExadataInfrastructureManagement(ctx context.Context, request DisableExternalExadataInfrastructureManagementRequest) (response DisableExternalExadataInfrastructureManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalExadataInfrastructureManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalExadataInfrastructureManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalExadataInfrastructureManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalExadataInfrastructureManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalExadataInfrastructureManagementResponse")
	}
	return
}

// disableExternalExadataInfrastructureManagement implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableExternalExadataInfrastructureManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataInfrastructures/{externalExadataInfrastructureId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalExadataInfrastructureManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/DisableExternalExadataInfrastructureManagement"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableExternalExadataInfrastructureManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableHighFrequencyAutomaticSpmEvolveAdvisorTask Disables the high-frequency Automatic SPM Evolve Advisor task.
// It is available only on Oracle Exadata Database Machine, Oracle Database Exadata
// Cloud Service (ExaCS) and Oracle Database Exadata Cloud@Customer (ExaCC).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableHighFrequencyAutomaticSpmEvolveAdvisorTask.go.html to see an example of how to use DisableHighFrequencyAutomaticSpmEvolveAdvisorTask API.
func (client DbManagementClient) DisableHighFrequencyAutomaticSpmEvolveAdvisorTask(ctx context.Context, request DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) (response DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableHighFrequencyAutomaticSpmEvolveAdvisorTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse")
	}
	return
}

// disableHighFrequencyAutomaticSpmEvolveAdvisorTask implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableHighFrequencyAutomaticSpmEvolveAdvisorTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/disableHighFrequencyAutomaticSpmEvolveAdvisorTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/DisableHighFrequencyAutomaticSpmEvolveAdvisorTask"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableHighFrequencyAutomaticSpmEvolveAdvisorTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableSqlPlanBaselinesUsage Disables the use of SQL plan baselines stored in SQL Management Base.
// When disabled, the optimizer does not use any SQL plan baselines.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DisableSqlPlanBaselinesUsage.go.html to see an example of how to use DisableSqlPlanBaselinesUsage API.
func (client DbManagementClient) DisableSqlPlanBaselinesUsage(ctx context.Context, request DisableSqlPlanBaselinesUsageRequest) (response DisableSqlPlanBaselinesUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableSqlPlanBaselinesUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableSqlPlanBaselinesUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableSqlPlanBaselinesUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableSqlPlanBaselinesUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableSqlPlanBaselinesUsageResponse")
	}
	return
}

// disableSqlPlanBaselinesUsage implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) disableSqlPlanBaselinesUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/disableSqlPlanBaselinesUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableSqlPlanBaselinesUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/DisableSqlPlanBaselinesUsage"
		err = common.PostProcessServiceError(err, "DbManagement", "DisableSqlPlanBaselinesUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DiscoverExternalExadataInfrastructure Completes the Exadata system prechecking on the following:
// - Verifies if the DB systems are valid RAC DB systems or return 400 status code with NON_RAC_DATABASE_SYSTEM error code.
// - Verifies if the ASM connector defined for each DB system or return 400 status code with CONNECTOR_NOT_DEFINED error code.
// - Verifies if the agents associated with ASM are valid and could be used for the Exadata storage servers or return 400 status code with
// INVALID_AGENT error code.
// - Verifies if it is an Exadata system or return 400 status code with INVALID_EXADATA_SYSTEM error code.
// Starts the discovery process for the Exadata system infrastructure. The following resources/components are discovered
// - Exadata storage servers from each DB systems
// - Exadata storage grid for all Exadata storage servers
// - Exadata infrastructure
// The same API covers both new discovery and rediscovery cases.
//
//	For the new discovery case, new managed resources/sub-resources are created or the existing ones are overridden.
//	For rediscovery case, the existing managed resources/sub-resources are checked to find out which ones should be added or which ones
//
// should be
//
//	removed based on the unique key defined for each resource/sub-resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DiscoverExternalExadataInfrastructure.go.html to see an example of how to use DiscoverExternalExadataInfrastructure API.
// A default retry strategy applies to this operation DiscoverExternalExadataInfrastructure()
func (client DbManagementClient) DiscoverExternalExadataInfrastructure(ctx context.Context, request DiscoverExternalExadataInfrastructureRequest) (response DiscoverExternalExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.discoverExternalExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DiscoverExternalExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DiscoverExternalExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DiscoverExternalExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DiscoverExternalExadataInfrastructureResponse")
	}
	return
}

// discoverExternalExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) discoverExternalExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataInfrastructures/actions/discoverExadataInfrastructure", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DiscoverExternalExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/DiscoverExternalExadataInfrastructure"
		err = common.PostProcessServiceError(err, "DbManagement", "DiscoverExternalExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DropSqlPlanBaselines Drops a single plan or all plans associated with a SQL statement.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DropSqlPlanBaselines.go.html to see an example of how to use DropSqlPlanBaselines API.
func (client DbManagementClient) DropSqlPlanBaselines(ctx context.Context, request DropSqlPlanBaselinesRequest) (response DropSqlPlanBaselinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.dropSqlPlanBaselines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DropSqlPlanBaselinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DropSqlPlanBaselinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DropSqlPlanBaselinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DropSqlPlanBaselinesResponse")
	}
	return
}

// dropSqlPlanBaselines implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) dropSqlPlanBaselines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/dropSqlPlanBaselines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DropSqlPlanBaselinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/DropSqlPlanBaselines"
		err = common.PostProcessServiceError(err, "DbManagement", "DropSqlPlanBaselines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DropTablespace Drops the tablespace specified by tablespaceName within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DropTablespace.go.html to see an example of how to use DropTablespace API.
func (client DbManagementClient) DropTablespace(ctx context.Context, request DropTablespaceRequest) (response DropTablespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.dropTablespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DropTablespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DropTablespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DropTablespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DropTablespaceResponse")
	}
	return
}

// dropTablespace implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) dropTablespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/tablespaces/{tablespaceName}/actions/dropTablespace", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DropTablespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/DropTablespace"
		err = common.PostProcessServiceError(err, "DbManagement", "DropTablespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAutomaticInitialPlanCapture Enables automatic initial plan capture. When enabled, the database checks whether
// executed SQL statements are eligible for automatic capture. It creates initial
// plan baselines for eligible statements.
// By default, the database creates a SQL plan baseline for every eligible repeatable
// statement, including all recursive SQL and monitoring SQL. Thus, automatic capture
// may result in the creation of an extremely large number of plan baselines. To limit
// the statements that are eligible for plan baselines, configure filters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableAutomaticInitialPlanCapture.go.html to see an example of how to use EnableAutomaticInitialPlanCapture API.
func (client DbManagementClient) EnableAutomaticInitialPlanCapture(ctx context.Context, request EnableAutomaticInitialPlanCaptureRequest) (response EnableAutomaticInitialPlanCaptureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableAutomaticInitialPlanCapture, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAutomaticInitialPlanCaptureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAutomaticInitialPlanCaptureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAutomaticInitialPlanCaptureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAutomaticInitialPlanCaptureResponse")
	}
	return
}

// enableAutomaticInitialPlanCapture implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableAutomaticInitialPlanCapture(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/enableAutomaticInitialPlanCapture", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAutomaticInitialPlanCaptureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/EnableAutomaticInitialPlanCapture"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableAutomaticInitialPlanCapture", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAutomaticSpmEvolveAdvisorTask Enables the Automatic SPM Evolve Advisor task. By default, the automatic task
// `SYS_AUTO_SPM_EVOLVE_TASK` runs every day in the scheduled maintenance window.
// The SPM Evolve Advisor performs the following tasks:
// - Checks AWR for top SQL
// - Looks for alternative plans in all available sources
// - Adds unaccepted plans to the plan history
// - Tests the execution of as many plans as possible during the maintenance window
// - Adds the alternative plan to the baseline if it performs better than the current plan
// One client controls both Automatic SQL Tuning Advisor and Automatic SPM Evolve Advisor.
// Thus, the same task enables or disables both.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableAutomaticSpmEvolveAdvisorTask.go.html to see an example of how to use EnableAutomaticSpmEvolveAdvisorTask API.
func (client DbManagementClient) EnableAutomaticSpmEvolveAdvisorTask(ctx context.Context, request EnableAutomaticSpmEvolveAdvisorTaskRequest) (response EnableAutomaticSpmEvolveAdvisorTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableAutomaticSpmEvolveAdvisorTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAutomaticSpmEvolveAdvisorTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAutomaticSpmEvolveAdvisorTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAutomaticSpmEvolveAdvisorTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAutomaticSpmEvolveAdvisorTaskResponse")
	}
	return
}

// enableAutomaticSpmEvolveAdvisorTask implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableAutomaticSpmEvolveAdvisorTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/enableAutomaticSpmEvolveAdvisorTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAutomaticSpmEvolveAdvisorTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/EnableAutomaticSpmEvolveAdvisorTask"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableAutomaticSpmEvolveAdvisorTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalDbSystemDatabaseManagement Enables Database Management service for all the components of the specified
// external DB system (except databases).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableExternalDbSystemDatabaseManagement.go.html to see an example of how to use EnableExternalDbSystemDatabaseManagement API.
// A default retry strategy applies to this operation EnableExternalDbSystemDatabaseManagement()
func (client DbManagementClient) EnableExternalDbSystemDatabaseManagement(ctx context.Context, request EnableExternalDbSystemDatabaseManagementRequest) (response EnableExternalDbSystemDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalDbSystemDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalDbSystemDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalDbSystemDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalDbSystemDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalDbSystemDatabaseManagementResponse")
	}
	return
}

// enableExternalDbSystemDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableExternalDbSystemDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystems/{externalDbSystemId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalDbSystemDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/EnableExternalDbSystemDatabaseManagement"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableExternalDbSystemDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalDbSystemStackMonitoring Enables Stack Monitoring for all the components of the specified
// external DB system (except databases).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableExternalDbSystemStackMonitoring.go.html to see an example of how to use EnableExternalDbSystemStackMonitoring API.
// A default retry strategy applies to this operation EnableExternalDbSystemStackMonitoring()
func (client DbManagementClient) EnableExternalDbSystemStackMonitoring(ctx context.Context, request EnableExternalDbSystemStackMonitoringRequest) (response EnableExternalDbSystemStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalDbSystemStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalDbSystemStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalDbSystemStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalDbSystemStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalDbSystemStackMonitoringResponse")
	}
	return
}

// enableExternalDbSystemStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableExternalDbSystemStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalDbSystems/{externalDbSystemId}/actions/enableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalDbSystemStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/EnableExternalDbSystemStackMonitoring"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableExternalDbSystemStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalExadataInfrastructureManagement Enables Database Management for the Exadata infrastructure specified by externalExadataInfrastructureId. It covers the following
// components:
// - Exadata infrastructure
// - Exadata storage grid
// - Exadata storage server
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableExternalExadataInfrastructureManagement.go.html to see an example of how to use EnableExternalExadataInfrastructureManagement API.
// A default retry strategy applies to this operation EnableExternalExadataInfrastructureManagement()
func (client DbManagementClient) EnableExternalExadataInfrastructureManagement(ctx context.Context, request EnableExternalExadataInfrastructureManagementRequest) (response EnableExternalExadataInfrastructureManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalExadataInfrastructureManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalExadataInfrastructureManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalExadataInfrastructureManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalExadataInfrastructureManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalExadataInfrastructureManagementResponse")
	}
	return
}

// enableExternalExadataInfrastructureManagement implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableExternalExadataInfrastructureManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalExadataInfrastructures/{externalExadataInfrastructureId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalExadataInfrastructureManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/EnableExternalExadataInfrastructureManagement"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableExternalExadataInfrastructureManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableHighFrequencyAutomaticSpmEvolveAdvisorTask Enables the high-frequency Automatic SPM Evolve Advisor task. The high-frequency
// task runs every hour and runs for no longer than 30 minutes. These settings
// are not configurable.
// The high-frequency task complements the standard Automatic SPM Evolve Advisor task.
// They are independent and are scheduled through two different frameworks.
// It is available only on Oracle Exadata Database Machine, Oracle Database Exadata
// Cloud Service (ExaCS) and Oracle Database Exadata Cloud@Customer (ExaCC).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableHighFrequencyAutomaticSpmEvolveAdvisorTask.go.html to see an example of how to use EnableHighFrequencyAutomaticSpmEvolveAdvisorTask API.
func (client DbManagementClient) EnableHighFrequencyAutomaticSpmEvolveAdvisorTask(ctx context.Context, request EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) (response EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableHighFrequencyAutomaticSpmEvolveAdvisorTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse")
	}
	return
}

// enableHighFrequencyAutomaticSpmEvolveAdvisorTask implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableHighFrequencyAutomaticSpmEvolveAdvisorTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/enableHighFrequencyAutomaticSpmEvolveAdvisorTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/EnableHighFrequencyAutomaticSpmEvolveAdvisorTask"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableHighFrequencyAutomaticSpmEvolveAdvisorTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableSqlPlanBaselinesUsage Enables the use of SQL plan baselines stored in SQL Management Base.
// When enabled, the optimizer uses SQL plan baselines to select plans
// to avoid potential performance regressions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableSqlPlanBaselinesUsage.go.html to see an example of how to use EnableSqlPlanBaselinesUsage API.
func (client DbManagementClient) EnableSqlPlanBaselinesUsage(ctx context.Context, request EnableSqlPlanBaselinesUsageRequest) (response EnableSqlPlanBaselinesUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableSqlPlanBaselinesUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableSqlPlanBaselinesUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableSqlPlanBaselinesUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableSqlPlanBaselinesUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableSqlPlanBaselinesUsageResponse")
	}
	return
}

// enableSqlPlanBaselinesUsage implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) enableSqlPlanBaselinesUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/enableSqlPlanBaselinesUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableSqlPlanBaselinesUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/EnableSqlPlanBaselinesUsage"
		err = common.PostProcessServiceError(err, "DbManagement", "EnableSqlPlanBaselinesUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateAwrSnapshot Creates an AWR snapshot for the target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GenerateAwrSnapshot.go.html to see an example of how to use GenerateAwrSnapshot API.
func (client DbManagementClient) GenerateAwrSnapshot(ctx context.Context, request GenerateAwrSnapshotRequest) (response GenerateAwrSnapshotResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateAwrSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateAwrSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateAwrSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateAwrSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateAwrSnapshotResponse")
	}
	return
}

// generateAwrSnapshot implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) generateAwrSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/generateAwrSnapshot", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateAwrSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SnapshotDetails/GenerateAwrSnapshot"
		err = common.PostProcessServiceError(err, "DbManagement", "GenerateAwrSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrDbReport Gets the AWR report for the specific database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetAwrDbReport"
		err = common.PostProcessServiceError(err, "DbManagement", "GetAwrDbReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrDbSqlReport Gets the SQL health check report for one SQL of the specific database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetAwrDbSqlReport"
		err = common.PostProcessServiceError(err, "DbManagement", "GetAwrDbSqlReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterCacheMetric Gets the metrics related to cluster cache for the Oracle
// Real Application Clusters (Oracle RAC) database specified
// by managedDatabaseId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ClusterCacheMetric/GetClusterCacheMetric"
		err = common.PostProcessServiceError(err, "DbManagement", "GetClusterCacheMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseFleetHealthMetrics Gets the health metrics for a fleet of databases in a compartment or in a Managed Database Group.
// Either the CompartmentId or the ManagedDatabaseGroupId query parameters must be provided to retrieve the health metrics.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DatabaseFleetHealthMetrics/GetDatabaseFleetHealthMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "GetDatabaseFleetHealthMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseHomeMetrics Gets a summary of the activity and resource usage metrics like DB Time, CPU, User I/O, Wait, Storage, and Memory for a Managed Database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DatabaseHomeMetrics/GetDatabaseHomeMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "GetDatabaseHomeMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbManagementPrivateEndpoint Gets the details of a specific Database Management private endpoint.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DbManagementPrivateEndpoint/GetDbManagementPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DbManagement", "GetDbManagementPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalAsm Gets the details for the external ASM specified by `externalAsmId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalAsm.go.html to see an example of how to use GetExternalAsm API.
// A default retry strategy applies to this operation GetExternalAsm()
func (client DbManagementClient) GetExternalAsm(ctx context.Context, request GetExternalAsmRequest) (response GetExternalAsmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalAsm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalAsmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalAsmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalAsmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalAsmResponse")
	}
	return
}

// getExternalAsm implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalAsm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsms/{externalAsmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalAsmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/GetExternalAsm"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalAsm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalAsmConfiguration Gets configuration details including disk groups for the external ASM specified by `externalAsmId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalAsmConfiguration.go.html to see an example of how to use GetExternalAsmConfiguration API.
// A default retry strategy applies to this operation GetExternalAsmConfiguration()
func (client DbManagementClient) GetExternalAsmConfiguration(ctx context.Context, request GetExternalAsmConfigurationRequest) (response GetExternalAsmConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalAsmConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalAsmConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalAsmConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalAsmConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalAsmConfigurationResponse")
	}
	return
}

// getExternalAsmConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalAsmConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsms/{externalAsmId}/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalAsmConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/GetExternalAsmConfiguration"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalAsmConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalAsmInstance Gets the details for the external ASM instance specified by `externalAsmInstanceId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalAsmInstance.go.html to see an example of how to use GetExternalAsmInstance API.
// A default retry strategy applies to this operation GetExternalAsmInstance()
func (client DbManagementClient) GetExternalAsmInstance(ctx context.Context, request GetExternalAsmInstanceRequest) (response GetExternalAsmInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalAsmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalAsmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalAsmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalAsmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalAsmInstanceResponse")
	}
	return
}

// getExternalAsmInstance implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalAsmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsmInstances/{externalAsmInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalAsmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsmInstance/GetExternalAsmInstance"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalAsmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalCluster Gets the details for the external cluster specified by `externalClusterId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalCluster.go.html to see an example of how to use GetExternalCluster API.
// A default retry strategy applies to this operation GetExternalCluster()
func (client DbManagementClient) GetExternalCluster(ctx context.Context, request GetExternalClusterRequest) (response GetExternalClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalClusterResponse")
	}
	return
}

// getExternalCluster implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalClusters/{externalClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalCluster/GetExternalCluster"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalClusterInstance Gets the details for the external cluster instance specified by `externalClusterInstanceId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalClusterInstance.go.html to see an example of how to use GetExternalClusterInstance API.
// A default retry strategy applies to this operation GetExternalClusterInstance()
func (client DbManagementClient) GetExternalClusterInstance(ctx context.Context, request GetExternalClusterInstanceRequest) (response GetExternalClusterInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalClusterInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalClusterInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalClusterInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalClusterInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalClusterInstanceResponse")
	}
	return
}

// getExternalClusterInstance implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalClusterInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalClusterInstances/{externalClusterInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalClusterInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalClusterInstance/GetExternalClusterInstance"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalClusterInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalDbHome Gets the details for the external DB home specified by `externalDbHomeId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalDbHome.go.html to see an example of how to use GetExternalDbHome API.
// A default retry strategy applies to this operation GetExternalDbHome()
func (client DbManagementClient) GetExternalDbHome(ctx context.Context, request GetExternalDbHomeRequest) (response GetExternalDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalDbHomeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalDbHomeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalDbHomeResponse")
	}
	return
}

// getExternalDbHome implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbHomes/{externalDbHomeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalDbHomeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbHome/GetExternalDbHome"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalDbHome", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalDbNode Gets the details for the external DB node specified by `externalDbNodeId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalDbNode.go.html to see an example of how to use GetExternalDbNode API.
// A default retry strategy applies to this operation GetExternalDbNode()
func (client DbManagementClient) GetExternalDbNode(ctx context.Context, request GetExternalDbNodeRequest) (response GetExternalDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalDbNodeResponse")
	}
	return
}

// getExternalDbNode implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbNodes/{externalDbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbNode/GetExternalDbNode"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalDbSystem Gets the details for the external DB system specified by `externalDbSystemId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalDbSystem.go.html to see an example of how to use GetExternalDbSystem API.
// A default retry strategy applies to this operation GetExternalDbSystem()
func (client DbManagementClient) GetExternalDbSystem(ctx context.Context, request GetExternalDbSystemRequest) (response GetExternalDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalDbSystemResponse")
	}
	return
}

// getExternalDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystems/{externalDbSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/GetExternalDbSystem"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalDbSystemConnector Gets the details for the external connector specified by `externalDbSystemConnectorId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalDbSystemConnector.go.html to see an example of how to use GetExternalDbSystemConnector API.
// A default retry strategy applies to this operation GetExternalDbSystemConnector()
func (client DbManagementClient) GetExternalDbSystemConnector(ctx context.Context, request GetExternalDbSystemConnectorRequest) (response GetExternalDbSystemConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalDbSystemConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalDbSystemConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalDbSystemConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalDbSystemConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalDbSystemConnectorResponse")
	}
	return
}

// getExternalDbSystemConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalDbSystemConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystemConnectors/{externalDbSystemConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalDbSystemConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemConnector/GetExternalDbSystemConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalDbSystemConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldbsystemconnector{})
	return response, err
}

// GetExternalDbSystemDiscovery Gets the details for the external DB system discovery resource specified by `externalDbSystemDiscoveryId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalDbSystemDiscovery.go.html to see an example of how to use GetExternalDbSystemDiscovery API.
// A default retry strategy applies to this operation GetExternalDbSystemDiscovery()
func (client DbManagementClient) GetExternalDbSystemDiscovery(ctx context.Context, request GetExternalDbSystemDiscoveryRequest) (response GetExternalDbSystemDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalDbSystemDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalDbSystemDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalDbSystemDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalDbSystemDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalDbSystemDiscoveryResponse")
	}
	return
}

// getExternalDbSystemDiscovery implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalDbSystemDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystemDiscoveries/{externalDbSystemDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalDbSystemDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemDiscovery/GetExternalDbSystemDiscovery"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalDbSystemDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalExadataInfrastructure Gets the details for the Exadata infrastructure specified by externalExadataInfrastructureId. It includes the DB systems and storage grid within the
// Exadata infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalExadataInfrastructure.go.html to see an example of how to use GetExternalExadataInfrastructure API.
// A default retry strategy applies to this operation GetExternalExadataInfrastructure()
func (client DbManagementClient) GetExternalExadataInfrastructure(ctx context.Context, request GetExternalExadataInfrastructureRequest) (response GetExternalExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalExadataInfrastructureResponse")
	}
	return
}

// getExternalExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataInfrastructures/{externalExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/GetExternalExadataInfrastructure"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalExadataStorageConnector Gets the details for the Exadata storage server connector specified by exadataStorageConnectorId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalExadataStorageConnector.go.html to see an example of how to use GetExternalExadataStorageConnector API.
// A default retry strategy applies to this operation GetExternalExadataStorageConnector()
func (client DbManagementClient) GetExternalExadataStorageConnector(ctx context.Context, request GetExternalExadataStorageConnectorRequest) (response GetExternalExadataStorageConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalExadataStorageConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalExadataStorageConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalExadataStorageConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalExadataStorageConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalExadataStorageConnectorResponse")
	}
	return
}

// getExternalExadataStorageConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalExadataStorageConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageConnectors/{externalExadataStorageConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalExadataStorageConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageConnector/GetExternalExadataStorageConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalExadataStorageConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalExadataStorageGrid Gets the details for the Exadata storage server grid specified by exadataStorageGridId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalExadataStorageGrid.go.html to see an example of how to use GetExternalExadataStorageGrid API.
// A default retry strategy applies to this operation GetExternalExadataStorageGrid()
func (client DbManagementClient) GetExternalExadataStorageGrid(ctx context.Context, request GetExternalExadataStorageGridRequest) (response GetExternalExadataStorageGridResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalExadataStorageGrid, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalExadataStorageGridResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalExadataStorageGridResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalExadataStorageGridResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalExadataStorageGridResponse")
	}
	return
}

// getExternalExadataStorageGrid implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalExadataStorageGrid(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageGrids/{externalExadataStorageGridId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalExadataStorageGridResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageGrid/GetExternalExadataStorageGrid"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalExadataStorageGrid", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalExadataStorageServer Gets the summary for the Exadata storage server specified by exadataStorageServerId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalExadataStorageServer.go.html to see an example of how to use GetExternalExadataStorageServer API.
// A default retry strategy applies to this operation GetExternalExadataStorageServer()
func (client DbManagementClient) GetExternalExadataStorageServer(ctx context.Context, request GetExternalExadataStorageServerRequest) (response GetExternalExadataStorageServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalExadataStorageServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalExadataStorageServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalExadataStorageServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalExadataStorageServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalExadataStorageServerResponse")
	}
	return
}

// getExternalExadataStorageServer implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalExadataStorageServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageServers/{externalExadataStorageServerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalExadataStorageServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageServer/GetExternalExadataStorageServer"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalExadataStorageServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalListener Gets the details for the external listener specified by `externalListenerId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExternalListener.go.html to see an example of how to use GetExternalListener API.
// A default retry strategy applies to this operation GetExternalListener()
func (client DbManagementClient) GetExternalListener(ctx context.Context, request GetExternalListenerRequest) (response GetExternalListenerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalListener, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalListenerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalListenerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalListenerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalListenerResponse")
	}
	return
}

// getExternalListener implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getExternalListener(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalListeners/{externalListenerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalListener/GetExternalListener"
		err = common.PostProcessServiceError(err, "DbManagement", "GetExternalListener", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIormPlan Get the IORM plan from the specific Exadata storage server.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetIormPlan.go.html to see an example of how to use GetIormPlan API.
// A default retry strategy applies to this operation GetIormPlan()
func (client DbManagementClient) GetIormPlan(ctx context.Context, request GetIormPlanRequest) (response GetIormPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIormPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIormPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIormPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIormPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIormPlanResponse")
	}
	return
}

// getIormPlan implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getIormPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageServers/{externalExadataStorageServerId}/iormPlan", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIormPlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageServer/GetIormPlan"
		err = common.PostProcessServiceError(err, "DbManagement", "GetIormPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJob Gets the details for the job specified by jobId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Job/GetJob"
		err = common.PostProcessServiceError(err, "DbManagement", "GetJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &job{})
	return response, err
}

// GetJobExecution Gets the details for the job execution specified by jobExecutionId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/JobExecution/GetJobExecution"
		err = common.PostProcessServiceError(err, "DbManagement", "GetJobExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobRun Gets the details for the job run specified by jobRunId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/JobRun/GetJobRun"
		err = common.PostProcessServiceError(err, "DbManagement", "GetJobRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedDatabase Gets the details for the Managed Database specified by managedDatabaseId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetManagedDatabase"
		err = common.PostProcessServiceError(err, "DbManagement", "GetManagedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedDatabaseGroup Gets the details for the Managed Database Group specified by managedDatabaseGroupId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/GetManagedDatabaseGroup"
		err = common.PostProcessServiceError(err, "DbManagement", "GetManagedDatabaseGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOpenAlertHistory Gets the open alerts from the specified Exadata storage server.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetOpenAlertHistory.go.html to see an example of how to use GetOpenAlertHistory API.
// A default retry strategy applies to this operation GetOpenAlertHistory()
func (client DbManagementClient) GetOpenAlertHistory(ctx context.Context, request GetOpenAlertHistoryRequest) (response GetOpenAlertHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOpenAlertHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOpenAlertHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOpenAlertHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOpenAlertHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOpenAlertHistoryResponse")
	}
	return
}

// getOpenAlertHistory implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getOpenAlertHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageServers/{externalExadataStorageServerId}/openAlertHistory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOpenAlertHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageServer/GetOpenAlertHistory"
		err = common.PostProcessServiceError(err, "DbManagement", "GetOpenAlertHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOptimizerStatisticsAdvisorExecution Gets a comprehensive report of the Optimizer Statistics Advisor execution, which includes details of the
// Managed Database, findings, recommendations, rationale, and examples.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetOptimizerStatisticsAdvisorExecution.go.html to see an example of how to use GetOptimizerStatisticsAdvisorExecution API.
func (client DbManagementClient) GetOptimizerStatisticsAdvisorExecution(ctx context.Context, request GetOptimizerStatisticsAdvisorExecutionRequest) (response GetOptimizerStatisticsAdvisorExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOptimizerStatisticsAdvisorExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOptimizerStatisticsAdvisorExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOptimizerStatisticsAdvisorExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOptimizerStatisticsAdvisorExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOptimizerStatisticsAdvisorExecutionResponse")
	}
	return
}

// getOptimizerStatisticsAdvisorExecution implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getOptimizerStatisticsAdvisorExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsAdvisorExecutions/{executionName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOptimizerStatisticsAdvisorExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetOptimizerStatisticsAdvisorExecution"
		err = common.PostProcessServiceError(err, "DbManagement", "GetOptimizerStatisticsAdvisorExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOptimizerStatisticsAdvisorExecutionScript Gets the Oracle system-generated script for the specified Optimizer Statistics Advisor execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetOptimizerStatisticsAdvisorExecutionScript.go.html to see an example of how to use GetOptimizerStatisticsAdvisorExecutionScript API.
func (client DbManagementClient) GetOptimizerStatisticsAdvisorExecutionScript(ctx context.Context, request GetOptimizerStatisticsAdvisorExecutionScriptRequest) (response GetOptimizerStatisticsAdvisorExecutionScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOptimizerStatisticsAdvisorExecutionScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOptimizerStatisticsAdvisorExecutionScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOptimizerStatisticsAdvisorExecutionScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOptimizerStatisticsAdvisorExecutionScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOptimizerStatisticsAdvisorExecutionScriptResponse")
	}
	return
}

// getOptimizerStatisticsAdvisorExecutionScript implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getOptimizerStatisticsAdvisorExecutionScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsAdvisorExecutions/{executionName}/script", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOptimizerStatisticsAdvisorExecutionScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetOptimizerStatisticsAdvisorExecutionScript"
		err = common.PostProcessServiceError(err, "DbManagement", "GetOptimizerStatisticsAdvisorExecutionScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOptimizerStatisticsCollectionOperation Gets a detailed report of the Optimizer Statistics Collection operation for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetOptimizerStatisticsCollectionOperation.go.html to see an example of how to use GetOptimizerStatisticsCollectionOperation API.
func (client DbManagementClient) GetOptimizerStatisticsCollectionOperation(ctx context.Context, request GetOptimizerStatisticsCollectionOperationRequest) (response GetOptimizerStatisticsCollectionOperationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOptimizerStatisticsCollectionOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOptimizerStatisticsCollectionOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOptimizerStatisticsCollectionOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOptimizerStatisticsCollectionOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOptimizerStatisticsCollectionOperationResponse")
	}
	return
}

// getOptimizerStatisticsCollectionOperation implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getOptimizerStatisticsCollectionOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsCollectionOperations/{optimizerStatisticsCollectionOperationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOptimizerStatisticsCollectionOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetOptimizerStatisticsCollectionOperation"
		err = common.PostProcessServiceError(err, "DbManagement", "GetOptimizerStatisticsCollectionOperation", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/PdbMetrics/GetPdbMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "GetPdbMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPreferredCredential Gets the preferred credential details for a Managed Database based on credentialName.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetPreferredCredential.go.html to see an example of how to use GetPreferredCredential API.
func (client DbManagementClient) GetPreferredCredential(ctx context.Context, request GetPreferredCredentialRequest) (response GetPreferredCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPreferredCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPreferredCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPreferredCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPreferredCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPreferredCredentialResponse")
	}
	return
}

// getPreferredCredential implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getPreferredCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/preferredCredentials/{credentialName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPreferredCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/PreferredCredential/GetPreferredCredential"
		err = common.PostProcessServiceError(err, "DbManagement", "GetPreferredCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &preferredcredential{})
	return response, err
}

// GetSqlPlanBaseline Gets the SQL plan baseline details for the specified planName.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlPlanBaseline.go.html to see an example of how to use GetSqlPlanBaseline API.
// A default retry strategy applies to this operation GetSqlPlanBaseline()
func (client DbManagementClient) GetSqlPlanBaseline(ctx context.Context, request GetSqlPlanBaselineRequest) (response GetSqlPlanBaselineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSqlPlanBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSqlPlanBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSqlPlanBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSqlPlanBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSqlPlanBaselineResponse")
	}
	return
}

// getSqlPlanBaseline implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getSqlPlanBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/{planName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSqlPlanBaselineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetSqlPlanBaseline"
		err = common.PostProcessServiceError(err, "DbManagement", "GetSqlPlanBaseline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSqlPlanBaselineConfiguration Gets the configuration details of SQL plan baselines for the specified
// Managed Database. The details include the settings for the capture and use of
// SQL plan baselines, SPM Evolve Advisor task, and SQL Management Base.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlPlanBaselineConfiguration.go.html to see an example of how to use GetSqlPlanBaselineConfiguration API.
// A default retry strategy applies to this operation GetSqlPlanBaselineConfiguration()
func (client DbManagementClient) GetSqlPlanBaselineConfiguration(ctx context.Context, request GetSqlPlanBaselineConfigurationRequest) (response GetSqlPlanBaselineConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSqlPlanBaselineConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSqlPlanBaselineConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSqlPlanBaselineConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSqlPlanBaselineConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSqlPlanBaselineConfigurationResponse")
	}
	return
}

// getSqlPlanBaselineConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getSqlPlanBaselineConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselineConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSqlPlanBaselineConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetSqlPlanBaselineConfiguration"
		err = common.PostProcessServiceError(err, "DbManagement", "GetSqlPlanBaselineConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTablespace Gets the details of the tablespace specified by tablespaceName within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetTablespace.go.html to see an example of how to use GetTablespace API.
func (client DbManagementClient) GetTablespace(ctx context.Context, request GetTablespaceRequest) (response GetTablespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTablespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTablespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTablespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTablespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTablespaceResponse")
	}
	return
}

// getTablespace implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getTablespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/tablespaces/{tablespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTablespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/GetTablespace"
		err = common.PostProcessServiceError(err, "DbManagement", "GetTablespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTopSqlCpuActivity Gets the SQL IDs with the top CPU activity from the Exadata storage server.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetTopSqlCpuActivity.go.html to see an example of how to use GetTopSqlCpuActivity API.
// A default retry strategy applies to this operation GetTopSqlCpuActivity()
func (client DbManagementClient) GetTopSqlCpuActivity(ctx context.Context, request GetTopSqlCpuActivityRequest) (response GetTopSqlCpuActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTopSqlCpuActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTopSqlCpuActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTopSqlCpuActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTopSqlCpuActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTopSqlCpuActivityResponse")
	}
	return
}

// getTopSqlCpuActivity implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) getTopSqlCpuActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageServers/{externalExadataStorageServerId}/topSqlCpuActivity", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTopSqlCpuActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageServer/GetTopSqlCpuActivity"
		err = common.PostProcessServiceError(err, "DbManagement", "GetTopSqlCpuActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUser Gets the details of the user specified by managedDatabaseId and userName.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetUser"
		err = common.PostProcessServiceError(err, "DbManagement", "GetUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given Work Request ID
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DbManagement", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImplementOptimizerStatisticsAdvisorRecommendations Asynchronously implements the findings and recommendations of the Optimizer Statistics Advisor execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ImplementOptimizerStatisticsAdvisorRecommendations.go.html to see an example of how to use ImplementOptimizerStatisticsAdvisorRecommendations API.
func (client DbManagementClient) ImplementOptimizerStatisticsAdvisorRecommendations(ctx context.Context, request ImplementOptimizerStatisticsAdvisorRecommendationsRequest) (response ImplementOptimizerStatisticsAdvisorRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.implementOptimizerStatisticsAdvisorRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImplementOptimizerStatisticsAdvisorRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImplementOptimizerStatisticsAdvisorRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImplementOptimizerStatisticsAdvisorRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImplementOptimizerStatisticsAdvisorRecommendationsResponse")
	}
	return
}

// implementOptimizerStatisticsAdvisorRecommendations implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) implementOptimizerStatisticsAdvisorRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsAdvisorExecutions/{executionName}/actions/implementRecommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImplementOptimizerStatisticsAdvisorRecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ImplementOptimizerStatisticsAdvisorRecommendations"
		err = common.PostProcessServiceError(err, "DbManagement", "ImplementOptimizerStatisticsAdvisorRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &job{})
	return response, err
}

// ListAsmProperties Gets the list of ASM properties for the specified managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAsmProperties.go.html to see an example of how to use ListAsmProperties API.
func (client DbManagementClient) ListAsmProperties(ctx context.Context, request ListAsmPropertiesRequest) (response ListAsmPropertiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAsmProperties, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAsmPropertiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAsmPropertiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAsmPropertiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAsmPropertiesResponse")
	}
	return
}

// listAsmProperties implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listAsmProperties(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/asmProperties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAsmPropertiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListAsmProperties"
		err = common.PostProcessServiceError(err, "DbManagement", "ListAsmProperties", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssociatedDatabases Gets the list of databases using a specific Database Management private endpoint.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/AssociatedDatabaseSummary/ListAssociatedDatabases"
		err = common.PostProcessServiceError(err, "DbManagement", "ListAssociatedDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrDbSnapshots Lists AWR snapshots for the specified database in the AWR.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListAwrDbSnapshots"
		err = common.PostProcessServiceError(err, "DbManagement", "ListAwrDbSnapshots", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrDbs Gets the list of databases and their snapshot summary details available in the AWR of the specified Managed Database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListAwrDbs"
		err = common.PostProcessServiceError(err, "DbManagement", "ListAwrDbs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConsumerGroupPrivileges Gets the list of consumer group privileges granted to a specific user.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListConsumerGroupPrivileges"
		err = common.PostProcessServiceError(err, "DbManagement", "ListConsumerGroupPrivileges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCursorCacheStatements Lists the SQL statements from shared SQL area, also called the cursor cache.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCursorCacheStatements.go.html to see an example of how to use ListCursorCacheStatements API.
// A default retry strategy applies to this operation ListCursorCacheStatements()
func (client DbManagementClient) ListCursorCacheStatements(ctx context.Context, request ListCursorCacheStatementsRequest) (response ListCursorCacheStatementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCursorCacheStatements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCursorCacheStatementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCursorCacheStatementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCursorCacheStatementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCursorCacheStatementsResponse")
	}
	return
}

// listCursorCacheStatements implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listCursorCacheStatements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/cursorCacheStatements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCursorCacheStatementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListCursorCacheStatements"
		err = common.PostProcessServiceError(err, "DbManagement", "ListCursorCacheStatements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataAccessContainers Gets the list of containers for a specific user. This is only applicable if ALL_CONTAINERS !='Y'.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListDataAccessContainers"
		err = common.PostProcessServiceError(err, "DbManagement", "ListDataAccessContainers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseParameters Gets the list of database parameters for the specified Managed Database. The parameters are listed in alphabetical order, along with their current values.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListDatabaseParameters"
		err = common.PostProcessServiceError(err, "DbManagement", "ListDatabaseParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbManagementPrivateEndpoints Gets a list of Database Management private endpoints.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DbManagementPrivateEndpoint/ListDbManagementPrivateEndpoints"
		err = common.PostProcessServiceError(err, "DbManagement", "ListDbManagementPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalAsmDiskGroups Lists ASM disk groups for the external ASM specified by `externalAsmId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsmDiskGroups.go.html to see an example of how to use ListExternalAsmDiskGroups API.
// A default retry strategy applies to this operation ListExternalAsmDiskGroups()
func (client DbManagementClient) ListExternalAsmDiskGroups(ctx context.Context, request ListExternalAsmDiskGroupsRequest) (response ListExternalAsmDiskGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalAsmDiskGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalAsmDiskGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalAsmDiskGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalAsmDiskGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalAsmDiskGroupsResponse")
	}
	return
}

// listExternalAsmDiskGroups implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalAsmDiskGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsms/{externalAsmId}/diskGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalAsmDiskGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/ListExternalAsmDiskGroups"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalAsmDiskGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalAsmInstances Lists the ASM instances in the specified external ASM.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsmInstances.go.html to see an example of how to use ListExternalAsmInstances API.
// A default retry strategy applies to this operation ListExternalAsmInstances()
func (client DbManagementClient) ListExternalAsmInstances(ctx context.Context, request ListExternalAsmInstancesRequest) (response ListExternalAsmInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalAsmInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalAsmInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalAsmInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalAsmInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalAsmInstancesResponse")
	}
	return
}

// listExternalAsmInstances implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalAsmInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsmInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalAsmInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsmInstance/ListExternalAsmInstances"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalAsmInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalAsmUsers Lists ASM users for the external ASM specified by `externalAsmId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsmUsers.go.html to see an example of how to use ListExternalAsmUsers API.
// A default retry strategy applies to this operation ListExternalAsmUsers()
func (client DbManagementClient) ListExternalAsmUsers(ctx context.Context, request ListExternalAsmUsersRequest) (response ListExternalAsmUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalAsmUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalAsmUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalAsmUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalAsmUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalAsmUsersResponse")
	}
	return
}

// listExternalAsmUsers implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalAsmUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsms/{externalAsmId}/users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalAsmUsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/ListExternalAsmUsers"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalAsmUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalAsms Lists the ASMs in the specified external DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsms.go.html to see an example of how to use ListExternalAsms API.
// A default retry strategy applies to this operation ListExternalAsms()
func (client DbManagementClient) ListExternalAsms(ctx context.Context, request ListExternalAsmsRequest) (response ListExternalAsmsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalAsms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalAsmsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalAsmsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalAsmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalAsmsResponse")
	}
	return
}

// listExternalAsms implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalAsms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalAsmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/ListExternalAsms"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalAsms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalClusterInstances Lists the cluster instances in the specified external cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalClusterInstances.go.html to see an example of how to use ListExternalClusterInstances API.
// A default retry strategy applies to this operation ListExternalClusterInstances()
func (client DbManagementClient) ListExternalClusterInstances(ctx context.Context, request ListExternalClusterInstancesRequest) (response ListExternalClusterInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalClusterInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalClusterInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalClusterInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalClusterInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalClusterInstancesResponse")
	}
	return
}

// listExternalClusterInstances implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalClusterInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalClusterInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalClusterInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalClusterInstance/ListExternalClusterInstances"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalClusterInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalClusters Lists the clusters in the specified external DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalClusters.go.html to see an example of how to use ListExternalClusters API.
// A default retry strategy applies to this operation ListExternalClusters()
func (client DbManagementClient) ListExternalClusters(ctx context.Context, request ListExternalClustersRequest) (response ListExternalClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalClustersResponse")
	}
	return
}

// listExternalClusters implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalCluster/ListExternalClusters"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalDatabases Lists the external databases in the specified compartment or in the specified DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDatabases.go.html to see an example of how to use ListExternalDatabases API.
// A default retry strategy applies to this operation ListExternalDatabases()
func (client DbManagementClient) ListExternalDatabases(ctx context.Context, request ListExternalDatabasesRequest) (response ListExternalDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDatabasesResponse")
	}
	return
}

// listExternalDatabases implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDatabaseCollection/ListExternalDatabases"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalDbHomes Lists the DB homes in the specified external DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbHomes.go.html to see an example of how to use ListExternalDbHomes API.
// A default retry strategy applies to this operation ListExternalDbHomes()
func (client DbManagementClient) ListExternalDbHomes(ctx context.Context, request ListExternalDbHomesRequest) (response ListExternalDbHomesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDbHomes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDbHomesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDbHomesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDbHomesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDbHomesResponse")
	}
	return
}

// listExternalDbHomes implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalDbHomes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbHomes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDbHomesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbHome/ListExternalDbHomes"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalDbHomes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalDbNodes Lists the external DB nodes in the specified external DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbNodes.go.html to see an example of how to use ListExternalDbNodes API.
// A default retry strategy applies to this operation ListExternalDbNodes()
func (client DbManagementClient) ListExternalDbNodes(ctx context.Context, request ListExternalDbNodesRequest) (response ListExternalDbNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDbNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDbNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDbNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDbNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDbNodesResponse")
	}
	return
}

// listExternalDbNodes implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalDbNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDbNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbNode/ListExternalDbNodes"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalDbNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalDbSystemConnectors Lists the external connectors in the specified external DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbSystemConnectors.go.html to see an example of how to use ListExternalDbSystemConnectors API.
// A default retry strategy applies to this operation ListExternalDbSystemConnectors()
func (client DbManagementClient) ListExternalDbSystemConnectors(ctx context.Context, request ListExternalDbSystemConnectorsRequest) (response ListExternalDbSystemConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDbSystemConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDbSystemConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDbSystemConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDbSystemConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDbSystemConnectorsResponse")
	}
	return
}

// listExternalDbSystemConnectors implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalDbSystemConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystemConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDbSystemConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemConnector/ListExternalDbSystemConnectors"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalDbSystemConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalDbSystemDiscoveries Lists the external DB system discovery resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbSystemDiscoveries.go.html to see an example of how to use ListExternalDbSystemDiscoveries API.
// A default retry strategy applies to this operation ListExternalDbSystemDiscoveries()
func (client DbManagementClient) ListExternalDbSystemDiscoveries(ctx context.Context, request ListExternalDbSystemDiscoveriesRequest) (response ListExternalDbSystemDiscoveriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDbSystemDiscoveries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDbSystemDiscoveriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDbSystemDiscoveriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDbSystemDiscoveriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDbSystemDiscoveriesResponse")
	}
	return
}

// listExternalDbSystemDiscoveries implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalDbSystemDiscoveries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystemDiscoveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDbSystemDiscoveriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemDiscovery/ListExternalDbSystemDiscoveries"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalDbSystemDiscoveries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalDbSystems Lists the external DB systems in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbSystems.go.html to see an example of how to use ListExternalDbSystems API.
// A default retry strategy applies to this operation ListExternalDbSystems()
func (client DbManagementClient) ListExternalDbSystems(ctx context.Context, request ListExternalDbSystemsRequest) (response ListExternalDbSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDbSystems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDbSystemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDbSystemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDbSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDbSystemsResponse")
	}
	return
}

// listExternalDbSystems implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalDbSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDbSystemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/ListExternalDbSystems"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalDbSystems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalExadataInfrastructures Lists the Exadata infrastructure resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalExadataInfrastructures.go.html to see an example of how to use ListExternalExadataInfrastructures API.
// A default retry strategy applies to this operation ListExternalExadataInfrastructures()
func (client DbManagementClient) ListExternalExadataInfrastructures(ctx context.Context, request ListExternalExadataInfrastructuresRequest) (response ListExternalExadataInfrastructuresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalExadataInfrastructures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalExadataInfrastructuresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalExadataInfrastructuresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalExadataInfrastructuresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalExadataInfrastructuresResponse")
	}
	return
}

// listExternalExadataInfrastructures implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalExadataInfrastructuresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/ListExternalExadataInfrastructures"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalExadataInfrastructures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalExadataStorageConnectors Lists the Exadata storage server connectors for the specified Exadata infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalExadataStorageConnectors.go.html to see an example of how to use ListExternalExadataStorageConnectors API.
// A default retry strategy applies to this operation ListExternalExadataStorageConnectors()
func (client DbManagementClient) ListExternalExadataStorageConnectors(ctx context.Context, request ListExternalExadataStorageConnectorsRequest) (response ListExternalExadataStorageConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalExadataStorageConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalExadataStorageConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalExadataStorageConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalExadataStorageConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalExadataStorageConnectorsResponse")
	}
	return
}

// listExternalExadataStorageConnectors implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalExadataStorageConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalExadataStorageConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageConnector/ListExternalExadataStorageConnectors"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalExadataStorageConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalExadataStorageServers Lists the Exadata storage servers for the specified Exadata infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalExadataStorageServers.go.html to see an example of how to use ListExternalExadataStorageServers API.
// A default retry strategy applies to this operation ListExternalExadataStorageServers()
func (client DbManagementClient) ListExternalExadataStorageServers(ctx context.Context, request ListExternalExadataStorageServersRequest) (response ListExternalExadataStorageServersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalExadataStorageServers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalExadataStorageServersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalExadataStorageServersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalExadataStorageServersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalExadataStorageServersResponse")
	}
	return
}

// listExternalExadataStorageServers implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalExadataStorageServers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalExadataStorageServers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalExadataStorageServersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageServer/ListExternalExadataStorageServers"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalExadataStorageServers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalListenerServices Lists the database services registered with the specified external listener
// for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalListenerServices.go.html to see an example of how to use ListExternalListenerServices API.
// A default retry strategy applies to this operation ListExternalListenerServices()
func (client DbManagementClient) ListExternalListenerServices(ctx context.Context, request ListExternalListenerServicesRequest) (response ListExternalListenerServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalListenerServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalListenerServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalListenerServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalListenerServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalListenerServicesResponse")
	}
	return
}

// listExternalListenerServices implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalListenerServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalListeners/{externalListenerId}/services", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalListenerServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalListener/ListExternalListenerServices"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalListenerServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalListeners Lists the listeners in the specified external DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalListeners.go.html to see an example of how to use ListExternalListeners API.
// A default retry strategy applies to this operation ListExternalListeners()
func (client DbManagementClient) ListExternalListeners(ctx context.Context, request ListExternalListenersRequest) (response ListExternalListenersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalListeners, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalListenersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalListenersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalListenersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalListenersResponse")
	}
	return
}

// listExternalListeners implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listExternalListeners(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalListeners", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalListenersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalListener/ListExternalListeners"
		err = common.PostProcessServiceError(err, "DbManagement", "ListExternalListeners", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/JobExecution/ListJobExecutions"
		err = common.PostProcessServiceError(err, "DbManagement", "ListJobExecutions", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/JobRun/ListJobRuns"
		err = common.PostProcessServiceError(err, "DbManagement", "ListJobRuns", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Job/ListJobs"
		err = common.PostProcessServiceError(err, "DbManagement", "ListJobs", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/ListManagedDatabaseGroups"
		err = common.PostProcessServiceError(err, "DbManagement", "ListManagedDatabaseGroups", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListManagedDatabases"
		err = common.PostProcessServiceError(err, "DbManagement", "ListManagedDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListObjectPrivileges Gets the list of object privileges granted to a specific user.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListObjectPrivileges"
		err = common.PostProcessServiceError(err, "DbManagement", "ListObjectPrivileges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOptimizerStatisticsAdvisorExecutions Lists the details of the Optimizer Statistics Advisor task executions, such as their duration, and the number of findings, if any.
// Optionally, you can specify a date-time range (of seven days) to obtain the list of executions that fall within the specified time range.
// If the date-time range is not specified, then the executions in the last seven days are listed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOptimizerStatisticsAdvisorExecutions.go.html to see an example of how to use ListOptimizerStatisticsAdvisorExecutions API.
func (client DbManagementClient) ListOptimizerStatisticsAdvisorExecutions(ctx context.Context, request ListOptimizerStatisticsAdvisorExecutionsRequest) (response ListOptimizerStatisticsAdvisorExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOptimizerStatisticsAdvisorExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOptimizerStatisticsAdvisorExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOptimizerStatisticsAdvisorExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOptimizerStatisticsAdvisorExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOptimizerStatisticsAdvisorExecutionsResponse")
	}
	return
}

// listOptimizerStatisticsAdvisorExecutions implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listOptimizerStatisticsAdvisorExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsAdvisorExecutions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOptimizerStatisticsAdvisorExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListOptimizerStatisticsAdvisorExecutions"
		err = common.PostProcessServiceError(err, "DbManagement", "ListOptimizerStatisticsAdvisorExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOptimizerStatisticsCollectionAggregations Gets a list of the optimizer statistics collection operations per hour, grouped by task or object status for the specified Managed Database.
// You must specify a value for GroupByQueryParam to determine whether the data should be grouped by task status or task object status.
// Optionally, you can specify a date-time range (of seven days) to obtain collection aggregations within the specified time range.
// If the date-time range is not specified, then the operations in the last seven days are listed.
// You can further filter the results by providing the optional type of TaskTypeQueryParam.
// If the task type not provided, then both Auto and Manual tasks are considered for aggregation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOptimizerStatisticsCollectionAggregations.go.html to see an example of how to use ListOptimizerStatisticsCollectionAggregations API.
func (client DbManagementClient) ListOptimizerStatisticsCollectionAggregations(ctx context.Context, request ListOptimizerStatisticsCollectionAggregationsRequest) (response ListOptimizerStatisticsCollectionAggregationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOptimizerStatisticsCollectionAggregations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOptimizerStatisticsCollectionAggregationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOptimizerStatisticsCollectionAggregationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOptimizerStatisticsCollectionAggregationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOptimizerStatisticsCollectionAggregationsResponse")
	}
	return
}

// listOptimizerStatisticsCollectionAggregations implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listOptimizerStatisticsCollectionAggregations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsCollectionAggregations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOptimizerStatisticsCollectionAggregationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListOptimizerStatisticsCollectionAggregations"
		err = common.PostProcessServiceError(err, "DbManagement", "ListOptimizerStatisticsCollectionAggregations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOptimizerStatisticsCollectionOperations Lists the Optimizer Statistics Collection (Auto and Manual) task operation summary for the specified Managed Database.
// The summary includes the details of each operation and the number of tasks grouped by status: Completed, In Progress, Failed, and so on.
// Optionally, you can specify a date-time range (of seven days) to obtain the list of operations that fall within the specified time range.
// If the date-time range is not specified, then the operations in the last seven days are listed.
// This API also enables the pagination of results and the opc-next-page response header indicates whether there is a next page.
// If you use the same header value in a consecutive request, the next page records are returned.
// To obtain the required results, you can apply the different types of filters supported by this API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOptimizerStatisticsCollectionOperations.go.html to see an example of how to use ListOptimizerStatisticsCollectionOperations API.
func (client DbManagementClient) ListOptimizerStatisticsCollectionOperations(ctx context.Context, request ListOptimizerStatisticsCollectionOperationsRequest) (response ListOptimizerStatisticsCollectionOperationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOptimizerStatisticsCollectionOperations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOptimizerStatisticsCollectionOperationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOptimizerStatisticsCollectionOperationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOptimizerStatisticsCollectionOperationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOptimizerStatisticsCollectionOperationsResponse")
	}
	return
}

// listOptimizerStatisticsCollectionOperations implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listOptimizerStatisticsCollectionOperations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/optimizerStatisticsCollectionOperations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOptimizerStatisticsCollectionOperationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListOptimizerStatisticsCollectionOperations"
		err = common.PostProcessServiceError(err, "DbManagement", "ListOptimizerStatisticsCollectionOperations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPreferredCredentials Gets the list of preferred credentials for a given Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListPreferredCredentials.go.html to see an example of how to use ListPreferredCredentials API.
func (client DbManagementClient) ListPreferredCredentials(ctx context.Context, request ListPreferredCredentialsRequest) (response ListPreferredCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPreferredCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPreferredCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPreferredCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPreferredCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPreferredCredentialsResponse")
	}
	return
}

// listPreferredCredentials implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listPreferredCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/preferredCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPreferredCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/PreferredCredential/ListPreferredCredentials"
		err = common.PostProcessServiceError(err, "DbManagement", "ListPreferredCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProxiedForUsers Gets the list of users on whose behalf the current user acts as proxy.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListProxiedForUsers"
		err = common.PostProcessServiceError(err, "DbManagement", "ListProxiedForUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProxyUsers Gets the list of proxy users for the current user.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListProxyUsers"
		err = common.PostProcessServiceError(err, "DbManagement", "ListProxyUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoles Gets the list of roles granted to a specific user.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListRoles"
		err = common.PostProcessServiceError(err, "DbManagement", "ListRoles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlPlanBaselineJobs Lists the database jobs used for loading SQL plan baselines in the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlPlanBaselineJobs.go.html to see an example of how to use ListSqlPlanBaselineJobs API.
// A default retry strategy applies to this operation ListSqlPlanBaselineJobs()
func (client DbManagementClient) ListSqlPlanBaselineJobs(ctx context.Context, request ListSqlPlanBaselineJobsRequest) (response ListSqlPlanBaselineJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlPlanBaselineJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlPlanBaselineJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlPlanBaselineJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlPlanBaselineJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlPlanBaselineJobsResponse")
	}
	return
}

// listSqlPlanBaselineJobs implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listSqlPlanBaselineJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselineJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlPlanBaselineJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSqlPlanBaselineJobs"
		err = common.PostProcessServiceError(err, "DbManagement", "ListSqlPlanBaselineJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlPlanBaselines Lists the SQL plan baselines for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlPlanBaselines.go.html to see an example of how to use ListSqlPlanBaselines API.
// A default retry strategy applies to this operation ListSqlPlanBaselines()
func (client DbManagementClient) ListSqlPlanBaselines(ctx context.Context, request ListSqlPlanBaselinesRequest) (response ListSqlPlanBaselinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlPlanBaselines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlPlanBaselinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlPlanBaselinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlPlanBaselinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlPlanBaselinesResponse")
	}
	return
}

// listSqlPlanBaselines implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listSqlPlanBaselines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlPlanBaselinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSqlPlanBaselines"
		err = common.PostProcessServiceError(err, "DbManagement", "ListSqlPlanBaselines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSystemPrivileges Gets the list of system privileges granted to a specific user.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSystemPrivileges"
		err = common.PostProcessServiceError(err, "DbManagement", "ListSystemPrivileges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTableStatistics Lists the database table statistics grouped by different statuses such as Not Stale Stats, Stale Stats, and No Stats.
// This also includes the percentage of each status.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListTableStatistics.go.html to see an example of how to use ListTableStatistics API.
func (client DbManagementClient) ListTableStatistics(ctx context.Context, request ListTableStatisticsRequest) (response ListTableStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTableStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTableStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTableStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTableStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTableStatisticsResponse")
	}
	return
}

// listTableStatistics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) listTableStatistics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/tableStatistics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTableStatisticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListTableStatistics"
		err = common.PostProcessServiceError(err, "DbManagement", "ListTableStatistics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTablespaces Gets the list of tablespaces for the specified managedDatabaseId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/ListTablespaces"
		err = common.PostProcessServiceError(err, "DbManagement", "ListTablespaces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUsers Gets the list of users for the specified managedDatabaseId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListUsers"
		err = common.PostProcessServiceError(err, "DbManagement", "ListUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a paginated list of errors for a given work request.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DbManagement", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for a given work request.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DbManagement", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests The list of work requests in a specific compartment was retrieved successfully.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DbManagement", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LoadSqlPlanBaselinesFromAwr Loads plans from Automatic Workload Repository (AWR) snapshots. You must
// specify the beginning and ending of the snapshot range. Optionally, you
// can apply a filter to load only plans that meet specified criteria. By
// default, the optimizer uses the loaded plans the next time that the database
// executes the SQL statements.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/LoadSqlPlanBaselinesFromAwr.go.html to see an example of how to use LoadSqlPlanBaselinesFromAwr API.
func (client DbManagementClient) LoadSqlPlanBaselinesFromAwr(ctx context.Context, request LoadSqlPlanBaselinesFromAwrRequest) (response LoadSqlPlanBaselinesFromAwrResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.loadSqlPlanBaselinesFromAwr, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LoadSqlPlanBaselinesFromAwrResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LoadSqlPlanBaselinesFromAwrResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LoadSqlPlanBaselinesFromAwrResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LoadSqlPlanBaselinesFromAwrResponse")
	}
	return
}

// loadSqlPlanBaselinesFromAwr implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) loadSqlPlanBaselinesFromAwr(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/loadSqlPlanBaselinesFromAwr", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LoadSqlPlanBaselinesFromAwrResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/LoadSqlPlanBaselinesFromAwr"
		err = common.PostProcessServiceError(err, "DbManagement", "LoadSqlPlanBaselinesFromAwr", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LoadSqlPlanBaselinesFromCursorCache Loads plans for statements directly from the shared SQL area, also called
// the cursor cache. By applying a filter on the module name, the schema, or
// the SQL ID you identify the SQL statement or set of SQL statements to load.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/LoadSqlPlanBaselinesFromCursorCache.go.html to see an example of how to use LoadSqlPlanBaselinesFromCursorCache API.
func (client DbManagementClient) LoadSqlPlanBaselinesFromCursorCache(ctx context.Context, request LoadSqlPlanBaselinesFromCursorCacheRequest) (response LoadSqlPlanBaselinesFromCursorCacheResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.loadSqlPlanBaselinesFromCursorCache, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LoadSqlPlanBaselinesFromCursorCacheResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LoadSqlPlanBaselinesFromCursorCacheResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LoadSqlPlanBaselinesFromCursorCacheResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LoadSqlPlanBaselinesFromCursorCacheResponse")
	}
	return
}

// loadSqlPlanBaselinesFromCursorCache implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) loadSqlPlanBaselinesFromCursorCache(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselines/actions/loadSqlPlanBaselinesFromCursorCache", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LoadSqlPlanBaselinesFromCursorCacheResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/LoadSqlPlanBaselinesFromCursorCache"
		err = common.PostProcessServiceError(err, "DbManagement", "LoadSqlPlanBaselinesFromCursorCache", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchExternalDbSystemDiscovery Patches the external DB system discovery specified by `externalDbSystemDiscoveryId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/PatchExternalDbSystemDiscovery.go.html to see an example of how to use PatchExternalDbSystemDiscovery API.
func (client DbManagementClient) PatchExternalDbSystemDiscovery(ctx context.Context, request PatchExternalDbSystemDiscoveryRequest) (response PatchExternalDbSystemDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchExternalDbSystemDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchExternalDbSystemDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchExternalDbSystemDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchExternalDbSystemDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchExternalDbSystemDiscoveryResponse")
	}
	return
}

// patchExternalDbSystemDiscovery implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) patchExternalDbSystemDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/externalDbSystemDiscoveries/{externalDbSystemDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchExternalDbSystemDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemDiscovery/PatchExternalDbSystemDiscovery"
		err = common.PostProcessServiceError(err, "DbManagement", "PatchExternalDbSystemDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDataFile Removes a data file or temp file from the tablespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/RemoveDataFile.go.html to see an example of how to use RemoveDataFile API.
func (client DbManagementClient) RemoveDataFile(ctx context.Context, request RemoveDataFileRequest) (response RemoveDataFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeDataFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDataFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDataFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDataFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDataFileResponse")
	}
	return
}

// removeDataFile implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) removeDataFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/tablespaces/{tablespaceName}/actions/removeDataFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDataFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/RemoveDataFile"
		err = common.PostProcessServiceError(err, "DbManagement", "RemoveDataFile", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/RemoveManagedDatabaseFromManagedDatabaseGroup"
		err = common.PostProcessServiceError(err, "DbManagement", "RemoveManagedDatabaseFromManagedDatabaseGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResetDatabaseParameters Resets database parameter values to their default or startup values.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ResetDatabaseParameters"
		err = common.PostProcessServiceError(err, "DbManagement", "ResetDatabaseParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResizeDataFile Resizes a data file or temp file within the tablespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ResizeDataFile.go.html to see an example of how to use ResizeDataFile API.
func (client DbManagementClient) ResizeDataFile(ctx context.Context, request ResizeDataFileRequest) (response ResizeDataFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resizeDataFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResizeDataFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResizeDataFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResizeDataFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResizeDataFileResponse")
	}
	return
}

// resizeDataFile implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) resizeDataFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/tablespaces/{tablespaceName}/actions/resizeDataFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResizeDataFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/ResizeDataFile"
		err = common.PostProcessServiceError(err, "DbManagement", "ResizeDataFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunHistoricAddm Creates and executes a historic ADDM task using the specified AWR snapshot IDs. If an existing ADDM task
// uses the provided awr snapshot IDs, the existing task will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/RunHistoricAddm.go.html to see an example of how to use RunHistoricAddm API.
func (client DbManagementClient) RunHistoricAddm(ctx context.Context, request RunHistoricAddmRequest) (response RunHistoricAddmResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.runHistoricAddm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunHistoricAddmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunHistoricAddmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunHistoricAddmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunHistoricAddmResponse")
	}
	return
}

// runHistoricAddm implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) runHistoricAddm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/runHistoricAddm", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunHistoricAddmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/HistoricAddmResult/RunHistoricAddm"
		err = common.PostProcessServiceError(err, "DbManagement", "RunHistoricAddm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbCpuUsages Summarizes the AWR CPU resource limits and metrics for the specified database in AWR.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbCpuUsages"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbCpuUsages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbMetrics Summarizes the metric samples for the specified database in the AWR. The metric samples are summarized based on the Time dimension for each metric.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbMetrics", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbParameterChanges"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbParameterChanges", apiReferenceLink)
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
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbParameters"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbSnapshotRanges Summarizes the AWR snapshot ranges that contain continuous snapshots, for the specified Managed Database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbSnapshotRanges"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbSnapshotRanges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbSysstats Summarizes the AWR SYSSTAT sample data for the specified database in AWR. The statistical data is summarized based on the Time dimension for each statistic.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbSysstats"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbSysstats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbTopWaitEvents Summarizes the AWR top wait events.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbTopWaitEvents"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbTopWaitEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbWaitEventBuckets Summarizes AWR wait event data into value buckets and frequency, for the specified database in the AWR.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbWaitEventBuckets"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbWaitEventBuckets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDbWaitEvents Summarizes the AWR wait event sample data for the specified database in the AWR. The event data is summarized based on the Time dimension for each event.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAwrDbWaitEvents"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeAwrDbWaitEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExternalAsmMetrics Gets metrics for the external ASM specified by `externalAsmId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeExternalAsmMetrics.go.html to see an example of how to use SummarizeExternalAsmMetrics API.
// A default retry strategy applies to this operation SummarizeExternalAsmMetrics()
func (client DbManagementClient) SummarizeExternalAsmMetrics(ctx context.Context, request SummarizeExternalAsmMetricsRequest) (response SummarizeExternalAsmMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExternalAsmMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExternalAsmMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExternalAsmMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExternalAsmMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExternalAsmMetricsResponse")
	}
	return
}

// summarizeExternalAsmMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeExternalAsmMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalAsms/{externalAsmId}/metrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExternalAsmMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/SummarizeExternalAsmMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeExternalAsmMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExternalClusterMetrics Gets metrics for the external cluster specified by `externalClusterId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeExternalClusterMetrics.go.html to see an example of how to use SummarizeExternalClusterMetrics API.
// A default retry strategy applies to this operation SummarizeExternalClusterMetrics()
func (client DbManagementClient) SummarizeExternalClusterMetrics(ctx context.Context, request SummarizeExternalClusterMetricsRequest) (response SummarizeExternalClusterMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExternalClusterMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExternalClusterMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExternalClusterMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExternalClusterMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExternalClusterMetricsResponse")
	}
	return
}

// summarizeExternalClusterMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeExternalClusterMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalClusters/{externalClusterId}/metrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExternalClusterMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalCluster/SummarizeExternalClusterMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeExternalClusterMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExternalDbNodeMetrics Gets metrics for the external DB node specified by `externalDbNodeId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeExternalDbNodeMetrics.go.html to see an example of how to use SummarizeExternalDbNodeMetrics API.
// A default retry strategy applies to this operation SummarizeExternalDbNodeMetrics()
func (client DbManagementClient) SummarizeExternalDbNodeMetrics(ctx context.Context, request SummarizeExternalDbNodeMetricsRequest) (response SummarizeExternalDbNodeMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExternalDbNodeMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExternalDbNodeMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExternalDbNodeMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExternalDbNodeMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExternalDbNodeMetricsResponse")
	}
	return
}

// summarizeExternalDbNodeMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeExternalDbNodeMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbNodes/{externalDbNodeId}/metrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExternalDbNodeMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbNode/SummarizeExternalDbNodeMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeExternalDbNodeMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExternalDbSystemAvailabilityMetrics Gets availability metrics for the components present in the external DB system specified by `externalDbSystemId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeExternalDbSystemAvailabilityMetrics.go.html to see an example of how to use SummarizeExternalDbSystemAvailabilityMetrics API.
// A default retry strategy applies to this operation SummarizeExternalDbSystemAvailabilityMetrics()
func (client DbManagementClient) SummarizeExternalDbSystemAvailabilityMetrics(ctx context.Context, request SummarizeExternalDbSystemAvailabilityMetricsRequest) (response SummarizeExternalDbSystemAvailabilityMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExternalDbSystemAvailabilityMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExternalDbSystemAvailabilityMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExternalDbSystemAvailabilityMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExternalDbSystemAvailabilityMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExternalDbSystemAvailabilityMetricsResponse")
	}
	return
}

// summarizeExternalDbSystemAvailabilityMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeExternalDbSystemAvailabilityMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalDbSystems/{externalDbSystemId}/availabilityMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExternalDbSystemAvailabilityMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/SummarizeExternalDbSystemAvailabilityMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeExternalDbSystemAvailabilityMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExternalListenerMetrics Gets metrics for the external listener specified by `externalListenerId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeExternalListenerMetrics.go.html to see an example of how to use SummarizeExternalListenerMetrics API.
// A default retry strategy applies to this operation SummarizeExternalListenerMetrics()
func (client DbManagementClient) SummarizeExternalListenerMetrics(ctx context.Context, request SummarizeExternalListenerMetricsRequest) (response SummarizeExternalListenerMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExternalListenerMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExternalListenerMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExternalListenerMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExternalListenerMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExternalListenerMetricsResponse")
	}
	return
}

// summarizeExternalListenerMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeExternalListenerMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalListeners/{externalListenerId}/metrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExternalListenerMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalListener/SummarizeExternalListenerMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeExternalListenerMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeJobExecutionsStatuses Gets the number of job executions grouped by status for a job, Managed Database, or Database Group in a specific compartment. Only one of the parameters, jobId, managedDatabaseId, or managedDatabaseGroupId should be provided.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/JobExecutionsStatusSummaryCollection/SummarizeJobExecutionsStatuses"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeJobExecutionsStatuses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeManagedDatabaseAvailabilityMetrics Gets the availability metrics related to managed database for the Oracle
// database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeManagedDatabaseAvailabilityMetrics.go.html to see an example of how to use SummarizeManagedDatabaseAvailabilityMetrics API.
// A default retry strategy applies to this operation SummarizeManagedDatabaseAvailabilityMetrics()
func (client DbManagementClient) SummarizeManagedDatabaseAvailabilityMetrics(ctx context.Context, request SummarizeManagedDatabaseAvailabilityMetricsRequest) (response SummarizeManagedDatabaseAvailabilityMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeManagedDatabaseAvailabilityMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeManagedDatabaseAvailabilityMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeManagedDatabaseAvailabilityMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeManagedDatabaseAvailabilityMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeManagedDatabaseAvailabilityMetricsResponse")
	}
	return
}

// summarizeManagedDatabaseAvailabilityMetrics implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeManagedDatabaseAvailabilityMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/availabilityMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeManagedDatabaseAvailabilityMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeManagedDatabaseAvailabilityMetrics"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeManagedDatabaseAvailabilityMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlPlanBaselines Gets the number of SQL plan baselines aggregated by their attributes.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeSqlPlanBaselines.go.html to see an example of how to use SummarizeSqlPlanBaselines API.
// A default retry strategy applies to this operation SummarizeSqlPlanBaselines()
func (client DbManagementClient) SummarizeSqlPlanBaselines(ctx context.Context, request SummarizeSqlPlanBaselinesRequest) (response SummarizeSqlPlanBaselinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlPlanBaselines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlPlanBaselinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlPlanBaselinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlPlanBaselinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlPlanBaselinesResponse")
	}
	return
}

// summarizeSqlPlanBaselines implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeSqlPlanBaselines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselineStats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlPlanBaselinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeSqlPlanBaselines"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeSqlPlanBaselines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlPlanBaselinesByLastExecution Gets the number of SQL plan baselines aggregated by the age of their last execution in weeks.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeSqlPlanBaselinesByLastExecution.go.html to see an example of how to use SummarizeSqlPlanBaselinesByLastExecution API.
// A default retry strategy applies to this operation SummarizeSqlPlanBaselinesByLastExecution()
func (client DbManagementClient) SummarizeSqlPlanBaselinesByLastExecution(ctx context.Context, request SummarizeSqlPlanBaselinesByLastExecutionRequest) (response SummarizeSqlPlanBaselinesByLastExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlPlanBaselinesByLastExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlPlanBaselinesByLastExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlPlanBaselinesByLastExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlPlanBaselinesByLastExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlPlanBaselinesByLastExecutionResponse")
	}
	return
}

// summarizeSqlPlanBaselinesByLastExecution implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) summarizeSqlPlanBaselinesByLastExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlPlanBaselineExecutionStats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlPlanBaselinesByLastExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeSqlPlanBaselinesByLastExecution"
		err = common.PostProcessServiceError(err, "DbManagement", "SummarizeSqlPlanBaselinesByLastExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestPreferredCredential Tests the preferred credential.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/TestPreferredCredential.go.html to see an example of how to use TestPreferredCredential API.
func (client DbManagementClient) TestPreferredCredential(ctx context.Context, request TestPreferredCredentialRequest) (response TestPreferredCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.testPreferredCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestPreferredCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestPreferredCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestPreferredCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestPreferredCredentialResponse")
	}
	return
}

// testPreferredCredential implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) testPreferredCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/preferredCredentials/{credentialName}/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestPreferredCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/PreferredCredential/TestPreferredCredential"
		err = common.PostProcessServiceError(err, "DbManagement", "TestPreferredCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbManagementPrivateEndpoint Updates one or more attributes of a specific Database Management private endpoint.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/DbManagementPrivateEndpoint/UpdateDbManagementPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateDbManagementPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalAsm Updates the external ASM specified by `externalAsmId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalAsm.go.html to see an example of how to use UpdateExternalAsm API.
func (client DbManagementClient) UpdateExternalAsm(ctx context.Context, request UpdateExternalAsmRequest) (response UpdateExternalAsmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalAsm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalAsmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalAsmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalAsmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalAsmResponse")
	}
	return
}

// updateExternalAsm implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalAsm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalAsms/{externalAsmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalAsmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalAsm/UpdateExternalAsm"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalAsm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalCluster Updates the external cluster specified by `externalClusterId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalCluster.go.html to see an example of how to use UpdateExternalCluster API.
func (client DbManagementClient) UpdateExternalCluster(ctx context.Context, request UpdateExternalClusterRequest) (response UpdateExternalClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalClusterResponse")
	}
	return
}

// updateExternalCluster implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalClusters/{externalClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalCluster/UpdateExternalCluster"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalClusterInstance Updates the external cluster instance specified by `externalClusterInstanceId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalClusterInstance.go.html to see an example of how to use UpdateExternalClusterInstance API.
func (client DbManagementClient) UpdateExternalClusterInstance(ctx context.Context, request UpdateExternalClusterInstanceRequest) (response UpdateExternalClusterInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalClusterInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalClusterInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalClusterInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalClusterInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalClusterInstanceResponse")
	}
	return
}

// updateExternalClusterInstance implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalClusterInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalClusterInstances/{externalClusterInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalClusterInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalClusterInstance/UpdateExternalClusterInstance"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalClusterInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalDbNode Updates the external DB node specified by `externalDbNodeId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalDbNode.go.html to see an example of how to use UpdateExternalDbNode API.
func (client DbManagementClient) UpdateExternalDbNode(ctx context.Context, request UpdateExternalDbNodeRequest) (response UpdateExternalDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalDbNodeResponse")
	}
	return
}

// updateExternalDbNode implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalDbNodes/{externalDbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbNode/UpdateExternalDbNode"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalDbSystem Updates the external DB system specified by `externalDbSystemId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalDbSystem.go.html to see an example of how to use UpdateExternalDbSystem API.
func (client DbManagementClient) UpdateExternalDbSystem(ctx context.Context, request UpdateExternalDbSystemRequest) (response UpdateExternalDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalDbSystemResponse")
	}
	return
}

// updateExternalDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalDbSystems/{externalDbSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystem/UpdateExternalDbSystem"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalDbSystemConnector Updates the external connector specified by `externalDbSystemConnectorId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalDbSystemConnector.go.html to see an example of how to use UpdateExternalDbSystemConnector API.
func (client DbManagementClient) UpdateExternalDbSystemConnector(ctx context.Context, request UpdateExternalDbSystemConnectorRequest) (response UpdateExternalDbSystemConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalDbSystemConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalDbSystemConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalDbSystemConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalDbSystemConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalDbSystemConnectorResponse")
	}
	return
}

// updateExternalDbSystemConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalDbSystemConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalDbSystemConnectors/{externalDbSystemConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalDbSystemConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemConnector/UpdateExternalDbSystemConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalDbSystemConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalDbSystemDiscovery Updates the external DB system discovery specified by `externalDbSystemDiscoveryId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalDbSystemDiscovery.go.html to see an example of how to use UpdateExternalDbSystemDiscovery API.
func (client DbManagementClient) UpdateExternalDbSystemDiscovery(ctx context.Context, request UpdateExternalDbSystemDiscoveryRequest) (response UpdateExternalDbSystemDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalDbSystemDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalDbSystemDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalDbSystemDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalDbSystemDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalDbSystemDiscoveryResponse")
	}
	return
}

// updateExternalDbSystemDiscovery implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalDbSystemDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalDbSystemDiscoveries/{externalDbSystemDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalDbSystemDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalDbSystemDiscovery/UpdateExternalDbSystemDiscovery"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalDbSystemDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalExadataInfrastructure Updates the details for the Exadata infrastructure specified by externalExadataInfrastructureId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalExadataInfrastructure.go.html to see an example of how to use UpdateExternalExadataInfrastructure API.
// A default retry strategy applies to this operation UpdateExternalExadataInfrastructure()
func (client DbManagementClient) UpdateExternalExadataInfrastructure(ctx context.Context, request UpdateExternalExadataInfrastructureRequest) (response UpdateExternalExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateExternalExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalExadataInfrastructureResponse")
	}
	return
}

// updateExternalExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalExadataInfrastructures/{externalExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataInfrastructure/UpdateExternalExadataInfrastructure"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalExadataStorageConnector Updates the Exadata storage server connector specified by exadataStorageConnectorId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalExadataStorageConnector.go.html to see an example of how to use UpdateExternalExadataStorageConnector API.
func (client DbManagementClient) UpdateExternalExadataStorageConnector(ctx context.Context, request UpdateExternalExadataStorageConnectorRequest) (response UpdateExternalExadataStorageConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalExadataStorageConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalExadataStorageConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalExadataStorageConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalExadataStorageConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalExadataStorageConnectorResponse")
	}
	return
}

// updateExternalExadataStorageConnector implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalExadataStorageConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalExadataStorageConnectors/{externalExadataStorageConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalExadataStorageConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalExadataStorageConnector/UpdateExternalExadataStorageConnector"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalExadataStorageConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalListener Updates the external listener specified by `externalListenerId`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateExternalListener.go.html to see an example of how to use UpdateExternalListener API.
func (client DbManagementClient) UpdateExternalListener(ctx context.Context, request UpdateExternalListenerRequest) (response UpdateExternalListenerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalListener, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalListenerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalListenerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalListenerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalListenerResponse")
	}
	return
}

// updateExternalListener implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateExternalListener(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalListeners/{externalListenerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalListener/UpdateExternalListener"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateExternalListener", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJob Updates the details for the recurring scheduled job specified by jobId. Note that non-recurring (one time) jobs cannot be updated.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Job/UpdateJob"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &job{})
	return response, err
}

// UpdateManagedDatabaseGroup Updates the Managed Database Group specified by managedDatabaseGroupId.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabaseGroup/UpdateManagedDatabaseGroup"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateManagedDatabaseGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePreferredCredential Updates the preferred credential based on the credentialName.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdatePreferredCredential.go.html to see an example of how to use UpdatePreferredCredential API.
func (client DbManagementClient) UpdatePreferredCredential(ctx context.Context, request UpdatePreferredCredentialRequest) (response UpdatePreferredCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePreferredCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePreferredCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePreferredCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePreferredCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePreferredCredentialResponse")
	}
	return
}

// updatePreferredCredential implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updatePreferredCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedDatabases/{managedDatabaseId}/preferredCredentials/{credentialName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePreferredCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/PreferredCredential/UpdatePreferredCredential"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdatePreferredCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &preferredcredential{})
	return response, err
}

// UpdateTablespace Updates the attributes of the tablespace specified by tablespaceName within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/UpdateTablespace.go.html to see an example of how to use UpdateTablespace API.
func (client DbManagementClient) UpdateTablespace(ctx context.Context, request UpdateTablespaceRequest) (response UpdateTablespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTablespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTablespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTablespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTablespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTablespaceResponse")
	}
	return
}

// updateTablespace implements the OCIOperation interface (enables retrying operations)
func (client DbManagementClient) updateTablespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedDatabases/{managedDatabaseId}/tablespaces/{tablespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTablespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/Tablespace/UpdateTablespace"
		err = common.PostProcessServiceError(err, "DbManagement", "UpdateTablespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
