// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DataScienceClient a client for DataScience
type DataScienceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataScienceClientWithConfigurationProvider Creates a new default DataScience client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataScienceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataScienceClient, err error) {
	if enabled := common.CheckForEnabledServices("datascience"); !enabled {
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
	return newDataScienceClientFromBaseClient(baseClient, provider)
}

// NewDataScienceClientWithOboToken Creates a new default DataScience client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDataScienceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataScienceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataScienceClientFromBaseClient(baseClient, configProvider)
}

func newDataScienceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataScienceClient, err error) {
	// DataScience service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DataScience"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataScienceClient{BaseClient: baseClient}
	client.BasePath = "20190101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataScienceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datascience", "https://datascience.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataScienceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataScienceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateModel Activates the model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ActivateModel.go.html to see an example of how to use ActivateModel API.
// A default retry strategy applies to this operation ActivateModel()
func (client DataScienceClient) ActivateModel(ctx context.Context, request ActivateModelRequest) (response ActivateModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.activateModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateModelResponse")
	}
	return
}

// activateModel implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) activateModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/ActivateModel"
		err = common.PostProcessServiceError(err, "DataScience", "ActivateModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ActivateModelDeployment Activates the model deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ActivateModelDeployment.go.html to see an example of how to use ActivateModelDeployment API.
func (client DataScienceClient) ActivateModelDeployment(ctx context.Context, request ActivateModelDeploymentRequest) (response ActivateModelDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.activateModelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateModelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateModelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateModelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateModelDeploymentResponse")
	}
	return
}

// activateModelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) activateModelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/modelDeployments/{modelDeploymentId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateModelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/ActivateModelDeployment"
		err = common.PostProcessServiceError(err, "DataScience", "ActivateModelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ActivateNotebookSession Activates the notebook session.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ActivateNotebookSession.go.html to see an example of how to use ActivateNotebookSession API.
func (client DataScienceClient) ActivateNotebookSession(ctx context.Context, request ActivateNotebookSessionRequest) (response ActivateNotebookSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.activateNotebookSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateNotebookSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateNotebookSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateNotebookSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateNotebookSessionResponse")
	}
	return
}

// activateNotebookSession implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) activateNotebookSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/notebookSessions/{notebookSessionId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateNotebookSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/ActivateNotebookSession"
		err = common.PostProcessServiceError(err, "DataScience", "ActivateNotebookSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ActivateSchedule Activate schedule.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ActivateSchedule.go.html to see an example of how to use ActivateSchedule API.
// A default retry strategy applies to this operation ActivateSchedule()
func (client DataScienceClient) ActivateSchedule(ctx context.Context, request ActivateScheduleRequest) (response ActivateScheduleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateScheduleResponse")
	}
	return
}

// activateSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) activateSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedules/{scheduleId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/ActivateSchedule"
		err = common.PostProcessServiceError(err, "DataScience", "ActivateSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelJobRun Cancels an IN_PROGRESS job run.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CancelJobRun.go.html to see an example of how to use CancelJobRun API.
// A default retry strategy applies to this operation CancelJobRun()
func (client DataScienceClient) CancelJobRun(ctx context.Context, request CancelJobRunRequest) (response CancelJobRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelJobRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelJobRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelJobRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelJobRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelJobRunResponse")
	}
	return
}

// cancelJobRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) cancelJobRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobRuns/{jobRunId}/actions/cancelJobRun", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelJobRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRun/CancelJobRun"
		err = common.PostProcessServiceError(err, "DataScience", "CancelJobRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelPipelineRun Cancel a PipelineRun.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CancelPipelineRun.go.html to see an example of how to use CancelPipelineRun API.
// A default retry strategy applies to this operation CancelPipelineRun()
func (client DataScienceClient) CancelPipelineRun(ctx context.Context, request CancelPipelineRunRequest) (response CancelPipelineRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelPipelineRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelPipelineRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelPipelineRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelPipelineRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelPipelineRunResponse")
	}
	return
}

// cancelPipelineRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) cancelPipelineRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelineRuns/{pipelineRunId}/actions/cancelPipelineRun", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelPipelineRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/CancelPipelineRun"
		err = common.PostProcessServiceError(err, "DataScience", "CancelPipelineRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancels a work request that has not started.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
func (client DataScienceClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DataScienceClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "DataScience", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDataSciencePrivateEndpointCompartment Moves a private endpoint into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeDataSciencePrivateEndpointCompartment.go.html to see an example of how to use ChangeDataSciencePrivateEndpointCompartment API.
// A default retry strategy applies to this operation ChangeDataSciencePrivateEndpointCompartment()
func (client DataScienceClient) ChangeDataSciencePrivateEndpointCompartment(ctx context.Context, request ChangeDataSciencePrivateEndpointCompartmentRequest) (response ChangeDataSciencePrivateEndpointCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDataSciencePrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDataSciencePrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDataSciencePrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDataSciencePrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDataSciencePrivateEndpointCompartmentResponse")
	}
	return
}

// changeDataSciencePrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeDataSciencePrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSciencePrivateEndpoints/{dataSciencePrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDataSciencePrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/DataSciencePrivateEndpoint/ChangeDataSciencePrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeDataSciencePrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeJobCompartment Changes a job's compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeJobCompartment.go.html to see an example of how to use ChangeJobCompartment API.
func (client DataScienceClient) ChangeJobCompartment(ctx context.Context, request ChangeJobCompartmentRequest) (response ChangeJobCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
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
func (client DataScienceClient) changeJobCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/ChangeJobCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeJobCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeJobRunCompartment Changes a job run's compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeJobRunCompartment.go.html to see an example of how to use ChangeJobRunCompartment API.
func (client DataScienceClient) ChangeJobRunCompartment(ctx context.Context, request ChangeJobRunCompartmentRequest) (response ChangeJobRunCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeJobRunCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeJobRunCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeJobRunCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeJobRunCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeJobRunCompartmentResponse")
	}
	return
}

// changeJobRunCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeJobRunCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobRuns/{jobRunId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeJobRunCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRun/ChangeJobRunCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeJobRunCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMlApplicationCompartment Moves a MlApplication resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeMlApplicationCompartment.go.html to see an example of how to use ChangeMlApplicationCompartment API.
// A default retry strategy applies to this operation ChangeMlApplicationCompartment()
func (client DataScienceClient) ChangeMlApplicationCompartment(ctx context.Context, request ChangeMlApplicationCompartmentRequest) (response ChangeMlApplicationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMlApplicationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMlApplicationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMlApplicationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMlApplicationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMlApplicationCompartmentResponse")
	}
	return
}

// changeMlApplicationCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeMlApplicationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplications/{mlApplicationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMlApplicationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplication/ChangeMlApplicationCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeMlApplicationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMlApplicationImplementationCompartment Moves a MlApplicationImplementation resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeMlApplicationImplementationCompartment.go.html to see an example of how to use ChangeMlApplicationImplementationCompartment API.
// A default retry strategy applies to this operation ChangeMlApplicationImplementationCompartment()
func (client DataScienceClient) ChangeMlApplicationImplementationCompartment(ctx context.Context, request ChangeMlApplicationImplementationCompartmentRequest) (response ChangeMlApplicationImplementationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMlApplicationImplementationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMlApplicationImplementationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMlApplicationImplementationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMlApplicationImplementationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMlApplicationImplementationCompartmentResponse")
	}
	return
}

// changeMlApplicationImplementationCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeMlApplicationImplementationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationImplementations/{mlApplicationImplementationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMlApplicationImplementationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/ChangeMlApplicationImplementationCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeMlApplicationImplementationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMlApplicationInstanceCompartment Moves a MlApplicationInstance resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeMlApplicationInstanceCompartment.go.html to see an example of how to use ChangeMlApplicationInstanceCompartment API.
// A default retry strategy applies to this operation ChangeMlApplicationInstanceCompartment()
func (client DataScienceClient) ChangeMlApplicationInstanceCompartment(ctx context.Context, request ChangeMlApplicationInstanceCompartmentRequest) (response ChangeMlApplicationInstanceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMlApplicationInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMlApplicationInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMlApplicationInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMlApplicationInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMlApplicationInstanceCompartmentResponse")
	}
	return
}

// changeMlApplicationInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeMlApplicationInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstances/{mlApplicationInstanceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMlApplicationInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/ChangeMlApplicationInstanceCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeMlApplicationInstanceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMlApplicationInstanceViewCompartment Moves a MlApplicationInstanceView resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeMlApplicationInstanceViewCompartment.go.html to see an example of how to use ChangeMlApplicationInstanceViewCompartment API.
// A default retry strategy applies to this operation ChangeMlApplicationInstanceViewCompartment()
func (client DataScienceClient) ChangeMlApplicationInstanceViewCompartment(ctx context.Context, request ChangeMlApplicationInstanceViewCompartmentRequest) (response ChangeMlApplicationInstanceViewCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMlApplicationInstanceViewCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMlApplicationInstanceViewCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMlApplicationInstanceViewCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMlApplicationInstanceViewCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMlApplicationInstanceViewCompartmentResponse")
	}
	return
}

// changeMlApplicationInstanceViewCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeMlApplicationInstanceViewCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMlApplicationInstanceViewCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/ChangeMlApplicationInstanceViewCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeMlApplicationInstanceViewCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeModelCompartment Moves a model resource into a different compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeModelCompartment.go.html to see an example of how to use ChangeModelCompartment API.
func (client DataScienceClient) ChangeModelCompartment(ctx context.Context, request ChangeModelCompartmentRequest) (response ChangeModelCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeModelCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeModelCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeModelCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeModelCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeModelCompartmentResponse")
	}
	return
}

// changeModelCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeModelCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeModelCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/ChangeModelCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeModelCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeModelDeploymentCompartment Moves a model deployment into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeModelDeploymentCompartment.go.html to see an example of how to use ChangeModelDeploymentCompartment API.
func (client DataScienceClient) ChangeModelDeploymentCompartment(ctx context.Context, request ChangeModelDeploymentCompartmentRequest) (response ChangeModelDeploymentCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeModelDeploymentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeModelDeploymentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeModelDeploymentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeModelDeploymentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeModelDeploymentCompartmentResponse")
	}
	return
}

// changeModelDeploymentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeModelDeploymentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/modelDeployments/{modelDeploymentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeModelDeploymentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/ChangeModelDeploymentCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeModelDeploymentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeModelVersionSetCompartment Moves a modelVersionSet resource into a different compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeModelVersionSetCompartment.go.html to see an example of how to use ChangeModelVersionSetCompartment API.
func (client DataScienceClient) ChangeModelVersionSetCompartment(ctx context.Context, request ChangeModelVersionSetCompartmentRequest) (response ChangeModelVersionSetCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeModelVersionSetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeModelVersionSetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeModelVersionSetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeModelVersionSetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeModelVersionSetCompartmentResponse")
	}
	return
}

// changeModelVersionSetCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeModelVersionSetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/modelVersionSets/{modelVersionSetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeModelVersionSetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelVersionSet/ChangeModelVersionSetCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeModelVersionSetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeNotebookSessionCompartment Moves a notebook session resource into a different compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeNotebookSessionCompartment.go.html to see an example of how to use ChangeNotebookSessionCompartment API.
func (client DataScienceClient) ChangeNotebookSessionCompartment(ctx context.Context, request ChangeNotebookSessionCompartmentRequest) (response ChangeNotebookSessionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeNotebookSessionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeNotebookSessionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeNotebookSessionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeNotebookSessionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeNotebookSessionCompartmentResponse")
	}
	return
}

// changeNotebookSessionCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeNotebookSessionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/notebookSessions/{notebookSessionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeNotebookSessionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/ChangeNotebookSessionCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeNotebookSessionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePipelineCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangePipelineCompartment.go.html to see an example of how to use ChangePipelineCompartment API.
func (client DataScienceClient) ChangePipelineCompartment(ctx context.Context, request ChangePipelineCompartmentRequest) (response ChangePipelineCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changePipelineCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePipelineCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePipelineCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePipelineCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePipelineCompartmentResponse")
	}
	return
}

// changePipelineCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changePipelineCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePipelineCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/ChangePipelineCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangePipelineCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePipelineRunCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangePipelineRunCompartment.go.html to see an example of how to use ChangePipelineRunCompartment API.
func (client DataScienceClient) ChangePipelineRunCompartment(ctx context.Context, request ChangePipelineRunCompartmentRequest) (response ChangePipelineRunCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changePipelineRunCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePipelineRunCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePipelineRunCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePipelineRunCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePipelineRunCompartmentResponse")
	}
	return
}

// changePipelineRunCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changePipelineRunCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelineRuns/{pipelineRunId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePipelineRunCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/ChangePipelineRunCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangePipelineRunCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeProjectCompartment Moves a project resource into a different compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeProjectCompartment.go.html to see an example of how to use ChangeProjectCompartment API.
func (client DataScienceClient) ChangeProjectCompartment(ctx context.Context, request ChangeProjectCompartmentRequest) (response ChangeProjectCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeProjectCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeProjectCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeProjectCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeProjectCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeProjectCompartmentResponse")
	}
	return
}

// changeProjectCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeProjectCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects/{projectId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeProjectCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Project/ChangeProjectCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeProjectCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeScheduleCompartment Moves a Schedule resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ChangeScheduleCompartment.go.html to see an example of how to use ChangeScheduleCompartment API.
// A default retry strategy applies to this operation ChangeScheduleCompartment()
func (client DataScienceClient) ChangeScheduleCompartment(ctx context.Context, request ChangeScheduleCompartmentRequest) (response ChangeScheduleCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeScheduleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeScheduleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeScheduleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeScheduleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeScheduleCompartmentResponse")
	}
	return
}

// changeScheduleCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) changeScheduleCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedules/{scheduleId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeScheduleCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/ChangeScheduleCompartment"
		err = common.PostProcessServiceError(err, "DataScience", "ChangeScheduleCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataSciencePrivateEndpoint Creates a Data Science private endpoint to be used by a Data Science resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateDataSciencePrivateEndpoint.go.html to see an example of how to use CreateDataSciencePrivateEndpoint API.
// A default retry strategy applies to this operation CreateDataSciencePrivateEndpoint()
func (client DataScienceClient) CreateDataSciencePrivateEndpoint(ctx context.Context, request CreateDataSciencePrivateEndpointRequest) (response CreateDataSciencePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataSciencePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataSciencePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataSciencePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataSciencePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataSciencePrivateEndpointResponse")
	}
	return
}

// createDataSciencePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createDataSciencePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSciencePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataSciencePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataScience", "CreateDataSciencePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJob Creates a job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateJob.go.html to see an example of how to use CreateJob API.
// A default retry strategy applies to this operation CreateJob()
func (client DataScienceClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
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
func (client DataScienceClient) createJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/CreateJob"
		err = common.PostProcessServiceError(err, "DataScience", "CreateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJobArtifact Uploads a job artifact.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateJobArtifact.go.html to see an example of how to use CreateJobArtifact API.
func (client DataScienceClient) CreateJobArtifact(ctx context.Context, request CreateJobArtifactRequest) (response CreateJobArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJobArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobArtifactResponse")
	}
	return
}

// createJobArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createJobArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs/{jobId}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJobArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/CreateJobArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "CreateJobArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJobRun Creates a job run.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateJobRun.go.html to see an example of how to use CreateJobRun API.
// A default retry strategy applies to this operation CreateJobRun()
func (client DataScienceClient) CreateJobRun(ctx context.Context, request CreateJobRunRequest) (response CreateJobRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createJobRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobRunResponse")
	}
	return
}

// createJobRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createJobRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJobRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRun/CreateJobRun"
		err = common.PostProcessServiceError(err, "DataScience", "CreateJobRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMlApplication Creates a new MlApplication.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateMlApplication.go.html to see an example of how to use CreateMlApplication API.
// A default retry strategy applies to this operation CreateMlApplication()
func (client DataScienceClient) CreateMlApplication(ctx context.Context, request CreateMlApplicationRequest) (response CreateMlApplicationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMlApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMlApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMlApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMlApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMlApplicationResponse")
	}
	return
}

// createMlApplication implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createMlApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMlApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplication/CreateMlApplication"
		err = common.PostProcessServiceError(err, "DataScience", "CreateMlApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMlApplicationImplementation Creates a new MlApplicationImplementation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateMlApplicationImplementation.go.html to see an example of how to use CreateMlApplicationImplementation API.
// A default retry strategy applies to this operation CreateMlApplicationImplementation()
func (client DataScienceClient) CreateMlApplicationImplementation(ctx context.Context, request CreateMlApplicationImplementationRequest) (response CreateMlApplicationImplementationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMlApplicationImplementation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMlApplicationImplementationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMlApplicationImplementationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMlApplicationImplementationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMlApplicationImplementationResponse")
	}
	return
}

// createMlApplicationImplementation implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createMlApplicationImplementation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationImplementations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMlApplicationImplementationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/CreateMlApplicationImplementation"
		err = common.PostProcessServiceError(err, "DataScience", "CreateMlApplicationImplementation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMlApplicationInstance Creates a new MlApplicationInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateMlApplicationInstance.go.html to see an example of how to use CreateMlApplicationInstance API.
// A default retry strategy applies to this operation CreateMlApplicationInstance()
func (client DataScienceClient) CreateMlApplicationInstance(ctx context.Context, request CreateMlApplicationInstanceRequest) (response CreateMlApplicationInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMlApplicationInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMlApplicationInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMlApplicationInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMlApplicationInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMlApplicationInstanceResponse")
	}
	return
}

// createMlApplicationInstance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createMlApplicationInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMlApplicationInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/CreateMlApplicationInstance"
		err = common.PostProcessServiceError(err, "DataScience", "CreateMlApplicationInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModel Creates a new model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModel.go.html to see an example of how to use CreateModel API.
// A default retry strategy applies to this operation CreateModel()
func (client DataScienceClient) CreateModel(ctx context.Context, request CreateModelRequest) (response CreateModelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelResponse")
	}
	return
}

// createModel implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/CreateModel"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModelArtifact Creates model artifact for specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModelArtifact.go.html to see an example of how to use CreateModelArtifact API.
// A default retry strategy applies to this operation CreateModelArtifact()
func (client DataScienceClient) CreateModelArtifact(ctx context.Context, request CreateModelArtifactRequest) (response CreateModelArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createModelArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelArtifactResponse")
	}
	return
}

// createModelArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModelArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/CreateModelArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModelArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModelCustomMetadatumArtifact Creates model custom metadata artifact for specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModelCustomMetadatumArtifact.go.html to see an example of how to use CreateModelCustomMetadatumArtifact API.
func (client DataScienceClient) CreateModelCustomMetadatumArtifact(ctx context.Context, request CreateModelCustomMetadatumArtifactRequest) (response CreateModelCustomMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createModelCustomMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelCustomMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelCustomMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelCustomMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelCustomMetadatumArtifactResponse")
	}
	return
}

// createModelCustomMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModelCustomMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/customMetadata/{metadatumKeyName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelCustomMetadatumArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/CreateModelCustomMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModelCustomMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModelDefinedMetadatumArtifact Creates model defined metadata artifact for specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModelDefinedMetadatumArtifact.go.html to see an example of how to use CreateModelDefinedMetadatumArtifact API.
func (client DataScienceClient) CreateModelDefinedMetadatumArtifact(ctx context.Context, request CreateModelDefinedMetadatumArtifactRequest) (response CreateModelDefinedMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createModelDefinedMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelDefinedMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelDefinedMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelDefinedMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelDefinedMetadatumArtifactResponse")
	}
	return
}

// createModelDefinedMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModelDefinedMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/definedMetadata/{metadatumKeyName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelDefinedMetadatumArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/CreateModelDefinedMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModelDefinedMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModelDeployment Creates a new model deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModelDeployment.go.html to see an example of how to use CreateModelDeployment API.
// A default retry strategy applies to this operation CreateModelDeployment()
func (client DataScienceClient) CreateModelDeployment(ctx context.Context, request CreateModelDeploymentRequest) (response CreateModelDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createModelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelDeploymentResponse")
	}
	return
}

// createModelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/modelDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/CreateModelDeployment"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModelProvenance Creates provenance information for the specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModelProvenance.go.html to see an example of how to use CreateModelProvenance API.
// A default retry strategy applies to this operation CreateModelProvenance()
func (client DataScienceClient) CreateModelProvenance(ctx context.Context, request CreateModelProvenanceRequest) (response CreateModelProvenanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createModelProvenance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelProvenanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelProvenanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelProvenanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelProvenanceResponse")
	}
	return
}

// createModelProvenance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModelProvenance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/provenance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelProvenanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/CreateModelProvenance"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModelProvenance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateModelVersionSet Creates a new modelVersionSet.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateModelVersionSet.go.html to see an example of how to use CreateModelVersionSet API.
// A default retry strategy applies to this operation CreateModelVersionSet()
func (client DataScienceClient) CreateModelVersionSet(ctx context.Context, request CreateModelVersionSetRequest) (response CreateModelVersionSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createModelVersionSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateModelVersionSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateModelVersionSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateModelVersionSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateModelVersionSetResponse")
	}
	return
}

// createModelVersionSet implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createModelVersionSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/modelVersionSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateModelVersionSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelVersionSet/CreateModelVersionSet"
		err = common.PostProcessServiceError(err, "DataScience", "CreateModelVersionSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNotebookSession Creates a new notebook session.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateNotebookSession.go.html to see an example of how to use CreateNotebookSession API.
// A default retry strategy applies to this operation CreateNotebookSession()
func (client DataScienceClient) CreateNotebookSession(ctx context.Context, request CreateNotebookSessionRequest) (response CreateNotebookSessionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNotebookSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNotebookSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNotebookSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNotebookSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNotebookSessionResponse")
	}
	return
}

