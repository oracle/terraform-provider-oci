// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DatabaseToolsClient a client for DatabaseTools
type DatabaseToolsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseToolsClientWithConfigurationProvider Creates a new default DatabaseTools client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseToolsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseToolsClient, err error) {
	if enabled := common.CheckForEnabledServices("databasetools"); !enabled {
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
	return newDatabaseToolsClientFromBaseClient(baseClient, provider)
}

// NewDatabaseToolsClientWithOboToken Creates a new default DatabaseTools client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatabaseToolsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseToolsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseToolsClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseToolsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseToolsClient, err error) {
	// DatabaseTools service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DatabaseTools"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatabaseToolsClient{BaseClient: baseClient}
	client.BasePath = "20201005"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseToolsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasetools", "https://dbtools.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseToolsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseToolsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddDatabaseToolsConnectionLock Adds a lock to a DatabaseToolsConnection resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsConnectionLock.go.html to see an example of how to use AddDatabaseToolsConnectionLock API.
func (client DatabaseToolsClient) AddDatabaseToolsConnectionLock(ctx context.Context, request AddDatabaseToolsConnectionLockRequest) (response AddDatabaseToolsConnectionLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsConnectionLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsConnectionLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsConnectionLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsConnectionLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsConnectionLockResponse")
	}
	return
}

// addDatabaseToolsConnectionLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsConnectionLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsConnectionLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsConnectionLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsConnectionLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsconnection{})
	return response, err
}

// AddDatabaseToolsDatabaseApiGatewayConfigLock Adds a lock to a DatabaseToolsDatabaseApiGatewayConfig resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsDatabaseApiGatewayConfigLock.go.html to see an example of how to use AddDatabaseToolsDatabaseApiGatewayConfigLock API.
func (client DatabaseToolsClient) AddDatabaseToolsDatabaseApiGatewayConfigLock(ctx context.Context, request AddDatabaseToolsDatabaseApiGatewayConfigLockRequest) (response AddDatabaseToolsDatabaseApiGatewayConfigLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsDatabaseApiGatewayConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsDatabaseApiGatewayConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsDatabaseApiGatewayConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsDatabaseApiGatewayConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsDatabaseApiGatewayConfigLockResponse")
	}
	return
}

// addDatabaseToolsDatabaseApiGatewayConfigLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsDatabaseApiGatewayConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsDatabaseApiGatewayConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsDatabaseApiGatewayConfigLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsDatabaseApiGatewayConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfig{})
	return response, err
}

// AddDatabaseToolsIdentityLock Adds a lock to a DatabaseToolsIdentity resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsIdentityLock.go.html to see an example of how to use AddDatabaseToolsIdentityLock API.
func (client DatabaseToolsClient) AddDatabaseToolsIdentityLock(ctx context.Context, request AddDatabaseToolsIdentityLockRequest) (response AddDatabaseToolsIdentityLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsIdentityLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsIdentityLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsIdentityLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsIdentityLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsIdentityLockResponse")
	}
	return
}

// addDatabaseToolsIdentityLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsIdentityLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities/{databaseToolsIdentityId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsIdentityLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsIdentityLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsIdentityLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsidentity{})
	return response, err
}

// AddDatabaseToolsMcpServerLock Adds a lock to a DatabaseToolsMcpServer resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsMcpServerLock.go.html to see an example of how to use AddDatabaseToolsMcpServerLock API.
func (client DatabaseToolsClient) AddDatabaseToolsMcpServerLock(ctx context.Context, request AddDatabaseToolsMcpServerLockRequest) (response AddDatabaseToolsMcpServerLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsMcpServerLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsMcpServerLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsMcpServerLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsMcpServerLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsMcpServerLockResponse")
	}
	return
}

// addDatabaseToolsMcpServerLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsMcpServerLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpServers/{databaseToolsMcpServerId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsMcpServerLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsMcpServerLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsMcpServerLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcpserver{})
	return response, err
}

// AddDatabaseToolsMcpToolsetLock Adds a lock to a DatabaseToolsMcpToolset resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsMcpToolsetLock.go.html to see an example of how to use AddDatabaseToolsMcpToolsetLock API.
func (client DatabaseToolsClient) AddDatabaseToolsMcpToolsetLock(ctx context.Context, request AddDatabaseToolsMcpToolsetLockRequest) (response AddDatabaseToolsMcpToolsetLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsMcpToolsetLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsMcpToolsetLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsMcpToolsetLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsMcpToolsetLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsMcpToolsetLockResponse")
	}
	return
}

// addDatabaseToolsMcpToolsetLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsMcpToolsetLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpToolsets/{databaseToolsMcpToolsetId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsMcpToolsetLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsMcpToolsetLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsMcpToolsetLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcptoolset{})
	return response, err
}

// AddDatabaseToolsPrivateEndpointLock Adds a lock to a DatabaseToolsPrivateEndpoint resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsPrivateEndpointLock.go.html to see an example of how to use AddDatabaseToolsPrivateEndpointLock API.
func (client DatabaseToolsClient) AddDatabaseToolsPrivateEndpointLock(ctx context.Context, request AddDatabaseToolsPrivateEndpointLockRequest) (response AddDatabaseToolsPrivateEndpointLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsPrivateEndpointLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsPrivateEndpointLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsPrivateEndpointLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsPrivateEndpointLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsPrivateEndpointLockResponse")
	}
	return
}

// addDatabaseToolsPrivateEndpointLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsPrivateEndpointLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsPrivateEndpoints/{databaseToolsPrivateEndpointId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsPrivateEndpointLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsPrivateEndpointLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsPrivateEndpointLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddDatabaseToolsSqlReportLock Adds a lock to a DatabaseToolsSqlReport resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/AddDatabaseToolsSqlReportLock.go.html to see an example of how to use AddDatabaseToolsSqlReportLock API.
func (client DatabaseToolsClient) AddDatabaseToolsSqlReportLock(ctx context.Context, request AddDatabaseToolsSqlReportLockRequest) (response AddDatabaseToolsSqlReportLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDatabaseToolsSqlReportLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatabaseToolsSqlReportLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatabaseToolsSqlReportLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatabaseToolsSqlReportLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatabaseToolsSqlReportLockResponse")
	}
	return
}

// addDatabaseToolsSqlReportLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) addDatabaseToolsSqlReportLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsSqlReports/{databaseToolsSqlReportId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatabaseToolsSqlReportLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "AddDatabaseToolsSqlReportLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "AddDatabaseToolsSqlReportLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolssqlreport{})
	return response, err
}

// CascadingDeleteDatabaseToolsMcpServer Deletes Database Tools McpServer resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CascadingDeleteDatabaseToolsMcpServer.go.html to see an example of how to use CascadingDeleteDatabaseToolsMcpServer API.
func (client DatabaseToolsClient) CascadingDeleteDatabaseToolsMcpServer(ctx context.Context, request CascadingDeleteDatabaseToolsMcpServerRequest) (response CascadingDeleteDatabaseToolsMcpServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cascadingDeleteDatabaseToolsMcpServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CascadingDeleteDatabaseToolsMcpServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CascadingDeleteDatabaseToolsMcpServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CascadingDeleteDatabaseToolsMcpServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CascadingDeleteDatabaseToolsMcpServerResponse")
	}
	return
}

// cascadingDeleteDatabaseToolsMcpServer implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) cascadingDeleteDatabaseToolsMcpServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpServers/{databaseToolsMcpServerId}/actions/cascadingDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CascadingDeleteDatabaseToolsMcpServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CascadingDeleteDatabaseToolsMcpServer")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CascadingDeleteDatabaseToolsMcpServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsConnectionCompartment Moves the specified Database Tools connection to a different compartment in the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsConnectionCompartment.go.html to see an example of how to use ChangeDatabaseToolsConnectionCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsConnectionCompartment(ctx context.Context, request ChangeDatabaseToolsConnectionCompartmentRequest) (response ChangeDatabaseToolsConnectionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsConnectionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsConnectionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsConnectionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsConnectionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsConnectionCompartmentResponse")
	}
	return
}

// changeDatabaseToolsConnectionCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsConnectionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsConnectionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsConnectionCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsConnectionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment Moves the specified Database Tools database API gateway config to a different compartment in the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/DbApiGatewayConfig/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment.go.html to see an example of how to use ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment(ctx context.Context, request ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentRequest) (response ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsDatabaseApiGatewayConfigCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentResponse")
	}
	return
}

// changeDatabaseToolsDatabaseApiGatewayConfigCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsDatabaseApiGatewayConfigCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsIdentityCompartment Moves the specified Database Tools identity to a different compartment in the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsIdentityCompartment.go.html to see an example of how to use ChangeDatabaseToolsIdentityCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsIdentityCompartment(ctx context.Context, request ChangeDatabaseToolsIdentityCompartmentRequest) (response ChangeDatabaseToolsIdentityCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsIdentityCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsIdentityCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsIdentityCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsIdentityCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsIdentityCompartmentResponse")
	}
	return
}

// changeDatabaseToolsIdentityCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsIdentityCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities/{databaseToolsIdentityId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsIdentityCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsIdentityCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsIdentityCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsMcpServerCompartment Moves the specified Database Tools mcpserver to a different compartment in the same tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsMcpServerCompartment.go.html to see an example of how to use ChangeDatabaseToolsMcpServerCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsMcpServerCompartment(ctx context.Context, request ChangeDatabaseToolsMcpServerCompartmentRequest) (response ChangeDatabaseToolsMcpServerCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsMcpServerCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsMcpServerCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsMcpServerCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsMcpServerCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsMcpServerCompartmentResponse")
	}
	return
}

// changeDatabaseToolsMcpServerCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsMcpServerCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpServers/{databaseToolsMcpServerId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsMcpServerCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsMcpServerCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsMcpServerCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsMcpToolsetCompartment Moves the specified Database Tools MCP Toolset to a different compartment in the same tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsMcpToolsetCompartment.go.html to see an example of how to use ChangeDatabaseToolsMcpToolsetCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsMcpToolsetCompartment(ctx context.Context, request ChangeDatabaseToolsMcpToolsetCompartmentRequest) (response ChangeDatabaseToolsMcpToolsetCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsMcpToolsetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsMcpToolsetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsMcpToolsetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsMcpToolsetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsMcpToolsetCompartmentResponse")
	}
	return
}

// changeDatabaseToolsMcpToolsetCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsMcpToolsetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpToolsets/{databaseToolsMcpToolsetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsMcpToolsetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsMcpToolsetCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsMcpToolsetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsPrivateEndpointCompartment Moves a Database Tools private endpoint into a different compartment in the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsPrivateEndpointCompartment.go.html to see an example of how to use ChangeDatabaseToolsPrivateEndpointCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsPrivateEndpointCompartment(ctx context.Context, request ChangeDatabaseToolsPrivateEndpointCompartmentRequest) (response ChangeDatabaseToolsPrivateEndpointCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsPrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsPrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsPrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsPrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsPrivateEndpointCompartmentResponse")
	}
	return
}

// changeDatabaseToolsPrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsPrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsPrivateEndpoints/{databaseToolsPrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsPrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsPrivateEndpointCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsPrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseToolsSqlReportCompartment Moves the specified Database Tools SQL Report to a different compartment in the same tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ChangeDatabaseToolsSqlReportCompartment.go.html to see an example of how to use ChangeDatabaseToolsSqlReportCompartment API.
func (client DatabaseToolsClient) ChangeDatabaseToolsSqlReportCompartment(ctx context.Context, request ChangeDatabaseToolsSqlReportCompartmentRequest) (response ChangeDatabaseToolsSqlReportCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseToolsSqlReportCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseToolsSqlReportCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseToolsSqlReportCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseToolsSqlReportCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseToolsSqlReportCompartmentResponse")
	}
	return
}

// changeDatabaseToolsSqlReportCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) changeDatabaseToolsSqlReportCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsSqlReports/{databaseToolsSqlReportId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseToolsSqlReportCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ChangeDatabaseToolsSqlReportCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ChangeDatabaseToolsSqlReportCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabaseToolsConnection Creates a new Database Tools connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsConnection.go.html to see an example of how to use CreateDatabaseToolsConnection API.
// A default retry strategy applies to this operation CreateDatabaseToolsConnection()
func (client DatabaseToolsClient) CreateDatabaseToolsConnection(ctx context.Context, request CreateDatabaseToolsConnectionRequest) (response CreateDatabaseToolsConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsConnectionResponse")
	}
	return
}

// createDatabaseToolsConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsconnection{})
	return response, err
}

// CreateDatabaseToolsDatabaseApiGatewayConfig Creates a new Database Tools database API gateway config.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsDatabaseApiGatewayConfig.go.html to see an example of how to use CreateDatabaseToolsDatabaseApiGatewayConfig API.
// A default retry strategy applies to this operation CreateDatabaseToolsDatabaseApiGatewayConfig()
func (client DatabaseToolsClient) CreateDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request CreateDatabaseToolsDatabaseApiGatewayConfigRequest) (response CreateDatabaseToolsDatabaseApiGatewayConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsDatabaseApiGatewayConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsDatabaseApiGatewayConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsDatabaseApiGatewayConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsDatabaseApiGatewayConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsDatabaseApiGatewayConfigResponse")
	}
	return
}

// createDatabaseToolsDatabaseApiGatewayConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsDatabaseApiGatewayConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsDatabaseApiGatewayConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsDatabaseApiGatewayConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfig{})
	return response, err
}

// CreateDatabaseToolsIdentity Creates a new Database Tools identity.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsIdentity.go.html to see an example of how to use CreateDatabaseToolsIdentity API.
// A default retry strategy applies to this operation CreateDatabaseToolsIdentity()
func (client DatabaseToolsClient) CreateDatabaseToolsIdentity(ctx context.Context, request CreateDatabaseToolsIdentityRequest) (response CreateDatabaseToolsIdentityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsIdentity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsIdentityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsIdentityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsIdentityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsIdentityResponse")
	}
	return
}

// createDatabaseToolsIdentity implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsIdentity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsIdentityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsIdentity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsIdentity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsidentity{})
	return response, err
}

// CreateDatabaseToolsMcpServer Creates a new Database Tools MCP server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsMcpServer.go.html to see an example of how to use CreateDatabaseToolsMcpServer API.
// A default retry strategy applies to this operation CreateDatabaseToolsMcpServer()
func (client DatabaseToolsClient) CreateDatabaseToolsMcpServer(ctx context.Context, request CreateDatabaseToolsMcpServerRequest) (response CreateDatabaseToolsMcpServerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsMcpServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsMcpServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsMcpServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsMcpServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsMcpServerResponse")
	}
	return
}

// createDatabaseToolsMcpServer implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsMcpServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpServers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsMcpServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsMcpServer")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsMcpServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcpserver{})
	return response, err
}

// CreateDatabaseToolsMcpToolset Creates a new Database Tools MCP Toolset.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsMcpToolset.go.html to see an example of how to use CreateDatabaseToolsMcpToolset API.
// A default retry strategy applies to this operation CreateDatabaseToolsMcpToolset()
func (client DatabaseToolsClient) CreateDatabaseToolsMcpToolset(ctx context.Context, request CreateDatabaseToolsMcpToolsetRequest) (response CreateDatabaseToolsMcpToolsetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsMcpToolset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsMcpToolsetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsMcpToolsetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsMcpToolsetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsMcpToolsetResponse")
	}
	return
}

// createDatabaseToolsMcpToolset implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsMcpToolset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpToolsets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsMcpToolsetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsMcpToolset")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsMcpToolset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcptoolset{})
	return response, err
}

// CreateDatabaseToolsPrivateEndpoint Creates a new Database Tools private endpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsPrivateEndpoint.go.html to see an example of how to use CreateDatabaseToolsPrivateEndpoint API.
// A default retry strategy applies to this operation CreateDatabaseToolsPrivateEndpoint()
func (client DatabaseToolsClient) CreateDatabaseToolsPrivateEndpoint(ctx context.Context, request CreateDatabaseToolsPrivateEndpointRequest) (response CreateDatabaseToolsPrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsPrivateEndpointResponse")
	}
	return
}

// createDatabaseToolsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabaseToolsSqlReport Creates a new Database Tools  Sql Report.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/CreateDatabaseToolsSqlReport.go.html to see an example of how to use CreateDatabaseToolsSqlReport API.
// A default retry strategy applies to this operation CreateDatabaseToolsSqlReport()
func (client DatabaseToolsClient) CreateDatabaseToolsSqlReport(ctx context.Context, request CreateDatabaseToolsSqlReportRequest) (response CreateDatabaseToolsSqlReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsSqlReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsSqlReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsSqlReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsSqlReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsSqlReportResponse")
	}
	return
}

// createDatabaseToolsSqlReport implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) createDatabaseToolsSqlReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsSqlReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsSqlReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "CreateDatabaseToolsSqlReport")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "CreateDatabaseToolsSqlReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolssqlreport{})
	return response, err
}

// DeleteDatabaseToolsConnection Deletes the specified Database Tools connection resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsConnection.go.html to see an example of how to use DeleteDatabaseToolsConnection API.
func (client DatabaseToolsClient) DeleteDatabaseToolsConnection(ctx context.Context, request DeleteDatabaseToolsConnectionRequest) (response DeleteDatabaseToolsConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsConnectionResponse")
	}
	return
}

// deleteDatabaseToolsConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsConnections/{databaseToolsConnectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsDatabaseApiGatewayConfig Deletes the specified Database Tools database API gateway config resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsDatabaseApiGatewayConfig.go.html to see an example of how to use DeleteDatabaseToolsDatabaseApiGatewayConfig API.
func (client DatabaseToolsClient) DeleteDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request DeleteDatabaseToolsDatabaseApiGatewayConfigRequest) (response DeleteDatabaseToolsDatabaseApiGatewayConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsDatabaseApiGatewayConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsDatabaseApiGatewayConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsDatabaseApiGatewayConfigResponse")
	}
	return
}

// deleteDatabaseToolsDatabaseApiGatewayConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsDatabaseApiGatewayConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsDatabaseApiGatewayConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsDatabaseApiGatewayConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsIdentity Deletes the specified Database Tools identity resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsIdentity.go.html to see an example of how to use DeleteDatabaseToolsIdentity API.
func (client DatabaseToolsClient) DeleteDatabaseToolsIdentity(ctx context.Context, request DeleteDatabaseToolsIdentityRequest) (response DeleteDatabaseToolsIdentityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsIdentity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsIdentityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsIdentityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsIdentityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsIdentityResponse")
	}
	return
}

// deleteDatabaseToolsIdentity implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsIdentity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsIdentities/{databaseToolsIdentityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsIdentityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsIdentity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsIdentity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsMcpServer Deletes the specified Database Tools MCP server resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsMcpServer.go.html to see an example of how to use DeleteDatabaseToolsMcpServer API.
func (client DatabaseToolsClient) DeleteDatabaseToolsMcpServer(ctx context.Context, request DeleteDatabaseToolsMcpServerRequest) (response DeleteDatabaseToolsMcpServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsMcpServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsMcpServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsMcpServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsMcpServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsMcpServerResponse")
	}
	return
}

// deleteDatabaseToolsMcpServer implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsMcpServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsMcpServers/{databaseToolsMcpServerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsMcpServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsMcpServer")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsMcpServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsMcpToolset Deletes the specified Database Tools MCP Toolset resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsMcpToolset.go.html to see an example of how to use DeleteDatabaseToolsMcpToolset API.
func (client DatabaseToolsClient) DeleteDatabaseToolsMcpToolset(ctx context.Context, request DeleteDatabaseToolsMcpToolsetRequest) (response DeleteDatabaseToolsMcpToolsetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsMcpToolset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsMcpToolsetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsMcpToolsetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsMcpToolsetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsMcpToolsetResponse")
	}
	return
}

// deleteDatabaseToolsMcpToolset implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsMcpToolset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsMcpToolsets/{databaseToolsMcpToolsetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsMcpToolsetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsMcpToolset")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsMcpToolset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsPrivateEndpoint Deletes the specified Database Tools private endpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsPrivateEndpoint.go.html to see an example of how to use DeleteDatabaseToolsPrivateEndpoint API.
func (client DatabaseToolsClient) DeleteDatabaseToolsPrivateEndpoint(ctx context.Context, request DeleteDatabaseToolsPrivateEndpointRequest) (response DeleteDatabaseToolsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsPrivateEndpointResponse")
	}
	return
}

// deleteDatabaseToolsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsPrivateEndpoints/{databaseToolsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsSqlReport Deletes the specified Database Tools SQL Report resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/DeleteDatabaseToolsSqlReport.go.html to see an example of how to use DeleteDatabaseToolsSqlReport API.
func (client DatabaseToolsClient) DeleteDatabaseToolsSqlReport(ctx context.Context, request DeleteDatabaseToolsSqlReportRequest) (response DeleteDatabaseToolsSqlReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsSqlReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsSqlReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsSqlReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsSqlReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsSqlReportResponse")
	}
	return
}

// deleteDatabaseToolsSqlReport implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) deleteDatabaseToolsSqlReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsSqlReports/{databaseToolsSqlReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsSqlReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "DeleteDatabaseToolsSqlReport")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "DeleteDatabaseToolsSqlReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseToolsConnection Gets details of the specified Database Tools connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsConnection.go.html to see an example of how to use GetDatabaseToolsConnection API.
// A default retry strategy applies to this operation GetDatabaseToolsConnection()
func (client DatabaseToolsClient) GetDatabaseToolsConnection(ctx context.Context, request GetDatabaseToolsConnectionRequest) (response GetDatabaseToolsConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsConnectionResponse")
	}
	return
}

// getDatabaseToolsConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsconnection{})
	return response, err
}

// GetDatabaseToolsDatabaseApiGatewayConfig Gets details of the specified Database Tools database API gateway config.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsDatabaseApiGatewayConfig.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfig API.
// A default retry strategy applies to this operation GetDatabaseToolsDatabaseApiGatewayConfig()
func (client DatabaseToolsClient) GetDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request GetDatabaseToolsDatabaseApiGatewayConfigRequest) (response GetDatabaseToolsDatabaseApiGatewayConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsDatabaseApiGatewayConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsDatabaseApiGatewayConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsDatabaseApiGatewayConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsDatabaseApiGatewayConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsDatabaseApiGatewayConfigResponse")
	}
	return
}

// getDatabaseToolsDatabaseApiGatewayConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsDatabaseApiGatewayConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsDatabaseApiGatewayConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsDatabaseApiGatewayConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfig{})
	return response, err
}

// GetDatabaseToolsEndpointService Gets details for the specified Database Tools endpoint service.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsEndpointService.go.html to see an example of how to use GetDatabaseToolsEndpointService API.
// A default retry strategy applies to this operation GetDatabaseToolsEndpointService()
func (client DatabaseToolsClient) GetDatabaseToolsEndpointService(ctx context.Context, request GetDatabaseToolsEndpointServiceRequest) (response GetDatabaseToolsEndpointServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsEndpointService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsEndpointServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsEndpointServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsEndpointServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsEndpointServiceResponse")
	}
	return
}

