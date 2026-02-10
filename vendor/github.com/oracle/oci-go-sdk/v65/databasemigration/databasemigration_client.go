// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseMigrationClient a client for DatabaseMigration
type DatabaseMigrationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseMigrationClientWithConfigurationProvider Creates a new default DatabaseMigration client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseMigrationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseMigrationClient, err error) {
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
	return newDatabaseMigrationClientFromBaseClient(baseClient, provider)
}

// NewDatabaseMigrationClientWithOboToken Creates a new default DatabaseMigration client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatabaseMigrationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseMigrationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseMigrationClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseMigrationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseMigrationClient, err error) {
	// DatabaseMigration service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DatabaseMigration"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatabaseMigrationClient{BaseClient: baseClient}
	client.BasePath = "20230518"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseMigrationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemigration", "https://odms.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseMigrationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseMigrationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortJob Aborts a Migration Job (either Evaluation or Migration).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/AbortJob.go.html to see an example of how to use AbortJob API.
// A default retry strategy applies to this operation AbortJob()
func (client DatabaseMigrationClient) AbortJob(ctx context.Context, request AbortJobRequest) (response AbortJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.abortJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AbortJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AbortJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AbortJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AbortJobResponse")
	}
	return
}

// abortJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) abortJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/actions/abort", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AbortJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/AbortJob"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "AbortJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddAssessmentObjects Add excluded/included object to the list.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/AddAssessmentObjects.go.html to see an example of how to use AddAssessmentObjects API.
// A default retry strategy applies to this operation AddAssessmentObjects()
func (client DatabaseMigrationClient) AddAssessmentObjects(ctx context.Context, request AddAssessmentObjectsRequest) (response AddAssessmentObjectsResponse, err error) {
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
func (client DatabaseMigrationClient) addAssessmentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "AddAssessmentObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddMigrationObjects Add excluded/included object to the list.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/AddMigrationObjects.go.html to see an example of how to use AddMigrationObjects API.
// A default retry strategy applies to this operation AddMigrationObjects()
func (client DatabaseMigrationClient) AddMigrationObjects(ctx context.Context, request AddMigrationObjectsRequest) (response AddMigrationObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addMigrationObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMigrationObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMigrationObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMigrationObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMigrationObjectsResponse")
	}
	return
}

// addMigrationObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) addMigrationObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/addMigrationObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMigrationObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/AddMigrationObjects"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "AddMigrationObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAssessmentCompartment Used to change the Assessment compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ChangeAssessmentCompartment.go.html to see an example of how to use ChangeAssessmentCompartment API.
// A default retry strategy applies to this operation ChangeAssessmentCompartment()
func (client DatabaseMigrationClient) ChangeAssessmentCompartment(ctx context.Context, request ChangeAssessmentCompartmentRequest) (response ChangeAssessmentCompartmentResponse, err error) {
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
func (client DatabaseMigrationClient) changeAssessmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ChangeAssessmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeConnectionCompartment Used to change the Database Connection compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ChangeConnectionCompartment.go.html to see an example of how to use ChangeConnectionCompartment API.
// A default retry strategy applies to this operation ChangeConnectionCompartment()
func (client DatabaseMigrationClient) ChangeConnectionCompartment(ctx context.Context, request ChangeConnectionCompartmentRequest) (response ChangeConnectionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeConnectionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeConnectionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeConnectionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeConnectionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeConnectionCompartmentResponse")
	}
	return
}

// changeConnectionCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) changeConnectionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeConnectionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Connection/ChangeConnectionCompartment"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ChangeConnectionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMigrationCompartment Used to change the Migration compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ChangeMigrationCompartment.go.html to see an example of how to use ChangeMigrationCompartment API.
// A default retry strategy applies to this operation ChangeMigrationCompartment()
func (client DatabaseMigrationClient) ChangeMigrationCompartment(ctx context.Context, request ChangeMigrationCompartmentRequest) (response ChangeMigrationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMigrationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMigrationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMigrationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMigrationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMigrationCompartmentResponse")
	}
	return
}

// changeMigrationCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) changeMigrationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMigrationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/ChangeMigrationCompartment"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ChangeMigrationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CloneAssessment Clone a configuration from an existing Assessment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CloneAssessment.go.html to see an example of how to use CloneAssessment API.
// A default retry strategy applies to this operation CloneAssessment()
func (client DatabaseMigrationClient) CloneAssessment(ctx context.Context, request CloneAssessmentRequest) (response CloneAssessmentResponse, err error) {
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
func (client DatabaseMigrationClient) cloneAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CloneAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessment{})
	return response, err
}

