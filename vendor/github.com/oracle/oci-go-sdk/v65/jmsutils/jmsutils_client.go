// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// JmsUtilsClient a client for JmsUtils
type JmsUtilsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewJmsUtilsClientWithConfigurationProvider Creates a new default JmsUtils client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewJmsUtilsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client JmsUtilsClient, err error) {
	if enabled := common.CheckForEnabledServices("jmsutils"); !enabled {
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
	return newJmsUtilsClientFromBaseClient(baseClient, provider)
}

// NewJmsUtilsClientWithOboToken Creates a new default JmsUtils client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewJmsUtilsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client JmsUtilsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newJmsUtilsClientFromBaseClient(baseClient, configProvider)
}

func newJmsUtilsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client JmsUtilsClient, err error) {
	// JmsUtils service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("JmsUtils"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = JmsUtilsClient{BaseClient: baseClient}
	client.BasePath = "20250521"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *JmsUtilsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("jmsutils", "https://javamanagement-utils.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *JmsUtilsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *JmsUtilsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client JmsUtilsClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "JmsUtils", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJavaMigrationAnalysis Deletes a Java Migration Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/DeleteJavaMigrationAnalysis.go.html to see an example of how to use DeleteJavaMigrationAnalysis API.
// A default retry strategy applies to this operation DeleteJavaMigrationAnalysis()
func (client JmsUtilsClient) DeleteJavaMigrationAnalysis(ctx context.Context, request DeleteJavaMigrationAnalysisRequest) (response DeleteJavaMigrationAnalysisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJavaMigrationAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJavaMigrationAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJavaMigrationAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJavaMigrationAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJavaMigrationAnalysisResponse")
	}
	return
}

// deleteJavaMigrationAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) deleteJavaMigrationAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/javaMigrationAnalysis/{javaMigrationAnalysisId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJavaMigrationAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/JavaMigrationAnalysis/DeleteJavaMigrationAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "DeleteJavaMigrationAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePerformanceTuningAnalysis Deletes a Performance Tuning Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/DeletePerformanceTuningAnalysis.go.html to see an example of how to use DeletePerformanceTuningAnalysis API.
// A default retry strategy applies to this operation DeletePerformanceTuningAnalysis()
func (client JmsUtilsClient) DeletePerformanceTuningAnalysis(ctx context.Context, request DeletePerformanceTuningAnalysisRequest) (response DeletePerformanceTuningAnalysisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePerformanceTuningAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePerformanceTuningAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePerformanceTuningAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePerformanceTuningAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePerformanceTuningAnalysisResponse")
	}
	return
}

// deletePerformanceTuningAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) deletePerformanceTuningAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/performanceTuningAnalysis/{performanceTuningAnalysisId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePerformanceTuningAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/PerformanceTuningAnalysis/DeletePerformanceTuningAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "DeletePerformanceTuningAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnalyzeApplicationsConfiguration Returns the configuration for analyzing applications.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/GetAnalyzeApplicationsConfiguration.go.html to see an example of how to use GetAnalyzeApplicationsConfiguration API.
// A default retry strategy applies to this operation GetAnalyzeApplicationsConfiguration()
func (client JmsUtilsClient) GetAnalyzeApplicationsConfiguration(ctx context.Context, request GetAnalyzeApplicationsConfigurationRequest) (response GetAnalyzeApplicationsConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnalyzeApplicationsConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnalyzeApplicationsConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnalyzeApplicationsConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnalyzeApplicationsConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnalyzeApplicationsConfigurationResponse")
	}
	return
}

// getAnalyzeApplicationsConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) getAnalyzeApplicationsConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyzeApplicationsConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnalyzeApplicationsConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/AnalyzeApplicationsConfiguration/GetAnalyzeApplicationsConfiguration"
		err = common.PostProcessServiceError(err, "JmsUtils", "GetAnalyzeApplicationsConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaMigrationAnalysis Gets information about a Java Migration Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/GetJavaMigrationAnalysis.go.html to see an example of how to use GetJavaMigrationAnalysis API.
// A default retry strategy applies to this operation GetJavaMigrationAnalysis()
func (client JmsUtilsClient) GetJavaMigrationAnalysis(ctx context.Context, request GetJavaMigrationAnalysisRequest) (response GetJavaMigrationAnalysisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaMigrationAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaMigrationAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaMigrationAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaMigrationAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaMigrationAnalysisResponse")
	}
	return
}

// getJavaMigrationAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) getJavaMigrationAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/javaMigrationAnalysis/{javaMigrationAnalysisId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaMigrationAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/JavaMigrationAnalysis/GetJavaMigrationAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "GetJavaMigrationAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPerformanceTuningAnalysis Gets information about a Performance Tuning Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/GetPerformanceTuningAnalysis.go.html to see an example of how to use GetPerformanceTuningAnalysis API.
// A default retry strategy applies to this operation GetPerformanceTuningAnalysis()
func (client JmsUtilsClient) GetPerformanceTuningAnalysis(ctx context.Context, request GetPerformanceTuningAnalysisRequest) (response GetPerformanceTuningAnalysisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPerformanceTuningAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPerformanceTuningAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPerformanceTuningAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPerformanceTuningAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPerformanceTuningAnalysisResponse")
	}
	return
}

// getPerformanceTuningAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) getPerformanceTuningAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/performanceTuningAnalysis/{performanceTuningAnalysisId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPerformanceTuningAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/PerformanceTuningAnalysis/GetPerformanceTuningAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "GetPerformanceTuningAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSubscriptionAcknowledgmentConfiguration Returns the configuration for subscription acknowledgment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/GetSubscriptionAcknowledgmentConfiguration.go.html to see an example of how to use GetSubscriptionAcknowledgmentConfiguration API.
// A default retry strategy applies to this operation GetSubscriptionAcknowledgmentConfiguration()
func (client JmsUtilsClient) GetSubscriptionAcknowledgmentConfiguration(ctx context.Context, request GetSubscriptionAcknowledgmentConfigurationRequest) (response GetSubscriptionAcknowledgmentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSubscriptionAcknowledgmentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSubscriptionAcknowledgmentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSubscriptionAcknowledgmentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSubscriptionAcknowledgmentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSubscriptionAcknowledgmentConfigurationResponse")
	}
	return
}

// getSubscriptionAcknowledgmentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) getSubscriptionAcknowledgmentConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptionAcknowledgmentConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSubscriptionAcknowledgmentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/SubscriptionAcknowledgmentConfiguration/GetSubscriptionAcknowledgmentConfiguration"
		err = common.PostProcessServiceError(err, "JmsUtils", "GetSubscriptionAcknowledgmentConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client JmsUtilsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "JmsUtils", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaMigrationAnalysis Gets a list of Java Migration Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListJavaMigrationAnalysis.go.html to see an example of how to use ListJavaMigrationAnalysis API.
// A default retry strategy applies to this operation ListJavaMigrationAnalysis()
func (client JmsUtilsClient) ListJavaMigrationAnalysis(ctx context.Context, request ListJavaMigrationAnalysisRequest) (response ListJavaMigrationAnalysisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaMigrationAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaMigrationAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaMigrationAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaMigrationAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaMigrationAnalysisResponse")
	}
	return
}

// listJavaMigrationAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) listJavaMigrationAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/javaMigrationAnalysis", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaMigrationAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/JavaMigrationAnalysis/ListJavaMigrationAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "ListJavaMigrationAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPerformanceTuningAnalysis Gets a list of Performance tuning Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListPerformanceTuningAnalysis.go.html to see an example of how to use ListPerformanceTuningAnalysis API.
// A default retry strategy applies to this operation ListPerformanceTuningAnalysis()
func (client JmsUtilsClient) ListPerformanceTuningAnalysis(ctx context.Context, request ListPerformanceTuningAnalysisRequest) (response ListPerformanceTuningAnalysisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPerformanceTuningAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPerformanceTuningAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPerformanceTuningAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPerformanceTuningAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPerformanceTuningAnalysisResponse")
	}
	return
}

// listPerformanceTuningAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) listPerformanceTuningAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/performanceTuningAnalysis", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPerformanceTuningAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/PerformanceTuningAnalysis/ListPerformanceTuningAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "ListPerformanceTuningAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkItems Retrieve a paginated list of work items for a specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListWorkItems.go.html to see an example of how to use ListWorkItems API.
// A default retry strategy applies to this operation ListWorkItems()
func (client JmsUtilsClient) ListWorkItems(ctx context.Context, request ListWorkItemsRequest) (response ListWorkItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkItemsResponse")
	}
	return
}

