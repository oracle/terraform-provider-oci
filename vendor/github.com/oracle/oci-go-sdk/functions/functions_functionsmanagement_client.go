// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//FunctionsManagementClient a client for FunctionsManagement
type FunctionsManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFunctionsManagementClientWithConfigurationProvider Creates a new default FunctionsManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFunctionsManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FunctionsManagementClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newFunctionsManagementClientFromBaseClient(baseClient, configProvider)
}

// NewFunctionsManagementClientWithOboToken Creates a new default FunctionsManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewFunctionsManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FunctionsManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newFunctionsManagementClientFromBaseClient(baseClient, configProvider)
}

func newFunctionsManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FunctionsManagementClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *FunctionsManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeApplicationCompartment Moves an application into a different compartment within the same tenancy.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client FunctionsManagementClient) ChangeApplicationCompartment(ctx context.Context, request ChangeApplicationCompartmentRequest) (response ChangeApplicationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) changeApplicationCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applications/{applicationId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeApplicationCompartmentResponse
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

// CreateApplication Creates a new application.
func (client FunctionsManagementClient) CreateApplication(ctx context.Context, request CreateApplicationRequest) (response CreateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) createApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applications")
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

// CreateFunction Creates a new function.
func (client FunctionsManagementClient) CreateFunction(ctx context.Context, request CreateFunctionRequest) (response CreateFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) createFunction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/functions")
	if err != nil {
		return nil, err
	}

	var response CreateFunctionResponse
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

// DeleteApplication Deletes an application.
func (client FunctionsManagementClient) DeleteApplication(ctx context.Context, request DeleteApplicationRequest) (response DeleteApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) deleteApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/applications/{applicationId}")
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

// DeleteFunction Deletes a function.
func (client FunctionsManagementClient) DeleteFunction(ctx context.Context, request DeleteFunctionRequest) (response DeleteFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) deleteFunction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/functions/{functionId}")
	if err != nil {
		return nil, err
	}

	var response DeleteFunctionResponse
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

// GetApplication Retrieves an application.
func (client FunctionsManagementClient) GetApplication(ctx context.Context, request GetApplicationRequest) (response GetApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) getApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications/{applicationId}")
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

// GetFunction Retrieves a function.
func (client FunctionsManagementClient) GetFunction(ctx context.Context, request GetFunctionRequest) (response GetFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) getFunction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/functions/{functionId}")
	if err != nil {
		return nil, err
	}

	var response GetFunctionResponse
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

// ListApplications Lists applications for a compartment.
func (client FunctionsManagementClient) ListApplications(ctx context.Context, request ListApplicationsRequest) (response ListApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) listApplications(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications")
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

// ListFunctions Lists functions for an application.
func (client FunctionsManagementClient) ListFunctions(ctx context.Context, request ListFunctionsRequest) (response ListFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) listFunctions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/functions")
	if err != nil {
		return nil, err
	}

	var response ListFunctionsResponse
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

// UpdateApplication Modifies an application
func (client FunctionsManagementClient) UpdateApplication(ctx context.Context, request UpdateApplicationRequest) (response UpdateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) updateApplication(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/applications/{applicationId}")
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

// UpdateFunction Modifies a function
func (client FunctionsManagementClient) UpdateFunction(ctx context.Context, request UpdateFunctionRequest) (response UpdateFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client FunctionsManagementClient) updateFunction(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/functions/{functionId}")
	if err != nil {
		return nil, err
	}

	var response UpdateFunctionResponse
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
