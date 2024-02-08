// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ApplicationDependencyManagementClient a client for ApplicationDependencyManagement
type ApplicationDependencyManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApplicationDependencyManagementClientWithConfigurationProvider Creates a new default ApplicationDependencyManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApplicationDependencyManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApplicationDependencyManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("adm"); !enabled {
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
	return newApplicationDependencyManagementClientFromBaseClient(baseClient, provider)
}

// NewApplicationDependencyManagementClientWithOboToken Creates a new default ApplicationDependencyManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewApplicationDependencyManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApplicationDependencyManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newApplicationDependencyManagementClientFromBaseClient(baseClient, configProvider)
}

func newApplicationDependencyManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApplicationDependencyManagementClient, err error) {
	// ApplicationDependencyManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ApplicationDependencyManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ApplicationDependencyManagementClient{BaseClient: baseClient}
	client.BasePath = "20220421"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApplicationDependencyManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("adm", "https://adm.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApplicationDependencyManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApplicationDependencyManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateRemediationRecipe Activates the specified Remediation Recipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ActivateRemediationRecipe.go.html to see an example of how to use ActivateRemediationRecipe API.
// A default retry strategy applies to this operation ActivateRemediationRecipe()
func (client ApplicationDependencyManagementClient) ActivateRemediationRecipe(ctx context.Context, request ActivateRemediationRecipeRequest) (response ActivateRemediationRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.activateRemediationRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateRemediationRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateRemediationRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateRemediationRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateRemediationRecipeResponse")
	}
	return
}

// activateRemediationRecipe implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) activateRemediationRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRecipes/{remediationRecipeId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateRemediationRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/ActivateRemediationRecipe"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ActivateRemediationRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelRemediationRun Cancels the specified remediation run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/CancelRemediationRun.go.html to see an example of how to use CancelRemediationRun API.
// A default retry strategy applies to this operation CancelRemediationRun()
func (client ApplicationDependencyManagementClient) CancelRemediationRun(ctx context.Context, request CancelRemediationRunRequest) (response CancelRemediationRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelRemediationRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelRemediationRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelRemediationRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelRemediationRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelRemediationRunResponse")
	}
	return
}

// cancelRemediationRun implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) cancelRemediationRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRuns/{remediationRunId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelRemediationRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/CancelRemediationRun"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "CancelRemediationRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancel work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client ApplicationDependencyManagementClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
func (client ApplicationDependencyManagementClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeKnowledgeBaseCompartment Moves a Knowledge Base from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ChangeKnowledgeBaseCompartment.go.html to see an example of how to use ChangeKnowledgeBaseCompartment API.
// A default retry strategy applies to this operation ChangeKnowledgeBaseCompartment()
func (client ApplicationDependencyManagementClient) ChangeKnowledgeBaseCompartment(ctx context.Context, request ChangeKnowledgeBaseCompartmentRequest) (response ChangeKnowledgeBaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeKnowledgeBaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeKnowledgeBaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeKnowledgeBaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeKnowledgeBaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeKnowledgeBaseCompartmentResponse")
	}
	return
}

// changeKnowledgeBaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) changeKnowledgeBaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/knowledgeBases/{knowledgeBaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeKnowledgeBaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/KnowledgeBase/ChangeKnowledgeBaseCompartment"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ChangeKnowledgeBaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeRemediationRecipeCompartment Moves a Remediation Recipe from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ChangeRemediationRecipeCompartment.go.html to see an example of how to use ChangeRemediationRecipeCompartment API.
// A default retry strategy applies to this operation ChangeRemediationRecipeCompartment()
func (client ApplicationDependencyManagementClient) ChangeRemediationRecipeCompartment(ctx context.Context, request ChangeRemediationRecipeCompartmentRequest) (response ChangeRemediationRecipeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRemediationRecipeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRemediationRecipeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRemediationRecipeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRemediationRecipeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRemediationRecipeCompartmentResponse")
	}
	return
}

// changeRemediationRecipeCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) changeRemediationRecipeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRecipes/{remediationRecipeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRemediationRecipeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/ChangeRemediationRecipeCompartment"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ChangeRemediationRecipeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeRemediationRunCompartment Moves a remediation run from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ChangeRemediationRunCompartment.go.html to see an example of how to use ChangeRemediationRunCompartment API.
// A default retry strategy applies to this operation ChangeRemediationRunCompartment()
func (client ApplicationDependencyManagementClient) ChangeRemediationRunCompartment(ctx context.Context, request ChangeRemediationRunCompartmentRequest) (response ChangeRemediationRunCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRemediationRunCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRemediationRunCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRemediationRunCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRemediationRunCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRemediationRunCompartmentResponse")
	}
	return
}

// changeRemediationRunCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) changeRemediationRunCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRuns/{remediationRunId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRemediationRunCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/ChangeRemediationRunCompartment"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ChangeRemediationRunCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVulnerabilityAuditCompartment Moves a Vulnerability Audit from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ChangeVulnerabilityAuditCompartment.go.html to see an example of how to use ChangeVulnerabilityAuditCompartment API.
// A default retry strategy applies to this operation ChangeVulnerabilityAuditCompartment()
func (client ApplicationDependencyManagementClient) ChangeVulnerabilityAuditCompartment(ctx context.Context, request ChangeVulnerabilityAuditCompartmentRequest) (response ChangeVulnerabilityAuditCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVulnerabilityAuditCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVulnerabilityAuditCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVulnerabilityAuditCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVulnerabilityAuditCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVulnerabilityAuditCompartmentResponse")
	}
	return
}

// changeVulnerabilityAuditCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) changeVulnerabilityAuditCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vulnerabilityAudits/{vulnerabilityAuditId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVulnerabilityAuditCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/ChangeVulnerabilityAuditCompartment"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ChangeVulnerabilityAuditCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateKnowledgeBase Creates a new Knowledge Base.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/CreateKnowledgeBase.go.html to see an example of how to use CreateKnowledgeBase API.
// A default retry strategy applies to this operation CreateKnowledgeBase()
func (client ApplicationDependencyManagementClient) CreateKnowledgeBase(ctx context.Context, request CreateKnowledgeBaseRequest) (response CreateKnowledgeBaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createKnowledgeBase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateKnowledgeBaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateKnowledgeBaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKnowledgeBaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKnowledgeBaseResponse")
	}
	return
}

// createKnowledgeBase implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) createKnowledgeBase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/knowledgeBases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateKnowledgeBaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/KnowledgeBase/CreateKnowledgeBase"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "CreateKnowledgeBase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRemediationRecipe Creates a new Remediation Recipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/CreateRemediationRecipe.go.html to see an example of how to use CreateRemediationRecipe API.
// A default retry strategy applies to this operation CreateRemediationRecipe()
func (client ApplicationDependencyManagementClient) CreateRemediationRecipe(ctx context.Context, request CreateRemediationRecipeRequest) (response CreateRemediationRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRemediationRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRemediationRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRemediationRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRemediationRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRemediationRecipeResponse")
	}
	return
}

// createRemediationRecipe implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) createRemediationRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRemediationRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/CreateRemediationRecipe"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "CreateRemediationRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRemediationRun Creates a new remediation run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/CreateRemediationRun.go.html to see an example of how to use CreateRemediationRun API.
// A default retry strategy applies to this operation CreateRemediationRun()
func (client ApplicationDependencyManagementClient) CreateRemediationRun(ctx context.Context, request CreateRemediationRunRequest) (response CreateRemediationRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRemediationRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRemediationRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRemediationRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRemediationRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRemediationRunResponse")
	}
	return
}

// createRemediationRun implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) createRemediationRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRemediationRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/CreateRemediationRun"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "CreateRemediationRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVulnerabilityAudit Creates a new Vulnerability Audit by providing a tree of Application Dependencies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/CreateVulnerabilityAudit.go.html to see an example of how to use CreateVulnerabilityAudit API.
// A default retry strategy applies to this operation CreateVulnerabilityAudit()
func (client ApplicationDependencyManagementClient) CreateVulnerabilityAudit(ctx context.Context, request CreateVulnerabilityAuditRequest) (response CreateVulnerabilityAuditResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVulnerabilityAudit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVulnerabilityAuditResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVulnerabilityAuditResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVulnerabilityAuditResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVulnerabilityAuditResponse")
	}
	return
}

