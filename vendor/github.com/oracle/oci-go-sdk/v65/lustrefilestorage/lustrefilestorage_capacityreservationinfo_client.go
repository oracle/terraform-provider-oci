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

// CapacityReservationInfoClient a client for CapacityReservationInfo
type CapacityReservationInfoClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCapacityReservationInfoClientWithConfigurationProvider Creates a new default CapacityReservationInfo client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCapacityReservationInfoClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CapacityReservationInfoClient, err error) {
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
	return newCapacityReservationInfoClientFromBaseClient(baseClient, provider)
}

// NewCapacityReservationInfoClientWithOboToken Creates a new default CapacityReservationInfo client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCapacityReservationInfoClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CapacityReservationInfoClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCapacityReservationInfoClientFromBaseClient(baseClient, configProvider)
}

func newCapacityReservationInfoClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CapacityReservationInfoClient, err error) {
	// CapacityReservationInfo service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CapacityReservationInfo"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CapacityReservationInfoClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CapacityReservationInfoClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CapacityReservationInfoClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CapacityReservationInfoClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateCapacityReservationInfo Creates a Capacity Reservation Info.
// A default retry strategy applies to this operation CreateCapacityReservationInfo()
func (client CapacityReservationInfoClient) CreateCapacityReservationInfo(ctx context.Context, request CreateCapacityReservationInfoRequest) (response CreateCapacityReservationInfoResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCapacityReservationInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCapacityReservationInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCapacityReservationInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCapacityReservationInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCapacityReservationInfoResponse")
	}
	return
}

// createCapacityReservationInfo implements the OCIOperation interface (enables retrying operations)
func (client CapacityReservationInfoClient) createCapacityReservationInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/capacityReservationInfos", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCapacityReservationInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CapacityReservationInfo/CreateCapacityReservationInfo"
		err = common.PostProcessServiceError(err, "CapacityReservationInfo", "CreateCapacityReservationInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCapacityReservationInfo Deletes a CapacityReservationInfo.
// A default retry strategy applies to this operation DeleteCapacityReservationInfo()
func (client CapacityReservationInfoClient) DeleteCapacityReservationInfo(ctx context.Context, request DeleteCapacityReservationInfoRequest) (response DeleteCapacityReservationInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCapacityReservationInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCapacityReservationInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCapacityReservationInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCapacityReservationInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCapacityReservationInfoResponse")
	}
	return
}

// deleteCapacityReservationInfo implements the OCIOperation interface (enables retrying operations)
func (client CapacityReservationInfoClient) deleteCapacityReservationInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/capacityReservationInfos/{capacityReservationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCapacityReservationInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CapacityReservationInfo/DeleteCapacityReservationInfo"
		err = common.PostProcessServiceError(err, "CapacityReservationInfo", "DeleteCapacityReservationInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCapacityReservationInfo Gets information about a CapacityReservationInfo.
// A default retry strategy applies to this operation GetCapacityReservationInfo()
func (client CapacityReservationInfoClient) GetCapacityReservationInfo(ctx context.Context, request GetCapacityReservationInfoRequest) (response GetCapacityReservationInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCapacityReservationInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCapacityReservationInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCapacityReservationInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCapacityReservationInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCapacityReservationInfoResponse")
	}
	return
}

// getCapacityReservationInfo implements the OCIOperation interface (enables retrying operations)
func (client CapacityReservationInfoClient) getCapacityReservationInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/capacityReservationInfos/{capacityReservationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCapacityReservationInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CapacityReservationInfo/GetCapacityReservationInfo"
		err = common.PostProcessServiceError(err, "CapacityReservationInfo", "GetCapacityReservationInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCapacityReservationInfos Gets a list of Capacity Reservation Info.
// A default retry strategy applies to this operation ListCapacityReservationInfos()
func (client CapacityReservationInfoClient) ListCapacityReservationInfos(ctx context.Context, request ListCapacityReservationInfosRequest) (response ListCapacityReservationInfosResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCapacityReservationInfos, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCapacityReservationInfosResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCapacityReservationInfosResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCapacityReservationInfosResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCapacityReservationInfosResponse")
	}
	return
}

// listCapacityReservationInfos implements the OCIOperation interface (enables retrying operations)
func (client CapacityReservationInfoClient) listCapacityReservationInfos(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/capacityReservationInfos", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCapacityReservationInfosResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CapacityReservationInfoCollection/ListCapacityReservationInfos"
		err = common.PostProcessServiceError(err, "CapacityReservationInfo", "ListCapacityReservationInfos", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCapacityReservationInfo Updates a CapacityReservationInfo.
// A default retry strategy applies to this operation UpdateCapacityReservationInfo()
func (client CapacityReservationInfoClient) UpdateCapacityReservationInfo(ctx context.Context, request UpdateCapacityReservationInfoRequest) (response UpdateCapacityReservationInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCapacityReservationInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCapacityReservationInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCapacityReservationInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCapacityReservationInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCapacityReservationInfoResponse")
	}
	return
}

// updateCapacityReservationInfo implements the OCIOperation interface (enables retrying operations)
func (client CapacityReservationInfoClient) updateCapacityReservationInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/capacityReservationInfos/{capacityReservationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCapacityReservationInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CapacityReservationInfo/UpdateCapacityReservationInfo"
		err = common.PostProcessServiceError(err, "CapacityReservationInfo", "UpdateCapacityReservationInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
