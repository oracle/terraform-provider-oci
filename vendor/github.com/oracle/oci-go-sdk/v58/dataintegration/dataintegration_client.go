// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//DataIntegrationClient a client for DataIntegration
type DataIntegrationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataIntegrationClientWithConfigurationProvider Creates a new default DataIntegration client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataIntegrationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataIntegrationClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDataIntegrationClientFromBaseClient(baseClient, provider)
}

// NewDataIntegrationClientWithOboToken Creates a new default DataIntegration client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataIntegrationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataIntegrationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataIntegrationClientFromBaseClient(baseClient, configProvider)
}

func newDataIntegrationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataIntegrationClient, err error) {
	// DataIntegration service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataIntegrationClient{BaseClient: baseClient}
	client.BasePath = "20200430"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataIntegrationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dataintegration", "https://dataintegration.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataIntegrationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataIntegrationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeCompartment Moves a workspace to a specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ChangeCompartment.go.html to see an example of how to use ChangeCompartment API.
func (client DataIntegrationClient) ChangeCompartment(ctx context.Context, request ChangeCompartmentRequest) (response ChangeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCompartmentResponse")
	}
	return
}

// changeCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) changeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateApplication Creates an application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateApplication.go.html to see an example of how to use CreateApplication API.
func (client DataIntegrationClient) CreateApplication(ctx context.Context, request CreateApplicationRequest) (response CreateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApplicationResponse")
	}
	return
}

// createApplication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConnection Creates a connection under an existing data asset.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateConnection.go.html to see an example of how to use CreateConnection API.
func (client DataIntegrationClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionResponse")
	}
	return
}

// createConnection implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// CreateConnectionValidation Creates a connection validation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateConnectionValidation.go.html to see an example of how to use CreateConnectionValidation API.
func (client DataIntegrationClient) CreateConnectionValidation(ctx context.Context, request CreateConnectionValidationRequest) (response CreateConnectionValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createConnectionValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionValidationResponse")
	}
	return
}

// createConnectionValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createConnectionValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/connectionValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataAsset Creates a data asset with default connection.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateDataAsset.go.html to see an example of how to use CreateDataAsset API.
func (client DataIntegrationClient) CreateDataAsset(ctx context.Context, request CreateDataAssetRequest) (response CreateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetResponse")
	}
	return
}

// createDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/dataAssets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dataasset{})
	return response, err
}

// CreateDataFlow Creates a new data flow in a project or folder ready for performing data integrations.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateDataFlow.go.html to see an example of how to use CreateDataFlow API.
func (client DataIntegrationClient) CreateDataFlow(ctx context.Context, request CreateDataFlowRequest) (response CreateDataFlowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataFlow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataFlowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataFlowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataFlowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataFlowResponse")
	}
	return
}

// createDataFlow implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createDataFlow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/dataFlows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataFlowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataFlowValidation Accepts the data flow definition in the request payload and creates a data flow validation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateDataFlowValidation.go.html to see an example of how to use CreateDataFlowValidation API.
func (client DataIntegrationClient) CreateDataFlowValidation(ctx context.Context, request CreateDataFlowValidationRequest) (response CreateDataFlowValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataFlowValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataFlowValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataFlowValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataFlowValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataFlowValidationResponse")
	}
	return
}

// createDataFlowValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createDataFlowValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/dataFlowValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataFlowValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEntityShape Creates the data entity shape using the shape from the data asset.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateEntityShape.go.html to see an example of how to use CreateEntityShape API.
func (client DataIntegrationClient) CreateEntityShape(ctx context.Context, request CreateEntityShapeRequest) (response CreateEntityShapeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntityShape, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEntityShapeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEntityShapeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityShapeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityShapeResponse")
	}
	return
}

// createEntityShape implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createEntityShape(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/connections/{connectionKey}/schemas/{schemaResourceName}/entityShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEntityShapeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &entityshape{})
	return response, err
}

// CreateExternalPublication Publish a DataFlow in a OCI DataFlow application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateExternalPublication.go.html to see an example of how to use CreateExternalPublication API.
func (client DataIntegrationClient) CreateExternalPublication(ctx context.Context, request CreateExternalPublicationRequest) (response CreateExternalPublicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createExternalPublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalPublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalPublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalPublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalPublicationResponse")
	}
	return
}

// createExternalPublication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createExternalPublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalPublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExternalPublicationValidation Validates a specific task.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateExternalPublicationValidation.go.html to see an example of how to use CreateExternalPublicationValidation API.
func (client DataIntegrationClient) CreateExternalPublicationValidation(ctx context.Context, request CreateExternalPublicationValidationRequest) (response CreateExternalPublicationValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createExternalPublicationValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExternalPublicationValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExternalPublicationValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExternalPublicationValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExternalPublicationValidationResponse")
	}
	return
}

// createExternalPublicationValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createExternalPublicationValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublicationValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExternalPublicationValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFolder Creates a folder in a project or in another folder, limited to two levels of folders. |
// Folders are used to organize your design-time resources, such as tasks or data flows.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateFolder.go.html to see an example of how to use CreateFolder API.
func (client DataIntegrationClient) CreateFolder(ctx context.Context, request CreateFolderRequest) (response CreateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderResponse")
	}
	return
}

// createFolder implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/folders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFunctionLibrary Creates a function library in a project or in another function library, limited to two levels of function libraries. |
// FunctionLibraries are used to organize your design-time resources, such as tasks or data flows.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateFunctionLibrary.go.html to see an example of how to use CreateFunctionLibrary API.
func (client DataIntegrationClient) CreateFunctionLibrary(ctx context.Context, request CreateFunctionLibraryRequest) (response CreateFunctionLibraryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFunctionLibrary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFunctionLibraryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFunctionLibraryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFunctionLibraryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFunctionLibraryResponse")
	}
	return
}

// createFunctionLibrary implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createFunctionLibrary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/functionLibraries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFunctionLibraryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePatch Creates a patch in an application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreatePatch.go.html to see an example of how to use CreatePatch API.
func (client DataIntegrationClient) CreatePatch(ctx context.Context, request CreatePatchRequest) (response CreatePatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePatchResponse")
	}
	return
}

// createPatch implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/applications/{applicationKey}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePipeline Creates a new pipeline in a project or folder ready for performing task orchestration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreatePipeline.go.html to see an example of how to use CreatePipeline API.
func (client DataIntegrationClient) CreatePipeline(ctx context.Context, request CreatePipelineRequest) (response CreatePipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePipelineResponse")
	}
	return
}

