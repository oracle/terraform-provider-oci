// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//DataSafeClient a client for DataSafe
type DataSafeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataSafeClientWithConfigurationProvider Creates a new default DataSafe client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataSafeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataSafeClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDataSafeClientFromBaseClient(baseClient, provider)
}

// NewDataSafeClientWithOboToken Creates a new default DataSafe client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataSafeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataSafeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataSafeClientFromBaseClient(baseClient, configProvider)
}

func newDataSafeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataSafeClient, err error) {
	// DataSafe service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataSafeClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataSafeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datasafe", "https://datasafe.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataSafeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataSafeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateTargetDatabase Reactivates a previously deactivated Data Safe target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ActivateTargetDatabase.go.html to see an example of how to use ActivateTargetDatabase API.
func (client DataSafeClient) ActivateTargetDatabase(ctx context.Context, request ActivateTargetDatabaseRequest) (response ActivateTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.activateTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateTargetDatabaseResponse")
	}
	return
}

// activateTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) activateTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDataSafePrivateEndpointCompartment Moves the Data Safe private endpoint and its dependent resources to the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeDataSafePrivateEndpointCompartment.go.html to see an example of how to use ChangeDataSafePrivateEndpointCompartment API.
func (client DataSafeClient) ChangeDataSafePrivateEndpointCompartment(ctx context.Context, request ChangeDataSafePrivateEndpointCompartmentRequest) (response ChangeDataSafePrivateEndpointCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDataSafePrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDataSafePrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDataSafePrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDataSafePrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDataSafePrivateEndpointCompartmentResponse")
	}
	return
}

// changeDataSafePrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeDataSafePrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDataSafePrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOnPremConnectorCompartment Moves the specified on-premises connector into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeOnPremConnectorCompartment.go.html to see an example of how to use ChangeOnPremConnectorCompartment API.
func (client DataSafeClient) ChangeOnPremConnectorCompartment(ctx context.Context, request ChangeOnPremConnectorCompartmentRequest) (response ChangeOnPremConnectorCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeOnPremConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOnPremConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOnPremConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOnPremConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOnPremConnectorCompartmentResponse")
	}
	return
}

// changeOnPremConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeOnPremConnectorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/onPremConnectors/{onPremConnectorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOnPremConnectorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityAssessmentCompartment Moves the specified saved security assessment or future scheduled assessments into a different compartment.
// To start, call first the operation ListSecurityAssessments with filters "type = save_schedule". This returns the scheduleAssessmentId. Then, call this changeCompartment with the scheduleAssessmentId.
// The existing saved security assessments created due to the schedule are not moved. However, all new saves will be associated with the new compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSecurityAssessmentCompartment.go.html to see an example of how to use ChangeSecurityAssessmentCompartment API.
func (client DataSafeClient) ChangeSecurityAssessmentCompartment(ctx context.Context, request ChangeSecurityAssessmentCompartmentRequest) (response ChangeSecurityAssessmentCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityAssessmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityAssessmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityAssessmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityAssessmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityAssessmentCompartmentResponse")
	}
	return
}

// changeSecurityAssessmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSecurityAssessmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityAssessmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTargetDatabaseCompartment Moves the Data Safe target database to the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeTargetDatabaseCompartment.go.html to see an example of how to use ChangeTargetDatabaseCompartment API.
func (client DataSafeClient) ChangeTargetDatabaseCompartment(ctx context.Context, request ChangeTargetDatabaseCompartmentRequest) (response ChangeTargetDatabaseCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeTargetDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTargetDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTargetDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTargetDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTargetDatabaseCompartmentResponse")
	}
	return
}

// changeTargetDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeTargetDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTargetDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeUserAssessmentCompartment Moves the specified saved user assessment or future scheduled assessments into a different compartment.
// To start storing scheduled user assessments on a different compartment, first call the operation ListUserAssessments with
// the filters "type = save_schedule". That call returns the scheduleAssessmentId. Then call
// ChangeUserAssessmentCompartment with the scheduleAssessmentId. The existing saved user assessments created per the schedule
// are not be moved. However, all new saves will be associated with the new compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeUserAssessmentCompartment.go.html to see an example of how to use ChangeUserAssessmentCompartment API.
func (client DataSafeClient) ChangeUserAssessmentCompartment(ctx context.Context, request ChangeUserAssessmentCompartmentRequest) (response ChangeUserAssessmentCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeUserAssessmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeUserAssessmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeUserAssessmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeUserAssessmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeUserAssessmentCompartmentResponse")
	}
	return
}

// changeUserAssessmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeUserAssessmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeUserAssessmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CompareSecurityAssessment Compares two security assessments. For this comparison, a security assessment can be a saved assessment, a latest assessment, or a baseline assessment.
// For example, you can compare saved assessment or a latest assessment against a baseline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CompareSecurityAssessment.go.html to see an example of how to use CompareSecurityAssessment API.
func (client DataSafeClient) CompareSecurityAssessment(ctx context.Context, request CompareSecurityAssessmentRequest) (response CompareSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.compareSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompareSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompareSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompareSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompareSecurityAssessmentResponse")
	}
	return
}

// compareSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) compareSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/compare", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompareSecurityAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CompareUserAssessment Compares two user assessments. For this comparison, a user assessment can be a saved, a latest assessment, or a baseline.
// As an example, it can be used to compare a user assessment saved or a latest assessment with a baseline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CompareUserAssessment.go.html to see an example of how to use CompareUserAssessment API.
func (client DataSafeClient) CompareUserAssessment(ctx context.Context, request CompareUserAssessmentRequest) (response CompareUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.compareUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompareUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompareUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompareUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompareUserAssessmentResponse")
	}
	return
}

// compareUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) compareUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/compare", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompareUserAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataSafePrivateEndpoint Creates a new Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateDataSafePrivateEndpoint.go.html to see an example of how to use CreateDataSafePrivateEndpoint API.
func (client DataSafeClient) CreateDataSafePrivateEndpoint(ctx context.Context, request CreateDataSafePrivateEndpointRequest) (response CreateDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataSafePrivateEndpointResponse")
	}
	return
}

// createDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSafePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataSafePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOnPremConnector Creates a new on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateOnPremConnector.go.html to see an example of how to use CreateOnPremConnector API.
func (client DataSafeClient) CreateOnPremConnector(ctx context.Context, request CreateOnPremConnectorRequest) (response CreateOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOnPremConnectorResponse")
	}
	return
}

// createOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/onPremConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOnPremConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityAssessment Creates a new saved security assessment for one or multiple targets in a compartment. When this operation is performed,
// it will save the latest assessments in the specified compartment. If a schedule is passed, it will persist the latest assessments,
// at the defined date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSecurityAssessment.go.html to see an example of how to use CreateSecurityAssessment API.
func (client DataSafeClient) CreateSecurityAssessment(ctx context.Context, request CreateSecurityAssessmentRequest) (response CreateSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityAssessmentResponse")
	}
	return
}

// createSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetDatabase Registers the specified database with Data Safe and creates a Data Safe target database in the Data Safe Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateTargetDatabase.go.html to see an example of how to use CreateTargetDatabase API.
func (client DataSafeClient) CreateTargetDatabase(ctx context.Context, request CreateTargetDatabaseRequest) (response CreateTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetDatabaseResponse")
	}
	return
}

// createTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUserAssessment Creates a new saved user assessment for one or multiple targets in a compartment. It saves the latest assessments in the
// specified compartment. If a scheduled is passed in, this operation persists the latest assessments that exist at the defined
// date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateUserAssessment.go.html to see an example of how to use CreateUserAssessment API.
func (client DataSafeClient) CreateUserAssessment(ctx context.Context, request CreateUserAssessmentRequest) (response CreateUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserAssessmentResponse")
	}
	return
}

// createUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateTargetDatabase Deactivates a target database in Data Safe.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeactivateTargetDatabase.go.html to see an example of how to use DeactivateTargetDatabase API.
func (client DataSafeClient) DeactivateTargetDatabase(ctx context.Context, request DeactivateTargetDatabaseRequest) (response DeactivateTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deactivateTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateTargetDatabaseResponse")
	}
	return
}

// deactivateTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deactivateTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataSafePrivateEndpoint Deletes the specified Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDataSafePrivateEndpoint.go.html to see an example of how to use DeleteDataSafePrivateEndpoint API.
func (client DataSafeClient) DeleteDataSafePrivateEndpoint(ctx context.Context, request DeleteDataSafePrivateEndpointRequest) (response DeleteDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataSafePrivateEndpointResponse")
	}
	return
}

// deleteDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataSafePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOnPremConnector Deletes the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteOnPremConnector.go.html to see an example of how to use DeleteOnPremConnector API.
func (client DataSafeClient) DeleteOnPremConnector(ctx context.Context, request DeleteOnPremConnectorRequest) (response DeleteOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOnPremConnectorResponse")
	}
	return
}

// deleteOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/onPremConnectors/{onPremConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOnPremConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityAssessment Deletes the specified saved security assessment or schedule. To delete a security assessment schedule,
// first call the operation ListSecurityAssessments with filters "type = save_schedule".
// That operation returns the scheduleAssessmentId. Then, call DeleteSecurityAssessment with the scheduleAssessmentId.
// If the assessment being deleted is the baseline for that compartment, then it will impact all baselines in the compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSecurityAssessment.go.html to see an example of how to use DeleteSecurityAssessment API.
func (client DataSafeClient) DeleteSecurityAssessment(ctx context.Context, request DeleteSecurityAssessmentRequest) (response DeleteSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityAssessmentResponse")
	}
	return
}

// deleteSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/securityAssessments/{securityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetDatabase Deregisters the specified database from Data Safe and removes the target database from the Data Safe Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteTargetDatabase.go.html to see an example of how to use DeleteTargetDatabase API.
func (client DataSafeClient) DeleteTargetDatabase(ctx context.Context, request DeleteTargetDatabaseRequest) (response DeleteTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetDatabaseResponse")
	}
	return
}

// deleteTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targetDatabases/{targetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUserAssessment Deletes the specified saved user assessment or schedule. To delete a user assessment schedule, first call the operation
// ListUserAssessments with filters "type = save_schedule".
// That call returns the scheduleAssessmentId. Then call DeleteUserAssessment with the scheduleAssessmentId.
// If the assessment being deleted is the baseline for that compartment, then it will impact all baselines in the compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteUserAssessment.go.html to see an example of how to use DeleteUserAssessment API.
func (client DataSafeClient) DeleteUserAssessment(ctx context.Context, request DeleteUserAssessmentRequest) (response DeleteUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserAssessmentResponse")
	}
	return
}

// deleteUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/userAssessments/{userAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadPrivilegeScript Downloads the privilege script to grant/revoke required roles from the Data Safe account on the target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadPrivilegeScript.go.html to see an example of how to use DownloadPrivilegeScript API.
func (client DataSafeClient) DownloadPrivilegeScript(ctx context.Context, request DownloadPrivilegeScriptRequest) (response DownloadPrivilegeScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.downloadPrivilegeScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadPrivilegeScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadPrivilegeScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadPrivilegeScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadPrivilegeScriptResponse")
	}
	return
}

// downloadPrivilegeScript implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadPrivilegeScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/downloadPrivilegeScript", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadPrivilegeScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadSecurityAssessmentReport Downloads the report of the specified security assessment. To download the security assessment report, it needs to be generated first.
// Please use GenerateSecurityAssessmentReport to generate a downloadable report in the preferred format (PDF, XLS).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadSecurityAssessmentReport.go.html to see an example of how to use DownloadSecurityAssessmentReport API.
func (client DataSafeClient) DownloadSecurityAssessmentReport(ctx context.Context, request DownloadSecurityAssessmentReportRequest) (response DownloadSecurityAssessmentReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.downloadSecurityAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadSecurityAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadSecurityAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadSecurityAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadSecurityAssessmentReportResponse")
	}
	return
}

// downloadSecurityAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadSecurityAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/downloadReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadSecurityAssessmentReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadUserAssessmentReport Downloads the report of the specified user assessment. To download the user assessment report, it needs to be generated first.
// Please use GenerateUserAssessmentReport to generate a downloadable report in the preferred format (PDF, XLS).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadUserAssessmentReport.go.html to see an example of how to use DownloadUserAssessmentReport API.
func (client DataSafeClient) DownloadUserAssessmentReport(ctx context.Context, request DownloadUserAssessmentReportRequest) (response DownloadUserAssessmentReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.downloadUserAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadUserAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadUserAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadUserAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadUserAssessmentReportResponse")
	}
	return
}

// downloadUserAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadUserAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/downloadReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadUserAssessmentReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDataSafeConfiguration Enables Data Safe in the tenancy and region.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/EnableDataSafeConfiguration.go.html to see an example of how to use EnableDataSafeConfiguration API.
func (client DataSafeClient) EnableDataSafeConfiguration(ctx context.Context, request EnableDataSafeConfigurationRequest) (response EnableDataSafeConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableDataSafeConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableDataSafeConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableDataSafeConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableDataSafeConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableDataSafeConfigurationResponse")
	}
	return
}

// enableDataSafeConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) enableDataSafeConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableDataSafeConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateOnPremConnectorConfiguration Creates and downloads the configuration of the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateOnPremConnectorConfiguration.go.html to see an example of how to use GenerateOnPremConnectorConfiguration API.
func (client DataSafeClient) GenerateOnPremConnectorConfiguration(ctx context.Context, request GenerateOnPremConnectorConfigurationRequest) (response GenerateOnPremConnectorConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.generateOnPremConnectorConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateOnPremConnectorConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateOnPremConnectorConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateOnPremConnectorConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateOnPremConnectorConfigurationResponse")
	}
	return
}

// generateOnPremConnectorConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateOnPremConnectorConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/onPremConnectors/{onPremConnectorId}/actions/generateConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateOnPremConnectorConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateSecurityAssessmentReport Generates the report of the specified security assessment. You can get the report in PDF or XLS format.
// After generating the report, use DownloadSecurityAssessmentReport to download it in the preferred format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateSecurityAssessmentReport.go.html to see an example of how to use GenerateSecurityAssessmentReport API.
func (client DataSafeClient) GenerateSecurityAssessmentReport(ctx context.Context, request GenerateSecurityAssessmentReportRequest) (response GenerateSecurityAssessmentReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.generateSecurityAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateSecurityAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateSecurityAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateSecurityAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateSecurityAssessmentReportResponse")
	}
	return
}

// generateSecurityAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateSecurityAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/generateReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateSecurityAssessmentReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateUserAssessmentReport Generates the report of the specified user assessment. The report is available in PDF or XLS format.
// After generating the report, use DownloadUserAssessmentReport to download it in the preferred format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateUserAssessmentReport.go.html to see an example of how to use GenerateUserAssessmentReport API.
func (client DataSafeClient) GenerateUserAssessmentReport(ctx context.Context, request GenerateUserAssessmentReportRequest) (response GenerateUserAssessmentReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.generateUserAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateUserAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateUserAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateUserAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateUserAssessmentReportResponse")
	}
	return
}

// generateUserAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateUserAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/generateReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateUserAssessmentReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataSafeConfiguration Gets the details of the Data Safe configuration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDataSafeConfiguration.go.html to see an example of how to use GetDataSafeConfiguration API.
func (client DataSafeClient) GetDataSafeConfiguration(ctx context.Context, request GetDataSafeConfigurationRequest) (response GetDataSafeConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataSafeConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataSafeConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataSafeConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataSafeConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataSafeConfigurationResponse")
	}
	return
}

// getDataSafeConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDataSafeConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataSafeConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataSafePrivateEndpoint Gets the details of the specified Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDataSafePrivateEndpoint.go.html to see an example of how to use GetDataSafePrivateEndpoint API.
func (client DataSafeClient) GetDataSafePrivateEndpoint(ctx context.Context, request GetDataSafePrivateEndpointRequest) (response GetDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataSafePrivateEndpointResponse")
	}
	return
}

// getDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataSafePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOnPremConnector Gets the details of the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetOnPremConnector.go.html to see an example of how to use GetOnPremConnector API.
func (client DataSafeClient) GetOnPremConnector(ctx context.Context, request GetOnPremConnectorRequest) (response GetOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOnPremConnectorResponse")
	}
	return
}

// getOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/onPremConnectors/{onPremConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOnPremConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAssessment Gets the details of the specified security assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessment.go.html to see an example of how to use GetSecurityAssessment API.
func (client DataSafeClient) GetSecurityAssessment(ctx context.Context, request GetSecurityAssessmentRequest) (response GetSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAssessmentResponse")
	}
	return
}

// getSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAssessmentComparison Gets the details of the comparison report on the security assessments submitted for comparison.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessmentComparison.go.html to see an example of how to use GetSecurityAssessmentComparison API.
func (client DataSafeClient) GetSecurityAssessmentComparison(ctx context.Context, request GetSecurityAssessmentComparisonRequest) (response GetSecurityAssessmentComparisonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAssessmentComparison, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAssessmentComparisonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAssessmentComparisonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAssessmentComparisonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAssessmentComparisonResponse")
	}
	return
}