// CloneMigration Clone a configuration from an existing Migration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CloneMigration.go.html to see an example of how to use CloneMigration API.
// A default retry strategy applies to this operation CloneMigration()
func (client DatabaseMigrationClient) CloneMigration(ctx context.Context, request CloneMigrationRequest) (response CloneMigrationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cloneMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneMigrationResponse")
	}
	return
}

// cloneMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) cloneMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/clone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/CloneMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CloneMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &migration{})
	return response, err
}

// CollectTraces Collects the DB trace and alert logs.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CollectTraces.go.html to see an example of how to use CollectTraces API.
// A default retry strategy applies to this operation CollectTraces()
func (client DatabaseMigrationClient) CollectTraces(ctx context.Context, request CollectTracesRequest) (response CollectTracesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.collectTraces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CollectTracesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CollectTracesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CollectTracesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CollectTracesResponse")
	}
	return
}

// collectTraces implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) collectTraces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/actions/collectTraces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CollectTracesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/CollectTraces"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CollectTraces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConnectionDiagnostics Perform connection test for a database connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ConnectionDiagnostics.go.html to see an example of how to use ConnectionDiagnostics API.
// A default retry strategy applies to this operation ConnectionDiagnostics()
func (client DatabaseMigrationClient) ConnectionDiagnostics(ctx context.Context, request ConnectionDiagnosticsRequest) (response ConnectionDiagnosticsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.connectionDiagnostics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConnectionDiagnosticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConnectionDiagnosticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConnectionDiagnosticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConnectionDiagnosticsResponse")
	}
	return
}

// connectionDiagnostics implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) connectionDiagnostics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/diagnostics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConnectionDiagnosticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Connection/ConnectionDiagnostics"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ConnectionDiagnostics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAssessment Create an Assessment resource that contains all the details to perform the
// database assessment operation, such as source and destination database
// details, network throughput, accepted downtime etc.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CreateAssessment.go.html to see an example of how to use CreateAssessment API.
// A default retry strategy applies to this operation CreateAssessment()
func (client DatabaseMigrationClient) CreateAssessment(ctx context.Context, request CreateAssessmentRequest) (response CreateAssessmentResponse, err error) {
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
func (client DatabaseMigrationClient) createAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CreateAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessment{})
	return response, err
}

// CreateConnection Create a Database Connection resource that contains the details to connect to either a Source or Target Database
// in the migration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CreateConnection.go.html to see an example of how to use CreateConnection API.
// A default retry strategy applies to this operation CreateConnection()
func (client DatabaseMigrationClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
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
func (client DatabaseMigrationClient) createConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CreateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// CreateMigration Create a Migration resource that contains all the details to perform the
// database migration operation, such as source and destination database
// details, credentials, etc.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CreateMigration.go.html to see an example of how to use CreateMigration API.
// A default retry strategy applies to this operation CreateMigration()
func (client DatabaseMigrationClient) CreateMigration(ctx context.Context, request CreateMigrationRequest) (response CreateMigrationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMigrationResponse")
	}
	return
}

// createMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) createMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/CreateMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CreateMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &migration{})
	return response, err
}

// CreateParameterFileVersion Creates a new version of the current parameter file contents to the specified value.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/CreateParameterFileVersion.go.html to see an example of how to use CreateParameterFileVersion API.
// A default retry strategy applies to this operation CreateParameterFileVersion()
func (client DatabaseMigrationClient) CreateParameterFileVersion(ctx context.Context, request CreateParameterFileVersionRequest) (response CreateParameterFileVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createParameterFileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateParameterFileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateParameterFileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateParameterFileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateParameterFileVersionResponse")
	}
	return
}

