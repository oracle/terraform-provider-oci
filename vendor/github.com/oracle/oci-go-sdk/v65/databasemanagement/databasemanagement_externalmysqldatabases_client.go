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

// ExternalMySqlDatabasesClient a client for ExternalMySqlDatabases
type ExternalMySqlDatabasesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewExternalMySqlDatabasesClientWithConfigurationProvider Creates a new default ExternalMySqlDatabases client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewExternalMySqlDatabasesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ExternalMySqlDatabasesClient, err error) {
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
	return newExternalMySqlDatabasesClientFromBaseClient(baseClient, provider)
}

// NewExternalMySqlDatabasesClientWithOboToken Creates a new default ExternalMySqlDatabases client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewExternalMySqlDatabasesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ExternalMySqlDatabasesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newExternalMySqlDatabasesClientFromBaseClient(baseClient, configProvider)
}

func newExternalMySqlDatabasesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ExternalMySqlDatabasesClient, err error) {
	// ExternalMySqlDatabases service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ExternalMySqlDatabases"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ExternalMySqlDatabasesClient{BaseClient: baseClient}
	client.BasePath = "20201101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ExternalMySqlDatabasesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemanagement", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ExternalMySqlDatabasesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ExternalMySqlDatabasesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CheckExternalMySqlDatabaseConnectorConnectionStatus Check the status of the external database connection specified in this connector.
// This operation will refresh the connectionStatus and timeConnectionStatusLastUpdated fields.
// A default retry strategy applies to this operation CheckExternalMySqlDatabaseConnectorConnectionStatus()
func (client ExternalMySqlDatabasesClient) CheckExternalMySqlDatabaseConnectorConnectionStatus(ctx context.Context, request CheckExternalMySqlDatabaseConnectorConnectionStatusRequest) (response CheckExternalMySqlDatabaseConnectorConnectionStatusResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.checkExternalMySqlDatabaseConnectorConnectionStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CheckExternalMySqlDatabaseConnectorConnectionStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CheckExternalMySqlDatabaseConnectorConnectionStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CheckExternalMySqlDatabaseConnectorConnectionStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CheckExternalMySqlDatabaseConnectorConnectionStatusResponse")
	}
	return
}

// checkExternalMySqlDatabaseConnectorConnectionStatus implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) checkExternalMySqlDatabaseConnectorConnectionStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalMySqlDatabaseConnectors/{externalMySqlDatabaseConnectorId}/actions/checkConnectionStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CheckExternalMySqlDatabaseConnectorConnectionStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabaseConnector/CheckExternalMySqlDatabaseConnectorConnectionStatus"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "CheckExternalMySqlDatabaseConnectorConnectionStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalMySqlDatabase Creates an external MySQL database.
// A default retry strategy applies to this operation CreateExternalMySqlDatabase()
func (client ExternalMySqlDatabasesClient) CreateExternalMySqlDatabase(ctx context.Context, request CreateExternalMySqlDatabaseRequest) (response CreateExternalMySqlDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalMySqlDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalMySqlDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalMySqlDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalMySqlDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalMySqlDatabaseResponse")
	}
	return
}

// createExternalMySqlDatabase implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) createExternalMySqlDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalMySqlDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalMySqlDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/CreateExternalMySqlDatabase"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "CreateExternalMySqlDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalMySqlDatabaseConnector Creates an external MySQL connector resource.
// A default retry strategy applies to this operation CreateExternalMySqlDatabaseConnector()
func (client ExternalMySqlDatabasesClient) CreateExternalMySqlDatabaseConnector(ctx context.Context, request CreateExternalMySqlDatabaseConnectorRequest) (response CreateExternalMySqlDatabaseConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExternalMySqlDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalMySqlDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalMySqlDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalMySqlDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalMySqlDatabaseConnectorResponse")
	}
	return
}

// createExternalMySqlDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) createExternalMySqlDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalMySqlDatabaseConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalMySqlDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabaseConnector/CreateExternalMySqlDatabaseConnector"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "CreateExternalMySqlDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalMySqlDatabase Deletes the Oracle Cloud Infrastructure resource representing an external MySQL database.
func (client ExternalMySqlDatabasesClient) DeleteExternalMySqlDatabase(ctx context.Context, request DeleteExternalMySqlDatabaseRequest) (response DeleteExternalMySqlDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalMySqlDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalMySqlDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalMySqlDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalMySqlDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalMySqlDatabaseResponse")
	}
	return
}

// deleteExternalMySqlDatabase implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) deleteExternalMySqlDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalMySqlDatabases/{externalMySqlDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalMySqlDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/DeleteExternalMySqlDatabase"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "DeleteExternalMySqlDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalMySqlDatabaseConnector Deletes the Oracle Cloud Infrastructure resource representing an external MySQL database connector.
func (client ExternalMySqlDatabasesClient) DeleteExternalMySqlDatabaseConnector(ctx context.Context, request DeleteExternalMySqlDatabaseConnectorRequest) (response DeleteExternalMySqlDatabaseConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalMySqlDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalMySqlDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalMySqlDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalMySqlDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalMySqlDatabaseConnectorResponse")
	}
	return
}

// deleteExternalMySqlDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) deleteExternalMySqlDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/externalMySqlDatabaseConnectors/{externalMySqlDatabaseConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalMySqlDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabaseConnector/DeleteExternalMySqlDatabaseConnector"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "DeleteExternalMySqlDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalMySqlDatabaseManagement Disables Database Management for an external MySQL Database.
// A default retry strategy applies to this operation DisableExternalMySqlDatabaseManagement()
func (client ExternalMySqlDatabasesClient) DisableExternalMySqlDatabaseManagement(ctx context.Context, request DisableExternalMySqlDatabaseManagementRequest) (response DisableExternalMySqlDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalMySqlDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalMySqlDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalMySqlDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalMySqlDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalMySqlDatabaseManagementResponse")
	}
	return
}

// disableExternalMySqlDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) disableExternalMySqlDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalMySqlDatabases/{externalMySqlDatabaseId}/actions/disableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalMySqlDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/DisableExternalMySqlDatabaseManagement"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "DisableExternalMySqlDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExternalMySqlDatabaseManagement Enables Database Management for an external MySQL Database.
// A default retry strategy applies to this operation EnableExternalMySqlDatabaseManagement()
func (client ExternalMySqlDatabasesClient) EnableExternalMySqlDatabaseManagement(ctx context.Context, request EnableExternalMySqlDatabaseManagementRequest) (response EnableExternalMySqlDatabaseManagementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExternalMySqlDatabaseManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExternalMySqlDatabaseManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExternalMySqlDatabaseManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExternalMySqlDatabaseManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExternalMySqlDatabaseManagementResponse")
	}
	return
}

// enableExternalMySqlDatabaseManagement implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) enableExternalMySqlDatabaseManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/externalMySqlDatabases/{externalMySqlDatabaseId}/actions/enableDatabaseManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExternalMySqlDatabaseManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/EnableExternalMySqlDatabaseManagement"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "EnableExternalMySqlDatabaseManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalMySqlDatabase Retrieves the external MySQL database information.
// A default retry strategy applies to this operation GetExternalMySqlDatabase()
func (client ExternalMySqlDatabasesClient) GetExternalMySqlDatabase(ctx context.Context, request GetExternalMySqlDatabaseRequest) (response GetExternalMySqlDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalMySqlDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalMySqlDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalMySqlDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalMySqlDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalMySqlDatabaseResponse")
	}
	return
}

// getExternalMySqlDatabase implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) getExternalMySqlDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalMySqlDatabases/{externalMySqlDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalMySqlDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/GetExternalMySqlDatabase"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "GetExternalMySqlDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalMySqlDatabaseConnector Retrieves the MySQL database connector.
// A default retry strategy applies to this operation GetExternalMySqlDatabaseConnector()
func (client ExternalMySqlDatabasesClient) GetExternalMySqlDatabaseConnector(ctx context.Context, request GetExternalMySqlDatabaseConnectorRequest) (response GetExternalMySqlDatabaseConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalMySqlDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalMySqlDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalMySqlDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalMySqlDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalMySqlDatabaseConnectorResponse")
	}
	return
}

// getExternalMySqlDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) getExternalMySqlDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalMySqlDatabaseConnectors/{externalMySqlDatabaseConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalMySqlDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/GetExternalMySqlDatabaseConnector"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "GetExternalMySqlDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalMySqlDatabases Gets the list of External MySQL Databases.
func (client ExternalMySqlDatabasesClient) ListExternalMySqlDatabases(ctx context.Context, request ListExternalMySqlDatabasesRequest) (response ListExternalMySqlDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalMySqlDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalMySqlDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalMySqlDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalMySqlDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalMySqlDatabasesResponse")
	}
	return
}

// listExternalMySqlDatabases implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) listExternalMySqlDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalMySqlDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalMySqlDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabaseCollection/ListExternalMySqlDatabases"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "ListExternalMySqlDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMySqlDatabaseConnectors Gets the list of External MySQL Database connectors.
func (client ExternalMySqlDatabasesClient) ListMySqlDatabaseConnectors(ctx context.Context, request ListMySqlDatabaseConnectorsRequest) (response ListMySqlDatabaseConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMySqlDatabaseConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMySqlDatabaseConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMySqlDatabaseConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMySqlDatabaseConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMySqlDatabaseConnectorsResponse")
	}
	return
}

// listMySqlDatabaseConnectors implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) listMySqlDatabaseConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalMySqlDatabaseConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMySqlDatabaseConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/MySqlConnectorCollection/ListMySqlDatabaseConnectors"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "ListMySqlDatabaseConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalMysqlDatabase Updates the External Mysql Database.
func (client ExternalMySqlDatabasesClient) UpdateExternalMysqlDatabase(ctx context.Context, request UpdateExternalMysqlDatabaseRequest) (response UpdateExternalMysqlDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalMysqlDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalMysqlDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalMysqlDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalMysqlDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalMysqlDatabaseResponse")
	}
	return
}

// updateExternalMysqlDatabase implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) updateExternalMysqlDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalMySqlDatabases/{externalMySqlDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalMysqlDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabase/UpdateExternalMysqlDatabase"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "UpdateExternalMysqlDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalMysqlDatabaseConnector Updates the External Mysql Database Connector.
func (client ExternalMySqlDatabasesClient) UpdateExternalMysqlDatabaseConnector(ctx context.Context, request UpdateExternalMysqlDatabaseConnectorRequest) (response UpdateExternalMysqlDatabaseConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalMysqlDatabaseConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalMysqlDatabaseConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalMysqlDatabaseConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalMysqlDatabaseConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalMysqlDatabaseConnectorResponse")
	}
	return
}

// updateExternalMysqlDatabaseConnector implements the OCIOperation interface (enables retrying operations)
func (client ExternalMySqlDatabasesClient) updateExternalMysqlDatabaseConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/externalMySqlDatabaseConnectors/{externalMySqlDatabaseConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalMysqlDatabaseConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ExternalMySqlDatabaseConnector/UpdateExternalMysqlDatabaseConnector"
		err = common.PostProcessServiceError(err, "ExternalMySqlDatabases", "UpdateExternalMysqlDatabaseConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
