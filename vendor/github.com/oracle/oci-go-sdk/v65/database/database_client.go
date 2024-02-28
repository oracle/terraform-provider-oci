// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DatabaseClient a client for Database
type DatabaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseClientWithConfigurationProvider Creates a new default Database client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
	if enabled := common.CheckForEnabledServices("database"); !enabled {
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
	return newDatabaseClientFromBaseClient(baseClient, provider)
}

// NewDatabaseClientWithOboToken Creates a new default Database client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatabaseClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
	// Database service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Database"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatabaseClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("database", "https://database.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateExadataInfrastructure Activates the specified Exadata infrastructure resource. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ActivateExadataInfrastructure.go.html to see an example of how to use ActivateExadataInfrastructure API.
func (client DatabaseClient) ActivateExadataInfrastructure(ctx context.Context, request ActivateExadataInfrastructureRequest) (response ActivateExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateExadataInfrastructureResponse")
	}
	return
}

// activateExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) activateExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/ActivateExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "ActivateExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddStorageCapacityCloudExadataInfrastructure Makes the storage capacity from additional storage servers available for Cloud VM Cluster consumption. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/AddStorageCapacityCloudExadataInfrastructure.go.html to see an example of how to use AddStorageCapacityCloudExadataInfrastructure API.
func (client DatabaseClient) AddStorageCapacityCloudExadataInfrastructure(ctx context.Context, request AddStorageCapacityCloudExadataInfrastructureRequest) (response AddStorageCapacityCloudExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addStorageCapacityCloudExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddStorageCapacityCloudExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddStorageCapacityCloudExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddStorageCapacityCloudExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddStorageCapacityCloudExadataInfrastructureResponse")
	}
	return
}

// addStorageCapacityCloudExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) addStorageCapacityCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}/actions/addStorageCapacity", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddStorageCapacityCloudExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/AddStorageCapacityCloudExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "AddStorageCapacityCloudExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddStorageCapacityExadataInfrastructure Makes the storage capacity from additional storage servers available for VM Cluster consumption. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/AddStorageCapacityExadataInfrastructure.go.html to see an example of how to use AddStorageCapacityExadataInfrastructure API.
func (client DatabaseClient) AddStorageCapacityExadataInfrastructure(ctx context.Context, request AddStorageCapacityExadataInfrastructureRequest) (response AddStorageCapacityExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addStorageCapacityExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddStorageCapacityExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddStorageCapacityExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddStorageCapacityExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddStorageCapacityExadataInfrastructureResponse")
	}
	return
}

// addStorageCapacityExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) addStorageCapacityExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/addStorageCapacity", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddStorageCapacityExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/AddStorageCapacityExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "AddStorageCapacityExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddVirtualMachineToCloudVmCluster Add Virtual Machines to the Cloud VM cluster. Applies to Exadata Cloud instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/AddVirtualMachineToCloudVmCluster.go.html to see an example of how to use AddVirtualMachineToCloudVmCluster API.
func (client DatabaseClient) AddVirtualMachineToCloudVmCluster(ctx context.Context, request AddVirtualMachineToCloudVmClusterRequest) (response AddVirtualMachineToCloudVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addVirtualMachineToCloudVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddVirtualMachineToCloudVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddVirtualMachineToCloudVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddVirtualMachineToCloudVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddVirtualMachineToCloudVmClusterResponse")
	}
	return
}

// addVirtualMachineToCloudVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) addVirtualMachineToCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudVmClusters/{cloudVmClusterId}/actions/addVirtualMachine", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddVirtualMachineToCloudVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/AddVirtualMachineToCloudVmCluster"
		err = common.PostProcessServiceError(err, "Database", "AddVirtualMachineToCloudVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddVirtualMachineToVmCluster Add Virtual Machines to the VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/AddVirtualMachineToVmCluster.go.html to see an example of how to use AddVirtualMachineToVmCluster API.
func (client DatabaseClient) AddVirtualMachineToVmCluster(ctx context.Context, request AddVirtualMachineToVmClusterRequest) (response AddVirtualMachineToVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addVirtualMachineToVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddVirtualMachineToVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddVirtualMachineToVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddVirtualMachineToVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddVirtualMachineToVmClusterResponse")
	}
	return
}

// addVirtualMachineToVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) addVirtualMachineToVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusters/{vmClusterId}/actions/addVirtualMachine", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddVirtualMachineToVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/AddVirtualMachineToVmCluster"
		err = common.PostProcessServiceError(err, "Database", "AddVirtualMachineToVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AutonomousDatabaseManualRefresh Initiates a data refresh for an Autonomous Database refreshable clone. Data is refreshed from the source database to the point of a specified timestamp.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/AutonomousDatabaseManualRefresh.go.html to see an example of how to use AutonomousDatabaseManualRefresh API.
func (client DatabaseClient) AutonomousDatabaseManualRefresh(ctx context.Context, request AutonomousDatabaseManualRefreshRequest) (response AutonomousDatabaseManualRefreshResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.autonomousDatabaseManualRefresh, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AutonomousDatabaseManualRefreshResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AutonomousDatabaseManualRefreshResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AutonomousDatabaseManualRefreshResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AutonomousDatabaseManualRefreshResponse")
	}
	return
}

// autonomousDatabaseManualRefresh implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) autonomousDatabaseManualRefresh(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AutonomousDatabaseManualRefreshResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/AutonomousDatabaseManualRefresh"
		err = common.PostProcessServiceError(err, "Database", "AutonomousDatabaseManualRefresh", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelBackup Cancel automatic/standalone full/incremental create backup workrequests specified by the backup Id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CancelBackup.go.html to see an example of how to use CancelBackup API.
func (client DatabaseClient) CancelBackup(ctx context.Context, request CancelBackupRequest) (response CancelBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelBackupResponse")
	}
	return
}

// cancelBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) cancelBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups/{backupId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Backup/CancelBackup"
		err = common.PostProcessServiceError(err, "Database", "CancelBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAutonomousContainerDatabaseCompartment Move the Autonomous Container Database and its dependent resources to the specified compartment.
// For more information about moving Autonomous Container Databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeAutonomousContainerDatabaseCompartment.go.html to see an example of how to use ChangeAutonomousContainerDatabaseCompartment API.
func (client DatabaseClient) ChangeAutonomousContainerDatabaseCompartment(ctx context.Context, request ChangeAutonomousContainerDatabaseCompartmentRequest) (response ChangeAutonomousContainerDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAutonomousContainerDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAutonomousContainerDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAutonomousContainerDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAutonomousContainerDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAutonomousContainerDatabaseCompartmentResponse")
	}
	return
}

// changeAutonomousContainerDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeAutonomousContainerDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousContainerDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/ChangeAutonomousContainerDatabaseCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeAutonomousContainerDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAutonomousDatabaseCompartment Move the Autonomous Database and its dependent resources to the specified compartment.
// For more information about moving Autonomous Databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeAutonomousDatabaseCompartment.go.html to see an example of how to use ChangeAutonomousDatabaseCompartment API.
func (client DatabaseClient) ChangeAutonomousDatabaseCompartment(ctx context.Context, request ChangeAutonomousDatabaseCompartmentRequest) (response ChangeAutonomousDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAutonomousDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAutonomousDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAutonomousDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAutonomousDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAutonomousDatabaseCompartmentResponse")
	}
	return
}

// changeAutonomousDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeAutonomousDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ChangeAutonomousDatabaseCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeAutonomousDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAutonomousExadataInfrastructureCompartment **Deprecated.** Use the ChangeCloudExadataInfrastructureCompartment operation to move an Exadata infrastructure resource to a different compartment and  ChangeCloudAutonomousVmClusterCompartment operation to move an Autonomous Exadata VM cluster to a different compartment.
// For more information, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeAutonomousExadataInfrastructureCompartment.go.html to see an example of how to use ChangeAutonomousExadataInfrastructureCompartment API.
func (client DatabaseClient) ChangeAutonomousExadataInfrastructureCompartment(ctx context.Context, request ChangeAutonomousExadataInfrastructureCompartmentRequest) (response ChangeAutonomousExadataInfrastructureCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAutonomousExadataInfrastructureCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAutonomousExadataInfrastructureCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAutonomousExadataInfrastructureCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAutonomousExadataInfrastructureCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAutonomousExadataInfrastructureCompartmentResponse")
	}
	return
}

// changeAutonomousExadataInfrastructureCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeAutonomousExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousExadataInfrastructureCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/ChangeAutonomousExadataInfrastructureCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeAutonomousExadataInfrastructureCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAutonomousVmClusterCompartment Moves an Autonomous VM cluster and its dependent resources to another compartment. Applies to Exadata Cloud@Customer  only. For systems in the Oracle cloud, see ChangeAutonomousVmClusterCompartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeAutonomousVmClusterCompartment.go.html to see an example of how to use ChangeAutonomousVmClusterCompartment API.
func (client DatabaseClient) ChangeAutonomousVmClusterCompartment(ctx context.Context, request ChangeAutonomousVmClusterCompartmentRequest) (response ChangeAutonomousVmClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAutonomousVmClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAutonomousVmClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAutonomousVmClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAutonomousVmClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAutonomousVmClusterCompartmentResponse")
	}
	return
}

// changeAutonomousVmClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeAutonomousVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousVmClusters/{autonomousVmClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousVmClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/ChangeAutonomousVmClusterCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeAutonomousVmClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeBackupDestinationCompartment Move the backup destination and its dependent resources to the specified compartment.
// For more information, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeBackupDestinationCompartment.go.html to see an example of how to use ChangeBackupDestinationCompartment API.
func (client DatabaseClient) ChangeBackupDestinationCompartment(ctx context.Context, request ChangeBackupDestinationCompartmentRequest) (response ChangeBackupDestinationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeBackupDestinationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBackupDestinationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBackupDestinationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBackupDestinationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBackupDestinationCompartmentResponse")
	}
	return
}

// changeBackupDestinationCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeBackupDestinationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backupDestinations/{backupDestinationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBackupDestinationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/BackupDestination/ChangeBackupDestinationCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeBackupDestinationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCloudAutonomousVmClusterCompartment Moves an Autonomous Exadata VM cluster in the Oracle cloud and its dependent resources to another compartment. For Exadata Cloud@Customer systems, see ChangeAutonomousVmClusterCompartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeCloudAutonomousVmClusterCompartment.go.html to see an example of how to use ChangeCloudAutonomousVmClusterCompartment API.
func (client DatabaseClient) ChangeCloudAutonomousVmClusterCompartment(ctx context.Context, request ChangeCloudAutonomousVmClusterCompartmentRequest) (response ChangeCloudAutonomousVmClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCloudAutonomousVmClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCloudAutonomousVmClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCloudAutonomousVmClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCloudAutonomousVmClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCloudAutonomousVmClusterCompartmentResponse")
	}
	return
}

// changeCloudAutonomousVmClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeCloudAutonomousVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCloudAutonomousVmClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/ChangeCloudAutonomousVmClusterCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeCloudAutonomousVmClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCloudExadataInfrastructureCompartment Moves a cloud Exadata infrastructure resource and its dependent resources to another compartment. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.For more information about moving resources to a different compartment, see Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeCloudExadataInfrastructureCompartment.go.html to see an example of how to use ChangeCloudExadataInfrastructureCompartment API.
func (client DatabaseClient) ChangeCloudExadataInfrastructureCompartment(ctx context.Context, request ChangeCloudExadataInfrastructureCompartmentRequest) (response ChangeCloudExadataInfrastructureCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCloudExadataInfrastructureCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCloudExadataInfrastructureCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCloudExadataInfrastructureCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCloudExadataInfrastructureCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCloudExadataInfrastructureCompartmentResponse")
	}
	return
}

// changeCloudExadataInfrastructureCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeCloudExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCloudExadataInfrastructureCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/ChangeCloudExadataInfrastructureCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeCloudExadataInfrastructureCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCloudVmClusterCompartment Moves a cloud VM cluster and its dependent resources to another compartment. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeCloudVmClusterCompartment.go.html to see an example of how to use ChangeCloudVmClusterCompartment API.
func (client DatabaseClient) ChangeCloudVmClusterCompartment(ctx context.Context, request ChangeCloudVmClusterCompartmentRequest) (response ChangeCloudVmClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCloudVmClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCloudVmClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCloudVmClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCloudVmClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCloudVmClusterCompartmentResponse")
	}
	return
}

// changeCloudVmClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeCloudVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudVmClusters/{cloudVmClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCloudVmClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/ChangeCloudVmClusterCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeCloudVmClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseSoftwareImageCompartment Move the Database Software Image and its dependent resources to the specified compartment.
// For more information about moving Databse Software Images, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeDatabaseSoftwareImageCompartment.go.html to see an example of how to use ChangeDatabaseSoftwareImageCompartment API.
func (client DatabaseClient) ChangeDatabaseSoftwareImageCompartment(ctx context.Context, request ChangeDatabaseSoftwareImageCompartmentRequest) (response ChangeDatabaseSoftwareImageCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseSoftwareImageCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseSoftwareImageCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseSoftwareImageCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseSoftwareImageCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseSoftwareImageCompartmentResponse")
	}
	return
}

// changeDatabaseSoftwareImageCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeDatabaseSoftwareImageCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseSoftwareImages/{databaseSoftwareImageId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseSoftwareImageCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseSoftwareImage/ChangeDatabaseSoftwareImageCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeDatabaseSoftwareImageCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDataguardRole Switch the Autonomous Container Database role between Standby and Snapshot Standby.
// For more information about changing Autonomous Container Databases Dataguard Role, see
// Convert Physical Standby to Snapshot Standby (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html#ADBCL-GUID-D3B503F1-0032-4B0D-9F00-ACAE8151AB80) and Convert Snapshot Standby to Physical Standby (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html#ADBCL-GUID-E8D7E0EE-8244-467D-B33A-1BC6F969A0A4).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeDataguardRole.go.html to see an example of how to use ChangeDataguardRole API.
func (client DatabaseClient) ChangeDataguardRole(ctx context.Context, request ChangeDataguardRoleRequest) (response ChangeDataguardRoleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDataguardRole, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDataguardRoleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDataguardRoleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDataguardRoleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDataguardRoleResponse")
	}
	return
}

// changeDataguardRole implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeDataguardRole(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/changeDataguardRole", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDataguardRoleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/ChangeDataguardRole"
		err = common.PostProcessServiceError(err, "Database", "ChangeDataguardRole", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDbSystemCompartment Moves the DB system and its dependent resources to the specified compartment.
// For more information about moving DB systems, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeDbSystemCompartment.go.html to see an example of how to use ChangeDbSystemCompartment API.
func (client DatabaseClient) ChangeDbSystemCompartment(ctx context.Context, request ChangeDbSystemCompartmentRequest) (response ChangeDbSystemCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDbSystemCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDbSystemCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDbSystemCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDbSystemCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDbSystemCompartmentResponse")
	}
	return
}

// changeDbSystemCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeDbSystemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDbSystemCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/ChangeDbSystemCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeDbSystemCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDisasterRecoveryConfiguration This operation updates the cross-region disaster recovery (DR) details of the standby Autonomous Database Serverless database, and must be run on the standby side.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeDisasterRecoveryConfiguration.go.html to see an example of how to use ChangeDisasterRecoveryConfiguration API.
func (client DatabaseClient) ChangeDisasterRecoveryConfiguration(ctx context.Context, request ChangeDisasterRecoveryConfigurationRequest) (response ChangeDisasterRecoveryConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDisasterRecoveryConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDisasterRecoveryConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDisasterRecoveryConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDisasterRecoveryConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDisasterRecoveryConfigurationResponse")
	}
	return
}

// changeDisasterRecoveryConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeDisasterRecoveryConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/{autonomousDatabaseId}/actions/changeDisasterRecoveryConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDisasterRecoveryConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ChangeDisasterRecoveryConfiguration"
		err = common.PostProcessServiceError(err, "Database", "ChangeDisasterRecoveryConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExadataInfrastructureCompartment Moves an Exadata infrastructure resource and its dependent resources to another compartment. Applies to Exadata Cloud@Customer instances only.
// To move an Exadata Cloud Service infrastructure resource to another compartment, use the  ChangeCloudExadataInfrastructureCompartment operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeExadataInfrastructureCompartment.go.html to see an example of how to use ChangeExadataInfrastructureCompartment API.
func (client DatabaseClient) ChangeExadataInfrastructureCompartment(ctx context.Context, request ChangeExadataInfrastructureCompartmentRequest) (response ChangeExadataInfrastructureCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExadataInfrastructureCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExadataInfrastructureCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExadataInfrastructureCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExadataInfrastructureCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExadataInfrastructureCompartmentResponse")
	}
	return
}

// changeExadataInfrastructureCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExadataInfrastructureCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/ChangeExadataInfrastructureCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeExadataInfrastructureCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExternalContainerDatabaseCompartment Move the CreateExternalContainerDatabaseDetails
// and its dependent resources to the specified compartment.
// For more information about moving external container databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeExternalContainerDatabaseCompartment.go.html to see an example of how to use ChangeExternalContainerDatabaseCompartment API.
func (client DatabaseClient) ChangeExternalContainerDatabaseCompartment(ctx context.Context, request ChangeExternalContainerDatabaseCompartmentRequest) (response ChangeExternalContainerDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExternalContainerDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExternalContainerDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExternalContainerDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExternalContainerDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExternalContainerDatabaseCompartmentResponse")
	}
	return
}

// changeExternalContainerDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeExternalContainerDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalContainerDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/ChangeExternalContainerDatabaseCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeExternalContainerDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExternalNonContainerDatabaseCompartment Move the external non-container database and its dependent resources to the specified compartment.
// For more information about moving external non-container databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeExternalNonContainerDatabaseCompartment.go.html to see an example of how to use ChangeExternalNonContainerDatabaseCompartment API.
func (client DatabaseClient) ChangeExternalNonContainerDatabaseCompartment(ctx context.Context, request ChangeExternalNonContainerDatabaseCompartmentRequest) (response ChangeExternalNonContainerDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExternalNonContainerDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExternalNonContainerDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExternalNonContainerDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExternalNonContainerDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExternalNonContainerDatabaseCompartmentResponse")
	}
	return
}

// changeExternalNonContainerDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeExternalNonContainerDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalNonContainerDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/ChangeExternalNonContainerDatabaseCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeExternalNonContainerDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExternalPluggableDatabaseCompartment Move the CreateExternalPluggableDatabaseDetails and
// its dependent resources to the specified compartment.
// For more information about moving external pluggable databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeExternalPluggableDatabaseCompartment.go.html to see an example of how to use ChangeExternalPluggableDatabaseCompartment API.
func (client DatabaseClient) ChangeExternalPluggableDatabaseCompartment(ctx context.Context, request ChangeExternalPluggableDatabaseCompartmentRequest) (response ChangeExternalPluggableDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExternalPluggableDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExternalPluggableDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExternalPluggableDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExternalPluggableDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExternalPluggableDatabaseCompartmentResponse")
	}
	return
}

// changeExternalPluggableDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeExternalPluggableDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalPluggableDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/ChangeExternalPluggableDatabaseCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeExternalPluggableDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeKeyStoreCompartment Move the key store resource to the specified compartment.
// For more information about moving key stores, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeKeyStoreCompartment.go.html to see an example of how to use ChangeKeyStoreCompartment API.
func (client DatabaseClient) ChangeKeyStoreCompartment(ctx context.Context, request ChangeKeyStoreCompartmentRequest) (response ChangeKeyStoreCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeKeyStoreCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeKeyStoreCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeKeyStoreCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeKeyStoreCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeKeyStoreCompartmentResponse")
	}
	return
}

// changeKeyStoreCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeKeyStoreCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keyStores/{keyStoreId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeKeyStoreCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/KeyStore/ChangeKeyStoreCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeKeyStoreCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeKeyStoreType Changes encryption key management type
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeKeyStoreType.go.html to see an example of how to use ChangeKeyStoreType API.
func (client DatabaseClient) ChangeKeyStoreType(ctx context.Context, request ChangeKeyStoreTypeRequest) (response ChangeKeyStoreTypeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeKeyStoreType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeKeyStoreTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeKeyStoreTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeKeyStoreTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeKeyStoreTypeResponse")
	}
	return
}

// changeKeyStoreType implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeKeyStoreType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/changeKeyStoreType", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeKeyStoreTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/ChangeKeyStoreType"
		err = common.PostProcessServiceError(err, "Database", "ChangeKeyStoreType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOneoffPatchCompartment Move the one-off patch to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeOneoffPatchCompartment.go.html to see an example of how to use ChangeOneoffPatchCompartment API.
func (client DatabaseClient) ChangeOneoffPatchCompartment(ctx context.Context, request ChangeOneoffPatchCompartmentRequest) (response ChangeOneoffPatchCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOneoffPatchCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOneoffPatchCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOneoffPatchCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOneoffPatchCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOneoffPatchCompartmentResponse")
	}
	return
}

// changeOneoffPatchCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeOneoffPatchCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oneoffPatches/{oneoffPatchId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOneoffPatchCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/ChangeOneoffPatchCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeOneoffPatchCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVmClusterCompartment Moves a VM cluster and its dependent resources to another compartment. Applies to Exadata Cloud@Customer instances only.
// To move a cloud VM cluster in an Exadata Cloud Service instance to another compartment, use the ChangeCloudVmClusterCompartment operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ChangeVmClusterCompartment.go.html to see an example of how to use ChangeVmClusterCompartment API.
func (client DatabaseClient) ChangeVmClusterCompartment(ctx context.Context, request ChangeVmClusterCompartmentRequest) (response ChangeVmClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVmClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVmClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVmClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVmClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVmClusterCompartmentResponse")
	}
	return
}

// changeVmClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) changeVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusters/{vmClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVmClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/ChangeVmClusterCompartment"
		err = common.PostProcessServiceError(err, "Database", "ChangeVmClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CheckExternalDatabaseConnectorConnectionStatus Check the status of the external database connection specified in this connector.
// This operation will refresh the connectionStatus and timeConnectionStatusLastUpdated fields.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CheckExternalDatabaseConnectorConnectionStatus.go.html to see an example of how to use CheckExternalDatabaseConnectorConnectionStatus API.
func (client DatabaseClient) CheckExternalDatabaseConnectorConnectionStatus(ctx context.Context, request CheckExternalDatabaseConnectorConnectionStatusRequest) (response CheckExternalDatabaseConnectorConnectionStatusResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.checkExternalDatabaseConnectorConnectionStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CheckExternalDatabaseConnectorConnectionStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CheckExternalDatabaseConnectorConnectionStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CheckExternalDatabaseConnectorConnectionStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CheckExternalDatabaseConnectorConnectionStatusResponse")
	}
	return
}

// checkExternalDatabaseConnectorConnectionStatus implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) checkExternalDatabaseConnectorConnectionStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externaldatabaseconnectors/{externalDatabaseConnectorId}/actions/checkConnectionStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CheckExternalDatabaseConnectorConnectionStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalDatabaseConnector/CheckExternalDatabaseConnectorConnectionStatus"
		err = common.PostProcessServiceError(err, "Database", "CheckExternalDatabaseConnectorConnectionStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CompleteExternalBackupJob Changes the status of the standalone backup resource to `ACTIVE` after the backup is created from the on-premises database and placed in Oracle Cloud Infrastructure Object Storage.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.cloud.oracle.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CompleteExternalBackupJob.go.html to see an example of how to use CompleteExternalBackupJob API.
func (client DatabaseClient) CompleteExternalBackupJob(ctx context.Context, request CompleteExternalBackupJobRequest) (response CompleteExternalBackupJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.completeExternalBackupJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompleteExternalBackupJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompleteExternalBackupJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompleteExternalBackupJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompleteExternalBackupJobResponse")
	}
	return
}

// completeExternalBackupJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) completeExternalBackupJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalBackupJobs/{backupId}/actions/complete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompleteExternalBackupJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalBackupJob/CompleteExternalBackupJob"
		err = common.PostProcessServiceError(err, "Database", "CompleteExternalBackupJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureAutonomousDatabaseVaultKey Configures the Autonomous Database Vault service key (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ConfigureAutonomousDatabaseVaultKey.go.html to see an example of how to use ConfigureAutonomousDatabaseVaultKey API.
func (client DatabaseClient) ConfigureAutonomousDatabaseVaultKey(ctx context.Context, request ConfigureAutonomousDatabaseVaultKeyRequest) (response ConfigureAutonomousDatabaseVaultKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureAutonomousDatabaseVaultKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureAutonomousDatabaseVaultKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureAutonomousDatabaseVaultKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureAutonomousDatabaseVaultKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureAutonomousDatabaseVaultKeyResponse")
	}
	return
}

// configureAutonomousDatabaseVaultKey implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) configureAutonomousDatabaseVaultKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/configureAutonomousDatabaseVaultKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureAutonomousDatabaseVaultKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ConfigureAutonomousDatabaseVaultKey"
		err = common.PostProcessServiceError(err, "Database", "ConfigureAutonomousDatabaseVaultKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureSaasAdminUser This operation updates SaaS administrative user configuration of the Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ConfigureSaasAdminUser.go.html to see an example of how to use ConfigureSaasAdminUser API.
func (client DatabaseClient) ConfigureSaasAdminUser(ctx context.Context, request ConfigureSaasAdminUserRequest) (response ConfigureSaasAdminUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.configureSaasAdminUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureSaasAdminUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureSaasAdminUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureSaasAdminUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureSaasAdminUserResponse")
	}
	return
}

// configureSaasAdminUser implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) configureSaasAdminUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/configureSaasAdminUser", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureSaasAdminUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ConfigureSaasAdminUser"
		err = common.PostProcessServiceError(err, "Database", "ConfigureSaasAdminUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConvertToPdb Converts a non-container database to a pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ConvertToPdb.go.html to see an example of how to use ConvertToPdb API.
func (client DatabaseClient) ConvertToPdb(ctx context.Context, request ConvertToPdbRequest) (response ConvertToPdbResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.convertToPdb, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConvertToPdbResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConvertToPdbResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConvertToPdbResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConvertToPdbResponse")
	}
	return
}

// convertToPdb implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) convertToPdb(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/convertToPdb", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConvertToPdbResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/ConvertToPdb"
		err = common.PostProcessServiceError(err, "Database", "ConvertToPdb", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConvertToRegularPluggableDatabase Converts a Refreshable clone to Regular pluggable database (PDB).
// Pluggable Database will be in `READ_WRITE` openmode after conversion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ConvertToRegularPluggableDatabase.go.html to see an example of how to use ConvertToRegularPluggableDatabase API.
func (client DatabaseClient) ConvertToRegularPluggableDatabase(ctx context.Context, request ConvertToRegularPluggableDatabaseRequest) (response ConvertToRegularPluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.convertToRegularPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConvertToRegularPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConvertToRegularPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConvertToRegularPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConvertToRegularPluggableDatabaseResponse")
	}
	return
}

// convertToRegularPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) convertToRegularPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/convertToRegular", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConvertToRegularPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/ConvertToRegularPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "ConvertToRegularPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateApplicationVip Creates a new application virtual IP (VIP) address in the specified cloud VM cluster based on the request parameters you provide.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateApplicationVip.go.html to see an example of how to use CreateApplicationVip API.
func (client DatabaseClient) CreateApplicationVip(ctx context.Context, request CreateApplicationVipRequest) (response CreateApplicationVipResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createApplicationVip, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApplicationVipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApplicationVipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApplicationVipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApplicationVipResponse")
	}
	return
}

