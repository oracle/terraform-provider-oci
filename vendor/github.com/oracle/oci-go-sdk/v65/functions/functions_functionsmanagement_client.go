// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FunctionsManagementClient a client for FunctionsManagement
type FunctionsManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFunctionsManagementClientWithConfigurationProvider Creates a new default FunctionsManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFunctionsManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FunctionsManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("functions"); !enabled {
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
	return newFunctionsManagementClientFromBaseClient(baseClient, provider)
}

// NewFunctionsManagementClientWithOboToken Creates a new default FunctionsManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFunctionsManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FunctionsManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFunctionsManagementClientFromBaseClient(baseClient, configProvider)
}

func newFunctionsManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FunctionsManagementClient, err error) {
	// FunctionsManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FunctionsManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FunctionsManagementClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FunctionsManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("functions", "https://functions.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FunctionsManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FunctionsManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeApplicationCompartment Moves an application into a different compartment within the same tenancy.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ChangeApplicationCompartment.go.html to see an example of how to use ChangeApplicationCompartment API.
// A default retry strategy applies to this operation ChangeApplicationCompartment()
func (client FunctionsManagementClient) ChangeApplicationCompartment(ctx context.Context, request ChangeApplicationCompartmentRequest) (response ChangeApplicationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeApplicationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeApplicationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeApplicationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeApplicationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeApplicationCompartmentResponse")
	}
	return
}

// changeApplicationCompartment implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) changeApplicationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applications/{applicationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeApplicationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Application/ChangeApplicationCompartment"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "ChangeApplicationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateApplication Creates a new application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/CreateApplication.go.html to see an example of how to use CreateApplication API.
// A default retry strategy applies to this operation CreateApplication()
func (client FunctionsManagementClient) CreateApplication(ctx context.Context, request CreateApplicationRequest) (response CreateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
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
func (client FunctionsManagementClient) createApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Application/CreateApplication"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "CreateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFunction Creates a new function.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/CreateFunction.go.html to see an example of how to use CreateFunction API.
// A default retry strategy applies to this operation CreateFunction()
func (client FunctionsManagementClient) CreateFunction(ctx context.Context, request CreateFunctionRequest) (response CreateFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFunctionResponse")
	}
	return
}

// createFunction implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) createFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/functions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Function/CreateFunction"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "CreateFunction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApplication Deletes an application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/DeleteApplication.go.html to see an example of how to use DeleteApplication API.
// A default retry strategy applies to this operation DeleteApplication()
func (client FunctionsManagementClient) DeleteApplication(ctx context.Context, request DeleteApplicationRequest) (response DeleteApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client FunctionsManagementClient) deleteApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/applications/{applicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Application/DeleteApplication"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "DeleteApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFunction Deletes a function.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/DeleteFunction.go.html to see an example of how to use DeleteFunction API.
// A default retry strategy applies to this operation DeleteFunction()
func (client FunctionsManagementClient) DeleteFunction(ctx context.Context, request DeleteFunctionRequest) (response DeleteFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFunctionResponse")
	}
	return
}

// deleteFunction implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) deleteFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/functions/{functionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Function/DeleteFunction"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "DeleteFunction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApplication Retrieves an application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/GetApplication.go.html to see an example of how to use GetApplication API.
// A default retry strategy applies to this operation GetApplication()
func (client FunctionsManagementClient) GetApplication(ctx context.Context, request GetApplicationRequest) (response GetApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client FunctionsManagementClient) getApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications/{applicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Application/GetApplication"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "GetApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFunction Retrieves a function.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/GetFunction.go.html to see an example of how to use GetFunction API.
// A default retry strategy applies to this operation GetFunction()
func (client FunctionsManagementClient) GetFunction(ctx context.Context, request GetFunctionRequest) (response GetFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFunctionResponse")
	}
	return
}

// getFunction implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) getFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/functions/{functionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Function/GetFunction"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "GetFunction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPbfListing Fetches a Pre-built Function(PBF) Listing. Returns a PbfListing response model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/GetPbfListing.go.html to see an example of how to use GetPbfListing API.
// A default retry strategy applies to this operation GetPbfListing()
func (client FunctionsManagementClient) GetPbfListing(ctx context.Context, request GetPbfListingRequest) (response GetPbfListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPbfListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPbfListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPbfListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPbfListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPbfListingResponse")
	}
	return
}

