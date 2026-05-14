// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DatabaseToolsRuntimeClient a client for DatabaseToolsRuntime
type DatabaseToolsRuntimeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseToolsRuntimeClientWithConfigurationProvider Creates a new default DatabaseToolsRuntime client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseToolsRuntimeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseToolsRuntimeClient, err error) {
	if enabled := common.CheckForEnabledServices("databasetoolsruntime"); !enabled {
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
	return newDatabaseToolsRuntimeClientFromBaseClient(baseClient, provider)
}

// NewDatabaseToolsRuntimeClientWithOboToken Creates a new default DatabaseToolsRuntime client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatabaseToolsRuntimeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseToolsRuntimeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseToolsRuntimeClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseToolsRuntimeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseToolsRuntimeClient, err error) {
	// DatabaseToolsRuntime service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DatabaseToolsRuntime"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatabaseToolsRuntimeClient{BaseClient: baseClient}
	client.BasePath = "20230222"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseToolsRuntimeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasetoolsruntime", "https://dbtools.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseToolsRuntimeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseToolsRuntimeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Attempts to cancel the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client DatabaseToolsRuntimeClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CancelWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCredential Creates a credential for the user specified by the key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateCredential.go.html to see an example of how to use CreateCredential API.
// A default retry strategy applies to this operation CreateCredential()
func (client DatabaseToolsRuntimeClient) CreateCredential(ctx context.Context, request CreateCredentialRequest) (response CreateCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCredentialResponse")
	}
	return
}

// createCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) createCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CreateCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CreateCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCredentialExecuteGrantee Grants the EXECUTE privilege on the credential to the user specified by the key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateCredentialExecuteGrantee.go.html to see an example of how to use CreateCredentialExecuteGrantee API.
// A default retry strategy applies to this operation CreateCredentialExecuteGrantee()
func (client DatabaseToolsRuntimeClient) CreateCredentialExecuteGrantee(ctx context.Context, request CreateCredentialExecuteGranteeRequest) (response CreateCredentialExecuteGranteeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createCredentialExecuteGrantee, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCredentialExecuteGranteeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCredentialExecuteGranteeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCredentialExecuteGranteeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCredentialExecuteGranteeResponse")
	}
	return
}

// createCredentialExecuteGrantee implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) createCredentialExecuteGrantee(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/executeGrantees", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCredentialExecuteGranteeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CreateCredentialExecuteGrantee")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CreateCredentialExecuteGrantee", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCredentialPublicSynonym Creates a public synonym for the given credentials
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateCredentialPublicSynonym.go.html to see an example of how to use CreateCredentialPublicSynonym API.
// A default retry strategy applies to this operation CreateCredentialPublicSynonym()
func (client DatabaseToolsRuntimeClient) CreateCredentialPublicSynonym(ctx context.Context, request CreateCredentialPublicSynonymRequest) (response CreateCredentialPublicSynonymResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createCredentialPublicSynonym, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCredentialPublicSynonymResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCredentialPublicSynonymResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCredentialPublicSynonymResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCredentialPublicSynonymResponse")
	}
	return
}

// createCredentialPublicSynonym implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) createCredentialPublicSynonym(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/publicSynonyms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCredentialPublicSynonymResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CreateCredentialPublicSynonym")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CreateCredentialPublicSynonym", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPool Create a Database Tools database API gateway config pool resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateDatabaseToolsDatabaseApiGatewayConfigPool.go.html to see an example of how to use CreateDatabaseToolsDatabaseApiGatewayConfigPool API.
// A default retry strategy applies to this operation CreateDatabaseToolsDatabaseApiGatewayConfigPool()
func (client DatabaseToolsRuntimeClient) CreateDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) (response CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsDatabaseApiGatewayConfigPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse")
	}
	return
}

// createDatabaseToolsDatabaseApiGatewayConfigPool implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) createDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CreateDatabaseToolsDatabaseApiGatewayConfigPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CreateDatabaseToolsDatabaseApiGatewayConfigPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpool{})
	return response, err
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec Create a Database Tools database API gateway config API spec resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec.go.html to see an example of how to use CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec API.
// A default retry strategy applies to this operation CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec()
func (client DatabaseToolsRuntimeClient) CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest) (response CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse")
	}
	return
}

// createDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) createDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/apiSpecs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpoolapispec{})
	return response, err
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec Create a Database Tools database API gateway config auto API spec resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.go.html to see an example of how to use CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec API.
// A default retry strategy applies to this operation CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec()
func (client DatabaseToolsRuntimeClient) CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) (response CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse")
	}
	return
}

// createDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) createDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/autoApiSpecs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpoolautoapispec{})
	return response, err
}

// DeleteCredential Delete credential
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/DeleteCredential.go.html to see an example of how to use DeleteCredential API.
// A default retry strategy applies to this operation DeleteCredential()
func (client DatabaseToolsRuntimeClient) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) (response DeleteCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCredentialResponse")
	}
	return
}

// deleteCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) deleteCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "DeleteCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "DeleteCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCredentialExecuteGrantee Delete execute grantee
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/DeleteCredentialExecuteGrantee.go.html to see an example of how to use DeleteCredentialExecuteGrantee API.
// A default retry strategy applies to this operation DeleteCredentialExecuteGrantee()
func (client DatabaseToolsRuntimeClient) DeleteCredentialExecuteGrantee(ctx context.Context, request DeleteCredentialExecuteGranteeRequest) (response DeleteCredentialExecuteGranteeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCredentialExecuteGrantee, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCredentialExecuteGranteeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCredentialExecuteGranteeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCredentialExecuteGranteeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCredentialExecuteGranteeResponse")
	}
	return
}

// deleteCredentialExecuteGrantee implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) deleteCredentialExecuteGrantee(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/executeGrantees/{executeGranteeKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCredentialExecuteGranteeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "DeleteCredentialExecuteGrantee")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "DeleteCredentialExecuteGrantee", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCredentialPublicSynonym Deletes the public synonym
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/DeleteCredentialPublicSynonym.go.html to see an example of how to use DeleteCredentialPublicSynonym API.
// A default retry strategy applies to this operation DeleteCredentialPublicSynonym()
func (client DatabaseToolsRuntimeClient) DeleteCredentialPublicSynonym(ctx context.Context, request DeleteCredentialPublicSynonymRequest) (response DeleteCredentialPublicSynonymResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCredentialPublicSynonym, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCredentialPublicSynonymResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCredentialPublicSynonymResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCredentialPublicSynonymResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCredentialPublicSynonymResponse")
	}
	return
}

// deleteCredentialPublicSynonym implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) deleteCredentialPublicSynonym(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/publicSynonyms/{publicSynonymKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCredentialPublicSynonymResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "DeleteCredentialPublicSynonym")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "DeleteCredentialPublicSynonym", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsDatabaseApiGatewayConfigPool Deletes the specified Database Tools database API gateway config pool resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/DeleteDatabaseToolsDatabaseApiGatewayConfigPool.go.html to see an example of how to use DeleteDatabaseToolsDatabaseApiGatewayConfigPool API.
// A default retry strategy applies to this operation DeleteDatabaseToolsDatabaseApiGatewayConfigPool()
func (client DatabaseToolsRuntimeClient) DeleteDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request DeleteDatabaseToolsDatabaseApiGatewayConfigPoolRequest) (response DeleteDatabaseToolsDatabaseApiGatewayConfigPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsDatabaseApiGatewayConfigPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsDatabaseApiGatewayConfigPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsDatabaseApiGatewayConfigPoolResponse")
	}
	return
}

// deleteDatabaseToolsDatabaseApiGatewayConfigPool implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) deleteDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsDatabaseApiGatewayConfigPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "DeleteDatabaseToolsDatabaseApiGatewayConfigPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "DeleteDatabaseToolsDatabaseApiGatewayConfigPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec Deletes the specified Database Tools database API gateway config API spec resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec.go.html to see an example of how to use DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec API.
// A default retry strategy applies to this operation DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec()
func (client DatabaseToolsRuntimeClient) DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest) (response DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse")
	}
	return
}

// deleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) deleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/apiSpecs/{apiSpecKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec Deletes the specified Database Tools database API gateway config auto API spec resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.go.html to see an example of how to use DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec API.
// A default retry strategy applies to this operation DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec()
func (client DatabaseToolsRuntimeClient) DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) (response DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse")
	}
	return
}

// deleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) deleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/autoApiSpecs/{autoApiSpecKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExecuteSqlDatabaseToolsConnection Execute statements on a database tools connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ExecuteSqlDatabaseToolsConnection.go.html to see an example of how to use ExecuteSqlDatabaseToolsConnection API.
// A default retry strategy applies to this operation ExecuteSqlDatabaseToolsConnection()
func (client DatabaseToolsRuntimeClient) ExecuteSqlDatabaseToolsConnection(ctx context.Context, request ExecuteSqlDatabaseToolsConnectionRequest) (response ExecuteSqlDatabaseToolsConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.executeSqlDatabaseToolsConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExecuteSqlDatabaseToolsConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExecuteSqlDatabaseToolsConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExecuteSqlDatabaseToolsConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExecuteSqlDatabaseToolsConnectionResponse")
	}
	return
}

// executeSqlDatabaseToolsConnection implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) executeSqlDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/actions/executeSql", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExecuteSqlDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ExecuteSqlDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ExecuteSqlDatabaseToolsConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &executesqlresponse{})
	return response, err
}

// GetCredential Get a credential
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetCredential.go.html to see an example of how to use GetCredential API.
// A default retry strategy applies to this operation GetCredential()
func (client DatabaseToolsRuntimeClient) GetCredential(ctx context.Context, request GetCredentialRequest) (response GetCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCredentialResponse")
	}
	return
}

// getCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCredentialExecuteGrantee Get a credential execute grantee
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetCredentialExecuteGrantee.go.html to see an example of how to use GetCredentialExecuteGrantee API.
// A default retry strategy applies to this operation GetCredentialExecuteGrantee()
func (client DatabaseToolsRuntimeClient) GetCredentialExecuteGrantee(ctx context.Context, request GetCredentialExecuteGranteeRequest) (response GetCredentialExecuteGranteeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCredentialExecuteGrantee, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCredentialExecuteGranteeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCredentialExecuteGranteeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCredentialExecuteGranteeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCredentialExecuteGranteeResponse")
	}
	return
}

// getCredentialExecuteGrantee implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getCredentialExecuteGrantee(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/executeGrantees/{executeGranteeKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCredentialExecuteGranteeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetCredentialExecuteGrantee")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetCredentialExecuteGrantee", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCredentialPublicSynonym Get a public synonym
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetCredentialPublicSynonym.go.html to see an example of how to use GetCredentialPublicSynonym API.
// A default retry strategy applies to this operation GetCredentialPublicSynonym()
func (client DatabaseToolsRuntimeClient) GetCredentialPublicSynonym(ctx context.Context, request GetCredentialPublicSynonymRequest) (response GetCredentialPublicSynonymResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCredentialPublicSynonym, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCredentialPublicSynonymResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCredentialPublicSynonymResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCredentialPublicSynonymResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCredentialPublicSynonymResponse")
	}
	return
}

// getCredentialPublicSynonym implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getCredentialPublicSynonym(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/publicSynonyms/{publicSynonymKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCredentialPublicSynonymResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetCredentialPublicSynonym")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetCredentialPublicSynonym", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseToolsDatabaseApiGatewayConfigContent Get the content of a Database Tools database API gateway config
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetDatabaseToolsDatabaseApiGatewayConfigContent.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfigContent API.
// A default retry strategy applies to this operation GetDatabaseToolsDatabaseApiGatewayConfigContent()
func (client DatabaseToolsRuntimeClient) GetDatabaseToolsDatabaseApiGatewayConfigContent(ctx context.Context, request GetDatabaseToolsDatabaseApiGatewayConfigContentRequest) (response GetDatabaseToolsDatabaseApiGatewayConfigContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsDatabaseApiGatewayConfigContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsDatabaseApiGatewayConfigContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsDatabaseApiGatewayConfigContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsDatabaseApiGatewayConfigContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsDatabaseApiGatewayConfigContentResponse")
	}
	return
}