// createApplicationVip implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createApplicationVip(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applicationVip", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApplicationVipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ApplicationVip/CreateApplicationVip"
		err = common.PostProcessServiceError(err, "Database", "CreateApplicationVip", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAutonomousContainerDatabase Creates an Autonomous Container Database in the specified Autonomous Exadata Infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateAutonomousContainerDatabase.go.html to see an example of how to use CreateAutonomousContainerDatabase API.
func (client DatabaseClient) CreateAutonomousContainerDatabase(ctx context.Context, request CreateAutonomousContainerDatabaseRequest) (response CreateAutonomousContainerDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutonomousContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutonomousContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousContainerDatabaseResponse")
	}
	return
}

// createAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/CreateAutonomousContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreateAutonomousContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAutonomousContainerDatabaseDataguardAssociation Create a new Autonomous Data Guard association. An Autonomous Data Guard association represents the replication relationship between the
// specified Autonomous Container database and a peer Autonomous Container database. For more information, see Using Oracle Data Guard (https://docs.cloud.oracle.com/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. For more information, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateAutonomousContainerDatabaseDataguardAssociation.go.html to see an example of how to use CreateAutonomousContainerDatabaseDataguardAssociation API.
func (client DatabaseClient) CreateAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request CreateAutonomousContainerDatabaseDataguardAssociationRequest) (response CreateAutonomousContainerDatabaseDataguardAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousContainerDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutonomousContainerDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutonomousContainerDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousContainerDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousContainerDatabaseDataguardAssociationResponse")
	}
	return
}

// createAutonomousContainerDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousContainerDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/CreateAutonomousContainerDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "CreateAutonomousContainerDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAutonomousDatabase Creates a new Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateAutonomousDatabase.go.html to see an example of how to use CreateAutonomousDatabase API.
func (client DatabaseClient) CreateAutonomousDatabase(ctx context.Context, request CreateAutonomousDatabaseRequest) (response CreateAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousDatabaseResponse")
	}
	return
}

// createAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/CreateAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreateAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAutonomousDatabaseBackup Creates a new Autonomous Database backup for the specified database based on the provided request parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateAutonomousDatabaseBackup.go.html to see an example of how to use CreateAutonomousDatabaseBackup API.
func (client DatabaseClient) CreateAutonomousDatabaseBackup(ctx context.Context, request CreateAutonomousDatabaseBackupRequest) (response CreateAutonomousDatabaseBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutonomousDatabaseBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutonomousDatabaseBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousDatabaseBackupResponse")
	}
	return
}

// createAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabaseBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDatabaseBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseBackup/CreateAutonomousDatabaseBackup"
		err = common.PostProcessServiceError(err, "Database", "CreateAutonomousDatabaseBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAutonomousVmCluster Creates an Autonomous VM cluster for Exadata Cloud@Customer. To create an Autonomous VM Cluster in the Oracle cloud, see CreateCloudAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateAutonomousVmCluster.go.html to see an example of how to use CreateAutonomousVmCluster API.
func (client DatabaseClient) CreateAutonomousVmCluster(ctx context.Context, request CreateAutonomousVmClusterRequest) (response CreateAutonomousVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutonomousVmClusterResponse")
	}
	return
}

// createAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/CreateAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "CreateAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBackup Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateBackup.go.html to see an example of how to use CreateBackup API.
func (client DatabaseClient) CreateBackup(ctx context.Context, request CreateBackupRequest) (response CreateBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBackupResponse")
	}
	return
}

// createBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Backup/CreateBackup"
		err = common.PostProcessServiceError(err, "Database", "CreateBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBackupDestination Creates a backup destination in an Exadata Cloud@Customer system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateBackupDestination.go.html to see an example of how to use CreateBackupDestination API.
func (client DatabaseClient) CreateBackupDestination(ctx context.Context, request CreateBackupDestinationRequest) (response CreateBackupDestinationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBackupDestination, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBackupDestinationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBackupDestinationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBackupDestinationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBackupDestinationResponse")
	}
	return
}

// createBackupDestination implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backupDestinations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBackupDestinationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/BackupDestination/CreateBackupDestination"
		err = common.PostProcessServiceError(err, "Database", "CreateBackupDestination", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCloudAutonomousVmCluster Creates an Autonomous Exadata VM cluster in the Oracle cloud. For Exadata Cloud@Customer systems, see CreateAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateCloudAutonomousVmCluster.go.html to see an example of how to use CreateCloudAutonomousVmCluster API.
func (client DatabaseClient) CreateCloudAutonomousVmCluster(ctx context.Context, request CreateCloudAutonomousVmClusterRequest) (response CreateCloudAutonomousVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCloudAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCloudAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCloudAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCloudAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCloudAutonomousVmClusterResponse")
	}
	return
}

// createCloudAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createCloudAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudAutonomousVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCloudAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/CreateCloudAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "CreateCloudAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCloudExadataInfrastructure Creates a cloud Exadata infrastructure resource. This resource is used to create either an Exadata Cloud Service (https://docs.cloud.oracle.com/Content/Database/Concepts/exaoverview.htm) instance or an Autonomous Database on dedicated Exadata infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateCloudExadataInfrastructure.go.html to see an example of how to use CreateCloudExadataInfrastructure API.
func (client DatabaseClient) CreateCloudExadataInfrastructure(ctx context.Context, request CreateCloudExadataInfrastructureRequest) (response CreateCloudExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCloudExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCloudExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCloudExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCloudExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCloudExadataInfrastructureResponse")
	}
	return
}

// createCloudExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudExadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCloudExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/CreateCloudExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "CreateCloudExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCloudVmCluster Creates a cloud VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateCloudVmCluster.go.html to see an example of how to use CreateCloudVmCluster API.
func (client DatabaseClient) CreateCloudVmCluster(ctx context.Context, request CreateCloudVmClusterRequest) (response CreateCloudVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCloudVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCloudVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCloudVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCloudVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCloudVmClusterResponse")
	}
	return
}

// createCloudVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCloudVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/CreateCloudVmCluster"
		err = common.PostProcessServiceError(err, "Database", "CreateCloudVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConsoleConnection Creates a new console connection to the specified database node.
// After the console connection has been created and is available,
// you connect to the console using SSH.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateConsoleConnection.go.html to see an example of how to use CreateConsoleConnection API.
func (client DatabaseClient) CreateConsoleConnection(ctx context.Context, request CreateConsoleConnectionRequest) (response CreateConsoleConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConsoleConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConsoleConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConsoleConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConsoleConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConsoleConnectionResponse")
	}
	return
}

// createConsoleConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbNodes/{dbNodeId}/consoleConnections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConsoleConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleConnection/CreateConsoleConnection"
		err = common.PostProcessServiceError(err, "Database", "CreateConsoleConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConsoleHistory Captures the most recent serial console data (up to a megabyte) for the specified database node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateConsoleHistory.go.html to see an example of how to use CreateConsoleHistory API.
func (client DatabaseClient) CreateConsoleHistory(ctx context.Context, request CreateConsoleHistoryRequest) (response CreateConsoleHistoryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConsoleHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConsoleHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConsoleHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConsoleHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConsoleHistoryResponse")
	}
	return
}

// createConsoleHistory implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createConsoleHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbNodes/{dbNodeId}/consoleHistories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConsoleHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleHistory/CreateConsoleHistory"
		err = common.PostProcessServiceError(err, "Database", "CreateConsoleHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataGuardAssociation Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
// specified database and a peer database. For more information, see Using Oracle Data Guard (https://docs.cloud.oracle.com/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. For more information, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateDataGuardAssociation.go.html to see an example of how to use CreateDataGuardAssociation API.
func (client DatabaseClient) CreateDataGuardAssociation(ctx context.Context, request CreateDataGuardAssociationRequest) (response CreateDataGuardAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataGuardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataGuardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataGuardAssociationResponse")
	}
	return
}

// createDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataGuardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/CreateDataGuardAssociation"
		err = common.PostProcessServiceError(err, "Database", "CreateDataGuardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabase Creates a new database in the specified Database Home. If the database version is provided, it must match the version of the Database Home. Applies to Exadata and Exadata Cloud@Customer systems.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateDatabase.go.html to see an example of how to use CreateDatabase API.
func (client DatabaseClient) CreateDatabase(ctx context.Context, request CreateDatabaseRequest) (response CreateDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseResponse")
	}
	return
}

// createDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/CreateDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreateDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabaseSoftwareImage create database software image in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateDatabaseSoftwareImage.go.html to see an example of how to use CreateDatabaseSoftwareImage API.
func (client DatabaseClient) CreateDatabaseSoftwareImage(ctx context.Context, request CreateDatabaseSoftwareImageRequest) (response CreateDatabaseSoftwareImageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseSoftwareImageResponse")
	}
	return
}

// createDatabaseSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseSoftwareImages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseSoftwareImage/CreateDatabaseSoftwareImage"
		err = common.PostProcessServiceError(err, "Database", "CreateDatabaseSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDbHome Creates a new Database Home in the specified database system based on the request parameters you provide. Applies to bare metal DB systems, Exadata systems, and Exadata Cloud@Customer systems.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateDbHome.go.html to see an example of how to use CreateDbHome API.
func (client DatabaseClient) CreateDbHome(ctx context.Context, request CreateDbHomeRequest) (response CreateDbHomeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDbHomeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDbHomeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbHomeResponse")
	}
	return
}

// createDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbHomes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDbHomeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbHome/CreateDbHome"
		err = common.PostProcessServiceError(err, "Database", "CreateDbHome", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExadataInfrastructure Creates an Exadata infrastructure resource. Applies to Exadata Cloud@Customer instances only.
// To create an Exadata Cloud Service infrastructure resource, use the  CreateCloudExadataInfrastructure operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateExadataInfrastructure.go.html to see an example of how to use CreateExadataInfrastructure API.
func (client DatabaseClient) CreateExadataInfrastructure(ctx context.Context, request CreateExadataInfrastructureRequest) (response CreateExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExadataInfrastructureResponse")
	}
	return
}

// createExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/CreateExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "CreateExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalBackupJob Creates a new backup resource and returns the information the caller needs to back up an on-premises Oracle Database to Oracle Cloud Infrastructure.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.cloud.oracle.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateExternalBackupJob.go.html to see an example of how to use CreateExternalBackupJob API.
func (client DatabaseClient) CreateExternalBackupJob(ctx context.Context, request CreateExternalBackupJobRequest) (response CreateExternalBackupJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalBackupJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalBackupJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalBackupJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalBackupJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalBackupJobResponse")
	}
	return
}

// createExternalBackupJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExternalBackupJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalBackupJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalBackupJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalBackupJob/CreateExternalBackupJob"
		err = common.PostProcessServiceError(err, "Database", "CreateExternalBackupJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalContainerDatabase Creates a new external container database resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateExternalContainerDatabase.go.html to see an example of how to use CreateExternalContainerDatabase API.
func (client DatabaseClient) CreateExternalContainerDatabase(ctx context.Context, request CreateExternalContainerDatabaseRequest) (response CreateExternalContainerDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalContainerDatabaseResponse")
	}
	return
}

// createExternalContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/CreateExternalContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreateExternalContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalDatabaseConnector Creates a new external database connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateExternalDatabaseConnector.go.html to see an example of how to use CreateExternalDatabaseConnector API.
func (client DatabaseClient) CreateExternalDatabaseConnector(ctx context.Context, request CreateExternalDatabaseConnectorRequest) (response CreateExternalDatabaseConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalDatabaseConnectorResponse")
	}
	return
}

// createExternalDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externaldatabaseconnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalDatabaseConnector/CreateExternalDatabaseConnector"
		err = common.PostProcessServiceError(err, "Database", "CreateExternalDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldatabaseconnector{})
	return response, err
}

// CreateExternalNonContainerDatabase Creates a new ExternalNonContainerDatabase resource
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateExternalNonContainerDatabase.go.html to see an example of how to use CreateExternalNonContainerDatabase API.
func (client DatabaseClient) CreateExternalNonContainerDatabase(ctx context.Context, request CreateExternalNonContainerDatabaseRequest) (response CreateExternalNonContainerDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalNonContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalNonContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalNonContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalNonContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalNonContainerDatabaseResponse")
	}
	return
}

// createExternalNonContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalNonContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/CreateExternalNonContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreateExternalNonContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalPluggableDatabase Registers a new CreateExternalPluggableDatabaseDetails
// resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateExternalPluggableDatabase.go.html to see an example of how to use CreateExternalPluggableDatabase API.
func (client DatabaseClient) CreateExternalPluggableDatabase(ctx context.Context, request CreateExternalPluggableDatabaseRequest) (response CreateExternalPluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalPluggableDatabaseResponse")
	}
	return
}

// createExternalPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/CreateExternalPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreateExternalPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateKeyStore Creates a Key Store.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateKeyStore.go.html to see an example of how to use CreateKeyStore API.
func (client DatabaseClient) CreateKeyStore(ctx context.Context, request CreateKeyStoreRequest) (response CreateKeyStoreResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createKeyStore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateKeyStoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateKeyStoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKeyStoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKeyStoreResponse")
	}
	return
}

// createKeyStore implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keyStores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateKeyStoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/KeyStore/CreateKeyStore"
		err = common.PostProcessServiceError(err, "Database", "CreateKeyStore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMaintenanceRun Creates a maintenance run with one of the following:
// 1. The latest available release update patch (RUP) for the Autonomous Container Database.
// 2. The latest available RUP and DST time-zone (TZ) file updates for the Autonomous Container Database.
// 3. The DST TZ file updates for the Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateMaintenanceRun.go.html to see an example of how to use CreateMaintenanceRun API.
func (client DatabaseClient) CreateMaintenanceRun(ctx context.Context, request CreateMaintenanceRunRequest) (response CreateMaintenanceRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMaintenanceRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMaintenanceRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMaintenanceRunResponse")
	}
	return
}

// createMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maintenanceRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMaintenanceRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/MaintenanceRun/CreateMaintenanceRun"
		err = common.PostProcessServiceError(err, "Database", "CreateMaintenanceRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOneoffPatch Creates one-off patch for specified database version to download.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateOneoffPatch.go.html to see an example of how to use CreateOneoffPatch API.
func (client DatabaseClient) CreateOneoffPatch(ctx context.Context, request CreateOneoffPatchRequest) (response CreateOneoffPatchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOneoffPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOneoffPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOneoffPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOneoffPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOneoffPatchResponse")
	}
	return
}

// createOneoffPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createOneoffPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oneoffPatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOneoffPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/CreateOneoffPatch"
		err = common.PostProcessServiceError(err, "Database", "CreateOneoffPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePluggableDatabase Creates and starts a pluggable database in the specified container database.
// Pluggable Database can be created using different operations (e.g. LocalClone, RemoteClone, Relocate ) with this API.
// Use the StartPluggableDatabase and StopPluggableDatabase APIs to start and stop the pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreatePluggableDatabase.go.html to see an example of how to use CreatePluggableDatabase API.
func (client DatabaseClient) CreatePluggableDatabase(ctx context.Context, request CreatePluggableDatabaseRequest) (response CreatePluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePluggableDatabaseResponse")
	}
	return
}

// createPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/CreatePluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "CreatePluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVmCluster Creates an Exadata Cloud@Customer VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateVmCluster.go.html to see an example of how to use CreateVmCluster API.
func (client DatabaseClient) CreateVmCluster(ctx context.Context, request CreateVmClusterRequest) (response CreateVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVmClusterResponse")
	}
	return
}

// createVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/CreateVmCluster"
		err = common.PostProcessServiceError(err, "Database", "CreateVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVmClusterNetwork Creates the VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To create a cloud VM cluster in an Exadata Cloud Service instance, use the CreateCloudVmCluster operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/CreateVmClusterNetwork.go.html to see an example of how to use CreateVmClusterNetwork API.
func (client DatabaseClient) CreateVmClusterNetwork(ctx context.Context, request CreateVmClusterNetworkRequest) (response CreateVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVmClusterNetworkResponse")
	}
	return
}

// createVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) createVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/CreateVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "CreateVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DbNodeAction Performs one of the following power actions on the specified DB node:
// - start - power on
// - stop - power off
// - softreset - ACPI shutdown and power on
// - reset - power off and power on
// **Note:** Stopping a node affects billing differently, depending on the type of DB system:
// *Bare metal and Exadata systems* - The _stop_ state has no effect on the resources you consume.
// Billing continues for DB nodes that you stop, and related resources continue
// to apply against any relevant quotas. You must terminate the DB system
// (TerminateDbSystem)
// to remove its resources from billing and quotas.
// *Virtual machine DB systems* - Stopping a node stops billing for all OCPUs associated with that node, and billing resumes when you restart the node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DbNodeAction.go.html to see an example of how to use DbNodeAction API.
func (client DatabaseClient) DbNodeAction(ctx context.Context, request DbNodeActionRequest) (response DbNodeActionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.dbNodeAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DbNodeActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DbNodeActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DbNodeActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DbNodeActionResponse")
	}
	return
}

// dbNodeAction implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) dbNodeAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbNodes/{dbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DbNodeActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbNode/DbNodeAction"
		err = common.PostProcessServiceError(err, "Database", "DbNodeAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApplicationVip Deletes and deregisters the specified application virtual IP (VIP) address.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteApplicationVip.go.html to see an example of how to use DeleteApplicationVip API.
func (client DatabaseClient) DeleteApplicationVip(ctx context.Context, request DeleteApplicationVipRequest) (response DeleteApplicationVipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApplicationVip, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApplicationVipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApplicationVipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplicationVipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplicationVipResponse")
	}
	return
}

// deleteApplicationVip implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteApplicationVip(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/applicationVip/{applicationVipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApplicationVipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ApplicationVip/DeleteApplicationVip"
		err = common.PostProcessServiceError(err, "Database", "DeleteApplicationVip", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAutonomousDatabase Deletes the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteAutonomousDatabase.go.html to see an example of how to use DeleteAutonomousDatabase API.
func (client DatabaseClient) DeleteAutonomousDatabase(ctx context.Context, request DeleteAutonomousDatabaseRequest) (response DeleteAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutonomousDatabaseResponse")
	}
	return
}

// deleteAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDatabases/{autonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/DeleteAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "DeleteAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAutonomousDatabaseBackup Deletes a long-term backup. You cannot delete other backups using this API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteAutonomousDatabaseBackup.go.html to see an example of how to use DeleteAutonomousDatabaseBackup API.
func (client DatabaseClient) DeleteAutonomousDatabaseBackup(ctx context.Context, request DeleteAutonomousDatabaseBackupRequest) (response DeleteAutonomousDatabaseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAutonomousDatabaseBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAutonomousDatabaseBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutonomousDatabaseBackupResponse")
	}
	return
}

// deleteAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDatabaseBackups/{autonomousDatabaseBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousDatabaseBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseBackup/DeleteAutonomousDatabaseBackup"
		err = common.PostProcessServiceError(err, "Database", "DeleteAutonomousDatabaseBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAutonomousVmCluster Deletes the specified Autonomous VM cluster in an Exadata Cloud@Customer system. To delete an Autonomous VM Cluster in the Oracle cloud, see DeleteCloudAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteAutonomousVmCluster.go.html to see an example of how to use DeleteAutonomousVmCluster API.
func (client DatabaseClient) DeleteAutonomousVmCluster(ctx context.Context, request DeleteAutonomousVmClusterRequest) (response DeleteAutonomousVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutonomousVmClusterResponse")
	}
	return
}

// deleteAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousVmClusters/{autonomousVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/DeleteAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "DeleteAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBackup Deletes a full backup. You cannot delete automatic backups using this API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteBackup.go.html to see an example of how to use DeleteBackup API.
func (client DatabaseClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest) (response DeleteBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBackupResponse")
	}
	return
}

// deleteBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/backups/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Backup/DeleteBackup"
		err = common.PostProcessServiceError(err, "Database", "DeleteBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBackupDestination Deletes a backup destination in an Exadata Cloud@Customer system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteBackupDestination.go.html to see an example of how to use DeleteBackupDestination API.
func (client DatabaseClient) DeleteBackupDestination(ctx context.Context, request DeleteBackupDestinationRequest) (response DeleteBackupDestinationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackupDestination, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBackupDestinationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBackupDestinationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBackupDestinationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBackupDestinationResponse")
	}
	return
}

// deleteBackupDestination implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/backupDestinations/{backupDestinationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBackupDestinationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/BackupDestination/DeleteBackupDestination"
		err = common.PostProcessServiceError(err, "Database", "DeleteBackupDestination", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCloudAutonomousVmCluster Deletes the specified Autonomous Exadata VM cluster in the Oracle cloud. For Exadata Cloud@Customer systems, see DeleteAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteCloudAutonomousVmCluster.go.html to see an example of how to use DeleteCloudAutonomousVmCluster API.
func (client DatabaseClient) DeleteCloudAutonomousVmCluster(ctx context.Context, request DeleteCloudAutonomousVmClusterRequest) (response DeleteCloudAutonomousVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCloudAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCloudAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCloudAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCloudAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCloudAutonomousVmClusterResponse")
	}
	return
}

// deleteCloudAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteCloudAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCloudAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/DeleteCloudAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "DeleteCloudAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCloudExadataInfrastructure Deletes the cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteCloudExadataInfrastructure.go.html to see an example of how to use DeleteCloudExadataInfrastructure API.
func (client DatabaseClient) DeleteCloudExadataInfrastructure(ctx context.Context, request DeleteCloudExadataInfrastructureRequest) (response DeleteCloudExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCloudExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCloudExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCloudExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCloudExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCloudExadataInfrastructureResponse")
	}
	return
}

// deleteCloudExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCloudExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/DeleteCloudExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "DeleteCloudExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCloudVmCluster Deletes the specified cloud VM cluster. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteCloudVmCluster.go.html to see an example of how to use DeleteCloudVmCluster API.
func (client DatabaseClient) DeleteCloudVmCluster(ctx context.Context, request DeleteCloudVmClusterRequest) (response DeleteCloudVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCloudVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCloudVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCloudVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCloudVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCloudVmClusterResponse")
	}
	return
}

// deleteCloudVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cloudVmClusters/{cloudVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCloudVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/DeleteCloudVmCluster"
		err = common.PostProcessServiceError(err, "Database", "DeleteCloudVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConsoleConnection Deletes the specified database node console connection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteConsoleConnection.go.html to see an example of how to use DeleteConsoleConnection API.
func (client DatabaseClient) DeleteConsoleConnection(ctx context.Context, request DeleteConsoleConnectionRequest) (response DeleteConsoleConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConsoleConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConsoleConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConsoleConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConsoleConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConsoleConnectionResponse")
	}
	return
}

// deleteConsoleConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbNodes/{dbNodeId}/consoleConnections/{consoleConnectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConsoleConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleConnection/DeleteConsoleConnection"
		err = common.PostProcessServiceError(err, "Database", "DeleteConsoleConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConsoleHistory Deletes the specified database node console history.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteConsoleHistory.go.html to see an example of how to use DeleteConsoleHistory API.
func (client DatabaseClient) DeleteConsoleHistory(ctx context.Context, request DeleteConsoleHistoryRequest) (response DeleteConsoleHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConsoleHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConsoleHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConsoleHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConsoleHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConsoleHistoryResponse")
	}
	return
}

// deleteConsoleHistory implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteConsoleHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbNodes/{dbNodeId}/consoleHistories/{consoleHistoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConsoleHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleHistory/DeleteConsoleHistory"
		err = common.PostProcessServiceError(err, "Database", "DeleteConsoleHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabase Deletes the specified database. Applies only to Exadata systems.
// The data in this database is local to the Exadata system and will be lost when the database is deleted. Oracle recommends that you back up any data in the Exadata system prior to deleting it. You can use the `performFinalBackup` parameter to have the Exadata system database backed up before it is deleted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteDatabase.go.html to see an example of how to use DeleteDatabase API.
func (client DatabaseClient) DeleteDatabase(ctx context.Context, request DeleteDatabaseRequest) (response DeleteDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseResponse")
	}
	return
}

// deleteDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databases/{databaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/DeleteDatabase"
		err = common.PostProcessServiceError(err, "Database", "DeleteDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseSoftwareImage Delete a database software image
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteDatabaseSoftwareImage.go.html to see an example of how to use DeleteDatabaseSoftwareImage API.
func (client DatabaseClient) DeleteDatabaseSoftwareImage(ctx context.Context, request DeleteDatabaseSoftwareImageRequest) (response DeleteDatabaseSoftwareImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseSoftwareImageResponse")
	}
	return
}

// deleteDatabaseSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseSoftwareImages/{databaseSoftwareImageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseSoftwareImage/DeleteDatabaseSoftwareImage"
		err = common.PostProcessServiceError(err, "Database", "DeleteDatabaseSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDbHome Deletes a Database Home. Applies to bare metal DB systems, Exadata Cloud Service, and Exadata Cloud@Customer systems.
// Oracle recommends that you use the `performFinalBackup` parameter to back up any data on a bare metal DB system before you delete a Database Home. On an Exadata Cloud@Customer system or an Exadata Cloud Service system, you can delete a Database Home only when there are no databases in it and therefore you cannot use the `performFinalBackup` parameter to back up data.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteDbHome.go.html to see an example of how to use DeleteDbHome API.
func (client DatabaseClient) DeleteDbHome(ctx context.Context, request DeleteDbHomeRequest) (response DeleteDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDbHomeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDbHomeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbHomeResponse")
	}
	return
}

// deleteDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbHomes/{dbHomeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDbHomeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbHome/DeleteDbHome"
		err = common.PostProcessServiceError(err, "Database", "DeleteDbHome", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExadataInfrastructure Deletes the Exadata Cloud@Customer infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteExadataInfrastructure.go.html to see an example of how to use DeleteExadataInfrastructure API.
func (client DatabaseClient) DeleteExadataInfrastructure(ctx context.Context, request DeleteExadataInfrastructureRequest) (response DeleteExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExadataInfrastructureResponse")
	}
	return
}

// deleteExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exadataInfrastructures/{exadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/DeleteExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "DeleteExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalContainerDatabase Deletes the CreateExternalContainerDatabaseDetails
// resource. Any external pluggable databases registered under this container database must be deleted in
// your Oracle Cloud Infrastructure tenancy prior to this operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteExternalContainerDatabase.go.html to see an example of how to use DeleteExternalContainerDatabase API.
func (client DatabaseClient) DeleteExternalContainerDatabase(ctx context.Context, request DeleteExternalContainerDatabaseRequest) (response DeleteExternalContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalContainerDatabaseResponse")
	}
	return
}

// deleteExternalContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalcontainerdatabases/{externalContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/DeleteExternalContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "DeleteExternalContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalDatabaseConnector Deletes an external database connector.
// Any services enabled using the external database connector must be
// deleted prior to this operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteExternalDatabaseConnector.go.html to see an example of how to use DeleteExternalDatabaseConnector API.
func (client DatabaseClient) DeleteExternalDatabaseConnector(ctx context.Context, request DeleteExternalDatabaseConnectorRequest) (response DeleteExternalDatabaseConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalDatabaseConnectorResponse")
	}
	return
}

// deleteExternalDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externaldatabaseconnectors/{externalDatabaseConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalDatabaseConnector/DeleteExternalDatabaseConnector"
		err = common.PostProcessServiceError(err, "Database", "DeleteExternalDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalNonContainerDatabase Deletes the Oracle Cloud Infrastructure resource representing an external non-container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteExternalNonContainerDatabase.go.html to see an example of how to use DeleteExternalNonContainerDatabase API.
func (client DatabaseClient) DeleteExternalNonContainerDatabase(ctx context.Context, request DeleteExternalNonContainerDatabaseRequest) (response DeleteExternalNonContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalNonContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalNonContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalNonContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalNonContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalNonContainerDatabaseResponse")
	}
	return
}

// deleteExternalNonContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalNonContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/DeleteExternalNonContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "DeleteExternalNonContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalPluggableDatabase Deletes the CreateExternalPluggableDatabaseDetails.
// resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteExternalPluggableDatabase.go.html to see an example of how to use DeleteExternalPluggableDatabase API.
func (client DatabaseClient) DeleteExternalPluggableDatabase(ctx context.Context, request DeleteExternalPluggableDatabaseRequest) (response DeleteExternalPluggableDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalPluggableDatabaseResponse")
	}
	return
}

// deleteExternalPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalpluggabledatabases/{externalPluggableDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/DeleteExternalPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "DeleteExternalPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteKeyStore Deletes a key store.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteKeyStore.go.html to see an example of how to use DeleteKeyStore API.
func (client DatabaseClient) DeleteKeyStore(ctx context.Context, request DeleteKeyStoreRequest) (response DeleteKeyStoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteKeyStore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteKeyStoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteKeyStoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteKeyStoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteKeyStoreResponse")
	}
	return
}

// deleteKeyStore implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/keyStores/{keyStoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteKeyStoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/KeyStore/DeleteKeyStore"
		err = common.PostProcessServiceError(err, "Database", "DeleteKeyStore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOneoffPatch Deletes a one-off patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteOneoffPatch.go.html to see an example of how to use DeleteOneoffPatch API.
func (client DatabaseClient) DeleteOneoffPatch(ctx context.Context, request DeleteOneoffPatchRequest) (response DeleteOneoffPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOneoffPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOneoffPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOneoffPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOneoffPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOneoffPatchResponse")
	}
	return
}

// deleteOneoffPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteOneoffPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oneoffPatches/{oneoffPatchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOneoffPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/DeleteOneoffPatch"
		err = common.PostProcessServiceError(err, "Database", "DeleteOneoffPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePluggableDatabase Deletes the specified pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeletePluggableDatabase.go.html to see an example of how to use DeletePluggableDatabase API.
func (client DatabaseClient) DeletePluggableDatabase(ctx context.Context, request DeletePluggableDatabaseRequest) (response DeletePluggableDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePluggableDatabaseResponse")
	}
	return
}

// deletePluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deletePluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/pluggableDatabases/{pluggableDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/DeletePluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "DeletePluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVmCluster Deletes the specified VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteVmCluster.go.html to see an example of how to use DeleteVmCluster API.
func (client DatabaseClient) DeleteVmCluster(ctx context.Context, request DeleteVmClusterRequest) (response DeleteVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVmClusterResponse")
	}
	return
}

// deleteVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/vmClusters/{vmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/DeleteVmCluster"
		err = common.PostProcessServiceError(err, "Database", "DeleteVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVmClusterNetwork Deletes the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To delete a cloud VM cluster in an Exadata Cloud Service instance, use the DeleteCloudVmCluster operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeleteVmClusterNetwork.go.html to see an example of how to use DeleteVmClusterNetwork API.
func (client DatabaseClient) DeleteVmClusterNetwork(ctx context.Context, request DeleteVmClusterNetworkRequest) (response DeleteVmClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVmClusterNetworkResponse")
	}
	return
}

// deleteVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deleteVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/DeleteVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "DeleteVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeregisterAutonomousDatabaseDataSafe Asynchronously deregisters this Autonomous Database with Data Safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DeregisterAutonomousDatabaseDataSafe.go.html to see an example of how to use DeregisterAutonomousDatabaseDataSafe API.
func (client DatabaseClient) DeregisterAutonomousDatabaseDataSafe(ctx context.Context, request DeregisterAutonomousDatabaseDataSafeRequest) (response DeregisterAutonomousDatabaseDataSafeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deregisterAutonomousDatabaseDataSafe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeregisterAutonomousDatabaseDataSafeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeregisterAutonomousDatabaseDataSafeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeregisterAutonomousDatabaseDataSafeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeregisterAutonomousDatabaseDataSafeResponse")
	}
	return
}

// deregisterAutonomousDatabaseDataSafe implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) deregisterAutonomousDatabaseDataSafe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/deregisterDataSafe", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeregisterAutonomousDatabaseDataSafeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/DeregisterAutonomousDatabaseDataSafe"
		err = common.PostProcessServiceError(err, "Database", "DeregisterAutonomousDatabaseDataSafe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAutonomousDatabaseManagement Disables Database Management for the Autonomous Database resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableAutonomousDatabaseManagement.go.html to see an example of how to use DisableAutonomousDatabaseManagement API.
func (client DatabaseClient) DisableAutonomousDatabaseManagement(ctx context.Context, request DisableAutonomousDatabaseManagementRequest) (response DisableAutonomousDatabaseManagementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableAutonomousDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAutonomousDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAutonomousDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAutonomousDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAutonomousDatabaseManagementResponse")
	}
	return
}

// disableAutonomousDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableAutonomousDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAutonomousDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/DisableAutonomousDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "DisableAutonomousDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAutonomousDatabaseOperationsInsights Disables Operations Insights for the Autonomous Database resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableAutonomousDatabaseOperationsInsights.go.html to see an example of how to use DisableAutonomousDatabaseOperationsInsights API.
func (client DatabaseClient) DisableAutonomousDatabaseOperationsInsights(ctx context.Context, request DisableAutonomousDatabaseOperationsInsightsRequest) (response DisableAutonomousDatabaseOperationsInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableAutonomousDatabaseOperationsInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAutonomousDatabaseOperationsInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAutonomousDatabaseOperationsInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAutonomousDatabaseOperationsInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAutonomousDatabaseOperationsInsightsResponse")
	}
	return
}

// disableAutonomousDatabaseOperationsInsights implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableAutonomousDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/disableOperationsInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAutonomousDatabaseOperationsInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/DisableAutonomousDatabaseOperationsInsights"
		err = common.PostProcessServiceError(err, "Database", "DisableAutonomousDatabaseOperationsInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableDatabaseManagement Disables the Database Management service for the database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableDatabaseManagement.go.html to see an example of how to use DisableDatabaseManagement API.
func (client DatabaseClient) DisableDatabaseManagement(ctx context.Context, request DisableDatabaseManagementRequest) (response DisableDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableDatabaseManagementResponse")
	}
	return
}

// disableDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/DisableDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "DisableDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalContainerDatabaseDatabaseManagement Disable Database Management service for the external container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalContainerDatabaseDatabaseManagement.go.html to see an example of how to use DisableExternalContainerDatabaseDatabaseManagement API.
func (client DatabaseClient) DisableExternalContainerDatabaseDatabaseManagement(ctx context.Context, request DisableExternalContainerDatabaseDatabaseManagementRequest) (response DisableExternalContainerDatabaseDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalContainerDatabaseDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalContainerDatabaseDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalContainerDatabaseDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalContainerDatabaseDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalContainerDatabaseDatabaseManagementResponse")
	}
	return
}

// disableExternalContainerDatabaseDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalContainerDatabaseDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/DisableExternalContainerDatabaseDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalContainerDatabaseDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalContainerDatabaseStackMonitoring Disable Stack Monitoring for the external container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalContainerDatabaseStackMonitoring.go.html to see an example of how to use DisableExternalContainerDatabaseStackMonitoring API.
func (client DatabaseClient) DisableExternalContainerDatabaseStackMonitoring(ctx context.Context, request DisableExternalContainerDatabaseStackMonitoringRequest) (response DisableExternalContainerDatabaseStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalContainerDatabaseStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalContainerDatabaseStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalContainerDatabaseStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalContainerDatabaseStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalContainerDatabaseStackMonitoringResponse")
	}
	return
}

// disableExternalContainerDatabaseStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalContainerDatabaseStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/disableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalContainerDatabaseStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/DisableExternalContainerDatabaseStackMonitoring"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalContainerDatabaseStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalNonContainerDatabaseDatabaseManagement Disable Database Management Service for the external non-container database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalNonContainerDatabaseDatabaseManagement.go.html to see an example of how to use DisableExternalNonContainerDatabaseDatabaseManagement API.
func (client DatabaseClient) DisableExternalNonContainerDatabaseDatabaseManagement(ctx context.Context, request DisableExternalNonContainerDatabaseDatabaseManagementRequest) (response DisableExternalNonContainerDatabaseDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalNonContainerDatabaseDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalNonContainerDatabaseDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalNonContainerDatabaseDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalNonContainerDatabaseDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalNonContainerDatabaseDatabaseManagementResponse")
	}
	return
}

// disableExternalNonContainerDatabaseDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalNonContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalNonContainerDatabaseDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/DisableExternalNonContainerDatabaseDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalNonContainerDatabaseDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalNonContainerDatabaseOperationsInsights Disable Operations Insights for the external non-container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalNonContainerDatabaseOperationsInsights.go.html to see an example of how to use DisableExternalNonContainerDatabaseOperationsInsights API.
func (client DatabaseClient) DisableExternalNonContainerDatabaseOperationsInsights(ctx context.Context, request DisableExternalNonContainerDatabaseOperationsInsightsRequest) (response DisableExternalNonContainerDatabaseOperationsInsightsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalNonContainerDatabaseOperationsInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalNonContainerDatabaseOperationsInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalNonContainerDatabaseOperationsInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalNonContainerDatabaseOperationsInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalNonContainerDatabaseOperationsInsightsResponse")
	}
	return
}

// disableExternalNonContainerDatabaseOperationsInsights implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalNonContainerDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/disableOperationsInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalNonContainerDatabaseOperationsInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/DisableExternalNonContainerDatabaseOperationsInsights"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalNonContainerDatabaseOperationsInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalNonContainerDatabaseStackMonitoring Disable Stack Monitoring for the external non-container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalNonContainerDatabaseStackMonitoring.go.html to see an example of how to use DisableExternalNonContainerDatabaseStackMonitoring API.
func (client DatabaseClient) DisableExternalNonContainerDatabaseStackMonitoring(ctx context.Context, request DisableExternalNonContainerDatabaseStackMonitoringRequest) (response DisableExternalNonContainerDatabaseStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalNonContainerDatabaseStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalNonContainerDatabaseStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalNonContainerDatabaseStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalNonContainerDatabaseStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalNonContainerDatabaseStackMonitoringResponse")
	}
	return
}

// disableExternalNonContainerDatabaseStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalNonContainerDatabaseStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/disableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalNonContainerDatabaseStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/DisableExternalNonContainerDatabaseStackMonitoring"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalNonContainerDatabaseStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalPluggableDatabaseDatabaseManagement Disable Database Management Service for the external pluggable database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalPluggableDatabaseDatabaseManagement.go.html to see an example of how to use DisableExternalPluggableDatabaseDatabaseManagement API.
func (client DatabaseClient) DisableExternalPluggableDatabaseDatabaseManagement(ctx context.Context, request DisableExternalPluggableDatabaseDatabaseManagementRequest) (response DisableExternalPluggableDatabaseDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalPluggableDatabaseDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalPluggableDatabaseDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalPluggableDatabaseDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalPluggableDatabaseDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalPluggableDatabaseDatabaseManagementResponse")
	}
	return
}

// disableExternalPluggableDatabaseDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalPluggableDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalPluggableDatabaseDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/DisableExternalPluggableDatabaseDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalPluggableDatabaseDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalPluggableDatabaseOperationsInsights Disable Operations Insights for the external pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalPluggableDatabaseOperationsInsights.go.html to see an example of how to use DisableExternalPluggableDatabaseOperationsInsights API.
func (client DatabaseClient) DisableExternalPluggableDatabaseOperationsInsights(ctx context.Context, request DisableExternalPluggableDatabaseOperationsInsightsRequest) (response DisableExternalPluggableDatabaseOperationsInsightsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalPluggableDatabaseOperationsInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalPluggableDatabaseOperationsInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalPluggableDatabaseOperationsInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalPluggableDatabaseOperationsInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalPluggableDatabaseOperationsInsightsResponse")
	}
	return
}

// disableExternalPluggableDatabaseOperationsInsights implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalPluggableDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/disableOperationsInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalPluggableDatabaseOperationsInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/DisableExternalPluggableDatabaseOperationsInsights"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalPluggableDatabaseOperationsInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalPluggableDatabaseStackMonitoring Disable Stack Monitoring for the external pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableExternalPluggableDatabaseStackMonitoring.go.html to see an example of how to use DisableExternalPluggableDatabaseStackMonitoring API.
func (client DatabaseClient) DisableExternalPluggableDatabaseStackMonitoring(ctx context.Context, request DisableExternalPluggableDatabaseStackMonitoringRequest) (response DisableExternalPluggableDatabaseStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalPluggableDatabaseStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalPluggableDatabaseStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalPluggableDatabaseStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalPluggableDatabaseStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalPluggableDatabaseStackMonitoringResponse")
	}
	return
}

// disableExternalPluggableDatabaseStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disableExternalPluggableDatabaseStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/disableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalPluggableDatabaseStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/DisableExternalPluggableDatabaseStackMonitoring"
		err = common.PostProcessServiceError(err, "Database", "DisableExternalPluggableDatabaseStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisablePluggableDatabaseManagement Disables the Database Management service for the pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisablePluggableDatabaseManagement.go.html to see an example of how to use DisablePluggableDatabaseManagement API.
func (client DatabaseClient) DisablePluggableDatabaseManagement(ctx context.Context, request DisablePluggableDatabaseManagementRequest) (response DisablePluggableDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disablePluggableDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisablePluggableDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisablePluggableDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisablePluggableDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisablePluggableDatabaseManagementResponse")
	}
	return
}

// disablePluggableDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) disablePluggableDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/disablePluggableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisablePluggableDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/DisablePluggableDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "DisablePluggableDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadExadataInfrastructureConfigFile Downloads the configuration file for the specified Exadata Cloud@Customer infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DownloadExadataInfrastructureConfigFile.go.html to see an example of how to use DownloadExadataInfrastructureConfigFile API.
func (client DatabaseClient) DownloadExadataInfrastructureConfigFile(ctx context.Context, request DownloadExadataInfrastructureConfigFileRequest) (response DownloadExadataInfrastructureConfigFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadExadataInfrastructureConfigFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadExadataInfrastructureConfigFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadExadataInfrastructureConfigFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadExadataInfrastructureConfigFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadExadataInfrastructureConfigFileResponse")
	}
	return
}

// downloadExadataInfrastructureConfigFile implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) downloadExadataInfrastructureConfigFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/downloadConfigFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadExadataInfrastructureConfigFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/DownloadExadataInfrastructureConfigFile"
		err = common.PostProcessServiceError(err, "Database", "DownloadExadataInfrastructureConfigFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadOneoffPatch Download one-off patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DownloadOneoffPatch.go.html to see an example of how to use DownloadOneoffPatch API.
func (client DatabaseClient) DownloadOneoffPatch(ctx context.Context, request DownloadOneoffPatchRequest) (response DownloadOneoffPatchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadOneoffPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadOneoffPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadOneoffPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadOneoffPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadOneoffPatchResponse")
	}
	return
}

// downloadOneoffPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) downloadOneoffPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oneoffPatches/{oneoffPatchId}/actions/downloadOneoffPatch", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadOneoffPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/DownloadOneoffPatch"
		err = common.PostProcessServiceError(err, "Database", "DownloadOneoffPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadValidationReport Downloads the network validation report file for the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DownloadValidationReport.go.html to see an example of how to use DownloadValidationReport API.
func (client DatabaseClient) DownloadValidationReport(ctx context.Context, request DownloadValidationReportRequest) (response DownloadValidationReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadValidationReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadValidationReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadValidationReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadValidationReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadValidationReportResponse")
	}
	return
}

// downloadValidationReport implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) downloadValidationReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}/actions/downloadValidationReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadValidationReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/DownloadValidationReport"
		err = common.PostProcessServiceError(err, "Database", "DownloadValidationReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadVmClusterNetworkConfigFile Downloads the configuration file for the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DownloadVmClusterNetworkConfigFile.go.html to see an example of how to use DownloadVmClusterNetworkConfigFile API.
func (client DatabaseClient) DownloadVmClusterNetworkConfigFile(ctx context.Context, request DownloadVmClusterNetworkConfigFileRequest) (response DownloadVmClusterNetworkConfigFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadVmClusterNetworkConfigFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadVmClusterNetworkConfigFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadVmClusterNetworkConfigFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadVmClusterNetworkConfigFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadVmClusterNetworkConfigFileResponse")
	}
	return
}

// downloadVmClusterNetworkConfigFile implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) downloadVmClusterNetworkConfigFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}/actions/downloadConfigFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadVmClusterNetworkConfigFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/DownloadVmClusterNetworkConfigFile"
		err = common.PostProcessServiceError(err, "Database", "DownloadVmClusterNetworkConfigFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAutonomousDatabaseManagement Enables Database Management for Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableAutonomousDatabaseManagement.go.html to see an example of how to use EnableAutonomousDatabaseManagement API.
func (client DatabaseClient) EnableAutonomousDatabaseManagement(ctx context.Context, request EnableAutonomousDatabaseManagementRequest) (response EnableAutonomousDatabaseManagementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableAutonomousDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAutonomousDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAutonomousDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAutonomousDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAutonomousDatabaseManagementResponse")
	}
	return
}

// enableAutonomousDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableAutonomousDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAutonomousDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/EnableAutonomousDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "EnableAutonomousDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAutonomousDatabaseOperationsInsights Enables the specified Autonomous Database with Operations Insights.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableAutonomousDatabaseOperationsInsights.go.html to see an example of how to use EnableAutonomousDatabaseOperationsInsights API.
func (client DatabaseClient) EnableAutonomousDatabaseOperationsInsights(ctx context.Context, request EnableAutonomousDatabaseOperationsInsightsRequest) (response EnableAutonomousDatabaseOperationsInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableAutonomousDatabaseOperationsInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAutonomousDatabaseOperationsInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAutonomousDatabaseOperationsInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAutonomousDatabaseOperationsInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAutonomousDatabaseOperationsInsightsResponse")
	}
	return
}

// enableAutonomousDatabaseOperationsInsights implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableAutonomousDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/enableOperationsInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAutonomousDatabaseOperationsInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/EnableAutonomousDatabaseOperationsInsights"
		err = common.PostProcessServiceError(err, "Database", "EnableAutonomousDatabaseOperationsInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDatabaseManagement Enables the Database Management service for an Oracle Database located in Oracle Cloud Infrastructure. This service allows the database to access tools including Metrics and Performance hub. Database Management is enabled at the container database (CDB) level.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableDatabaseManagement.go.html to see an example of how to use EnableDatabaseManagement API.
func (client DatabaseClient) EnableDatabaseManagement(ctx context.Context, request EnableDatabaseManagementRequest) (response EnableDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableDatabaseManagementResponse")
	}
	return
}

// enableDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/EnableDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "EnableDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalContainerDatabaseDatabaseManagement Enables Database Management Service for the external container database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalContainerDatabaseDatabaseManagement.go.html to see an example of how to use EnableExternalContainerDatabaseDatabaseManagement API.
func (client DatabaseClient) EnableExternalContainerDatabaseDatabaseManagement(ctx context.Context, request EnableExternalContainerDatabaseDatabaseManagementRequest) (response EnableExternalContainerDatabaseDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalContainerDatabaseDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalContainerDatabaseDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalContainerDatabaseDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalContainerDatabaseDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalContainerDatabaseDatabaseManagementResponse")
	}
	return
}

// enableExternalContainerDatabaseDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalContainerDatabaseDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/EnableExternalContainerDatabaseDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalContainerDatabaseDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalContainerDatabaseStackMonitoring Enable Stack Monitoring for the external container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalContainerDatabaseStackMonitoring.go.html to see an example of how to use EnableExternalContainerDatabaseStackMonitoring API.
func (client DatabaseClient) EnableExternalContainerDatabaseStackMonitoring(ctx context.Context, request EnableExternalContainerDatabaseStackMonitoringRequest) (response EnableExternalContainerDatabaseStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalContainerDatabaseStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalContainerDatabaseStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalContainerDatabaseStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalContainerDatabaseStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalContainerDatabaseStackMonitoringResponse")
	}
	return
}

// enableExternalContainerDatabaseStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalContainerDatabaseStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/enableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalContainerDatabaseStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/EnableExternalContainerDatabaseStackMonitoring"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalContainerDatabaseStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalNonContainerDatabaseDatabaseManagement Enable Database Management Service for the external non-container database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalNonContainerDatabaseDatabaseManagement.go.html to see an example of how to use EnableExternalNonContainerDatabaseDatabaseManagement API.
func (client DatabaseClient) EnableExternalNonContainerDatabaseDatabaseManagement(ctx context.Context, request EnableExternalNonContainerDatabaseDatabaseManagementRequest) (response EnableExternalNonContainerDatabaseDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalNonContainerDatabaseDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalNonContainerDatabaseDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalNonContainerDatabaseDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalNonContainerDatabaseDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalNonContainerDatabaseDatabaseManagementResponse")
	}
	return
}

// enableExternalNonContainerDatabaseDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalNonContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalNonContainerDatabaseDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/EnableExternalNonContainerDatabaseDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalNonContainerDatabaseDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalNonContainerDatabaseOperationsInsights Enable Operations Insights for the external non-container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalNonContainerDatabaseOperationsInsights.go.html to see an example of how to use EnableExternalNonContainerDatabaseOperationsInsights API.
func (client DatabaseClient) EnableExternalNonContainerDatabaseOperationsInsights(ctx context.Context, request EnableExternalNonContainerDatabaseOperationsInsightsRequest) (response EnableExternalNonContainerDatabaseOperationsInsightsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalNonContainerDatabaseOperationsInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalNonContainerDatabaseOperationsInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalNonContainerDatabaseOperationsInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalNonContainerDatabaseOperationsInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalNonContainerDatabaseOperationsInsightsResponse")
	}
	return
}

// enableExternalNonContainerDatabaseOperationsInsights implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalNonContainerDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/enableOperationsInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalNonContainerDatabaseOperationsInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/EnableExternalNonContainerDatabaseOperationsInsights"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalNonContainerDatabaseOperationsInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalNonContainerDatabaseStackMonitoring Enable Stack Monitoring for the external non-container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalNonContainerDatabaseStackMonitoring.go.html to see an example of how to use EnableExternalNonContainerDatabaseStackMonitoring API.
func (client DatabaseClient) EnableExternalNonContainerDatabaseStackMonitoring(ctx context.Context, request EnableExternalNonContainerDatabaseStackMonitoringRequest) (response EnableExternalNonContainerDatabaseStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalNonContainerDatabaseStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalNonContainerDatabaseStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalNonContainerDatabaseStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalNonContainerDatabaseStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalNonContainerDatabaseStackMonitoringResponse")
	}
	return
}

// enableExternalNonContainerDatabaseStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalNonContainerDatabaseStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/enableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalNonContainerDatabaseStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/EnableExternalNonContainerDatabaseStackMonitoring"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalNonContainerDatabaseStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalPluggableDatabaseDatabaseManagement Enable Database Management Service for the external pluggable database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalPluggableDatabaseDatabaseManagement.go.html to see an example of how to use EnableExternalPluggableDatabaseDatabaseManagement API.
func (client DatabaseClient) EnableExternalPluggableDatabaseDatabaseManagement(ctx context.Context, request EnableExternalPluggableDatabaseDatabaseManagementRequest) (response EnableExternalPluggableDatabaseDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalPluggableDatabaseDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalPluggableDatabaseDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalPluggableDatabaseDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalPluggableDatabaseDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalPluggableDatabaseDatabaseManagementResponse")
	}
	return
}

// enableExternalPluggableDatabaseDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalPluggableDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalPluggableDatabaseDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/EnableExternalPluggableDatabaseDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalPluggableDatabaseDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalPluggableDatabaseOperationsInsights Enable Operations Insights for the external pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalPluggableDatabaseOperationsInsights.go.html to see an example of how to use EnableExternalPluggableDatabaseOperationsInsights API.
func (client DatabaseClient) EnableExternalPluggableDatabaseOperationsInsights(ctx context.Context, request EnableExternalPluggableDatabaseOperationsInsightsRequest) (response EnableExternalPluggableDatabaseOperationsInsightsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalPluggableDatabaseOperationsInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalPluggableDatabaseOperationsInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalPluggableDatabaseOperationsInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalPluggableDatabaseOperationsInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalPluggableDatabaseOperationsInsightsResponse")
	}
	return
}

// enableExternalPluggableDatabaseOperationsInsights implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalPluggableDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/enableOperationsInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalPluggableDatabaseOperationsInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/EnableExternalPluggableDatabaseOperationsInsights"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalPluggableDatabaseOperationsInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalPluggableDatabaseStackMonitoring Enable Stack Monitoring for the external pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnableExternalPluggableDatabaseStackMonitoring.go.html to see an example of how to use EnableExternalPluggableDatabaseStackMonitoring API.
func (client DatabaseClient) EnableExternalPluggableDatabaseStackMonitoring(ctx context.Context, request EnableExternalPluggableDatabaseStackMonitoringRequest) (response EnableExternalPluggableDatabaseStackMonitoringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalPluggableDatabaseStackMonitoring, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalPluggableDatabaseStackMonitoringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalPluggableDatabaseStackMonitoringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalPluggableDatabaseStackMonitoringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalPluggableDatabaseStackMonitoringResponse")
	}
	return
}

// enableExternalPluggableDatabaseStackMonitoring implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enableExternalPluggableDatabaseStackMonitoring(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/enableStackMonitoring", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalPluggableDatabaseStackMonitoringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/EnableExternalPluggableDatabaseStackMonitoring"
		err = common.PostProcessServiceError(err, "Database", "EnableExternalPluggableDatabaseStackMonitoring", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnablePluggableDatabaseManagement Enables the Database Management service for an Oracle Pluggable Database located in Oracle Cloud Infrastructure. This service allows the pluggable database to access tools including Metrics and Performance hub. Database Management is enabled at the pluggable database (PDB) level.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/EnablePluggableDatabaseManagement.go.html to see an example of how to use EnablePluggableDatabaseManagement API.
func (client DatabaseClient) EnablePluggableDatabaseManagement(ctx context.Context, request EnablePluggableDatabaseManagementRequest) (response EnablePluggableDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enablePluggableDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnablePluggableDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnablePluggableDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnablePluggableDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnablePluggableDatabaseManagementResponse")
	}
	return
}

// enablePluggableDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) enablePluggableDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/enablePluggableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnablePluggableDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/EnablePluggableDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "EnablePluggableDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FailOverAutonomousDatabase Initiates a failover of the specified Autonomous Database to the associated peer database. Applicable only to databases with Disaster Recovery enabled.
// This API should be called in the remote region where the peer database resides.
// Below parameter is optional:
//   - `peerDbId`
//     Use this parameter to specify the database OCID of the Disaster Recovery peer, which is located in a different (remote) region from the current peer database.
//     If this parameter is not provided, the failover will happen in the same region.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/FailOverAutonomousDatabase.go.html to see an example of how to use FailOverAutonomousDatabase API.
func (client DatabaseClient) FailOverAutonomousDatabase(ctx context.Context, request FailOverAutonomousDatabaseRequest) (response FailOverAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.failOverAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FailOverAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FailOverAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FailOverAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FailOverAutonomousDatabaseResponse")
	}
	return
}

// failOverAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) failOverAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/failover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FailOverAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/FailOverAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "FailOverAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FailoverAutonomousContainerDatabaseDataguardAssociation Fails over the standby Autonomous Container Database identified by the autonomousContainerDatabaseId parameter to the primary Autonomous Container Database after the existing primary Autonomous Container Database fails or becomes unreachable.
// A failover can result in data loss, depending on the protection mode in effect at the time the primary Autonomous Container Database fails.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/FailoverAutonomousContainerDatabaseDataguardAssociation.go.html to see an example of how to use FailoverAutonomousContainerDatabaseDataguardAssociation API.
func (client DatabaseClient) FailoverAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request FailoverAutonomousContainerDatabaseDataguardAssociationRequest) (response FailoverAutonomousContainerDatabaseDataguardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.failoverAutonomousContainerDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FailoverAutonomousContainerDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FailoverAutonomousContainerDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FailoverAutonomousContainerDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FailoverAutonomousContainerDatabaseDataguardAssociationResponse")
	}
	return
}

// failoverAutonomousContainerDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) failoverAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}/actions/failover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FailoverAutonomousContainerDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/FailoverAutonomousContainerDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "FailoverAutonomousContainerDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FailoverDataGuardAssociation Performs a failover to transition the standby database identified by the `databaseId` parameter into the
// specified Data Guard association's primary role after the existing primary database fails or becomes unreachable.
// A failover might result in data loss depending on the protection mode in effect at the time of the primary
// database failure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/FailoverDataGuardAssociation.go.html to see an example of how to use FailoverDataGuardAssociation API.
func (client DatabaseClient) FailoverDataGuardAssociation(ctx context.Context, request FailoverDataGuardAssociationRequest) (response FailoverDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.failoverDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FailoverDataGuardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FailoverDataGuardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FailoverDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FailoverDataGuardAssociationResponse")
	}
	return
}

// failoverDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) failoverDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/failover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FailoverDataGuardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/FailoverDataGuardAssociation"
		err = common.PostProcessServiceError(err, "Database", "FailoverDataGuardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateAutonomousDatabaseWallet Creates and downloads a wallet for the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GenerateAutonomousDatabaseWallet.go.html to see an example of how to use GenerateAutonomousDatabaseWallet API.
func (client DatabaseClient) GenerateAutonomousDatabaseWallet(ctx context.Context, request GenerateAutonomousDatabaseWalletRequest) (response GenerateAutonomousDatabaseWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateAutonomousDatabaseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateAutonomousDatabaseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateAutonomousDatabaseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateAutonomousDatabaseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateAutonomousDatabaseWalletResponse")
	}
	return
}

// generateAutonomousDatabaseWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) generateAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/generateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateAutonomousDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/GenerateAutonomousDatabaseWallet"
		err = common.PostProcessServiceError(err, "Database", "GenerateAutonomousDatabaseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateRecommendedVmClusterNetwork Generates a recommended Cloud@Customer VM cluster network configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GenerateRecommendedVmClusterNetwork.go.html to see an example of how to use GenerateRecommendedVmClusterNetwork API.
func (client DatabaseClient) GenerateRecommendedVmClusterNetwork(ctx context.Context, request GenerateRecommendedVmClusterNetworkRequest) (response GenerateRecommendedVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateRecommendedVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateRecommendedVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateRecommendedVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateRecommendedVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateRecommendedVmClusterNetworkResponse")
	}
	return
}

// generateRecommendedVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) generateRecommendedVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/actions/generateRecommendedNetwork", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateRecommendedVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/GenerateRecommendedVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "GenerateRecommendedVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApplicationVip Gets information about a specified application virtual IP (VIP) address.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetApplicationVip.go.html to see an example of how to use GetApplicationVip API.
func (client DatabaseClient) GetApplicationVip(ctx context.Context, request GetApplicationVipRequest) (response GetApplicationVipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplicationVip, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApplicationVipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApplicationVipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplicationVipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplicationVipResponse")
	}
	return
}

// getApplicationVip implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getApplicationVip(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applicationVip/{applicationVipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApplicationVipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ApplicationVip/GetApplicationVip"
		err = common.PostProcessServiceError(err, "Database", "GetApplicationVip", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousContainerDatabase Gets information about the specified Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousContainerDatabase.go.html to see an example of how to use GetAutonomousContainerDatabase API.
func (client DatabaseClient) GetAutonomousContainerDatabase(ctx context.Context, request GetAutonomousContainerDatabaseRequest) (response GetAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousContainerDatabaseResponse")
	}
	return
}

// getAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/GetAutonomousContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousContainerDatabaseDataguardAssociation Gets an Autonomous Container Database enabled with Autonomous Data Guard associated with the specified Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousContainerDatabaseDataguardAssociation.go.html to see an example of how to use GetAutonomousContainerDatabaseDataguardAssociation API.
func (client DatabaseClient) GetAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request GetAutonomousContainerDatabaseDataguardAssociationRequest) (response GetAutonomousContainerDatabaseDataguardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousContainerDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousContainerDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousContainerDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousContainerDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousContainerDatabaseDataguardAssociationResponse")
	}
	return
}

// getAutonomousContainerDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/GetAutonomousContainerDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousContainerDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousContainerDatabaseResourceUsage Get resource usage details for the specified Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousContainerDatabaseResourceUsage.go.html to see an example of how to use GetAutonomousContainerDatabaseResourceUsage API.
func (client DatabaseClient) GetAutonomousContainerDatabaseResourceUsage(ctx context.Context, request GetAutonomousContainerDatabaseResourceUsageRequest) (response GetAutonomousContainerDatabaseResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousContainerDatabaseResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousContainerDatabaseResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousContainerDatabaseResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousContainerDatabaseResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousContainerDatabaseResourceUsageResponse")
	}
	return
}

// getAutonomousContainerDatabaseResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousContainerDatabaseResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/resourceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseResourceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/GetAutonomousContainerDatabaseResourceUsage"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousContainerDatabaseResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousDatabase Gets the details of the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousDatabase.go.html to see an example of how to use GetAutonomousDatabase API.
func (client DatabaseClient) GetAutonomousDatabase(ctx context.Context, request GetAutonomousDatabaseRequest) (response GetAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseResponse")
	}
	return
}

// getAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/GetAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousDatabaseBackup Gets information about the specified Autonomous Database backup.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousDatabaseBackup.go.html to see an example of how to use GetAutonomousDatabaseBackup API.
func (client DatabaseClient) GetAutonomousDatabaseBackup(ctx context.Context, request GetAutonomousDatabaseBackupRequest) (response GetAutonomousDatabaseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousDatabaseBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousDatabaseBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseBackupResponse")
	}
	return
}

// getAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseBackups/{autonomousDatabaseBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseBackup/GetAutonomousDatabaseBackup"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousDatabaseBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousDatabaseDataguardAssociation Gets an Autonomous Data Guard-enabled database associated with the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousDatabaseDataguardAssociation.go.html to see an example of how to use GetAutonomousDatabaseDataguardAssociation API.
func (client DatabaseClient) GetAutonomousDatabaseDataguardAssociation(ctx context.Context, request GetAutonomousDatabaseDataguardAssociationRequest) (response GetAutonomousDatabaseDataguardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseDataguardAssociationResponse")
	}
	return
}

// getAutonomousDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/autonomousDatabaseDataguardAssociations/{autonomousDatabaseDataguardAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseDataguardAssociation/GetAutonomousDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousDatabaseRegionalWallet Gets the Autonomous Database regional wallet details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousDatabaseRegionalWallet.go.html to see an example of how to use GetAutonomousDatabaseRegionalWallet API.
func (client DatabaseClient) GetAutonomousDatabaseRegionalWallet(ctx context.Context, request GetAutonomousDatabaseRegionalWalletRequest) (response GetAutonomousDatabaseRegionalWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabaseRegionalWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousDatabaseRegionalWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousDatabaseRegionalWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseRegionalWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseRegionalWalletResponse")
	}
	return
}

// getAutonomousDatabaseRegionalWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabaseRegionalWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/wallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseRegionalWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseWallet/GetAutonomousDatabaseRegionalWallet"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousDatabaseRegionalWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousDatabaseWallet Gets the wallet details for the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousDatabaseWallet.go.html to see an example of how to use GetAutonomousDatabaseWallet API.
func (client DatabaseClient) GetAutonomousDatabaseWallet(ctx context.Context, request GetAutonomousDatabaseWalletRequest) (response GetAutonomousDatabaseWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousDatabaseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousDatabaseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousDatabaseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousDatabaseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousDatabaseWalletResponse")
	}
	return
}

// getAutonomousDatabaseWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/wallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseWallet/GetAutonomousDatabaseWallet"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousDatabaseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousExadataInfrastructure **Deprecated.** Use the GetCloudExadataInfrastructure operation to get details of an Exadata Infrastructure resource and the GetCloudAutonomousVmCluster operation to get details of an Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousExadataInfrastructure.go.html to see an example of how to use GetAutonomousExadataInfrastructure API.
func (client DatabaseClient) GetAutonomousExadataInfrastructure(ctx context.Context, request GetAutonomousExadataInfrastructureRequest) (response GetAutonomousExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousExadataInfrastructureResponse")
	}
	return
}

// getAutonomousExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/GetAutonomousExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousPatch Gets information about a specific autonomous patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousPatch.go.html to see an example of how to use GetAutonomousPatch API.
func (client DatabaseClient) GetAutonomousPatch(ctx context.Context, request GetAutonomousPatchRequest) (response GetAutonomousPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousPatchResponse")
	}
	return
}

// getAutonomousPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousPatches/{autonomousPatchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousPatch/GetAutonomousPatch"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousVirtualMachine Gets the details of specific Autonomous Virtual Machine.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousVirtualMachine.go.html to see an example of how to use GetAutonomousVirtualMachine API.
func (client DatabaseClient) GetAutonomousVirtualMachine(ctx context.Context, request GetAutonomousVirtualMachineRequest) (response GetAutonomousVirtualMachineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousVirtualMachine, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousVirtualMachineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousVirtualMachineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousVirtualMachineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousVirtualMachineResponse")
	}
	return
}

// getAutonomousVirtualMachine implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousVirtualMachine(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVirtualMachines/{autonomousVirtualMachineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousVirtualMachineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVirtualMachine/GetAutonomousVirtualMachine"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousVirtualMachine", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousVmCluster Gets information about the specified Autonomous VM cluster for an Exadata Cloud@Customer system. To get information about an Autonomous VM Cluster in the Oracle cloud, see GetCloudAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousVmCluster.go.html to see an example of how to use GetAutonomousVmCluster API.
func (client DatabaseClient) GetAutonomousVmCluster(ctx context.Context, request GetAutonomousVmClusterRequest) (response GetAutonomousVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousVmClusterResponse")
	}
	return
}

// getAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVmClusters/{autonomousVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/GetAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutonomousVmClusterResourceUsage Get the resource usage details for the specified Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousVmClusterResourceUsage.go.html to see an example of how to use GetAutonomousVmClusterResourceUsage API.
func (client DatabaseClient) GetAutonomousVmClusterResourceUsage(ctx context.Context, request GetAutonomousVmClusterResourceUsageRequest) (response GetAutonomousVmClusterResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutonomousVmClusterResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutonomousVmClusterResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutonomousVmClusterResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutonomousVmClusterResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutonomousVmClusterResourceUsageResponse")
	}
	return
}

// getAutonomousVmClusterResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getAutonomousVmClusterResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVmClusters/{autonomousVmClusterId}/resourceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousVmClusterResourceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/GetAutonomousVmClusterResourceUsage"
		err = common.PostProcessServiceError(err, "Database", "GetAutonomousVmClusterResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackup Gets information about the specified backup.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetBackup.go.html to see an example of how to use GetBackup API.
func (client DatabaseClient) GetBackup(ctx context.Context, request GetBackupRequest) (response GetBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackupResponse")
	}
	return
}

// getBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Backup/GetBackup"
		err = common.PostProcessServiceError(err, "Database", "GetBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackupDestination Gets information about the specified backup destination in an Exadata Cloud@Customer system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetBackupDestination.go.html to see an example of how to use GetBackupDestination API.
func (client DatabaseClient) GetBackupDestination(ctx context.Context, request GetBackupDestinationRequest) (response GetBackupDestinationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackupDestination, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackupDestinationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackupDestinationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackupDestinationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackupDestinationResponse")
	}
	return
}

// getBackupDestination implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backupDestinations/{backupDestinationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackupDestinationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/BackupDestination/GetBackupDestination"
		err = common.PostProcessServiceError(err, "Database", "GetBackupDestination", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudAutonomousVmCluster Gets information about the specified Autonomous Exadata VM cluster in the Oracle cloud. For Exadata Cloud@Custustomer systems, see GetAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudAutonomousVmCluster.go.html to see an example of how to use GetCloudAutonomousVmCluster API.
func (client DatabaseClient) GetCloudAutonomousVmCluster(ctx context.Context, request GetCloudAutonomousVmClusterRequest) (response GetCloudAutonomousVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudAutonomousVmClusterResponse")
	}
	return
}

// getCloudAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/GetCloudAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "GetCloudAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudAutonomousVmClusterResourceUsage Get the resource usage details for the specified Cloud Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudAutonomousVmClusterResourceUsage.go.html to see an example of how to use GetCloudAutonomousVmClusterResourceUsage API.
func (client DatabaseClient) GetCloudAutonomousVmClusterResourceUsage(ctx context.Context, request GetCloudAutonomousVmClusterResourceUsageRequest) (response GetCloudAutonomousVmClusterResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudAutonomousVmClusterResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudAutonomousVmClusterResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudAutonomousVmClusterResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudAutonomousVmClusterResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudAutonomousVmClusterResourceUsageResponse")
	}
	return
}

// getCloudAutonomousVmClusterResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudAutonomousVmClusterResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}/resourceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudAutonomousVmClusterResourceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/GetCloudAutonomousVmClusterResourceUsage"
		err = common.PostProcessServiceError(err, "Database", "GetCloudAutonomousVmClusterResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudExadataInfrastructure Gets information about the specified cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudExadataInfrastructure.go.html to see an example of how to use GetCloudExadataInfrastructure API.
func (client DatabaseClient) GetCloudExadataInfrastructure(ctx context.Context, request GetCloudExadataInfrastructureRequest) (response GetCloudExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudExadataInfrastructureResponse")
	}
	return
}

// getCloudExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/GetCloudExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "GetCloudExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudExadataInfrastructureUnallocatedResources Gets unallocated resources information for the specified Cloud Exadata infrastructure.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudExadataInfrastructureUnallocatedResources.go.html to see an example of how to use GetCloudExadataInfrastructureUnallocatedResources API.
func (client DatabaseClient) GetCloudExadataInfrastructureUnallocatedResources(ctx context.Context, request GetCloudExadataInfrastructureUnallocatedResourcesRequest) (response GetCloudExadataInfrastructureUnallocatedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudExadataInfrastructureUnallocatedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudExadataInfrastructureUnallocatedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudExadataInfrastructureUnallocatedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudExadataInfrastructureUnallocatedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudExadataInfrastructureUnallocatedResourcesResponse")
	}
	return
}

// getCloudExadataInfrastructureUnallocatedResources implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudExadataInfrastructureUnallocatedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}/unAllocatedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudExadataInfrastructureUnallocatedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructureUnallocatedResources/GetCloudExadataInfrastructureUnallocatedResources"
		err = common.PostProcessServiceError(err, "Database", "GetCloudExadataInfrastructureUnallocatedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudVmCluster Gets information about the specified cloud VM cluster. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudVmCluster.go.html to see an example of how to use GetCloudVmCluster API.
func (client DatabaseClient) GetCloudVmCluster(ctx context.Context, request GetCloudVmClusterRequest) (response GetCloudVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudVmClusterResponse")
	}
	return
}

// getCloudVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/GetCloudVmCluster"
		err = common.PostProcessServiceError(err, "Database", "GetCloudVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudVmClusterIormConfig Gets the IORM configuration for the specified cloud VM cluster in an Exadata Cloud Service instance.
// If you have not specified an IORM configuration, the default configuration is returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudVmClusterIormConfig.go.html to see an example of how to use GetCloudVmClusterIormConfig API.
func (client DatabaseClient) GetCloudVmClusterIormConfig(ctx context.Context, request GetCloudVmClusterIormConfigRequest) (response GetCloudVmClusterIormConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudVmClusterIormConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudVmClusterIormConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudVmClusterIormConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudVmClusterIormConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudVmClusterIormConfigResponse")
	}
	return
}

// getCloudVmClusterIormConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudVmClusterIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/CloudVmClusterIormConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterIormConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/GetCloudVmClusterIormConfig"
		err = common.PostProcessServiceError(err, "Database", "GetCloudVmClusterIormConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudVmClusterUpdate Gets information about a specified maintenance update package for a cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudVmClusterUpdate.go.html to see an example of how to use GetCloudVmClusterUpdate API.
func (client DatabaseClient) GetCloudVmClusterUpdate(ctx context.Context, request GetCloudVmClusterUpdateRequest) (response GetCloudVmClusterUpdateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudVmClusterUpdate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudVmClusterUpdateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudVmClusterUpdateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudVmClusterUpdateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudVmClusterUpdateResponse")
	}
	return
}

// getCloudVmClusterUpdate implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudVmClusterUpdate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updates/{updateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterUpdateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Update/GetCloudVmClusterUpdate"
		err = common.PostProcessServiceError(err, "Database", "GetCloudVmClusterUpdate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCloudVmClusterUpdateHistoryEntry Gets the maintenance update history details for the specified update history entry. Applies to Exadata Cloud Service instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetCloudVmClusterUpdateHistoryEntry.go.html to see an example of how to use GetCloudVmClusterUpdateHistoryEntry API.
func (client DatabaseClient) GetCloudVmClusterUpdateHistoryEntry(ctx context.Context, request GetCloudVmClusterUpdateHistoryEntryRequest) (response GetCloudVmClusterUpdateHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCloudVmClusterUpdateHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCloudVmClusterUpdateHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCloudVmClusterUpdateHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCloudVmClusterUpdateHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCloudVmClusterUpdateHistoryEntryResponse")
	}
	return
}

// getCloudVmClusterUpdateHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getCloudVmClusterUpdateHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updateHistoryEntries/{updateHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterUpdateHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/UpdateHistoryEntry/GetCloudVmClusterUpdateHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetCloudVmClusterUpdateHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConsoleConnection Gets the specified database node console connection's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetConsoleConnection.go.html to see an example of how to use GetConsoleConnection API.
func (client DatabaseClient) GetConsoleConnection(ctx context.Context, request GetConsoleConnectionRequest) (response GetConsoleConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConsoleConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConsoleConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConsoleConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConsoleConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConsoleConnectionResponse")
	}
	return
}

// getConsoleConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleConnections/{consoleConnectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConsoleConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleConnection/GetConsoleConnection"
		err = common.PostProcessServiceError(err, "Database", "GetConsoleConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConsoleHistory Gets information about the specified database node console history.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetConsoleHistory.go.html to see an example of how to use GetConsoleHistory API.
func (client DatabaseClient) GetConsoleHistory(ctx context.Context, request GetConsoleHistoryRequest) (response GetConsoleHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConsoleHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConsoleHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConsoleHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConsoleHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConsoleHistoryResponse")
	}
	return
}

// getConsoleHistory implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getConsoleHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleHistories/{consoleHistoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConsoleHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleHistory/GetConsoleHistory"
		err = common.PostProcessServiceError(err, "Database", "GetConsoleHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConsoleHistoryContent Retrieves the specified database node console history contents upto a megabyte.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetConsoleHistoryContent.go.html to see an example of how to use GetConsoleHistoryContent API.
func (client DatabaseClient) GetConsoleHistoryContent(ctx context.Context, request GetConsoleHistoryContentRequest) (response GetConsoleHistoryContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConsoleHistoryContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConsoleHistoryContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConsoleHistoryContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConsoleHistoryContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConsoleHistoryContentResponse")
	}
	return
}

// getConsoleHistoryContent implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getConsoleHistoryContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleHistories/{consoleHistoryId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConsoleHistoryContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleHistory/GetConsoleHistoryContent"
		err = common.PostProcessServiceError(err, "Database", "GetConsoleHistoryContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataGuardAssociation Gets the specified Data Guard association's configuration information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDataGuardAssociation.go.html to see an example of how to use GetDataGuardAssociation API.
func (client DatabaseClient) GetDataGuardAssociation(ctx context.Context, request GetDataGuardAssociationRequest) (response GetDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataGuardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataGuardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataGuardAssociationResponse")
	}
	return
}

// getDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataGuardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/GetDataGuardAssociation"
		err = common.PostProcessServiceError(err, "Database", "GetDataGuardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabase Gets information about the specified database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDatabase.go.html to see an example of how to use GetDatabase API.
func (client DatabaseClient) GetDatabase(ctx context.Context, request GetDatabaseRequest) (response GetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseResponse")
	}
	return
}

// getDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/GetDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseSoftwareImage Gets information about the specified database software image.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDatabaseSoftwareImage.go.html to see an example of how to use GetDatabaseSoftwareImage API.
func (client DatabaseClient) GetDatabaseSoftwareImage(ctx context.Context, request GetDatabaseSoftwareImageRequest) (response GetDatabaseSoftwareImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseSoftwareImageResponse")
	}
	return
}

// getDatabaseSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseSoftwareImages/{databaseSoftwareImageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseSoftwareImage/GetDatabaseSoftwareImage"
		err = common.PostProcessServiceError(err, "Database", "GetDatabaseSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseUpgradeHistoryEntry gets the upgrade history for a specified database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDatabaseUpgradeHistoryEntry.go.html to see an example of how to use GetDatabaseUpgradeHistoryEntry API.
func (client DatabaseClient) GetDatabaseUpgradeHistoryEntry(ctx context.Context, request GetDatabaseUpgradeHistoryEntryRequest) (response GetDatabaseUpgradeHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseUpgradeHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseUpgradeHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseUpgradeHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseUpgradeHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseUpgradeHistoryEntryResponse")
	}
	return
}

// getDatabaseUpgradeHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDatabaseUpgradeHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/upgradeHistoryEntries/{upgradeHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseUpgradeHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseUpgradeHistoryEntry/GetDatabaseUpgradeHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetDatabaseUpgradeHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbHome Gets information about the specified Database Home.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbHome.go.html to see an example of how to use GetDbHome API.
func (client DatabaseClient) GetDbHome(ctx context.Context, request GetDbHomeRequest) (response GetDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbHomeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbHomeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbHomeResponse")
	}
	return
}

// getDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbHomeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbHome/GetDbHome"
		err = common.PostProcessServiceError(err, "Database", "GetDbHome", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbHomePatch Gets information about a specified patch package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbHomePatch.go.html to see an example of how to use GetDbHomePatch API.
func (client DatabaseClient) GetDbHomePatch(ctx context.Context, request GetDbHomePatchRequest) (response GetDbHomePatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbHomePatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbHomePatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbHomePatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbHomePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbHomePatchResponse")
	}
	return
}

// getDbHomePatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbHomePatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches/{patchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Patch/GetDbHomePatch"
		err = common.PostProcessServiceError(err, "Database", "GetDbHomePatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbHomePatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbHomePatchHistoryEntry.go.html to see an example of how to use GetDbHomePatchHistoryEntry API.
func (client DatabaseClient) GetDbHomePatchHistoryEntry(ctx context.Context, request GetDbHomePatchHistoryEntryRequest) (response GetDbHomePatchHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbHomePatchHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbHomePatchHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbHomePatchHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbHomePatchHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbHomePatchHistoryEntryResponse")
	}
	return
}

// getDbHomePatchHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbHomePatchHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries/{patchHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PatchHistoryEntry/GetDbHomePatchHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetDbHomePatchHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbNode Gets information about the specified database node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbNode.go.html to see an example of how to use GetDbNode API.
func (client DatabaseClient) GetDbNode(ctx context.Context, request GetDbNodeRequest) (response GetDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbNodeResponse")
	}
	return
}

// getDbNode implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbNode/GetDbNode"
		err = common.PostProcessServiceError(err, "Database", "GetDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbServer Gets information about the Exadata Db server.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbServer.go.html to see an example of how to use GetDbServer API.
func (client DatabaseClient) GetDbServer(ctx context.Context, request GetDbServerRequest) (response GetDbServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbServerResponse")
	}
	return
}

// getDbServer implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbServers/{dbServerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbServer/GetDbServer"
		err = common.PostProcessServiceError(err, "Database", "GetDbServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbSystem Gets information about the specified DB system.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbSystem.go.html to see an example of how to use GetDbSystem API.
func (client DatabaseClient) GetDbSystem(ctx context.Context, request GetDbSystemRequest) (response GetDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemResponse")
	}
	return
}

// getDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/GetDbSystem"
		err = common.PostProcessServiceError(err, "Database", "GetDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbSystemPatch Gets information the specified patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbSystemPatch.go.html to see an example of how to use GetDbSystemPatch API.
func (client DatabaseClient) GetDbSystemPatch(ctx context.Context, request GetDbSystemPatchRequest) (response GetDbSystemPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystemPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbSystemPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbSystemPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemPatchResponse")
	}
	return
}

// getDbSystemPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystemPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches/{patchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Patch/GetDbSystemPatch"
		err = common.PostProcessServiceError(err, "Database", "GetDbSystemPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbSystemPatchHistoryEntry Gets the details of the specified patch operation on the specified DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbSystemPatchHistoryEntry.go.html to see an example of how to use GetDbSystemPatchHistoryEntry API.
func (client DatabaseClient) GetDbSystemPatchHistoryEntry(ctx context.Context, request GetDbSystemPatchHistoryEntryRequest) (response GetDbSystemPatchHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystemPatchHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbSystemPatchHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbSystemPatchHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemPatchHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemPatchHistoryEntryResponse")
	}
	return
}

// getDbSystemPatchHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystemPatchHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries/{patchHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PatchHistoryEntry/GetDbSystemPatchHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetDbSystemPatchHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbSystemUpgradeHistoryEntry Gets the details of the specified operating system upgrade operation for the specified DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetDbSystemUpgradeHistoryEntry.go.html to see an example of how to use GetDbSystemUpgradeHistoryEntry API.
func (client DatabaseClient) GetDbSystemUpgradeHistoryEntry(ctx context.Context, request GetDbSystemUpgradeHistoryEntryRequest) (response GetDbSystemUpgradeHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystemUpgradeHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbSystemUpgradeHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbSystemUpgradeHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemUpgradeHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemUpgradeHistoryEntryResponse")
	}
	return
}

// getDbSystemUpgradeHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getDbSystemUpgradeHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/upgradeHistoryEntries/{upgradeHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemUpgradeHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystemUpgradeHistoryEntry/GetDbSystemUpgradeHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetDbSystemUpgradeHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExadataInfrastructure Gets information about the specified Exadata infrastructure. Applies to Exadata Cloud@Customer instances only.
// To get information on an Exadata Cloud Service infrastructure resource, use the  GetCloudExadataInfrastructure operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExadataInfrastructure.go.html to see an example of how to use GetExadataInfrastructure API.
func (client DatabaseClient) GetExadataInfrastructure(ctx context.Context, request GetExadataInfrastructureRequest) (response GetExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExadataInfrastructureResponse")
	}
	return
}

// getExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/GetExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "GetExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExadataInfrastructureOcpus Gets details of the available and consumed OCPUs for the specified Autonomous Exadata Infrastructure resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExadataInfrastructureOcpus.go.html to see an example of how to use GetExadataInfrastructureOcpus API.
func (client DatabaseClient) GetExadataInfrastructureOcpus(ctx context.Context, request GetExadataInfrastructureOcpusRequest) (response GetExadataInfrastructureOcpusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExadataInfrastructureOcpus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExadataInfrastructureOcpusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExadataInfrastructureOcpusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExadataInfrastructureOcpusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExadataInfrastructureOcpusResponse")
	}
	return
}

// getExadataInfrastructureOcpus implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExadataInfrastructureOcpus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/ocpus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExadataInfrastructureOcpusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OCPUs/GetExadataInfrastructureOcpus"
		err = common.PostProcessServiceError(err, "Database", "GetExadataInfrastructureOcpus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExadataInfrastructureUnAllocatedResources Gets un allocated resources information for the specified Exadata infrastructure. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExadataInfrastructureUnAllocatedResources.go.html to see an example of how to use GetExadataInfrastructureUnAllocatedResources API.
func (client DatabaseClient) GetExadataInfrastructureUnAllocatedResources(ctx context.Context, request GetExadataInfrastructureUnAllocatedResourcesRequest) (response GetExadataInfrastructureUnAllocatedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExadataInfrastructureUnAllocatedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExadataInfrastructureUnAllocatedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExadataInfrastructureUnAllocatedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExadataInfrastructureUnAllocatedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExadataInfrastructureUnAllocatedResourcesResponse")
	}
	return
}

// getExadataInfrastructureUnAllocatedResources implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExadataInfrastructureUnAllocatedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}/unAllocatedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExadataInfrastructureUnAllocatedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructureUnAllocatedResources/GetExadataInfrastructureUnAllocatedResources"
		err = common.PostProcessServiceError(err, "Database", "GetExadataInfrastructureUnAllocatedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExadataIormConfig Gets the IORM configuration settings for the specified cloud Exadata DB system.
// All Exadata service instances have default IORM settings.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
// The GetCloudVmClusterIormConfig API is used for this operation with Exadata systems using the
// new resource model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExadataIormConfig.go.html to see an example of how to use GetExadataIormConfig API.
func (client DatabaseClient) GetExadataIormConfig(ctx context.Context, request GetExadataIormConfigRequest) (response GetExadataIormConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExadataIormConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExadataIormConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExadataIormConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExadataIormConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExadataIormConfigResponse")
	}
	return
}

// getExadataIormConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExadataIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/ExadataIormConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExadataIormConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/GetExadataIormConfig"
		err = common.PostProcessServiceError(err, "Database", "GetExadataIormConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalBackupJob Gets information about the specified external backup job.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.cloud.oracle.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExternalBackupJob.go.html to see an example of how to use GetExternalBackupJob API.
func (client DatabaseClient) GetExternalBackupJob(ctx context.Context, request GetExternalBackupJobRequest) (response GetExternalBackupJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalBackupJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalBackupJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalBackupJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalBackupJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalBackupJobResponse")
	}
	return
}

// getExternalBackupJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExternalBackupJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalBackupJobs/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalBackupJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalBackupJob/GetExternalBackupJob"
		err = common.PostProcessServiceError(err, "Database", "GetExternalBackupJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalContainerDatabase Gets information about the specified external container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExternalContainerDatabase.go.html to see an example of how to use GetExternalContainerDatabase API.
func (client DatabaseClient) GetExternalContainerDatabase(ctx context.Context, request GetExternalContainerDatabaseRequest) (response GetExternalContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalContainerDatabaseResponse")
	}
	return
}

// getExternalContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalcontainerdatabases/{externalContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/GetExternalContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetExternalContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalDatabaseConnector Gets information about the specified external database connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExternalDatabaseConnector.go.html to see an example of how to use GetExternalDatabaseConnector API.
func (client DatabaseClient) GetExternalDatabaseConnector(ctx context.Context, request GetExternalDatabaseConnectorRequest) (response GetExternalDatabaseConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalDatabaseConnectorResponse")
	}
	return
}

// getExternalDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externaldatabaseconnectors/{externalDatabaseConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalDatabaseConnector/GetExternalDatabaseConnector"
		err = common.PostProcessServiceError(err, "Database", "GetExternalDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldatabaseconnector{})
	return response, err
}

// GetExternalNonContainerDatabase Gets information about a specific external non-container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExternalNonContainerDatabase.go.html to see an example of how to use GetExternalNonContainerDatabase API.
func (client DatabaseClient) GetExternalNonContainerDatabase(ctx context.Context, request GetExternalNonContainerDatabaseRequest) (response GetExternalNonContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalNonContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalNonContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalNonContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalNonContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalNonContainerDatabaseResponse")
	}
	return
}

// getExternalNonContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalNonContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/GetExternalNonContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetExternalNonContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalPluggableDatabase Gets information about a specific
// CreateExternalPluggableDatabaseDetails resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExternalPluggableDatabase.go.html to see an example of how to use GetExternalPluggableDatabase API.
func (client DatabaseClient) GetExternalPluggableDatabase(ctx context.Context, request GetExternalPluggableDatabaseRequest) (response GetExternalPluggableDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalPluggableDatabaseResponse")
	}
	return
}

// getExternalPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalpluggabledatabases/{externalPluggableDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/GetExternalPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetExternalPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInfrastructureTargetVersions Gets details of the Exadata Infrastructure target system software versions that can be applied to the specified infrastructure resource for maintenance updates.
// Applies to Exadata Cloud@Customer and Exadata Cloud instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetInfrastructureTargetVersions.go.html to see an example of how to use GetInfrastructureTargetVersions API.
func (client DatabaseClient) GetInfrastructureTargetVersions(ctx context.Context, request GetInfrastructureTargetVersionsRequest) (response GetInfrastructureTargetVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInfrastructureTargetVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInfrastructureTargetVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInfrastructureTargetVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInfrastructureTargetVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInfrastructureTargetVersionsResponse")
	}
	return
}

// getInfrastructureTargetVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getInfrastructureTargetVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/infrastructureTargetVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInfrastructureTargetVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/InfrastructureTargetVersion/GetInfrastructureTargetVersions"
		err = common.PostProcessServiceError(err, "Database", "GetInfrastructureTargetVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetKeyStore Gets information about the specified key store.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetKeyStore.go.html to see an example of how to use GetKeyStore API.
func (client DatabaseClient) GetKeyStore(ctx context.Context, request GetKeyStoreRequest) (response GetKeyStoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getKeyStore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetKeyStoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetKeyStoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetKeyStoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetKeyStoreResponse")
	}
	return
}

// getKeyStore implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keyStores/{keyStoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetKeyStoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/KeyStore/GetKeyStore"
		err = common.PostProcessServiceError(err, "Database", "GetKeyStore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaintenanceRun Gets information about the specified maintenance run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetMaintenanceRun.go.html to see an example of how to use GetMaintenanceRun API.
func (client DatabaseClient) GetMaintenanceRun(ctx context.Context, request GetMaintenanceRunRequest) (response GetMaintenanceRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaintenanceRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaintenanceRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceRunResponse")
	}
	return
}

// getMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns/{maintenanceRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/MaintenanceRun/GetMaintenanceRun"
		err = common.PostProcessServiceError(err, "Database", "GetMaintenanceRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaintenanceRunHistory Gets information about the specified maintenance run history.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetMaintenanceRunHistory.go.html to see an example of how to use GetMaintenanceRunHistory API.
func (client DatabaseClient) GetMaintenanceRunHistory(ctx context.Context, request GetMaintenanceRunHistoryRequest) (response GetMaintenanceRunHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceRunHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaintenanceRunHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaintenanceRunHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceRunHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceRunHistoryResponse")
	}
	return
}

// getMaintenanceRunHistory implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getMaintenanceRunHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRunHistory/{maintenanceRunHistoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceRunHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/MaintenanceRunHistory/GetMaintenanceRunHistory"
		err = common.PostProcessServiceError(err, "Database", "GetMaintenanceRunHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOneoffPatch Gets information about the specified one-off patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetOneoffPatch.go.html to see an example of how to use GetOneoffPatch API.
func (client DatabaseClient) GetOneoffPatch(ctx context.Context, request GetOneoffPatchRequest) (response GetOneoffPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOneoffPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOneoffPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOneoffPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOneoffPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOneoffPatchResponse")
	}
	return
}

// getOneoffPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getOneoffPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oneoffPatches/{oneoffPatchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOneoffPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/GetOneoffPatch"
		err = common.PostProcessServiceError(err, "Database", "GetOneoffPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPdbConversionHistoryEntry Gets the details of operations performed to convert the specified database from non-container (non-CDB) to pluggable (PDB).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetPdbConversionHistoryEntry.go.html to see an example of how to use GetPdbConversionHistoryEntry API.
func (client DatabaseClient) GetPdbConversionHistoryEntry(ctx context.Context, request GetPdbConversionHistoryEntryRequest) (response GetPdbConversionHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPdbConversionHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPdbConversionHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPdbConversionHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPdbConversionHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPdbConversionHistoryEntryResponse")
	}
	return
}

// getPdbConversionHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getPdbConversionHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/pdbConversionHistoryEntries/{pdbConversionHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPdbConversionHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PdbConversionHistoryEntry/GetPdbConversionHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetPdbConversionHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPluggableDatabase Gets information about the specified pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetPluggableDatabase.go.html to see an example of how to use GetPluggableDatabase API.
func (client DatabaseClient) GetPluggableDatabase(ctx context.Context, request GetPluggableDatabaseRequest) (response GetPluggableDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPluggableDatabaseResponse")
	}
	return
}

// getPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pluggableDatabases/{pluggableDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/GetPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "GetPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmCluster Gets information about the VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetVmCluster.go.html to see an example of how to use GetVmCluster API.
func (client DatabaseClient) GetVmCluster(ctx context.Context, request GetVmClusterRequest) (response GetVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterResponse")
	}
	return
}

// getVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/GetVmCluster"
		err = common.PostProcessServiceError(err, "Database", "GetVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmClusterNetwork Gets information about the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To get information about a cloud VM cluster in an Exadata Cloud Service instance, use the GetCloudVmCluster operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetVmClusterNetwork.go.html to see an example of how to use GetVmClusterNetwork API.
func (client DatabaseClient) GetVmClusterNetwork(ctx context.Context, request GetVmClusterNetworkRequest) (response GetVmClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterNetworkResponse")
	}
	return
}

// getVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/GetVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "GetVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmClusterPatch Gets information about a specified patch package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetVmClusterPatch.go.html to see an example of how to use GetVmClusterPatch API.
func (client DatabaseClient) GetVmClusterPatch(ctx context.Context, request GetVmClusterPatchRequest) (response GetVmClusterPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmClusterPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterPatchResponse")
	}
	return
}

// getVmClusterPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getVmClusterPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patches/{patchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Patch/GetVmClusterPatch"
		err = common.PostProcessServiceError(err, "Database", "GetVmClusterPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmClusterPatchHistoryEntry Gets the patch history details for the specified patch history entry.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetVmClusterPatchHistoryEntry.go.html to see an example of how to use GetVmClusterPatchHistoryEntry API.
func (client DatabaseClient) GetVmClusterPatchHistoryEntry(ctx context.Context, request GetVmClusterPatchHistoryEntryRequest) (response GetVmClusterPatchHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmClusterPatchHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterPatchHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterPatchHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterPatchHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterPatchHistoryEntryResponse")
	}
	return
}

// getVmClusterPatchHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getVmClusterPatchHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patchHistoryEntries/{patchHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterPatchHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PatchHistoryEntry/GetVmClusterPatchHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetVmClusterPatchHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmClusterUpdate Gets information about a specified maintenance update package for a VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetVmClusterUpdate.go.html to see an example of how to use GetVmClusterUpdate API.
func (client DatabaseClient) GetVmClusterUpdate(ctx context.Context, request GetVmClusterUpdateRequest) (response GetVmClusterUpdateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmClusterUpdate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterUpdateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterUpdateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterUpdateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterUpdateResponse")
	}
	return
}

// getVmClusterUpdate implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getVmClusterUpdate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/updates/{updateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterUpdateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterUpdate/GetVmClusterUpdate"
		err = common.PostProcessServiceError(err, "Database", "GetVmClusterUpdate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmClusterUpdateHistoryEntry Gets the maintenance update history details for the specified update history entry. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetVmClusterUpdateHistoryEntry.go.html to see an example of how to use GetVmClusterUpdateHistoryEntry API.
func (client DatabaseClient) GetVmClusterUpdateHistoryEntry(ctx context.Context, request GetVmClusterUpdateHistoryEntryRequest) (response GetVmClusterUpdateHistoryEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmClusterUpdateHistoryEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterUpdateHistoryEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterUpdateHistoryEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterUpdateHistoryEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterUpdateHistoryEntryResponse")
	}
	return
}

// getVmClusterUpdateHistoryEntry implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) getVmClusterUpdateHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/updateHistoryEntries/{updateHistoryEntryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterUpdateHistoryEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterUpdateHistoryEntry/GetVmClusterUpdateHistoryEntry"
		err = common.PostProcessServiceError(err, "Database", "GetVmClusterUpdateHistoryEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LaunchAutonomousExadataInfrastructure **Deprecated** To create a new Autonomous Database system on dedicated Exadata Infrastructure, use the CreateCloudExadataInfrastructure and CreateCloudAutonomousVmCluster operations instead. Note that to create an Autonomous VM cluster, you must have an existing Exadata Infrastructure resource to contain the VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/LaunchAutonomousExadataInfrastructure.go.html to see an example of how to use LaunchAutonomousExadataInfrastructure API.
func (client DatabaseClient) LaunchAutonomousExadataInfrastructure(ctx context.Context, request LaunchAutonomousExadataInfrastructureRequest) (response LaunchAutonomousExadataInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.launchAutonomousExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LaunchAutonomousExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LaunchAutonomousExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LaunchAutonomousExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LaunchAutonomousExadataInfrastructureResponse")
	}
	return
}

// launchAutonomousExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) launchAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LaunchAutonomousExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/LaunchAutonomousExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "LaunchAutonomousExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LaunchDbSystem Creates a new DB system in the specified compartment and availability domain. The Oracle
// Database edition that you specify applies to all the databases on that DB system. The selected edition cannot be changed.
// An initial database is created on the DB system based on the request parameters you provide and some default
// options. For detailed information about default options, see Bare metal and virtual machine DB system default options. (https://docs.cloud.oracle.com/Content/Database/Tasks/creatingDBsystem.htm#Default)
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
// Use the CreateCloudExadataInfrastructure and CreateCloudVmCluster APIs to provision a new Exadata Cloud Service instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/LaunchDbSystem.go.html to see an example of how to use LaunchDbSystem API.
func (client DatabaseClient) LaunchDbSystem(ctx context.Context, request LaunchDbSystemRequest) (response LaunchDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.launchDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LaunchDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LaunchDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LaunchDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LaunchDbSystemResponse")
	}
	return
}

// launchDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) launchDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LaunchDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/LaunchDbSystem"
		err = common.PostProcessServiceError(err, "Database", "LaunchDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplicationVips Gets a list of application virtual IP (VIP) addresses on a cloud VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListApplicationVips.go.html to see an example of how to use ListApplicationVips API.
func (client DatabaseClient) ListApplicationVips(ctx context.Context, request ListApplicationVipsRequest) (response ListApplicationVipsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplicationVips, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicationVipsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicationVipsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationVipsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationVipsResponse")
	}
	return
}

// listApplicationVips implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listApplicationVips(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applicationVip", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationVipsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ApplicationVipSummary/ListApplicationVips"
		err = common.PostProcessServiceError(err, "Database", "ListApplicationVips", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousContainerDatabaseDataguardAssociations Gets a list of the Autonomous Container Databases with Autonomous Data Guard-enabled associated with the specified Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousContainerDatabaseDataguardAssociations.go.html to see an example of how to use ListAutonomousContainerDatabaseDataguardAssociations API.
func (client DatabaseClient) ListAutonomousContainerDatabaseDataguardAssociations(ctx context.Context, request ListAutonomousContainerDatabaseDataguardAssociationsRequest) (response ListAutonomousContainerDatabaseDataguardAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousContainerDatabaseDataguardAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousContainerDatabaseDataguardAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousContainerDatabaseDataguardAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousContainerDatabaseDataguardAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousContainerDatabaseDataguardAssociationsResponse")
	}
	return
}

// listAutonomousContainerDatabaseDataguardAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousContainerDatabaseDataguardAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabaseDataguardAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/ListAutonomousContainerDatabaseDataguardAssociations"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousContainerDatabaseDataguardAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousContainerDatabaseVersions Gets a list of supported Autonomous Container Database versions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousContainerDatabaseVersions.go.html to see an example of how to use ListAutonomousContainerDatabaseVersions API.
func (client DatabaseClient) ListAutonomousContainerDatabaseVersions(ctx context.Context, request ListAutonomousContainerDatabaseVersionsRequest) (response ListAutonomousContainerDatabaseVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousContainerDatabaseVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousContainerDatabaseVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousContainerDatabaseVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousContainerDatabaseVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousContainerDatabaseVersionsResponse")
	}
	return
}

// listAutonomousContainerDatabaseVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousContainerDatabaseVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabaseVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabaseVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseVersionSummary/ListAutonomousContainerDatabaseVersions"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousContainerDatabaseVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousContainerDatabases Gets a list of the Autonomous Container Databases in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousContainerDatabases.go.html to see an example of how to use ListAutonomousContainerDatabases API.
func (client DatabaseClient) ListAutonomousContainerDatabases(ctx context.Context, request ListAutonomousContainerDatabasesRequest) (response ListAutonomousContainerDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousContainerDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousContainerDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousContainerDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousContainerDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousContainerDatabasesResponse")
	}
	return
}

// listAutonomousContainerDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousContainerDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/ListAutonomousContainerDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousContainerDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDatabaseBackups Gets a list of Autonomous Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseBackups.go.html to see an example of how to use ListAutonomousDatabaseBackups API.
func (client DatabaseClient) ListAutonomousDatabaseBackups(ctx context.Context, request ListAutonomousDatabaseBackupsRequest) (response ListAutonomousDatabaseBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDatabaseBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDatabaseBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseBackupsResponse")
	}
	return
}

// listAutonomousDatabaseBackups implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseBackup/ListAutonomousDatabaseBackups"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDatabaseBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDatabaseCharacterSets Gets a list of supported character sets.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseCharacterSets.go.html to see an example of how to use ListAutonomousDatabaseCharacterSets API.
func (client DatabaseClient) ListAutonomousDatabaseCharacterSets(ctx context.Context, request ListAutonomousDatabaseCharacterSetsRequest) (response ListAutonomousDatabaseCharacterSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseCharacterSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDatabaseCharacterSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDatabaseCharacterSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseCharacterSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseCharacterSetsResponse")
	}
	return
}

// listAutonomousDatabaseCharacterSets implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseCharacterSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseCharacterSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseCharacterSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseCharacterSets/ListAutonomousDatabaseCharacterSets"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDatabaseCharacterSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDatabaseClones Lists the Autonomous Database clones for the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseClones.go.html to see an example of how to use ListAutonomousDatabaseClones API.
func (client DatabaseClient) ListAutonomousDatabaseClones(ctx context.Context, request ListAutonomousDatabaseClonesRequest) (response ListAutonomousDatabaseClonesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseClones, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDatabaseClonesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDatabaseClonesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseClonesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseClonesResponse")
	}
	return
}

// listAutonomousDatabaseClones implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseClones(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/clones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseClonesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ListAutonomousDatabaseClones"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDatabaseClones", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDatabaseDataguardAssociations Gets a list of the Autonomous Data Guard-enabled databases associated with the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseDataguardAssociations.go.html to see an example of how to use ListAutonomousDatabaseDataguardAssociations API.
func (client DatabaseClient) ListAutonomousDatabaseDataguardAssociations(ctx context.Context, request ListAutonomousDatabaseDataguardAssociationsRequest) (response ListAutonomousDatabaseDataguardAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseDataguardAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDatabaseDataguardAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDatabaseDataguardAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseDataguardAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseDataguardAssociationsResponse")
	}
	return
}

// listAutonomousDatabaseDataguardAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseDataguardAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/autonomousDatabaseDataguardAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseDataguardAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseDataguardAssociation/ListAutonomousDatabaseDataguardAssociations"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDatabaseDataguardAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDatabaseRefreshableClones Lists the OCIDs of the Autonomous Database local and connected remote refreshable clones with the region where they exist for the specified source database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseRefreshableClones.go.html to see an example of how to use ListAutonomousDatabaseRefreshableClones API.
func (client DatabaseClient) ListAutonomousDatabaseRefreshableClones(ctx context.Context, request ListAutonomousDatabaseRefreshableClonesRequest) (response ListAutonomousDatabaseRefreshableClonesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabaseRefreshableClones, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDatabaseRefreshableClonesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDatabaseRefreshableClonesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabaseRefreshableClonesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabaseRefreshableClonesResponse")
	}
	return
}

// listAutonomousDatabaseRefreshableClones implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabaseRefreshableClones(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/refreshableClones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseRefreshableClonesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ListAutonomousDatabaseRefreshableClones"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDatabaseRefreshableClones", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDatabases Gets a list of Autonomous Databases based on the query parameters specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabases.go.html to see an example of how to use ListAutonomousDatabases API.
func (client DatabaseClient) ListAutonomousDatabases(ctx context.Context, request ListAutonomousDatabasesRequest) (response ListAutonomousDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDatabasesResponse")
	}
	return
}

// listAutonomousDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ListAutonomousDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDbPreviewVersions Gets a list of supported Autonomous Database versions. Note that preview version software is only available for
// Autonomous Database Serverless (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDbPreviewVersions.go.html to see an example of how to use ListAutonomousDbPreviewVersions API.
func (client DatabaseClient) ListAutonomousDbPreviewVersions(ctx context.Context, request ListAutonomousDbPreviewVersionsRequest) (response ListAutonomousDbPreviewVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDbPreviewVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDbPreviewVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDbPreviewVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDbPreviewVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDbPreviewVersionsResponse")
	}
	return
}

// listAutonomousDbPreviewVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDbPreviewVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbPreviewVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDbPreviewVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDbPreviewVersionSummary/ListAutonomousDbPreviewVersions"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDbPreviewVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousDbVersions Gets a list of supported Autonomous Database versions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDbVersions.go.html to see an example of how to use ListAutonomousDbVersions API.
func (client DatabaseClient) ListAutonomousDbVersions(ctx context.Context, request ListAutonomousDbVersionsRequest) (response ListAutonomousDbVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousDbVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousDbVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousDbVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousDbVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousDbVersionsResponse")
	}
	return
}

// listAutonomousDbVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousDbVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDbVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDbVersionSummary/ListAutonomousDbVersions"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousDbVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousExadataInfrastructureShapes **Deprecated.**
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousExadataInfrastructureShapes.go.html to see an example of how to use ListAutonomousExadataInfrastructureShapes API.
func (client DatabaseClient) ListAutonomousExadataInfrastructureShapes(ctx context.Context, request ListAutonomousExadataInfrastructureShapesRequest) (response ListAutonomousExadataInfrastructureShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousExadataInfrastructureShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousExadataInfrastructureShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousExadataInfrastructureShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousExadataInfrastructureShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousExadataInfrastructureShapesResponse")
	}
	return
}

// listAutonomousExadataInfrastructureShapes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousExadataInfrastructureShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructureShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousExadataInfrastructureShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructureShapeSummary/ListAutonomousExadataInfrastructureShapes"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousExadataInfrastructureShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousExadataInfrastructures **Deprecated.** Use the ListCloudExadataInfrastructures operation to list Exadata Infrastructures in the Oracle cloud and the  ListCloudAutonomousVmClusters operation to list Autonomous Exadata VM clusters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousExadataInfrastructures.go.html to see an example of how to use ListAutonomousExadataInfrastructures API.
func (client DatabaseClient) ListAutonomousExadataInfrastructures(ctx context.Context, request ListAutonomousExadataInfrastructuresRequest) (response ListAutonomousExadataInfrastructuresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousExadataInfrastructures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousExadataInfrastructuresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousExadataInfrastructuresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousExadataInfrastructuresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousExadataInfrastructuresResponse")
	}
	return
}

// listAutonomousExadataInfrastructures implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousExadataInfrastructuresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/ListAutonomousExadataInfrastructures"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousExadataInfrastructures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousVirtualMachines Lists the Autonomous Virtual Machines in the specified Autonomous VM Cluster and Compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousVirtualMachines.go.html to see an example of how to use ListAutonomousVirtualMachines API.
func (client DatabaseClient) ListAutonomousVirtualMachines(ctx context.Context, request ListAutonomousVirtualMachinesRequest) (response ListAutonomousVirtualMachinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousVirtualMachines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousVirtualMachinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousVirtualMachinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousVirtualMachinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousVirtualMachinesResponse")
	}
	return
}

// listAutonomousVirtualMachines implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousVirtualMachines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVirtualMachines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousVirtualMachinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVirtualMachine/ListAutonomousVirtualMachines"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousVirtualMachines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousVmClusterAcdResourceUsage Gets the list of resource usage details for all the Autonomous Container Database in the specified Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousVmClusterAcdResourceUsage.go.html to see an example of how to use ListAutonomousVmClusterAcdResourceUsage API.
func (client DatabaseClient) ListAutonomousVmClusterAcdResourceUsage(ctx context.Context, request ListAutonomousVmClusterAcdResourceUsageRequest) (response ListAutonomousVmClusterAcdResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousVmClusterAcdResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousVmClusterAcdResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousVmClusterAcdResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousVmClusterAcdResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousVmClusterAcdResourceUsageResponse")
	}
	return
}

// listAutonomousVmClusterAcdResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousVmClusterAcdResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVmClusters/{autonomousVmClusterId}/acdResourceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousVmClusterAcdResourceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/ListAutonomousVmClusterAcdResourceUsage"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousVmClusterAcdResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutonomousVmClusters Gets a list of Exadata Cloud@Customer Autonomous VM clusters in the specified compartment. To list Autonomous VM Clusters in the Oracle Cloud, see ListCloudAutonomousVmClusters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousVmClusters.go.html to see an example of how to use ListAutonomousVmClusters API.
func (client DatabaseClient) ListAutonomousVmClusters(ctx context.Context, request ListAutonomousVmClustersRequest) (response ListAutonomousVmClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutonomousVmClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutonomousVmClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutonomousVmClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutonomousVmClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutonomousVmClustersResponse")
	}
	return
}

// listAutonomousVmClusters implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listAutonomousVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousVmClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/ListAutonomousVmClusters"
		err = common.PostProcessServiceError(err, "Database", "ListAutonomousVmClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBackupDestination Gets a list of backup destinations in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListBackupDestination.go.html to see an example of how to use ListBackupDestination API.
func (client DatabaseClient) ListBackupDestination(ctx context.Context, request ListBackupDestinationRequest) (response ListBackupDestinationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackupDestination, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBackupDestinationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBackupDestinationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBackupDestinationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBackupDestinationResponse")
	}
	return
}

// listBackupDestination implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backupDestinations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBackupDestinationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/BackupDestinationSummary/ListBackupDestination"
		err = common.PostProcessServiceError(err, "Database", "ListBackupDestination", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBackups Gets a list of backups based on the `databaseId` or `compartmentId` specified. Either one of these query parameters must be provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListBackups.go.html to see an example of how to use ListBackups API.
func (client DatabaseClient) ListBackups(ctx context.Context, request ListBackupsRequest) (response ListBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBackupsResponse")
	}
	return
}