// createParameterFileVersion implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) createParameterFileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/parameterFileVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateParameterFileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/CreateParameterFileVersion"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "CreateParameterFileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAssessment Deletes the Assessment represented by the specified assessment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/DeleteAssessment.go.html to see an example of how to use DeleteAssessment API.
// A default retry strategy applies to this operation DeleteAssessment()
func (client DatabaseMigrationClient) DeleteAssessment(ctx context.Context, request DeleteAssessmentRequest) (response DeleteAssessmentResponse, err error) {
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
func (client DatabaseMigrationClient) deleteAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "DeleteAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnection Deletes the Database Connection represented by the specified connection ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/DeleteConnection.go.html to see an example of how to use DeleteConnection API.
// A default retry strategy applies to this operation DeleteConnection()
func (client DatabaseMigrationClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) deleteConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Connection/DeleteConnection"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "DeleteConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJob Deletes the migration job represented by the given job ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/DeleteJob.go.html to see an example of how to use DeleteJob API.
// A default retry strategy applies to this operation DeleteJob()
func (client DatabaseMigrationClient) DeleteJob(ctx context.Context, request DeleteJobRequest) (response DeleteJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) deleteJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/DeleteJob"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "DeleteJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMigration Deletes the Migration represented by the specified migration ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/DeleteMigration.go.html to see an example of how to use DeleteMigration API.
// A default retry strategy applies to this operation DeleteMigration()
func (client DatabaseMigrationClient) DeleteMigration(ctx context.Context, request DeleteMigrationRequest) (response DeleteMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMigrationResponse")
	}
	return
}

// deleteMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) deleteMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/migrations/{migrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/DeleteMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "DeleteMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteParameterFileVersion Deletes the given parameter file version
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/DeleteParameterFileVersion.go.html to see an example of how to use DeleteParameterFileVersion API.
// A default retry strategy applies to this operation DeleteParameterFileVersion()
func (client DatabaseMigrationClient) DeleteParameterFileVersion(ctx context.Context, request DeleteParameterFileVersionRequest) (response DeleteParameterFileVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteParameterFileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteParameterFileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteParameterFileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteParameterFileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteParameterFileVersionResponse")
	}
	return
}

// deleteParameterFileVersion implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) deleteParameterFileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/jobs/{jobId}/parameterFileVersions/{parameterFileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteParameterFileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/DeleteParameterFileVersion"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "DeleteParameterFileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EvaluateMigration Start Validate Migration job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/EvaluateMigration.go.html to see an example of how to use EvaluateMigration API.
// A default retry strategy applies to this operation EvaluateMigration()
func (client DatabaseMigrationClient) EvaluateMigration(ctx context.Context, request EvaluateMigrationRequest) (response EvaluateMigrationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.evaluateMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EvaluateMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EvaluateMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EvaluateMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EvaluateMigrationResponse")
	}
	return
}

// evaluateMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) evaluateMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EvaluateMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/EvaluateMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "EvaluateMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAdvisorReport Get the Pre-Migration Advisor report details
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetAdvisorReport.go.html to see an example of how to use GetAdvisorReport API.
// A default retry strategy applies to this operation GetAdvisorReport()
func (client DatabaseMigrationClient) GetAdvisorReport(ctx context.Context, request GetAdvisorReportRequest) (response GetAdvisorReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdvisorReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdvisorReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdvisorReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdvisorReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdvisorReportResponse")
	}
	return
}

// getAdvisorReport implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) getAdvisorReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/advisorReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdvisorReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/GetAdvisorReport"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetAdvisorReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssessment Display Assessment details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetAssessment.go.html to see an example of how to use GetAssessment API.
// A default retry strategy applies to this operation GetAssessment()
func (client DatabaseMigrationClient) GetAssessment(ctx context.Context, request GetAssessmentRequest) (response GetAssessmentResponse, err error) {
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
func (client DatabaseMigrationClient) getAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessment{})
	return response, err
}