// createPipeline implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/pipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePipelineValidation Accepts the data flow definition in the request payload and creates a pipeline validation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreatePipelineValidation.go.html to see an example of how to use CreatePipelineValidation API.
func (client DataIntegrationClient) CreatePipelineValidation(ctx context.Context, request CreatePipelineValidationRequest) (response CreatePipelineValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createPipelineValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePipelineValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePipelineValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePipelineValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePipelineValidationResponse")
	}
	return
}

// createPipelineValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createPipelineValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/pipelineValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePipelineValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProject Creates a project. Projects are organizational constructs within a workspace that you use to organize your design-time resources, such as tasks or data flows. Projects can be organized into folders.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateProject.go.html to see an example of how to use CreateProject API.
func (client DataIntegrationClient) CreateProject(ctx context.Context, request CreateProjectRequest) (response CreateProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProjectResponse")
	}
	return
}

// createProject implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/projects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSchedule Endpoint to create a new schedule
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateSchedule.go.html to see an example of how to use CreateSchedule API.
func (client DataIntegrationClient) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (response CreateScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduleResponse")
	}
	return
}

// createSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/applications/{applicationKey}/schedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTask Creates a new task ready for performing data integrations. There are specialized types of tasks that include data loader and integration tasks.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateTask.go.html to see an example of how to use CreateTask API.
func (client DataIntegrationClient) CreateTask(ctx context.Context, request CreateTaskRequest) (response CreateTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTaskResponse")
	}
	return
}

// createTask implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/tasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &task{})
	return response, err
}

// CreateTaskRun Creates a data integration task run for the specified task.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateTaskRun.go.html to see an example of how to use CreateTaskRun API.
func (client DataIntegrationClient) CreateTaskRun(ctx context.Context, request CreateTaskRunRequest) (response CreateTaskRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTaskRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTaskRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTaskRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTaskRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTaskRunResponse")
	}
	return
}

// createTaskRun implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createTaskRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/applications/{applicationKey}/taskRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTaskRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTaskSchedule Endpoint to be used create TaskSchedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateTaskSchedule.go.html to see an example of how to use CreateTaskSchedule API.
func (client DataIntegrationClient) CreateTaskSchedule(ctx context.Context, request CreateTaskScheduleRequest) (response CreateTaskScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTaskSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTaskScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTaskScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTaskScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTaskScheduleResponse")
	}
	return
}

// createTaskSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createTaskSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/applications/{applicationKey}/taskSchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTaskScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTaskValidation Validates a specific task.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateTaskValidation.go.html to see an example of how to use CreateTaskValidation API.
func (client DataIntegrationClient) CreateTaskValidation(ctx context.Context, request CreateTaskValidationRequest) (response CreateTaskValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTaskValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTaskValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTaskValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTaskValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTaskValidationResponse")
	}
	return
}

// createTaskValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createTaskValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/taskValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTaskValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUserDefinedFunction Creates a new UserDefinedFunction in a function library ready for performing data integrations.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateUserDefinedFunction.go.html to see an example of how to use CreateUserDefinedFunction API.
func (client DataIntegrationClient) CreateUserDefinedFunction(ctx context.Context, request CreateUserDefinedFunctionRequest) (response CreateUserDefinedFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createUserDefinedFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserDefinedFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserDefinedFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserDefinedFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserDefinedFunctionResponse")
	}
	return
}

// createUserDefinedFunction implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createUserDefinedFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/userDefinedFunctions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserDefinedFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUserDefinedFunctionValidation Accepts the UserDefinedFunction definition in the request payload and creates a UserDefinedFunction validation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateUserDefinedFunctionValidation.go.html to see an example of how to use CreateUserDefinedFunctionValidation API.
func (client DataIntegrationClient) CreateUserDefinedFunctionValidation(ctx context.Context, request CreateUserDefinedFunctionValidationRequest) (response CreateUserDefinedFunctionValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createUserDefinedFunctionValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserDefinedFunctionValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserDefinedFunctionValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserDefinedFunctionValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserDefinedFunctionValidationResponse")
	}
	return
}

// createUserDefinedFunctionValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createUserDefinedFunctionValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/userDefinedFunctionValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserDefinedFunctionValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateWorkspace Creates a new Data Integration workspace ready for performing data integration tasks. To retrieve the OCID for the new workspace, use the opc-work-request-id returned by this API and call the GetWorkRequest API.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/CreateWorkspace.go.html to see an example of how to use CreateWorkspace API.
func (client DataIntegrationClient) CreateWorkspace(ctx context.Context, request CreateWorkspaceRequest) (response CreateWorkspaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createWorkspace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateWorkspaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateWorkspaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateWorkspaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateWorkspaceResponse")
	}
	return
}

// createWorkspace implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) createWorkspace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateWorkspaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApplication Removes an application using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteApplication.go.html to see an example of how to use DeleteApplication API.
func (client DataIntegrationClient) DeleteApplication(ctx context.Context, request DeleteApplicationRequest) (response DeleteApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplicationResponse")
	}
	return
}

// deleteApplication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/applications/{applicationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnection Removes a connection using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteConnection.go.html to see an example of how to use DeleteConnection API.
func (client DataIntegrationClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionResponse")
	}
	return
}

// deleteConnection implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/connections/{connectionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnectionValidation Deletes a connection validation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteConnectionValidation.go.html to see an example of how to use DeleteConnectionValidation API.
func (client DataIntegrationClient) DeleteConnectionValidation(ctx context.Context, request DeleteConnectionValidationRequest) (response DeleteConnectionValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnectionValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionValidationResponse")
	}
	return
}

// deleteConnectionValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteConnectionValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/connectionValidations/{connectionValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataAsset Removes a data asset using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteDataAsset.go.html to see an example of how to use DeleteDataAsset API.
func (client DataIntegrationClient) DeleteDataAsset(ctx context.Context, request DeleteDataAssetRequest) (response DeleteDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetResponse")
	}
	return
}

// deleteDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/dataAssets/{dataAssetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataFlow Removes a data flow from a project or folder using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteDataFlow.go.html to see an example of how to use DeleteDataFlow API.
func (client DataIntegrationClient) DeleteDataFlow(ctx context.Context, request DeleteDataFlowRequest) (response DeleteDataFlowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataFlow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataFlowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataFlowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataFlowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataFlowResponse")
	}
	return
}

