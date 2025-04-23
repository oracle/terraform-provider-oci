// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ConfigClient a client for Config
type ConfigClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewConfigClientWithConfigurationProvider Creates a new default Config client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewConfigClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ConfigClient, err error) {
	if enabled := common.CheckForEnabledServices("apmconfig"); !enabled {
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
	return newConfigClientFromBaseClient(baseClient, provider)
}

// NewConfigClientWithOboToken Creates a new default Config client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewConfigClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ConfigClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newConfigClientFromBaseClient(baseClient, configProvider)
}

func newConfigClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ConfigClient, err error) {
	// Config service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Config"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ConfigClient{BaseClient: baseClient}
	client.BasePath = "20210201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ConfigClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmconfig", "https://apm-config.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ConfigClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ConfigClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CopyConfiguration Fast importing configuration items to a destination APM domain ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/CopyConfiguration.go.html to see an example of how to use CopyConfiguration API.
// A default retry strategy applies to this operation CopyConfiguration()
func (client ConfigClient) CopyConfiguration(ctx context.Context, request CopyConfigurationRequest) (response CopyConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.copyConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyConfigurationResponse")
	}
	return
}

// copyConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) copyConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/copyConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/ExportConfigurationDetails/CopyConfiguration"
		err = common.PostProcessServiceError(err, "Config", "CopyConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConfig Creates a new configuration item.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/CreateConfig.go.html to see an example of how to use CreateConfig API.
// A default retry strategy applies to this operation CreateConfig()
func (client ConfigClient) CreateConfig(ctx context.Context, request CreateConfigRequest) (response CreateConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConfigResponse")
	}
	return
}

// createConfig implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) createConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/Config/CreateConfig"
		err = common.PostProcessServiceError(err, "Config", "CreateConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &config{})
	return response, err
}

// DeleteConfig Deletes the configuration item identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/DeleteConfig.go.html to see an example of how to use DeleteConfig API.
// A default retry strategy applies to this operation DeleteConfig()
func (client ConfigClient) DeleteConfig(ctx context.Context, request DeleteConfigRequest) (response DeleteConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConfigResponse")
	}
	return
}

// deleteConfig implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) deleteConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/configs/{configId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/Config/DeleteConfig"
		err = common.PostProcessServiceError(err, "Config", "DeleteConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportConfiguration Exports configurations for the whole domain by domainId.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/ExportConfiguration.go.html to see an example of how to use ExportConfiguration API.
// A default retry strategy applies to this operation ExportConfiguration()
func (client ConfigClient) ExportConfiguration(ctx context.Context, request ExportConfigurationRequest) (response ExportConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportConfigurationResponse")
	}
	return
}

// exportConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) exportConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/exportConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/ExportConfigurationDetails/ExportConfiguration"
		err = common.PostProcessServiceError(err, "Config", "ExportConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfig Gets the configuration item identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/GetConfig.go.html to see an example of how to use GetConfig API.
// A default retry strategy applies to this operation GetConfig()
func (client ConfigClient) GetConfig(ctx context.Context, request GetConfigRequest) (response GetConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigResponse")
	}
	return
}

// getConfig implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) getConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configs/{configId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/Config/GetConfig"
		err = common.PostProcessServiceError(err, "Config", "GetConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &config{})
	return response, err
}

// ImportConfiguration Import configurations Item(s) with its dependencies into a destination domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/ImportConfiguration.go.html to see an example of how to use ImportConfiguration API.
// A default retry strategy applies to this operation ImportConfiguration()
func (client ConfigClient) ImportConfiguration(ctx context.Context, request ImportConfigurationRequest) (response ImportConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportConfigurationResponse")
	}
	return
}

// importConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) importConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/importConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/ImportConfigurationDetails/ImportConfiguration"
		err = common.PostProcessServiceError(err, "Config", "ImportConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConfigs Returns all configuration items, which can optionally be filtered by configuration type.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/ListConfigs.go.html to see an example of how to use ListConfigs API.
// A default retry strategy applies to this operation ListConfigs()
func (client ConfigClient) ListConfigs(ctx context.Context, request ListConfigsRequest) (response ListConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigsResponse")
	}
	return
}

// listConfigs implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) listConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/ConfigCollection/ListConfigs"
		err = common.PostProcessServiceError(err, "Config", "ListConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveNamespaceMetrics Returns all metrics associated with the specified namespace.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/RetrieveNamespaceMetrics.go.html to see an example of how to use RetrieveNamespaceMetrics API.
// A default retry strategy applies to this operation RetrieveNamespaceMetrics()
func (client ConfigClient) RetrieveNamespaceMetrics(ctx context.Context, request RetrieveNamespaceMetricsRequest) (response RetrieveNamespaceMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveNamespaceMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveNamespaceMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveNamespaceMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveNamespaceMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveNamespaceMetricsResponse")
	}
	return
}

// retrieveNamespaceMetrics implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) retrieveNamespaceMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/retrieveNamespaceMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveNamespaceMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/MetricGroup/RetrieveNamespaceMetrics"
		err = common.PostProcessServiceError(err, "Config", "RetrieveNamespaceMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveNamespaces Returns all namespaces available in APM.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/RetrieveNamespaces.go.html to see an example of how to use RetrieveNamespaces API.
// A default retry strategy applies to this operation RetrieveNamespaces()
func (client ConfigClient) RetrieveNamespaces(ctx context.Context, request RetrieveNamespacesRequest) (response RetrieveNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveNamespacesResponse")
	}
	return
}

// retrieveNamespaces implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) retrieveNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/retrieveNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveNamespacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/MetricGroup/RetrieveNamespaces"
		err = common.PostProcessServiceError(err, "Config", "RetrieveNamespaces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Test Tests a data processing operation on the provided input, returning the potentially modified
// input as output. Returns 200 on success, 422 when the input can not be processed.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/Test.go.html to see an example of how to use Test API.
// A default retry strategy applies to this operation Test()
func (client ConfigClient) Test(ctx context.Context, request TestRequest) (response TestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.test, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestResponse")
	}
	return
}

// test implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) test(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/TestOutput/Test"
		err = common.PostProcessServiceError(err, "Config", "Test", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &testoutput{})
	return response, err
}

// UpdateConfig Updates the details of the configuration item identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/UpdateConfig.go.html to see an example of how to use UpdateConfig API.
// A default retry strategy applies to this operation UpdateConfig()
func (client ConfigClient) UpdateConfig(ctx context.Context, request UpdateConfigRequest) (response UpdateConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigResponse")
	}
	return
}

// updateConfig implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) updateConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configs/{configId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/Config/UpdateConfig"
		err = common.PostProcessServiceError(err, "Config", "UpdateConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &config{})
	return response, err
}

// ValidateSpanFilterPattern Validates the Span Filter pattern (filterText) for syntactic correctness.
// Returns 204 on success, 422 when validation fails.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/ValidateSpanFilterPattern.go.html to see an example of how to use ValidateSpanFilterPattern API.
// A default retry strategy applies to this operation ValidateSpanFilterPattern()
func (client ConfigClient) ValidateSpanFilterPattern(ctx context.Context, request ValidateSpanFilterPatternRequest) (response ValidateSpanFilterPatternResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateSpanFilterPattern, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSpanFilterPatternResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSpanFilterPatternResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSpanFilterPatternResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSpanFilterPatternResponse")
	}
	return
}

// validateSpanFilterPattern implements the OCIOperation interface (enables retrying operations)
func (client ConfigClient) validateSpanFilterPattern(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/validateSpanFilterPattern", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateSpanFilterPatternResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-config/20210201/SpanFilter/ValidateSpanFilterPattern"
		err = common.PostProcessServiceError(err, "Config", "ValidateSpanFilterPattern", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