// listBackups implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Backup/ListBackups"
		err = common.PostProcessServiceError(err, "Database", "ListBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCloudAutonomousVmClusterAcdResourceUsage Gets the list of resource usage details for all the Cloud Autonomous Container Database
// in the specified Cloud Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudAutonomousVmClusterAcdResourceUsage.go.html to see an example of how to use ListCloudAutonomousVmClusterAcdResourceUsage API.
func (client DatabaseClient) ListCloudAutonomousVmClusterAcdResourceUsage(ctx context.Context, request ListCloudAutonomousVmClusterAcdResourceUsageRequest) (response ListCloudAutonomousVmClusterAcdResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCloudAutonomousVmClusterAcdResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCloudAutonomousVmClusterAcdResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCloudAutonomousVmClusterAcdResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCloudAutonomousVmClusterAcdResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCloudAutonomousVmClusterAcdResourceUsageResponse")
	}
	return
}

// listCloudAutonomousVmClusterAcdResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listCloudAutonomousVmClusterAcdResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}/acdResourceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCloudAutonomousVmClusterAcdResourceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/ListCloudAutonomousVmClusterAcdResourceUsage"
		err = common.PostProcessServiceError(err, "Database", "ListCloudAutonomousVmClusterAcdResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCloudAutonomousVmClusters Lists Autonomous Exadata VM clusters in the Oracle cloud. For Exadata Cloud@Customer systems, see ListAutonomousVmClusters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudAutonomousVmClusters.go.html to see an example of how to use ListCloudAutonomousVmClusters API.
func (client DatabaseClient) ListCloudAutonomousVmClusters(ctx context.Context, request ListCloudAutonomousVmClustersRequest) (response ListCloudAutonomousVmClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCloudAutonomousVmClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCloudAutonomousVmClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCloudAutonomousVmClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCloudAutonomousVmClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCloudAutonomousVmClustersResponse")
	}
	return
}

// listCloudAutonomousVmClusters implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listCloudAutonomousVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudAutonomousVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCloudAutonomousVmClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/ListCloudAutonomousVmClusters"
		err = common.PostProcessServiceError(err, "Database", "ListCloudAutonomousVmClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCloudExadataInfrastructures Gets a list of the cloud Exadata infrastructure resources in the specified compartment. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudExadataInfrastructures.go.html to see an example of how to use ListCloudExadataInfrastructures API.
func (client DatabaseClient) ListCloudExadataInfrastructures(ctx context.Context, request ListCloudExadataInfrastructuresRequest) (response ListCloudExadataInfrastructuresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCloudExadataInfrastructures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCloudExadataInfrastructuresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCloudExadataInfrastructuresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCloudExadataInfrastructuresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCloudExadataInfrastructuresResponse")
	}
	return
}

// listCloudExadataInfrastructures implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listCloudExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudExadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCloudExadataInfrastructuresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/ListCloudExadataInfrastructures"
		err = common.PostProcessServiceError(err, "Database", "ListCloudExadataInfrastructures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCloudVmClusterUpdateHistoryEntries Gets the history of the maintenance update actions performed on the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudVmClusterUpdateHistoryEntries.go.html to see an example of how to use ListCloudVmClusterUpdateHistoryEntries API.
func (client DatabaseClient) ListCloudVmClusterUpdateHistoryEntries(ctx context.Context, request ListCloudVmClusterUpdateHistoryEntriesRequest) (response ListCloudVmClusterUpdateHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCloudVmClusterUpdateHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCloudVmClusterUpdateHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCloudVmClusterUpdateHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCloudVmClusterUpdateHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCloudVmClusterUpdateHistoryEntriesResponse")
	}
	return
}

// listCloudVmClusterUpdateHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listCloudVmClusterUpdateHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updateHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCloudVmClusterUpdateHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/UpdateHistoryEntry/ListCloudVmClusterUpdateHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListCloudVmClusterUpdateHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCloudVmClusterUpdates Lists the maintenance updates that can be applied to the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudVmClusterUpdates.go.html to see an example of how to use ListCloudVmClusterUpdates API.
func (client DatabaseClient) ListCloudVmClusterUpdates(ctx context.Context, request ListCloudVmClusterUpdatesRequest) (response ListCloudVmClusterUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCloudVmClusterUpdates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCloudVmClusterUpdatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCloudVmClusterUpdatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCloudVmClusterUpdatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCloudVmClusterUpdatesResponse")
	}
	return
}

// listCloudVmClusterUpdates implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listCloudVmClusterUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCloudVmClusterUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Update/ListCloudVmClusterUpdates"
		err = common.PostProcessServiceError(err, "Database", "ListCloudVmClusterUpdates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCloudVmClusters Gets a list of the cloud VM clusters in the specified compartment. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudVmClusters.go.html to see an example of how to use ListCloudVmClusters API.
func (client DatabaseClient) ListCloudVmClusters(ctx context.Context, request ListCloudVmClustersRequest) (response ListCloudVmClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCloudVmClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCloudVmClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCloudVmClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCloudVmClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCloudVmClustersResponse")
	}
	return
}

// listCloudVmClusters implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listCloudVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCloudVmClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/ListCloudVmClusters"
		err = common.PostProcessServiceError(err, "Database", "ListCloudVmClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConsoleConnections Lists the console connections for the specified database node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListConsoleConnections.go.html to see an example of how to use ListConsoleConnections API.
func (client DatabaseClient) ListConsoleConnections(ctx context.Context, request ListConsoleConnectionsRequest) (response ListConsoleConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConsoleConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConsoleConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConsoleConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConsoleConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConsoleConnectionsResponse")
	}
	return
}

// listConsoleConnections implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listConsoleConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleConnections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConsoleConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleConnection/ListConsoleConnections"
		err = common.PostProcessServiceError(err, "Database", "ListConsoleConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConsoleHistories Lists the console histories for the specified database node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListConsoleHistories.go.html to see an example of how to use ListConsoleHistories API.
func (client DatabaseClient) ListConsoleHistories(ctx context.Context, request ListConsoleHistoriesRequest) (response ListConsoleHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConsoleHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConsoleHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConsoleHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConsoleHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConsoleHistoriesResponse")
	}
	return
}

// listConsoleHistories implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listConsoleHistories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleHistories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConsoleHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleHistory/ListConsoleHistories"
		err = common.PostProcessServiceError(err, "Database", "ListConsoleHistories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListContainerDatabasePatches Lists the patches applicable to the requested container database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListContainerDatabasePatches.go.html to see an example of how to use ListContainerDatabasePatches API.
func (client DatabaseClient) ListContainerDatabasePatches(ctx context.Context, request ListContainerDatabasePatchesRequest) (response ListContainerDatabasePatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listContainerDatabasePatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListContainerDatabasePatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListContainerDatabasePatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListContainerDatabasePatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListContainerDatabasePatchesResponse")
	}
	return
}

// listContainerDatabasePatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listContainerDatabasePatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListContainerDatabasePatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousPatch/ListContainerDatabasePatches"
		err = common.PostProcessServiceError(err, "Database", "ListContainerDatabasePatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataGuardAssociations Lists all Data Guard associations for the specified database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDataGuardAssociations.go.html to see an example of how to use ListDataGuardAssociations API.
func (client DatabaseClient) ListDataGuardAssociations(ctx context.Context, request ListDataGuardAssociationsRequest) (response ListDataGuardAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataGuardAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataGuardAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataGuardAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataGuardAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataGuardAssociationsResponse")
	}
	return
}

// listDataGuardAssociations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDataGuardAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataGuardAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/ListDataGuardAssociations"
		err = common.PostProcessServiceError(err, "Database", "ListDataGuardAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseSoftwareImages Gets a list of the database software images in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDatabaseSoftwareImages.go.html to see an example of how to use ListDatabaseSoftwareImages API.
func (client DatabaseClient) ListDatabaseSoftwareImages(ctx context.Context, request ListDatabaseSoftwareImagesRequest) (response ListDatabaseSoftwareImagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseSoftwareImages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseSoftwareImagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseSoftwareImagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseSoftwareImagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseSoftwareImagesResponse")
	}
	return
}

// listDatabaseSoftwareImages implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDatabaseSoftwareImages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseSoftwareImages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseSoftwareImagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseSoftwareImage/ListDatabaseSoftwareImages"
		err = common.PostProcessServiceError(err, "Database", "ListDatabaseSoftwareImages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseUpgradeHistoryEntries Gets the upgrade history for a specified database in a bare metal or virtual machine DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDatabaseUpgradeHistoryEntries.go.html to see an example of how to use ListDatabaseUpgradeHistoryEntries API.
func (client DatabaseClient) ListDatabaseUpgradeHistoryEntries(ctx context.Context, request ListDatabaseUpgradeHistoryEntriesRequest) (response ListDatabaseUpgradeHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseUpgradeHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseUpgradeHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseUpgradeHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseUpgradeHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseUpgradeHistoryEntriesResponse")
	}
	return
}

// listDatabaseUpgradeHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDatabaseUpgradeHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/upgradeHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseUpgradeHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/ListDatabaseUpgradeHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListDatabaseUpgradeHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabases Gets a list of the databases in the specified Database Home.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDatabases.go.html to see an example of how to use ListDatabases API.
func (client DatabaseClient) ListDatabases(ctx context.Context, request ListDatabasesRequest) (response ListDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabasesResponse")
	}
	return
}

// listDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/ListDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbHomePatchHistoryEntries Lists the history of patch operations on the specified Database Home.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbHomePatchHistoryEntries.go.html to see an example of how to use ListDbHomePatchHistoryEntries API.
func (client DatabaseClient) ListDbHomePatchHistoryEntries(ctx context.Context, request ListDbHomePatchHistoryEntriesRequest) (response ListDbHomePatchHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbHomePatchHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbHomePatchHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbHomePatchHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbHomePatchHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbHomePatchHistoryEntriesResponse")
	}
	return
}

// listDbHomePatchHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbHomePatchHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PatchHistoryEntry/ListDbHomePatchHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListDbHomePatchHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbHomePatches Lists patches applicable to the requested Database Home.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbHomePatches.go.html to see an example of how to use ListDbHomePatches API.
func (client DatabaseClient) ListDbHomePatches(ctx context.Context, request ListDbHomePatchesRequest) (response ListDbHomePatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbHomePatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbHomePatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbHomePatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbHomePatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbHomePatchesResponse")
	}
	return
}

// listDbHomePatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbHomePatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Patch/ListDbHomePatches"
		err = common.PostProcessServiceError(err, "Database", "ListDbHomePatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbHomes Lists the Database Homes in the specified DB system and compartment. A Database Home is a directory where Oracle Database software is installed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbHomes.go.html to see an example of how to use ListDbHomes API.
func (client DatabaseClient) ListDbHomes(ctx context.Context, request ListDbHomesRequest) (response ListDbHomesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbHomes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbHomesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbHomesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbHomesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbHomesResponse")
	}
	return
}

// listDbHomes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbHomes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbHomesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbHome/ListDbHomes"
		err = common.PostProcessServiceError(err, "Database", "ListDbHomes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbNodes Lists the database nodes in the specified DB system and compartment. A database node is a server running database software.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbNodes.go.html to see an example of how to use ListDbNodes API.
func (client DatabaseClient) ListDbNodes(ctx context.Context, request ListDbNodesRequest) (response ListDbNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbNodesResponse")
	}
	return
}

// listDbNodes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbNode/ListDbNodes"
		err = common.PostProcessServiceError(err, "Database", "ListDbNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbServers Lists the Exadata DB servers in the ExadataInfrastructureId and specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbServers.go.html to see an example of how to use ListDbServers API.
func (client DatabaseClient) ListDbServers(ctx context.Context, request ListDbServersRequest) (response ListDbServersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbServers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbServersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbServersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbServersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbServersResponse")
	}
	return
}

// listDbServers implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbServers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbServers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbServersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbServer/ListDbServers"
		err = common.PostProcessServiceError(err, "Database", "ListDbServers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystemComputePerformances Gets a list of expected compute performance parameters for a virtual machine DB system based on system configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemComputePerformances.go.html to see an example of how to use ListDbSystemComputePerformances API.
func (client DatabaseClient) ListDbSystemComputePerformances(ctx context.Context, request ListDbSystemComputePerformancesRequest) (response ListDbSystemComputePerformancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemComputePerformances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemComputePerformancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemComputePerformancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemComputePerformancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemComputePerformancesResponse")
	}
	return
}

// listDbSystemComputePerformances implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemComputePerformances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemComputePerformance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemComputePerformancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/ListDbSystemComputePerformances"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystemComputePerformances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystemPatchHistoryEntries Gets the history of the patch actions performed on the specified DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemPatchHistoryEntries.go.html to see an example of how to use ListDbSystemPatchHistoryEntries API.
func (client DatabaseClient) ListDbSystemPatchHistoryEntries(ctx context.Context, request ListDbSystemPatchHistoryEntriesRequest) (response ListDbSystemPatchHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemPatchHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemPatchHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemPatchHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemPatchHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemPatchHistoryEntriesResponse")
	}
	return
}

// listDbSystemPatchHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemPatchHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PatchHistoryEntry/ListDbSystemPatchHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystemPatchHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystemPatches Lists the patches applicable to the specified DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemPatches.go.html to see an example of how to use ListDbSystemPatches API.
func (client DatabaseClient) ListDbSystemPatches(ctx context.Context, request ListDbSystemPatchesRequest) (response ListDbSystemPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemPatchesResponse")
	}
	return
}

// listDbSystemPatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Patch/ListDbSystemPatches"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystemPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystemShapes Gets a list of the shapes that can be used to launch a new DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemShapes.go.html to see an example of how to use ListDbSystemShapes API.
func (client DatabaseClient) ListDbSystemShapes(ctx context.Context, request ListDbSystemShapesRequest) (response ListDbSystemShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemShapesResponse")
	}
	return
}

// listDbSystemShapes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystemShapeSummary/ListDbSystemShapes"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystemShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystemStoragePerformances Gets a list of possible expected storage performance parameters of a VMDB System based on Configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemStoragePerformances.go.html to see an example of how to use ListDbSystemStoragePerformances API.
func (client DatabaseClient) ListDbSystemStoragePerformances(ctx context.Context, request ListDbSystemStoragePerformancesRequest) (response ListDbSystemStoragePerformancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemStoragePerformances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemStoragePerformancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemStoragePerformancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemStoragePerformancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemStoragePerformancesResponse")
	}
	return
}

// listDbSystemStoragePerformances implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemStoragePerformances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemStoragePerformance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemStoragePerformancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/ListDbSystemStoragePerformances"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystemStoragePerformances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystemUpgradeHistoryEntries Gets the history of the upgrade actions performed on the specified DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemUpgradeHistoryEntries.go.html to see an example of how to use ListDbSystemUpgradeHistoryEntries API.
func (client DatabaseClient) ListDbSystemUpgradeHistoryEntries(ctx context.Context, request ListDbSystemUpgradeHistoryEntriesRequest) (response ListDbSystemUpgradeHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystemUpgradeHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemUpgradeHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemUpgradeHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemUpgradeHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemUpgradeHistoryEntriesResponse")
	}
	return
}

// listDbSystemUpgradeHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystemUpgradeHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/upgradeHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemUpgradeHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystemUpgradeHistoryEntry/ListDbSystemUpgradeHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystemUpgradeHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbSystems Lists the DB systems in the specified compartment. You can specify a `backupId` to list only the DB systems that support creating a database using this backup in this compartment.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystems.go.html to see an example of how to use ListDbSystems API.
func (client DatabaseClient) ListDbSystems(ctx context.Context, request ListDbSystemsRequest) (response ListDbSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemsResponse")
	}
	return
}

// listDbSystems implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/ListDbSystems"
		err = common.PostProcessServiceError(err, "Database", "ListDbSystems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbVersions Gets a list of supported Oracle Database versions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbVersions.go.html to see an example of how to use ListDbVersions API.
func (client DatabaseClient) ListDbVersions(ctx context.Context, request ListDbVersionsRequest) (response ListDbVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbVersionsResponse")
	}
	return
}

// listDbVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listDbVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions"
		err = common.PostProcessServiceError(err, "Database", "ListDbVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExadataInfrastructures Lists the Exadata infrastructure resources in the specified compartment. Applies to Exadata Cloud@Customer instances only.
// To list the Exadata Cloud Service infrastructure resources in a compartment, use the  ListCloudExadataInfrastructures operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExadataInfrastructures.go.html to see an example of how to use ListExadataInfrastructures API.
func (client DatabaseClient) ListExadataInfrastructures(ctx context.Context, request ListExadataInfrastructuresRequest) (response ListExadataInfrastructuresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExadataInfrastructures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExadataInfrastructuresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExadataInfrastructuresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExadataInfrastructuresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExadataInfrastructuresResponse")
	}
	return
}

// listExadataInfrastructures implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExadataInfrastructuresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/ListExadataInfrastructures"
		err = common.PostProcessServiceError(err, "Database", "ListExadataInfrastructures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalContainerDatabases Gets a list of the external container databases in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalContainerDatabases.go.html to see an example of how to use ListExternalContainerDatabases API.
func (client DatabaseClient) ListExternalContainerDatabases(ctx context.Context, request ListExternalContainerDatabasesRequest) (response ListExternalContainerDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalContainerDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalContainerDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalContainerDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalContainerDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalContainerDatabasesResponse")
	}
	return
}

// listExternalContainerDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listExternalContainerDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalcontainerdatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalContainerDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/ListExternalContainerDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListExternalContainerDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// listexternaldatabaseconnectorsummary allows to unmarshal list of polymorphic ExternalDatabaseConnectorSummary
type listexternaldatabaseconnectorsummary []externaldatabaseconnectorsummary

// UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
func (m *listexternaldatabaseconnectorsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	res := make([]ExternalDatabaseConnectorSummary, len(*m))
	for i, v := range *m {
		nn, err := v.UnmarshalPolymorphicJSON(v.JsonData)
		if err != nil {
			return nil, err
		}
		res[i] = nn.(ExternalDatabaseConnectorSummary)
	}
	return res, nil
}

// ListExternalDatabaseConnectors Gets a list of the external database connectors in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalDatabaseConnectors.go.html to see an example of how to use ListExternalDatabaseConnectors API.
func (client DatabaseClient) ListExternalDatabaseConnectors(ctx context.Context, request ListExternalDatabaseConnectorsRequest) (response ListExternalDatabaseConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalDatabaseConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalDatabaseConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalDatabaseConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalDatabaseConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalDatabaseConnectorsResponse")
	}
	return
}

// listExternalDatabaseConnectors implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listExternalDatabaseConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externaldatabaseconnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalDatabaseConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalDatabaseConnector/ListExternalDatabaseConnectors"
		err = common.PostProcessServiceError(err, "Database", "ListExternalDatabaseConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listexternaldatabaseconnectorsummary{})
	return response, err
}

// ListExternalNonContainerDatabases Gets a list of the ExternalNonContainerDatabases in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalNonContainerDatabases.go.html to see an example of how to use ListExternalNonContainerDatabases API.
func (client DatabaseClient) ListExternalNonContainerDatabases(ctx context.Context, request ListExternalNonContainerDatabasesRequest) (response ListExternalNonContainerDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalNonContainerDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalNonContainerDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalNonContainerDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalNonContainerDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalNonContainerDatabasesResponse")
	}
	return
}

// listExternalNonContainerDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listExternalNonContainerDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalnoncontainerdatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalNonContainerDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/ListExternalNonContainerDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListExternalNonContainerDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalPluggableDatabases Gets a list of the CreateExternalPluggableDatabaseDetails
// resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalPluggableDatabases.go.html to see an example of how to use ListExternalPluggableDatabases API.
func (client DatabaseClient) ListExternalPluggableDatabases(ctx context.Context, request ListExternalPluggableDatabasesRequest) (response ListExternalPluggableDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalPluggableDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalPluggableDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalPluggableDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalPluggableDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalPluggableDatabasesResponse")
	}
	return
}

// listExternalPluggableDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listExternalPluggableDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalpluggabledatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalPluggableDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/ListExternalPluggableDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListExternalPluggableDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFlexComponents Gets a list of the flex components that can be used to launch a new DB system. The flex component determines resources to allocate to the DB system - Database Servers and Storage Servers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListFlexComponents.go.html to see an example of how to use ListFlexComponents API.
func (client DatabaseClient) ListFlexComponents(ctx context.Context, request ListFlexComponentsRequest) (response ListFlexComponentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFlexComponents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFlexComponentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFlexComponentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFlexComponentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFlexComponentsResponse")
	}
	return
}

// listFlexComponents implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listFlexComponents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemShapes/flexComponents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFlexComponentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/FlexComponentCollection/ListFlexComponents"
		err = common.PostProcessServiceError(err, "Database", "ListFlexComponents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGiVersions Gets a list of supported GI versions for the Exadata Cloud@Customer VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListGiVersions.go.html to see an example of how to use ListGiVersions API.
func (client DatabaseClient) ListGiVersions(ctx context.Context, request ListGiVersionsRequest) (response ListGiVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGiVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGiVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGiVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGiVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGiVersionsResponse")
	}
	return
}

// listGiVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listGiVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/giVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGiVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/GiVersionSummary/ListGiVersions"
		err = common.PostProcessServiceError(err, "Database", "ListGiVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListKeyStores Gets a list of key stores in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListKeyStores.go.html to see an example of how to use ListKeyStores API.
func (client DatabaseClient) ListKeyStores(ctx context.Context, request ListKeyStoresRequest) (response ListKeyStoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listKeyStores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListKeyStoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListKeyStoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListKeyStoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKeyStoresResponse")
	}
	return
}

// listKeyStores implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listKeyStores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keyStores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListKeyStoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/KeyStoreSummary/ListKeyStores"
		err = common.PostProcessServiceError(err, "Database", "ListKeyStores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaintenanceRunHistory Gets a list of the maintenance run histories in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListMaintenanceRunHistory.go.html to see an example of how to use ListMaintenanceRunHistory API.
func (client DatabaseClient) ListMaintenanceRunHistory(ctx context.Context, request ListMaintenanceRunHistoryRequest) (response ListMaintenanceRunHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceRunHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaintenanceRunHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaintenanceRunHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceRunHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceRunHistoryResponse")
	}
	return
}

// listMaintenanceRunHistory implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listMaintenanceRunHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRunHistory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceRunHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/MaintenanceRunHistory/ListMaintenanceRunHistory"
		err = common.PostProcessServiceError(err, "Database", "ListMaintenanceRunHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaintenanceRuns Gets a list of the maintenance runs in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListMaintenanceRuns.go.html to see an example of how to use ListMaintenanceRuns API.
func (client DatabaseClient) ListMaintenanceRuns(ctx context.Context, request ListMaintenanceRunsRequest) (response ListMaintenanceRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaintenanceRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaintenanceRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceRunsResponse")
	}
	return
}

// listMaintenanceRuns implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listMaintenanceRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/MaintenanceRun/ListMaintenanceRuns"
		err = common.PostProcessServiceError(err, "Database", "ListMaintenanceRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOneoffPatches Lists one-off patches in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListOneoffPatches.go.html to see an example of how to use ListOneoffPatches API.
func (client DatabaseClient) ListOneoffPatches(ctx context.Context, request ListOneoffPatchesRequest) (response ListOneoffPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOneoffPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOneoffPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOneoffPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOneoffPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOneoffPatchesResponse")
	}
	return
}

// listOneoffPatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listOneoffPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oneoffPatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOneoffPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/ListOneoffPatches"
		err = common.PostProcessServiceError(err, "Database", "ListOneoffPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPdbConversionHistoryEntries Gets the pluggable database conversion history for a specified database in a bare metal or virtual machine DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListPdbConversionHistoryEntries.go.html to see an example of how to use ListPdbConversionHistoryEntries API.
func (client DatabaseClient) ListPdbConversionHistoryEntries(ctx context.Context, request ListPdbConversionHistoryEntriesRequest) (response ListPdbConversionHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPdbConversionHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPdbConversionHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPdbConversionHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPdbConversionHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPdbConversionHistoryEntriesResponse")
	}
	return
}

// listPdbConversionHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listPdbConversionHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/pdbConversionHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPdbConversionHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/ListPdbConversionHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListPdbConversionHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPluggableDatabases Gets a list of the pluggable databases in a database or compartment. You must provide either a `databaseId` or `compartmentId` value.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListPluggableDatabases.go.html to see an example of how to use ListPluggableDatabases API.
func (client DatabaseClient) ListPluggableDatabases(ctx context.Context, request ListPluggableDatabasesRequest) (response ListPluggableDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPluggableDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPluggableDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPluggableDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPluggableDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPluggableDatabasesResponse")
	}
	return
}

// listPluggableDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listPluggableDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pluggableDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPluggableDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/ListPluggableDatabases"
		err = common.PostProcessServiceError(err, "Database", "ListPluggableDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSystemVersions Gets a list of supported Exadata system versions for a given shape and GI version.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListSystemVersions.go.html to see an example of how to use ListSystemVersions API.
func (client DatabaseClient) ListSystemVersions(ctx context.Context, request ListSystemVersionsRequest) (response ListSystemVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSystemVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSystemVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSystemVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSystemVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSystemVersionsResponse")
	}
	return
}

// listSystemVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listSystemVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/systemVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSystemVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/SystemVersionCollection/ListSystemVersions"
		err = common.PostProcessServiceError(err, "Database", "ListSystemVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusterNetworks Gets a list of the VM cluster networks in the specified compartment. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterNetworks.go.html to see an example of how to use ListVmClusterNetworks API.
func (client DatabaseClient) ListVmClusterNetworks(ctx context.Context, request ListVmClusterNetworksRequest) (response ListVmClusterNetworksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusterNetworks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClusterNetworksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClusterNetworksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClusterNetworksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClusterNetworksResponse")
	}
	return
}

// listVmClusterNetworks implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listVmClusterNetworks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterNetworksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/ListVmClusterNetworks"
		err = common.PostProcessServiceError(err, "Database", "ListVmClusterNetworks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusterPatchHistoryEntries Gets the history of the patch actions performed on the specified VM cluster in an Exadata Cloud@Customer system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterPatchHistoryEntries.go.html to see an example of how to use ListVmClusterPatchHistoryEntries API.
func (client DatabaseClient) ListVmClusterPatchHistoryEntries(ctx context.Context, request ListVmClusterPatchHistoryEntriesRequest) (response ListVmClusterPatchHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusterPatchHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClusterPatchHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClusterPatchHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClusterPatchHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClusterPatchHistoryEntriesResponse")
	}
	return
}

// listVmClusterPatchHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listVmClusterPatchHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patchHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterPatchHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PatchHistoryEntry/ListVmClusterPatchHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListVmClusterPatchHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusterPatches Lists the patches applicable to the specified VM cluster in an Exadata Cloud@Customer system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterPatches.go.html to see an example of how to use ListVmClusterPatches API.
func (client DatabaseClient) ListVmClusterPatches(ctx context.Context, request ListVmClusterPatchesRequest) (response ListVmClusterPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusterPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClusterPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClusterPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClusterPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClusterPatchesResponse")
	}
	return
}

// listVmClusterPatches implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listVmClusterPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Patch/ListVmClusterPatches"
		err = common.PostProcessServiceError(err, "Database", "ListVmClusterPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusterUpdateHistoryEntries Gets the history of the maintenance update actions performed on the specified VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterUpdateHistoryEntries.go.html to see an example of how to use ListVmClusterUpdateHistoryEntries API.
func (client DatabaseClient) ListVmClusterUpdateHistoryEntries(ctx context.Context, request ListVmClusterUpdateHistoryEntriesRequest) (response ListVmClusterUpdateHistoryEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusterUpdateHistoryEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClusterUpdateHistoryEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClusterUpdateHistoryEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClusterUpdateHistoryEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClusterUpdateHistoryEntriesResponse")
	}
	return
}

