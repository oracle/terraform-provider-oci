// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// JavaManagementServiceClient a client for JavaManagementService
type JavaManagementServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewJavaManagementServiceClientWithConfigurationProvider Creates a new default JavaManagementService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewJavaManagementServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client JavaManagementServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("jms"); !enabled {
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
	return newJavaManagementServiceClientFromBaseClient(baseClient, provider)
}

// NewJavaManagementServiceClientWithOboToken Creates a new default JavaManagementService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewJavaManagementServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client JavaManagementServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newJavaManagementServiceClientFromBaseClient(baseClient, configProvider)
}

func newJavaManagementServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client JavaManagementServiceClient, err error) {
	// JavaManagementService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("JavaManagementService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = JavaManagementServiceClient{BaseClient: baseClient}
	client.BasePath = "20210610"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *JavaManagementServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("jms", "https://javamanagement.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *JavaManagementServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *JavaManagementServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddFleetInstallationSites Add Java installation sites in a Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/AddFleetInstallationSites.go.html to see an example of how to use AddFleetInstallationSites API.
// A default retry strategy applies to this operation AddFleetInstallationSites()
func (client JavaManagementServiceClient) AddFleetInstallationSites(ctx context.Context, request AddFleetInstallationSitesRequest) (response AddFleetInstallationSitesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addFleetInstallationSites, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddFleetInstallationSitesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddFleetInstallationSitesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddFleetInstallationSitesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddFleetInstallationSitesResponse")
	}
	return
}

// addFleetInstallationSites implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) addFleetInstallationSites(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/addInstallationSites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddFleetInstallationSitesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationSiteSummary/AddFleetInstallationSites"
		err = common.PostProcessServiceError(err, "JavaManagementService", "AddFleetInstallationSites", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Deletes the work request specified by an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client JavaManagementServiceClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
func (client JavaManagementServiceClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFleetCompartment Move a specified Fleet into the compartment identified in the POST form. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ChangeFleetCompartment.go.html to see an example of how to use ChangeFleetCompartment API.
// A default retry strategy applies to this operation ChangeFleetCompartment()
func (client JavaManagementServiceClient) ChangeFleetCompartment(ctx context.Context, request ChangeFleetCompartmentRequest) (response ChangeFleetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFleetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFleetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFleetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFleetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFleetCompartmentResponse")
	}
	return
}

// changeFleetCompartment implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) changeFleetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFleetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/ChangeFleetCompartment"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ChangeFleetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBlocklist Add a new record to the fleet blocklist.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/CreateBlocklist.go.html to see an example of how to use CreateBlocklist API.
// A default retry strategy applies to this operation CreateBlocklist()
func (client JavaManagementServiceClient) CreateBlocklist(ctx context.Context, request CreateBlocklistRequest) (response CreateBlocklistResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBlocklist, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBlocklistResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBlocklistResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBlocklistResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBlocklistResponse")
	}
	return
}

// createBlocklist implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) createBlocklist(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/blocklists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBlocklistResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Blocklist/CreateBlocklist"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CreateBlocklist", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDrsFile Request to perform validaition of the DRS file and create the file to the Object Storage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/CreateDrsFile.go.html to see an example of how to use CreateDrsFile API.
// A default retry strategy applies to this operation CreateDrsFile()
func (client JavaManagementServiceClient) CreateDrsFile(ctx context.Context, request CreateDrsFileRequest) (response CreateDrsFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDrsFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDrsFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDrsFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDrsFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDrsFileResponse")
	}
	return
}

// createDrsFile implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) createDrsFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/drsFiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDrsFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/CreateDrsFile"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CreateDrsFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFleet Create a new Fleet using the information provided.
// `inventoryLog` is now a required parameter for CreateFleet API.
// Update existing applications using this API
// before July 15, 2022 to ensure the applications continue to work.
// See the Service Change Notice (https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#JMS) for more details.
// Migrate existing fleets using the `UpdateFleet` API to set the `inventoryLog` parameter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/CreateFleet.go.html to see an example of how to use CreateFleet API.
// A default retry strategy applies to this operation CreateFleet()
func (client JavaManagementServiceClient) CreateFleet(ctx context.Context, request CreateFleetRequest) (response CreateFleetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFleetResponse")
	}
	return
}

// createFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) createFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/CreateFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CreateFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBlocklist Deletes the blocklist record specified by an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DeleteBlocklist.go.html to see an example of how to use DeleteBlocklist API.
// A default retry strategy applies to this operation DeleteBlocklist()
func (client JavaManagementServiceClient) DeleteBlocklist(ctx context.Context, request DeleteBlocklistRequest) (response DeleteBlocklistResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBlocklist, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBlocklistResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBlocklistResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBlocklistResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBlocklistResponse")
	}
	return
}

// deleteBlocklist implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteBlocklist(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}/blocklists/{blocklistKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBlocklistResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Blocklist/DeleteBlocklist"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteBlocklist", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCryptoAnalysisResult Deletes the metadata for the result of a Crypto event analysis. The actual report shall remain in the object storage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DeleteCryptoAnalysisResult.go.html to see an example of how to use DeleteCryptoAnalysisResult API.
// A default retry strategy applies to this operation DeleteCryptoAnalysisResult()
func (client JavaManagementServiceClient) DeleteCryptoAnalysisResult(ctx context.Context, request DeleteCryptoAnalysisResultRequest) (response DeleteCryptoAnalysisResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCryptoAnalysisResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCryptoAnalysisResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCryptoAnalysisResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCryptoAnalysisResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCryptoAnalysisResultResponse")
	}
	return
}

