// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//GatewayClient a client for Gateway
type GatewayClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGatewayClientWithConfigurationProvider Creates a new default Gateway client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGatewayClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GatewayClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newGatewayClientFromBaseClient(baseClient, provider)
}

// NewGatewayClientWithOboToken Creates a new default Gateway client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewGatewayClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GatewayClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGatewayClientFromBaseClient(baseClient, configProvider)
}

func newGatewayClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GatewayClient, err error) {
	// Gateway service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GatewayClient{BaseClient: baseClient}
	client.BasePath = "20190501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GatewayClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apigateway", "https://apigateway.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GatewayClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GatewayClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeGatewayCompartment Changes the gateway compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ChangeGatewayCompartment.go.html to see an example of how to use ChangeGatewayCompartment API.
func (client GatewayClient) ChangeGatewayCompartment(ctx context.Context, request ChangeGatewayCompartmentRequest) (response ChangeGatewayCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeGatewayCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeGatewayCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeGatewayCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeGatewayCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeGatewayCompartmentResponse")
	}
	return
}

// changeGatewayCompartment implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) changeGatewayCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gateways/{gatewayId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeGatewayCompartmentResponse
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

// CreateGateway Creates a new gateway.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/CreateGateway.go.html to see an example of how to use CreateGateway API.
func (client GatewayClient) CreateGateway(ctx context.Context, request CreateGatewayRequest) (response CreateGatewayResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGatewayResponse")
	}
	return
}

// createGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) createGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gateways", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGatewayResponse
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

// DeleteGateway Deletes the gateway with the given identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/DeleteGateway.go.html to see an example of how to use DeleteGateway API.
func (client GatewayClient) DeleteGateway(ctx context.Context, request DeleteGatewayRequest) (response DeleteGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGatewayResponse")
	}
	return
}

// deleteGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) deleteGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/gateways/{gatewayId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGatewayResponse
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

// GetGateway Gets a gateway by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/GetGateway.go.html to see an example of how to use GetGateway API.
func (client GatewayClient) GetGateway(ctx context.Context, request GetGatewayRequest) (response GetGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGatewayResponse")
	}
	return
}

// getGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) getGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gateways/{gatewayId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGatewayResponse
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

// ListGateways Returns a list of gateways.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListGateways.go.html to see an example of how to use ListGateways API.
func (client GatewayClient) ListGateways(ctx context.Context, request ListGatewaysRequest) (response ListGatewaysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGateways, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGatewaysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGatewaysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGatewaysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGatewaysResponse")
	}
	return
}

// listGateways implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) listGateways(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gateways", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGatewaysResponse
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

// UpdateGateway Updates the gateway with the given identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/UpdateGateway.go.html to see an example of how to use UpdateGateway API.
func (client GatewayClient) UpdateGateway(ctx context.Context, request UpdateGatewayRequest) (response UpdateGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGatewayResponse")
	}
	return
}

// updateGateway implements the OCIOperation interface (enables retrying operations)
func (client GatewayClient) updateGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/gateways/{gatewayId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGatewayResponse
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