// deleteDataFlow implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteDataFlow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/dataFlows/{dataFlowKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataFlowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataFlowValidation Removes a data flow validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteDataFlowValidation.go.html to see an example of how to use DeleteDataFlowValidation API.
func (client DataIntegrationClient) DeleteDataFlowValidation(ctx context.Context, request DeleteDataFlowValidationRequest) (response DeleteDataFlowValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataFlowValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataFlowValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataFlowValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataFlowValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataFlowValidationResponse")
	}
	return
}

// deleteDataFlowValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteDataFlowValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/dataFlowValidations/{dataFlowValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataFlowValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalPublication Removes a published object using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteExternalPublication.go.html to see an example of how to use DeleteExternalPublication API.
func (client DataIntegrationClient) DeleteExternalPublication(ctx context.Context, request DeleteExternalPublicationRequest) (response DeleteExternalPublicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalPublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalPublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalPublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalPublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalPublicationResponse")
	}
	return
}

// deleteExternalPublication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteExternalPublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublications/{externalPublicationsKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalPublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExternalPublicationValidation Removes a task validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteExternalPublicationValidation.go.html to see an example of how to use DeleteExternalPublicationValidation API.
func (client DataIntegrationClient) DeleteExternalPublicationValidation(ctx context.Context, request DeleteExternalPublicationValidationRequest) (response DeleteExternalPublicationValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExternalPublicationValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExternalPublicationValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExternalPublicationValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExternalPublicationValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExternalPublicationValidationResponse")
	}
	return
}

// deleteExternalPublicationValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteExternalPublicationValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublicationValidations/{externalPublicationValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExternalPublicationValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFolder Removes a folder from a project using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteFolder.go.html to see an example of how to use DeleteFolder API.
func (client DataIntegrationClient) DeleteFolder(ctx context.Context, request DeleteFolderRequest) (response DeleteFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderResponse")
	}
	return
}

// deleteFolder implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/folders/{folderKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFunctionLibrary Removes a Function Library from a project using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteFunctionLibrary.go.html to see an example of how to use DeleteFunctionLibrary API.
func (client DataIntegrationClient) DeleteFunctionLibrary(ctx context.Context, request DeleteFunctionLibraryRequest) (response DeleteFunctionLibraryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFunctionLibrary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFunctionLibraryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFunctionLibraryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFunctionLibraryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFunctionLibraryResponse")
	}
	return
}

// deleteFunctionLibrary implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteFunctionLibrary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/functionLibraries/{functionLibraryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFunctionLibraryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePatch Removes a patch using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeletePatch.go.html to see an example of how to use DeletePatch API.
func (client DataIntegrationClient) DeletePatch(ctx context.Context, request DeletePatchRequest) (response DeletePatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePatchResponse")
	}
	return
}

// deletePatch implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deletePatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/applications/{applicationKey}/patches/{patchKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePipeline Removes a pipeline from a project or folder using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeletePipeline.go.html to see an example of how to use DeletePipeline API.
func (client DataIntegrationClient) DeletePipeline(ctx context.Context, request DeletePipelineRequest) (response DeletePipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePipelineResponse")
	}
	return
}

// deletePipeline implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deletePipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/pipelines/{pipelineKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePipelineValidation Removes a pipeline validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeletePipelineValidation.go.html to see an example of how to use DeletePipelineValidation API.
func (client DataIntegrationClient) DeletePipelineValidation(ctx context.Context, request DeletePipelineValidationRequest) (response DeletePipelineValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePipelineValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePipelineValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePipelineValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePipelineValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePipelineValidationResponse")
	}
	return
}

// deletePipelineValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deletePipelineValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/pipelineValidations/{pipelineValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePipelineValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProject Removes a project from the workspace using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteProject.go.html to see an example of how to use DeleteProject API.
func (client DataIntegrationClient) DeleteProject(ctx context.Context, request DeleteProjectRequest) (response DeleteProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProjectResponse")
	}
	return
}

// deleteProject implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/projects/{projectKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSchedule Endpoint to delete schedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteSchedule.go.html to see an example of how to use DeleteSchedule API.
func (client DataIntegrationClient) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) (response DeleteScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduleResponse")
	}
	return
}

// deleteSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/applications/{applicationKey}/schedules/{scheduleKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTask Removes a task using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteTask.go.html to see an example of how to use DeleteTask API.
func (client DataIntegrationClient) DeleteTask(ctx context.Context, request DeleteTaskRequest) (response DeleteTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTaskResponse")
	}
	return
}

// deleteTask implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/tasks/{taskKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTaskRun Deletes a task run using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteTaskRun.go.html to see an example of how to use DeleteTaskRun API.
func (client DataIntegrationClient) DeleteTaskRun(ctx context.Context, request DeleteTaskRunRequest) (response DeleteTaskRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTaskRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTaskRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTaskRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTaskRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTaskRunResponse")
	}
	return
}

// deleteTaskRun implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteTaskRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/applications/{applicationKey}/taskRuns/{taskRunKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTaskRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTaskSchedule Endpoint to delete TaskSchedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteTaskSchedule.go.html to see an example of how to use DeleteTaskSchedule API.
func (client DataIntegrationClient) DeleteTaskSchedule(ctx context.Context, request DeleteTaskScheduleRequest) (response DeleteTaskScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTaskSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTaskScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTaskScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTaskScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTaskScheduleResponse")
	}
	return
}

// deleteTaskSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteTaskSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/applications/{applicationKey}/taskSchedules/{taskScheduleKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTaskScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTaskValidation Removes a task validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteTaskValidation.go.html to see an example of how to use DeleteTaskValidation API.
func (client DataIntegrationClient) DeleteTaskValidation(ctx context.Context, request DeleteTaskValidationRequest) (response DeleteTaskValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTaskValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTaskValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTaskValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTaskValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTaskValidationResponse")
	}
	return
}

// deleteTaskValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteTaskValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/taskValidations/{taskValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTaskValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUserDefinedFunction Removes a UserDefinedFunction from a function library using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteUserDefinedFunction.go.html to see an example of how to use DeleteUserDefinedFunction API.
func (client DataIntegrationClient) DeleteUserDefinedFunction(ctx context.Context, request DeleteUserDefinedFunctionRequest) (response DeleteUserDefinedFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUserDefinedFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserDefinedFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserDefinedFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserDefinedFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserDefinedFunctionResponse")
	}
	return
}