// createNotebookSession implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createNotebookSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/notebookSessions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNotebookSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/CreateNotebookSession"
		err = common.PostProcessServiceError(err, "DataScience", "CreateNotebookSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePipeline Creates a new Pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreatePipeline.go.html to see an example of how to use CreatePipeline API.
// A default retry strategy applies to this operation CreatePipeline()
func (client DataScienceClient) CreatePipeline(ctx context.Context, request CreatePipelineRequest) (response CreatePipelineResponse, err error) {
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
func (client DataScienceClient) createPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/CreatePipeline"
		err = common.PostProcessServiceError(err, "DataScience", "CreatePipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePipelineRun Creates a new PipelineRun.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreatePipelineRun.go.html to see an example of how to use CreatePipelineRun API.
// A default retry strategy applies to this operation CreatePipelineRun()
func (client DataScienceClient) CreatePipelineRun(ctx context.Context, request CreatePipelineRunRequest) (response CreatePipelineRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPipelineRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePipelineRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePipelineRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePipelineRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePipelineRunResponse")
	}
	return
}

// createPipelineRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createPipelineRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelineRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePipelineRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/CreatePipelineRun"
		err = common.PostProcessServiceError(err, "DataScience", "CreatePipelineRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProject Creates a new project.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateProject.go.html to see an example of how to use CreateProject API.
// A default retry strategy applies to this operation CreateProject()
func (client DataScienceClient) CreateProject(ctx context.Context, request CreateProjectRequest) (response CreateProjectResponse, err error) {
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
func (client DataScienceClient) createProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Project/CreateProject"
		err = common.PostProcessServiceError(err, "DataScience", "CreateProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSchedule Creates a new Schedule.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateSchedule.go.html to see an example of how to use CreateSchedule API.
// A default retry strategy applies to this operation CreateSchedule()
func (client DataScienceClient) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (response CreateScheduleResponse, err error) {
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
func (client DataScienceClient) createSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/CreateSchedule"
		err = common.PostProcessServiceError(err, "DataScience", "CreateSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateStepArtifact Upload the artifact for a step in the pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/CreateStepArtifact.go.html to see an example of how to use CreateStepArtifact API.
func (client DataScienceClient) CreateStepArtifact(ctx context.Context, request CreateStepArtifactRequest) (response CreateStepArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createStepArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStepArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStepArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStepArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStepArtifactResponse")
	}
	return
}

// createStepArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) createStepArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/steps/{stepName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStepArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/CreateStepArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "CreateStepArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateModel Deactivates the model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeactivateModel.go.html to see an example of how to use DeactivateModel API.
// A default retry strategy applies to this operation DeactivateModel()
func (client DataScienceClient) DeactivateModel(ctx context.Context, request DeactivateModelRequest) (response DeactivateModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deactivateModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateModelResponse")
	}
	return
}

// deactivateModel implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deactivateModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/DeactivateModel"
		err = common.PostProcessServiceError(err, "DataScience", "DeactivateModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateModelDeployment Deactivates the model deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeactivateModelDeployment.go.html to see an example of how to use DeactivateModelDeployment API.
func (client DataScienceClient) DeactivateModelDeployment(ctx context.Context, request DeactivateModelDeploymentRequest) (response DeactivateModelDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deactivateModelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateModelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateModelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateModelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateModelDeploymentResponse")
	}
	return
}

// deactivateModelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deactivateModelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/modelDeployments/{modelDeploymentId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateModelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/DeactivateModelDeployment"
		err = common.PostProcessServiceError(err, "DataScience", "DeactivateModelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateNotebookSession Deactivates the notebook session.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeactivateNotebookSession.go.html to see an example of how to use DeactivateNotebookSession API.
func (client DataScienceClient) DeactivateNotebookSession(ctx context.Context, request DeactivateNotebookSessionRequest) (response DeactivateNotebookSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deactivateNotebookSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateNotebookSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateNotebookSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateNotebookSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateNotebookSessionResponse")
	}
	return
}

// deactivateNotebookSession implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deactivateNotebookSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/notebookSessions/{notebookSessionId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateNotebookSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/DeactivateNotebookSession"
		err = common.PostProcessServiceError(err, "DataScience", "DeactivateNotebookSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateSchedule Deactivate schedule.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeactivateSchedule.go.html to see an example of how to use DeactivateSchedule API.
// A default retry strategy applies to this operation DeactivateSchedule()
func (client DataScienceClient) DeactivateSchedule(ctx context.Context, request DeactivateScheduleRequest) (response DeactivateScheduleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deactivateSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateScheduleResponse")
	}
	return
}

// deactivateSchedule implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deactivateSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedules/{scheduleId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/DeactivateSchedule"
		err = common.PostProcessServiceError(err, "DataScience", "DeactivateSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataSciencePrivateEndpoint Deletes a private endpoint using `privateEndpointId`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteDataSciencePrivateEndpoint.go.html to see an example of how to use DeleteDataSciencePrivateEndpoint API.
func (client DataScienceClient) DeleteDataSciencePrivateEndpoint(ctx context.Context, request DeleteDataSciencePrivateEndpointRequest) (response DeleteDataSciencePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataSciencePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataSciencePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataSciencePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataSciencePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataSciencePrivateEndpointResponse")
	}
	return
}

// deleteDataSciencePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteDataSciencePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dataSciencePrivateEndpoints/{dataSciencePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataSciencePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/DataSciencePrivateEndpoint/DeleteDataSciencePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteDataSciencePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJob Deletes a job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteJob.go.html to see an example of how to use DeleteJob API.
// A default retry strategy applies to this operation DeleteJob()
func (client DataScienceClient) DeleteJob(ctx context.Context, request DeleteJobRequest) (response DeleteJobResponse, err error) {
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
func (client DataScienceClient) deleteJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/DeleteJob"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJobRun Deletes a job run.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteJobRun.go.html to see an example of how to use DeleteJobRun API.
// A default retry strategy applies to this operation DeleteJobRun()
func (client DataScienceClient) DeleteJobRun(ctx context.Context, request DeleteJobRunRequest) (response DeleteJobRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJobRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJobRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJobRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobRunResponse")
	}
	return
}

// deleteJobRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteJobRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/jobRuns/{jobRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJobRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRun/DeleteJobRun"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteJobRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMlApplication Deletes a MlApplication resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteMlApplication.go.html to see an example of how to use DeleteMlApplication API.
// A default retry strategy applies to this operation DeleteMlApplication()
func (client DataScienceClient) DeleteMlApplication(ctx context.Context, request DeleteMlApplicationRequest) (response DeleteMlApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMlApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMlApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMlApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMlApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMlApplicationResponse")
	}
	return
}

// deleteMlApplication implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteMlApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mlApplications/{mlApplicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMlApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplication/DeleteMlApplication"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteMlApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMlApplicationImplementation Deletes a MlApplicationImplementation resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteMlApplicationImplementation.go.html to see an example of how to use DeleteMlApplicationImplementation API.
// A default retry strategy applies to this operation DeleteMlApplicationImplementation()
func (client DataScienceClient) DeleteMlApplicationImplementation(ctx context.Context, request DeleteMlApplicationImplementationRequest) (response DeleteMlApplicationImplementationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMlApplicationImplementation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMlApplicationImplementationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMlApplicationImplementationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMlApplicationImplementationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMlApplicationImplementationResponse")
	}
	return
}

// deleteMlApplicationImplementation implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteMlApplicationImplementation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mlApplicationImplementations/{mlApplicationImplementationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMlApplicationImplementationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/DeleteMlApplicationImplementation"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteMlApplicationImplementation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMlApplicationInstance Deletes a MlApplicationInstance resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteMlApplicationInstance.go.html to see an example of how to use DeleteMlApplicationInstance API.
// A default retry strategy applies to this operation DeleteMlApplicationInstance()
func (client DataScienceClient) DeleteMlApplicationInstance(ctx context.Context, request DeleteMlApplicationInstanceRequest) (response DeleteMlApplicationInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMlApplicationInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMlApplicationInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMlApplicationInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMlApplicationInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMlApplicationInstanceResponse")
	}
	return
}

// deleteMlApplicationInstance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteMlApplicationInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mlApplicationInstances/{mlApplicationInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMlApplicationInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/DeleteMlApplicationInstance"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteMlApplicationInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteModel Deletes the specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteModel.go.html to see an example of how to use DeleteModel API.
// A default retry strategy applies to this operation DeleteModel()
func (client DataScienceClient) DeleteModel(ctx context.Context, request DeleteModelRequest) (response DeleteModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteModelResponse")
	}
	return
}

// deleteModel implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/models/{modelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/DeleteModel"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteModelCustomMetadatumArtifact Deletes model custom metadata artifact for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteModelCustomMetadatumArtifact.go.html to see an example of how to use DeleteModelCustomMetadatumArtifact API.
// A default retry strategy applies to this operation DeleteModelCustomMetadatumArtifact()
func (client DataScienceClient) DeleteModelCustomMetadatumArtifact(ctx context.Context, request DeleteModelCustomMetadatumArtifactRequest) (response DeleteModelCustomMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteModelCustomMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteModelCustomMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteModelCustomMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteModelCustomMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteModelCustomMetadatumArtifactResponse")
	}
	return
}

// deleteModelCustomMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteModelCustomMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/models/{modelId}/customMetadata/{metadatumKeyName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteModelCustomMetadatumArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/DeleteModelCustomMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteModelCustomMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteModelDefinedMetadatumArtifact Deletes model defined metadata artifact for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteModelDefinedMetadatumArtifact.go.html to see an example of how to use DeleteModelDefinedMetadatumArtifact API.
// A default retry strategy applies to this operation DeleteModelDefinedMetadatumArtifact()
func (client DataScienceClient) DeleteModelDefinedMetadatumArtifact(ctx context.Context, request DeleteModelDefinedMetadatumArtifactRequest) (response DeleteModelDefinedMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteModelDefinedMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteModelDefinedMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteModelDefinedMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteModelDefinedMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteModelDefinedMetadatumArtifactResponse")
	}
	return
}

// deleteModelDefinedMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteModelDefinedMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/models/{modelId}/definedMetadata/{metadatumKeyName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteModelDefinedMetadatumArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/DeleteModelDefinedMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteModelDefinedMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteModelDeployment Deletes the specified model deployment. Any unsaved work in this model deployment is lost.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteModelDeployment.go.html to see an example of how to use DeleteModelDeployment API.
// A default retry strategy applies to this operation DeleteModelDeployment()
func (client DataScienceClient) DeleteModelDeployment(ctx context.Context, request DeleteModelDeploymentRequest) (response DeleteModelDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteModelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteModelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteModelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteModelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteModelDeploymentResponse")
	}
	return
}

// deleteModelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteModelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/modelDeployments/{modelDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteModelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/DeleteModelDeployment"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteModelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteModelVersionSet Deletes the specified modelVersionSet.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteModelVersionSet.go.html to see an example of how to use DeleteModelVersionSet API.
// A default retry strategy applies to this operation DeleteModelVersionSet()
func (client DataScienceClient) DeleteModelVersionSet(ctx context.Context, request DeleteModelVersionSetRequest) (response DeleteModelVersionSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteModelVersionSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteModelVersionSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteModelVersionSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteModelVersionSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteModelVersionSetResponse")
	}
	return
}