// GetAssessor Display Assessor details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetAssessor.go.html to see an example of how to use GetAssessor API.
// A default retry strategy applies to this operation GetAssessor()
func (client DatabaseMigrationClient) GetAssessor(ctx context.Context, request GetAssessorRequest) (response GetAssessorResponse, err error) {
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
func (client DatabaseMigrationClient) getAssessor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetAssessor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssessorCheck Get Assessor Check details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetAssessorCheck.go.html to see an example of how to use GetAssessorCheck API.
// A default retry strategy applies to this operation GetAssessorCheck()
func (client DatabaseMigrationClient) GetAssessorCheck(ctx context.Context, request GetAssessorCheckRequest) (response GetAssessorCheckResponse, err error) {
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
func (client DatabaseMigrationClient) getAssessorCheck(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetAssessorCheck", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConnection Display Database Connection details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetConnection.go.html to see an example of how to use GetConnection API.
// A default retry strategy applies to this operation GetConnection()
func (client DatabaseMigrationClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) getConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Connection/GetConnection"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// GetJob Get a migration job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetJob.go.html to see an example of how to use GetJob API.
// A default retry strategy applies to this operation GetJob()
func (client DatabaseMigrationClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) getJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/GetJob"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobOutputContent Get the migration Job Output content as a String.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetJobOutputContent.go.html to see an example of how to use GetJobOutputContent API.
// A default retry strategy applies to this operation GetJobOutputContent()
func (client DatabaseMigrationClient) GetJobOutputContent(ctx context.Context, request GetJobOutputContentRequest) (response GetJobOutputContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobOutputContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobOutputContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobOutputContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobOutputContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobOutputContentResponse")
	}
	return
}

// getJobOutputContent implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) getJobOutputContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/output/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobOutputContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/GetJobOutputContent"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetJobOutputContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMigration Display Migration details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetMigration.go.html to see an example of how to use GetMigration API.
// A default retry strategy applies to this operation GetMigration()
func (client DatabaseMigrationClient) GetMigration(ctx context.Context, request GetMigrationRequest) (response GetMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMigrationResponse")
	}
	return
}

// getMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) getMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations/{migrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/GetMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &migration{})
	return response, err
}

// GetParameterFileVersion Obtain the parameter file version contents for the specified parameter file name and the associated job. This operation will
// be allowed only if the job is certain acceptable lifecycle states.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetParameterFileVersion.go.html to see an example of how to use GetParameterFileVersion API.
// A default retry strategy applies to this operation GetParameterFileVersion()
func (client DatabaseMigrationClient) GetParameterFileVersion(ctx context.Context, request GetParameterFileVersionRequest) (response GetParameterFileVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getParameterFileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetParameterFileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetParameterFileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetParameterFileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetParameterFileVersionResponse")
	}
	return
}

// getParameterFileVersion implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) getParameterFileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/parameterFileVersions/{parameterFileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetParameterFileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/GetParameterFileVersion"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetParameterFileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScript Download DMS script.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetScript.go.html to see an example of how to use GetScript API.
// A default retry strategy applies to this operation GetScript()
func (client DatabaseMigrationClient) GetScript(ctx context.Context, request GetScriptRequest) (response GetScriptResponse, err error) {
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
func (client DatabaseMigrationClient) getScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DatabaseMigrationClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DatabaseMigrationClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdvisorReportCheckObjects Get the Pre-Migration extended Advisor report object list.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAdvisorReportCheckObjects.go.html to see an example of how to use ListAdvisorReportCheckObjects API.
// A default retry strategy applies to this operation ListAdvisorReportCheckObjects()
func (client DatabaseMigrationClient) ListAdvisorReportCheckObjects(ctx context.Context, request ListAdvisorReportCheckObjectsRequest) (response ListAdvisorReportCheckObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdvisorReportCheckObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdvisorReportCheckObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdvisorReportCheckObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdvisorReportCheckObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdvisorReportCheckObjectsResponse")
	}
	return
}

// listAdvisorReportCheckObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listAdvisorReportCheckObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/advisorReportChecks/{advisorReportCheckId}/objects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdvisorReportCheckObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/ListAdvisorReportCheckObjects"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAdvisorReportCheckObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdvisorReportChecks List of Pre-Migration checks from the advisor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAdvisorReportChecks.go.html to see an example of how to use ListAdvisorReportChecks API.
// A default retry strategy applies to this operation ListAdvisorReportChecks()
func (client DatabaseMigrationClient) ListAdvisorReportChecks(ctx context.Context, request ListAdvisorReportChecksRequest) (response ListAdvisorReportChecksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdvisorReportChecks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdvisorReportChecksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdvisorReportChecksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdvisorReportChecksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdvisorReportChecksResponse")
	}
	return
}

