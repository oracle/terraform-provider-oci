// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v41/common"
	"github.com/oracle/oci-go-sdk/v41/common/auth"
	"net/http"
)

//DatabaseClient a client for Database
type DatabaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseClientWithConfigurationProvider Creates a new default Database client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
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
//  as well as reading the region
func NewDatabaseClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DatabaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateExadataInfrastructure Activates the specified Exadata infrastructure resource. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) activateExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/activate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ActivateExadataInfrastructureResponse
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

// AutonomousDatabaseManualRefresh Initiates a data refresh for an Autonomous Database refreshable clone. Data is refreshed from the source database to the point of a specified timestamp.
//
// See also
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
func (client DatabaseClient) autonomousDatabaseManualRefresh(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/refresh", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AutonomousDatabaseManualRefreshResponse
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

// ChangeAutonomousContainerDatabaseCompartment Move the Autonomous Container Database and its dependent resources to the specified compartment.
// For more information about moving Autonomous Container Databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeAutonomousContainerDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousContainerDatabaseCompartmentResponse
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

// ChangeAutonomousDatabaseCompartment Move the Autonomous Database and its dependent resources to the specified compartment.
// For more information about moving Autonomous Databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeAutonomousDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousDatabaseCompartmentResponse
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

// ChangeAutonomousExadataInfrastructureCompartment Moves the Autonomous Exadata Infrastructure resource and its dependent resources to the specified compartment.
// For more information, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeAutonomousExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousExadataInfrastructureCompartmentResponse
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

// ChangeAutonomousVmClusterCompartment To move an Autonomous VM cluster and its dependent resources to another compartment, use the
// ChangeAutonomousVmClusterCompartment operation.
//
// See also
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
func (client DatabaseClient) changeAutonomousVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousVmClusters/{autonomousVmClusterId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousVmClusterCompartmentResponse
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

// ChangeBackupDestinationCompartment Move the backup destination and its dependent resources to the specified compartment.
// For more information, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeBackupDestinationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backupDestinations/{backupDestinationId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeBackupDestinationCompartmentResponse
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

// ChangeCloudExadataInfrastructureCompartment Moves a cloud Exadata infrastructure resource and its dependent resources to another compartment. Applies to Exadata Cloud Service instances only. For more information about moving resources to a different compartment, see Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeCloudExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeCloudExadataInfrastructureCompartmentResponse
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

// ChangeCloudVmClusterCompartment Moves a cloud VM cluster and its dependent resources to another compartment. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) changeCloudVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudVmClusters/{cloudVmClusterId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeCloudVmClusterCompartmentResponse
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

// ChangeDatabaseSoftwareImageCompartment Move the Database Software Image and its dependent resources to the specified compartment.
// For more information about moving Databse Software Images, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeDatabaseSoftwareImageCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseSoftwareImages/{databaseSoftwareImageId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseSoftwareImageCompartmentResponse
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

// ChangeDbSystemCompartment Moves the DB system and its dependent resources to the specified compartment.
// For more information about moving DB systems, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeDbSystemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeDbSystemCompartmentResponse
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

// ChangeExadataInfrastructureCompartment Moves an Exadata infrastructure resource and its dependent resources to another compartment. Applies to Exadata Cloud@Customer instances only.
// To move an Exadata Cloud Service infrastructure resource to another compartment, use the  ChangeCloudExadataInfrastructureCompartment operation.
//
// See also
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
func (client DatabaseClient) changeExadataInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeExadataInfrastructureCompartmentResponse
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

// ChangeExternalContainerDatabaseCompartment Move the CreateExternalContainerDatabaseDetails
// and its dependent resources to the specified compartment.
// For more information about moving external container databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeExternalContainerDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalContainerDatabaseCompartmentResponse
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

// ChangeExternalNonContainerDatabaseCompartment Move the external non-container database and its dependent resources to the specified compartment.
// For more information about moving external non-container databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeExternalNonContainerDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalNonContainerDatabaseCompartmentResponse
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

// ChangeExternalPluggableDatabaseCompartment Move the CreateExternalPluggableDatabaseDetails and
// its dependent resources to the specified compartment.
// For more information about moving external pluggable databases, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeExternalPluggableDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeExternalPluggableDatabaseCompartmentResponse
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

// ChangeKeyStoreCompartment Move the key store resource to the specified compartment.
// For more information about moving key stores, see
// Moving Database Resources to a Different Compartment (https://docs.cloud.oracle.com/Content/Database/Concepts/databaseoverview.htm#moveRes).
//
// See also
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
func (client DatabaseClient) changeKeyStoreCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keyStores/{keyStoreId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeKeyStoreCompartmentResponse
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

// ChangeVmClusterCompartment Moves a VM cluster and its dependent resources to another compartment. Applies to Exadata Cloud@Customer instances only.
// To move a cloud VM cluster in an Exadata Cloud Service instance to another compartment, use the ChangeCloudVmClusterCompartment operation.
//
// See also
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
func (client DatabaseClient) changeVmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusters/{vmClusterId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeVmClusterCompartmentResponse
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

// CheckExternalDatabaseConnectorConnectionStatus Check the status of the external database connection specified in this connector.
// This operation will refresh the connectionStatus and timeConnectionStatusLastUpdated fields.
//
// See also
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
func (client DatabaseClient) checkExternalDatabaseConnectorConnectionStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externaldatabaseconnectors/{externalDatabaseConnectorId}/actions/checkConnectionStatus", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CheckExternalDatabaseConnectorConnectionStatusResponse
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

// CompleteExternalBackupJob Changes the status of the standalone backup resource to `ACTIVE` after the backup is created from the on-premises database and placed in Oracle Cloud Infrastructure Object Storage.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.cloud.oracle.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
//
// See also
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
func (client DatabaseClient) completeExternalBackupJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalBackupJobs/{backupId}/actions/complete", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CompleteExternalBackupJobResponse
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

// ConfigureAutonomousDatabaseVaultKey Configures the Autonomous Database Vault service key (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
//
// See also
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
func (client DatabaseClient) configureAutonomousDatabaseVaultKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/configureAutonomousDatabaseVaultKey", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ConfigureAutonomousDatabaseVaultKeyResponse
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

// CreateAutonomousContainerDatabase Creates an Autonomous Container Database in the specified Autonomous Exadata Infrastructure.
//
// See also
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
func (client DatabaseClient) createAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousContainerDatabaseResponse
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

// CreateAutonomousDatabase Creates a new Autonomous Database.
//
// See also
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
func (client DatabaseClient) createAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDatabaseResponse
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

// CreateAutonomousDatabaseBackup Creates a new Autonomous Database backup for the specified database based on the provided request parameters.
//
// See also
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
func (client DatabaseClient) createAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabaseBackups", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousDatabaseBackupResponse
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

// CreateAutonomousVmCluster Creates an Autonomous VM cluster for Exadata Cloud@Customer.
//
// See also
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
func (client DatabaseClient) createAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousVmClusters", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateAutonomousVmClusterResponse
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

// CreateBackup Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
//
// See also
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
func (client DatabaseClient) createBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateBackupResponse
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

// CreateBackupDestination Creates a backup destination in an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) createBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backupDestinations", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateBackupDestinationResponse
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

// CreateCloudExadataInfrastructure Creates a cloud Exadata infrastructure resource. This resource is used to create an Exadata Cloud Service (https://docs.cloud.oracle.com/Content/Database/Concepts/exaoverview.htm) instance.
//
// See also
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
func (client DatabaseClient) createCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudExadataInfrastructures", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateCloudExadataInfrastructureResponse
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

// CreateCloudVmCluster Creates a cloud VM cluster.
//
// See also
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
func (client DatabaseClient) createCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cloudVmClusters", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateCloudVmClusterResponse
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

// CreateConsoleConnection Creates a new console connection to the specified database node.
// After the console connection has been created and is available,
// you connect to the console using SSH.
//
// See also
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
func (client DatabaseClient) createConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbNodes/{dbNodeId}/consoleConnections", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateConsoleConnectionResponse
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

// CreateDataGuardAssociation Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
// specified database and a peer database. For more information, see Using Oracle Data Guard (https://docs.cloud.oracle.com/Content/Database/Tasks/usingdataguard.htm).
// All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
// called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
// resource in the Console. For more information, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// See also
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
func (client DatabaseClient) createDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDataGuardAssociationResponse
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

// CreateDatabase Creates a new database in the specified Database Home. If the database version is provided, it must match the version of the Database Home. Applies to Exadata and Exadata Cloud@Customer systems.
//
// See also
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
func (client DatabaseClient) createDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseResponse
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

// CreateDatabaseSoftwareImage create database software image in the specified compartment.
//
// See also
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
func (client DatabaseClient) createDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseSoftwareImages", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseSoftwareImageResponse
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

// CreateDbHome Creates a new Database Home in the specified database system based on the request parameters you provide. Applies to bare metal DB systems, Exadata systems, and Exadata Cloud@Customer systems.
//
// See also
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
func (client DatabaseClient) createDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbHomes", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDbHomeResponse
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

// CreateExadataInfrastructure Creates an Exadata infrastructure resource. Applies to Exadata Cloud@Customer instances only.
// To create an Exadata Cloud Service infrastructure resource, use the  CreateCloudExadataInfrastructure operation.
//
// See also
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
func (client DatabaseClient) createExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateExadataInfrastructureResponse
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

// CreateExternalBackupJob Creates a new backup resource and returns the information the caller needs to back up an on-premises Oracle Database to Oracle Cloud Infrastructure.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.cloud.oracle.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
//
// See also
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
func (client DatabaseClient) createExternalBackupJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalBackupJobs", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateExternalBackupJobResponse
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

// CreateExternalContainerDatabase Creates a new external container database resource.
//
// See also
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
func (client DatabaseClient) createExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateExternalContainerDatabaseResponse
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

// CreateExternalDatabaseConnector Creates a new external database connector.
//
// See also
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
func (client DatabaseClient) createExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externaldatabaseconnectors", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldatabaseconnector{})
	return response, err
}

// CreateExternalNonContainerDatabase Creates a new ExternalNonContainerDatabase resource
//
// See also
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
func (client DatabaseClient) createExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateExternalNonContainerDatabaseResponse
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

// CreateExternalPluggableDatabase Registers a new CreateExternalPluggableDatabaseDetails
// resource.
//
// See also
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
func (client DatabaseClient) createExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateExternalPluggableDatabaseResponse
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

// CreateKeyStore Creates a Key Store.
//
// See also
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
func (client DatabaseClient) createKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keyStores", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateKeyStoreResponse
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

// CreateVmCluster Creates an Exadata Cloud@Customer VM cluster.
//
// See also
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
func (client DatabaseClient) createVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusters", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateVmClusterResponse
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

// CreateVmClusterNetwork Creates the VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To create a cloud VM cluster in an Exadata Cloud Service instance, use the CreateCloudVmCluster operation.
//
// See also
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
func (client DatabaseClient) createVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateVmClusterNetworkResponse
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
// See also
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
func (client DatabaseClient) dbNodeAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbNodes/{dbNodeId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DbNodeActionResponse
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

// DeleteAutonomousDatabase Deletes the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) deleteAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousDatabases/{autonomousDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousDatabaseResponse
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

// DeleteAutonomousVmCluster Deletes the specified Autonomous VM cluster in an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) deleteAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousVmClusters/{autonomousVmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteAutonomousVmClusterResponse
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

// DeleteBackup Deletes a full backup. You cannot delete automatic backups using this API.
//
// See also
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
func (client DatabaseClient) deleteBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/backups/{backupId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteBackupResponse
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

// DeleteBackupDestination Deletes a backup destination in an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) deleteBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/backupDestinations/{backupDestinationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteBackupDestinationResponse
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

// DeleteCloudExadataInfrastructure Deletes the cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) deleteCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteCloudExadataInfrastructureResponse
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

// DeleteCloudVmCluster Deletes the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) deleteCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cloudVmClusters/{cloudVmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteCloudVmClusterResponse
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

// DeleteConsoleConnection Deletes the specified database node console connection.
//
// See also
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
func (client DatabaseClient) deleteConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbNodes/{dbNodeId}/consoleConnections/{consoleConnectionId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteConsoleConnectionResponse
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

// DeleteDatabase Deletes the specified database. Applies only to Exadata systems.
// The data in this database is local to the Exadata system and will be lost when the database is deleted. Oracle recommends that you back up any data in the Exadata system prior to deleting it. You can use the `performFinalBackup` parameter to have the Exadata system database backed up before it is deleted.
//
// See also
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
func (client DatabaseClient) deleteDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databases/{databaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseResponse
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

// DeleteDatabaseSoftwareImage Delete a database software image
//
// See also
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
func (client DatabaseClient) deleteDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseSoftwareImages/{databaseSoftwareImageId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseSoftwareImageResponse
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

// DeleteDbHome Deletes a Database Home. Applies to bare metal DB systems, Exadata Cloud Service, and Exadata Cloud@Customer systems.
// Oracle recommends that you use the `performFinalBackup` parameter to back up any data on a bare metal DB system before you delete a Database Home. On an Exadata Cloud@Customer system or an Exadata Cloud Service system, you can delete a Database Home only when there are no databases in it and therefore you cannot use the `performFinalBackup` parameter to back up data.
//
// See also
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
func (client DatabaseClient) deleteDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbHomes/{dbHomeId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDbHomeResponse
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

// DeleteExadataInfrastructure Deletes the Exadata Cloud@Customer infrastructure.
//
// See also
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
func (client DatabaseClient) deleteExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exadataInfrastructures/{exadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteExadataInfrastructureResponse
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

// DeleteExternalContainerDatabase Deletes the CreateExternalContainerDatabaseDetails
// resource. Any external pluggable databases registered under this container database must be deleted in
// your Oracle Cloud Infrastructure tenancy prior to this operation.
//
// See also
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
func (client DatabaseClient) deleteExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalcontainerdatabases/{externalContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalContainerDatabaseResponse
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

// DeleteExternalDatabaseConnector Deletes an external database connector.
// Any services enabled using the external database connector must be
// deleted prior to this operation.
//
// See also
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
func (client DatabaseClient) deleteExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externaldatabaseconnectors/{externalDatabaseConnectorId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalDatabaseConnectorResponse
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

// DeleteExternalNonContainerDatabase Deletes the Oracle Cloud Infrastructure resource representing an external non-container database.
//
// See also
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
func (client DatabaseClient) deleteExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalNonContainerDatabaseResponse
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

// DeleteExternalPluggableDatabase Deletes the CreateExternalPluggableDatabaseDetails.
// resource.
//
// See also
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
func (client DatabaseClient) deleteExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalpluggabledatabases/{externalPluggableDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalPluggableDatabaseResponse
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

// DeleteKeyStore Deletes a key store.
//
// See also
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
func (client DatabaseClient) deleteKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/keyStores/{keyStoreId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteKeyStoreResponse
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

// DeleteVmCluster Deletes the specified VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) deleteVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/vmClusters/{vmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteVmClusterResponse
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

// DeleteVmClusterNetwork Deletes the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To delete a cloud VM cluster in an Exadata Cloud Service instance, use the DeleteCloudVmCluster operation.
//
// See also
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
func (client DatabaseClient) deleteVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteVmClusterNetworkResponse
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

// DeregisterAutonomousDatabaseDataSafe Asynchronously deregisters this Autonomous Database with Data Safe.
//
// See also
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
func (client DatabaseClient) deregisterAutonomousDatabaseDataSafe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/deregisterDataSafe", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeregisterAutonomousDatabaseDataSafeResponse
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

// DisableAutonomousDatabaseOperationsInsights Disables Operations Insights for the Autonomous Database resource.
//
// See also
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
func (client DatabaseClient) disableAutonomousDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/disableOperationsInsights", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DisableAutonomousDatabaseOperationsInsightsResponse
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

// DisableExternalContainerDatabaseDatabaseManagement Disable Database Management service for the external container database.
//
// See also
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
func (client DatabaseClient) disableExternalContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/disableDatabaseManagement", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DisableExternalContainerDatabaseDatabaseManagementResponse
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

// DisableExternalNonContainerDatabaseDatabaseManagement Disable Database Management Service for the external non-container database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// See also
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
func (client DatabaseClient) disableExternalNonContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/disableDatabaseManagement", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DisableExternalNonContainerDatabaseDatabaseManagementResponse
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

// DisableExternalNonContainerDatabaseOperationsInsights Disable Operations Insights for the external non-container database.
//
// See also
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
func (client DatabaseClient) disableExternalNonContainerDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/disableOperationsInsights", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DisableExternalNonContainerDatabaseOperationsInsightsResponse
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

// DisableExternalPluggableDatabaseDatabaseManagement Disable Database Management Service for the external pluggable database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// See also
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
func (client DatabaseClient) disableExternalPluggableDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/disableDatabaseManagement", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DisableExternalPluggableDatabaseDatabaseManagementResponse
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

// DisableExternalPluggableDatabaseOperationsInsights Disable Operations Insights for the external pluggable database.
//
// See also
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
func (client DatabaseClient) disableExternalPluggableDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/disableOperationsInsights", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DisableExternalPluggableDatabaseOperationsInsightsResponse
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

// DownloadExadataInfrastructureConfigFile Downloads the configuration file for the specified Exadata Cloud@Customer infrastructure.
//
// See also
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
func (client DatabaseClient) downloadExadataInfrastructureConfigFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/actions/downloadConfigFile", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DownloadExadataInfrastructureConfigFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadVmClusterNetworkConfigFile Downloads the configuration file for the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) downloadVmClusterNetworkConfigFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}/actions/downloadConfigFile", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DownloadVmClusterNetworkConfigFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAutonomousDatabaseOperationsInsights Enables the specified Autonomous Database with Operations Insights.
//
// See also
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
func (client DatabaseClient) enableAutonomousDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/enableOperationsInsights", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response EnableAutonomousDatabaseOperationsInsightsResponse
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

// EnableExternalContainerDatabaseDatabaseManagement Enables Database Management Service for the external container database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// See also
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
func (client DatabaseClient) enableExternalContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/enableDatabaseManagement", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response EnableExternalContainerDatabaseDatabaseManagementResponse
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

// EnableExternalNonContainerDatabaseDatabaseManagement Enable Database Management Service for the external non-container database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// See also
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
func (client DatabaseClient) enableExternalNonContainerDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/enableDatabaseManagement", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response EnableExternalNonContainerDatabaseDatabaseManagementResponse
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

// EnableExternalNonContainerDatabaseOperationsInsights Enable Operations Insights for the external non-container database.
//
// See also
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
func (client DatabaseClient) enableExternalNonContainerDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}/actions/enableOperationsInsights", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response EnableExternalNonContainerDatabaseOperationsInsightsResponse
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

// EnableExternalPluggableDatabaseDatabaseManagement Enable Database Management Service for the external pluggable database.
// For more information about the Database Management Service, see
// Database Management Service (https://docs.cloud.oracle.com/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).
//
// See also
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
func (client DatabaseClient) enableExternalPluggableDatabaseDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/enableDatabaseManagement", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response EnableExternalPluggableDatabaseDatabaseManagementResponse
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

// EnableExternalPluggableDatabaseOperationsInsights Enable Operations Insights for the external pluggable database.
//
// See also
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
func (client DatabaseClient) enableExternalPluggableDatabaseOperationsInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalpluggabledatabases/{externalPluggableDatabaseId}/actions/enableOperationsInsights", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response EnableExternalPluggableDatabaseOperationsInsightsResponse
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

// FailOverAutonomousDatabase Initiates a failover the specified Autonomous Database to a standby.
//
// See also
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
func (client DatabaseClient) failOverAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/failover", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response FailOverAutonomousDatabaseResponse
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

// FailoverAutonomousContainerDatabaseDataguardAssociation Fails over the standby Autonomous Container Database identified by the autonomousContainerDatabaseId parameter to the primary Autonomous Container Database after the existing primary Autonomous Container Database fails or becomes unreachable.
// A failover can result in data loss, depending on the protection mode in effect at the time the primary Autonomous Container Database fails.
//
// See also
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
func (client DatabaseClient) failoverAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}/actions/failover", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response FailoverAutonomousContainerDatabaseDataguardAssociationResponse
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

// FailoverDataGuardAssociation Performs a failover to transition the standby database identified by the `databaseId` parameter into the
// specified Data Guard association's primary role after the existing primary database fails or becomes unreachable.
// A failover might result in data loss depending on the protection mode in effect at the time of the primary
// database failure.
//
// See also
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
func (client DatabaseClient) failoverDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/failover", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response FailoverDataGuardAssociationResponse
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

// GenerateAutonomousDatabaseWallet Creates and downloads a wallet for the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) generateAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/generateWallet", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GenerateAutonomousDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateRecommendedVmClusterNetwork Generates a recommended Cloud@Customer VM cluster network configuration.
//
// See also
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
func (client DatabaseClient) generateRecommendedVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/actions/generateRecommendedNetwork", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GenerateRecommendedVmClusterNetworkResponse
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

// GetAutonomousContainerDatabase Gets information about the specified Autonomous Container Database.
//
// See also
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
func (client DatabaseClient) getAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseResponse
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

// GetAutonomousContainerDatabaseDataguardAssociation Gets an Autonomous Container Database enabled with Autonomous Data Guard associated with the specified Autonomous Container Database.
//
// See also
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
func (client DatabaseClient) getAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousContainerDatabaseDataguardAssociationResponse
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

// GetAutonomousDatabase Gets the details of the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) getAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseResponse
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

// GetAutonomousDatabaseBackup Gets information about the specified Autonomous Database backup.
//
// See also
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
func (client DatabaseClient) getAutonomousDatabaseBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseBackups/{autonomousDatabaseBackupId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseBackupResponse
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

// GetAutonomousDatabaseDataguardAssociation Gets an Autonomous Data Guard-enabled database associated with the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) getAutonomousDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/autonomousDatabaseDataguardAssociations/{autonomousDatabaseDataguardAssociationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseDataguardAssociationResponse
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

// GetAutonomousDatabaseRegionalWallet Gets the Autonomous Database regional wallet details.
//
// See also
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
func (client DatabaseClient) getAutonomousDatabaseRegionalWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/wallet", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseRegionalWalletResponse
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

// GetAutonomousDatabaseWallet Gets the wallet details for the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) getAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/wallet", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousDatabaseWalletResponse
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

// GetAutonomousExadataInfrastructure Gets information about the specified Autonomous Exadata Infrastructure resource.
//
// See also
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
func (client DatabaseClient) getAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousExadataInfrastructureResponse
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

// GetAutonomousPatch Gets information about a specific autonomous patch.
//
// See also
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
func (client DatabaseClient) getAutonomousPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousPatches/{autonomousPatchId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousPatchResponse
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

// GetAutonomousVmCluster Gets information about the specified Autonomous VM cluster for an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) getAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVmClusters/{autonomousVmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutonomousVmClusterResponse
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

// GetBackup Gets information about the specified backup.
//
// See also
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
func (client DatabaseClient) getBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups/{backupId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetBackupResponse
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

// GetBackupDestination Gets information about the specified backup destination in an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) getBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backupDestinations/{backupDestinationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetBackupDestinationResponse
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

// GetCloudExadataInfrastructure Gets information about the specified cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) getCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetCloudExadataInfrastructureResponse
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

// GetCloudVmCluster Gets information about the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) getCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterResponse
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

// GetCloudVmClusterIormConfig Gets the IORM configuration for the specified cloud VM cluster in an Exadata Cloud Service instance.
// If you have not specified an IORM configuration, the default configuration is returned.
//
// See also
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
func (client DatabaseClient) getCloudVmClusterIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/CloudVmClusterIormConfig", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterIormConfigResponse
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

// GetCloudVmClusterUpdate Gets information about a specified maintenance update package for a cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) getCloudVmClusterUpdate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updates/{updateId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterUpdateResponse
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

// GetCloudVmClusterUpdateHistoryEntry Gets the maintenance update history details for the specified update history entry. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) getCloudVmClusterUpdateHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updateHistoryEntries/{updateHistoryEntryId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetCloudVmClusterUpdateHistoryEntryResponse
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

// GetConsoleConnection Gets the specified database node console connection's information.
//
// See also
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
func (client DatabaseClient) getConsoleConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleConnections/{consoleConnectionId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetConsoleConnectionResponse
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

// GetDataGuardAssociation Gets the specified Data Guard association's configuration information.
//
// See also
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
func (client DatabaseClient) getDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDataGuardAssociationResponse
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

// GetDatabase Gets information about the specified database.
//
// See also
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
func (client DatabaseClient) getDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseResponse
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

// GetDatabaseSoftwareImage Gets information about the specified database software image.
//
// See also
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
func (client DatabaseClient) getDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseSoftwareImages/{databaseSoftwareImageId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseSoftwareImageResponse
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

// GetDatabaseUpgradeHistoryEntry gets the upgrade history for a specified database.
//
// See also
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
func (client DatabaseClient) getDatabaseUpgradeHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/upgradeHistoryEntries/{upgradeHistoryEntryId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseUpgradeHistoryEntryResponse
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

// GetDbHome Gets information about the specified Database Home.
//
// See also
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
func (client DatabaseClient) getDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbHomeResponse
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

// GetDbHomePatch Gets information about a specified patch package.
//
// See also
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
func (client DatabaseClient) getDbHomePatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches/{patchId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchResponse
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

// GetDbHomePatchHistoryEntry Gets the patch history details for the specified patchHistoryEntryId
//
// See also
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
func (client DatabaseClient) getDbHomePatchHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries/{patchHistoryEntryId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbHomePatchHistoryEntryResponse
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

// GetDbNode Gets information about the specified database node.
//
// See also
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
func (client DatabaseClient) getDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbNodeResponse
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

// GetDbSystem Gets information about the specified DB system.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// See also
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
func (client DatabaseClient) getDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemResponse
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

// GetDbSystemPatch Gets information the specified patch.
//
// See also
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
func (client DatabaseClient) getDbSystemPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches/{patchId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchResponse
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

// GetDbSystemPatchHistoryEntry Gets the details of the specified patch operation on the specified DB system.
//
// See also
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
func (client DatabaseClient) getDbSystemPatchHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries/{patchHistoryEntryId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemPatchHistoryEntryResponse
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

// GetExadataInfrastructure Gets information about the specified Exadata infrastructure. Applies to Exadata Cloud@Customer instances only.
// To get information on an Exadata Cloud Service infrastructure resource, use the  GetCloudExadataInfrastructure operation.
//
// See also
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
func (client DatabaseClient) getExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExadataInfrastructureResponse
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

// GetExadataInfrastructureOcpus Gets details of the available and consumed OCPUs for the specified Autonomous Exadata Infrastructure resource.
//
// See also
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
func (client DatabaseClient) getExadataInfrastructureOcpus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/ocpus", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExadataInfrastructureOcpusResponse
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

// GetExadataIormConfig Gets the IORM configuration settings for the specified cloud Exadata DB system.
// All Exadata service instances have default IORM settings.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
// The GetCloudVmClusterIormConfig API is used for this operation with Exadata systems using the
// new resource model.
//
// See also
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
func (client DatabaseClient) getExadataIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/ExadataIormConfig", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExadataIormConfigResponse
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

// GetExternalBackupJob Gets information about the specified external backup job.
// **Note:** This API is used by an Oracle Cloud Infrastructure Python script that is packaged with the Oracle Cloud Infrastructure CLI. Oracle recommends that you use the script instead using the API directly. See Migrating an On-Premises Database to Oracle Cloud Infrastructure by Creating a Backup in the Cloud (https://docs.cloud.oracle.com/Content/Database/Tasks/mig-onprembackup.htm) for more information.
//
// See also
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
func (client DatabaseClient) getExternalBackupJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalBackupJobs/{backupId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExternalBackupJobResponse
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

// GetExternalContainerDatabase Gets information about the specified external container database.
//
// See also
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
func (client DatabaseClient) getExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalcontainerdatabases/{externalContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExternalContainerDatabaseResponse
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

// GetExternalDatabaseConnector Gets information about the specified external database connector.
//
// See also
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
func (client DatabaseClient) getExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externaldatabaseconnectors/{externalDatabaseConnectorId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldatabaseconnector{})
	return response, err
}

// GetExternalNonContainerDatabase Gets information about a specific external non-container database.
//
// See also
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
func (client DatabaseClient) getExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExternalNonContainerDatabaseResponse
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

// GetExternalPluggableDatabase Gets information about a specific
// CreateExternalPluggableDatabaseDetails resource.
//
// See also
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
func (client DatabaseClient) getExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalpluggabledatabases/{externalPluggableDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetExternalPluggableDatabaseResponse
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

// GetKeyStore Gets information about the specified key store.
//
// See also
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
func (client DatabaseClient) getKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keyStores/{keyStoreId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetKeyStoreResponse
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

// GetMaintenanceRun Gets information about the specified maintenance run.
//
// See also
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
func (client DatabaseClient) getMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns/{maintenanceRunId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceRunResponse
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

// GetVmCluster Gets information about the VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) getVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterResponse
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

// GetVmClusterNetwork Gets information about the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To get information about a cloud VM cluster in an Exadata Cloud Service instance, use the GetCloudVmCluster operation.
//
// See also
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
func (client DatabaseClient) getVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterNetworkResponse
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

// GetVmClusterPatch Gets information about a specified patch package.
//
// See also
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
func (client DatabaseClient) getVmClusterPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patches/{patchId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterPatchResponse
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

// GetVmClusterPatchHistoryEntry Gets the patch history details for the specified patch history entry.
//
// See also
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
func (client DatabaseClient) getVmClusterPatchHistoryEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patchHistoryEntries/{patchHistoryEntryId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterPatchHistoryEntryResponse
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

// LaunchAutonomousExadataInfrastructure Creates a new Autonomous Exadata Infrastructure in the specified compartment and availability domain.
//
// See also
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
func (client DatabaseClient) launchAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response LaunchAutonomousExadataInfrastructureResponse
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

// LaunchDbSystem Creates a new DB system in the specified compartment and availability domain. The Oracle
// Database edition that you specify applies to all the databases on that DB system. The selected edition cannot be changed.
// An initial database is created on the DB system based on the request parameters you provide and some default
// options. For detailed information about default options, see Bare metal and virtual machine DB system default options. (https://docs.cloud.oracle.com/Content/Database/Tasks/creatingDBsystem.htm#Default)
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
// Use the CreateCloudExadataInfrastructure and CreateCloudVmCluster APIs to provision a new Exadata Cloud Service instance.
//
// See also
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
func (client DatabaseClient) launchDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response LaunchDbSystemResponse
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

// ListAutonomousContainerDatabaseDataguardAssociations Gets a list of the Autonomous Container Databases with Autonomous Data Guard-enabled associated with the specified Autonomous Container Database.
//
// See also
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
func (client DatabaseClient) listAutonomousContainerDatabaseDataguardAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabaseDataguardAssociationsResponse
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

// ListAutonomousContainerDatabases Gets a list of the Autonomous Container Databases in the specified compartment.
//
// See also
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
func (client DatabaseClient) listAutonomousContainerDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousContainerDatabasesResponse
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

// ListAutonomousDatabaseBackups Gets a list of Autonomous Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.
//
// See also
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
func (client DatabaseClient) listAutonomousDatabaseBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabaseBackups", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseBackupsResponse
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

// ListAutonomousDatabaseClones Lists the Autonomous Database clones for the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) listAutonomousDatabaseClones(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/clones", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseClonesResponse
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

// ListAutonomousDatabaseDataguardAssociations Gets a list of the Autonomous Data Guard-enabled databases associated with the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) listAutonomousDatabaseDataguardAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases/{autonomousDatabaseId}/autonomousDatabaseDataguardAssociations", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabaseDataguardAssociationsResponse
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

// ListAutonomousDatabases Gets a list of Autonomous Databases based on the query parameters specified.
//
// See also
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
func (client DatabaseClient) listAutonomousDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDatabasesResponse
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

// ListAutonomousDbPreviewVersions Gets a list of supported Autonomous Database versions. Note that preview version software is only available for
// databases with shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI).
//
// See also
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
func (client DatabaseClient) listAutonomousDbPreviewVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbPreviewVersions", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDbPreviewVersionsResponse
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

// ListAutonomousDbVersions Gets a list of supported Autonomous Database versions.
//
// See also
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
func (client DatabaseClient) listAutonomousDbVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousDbVersions", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousDbVersionsResponse
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

// ListAutonomousExadataInfrastructureShapes Gets a list of the shapes that can be used to launch a new Autonomous Exadata Infrastructure resource. The shape determines resources to allocate (CPU cores, memory and storage).
//
// See also
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
func (client DatabaseClient) listAutonomousExadataInfrastructureShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructureShapes", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousExadataInfrastructureShapesResponse
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

// ListAutonomousExadataInfrastructures Gets a list of the Autonomous Exadata Infrastructures in the specified compartment.
//
// See also
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
func (client DatabaseClient) listAutonomousExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousExadataInfrastructures", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousExadataInfrastructuresResponse
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

// ListAutonomousVmClusters Gets a list of Exadata Cloud@Customer Autonomous VM clusters in the specified compartment.
//
// See also
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
func (client DatabaseClient) listAutonomousVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousVmClusters", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutonomousVmClustersResponse
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

// ListBackupDestination Gets a list of backup destinations in the specified compartment.
//
// See also
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
func (client DatabaseClient) listBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backupDestinations", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListBackupDestinationResponse
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

// ListBackups Gets a list of backups based on the `databaseId` or `compartmentId` specified. Either one of these query parameters must be provided.
//
// See also
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
func (client DatabaseClient) listBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListBackupsResponse
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

// ListCloudExadataInfrastructures Gets a list of the cloud Exadata infrastructure resources in the specified compartment. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) listCloudExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudExadataInfrastructures", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListCloudExadataInfrastructuresResponse
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

// ListCloudVmClusterUpdateHistoryEntries Gets the history of the maintenance update actions performed on the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) listCloudVmClusterUpdateHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updateHistoryEntries", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListCloudVmClusterUpdateHistoryEntriesResponse
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

// ListCloudVmClusterUpdates Lists the maintenance updates that can be applied to the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) listCloudVmClusterUpdates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters/{cloudVmClusterId}/updates", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListCloudVmClusterUpdatesResponse
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

// ListCloudVmClusters Gets a list of the cloud VM clusters in the specified compartment. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) listCloudVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cloudVmClusters", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListCloudVmClustersResponse
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

// ListConsoleConnections Lists the console connections for the specified database node.
//
// See also
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
func (client DatabaseClient) listConsoleConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes/{dbNodeId}/consoleConnections", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListConsoleConnectionsResponse
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

// ListContainerDatabasePatches Lists the patches applicable to the requested container database.
//
// See also
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
func (client DatabaseClient) listContainerDatabasePatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/patches", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListContainerDatabasePatchesResponse
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

// ListDataGuardAssociations Lists all Data Guard associations for the specified database.
//
// See also
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
func (client DatabaseClient) listDataGuardAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/dataGuardAssociations", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDataGuardAssociationsResponse
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

// ListDatabaseSoftwareImages Gets a list of the database software images in the specified compartment.
//
// See also
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
func (client DatabaseClient) listDatabaseSoftwareImages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseSoftwareImages", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseSoftwareImagesResponse
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

// ListDatabaseUpgradeHistoryEntries Gets the upgrade history for a specified database in a bare metal or virtual machine DB system.
//
// See also
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
func (client DatabaseClient) listDatabaseUpgradeHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases/{databaseId}/upgradeHistoryEntries", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseUpgradeHistoryEntriesResponse
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

// ListDatabases Gets a list of the databases in the specified Database Home.
//
// See also
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
func (client DatabaseClient) listDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDatabasesResponse
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

// ListDbHomePatchHistoryEntries Lists the history of patch operations on the specified Database Home.
//
// See also
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
func (client DatabaseClient) listDbHomePatchHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patchHistoryEntries", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchHistoryEntriesResponse
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

// ListDbHomePatches Lists patches applicable to the requested Database Home.
//
// See also
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
func (client DatabaseClient) listDbHomePatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes/{dbHomeId}/patches", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbHomePatchesResponse
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

// ListDbHomes Lists the Database Homes in the specified DB system and compartment. A Database Home is a directory where Oracle Database software is installed.
//
// See also
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
func (client DatabaseClient) listDbHomes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbHomes", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbHomesResponse
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

// ListDbNodes Lists the database nodes in the specified DB system and compartment. A database node is a server running database software.
//
// See also
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
func (client DatabaseClient) listDbNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbNodes", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbNodesResponse
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

// ListDbSystemPatchHistoryEntries Gets the history of the patch actions performed on the specified DB system.
//
// See also
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
func (client DatabaseClient) listDbSystemPatchHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patchHistoryEntries", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchHistoryEntriesResponse
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

// ListDbSystemPatches Lists the patches applicable to the specified DB system.
//
// See also
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
func (client DatabaseClient) listDbSystemPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/patches", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemPatchesResponse
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

// ListDbSystemShapes Gets a list of the shapes that can be used to launch a new DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
//
// See also
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
func (client DatabaseClient) listDbSystemShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemShapes", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemShapesResponse
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

// ListDbSystems Lists the DB systems in the specified compartment. You can specify a `backupId` to list only the DB systems that support creating a database using this backup in this compartment.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// See also
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
func (client DatabaseClient) listDbSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemsResponse
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

// ListDbVersions Gets a list of supported Oracle Database versions.
//
// See also
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
func (client DatabaseClient) listDbVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbVersions", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbVersionsResponse
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

// ListExadataInfrastructures Lists the Exadata infrastructure resources in the specified compartment. Applies to Exadata Cloud@Customer instances only.
// To list the Exadata Cloud Service infrastructure resources in a compartment, use the  ListCloudExadataInfrastructures operation.
//
// See also
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
func (client DatabaseClient) listExadataInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListExadataInfrastructuresResponse
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

// ListExternalContainerDatabases Gets a list of the external container databases in the specified compartment.
//
// See also
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
func (client DatabaseClient) listExternalContainerDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalcontainerdatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListExternalContainerDatabasesResponse
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

//listexternaldatabaseconnectorsummary allows to unmarshal list of polymorphic ExternalDatabaseConnectorSummary
type listexternaldatabaseconnectorsummary []externaldatabaseconnectorsummary

//UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
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
// See also
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
func (client DatabaseClient) listExternalDatabaseConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externaldatabaseconnectors", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListExternalDatabaseConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listexternaldatabaseconnectorsummary{})
	return response, err
}

// ListExternalNonContainerDatabases Gets a list of the ExternalNonContainerDatabases in the specified compartment.
//
// See also
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
func (client DatabaseClient) listExternalNonContainerDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalnoncontainerdatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListExternalNonContainerDatabasesResponse
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

// ListExternalPluggableDatabases Gets a list of the CreateExternalPluggableDatabaseDetails
// resources in the specified compartment.
//
// See also
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
func (client DatabaseClient) listExternalPluggableDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalpluggabledatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListExternalPluggableDatabasesResponse
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

// ListFlexComponents Gets a list of the flex components that can be used to launch a new DB system. The flex component determines resources to allocate to the DB system - Database Servers and Storage Servers.
//
// See also
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
func (client DatabaseClient) listFlexComponents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystemShapes/flexComponents", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListFlexComponentsResponse
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

// ListGiVersions Gets a list of supported GI versions for the Exadata Cloud@Customer VM cluster.
//
// See also
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
func (client DatabaseClient) listGiVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/giVersions", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListGiVersionsResponse
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

// ListKeyStores Gets a list of key stores in the specified compartment.
//
// See also
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
func (client DatabaseClient) listKeyStores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keyStores", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListKeyStoresResponse
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

// ListMaintenanceRuns Gets a list of the maintenance runs in the specified compartment.
//
// See also
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
func (client DatabaseClient) listMaintenanceRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceRunsResponse
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

// ListVmClusterNetworks Gets a list of the VM cluster networks in the specified compartment. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) listVmClusterNetworks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterNetworksResponse
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

// ListVmClusterPatchHistoryEntries Gets the history of the patch actions performed on the specified VM cluster in an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) listVmClusterPatchHistoryEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patchHistoryEntries", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterPatchHistoryEntriesResponse
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

// ListVmClusterPatches Lists the patches applicable to the specified VM cluster in an Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) listVmClusterPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters/{vmClusterId}/patches", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterPatchesResponse
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

// ListVmClusters Lists the VM clusters in the specified compartment. Applies to Exadata Cloud@Customer instances only.
// To list the cloud VM clusters in an Exadata Cloud Service instance, use the ListCloudVmClusters operation.
//
// See also
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
func (client DatabaseClient) listVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusters", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListVmClustersResponse
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

// MigrateExadataDbSystemResourceModel Migrates the Exadata DB system to the new Exadata resource model (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model).
// All related resources will be migrated.
//
// See also
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
func (client DatabaseClient) migrateExadataDbSystemResourceModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/migration", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response MigrateExadataDbSystemResourceModelResponse
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

// MigrateVaultKey Changes encryption key management from customer-managed, using the Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm), to Oracle-managed.
//
// See also
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
func (client DatabaseClient) migrateVaultKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/migrateKey", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response MigrateVaultKeyResponse
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

// RegisterAutonomousDatabaseDataSafe Asynchronously registers this Autonomous Database with Data Safe.
//
// See also
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
func (client DatabaseClient) registerAutonomousDatabaseDataSafe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/registerDataSafe", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RegisterAutonomousDatabaseDataSafeResponse
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

// ReinstateAutonomousContainerDatabaseDataguardAssociation Reinstates a disabled standby Autonomous Container Database, identified by the autonomousContainerDatabaseId parameter, to an active standby Autonomous Container Database.
//
// See also
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
func (client DatabaseClient) reinstateAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}/actions/reinstate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ReinstateAutonomousContainerDatabaseDataguardAssociationResponse
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

// ReinstateDataGuardAssociation Reinstates the database identified by the `databaseId` parameter into the standby role in a Data Guard association.
//
// See also
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
func (client DatabaseClient) reinstateDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/reinstate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ReinstateDataGuardAssociationResponse
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

// RestartAutonomousContainerDatabase Rolling restarts the specified Autonomous Container Database.
//
// See also
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
func (client DatabaseClient) restartAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/restart", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestartAutonomousContainerDatabaseResponse
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

// RestartAutonomousDatabase Restarts the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) restartAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/restart", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestartAutonomousDatabaseResponse
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

// RestoreAutonomousDatabase Restores an Autonomous Database based on the provided request parameters.
//
// See also
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
func (client DatabaseClient) restoreAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/restore", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestoreAutonomousDatabaseResponse
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

// RestoreDatabase Restore a Database based on the request parameters you provide.
//
// See also
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
func (client DatabaseClient) restoreDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/restore", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestoreDatabaseResponse
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

// RotateAutonomousContainerDatabaseEncryptionKey Creates a new version of an existing Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) key.
//
// See also
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
func (client DatabaseClient) rotateAutonomousContainerDatabaseEncryptionKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/actions/rotateKey", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RotateAutonomousContainerDatabaseEncryptionKeyResponse
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

// RotateAutonomousDatabaseEncryptionKey Rotate existing AutonomousDatabase Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) key.
//
// See also
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
func (client DatabaseClient) rotateAutonomousDatabaseEncryptionKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/rotateKey", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RotateAutonomousDatabaseEncryptionKeyResponse
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

// RotateOrdsCerts Rotates Oracle REST Data Services (ORDS) certs for an Autonomous Exadata Infrastructure resource.
//
// See also
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
func (client DatabaseClient) rotateOrdsCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/actions/rotateOrdsCerts", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RotateOrdsCertsResponse
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

// RotateSslCerts Rotates SSL certs for an Autonomous Exadata Infrastructure resource.
//
// See also
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
func (client DatabaseClient) rotateSslCerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}/actions/rotateSslCerts", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RotateSslCertsResponse
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

// RotateVaultKey Creates a new version of an existing Vault service (https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) key.
//
// See also
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
func (client DatabaseClient) rotateVaultKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/rotateKey", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RotateVaultKeyResponse
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

// ScanExternalContainerDatabasePluggableDatabases Scans for pluggable databases in the specified external container database.
// This operation will return un-registered pluggable databases in the `GetWorkRequest` operation.
//
// See also
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
func (client DatabaseClient) scanExternalContainerDatabasePluggableDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalcontainerdatabases/{externalContainerDatabaseId}/actions/scanPluggableDatabases", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ScanExternalContainerDatabasePluggableDatabasesResponse
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

// StartAutonomousDatabase Starts the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) startAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/start", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StartAutonomousDatabaseResponse
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

// StopAutonomousDatabase Stops the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) stopAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/stop", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StopAutonomousDatabaseResponse
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

// SwitchoverAutonomousContainerDatabaseDataguardAssociation Switches over the primary Autonomous Container Database of an Autonomous Data Guard peer association to standby role. The standby Autonomous Container Database associated with autonomousContainerDatabaseDataguardAssociationId assumes the primary Autonomous Container Database role.
// A switchover incurs no data loss.
//
// See also
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
func (client DatabaseClient) switchoverAutonomousContainerDatabaseDataguardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}/actions/switchover", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response SwitchoverAutonomousContainerDatabaseDataguardAssociationResponse
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

// SwitchoverAutonomousDatabase Initiates a switchover of the specified Autonomous Database to the associated standby database. Applicable only to databases with Autonomous Data Guard enabled.
//
// See also
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
func (client DatabaseClient) switchoverAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autonomousDatabases/{autonomousDatabaseId}/actions/switchover", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response SwitchoverAutonomousDatabaseResponse
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

// SwitchoverDataGuardAssociation Performs a switchover to transition the primary database of a Data Guard association into a standby role. The
// standby database associated with the `dataGuardAssociationId` assumes the primary database role.
// A switchover guarantees no data loss.
//
// See also
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
func (client DatabaseClient) switchoverDataGuardAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/dataGuardAssociations/{dataGuardAssociationId}/actions/switchover", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response SwitchoverDataGuardAssociationResponse
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

// TerminateAutonomousContainerDatabase Terminates an Autonomous Container Database, which permanently deletes the container database and any databases within the container database. The database data is local to the Autonomous Exadata Infrastructure and will be lost when the container database is terminated. Oracle recommends that you back up any data in the Autonomous Container Database prior to terminating it.
//
// See also
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
func (client DatabaseClient) terminateAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousContainerDatabaseResponse
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

// TerminateAutonomousExadataInfrastructure Terminates an Autonomous Exadata Infrastructure, which permanently deletes the infrastructure resource and any container databases and databases contained in the resource. The database data is local to the Autonomous Exadata Infrastructure and will be lost when the system is terminated. Oracle recommends that you back up any data in the Autonomous Exadata Infrastructure prior to terminating it.
//
// See also
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
func (client DatabaseClient) terminateAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response TerminateAutonomousExadataInfrastructureResponse
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

// TerminateDbSystem Terminates a DB system and permanently deletes it and any databases running on it, and any storage volumes attached to it. The database data is local to the DB system and will be lost when the system is terminated. Oracle recommends that you back up any data in the DB system prior to terminating it.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// See also
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
func (client DatabaseClient) terminateDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbSystems/{dbSystemId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response TerminateDbSystemResponse
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

// UpdateAutonomousContainerDatabase Updates the properties of an Autonomous Container Database, such as the OCPU core count and storage size.
//
// See also
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
func (client DatabaseClient) updateAutonomousContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousContainerDatabases/{autonomousContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousContainerDatabaseResponse
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

// UpdateAutonomousDatabase Updates one or more attributes of the specified Autonomous Database. See the UpdateAutonomousDatabaseDetails resource for a full list of attributes that can be updated.
//
// See also
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
func (client DatabaseClient) updateAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/{autonomousDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseResponse
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

// UpdateAutonomousDatabaseRegionalWallet Updates the Autonomous Database regional wallet.
//
// See also
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
func (client DatabaseClient) updateAutonomousDatabaseRegionalWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/wallet", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseRegionalWalletResponse
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

// UpdateAutonomousDatabaseWallet Updates the wallet for the specified Autonomous Database.
//
// See also
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
func (client DatabaseClient) updateAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousDatabases/{autonomousDatabaseId}/wallet", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousDatabaseWalletResponse
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

// UpdateAutonomousExadataInfrastructure Updates the properties of an Autonomous Exadata Infrastructure, such as the CPU core count.
//
// See also
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
func (client DatabaseClient) updateAutonomousExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousExadataInfrastructures/{autonomousExadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousExadataInfrastructureResponse
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

// UpdateAutonomousVmCluster Updates the specified Autonomous VM cluster for the Exadata Cloud@Customer system.
//
// See also
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
func (client DatabaseClient) updateAutonomousVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autonomousVmClusters/{autonomousVmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutonomousVmClusterResponse
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

// UpdateBackupDestination If no database is associated with the backup destination:
// - For a RECOVERY_APPLIANCE backup destination, updates the connection string and/or the list of VPC users.
// - For an NFS backup destination, updates the NFS location.
//
// See also
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
func (client DatabaseClient) updateBackupDestination(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/backupDestinations/{backupDestinationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateBackupDestinationResponse
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

// UpdateCloudExadataInfrastructure Updates the Cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) updateCloudExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudExadataInfrastructures/{cloudExadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudExadataInfrastructureResponse
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

// UpdateCloudVmCluster Updates the specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
//
// See also
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
func (client DatabaseClient) updateCloudVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudVmClusters/{cloudVmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudVmClusterResponse
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

// UpdateCloudVmClusterIormConfig Updates the IORM settings for the specified cloud VM cluster in an Exadata Cloud Service instance.
//
// See also
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
func (client DatabaseClient) updateCloudVmClusterIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cloudVmClusters/{cloudVmClusterId}/CloudVmClusterIormConfig", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateCloudVmClusterIormConfigResponse
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

// UpdateDatabase Update the specified database based on the request parameters provided.
//
// See also
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
func (client DatabaseClient) updateDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databases/{databaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseResponse
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

// UpdateDatabaseSoftwareImage Updates the properties of a Database Software Image, like Display Nmae
//
// See also
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
func (client DatabaseClient) updateDatabaseSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseSoftwareImages/{databaseSoftwareImageId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseSoftwareImageResponse
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

// UpdateDbHome Patches the specified Database Home.
//
// See also
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
func (client DatabaseClient) updateDbHome(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbHomes/{dbHomeId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDbHomeResponse
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

// UpdateDbSystem Updates the properties of the specified DB system.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// See also
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
func (client DatabaseClient) updateDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDbSystemResponse
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

// UpdateExadataInfrastructure Updates the Exadata infrastructure resource. Applies to Exadata Cloud@Customer instances only.
// To update an Exadata Cloud Service infrastructure resource, use the  UpdateCloudExadataInfrastructure operation.
//
// See also
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
func (client DatabaseClient) updateExadataInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exadataInfrastructures/{exadataInfrastructureId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateExadataInfrastructureResponse
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

// UpdateExadataIormConfig Updates IORM settings for the specified Exadata DB system.
// **Note:** Deprecated for Exadata Cloud Service systems. Use the new resource model APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See Switching an Exadata DB System to the New Resource Model and APIs (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
// The UpdateCloudVmClusterIormConfig API is used for Exadata systems using the
// new resource model.
//
// See also
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
func (client DatabaseClient) updateExadataIormConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}/ExadataIormConfig", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateExadataIormConfigResponse
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

// UpdateExternalContainerDatabase Updates the properties of
// an CreateExternalContainerDatabaseDetails resource,
// such as the display name.
//
// See also
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
func (client DatabaseClient) updateExternalContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalcontainerdatabases/{externalContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalContainerDatabaseResponse
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

// UpdateExternalDatabaseConnector Updates the properties of an external database connector, such as the display name.
//
// See also
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
func (client DatabaseClient) updateExternalDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externaldatabaseconnectors/{externalDatabaseConnectorId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &externaldatabaseconnector{})
	return response, err
}

// UpdateExternalNonContainerDatabase Updates the properties of an external non-container database, such as the display name.
//
// See also
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
func (client DatabaseClient) updateExternalNonContainerDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalnoncontainerdatabases/{externalNonContainerDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalNonContainerDatabaseResponse
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

// UpdateExternalPluggableDatabase Updates the properties of an
// CreateExternalPluggableDatabaseDetails resource,
// such as the display name.
//
// See also
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
func (client DatabaseClient) updateExternalPluggableDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalpluggabledatabases/{externalPluggableDatabaseId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalPluggableDatabaseResponse
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

// UpdateKeyStore If no database is associated with the key store, edit the key store.
//
// See also
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
func (client DatabaseClient) updateKeyStore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/keyStores/{keyStoreId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateKeyStoreResponse
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

// UpdateMaintenanceRun Updates the properties of a maintenance run, such as the state of a maintenance run.
//
// See also
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
func (client DatabaseClient) updateMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maintenanceRuns/{maintenanceRunId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateMaintenanceRunResponse
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

// UpdateVmCluster Updates the specified VM cluster. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) updateVmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/vmClusters/{vmClusterId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateVmClusterResponse
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

// UpdateVmClusterNetwork Updates the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To update a cloud VM cluster in an Exadata Cloud Service instance, use the UpdateCloudVmCluster operation.
//
// See also
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
func (client DatabaseClient) updateVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateVmClusterNetworkResponse
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

// UpgradeDatabase Upgrades the specified Oracle Database instance.
//
// See also
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
func (client DatabaseClient) upgradeDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databases/{databaseId}/actions/upgrade", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpgradeDatabaseResponse
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

// ValidateVmClusterNetwork Validates the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
//
// See also
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
func (client DatabaseClient) validateVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}/actions/validate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ValidateVmClusterNetworkResponse
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