// deleteUserDefinedFunction implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteUserDefinedFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/userDefinedFunctions/{userDefinedFunctionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserDefinedFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUserDefinedFunctionValidation Removes a UserDefinedFunction validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteUserDefinedFunctionValidation.go.html to see an example of how to use DeleteUserDefinedFunctionValidation API.
func (client DataIntegrationClient) DeleteUserDefinedFunctionValidation(ctx context.Context, request DeleteUserDefinedFunctionValidationRequest) (response DeleteUserDefinedFunctionValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUserDefinedFunctionValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserDefinedFunctionValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserDefinedFunctionValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserDefinedFunctionValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserDefinedFunctionValidationResponse")
	}
	return
}

// deleteUserDefinedFunctionValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteUserDefinedFunctionValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}/userDefinedFunctionValidations/{userDefinedFunctionValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserDefinedFunctionValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkspace Deletes a Data Integration workspace resource using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteWorkspace.go.html to see an example of how to use DeleteWorkspace API.
func (client DataIntegrationClient) DeleteWorkspace(ctx context.Context, request DeleteWorkspaceRequest) (response DeleteWorkspaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkspace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkspaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkspaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkspaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkspaceResponse")
	}
	return
}

// deleteWorkspace implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) deleteWorkspace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workspaces/{workspaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWorkspaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApplication Retrieves an application using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetApplication.go.html to see an example of how to use GetApplication API.
func (client DataIntegrationClient) GetApplication(ctx context.Context, request GetApplicationRequest) (response GetApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplicationResponse")
	}
	return
}

// getApplication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConnection Retrieves the connection details using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetConnection.go.html to see an example of how to use GetConnection API.
func (client DataIntegrationClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionResponse")
	}
	return
}

// getConnection implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connections/{connectionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// GetConnectionValidation Retrieves a connection validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetConnectionValidation.go.html to see an example of how to use GetConnectionValidation API.
func (client DataIntegrationClient) GetConnectionValidation(ctx context.Context, request GetConnectionValidationRequest) (response GetConnectionValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnectionValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionValidationResponse")
	}
	return
}

// getConnectionValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getConnectionValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connectionValidations/{connectionValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCountStatistic Retrieves statistics on a workspace. It returns an object with an array of property values, such as the number of projects, |
//        applications, data assets, and so on.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetCountStatistic.go.html to see an example of how to use GetCountStatistic API.
func (client DataIntegrationClient) GetCountStatistic(ctx context.Context, request GetCountStatisticRequest) (response GetCountStatisticResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCountStatistic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCountStatisticResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCountStatisticResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCountStatisticResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCountStatisticResponse")
	}
	return
}

// getCountStatistic implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getCountStatistic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/countStatistics/{countStatisticKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCountStatisticResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataAsset Retrieves details of a data asset using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetDataAsset.go.html to see an example of how to use GetDataAsset API.
func (client DataIntegrationClient) GetDataAsset(ctx context.Context, request GetDataAssetRequest) (response GetDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetResponse")
	}
	return
}

// getDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/dataAssets/{dataAssetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dataasset{})
	return response, err
}

// GetDataEntity Retrieves the data entity details with the given name from live schema.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetDataEntity.go.html to see an example of how to use GetDataEntity API.
func (client DataIntegrationClient) GetDataEntity(ctx context.Context, request GetDataEntityRequest) (response GetDataEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataEntityResponse")
	}
	return
}

// getDataEntity implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getDataEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connections/{connectionKey}/schemas/{schemaResourceName}/dataEntities/{dataEntityKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dataentity{})
	return response, err
}

// GetDataFlow Retrieves a data flow using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetDataFlow.go.html to see an example of how to use GetDataFlow API.
func (client DataIntegrationClient) GetDataFlow(ctx context.Context, request GetDataFlowRequest) (response GetDataFlowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataFlow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataFlowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataFlowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataFlowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataFlowResponse")
	}
	return
}

// getDataFlow implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getDataFlow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/dataFlows/{dataFlowKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataFlowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataFlowValidation Retrieves a data flow validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetDataFlowValidation.go.html to see an example of how to use GetDataFlowValidation API.
func (client DataIntegrationClient) GetDataFlowValidation(ctx context.Context, request GetDataFlowValidationRequest) (response GetDataFlowValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataFlowValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataFlowValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataFlowValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataFlowValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataFlowValidationResponse")
	}
	return
}

// getDataFlowValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getDataFlowValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/dataFlowValidations/{dataFlowValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataFlowValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDependentObject Retrieves the details of a dependent object from an application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetDependentObject.go.html to see an example of how to use GetDependentObject API.
func (client DataIntegrationClient) GetDependentObject(ctx context.Context, request GetDependentObjectRequest) (response GetDependentObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDependentObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDependentObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDependentObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDependentObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDependentObjectResponse")
	}
	return
}

// getDependentObject implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getDependentObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/dependentObjects/{dependentObjectKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDependentObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalPublication Retrieves a publshed object in an task using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetExternalPublication.go.html to see an example of how to use GetExternalPublication API.
func (client DataIntegrationClient) GetExternalPublication(ctx context.Context, request GetExternalPublicationRequest) (response GetExternalPublicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalPublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalPublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalPublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalPublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalPublicationResponse")
	}
	return
}

// getExternalPublication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getExternalPublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublications/{externalPublicationsKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalPublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExternalPublicationValidation Retrieves an external publication validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetExternalPublicationValidation.go.html to see an example of how to use GetExternalPublicationValidation API.
func (client DataIntegrationClient) GetExternalPublicationValidation(ctx context.Context, request GetExternalPublicationValidationRequest) (response GetExternalPublicationValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExternalPublicationValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExternalPublicationValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExternalPublicationValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExternalPublicationValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExternalPublicationValidationResponse")
	}
	return
}

// getExternalPublicationValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getExternalPublicationValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublicationValidations/{externalPublicationValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExternalPublicationValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFolder Retrieves a folder using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetFolder.go.html to see an example of how to use GetFolder API.
func (client DataIntegrationClient) GetFolder(ctx context.Context, request GetFolderRequest) (response GetFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderResponse")
	}
	return
}

// getFolder implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/folders/{folderKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFunctionLibrary Retrieves a Function Library using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetFunctionLibrary.go.html to see an example of how to use GetFunctionLibrary API.
func (client DataIntegrationClient) GetFunctionLibrary(ctx context.Context, request GetFunctionLibraryRequest) (response GetFunctionLibraryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFunctionLibrary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFunctionLibraryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFunctionLibraryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFunctionLibraryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFunctionLibraryResponse")
	}
	return
}

// getFunctionLibrary implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getFunctionLibrary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/functionLibraries/{functionLibraryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFunctionLibraryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPatch Retrieves a patch in an application using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetPatch.go.html to see an example of how to use GetPatch API.
func (client DataIntegrationClient) GetPatch(ctx context.Context, request GetPatchRequest) (response GetPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatchResponse")
	}
	return
}