// listAdvisorReportChecks implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listAdvisorReportChecks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/advisorReportChecks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdvisorReportChecksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/ListAdvisorReportChecks"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAdvisorReportChecks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAffectedObjects Display Check Affected objects.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAffectedObjects.go.html to see an example of how to use ListAffectedObjects API.
// A default retry strategy applies to this operation ListAffectedObjects()
func (client DatabaseMigrationClient) ListAffectedObjects(ctx context.Context, request ListAffectedObjectsRequest) (response ListAffectedObjectsResponse, err error) {
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
func (client DatabaseMigrationClient) listAffectedObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAffectedObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessmentObjectTypes Display sample object types to exclude or include for an Assessment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessmentObjectTypes.go.html to see an example of how to use ListAssessmentObjectTypes API.
// A default retry strategy applies to this operation ListAssessmentObjectTypes()
func (client DatabaseMigrationClient) ListAssessmentObjectTypes(ctx context.Context, request ListAssessmentObjectTypesRequest) (response ListAssessmentObjectTypesResponse, err error) {
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
func (client DatabaseMigrationClient) listAssessmentObjectTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAssessmentObjectTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessmentObjects Display excluded/included objects.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessmentObjects.go.html to see an example of how to use ListAssessmentObjects API.
// A default retry strategy applies to this operation ListAssessmentObjects()
func (client DatabaseMigrationClient) ListAssessmentObjects(ctx context.Context, request ListAssessmentObjectsRequest) (response ListAssessmentObjectsResponse, err error) {
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
func (client DatabaseMigrationClient) listAssessmentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAssessmentObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assessmentobjectcollection{})
	return response, err
}

// ListAssessments List all Assessments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessments.go.html to see an example of how to use ListAssessments API.
// A default retry strategy applies to this operation ListAssessments()
func (client DatabaseMigrationClient) ListAssessments(ctx context.Context, request ListAssessmentsRequest) (response ListAssessmentsResponse, err error) {
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
func (client DatabaseMigrationClient) listAssessments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAssessments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessorChecks List Assessor Check Summaries.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessorChecks.go.html to see an example of how to use ListAssessorChecks API.
// A default retry strategy applies to this operation ListAssessorChecks()
func (client DatabaseMigrationClient) ListAssessorChecks(ctx context.Context, request ListAssessorChecksRequest) (response ListAssessorChecksResponse, err error) {
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
func (client DatabaseMigrationClient) listAssessorChecks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAssessorChecks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssessors List all Assessors.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessors.go.html to see an example of how to use ListAssessors API.
// A default retry strategy applies to this operation ListAssessors()
func (client DatabaseMigrationClient) ListAssessors(ctx context.Context, request ListAssessorsRequest) (response ListAssessorsResponse, err error) {
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
func (client DatabaseMigrationClient) listAssessors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListAssessors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnections List all Database Connections.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListConnections.go.html to see an example of how to use ListConnections API.
// A default retry strategy applies to this operation ListConnections()
func (client DatabaseMigrationClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) listConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/ConnectionSummary/ListConnections"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseConnectionType List supported Database Types, Sub-types and Versions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListDatabaseConnectionType.go.html to see an example of how to use ListDatabaseConnectionType API.
// A default retry strategy applies to this operation ListDatabaseConnectionType()
func (client DatabaseMigrationClient) ListDatabaseConnectionType(ctx context.Context, request ListDatabaseConnectionTypeRequest) (response ListDatabaseConnectionTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseConnectionType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseConnectionTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseConnectionTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseConnectionTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseConnectionTypeResponse")
	}
	return
}

// listDatabaseConnectionType implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listDatabaseConnectionType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections/databaseconnectiontype", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseConnectionTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/DatabaseConnectionTypeSummary/ListDatabaseConnectionType"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListDatabaseConnectionType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExcludedObjects List the excluded database objects.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListExcludedObjects.go.html to see an example of how to use ListExcludedObjects API.
// A default retry strategy applies to this operation ListExcludedObjects()
func (client DatabaseMigrationClient) ListExcludedObjects(ctx context.Context, request ListExcludedObjectsRequest) (response ListExcludedObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExcludedObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExcludedObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExcludedObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExcludedObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExcludedObjectsResponse")
	}
	return
}

// listExcludedObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listExcludedObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/excludedObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExcludedObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/ExcludedObjectSummary/ListExcludedObjects"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListExcludedObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobOutputs List the Job Outputs
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListJobOutputs.go.html to see an example of how to use ListJobOutputs API.
// A default retry strategy applies to this operation ListJobOutputs()
func (client DatabaseMigrationClient) ListJobOutputs(ctx context.Context, request ListJobOutputsRequest) (response ListJobOutputsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobOutputs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobOutputsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobOutputsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobOutputsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobOutputsResponse")
	}
	return
}

// listJobOutputs implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listJobOutputs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/output", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobOutputsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/JobOutputSummary/ListJobOutputs"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListJobOutputs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobs List all the names of the Migration jobs associated to the specified
// migration site.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListJobs.go.html to see an example of how to use ListJobs API.
// A default retry strategy applies to this operation ListJobs()
func (client DatabaseMigrationClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) listJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/JobSummary/ListJobs"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMigrationObjectTypes Display sample object types to exclude or include for a Migration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrationObjectTypes.go.html to see an example of how to use ListMigrationObjectTypes API.
// A default retry strategy applies to this operation ListMigrationObjectTypes()
func (client DatabaseMigrationClient) ListMigrationObjectTypes(ctx context.Context, request ListMigrationObjectTypesRequest) (response ListMigrationObjectTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMigrationObjectTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMigrationObjectTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMigrationObjectTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMigrationObjectTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMigrationObjectTypesResponse")
	}
	return
}