// getSecurityAssessmentComparison implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityAssessmentComparison(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}/comparison/{comparisonSecurityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAssessmentComparisonResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetDatabase Returns the details of the specified Data Safe target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetTargetDatabase.go.html to see an example of how to use GetTargetDatabase API.
func (client DataSafeClient) GetTargetDatabase(ctx context.Context, request GetTargetDatabaseRequest) (response GetTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetDatabaseResponse")
	}
	return
}

// getTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserAssessment Gets a user assessment by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetUserAssessment.go.html to see an example of how to use GetUserAssessment API.
func (client DataSafeClient) GetUserAssessment(ctx context.Context, request GetUserAssessmentRequest) (response GetUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserAssessmentResponse")
	}
	return
}

// getUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserAssessmentComparison Gets the details of the comparison report for the user assessments provided.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetUserAssessmentComparison.go.html to see an example of how to use GetUserAssessmentComparison API.
func (client DataSafeClient) GetUserAssessmentComparison(ctx context.Context, request GetUserAssessmentComparisonRequest) (response GetUserAssessmentComparisonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserAssessmentComparison, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserAssessmentComparisonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserAssessmentComparisonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserAssessmentComparisonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserAssessmentComparisonResponse")
	}
	return
}

// getUserAssessmentComparison implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getUserAssessmentComparison(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/comparison/{comparisonUserAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserAssessmentComparisonResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client DataSafeClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DataSafeClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// ListDataSafePrivateEndpoints Gets a list of Data Safe private endpoints.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDataSafePrivateEndpoints.go.html to see an example of how to use ListDataSafePrivateEndpoints API.
func (client DataSafeClient) ListDataSafePrivateEndpoints(ctx context.Context, request ListDataSafePrivateEndpointsRequest) (response ListDataSafePrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataSafePrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataSafePrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataSafePrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataSafePrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataSafePrivateEndpointsResponse")
	}
	return
}

// listDataSafePrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDataSafePrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSafePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataSafePrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFindings List all the findings from all the targets in the specified assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindings.go.html to see an example of how to use ListFindings API.
func (client DataSafeClient) ListFindings(ctx context.Context, request ListFindingsRequest) (response ListFindingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFindings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFindingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFindingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFindingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFindingsResponse")
	}
	return
}

// listFindings implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listFindings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}/findings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFindingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGrants Gets a list of grants for a particular user in the specified user assessment. A user grant contains details such as the
// privilege name, type, category, and depth level. The depth level indicates how deep in the hierarchy of roles granted to
// roles a privilege grant is. The userKey in this operation is a system-generated identifier. Perform the operation ListUsers
// to get the userKey for a particular user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListGrants.go.html to see an example of how to use ListGrants API.
func (client DataSafeClient) ListGrants(ctx context.Context, request ListGrantsRequest) (response ListGrantsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGrants, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGrantsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGrantsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGrantsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGrantsResponse")
	}
	return
}

// listGrants implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listGrants(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/users/{userKey}/grants", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGrantsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOnPremConnectors Gets a list of on-premises connectors.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListOnPremConnectors.go.html to see an example of how to use ListOnPremConnectors API.
func (client DataSafeClient) ListOnPremConnectors(ctx context.Context, request ListOnPremConnectorsRequest) (response ListOnPremConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOnPremConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOnPremConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOnPremConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOnPremConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOnPremConnectorsResponse")
	}
	return
}

// listOnPremConnectors implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listOnPremConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/onPremConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOnPremConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityAssessments Gets a list of security assessments.
// The ListSecurityAssessments operation returns only the assessments in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSecurityAssessments on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityAssessments.go.html to see an example of how to use ListSecurityAssessments API.
func (client DataSafeClient) ListSecurityAssessments(ctx context.Context, request ListSecurityAssessmentsRequest) (response ListSecurityAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAssessments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAssessmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAssessmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAssessmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAssessmentsResponse")
	}
	return
}

// listSecurityAssessments implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityAssessments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAssessmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetDatabases Returns the list of registered target databases in Data Safe.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetDatabases.go.html to see an example of how to use ListTargetDatabases API.
func (client DataSafeClient) ListTargetDatabases(ctx context.Context, request ListTargetDatabasesRequest) (response ListTargetDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetDatabasesResponse")
	}
	return
}