// deleteModelVersionSet implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteModelVersionSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/modelVersionSets/{modelVersionSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteModelVersionSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelVersionSet/DeleteModelVersionSet"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteModelVersionSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNotebookSession Deletes the specified notebook session. Any unsaved work in this notebook session are lost.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteNotebookSession.go.html to see an example of how to use DeleteNotebookSession API.
// A default retry strategy applies to this operation DeleteNotebookSession()
func (client DataScienceClient) DeleteNotebookSession(ctx context.Context, request DeleteNotebookSessionRequest) (response DeleteNotebookSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNotebookSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNotebookSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNotebookSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNotebookSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNotebookSessionResponse")
	}
	return
}

// deleteNotebookSession implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deleteNotebookSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/notebookSessions/{notebookSessionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNotebookSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/DeleteNotebookSession"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteNotebookSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePipeline Deletes a Pipeline resource by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeletePipeline.go.html to see an example of how to use DeletePipeline API.
func (client DataScienceClient) DeletePipeline(ctx context.Context, request DeletePipelineRequest) (response DeletePipelineResponse, err error) {
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
func (client DataScienceClient) deletePipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/pipelines/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/DeletePipeline"
		err = common.PostProcessServiceError(err, "DataScience", "DeletePipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePipelineRun Deletes a PipelineRun resource by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeletePipelineRun.go.html to see an example of how to use DeletePipelineRun API.
func (client DataScienceClient) DeletePipelineRun(ctx context.Context, request DeletePipelineRunRequest) (response DeletePipelineRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePipelineRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePipelineRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePipelineRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePipelineRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePipelineRunResponse")
	}
	return
}

// deletePipelineRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) deletePipelineRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/pipelineRuns/{pipelineRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePipelineRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/DeletePipelineRun"
		err = common.PostProcessServiceError(err, "DataScience", "DeletePipelineRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProject Deletes the specified project. This operation fails unless all associated resources (notebook sessions or models) are in a DELETED state. You must delete all associated resources before deleting a project.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteProject.go.html to see an example of how to use DeleteProject API.
// A default retry strategy applies to this operation DeleteProject()
func (client DataScienceClient) DeleteProject(ctx context.Context, request DeleteProjectRequest) (response DeleteProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) deleteProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/projects/{projectId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Project/DeleteProject"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSchedule Deletes a Schedule resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DeleteSchedule.go.html to see an example of how to use DeleteSchedule API.
func (client DataScienceClient) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) (response DeleteScheduleResponse, err error) {
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
func (client DataScienceClient) deleteSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/schedules/{scheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/DeleteSchedule"
		err = common.PostProcessServiceError(err, "DataScience", "DeleteSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableMlApplicationInstanceViewTrigger Disable trigger of given name for given ML Application Instance View flow
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/DisableMlApplicationInstanceViewTrigger.go.html to see an example of how to use DisableMlApplicationInstanceViewTrigger API.
// A default retry strategy applies to this operation DisableMlApplicationInstanceViewTrigger()
func (client DataScienceClient) DisableMlApplicationInstanceViewTrigger(ctx context.Context, request DisableMlApplicationInstanceViewTriggerRequest) (response DisableMlApplicationInstanceViewTriggerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableMlApplicationInstanceViewTrigger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableMlApplicationInstanceViewTriggerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableMlApplicationInstanceViewTriggerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableMlApplicationInstanceViewTriggerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableMlApplicationInstanceViewTriggerResponse")
	}
	return
}

// disableMlApplicationInstanceViewTrigger implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) disableMlApplicationInstanceViewTrigger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}/actions/disableTrigger", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableMlApplicationInstanceViewTriggerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/DisableMlApplicationInstanceViewTrigger"
		err = common.PostProcessServiceError(err, "DataScience", "DisableMlApplicationInstanceViewTrigger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableMlApplicationInstanceViewTrigger Enable trigger of given name for given ML Application Instance View flow
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/EnableMlApplicationInstanceViewTrigger.go.html to see an example of how to use EnableMlApplicationInstanceViewTrigger API.
// A default retry strategy applies to this operation EnableMlApplicationInstanceViewTrigger()
func (client DataScienceClient) EnableMlApplicationInstanceViewTrigger(ctx context.Context, request EnableMlApplicationInstanceViewTriggerRequest) (response EnableMlApplicationInstanceViewTriggerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableMlApplicationInstanceViewTrigger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableMlApplicationInstanceViewTriggerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableMlApplicationInstanceViewTriggerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableMlApplicationInstanceViewTriggerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableMlApplicationInstanceViewTriggerResponse")
	}
	return
}

// enableMlApplicationInstanceViewTrigger implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) enableMlApplicationInstanceViewTrigger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}/actions/enableTrigger", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableMlApplicationInstanceViewTriggerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/EnableMlApplicationInstanceViewTrigger"
		err = common.PostProcessServiceError(err, "DataScience", "EnableMlApplicationInstanceViewTrigger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportModelArtifact Export model artifact from source to the service bucket
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ExportModelArtifact.go.html to see an example of how to use ExportModelArtifact API.
// A default retry strategy applies to this operation ExportModelArtifact()
func (client DataScienceClient) ExportModelArtifact(ctx context.Context, request ExportModelArtifactRequest) (response ExportModelArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportModelArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportModelArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportModelArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportModelArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportModelArtifactResponse")
	}
	return
}

// exportModelArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) exportModelArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/exportArtifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportModelArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/ExportModelArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "ExportModelArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataSciencePrivateEndpoint Retrieves an private endpoint using a `privateEndpointId`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetDataSciencePrivateEndpoint.go.html to see an example of how to use GetDataSciencePrivateEndpoint API.
// A default retry strategy applies to this operation GetDataSciencePrivateEndpoint()
func (client DataScienceClient) GetDataSciencePrivateEndpoint(ctx context.Context, request GetDataSciencePrivateEndpointRequest) (response GetDataSciencePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataSciencePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataSciencePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataSciencePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataSciencePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataSciencePrivateEndpointResponse")
	}
	return
}

// getDataSciencePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getDataSciencePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSciencePrivateEndpoints/{dataSciencePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataSciencePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/DataSciencePrivateEndpoint/GetDataSciencePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataScience", "GetDataSciencePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJob Gets a job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetJob.go.html to see an example of how to use GetJob API.
// A default retry strategy applies to this operation GetJob()
func (client DataScienceClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
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
func (client DataScienceClient) getJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/GetJob"
		err = common.PostProcessServiceError(err, "DataScience", "GetJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobArtifactContent Downloads job artifact content for specified job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetJobArtifactContent.go.html to see an example of how to use GetJobArtifactContent API.
// A default retry strategy applies to this operation GetJobArtifactContent()
func (client DataScienceClient) GetJobArtifactContent(ctx context.Context, request GetJobArtifactContentRequest) (response GetJobArtifactContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobArtifactContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobArtifactContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobArtifactContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobArtifactContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobArtifactContentResponse")
	}
	return
}

// getJobArtifactContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getJobArtifactContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobArtifactContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/GetJobArtifactContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetJobArtifactContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobRun Gets a job run.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetJobRun.go.html to see an example of how to use GetJobRun API.
// A default retry strategy applies to this operation GetJobRun()
func (client DataScienceClient) GetJobRun(ctx context.Context, request GetJobRunRequest) (response GetJobRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) getJobRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRun/GetJobRun"
		err = common.PostProcessServiceError(err, "DataScience", "GetJobRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplication Gets a MlApplication by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplication.go.html to see an example of how to use GetMlApplication API.
// A default retry strategy applies to this operation GetMlApplication()
func (client DataScienceClient) GetMlApplication(ctx context.Context, request GetMlApplicationRequest) (response GetMlApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationResponse")
	}
	return
}

// getMlApplication implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplications/{mlApplicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplication/GetMlApplication"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplicationHistoricalPackageContent Retrieves ML Application package for MlApplicationImplementationVersion with given id.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplicationHistoricalPackageContent.go.html to see an example of how to use GetMlApplicationHistoricalPackageContent API.
// A default retry strategy applies to this operation GetMlApplicationHistoricalPackageContent()
func (client DataScienceClient) GetMlApplicationHistoricalPackageContent(ctx context.Context, request GetMlApplicationHistoricalPackageContentRequest) (response GetMlApplicationHistoricalPackageContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplicationHistoricalPackageContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationHistoricalPackageContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationHistoricalPackageContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationHistoricalPackageContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationHistoricalPackageContentResponse")
	}
	return
}

// getMlApplicationHistoricalPackageContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplicationHistoricalPackageContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationImplementationVersions/{mlApplicationImplementationVersionId}/mlApplicationHistoricalPackage/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationHistoricalPackageContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementationVersion/GetMlApplicationHistoricalPackageContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplicationHistoricalPackageContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplicationImplementation Gets a MlApplicationImplementation by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplicationImplementation.go.html to see an example of how to use GetMlApplicationImplementation API.
// A default retry strategy applies to this operation GetMlApplicationImplementation()
func (client DataScienceClient) GetMlApplicationImplementation(ctx context.Context, request GetMlApplicationImplementationRequest) (response GetMlApplicationImplementationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplicationImplementation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationImplementationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationImplementationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationImplementationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationImplementationResponse")
	}
	return
}

// getMlApplicationImplementation implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplicationImplementation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationImplementations/{mlApplicationImplementationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationImplementationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/GetMlApplicationImplementation"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplicationImplementation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplicationImplementationVersion Gets a MlApplicationImplementationVersion by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplicationImplementationVersion.go.html to see an example of how to use GetMlApplicationImplementationVersion API.
// A default retry strategy applies to this operation GetMlApplicationImplementationVersion()
func (client DataScienceClient) GetMlApplicationImplementationVersion(ctx context.Context, request GetMlApplicationImplementationVersionRequest) (response GetMlApplicationImplementationVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplicationImplementationVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationImplementationVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationImplementationVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationImplementationVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationImplementationVersionResponse")
	}
	return
}

// getMlApplicationImplementationVersion implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplicationImplementationVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationImplementationVersions/{mlApplicationImplementationVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationImplementationVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementationVersion/GetMlApplicationImplementationVersion"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplicationImplementationVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplicationInstance Gets a MlApplicationInstance by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplicationInstance.go.html to see an example of how to use GetMlApplicationInstance API.
// A default retry strategy applies to this operation GetMlApplicationInstance()
func (client DataScienceClient) GetMlApplicationInstance(ctx context.Context, request GetMlApplicationInstanceRequest) (response GetMlApplicationInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplicationInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationInstanceResponse")
	}
	return
}

// getMlApplicationInstance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplicationInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationInstances/{mlApplicationInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/GetMlApplicationInstance"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplicationInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplicationInstanceView Gets a MlApplicationInstanceView by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplicationInstanceView.go.html to see an example of how to use GetMlApplicationInstanceView API.
// A default retry strategy applies to this operation GetMlApplicationInstanceView()
func (client DataScienceClient) GetMlApplicationInstanceView(ctx context.Context, request GetMlApplicationInstanceViewRequest) (response GetMlApplicationInstanceViewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplicationInstanceView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationInstanceViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationInstanceViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationInstanceViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationInstanceViewResponse")
	}
	return
}