// listMigrationObjectTypes implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listMigrationObjectTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrationObjectTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMigrationObjectTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/MigrationObjectTypeSummary/ListMigrationObjectTypes"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListMigrationObjectTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMigrationObjects Display excluded/included objects.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrationObjects.go.html to see an example of how to use ListMigrationObjects API.
// A default retry strategy applies to this operation ListMigrationObjects()
func (client DatabaseMigrationClient) ListMigrationObjects(ctx context.Context, request ListMigrationObjectsRequest) (response ListMigrationObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMigrationObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMigrationObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMigrationObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMigrationObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMigrationObjectsResponse")
	}
	return
}

// listMigrationObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listMigrationObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations/{migrationId}/migrationObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMigrationObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/MigrationObjectCollection/ListMigrationObjects"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListMigrationObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &migrationobjectcollection{})
	return response, err
}

// ListMigrationParameters List of parameters that can be used to customize migrations.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrationParameters.go.html to see an example of how to use ListMigrationParameters API.
// A default retry strategy applies to this operation ListMigrationParameters()
func (client DatabaseMigrationClient) ListMigrationParameters(ctx context.Context, request ListMigrationParametersRequest) (response ListMigrationParametersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMigrationParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMigrationParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMigrationParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMigrationParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMigrationParametersResponse")
	}
	return
}

// listMigrationParameters implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listMigrationParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrationParameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMigrationParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/MigrationParameterSummary/ListMigrationParameters"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListMigrationParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMigrations List all Migrations.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrations.go.html to see an example of how to use ListMigrations API.
// A default retry strategy applies to this operation ListMigrations()
func (client DatabaseMigrationClient) ListMigrations(ctx context.Context, request ListMigrationsRequest) (response ListMigrationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMigrations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMigrationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMigrationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMigrationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMigrationsResponse")
	}
	return
}

// listMigrations implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listMigrations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/migrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMigrationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/MigrationSummary/ListMigrations"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListMigrations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListParameterFileVersions Return a list of the parameter file metadata of the migration execution of the specified job.  This will
// only be acceptable if the job is in particular state. It will be accessible if the job is in
// the FAILED, PAUSED or SUSPENDED state.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListParameterFileVersions.go.html to see an example of how to use ListParameterFileVersions API.
// A default retry strategy applies to this operation ListParameterFileVersions()
func (client DatabaseMigrationClient) ListParameterFileVersions(ctx context.Context, request ListParameterFileVersionsRequest) (response ListParameterFileVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParameterFileVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParameterFileVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParameterFileVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParameterFileVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParameterFileVersionsResponse")
	}
	return
}

// listParameterFileVersions implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) listParameterFileVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/parameterFileVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListParameterFileVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/ListParameterFileVersions"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListParameterFileVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Gets the errors for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DatabaseMigrationClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DatabaseMigrationClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Gets the logs for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DatabaseMigrationClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DatabaseMigrationClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment or for a specified resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DatabaseMigrationClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DatabaseMigrationClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MakeCurrentParameterFileVersion Make current the given parameter file version
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/MakeCurrentParameterFileVersion.go.html to see an example of how to use MakeCurrentParameterFileVersion API.
// A default retry strategy applies to this operation MakeCurrentParameterFileVersion()
func (client DatabaseMigrationClient) MakeCurrentParameterFileVersion(ctx context.Context, request MakeCurrentParameterFileVersionRequest) (response MakeCurrentParameterFileVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.makeCurrentParameterFileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MakeCurrentParameterFileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MakeCurrentParameterFileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MakeCurrentParameterFileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MakeCurrentParameterFileVersionResponse")
	}
	return
}