// deleteCryptoAnalysisResult implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteCryptoAnalysisResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}/cryptoAnalysisResults/{cryptoAnalysisResultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCryptoAnalysisResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/CryptoAnalysisResult/DeleteCryptoAnalysisResult"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteCryptoAnalysisResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDrsFile Request to delete the DRS file from the Object Storage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DeleteDrsFile.go.html to see an example of how to use DeleteDrsFile API.
// A default retry strategy applies to this operation DeleteDrsFile()
func (client JavaManagementServiceClient) DeleteDrsFile(ctx context.Context, request DeleteDrsFileRequest) (response DeleteDrsFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDrsFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDrsFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDrsFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDrsFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDrsFileResponse")
	}
	return
}

// deleteDrsFile implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteDrsFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}/drsFiles/{drsFileKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDrsFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/DrsFile/DeleteDrsFile"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteDrsFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFleet Deletes the Fleet specified by an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DeleteFleet.go.html to see an example of how to use DeleteFleet API.
// A default retry strategy applies to this operation DeleteFleet()
func (client JavaManagementServiceClient) DeleteFleet(ctx context.Context, request DeleteFleetRequest) (response DeleteFleetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFleetResponse")
	}
	return
}

// deleteFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/DeleteFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJavaMigrationAnalysisResult Delete the Java migration analysis result. The actual report will remain in the Object Storage bucket.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DeleteJavaMigrationAnalysisResult.go.html to see an example of how to use DeleteJavaMigrationAnalysisResult API.
// A default retry strategy applies to this operation DeleteJavaMigrationAnalysisResult()
func (client JavaManagementServiceClient) DeleteJavaMigrationAnalysisResult(ctx context.Context, request DeleteJavaMigrationAnalysisResultRequest) (response DeleteJavaMigrationAnalysisResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJavaMigrationAnalysisResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJavaMigrationAnalysisResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJavaMigrationAnalysisResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJavaMigrationAnalysisResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJavaMigrationAnalysisResultResponse")
	}
	return
}

// deleteJavaMigrationAnalysisResult implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteJavaMigrationAnalysisResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}/javaMigrationAnalysisResults/{javaMigrationAnalysisResultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJavaMigrationAnalysisResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaMigrationAnalysisResult/DeleteJavaMigrationAnalysisResult"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteJavaMigrationAnalysisResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePerformanceTuningAnalysisResult Deletes only the metadata of the Performance Tuning Analysis result, but the file remains in the object storage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DeletePerformanceTuningAnalysisResult.go.html to see an example of how to use DeletePerformanceTuningAnalysisResult API.
// A default retry strategy applies to this operation DeletePerformanceTuningAnalysisResult()
func (client JavaManagementServiceClient) DeletePerformanceTuningAnalysisResult(ctx context.Context, request DeletePerformanceTuningAnalysisResultRequest) (response DeletePerformanceTuningAnalysisResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePerformanceTuningAnalysisResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePerformanceTuningAnalysisResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePerformanceTuningAnalysisResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePerformanceTuningAnalysisResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePerformanceTuningAnalysisResultResponse")
	}
	return
}

// deletePerformanceTuningAnalysisResult implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deletePerformanceTuningAnalysisResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}/performanceTuningAnalysisResults/{performanceTuningAnalysisResultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePerformanceTuningAnalysisResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/PerformanceTuningAnalysisResult/DeletePerformanceTuningAnalysisResult"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeletePerformanceTuningAnalysisResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableDrs Request to disable the DRS in the selected target in the Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/DisableDrs.go.html to see an example of how to use DisableDrs API.
// A default retry strategy applies to this operation DisableDrs()
func (client JavaManagementServiceClient) DisableDrs(ctx context.Context, request DisableDrsRequest) (response DisableDrsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableDrs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableDrsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableDrsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableDrsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableDrsResponse")
	}
	return
}

// disableDrs implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) disableDrs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/disableDrs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableDrsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/DisableDrs"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DisableDrs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDrs Request to enable the DRS in the selected target in the Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/EnableDrs.go.html to see an example of how to use EnableDrs API.
// A default retry strategy applies to this operation EnableDrs()
func (client JavaManagementServiceClient) EnableDrs(ctx context.Context, request EnableDrsRequest) (response EnableDrsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableDrs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableDrsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableDrsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableDrsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableDrsResponse")
	}
	return
}

// enableDrs implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) enableDrs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/enableDrs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableDrsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/EnableDrs"
		err = common.PostProcessServiceError(err, "JavaManagementService", "EnableDrs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateAgentDeployScript Generates Agent Deploy Script for Fleet using the information provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GenerateAgentDeployScript.go.html to see an example of how to use GenerateAgentDeployScript API.
// A default retry strategy applies to this operation GenerateAgentDeployScript()
func (client JavaManagementServiceClient) GenerateAgentDeployScript(ctx context.Context, request GenerateAgentDeployScriptRequest) (response GenerateAgentDeployScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateAgentDeployScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateAgentDeployScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateAgentDeployScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateAgentDeployScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateAgentDeployScriptResponse")
	}
	return
}

// generateAgentDeployScript implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) generateAgentDeployScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/generateAgentDeployScript", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateAgentDeployScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/GenerateAgentDeployScript"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GenerateAgentDeployScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCryptoAnalysisResult Retrieve the metadata for the result of a Crypto event analysis.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetCryptoAnalysisResult.go.html to see an example of how to use GetCryptoAnalysisResult API.
// A default retry strategy applies to this operation GetCryptoAnalysisResult()
func (client JavaManagementServiceClient) GetCryptoAnalysisResult(ctx context.Context, request GetCryptoAnalysisResultRequest) (response GetCryptoAnalysisResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCryptoAnalysisResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCryptoAnalysisResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCryptoAnalysisResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCryptoAnalysisResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCryptoAnalysisResultResponse")
	}
	return
}

