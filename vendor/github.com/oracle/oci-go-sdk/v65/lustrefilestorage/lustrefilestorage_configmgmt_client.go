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

// ConfigMgmtClient a client for ConfigMgmt
type ConfigMgmtClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewConfigMgmtClientWithConfigurationProvider Creates a new default ConfigMgmt client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewConfigMgmtClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ConfigMgmtClient, err error) {
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
	return newConfigMgmtClientFromBaseClient(baseClient, provider)
}

// NewConfigMgmtClientWithOboToken Creates a new default ConfigMgmt client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewConfigMgmtClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ConfigMgmtClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newConfigMgmtClientFromBaseClient(baseClient, configProvider)
}

func newConfigMgmtClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ConfigMgmtClient, err error) {
	// ConfigMgmt service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ConfigMgmt"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ConfigMgmtClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ConfigMgmtClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("lustrefilestorage", "https://lustre-file-storage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ConfigMgmtClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ConfigMgmtClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DeleteTenancyConfiguration Delete the tenancy configuration mapping for given tenancy
// A default retry strategy applies to this operation DeleteTenancyConfiguration()
func (client ConfigMgmtClient) DeleteTenancyConfiguration(ctx context.Context, request DeleteTenancyConfigurationRequest) (response DeleteTenancyConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTenancyConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTenancyConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTenancyConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTenancyConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTenancyConfigurationResponse")
	}
	return
}

// deleteTenancyConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ConfigMgmtClient) deleteTenancyConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/configManagement/tenancyConfigurations/{tenancyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTenancyConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/TenancyConfigurationSummary/DeleteTenancyConfiguration"
		err = common.PostProcessServiceError(err, "ConfigMgmt", "DeleteTenancyConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfiles Retrieve all available profiles for the file system configuration.
// A default retry strategy applies to this operation ListProfiles()
func (client ConfigMgmtClient) ListProfiles(ctx context.Context, request ListProfilesRequest) (response ListProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfilesResponse")
	}
	return
}

// listProfiles implements the OCIOperation interface (enables retrying operations)
func (client ConfigMgmtClient) listProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configManagement/profiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/ProfileCollection/ListProfiles"
		err = common.PostProcessServiceError(err, "ConfigMgmt", "ListProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTenancyConfigurations Returns all the tenancy configuration along with profile and capacity reservation associated.
// A default retry strategy applies to this operation ListTenancyConfigurations()
func (client ConfigMgmtClient) ListTenancyConfigurations(ctx context.Context, request ListTenancyConfigurationsRequest) (response ListTenancyConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTenancyConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTenancyConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTenancyConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTenancyConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTenancyConfigurationsResponse")
	}
	return
}

// listTenancyConfigurations implements the OCIOperation interface (enables retrying operations)
func (client ConfigMgmtClient) listTenancyConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configManagement/tenancyConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTenancyConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/MapTenancyConfigurationCollection/ListTenancyConfigurations"
		err = common.PostProcessServiceError(err, "ConfigMgmt", "ListTenancyConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MapTenancyConfiguration Assigns given configuration to a tenancy.
// A default retry strategy applies to this operation MapTenancyConfiguration()
func (client ConfigMgmtClient) MapTenancyConfiguration(ctx context.Context, request MapTenancyConfigurationRequest) (response MapTenancyConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.mapTenancyConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MapTenancyConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MapTenancyConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MapTenancyConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MapTenancyConfigurationResponse")
	}
	return
}

// mapTenancyConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ConfigMgmtClient) mapTenancyConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configManagement/tenancyConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MapTenancyConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/lustre/20250228/CreateMapTenancyConfiguration/MapTenancyConfiguration"
		err = common.PostProcessServiceError(err, "ConfigMgmt", "MapTenancyConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
