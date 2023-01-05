// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//LoggingManagementClient a client for LoggingManagement
type LoggingManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLoggingManagementClientWithConfigurationProvider Creates a new default LoggingManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLoggingManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LoggingManagementClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newLoggingManagementClientFromBaseClient(baseClient, provider)
}

// NewLoggingManagementClientWithOboToken Creates a new default LoggingManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLoggingManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LoggingManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLoggingManagementClientFromBaseClient(baseClient, configProvider)
}

func newLoggingManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LoggingManagementClient, err error) {
	// LoggingManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LoggingManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LoggingManagementClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LoggingManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("notification", "https://notification.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LoggingManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *LoggingManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetLogging Get resource-specific logging configuration.
// A default retry strategy applies to this operation GetLogging()
func (client LoggingManagementClient) GetLogging(ctx context.Context, request GetLoggingRequest) (response GetLoggingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLoggingResponse")
	}
	return
}

// getLogging implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) getLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logging/{logId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/PublicLoggingDetails/GetLogging"
		err = common.PostProcessServiceError(err, "LoggingManagement", "GetLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartLogging Enables logging for a resource.
// A default retry strategy applies to this operation StartLogging()
func (client LoggingManagementClient) StartLogging(ctx context.Context, request StartLoggingRequest) (response StartLoggingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartLoggingResponse")
	}
	return
}

// startLogging implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) startLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logging", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/PublicLoggingDetails/StartLogging"
		err = common.PostProcessServiceError(err, "LoggingManagement", "StartLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopLogging Disable a Logging resource by identifier
// A default retry strategy applies to this operation StopLogging()
func (client LoggingManagementClient) StopLogging(ctx context.Context, request StopLoggingRequest) (response StopLoggingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopLoggingResponse")
	}
	return
}

// stopLogging implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) stopLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/logging/{logId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/PublicLoggingDetails/StopLogging"
		err = common.PostProcessServiceError(err, "LoggingManagement", "StopLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogging Updates the category parameters in logging resource
// A default retry strategy applies to this operation UpdateLogging()
func (client LoggingManagementClient) UpdateLogging(ctx context.Context, request UpdateLoggingRequest) (response UpdateLoggingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLoggingResponse")
	}
	return
}

// updateLogging implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) updateLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/logging", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/notification/20181201/PublicLoggingDetails/UpdateLogging"
		err = common.PostProcessServiceError(err, "LoggingManagement", "UpdateLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
