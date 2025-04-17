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

// CpgOverrideClient a client for CpgOverride
type CpgOverrideClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCpgOverrideClientWithConfigurationProvider Creates a new default CpgOverride client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCpgOverrideClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CpgOverrideClient, err error) {
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
	return newCpgOverrideClientFromBaseClient(baseClient, provider)
}

// NewCpgOverrideClientWithOboToken Creates a new default CpgOverride client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCpgOverrideClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CpgOverrideClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCpgOverrideClientFromBaseClient(baseClient, configProvider)
}

func newCpgOverrideClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CpgOverrideClient, err error) {
	// CpgOverride service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CpgOverride"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CpgOverrideClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CpgOverrideClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CpgOverrideClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CpgOverrideClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateCpgOverride Creates a CPG Override.
// A default retry strategy applies to this operation CreateCpgOverride()
func (client CpgOverrideClient) CreateCpgOverride(ctx context.Context, request CreateCpgOverrideRequest) (response CreateCpgOverrideResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCpgOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCpgOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCpgOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCpgOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCpgOverrideResponse")
	}
	return
}

// createCpgOverride implements the OCIOperation interface (enables retrying operations)
func (client CpgOverrideClient) createCpgOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cpgOverrides", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCpgOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CpgOverride/CreateCpgOverride"
		err = common.PostProcessServiceError(err, "CpgOverride", "CreateCpgOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCpgOverride Deletes a CPG Override.
// A default retry strategy applies to this operation DeleteCpgOverride()
func (client CpgOverrideClient) DeleteCpgOverride(ctx context.Context, request DeleteCpgOverrideRequest) (response DeleteCpgOverrideResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCpgOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCpgOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCpgOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCpgOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCpgOverrideResponse")
	}
	return
}

// deleteCpgOverride implements the OCIOperation interface (enables retrying operations)
func (client CpgOverrideClient) deleteCpgOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cpgOverrides/{customerCpgId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCpgOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CpgOverride/DeleteCpgOverride"
		err = common.PostProcessServiceError(err, "CpgOverride", "DeleteCpgOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCpgOverride Gets information about a CPG Override.
// A default retry strategy applies to this operation GetCpgOverride()
func (client CpgOverrideClient) GetCpgOverride(ctx context.Context, request GetCpgOverrideRequest) (response GetCpgOverrideResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCpgOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCpgOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCpgOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCpgOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCpgOverrideResponse")
	}
	return
}

// getCpgOverride implements the OCIOperation interface (enables retrying operations)
func (client CpgOverrideClient) getCpgOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cpgOverrides/{customerCpgId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCpgOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CpgOverride/GetCpgOverride"
		err = common.PostProcessServiceError(err, "CpgOverride", "GetCpgOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCpgOverrides Gets a list of CPG Overrides.
// A default retry strategy applies to this operation ListCpgOverrides()
func (client CpgOverrideClient) ListCpgOverrides(ctx context.Context, request ListCpgOverridesRequest) (response ListCpgOverridesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCpgOverrides, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCpgOverridesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCpgOverridesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCpgOverridesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCpgOverridesResponse")
	}
	return
}

// listCpgOverrides implements the OCIOperation interface (enables retrying operations)
func (client CpgOverrideClient) listCpgOverrides(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cpgOverrides", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCpgOverridesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CpgOverrideCollection/ListCpgOverrides"
		err = common.PostProcessServiceError(err, "CpgOverride", "ListCpgOverrides", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCpgOverride Updates a CPG Override.
// A default retry strategy applies to this operation UpdateCpgOverride()
func (client CpgOverrideClient) UpdateCpgOverride(ctx context.Context, request UpdateCpgOverrideRequest) (response UpdateCpgOverrideResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCpgOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCpgOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCpgOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCpgOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCpgOverrideResponse")
	}
	return
}

// updateCpgOverride implements the OCIOperation interface (enables retrying operations)
func (client CpgOverrideClient) updateCpgOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cpgOverrides/{customerCpgId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCpgOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CpgOverride/UpdateCpgOverride"
		err = common.PostProcessServiceError(err, "CpgOverride", "UpdateCpgOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
