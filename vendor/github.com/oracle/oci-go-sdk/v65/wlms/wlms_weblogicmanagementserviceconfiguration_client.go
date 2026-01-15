// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// WeblogicManagementServiceConfigurationClient a client for WeblogicManagementServiceConfiguration
type WeblogicManagementServiceConfigurationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewWeblogicManagementServiceConfigurationClientWithConfigurationProvider Creates a new default WeblogicManagementServiceConfiguration client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewWeblogicManagementServiceConfigurationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client WeblogicManagementServiceConfigurationClient, err error) {
	if enabled := common.CheckForEnabledServices("wlms"); !enabled {
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
	return newWeblogicManagementServiceConfigurationClientFromBaseClient(baseClient, provider)
}

// NewWeblogicManagementServiceConfigurationClientWithOboToken Creates a new default WeblogicManagementServiceConfiguration client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewWeblogicManagementServiceConfigurationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client WeblogicManagementServiceConfigurationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newWeblogicManagementServiceConfigurationClientFromBaseClient(baseClient, configProvider)
}

func newWeblogicManagementServiceConfigurationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client WeblogicManagementServiceConfigurationClient, err error) {
	// WeblogicManagementServiceConfiguration service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("WeblogicManagementServiceConfiguration"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = WeblogicManagementServiceConfigurationClient{BaseClient: baseClient}
	client.BasePath = "20241101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *WeblogicManagementServiceConfigurationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("wlms", "https://api.weblogicmanagement.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *WeblogicManagementServiceConfigurationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *WeblogicManagementServiceConfigurationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetConfiguration Gets service configuration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetConfiguration.go.html to see an example of how to use GetConfiguration API.
// A default retry strategy applies to this operation GetConfiguration()
func (client WeblogicManagementServiceConfigurationClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationResponse")
	}
	return
}

// getConfiguration implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceConfigurationClient) getConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/Configuration/GetConfiguration"
		err = common.PostProcessServiceError(err, "WeblogicManagementServiceConfiguration", "GetConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfiguration Updates service configuration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/UpdateConfiguration.go.html to see an example of how to use UpdateConfiguration API.
// A default retry strategy applies to this operation UpdateConfiguration()
func (client WeblogicManagementServiceConfigurationClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationResponse")
	}
	return
}

// updateConfiguration implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceConfigurationClient) updateConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/Configuration/UpdateConfiguration"
		err = common.PostProcessServiceError(err, "WeblogicManagementServiceConfiguration", "UpdateConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
