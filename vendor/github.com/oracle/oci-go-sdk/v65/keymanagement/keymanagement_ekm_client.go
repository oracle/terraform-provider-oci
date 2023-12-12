// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// EkmClient a client for Ekm
type EkmClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEkmClientWithConfigurationProvider Creates a new default Ekm client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEkmClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EkmClient, err error) {
	if enabled := common.CheckForEnabledServices("keymanagement"); !enabled {
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
	return newEkmClientFromBaseClient(baseClient, provider)
}

// NewEkmClientWithOboToken Creates a new default Ekm client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewEkmClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EkmClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEkmClientFromBaseClient(baseClient, configProvider)
}

func newEkmClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EkmClient, err error) {
	// Ekm service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Ekm"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EkmClient{BaseClient: baseClient}
	client.BasePath = ""
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EkmClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("kms", "https://kms.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EkmClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EkmClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateEkmsPrivateEndpoint Create a new EKMS private endpoint used to connect to external key manager system
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/CreateEkmsPrivateEndpoint.go.html to see an example of how to use CreateEkmsPrivateEndpoint API.
func (client EkmClient) CreateEkmsPrivateEndpoint(ctx context.Context, request CreateEkmsPrivateEndpointRequest) (response CreateEkmsPrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEkmsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEkmsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEkmsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEkmsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEkmsPrivateEndpointResponse")
	}
	return
}

// createEkmsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EkmClient) createEkmsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/ekmsPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEkmsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/EkmsPrivateEndpoint/CreateEkmsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Ekm", "CreateEkmsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEkmsPrivateEndpoint Deletes EKMS private endpoint by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/DeleteEkmsPrivateEndpoint.go.html to see an example of how to use DeleteEkmsPrivateEndpoint API.
func (client EkmClient) DeleteEkmsPrivateEndpoint(ctx context.Context, request DeleteEkmsPrivateEndpointRequest) (response DeleteEkmsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEkmsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEkmsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEkmsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEkmsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEkmsPrivateEndpointResponse")
	}
	return
}

// deleteEkmsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EkmClient) deleteEkmsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/20180608/ekmsPrivateEndpoints/{ekmsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEkmsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/EkmsPrivateEndpoint/DeleteEkmsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Ekm", "DeleteEkmsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEkmsPrivateEndpoint Gets a specific EKMS private by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/GetEkmsPrivateEndpoint.go.html to see an example of how to use GetEkmsPrivateEndpoint API.
func (client EkmClient) GetEkmsPrivateEndpoint(ctx context.Context, request GetEkmsPrivateEndpointRequest) (response GetEkmsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEkmsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEkmsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEkmsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEkmsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEkmsPrivateEndpointResponse")
	}
	return
}

// getEkmsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EkmClient) getEkmsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/ekmsPrivateEndpoints/{ekmsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEkmsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/EkmsPrivateEndpoint/GetEkmsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Ekm", "GetEkmsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEkmsPrivateEndpoints Returns a list of all the EKMS private endpoints in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ListEkmsPrivateEndpoints.go.html to see an example of how to use ListEkmsPrivateEndpoints API.
func (client EkmClient) ListEkmsPrivateEndpoints(ctx context.Context, request ListEkmsPrivateEndpointsRequest) (response ListEkmsPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEkmsPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEkmsPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEkmsPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEkmsPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEkmsPrivateEndpointsResponse")
	}
	return
}

// listEkmsPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client EkmClient) listEkmsPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/ekmsPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEkmsPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/EkmsPrivateEndpointSummary/ListEkmsPrivateEndpoints"
		err = common.PostProcessServiceError(err, "Ekm", "ListEkmsPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEkmsPrivateEndpoint Updates EKMS private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/UpdateEkmsPrivateEndpoint.go.html to see an example of how to use UpdateEkmsPrivateEndpoint API.
func (client EkmClient) UpdateEkmsPrivateEndpoint(ctx context.Context, request UpdateEkmsPrivateEndpointRequest) (response UpdateEkmsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEkmsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEkmsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEkmsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEkmsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEkmsPrivateEndpointResponse")
	}
	return
}

// updateEkmsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EkmClient) updateEkmsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/20180608/ekmsPrivateEndpoints/{ekmsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEkmsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/EkmsPrivateEndpoint/UpdateEkmsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Ekm", "UpdateEkmsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