// listTargetDatabases implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listTargetDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserAnalytics Gets a list of aggregated user details from the specified user assessment. This provides information about the overall state
// of database user security.  For example, the user details include how many users have the DBA role and how many users are in
// the critical category. This data is especially useful content for dashboards or to support analytics.
// When you perform the ListUserAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
// parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
// permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
// root compartment. If the requestor does not have access to at least one subcompartment of the compartment specified by
// compartmentId, then "Not Authorized" is returned.
// The parameter compartmentIdInSubtree applies when you perform ListUserAnalytics on the compartmentId passed and when it is
// set to true, the entire hierarchy of compartments can be returned.
// To use ListUserAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
// set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAnalytics.go.html to see an example of how to use ListUserAnalytics API.
func (client DataSafeClient) ListUserAnalytics(ctx context.Context, request ListUserAnalyticsRequest) (response ListUserAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserAnalyticsResponse")
	}
	return
}

// listUserAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listUserAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/userAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserAssessments Gets a list of user assessments.
// The ListUserAssessments operation returns only the assessments in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListUserAssessments on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAssessments.go.html to see an example of how to use ListUserAssessments API.
func (client DataSafeClient) ListUserAssessments(ctx context.Context, request ListUserAssessmentsRequest) (response ListUserAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserAssessments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserAssessmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserAssessmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserAssessmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserAssessmentsResponse")
	}
	return
}

// listUserAssessments implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listUserAssessments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserAssessmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUsers Gets a list of users of the specified user assessment. The result contains the database user details for each user, such
// as user type, account status, last login time, user creation time, authentication type, user profile, and the date and time
// of the latest password change. It also contains the user category derived from these user details as well as privileges
// granted to each user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUsers.go.html to see an example of how to use ListUsers API.
func (client DataSafeClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
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
func (client DataSafeClient) listUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/users", binaryReqBody, extraHeaders)
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

// ListWorkRequestErrors Gets a list of errors for the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client DataSafeClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DataSafeClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// ListWorkRequestLogs Gets a list of log entries for the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client DataSafeClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DataSafeClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// ListWorkRequests Gets a list of work requests.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client DataSafeClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DataSafeClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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

// RefreshSecurityAssessment Runs a security assessment, refreshes the latest assessment, and saves it for future reference.
// The assessment runs with a securityAssessmentId of type LATEST. Before you start, first call the ListSecurityAssessments operation with filter "type = latest" to get the security assessment id for the target's latest assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshSecurityAssessment.go.html to see an example of how to use RefreshSecurityAssessment API.
func (client DataSafeClient) RefreshSecurityAssessment(ctx context.Context, request RefreshSecurityAssessmentRequest) (response RefreshSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.refreshSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshSecurityAssessmentResponse")
	}
	return
}

// refreshSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshSecurityAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshUserAssessment Refreshes the latest assessment and saves it for future reference. This operation runs with a userAssessmentId of type LATEST.
// Before you start, first call the ListUserAssessments operation with filter "type = latest" to get the user assessment ID for
// the target's latest assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshUserAssessment.go.html to see an example of how to use RefreshUserAssessment API.
func (client DataSafeClient) RefreshUserAssessment(ctx context.Context, request RefreshUserAssessmentRequest) (response RefreshUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.refreshUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshUserAssessmentResponse")
	}
	return
}

// refreshUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshUserAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetSecurityAssessmentBaseline Sets the saved security assessment as the baseline in the compartment where the the specified assessment resides. The security assessment needs to be of type 'SAVED'.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SetSecurityAssessmentBaseline.go.html to see an example of how to use SetSecurityAssessmentBaseline API.
func (client DataSafeClient) SetSecurityAssessmentBaseline(ctx context.Context, request SetSecurityAssessmentBaselineRequest) (response SetSecurityAssessmentBaselineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.setSecurityAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetSecurityAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetSecurityAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetSecurityAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetSecurityAssessmentBaselineResponse")
	}
	return
}

// setSecurityAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) setSecurityAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/setBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetSecurityAssessmentBaselineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetUserAssessmentBaseline Sets the saved user assessment as the baseline in the compartment where the specified assessment resides. The user assessment needs to be of type 'SAVED'.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SetUserAssessmentBaseline.go.html to see an example of how to use SetUserAssessmentBaseline API.
func (client DataSafeClient) SetUserAssessmentBaseline(ctx context.Context, request SetUserAssessmentBaselineRequest) (response SetUserAssessmentBaselineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.setUserAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetUserAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetUserAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetUserAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetUserAssessmentBaselineResponse")
	}
	return
}

// setUserAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) setUserAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/setBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetUserAssessmentBaselineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UnsetSecurityAssessmentBaseline Removes the baseline setting for the saved security assessment. The saved security assessment is no longer considered a baseline.
// Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UnsetSecurityAssessmentBaseline.go.html to see an example of how to use UnsetSecurityAssessmentBaseline API.
func (client DataSafeClient) UnsetSecurityAssessmentBaseline(ctx context.Context, request UnsetSecurityAssessmentBaselineRequest) (response UnsetSecurityAssessmentBaselineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.unsetSecurityAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnsetSecurityAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnsetSecurityAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnsetSecurityAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnsetSecurityAssessmentBaselineResponse")
	}
	return
}

// unsetSecurityAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) unsetSecurityAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/unsetBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnsetSecurityAssessmentBaselineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UnsetUserAssessmentBaseline Removes the baseline setting for the saved user assessment. The saved user assessment is no longer considered a baseline.
// Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UnsetUserAssessmentBaseline.go.html to see an example of how to use UnsetUserAssessmentBaseline API.
func (client DataSafeClient) UnsetUserAssessmentBaseline(ctx context.Context, request UnsetUserAssessmentBaselineRequest) (response UnsetUserAssessmentBaselineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.unsetUserAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnsetUserAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnsetUserAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnsetUserAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnsetUserAssessmentBaselineResponse")
	}
	return
}

// unsetUserAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) unsetUserAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/unsetBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnsetUserAssessmentBaselineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataSafePrivateEndpoint Updates one or more attributes of the specified Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateDataSafePrivateEndpoint.go.html to see an example of how to use UpdateDataSafePrivateEndpoint API.
func (client DataSafeClient) UpdateDataSafePrivateEndpoint(ctx context.Context, request UpdateDataSafePrivateEndpointRequest) (response UpdateDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataSafePrivateEndpointResponse")
	}
	return
}

// updateDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataSafePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOnPremConnector Updates one or more attributes of the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateOnPremConnector.go.html to see an example of how to use UpdateOnPremConnector API.
func (client DataSafeClient) UpdateOnPremConnector(ctx context.Context, request UpdateOnPremConnectorRequest) (response UpdateOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOnPremConnectorResponse")
	}
	return
}

// updateOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/onPremConnectors/{onPremConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOnPremConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOnPremConnectorWallet Updates the wallet for the specified on-premises connector to a new version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateOnPremConnectorWallet.go.html to see an example of how to use UpdateOnPremConnectorWallet API.
func (client DataSafeClient) UpdateOnPremConnectorWallet(ctx context.Context, request UpdateOnPremConnectorWalletRequest) (response UpdateOnPremConnectorWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateOnPremConnectorWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOnPremConnectorWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOnPremConnectorWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOnPremConnectorWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOnPremConnectorWalletResponse")
	}
	return
}

// updateOnPremConnectorWallet implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateOnPremConnectorWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/onPremConnectors/{onPremConnectorId}/wallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOnPremConnectorWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityAssessment Updates one or more attributes of the specified security assessment. This operation allows to update the security assessment displayName, description, or schedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSecurityAssessment.go.html to see an example of how to use UpdateSecurityAssessment API.
func (client DataSafeClient) UpdateSecurityAssessment(ctx context.Context, request UpdateSecurityAssessmentRequest) (response UpdateSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityAssessmentResponse")
	}
	return
}

// updateSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityAssessments/{securityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetDatabase Updates one or more attributes of the specified Data Safe target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateTargetDatabase.go.html to see an example of how to use UpdateTargetDatabase API.
func (client DataSafeClient) UpdateTargetDatabase(ctx context.Context, request UpdateTargetDatabaseRequest) (response UpdateTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetDatabaseResponse")
	}
	return
}

// updateTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targetDatabases/{targetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUserAssessment Updates one or more attributes of the specified user assessment. This operation allows to update the user assessment displayName, description, or schedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateUserAssessment.go.html to see an example of how to use UpdateUserAssessment API.
func (client DataSafeClient) UpdateUserAssessment(ctx context.Context, request UpdateUserAssessmentRequest) (response UpdateUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUserAssessmentResponse")
	}
	return
}

// updateUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/userAssessments/{userAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUserAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