// getDatabaseToolsDatabaseApiGatewayConfigContent implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getDatabaseToolsDatabaseApiGatewayConfigContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsDatabaseApiGatewayConfigContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigContent")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseToolsDatabaseApiGatewayConfigGlobal Get a Database Tools database API gateway config global resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetDatabaseToolsDatabaseApiGatewayConfigGlobal.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfigGlobal API.
// A default retry strategy applies to this operation GetDatabaseToolsDatabaseApiGatewayConfigGlobal()
func (client DatabaseToolsRuntimeClient) GetDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx context.Context, request GetDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) (response GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsDatabaseApiGatewayConfigGlobal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse")
	}
	return
}

// getDatabaseToolsDatabaseApiGatewayConfigGlobal implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/globals/{globalKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsDatabaseApiGatewayConfigGlobalResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigGlobal")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigGlobal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigglobal{})
	return response, err
}

// GetDatabaseToolsDatabaseApiGatewayConfigPool Get a Database Tools database API gateway config pool resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetDatabaseToolsDatabaseApiGatewayConfigPool.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfigPool API.
// A default retry strategy applies to this operation GetDatabaseToolsDatabaseApiGatewayConfigPool()
func (client DatabaseToolsRuntimeClient) GetDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request GetDatabaseToolsDatabaseApiGatewayConfigPoolRequest) (response GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsDatabaseApiGatewayConfigPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse")
	}
	return
}

// getDatabaseToolsDatabaseApiGatewayConfigPool implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpool{})
	return response, err
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec Get a Database Tools database API gateway config API spec resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec API.
// A default retry strategy applies to this operation GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec()
func (client DatabaseToolsRuntimeClient) GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest) (response GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse")
	}
	return
}

// getDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/apiSpecs/{apiSpecKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpoolapispec{})
	return response, err
}

// GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec Get a Database Tools database API gateway config auto API spec resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec API.
// A default retry strategy applies to this operation GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec()
func (client DatabaseToolsRuntimeClient) GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) (response GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse")
	}
	return
}

// getDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/autoApiSpecs/{autoApiSpecKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpoolautoapispec{})
	return response, err
}

// GetPropertySet Get a property set
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetPropertySet.go.html to see an example of how to use GetPropertySet API.
// A default retry strategy applies to this operation GetPropertySet()
func (client DatabaseToolsRuntimeClient) GetPropertySet(ctx context.Context, request GetPropertySetRequest) (response GetPropertySetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPropertySet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPropertySetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPropertySetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPropertySetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPropertySetResponse")
	}
	return
}

// getPropertySet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getPropertySet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/propertySets/{propertySetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPropertySetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetPropertySet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetPropertySet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &propertyset{})
	return response, err
}

// GetUserCredential Get a user credential
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetUserCredential.go.html to see an example of how to use GetUserCredential API.
// A default retry strategy applies to this operation GetUserCredential()
func (client DatabaseToolsRuntimeClient) GetUserCredential(ctx context.Context, request GetUserCredentialRequest) (response GetUserCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserCredentialResponse")
	}
	return
}

// getUserCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) getUserCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/users/{userKey}/credentials/{credentialKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetUserCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetUserCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DatabaseToolsRuntimeClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DatabaseToolsRuntimeClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "GetWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCredentialExecuteGrantees Get a list of all execute grantees
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListCredentialExecuteGrantees.go.html to see an example of how to use ListCredentialExecuteGrantees API.
// A default retry strategy applies to this operation ListCredentialExecuteGrantees()
func (client DatabaseToolsRuntimeClient) ListCredentialExecuteGrantees(ctx context.Context, request ListCredentialExecuteGranteesRequest) (response ListCredentialExecuteGranteesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCredentialExecuteGrantees, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCredentialExecuteGranteesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCredentialExecuteGranteesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCredentialExecuteGranteesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCredentialExecuteGranteesResponse")
	}
	return
}