// getPatch implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/patches/{patchKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPipeline Retrieves a pipeline using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetPipeline.go.html to see an example of how to use GetPipeline API.
func (client DataIntegrationClient) GetPipeline(ctx context.Context, request GetPipelineRequest) (response GetPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPipelineResponse")
	}
	return
}

// getPipeline implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/pipelines/{pipelineKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPipelineValidation Retrieves a pipeline validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetPipelineValidation.go.html to see an example of how to use GetPipelineValidation API.
func (client DataIntegrationClient) GetPipelineValidation(ctx context.Context, request GetPipelineValidationRequest) (response GetPipelineValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPipelineValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPipelineValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPipelineValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPipelineValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPipelineValidationResponse")
	}
	return
}

// getPipelineValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getPipelineValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/pipelineValidations/{pipelineValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPipelineValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProject Retrieves a project using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetProject.go.html to see an example of how to use GetProject API.
func (client DataIntegrationClient) GetProject(ctx context.Context, request GetProjectRequest) (response GetProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProjectResponse")
	}
	return
}

// getProject implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/projects/{projectKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPublishedObject Retrieves the details of a published object from an application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetPublishedObject.go.html to see an example of how to use GetPublishedObject API.
func (client DataIntegrationClient) GetPublishedObject(ctx context.Context, request GetPublishedObjectRequest) (response GetPublishedObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPublishedObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPublishedObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPublishedObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPublishedObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPublishedObjectResponse")
	}
	return
}

// getPublishedObject implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getPublishedObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/publishedObjects/{publishedObjectKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPublishedObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &publishedobject{})
	return response, err
}

// GetReference Retrieves a reference in an application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetReference.go.html to see an example of how to use GetReference API.
func (client DataIntegrationClient) GetReference(ctx context.Context, request GetReferenceRequest) (response GetReferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReferenceResponse")
	}
	return
}

// getReference implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getReference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/references/{referenceKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSchedule Retrieves schedule by schedule key
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetSchedule.go.html to see an example of how to use GetSchedule API.
func (client DataIntegrationClient) GetSchedule(ctx context.Context, request GetScheduleRequest) (response GetScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduleResponse")
	}
	return
}

// getSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/schedules/{scheduleKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSchema Retrieves a schema that can be accessed using the specified connection.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetSchema.go.html to see an example of how to use GetSchema API.
func (client DataIntegrationClient) GetSchema(ctx context.Context, request GetSchemaRequest) (response GetSchemaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSchema, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSchemaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSchemaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSchemaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSchemaResponse")
	}
	return
}

// getSchema implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getSchema(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connections/{connectionKey}/schemas/{schemaResourceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSchemaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTask Retrieves a task using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetTask.go.html to see an example of how to use GetTask API.
func (client DataIntegrationClient) GetTask(ctx context.Context, request GetTaskRequest) (response GetTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaskResponse")
	}
	return
}

// getTask implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/tasks/{taskKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &task{})
	return response, err
}

// GetTaskRun Retrieves a task run using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetTaskRun.go.html to see an example of how to use GetTaskRun API.
func (client DataIntegrationClient) GetTaskRun(ctx context.Context, request GetTaskRunRequest) (response GetTaskRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTaskRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaskRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaskRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaskRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaskRunResponse")
	}
	return
}

// getTaskRun implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getTaskRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/taskRuns/{taskRunKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaskRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTaskSchedule Endpoint used to get taskSchedule by its key
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetTaskSchedule.go.html to see an example of how to use GetTaskSchedule API.
func (client DataIntegrationClient) GetTaskSchedule(ctx context.Context, request GetTaskScheduleRequest) (response GetTaskScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTaskSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaskScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaskScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaskScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaskScheduleResponse")
	}
	return
}

// getTaskSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getTaskSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/taskSchedules/{taskScheduleKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaskScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTaskValidation Retrieves a task validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetTaskValidation.go.html to see an example of how to use GetTaskValidation API.
func (client DataIntegrationClient) GetTaskValidation(ctx context.Context, request GetTaskValidationRequest) (response GetTaskValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTaskValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaskValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaskValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaskValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaskValidationResponse")
	}
	return
}

// getTaskValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getTaskValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/taskValidations/{taskValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaskValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserDefinedFunction Retrieves a UserDefinedFunction using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetUserDefinedFunction.go.html to see an example of how to use GetUserDefinedFunction API.
func (client DataIntegrationClient) GetUserDefinedFunction(ctx context.Context, request GetUserDefinedFunctionRequest) (response GetUserDefinedFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserDefinedFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserDefinedFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserDefinedFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserDefinedFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserDefinedFunctionResponse")
	}
	return
}

// getUserDefinedFunction implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getUserDefinedFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/userDefinedFunctions/{userDefinedFunctionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserDefinedFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserDefinedFunctionValidation Retrieves a UserDefinedFunction validation using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetUserDefinedFunctionValidation.go.html to see an example of how to use GetUserDefinedFunctionValidation API.
func (client DataIntegrationClient) GetUserDefinedFunctionValidation(ctx context.Context, request GetUserDefinedFunctionValidationRequest) (response GetUserDefinedFunctionValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserDefinedFunctionValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserDefinedFunctionValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserDefinedFunctionValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserDefinedFunctionValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserDefinedFunctionValidationResponse")
	}
	return
}

// getUserDefinedFunctionValidation implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getUserDefinedFunctionValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/userDefinedFunctionValidations/{userDefinedFunctionValidationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserDefinedFunctionValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Retrieves the status of the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client DataIntegrationClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DataIntegrationClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// GetWorkspace Retrieves a Data Integration workspace using the specified identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetWorkspace.go.html to see an example of how to use GetWorkspace API.
func (client DataIntegrationClient) GetWorkspace(ctx context.Context, request GetWorkspaceRequest) (response GetWorkspaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkspace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkspaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkspaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkspaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkspaceResponse")
	}
	return
}

// getWorkspace implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) getWorkspace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkspaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplications Retrieves a list of applications and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListApplications.go.html to see an example of how to use ListApplications API.
func (client DataIntegrationClient) ListApplications(ctx context.Context, request ListApplicationsRequest) (response ListApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationsResponse")
	}
	return
}

// listApplications implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnectionValidations Retrieves a list of connection validations within the specified workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListConnectionValidations.go.html to see an example of how to use ListConnectionValidations API.
func (client DataIntegrationClient) ListConnectionValidations(ctx context.Context, request ListConnectionValidationsRequest) (response ListConnectionValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnectionValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionValidationsResponse")
	}
	return
}