// createVulnerabilityAudit implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) createVulnerabilityAudit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vulnerabilityAudits", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVulnerabilityAuditResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/CreateVulnerabilityAudit"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "CreateVulnerabilityAudit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateRemediationRecipe Deactivates the specified Remediation Recipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/DeactivateRemediationRecipe.go.html to see an example of how to use DeactivateRemediationRecipe API.
// A default retry strategy applies to this operation DeactivateRemediationRecipe()
func (client ApplicationDependencyManagementClient) DeactivateRemediationRecipe(ctx context.Context, request DeactivateRemediationRecipeRequest) (response DeactivateRemediationRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deactivateRemediationRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateRemediationRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateRemediationRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateRemediationRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateRemediationRecipeResponse")
	}
	return
}

// deactivateRemediationRecipe implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) deactivateRemediationRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/remediationRecipes/{remediationRecipeId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateRemediationRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/DeactivateRemediationRecipe"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "DeactivateRemediationRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteKnowledgeBase Deletes the specified Knowledge Base.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/DeleteKnowledgeBase.go.html to see an example of how to use DeleteKnowledgeBase API.
// A default retry strategy applies to this operation DeleteKnowledgeBase()
func (client ApplicationDependencyManagementClient) DeleteKnowledgeBase(ctx context.Context, request DeleteKnowledgeBaseRequest) (response DeleteKnowledgeBaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteKnowledgeBase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteKnowledgeBaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteKnowledgeBaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteKnowledgeBaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteKnowledgeBaseResponse")
	}
	return
}

// deleteKnowledgeBase implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) deleteKnowledgeBase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/knowledgeBases/{knowledgeBaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteKnowledgeBaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/KnowledgeBase/DeleteKnowledgeBase"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "DeleteKnowledgeBase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRemediationRecipe Deletes the specified Remediation Recipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/DeleteRemediationRecipe.go.html to see an example of how to use DeleteRemediationRecipe API.
// A default retry strategy applies to this operation DeleteRemediationRecipe()
func (client ApplicationDependencyManagementClient) DeleteRemediationRecipe(ctx context.Context, request DeleteRemediationRecipeRequest) (response DeleteRemediationRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRemediationRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRemediationRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRemediationRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRemediationRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRemediationRecipeResponse")
	}
	return
}

// deleteRemediationRecipe implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) deleteRemediationRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/remediationRecipes/{remediationRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRemediationRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/DeleteRemediationRecipe"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "DeleteRemediationRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRemediationRun Deletes the specified remediation run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/DeleteRemediationRun.go.html to see an example of how to use DeleteRemediationRun API.
// A default retry strategy applies to this operation DeleteRemediationRun()
func (client ApplicationDependencyManagementClient) DeleteRemediationRun(ctx context.Context, request DeleteRemediationRunRequest) (response DeleteRemediationRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRemediationRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRemediationRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRemediationRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRemediationRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRemediationRunResponse")
	}
	return
}

// deleteRemediationRun implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) deleteRemediationRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/remediationRuns/{remediationRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRemediationRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/DeleteRemediationRun"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "DeleteRemediationRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVulnerabilityAudit Deletes the specified Vulnerability Audit.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/DeleteVulnerabilityAudit.go.html to see an example of how to use DeleteVulnerabilityAudit API.
// A default retry strategy applies to this operation DeleteVulnerabilityAudit()
func (client ApplicationDependencyManagementClient) DeleteVulnerabilityAudit(ctx context.Context, request DeleteVulnerabilityAuditRequest) (response DeleteVulnerabilityAuditResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVulnerabilityAudit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVulnerabilityAuditResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVulnerabilityAuditResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVulnerabilityAuditResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVulnerabilityAuditResponse")
	}
	return
}