// listWorkItems implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) listWorkItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/workItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/WorkItemSummary/ListWorkItems"
		err = common.PostProcessServiceError(err, "JmsUtils", "ListWorkItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client JmsUtilsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "JmsUtils", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client JmsUtilsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "JmsUtils", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client JmsUtilsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "JmsUtils", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestJavaMigrationAnalysis Requests Java Migration Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/RequestJavaMigrationAnalysis.go.html to see an example of how to use RequestJavaMigrationAnalysis API.
// A default retry strategy applies to this operation RequestJavaMigrationAnalysis()
func (client JmsUtilsClient) RequestJavaMigrationAnalysis(ctx context.Context, request RequestJavaMigrationAnalysisRequest) (response RequestJavaMigrationAnalysisResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.requestJavaMigrationAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestJavaMigrationAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestJavaMigrationAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestJavaMigrationAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestJavaMigrationAnalysisResponse")
	}
	return
}

// requestJavaMigrationAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) requestJavaMigrationAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/javaMigrationAnalysis/actions/requestJavaMigrationAnalysis", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestJavaMigrationAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/JavaMigrationAnalysis/RequestJavaMigrationAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "RequestJavaMigrationAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestPerformanceTuningAnalysis Requests Performance Tuning Analysis.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/RequestPerformanceTuningAnalysis.go.html to see an example of how to use RequestPerformanceTuningAnalysis API.
// A default retry strategy applies to this operation RequestPerformanceTuningAnalysis()
func (client JmsUtilsClient) RequestPerformanceTuningAnalysis(ctx context.Context, request RequestPerformanceTuningAnalysisRequest) (response RequestPerformanceTuningAnalysisResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.requestPerformanceTuningAnalysis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestPerformanceTuningAnalysisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestPerformanceTuningAnalysisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestPerformanceTuningAnalysisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestPerformanceTuningAnalysisResponse")
	}
	return
}

// requestPerformanceTuningAnalysis implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) requestPerformanceTuningAnalysis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/performanceTuningAnalysis/actions/requestPerformanceTuningAnalysis", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestPerformanceTuningAnalysisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/PerformanceTuningAnalysis/RequestPerformanceTuningAnalysis"
		err = common.PostProcessServiceError(err, "JmsUtils", "RequestPerformanceTuningAnalysis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAnalyzeApplicationsConfiguration Updates the configuration for analyze application.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/UpdateAnalyzeApplicationsConfiguration.go.html to see an example of how to use UpdateAnalyzeApplicationsConfiguration API.
// A default retry strategy applies to this operation UpdateAnalyzeApplicationsConfiguration()
func (client JmsUtilsClient) UpdateAnalyzeApplicationsConfiguration(ctx context.Context, request UpdateAnalyzeApplicationsConfigurationRequest) (response UpdateAnalyzeApplicationsConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnalyzeApplicationsConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnalyzeApplicationsConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnalyzeApplicationsConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnalyzeApplicationsConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnalyzeApplicationsConfigurationResponse")
	}
	return
}

// updateAnalyzeApplicationsConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) updateAnalyzeApplicationsConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/analyzeApplicationsConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAnalyzeApplicationsConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/AnalyzeApplicationsConfiguration/UpdateAnalyzeApplicationsConfiguration"
		err = common.PostProcessServiceError(err, "JmsUtils", "UpdateAnalyzeApplicationsConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSubscriptionAcknowledgmentConfiguration Updates the configuration for subscription acknowledgment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/UpdateSubscriptionAcknowledgmentConfiguration.go.html to see an example of how to use UpdateSubscriptionAcknowledgmentConfiguration API.
// A default retry strategy applies to this operation UpdateSubscriptionAcknowledgmentConfiguration()
func (client JmsUtilsClient) UpdateSubscriptionAcknowledgmentConfiguration(ctx context.Context, request UpdateSubscriptionAcknowledgmentConfigurationRequest) (response UpdateSubscriptionAcknowledgmentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSubscriptionAcknowledgmentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSubscriptionAcknowledgmentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSubscriptionAcknowledgmentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSubscriptionAcknowledgmentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSubscriptionAcknowledgmentConfigurationResponse")
	}
	return
}

// updateSubscriptionAcknowledgmentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JmsUtilsClient) updateSubscriptionAcknowledgmentConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/subscriptionAcknowledgmentConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSubscriptionAcknowledgmentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-utils/20250521/SubscriptionAcknowledgmentConfiguration/UpdateSubscriptionAcknowledgmentConfiguration"
		err = common.PostProcessServiceError(err, "JmsUtils", "UpdateSubscriptionAcknowledgmentConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