// listConnectionValidations implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listConnectionValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connectionValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnections Retrieves a list of all connections.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListConnections.go.html to see an example of how to use ListConnections API.
func (client DataIntegrationClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionsResponse")
	}
	return
}

// listConnections implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataAssets Retrieves a list of all data asset summaries.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDataAssets.go.html to see an example of how to use ListDataAssets API.
func (client DataIntegrationClient) ListDataAssets(ctx context.Context, request ListDataAssetsRequest) (response ListDataAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataAssetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataAssetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetsResponse")
	}
	return
}

// listDataAssets implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listDataAssets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/dataAssets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataAssetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataEntities Lists a summary of data entities from the data asset using the specified connection.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDataEntities.go.html to see an example of how to use ListDataEntities API.
func (client DataIntegrationClient) ListDataEntities(ctx context.Context, request ListDataEntitiesRequest) (response ListDataEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataEntitiesResponse")
	}
	return
}

// listDataEntities implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listDataEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connections/{connectionKey}/schemas/{schemaResourceName}/dataEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataFlowValidations Retrieves a list of data flow validations within the specified workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDataFlowValidations.go.html to see an example of how to use ListDataFlowValidations API.
func (client DataIntegrationClient) ListDataFlowValidations(ctx context.Context, request ListDataFlowValidationsRequest) (response ListDataFlowValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataFlowValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataFlowValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataFlowValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataFlowValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataFlowValidationsResponse")
	}
	return
}

// listDataFlowValidations implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listDataFlowValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/dataFlowValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataFlowValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataFlows Retrieves a list of data flows in a project or folder.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDataFlows.go.html to see an example of how to use ListDataFlows API.
func (client DataIntegrationClient) ListDataFlows(ctx context.Context, request ListDataFlowsRequest) (response ListDataFlowsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataFlows, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataFlowsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataFlowsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataFlowsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataFlowsResponse")
	}
	return
}

// listDataFlows implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listDataFlows(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/dataFlows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataFlowsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDependentObjects Retrieves a list of all dependent objects for a specific application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDependentObjects.go.html to see an example of how to use ListDependentObjects API.
func (client DataIntegrationClient) ListDependentObjects(ctx context.Context, request ListDependentObjectsRequest) (response ListDependentObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDependentObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDependentObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDependentObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDependentObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDependentObjectsResponse")
	}
	return
}

// listDependentObjects implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listDependentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/dependentObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDependentObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalPublicationValidations Retrieves a lists of external publication validations in a workspace and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListExternalPublicationValidations.go.html to see an example of how to use ListExternalPublicationValidations API.
func (client DataIntegrationClient) ListExternalPublicationValidations(ctx context.Context, request ListExternalPublicationValidationsRequest) (response ListExternalPublicationValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalPublicationValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalPublicationValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalPublicationValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalPublicationValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalPublicationValidationsResponse")
	}
	return
}

// listExternalPublicationValidations implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listExternalPublicationValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublicationValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalPublicationValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalPublications Retrieves a list of external publications in an application and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListExternalPublications.go.html to see an example of how to use ListExternalPublications API.
func (client DataIntegrationClient) ListExternalPublications(ctx context.Context, request ListExternalPublicationsRequest) (response ListExternalPublicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalPublications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalPublicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalPublicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalPublicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalPublicationsResponse")
	}
	return
}

// listExternalPublications implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listExternalPublications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalPublicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFolders Retrieves a list of folders in a project and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListFolders.go.html to see an example of how to use ListFolders API.
func (client DataIntegrationClient) ListFolders(ctx context.Context, request ListFoldersRequest) (response ListFoldersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFoldersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFoldersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFoldersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFoldersResponse")
	}
	return
}

// listFolders implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listFolders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/folders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFoldersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFunctionLibraries Retrieves a list of function libraries in a project and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListFunctionLibraries.go.html to see an example of how to use ListFunctionLibraries API.
func (client DataIntegrationClient) ListFunctionLibraries(ctx context.Context, request ListFunctionLibrariesRequest) (response ListFunctionLibrariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFunctionLibraries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFunctionLibrariesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFunctionLibrariesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFunctionLibrariesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFunctionLibrariesResponse")
	}
	return
}

// listFunctionLibraries implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listFunctionLibraries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/functionLibraries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFunctionLibrariesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatchChanges Retrieves a list of patches in an application and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPatchChanges.go.html to see an example of how to use ListPatchChanges API.
func (client DataIntegrationClient) ListPatchChanges(ctx context.Context, request ListPatchChangesRequest) (response ListPatchChangesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatchChanges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchChangesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchChangesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchChangesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchChangesResponse")
	}
	return
}

// listPatchChanges implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listPatchChanges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/patchChanges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchChangesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatches Retrieves a list of patches in an application and provides options to filter the list. For listing changes based on a period and logical objects changed, see ListPatchChanges API.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPatches.go.html to see an example of how to use ListPatches API.
func (client DataIntegrationClient) ListPatches(ctx context.Context, request ListPatchesRequest) (response ListPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchesResponse")
	}
	return
}

// listPatches implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelineValidations Retrieves a list of pipeline validations within the specified workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPipelineValidations.go.html to see an example of how to use ListPipelineValidations API.
func (client DataIntegrationClient) ListPipelineValidations(ctx context.Context, request ListPipelineValidationsRequest) (response ListPipelineValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelineValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelineValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelineValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelineValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelineValidationsResponse")
	}
	return
}

// listPipelineValidations implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listPipelineValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/pipelineValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelineValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelines Retrieves a list of pipelines in a project or folder from within a workspace, the query parameter specifies the project or folder.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPipelines.go.html to see an example of how to use ListPipelines API.
func (client DataIntegrationClient) ListPipelines(ctx context.Context, request ListPipelinesRequest) (response ListPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelinesResponse")
	}
	return
}

// listPipelines implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/pipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProjects Retrieves a lists of projects in a workspace and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListProjects.go.html to see an example of how to use ListProjects API.
func (client DataIntegrationClient) ListProjects(ctx context.Context, request ListProjectsRequest) (response ListProjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProjectsResponse")
	}
	return
}

// listProjects implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listProjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/projects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPublishedObjects Retrieves a list of all the published objects for a specified application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPublishedObjects.go.html to see an example of how to use ListPublishedObjects API.
func (client DataIntegrationClient) ListPublishedObjects(ctx context.Context, request ListPublishedObjectsRequest) (response ListPublishedObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublishedObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublishedObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublishedObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublishedObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublishedObjectsResponse")
	}
	return
}