// getPbfListing implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) getPbfListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pbfListings/{pbfListingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPbfListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/PbfListing/GetPbfListing"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "GetPbfListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPbfListingVersion Gets a PbfListingVersion by identifier for a PbfListing.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/GetPbfListingVersion.go.html to see an example of how to use GetPbfListingVersion API.
// A default retry strategy applies to this operation GetPbfListingVersion()
func (client FunctionsManagementClient) GetPbfListingVersion(ctx context.Context, request GetPbfListingVersionRequest) (response GetPbfListingVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPbfListingVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPbfListingVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPbfListingVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPbfListingVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPbfListingVersionResponse")
	}
	return
}

// getPbfListingVersion implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) getPbfListingVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pbfListingVersions/{pbfListingVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPbfListingVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/PbfListingVersion/GetPbfListingVersion"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "GetPbfListingVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplications Lists applications for a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListApplications.go.html to see an example of how to use ListApplications API.
// A default retry strategy applies to this operation ListApplications()
func (client FunctionsManagementClient) ListApplications(ctx context.Context, request ListApplicationsRequest) (response ListApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client FunctionsManagementClient) listApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/ApplicationSummary/ListApplications"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "ListApplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFunctions Lists functions for an application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListFunctions.go.html to see an example of how to use ListFunctions API.
// A default retry strategy applies to this operation ListFunctions()
func (client FunctionsManagementClient) ListFunctions(ctx context.Context, request ListFunctionsRequest) (response ListFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFunctions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFunctionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFunctionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFunctionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFunctionsResponse")
	}
	return
}

// listFunctions implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) listFunctions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/functions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFunctionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/FunctionSummary/ListFunctions"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "ListFunctions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPbfListingVersions Fetches a wrapped list of all Pre-built Function(PBF) Listing versions. Returns a PbfListingVersionCollection
// containing an array of PbfListingVersionSummary response models.
// Note that the PbfListingIdentifier must be provided as a query parameter, otherwise an exception shall
// be thrown.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListPbfListingVersions.go.html to see an example of how to use ListPbfListingVersions API.
// A default retry strategy applies to this operation ListPbfListingVersions()
func (client FunctionsManagementClient) ListPbfListingVersions(ctx context.Context, request ListPbfListingVersionsRequest) (response ListPbfListingVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPbfListingVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPbfListingVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPbfListingVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPbfListingVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPbfListingVersionsResponse")
	}
	return
}

// listPbfListingVersions implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) listPbfListingVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pbfListingVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPbfListingVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/PbfListingVersion/ListPbfListingVersions"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "ListPbfListingVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPbfListings Fetches a wrapped list of all Pre-built Function(PBF) Listings. Returns a PbfListingCollection containing
// an array of PbfListingSummary response models.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListPbfListings.go.html to see an example of how to use ListPbfListings API.
// A default retry strategy applies to this operation ListPbfListings()
func (client FunctionsManagementClient) ListPbfListings(ctx context.Context, request ListPbfListingsRequest) (response ListPbfListingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPbfListings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPbfListingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPbfListingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPbfListingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPbfListingsResponse")
	}
	return
}

// listPbfListings implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) listPbfListings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pbfListings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPbfListingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/PbfListing/ListPbfListings"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "ListPbfListings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTriggers Returns a list of Triggers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListTriggers.go.html to see an example of how to use ListTriggers API.
// A default retry strategy applies to this operation ListTriggers()
func (client FunctionsManagementClient) ListTriggers(ctx context.Context, request ListTriggersRequest) (response ListTriggersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTriggers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTriggersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTriggersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTriggersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTriggersResponse")
	}
	return
}

// listTriggers implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) listTriggers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pbfListings/triggers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTriggersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/TriggersCollection/ListTriggers"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "ListTriggers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateApplication Modifies an application
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/UpdateApplication.go.html to see an example of how to use UpdateApplication API.
// A default retry strategy applies to this operation UpdateApplication()
func (client FunctionsManagementClient) UpdateApplication(ctx context.Context, request UpdateApplicationRequest) (response UpdateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client FunctionsManagementClient) updateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/applications/{applicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Application/UpdateApplication"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "UpdateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFunction Modifies a function
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/UpdateFunction.go.html to see an example of how to use UpdateFunction API.
// A default retry strategy applies to this operation UpdateFunction()
func (client FunctionsManagementClient) UpdateFunction(ctx context.Context, request UpdateFunctionRequest) (response UpdateFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFunctionResponse")
	}
	return
}

// updateFunction implements the OCIOperation interface (enables retrying operations)
func (client FunctionsManagementClient) updateFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/functions/{functionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20181201/Function/UpdateFunction"
		err = common.PostProcessServiceError(err, "FunctionsManagement", "UpdateFunction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