// makeCurrentParameterFileVersion implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) makeCurrentParameterFileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/parameterFileVersions/{parameterFileName}/actions/makeCurrent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MakeCurrentParameterFileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/MakeCurrentParameterFileVersion"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "MakeCurrentParameterFileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PerformAssessorAction Assessor Action.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/PerformAssessorAction.go.html to see an example of how to use PerformAssessorAction API.
// A default retry strategy applies to this operation PerformAssessorAction()
func (client DatabaseMigrationClient) PerformAssessorAction(ctx context.Context, request PerformAssessorActionRequest) (response PerformAssessorActionResponse, err error) {
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
func (client DatabaseMigrationClient) performAssessorAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "PerformAssessorAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PerformAssessorActionDownloadSql Download SQL script Assessor Action.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/PerformAssessorActionDownloadSql.go.html to see an example of how to use PerformAssessorActionDownloadSql API.
// A default retry strategy applies to this operation PerformAssessorActionDownloadSql()
func (client DatabaseMigrationClient) PerformAssessorActionDownloadSql(ctx context.Context, request PerformAssessorActionDownloadSqlRequest) (response PerformAssessorActionDownloadSqlResponse, err error) {
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
func (client DatabaseMigrationClient) performAssessorActionDownloadSql(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "PerformAssessorActionDownloadSql", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PerformAssessorCheckAction Assessor Check Action.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/PerformAssessorCheckAction.go.html to see an example of how to use PerformAssessorCheckAction API.
// A default retry strategy applies to this operation PerformAssessorCheckAction()
func (client DatabaseMigrationClient) PerformAssessorCheckAction(ctx context.Context, request PerformAssessorCheckActionRequest) (response PerformAssessorCheckActionResponse, err error) {
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
func (client DatabaseMigrationClient) performAssessorCheckAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "PerformAssessorCheckAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAssessmentObjects Remove excluded/included objects.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/RemoveAssessmentObjects.go.html to see an example of how to use RemoveAssessmentObjects API.
// A default retry strategy applies to this operation RemoveAssessmentObjects()
func (client DatabaseMigrationClient) RemoveAssessmentObjects(ctx context.Context, request RemoveAssessmentObjectsRequest) (response RemoveAssessmentObjectsResponse, err error) {
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
func (client DatabaseMigrationClient) removeAssessmentObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "RemoveAssessmentObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveMigrationObjects Remove excluded/included objects.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/RemoveMigrationObjects.go.html to see an example of how to use RemoveMigrationObjects API.
// A default retry strategy applies to this operation RemoveMigrationObjects()
func (client DatabaseMigrationClient) RemoveMigrationObjects(ctx context.Context, request RemoveMigrationObjectsRequest) (response RemoveMigrationObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeMigrationObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveMigrationObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveMigrationObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveMigrationObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveMigrationObjectsResponse")
	}
	return
}

// removeMigrationObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) removeMigrationObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/removeMigrationObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveMigrationObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/RemoveMigrationObjects"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "RemoveMigrationObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResumeJob Resume a migration Job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ResumeJob.go.html to see an example of how to use ResumeJob API.
// A default retry strategy applies to this operation ResumeJob()
func (client DatabaseMigrationClient) ResumeJob(ctx context.Context, request ResumeJobRequest) (response ResumeJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resumeJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeJobResponse")
	}
	return
}

// resumeJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) resumeJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/ResumeJob"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "ResumeJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveSupportedPhases Display Migration Phases for a specified migration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/RetrieveSupportedPhases.go.html to see an example of how to use RetrieveSupportedPhases API.
// A default retry strategy applies to this operation RetrieveSupportedPhases()
func (client DatabaseMigrationClient) RetrieveSupportedPhases(ctx context.Context, request RetrieveSupportedPhasesRequest) (response RetrieveSupportedPhasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveSupportedPhases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveSupportedPhasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveSupportedPhasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveSupportedPhasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveSupportedPhasesResponse")
	}
	return
}

// retrieveSupportedPhases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) retrieveSupportedPhases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/getSupportedPhases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveSupportedPhasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/RetrieveSupportedPhases"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "RetrieveSupportedPhases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartMigration Start Migration job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/StartMigration.go.html to see an example of how to use StartMigration API.
// A default retry strategy applies to this operation StartMigration()
func (client DatabaseMigrationClient) StartMigration(ctx context.Context, request StartMigrationRequest) (response StartMigrationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartMigrationResponse")
	}
	return
}

// startMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) startMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/migrations/{migrationId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/StartMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "StartMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SuspendJob Place the currently executing migration Job in a Suspended State.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/SuspendJob.go.html to see an example of how to use SuspendJob API.
// A default retry strategy applies to this operation SuspendJob()
func (client DatabaseMigrationClient) SuspendJob(ctx context.Context, request SuspendJobRequest) (response SuspendJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.suspendJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuspendJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuspendJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuspendJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuspendJobResponse")
	}
	return
}