// deleteVulnerabilityAudit implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) deleteVulnerabilityAudit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/vulnerabilityAudits/{vulnerabilityAuditId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVulnerabilityAuditResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/DeleteVulnerabilityAudit"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "DeleteVulnerabilityAudit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetKnowledgeBase Returns the details of the specified Knowledge Base.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/GetKnowledgeBase.go.html to see an example of how to use GetKnowledgeBase API.
// A default retry strategy applies to this operation GetKnowledgeBase()
func (client ApplicationDependencyManagementClient) GetKnowledgeBase(ctx context.Context, request GetKnowledgeBaseRequest) (response GetKnowledgeBaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getKnowledgeBase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetKnowledgeBaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetKnowledgeBaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetKnowledgeBaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetKnowledgeBaseResponse")
	}
	return
}

// getKnowledgeBase implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) getKnowledgeBase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/knowledgeBases/{knowledgeBaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetKnowledgeBaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/KnowledgeBase/GetKnowledgeBase"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "GetKnowledgeBase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRemediationRecipe Returns the details of the specified RemediationRecipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/GetRemediationRecipe.go.html to see an example of how to use GetRemediationRecipe API.
// A default retry strategy applies to this operation GetRemediationRecipe()
func (client ApplicationDependencyManagementClient) GetRemediationRecipe(ctx context.Context, request GetRemediationRecipeRequest) (response GetRemediationRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRemediationRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRemediationRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRemediationRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRemediationRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRemediationRecipeResponse")
	}
	return
}

// getRemediationRecipe implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) getRemediationRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRecipes/{remediationRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRemediationRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/GetRemediationRecipe"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "GetRemediationRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRemediationRun Returns the details of the specified remediation run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/GetRemediationRun.go.html to see an example of how to use GetRemediationRun API.
// A default retry strategy applies to this operation GetRemediationRun()
func (client ApplicationDependencyManagementClient) GetRemediationRun(ctx context.Context, request GetRemediationRunRequest) (response GetRemediationRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRemediationRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRemediationRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRemediationRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRemediationRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRemediationRunResponse")
	}
	return
}

// getRemediationRun implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) getRemediationRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRuns/{remediationRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRemediationRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/GetRemediationRun"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "GetRemediationRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStage Returns the details of the specified Remediation Run Stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/GetStage.go.html to see an example of how to use GetStage API.
// A default retry strategy applies to this operation GetStage()
func (client ApplicationDependencyManagementClient) GetStage(ctx context.Context, request GetStageRequest) (response GetStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStageResponse")
	}
	return
}

// getStage implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) getStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRuns/{remediationRunId}/stages/{stageType}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRunStage/GetStage"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "GetStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &remediationrunstage{})
	return response, err
}

// GetVulnerabilityAudit Returns the details of the specified Vulnerability Audit.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/GetVulnerabilityAudit.go.html to see an example of how to use GetVulnerabilityAudit API.
// A default retry strategy applies to this operation GetVulnerabilityAudit()
func (client ApplicationDependencyManagementClient) GetVulnerabilityAudit(ctx context.Context, request GetVulnerabilityAuditRequest) (response GetVulnerabilityAuditResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVulnerabilityAudit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVulnerabilityAuditResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVulnerabilityAuditResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVulnerabilityAuditResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVulnerabilityAuditResponse")
	}
	return
}

// getVulnerabilityAudit implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) getVulnerabilityAudit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerabilityAudits/{vulnerabilityAuditId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVulnerabilityAuditResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/GetVulnerabilityAudit"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "GetVulnerabilityAudit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ApplicationDependencyManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client ApplicationDependencyManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplicationDependencyRecommendations Returns a list of application dependency with their associated recommendations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListApplicationDependencyRecommendations.go.html to see an example of how to use ListApplicationDependencyRecommendations API.
// A default retry strategy applies to this operation ListApplicationDependencyRecommendations()
func (client ApplicationDependencyManagementClient) ListApplicationDependencyRecommendations(ctx context.Context, request ListApplicationDependencyRecommendationsRequest) (response ListApplicationDependencyRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplicationDependencyRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicationDependencyRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicationDependencyRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationDependencyRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationDependencyRecommendationsResponse")
	}
	return
}