// getDatabaseToolsEndpointService implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsEndpointService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsEndpointServices/{databaseToolsEndpointServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsEndpointServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsEndpointService")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsEndpointService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseToolsIdentity Gets details of the specified Database Tools identity.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsIdentity.go.html to see an example of how to use GetDatabaseToolsIdentity API.
// A default retry strategy applies to this operation GetDatabaseToolsIdentity()
func (client DatabaseToolsClient) GetDatabaseToolsIdentity(ctx context.Context, request GetDatabaseToolsIdentityRequest) (response GetDatabaseToolsIdentityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsIdentity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsIdentityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsIdentityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsIdentityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsIdentityResponse")
	}
	return
}

// getDatabaseToolsIdentity implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsIdentity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsIdentities/{databaseToolsIdentityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsIdentityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsIdentity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsIdentity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsidentity{})
	return response, err
}

// GetDatabaseToolsMcpServer Gets details of the specified Database Tools MCP server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsMcpServer.go.html to see an example of how to use GetDatabaseToolsMcpServer API.
// A default retry strategy applies to this operation GetDatabaseToolsMcpServer()
func (client DatabaseToolsClient) GetDatabaseToolsMcpServer(ctx context.Context, request GetDatabaseToolsMcpServerRequest) (response GetDatabaseToolsMcpServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsMcpServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsMcpServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsMcpServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsMcpServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsMcpServerResponse")
	}
	return
}

// getDatabaseToolsMcpServer implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsMcpServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsMcpServers/{databaseToolsMcpServerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsMcpServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsMcpServer")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsMcpServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcpserver{})
	return response, err
}

// GetDatabaseToolsMcpToolset Gets details of the specified Database Tools MCP Toolset.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsMcpToolset.go.html to see an example of how to use GetDatabaseToolsMcpToolset API.
// A default retry strategy applies to this operation GetDatabaseToolsMcpToolset()
func (client DatabaseToolsClient) GetDatabaseToolsMcpToolset(ctx context.Context, request GetDatabaseToolsMcpToolsetRequest) (response GetDatabaseToolsMcpToolsetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsMcpToolset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsMcpToolsetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsMcpToolsetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsMcpToolsetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsMcpToolsetResponse")
	}
	return
}

// getDatabaseToolsMcpToolset implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsMcpToolset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsMcpToolsets/{databaseToolsMcpToolsetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsMcpToolsetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsMcpToolset")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsMcpToolset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcptoolset{})
	return response, err
}

// GetDatabaseToolsPrivateEndpoint Gets details of a specified Database Tools private endpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsPrivateEndpoint.go.html to see an example of how to use GetDatabaseToolsPrivateEndpoint API.
// A default retry strategy applies to this operation GetDatabaseToolsPrivateEndpoint()
func (client DatabaseToolsClient) GetDatabaseToolsPrivateEndpoint(ctx context.Context, request GetDatabaseToolsPrivateEndpointRequest) (response GetDatabaseToolsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsPrivateEndpointResponse")
	}
	return
}

// getDatabaseToolsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsPrivateEndpoints/{databaseToolsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseToolsSqlReport Gets details of the specified Database Tools SQL report.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetDatabaseToolsSqlReport.go.html to see an example of how to use GetDatabaseToolsSqlReport API.
// A default retry strategy applies to this operation GetDatabaseToolsSqlReport()
func (client DatabaseToolsClient) GetDatabaseToolsSqlReport(ctx context.Context, request GetDatabaseToolsSqlReportRequest) (response GetDatabaseToolsSqlReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsSqlReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsSqlReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsSqlReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsSqlReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsSqlReportResponse")
	}
	return
}

// getDatabaseToolsSqlReport implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) getDatabaseToolsSqlReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsSqlReports/{databaseToolsSqlReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsSqlReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetDatabaseToolsSqlReport")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetDatabaseToolsSqlReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolssqlreport{})
	return response, err
}

// GetWorkRequest Gets the status of the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DatabaseToolsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseToolsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "GetWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsConnections Returns a list of Database Tools connections.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsConnections.go.html to see an example of how to use ListDatabaseToolsConnections API.
// A default retry strategy applies to this operation ListDatabaseToolsConnections()
func (client DatabaseToolsClient) ListDatabaseToolsConnections(ctx context.Context, request ListDatabaseToolsConnectionsRequest) (response ListDatabaseToolsConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsConnectionsResponse")
	}
	return
}

// listDatabaseToolsConnections implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsConnections")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsDatabaseApiGatewayConfigs Returns a list of Database Tools database API gateway configs.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsDatabaseApiGatewayConfigs.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigs API.
// A default retry strategy applies to this operation ListDatabaseToolsDatabaseApiGatewayConfigs()
func (client DatabaseToolsClient) ListDatabaseToolsDatabaseApiGatewayConfigs(ctx context.Context, request ListDatabaseToolsDatabaseApiGatewayConfigsRequest) (response ListDatabaseToolsDatabaseApiGatewayConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsDatabaseApiGatewayConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsDatabaseApiGatewayConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsDatabaseApiGatewayConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsDatabaseApiGatewayConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsDatabaseApiGatewayConfigsResponse")
	}
	return
}

// listDatabaseToolsDatabaseApiGatewayConfigs implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsDatabaseApiGatewayConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsDatabaseApiGatewayConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsDatabaseApiGatewayConfigs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsDatabaseApiGatewayConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsEndpointServices Returns a list of Database Tools endpoint services.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsEndpointServices.go.html to see an example of how to use ListDatabaseToolsEndpointServices API.
// A default retry strategy applies to this operation ListDatabaseToolsEndpointServices()
func (client DatabaseToolsClient) ListDatabaseToolsEndpointServices(ctx context.Context, request ListDatabaseToolsEndpointServicesRequest) (response ListDatabaseToolsEndpointServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsEndpointServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsEndpointServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsEndpointServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsEndpointServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsEndpointServicesResponse")
	}
	return
}