// getCryptoAnalysisResult implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getCryptoAnalysisResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/cryptoAnalysisResults/{cryptoAnalysisResultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCryptoAnalysisResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/CryptoAnalysisResult/GetCryptoAnalysisResult"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetCryptoAnalysisResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDrsFile Get the detail about the created DRS file in the Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetDrsFile.go.html to see an example of how to use GetDrsFile API.
// A default retry strategy applies to this operation GetDrsFile()
func (client JavaManagementServiceClient) GetDrsFile(ctx context.Context, request GetDrsFileRequest) (response GetDrsFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDrsFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDrsFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDrsFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDrsFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDrsFileResponse")
	}
	return
}

// getDrsFile implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getDrsFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/drsFiles/{drsFileKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDrsFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/DrsFile/GetDrsFile"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetDrsFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExportSetting Returns export setting for the specified Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetExportSetting.go.html to see an example of how to use GetExportSetting API.
// A default retry strategy applies to this operation GetExportSetting()
func (client JavaManagementServiceClient) GetExportSetting(ctx context.Context, request GetExportSettingRequest) (response GetExportSettingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExportSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExportSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExportSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExportSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExportSettingResponse")
	}
	return
}

// getExportSetting implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getExportSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/exportSetting", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExportSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ExportSetting/GetExportSetting"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetExportSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExportStatus Returns last export status for the specified Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetExportStatus.go.html to see an example of how to use GetExportStatus API.
// A default retry strategy applies to this operation GetExportStatus()
func (client JavaManagementServiceClient) GetExportStatus(ctx context.Context, request GetExportStatusRequest) (response GetExportStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExportStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExportStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExportStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExportStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExportStatusResponse")
	}
	return
}

// getExportStatus implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getExportStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/exportStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExportStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ExportStatus/GetExportStatus"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetExportStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFleet Retrieve a Fleet with the specified identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetFleet.go.html to see an example of how to use GetFleet API.
// A default retry strategy applies to this operation GetFleet()
func (client JavaManagementServiceClient) GetFleet(ctx context.Context, request GetFleetRequest) (response GetFleetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFleetResponse")
	}
	return
}

// getFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/GetFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFleetAdvancedFeatureConfiguration Returns Fleet level advanced feature configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetFleetAdvancedFeatureConfiguration.go.html to see an example of how to use GetFleetAdvancedFeatureConfiguration API.
// A default retry strategy applies to this operation GetFleetAdvancedFeatureConfiguration()
func (client JavaManagementServiceClient) GetFleetAdvancedFeatureConfiguration(ctx context.Context, request GetFleetAdvancedFeatureConfigurationRequest) (response GetFleetAdvancedFeatureConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFleetAdvancedFeatureConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFleetAdvancedFeatureConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFleetAdvancedFeatureConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFleetAdvancedFeatureConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFleetAdvancedFeatureConfigurationResponse")
	}
	return
}

// getFleetAdvancedFeatureConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getFleetAdvancedFeatureConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/advancedFeatureConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFleetAdvancedFeatureConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetAdvancedFeatureConfiguration/GetFleetAdvancedFeatureConfiguration"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetFleetAdvancedFeatureConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFleetAgentConfiguration Retrieve a Fleet Agent Configuration for the specified Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetFleetAgentConfiguration.go.html to see an example of how to use GetFleetAgentConfiguration API.
// A default retry strategy applies to this operation GetFleetAgentConfiguration()
func (client JavaManagementServiceClient) GetFleetAgentConfiguration(ctx context.Context, request GetFleetAgentConfigurationRequest) (response GetFleetAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFleetAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFleetAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFleetAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFleetAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFleetAgentConfigurationResponse")
	}
	return
}

// getFleetAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getFleetAgentConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/agentConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFleetAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetAgentConfiguration/GetFleetAgentConfiguration"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetFleetAgentConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaFamily Returns metadata associated with a specific Java release family.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetJavaFamily.go.html to see an example of how to use GetJavaFamily API.
// A default retry strategy applies to this operation GetJavaFamily()
func (client JavaManagementServiceClient) GetJavaFamily(ctx context.Context, request GetJavaFamilyRequest) (response GetJavaFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaFamilyResponse")
	}
	return
}

// getJavaFamily implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getJavaFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/javaFamilies/{familyVersion}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaFamily/GetJavaFamily"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetJavaFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaMigrationAnalysisResult Retrieve Java Migration Analysis result.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetJavaMigrationAnalysisResult.go.html to see an example of how to use GetJavaMigrationAnalysisResult API.
// A default retry strategy applies to this operation GetJavaMigrationAnalysisResult()
func (client JavaManagementServiceClient) GetJavaMigrationAnalysisResult(ctx context.Context, request GetJavaMigrationAnalysisResultRequest) (response GetJavaMigrationAnalysisResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaMigrationAnalysisResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaMigrationAnalysisResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaMigrationAnalysisResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaMigrationAnalysisResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaMigrationAnalysisResultResponse")
	}
	return
}

// getJavaMigrationAnalysisResult implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getJavaMigrationAnalysisResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/javaMigrationAnalysisResults/{javaMigrationAnalysisResultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaMigrationAnalysisResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaMigrationAnalysisResult/GetJavaMigrationAnalysisResult"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetJavaMigrationAnalysisResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaRelease Returns detail of a Java release.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetJavaRelease.go.html to see an example of how to use GetJavaRelease API.
// A default retry strategy applies to this operation GetJavaRelease()
func (client JavaManagementServiceClient) GetJavaRelease(ctx context.Context, request GetJavaReleaseRequest) (response GetJavaReleaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaRelease, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaReleaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaReleaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaReleaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaReleaseResponse")
	}
	return
}

