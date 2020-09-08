// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
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
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newLoggingManagementClientFromBaseClient(baseClient, configProvider)
}

// NewLoggingManagementClientWithOboToken Creates a new default LoggingManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLoggingManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LoggingManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newLoggingManagementClientFromBaseClient(baseClient, configProvider)
}

func newLoggingManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LoggingManagementClient, err error) {
	client = LoggingManagementClient{BaseClient: baseClient}
	client.BasePath = "20200531"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LoggingManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("logging", "https://logging.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LoggingManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *LoggingManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeLogGroupCompartment Moves a log group into a different compartment within the same tenancy.  When provided, If-Match is checked against ETag values of the resource.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client LoggingManagementClient) ChangeLogGroupCompartment(ctx context.Context, request ChangeLogGroupCompartmentRequest) (response ChangeLogGroupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeLogGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogGroupCompartmentResponse")
	}
	return
}

// changeLogGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) changeLogGroupCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logGroups/{logGroupId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeLogGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLogLogGroup Moves a log into a different log group within the same tenancy.  When provided, If-Match is checked against ETag values of the resource.
func (client LoggingManagementClient) ChangeLogLogGroup(ctx context.Context, request ChangeLogLogGroupRequest) (response ChangeLogLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeLogLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogLogGroupResponse")
	}
	return
}

// changeLogLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) changeLogLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logGroups/{logGroupId}/logs/{logId}/actions/changeLogGroup")
	if err != nil {
		return nil, err
	}

	var response ChangeLogLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLogSavedSearchCompartment Moves a saved search into a different compartment within the same tenancy. For information about moving
// resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client LoggingManagementClient) ChangeLogSavedSearchCompartment(ctx context.Context, request ChangeLogSavedSearchCompartmentRequest) (response ChangeLogSavedSearchCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeLogSavedSearchCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogSavedSearchCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogSavedSearchCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogSavedSearchCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogSavedSearchCompartmentResponse")
	}
	return
}

// changeLogSavedSearchCompartment implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) changeLogSavedSearchCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logSavedSearches/{logSavedSearchId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeLogSavedSearchCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeUnifiedAgentConfigurationCompartment Moves unified agent configuration into a different compartment within the same tenancy.  When provided, If-Match is checked against ETag values of the resource.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client LoggingManagementClient) ChangeUnifiedAgentConfigurationCompartment(ctx context.Context, request ChangeUnifiedAgentConfigurationCompartmentRequest) (response ChangeUnifiedAgentConfigurationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeUnifiedAgentConfigurationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeUnifiedAgentConfigurationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeUnifiedAgentConfigurationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeUnifiedAgentConfigurationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeUnifiedAgentConfigurationCompartmentResponse")
	}
	return
}

// changeUnifiedAgentConfigurationCompartment implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) changeUnifiedAgentConfigurationCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/unifiedAgentConfigurations/{unifiedAgentConfigurationId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeUnifiedAgentConfigurationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLog Creates a log within specified log group. This call fails if log group is already created
// with same displayName or (service, resource, category) triplet.
func (client LoggingManagementClient) CreateLog(ctx context.Context, request CreateLogRequest) (response CreateLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogResponse")
	}
	return
}

// createLog implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) createLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logGroups/{logGroupId}/logs")
	if err != nil {
		return nil, err
	}

	var response CreateLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLogGroup Create new log group with unique display name. This call fails
// if log group is already created with same displayName in the compartment.
func (client LoggingManagementClient) CreateLogGroup(ctx context.Context, request CreateLogGroupRequest) (response CreateLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogGroupResponse")
	}
	return
}

// createLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) createLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logGroups")
	if err != nil {
		return nil, err
	}

	var response CreateLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLogSavedSearch Creates a new LogSavedSearch.
func (client LoggingManagementClient) CreateLogSavedSearch(ctx context.Context, request CreateLogSavedSearchRequest) (response CreateLogSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogSavedSearchResponse")
	}
	return
}

// createLogSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) createLogSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logSavedSearches")
	if err != nil {
		return nil, err
	}

	var response CreateLogSavedSearchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUnifiedAgentConfiguration Create unified agent config registration
func (client LoggingManagementClient) CreateUnifiedAgentConfiguration(ctx context.Context, request CreateUnifiedAgentConfigurationRequest) (response CreateUnifiedAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createUnifiedAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUnifiedAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUnifiedAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUnifiedAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUnifiedAgentConfigurationResponse")
	}
	return
}

// createUnifiedAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) createUnifiedAgentConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/unifiedAgentConfigurations")
	if err != nil {
		return nil, err
	}

	var response CreateUnifiedAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLog Deletes the log object in a log group.
func (client LoggingManagementClient) DeleteLog(ctx context.Context, request DeleteLogRequest) (response DeleteLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogResponse")
	}
	return
}

// deleteLog implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) deleteLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/logGroups/{logGroupId}/logs/{logId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLogGroup Deletes the specified log group.
func (client LoggingManagementClient) DeleteLogGroup(ctx context.Context, request DeleteLogGroupRequest) (response DeleteLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogGroupResponse")
	}
	return
}

// deleteLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) deleteLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/logGroups/{logGroupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLogSavedSearch Deletes the specified log saved search.
func (client LoggingManagementClient) DeleteLogSavedSearch(ctx context.Context, request DeleteLogSavedSearchRequest) (response DeleteLogSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogSavedSearchResponse")
	}
	return
}

// deleteLogSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) deleteLogSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/logSavedSearches/{logSavedSearchId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogSavedSearchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUnifiedAgentConfiguration Delete unified agent configuration
func (client LoggingManagementClient) DeleteUnifiedAgentConfiguration(ctx context.Context, request DeleteUnifiedAgentConfigurationRequest) (response DeleteUnifiedAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUnifiedAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUnifiedAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUnifiedAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUnifiedAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUnifiedAgentConfigurationResponse")
	}
	return
}

// deleteUnifiedAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) deleteUnifiedAgentConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/unifiedAgentConfigurations/{unifiedAgentConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response DeleteUnifiedAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkRequest Cancel a work request that has not started yet.
func (client LoggingManagementClient) DeleteWorkRequest(ctx context.Context, request DeleteWorkRequestRequest) (response DeleteWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkRequestResponse")
	}
	return
}

// deleteWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) deleteWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response DeleteWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLog Gets the log object config for log object OCID.
func (client LoggingManagementClient) GetLog(ctx context.Context, request GetLogRequest) (response GetLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogResponse")
	}
	return
}

// getLog implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) getLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logGroups/{logGroupId}/logs/{logId}")
	if err != nil {
		return nil, err
	}

	var response GetLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogGroup Get the specified log group's information.
func (client LoggingManagementClient) GetLogGroup(ctx context.Context, request GetLogGroupRequest) (response GetLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogGroupResponse")
	}
	return
}

// getLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) getLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logGroups/{logGroupId}")
	if err != nil {
		return nil, err
	}

	var response GetLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogIncludedSearch Retrieves a LogIncludedSearch.
func (client LoggingManagementClient) GetLogIncludedSearch(ctx context.Context, request GetLogIncludedSearchRequest) (response GetLogIncludedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogIncludedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogIncludedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogIncludedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogIncludedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogIncludedSearchResponse")
	}
	return
}

// getLogIncludedSearch implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) getLogIncludedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logIncludedSearch/{logIncludedSearchId}")
	if err != nil {
		return nil, err
	}

	var response GetLogIncludedSearchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogSavedSearch Retrieves a log saved search.
func (client LoggingManagementClient) GetLogSavedSearch(ctx context.Context, request GetLogSavedSearchRequest) (response GetLogSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogSavedSearchResponse")
	}
	return
}

// getLogSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) getLogSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logSavedSearches/{logSavedSearchId}")
	if err != nil {
		return nil, err
	}

	var response GetLogSavedSearchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUnifiedAgentConfiguration Get unified agent configuration for an id
func (client LoggingManagementClient) GetUnifiedAgentConfiguration(ctx context.Context, request GetUnifiedAgentConfigurationRequest) (response GetUnifiedAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUnifiedAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUnifiedAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUnifiedAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUnifiedAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUnifiedAgentConfigurationResponse")
	}
	return
}

// getUnifiedAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) getUnifiedAgentConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/unifiedAgentConfigurations/{unifiedAgentConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response GetUnifiedAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of the work request with the given ID.
func (client LoggingManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LoggingManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogGroups Lists all log groups for the specified compartment or tenancy.
func (client LoggingManagementClient) ListLogGroups(ctx context.Context, request ListLogGroupsRequest) (response ListLogGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogGroupsResponse")
	}
	return
}