// listPublishedObjects implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listPublishedObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/publishedObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPublishedObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReferences Retrieves a list of references in an application. Reference objects are created when dataflows and tasks use objects, such as data assets and connections.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListReferences.go.html to see an example of how to use ListReferences API.
func (client DataIntegrationClient) ListReferences(ctx context.Context, request ListReferencesRequest) (response ListReferencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReferences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReferencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReferencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReferencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReferencesResponse")
	}
	return
}

// listReferences implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listReferences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/references", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReferencesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchedules Use this endpoint to list schedules.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListSchedules.go.html to see an example of how to use ListSchedules API.
func (client DataIntegrationClient) ListSchedules(ctx context.Context, request ListSchedulesRequest) (response ListSchedulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSchedules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSchedulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSchedulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSchedulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSchedulesResponse")
	}
	return
}

// listSchedules implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listSchedules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/schedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchedulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchemas Retrieves a list of all the schemas that can be accessed using the specified connection.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListSchemas.go.html to see an example of how to use ListSchemas API.
func (client DataIntegrationClient) ListSchemas(ctx context.Context, request ListSchemasRequest) (response ListSchemasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSchemas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSchemasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSchemasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSchemasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSchemasResponse")
	}
	return
}

// listSchemas implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listSchemas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/connections/{connectionKey}/schemas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchemasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskRunLogs Gets log entries for task runs using its key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskRunLogs.go.html to see an example of how to use ListTaskRunLogs API.
func (client DataIntegrationClient) ListTaskRunLogs(ctx context.Context, request ListTaskRunLogsRequest) (response ListTaskRunLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaskRunLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaskRunLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaskRunLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaskRunLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaskRunLogsResponse")
	}
	return
}

// listTaskRunLogs implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listTaskRunLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/taskRuns/{taskRunKey}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaskRunLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskRuns Retrieves a list of task runs and provides options to filter the list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskRuns.go.html to see an example of how to use ListTaskRuns API.
func (client DataIntegrationClient) ListTaskRuns(ctx context.Context, request ListTaskRunsRequest) (response ListTaskRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaskRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaskRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaskRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaskRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaskRunsResponse")
	}
	return
}

// listTaskRuns implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listTaskRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/taskRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaskRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskSchedules This endpoint can be used to get the list of all the TaskSchedule objects.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskSchedules.go.html to see an example of how to use ListTaskSchedules API.
func (client DataIntegrationClient) ListTaskSchedules(ctx context.Context, request ListTaskSchedulesRequest) (response ListTaskSchedulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaskSchedules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaskSchedulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaskSchedulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaskSchedulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaskSchedulesResponse")
	}
	return
}

// listTaskSchedules implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listTaskSchedules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/applications/{applicationKey}/taskSchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaskSchedulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskValidations Retrieves a list of task validations within the specified workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskValidations.go.html to see an example of how to use ListTaskValidations API.
func (client DataIntegrationClient) ListTaskValidations(ctx context.Context, request ListTaskValidationsRequest) (response ListTaskValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaskValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaskValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaskValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaskValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaskValidationsResponse")
	}
	return
}

// listTaskValidations implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listTaskValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/taskValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaskValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTasks Retrieves a list of all tasks in a specified project or folder.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTasks.go.html to see an example of how to use ListTasks API.
func (client DataIntegrationClient) ListTasks(ctx context.Context, request ListTasksRequest) (response ListTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTasksResponse")
	}
	return
}

// listTasks implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/tasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTasksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserDefinedFunctionValidations Retrieves a list of UserDefinedFunctionvalidations within the specified workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListUserDefinedFunctionValidations.go.html to see an example of how to use ListUserDefinedFunctionValidations API.
func (client DataIntegrationClient) ListUserDefinedFunctionValidations(ctx context.Context, request ListUserDefinedFunctionValidationsRequest) (response ListUserDefinedFunctionValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserDefinedFunctionValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserDefinedFunctionValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserDefinedFunctionValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserDefinedFunctionValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserDefinedFunctionValidationsResponse")
	}
	return
}

// listUserDefinedFunctionValidations implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listUserDefinedFunctionValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/userDefinedFunctionValidations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserDefinedFunctionValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserDefinedFunctions Retrieves a list of UserDefinedFunctions in a function library.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListUserDefinedFunctions.go.html to see an example of how to use ListUserDefinedFunctions API.
func (client DataIntegrationClient) ListUserDefinedFunctions(ctx context.Context, request ListUserDefinedFunctionsRequest) (response ListUserDefinedFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserDefinedFunctions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserDefinedFunctionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserDefinedFunctionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserDefinedFunctionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserDefinedFunctionsResponse")
	}
	return
}

// listUserDefinedFunctions implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listUserDefinedFunctions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces/{workspaceId}/userDefinedFunctions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserDefinedFunctionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Retrieves a paginated list of errors for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client DataIntegrationClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DataIntegrationClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/workRequestErrors", binaryReqBody, extraHeaders)
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

// ListWorkRequestLogs Retrieves a paginated list of logs for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client DataIntegrationClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DataIntegrationClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// ListWorkRequests Lists the work requests in a compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client DataIntegrationClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DataIntegrationClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// ListWorkspaces Retrieves a list of Data Integration workspaces.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkspaces.go.html to see an example of how to use ListWorkspaces API.
func (client DataIntegrationClient) ListWorkspaces(ctx context.Context, request ListWorkspacesRequest) (response ListWorkspacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkspaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkspacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkspacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkspacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkspacesResponse")
	}
	return
}

// listWorkspaces implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) listWorkspaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workspaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkspacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartWorkspace Starts a workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/StartWorkspace.go.html to see an example of how to use StartWorkspace API.
func (client DataIntegrationClient) StartWorkspace(ctx context.Context, request StartWorkspaceRequest) (response StartWorkspaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startWorkspace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartWorkspaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartWorkspaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartWorkspaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartWorkspaceResponse")
	}
	return
}

// startWorkspace implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) startWorkspace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartWorkspaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopWorkspace Stops a workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/StopWorkspace.go.html to see an example of how to use StopWorkspace API.
func (client DataIntegrationClient) StopWorkspace(ctx context.Context, request StopWorkspaceRequest) (response StopWorkspaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopWorkspace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopWorkspaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopWorkspaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopWorkspaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopWorkspaceResponse")
	}
	return
}

// stopWorkspace implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) stopWorkspace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workspaces/{workspaceId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopWorkspaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateApplication Updates an application.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateApplication.go.html to see an example of how to use UpdateApplication API.
func (client DataIntegrationClient) UpdateApplication(ctx context.Context, request UpdateApplicationRequest) (response UpdateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApplicationResponse")
	}
	return
}

