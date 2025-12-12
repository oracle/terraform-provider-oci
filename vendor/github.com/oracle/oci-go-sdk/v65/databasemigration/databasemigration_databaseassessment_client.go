// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DatabaseAssessmentClient a client for DatabaseAssessment
type DatabaseAssessmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseAssessmentClientWithConfigurationProvider Creates a new default DatabaseAssessment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseAssessmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseAssessmentClient, err error) {
	if enabled := common.CheckForEnabledServices("databasemigration"); !enabled {
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
	return newDatabaseAssessmentClientFromBaseClient(baseClient, provider)
}

// NewDatabaseAssessmentClientWithOboToken Creates a new default DatabaseAssessment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatabaseAssessmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseAssessmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseAssessmentClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseAssessmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseAssessmentClient, err error) {
	// DatabaseAssessment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DatabaseAssessment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatabaseAssessmentClient{BaseClient: baseClient}
	client.BasePath = "20230518"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseAssessmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemigration", "https://odms.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseAssessmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseAssessmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddAssessmentObjects Add excluded/included object to the list.
// A default retry strategy applies to this operation AddAssessmentObjects()
func (client DatabaseAssessmentClient) AddAssessmentObjects(ctx context.Context, request AddAssessmentObjectsRequest) (response AddAssessmentObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addAssessmentObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddAssessmentObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddAssessmentObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddAssessmentObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddAssessmentObjectsResponse")
	}
	return
}

// addAssessmentObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) addAssessmentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/actions/addAssessmentObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddAssessmentObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/AddAssessmentObjects"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "AddAssessmentObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAssessmentCompartment Used to change the Assessment compartment.
// A default retry strategy applies to this operation ChangeAssessmentCompartment()
func (client DatabaseAssessmentClient) ChangeAssessmentCompartment(ctx context.Context, request ChangeAssessmentCompartmentRequest) (response ChangeAssessmentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAssessmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAssessmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAssessmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAssessmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAssessmentCompartmentResponse")
	}
	return
}

// changeAssessmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) changeAssessmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAssessmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/ChangeAssessmentCompartment"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ChangeAssessmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CloneAssessment Clone a configuration from an existing Assessment.
// A default retry strategy applies to this operation CloneAssessment()
func (client DatabaseAssessmentClient) CloneAssessment(ctx context.Context, request CloneAssessmentRequest) (response CloneAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cloneAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneAssessmentResponse")
	}
	return
}

// cloneAssessment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) cloneAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/actions/clone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/CloneAssessment"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "CloneAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessment{})
	return response, err
}

// CreateAssessment Create an Assessment resource that contains all the details to perform the
// database assessment operation, such as source and destination database
// details, network throughput, accepted downtime etc.
// A default retry strategy applies to this operation CreateAssessment()
func (client DatabaseAssessmentClient) CreateAssessment(ctx context.Context, request CreateAssessmentRequest) (response CreateAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAssessmentResponse")
	}
	return
}

// createAssessment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) createAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/CreateAssessment"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "CreateAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessment{})
	return response, err
}

// DeleteAssessment Deletes the Assessment represented by the specified assessment ID.
// A default retry strategy applies to this operation DeleteAssessment()
func (client DatabaseAssessmentClient) DeleteAssessment(ctx context.Context, request DeleteAssessmentRequest) (response DeleteAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAssessmentResponse")
	}
	return
}

// deleteAssessment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) deleteAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/assessments/{assessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/DeleteAssessment"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "DeleteAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssessment Display Assessment details.
// A default retry strategy applies to this operation GetAssessment()
func (client DatabaseAssessmentClient) GetAssessment(ctx context.Context, request GetAssessmentRequest) (response GetAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssessmentResponse")
	}
	return
}

// getAssessment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) getAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/GetAssessment"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "GetAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessment{})
	return response, err
}

// GetAssessor Display Assessor details.
// A default retry strategy applies to this operation GetAssessor()
func (client DatabaseAssessmentClient) GetAssessor(ctx context.Context, request GetAssessorRequest) (response GetAssessorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssessor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssessorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssessorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssessorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssessorResponse")
	}
	return
}