// suspendJob implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) suspendJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/actions/suspend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SuspendJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/SuspendJob"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "SuspendJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAdvisorReportCheck Update the premigration extended Advisor report check.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateAdvisorReportCheck.go.html to see an example of how to use UpdateAdvisorReportCheck API.
// A default retry strategy applies to this operation UpdateAdvisorReportCheck()
func (client DatabaseMigrationClient) UpdateAdvisorReportCheck(ctx context.Context, request UpdateAdvisorReportCheckRequest) (response UpdateAdvisorReportCheckResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAdvisorReportCheck, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAdvisorReportCheckResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAdvisorReportCheckResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAdvisorReportCheckResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAdvisorReportCheckResponse")
	}
	return
}

// updateAdvisorReportCheck implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) updateAdvisorReportCheck(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/jobs/{jobId}/advisorReportChecks/{advisorReportCheckId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAdvisorReportCheckResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateAdvisorReportCheck", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAdvisorReportCheckObjects Update the Pre-Migration extended Advisor report object list.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateAdvisorReportCheckObjects.go.html to see an example of how to use UpdateAdvisorReportCheckObjects API.
// A default retry strategy applies to this operation UpdateAdvisorReportCheckObjects()
func (client DatabaseMigrationClient) UpdateAdvisorReportCheckObjects(ctx context.Context, request UpdateAdvisorReportCheckObjectsRequest) (response UpdateAdvisorReportCheckObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAdvisorReportCheckObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAdvisorReportCheckObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAdvisorReportCheckObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAdvisorReportCheckObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAdvisorReportCheckObjectsResponse")
	}
	return
}

// updateAdvisorReportCheckObjects implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) updateAdvisorReportCheckObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/advisorReportChecks/{advisorReportCheckId}/actions/updateObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAdvisorReportCheckObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/AdvisorReportCheckCollection/UpdateAdvisorReportCheckObjects"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateAdvisorReportCheckObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAssessment Update Assessment resource details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateAssessment.go.html to see an example of how to use UpdateAssessment API.
// A default retry strategy applies to this operation UpdateAssessment()
func (client DatabaseMigrationClient) UpdateAssessment(ctx context.Context, request UpdateAssessmentRequest) (response UpdateAssessmentResponse, err error) {
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
func (client DatabaseMigrationClient) updateAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCheckActionUpdateObject Update the advisor report object list.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateCheckActionUpdateObject.go.html to see an example of how to use UpdateCheckActionUpdateObject API.
// A default retry strategy applies to this operation UpdateCheckActionUpdateObject()
func (client DatabaseMigrationClient) UpdateCheckActionUpdateObject(ctx context.Context, request UpdateCheckActionUpdateObjectRequest) (response UpdateCheckActionUpdateObjectResponse, err error) {
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
func (client DatabaseMigrationClient) updateCheckActionUpdateObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateCheckActionUpdateObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConnection Update Database Connection resource details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateConnection.go.html to see an example of how to use UpdateConnection API.
// A default retry strategy applies to this operation UpdateConnection()
func (client DatabaseMigrationClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) updateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Connection/UpdateConnection"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJob Update Migration Job resource details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateJob.go.html to see an example of how to use UpdateJob API.
// A default retry strategy applies to this operation UpdateJob()
func (client DatabaseMigrationClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DatabaseMigrationClient) updateJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Job/UpdateJob"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMigration Update Migration resource details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateMigration.go.html to see an example of how to use UpdateMigration API.
// A default retry strategy applies to this operation UpdateMigration()
func (client DatabaseMigrationClient) UpdateMigration(ctx context.Context, request UpdateMigrationRequest) (response UpdateMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMigrationResponse")
	}
	return
}

// updateMigration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseMigrationClient) updateMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/migrations/{migrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-migration/20230518/Migration/UpdateMigration"
		err = common.PostProcessServiceError(err, "DatabaseMigration", "UpdateMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