// updateApplication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/applications/{applicationKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConnection Updates a connection under a data asset.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateConnection.go.html to see an example of how to use UpdateConnection API.
func (client DataIntegrationClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectionResponse")
	}
	return
}

// updateConnection implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/connections/{connectionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// UpdateDataAsset Updates a specific data asset with default connection.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateDataAsset.go.html to see an example of how to use UpdateDataAsset API.
func (client DataIntegrationClient) UpdateDataAsset(ctx context.Context, request UpdateDataAssetRequest) (response UpdateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataAssetResponse")
	}
	return
}

// updateDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/dataAssets/{dataAssetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dataasset{})
	return response, err
}

// UpdateDataFlow Updates a specific data flow.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateDataFlow.go.html to see an example of how to use UpdateDataFlow API.
func (client DataIntegrationClient) UpdateDataFlow(ctx context.Context, request UpdateDataFlowRequest) (response UpdateDataFlowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataFlow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataFlowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataFlowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataFlowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataFlowResponse")
	}
	return
}

// updateDataFlow implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateDataFlow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/dataFlows/{dataFlowKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataFlowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExternalPublication Updates the external publication object.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateExternalPublication.go.html to see an example of how to use UpdateExternalPublication API.
func (client DataIntegrationClient) UpdateExternalPublication(ctx context.Context, request UpdateExternalPublicationRequest) (response UpdateExternalPublicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExternalPublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExternalPublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExternalPublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExternalPublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExternalPublicationResponse")
	}
	return
}

// updateExternalPublication implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateExternalPublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/tasks/{taskKey}/externalPublications/{externalPublicationsKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExternalPublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFolder Updates a specific folder.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateFolder.go.html to see an example of how to use UpdateFolder API.
func (client DataIntegrationClient) UpdateFolder(ctx context.Context, request UpdateFolderRequest) (response UpdateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFolderResponse")
	}
	return
}

// updateFolder implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/folders/{folderKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFunctionLibrary Updates a specific Function Library.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateFunctionLibrary.go.html to see an example of how to use UpdateFunctionLibrary API.
func (client DataIntegrationClient) UpdateFunctionLibrary(ctx context.Context, request UpdateFunctionLibraryRequest) (response UpdateFunctionLibraryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFunctionLibrary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFunctionLibraryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFunctionLibraryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFunctionLibraryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFunctionLibraryResponse")
	}
	return
}

// updateFunctionLibrary implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateFunctionLibrary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/functionLibraries/{functionLibraryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFunctionLibraryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePipeline Updates a specific pipeline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdatePipeline.go.html to see an example of how to use UpdatePipeline API.
func (client DataIntegrationClient) UpdatePipeline(ctx context.Context, request UpdatePipelineRequest) (response UpdatePipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePipelineResponse")
	}
	return
}

// updatePipeline implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updatePipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/pipelines/{pipelineKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProject Updates a specific project.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateProject.go.html to see an example of how to use UpdateProject API.
func (client DataIntegrationClient) UpdateProject(ctx context.Context, request UpdateProjectRequest) (response UpdateProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProjectResponse")
	}
	return
}

// updateProject implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/projects/{projectKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateReference Updates the application references. For example, to map a data asset to a different target object.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateReference.go.html to see an example of how to use UpdateReference API.
func (client DataIntegrationClient) UpdateReference(ctx context.Context, request UpdateReferenceRequest) (response UpdateReferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateReference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateReferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateReferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateReferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateReferenceResponse")
	}
	return
}

// updateReference implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateReference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/applications/{applicationKey}/references/{referenceKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateReferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSchedule Endpoint used to update the schedule
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateSchedule.go.html to see an example of how to use UpdateSchedule API.
func (client DataIntegrationClient) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (response UpdateScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduleResponse")
	}
	return
}

// updateSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/applications/{applicationKey}/schedules/{scheduleKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTask Updates a specific task. For example, you can update the task description or move the task to a different folder by changing the `aggregatorKey` to a different folder in the registry.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateTask.go.html to see an example of how to use UpdateTask API.
func (client DataIntegrationClient) UpdateTask(ctx context.Context, request UpdateTaskRequest) (response UpdateTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTaskResponse")
	}
	return
}

// updateTask implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/tasks/{taskKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &task{})
	return response, err
}

// UpdateTaskRun Updates the status of the task run. For example, aborts a task run.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateTaskRun.go.html to see an example of how to use UpdateTaskRun API.
func (client DataIntegrationClient) UpdateTaskRun(ctx context.Context, request UpdateTaskRunRequest) (response UpdateTaskRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTaskRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTaskRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTaskRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTaskRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTaskRunResponse")
	}
	return
}

// updateTaskRun implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateTaskRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/applications/{applicationKey}/taskRuns/{taskRunKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTaskRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTaskSchedule Endpoint used to update the TaskSchedule
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateTaskSchedule.go.html to see an example of how to use UpdateTaskSchedule API.
func (client DataIntegrationClient) UpdateTaskSchedule(ctx context.Context, request UpdateTaskScheduleRequest) (response UpdateTaskScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTaskSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTaskScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTaskScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTaskScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTaskScheduleResponse")
	}
	return
}

// updateTaskSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateTaskSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/applications/{applicationKey}/taskSchedules/{taskScheduleKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTaskScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUserDefinedFunction Updates a specific UserDefinedFunction.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateUserDefinedFunction.go.html to see an example of how to use UpdateUserDefinedFunction API.
func (client DataIntegrationClient) UpdateUserDefinedFunction(ctx context.Context, request UpdateUserDefinedFunctionRequest) (response UpdateUserDefinedFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUserDefinedFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUserDefinedFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUserDefinedFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUserDefinedFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUserDefinedFunctionResponse")
	}
	return
}

// updateUserDefinedFunction implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateUserDefinedFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}/userDefinedFunctions/{userDefinedFunctionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUserDefinedFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWorkspace Updates the specified Data Integration workspace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/UpdateWorkspace.go.html to see an example of how to use UpdateWorkspace API.
func (client DataIntegrationClient) UpdateWorkspace(ctx context.Context, request UpdateWorkspaceRequest) (response UpdateWorkspaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWorkspace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWorkspaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWorkspaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWorkspaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWorkspaceResponse")
	}
	return
}

// updateWorkspace implements the OCIOperation interface (enables retrying operations)
func (client DataIntegrationClient) updateWorkspace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/workspaces/{workspaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWorkspaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
