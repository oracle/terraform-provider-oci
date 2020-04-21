// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//RedirectClient a client for Redirect
type RedirectClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRedirectClientWithConfigurationProvider Creates a new default Redirect client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRedirectClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RedirectClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newRedirectClientFromBaseClient(baseClient, configProvider)
}

// NewRedirectClientWithOboToken Creates a new default Redirect client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewRedirectClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RedirectClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newRedirectClientFromBaseClient(baseClient, configProvider)
}

func newRedirectClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RedirectClient, err error) {
	client = RedirectClient{BaseClient: baseClient}
	client.BasePath = "20181116"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RedirectClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("waas", "https://waas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RedirectClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RedirectClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeHttpRedirectCompartment Moves HTTP Redirect into a different compartment. When provided, If-Match is checked against ETag values of the WAAS policy.
func (client RedirectClient) ChangeHttpRedirectCompartment(ctx context.Context, request ChangeHttpRedirectCompartmentRequest) (response ChangeHttpRedirectCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeHttpRedirectCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeHttpRedirectCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeHttpRedirectCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeHttpRedirectCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeHttpRedirectCompartmentResponse")
	}
	return
}

// changeHttpRedirectCompartment implements the OCIOperation interface (enables retrying operations)
func (client RedirectClient) changeHttpRedirectCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/httpRedirects/{httpRedirectId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeHttpRedirectCompartmentResponse
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

// CreateHttpRedirect Creates a new HTTP Redirect on the WAF edge.
func (client RedirectClient) CreateHttpRedirect(ctx context.Context, request CreateHttpRedirectRequest) (response CreateHttpRedirectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createHttpRedirect, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateHttpRedirectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateHttpRedirectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateHttpRedirectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateHttpRedirectResponse")
	}
	return
}

// createHttpRedirect implements the OCIOperation interface (enables retrying operations)
func (client RedirectClient) createHttpRedirect(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/httpRedirects")
	if err != nil {
		return nil, err
	}

	var response CreateHttpRedirectResponse
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

// DeleteHttpRedirect Deletes a redirect.
func (client RedirectClient) DeleteHttpRedirect(ctx context.Context, request DeleteHttpRedirectRequest) (response DeleteHttpRedirectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteHttpRedirect, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteHttpRedirectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteHttpRedirectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteHttpRedirectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteHttpRedirectResponse")
	}
	return
}

// deleteHttpRedirect implements the OCIOperation interface (enables retrying operations)
func (client RedirectClient) deleteHttpRedirect(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/httpRedirects/{httpRedirectId}")
	if err != nil {
		return nil, err
	}

	var response DeleteHttpRedirectResponse
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

// GetHttpRedirect Gets the details of a HTTP Redirect.
func (client RedirectClient) GetHttpRedirect(ctx context.Context, request GetHttpRedirectRequest) (response GetHttpRedirectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHttpRedirect, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHttpRedirectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHttpRedirectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHttpRedirectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHttpRedirectResponse")
	}
	return
}

// getHttpRedirect implements the OCIOperation interface (enables retrying operations)
func (client RedirectClient) getHttpRedirect(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/httpRedirects/{httpRedirectId}")
	if err != nil {
		return nil, err
	}

	var response GetHttpRedirectResponse
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

// ListHttpRedirects Gets a list of HTTP Redirects.
func (client RedirectClient) ListHttpRedirects(ctx context.Context, request ListHttpRedirectsRequest) (response ListHttpRedirectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHttpRedirects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHttpRedirectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHttpRedirectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHttpRedirectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHttpRedirectsResponse")
	}
	return
}

// listHttpRedirects implements the OCIOperation interface (enables retrying operations)
func (client RedirectClient) listHttpRedirects(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/httpRedirects")
	if err != nil {
		return nil, err
	}

	var response ListHttpRedirectsResponse
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

// UpdateHttpRedirect Updates the details of a HTTP Redirect, including target and tags. Only the fields specified in the request body will be updated; all other properties will remain unchanged.
func (client RedirectClient) UpdateHttpRedirect(ctx context.Context, request UpdateHttpRedirectRequest) (response UpdateHttpRedirectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateHttpRedirect, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHttpRedirectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHttpRedirectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHttpRedirectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHttpRedirectResponse")
	}
	return
}

// updateHttpRedirect implements the OCIOperation interface (enables retrying operations)
func (client RedirectClient) updateHttpRedirect(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/httpRedirects/{httpRedirectId}")
	if err != nil {
		return nil, err
	}

	var response UpdateHttpRedirectResponse
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