// listLogGroups implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) listLogGroups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logGroups")
	if err != nil {
		return nil, err
	}

	var response ListLogGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogIncludedSearches Lists Logging Included Searches for this compartment.
func (client LoggingManagementClient) ListLogIncludedSearches(ctx context.Context, request ListLogIncludedSearchesRequest) (response ListLogIncludedSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogIncludedSearches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogIncludedSearchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogIncludedSearchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogIncludedSearchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogIncludedSearchesResponse")
	}
	return
}

// listLogIncludedSearches implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) listLogIncludedSearches(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logIncludedSearches")
	if err != nil {
		return nil, err
	}

	var response ListLogIncludedSearchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogSavedSearches Lists Logging Saved Searches for this compartment.
func (client LoggingManagementClient) ListLogSavedSearches(ctx context.Context, request ListLogSavedSearchesRequest) (response ListLogSavedSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogSavedSearches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogSavedSearchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogSavedSearchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogSavedSearchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogSavedSearchesResponse")
	}
	return
}

// listLogSavedSearches implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) listLogSavedSearches(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logSavedSearches")
	if err != nil {
		return nil, err
	}

	var response ListLogSavedSearchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogs Lists the specified log group's log objects.
func (client LoggingManagementClient) ListLogs(ctx context.Context, request ListLogsRequest) (response ListLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogsResponse")
	}
	return
}

// listLogs implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) listLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/logGroups/{logGroupId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServices Lists all services supporting logging.
func (client LoggingManagementClient) ListServices(ctx context.Context, request ListServicesRequest) (response ListServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServicesResponse")
	}
	return
}

// listServices implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) listServices(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/registry/services")
	if err != nil {
		return nil, err
	}

	var response ListServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUnifiedAgentConfigurations Lists all unified agent configurations in the specified compartment
func (client LoggingManagementClient) ListUnifiedAgentConfigurations(ctx context.Context, request ListUnifiedAgentConfigurationsRequest) (response ListUnifiedAgentConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUnifiedAgentConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUnifiedAgentConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUnifiedAgentConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUnifiedAgentConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUnifiedAgentConfigurationsResponse")
	}
	return
}

// listUnifiedAgentConfigurations implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) listUnifiedAgentConfigurations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/unifiedAgentConfigurations")
	if err != nil {
		return nil, err
	}

	var response ListUnifiedAgentConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a list of errors for a given work request.
func (client LoggingManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LoggingManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a list of logs for a given work request.
func (client LoggingManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LoggingManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
func (client LoggingManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LoggingManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLog Updates existing log object with the associated config. This call
//       fails if log object does not exist.
func (client LoggingManagementClient) UpdateLog(ctx context.Context, request UpdateLogRequest) (response UpdateLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogResponse")
	}
	return
}

// updateLog implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) updateLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/logGroups/{logGroupId}/logs/{logId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogGroup Updates existing log group with the associated config. This call
//       fails if log group does not exist.
func (client LoggingManagementClient) UpdateLogGroup(ctx context.Context, request UpdateLogGroupRequest) (response UpdateLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogGroupResponse")
	}
	return
}

// updateLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) updateLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/logGroups/{logGroupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogSavedSearch Updates an  existing log saved search.
func (client LoggingManagementClient) UpdateLogSavedSearch(ctx context.Context, request UpdateLogSavedSearchRequest) (response UpdateLogSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogSavedSearchResponse")
	}
	return
}

// updateLogSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) updateLogSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/logSavedSearches/{logSavedSearchId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogSavedSearchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUnifiedAgentConfiguration Update an existing unified agent configuration. This call
//       fails if log group does not exist.
func (client LoggingManagementClient) UpdateUnifiedAgentConfiguration(ctx context.Context, request UpdateUnifiedAgentConfigurationRequest) (response UpdateUnifiedAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUnifiedAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUnifiedAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUnifiedAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUnifiedAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUnifiedAgentConfigurationResponse")
	}
	return
}

// updateUnifiedAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client LoggingManagementClient) updateUnifiedAgentConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/unifiedAgentConfigurations/{unifiedAgentConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response UpdateUnifiedAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
