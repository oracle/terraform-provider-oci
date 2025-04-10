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

// LfsCpgInfoClient a client for LfsCpgInfo
type LfsCpgInfoClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLfsCpgInfoClientWithConfigurationProvider Creates a new default LfsCpgInfo client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLfsCpgInfoClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LfsCpgInfoClient, err error) {
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
	return newLfsCpgInfoClientFromBaseClient(baseClient, provider)
}

// NewLfsCpgInfoClientWithOboToken Creates a new default LfsCpgInfo client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLfsCpgInfoClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LfsCpgInfoClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLfsCpgInfoClientFromBaseClient(baseClient, configProvider)
}

func newLfsCpgInfoClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LfsCpgInfoClient, err error) {
	// LfsCpgInfo service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LfsCpgInfo"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LfsCpgInfoClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LfsCpgInfoClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LfsCpgInfoClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LfsCpgInfoClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateLfsCpgInfo Creates a CPG lfsCpgInfo.
// A default retry strategy applies to this operation CreateLfsCpgInfo()
func (client LfsCpgInfoClient) CreateLfsCpgInfo(ctx context.Context, request CreateLfsCpgInfoRequest) (response CreateLfsCpgInfoResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createLfsCpgInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLfsCpgInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLfsCpgInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLfsCpgInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLfsCpgInfoResponse")
	}
	return
}

// createLfsCpgInfo implements the OCIOperation interface (enables retrying operations)
func (client LfsCpgInfoClient) createLfsCpgInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lfsCpgInfos", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLfsCpgInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LfsCpgInfo/CreateLfsCpgInfo"
		err = common.PostProcessServiceError(err, "LfsCpgInfo", "CreateLfsCpgInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLfsCpgInfo Deletes a LFS CPG Info.
// A default retry strategy applies to this operation DeleteLfsCpgInfo()
func (client LfsCpgInfoClient) DeleteLfsCpgInfo(ctx context.Context, request DeleteLfsCpgInfoRequest) (response DeleteLfsCpgInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLfsCpgInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLfsCpgInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLfsCpgInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLfsCpgInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLfsCpgInfoResponse")
	}
	return
}

// deleteLfsCpgInfo implements the OCIOperation interface (enables retrying operations)
func (client LfsCpgInfoClient) deleteLfsCpgInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/lfsCpgInfos/{lfsCpgId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLfsCpgInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LfsCpgInfo/DeleteLfsCpgInfo"
		err = common.PostProcessServiceError(err, "LfsCpgInfo", "DeleteLfsCpgInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLfsCpgInfo Gets information about a LFS CPG
// A default retry strategy applies to this operation GetLfsCpgInfo()
func (client LfsCpgInfoClient) GetLfsCpgInfo(ctx context.Context, request GetLfsCpgInfoRequest) (response GetLfsCpgInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLfsCpgInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLfsCpgInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLfsCpgInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLfsCpgInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLfsCpgInfoResponse")
	}
	return
}

// getLfsCpgInfo implements the OCIOperation interface (enables retrying operations)
func (client LfsCpgInfoClient) getLfsCpgInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lfsCpgInfos/{lfsCpgId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLfsCpgInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LfsCpgInfo/GetLfsCpgInfo"
		err = common.PostProcessServiceError(err, "LfsCpgInfo", "GetLfsCpgInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLfsCpgInfos Gets a list of lfsCpgInfo.
// A default retry strategy applies to this operation ListLfsCpgInfos()
func (client LfsCpgInfoClient) ListLfsCpgInfos(ctx context.Context, request ListLfsCpgInfosRequest) (response ListLfsCpgInfosResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLfsCpgInfos, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLfsCpgInfosResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLfsCpgInfosResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLfsCpgInfosResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLfsCpgInfosResponse")
	}
	return
}

// listLfsCpgInfos implements the OCIOperation interface (enables retrying operations)
func (client LfsCpgInfoClient) listLfsCpgInfos(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lfsCpgInfos", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLfsCpgInfosResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LfsCpgInfoCollection/ListLfsCpgInfos"
		err = common.PostProcessServiceError(err, "LfsCpgInfo", "ListLfsCpgInfos", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLfsCpgInfo Updates a a LfsCpgInfo.
// A default retry strategy applies to this operation UpdateLfsCpgInfo()
func (client LfsCpgInfoClient) UpdateLfsCpgInfo(ctx context.Context, request UpdateLfsCpgInfoRequest) (response UpdateLfsCpgInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLfsCpgInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLfsCpgInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLfsCpgInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLfsCpgInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLfsCpgInfoResponse")
	}
	return
}

// updateLfsCpgInfo implements the OCIOperation interface (enables retrying operations)
func (client LfsCpgInfoClient) updateLfsCpgInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/lfsCpgInfos/{lfsCpgId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLfsCpgInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/LfsCpgInfo/UpdateLfsCpgInfo"
		err = common.PostProcessServiceError(err, "LfsCpgInfo", "UpdateLfsCpgInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