// getJavaRelease implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getJavaRelease(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/javaReleases/{releaseVersion}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaReleaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaRelease/GetJavaRelease"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetJavaRelease", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPerformanceTuningAnalysisResult Retrieve metadata of the Performance Tuning Analysis result.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetPerformanceTuningAnalysisResult.go.html to see an example of how to use GetPerformanceTuningAnalysisResult API.
// A default retry strategy applies to this operation GetPerformanceTuningAnalysisResult()
func (client JavaManagementServiceClient) GetPerformanceTuningAnalysisResult(ctx context.Context, request GetPerformanceTuningAnalysisResultRequest) (response GetPerformanceTuningAnalysisResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPerformanceTuningAnalysisResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPerformanceTuningAnalysisResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPerformanceTuningAnalysisResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPerformanceTuningAnalysisResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPerformanceTuningAnalysisResultResponse")
	}
	return
}

// getPerformanceTuningAnalysisResult implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getPerformanceTuningAnalysisResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/performanceTuningAnalysisResults/{performanceTuningAnalysisResultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPerformanceTuningAnalysisResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/PerformanceTuningAnalysisResult/GetPerformanceTuningAnalysisResult"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetPerformanceTuningAnalysisResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Retrieve the details of a work request with the specified ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client JavaManagementServiceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client JavaManagementServiceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAnnouncements Return a list of AnnouncementSummary items
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListAnnouncements.go.html to see an example of how to use ListAnnouncements API.
// A default retry strategy applies to this operation ListAnnouncements()
func (client JavaManagementServiceClient) ListAnnouncements(ctx context.Context, request ListAnnouncementsRequest) (response ListAnnouncementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnnouncements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnnouncementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnnouncementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnnouncementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnnouncementsResponse")
	}
	return
}

// listAnnouncements implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listAnnouncements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/announcements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAnnouncementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/AnnouncementCollection/ListAnnouncements"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListAnnouncements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBlocklists Returns a list of blocklist entities contained by a fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListBlocklists.go.html to see an example of how to use ListBlocklists API.
// A default retry strategy applies to this operation ListBlocklists()
func (client JavaManagementServiceClient) ListBlocklists(ctx context.Context, request ListBlocklistsRequest) (response ListBlocklistsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlocklists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlocklistsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlocklistsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlocklistsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlocklistsResponse")
	}
	return
}

// listBlocklists implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listBlocklists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/blocklists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBlocklistsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Blocklist/ListBlocklists"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListBlocklists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCryptoAnalysisResults Lists the results of a Crypto event analysis.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListCryptoAnalysisResults.go.html to see an example of how to use ListCryptoAnalysisResults API.
// A default retry strategy applies to this operation ListCryptoAnalysisResults()
func (client JavaManagementServiceClient) ListCryptoAnalysisResults(ctx context.Context, request ListCryptoAnalysisResultsRequest) (response ListCryptoAnalysisResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCryptoAnalysisResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCryptoAnalysisResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCryptoAnalysisResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCryptoAnalysisResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCryptoAnalysisResultsResponse")
	}
	return
}

// listCryptoAnalysisResults implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listCryptoAnalysisResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/cryptoAnalysisResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCryptoAnalysisResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/CryptoAnalysisResult/ListCryptoAnalysisResults"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListCryptoAnalysisResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDrsFiles List the details about the created DRS files in the Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListDrsFiles.go.html to see an example of how to use ListDrsFiles API.
// A default retry strategy applies to this operation ListDrsFiles()
func (client JavaManagementServiceClient) ListDrsFiles(ctx context.Context, request ListDrsFilesRequest) (response ListDrsFilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDrsFiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDrsFilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDrsFilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDrsFilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDrsFilesResponse")
	}
	return
}

// listDrsFiles implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listDrsFiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/drsFiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDrsFilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/DrsFileCollection/ListDrsFiles"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListDrsFiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFleetDiagnoses List potential diagnoses that would put a fleet into FAILED or NEEDS_ATTENTION lifecycle state.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListFleetDiagnoses.go.html to see an example of how to use ListFleetDiagnoses API.
// A default retry strategy applies to this operation ListFleetDiagnoses()
func (client JavaManagementServiceClient) ListFleetDiagnoses(ctx context.Context, request ListFleetDiagnosesRequest) (response ListFleetDiagnosesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFleetDiagnoses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFleetDiagnosesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFleetDiagnosesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFleetDiagnosesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFleetDiagnosesResponse")
	}
	return
}

// listFleetDiagnoses implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listFleetDiagnoses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/diagnoses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFleetDiagnosesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetDiagnosisSummary/ListFleetDiagnoses"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListFleetDiagnoses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFleets Returns a list of all the Fleets contained by a compartment. The query parameter `compartmentId`
// is required unless the query parameter `id` is specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListFleets.go.html to see an example of how to use ListFleets API.
// A default retry strategy applies to this operation ListFleets()
func (client JavaManagementServiceClient) ListFleets(ctx context.Context, request ListFleetsRequest) (response ListFleetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFleets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFleetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFleetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFleetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFleetsResponse")
	}
	return
}

// listFleets implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listFleets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFleetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/ListFleets"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListFleets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstallationSites List Java installation sites in a Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListInstallationSites.go.html to see an example of how to use ListInstallationSites API.
// A default retry strategy applies to this operation ListInstallationSites()
func (client JavaManagementServiceClient) ListInstallationSites(ctx context.Context, request ListInstallationSitesRequest) (response ListInstallationSitesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstallationSites, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstallationSitesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstallationSitesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstallationSitesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstallationSitesResponse")
	}
	return
}

