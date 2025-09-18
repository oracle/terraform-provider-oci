// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LfsPoolClient a client for LfsPool
type LfsPoolClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLfsPoolClientWithConfigurationProvider Creates a new default LfsPool client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLfsPoolClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LfsPoolClient, err error) {
	if enabled := common.CheckForEnabledServices("lustrefilestorage"); !enabled {
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
	return newLfsPoolClientFromBaseClient(baseClient, provider)
}

// NewLfsPoolClientWithOboToken Creates a new default LfsPool client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLfsPoolClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LfsPoolClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLfsPoolClientFromBaseClient(baseClient, configProvider)
}

func newLfsPoolClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LfsPoolClient, err error) {
	// LfsPool service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LfsPool"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LfsPoolClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LfsPoolClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LfsPoolClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LfsPoolClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreatePool Creates a Pool.
// A default retry strategy applies to this operation CreatePool()
func (client LfsPoolClient) CreatePool(ctx context.Context, request CreatePoolRequest) (response CreatePoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePoolResponse")
	}
	return
}

// createPool implements the OCIOperation interface (enables retrying operations)
func (client LfsPoolClient) createPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/Pool/CreatePool"
		err = common.PostProcessServiceError(err, "LfsPool", "CreatePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePool Deletes the pool.
// A default retry strategy applies to this operation DeletePool()
func (client LfsPoolClient) DeletePool(ctx context.Context, request DeletePoolRequest) (response DeletePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePoolResponse")
	}
	return
}

// deletePool implements the OCIOperation interface (enables retrying operations)
func (client LfsPoolClient) deletePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/pools/{poolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/Pool/DeletePool"
		err = common.PostProcessServiceError(err, "LfsPool", "DeletePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPool Gets information about a Pool.
// A default retry strategy applies to this operation GetPool()
func (client LfsPoolClient) GetPool(ctx context.Context, request GetPoolRequest) (response GetPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPoolResponse")
	}
	return
}

// getPool implements the OCIOperation interface (enables retrying operations)
func (client LfsPoolClient) getPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pools/{poolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/Pool/GetPool"
		err = common.PostProcessServiceError(err, "LfsPool", "GetPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPools Gets a list of Pools.
// A default retry strategy applies to this operation ListPools()
func (client LfsPoolClient) ListPools(ctx context.Context, request ListPoolsRequest) (response ListPoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPoolsResponse")
	}
	return
}

// listPools implements the OCIOperation interface (enables retrying operations)
func (client LfsPoolClient) listPools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/PoolCollection/ListPools"
		err = common.PostProcessServiceError(err, "LfsPool", "ListPools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePool Updates a Pool.
// A default retry strategy applies to this operation UpdatePool()
func (client LfsPoolClient) UpdatePool(ctx context.Context, request UpdatePoolRequest) (response UpdatePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePoolResponse")
	}
	return
}

// updatePool implements the OCIOperation interface (enables retrying operations)
func (client LfsPoolClient) updatePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/pools/{poolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/Pool/UpdatePool"
		err = common.PostProcessServiceError(err, "LfsPool", "UpdatePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