// getAssessor implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) getAssessor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}/assessors/{assessorName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssessorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessor/GetAssessor"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "GetAssessor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssessorCheck Get Assessor Check details.
// A default retry strategy applies to this operation GetAssessorCheck()
func (client DatabaseAssessmentClient) GetAssessorCheck(ctx context.Context, request GetAssessorCheckRequest) (response GetAssessorCheckResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssessorCheck, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssessorCheckResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssessorCheckResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssessorCheckResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssessorCheckResponse")
	}
	return
}

// getAssessorCheck implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) getAssessorCheck(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}/assessors/{assessorName}/checks/{checkName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssessorCheckResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessorCheck/GetAssessorCheck"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "GetAssessorCheck", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScript Download DMS script.
// A default retry strategy applies to this operation GetScript()
func (client DatabaseAssessmentClient) GetScript(ctx context.Context, request GetScriptRequest) (response GetScriptResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScriptResponse")
	}
	return
}

// getScript implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) getScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scripts/{scriptId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "GetScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAffectedObjects Display Check Affected objects.
// A default retry strategy applies to this operation ListAffectedObjects()
func (client DatabaseAssessmentClient) ListAffectedObjects(ctx context.Context, request ListAffectedObjectsRequest) (response ListAffectedObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAffectedObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAffectedObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAffectedObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAffectedObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAffectedObjectsResponse")
	}
	return
}

// listAffectedObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) listAffectedObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}/assessors/{assessorName}/checks/{checkName}/affectedObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAffectedObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessorCheck/ListAffectedObjects"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ListAffectedObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessmentObjectTypes Display sample object types to exclude or include for an Assessment.
// A default retry strategy applies to this operation ListAssessmentObjectTypes()
func (client DatabaseAssessmentClient) ListAssessmentObjectTypes(ctx context.Context, request ListAssessmentObjectTypesRequest) (response ListAssessmentObjectTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssessmentObjectTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssessmentObjectTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssessmentObjectTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssessmentObjectTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssessmentObjectTypesResponse")
	}
	return
}

// listAssessmentObjectTypes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) listAssessmentObjectTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessmentObjectTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssessmentObjectTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessmentObjectTypeSummary/ListAssessmentObjectTypes"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ListAssessmentObjectTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessmentObjects Display excluded/included objects.
// A default retry strategy applies to this operation ListAssessmentObjects()
func (client DatabaseAssessmentClient) ListAssessmentObjects(ctx context.Context, request ListAssessmentObjectsRequest) (response ListAssessmentObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssessmentObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssessmentObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssessmentObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssessmentObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssessmentObjectsResponse")
	}
	return
}

// listAssessmentObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) listAssessmentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}/assessmentObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssessmentObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessmentObjectCollection/ListAssessmentObjects"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ListAssessmentObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessmentobjectcollection{})
	return response, err
}

// ListAssessments List all Assessments.
// A default retry strategy applies to this operation ListAssessments()
func (client DatabaseAssessmentClient) ListAssessments(ctx context.Context, request ListAssessmentsRequest) (response ListAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssessments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssessmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssessmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssessmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssessmentsResponse")
	}
	return
}

// listAssessments implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) listAssessments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssessmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessmentSummary/ListAssessments"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ListAssessments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessorChecks List Assessor Check Summaries.
// A default retry strategy applies to this operation ListAssessorChecks()
func (client DatabaseAssessmentClient) ListAssessorChecks(ctx context.Context, request ListAssessorChecksRequest) (response ListAssessorChecksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssessorChecks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssessorChecksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssessorChecksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssessorChecksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssessorChecksResponse")
	}
	return
}

// listAssessorChecks implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) listAssessorChecks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}/assessors/{assessorName}/checks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssessorChecksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessorCheckSummary/ListAssessorChecks"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ListAssessorChecks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessors List all Assessors.
// A default retry strategy applies to this operation ListAssessors()
func (client DatabaseAssessmentClient) ListAssessors(ctx context.Context, request ListAssessorsRequest) (response ListAssessorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssessors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssessorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssessorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssessorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssessorsResponse")
	}
	return
}