// listInstallationSites implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listInstallationSites(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/installationSites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstallationSitesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationSiteSummary/ListInstallationSites"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListInstallationSites", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaFamilies Returns a list of the Java release family information.
// A Java release family is typically a major version in the Java version identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJavaFamilies.go.html to see an example of how to use ListJavaFamilies API.
// A default retry strategy applies to this operation ListJavaFamilies()
func (client JavaManagementServiceClient) ListJavaFamilies(ctx context.Context, request ListJavaFamiliesRequest) (response ListJavaFamiliesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaFamilies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaFamiliesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaFamiliesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaFamiliesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaFamiliesResponse")
	}
	return
}

// listJavaFamilies implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listJavaFamilies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/javaFamilies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaFamiliesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaFamily/ListJavaFamilies"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListJavaFamilies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaMigrationAnalysisResults Lists the results of a Java migration analysis.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJavaMigrationAnalysisResults.go.html to see an example of how to use ListJavaMigrationAnalysisResults API.
// A default retry strategy applies to this operation ListJavaMigrationAnalysisResults()
func (client JavaManagementServiceClient) ListJavaMigrationAnalysisResults(ctx context.Context, request ListJavaMigrationAnalysisResultsRequest) (response ListJavaMigrationAnalysisResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaMigrationAnalysisResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaMigrationAnalysisResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaMigrationAnalysisResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaMigrationAnalysisResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaMigrationAnalysisResultsResponse")
	}
	return
}

// listJavaMigrationAnalysisResults implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listJavaMigrationAnalysisResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/javaMigrationAnalysisResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaMigrationAnalysisResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaMigrationAnalysisResult/ListJavaMigrationAnalysisResults"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListJavaMigrationAnalysisResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaReleases Returns a list of Java releases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJavaReleases.go.html to see an example of how to use ListJavaReleases API.
// A default retry strategy applies to this operation ListJavaReleases()
func (client JavaManagementServiceClient) ListJavaReleases(ctx context.Context, request ListJavaReleasesRequest) (response ListJavaReleasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaReleases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaReleasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaReleasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaReleasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaReleasesResponse")
	}
	return
}

// listJavaReleases implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listJavaReleases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/javaReleases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaReleasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaRelease/ListJavaReleases"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListJavaReleases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJreUsage List Java Runtime usage in a specified host filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJreUsage.go.html to see an example of how to use ListJreUsage API.
// A default retry strategy applies to this operation ListJreUsage()
func (client JavaManagementServiceClient) ListJreUsage(ctx context.Context, request ListJreUsageRequest) (response ListJreUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJreUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJreUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJreUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJreUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJreUsageResponse")
	}
	return
}

// listJreUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listJreUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listJreUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJreUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JreUsage/ListJreUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListJreUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPerformanceTuningAnalysisResults List Performance Tuning Analysis results.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListPerformanceTuningAnalysisResults.go.html to see an example of how to use ListPerformanceTuningAnalysisResults API.
// A default retry strategy applies to this operation ListPerformanceTuningAnalysisResults()
func (client JavaManagementServiceClient) ListPerformanceTuningAnalysisResults(ctx context.Context, request ListPerformanceTuningAnalysisResultsRequest) (response ListPerformanceTuningAnalysisResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPerformanceTuningAnalysisResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPerformanceTuningAnalysisResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPerformanceTuningAnalysisResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPerformanceTuningAnalysisResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPerformanceTuningAnalysisResultsResponse")
	}
	return
}

// listPerformanceTuningAnalysisResults implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listPerformanceTuningAnalysisResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/performanceTuningAnalysisResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPerformanceTuningAnalysisResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/PerformanceTuningAnalysisResult/ListPerformanceTuningAnalysisResults"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListPerformanceTuningAnalysisResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkItems Retrieve a paginated list of work items for a specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListWorkItems.go.html to see an example of how to use ListWorkItems API.
// A default retry strategy applies to this operation ListWorkItems()
func (client JavaManagementServiceClient) ListWorkItems(ctx context.Context, request ListWorkItemsRequest) (response ListWorkItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkItemsResponse")
	}
	return
}

// listWorkItems implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listWorkItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/workItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkItemSummary/ListWorkItems"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Retrieve a (paginated) list of errors for a specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client JavaManagementServiceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client JavaManagementServiceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Retrieve a paginated list of logs for a specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client JavaManagementServiceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client JavaManagementServiceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests List the work requests in a compartment. The query parameter `compartmentId` is required unless the query parameter `id` or `fleetId` is specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client JavaManagementServiceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client JavaManagementServiceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveFleetInstallationSites Remove Java installation sites in a Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/RemoveFleetInstallationSites.go.html to see an example of how to use RemoveFleetInstallationSites API.
// A default retry strategy applies to this operation RemoveFleetInstallationSites()
func (client JavaManagementServiceClient) RemoveFleetInstallationSites(ctx context.Context, request RemoveFleetInstallationSitesRequest) (response RemoveFleetInstallationSitesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeFleetInstallationSites, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveFleetInstallationSitesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveFleetInstallationSitesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveFleetInstallationSitesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveFleetInstallationSitesResponse")
	}
	return
}