// getMlApplicationInstanceView implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplicationInstanceView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationInstanceViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/GetMlApplicationInstanceView"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplicationInstanceView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMlApplicationPackageContent Retrieves last ML Application package uploaded for given ML Application Implementation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetMlApplicationPackageContent.go.html to see an example of how to use GetMlApplicationPackageContent API.
// A default retry strategy applies to this operation GetMlApplicationPackageContent()
func (client DataScienceClient) GetMlApplicationPackageContent(ctx context.Context, request GetMlApplicationPackageContentRequest) (response GetMlApplicationPackageContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMlApplicationPackageContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMlApplicationPackageContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMlApplicationPackageContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMlApplicationPackageContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMlApplicationPackageContentResponse")
	}
	return
}

// getMlApplicationPackageContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getMlApplicationPackageContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationImplementations/{mlApplicationImplementationId}/mlApplicationPackage/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMlApplicationPackageContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/GetMlApplicationPackageContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetMlApplicationPackageContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModel Gets the specified model's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModel.go.html to see an example of how to use GetModel API.
// A default retry strategy applies to this operation GetModel()
func (client DataScienceClient) GetModel(ctx context.Context, request GetModelRequest) (response GetModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelResponse")
	}
	return
}

// getModel implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/models/{modelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/GetModel"
		err = common.PostProcessServiceError(err, "DataScience", "GetModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModelArtifactContent Downloads model artifact content for specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModelArtifactContent.go.html to see an example of how to use GetModelArtifactContent API.
// A default retry strategy applies to this operation GetModelArtifactContent()
func (client DataScienceClient) GetModelArtifactContent(ctx context.Context, request GetModelArtifactContentRequest) (response GetModelArtifactContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModelArtifactContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelArtifactContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelArtifactContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelArtifactContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelArtifactContentResponse")
	}
	return
}

// getModelArtifactContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModelArtifactContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/models/{modelId}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelArtifactContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/GetModelArtifactContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetModelArtifactContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModelCustomMetadatumArtifactContent Downloads model custom metadata artifact content for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModelCustomMetadatumArtifactContent.go.html to see an example of how to use GetModelCustomMetadatumArtifactContent API.
// A default retry strategy applies to this operation GetModelCustomMetadatumArtifactContent()
func (client DataScienceClient) GetModelCustomMetadatumArtifactContent(ctx context.Context, request GetModelCustomMetadatumArtifactContentRequest) (response GetModelCustomMetadatumArtifactContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModelCustomMetadatumArtifactContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelCustomMetadatumArtifactContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelCustomMetadatumArtifactContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelCustomMetadatumArtifactContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelCustomMetadatumArtifactContentResponse")
	}
	return
}

// getModelCustomMetadatumArtifactContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModelCustomMetadatumArtifactContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/models/{modelId}/customMetadata/{metadatumKeyName}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelCustomMetadatumArtifactContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/GetModelCustomMetadatumArtifactContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetModelCustomMetadatumArtifactContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModelDefinedMetadatumArtifactContent Downloads model defined metadata artifact content for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModelDefinedMetadatumArtifactContent.go.html to see an example of how to use GetModelDefinedMetadatumArtifactContent API.
// A default retry strategy applies to this operation GetModelDefinedMetadatumArtifactContent()
func (client DataScienceClient) GetModelDefinedMetadatumArtifactContent(ctx context.Context, request GetModelDefinedMetadatumArtifactContentRequest) (response GetModelDefinedMetadatumArtifactContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModelDefinedMetadatumArtifactContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelDefinedMetadatumArtifactContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelDefinedMetadatumArtifactContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelDefinedMetadatumArtifactContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelDefinedMetadatumArtifactContentResponse")
	}
	return
}

// getModelDefinedMetadatumArtifactContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModelDefinedMetadatumArtifactContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/models/{modelId}/definedMetadata/{metadatumKeyName}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelDefinedMetadatumArtifactContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/GetModelDefinedMetadatumArtifactContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetModelDefinedMetadatumArtifactContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModelDeployment Retrieves the model deployment for the specified `modelDeploymentId`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModelDeployment.go.html to see an example of how to use GetModelDeployment API.
// A default retry strategy applies to this operation GetModelDeployment()
func (client DataScienceClient) GetModelDeployment(ctx context.Context, request GetModelDeploymentRequest) (response GetModelDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelDeploymentResponse")
	}
	return
}

// getModelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/modelDeployments/{modelDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/GetModelDeployment"
		err = common.PostProcessServiceError(err, "DataScience", "GetModelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModelProvenance Gets provenance information for specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModelProvenance.go.html to see an example of how to use GetModelProvenance API.
// A default retry strategy applies to this operation GetModelProvenance()
func (client DataScienceClient) GetModelProvenance(ctx context.Context, request GetModelProvenanceRequest) (response GetModelProvenanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModelProvenance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelProvenanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelProvenanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelProvenanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelProvenanceResponse")
	}
	return
}

// getModelProvenance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModelProvenance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/models/{modelId}/provenance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelProvenanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/GetModelProvenance"
		err = common.PostProcessServiceError(err, "DataScience", "GetModelProvenance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetModelVersionSet Gets the specified model version set information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetModelVersionSet.go.html to see an example of how to use GetModelVersionSet API.
// A default retry strategy applies to this operation GetModelVersionSet()
func (client DataScienceClient) GetModelVersionSet(ctx context.Context, request GetModelVersionSetRequest) (response GetModelVersionSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getModelVersionSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetModelVersionSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetModelVersionSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetModelVersionSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetModelVersionSetResponse")
	}
	return
}

// getModelVersionSet implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getModelVersionSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/modelVersionSets/{modelVersionSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetModelVersionSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelVersionSet/GetModelVersionSet"
		err = common.PostProcessServiceError(err, "DataScience", "GetModelVersionSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNotebookSession Gets the specified notebook session's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetNotebookSession.go.html to see an example of how to use GetNotebookSession API.
// A default retry strategy applies to this operation GetNotebookSession()
func (client DataScienceClient) GetNotebookSession(ctx context.Context, request GetNotebookSessionRequest) (response GetNotebookSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNotebookSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNotebookSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNotebookSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNotebookSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNotebookSessionResponse")
	}
	return
}

// getNotebookSession implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getNotebookSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/notebookSessions/{notebookSessionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNotebookSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/GetNotebookSession"
		err = common.PostProcessServiceError(err, "DataScience", "GetNotebookSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPipeline Gets a Pipeline by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetPipeline.go.html to see an example of how to use GetPipeline API.
// A default retry strategy applies to this operation GetPipeline()
func (client DataScienceClient) GetPipeline(ctx context.Context, request GetPipelineRequest) (response GetPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) getPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/GetPipeline"
		err = common.PostProcessServiceError(err, "DataScience", "GetPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPipelineRun Gets a PipelineRun by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetPipelineRun.go.html to see an example of how to use GetPipelineRun API.
// A default retry strategy applies to this operation GetPipelineRun()
func (client DataScienceClient) GetPipelineRun(ctx context.Context, request GetPipelineRunRequest) (response GetPipelineRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPipelineRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPipelineRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPipelineRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPipelineRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPipelineRunResponse")
	}
	return
}

// getPipelineRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getPipelineRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelineRuns/{pipelineRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPipelineRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/GetPipelineRun"
		err = common.PostProcessServiceError(err, "DataScience", "GetPipelineRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProject Gets the specified project's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetProject.go.html to see an example of how to use GetProject API.
// A default retry strategy applies to this operation GetProject()
func (client DataScienceClient) GetProject(ctx context.Context, request GetProjectRequest) (response GetProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) getProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/projects/{projectId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Project/GetProject"
		err = common.PostProcessServiceError(err, "DataScience", "GetProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSchedule Gets a Schedule by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetSchedule.go.html to see an example of how to use GetSchedule API.
// A default retry strategy applies to this operation GetSchedule()
func (client DataScienceClient) GetSchedule(ctx context.Context, request GetScheduleRequest) (response GetScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) getSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedules/{scheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/GetSchedule"
		err = common.PostProcessServiceError(err, "DataScience", "GetSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStepArtifactContent Download the artifact for a step in the pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetStepArtifactContent.go.html to see an example of how to use GetStepArtifactContent API.
// A default retry strategy applies to this operation GetStepArtifactContent()
func (client DataScienceClient) GetStepArtifactContent(ctx context.Context, request GetStepArtifactContentRequest) (response GetStepArtifactContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStepArtifactContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStepArtifactContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStepArtifactContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStepArtifactContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStepArtifactContentResponse")
	}
	return
}

// getStepArtifactContent implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) getStepArtifactContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}/steps/{stepName}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStepArtifactContentResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/GetStepArtifactContent"
		err = common.PostProcessServiceError(err, "DataScience", "GetStepArtifactContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the specified work request's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DataScienceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DataScienceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DataScience", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadJobArtifact Gets job artifact metadata.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/HeadJobArtifact.go.html to see an example of how to use HeadJobArtifact API.
// A default retry strategy applies to this operation HeadJobArtifact()
func (client DataScienceClient) HeadJobArtifact(ctx context.Context, request HeadJobArtifactRequest) (response HeadJobArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headJobArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadJobArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadJobArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadJobArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadJobArtifactResponse")
	}
	return
}

// headJobArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) headJobArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/jobs/{jobId}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadJobArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/HeadJobArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "HeadJobArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadModelArtifact Gets model artifact metadata for specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/HeadModelArtifact.go.html to see an example of how to use HeadModelArtifact API.
// A default retry strategy applies to this operation HeadModelArtifact()
func (client DataScienceClient) HeadModelArtifact(ctx context.Context, request HeadModelArtifactRequest) (response HeadModelArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headModelArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadModelArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadModelArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadModelArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadModelArtifactResponse")
	}
	return
}

// headModelArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) headModelArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/models/{modelId}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadModelArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/HeadModelArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "HeadModelArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadModelCustomMetadatumArtifact Gets custom metadata artifact metadata for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/HeadModelCustomMetadatumArtifact.go.html to see an example of how to use HeadModelCustomMetadatumArtifact API.
// A default retry strategy applies to this operation HeadModelCustomMetadatumArtifact()
func (client DataScienceClient) HeadModelCustomMetadatumArtifact(ctx context.Context, request HeadModelCustomMetadatumArtifactRequest) (response HeadModelCustomMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headModelCustomMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadModelCustomMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadModelCustomMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadModelCustomMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadModelCustomMetadatumArtifactResponse")
	}
	return
}

// headModelCustomMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) headModelCustomMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/models/{modelId}/customMetadata/{metadatumKeyName}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadModelCustomMetadatumArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/HeadModelCustomMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "HeadModelCustomMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadModelDefinedMetadatumArtifact Gets defined metadata artifact metadata for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/HeadModelDefinedMetadatumArtifact.go.html to see an example of how to use HeadModelDefinedMetadatumArtifact API.
// A default retry strategy applies to this operation HeadModelDefinedMetadatumArtifact()
func (client DataScienceClient) HeadModelDefinedMetadatumArtifact(ctx context.Context, request HeadModelDefinedMetadatumArtifactRequest) (response HeadModelDefinedMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headModelDefinedMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadModelDefinedMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadModelDefinedMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadModelDefinedMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadModelDefinedMetadatumArtifactResponse")
	}
	return
}

// headModelDefinedMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) headModelDefinedMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/models/{modelId}/definedMetadata/{metadatumKeyName}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadModelDefinedMetadatumArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/HeadModelDefinedMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "HeadModelDefinedMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadStepArtifact Get the artifact metadata for a step in the pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/HeadStepArtifact.go.html to see an example of how to use HeadStepArtifact API.
// A default retry strategy applies to this operation HeadStepArtifact()
func (client DataScienceClient) HeadStepArtifact(ctx context.Context, request HeadStepArtifactRequest) (response HeadStepArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headStepArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadStepArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadStepArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadStepArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadStepArtifactResponse")
	}
	return
}

// headStepArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) headStepArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/pipelines/{pipelineId}/steps/{stepName}/artifact/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadStepArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/HeadStepArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "HeadStepArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportModelArtifact Import model artifact from service bucket
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ImportModelArtifact.go.html to see an example of how to use ImportModelArtifact API.
func (client DataScienceClient) ImportModelArtifact(ctx context.Context, request ImportModelArtifactRequest) (response ImportModelArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.importModelArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportModelArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportModelArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportModelArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportModelArtifactResponse")
	}
	return
}

// importModelArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) importModelArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/importArtifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportModelArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/ImportModelArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "ImportModelArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListContainers List containers.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListContainers.go.html to see an example of how to use ListContainers API.
// A default retry strategy applies to this operation ListContainers()
func (client DataScienceClient) ListContainers(ctx context.Context, request ListContainersRequest) (response ListContainersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listContainers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListContainersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListContainersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListContainersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListContainersResponse")
	}
	return
}

// listContainers implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listContainers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/containers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListContainersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ContainerSummary/ListContainers"
		err = common.PostProcessServiceError(err, "DataScience", "ListContainers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataSciencePrivateEndpoints Lists all Data Science private endpoints in the specified compartment. The query must include compartmentId. The query can also include one other parameter. If the query doesn't include compartmentId, or includes compartmentId with two or more other parameters, then an error is returned.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListDataSciencePrivateEndpoints.go.html to see an example of how to use ListDataSciencePrivateEndpoints API.
// A default retry strategy applies to this operation ListDataSciencePrivateEndpoints()
func (client DataScienceClient) ListDataSciencePrivateEndpoints(ctx context.Context, request ListDataSciencePrivateEndpointsRequest) (response ListDataSciencePrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataSciencePrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataSciencePrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataSciencePrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataSciencePrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataSciencePrivateEndpointsResponse")
	}
	return
}

// listDataSciencePrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listDataSciencePrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSciencePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataSciencePrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/DataSciencePrivateEndpoint/ListDataSciencePrivateEndpoints"
		err = common.PostProcessServiceError(err, "DataScience", "ListDataSciencePrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFastLaunchJobConfigs List fast launch capable job configs in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListFastLaunchJobConfigs.go.html to see an example of how to use ListFastLaunchJobConfigs API.
// A default retry strategy applies to this operation ListFastLaunchJobConfigs()
func (client DataScienceClient) ListFastLaunchJobConfigs(ctx context.Context, request ListFastLaunchJobConfigsRequest) (response ListFastLaunchJobConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFastLaunchJobConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFastLaunchJobConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFastLaunchJobConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFastLaunchJobConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFastLaunchJobConfigsResponse")
	}
	return
}