// listCredentialExecuteGrantees implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listCredentialExecuteGrantees(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/executeGrantees", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCredentialExecuteGranteesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListCredentialExecuteGrantees")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListCredentialExecuteGrantees", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCredentialPublicSynonyms Get a list of all public synonyms for the given credential
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListCredentialPublicSynonyms.go.html to see an example of how to use ListCredentialPublicSynonyms API.
// A default retry strategy applies to this operation ListCredentialPublicSynonyms()
func (client DatabaseToolsRuntimeClient) ListCredentialPublicSynonyms(ctx context.Context, request ListCredentialPublicSynonymsRequest) (response ListCredentialPublicSynonymsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCredentialPublicSynonyms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCredentialPublicSynonymsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCredentialPublicSynonymsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCredentialPublicSynonymsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCredentialPublicSynonymsResponse")
	}
	return
}

// listCredentialPublicSynonyms implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listCredentialPublicSynonyms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}/publicSynonyms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCredentialPublicSynonymsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListCredentialPublicSynonyms")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListCredentialPublicSynonyms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCredentials Returns a paginated list of `CredentialSummary` for the specified request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListCredentials.go.html to see an example of how to use ListCredentials API.
// A default retry strategy applies to this operation ListCredentials()
func (client DatabaseToolsRuntimeClient) ListCredentials(ctx context.Context, request ListCredentialsRequest) (response ListCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCredentialsResponse")
	}
	return
}

// listCredentials implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListCredentials")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties Returns list of database API gateway config setting descriptions to be provided as advanced properties.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties API.
// A default retry strategy applies to this operation ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties()
func (client DatabaseToolsRuntimeClient) ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties(ctx context.Context, request ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesRequest) (response ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse")
	}
	return
}

// listDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigAdvancedProperties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs Returns a list of Database Tools database API gateway config API spec resources
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs API.
// A default retry strategy applies to this operation ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs()
func (client DatabaseToolsRuntimeClient) ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs(ctx context.Context, request ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest) (response ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse")
	}
	return
}

// listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/apiSpecs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs Returns a list of Database Tools database API gateway config auto API spec resources
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs API.
// A default retry strategy applies to this operation ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs()
func (client DatabaseToolsRuntimeClient) ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs(ctx context.Context, request ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest) (response ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse")
	}
	return
}

// listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/autoApiSpecs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseToolsDatabaseApiGatewayConfigPools Returns a list of Database Tools database API gateway config pool resources
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListDatabaseToolsDatabaseApiGatewayConfigPools.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigPools API.
// A default retry strategy applies to this operation ListDatabaseToolsDatabaseApiGatewayConfigPools()
func (client DatabaseToolsRuntimeClient) ListDatabaseToolsDatabaseApiGatewayConfigPools(ctx context.Context, request ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest) (response ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseToolsDatabaseApiGatewayConfigPools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse")
	}
	return
}

// listDatabaseToolsDatabaseApiGatewayConfigPools implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listDatabaseToolsDatabaseApiGatewayConfigPools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseToolsDatabaseApiGatewayConfigPoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigPools")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListDatabaseToolsDatabaseApiGatewayConfigPools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserCredentials Returns a paginated list of user `UserCredentialSummary` for the specified request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListUserCredentials.go.html to see an example of how to use ListUserCredentials API.
// A default retry strategy applies to this operation ListUserCredentials()
func (client DatabaseToolsRuntimeClient) ListUserCredentials(ctx context.Context, request ListUserCredentialsRequest) (response ListUserCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserCredentialsResponse")
	}
	return
}

// listUserCredentials implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) listUserCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseToolsConnections/{databaseToolsConnectionId}/users/{userKey}/credentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListUserCredentials")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListUserCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a paginated list of errors for the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DatabaseToolsRuntimeClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DatabaseToolsRuntimeClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DatabaseToolsRuntimeClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DatabaseToolsRuntimeClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DatabaseToolsRuntimeClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DatabaseToolsRuntimeClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ListWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCredential Update a credential
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateCredential.go.html to see an example of how to use UpdateCredential API.
// A default retry strategy applies to this operation UpdateCredential()
func (client DatabaseToolsRuntimeClient) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) (response UpdateCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCredentialResponse")
	}
	return
}