// removeFleetInstallationSites implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) removeFleetInstallationSites(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/removeInstallationSites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveFleetInstallationSitesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationSiteSummary/RemoveFleetInstallationSites"
		err = common.PostProcessServiceError(err, "JavaManagementService", "RemoveFleetInstallationSites", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestCryptoAnalyses Request to perform crypto analysis on one or more selected targets in the Fleet. The result of the crypto analysis will be uploaded to the object storage bucket created by JMS on enabling the Crypto Event Analysis feature in the Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/RequestCryptoAnalyses.go.html to see an example of how to use RequestCryptoAnalyses API.
// A default retry strategy applies to this operation RequestCryptoAnalyses()
func (client JavaManagementServiceClient) RequestCryptoAnalyses(ctx context.Context, request RequestCryptoAnalysesRequest) (response RequestCryptoAnalysesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestCryptoAnalyses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestCryptoAnalysesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestCryptoAnalysesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestCryptoAnalysesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestCryptoAnalysesResponse")
	}
	return
}

// requestCryptoAnalyses implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) requestCryptoAnalyses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/requestCryptoAnalyses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestCryptoAnalysesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/RequestCryptoAnalyses"
		err = common.PostProcessServiceError(err, "JavaManagementService", "RequestCryptoAnalyses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestJavaMigrationAnalyses Request to perform a Java migration analysis. The results of the Java migration analysis will be uploaded to the
// Object Storage bucket that you designate when you enable the Java Migration Analysis feature.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/RequestJavaMigrationAnalyses.go.html to see an example of how to use RequestJavaMigrationAnalyses API.
// A default retry strategy applies to this operation RequestJavaMigrationAnalyses()
func (client JavaManagementServiceClient) RequestJavaMigrationAnalyses(ctx context.Context, request RequestJavaMigrationAnalysesRequest) (response RequestJavaMigrationAnalysesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestJavaMigrationAnalyses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestJavaMigrationAnalysesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestJavaMigrationAnalysesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestJavaMigrationAnalysesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestJavaMigrationAnalysesResponse")
	}
	return
}

// requestJavaMigrationAnalyses implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) requestJavaMigrationAnalyses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/requestJavaMigrationAnalyses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestJavaMigrationAnalysesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/RequestJavaMigrationAnalyses"
		err = common.PostProcessServiceError(err, "JavaManagementService", "RequestJavaMigrationAnalyses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestJfrRecordings Request to collect the JFR recordings on the selected target in the Fleet. The JFR files are uploaded to the object storage bucket created by JMS on enabling Generic JFR feature in the Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/RequestJfrRecordings.go.html to see an example of how to use RequestJfrRecordings API.
// A default retry strategy applies to this operation RequestJfrRecordings()
func (client JavaManagementServiceClient) RequestJfrRecordings(ctx context.Context, request RequestJfrRecordingsRequest) (response RequestJfrRecordingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestJfrRecordings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestJfrRecordingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestJfrRecordingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestJfrRecordingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestJfrRecordingsResponse")
	}
	return
}

// requestJfrRecordings implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) requestJfrRecordings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/requestJfrRecordings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestJfrRecordingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/RequestJfrRecordings"
		err = common.PostProcessServiceError(err, "JavaManagementService", "RequestJfrRecordings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestPerformanceTuningAnalyses Request to perform performance tuning analyses. The result of performance tuning analysis will be uploaded to the
// object storage bucket that you designated when you enabled the recording feature.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/RequestPerformanceTuningAnalyses.go.html to see an example of how to use RequestPerformanceTuningAnalyses API.
// A default retry strategy applies to this operation RequestPerformanceTuningAnalyses()
func (client JavaManagementServiceClient) RequestPerformanceTuningAnalyses(ctx context.Context, request RequestPerformanceTuningAnalysesRequest) (response RequestPerformanceTuningAnalysesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestPerformanceTuningAnalyses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestPerformanceTuningAnalysesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestPerformanceTuningAnalysesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestPerformanceTuningAnalysesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestPerformanceTuningAnalysesResponse")
	}
	return
}

// requestPerformanceTuningAnalyses implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) requestPerformanceTuningAnalyses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/requestPerformanceTuningAnalyses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestPerformanceTuningAnalysesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/RequestPerformanceTuningAnalyses"
		err = common.PostProcessServiceError(err, "JavaManagementService", "RequestPerformanceTuningAnalyses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScanJavaServerUsage Scan Java Server usage in a fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ScanJavaServerUsage.go.html to see an example of how to use ScanJavaServerUsage API.
// A default retry strategy applies to this operation ScanJavaServerUsage()
func (client JavaManagementServiceClient) ScanJavaServerUsage(ctx context.Context, request ScanJavaServerUsageRequest) (response ScanJavaServerUsageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scanJavaServerUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScanJavaServerUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScanJavaServerUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScanJavaServerUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScanJavaServerUsageResponse")
	}
	return
}

// scanJavaServerUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) scanJavaServerUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/scanJavaServerUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScanJavaServerUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaServerUsage/ScanJavaServerUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ScanJavaServerUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScanLibraryUsage Scan library usage in a fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ScanLibraryUsage.go.html to see an example of how to use ScanLibraryUsage API.
// A default retry strategy applies to this operation ScanLibraryUsage()
func (client JavaManagementServiceClient) ScanLibraryUsage(ctx context.Context, request ScanLibraryUsageRequest) (response ScanLibraryUsageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scanLibraryUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScanLibraryUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScanLibraryUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScanLibraryUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScanLibraryUsageResponse")
	}
	return
}

// scanLibraryUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) scanLibraryUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/scanLibraryUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScanLibraryUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/LibraryUsage/ScanLibraryUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ScanLibraryUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeApplicationInstallationUsage Summarizes the application installation usage in a Fleet filtered by query parameters. In contrast to SummarizeApplicationUsage, which provides only information aggregated by application name, this operation provides installation details. This allows for better focusing of actions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeApplicationInstallationUsage.go.html to see an example of how to use SummarizeApplicationInstallationUsage API.
// A default retry strategy applies to this operation SummarizeApplicationInstallationUsage()
func (client JavaManagementServiceClient) SummarizeApplicationInstallationUsage(ctx context.Context, request SummarizeApplicationInstallationUsageRequest) (response SummarizeApplicationInstallationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeApplicationInstallationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeApplicationInstallationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeApplicationInstallationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeApplicationInstallationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeApplicationInstallationUsageResponse")
	}
	return
}

// summarizeApplicationInstallationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeApplicationInstallationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeApplicationInstallationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeApplicationInstallationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ApplicationInstallationUsageSummary/SummarizeApplicationInstallationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeApplicationInstallationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeApplicationUsage List application usage in a Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeApplicationUsage.go.html to see an example of how to use SummarizeApplicationUsage API.
// A default retry strategy applies to this operation SummarizeApplicationUsage()
func (client JavaManagementServiceClient) SummarizeApplicationUsage(ctx context.Context, request SummarizeApplicationUsageRequest) (response SummarizeApplicationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeApplicationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeApplicationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeApplicationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeApplicationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeApplicationUsageResponse")
	}
	return
}

// summarizeApplicationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeApplicationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeApplicationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeApplicationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ApplicationUsage/SummarizeApplicationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeApplicationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDeployedApplicationInstallationUsage Summarize installation usage of an application deployed on Java servers in a fleet filtered by query parameters. In contrast to SummarizeDeployedApplicationUsage, which provides only information aggregated by the deployment information, this operation provides installation details and allows for better focusing of actions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeDeployedApplicationInstallationUsage.go.html to see an example of how to use SummarizeDeployedApplicationInstallationUsage API.
// A default retry strategy applies to this operation SummarizeDeployedApplicationInstallationUsage()
func (client JavaManagementServiceClient) SummarizeDeployedApplicationInstallationUsage(ctx context.Context, request SummarizeDeployedApplicationInstallationUsageRequest) (response SummarizeDeployedApplicationInstallationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDeployedApplicationInstallationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDeployedApplicationInstallationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDeployedApplicationInstallationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDeployedApplicationInstallationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDeployedApplicationInstallationUsageResponse")
	}
	return
}

// summarizeDeployedApplicationInstallationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeDeployedApplicationInstallationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeDeployedApplicationInstallationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDeployedApplicationInstallationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/DeployedApplicationInstallationUsageSummary/SummarizeDeployedApplicationInstallationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeDeployedApplicationInstallationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDeployedApplicationUsage List of deployed applications in a Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeDeployedApplicationUsage.go.html to see an example of how to use SummarizeDeployedApplicationUsage API.
// A default retry strategy applies to this operation SummarizeDeployedApplicationUsage()
func (client JavaManagementServiceClient) SummarizeDeployedApplicationUsage(ctx context.Context, request SummarizeDeployedApplicationUsageRequest) (response SummarizeDeployedApplicationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDeployedApplicationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDeployedApplicationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDeployedApplicationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDeployedApplicationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDeployedApplicationUsageResponse")
	}
	return
}

// summarizeDeployedApplicationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeDeployedApplicationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeDeployedApplicationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDeployedApplicationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/DeployedApplicationUsage/SummarizeDeployedApplicationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeDeployedApplicationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeInstallationUsage List Java installation usage in a Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeInstallationUsage.go.html to see an example of how to use SummarizeInstallationUsage API.
// A default retry strategy applies to this operation SummarizeInstallationUsage()
func (client JavaManagementServiceClient) SummarizeInstallationUsage(ctx context.Context, request SummarizeInstallationUsageRequest) (response SummarizeInstallationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeInstallationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeInstallationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeInstallationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeInstallationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeInstallationUsageResponse")
	}
	return
}

// summarizeInstallationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeInstallationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeInstallationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeInstallationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationUsage/SummarizeInstallationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeInstallationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeJavaServerInstanceUsage List Java Server instances in a fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJavaServerInstanceUsage.go.html to see an example of how to use SummarizeJavaServerInstanceUsage API.
// A default retry strategy applies to this operation SummarizeJavaServerInstanceUsage()
func (client JavaManagementServiceClient) SummarizeJavaServerInstanceUsage(ctx context.Context, request SummarizeJavaServerInstanceUsageRequest) (response SummarizeJavaServerInstanceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeJavaServerInstanceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeJavaServerInstanceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeJavaServerInstanceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeJavaServerInstanceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeJavaServerInstanceUsageResponse")
	}
	return
}

// summarizeJavaServerInstanceUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeJavaServerInstanceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeJavaServerInstanceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeJavaServerInstanceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaServerInstanceUsage/SummarizeJavaServerInstanceUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeJavaServerInstanceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeJavaServerUsage List of Java servers in a Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJavaServerUsage.go.html to see an example of how to use SummarizeJavaServerUsage API.
// A default retry strategy applies to this operation SummarizeJavaServerUsage()
func (client JavaManagementServiceClient) SummarizeJavaServerUsage(ctx context.Context, request SummarizeJavaServerUsageRequest) (response SummarizeJavaServerUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeJavaServerUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeJavaServerUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeJavaServerUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeJavaServerUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeJavaServerUsageResponse")
	}
	return
}

// summarizeJavaServerUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeJavaServerUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeJavaServerUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeJavaServerUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JavaServerUsage/SummarizeJavaServerUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeJavaServerUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeJreUsage List Java Runtime usage in a specified Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJreUsage.go.html to see an example of how to use SummarizeJreUsage API.
// A default retry strategy applies to this operation SummarizeJreUsage()
func (client JavaManagementServiceClient) SummarizeJreUsage(ctx context.Context, request SummarizeJreUsageRequest) (response SummarizeJreUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeJreUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeJreUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeJreUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeJreUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeJreUsageResponse")
	}
	return
}

// summarizeJreUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeJreUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeJreUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeJreUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JreUsage/SummarizeJreUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeJreUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeLibraryUsage List libraries in a fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeLibraryUsage.go.html to see an example of how to use SummarizeLibraryUsage API.
// A default retry strategy applies to this operation SummarizeLibraryUsage()
func (client JavaManagementServiceClient) SummarizeLibraryUsage(ctx context.Context, request SummarizeLibraryUsageRequest) (response SummarizeLibraryUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeLibraryUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeLibraryUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeLibraryUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeLibraryUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeLibraryUsageResponse")
	}
	return
}

// summarizeLibraryUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeLibraryUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeLibraryUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeLibraryUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/LibraryUsage/SummarizeLibraryUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeLibraryUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeManagedInstanceUsage List managed instance usage in a Fleet filtered by query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeManagedInstanceUsage.go.html to see an example of how to use SummarizeManagedInstanceUsage API.
// A default retry strategy applies to this operation SummarizeManagedInstanceUsage()
func (client JavaManagementServiceClient) SummarizeManagedInstanceUsage(ctx context.Context, request SummarizeManagedInstanceUsageRequest) (response SummarizeManagedInstanceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeManagedInstanceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeManagedInstanceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeManagedInstanceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeManagedInstanceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeManagedInstanceUsageResponse")
	}
	return
}

// summarizeManagedInstanceUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeManagedInstanceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeManagedInstanceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeManagedInstanceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ManagedInstanceUsage/SummarizeManagedInstanceUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeManagedInstanceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeResourceInventory Retrieve the inventory of JMS resources in the specified compartment: a list of the number of _active_ fleets, managed instances, Java Runtimes, Java installations, and applications.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeResourceInventory.go.html to see an example of how to use SummarizeResourceInventory API.
// A default retry strategy applies to this operation SummarizeResourceInventory()
func (client JavaManagementServiceClient) SummarizeResourceInventory(ctx context.Context, request SummarizeResourceInventoryRequest) (response SummarizeResourceInventoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeResourceInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeResourceInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeResourceInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeResourceInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeResourceInventoryResponse")
	}
	return
}

// summarizeResourceInventory implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeResourceInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/summarizeResourceInventory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeResourceInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/SummarizeResourceInventory"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeResourceInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDrsFile Request to perform validaition of the DRS file and update the existing file in the Object Storage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/UpdateDrsFile.go.html to see an example of how to use UpdateDrsFile API.
// A default retry strategy applies to this operation UpdateDrsFile()
func (client JavaManagementServiceClient) UpdateDrsFile(ctx context.Context, request UpdateDrsFileRequest) (response UpdateDrsFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateDrsFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDrsFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDrsFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDrsFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDrsFileResponse")
	}
	return
}

// updateDrsFile implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateDrsFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}/drsFiles/{drsFileKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDrsFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/UpdateDrsFile"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateDrsFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExportSetting Updates existing export setting for the specified Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/UpdateExportSetting.go.html to see an example of how to use UpdateExportSetting API.
// A default retry strategy applies to this operation UpdateExportSetting()
func (client JavaManagementServiceClient) UpdateExportSetting(ctx context.Context, request UpdateExportSettingRequest) (response UpdateExportSettingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExportSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExportSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExportSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExportSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExportSettingResponse")
	}
	return
}

// updateExportSetting implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateExportSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}/exportSetting", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExportSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ExportSetting/UpdateExportSetting"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateExportSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFleet Update the Fleet specified by an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/UpdateFleet.go.html to see an example of how to use UpdateFleet API.
// A default retry strategy applies to this operation UpdateFleet()
func (client JavaManagementServiceClient) UpdateFleet(ctx context.Context, request UpdateFleetRequest) (response UpdateFleetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFleetResponse")
	}
	return
}

// updateFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/UpdateFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFleetAdvancedFeatureConfiguration Update advanced feature configurations for the Fleet.
// Ensure that the namespace and bucket storage are created prior to turning on the JfrRecording or CryptoEventAnalysis feature.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/UpdateFleetAdvancedFeatureConfiguration.go.html to see an example of how to use UpdateFleetAdvancedFeatureConfiguration API.
// A default retry strategy applies to this operation UpdateFleetAdvancedFeatureConfiguration()
func (client JavaManagementServiceClient) UpdateFleetAdvancedFeatureConfiguration(ctx context.Context, request UpdateFleetAdvancedFeatureConfigurationRequest) (response UpdateFleetAdvancedFeatureConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateFleetAdvancedFeatureConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFleetAdvancedFeatureConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFleetAdvancedFeatureConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFleetAdvancedFeatureConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFleetAdvancedFeatureConfigurationResponse")
	}
	return
}

// updateFleetAdvancedFeatureConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateFleetAdvancedFeatureConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}/advancedFeatureConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFleetAdvancedFeatureConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetAdvancedFeatureConfiguration/UpdateFleetAdvancedFeatureConfiguration"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateFleetAdvancedFeatureConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFleetAgentConfiguration Update the Fleet Agent Configuration for the specified Fleet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/UpdateFleetAgentConfiguration.go.html to see an example of how to use UpdateFleetAgentConfiguration API.
// A default retry strategy applies to this operation UpdateFleetAgentConfiguration()
func (client JavaManagementServiceClient) UpdateFleetAgentConfiguration(ctx context.Context, request UpdateFleetAgentConfigurationRequest) (response UpdateFleetAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFleetAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFleetAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFleetAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFleetAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFleetAgentConfigurationResponse")
	}
	return
}

// updateFleetAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateFleetAgentConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}/agentConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFleetAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetAgentConfiguration/UpdateFleetAgentConfiguration"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateFleetAgentConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