// listAssessors implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) listAssessors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assessments/{assessmentId}/assessors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssessorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AssessorSummary/ListAssessors"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "ListAssessors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PerformAssessorAction Assessor Action.
// A default retry strategy applies to this operation PerformAssessorAction()
func (client DatabaseAssessmentClient) PerformAssessorAction(ctx context.Context, request PerformAssessorActionRequest) (response PerformAssessorActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.performAssessorAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PerformAssessorActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PerformAssessorActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PerformAssessorActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PerformAssessorActionResponse")
	}
	return
}

// performAssessorAction implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) performAssessorAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/assessors/{assessorName}/assessorActions/{assessorAction}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PerformAssessorActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/PerformAssessorAction"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "PerformAssessorAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PerformAssessorActionDownloadSql Download SQL script Assessor Action.
// A default retry strategy applies to this operation PerformAssessorActionDownloadSql()
func (client DatabaseAssessmentClient) PerformAssessorActionDownloadSql(ctx context.Context, request PerformAssessorActionDownloadSqlRequest) (response PerformAssessorActionDownloadSqlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.performAssessorActionDownloadSql, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PerformAssessorActionDownloadSqlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PerformAssessorActionDownloadSqlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PerformAssessorActionDownloadSqlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PerformAssessorActionDownloadSqlResponse")
	}
	return
}

// performAssessorActionDownloadSql implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) performAssessorActionDownloadSql(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/assessors/{assessorName}/actions/download_sql", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PerformAssessorActionDownloadSqlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/PerformAssessorActionDownloadSql"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "PerformAssessorActionDownloadSql", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PerformAssessorCheckAction Assessor Check Action.
// A default retry strategy applies to this operation PerformAssessorCheckAction()
func (client DatabaseAssessmentClient) PerformAssessorCheckAction(ctx context.Context, request PerformAssessorCheckActionRequest) (response PerformAssessorCheckActionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.performAssessorCheckAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PerformAssessorCheckActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PerformAssessorCheckActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PerformAssessorCheckActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PerformAssessorCheckActionResponse")
	}
	return
}

// performAssessorCheckAction implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) performAssessorCheckAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/assessors/{assessorName}/checks/{checkName}/actions/{assessorCheckAction}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PerformAssessorCheckActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/PerformAssessorCheckAction"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "PerformAssessorCheckAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAssessmentObjects Remove excluded/included objects.
// A default retry strategy applies to this operation RemoveAssessmentObjects()
func (client DatabaseAssessmentClient) RemoveAssessmentObjects(ctx context.Context, request RemoveAssessmentObjectsRequest) (response RemoveAssessmentObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeAssessmentObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveAssessmentObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveAssessmentObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveAssessmentObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveAssessmentObjectsResponse")
	}
	return
}

// removeAssessmentObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) removeAssessmentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assessments/{assessmentId}/actions/removeAssessmentObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveAssessmentObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/RemoveAssessmentObjects"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "RemoveAssessmentObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAssessment Update Assessment resource details.
// A default retry strategy applies to this operation UpdateAssessment()
func (client DatabaseAssessmentClient) UpdateAssessment(ctx context.Context, request UpdateAssessmentRequest) (response UpdateAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAssessmentResponse")
	}
	return
}

// updateAssessment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) updateAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/assessments/{assessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAssessmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Assessment/UpdateAssessment"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "UpdateAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCheckActionUpdateObject Update the advisor report object list.
// A default retry strategy applies to this operation UpdateCheckActionUpdateObject()
func (client DatabaseAssessmentClient) UpdateCheckActionUpdateObject(ctx context.Context, request UpdateCheckActionUpdateObjectRequest) (response UpdateCheckActionUpdateObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCheckActionUpdateObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCheckActionUpdateObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCheckActionUpdateObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCheckActionUpdateObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCheckActionUpdateObjectResponse")
	}
	return
}

// updateCheckActionUpdateObject implements the OCIOperation interface (enables retrying operations)
func (client DatabaseAssessmentClient) updateCheckActionUpdateObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/assessments/{assessmentId}/assessors/{assessorName}/checks/{checkName}/actions/updateObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCheckActionUpdateObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AdvisorReportCheckCollection/UpdateCheckActionUpdateObject"
		err = common.PostProcessServiceError(err, "DatabaseAssessment", "UpdateCheckActionUpdateObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