// listFastLaunchJobConfigs implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listFastLaunchJobConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fastLaunchJobConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFastLaunchJobConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/FastLaunchJobConfigSummary/ListFastLaunchJobConfigs"
		err = common.PostProcessServiceError(err, "DataScience", "ListFastLaunchJobConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobRuns List out job runs.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListJobRuns.go.html to see an example of how to use ListJobRuns API.
// A default retry strategy applies to this operation ListJobRuns()
func (client DataScienceClient) ListJobRuns(ctx context.Context, request ListJobRunsRequest) (response ListJobRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) listJobRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRunSummary/ListJobRuns"
		err = common.PostProcessServiceError(err, "DataScience", "ListJobRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobShapes List job shapes available in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListJobShapes.go.html to see an example of how to use ListJobShapes API.
// A default retry strategy applies to this operation ListJobShapes()
func (client DataScienceClient) ListJobShapes(ctx context.Context, request ListJobShapesRequest) (response ListJobShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobShapesResponse")
	}
	return
}

// listJobShapes implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listJobShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobShapeSummary/ListJobShapes"
		err = common.PostProcessServiceError(err, "DataScience", "ListJobShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobs List jobs in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListJobs.go.html to see an example of how to use ListJobs API.
// A default retry strategy applies to this operation ListJobs()
func (client DataScienceClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
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
func (client DataScienceClient) listJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobSummary/ListJobs"
		err = common.PostProcessServiceError(err, "DataScience", "ListJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMlApplicationImplementationVersions Returns a list of MlApplicationImplementationVersions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationImplementationVersions.go.html to see an example of how to use ListMlApplicationImplementationVersions API.
// A default retry strategy applies to this operation ListMlApplicationImplementationVersions()
func (client DataScienceClient) ListMlApplicationImplementationVersions(ctx context.Context, request ListMlApplicationImplementationVersionsRequest) (response ListMlApplicationImplementationVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMlApplicationImplementationVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMlApplicationImplementationVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMlApplicationImplementationVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMlApplicationImplementationVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMlApplicationImplementationVersionsResponse")
	}
	return
}

// listMlApplicationImplementationVersions implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listMlApplicationImplementationVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationImplementationVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMlApplicationImplementationVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementationVersion/ListMlApplicationImplementationVersions"
		err = common.PostProcessServiceError(err, "DataScience", "ListMlApplicationImplementationVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMlApplicationImplementations Returns a list of MlApplicationImplementations.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationImplementations.go.html to see an example of how to use ListMlApplicationImplementations API.
// A default retry strategy applies to this operation ListMlApplicationImplementations()
func (client DataScienceClient) ListMlApplicationImplementations(ctx context.Context, request ListMlApplicationImplementationsRequest) (response ListMlApplicationImplementationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMlApplicationImplementations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMlApplicationImplementationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMlApplicationImplementationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMlApplicationImplementationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMlApplicationImplementationsResponse")
	}
	return
}

// listMlApplicationImplementations implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listMlApplicationImplementations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationImplementations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMlApplicationImplementationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/ListMlApplicationImplementations"
		err = common.PostProcessServiceError(err, "DataScience", "ListMlApplicationImplementations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMlApplicationInstanceViews Returns a list of MlApplicationInstanceViews.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationInstanceViews.go.html to see an example of how to use ListMlApplicationInstanceViews API.
// A default retry strategy applies to this operation ListMlApplicationInstanceViews()
func (client DataScienceClient) ListMlApplicationInstanceViews(ctx context.Context, request ListMlApplicationInstanceViewsRequest) (response ListMlApplicationInstanceViewsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMlApplicationInstanceViews, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMlApplicationInstanceViewsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMlApplicationInstanceViewsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMlApplicationInstanceViewsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMlApplicationInstanceViewsResponse")
	}
	return
}

// listMlApplicationInstanceViews implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listMlApplicationInstanceViews(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationInstanceViews", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMlApplicationInstanceViewsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/ListMlApplicationInstanceViews"
		err = common.PostProcessServiceError(err, "DataScience", "ListMlApplicationInstanceViews", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMlApplicationInstances Returns a list of MlApplicationsInstances.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationInstances.go.html to see an example of how to use ListMlApplicationInstances API.
// A default retry strategy applies to this operation ListMlApplicationInstances()
func (client DataScienceClient) ListMlApplicationInstances(ctx context.Context, request ListMlApplicationInstancesRequest) (response ListMlApplicationInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMlApplicationInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMlApplicationInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMlApplicationInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMlApplicationInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMlApplicationInstancesResponse")
	}
	return
}

// listMlApplicationInstances implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listMlApplicationInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplicationInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMlApplicationInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/ListMlApplicationInstances"
		err = common.PostProcessServiceError(err, "DataScience", "ListMlApplicationInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMlApplications Returns a list of MlApplications.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplications.go.html to see an example of how to use ListMlApplications API.
// A default retry strategy applies to this operation ListMlApplications()
func (client DataScienceClient) ListMlApplications(ctx context.Context, request ListMlApplicationsRequest) (response ListMlApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMlApplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMlApplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMlApplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMlApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMlApplicationsResponse")
	}
	return
}

// listMlApplications implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listMlApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mlApplications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMlApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplication/ListMlApplications"
		err = common.PostProcessServiceError(err, "DataScience", "ListMlApplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModelDeploymentShapes Lists the valid model deployment shapes.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelDeploymentShapes.go.html to see an example of how to use ListModelDeploymentShapes API.
// A default retry strategy applies to this operation ListModelDeploymentShapes()
func (client DataScienceClient) ListModelDeploymentShapes(ctx context.Context, request ListModelDeploymentShapesRequest) (response ListModelDeploymentShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModelDeploymentShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModelDeploymentShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModelDeploymentShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModelDeploymentShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModelDeploymentShapesResponse")
	}
	return
}

// listModelDeploymentShapes implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listModelDeploymentShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/modelDeploymentShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModelDeploymentShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeploymentShapeSummary/ListModelDeploymentShapes"
		err = common.PostProcessServiceError(err, "DataScience", "ListModelDeploymentShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModelDeployments Lists all model deployments in the specified compartment. Only one parameter other than compartmentId may also be included in a query. The query must include compartmentId. If the query does not include compartmentId, or includes compartmentId but two or more other parameters an error is returned.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelDeployments.go.html to see an example of how to use ListModelDeployments API.
// A default retry strategy applies to this operation ListModelDeployments()
func (client DataScienceClient) ListModelDeployments(ctx context.Context, request ListModelDeploymentsRequest) (response ListModelDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModelDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModelDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModelDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModelDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModelDeploymentsResponse")
	}
	return
}

// listModelDeployments implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listModelDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/modelDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModelDeploymentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeploymentSummary/ListModelDeployments"
		err = common.PostProcessServiceError(err, "DataScience", "ListModelDeployments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModelVersionSets Lists model version sets in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelVersionSets.go.html to see an example of how to use ListModelVersionSets API.
// A default retry strategy applies to this operation ListModelVersionSets()
func (client DataScienceClient) ListModelVersionSets(ctx context.Context, request ListModelVersionSetsRequest) (response ListModelVersionSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModelVersionSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModelVersionSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModelVersionSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModelVersionSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModelVersionSetsResponse")
	}
	return
}

// listModelVersionSets implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listModelVersionSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/modelVersionSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModelVersionSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelVersionSetSummary/ListModelVersionSets"
		err = common.PostProcessServiceError(err, "DataScience", "ListModelVersionSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListModels Lists models in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModels.go.html to see an example of how to use ListModels API.
// A default retry strategy applies to this operation ListModels()
func (client DataScienceClient) ListModels(ctx context.Context, request ListModelsRequest) (response ListModelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listModels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListModelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListModelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListModelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListModelsResponse")
	}
	return
}

// listModels implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listModels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/models", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListModelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelSummary/ListModels"
		err = common.PostProcessServiceError(err, "DataScience", "ListModels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNotebookSessionShapes Lists the valid notebook session shapes.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListNotebookSessionShapes.go.html to see an example of how to use ListNotebookSessionShapes API.
// A default retry strategy applies to this operation ListNotebookSessionShapes()
func (client DataScienceClient) ListNotebookSessionShapes(ctx context.Context, request ListNotebookSessionShapesRequest) (response ListNotebookSessionShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNotebookSessionShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNotebookSessionShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNotebookSessionShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNotebookSessionShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNotebookSessionShapesResponse")
	}
	return
}

// listNotebookSessionShapes implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listNotebookSessionShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/notebookSessionShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNotebookSessionShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSessionShapeSummary/ListNotebookSessionShapes"
		err = common.PostProcessServiceError(err, "DataScience", "ListNotebookSessionShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNotebookSessions Lists the notebook sessions in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListNotebookSessions.go.html to see an example of how to use ListNotebookSessions API.
// A default retry strategy applies to this operation ListNotebookSessions()
func (client DataScienceClient) ListNotebookSessions(ctx context.Context, request ListNotebookSessionsRequest) (response ListNotebookSessionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNotebookSessions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNotebookSessionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNotebookSessionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNotebookSessionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNotebookSessionsResponse")
	}
	return
}

// listNotebookSessions implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listNotebookSessions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/notebookSessions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNotebookSessionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSessionSummary/ListNotebookSessions"
		err = common.PostProcessServiceError(err, "DataScience", "ListNotebookSessions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelineRuns Returns a list of PipelineRuns.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListPipelineRuns.go.html to see an example of how to use ListPipelineRuns API.
// A default retry strategy applies to this operation ListPipelineRuns()
func (client DataScienceClient) ListPipelineRuns(ctx context.Context, request ListPipelineRunsRequest) (response ListPipelineRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelineRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelineRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelineRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelineRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelineRunsResponse")
	}
	return
}

// listPipelineRuns implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) listPipelineRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelineRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelineRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/ListPipelineRuns"
		err = common.PostProcessServiceError(err, "DataScience", "ListPipelineRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelines Returns a list of Pipelines.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListPipelines.go.html to see an example of how to use ListPipelines API.
// A default retry strategy applies to this operation ListPipelines()
func (client DataScienceClient) ListPipelines(ctx context.Context, request ListPipelinesRequest) (response ListPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) listPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/ListPipelines"
		err = common.PostProcessServiceError(err, "DataScience", "ListPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProjects Lists projects in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListProjects.go.html to see an example of how to use ListProjects API.
// A default retry strategy applies to this operation ListProjects()
func (client DataScienceClient) ListProjects(ctx context.Context, request ListProjectsRequest) (response ListProjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) listProjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/projects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ProjectSummary/ListProjects"
		err = common.PostProcessServiceError(err, "DataScience", "ListProjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchedules Returns a list of Schedules.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListSchedules.go.html to see an example of how to use ListSchedules API.
// A default retry strategy applies to this operation ListSchedules()
func (client DataScienceClient) ListSchedules(ctx context.Context, request ListSchedulesRequest) (response ListSchedulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) listSchedules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchedulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/ListSchedules"
		err = common.PostProcessServiceError(err, "DataScience", "ListSchedules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists work request errors for the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DataScienceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DataScienceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/WorkRequest/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DataScience", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists work request logs for the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DataScienceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DataScienceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/WorkRequest/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DataScience", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists work requests in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DataScienceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DataScienceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DataScience", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutMlApplicationPackage Upload ML Application Package
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/PutMlApplicationPackage.go.html to see an example of how to use PutMlApplicationPackage API.
// A default retry strategy applies to this operation PutMlApplicationPackage()
func (client DataScienceClient) PutMlApplicationPackage(ctx context.Context, request PutMlApplicationPackageRequest) (response PutMlApplicationPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putMlApplicationPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutMlApplicationPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutMlApplicationPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutMlApplicationPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutMlApplicationPackageResponse")
	}
	return
}

// putMlApplicationPackage implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) putMlApplicationPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mlApplicationImplementations/{mlApplicationImplementationId}/mlApplicationPackage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutMlApplicationPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/PutMlApplicationPackage"
		err = common.PostProcessServiceError(err, "DataScience", "PutMlApplicationPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RecoverMlApplicationInstanceView Provider can initiate recovery of the resource only if MlApplicationInstanceView is in one of the recoverable sub-states (RECOVERABLE_PROVIDER_ISSUE, RECOVERABLE_SERVICE_ISSUE).
// Provider should investigate (using MlApplicationInstanceView lifecycleDetails, relevant logs and metrics) and fix the issue before the recovery is initiated.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/RecoverMlApplicationInstanceView.go.html to see an example of how to use RecoverMlApplicationInstanceView API.
// A default retry strategy applies to this operation RecoverMlApplicationInstanceView()
func (client DataScienceClient) RecoverMlApplicationInstanceView(ctx context.Context, request RecoverMlApplicationInstanceViewRequest) (response RecoverMlApplicationInstanceViewResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.recoverMlApplicationInstanceView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecoverMlApplicationInstanceViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecoverMlApplicationInstanceViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecoverMlApplicationInstanceViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecoverMlApplicationInstanceViewResponse")
	}
	return
}

// recoverMlApplicationInstanceView implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) recoverMlApplicationInstanceView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}/actions/recover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RecoverMlApplicationInstanceViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/RecoverMlApplicationInstanceView"
		err = common.PostProcessServiceError(err, "DataScience", "RecoverMlApplicationInstanceView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RegisterModelArtifactReference Registers model artifact reference metadata
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/RegisterModelArtifactReference.go.html to see an example of how to use RegisterModelArtifactReference API.
// A default retry strategy applies to this operation RegisterModelArtifactReference()
func (client DataScienceClient) RegisterModelArtifactReference(ctx context.Context, request RegisterModelArtifactReferenceRequest) (response RegisterModelArtifactReferenceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.registerModelArtifactReference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RegisterModelArtifactReferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RegisterModelArtifactReferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RegisterModelArtifactReferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RegisterModelArtifactReferenceResponse")
	}
	return
}

// registerModelArtifactReference implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) registerModelArtifactReference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/registerArtifactReference", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RegisterModelArtifactReferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/RegisterModelArtifactReferenceDetails/RegisterModelArtifactReference"
		err = common.PostProcessServiceError(err, "DataScience", "RegisterModelArtifactReference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreArchivedModelArtifact Restore archived model artifact
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/RestoreArchivedModelArtifact.go.html to see an example of how to use RestoreArchivedModelArtifact API.
// A default retry strategy applies to this operation RestoreArchivedModelArtifact()
func (client DataScienceClient) RestoreArchivedModelArtifact(ctx context.Context, request RestoreArchivedModelArtifactRequest) (response RestoreArchivedModelArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restoreArchivedModelArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreArchivedModelArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreArchivedModelArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreArchivedModelArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreArchivedModelArtifactResponse")
	}
	return
}

// restoreArchivedModelArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) restoreArchivedModelArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/models/{modelId}/actions/restore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreArchivedModelArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/RestoreArchivedModelArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "RestoreArchivedModelArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TriggerMlApplicationInstanceFlow Trigger ML Application Instance flow if possible
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/TriggerMlApplicationInstanceFlow.go.html to see an example of how to use TriggerMlApplicationInstanceFlow API.
// A default retry strategy applies to this operation TriggerMlApplicationInstanceFlow()
func (client DataScienceClient) TriggerMlApplicationInstanceFlow(ctx context.Context, request TriggerMlApplicationInstanceFlowRequest) (response TriggerMlApplicationInstanceFlowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.triggerMlApplicationInstanceFlow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TriggerMlApplicationInstanceFlowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TriggerMlApplicationInstanceFlowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TriggerMlApplicationInstanceFlowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TriggerMlApplicationInstanceFlowResponse")
	}
	return
}

// triggerMlApplicationInstanceFlow implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) triggerMlApplicationInstanceFlow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstances/{mlApplicationInstanceId}/actions/trigger", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TriggerMlApplicationInstanceFlowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/TriggerMlApplicationInstanceFlow"
		err = common.PostProcessServiceError(err, "DataScience", "TriggerMlApplicationInstanceFlow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TriggerMlApplicationInstanceViewFlow Trigger ML Application Instance View flow if possible
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/TriggerMlApplicationInstanceViewFlow.go.html to see an example of how to use TriggerMlApplicationInstanceViewFlow API.
// A default retry strategy applies to this operation TriggerMlApplicationInstanceViewFlow()
func (client DataScienceClient) TriggerMlApplicationInstanceViewFlow(ctx context.Context, request TriggerMlApplicationInstanceViewFlowRequest) (response TriggerMlApplicationInstanceViewFlowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.triggerMlApplicationInstanceViewFlow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TriggerMlApplicationInstanceViewFlowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TriggerMlApplicationInstanceViewFlowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TriggerMlApplicationInstanceViewFlowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TriggerMlApplicationInstanceViewFlowResponse")
	}
	return
}

// triggerMlApplicationInstanceViewFlow implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) triggerMlApplicationInstanceViewFlow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}/actions/trigger", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TriggerMlApplicationInstanceViewFlowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/TriggerMlApplicationInstanceViewFlow"
		err = common.PostProcessServiceError(err, "DataScience", "TriggerMlApplicationInstanceViewFlow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataSciencePrivateEndpoint Updates a private endpoint using a `privateEndpointId`.  If changes to a private endpoint match
// a previously defined private endpoint, then a 409 status code is returned.  This indicates
// that a conflict has been detected.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateDataSciencePrivateEndpoint.go.html to see an example of how to use UpdateDataSciencePrivateEndpoint API.
func (client DataScienceClient) UpdateDataSciencePrivateEndpoint(ctx context.Context, request UpdateDataSciencePrivateEndpointRequest) (response UpdateDataSciencePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataSciencePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataSciencePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataSciencePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataSciencePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataSciencePrivateEndpointResponse")
	}
	return
}

// updateDataSciencePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateDataSciencePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dataSciencePrivateEndpoints/{dataSciencePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataSciencePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/DataSciencePrivateEndpoint/UpdateDataSciencePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateDataSciencePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJob Updates a job.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateJob.go.html to see an example of how to use UpdateJob API.
// A default retry strategy applies to this operation UpdateJob()
func (client DataScienceClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
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
func (client DataScienceClient) updateJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Job/UpdateJob"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJobRun Updates a job run.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateJobRun.go.html to see an example of how to use UpdateJobRun API.
// A default retry strategy applies to this operation UpdateJobRun()
func (client DataScienceClient) UpdateJobRun(ctx context.Context, request UpdateJobRunRequest) (response UpdateJobRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJobRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobRunResponse")
	}
	return
}

// updateJobRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateJobRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/jobRuns/{jobRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJobRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/JobRun/UpdateJobRun"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateJobRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMlApplication Updates the MlApplication
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateMlApplication.go.html to see an example of how to use UpdateMlApplication API.
// A default retry strategy applies to this operation UpdateMlApplication()
func (client DataScienceClient) UpdateMlApplication(ctx context.Context, request UpdateMlApplicationRequest) (response UpdateMlApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMlApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMlApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMlApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMlApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMlApplicationResponse")
	}
	return
}

// updateMlApplication implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateMlApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mlApplications/{mlApplicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMlApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplication/UpdateMlApplication"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateMlApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMlApplicationImplementation Updates the MlApplicationImplementation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateMlApplicationImplementation.go.html to see an example of how to use UpdateMlApplicationImplementation API.
// A default retry strategy applies to this operation UpdateMlApplicationImplementation()
func (client DataScienceClient) UpdateMlApplicationImplementation(ctx context.Context, request UpdateMlApplicationImplementationRequest) (response UpdateMlApplicationImplementationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMlApplicationImplementation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMlApplicationImplementationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMlApplicationImplementationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMlApplicationImplementationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMlApplicationImplementationResponse")
	}
	return
}

// updateMlApplicationImplementation implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateMlApplicationImplementation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mlApplicationImplementations/{mlApplicationImplementationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMlApplicationImplementationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementation/UpdateMlApplicationImplementation"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateMlApplicationImplementation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMlApplicationImplementationVersion Updates the MlApplicationImplementationVersion
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateMlApplicationImplementationVersion.go.html to see an example of how to use UpdateMlApplicationImplementationVersion API.
// A default retry strategy applies to this operation UpdateMlApplicationImplementationVersion()
func (client DataScienceClient) UpdateMlApplicationImplementationVersion(ctx context.Context, request UpdateMlApplicationImplementationVersionRequest) (response UpdateMlApplicationImplementationVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMlApplicationImplementationVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMlApplicationImplementationVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMlApplicationImplementationVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMlApplicationImplementationVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMlApplicationImplementationVersionResponse")
	}
	return
}

// updateMlApplicationImplementationVersion implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateMlApplicationImplementationVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mlApplicationImplementationVersions/{mlApplicationImplementationVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMlApplicationImplementationVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationImplementationVersion/UpdateMlApplicationImplementationVersion"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateMlApplicationImplementationVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMlApplicationInstance Updates the MlApplicationInstance
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateMlApplicationInstance.go.html to see an example of how to use UpdateMlApplicationInstance API.
// A default retry strategy applies to this operation UpdateMlApplicationInstance()
func (client DataScienceClient) UpdateMlApplicationInstance(ctx context.Context, request UpdateMlApplicationInstanceRequest) (response UpdateMlApplicationInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMlApplicationInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMlApplicationInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMlApplicationInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMlApplicationInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMlApplicationInstanceResponse")
	}
	return
}

// updateMlApplicationInstance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateMlApplicationInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mlApplicationInstances/{mlApplicationInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMlApplicationInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstance/UpdateMlApplicationInstance"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateMlApplicationInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMlApplicationInstanceView Updates the MlApplicationInstanceView
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateMlApplicationInstanceView.go.html to see an example of how to use UpdateMlApplicationInstanceView API.
func (client DataScienceClient) UpdateMlApplicationInstanceView(ctx context.Context, request UpdateMlApplicationInstanceViewRequest) (response UpdateMlApplicationInstanceViewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMlApplicationInstanceView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMlApplicationInstanceViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMlApplicationInstanceViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMlApplicationInstanceViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMlApplicationInstanceViewResponse")
	}
	return
}