// listApplicationDependencyRecommendations implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listApplicationDependencyRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRuns/{remediationRunId}/applicationDependencyRecommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationDependencyRecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/ListApplicationDependencyRecommendations"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListApplicationDependencyRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplicationDependencyVulnerabilities Returns a list of Application Dependencies with their associated vulnerabilities.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListApplicationDependencyVulnerabilities.go.html to see an example of how to use ListApplicationDependencyVulnerabilities API.
// A default retry strategy applies to this operation ListApplicationDependencyVulnerabilities()
func (client ApplicationDependencyManagementClient) ListApplicationDependencyVulnerabilities(ctx context.Context, request ListApplicationDependencyVulnerabilitiesRequest) (response ListApplicationDependencyVulnerabilitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplicationDependencyVulnerabilities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicationDependencyVulnerabilitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicationDependencyVulnerabilitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationDependencyVulnerabilitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationDependencyVulnerabilitiesResponse")
	}
	return
}

// listApplicationDependencyVulnerabilities implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listApplicationDependencyVulnerabilities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerabilityAudits/{vulnerabilityAuditId}/applicationDependencyVulnerabilities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationDependencyVulnerabilitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/ListApplicationDependencyVulnerabilities"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListApplicationDependencyVulnerabilities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListKnowledgeBases Returns a list of KnowledgeBases based on the specified query parameters.
// At least id or compartmentId query parameter must be provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListKnowledgeBases.go.html to see an example of how to use ListKnowledgeBases API.
// A default retry strategy applies to this operation ListKnowledgeBases()
func (client ApplicationDependencyManagementClient) ListKnowledgeBases(ctx context.Context, request ListKnowledgeBasesRequest) (response ListKnowledgeBasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listKnowledgeBases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListKnowledgeBasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListKnowledgeBasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListKnowledgeBasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKnowledgeBasesResponse")
	}
	return
}

// listKnowledgeBases implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listKnowledgeBases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/knowledgeBases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListKnowledgeBasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/KnowledgeBase/ListKnowledgeBases"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListKnowledgeBases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRemediationRecipes Returns a list of Remediation Recipes based on the specified query parameters.
// The query parameters `compartmentId` or `id` must be provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListRemediationRecipes.go.html to see an example of how to use ListRemediationRecipes API.
// A default retry strategy applies to this operation ListRemediationRecipes()
func (client ApplicationDependencyManagementClient) ListRemediationRecipes(ctx context.Context, request ListRemediationRecipesRequest) (response ListRemediationRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRemediationRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRemediationRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRemediationRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRemediationRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRemediationRecipesResponse")
	}
	return
}

// listRemediationRecipes implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listRemediationRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRemediationRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/ListRemediationRecipes"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListRemediationRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRemediationRuns Returns a list of remediation runs contained by a compartment.
// The query parameter `compartmentId` is required unless the query parameter `id` is specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListRemediationRuns.go.html to see an example of how to use ListRemediationRuns API.
// A default retry strategy applies to this operation ListRemediationRuns()
func (client ApplicationDependencyManagementClient) ListRemediationRuns(ctx context.Context, request ListRemediationRunsRequest) (response ListRemediationRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRemediationRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRemediationRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRemediationRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRemediationRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRemediationRunsResponse")
	}
	return
}

// listRemediationRuns implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listRemediationRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRemediationRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/ListRemediationRuns"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListRemediationRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStages Returns a list of Remediation Run Stages based on the specified query parameters and Remediation Run identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListStages.go.html to see an example of how to use ListStages API.
// A default retry strategy applies to this operation ListStages()
func (client ApplicationDependencyManagementClient) ListStages(ctx context.Context, request ListStagesRequest) (response ListStagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStagesResponse")
	}
	return
}