// listVmClusterUpdateHistoryEntries implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listVmClusterUpdateHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/updateHistoryEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterUpdateHistoryEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterUpdateHistoryEntry/ListVmClusterUpdateHistoryEntries"
		err = common.PostProcessServiceError(err, "Database", "ListVmClusterUpdateHistoryEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusterUpdates Lists the maintenance updates that can be applied to the specified VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterUpdates.go.html to see an example of how to use ListVmClusterUpdates API.
func (client DatabaseClient) ListVmClusterUpdates(ctx context.Context, request ListVmClusterUpdatesRequest) (response ListVmClusterUpdatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusterUpdates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClusterUpdatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClusterUpdatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClusterUpdatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClusterUpdatesResponse")
	}
	return
}

// listVmClusterUpdates implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listVmClusterUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/updates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterUpdatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterUpdate/ListVmClusterUpdates"
		err = common.PostProcessServiceError(err, "Database", "ListVmClusterUpdates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusters Lists the VM clusters in the specified compartment. Applies to Exadata Cloud@Customer instances only.
// To list the cloud VM clusters in an Exadata Cloud Service instance, use the ListCloudVmClusters operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusters.go.html to see an example of how to use ListVmClusters API.
func (client DatabaseClient) ListVmClusters(ctx context.Context, request ListVmClustersRequest) (response ListVmClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClustersResponse")
	}
	return
}

// listVmClusters implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) listVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/ListVmClusters"
		err = common.PostProcessServiceError(err, "Database", "ListVmClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LocalClonePluggableDatabase **Deprecated.** Use CreatePluggableDatabase for Pluggable Database LocalClone Operation.
// Clones and starts a pluggable database (PDB) in the same database (CDB) as the source PDB. The source PDB must be in the `READ_WRITE` openMode to perform the clone operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/LocalClonePluggableDatabase.go.html to see an example of how to use LocalClonePluggableDatabase API.
func (client DatabaseClient) LocalClonePluggableDatabase(ctx context.Context, request LocalClonePluggableDatabaseRequest) (response LocalClonePluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.localClonePluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LocalClonePluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LocalClonePluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LocalClonePluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LocalClonePluggableDatabaseResponse")
	}
	return
}

// localClonePluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) localClonePluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/localClone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LocalClonePluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/LocalClonePluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "LocalClonePluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MigrateExadataDbSystemResourceModel Migrates the Exadata DB system to the new Exadata resource model (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model).
// All related resources will be migrated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/MigrateExadataDbSystemResourceModel.go.html to see an example of how to use MigrateExadataDbSystemResourceModel API.
func (client DatabaseClient) MigrateExadataDbSystemResourceModel(ctx context.Context, request MigrateExadataDbSystemResourceModelRequest) (response MigrateExadataDbSystemResourceModelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.migrateExadataDbSystemResourceModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MigrateExadataDbSystemResourceModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MigrateExadataDbSystemResourceModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MigrateExadataDbSystemResourceModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MigrateExadataDbSystemResourceModelResponse")
	}
	return
}

// migrateExadataDbSystemResourceModel implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) migrateExadataDbSystemResourceModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/migration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MigrateExadataDbSystemResourceModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/MigrateExadataDbSystemResourceModel"
		err = common.PostProcessServiceError(err, "Database", "MigrateExadataDbSystemResourceModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MigrateVaultKey Changes encryption key management from customer-managed, using the Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm), to Oracle-managed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/MigrateVaultKey.go.html to see an example of how to use MigrateVaultKey API.
func (client DatabaseClient) MigrateVaultKey(ctx context.Context, request MigrateVaultKeyRequest) (response MigrateVaultKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.migrateVaultKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MigrateVaultKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MigrateVaultKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MigrateVaultKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MigrateVaultKeyResponse")
	}
	return
}

// migrateVaultKey implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) migrateVaultKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/migrateKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MigrateVaultKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/MigrateVaultKey"
		err = common.PostProcessServiceError(err, "Database", "MigrateVaultKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ModifyDatabaseManagement Updates one or more attributes of the Database Management service for the database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ModifyDatabaseManagement.go.html to see an example of how to use ModifyDatabaseManagement API.
func (client DatabaseClient) ModifyDatabaseManagement(ctx context.Context, request ModifyDatabaseManagementRequest) (response ModifyDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.modifyDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ModifyDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ModifyDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ModifyDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ModifyDatabaseManagementResponse")
	}
	return
}

// modifyDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) modifyDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/modifyDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ModifyDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/ModifyDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "ModifyDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ModifyPluggableDatabaseManagement Updates one or more attributes of the Database Management service for the pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ModifyPluggableDatabaseManagement.go.html to see an example of how to use ModifyPluggableDatabaseManagement API.
func (client DatabaseClient) ModifyPluggableDatabaseManagement(ctx context.Context, request ModifyPluggableDatabaseManagementRequest) (response ModifyPluggableDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.modifyPluggableDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ModifyPluggableDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ModifyPluggableDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ModifyPluggableDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ModifyPluggableDatabaseManagementResponse")
	}
	return
}

// modifyPluggableDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) modifyPluggableDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/modifyPluggableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ModifyPluggableDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/ModifyPluggableDatabaseManagement"
		err = common.PostProcessServiceError(err, "Database", "ModifyPluggableDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshPluggableDatabase Refreshes a pluggable database (PDB) Refreshable clone.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RefreshPluggableDatabase.go.html to see an example of how to use RefreshPluggableDatabase API.
func (client DatabaseClient) RefreshPluggableDatabase(ctx context.Context, request RefreshPluggableDatabaseRequest) (response RefreshPluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshPluggableDatabaseResponse")
	}
	return
}

// refreshPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) refreshPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/RefreshPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "RefreshPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RegisterAutonomousDatabaseDataSafe Asynchronously registers this Autonomous Database with Data Safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RegisterAutonomousDatabaseDataSafe.go.html to see an example of how to use RegisterAutonomousDatabaseDataSafe API.
func (client DatabaseClient) RegisterAutonomousDatabaseDataSafe(ctx context.Context, request RegisterAutonomousDatabaseDataSafeRequest) (response RegisterAutonomousDatabaseDataSafeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.registerAutonomousDatabaseDataSafe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RegisterAutonomousDatabaseDataSafeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RegisterAutonomousDatabaseDataSafeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RegisterAutonomousDatabaseDataSafeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RegisterAutonomousDatabaseDataSafeResponse")
	}
	return
}

// registerAutonomousDatabaseDataSafe implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) registerAutonomousDatabaseDataSafe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/registerDataSafe", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RegisterAutonomousDatabaseDataSafeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/RegisterAutonomousDatabaseDataSafe"
		err = common.PostProcessServiceError(err, "Database", "RegisterAutonomousDatabaseDataSafe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReinstateAutonomousContainerDatabaseDataguardAssociation Reinstates a disabled standby Autonomous Container Database, identified by the autonomousContainerDatabaseId parameter, to an active standby Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ReinstateAutonomousContainerDatabaseDataguardAssociation.go.html to see an example of how to use ReinstateAutonomousContainerDatabaseDataguardAssociation API.
func (client DatabaseClient) ReinstateAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request ReinstateAutonomousContainerDatabaseDataguardAssociationRequest) (response ReinstateAutonomousContainerDatabaseDataguardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reinstateAutonomousContainerDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReinstateAutonomousContainerDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReinstateAutonomousContainerDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReinstateAutonomousContainerDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReinstateAutonomousContainerDatabaseDataguardAssociationResponse")
	}
	return
}

// reinstateAutonomousContainerDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) reinstateAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}/actions/reinstate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReinstateAutonomousContainerDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/ReinstateAutonomousContainerDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "ReinstateAutonomousContainerDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReinstateDataGuardAssociation Reinstates the database identified by the `databaseId` parameter into the standby role in a Data Guard association.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ReinstateDataGuardAssociation.go.html to see an example of how to use ReinstateDataGuardAssociation API.
func (client DatabaseClient) ReinstateDataGuardAssociation(ctx context.Context, request ReinstateDataGuardAssociationRequest) (response ReinstateDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reinstateDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReinstateDataGuardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReinstateDataGuardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReinstateDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReinstateDataGuardAssociationResponse")
	}
	return
}

// reinstateDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) reinstateDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/reinstate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReinstateDataGuardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/ReinstateDataGuardAssociation"
		err = common.PostProcessServiceError(err, "Database", "ReinstateDataGuardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoteClonePluggableDatabase **Deprecated.** Use CreatePluggableDatabase for Pluggable Database RemoteClone Operation.
// Clones a pluggable database (PDB) to a different database from the source PDB. The cloned PDB will be started upon completion of the clone operation. The source PDB must be in the `READ_WRITE` openMode when performing the clone.
// For Exadata Cloud@Customer instances, the source pluggable database (PDB) must be on the same Exadata Infrastructure as the target container database (CDB) to create a remote clone.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RemoteClonePluggableDatabase.go.html to see an example of how to use RemoteClonePluggableDatabase API.
func (client DatabaseClient) RemoteClonePluggableDatabase(ctx context.Context, request RemoteClonePluggableDatabaseRequest) (response RemoteClonePluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.remoteClonePluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoteClonePluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoteClonePluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoteClonePluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoteClonePluggableDatabaseResponse")
	}
	return
}

// remoteClonePluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) remoteClonePluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/remoteClone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoteClonePluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/RemoteClonePluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "RemoteClonePluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveVirtualMachineFromCloudVmCluster Remove Virtual Machines from the Cloud VM cluster. Applies to Exadata Cloud instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RemoveVirtualMachineFromCloudVmCluster.go.html to see an example of how to use RemoveVirtualMachineFromCloudVmCluster API.
func (client DatabaseClient) RemoveVirtualMachineFromCloudVmCluster(ctx context.Context, request RemoveVirtualMachineFromCloudVmClusterRequest) (response RemoveVirtualMachineFromCloudVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeVirtualMachineFromCloudVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveVirtualMachineFromCloudVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveVirtualMachineFromCloudVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveVirtualMachineFromCloudVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveVirtualMachineFromCloudVmClusterResponse")
	}
	return
}

// removeVirtualMachineFromCloudVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) removeVirtualMachineFromCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudVmClusters/{cloudVmClusterId}/actions/removeVirtualMachine", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveVirtualMachineFromCloudVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/RemoveVirtualMachineFromCloudVmCluster"
		err = common.PostProcessServiceError(err, "Database", "RemoveVirtualMachineFromCloudVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveVirtualMachineFromVmCluster Remove Virtual Machines from the VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RemoveVirtualMachineFromVmCluster.go.html to see an example of how to use RemoveVirtualMachineFromVmCluster API.
func (client DatabaseClient) RemoveVirtualMachineFromVmCluster(ctx context.Context, request RemoveVirtualMachineFromVmClusterRequest) (response RemoveVirtualMachineFromVmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeVirtualMachineFromVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveVirtualMachineFromVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveVirtualMachineFromVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveVirtualMachineFromVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveVirtualMachineFromVmClusterResponse")
	}
	return
}

// removeVirtualMachineFromVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) removeVirtualMachineFromVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusters/{vmClusterId}/actions/removeVirtualMachine", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveVirtualMachineFromVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/RemoveVirtualMachineFromVmCluster"
		err = common.PostProcessServiceError(err, "Database", "RemoveVirtualMachineFromVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResizeVmClusterNetwork Adds or removes Db server network nodes to extend or shrink the existing VM cluster network. Applies to Exadata
// Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ResizeVmClusterNetwork.go.html to see an example of how to use ResizeVmClusterNetwork API.
func (client DatabaseClient) ResizeVmClusterNetwork(ctx context.Context, request ResizeVmClusterNetworkRequest) (response ResizeVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resizeVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResizeVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResizeVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResizeVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResizeVmClusterNetworkResponse")
	}
	return
}

// resizeVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) resizeVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}/actions/resize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResizeVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/ResizeVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "ResizeVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResourcePoolShapes Lists available resource pools shapes.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ResourcePoolShapes.go.html to see an example of how to use ResourcePoolShapes API.
func (client DatabaseClient) ResourcePoolShapes(ctx context.Context, request ResourcePoolShapesRequest) (response ResourcePoolShapesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resourcePoolShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResourcePoolShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResourcePoolShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResourcePoolShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResourcePoolShapesResponse")
	}
	return
}

// resourcePoolShapes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) resourcePoolShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/actions/listResourcePoolShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResourcePoolShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ResourcePoolShapes"
		err = common.PostProcessServiceError(err, "Database", "ResourcePoolShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartAutonomousContainerDatabase Rolling restarts the specified Autonomous Container Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RestartAutonomousContainerDatabase.go.html to see an example of how to use RestartAutonomousContainerDatabase API.
func (client DatabaseClient) RestartAutonomousContainerDatabase(ctx context.Context, request RestartAutonomousContainerDatabaseRequest) (response RestartAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restartAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartAutonomousContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartAutonomousContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartAutonomousContainerDatabaseResponse")
	}
	return
}

// restartAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restartAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/restart", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestartAutonomousContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/RestartAutonomousContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "RestartAutonomousContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartAutonomousDatabase Restarts the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RestartAutonomousDatabase.go.html to see an example of how to use RestartAutonomousDatabase API.
func (client DatabaseClient) RestartAutonomousDatabase(ctx context.Context, request RestartAutonomousDatabaseRequest) (response RestartAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restartAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartAutonomousDatabaseResponse")
	}
	return
}

// restartAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restartAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/restart", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestartAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/RestartAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "RestartAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreAutonomousDatabase Restores an Autonomous Database based on the provided request parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RestoreAutonomousDatabase.go.html to see an example of how to use RestoreAutonomousDatabase API.
func (client DatabaseClient) RestoreAutonomousDatabase(ctx context.Context, request RestoreAutonomousDatabaseRequest) (response RestoreAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreAutonomousDatabaseResponse")
	}
	return
}

// restoreAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restoreAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/restore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/RestoreAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "RestoreAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreDatabase Restore a Database based on the request parameters you provide.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RestoreDatabase.go.html to see an example of how to use RestoreDatabase API.
func (client DatabaseClient) RestoreDatabase(ctx context.Context, request RestoreDatabaseRequest) (response RestoreDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreDatabaseResponse")
	}
	return
}

// restoreDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) restoreDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/restore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/RestoreDatabase"
		err = common.PostProcessServiceError(err, "Database", "RestoreDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateAutonomousContainerDatabaseEncryptionKey Creates a new version of an existing Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateAutonomousContainerDatabaseEncryptionKey.go.html to see an example of how to use RotateAutonomousContainerDatabaseEncryptionKey API.
func (client DatabaseClient) RotateAutonomousContainerDatabaseEncryptionKey(ctx context.Context, request RotateAutonomousContainerDatabaseEncryptionKeyRequest) (response RotateAutonomousContainerDatabaseEncryptionKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateAutonomousContainerDatabaseEncryptionKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateAutonomousContainerDatabaseEncryptionKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateAutonomousContainerDatabaseEncryptionKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateAutonomousContainerDatabaseEncryptionKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateAutonomousContainerDatabaseEncryptionKeyResponse")
	}
	return
}

// rotateAutonomousContainerDatabaseEncryptionKey implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateAutonomousContainerDatabaseEncryptionKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/rotateKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateAutonomousContainerDatabaseEncryptionKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/RotateAutonomousContainerDatabaseEncryptionKey"
		err = common.PostProcessServiceError(err, "Database", "RotateAutonomousContainerDatabaseEncryptionKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateAutonomousDatabaseEncryptionKey Rotate existing AutonomousDatabase Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateAutonomousDatabaseEncryptionKey.go.html to see an example of how to use RotateAutonomousDatabaseEncryptionKey API.
func (client DatabaseClient) RotateAutonomousDatabaseEncryptionKey(ctx context.Context, request RotateAutonomousDatabaseEncryptionKeyRequest) (response RotateAutonomousDatabaseEncryptionKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateAutonomousDatabaseEncryptionKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateAutonomousDatabaseEncryptionKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateAutonomousDatabaseEncryptionKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateAutonomousDatabaseEncryptionKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateAutonomousDatabaseEncryptionKeyResponse")
	}
	return
}

// rotateAutonomousDatabaseEncryptionKey implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateAutonomousDatabaseEncryptionKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/rotateKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateAutonomousDatabaseEncryptionKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/RotateAutonomousDatabaseEncryptionKey"
		err = common.PostProcessServiceError(err, "Database", "RotateAutonomousDatabaseEncryptionKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateAutonomousVmClusterOrdsCerts Rotates the Oracle REST Data Services (ORDS) certificates for Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateAutonomousVmClusterOrdsCerts.go.html to see an example of how to use RotateAutonomousVmClusterOrdsCerts API.
func (client DatabaseClient) RotateAutonomousVmClusterOrdsCerts(ctx context.Context, request RotateAutonomousVmClusterOrdsCertsRequest) (response RotateAutonomousVmClusterOrdsCertsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateAutonomousVmClusterOrdsCerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateAutonomousVmClusterOrdsCertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateAutonomousVmClusterOrdsCertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateAutonomousVmClusterOrdsCertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateAutonomousVmClusterOrdsCertsResponse")
	}
	return
}

// rotateAutonomousVmClusterOrdsCerts implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateAutonomousVmClusterOrdsCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousVmClusters/{autonomousVmClusterId}/actions/rotateOrdsCerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateAutonomousVmClusterOrdsCertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/RotateAutonomousVmClusterOrdsCerts"
		err = common.PostProcessServiceError(err, "Database", "RotateAutonomousVmClusterOrdsCerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateAutonomousVmClusterSslCerts Rotates the SSL certificates for Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateAutonomousVmClusterSslCerts.go.html to see an example of how to use RotateAutonomousVmClusterSslCerts API.
func (client DatabaseClient) RotateAutonomousVmClusterSslCerts(ctx context.Context, request RotateAutonomousVmClusterSslCertsRequest) (response RotateAutonomousVmClusterSslCertsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateAutonomousVmClusterSslCerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateAutonomousVmClusterSslCertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateAutonomousVmClusterSslCertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateAutonomousVmClusterSslCertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateAutonomousVmClusterSslCertsResponse")
	}
	return
}

// rotateAutonomousVmClusterSslCerts implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateAutonomousVmClusterSslCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousVmClusters/{autonomousVmClusterId}/actions/rotateSslCerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateAutonomousVmClusterSslCertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/RotateAutonomousVmClusterSslCerts"
		err = common.PostProcessServiceError(err, "Database", "RotateAutonomousVmClusterSslCerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateCloudAutonomousVmClusterOrdsCerts Rotates the Oracle REST Data Services (ORDS) certificates for a cloud Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateCloudAutonomousVmClusterOrdsCerts.go.html to see an example of how to use RotateCloudAutonomousVmClusterOrdsCerts API.
func (client DatabaseClient) RotateCloudAutonomousVmClusterOrdsCerts(ctx context.Context, request RotateCloudAutonomousVmClusterOrdsCertsRequest) (response RotateCloudAutonomousVmClusterOrdsCertsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateCloudAutonomousVmClusterOrdsCerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateCloudAutonomousVmClusterOrdsCertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateCloudAutonomousVmClusterOrdsCertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateCloudAutonomousVmClusterOrdsCertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateCloudAutonomousVmClusterOrdsCertsResponse")
	}
	return
}

// rotateCloudAutonomousVmClusterOrdsCerts implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateCloudAutonomousVmClusterOrdsCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}/actions/rotateOrdsCerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateCloudAutonomousVmClusterOrdsCertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/RotateCloudAutonomousVmClusterOrdsCerts"
		err = common.PostProcessServiceError(err, "Database", "RotateCloudAutonomousVmClusterOrdsCerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateCloudAutonomousVmClusterSslCerts Rotates the SSL certficates for a cloud Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateCloudAutonomousVmClusterSslCerts.go.html to see an example of how to use RotateCloudAutonomousVmClusterSslCerts API.
func (client DatabaseClient) RotateCloudAutonomousVmClusterSslCerts(ctx context.Context, request RotateCloudAutonomousVmClusterSslCertsRequest) (response RotateCloudAutonomousVmClusterSslCertsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateCloudAutonomousVmClusterSslCerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateCloudAutonomousVmClusterSslCertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateCloudAutonomousVmClusterSslCertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateCloudAutonomousVmClusterSslCertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateCloudAutonomousVmClusterSslCertsResponse")
	}
	return
}

// rotateCloudAutonomousVmClusterSslCerts implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateCloudAutonomousVmClusterSslCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}/actions/rotateSslCerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateCloudAutonomousVmClusterSslCertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/RotateCloudAutonomousVmClusterSslCerts"
		err = common.PostProcessServiceError(err, "Database", "RotateCloudAutonomousVmClusterSslCerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateOrdsCerts **Deprecated.** Use the RotateCloudAutonomousVmClusterOrdsCerts to rotate Oracle REST Data Services (ORDS) certs for an Autonomous Exadata VM cluster instead.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateOrdsCerts.go.html to see an example of how to use RotateOrdsCerts API.
func (client DatabaseClient) RotateOrdsCerts(ctx context.Context, request RotateOrdsCertsRequest) (response RotateOrdsCertsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateOrdsCerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateOrdsCertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateOrdsCertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateOrdsCertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateOrdsCertsResponse")
	}
	return
}

// rotateOrdsCerts implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateOrdsCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/actions/rotateOrdsCerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateOrdsCertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/RotateOrdsCerts"
		err = common.PostProcessServiceError(err, "Database", "RotateOrdsCerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotatePluggableDatabaseEncryptionKey Create a new version of the existing encryption key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotatePluggableDatabaseEncryptionKey.go.html to see an example of how to use RotatePluggableDatabaseEncryptionKey API.
func (client DatabaseClient) RotatePluggableDatabaseEncryptionKey(ctx context.Context, request RotatePluggableDatabaseEncryptionKeyRequest) (response RotatePluggableDatabaseEncryptionKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotatePluggableDatabaseEncryptionKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotatePluggableDatabaseEncryptionKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotatePluggableDatabaseEncryptionKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotatePluggableDatabaseEncryptionKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotatePluggableDatabaseEncryptionKeyResponse")
	}
	return
}

// rotatePluggableDatabaseEncryptionKey implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotatePluggableDatabaseEncryptionKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/rotateKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotatePluggableDatabaseEncryptionKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/RotatePluggableDatabaseEncryptionKey"
		err = common.PostProcessServiceError(err, "Database", "RotatePluggableDatabaseEncryptionKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateSslCerts **Deprecated.** Use the RotateCloudAutonomousVmClusterSslCerts to rotate SSL certs for an Autonomous Exadata VM cluster instead.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateSslCerts.go.html to see an example of how to use RotateSslCerts API.
func (client DatabaseClient) RotateSslCerts(ctx context.Context, request RotateSslCertsRequest) (response RotateSslCertsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateSslCerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateSslCertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateSslCertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateSslCertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateSslCertsResponse")
	}
	return
}

// rotateSslCerts implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateSslCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/actions/rotateSslCerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateSslCertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/RotateSslCerts"
		err = common.PostProcessServiceError(err, "Database", "RotateSslCerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateVaultKey Creates a new version of an existing Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/RotateVaultKey.go.html to see an example of how to use RotateVaultKey API.
func (client DatabaseClient) RotateVaultKey(ctx context.Context, request RotateVaultKeyRequest) (response RotateVaultKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateVaultKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateVaultKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateVaultKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateVaultKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateVaultKeyResponse")
	}
	return
}

// rotateVaultKey implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) rotateVaultKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/rotateKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateVaultKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/RotateVaultKey"
		err = common.PostProcessServiceError(err, "Database", "RotateVaultKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SaasAdminUserStatus This operation gets SaaS administrative user status of the Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/SaasAdminUserStatus.go.html to see an example of how to use SaasAdminUserStatus API.
func (client DatabaseClient) SaasAdminUserStatus(ctx context.Context, request SaasAdminUserStatusRequest) (response SaasAdminUserStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.saasAdminUserStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SaasAdminUserStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SaasAdminUserStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SaasAdminUserStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SaasAdminUserStatusResponse")
	}
	return
}

// saasAdminUserStatus implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) saasAdminUserStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/getSaasAdminUserStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SaasAdminUserStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/SaasAdminUserStatus"
		err = common.PostProcessServiceError(err, "Database", "SaasAdminUserStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScanExternalContainerDatabasePluggableDatabases Scans for pluggable databases in the specified external container database.
// This operation will return un-registered pluggable databases in the GetWorkRequest operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ScanExternalContainerDatabasePluggableDatabases.go.html to see an example of how to use ScanExternalContainerDatabasePluggableDatabases API.
func (client DatabaseClient) ScanExternalContainerDatabasePluggableDatabases(ctx context.Context, request ScanExternalContainerDatabasePluggableDatabasesRequest) (response ScanExternalContainerDatabasePluggableDatabasesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scanExternalContainerDatabasePluggableDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScanExternalContainerDatabasePluggableDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScanExternalContainerDatabasePluggableDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScanExternalContainerDatabasePluggableDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScanExternalContainerDatabasePluggableDatabasesResponse")
	}
	return
}

// scanExternalContainerDatabasePluggableDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) scanExternalContainerDatabasePluggableDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/scanPluggableDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScanExternalContainerDatabasePluggableDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/ScanExternalContainerDatabasePluggableDatabases"
		err = common.PostProcessServiceError(err, "Database", "ScanExternalContainerDatabasePluggableDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ShrinkAutonomousDatabase This operation shrinks the current allocated storage down to the current actual used data storage (actualUsedDataStorageSizeInTBs). The if the base storage value for the database (dataStorageSizeInTBs) is larger than the actualUsedDataStorageSizeInTBs value, you are billed for the base storage value.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ShrinkAutonomousDatabase.go.html to see an example of how to use ShrinkAutonomousDatabase API.
func (client DatabaseClient) ShrinkAutonomousDatabase(ctx context.Context, request ShrinkAutonomousDatabaseRequest) (response ShrinkAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.shrinkAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ShrinkAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ShrinkAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ShrinkAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ShrinkAutonomousDatabaseResponse")
	}
	return
}

// shrinkAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) shrinkAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/shrink", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ShrinkAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/ShrinkAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "ShrinkAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartAutonomousDatabase Starts the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/StartAutonomousDatabase.go.html to see an example of how to use StartAutonomousDatabase API.
func (client DatabaseClient) StartAutonomousDatabase(ctx context.Context, request StartAutonomousDatabaseRequest) (response StartAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAutonomousDatabaseResponse")
	}
	return
}

// startAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) startAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/StartAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "StartAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartPluggableDatabase Starts a stopped pluggable database. The `openMode` value of the pluggable database will be `READ_WRITE` upon completion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/StartPluggableDatabase.go.html to see an example of how to use StartPluggableDatabase API.
func (client DatabaseClient) StartPluggableDatabase(ctx context.Context, request StartPluggableDatabaseRequest) (response StartPluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartPluggableDatabaseResponse")
	}
	return
}

// startPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) startPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/StartPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "StartPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopAutonomousDatabase Stops the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/StopAutonomousDatabase.go.html to see an example of how to use StopAutonomousDatabase API.
func (client DatabaseClient) StopAutonomousDatabase(ctx context.Context, request StopAutonomousDatabaseRequest) (response StopAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAutonomousDatabaseResponse")
	}
	return
}

// stopAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) stopAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/StopAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "StopAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopPluggableDatabase Stops a pluggable database. The `openMode` value of the pluggable database will be `MOUNTED` upon completion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/StopPluggableDatabase.go.html to see an example of how to use StopPluggableDatabase API.
func (client DatabaseClient) StopPluggableDatabase(ctx context.Context, request StopPluggableDatabaseRequest) (response StopPluggableDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopPluggableDatabaseResponse")
	}
	return
}

// stopPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) stopPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pluggableDatabases/{pluggableDatabaseId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/StopPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "StopPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchoverAutonomousContainerDatabaseDataguardAssociation Switches over the primary Autonomous Container Database of an Autonomous Data Guard peer association to standby role. The standby Autonomous Container Database associated with autonomousContainerDatabaseDataguardAssociationId assumes the primary Autonomous Container Database role.
// A switchover incurs no data loss.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/SwitchoverAutonomousContainerDatabaseDataguardAssociation.go.html to see an example of how to use SwitchoverAutonomousContainerDatabaseDataguardAssociation API.
func (client DatabaseClient) SwitchoverAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request SwitchoverAutonomousContainerDatabaseDataguardAssociationRequest) (response SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.switchoverAutonomousContainerDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse")
	}
	return
}

// switchoverAutonomousContainerDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) switchoverAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}/actions/switchover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/SwitchoverAutonomousContainerDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "SwitchoverAutonomousContainerDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchoverAutonomousDatabase Initiates a switchover of the specified Autonomous Database to the associated peer database. Applicable only to databases with Disaster Recovery enabled.
// This API should be called in the remote region where the peer database resides.
// Below parameter is optional:
//   - `peerDbId`
//     Use this parameter to specify the database OCID of the Disaster Recovery peer, which is located in a different (remote) region from the current peer database.
//     If this parameter is not provided, the switchover will happen in the same region.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/SwitchoverAutonomousDatabase.go.html to see an example of how to use SwitchoverAutonomousDatabase API.
func (client DatabaseClient) SwitchoverAutonomousDatabase(ctx context.Context, request SwitchoverAutonomousDatabaseRequest) (response SwitchoverAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.switchoverAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchoverAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchoverAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverAutonomousDatabaseResponse")
	}
	return
}

// switchoverAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) switchoverAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/switchover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchoverAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/SwitchoverAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "SwitchoverAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchoverDataGuardAssociation Performs a switchover to transition the primary database of a Data Guard association into a standby role. The
// standby database associated with the `dataGuardAssociationId` assumes the primary database role.
// A switchover guarantees no data loss.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/SwitchoverDataGuardAssociation.go.html to see an example of how to use SwitchoverDataGuardAssociation API.
func (client DatabaseClient) SwitchoverDataGuardAssociation(ctx context.Context, request SwitchoverDataGuardAssociationRequest) (response SwitchoverDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.switchoverDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchoverDataGuardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchoverDataGuardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverDataGuardAssociationResponse")
	}
	return
}

// switchoverDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) switchoverDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/switchover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchoverDataGuardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/SwitchoverDataGuardAssociation"
		err = common.PostProcessServiceError(err, "Database", "SwitchoverDataGuardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TerminateAutonomousContainerDatabase Terminates an Autonomous Container Database, which permanently deletes the container database and any databases within the container database. The database data is local to the Autonomous Exadata Infrastructure and will be lost when the container database is terminated. Oracle recommends that you back up any data in the Autonomous Container Database prior to terminating it.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/TerminateAutonomousContainerDatabase.go.html to see an example of how to use TerminateAutonomousContainerDatabase API.
func (client DatabaseClient) TerminateAutonomousContainerDatabase(ctx context.Context, request TerminateAutonomousContainerDatabaseRequest) (response TerminateAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TerminateAutonomousContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TerminateAutonomousContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateAutonomousContainerDatabaseResponse")
	}
	return
}

// terminateAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/TerminateAutonomousContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "TerminateAutonomousContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TerminateAutonomousExadataInfrastructure **Deprecated.** To terminate an Exadata Infrastructure resource in the Oracle cloud, use the DeleteCloudExadataInfrastructure operation. To delete an Autonomous Exadata VM cluster in the Oracle cloud, use the DeleteCloudAutonomousVmCluster operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/TerminateAutonomousExadataInfrastructure.go.html to see an example of how to use TerminateAutonomousExadataInfrastructure API.
func (client DatabaseClient) TerminateAutonomousExadataInfrastructure(ctx context.Context, request TerminateAutonomousExadataInfrastructureRequest) (response TerminateAutonomousExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateAutonomousExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TerminateAutonomousExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TerminateAutonomousExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateAutonomousExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateAutonomousExadataInfrastructureResponse")
	}
	return
}

// terminateAutonomousExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/TerminateAutonomousExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "TerminateAutonomousExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TerminateDbSystem Terminates a DB system and permanently deletes it and any databases running on it, and any storage volumes attached to it. The database data is local to the DB system and will be lost when the system is terminated. Oracle recommends that you back up any data in the DB system prior to terminating it.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/TerminateDbSystem.go.html to see an example of how to use TerminateDbSystem API.
func (client DatabaseClient) TerminateDbSystem(ctx context.Context, request TerminateDbSystemRequest) (response TerminateDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TerminateDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TerminateDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateDbSystemResponse")
	}
	return
}

// terminateDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) terminateDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbSystems/{dbSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TerminateDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/TerminateDbSystem"
		err = common.PostProcessServiceError(err, "Database", "TerminateDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousContainerDatabase Updates the properties of an Autonomous Container Database, such as display name, maintenance preference, backup retention, and tags.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousContainerDatabase.go.html to see an example of how to use UpdateAutonomousContainerDatabase API.
func (client DatabaseClient) UpdateAutonomousContainerDatabase(ctx context.Context, request UpdateAutonomousContainerDatabaseRequest) (response UpdateAutonomousContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousContainerDatabaseResponse")
	}
	return
}

// updateAutonomousContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabase/UpdateAutonomousContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousContainerDatabaseDataguardAssociation Update Autonomous Data Guard association.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousContainerDatabaseDataguardAssociation.go.html to see an example of how to use UpdateAutonomousContainerDatabaseDataguardAssociation API.
func (client DatabaseClient) UpdateAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request UpdateAutonomousContainerDatabaseDataguardAssociationRequest) (response UpdateAutonomousContainerDatabaseDataguardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousContainerDatabaseDataguardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousContainerDatabaseDataguardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousContainerDatabaseDataguardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousContainerDatabaseDataguardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousContainerDatabaseDataguardAssociationResponse")
	}
	return
}

// updateAutonomousContainerDatabaseDataguardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousContainerDatabaseDataguardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousContainerDatabaseDataguardAssociation/UpdateAutonomousContainerDatabaseDataguardAssociation"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousContainerDatabaseDataguardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousDatabase Updates one or more attributes of the specified Autonomous Database. See the UpdateAutonomousDatabaseDetails resource for a full list of attributes that can be updated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousDatabase.go.html to see an example of how to use UpdateAutonomousDatabase API.
func (client DatabaseClient) UpdateAutonomousDatabase(ctx context.Context, request UpdateAutonomousDatabaseRequest) (response UpdateAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDatabaseResponse")
	}
	return
}

// updateAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/{autonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabase/UpdateAutonomousDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousDatabaseBackup Updates the Autonomous Database backup of the specified database based on the request parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousDatabaseBackup.go.html to see an example of how to use UpdateAutonomousDatabaseBackup API.
func (client DatabaseClient) UpdateAutonomousDatabaseBackup(ctx context.Context, request UpdateAutonomousDatabaseBackupRequest) (response UpdateAutonomousDatabaseBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDatabaseBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousDatabaseBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousDatabaseBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDatabaseBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDatabaseBackupResponse")
	}
	return
}

// updateAutonomousDatabaseBackup implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabaseBackups/{autonomousDatabaseBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseBackup/UpdateAutonomousDatabaseBackup"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousDatabaseBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousDatabaseRegionalWallet Updates the Autonomous Database regional wallet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousDatabaseRegionalWallet.go.html to see an example of how to use UpdateAutonomousDatabaseRegionalWallet API.
func (client DatabaseClient) UpdateAutonomousDatabaseRegionalWallet(ctx context.Context, request UpdateAutonomousDatabaseRegionalWalletRequest) (response UpdateAutonomousDatabaseRegionalWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDatabaseRegionalWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousDatabaseRegionalWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousDatabaseRegionalWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDatabaseRegionalWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDatabaseRegionalWalletResponse")
	}
	return
}

// updateAutonomousDatabaseRegionalWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDatabaseRegionalWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/wallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseRegionalWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseWallet/UpdateAutonomousDatabaseRegionalWallet"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousDatabaseRegionalWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousDatabaseWallet Updates the wallet for the specified Autonomous Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousDatabaseWallet.go.html to see an example of how to use UpdateAutonomousDatabaseWallet API.
func (client DatabaseClient) UpdateAutonomousDatabaseWallet(ctx context.Context, request UpdateAutonomousDatabaseWalletRequest) (response UpdateAutonomousDatabaseWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousDatabaseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousDatabaseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousDatabaseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousDatabaseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousDatabaseWalletResponse")
	}
	return
}

// updateAutonomousDatabaseWallet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/{autonomousDatabaseId}/wallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousDatabaseWallet/UpdateAutonomousDatabaseWallet"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousDatabaseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousExadataInfrastructure **Deprecated.** Use the UpdateCloudExadataInfrastructure operation to update an Exadata Infrastructure resource and  UpdateCloudAutonomousVmCluster operation to update an Autonomous Exadata VM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousExadataInfrastructure.go.html to see an example of how to use UpdateAutonomousExadataInfrastructure API.
func (client DatabaseClient) UpdateAutonomousExadataInfrastructure(ctx context.Context, request UpdateAutonomousExadataInfrastructureRequest) (response UpdateAutonomousExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousExadataInfrastructureResponse")
	}
	return
}

// updateAutonomousExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousExadataInfrastructure/UpdateAutonomousExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutonomousVmCluster Updates the specified Autonomous VM cluster for the Exadata Cloud@Customer system.To update an Autonomous VM Cluster in the Oracle cloud, see UpdateCloudAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateAutonomousVmCluster.go.html to see an example of how to use UpdateAutonomousVmCluster API.
func (client DatabaseClient) UpdateAutonomousVmCluster(ctx context.Context, request UpdateAutonomousVmClusterRequest) (response UpdateAutonomousVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutonomousVmClusterResponse")
	}
	return
}

// updateAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousVmClusters/{autonomousVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/AutonomousVmCluster/UpdateAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "UpdateAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBackupDestination If no database is associated with the backup destination:
// - For a RECOVERY_APPLIANCE backup destination, updates the connection string and/or the list of VPC users.
// - For an NFS backup destination, updates the NFS location.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateBackupDestination.go.html to see an example of how to use UpdateBackupDestination API.
func (client DatabaseClient) UpdateBackupDestination(ctx context.Context, request UpdateBackupDestinationRequest) (response UpdateBackupDestinationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBackupDestination, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBackupDestinationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBackupDestinationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBackupDestinationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBackupDestinationResponse")
	}
	return
}

// updateBackupDestination implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/backupDestinations/{backupDestinationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBackupDestinationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/BackupDestination/UpdateBackupDestination"
		err = common.PostProcessServiceError(err, "Database", "UpdateBackupDestination", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCloudAutonomousVmCluster Updates the specified Autonomous Exadata VM cluster in the Oracle cloud. For Exadata Cloud@Customer systems, see UpdateAutonomousVmCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateCloudAutonomousVmCluster.go.html to see an example of how to use UpdateCloudAutonomousVmCluster API.
func (client DatabaseClient) UpdateCloudAutonomousVmCluster(ctx context.Context, request UpdateCloudAutonomousVmClusterRequest) (response UpdateCloudAutonomousVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCloudAutonomousVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCloudAutonomousVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCloudAutonomousVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCloudAutonomousVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCloudAutonomousVmClusterResponse")
	}
	return
}

// updateCloudAutonomousVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateCloudAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudAutonomousVmClusters/{cloudAutonomousVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudAutonomousVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudAutonomousVmCluster/UpdateCloudAutonomousVmCluster"
		err = common.PostProcessServiceError(err, "Database", "UpdateCloudAutonomousVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCloudExadataInfrastructure Updates the Cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateCloudExadataInfrastructure.go.html to see an example of how to use UpdateCloudExadataInfrastructure API.
func (client DatabaseClient) UpdateCloudExadataInfrastructure(ctx context.Context, request UpdateCloudExadataInfrastructureRequest) (response UpdateCloudExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCloudExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCloudExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCloudExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCloudExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCloudExadataInfrastructureResponse")
	}
	return
}

// updateCloudExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudExadataInfrastructure/UpdateCloudExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "UpdateCloudExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCloudVmCluster Updates the specified cloud VM cluster. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateCloudVmCluster.go.html to see an example of how to use UpdateCloudVmCluster API.
func (client DatabaseClient) UpdateCloudVmCluster(ctx context.Context, request UpdateCloudVmClusterRequest) (response UpdateCloudVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCloudVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCloudVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCloudVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCloudVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCloudVmClusterResponse")
	}
	return
}

// updateCloudVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudVmClusters/{cloudVmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/UpdateCloudVmCluster"
		err = common.PostProcessServiceError(err, "Database", "UpdateCloudVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCloudVmClusterIormConfig Updates the IORM settings for the specified cloud VM cluster in an Exadata Cloud Service instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateCloudVmClusterIormConfig.go.html to see an example of how to use UpdateCloudVmClusterIormConfig API.
func (client DatabaseClient) UpdateCloudVmClusterIormConfig(ctx context.Context, request UpdateCloudVmClusterIormConfigRequest) (response UpdateCloudVmClusterIormConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCloudVmClusterIormConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCloudVmClusterIormConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCloudVmClusterIormConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCloudVmClusterIormConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCloudVmClusterIormConfigResponse")
	}
	return
}

// updateCloudVmClusterIormConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateCloudVmClusterIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudVmClusters/{cloudVmClusterId}/CloudVmClusterIormConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudVmClusterIormConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/CloudVmCluster/UpdateCloudVmClusterIormConfig"
		err = common.PostProcessServiceError(err, "Database", "UpdateCloudVmClusterIormConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConsoleConnection Updates the specified database node console connection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateConsoleConnection.go.html to see an example of how to use UpdateConsoleConnection API.
func (client DatabaseClient) UpdateConsoleConnection(ctx context.Context, request UpdateConsoleConnectionRequest) (response UpdateConsoleConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConsoleConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConsoleConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConsoleConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConsoleConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConsoleConnectionResponse")
	}
	return
}

// updateConsoleConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbNodes/{dbNodeId}/consoleConnections/{consoleConnectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConsoleConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleConnection/UpdateConsoleConnection"
		err = common.PostProcessServiceError(err, "Database", "UpdateConsoleConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConsoleHistory Updates the specified database node console history.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateConsoleHistory.go.html to see an example of how to use UpdateConsoleHistory API.
func (client DatabaseClient) UpdateConsoleHistory(ctx context.Context, request UpdateConsoleHistoryRequest) (response UpdateConsoleHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConsoleHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConsoleHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConsoleHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConsoleHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConsoleHistoryResponse")
	}
	return
}

// updateConsoleHistory implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateConsoleHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbNodes/{dbNodeId}/consoleHistories/{consoleHistoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConsoleHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ConsoleHistory/UpdateConsoleHistory"
		err = common.PostProcessServiceError(err, "Database", "UpdateConsoleHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataGuardAssociation Updates the Data Guard association the specified database. This API can be used to change the `protectionMode` and `transportType` of the Data Guard association.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateDataGuardAssociation.go.html to see an example of how to use UpdateDataGuardAssociation API.
func (client DatabaseClient) UpdateDataGuardAssociation(ctx context.Context, request UpdateDataGuardAssociationRequest) (response UpdateDataGuardAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataGuardAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataGuardAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataGuardAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataGuardAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataGuardAssociationResponse")
	}
	return
}

// updateDataGuardAssociation implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataGuardAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DataGuardAssociation/UpdateDataGuardAssociation"
		err = common.PostProcessServiceError(err, "Database", "UpdateDataGuardAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabase Update the specified database based on the request parameters provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateDatabase.go.html to see an example of how to use UpdateDatabase API.
func (client DatabaseClient) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (response UpdateDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseResponse")
	}
	return
}

// updateDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databases/{databaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/UpdateDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdateDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseSoftwareImage Updates the properties of a Database Software Image, like Display Nmae
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateDatabaseSoftwareImage.go.html to see an example of how to use UpdateDatabaseSoftwareImage API.
func (client DatabaseClient) UpdateDatabaseSoftwareImage(ctx context.Context, request UpdateDatabaseSoftwareImageRequest) (response UpdateDatabaseSoftwareImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseSoftwareImageResponse")
	}
	return
}

// updateDatabaseSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseSoftwareImages/{databaseSoftwareImageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DatabaseSoftwareImage/UpdateDatabaseSoftwareImage"
		err = common.PostProcessServiceError(err, "Database", "UpdateDatabaseSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbHome Patches the specified Database Home.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateDbHome.go.html to see an example of how to use UpdateDbHome API.
func (client DatabaseClient) UpdateDbHome(ctx context.Context, request UpdateDbHomeRequest) (response UpdateDbHomeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbHome, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbHomeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbHomeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbHomeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbHomeResponse")
	}
	return
}

// updateDbHome implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbHomes/{dbHomeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbHomeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbHome/UpdateDbHome"
		err = common.PostProcessServiceError(err, "Database", "UpdateDbHome", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbNode Updates the specified database node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateDbNode.go.html to see an example of how to use UpdateDbNode API.
func (client DatabaseClient) UpdateDbNode(ctx context.Context, request UpdateDbNodeRequest) (response UpdateDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbNodeResponse")
	}
	return
}

// updateDbNode implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbNodes/{dbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbNode/UpdateDbNode"
		err = common.PostProcessServiceError(err, "Database", "UpdateDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbSystem Updates the properties of the specified DB system.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateDbSystem.go.html to see an example of how to use UpdateDbSystem API.
func (client DatabaseClient) UpdateDbSystem(ctx context.Context, request UpdateDbSystemRequest) (response UpdateDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbSystemResponse")
	}
	return
}

// updateDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/UpdateDbSystem"
		err = common.PostProcessServiceError(err, "Database", "UpdateDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExadataInfrastructure Updates the Exadata infrastructure resource. Applies to Exadata Cloud@Customer instances only.
// To update an Exadata Cloud Service infrastructure resource, use the  UpdateCloudExadataInfrastructure operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateExadataInfrastructure.go.html to see an example of how to use UpdateExadataInfrastructure API.
func (client DatabaseClient) UpdateExadataInfrastructure(ctx context.Context, request UpdateExadataInfrastructureRequest) (response UpdateExadataInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExadataInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExadataInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExadataInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExadataInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExadataInfrastructureResponse")
	}
	return
}

// updateExadataInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exadataInfrastructures/{exadataInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExadataInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExadataInfrastructure/UpdateExadataInfrastructure"
		err = common.PostProcessServiceError(err, "Database", "UpdateExadataInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExadataIormConfig Updates IORM settings for the specified Exadata DB system.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
// The UpdateCloudVmClusterIormConfig API is used for Exadata systems using the
// new resource model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateExadataIormConfig.go.html to see an example of how to use UpdateExadataIormConfig API.
func (client DatabaseClient) UpdateExadataIormConfig(ctx context.Context, request UpdateExadataIormConfigRequest) (response UpdateExadataIormConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExadataIormConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExadataIormConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExadataIormConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExadataIormConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExadataIormConfigResponse")
	}
	return
}

// updateExadataIormConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExadataIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}/ExadataIormConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExadataIormConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/UpdateExadataIormConfig"
		err = common.PostProcessServiceError(err, "Database", "UpdateExadataIormConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalContainerDatabase Updates the properties of
// an CreateExternalContainerDatabaseDetails resource,
// such as the display name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateExternalContainerDatabase.go.html to see an example of how to use UpdateExternalContainerDatabase API.
func (client DatabaseClient) UpdateExternalContainerDatabase(ctx context.Context, request UpdateExternalContainerDatabaseRequest) (response UpdateExternalContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalContainerDatabaseResponse")
	}
	return
}

// updateExternalContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalcontainerdatabases/{externalContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalContainerDatabase/UpdateExternalContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdateExternalContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalDatabaseConnector Updates the properties of an external database connector, such as the display name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateExternalDatabaseConnector.go.html to see an example of how to use UpdateExternalDatabaseConnector API.
func (client DatabaseClient) UpdateExternalDatabaseConnector(ctx context.Context, request UpdateExternalDatabaseConnectorRequest) (response UpdateExternalDatabaseConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalDatabaseConnectorResponse")
	}
	return
}

// updateExternalDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externaldatabaseconnectors/{externalDatabaseConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalDatabaseConnector/UpdateExternalDatabaseConnector"
		err = common.PostProcessServiceError(err, "Database", "UpdateExternalDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldatabaseconnector{})
	return response, err
}

// UpdateExternalNonContainerDatabase Updates the properties of an external non-container database, such as the display name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateExternalNonContainerDatabase.go.html to see an example of how to use UpdateExternalNonContainerDatabase API.
func (client DatabaseClient) UpdateExternalNonContainerDatabase(ctx context.Context, request UpdateExternalNonContainerDatabaseRequest) (response UpdateExternalNonContainerDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalNonContainerDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalNonContainerDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalNonContainerDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalNonContainerDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalNonContainerDatabaseResponse")
	}
	return
}

// updateExternalNonContainerDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalNonContainerDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalNonContainerDatabase/UpdateExternalNonContainerDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdateExternalNonContainerDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalPluggableDatabase Updates the properties of an
// CreateExternalPluggableDatabaseDetails resource,
// such as the display name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateExternalPluggableDatabase.go.html to see an example of how to use UpdateExternalPluggableDatabase API.
func (client DatabaseClient) UpdateExternalPluggableDatabase(ctx context.Context, request UpdateExternalPluggableDatabaseRequest) (response UpdateExternalPluggableDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalPluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalPluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalPluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalPluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalPluggableDatabaseResponse")
	}
	return
}

// updateExternalPluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalpluggabledatabases/{externalPluggableDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalPluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/ExternalPluggableDatabase/UpdateExternalPluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdateExternalPluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateKeyStore If no database is associated with the key store, edit the key store.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateKeyStore.go.html to see an example of how to use UpdateKeyStore API.
func (client DatabaseClient) UpdateKeyStore(ctx context.Context, request UpdateKeyStoreRequest) (response UpdateKeyStoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateKeyStore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateKeyStoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateKeyStoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateKeyStoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateKeyStoreResponse")
	}
	return
}

// updateKeyStore implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/keyStores/{keyStoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateKeyStoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/KeyStore/UpdateKeyStore"
		err = common.PostProcessServiceError(err, "Database", "UpdateKeyStore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMaintenanceRun Updates the properties of a maintenance run, such as the state of a maintenance run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateMaintenanceRun.go.html to see an example of how to use UpdateMaintenanceRun API.
func (client DatabaseClient) UpdateMaintenanceRun(ctx context.Context, request UpdateMaintenanceRunRequest) (response UpdateMaintenanceRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMaintenanceRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMaintenanceRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaintenanceRunResponse")
	}
	return
}

// updateMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maintenanceRuns/{maintenanceRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMaintenanceRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/MaintenanceRun/UpdateMaintenanceRun"
		err = common.PostProcessServiceError(err, "Database", "UpdateMaintenanceRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOneoffPatch Updates the properties of the specified one-off patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateOneoffPatch.go.html to see an example of how to use UpdateOneoffPatch API.
func (client DatabaseClient) UpdateOneoffPatch(ctx context.Context, request UpdateOneoffPatchRequest) (response UpdateOneoffPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOneoffPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOneoffPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOneoffPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOneoffPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOneoffPatchResponse")
	}
	return
}

// updateOneoffPatch implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateOneoffPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oneoffPatches/{oneoffPatchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOneoffPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/OneoffPatch/UpdateOneoffPatch"
		err = common.PostProcessServiceError(err, "Database", "UpdateOneoffPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePluggableDatabase Updates the specified pluggable database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdatePluggableDatabase.go.html to see an example of how to use UpdatePluggableDatabase API.
func (client DatabaseClient) UpdatePluggableDatabase(ctx context.Context, request UpdatePluggableDatabaseRequest) (response UpdatePluggableDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePluggableDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePluggableDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePluggableDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePluggableDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePluggableDatabaseResponse")
	}
	return
}

// updatePluggableDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updatePluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/pluggableDatabases/{pluggableDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePluggableDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/PluggableDatabase/UpdatePluggableDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpdatePluggableDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVmCluster Updates the specified VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateVmCluster.go.html to see an example of how to use UpdateVmCluster API.
func (client DatabaseClient) UpdateVmCluster(ctx context.Context, request UpdateVmClusterRequest) (response UpdateVmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVmClusterResponse")
	}
	return
}

// updateVmCluster implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/vmClusters/{vmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmCluster/UpdateVmCluster"
		err = common.PostProcessServiceError(err, "Database", "UpdateVmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVmClusterNetwork Updates the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To update a cloud VM cluster in an Exadata Cloud Service instance, use the UpdateCloudVmCluster operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateVmClusterNetwork.go.html to see an example of how to use UpdateVmClusterNetwork API.
func (client DatabaseClient) UpdateVmClusterNetwork(ctx context.Context, request UpdateVmClusterNetworkRequest) (response UpdateVmClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVmClusterNetworkResponse")
	}
	return
}

// updateVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) updateVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/UpdateVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "UpdateVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpgradeDatabase Upgrades the specified Oracle Database instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpgradeDatabase.go.html to see an example of how to use UpgradeDatabase API.
func (client DatabaseClient) UpgradeDatabase(ctx context.Context, request UpgradeDatabaseRequest) (response UpgradeDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.upgradeDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpgradeDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpgradeDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpgradeDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpgradeDatabaseResponse")
	}
	return
}

// upgradeDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) upgradeDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/upgrade", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpgradeDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/Database/UpgradeDatabase"
		err = common.PostProcessServiceError(err, "Database", "UpgradeDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpgradeDbSystem Upgrades the operating system and grid infrastructure of the DB system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpgradeDbSystem.go.html to see an example of how to use UpgradeDbSystem API.
func (client DatabaseClient) UpgradeDbSystem(ctx context.Context, request UpgradeDbSystemRequest) (response UpgradeDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.upgradeDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpgradeDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpgradeDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpgradeDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpgradeDbSystemResponse")
	}
	return
}

// upgradeDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) upgradeDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/upgrade", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpgradeDbSystemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/DbSystem/UpgradeDbSystem"
		err = common.PostProcessServiceError(err, "Database", "UpgradeDbSystem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateVmClusterNetwork Validates the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ValidateVmClusterNetwork.go.html to see an example of how to use ValidateVmClusterNetwork API.
func (client DatabaseClient) ValidateVmClusterNetwork(ctx context.Context, request ValidateVmClusterNetworkRequest) (response ValidateVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateVmClusterNetworkResponse")
	}
	return
}

// validateVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client DatabaseClient) validateVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}/actions/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database/20160918/VmClusterNetwork/ValidateVmClusterNetwork"
		err = common.PostProcessServiceError(err, "Database", "ValidateVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