// updateCredential implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) updateCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "UpdateCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "UpdateCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal Update a Database Tools database API gateway config global resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal API.
// A default retry strategy applies to this operation UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal()
func (client DatabaseToolsRuntimeClient) UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx context.Context, request UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) (response UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsDatabaseApiGatewayConfigGlobal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse")
	}
	return
}

// updateDatabaseToolsDatabaseApiGatewayConfigGlobal implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) updateDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/globals/{globalKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigglobal{})
	return response, err
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPool Update a Database Tools database API gateway config pool resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateDatabaseToolsDatabaseApiGatewayConfigPool.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfigPool API.
// A default retry strategy applies to this operation UpdateDatabaseToolsDatabaseApiGatewayConfigPool()
func (client DatabaseToolsRuntimeClient) UpdateDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) (response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsDatabaseApiGatewayConfigPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsDatabaseApiGatewayConfigPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsDatabaseApiGatewayConfigPoolResponse")
	}
	return
}

// updateDatabaseToolsDatabaseApiGatewayConfigPool implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) updateDatabaseToolsDatabaseApiGatewayConfigPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpool{})
	return response, err
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec Update a Database Tools database API gateway config API spec resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec API.
// A default retry strategy applies to this operation UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec()
func (client DatabaseToolsRuntimeClient) UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest) (response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse")
	}
	return
}

// updateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) updateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/apiSpecs/{apiSpecKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpoolapispec{})
	return response, err
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec Update a Database Tools database API gateway config auto API spec resource
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec API.
// A default retry strategy applies to this operation UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec()
func (client DatabaseToolsRuntimeClient) UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) (response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse")
	}
	return
}

// updateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) updateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/autoApiSpecs/{autoApiSpecKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databasetoolsdatabaseapigatewayconfigpoolautoapispec{})
	return response, err
}

// UpdatePropertySet Update a property set
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdatePropertySet.go.html to see an example of how to use UpdatePropertySet API.
// A default retry strategy applies to this operation UpdatePropertySet()
func (client DatabaseToolsRuntimeClient) UpdatePropertySet(ctx context.Context, request UpdatePropertySetRequest) (response UpdatePropertySetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePropertySet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePropertySetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePropertySetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePropertySetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePropertySetResponse")
	}
	return
}

// updatePropertySet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseToolsRuntimeClient) updatePropertySet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseToolsConnections/{databaseToolsConnectionId}/propertySets/{propertySetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePropertySetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "UpdatePropertySet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "UpdatePropertySet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &propertyset{})
	return response, err
}

// ValidateDatabaseToolsConnection Validates the specified Database Tools connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ValidateDatabaseToolsConnection.go.html to see an example of how to use ValidateDatabaseToolsConnection API.
// A default retry strategy applies to this operation ValidateDatabaseToolsConnection()
func (client DatabaseToolsRuntimeClient) ValidateDatabaseToolsConnection(ctx context.Context, request ValidateDatabaseToolsConnectionRequest) (response ValidateDatabaseToolsConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseToolsRuntimeClient) validateDatabaseToolsConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsConnections/{databaseToolsConnectionId}/actions/validateConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDatabaseToolsConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ValidateDatabaseToolsConnection")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ValidateDatabaseToolsConnection", apiReferenceLink)
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
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ValidateDatabaseToolsIdentityCredential.go.html to see an example of how to use ValidateDatabaseToolsIdentityCredential API.
// A default retry strategy applies to this operation ValidateDatabaseToolsIdentityCredential()
func (client DatabaseToolsRuntimeClient) ValidateDatabaseToolsIdentityCredential(ctx context.Context, request ValidateDatabaseToolsIdentityCredentialRequest) (response ValidateDatabaseToolsIdentityCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseToolsRuntimeClient) validateDatabaseToolsIdentityCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseToolsIdentities/{databaseToolsIdentityId}/actions/validateCredential", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDatabaseToolsIdentityCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "databaseToolsRuntime", "ValidateDatabaseToolsIdentityCredential")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseToolsRuntime", "ValidateDatabaseToolsIdentityCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &validatedatabasetoolsidentitycredentialresult{})
	return response, err
}