// listStages implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listStages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/remediationRuns/{remediationRunId}/stages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRunStage/ListStages"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListStages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVulnerabilityAudits Returns a list of Vulnerability Audits based on the specified query parameters.
// At least one of id, compartmentId query parameter must be provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListVulnerabilityAudits.go.html to see an example of how to use ListVulnerabilityAudits API.
// A default retry strategy applies to this operation ListVulnerabilityAudits()
func (client ApplicationDependencyManagementClient) ListVulnerabilityAudits(ctx context.Context, request ListVulnerabilityAuditsRequest) (response ListVulnerabilityAuditsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVulnerabilityAudits, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVulnerabilityAuditsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVulnerabilityAuditsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVulnerabilityAuditsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVulnerabilityAuditsResponse")
	}
	return
}

// listVulnerabilityAudits implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) listVulnerabilityAudits(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerabilityAudits", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVulnerabilityAuditsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/ListVulnerabilityAudits"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListVulnerabilityAudits", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ApplicationDependencyManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client ApplicationDependencyManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ApplicationDependencyManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client ApplicationDependencyManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ApplicationDependencyManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client ApplicationDependencyManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateKnowledgeBase Updates one or more attributes of the specified Knowledge Base.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/UpdateKnowledgeBase.go.html to see an example of how to use UpdateKnowledgeBase API.
// A default retry strategy applies to this operation UpdateKnowledgeBase()
func (client ApplicationDependencyManagementClient) UpdateKnowledgeBase(ctx context.Context, request UpdateKnowledgeBaseRequest) (response UpdateKnowledgeBaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateKnowledgeBase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateKnowledgeBaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateKnowledgeBaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateKnowledgeBaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateKnowledgeBaseResponse")
	}
	return
}

// updateKnowledgeBase implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) updateKnowledgeBase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/knowledgeBases/{knowledgeBaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateKnowledgeBaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/KnowledgeBase/UpdateKnowledgeBase"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "UpdateKnowledgeBase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRemediationRecipe Updates one or more attributes of the specified Remediation Recipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/UpdateRemediationRecipe.go.html to see an example of how to use UpdateRemediationRecipe API.
// A default retry strategy applies to this operation UpdateRemediationRecipe()
func (client ApplicationDependencyManagementClient) UpdateRemediationRecipe(ctx context.Context, request UpdateRemediationRecipeRequest) (response UpdateRemediationRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRemediationRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRemediationRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRemediationRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRemediationRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRemediationRecipeResponse")
	}
	return
}

// updateRemediationRecipe implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) updateRemediationRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/remediationRecipes/{remediationRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRemediationRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRecipe/UpdateRemediationRecipe"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "UpdateRemediationRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRemediationRun Updates by identifier one or more attributes of the specified remediation run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/UpdateRemediationRun.go.html to see an example of how to use UpdateRemediationRun API.
// A default retry strategy applies to this operation UpdateRemediationRun()
func (client ApplicationDependencyManagementClient) UpdateRemediationRun(ctx context.Context, request UpdateRemediationRunRequest) (response UpdateRemediationRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRemediationRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRemediationRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRemediationRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRemediationRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRemediationRunResponse")
	}
	return
}

// updateRemediationRun implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) updateRemediationRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/remediationRuns/{remediationRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRemediationRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/RemediationRun/UpdateRemediationRun"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "UpdateRemediationRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVulnerabilityAudit Updates one or more attributes of the specified Vulnerability Audit.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/UpdateVulnerabilityAudit.go.html to see an example of how to use UpdateVulnerabilityAudit API.
// A default retry strategy applies to this operation UpdateVulnerabilityAudit()
func (client ApplicationDependencyManagementClient) UpdateVulnerabilityAudit(ctx context.Context, request UpdateVulnerabilityAuditRequest) (response UpdateVulnerabilityAuditResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVulnerabilityAudit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVulnerabilityAuditResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVulnerabilityAuditResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVulnerabilityAuditResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVulnerabilityAuditResponse")
	}
	return
}

// updateVulnerabilityAudit implements the OCIOperation interface (enables retrying operations)
func (client ApplicationDependencyManagementClient) updateVulnerabilityAudit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/vulnerabilityAudits/{vulnerabilityAuditId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVulnerabilityAuditResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/VulnerabilityAudit/UpdateVulnerabilityAudit"
		err = common.PostProcessServiceError(err, "ApplicationDependencyManagement", "UpdateVulnerabilityAudit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
