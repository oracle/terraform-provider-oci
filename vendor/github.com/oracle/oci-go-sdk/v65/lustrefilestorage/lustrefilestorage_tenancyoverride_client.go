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

// TenancyOverrideClient a client for TenancyOverride
type TenancyOverrideClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTenancyOverrideClientWithConfigurationProvider Creates a new default TenancyOverride client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTenancyOverrideClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TenancyOverrideClient, err error) {
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
	return newTenancyOverrideClientFromBaseClient(baseClient, provider)
}

// NewTenancyOverrideClientWithOboToken Creates a new default TenancyOverride client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewTenancyOverrideClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TenancyOverrideClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTenancyOverrideClientFromBaseClient(baseClient, configProvider)
}

func newTenancyOverrideClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TenancyOverrideClient, err error) {
	// TenancyOverride service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("TenancyOverride"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = TenancyOverrideClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TenancyOverrideClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TenancyOverrideClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TenancyOverrideClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateTenancyOverride Creates a Tenancy Override.
// A default retry strategy applies to this operation CreateTenancyOverride()
func (client TenancyOverrideClient) CreateTenancyOverride(ctx context.Context, request CreateTenancyOverrideRequest) (response CreateTenancyOverrideResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTenancyOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTenancyOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTenancyOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTenancyOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTenancyOverrideResponse")
	}
	return
}

// createTenancyOverride implements the OCIOperation interface (enables retrying operations)
func (client TenancyOverrideClient) createTenancyOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tenancyOverrides", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTenancyOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyOverride/CreateTenancyOverride"
		err = common.PostProcessServiceError(err, "TenancyOverride", "CreateTenancyOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAllTenancyOverridesForTenant Deletes all Tenancy Overrides for Tenant.
// A default retry strategy applies to this operation DeleteAllTenancyOverridesForTenant()
func (client TenancyOverrideClient) DeleteAllTenancyOverridesForTenant(ctx context.Context, request DeleteAllTenancyOverridesForTenantRequest) (response DeleteAllTenancyOverridesForTenantResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAllTenancyOverridesForTenant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAllTenancyOverridesForTenantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAllTenancyOverridesForTenantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAllTenancyOverridesForTenantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAllTenancyOverridesForTenantResponse")
	}
	return
}

// deleteAllTenancyOverridesForTenant implements the OCIOperation interface (enables retrying operations)
func (client TenancyOverrideClient) deleteAllTenancyOverridesForTenant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tenancyOverrides/{tenantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAllTenancyOverridesForTenantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyOverride/DeleteAllTenancyOverridesForTenant"
		err = common.PostProcessServiceError(err, "TenancyOverride", "DeleteAllTenancyOverridesForTenant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTenancyOverride Deletes a Tenancy Override.
// A default retry strategy applies to this operation DeleteTenancyOverride()
func (client TenancyOverrideClient) DeleteTenancyOverride(ctx context.Context, request DeleteTenancyOverrideRequest) (response DeleteTenancyOverrideResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTenancyOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTenancyOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTenancyOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTenancyOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTenancyOverrideResponse")
	}
	return
}

// deleteTenancyOverride implements the OCIOperation interface (enables retrying operations)
func (client TenancyOverrideClient) deleteTenancyOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tenancyOverrides/{tenantId}/{overrideId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTenancyOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyOverride/DeleteTenancyOverride"
		err = common.PostProcessServiceError(err, "TenancyOverride", "DeleteTenancyOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTenancyOverride Gets information about a Tenancy Override.
// A default retry strategy applies to this operation GetTenancyOverride()
func (client TenancyOverrideClient) GetTenancyOverride(ctx context.Context, request GetTenancyOverrideRequest) (response GetTenancyOverrideResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTenancyOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTenancyOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTenancyOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTenancyOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTenancyOverrideResponse")
	}
	return
}

// getTenancyOverride implements the OCIOperation interface (enables retrying operations)
func (client TenancyOverrideClient) getTenancyOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancyOverrides/{tenantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTenancyOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyOverride/GetTenancyOverride"
		err = common.PostProcessServiceError(err, "TenancyOverride", "GetTenancyOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTenancyOverrides Gets a list of Tenancy Overrides.
// A default retry strategy applies to this operation ListTenancyOverrides()
func (client TenancyOverrideClient) ListTenancyOverrides(ctx context.Context, request ListTenancyOverridesRequest) (response ListTenancyOverridesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTenancyOverrides, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTenancyOverridesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTenancyOverridesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTenancyOverridesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTenancyOverridesResponse")
	}
	return
}

// listTenancyOverrides implements the OCIOperation interface (enables retrying operations)
func (client TenancyOverrideClient) listTenancyOverrides(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancyOverrides", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTenancyOverridesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyOverrideCollection/ListTenancyOverrides"
		err = common.PostProcessServiceError(err, "TenancyOverride", "ListTenancyOverrides", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTenancyOverride Updates a Tenancy Override.
// A default retry strategy applies to this operation UpdateTenancyOverride()
func (client TenancyOverrideClient) UpdateTenancyOverride(ctx context.Context, request UpdateTenancyOverrideRequest) (response UpdateTenancyOverrideResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTenancyOverride, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTenancyOverrideResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTenancyOverrideResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTenancyOverrideResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTenancyOverrideResponse")
	}
	return
}

// updateTenancyOverride implements the OCIOperation interface (enables retrying operations)
func (client TenancyOverrideClient) updateTenancyOverride(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/tenancyOverrides/{tenantId}/{overrideId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTenancyOverrideResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyOverride/UpdateTenancyOverride"
		err = common.PostProcessServiceError(err, "TenancyOverride", "UpdateTenancyOverride", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