// listDatabaseToolsEndpointServices implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsEndpointServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsEndpointServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsEndpointServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsEndpointServices")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsEndpointServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsIdentities Returns a list of Database Tools identities.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsIdentities.go.html to see an example of how to use ListDatabaseToolsIdentities API.
// A default retry strategy applies to this operation ListDatabaseToolsIdentities()
func (client DatabaseToolsClient) ListDatabaseToolsIdentities(ctx context.Context, request ListDatabaseToolsIdentitiesRequest) (response ListDatabaseToolsIdentitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsIdentities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsIdentitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsIdentitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsIdentitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsIdentitiesResponse")
	}
	return
}

// listDatabaseToolsIdentities implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsIdentities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsIdentities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsIdentitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsIdentities")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsIdentities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsMcpServers Returns a list of Database Tools MCP servers.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsMcpServers.go.html to see an example of how to use ListDatabaseToolsMcpServers API.
// A default retry strategy applies to this operation ListDatabaseToolsMcpServers()
func (client DatabaseToolsClient) ListDatabaseToolsMcpServers(ctx context.Context, request ListDatabaseToolsMcpServersRequest) (response ListDatabaseToolsMcpServersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsMcpServers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsMcpServersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsMcpServersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsMcpServersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsMcpServersResponse")
	}
	return
}

// listDatabaseToolsMcpServers implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsMcpServers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsMcpServers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsMcpServersResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsMcpServers")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsMcpServers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsMcpToolsetVersions Returns a list of Database Tools Toolset versions
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsMcpToolsetVersions.go.html to see an example of how to use ListDatabaseToolsMcpToolsetVersions API.
// A default retry strategy applies to this operation ListDatabaseToolsMcpToolsetVersions()
func (client DatabaseToolsClient) ListDatabaseToolsMcpToolsetVersions(ctx context.Context, request ListDatabaseToolsMcpToolsetVersionsRequest) (response ListDatabaseToolsMcpToolsetVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsMcpToolsetVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsMcpToolsetVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsMcpToolsetVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsMcpToolsetVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsMcpToolsetVersionsResponse")
	}
	return
}

// listDatabaseToolsMcpToolsetVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsMcpToolsetVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsMcpToolsetVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsMcpToolsetVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsMcpToolsetVersions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsMcpToolsetVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsMcpToolsets Returns a list of Database Tools Toolsets.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsMcpToolsets.go.html to see an example of how to use ListDatabaseToolsMcpToolsets API.
// A default retry strategy applies to this operation ListDatabaseToolsMcpToolsets()
func (client DatabaseToolsClient) ListDatabaseToolsMcpToolsets(ctx context.Context, request ListDatabaseToolsMcpToolsetsRequest) (response ListDatabaseToolsMcpToolsetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsMcpToolsets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsMcpToolsetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsMcpToolsetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsMcpToolsetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsMcpToolsetsResponse")
	}
	return
}

// listDatabaseToolsMcpToolsets implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsMcpToolsets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsMcpToolsets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsMcpToolsetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsMcpToolsets")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsMcpToolsets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsPrivateEndpoints Returns a list of Database Tools private endpoints.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsPrivateEndpoints.go.html to see an example of how to use ListDatabaseToolsPrivateEndpoints API.
// A default retry strategy applies to this operation ListDatabaseToolsPrivateEndpoints()
func (client DatabaseToolsClient) ListDatabaseToolsPrivateEndpoints(ctx context.Context, request ListDatabaseToolsPrivateEndpointsRequest) (response ListDatabaseToolsPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsPrivateEndpointsResponse")
	}
	return
}

// listDatabaseToolsPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsPrivateEndpoints")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsSqlReports Returns a list of Database Tools SQL reports.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsSqlReports.go.html to see an example of how to use ListDatabaseToolsSqlReports API.
// A default retry strategy applies to this operation ListDatabaseToolsSqlReports()
func (client DatabaseToolsClient) ListDatabaseToolsSqlReports(ctx context.Context, request ListDatabaseToolsSqlReportsRequest) (response ListDatabaseToolsSqlReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsSqlReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsSqlReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsSqlReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsSqlReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsSqlReportsResponse")
	}
	return
}

// listDatabaseToolsSqlReports implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) listDatabaseToolsSqlReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsSqlReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsSqlReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListDatabaseToolsSqlReports")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListDatabaseToolsSqlReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a paginated list of errors for the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DatabaseToolsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseToolsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DatabaseToolsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseToolsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DatabaseToolsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseToolsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ListWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshDatabaseToolsIdentityCredential Refresh Database Tools identity credential.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RefreshDatabaseToolsIdentityCredential.go.html to see an example of how to use RefreshDatabaseToolsIdentityCredential API.
func (client DatabaseToolsClient) RefreshDatabaseToolsIdentityCredential(ctx context.Context, request RefreshDatabaseToolsIdentityCredentialRequest) (response RefreshDatabaseToolsIdentityCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.refreshDatabaseToolsIdentityCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshDatabaseToolsIdentityCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshDatabaseToolsIdentityCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshDatabaseToolsIdentityCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshDatabaseToolsIdentityCredentialResponse")
	}
	return
}

// refreshDatabaseToolsIdentityCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) refreshDatabaseToolsIdentityCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities/{databaseToolsIdentityId}/actions/refreshCredential", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshDatabaseToolsIdentityCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RefreshDatabaseToolsIdentityCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RefreshDatabaseToolsIdentityCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDatabaseToolsConnectionLock Removes a lock from a DatabaseToolsConnection resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsConnectionLock.go.html to see an example of how to use RemoveDatabaseToolsConnectionLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsConnectionLock(ctx context.Context, request RemoveDatabaseToolsConnectionLockRequest) (response RemoveDatabaseToolsConnectionLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsConnectionLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsConnectionLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsConnectionLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsConnectionLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsConnectionLockResponse")
	}
	return
}

// removeDatabaseToolsConnectionLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsConnectionLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsConnectionLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsConnectionLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsConnectionLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsconnection{})
	return response, err
}

// RemoveDatabaseToolsDatabaseApiGatewayConfigLock Removes a lock from a DatabaseToolsDatabaseApiGatewayConfig resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsDatabaseApiGatewayConfigLock.go.html to see an example of how to use RemoveDatabaseToolsDatabaseApiGatewayConfigLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsDatabaseApiGatewayConfigLock(ctx context.Context, request RemoveDatabaseToolsDatabaseApiGatewayConfigLockRequest) (response RemoveDatabaseToolsDatabaseApiGatewayConfigLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsDatabaseApiGatewayConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsDatabaseApiGatewayConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsDatabaseApiGatewayConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsDatabaseApiGatewayConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsDatabaseApiGatewayConfigLockResponse")
	}
	return
}

// removeDatabaseToolsDatabaseApiGatewayConfigLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsDatabaseApiGatewayConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsDatabaseApiGatewayConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsDatabaseApiGatewayConfigLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsDatabaseApiGatewayConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfig{})
	return response, err
}

// RemoveDatabaseToolsIdentityLock Removes a lock from a DatabaseToolsIdentity resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsIdentityLock.go.html to see an example of how to use RemoveDatabaseToolsIdentityLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsIdentityLock(ctx context.Context, request RemoveDatabaseToolsIdentityLockRequest) (response RemoveDatabaseToolsIdentityLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsIdentityLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsIdentityLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsIdentityLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsIdentityLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsIdentityLockResponse")
	}
	return
}

// removeDatabaseToolsIdentityLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsIdentityLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities/{databaseToolsIdentityId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsIdentityLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsIdentityLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsIdentityLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsidentity{})
	return response, err
}

// RemoveDatabaseToolsMcpServerLock Removes a lock from a DatabaseToolsMcpServer resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsMcpServerLock.go.html to see an example of how to use RemoveDatabaseToolsMcpServerLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsMcpServerLock(ctx context.Context, request RemoveDatabaseToolsMcpServerLockRequest) (response RemoveDatabaseToolsMcpServerLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsMcpServerLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsMcpServerLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsMcpServerLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsMcpServerLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsMcpServerLockResponse")
	}
	return
}

// removeDatabaseToolsMcpServerLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsMcpServerLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpServers/{databaseToolsMcpServerId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsMcpServerLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsMcpServerLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsMcpServerLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcpserver{})
	return response, err
}

// RemoveDatabaseToolsMcpToolsetLock Removes a lock from a DatabaseToolsMcpToolset resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsMcpToolsetLock.go.html to see an example of how to use RemoveDatabaseToolsMcpToolsetLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsMcpToolsetLock(ctx context.Context, request RemoveDatabaseToolsMcpToolsetLockRequest) (response RemoveDatabaseToolsMcpToolsetLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsMcpToolsetLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsMcpToolsetLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsMcpToolsetLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsMcpToolsetLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsMcpToolsetLockResponse")
	}
	return
}

// removeDatabaseToolsMcpToolsetLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsMcpToolsetLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsMcpToolsets/{databaseToolsMcpToolsetId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsMcpToolsetLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsMcpToolsetLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsMcpToolsetLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsmcptoolset{})
	return response, err
}

// RemoveDatabaseToolsPrivateEndpointLock Removes a lock from a DatabaseToolsPrivateEndpoint resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsPrivateEndpointLock.go.html to see an example of how to use RemoveDatabaseToolsPrivateEndpointLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsPrivateEndpointLock(ctx context.Context, request RemoveDatabaseToolsPrivateEndpointLockRequest) (response RemoveDatabaseToolsPrivateEndpointLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsPrivateEndpointLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsPrivateEndpointLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsPrivateEndpointLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsPrivateEndpointLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsPrivateEndpointLockResponse")
	}
	return
}

// removeDatabaseToolsPrivateEndpointLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsPrivateEndpointLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsPrivateEndpoints/{databaseToolsPrivateEndpointId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsPrivateEndpointLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsPrivateEndpointLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsPrivateEndpointLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDatabaseToolsSqlReportLock Removes a lock from a DatabaseToolsSqlReport resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/RemoveDatabaseToolsSqlReportLock.go.html to see an example of how to use RemoveDatabaseToolsSqlReportLock API.
func (client DatabaseToolsClient) RemoveDatabaseToolsSqlReportLock(ctx context.Context, request RemoveDatabaseToolsSqlReportLockRequest) (response RemoveDatabaseToolsSqlReportLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDatabaseToolsSqlReportLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatabaseToolsSqlReportLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatabaseToolsSqlReportLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatabaseToolsSqlReportLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatabaseToolsSqlReportLockResponse")
	}
	return
}

// removeDatabaseToolsSqlReportLock implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) removeDatabaseToolsSqlReportLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsSqlReports/{databaseToolsSqlReportId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatabaseToolsSqlReportLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "RemoveDatabaseToolsSqlReportLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "RemoveDatabaseToolsSqlReportLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolssqlreport{})
	return response, err
}

// UpdateDatabaseToolsConnection Updates the specified Database Tools connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsConnection.go.html to see an example of how to use UpdateDatabaseToolsConnection API.
func (client DatabaseToolsClient) UpdateDatabaseToolsConnection(ctx context.Context, request UpdateDatabaseToolsConnectionRequest) (response UpdateDatabaseToolsConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsConnectionResponse")
	}
	return
}

// updateDatabaseToolsConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsConnections/{databaseToolsConnectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseToolsDatabaseApiGatewayConfig Updates the specified Database Tools database API gateway config.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsDatabaseApiGatewayConfig.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfig API.
func (client DatabaseToolsClient) UpdateDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request UpdateDatabaseToolsDatabaseApiGatewayConfigRequest) (response UpdateDatabaseToolsDatabaseApiGatewayConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsDatabaseApiGatewayConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsDatabaseApiGatewayConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsDatabaseApiGatewayConfigResponse")
	}
	return
}

// updateDatabaseToolsDatabaseApiGatewayConfig implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsDatabaseApiGatewayConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsDatabaseApiGatewayConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsDatabaseApiGatewayConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsDatabaseApiGatewayConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfig{})
	return response, err
}

// UpdateDatabaseToolsIdentity Updates the specified Database Tools identity.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsIdentity.go.html to see an example of how to use UpdateDatabaseToolsIdentity API.
func (client DatabaseToolsClient) UpdateDatabaseToolsIdentity(ctx context.Context, request UpdateDatabaseToolsIdentityRequest) (response UpdateDatabaseToolsIdentityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsIdentity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsIdentityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsIdentityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsIdentityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsIdentityResponse")
	}
	return
}

// updateDatabaseToolsIdentity implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsIdentity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsIdentities/{databaseToolsIdentityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsIdentityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsIdentity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsIdentity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseToolsMcpServer Updates the specified Database Tools MCP server.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsMcpServer.go.html to see an example of how to use UpdateDatabaseToolsMcpServer API.
func (client DatabaseToolsClient) UpdateDatabaseToolsMcpServer(ctx context.Context, request UpdateDatabaseToolsMcpServerRequest) (response UpdateDatabaseToolsMcpServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsMcpServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsMcpServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsMcpServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsMcpServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsMcpServerResponse")
	}
	return
}

// updateDatabaseToolsMcpServer implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsMcpServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsMcpServers/{databaseToolsMcpServerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsMcpServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsMcpServer")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsMcpServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseToolsMcpToolset Updates the specified Database Tools MCP Toolset.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsMcpToolset.go.html to see an example of how to use UpdateDatabaseToolsMcpToolset API.
func (client DatabaseToolsClient) UpdateDatabaseToolsMcpToolset(ctx context.Context, request UpdateDatabaseToolsMcpToolsetRequest) (response UpdateDatabaseToolsMcpToolsetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsMcpToolset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsMcpToolsetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsMcpToolsetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsMcpToolsetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsMcpToolsetResponse")
	}
	return
}

// updateDatabaseToolsMcpToolset implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsMcpToolset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsMcpToolsets/{databaseToolsMcpToolsetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsMcpToolsetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsMcpToolset")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsMcpToolset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseToolsPrivateEndpoint Updates the specified Database Tools private endpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsPrivateEndpoint.go.html to see an example of how to use UpdateDatabaseToolsPrivateEndpoint API.
func (client DatabaseToolsClient) UpdateDatabaseToolsPrivateEndpoint(ctx context.Context, request UpdateDatabaseToolsPrivateEndpointRequest) (response UpdateDatabaseToolsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsPrivateEndpointResponse")
	}
	return
}

// updateDatabaseToolsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsPrivateEndpoints/{databaseToolsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseToolsSqlReport Updates the specified Database Tools SQL Report.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/UpdateDatabaseToolsSqlReport.go.html to see an example of how to use UpdateDatabaseToolsSqlReport API.
func (client DatabaseToolsClient) UpdateDatabaseToolsSqlReport(ctx context.Context, request UpdateDatabaseToolsSqlReportRequest) (response UpdateDatabaseToolsSqlReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsSqlReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsSqlReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsSqlReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsSqlReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsSqlReportResponse")
	}
	return
}

// updateDatabaseToolsSqlReport implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) updateDatabaseToolsSqlReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsSqlReports/{databaseToolsSqlReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsSqlReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "UpdateDatabaseToolsSqlReport")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "UpdateDatabaseToolsSqlReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolssqlreport{})
	return response, err
}

// ValidateDatabaseToolsConnection Validates the Database Tools connection details by establishing a connection to the database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ValidateDatabaseToolsConnection.go.html to see an example of how to use ValidateDatabaseToolsConnection API.
func (client DatabaseToolsClient) ValidateDatabaseToolsConnection(ctx context.Context, request ValidateDatabaseToolsConnectionRequest) (response ValidateDatabaseToolsConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateDatabaseToolsConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateDatabaseToolsConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateDatabaseToolsConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateDatabaseToolsConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateDatabaseToolsConnectionResponse")
	}
	return
}

// validateDatabaseToolsConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) validateDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/actions/validateConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ValidateDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ValidateDatabaseToolsConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &validatedatabasetoolsconnectionresult{})
	return response, err
}

// ValidateDatabaseToolsIdentityCredential Validates the Database Tools identity credentials by establishing a connection to the customer database
// and executing the dbms_cloud.send_request to validate the credential.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ValidateDatabaseToolsIdentityCredential.go.html to see an example of how to use ValidateDatabaseToolsIdentityCredential API.
func (client DatabaseToolsClient) ValidateDatabaseToolsIdentityCredential(ctx context.Context, request ValidateDatabaseToolsIdentityCredentialRequest) (response ValidateDatabaseToolsIdentityCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateDatabaseToolsIdentityCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateDatabaseToolsIdentityCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateDatabaseToolsIdentityCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateDatabaseToolsIdentityCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateDatabaseToolsIdentityCredentialResponse")
	}
	return
}

// validateDatabaseToolsIdentityCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsClient) validateDatabaseToolsIdentityCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities/{databaseToolsIdentityId}/actions/validateCredential", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDatabaseToolsIdentityCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseTools", "ValidateDatabaseToolsIdentityCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseTools", "ValidateDatabaseToolsIdentityCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &validatedatabasetoolsidentitycredentialresult{})
	return response, err
}