// updateMlApplicationInstanceView implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateMlApplicationInstanceView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mlApplicationInstanceViews/{mlApplicationInstanceViewId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMlApplicationInstanceViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/MlApplicationInstanceView/UpdateMlApplicationInstanceView"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateMlApplicationInstanceView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateModel Updates the properties of a model. You can update the `displayName`, `description`, `freeformTags`, and `definedTags` properties.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateModel.go.html to see an example of how to use UpdateModel API.
// A default retry strategy applies to this operation UpdateModel()
func (client DataScienceClient) UpdateModel(ctx context.Context, request UpdateModelRequest) (response UpdateModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateModelResponse")
	}
	return
}

// updateModel implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/models/{modelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/UpdateModel"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateModelCustomMetadatumArtifact Updates model custom metadata artifact for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateModelCustomMetadatumArtifact.go.html to see an example of how to use UpdateModelCustomMetadatumArtifact API.
// A default retry strategy applies to this operation UpdateModelCustomMetadatumArtifact()
func (client DataScienceClient) UpdateModelCustomMetadatumArtifact(ctx context.Context, request UpdateModelCustomMetadatumArtifactRequest) (response UpdateModelCustomMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateModelCustomMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateModelCustomMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateModelCustomMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateModelCustomMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateModelCustomMetadatumArtifactResponse")
	}
	return
}

// updateModelCustomMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateModelCustomMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/models/{modelId}/customMetadata/{metadatumKeyName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateModelCustomMetadatumArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/UpdateModelCustomMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateModelCustomMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateModelDefinedMetadatumArtifact Updates model defined metadata artifact for specified model metadata key.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateModelDefinedMetadatumArtifact.go.html to see an example of how to use UpdateModelDefinedMetadatumArtifact API.
// A default retry strategy applies to this operation UpdateModelDefinedMetadatumArtifact()
func (client DataScienceClient) UpdateModelDefinedMetadatumArtifact(ctx context.Context, request UpdateModelDefinedMetadatumArtifactRequest) (response UpdateModelDefinedMetadatumArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateModelDefinedMetadatumArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateModelDefinedMetadatumArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateModelDefinedMetadatumArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateModelDefinedMetadatumArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateModelDefinedMetadatumArtifactResponse")
	}
	return
}

// updateModelDefinedMetadatumArtifact implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateModelDefinedMetadatumArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/models/{modelId}/definedMetadata/{metadatumKeyName}/artifact", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateModelDefinedMetadatumArtifactResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/UpdateModelDefinedMetadatumArtifact"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateModelDefinedMetadatumArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateModelDeployment Updates the properties of a model deployment. Some of the properties of `modelDeploymentConfigurationDetails` or `CategoryLogDetails` can also be updated with zero down time
// when the model deployment's lifecycle state is ACTIVE or NEEDS_ATTENTION i.e `instanceShapeName`, `instanceCount` and `modelId`, separately `loadBalancerShape` or `CategoryLogDetails`
// can also be updated independently. All of the fields can be updated when the deployment is in the INACTIVE lifecycle state. Changes will take effect the next time the model
// deployment is activated.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateModelDeployment.go.html to see an example of how to use UpdateModelDeployment API.
// A default retry strategy applies to this operation UpdateModelDeployment()
func (client DataScienceClient) UpdateModelDeployment(ctx context.Context, request UpdateModelDeploymentRequest) (response UpdateModelDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateModelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateModelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateModelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateModelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateModelDeploymentResponse")
	}
	return
}

// updateModelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateModelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/modelDeployments/{modelDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateModelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelDeployment/UpdateModelDeployment"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateModelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateModelProvenance Updates the provenance information for the specified model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateModelProvenance.go.html to see an example of how to use UpdateModelProvenance API.
// A default retry strategy applies to this operation UpdateModelProvenance()
func (client DataScienceClient) UpdateModelProvenance(ctx context.Context, request UpdateModelProvenanceRequest) (response UpdateModelProvenanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateModelProvenance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateModelProvenanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateModelProvenanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateModelProvenanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateModelProvenanceResponse")
	}
	return
}

// updateModelProvenance implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateModelProvenance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/models/{modelId}/provenance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateModelProvenanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Model/UpdateModelProvenance"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateModelProvenance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateModelVersionSet Updates the properties of a model version set. User can update the `description` property.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateModelVersionSet.go.html to see an example of how to use UpdateModelVersionSet API.
// A default retry strategy applies to this operation UpdateModelVersionSet()
func (client DataScienceClient) UpdateModelVersionSet(ctx context.Context, request UpdateModelVersionSetRequest) (response UpdateModelVersionSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateModelVersionSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateModelVersionSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateModelVersionSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateModelVersionSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateModelVersionSetResponse")
	}
	return
}

// updateModelVersionSet implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateModelVersionSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/modelVersionSets/{modelVersionSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateModelVersionSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/ModelVersionSet/UpdateModelVersionSet"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateModelVersionSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNotebookSession Updates the properties of a notebook session. You can update the `displayName`, `freeformTags`, and `definedTags` properties.
// When the notebook session is in the INACTIVE lifecycle state, you can update `notebookSessionConfigurationDetails` and change `shape`, `subnetId`, and `blockStorageSizeInGBs`.
// Changes to the `notebookSessionConfigurationDetails` take effect the next time the `ActivateNotebookSession` action is invoked on the notebook session resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateNotebookSession.go.html to see an example of how to use UpdateNotebookSession API.
// A default retry strategy applies to this operation UpdateNotebookSession()
func (client DataScienceClient) UpdateNotebookSession(ctx context.Context, request UpdateNotebookSessionRequest) (response UpdateNotebookSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNotebookSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNotebookSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNotebookSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNotebookSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNotebookSessionResponse")
	}
	return
}

// updateNotebookSession implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updateNotebookSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/notebookSessions/{notebookSessionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNotebookSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/NotebookSession/UpdateNotebookSession"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateNotebookSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePipeline Updates the Pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdatePipeline.go.html to see an example of how to use UpdatePipeline API.
func (client DataScienceClient) UpdatePipeline(ctx context.Context, request UpdatePipelineRequest) (response UpdatePipelineResponse, err error) {
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
func (client DataScienceClient) updatePipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/pipelines/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Pipeline/UpdatePipeline"
		err = common.PostProcessServiceError(err, "DataScience", "UpdatePipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePipelineRun Updates the PipelineRun.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdatePipelineRun.go.html to see an example of how to use UpdatePipelineRun API.
func (client DataScienceClient) UpdatePipelineRun(ctx context.Context, request UpdatePipelineRunRequest) (response UpdatePipelineRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePipelineRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePipelineRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePipelineRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePipelineRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePipelineRunResponse")
	}
	return
}

// updatePipelineRun implements the OCIOperation interface (enables retrying operations)
func (client DataScienceClient) updatePipelineRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/pipelineRuns/{pipelineRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePipelineRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/PipelineRun/UpdatePipelineRun"
		err = common.PostProcessServiceError(err, "DataScience", "UpdatePipelineRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProject Updates the properties of a project. You can update the `displayName`, `description`, `freeformTags`, and `definedTags` properties.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateProject.go.html to see an example of how to use UpdateProject API.
// A default retry strategy applies to this operation UpdateProject()
func (client DataScienceClient) UpdateProject(ctx context.Context, request UpdateProjectRequest) (response UpdateProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client DataScienceClient) updateProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/projects/{projectId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Project/UpdateProject"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSchedule Updates the Schedule
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/UpdateSchedule.go.html to see an example of how to use UpdateSchedule API.
func (client DataScienceClient) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (response UpdateScheduleResponse, err error) {
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
func (client DataScienceClient) updateSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/schedules/{scheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-science/20190101/Schedule/UpdateSchedule"
		err = common.PostProcessServiceError(err, "DataScience", "UpdateSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
