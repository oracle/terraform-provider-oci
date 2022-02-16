// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//LogAnalyticsClient a client for LogAnalytics
type LogAnalyticsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLogAnalyticsClientWithConfigurationProvider Creates a new default LogAnalytics client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLogAnalyticsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LogAnalyticsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newLogAnalyticsClientFromBaseClient(baseClient, provider)
}

// NewLogAnalyticsClientWithOboToken Creates a new default LogAnalytics client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLogAnalyticsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LogAnalyticsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLogAnalyticsClientFromBaseClient(baseClient, configProvider)
}

func newLogAnalyticsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LogAnalyticsClient, err error) {
	// LogAnalytics service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LogAnalyticsClient{BaseClient: baseClient}
	client.BasePath = "20200601"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LogAnalyticsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("loganalytics", "https://loganalytics.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LogAnalyticsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LogAnalyticsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddEntityAssociation Adds association between input source log analytics entity and one or more existing destination entities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/AddEntityAssociation.go.html to see an example of how to use AddEntityAssociation API.
func (client LogAnalyticsClient) AddEntityAssociation(ctx context.Context, request AddEntityAssociationRequest) (response AddEntityAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addEntityAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddEntityAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddEntityAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddEntityAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddEntityAssociationResponse")
	}
	return
}

// addEntityAssociation implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) addEntityAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/actions/addEntityAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddEntityAssociationResponse
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

// AddSourceEventTypes Add one or more event types to a source. An event type and version can be enabled only on one source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/AddSourceEventTypes.go.html to see an example of how to use AddSourceEventTypes API.
func (client LogAnalyticsClient) AddSourceEventTypes(ctx context.Context, request AddSourceEventTypesRequest) (response AddSourceEventTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addSourceEventTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddSourceEventTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddSourceEventTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddSourceEventTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddSourceEventTypesResponse")
	}
	return
}

// addSourceEventTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) addSourceEventTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/{sourceName}/actions/addEventTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddSourceEventTypesResponse
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

// AppendLookupData Appends data to the lookup content. The csv file containing the content to be appended is passed in as binary data in the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/AppendLookupData.go.html to see an example of how to use AppendLookupData API.
func (client LogAnalyticsClient) AppendLookupData(ctx context.Context, request AppendLookupDataRequest) (response AppendLookupDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.appendLookupData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AppendLookupDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AppendLookupDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AppendLookupDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AppendLookupDataResponse")
	}
	return
}

// appendLookupData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) appendLookupData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/lookups/{lookupName}/actions/appendData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AppendLookupDataResponse
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

// BatchGetBasicInfo Lists basic information about a specified set of labels in batch.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/BatchGetBasicInfo.go.html to see an example of how to use BatchGetBasicInfo API.
func (client LogAnalyticsClient) BatchGetBasicInfo(ctx context.Context, request BatchGetBasicInfoRequest) (response BatchGetBasicInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchGetBasicInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchGetBasicInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchGetBasicInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchGetBasicInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchGetBasicInfoResponse")
	}
	return
}

// batchGetBasicInfo implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) batchGetBasicInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/labels/actions/basicInfo", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BatchGetBasicInfoResponse
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

// CancelQueryWorkRequest Cancel/Remove query job work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CancelQueryWorkRequest.go.html to see an example of how to use CancelQueryWorkRequest API.
func (client LogAnalyticsClient) CancelQueryWorkRequest(ctx context.Context, request CancelQueryWorkRequestRequest) (response CancelQueryWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelQueryWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelQueryWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelQueryWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelQueryWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelQueryWorkRequestResponse")
	}
	return
}

// cancelQueryWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) cancelQueryWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/queryWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelQueryWorkRequestResponse
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

// ChangeLogAnalyticsEmBridgeCompartment Update the compartment of the log analytics enterprise manager bridge with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ChangeLogAnalyticsEmBridgeCompartment.go.html to see an example of how to use ChangeLogAnalyticsEmBridgeCompartment API.
func (client LogAnalyticsClient) ChangeLogAnalyticsEmBridgeCompartment(ctx context.Context, request ChangeLogAnalyticsEmBridgeCompartmentRequest) (response ChangeLogAnalyticsEmBridgeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsEmBridgeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsEmBridgeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsEmBridgeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsEmBridgeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsEmBridgeCompartmentResponse")
	}
	return
}

// changeLogAnalyticsEmBridgeCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsEmBridgeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEmBridges/{logAnalyticsEmBridgeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsEmBridgeCompartmentResponse
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

// ChangeLogAnalyticsEntityCompartment Update the compartment of the log analytics entity with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ChangeLogAnalyticsEntityCompartment.go.html to see an example of how to use ChangeLogAnalyticsEntityCompartment API.
func (client LogAnalyticsClient) ChangeLogAnalyticsEntityCompartment(ctx context.Context, request ChangeLogAnalyticsEntityCompartmentRequest) (response ChangeLogAnalyticsEntityCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsEntityCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsEntityCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsEntityCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsEntityCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsEntityCompartmentResponse")
	}
	return
}

// changeLogAnalyticsEntityCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsEntityCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsEntityCompartmentResponse
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

// ChangeLogAnalyticsLogGroupCompartment Moves the specified log group to a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ChangeLogAnalyticsLogGroupCompartment.go.html to see an example of how to use ChangeLogAnalyticsLogGroupCompartment API.
func (client LogAnalyticsClient) ChangeLogAnalyticsLogGroupCompartment(ctx context.Context, request ChangeLogAnalyticsLogGroupCompartmentRequest) (response ChangeLogAnalyticsLogGroupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsLogGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsLogGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsLogGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsLogGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsLogGroupCompartmentResponse")
	}
	return
}

// changeLogAnalyticsLogGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsLogGroupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsLogGroupCompartmentResponse
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

// ChangeLogAnalyticsObjectCollectionRuleCompartment Move the rule from it's current compartment to the given compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ChangeLogAnalyticsObjectCollectionRuleCompartment.go.html to see an example of how to use ChangeLogAnalyticsObjectCollectionRuleCompartment API.
func (client LogAnalyticsClient) ChangeLogAnalyticsObjectCollectionRuleCompartment(ctx context.Context, request ChangeLogAnalyticsObjectCollectionRuleCompartmentRequest) (response ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsObjectCollectionRuleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse")
	}
	return
}

// changeLogAnalyticsObjectCollectionRuleCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsObjectCollectionRuleCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse
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

// ChangeScheduledTaskCompartment Move the scheduled task into a different compartment within the same tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ChangeScheduledTaskCompartment.go.html to see an example of how to use ChangeScheduledTaskCompartment API.
func (client LogAnalyticsClient) ChangeScheduledTaskCompartment(ctx context.Context, request ChangeScheduledTaskCompartmentRequest) (response ChangeScheduledTaskCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeScheduledTaskCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeScheduledTaskCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeScheduledTaskCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeScheduledTaskCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeScheduledTaskCompartmentResponse")
	}
	return
}

// changeScheduledTaskCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeScheduledTaskCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeScheduledTaskCompartmentResponse
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

// Clean Clean accumulated acceleration data stored for the accelerated saved search.
// The ScheduledTask taskType must be ACCELERATION.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/Clean.go.html to see an example of how to use Clean API.
func (client LogAnalyticsClient) Clean(ctx context.Context, request CleanRequest) (response CleanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.clean, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CleanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CleanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CleanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CleanResponse")
	}
	return
}

// clean implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) clean(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/clean", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CleanResponse
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

// CompareContent Returns the difference between the two input payloads, including intraline differences.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CompareContent.go.html to see an example of how to use CompareContent API.
func (client LogAnalyticsClient) CompareContent(ctx context.Context, request CompareContentRequest) (response CompareContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.compareContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompareContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompareContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompareContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompareContentResponse")
	}
	return
}

// compareContent implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) compareContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/compareContent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompareContentResponse
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

// CreateLogAnalyticsEmBridge Add configuration for enterprise manager bridge. Enterprise manager bridge is used to automatically add selected entities from enterprise manager cloud control. A corresponding OCI bridge configuration is required in enterprise manager.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CreateLogAnalyticsEmBridge.go.html to see an example of how to use CreateLogAnalyticsEmBridge API.
func (client LogAnalyticsClient) CreateLogAnalyticsEmBridge(ctx context.Context, request CreateLogAnalyticsEmBridgeRequest) (response CreateLogAnalyticsEmBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsEmBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsEmBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsEmBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsEmBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsEmBridgeResponse")
	}
	return
}

// createLogAnalyticsEmBridge implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsEmBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEmBridges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsEmBridgeResponse
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

// CreateLogAnalyticsEntity Create a new log analytics entity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CreateLogAnalyticsEntity.go.html to see an example of how to use CreateLogAnalyticsEntity API.
func (client LogAnalyticsClient) CreateLogAnalyticsEntity(ctx context.Context, request CreateLogAnalyticsEntityRequest) (response CreateLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsEntityResponse")
	}
	return
}

// createLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsEntityResponse
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

// CreateLogAnalyticsEntityType Add custom log analytics entity type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CreateLogAnalyticsEntityType.go.html to see an example of how to use CreateLogAnalyticsEntityType API.
func (client LogAnalyticsClient) CreateLogAnalyticsEntityType(ctx context.Context, request CreateLogAnalyticsEntityTypeRequest) (response CreateLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsEntityTypeResponse")
	}
	return
}

// createLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntityTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsEntityTypeResponse
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

// CreateLogAnalyticsLogGroup Creates a new log group in the specified compartment with the input display name. You may also specify optional information such as description, defined tags, and free-form tags.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CreateLogAnalyticsLogGroup.go.html to see an example of how to use CreateLogAnalyticsLogGroup API.
func (client LogAnalyticsClient) CreateLogAnalyticsLogGroup(ctx context.Context, request CreateLogAnalyticsLogGroupRequest) (response CreateLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsLogGroupResponse")
	}
	return
}

// createLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsLogGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsLogGroupResponse
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

// CreateLogAnalyticsObjectCollectionRule Creates a rule to collect logs from an object storage bucket.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CreateLogAnalyticsObjectCollectionRule.go.html to see an example of how to use CreateLogAnalyticsObjectCollectionRule API.
func (client LogAnalyticsClient) CreateLogAnalyticsObjectCollectionRule(ctx context.Context, request CreateLogAnalyticsObjectCollectionRuleRequest) (response CreateLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// createLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsObjectCollectionRuleResponse
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

// CreateScheduledTask Schedule a task as specified and return task info.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/CreateScheduledTask.go.html to see an example of how to use CreateScheduledTask API.
func (client LogAnalyticsClient) CreateScheduledTask(ctx context.Context, request CreateScheduledTaskRequest) (response CreateScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduledTaskResponse")
	}
	return
}

// createScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createScheduledTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &scheduledtask{})
	return response, err
}

// DeleteAssociations Deletes the associations between the sources and entities specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteAssociations.go.html to see an example of how to use DeleteAssociations API.
func (client LogAnalyticsClient) DeleteAssociations(ctx context.Context, request DeleteAssociationsRequest) (response DeleteAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAssociationsResponse")
	}
	return
}

// deleteAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/associations/actions/delete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAssociationsResponse
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

// DeleteField Deletes field with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteField.go.html to see an example of how to use DeleteField API.
func (client LogAnalyticsClient) DeleteField(ctx context.Context, request DeleteFieldRequest) (response DeleteFieldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteField, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFieldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFieldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFieldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFieldResponse")
	}
	return
}

// deleteField implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteField(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/fields/{fieldName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFieldResponse
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

// DeleteLabel Deletes label with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLabel.go.html to see an example of how to use DeleteLabel API.
func (client LogAnalyticsClient) DeleteLabel(ctx context.Context, request DeleteLabelRequest) (response DeleteLabelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteLabel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLabelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLabelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLabelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLabelResponse")
	}
	return
}

// deleteLabel implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLabel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/labels/{labelName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLabelResponse
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

// DeleteLogAnalyticsEmBridge Delete log analytics enterprise manager bridge with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLogAnalyticsEmBridge.go.html to see an example of how to use DeleteLogAnalyticsEmBridge API.
func (client LogAnalyticsClient) DeleteLogAnalyticsEmBridge(ctx context.Context, request DeleteLogAnalyticsEmBridgeRequest) (response DeleteLogAnalyticsEmBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsEmBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsEmBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsEmBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsEmBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsEmBridgeResponse")
	}
	return
}

// deleteLogAnalyticsEmBridge implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsEmBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsEmBridges/{logAnalyticsEmBridgeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsEmBridgeResponse
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

// DeleteLogAnalyticsEntity Delete log analytics entity with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLogAnalyticsEntity.go.html to see an example of how to use DeleteLogAnalyticsEntity API.
func (client LogAnalyticsClient) DeleteLogAnalyticsEntity(ctx context.Context, request DeleteLogAnalyticsEntityRequest) (response DeleteLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsEntityResponse")
	}
	return
}

// deleteLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsEntityResponse
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

// DeleteLogAnalyticsEntityType Delete log analytics entity type with the given name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLogAnalyticsEntityType.go.html to see an example of how to use DeleteLogAnalyticsEntityType API.
func (client LogAnalyticsClient) DeleteLogAnalyticsEntityType(ctx context.Context, request DeleteLogAnalyticsEntityTypeRequest) (response DeleteLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsEntityTypeResponse")
	}
	return
}

// deleteLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsEntityTypes/{entityTypeName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsEntityTypeResponse
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

// DeleteLogAnalyticsLogGroup Deletes the specified log group. The log group cannot be part of an active association or have an active upload.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLogAnalyticsLogGroup.go.html to see an example of how to use DeleteLogAnalyticsLogGroup API.
func (client LogAnalyticsClient) DeleteLogAnalyticsLogGroup(ctx context.Context, request DeleteLogAnalyticsLogGroupRequest) (response DeleteLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsLogGroupResponse")
	}
	return
}

// deleteLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsLogGroupResponse
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

// DeleteLogAnalyticsObjectCollectionRule Deletes the configured object storage bucket based collection rule and stop the log collection.
// It will not delete the existing processed data associated with this bucket from logging analytics storage.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLogAnalyticsObjectCollectionRule.go.html to see an example of how to use DeleteLogAnalyticsObjectCollectionRule API.
func (client LogAnalyticsClient) DeleteLogAnalyticsObjectCollectionRule(ctx context.Context, request DeleteLogAnalyticsObjectCollectionRuleRequest) (response DeleteLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// deleteLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsObjectCollectionRuleResponse
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

// DeleteLookup Deletes lookup with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteLookup.go.html to see an example of how to use DeleteLookup API.
func (client LogAnalyticsClient) DeleteLookup(ctx context.Context, request DeleteLookupRequest) (response DeleteLookupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteLookup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLookupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLookupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLookupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLookupResponse")
	}
	return
}

// deleteLookup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLookup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/lookups/{lookupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLookupResponse
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

// DeleteParser Deletes parser with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteParser.go.html to see an example of how to use DeleteParser API.
func (client LogAnalyticsClient) DeleteParser(ctx context.Context, request DeleteParserRequest) (response DeleteParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteParserResponse")
	}
	return
}

// deleteParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteParser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/parsers/{parserName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteParserResponse
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

// DeleteScheduledTask Delete the scheduled task.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteScheduledTask.go.html to see an example of how to use DeleteScheduledTask API.
func (client LogAnalyticsClient) DeleteScheduledTask(ctx context.Context, request DeleteScheduledTaskRequest) (response DeleteScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduledTaskResponse")
	}
	return
}

// deleteScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteScheduledTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduledTaskResponse
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

// DeleteSource Deletes source with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteSource.go.html to see an example of how to use DeleteSource API.
func (client LogAnalyticsClient) DeleteSource(ctx context.Context, request DeleteSourceRequest) (response DeleteSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSourceResponse")
	}
	return
}

// deleteSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/sources/{sourceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSourceResponse
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

// DeleteUpload Deletes an Upload by its reference.
// It deletes all the logs in storage asscoiated with the upload and the corresponding upload metadata.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteUpload.go.html to see an example of how to use DeleteUpload API.
func (client LogAnalyticsClient) DeleteUpload(ctx context.Context, request DeleteUploadRequest) (response DeleteUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUploadResponse")
	}
	return
}

// deleteUpload implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/uploads/{uploadReference}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUploadResponse
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

// DeleteUploadFile Deletes a specific log file inside an upload by upload file reference.
// It deletes all the logs from storage associated with the file and the corresponding metadata.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteUploadFile.go.html to see an example of how to use DeleteUploadFile API.
func (client LogAnalyticsClient) DeleteUploadFile(ctx context.Context, request DeleteUploadFileRequest) (response DeleteUploadFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUploadFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUploadFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUploadFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUploadFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUploadFileResponse")
	}
	return
}

// deleteUploadFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteUploadFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/uploads/{uploadReference}/files/{fileReference}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUploadFileResponse
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

// DeleteUploadWarning Suppresses a specific warning inside an upload.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DeleteUploadWarning.go.html to see an example of how to use DeleteUploadWarning API.
func (client LogAnalyticsClient) DeleteUploadWarning(ctx context.Context, request DeleteUploadWarningRequest) (response DeleteUploadWarningResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUploadWarning, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUploadWarningResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUploadWarningResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUploadWarningResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUploadWarningResponse")
	}
	return
}

// deleteUploadWarning implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteUploadWarning(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/uploads/{uploadReference}/warnings/{warningReference}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUploadWarningResponse
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

// DisableArchiving This API disables archiving.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DisableArchiving.go.html to see an example of how to use DisableArchiving API.
func (client LogAnalyticsClient) DisableArchiving(ctx context.Context, request DisableArchivingRequest) (response DisableArchivingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableArchiving, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableArchivingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableArchivingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableArchivingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableArchivingResponse")
	}
	return
}

// disableArchiving implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) disableArchiving(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/disableArchiving", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableArchivingResponse
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

// DisableAutoAssociation Disables auto association for a log source. In the future, this log source would not be automatically
// associated with any entity that becomes eligible for association. In addition, you may also optionally
// remove all existing associations for this log source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DisableAutoAssociation.go.html to see an example of how to use DisableAutoAssociation API.
func (client LogAnalyticsClient) DisableAutoAssociation(ctx context.Context, request DisableAutoAssociationRequest) (response DisableAutoAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.disableAutoAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAutoAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAutoAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAutoAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAutoAssociationResponse")
	}
	return
}

// disableAutoAssociation implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) disableAutoAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/{sourceName}/actions/disableAutoAssociation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAutoAssociationResponse
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

// DisableSourceEventTypes Disable one or more event types in a source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/DisableSourceEventTypes.go.html to see an example of how to use DisableSourceEventTypes API.
func (client LogAnalyticsClient) DisableSourceEventTypes(ctx context.Context, request DisableSourceEventTypesRequest) (response DisableSourceEventTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.disableSourceEventTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableSourceEventTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableSourceEventTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableSourceEventTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableSourceEventTypesResponse")
	}
	return
}

// disableSourceEventTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) disableSourceEventTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/{sourceName}/actions/disableEventTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableSourceEventTypesResponse
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

// EnableArchiving THis API enables archiving.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/EnableArchiving.go.html to see an example of how to use EnableArchiving API.
func (client LogAnalyticsClient) EnableArchiving(ctx context.Context, request EnableArchivingRequest) (response EnableArchivingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableArchiving, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableArchivingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableArchivingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableArchivingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableArchivingResponse")
	}
	return
}

// enableArchiving implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) enableArchiving(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/enableArchiving", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableArchivingResponse
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

// EnableAutoAssociation Enables auto association for a log source. This would initiate immediate association of the source
// to any eligible entities it is not already associated with, and would also ensure the log source gets
// associated with entities that are added or become eligible in the future.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/EnableAutoAssociation.go.html to see an example of how to use EnableAutoAssociation API.
func (client LogAnalyticsClient) EnableAutoAssociation(ctx context.Context, request EnableAutoAssociationRequest) (response EnableAutoAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.enableAutoAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAutoAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAutoAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAutoAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAutoAssociationResponse")
	}
	return
}

// enableAutoAssociation implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) enableAutoAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/{sourceName}/actions/enableAutoAssociation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAutoAssociationResponse
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

// EnableSourceEventTypes Enable one or more event types in a source. An event type and version can be enabled only in one source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/EnableSourceEventTypes.go.html to see an example of how to use EnableSourceEventTypes API.
func (client LogAnalyticsClient) EnableSourceEventTypes(ctx context.Context, request EnableSourceEventTypesRequest) (response EnableSourceEventTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.enableSourceEventTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableSourceEventTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableSourceEventTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableSourceEventTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableSourceEventTypesResponse")
	}
	return
}

// enableSourceEventTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) enableSourceEventTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/{sourceName}/actions/enableEventTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableSourceEventTypesResponse
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

// EstimatePurgeDataSize This API estimates the size of data to be purged based based on time interval, purge query etc.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/EstimatePurgeDataSize.go.html to see an example of how to use EstimatePurgeDataSize API.
func (client LogAnalyticsClient) EstimatePurgeDataSize(ctx context.Context, request EstimatePurgeDataSizeRequest) (response EstimatePurgeDataSizeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.estimatePurgeDataSize, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EstimatePurgeDataSizeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EstimatePurgeDataSizeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EstimatePurgeDataSizeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EstimatePurgeDataSizeResponse")
	}
	return
}

// estimatePurgeDataSize implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) estimatePurgeDataSize(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/estimatePurgeDataSize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EstimatePurgeDataSizeResponse
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

// EstimateRecallDataSize This API gives an active storage usage estimate for archived data to be recalled and the time range of such data.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/EstimateRecallDataSize.go.html to see an example of how to use EstimateRecallDataSize API.
func (client LogAnalyticsClient) EstimateRecallDataSize(ctx context.Context, request EstimateRecallDataSizeRequest) (response EstimateRecallDataSizeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.estimateRecallDataSize, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EstimateRecallDataSizeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EstimateRecallDataSizeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EstimateRecallDataSizeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EstimateRecallDataSizeResponse")
	}
	return
}

// estimateRecallDataSize implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) estimateRecallDataSize(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/estimateRecallDataSize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EstimateRecallDataSizeResponse
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

// EstimateReleaseDataSize This API gives an active storage usage estimate for recalled data to be released and the time range of such data.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/EstimateReleaseDataSize.go.html to see an example of how to use EstimateReleaseDataSize API.
func (client LogAnalyticsClient) EstimateReleaseDataSize(ctx context.Context, request EstimateReleaseDataSizeRequest) (response EstimateReleaseDataSizeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.estimateReleaseDataSize, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EstimateReleaseDataSizeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EstimateReleaseDataSizeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EstimateReleaseDataSizeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EstimateReleaseDataSizeResponse")
	}
	return
}

// estimateReleaseDataSize implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) estimateReleaseDataSize(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/estimateReleaseDataSize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EstimateReleaseDataSizeResponse
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

// ExportCustomContent Exports all custom details of the specified sources, parsers, fields and labels, in zip format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ExportCustomContent.go.html to see an example of how to use ExportCustomContent API.
func (client LogAnalyticsClient) ExportCustomContent(ctx context.Context, request ExportCustomContentRequest) (response ExportCustomContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.exportCustomContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportCustomContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportCustomContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportCustomContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportCustomContentResponse")
	}
	return
}

// exportCustomContent implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) exportCustomContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/contents/actions/exportCustomContent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportCustomContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportQueryResult Export data based on query. Endpoint returns a stream of data. Endpoint is synchronous. Queries must deliver first result within 60 seconds or calls are subject to timeout.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ExportQueryResult.go.html to see an example of how to use ExportQueryResult API.
func (client LogAnalyticsClient) ExportQueryResult(ctx context.Context, request ExportQueryResultRequest) (response ExportQueryResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.exportQueryResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportQueryResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportQueryResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportQueryResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportQueryResultResponse")
	}
	return
}

// exportQueryResult implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) exportQueryResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportQueryResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExtractStructuredLogFieldPaths Extracts the field paths from the example json or xml content.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ExtractStructuredLogFieldPaths.go.html to see an example of how to use ExtractStructuredLogFieldPaths API.
func (client LogAnalyticsClient) ExtractStructuredLogFieldPaths(ctx context.Context, request ExtractStructuredLogFieldPathsRequest) (response ExtractStructuredLogFieldPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.extractStructuredLogFieldPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExtractStructuredLogFieldPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExtractStructuredLogFieldPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExtractStructuredLogFieldPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExtractStructuredLogFieldPathsResponse")
	}
	return
}

// extractStructuredLogFieldPaths implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) extractStructuredLogFieldPaths(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/extractLogFieldPaths", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExtractStructuredLogFieldPathsResponse
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

// ExtractStructuredLogHeaderPaths Extracts the header paths from the example json or xml content.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ExtractStructuredLogHeaderPaths.go.html to see an example of how to use ExtractStructuredLogHeaderPaths API.
func (client LogAnalyticsClient) ExtractStructuredLogHeaderPaths(ctx context.Context, request ExtractStructuredLogHeaderPathsRequest) (response ExtractStructuredLogHeaderPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.extractStructuredLogHeaderPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExtractStructuredLogHeaderPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExtractStructuredLogHeaderPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExtractStructuredLogHeaderPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExtractStructuredLogHeaderPathsResponse")
	}
	return
}

// extractStructuredLogHeaderPaths implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) extractStructuredLogHeaderPaths(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/extractLogHeaderPaths", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExtractStructuredLogHeaderPathsResponse
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

// Filter Each filter specifies an operator, a field and one or more values to be inserted into the provided query as criteria.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/Filter.go.html to see an example of how to use Filter API.
func (client LogAnalyticsClient) Filter(ctx context.Context, request FilterRequest) (response FilterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.filter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FilterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FilterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FilterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FilterResponse")
	}
	return
}

// filter implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) filter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/filter", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FilterResponse
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

// GetAssociationSummary Returns the count of source associations for entities in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetAssociationSummary.go.html to see an example of how to use GetAssociationSummary API.
func (client LogAnalyticsClient) GetAssociationSummary(ctx context.Context, request GetAssociationSummaryRequest) (response GetAssociationSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssociationSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssociationSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssociationSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssociationSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssociationSummaryResponse")
	}
	return
}

// getAssociationSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getAssociationSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/associationSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssociationSummaryResponse
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

// GetCategory Gets detailed information about the category with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetCategory.go.html to see an example of how to use GetCategory API.
func (client LogAnalyticsClient) GetCategory(ctx context.Context, request GetCategoryRequest) (response GetCategoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCategory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCategoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCategoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCategoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCategoryResponse")
	}
	return
}

// getCategory implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getCategory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/categories/{categoryName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCategoryResponse
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

// GetColumnNames Extracts column names from the input SQL query.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetColumnNames.go.html to see an example of how to use GetColumnNames API.
func (client LogAnalyticsClient) GetColumnNames(ctx context.Context, request GetColumnNamesRequest) (response GetColumnNamesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getColumnNames, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetColumnNamesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetColumnNamesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetColumnNamesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetColumnNamesResponse")
	}
	return
}

// getColumnNames implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getColumnNames(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/sqlColumnNames", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetColumnNamesResponse
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

// GetConfigWorkRequest Returns detailed information about the configuration work request with the specified id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetConfigWorkRequest.go.html to see an example of how to use GetConfigWorkRequest API.
func (client LogAnalyticsClient) GetConfigWorkRequest(ctx context.Context, request GetConfigWorkRequestRequest) (response GetConfigWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfigWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigWorkRequestResponse")
	}
	return
}

// getConfigWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getConfigWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/configWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigWorkRequestResponse
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

// GetField Gets detailed information about the field with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetField.go.html to see an example of how to use GetField API.
func (client LogAnalyticsClient) GetField(ctx context.Context, request GetFieldRequest) (response GetFieldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getField, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFieldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFieldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFieldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFieldResponse")
	}
	return
}

// getField implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getField(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/fields/{fieldName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFieldResponse
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

// GetFieldsSummary Returns the count of fields. You may optionally specify isShowDetail=true to view a summary of each field data type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetFieldsSummary.go.html to see an example of how to use GetFieldsSummary API.
func (client LogAnalyticsClient) GetFieldsSummary(ctx context.Context, request GetFieldsSummaryRequest) (response GetFieldsSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFieldsSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFieldsSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFieldsSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFieldsSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFieldsSummaryResponse")
	}
	return
}

// getFieldsSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getFieldsSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/fieldSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFieldsSummaryResponse
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

// GetLabel Gets detailed information about the label with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLabel.go.html to see an example of how to use GetLabel API.
func (client LogAnalyticsClient) GetLabel(ctx context.Context, request GetLabelRequest) (response GetLabelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLabel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLabelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLabelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLabelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLabelResponse")
	}
	return
}

// getLabel implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLabel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labels/{labelName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLabelResponse
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

// GetLabelSummary Returns the count of labels.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLabelSummary.go.html to see an example of how to use GetLabelSummary API.
func (client LogAnalyticsClient) GetLabelSummary(ctx context.Context, request GetLabelSummaryRequest) (response GetLabelSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLabelSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLabelSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLabelSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLabelSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLabelSummaryResponse")
	}
	return
}

// getLabelSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLabelSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labelSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLabelSummaryResponse
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

// GetLogAnalyticsEmBridge Retrieve the log analytics enterprise manager bridge with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsEmBridge.go.html to see an example of how to use GetLogAnalyticsEmBridge API.
func (client LogAnalyticsClient) GetLogAnalyticsEmBridge(ctx context.Context, request GetLogAnalyticsEmBridgeRequest) (response GetLogAnalyticsEmBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEmBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEmBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEmBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEmBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEmBridgeResponse")
	}
	return
}

// getLogAnalyticsEmBridge implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEmBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEmBridges/{logAnalyticsEmBridgeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEmBridgeResponse
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

// GetLogAnalyticsEmBridgeSummary Returns log analytics enterprise manager bridges summary report.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsEmBridgeSummary.go.html to see an example of how to use GetLogAnalyticsEmBridgeSummary API.
func (client LogAnalyticsClient) GetLogAnalyticsEmBridgeSummary(ctx context.Context, request GetLogAnalyticsEmBridgeSummaryRequest) (response GetLogAnalyticsEmBridgeSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEmBridgeSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEmBridgeSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEmBridgeSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEmBridgeSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEmBridgeSummaryResponse")
	}
	return
}

// getLogAnalyticsEmBridgeSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEmBridgeSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEmBridges/emBridgeSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEmBridgeSummaryResponse
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

// GetLogAnalyticsEntitiesSummary Returns log analytics entities count summary report.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsEntitiesSummary.go.html to see an example of how to use GetLogAnalyticsEntitiesSummary API.
func (client LogAnalyticsClient) GetLogAnalyticsEntitiesSummary(ctx context.Context, request GetLogAnalyticsEntitiesSummaryRequest) (response GetLogAnalyticsEntitiesSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEntitiesSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEntitiesSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEntitiesSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEntitiesSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEntitiesSummaryResponse")
	}
	return
}

// getLogAnalyticsEntitiesSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEntitiesSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/entitySummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEntitiesSummaryResponse
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

// GetLogAnalyticsEntity Retrieve the log analytics entity with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsEntity.go.html to see an example of how to use GetLogAnalyticsEntity API.
func (client LogAnalyticsClient) GetLogAnalyticsEntity(ctx context.Context, request GetLogAnalyticsEntityRequest) (response GetLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEntityResponse")
	}
	return
}

// getLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEntityResponse
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

// GetLogAnalyticsEntityType Retrieve the log analytics entity type with the given name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsEntityType.go.html to see an example of how to use GetLogAnalyticsEntityType API.
func (client LogAnalyticsClient) GetLogAnalyticsEntityType(ctx context.Context, request GetLogAnalyticsEntityTypeRequest) (response GetLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEntityTypeResponse")
	}
	return
}

// getLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntityTypes/{entityTypeName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEntityTypeResponse
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

// GetLogAnalyticsLogGroup Gets detailed information about the specified log group such as display name, description, defined tags, and free-form tags.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsLogGroup.go.html to see an example of how to use GetLogAnalyticsLogGroup API.
func (client LogAnalyticsClient) GetLogAnalyticsLogGroup(ctx context.Context, request GetLogAnalyticsLogGroupRequest) (response GetLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsLogGroupResponse")
	}
	return
}

// getLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsLogGroupResponse
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

// GetLogAnalyticsLogGroupsSummary Returns the count of log groups in a compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsLogGroupsSummary.go.html to see an example of how to use GetLogAnalyticsLogGroupsSummary API.
func (client LogAnalyticsClient) GetLogAnalyticsLogGroupsSummary(ctx context.Context, request GetLogAnalyticsLogGroupsSummaryRequest) (response GetLogAnalyticsLogGroupsSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsLogGroupsSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsLogGroupsSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsLogGroupsSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsLogGroupsSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsLogGroupsSummaryResponse")
	}
	return
}

// getLogAnalyticsLogGroupsSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsLogGroupsSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsLogGroupsSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsLogGroupsSummaryResponse
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

// GetLogAnalyticsObjectCollectionRule Gets a configured object storage based collection rule by given id
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsObjectCollectionRule.go.html to see an example of how to use GetLogAnalyticsObjectCollectionRule API.
func (client LogAnalyticsClient) GetLogAnalyticsObjectCollectionRule(ctx context.Context, request GetLogAnalyticsObjectCollectionRuleRequest) (response GetLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// getLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsObjectCollectionRuleResponse
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

// GetLogSetsCount This API returns the count of distinct log sets.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogSetsCount.go.html to see an example of how to use GetLogSetsCount API.
func (client LogAnalyticsClient) GetLogSetsCount(ctx context.Context, request GetLogSetsCountRequest) (response GetLogSetsCountResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogSetsCount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogSetsCountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogSetsCountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogSetsCountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogSetsCountResponse")
	}
	return
}

// getLogSetsCount implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogSetsCount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage/logSetsCount", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLogSetsCountResponse
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

// GetLookup Gets detailed information about the lookup with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLookup.go.html to see an example of how to use GetLookup API.
func (client LogAnalyticsClient) GetLookup(ctx context.Context, request GetLookupRequest) (response GetLookupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLookup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLookupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLookupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLookupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLookupResponse")
	}
	return
}

// getLookup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLookup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/lookups/{lookupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLookupResponse
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

// GetLookupSummary Returns the count of user created and oracle defined lookups.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLookupSummary.go.html to see an example of how to use GetLookupSummary API.
func (client LogAnalyticsClient) GetLookupSummary(ctx context.Context, request GetLookupSummaryRequest) (response GetLookupSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLookupSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLookupSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLookupSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLookupSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLookupSummaryResponse")
	}
	return
}

// getLookupSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLookupSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/lookupSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLookupSummaryResponse
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

// GetNamespace This API gets the namespace details of a tenancy already onboarded in Logging Analytics Application
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetNamespace.go.html to see an example of how to use GetNamespace API.
func (client LogAnalyticsClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceResponse")
	}
	return
}

// getNamespace implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNamespaceResponse
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

// GetParser Gets detailed information about the parser with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetParser.go.html to see an example of how to use GetParser API.
func (client LogAnalyticsClient) GetParser(ctx context.Context, request GetParserRequest) (response GetParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetParserResponse")
	}
	return
}

// getParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getParser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parsers/{parserName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetParserResponse
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

// GetParserSummary Returns the count of parsers.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetParserSummary.go.html to see an example of how to use GetParserSummary API.
func (client LogAnalyticsClient) GetParserSummary(ctx context.Context, request GetParserSummaryRequest) (response GetParserSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getParserSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetParserSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetParserSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetParserSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetParserSummaryResponse")
	}
	return
}

// getParserSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getParserSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parsersSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetParserSummaryResponse
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

// GetPreferences Lists the preferences of the tenant. Currently, only "DEFAULT_HOMEPAGE" is supported.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetPreferences.go.html to see an example of how to use GetPreferences API.
func (client LogAnalyticsClient) GetPreferences(ctx context.Context, request GetPreferencesRequest) (response GetPreferencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPreferences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPreferencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPreferencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPreferencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPreferencesResponse")
	}
	return
}

// getPreferences implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getPreferences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/preferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPreferencesResponse
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

// GetQueryResult Returns the intermediate results for a query that was specified to run asynchronously if the query has not completed,
// otherwise the final query results identified by a queryWorkRequestId returned when submitting the query execute asynchronously.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetQueryResult.go.html to see an example of how to use GetQueryResult API.
func (client LogAnalyticsClient) GetQueryResult(ctx context.Context, request GetQueryResultRequest) (response GetQueryResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQueryResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQueryResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQueryResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQueryResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQueryResultResponse")
	}
	return
}

// getQueryResult implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getQueryResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/search/actions/query", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetQueryResultResponse
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

// GetQueryWorkRequest Retrieve work request details by workRequestId. This endpoint can be polled for status tracking of work request. Clients should poll using the interval returned in the retry-after header.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetQueryWorkRequest.go.html to see an example of how to use GetQueryWorkRequest API.
func (client LogAnalyticsClient) GetQueryWorkRequest(ctx context.Context, request GetQueryWorkRequestRequest) (response GetQueryWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQueryWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQueryWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQueryWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQueryWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQueryWorkRequestResponse")
	}
	return
}

// getQueryWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getQueryWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/queryWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetQueryWorkRequestResponse
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

// GetScheduledTask Get the scheduled task for the specified task identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetScheduledTask.go.html to see an example of how to use GetScheduledTask API.
func (client LogAnalyticsClient) GetScheduledTask(ctx context.Context, request GetScheduledTaskRequest) (response GetScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledTaskResponse")
	}
	return
}

// getScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getScheduledTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &scheduledtask{})
	return response, err
}

// GetSource Gets detailed information about the source with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetSource.go.html to see an example of how to use GetSource API.
func (client LogAnalyticsClient) GetSource(ctx context.Context, request GetSourceRequest) (response GetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSourceResponse")
	}
	return
}

// getSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSourceResponse
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

// GetSourceSummary Returns the count of sources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetSourceSummary.go.html to see an example of how to use GetSourceSummary API.
func (client LogAnalyticsClient) GetSourceSummary(ctx context.Context, request GetSourceSummaryRequest) (response GetSourceSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSourceSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSourceSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSourceSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSourceSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSourceSummaryResponse")
	}
	return
}

// getSourceSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getSourceSummary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSourceSummaryResponse
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

// GetStorage This API gets the storage configuration of a tenancy
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetStorage.go.html to see an example of how to use GetStorage API.
func (client LogAnalyticsClient) GetStorage(ctx context.Context, request GetStorageRequest) (response GetStorageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStorageResponse")
	}
	return
}

// getStorage implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getStorage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStorageResponse
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

// GetStorageUsage This API gets storage usage information of a tenancy.  Storage usage information includes active, archived or recalled
// data.  The unit of return data is in bytes.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetStorageUsage.go.html to see an example of how to use GetStorageUsage API.
func (client LogAnalyticsClient) GetStorageUsage(ctx context.Context, request GetStorageUsageRequest) (response GetStorageUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStorageUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStorageUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStorageUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStorageUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStorageUsageResponse")
	}
	return
}

// getStorageUsage implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getStorageUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage/usage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStorageUsageResponse
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

// GetStorageWorkRequest This API returns work request details specified by {workRequestId}. This API can be polled for status tracking of
// work request.  Clients should poll using the interval returned in retry-after header.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetStorageWorkRequest.go.html to see an example of how to use GetStorageWorkRequest API.
func (client LogAnalyticsClient) GetStorageWorkRequest(ctx context.Context, request GetStorageWorkRequestRequest) (response GetStorageWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStorageWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStorageWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStorageWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStorageWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStorageWorkRequestResponse")
	}
	return
}

// getStorageWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getStorageWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storageWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStorageWorkRequestResponse
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

// GetUnprocessedDataBucket This API retrieves details of the configured bucket that stores unprocessed payloads.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetUnprocessedDataBucket.go.html to see an example of how to use GetUnprocessedDataBucket API.
func (client LogAnalyticsClient) GetUnprocessedDataBucket(ctx context.Context, request GetUnprocessedDataBucketRequest) (response GetUnprocessedDataBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUnprocessedDataBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUnprocessedDataBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUnprocessedDataBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUnprocessedDataBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUnprocessedDataBucketResponse")
	}
	return
}

// getUnprocessedDataBucket implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getUnprocessedDataBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/unprocessedDataBucket", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUnprocessedDataBucketResponse
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

// GetUpload Gets an On-Demand Upload info by reference.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetUpload.go.html to see an example of how to use GetUpload API.
func (client LogAnalyticsClient) GetUpload(ctx context.Context, request GetUploadRequest) (response GetUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUploadResponse")
	}
	return
}

// getUpload implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads/{uploadReference}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUploadResponse
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

// GetWorkRequest Gets the status of the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client LogAnalyticsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LogAnalyticsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
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

// ImportCustomContent Imports the specified custom content from the input in zip format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ImportCustomContent.go.html to see an example of how to use ImportCustomContent API.
func (client LogAnalyticsClient) ImportCustomContent(ctx context.Context, request ImportCustomContentRequest) (response ImportCustomContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importCustomContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportCustomContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportCustomContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportCustomContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportCustomContentResponse")
	}
	return
}

// importCustomContent implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) importCustomContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/contents/actions/importCustomContent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportCustomContentResponse
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

// ListAssociableEntities Lists the entities in the specified compartment which are (in)eligible for association with this source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAssociableEntities.go.html to see an example of how to use ListAssociableEntities API.
func (client LogAnalyticsClient) ListAssociableEntities(ctx context.Context, request ListAssociableEntitiesRequest) (response ListAssociableEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssociableEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssociableEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssociableEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssociableEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssociableEntitiesResponse")
	}
	return
}

// listAssociableEntities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listAssociableEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/associableEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssociableEntitiesResponse
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

// ListAssociatedEntities Lists the association details of entities in the specified compartment that are associated with at least one source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAssociatedEntities.go.html to see an example of how to use ListAssociatedEntities API.
func (client LogAnalyticsClient) ListAssociatedEntities(ctx context.Context, request ListAssociatedEntitiesRequest) (response ListAssociatedEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssociatedEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssociatedEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssociatedEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssociatedEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssociatedEntitiesResponse")
	}
	return
}

// listAssociatedEntities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listAssociatedEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/associatedEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssociatedEntitiesResponse
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

// ListAutoAssociations Gets information related to auto association for the source with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAutoAssociations.go.html to see an example of how to use ListAutoAssociations API.
func (client LogAnalyticsClient) ListAutoAssociations(ctx context.Context, request ListAutoAssociationsRequest) (response ListAutoAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutoAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutoAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutoAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutoAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutoAssociationsResponse")
	}
	return
}

// listAutoAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listAutoAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/autoAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutoAssociationsResponse
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

// ListCategories Returns a list of categories, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as category name or description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListCategories.go.html to see an example of how to use ListCategories API.
func (client LogAnalyticsClient) ListCategories(ctx context.Context, request ListCategoriesRequest) (response ListCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCategoriesResponse")
	}
	return
}

// listCategories implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/categories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCategoriesResponse
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

// ListConfigWorkRequests Returns the list of configuration work requests such as association or lookup operations, containing detailed information about them. You may paginate or limit the number of results.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListConfigWorkRequests.go.html to see an example of how to use ListConfigWorkRequests API.
func (client LogAnalyticsClient) ListConfigWorkRequests(ctx context.Context, request ListConfigWorkRequestsRequest) (response ListConfigWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConfigWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConfigWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigWorkRequestsResponse")
	}
	return
}

// listConfigWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listConfigWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/configWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConfigWorkRequestsResponse
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

// ListEntityAssociations Return a list of log analytics entities associated with input source log analytics entity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListEntityAssociations.go.html to see an example of how to use ListEntityAssociations API.
func (client LogAnalyticsClient) ListEntityAssociations(ctx context.Context, request ListEntityAssociationsRequest) (response ListEntityAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntityAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntityAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntityAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntityAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntityAssociationsResponse")
	}
	return
}

// listEntityAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listEntityAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/entityAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEntityAssociationsResponse
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

// ListEntitySourceAssociations Returns the list of source associations for the specified entity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListEntitySourceAssociations.go.html to see an example of how to use ListEntitySourceAssociations API.
func (client LogAnalyticsClient) ListEntitySourceAssociations(ctx context.Context, request ListEntitySourceAssociationsRequest) (response ListEntitySourceAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntitySourceAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntitySourceAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntitySourceAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntitySourceAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntitySourceAssociationsResponse")
	}
	return
}

// listEntitySourceAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listEntitySourceAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/entityAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEntitySourceAssociationsResponse
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

// ListFields Returns a list of log fields, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by specifying various options including parser and source names.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListFields.go.html to see an example of how to use ListFields API.
func (client LogAnalyticsClient) ListFields(ctx context.Context, request ListFieldsRequest) (response ListFieldsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFields, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFieldsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFieldsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFieldsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFieldsResponse")
	}
	return
}

// listFields implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listFields(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/fields", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFieldsResponse
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

// ListLabelPriorities Lists the available problem priorities that could be associated with a label.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLabelPriorities.go.html to see an example of how to use ListLabelPriorities API.
func (client LogAnalyticsClient) ListLabelPriorities(ctx context.Context, request ListLabelPrioritiesRequest) (response ListLabelPrioritiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLabelPriorities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLabelPrioritiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLabelPrioritiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLabelPrioritiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLabelPrioritiesResponse")
	}
	return
}

// listLabelPriorities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLabelPriorities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labelPriorities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLabelPrioritiesResponse
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

// ListLabelSourceDetails Lists sources using the label, along with configuration details like base field, operator and condition.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLabelSourceDetails.go.html to see an example of how to use ListLabelSourceDetails API.
func (client LogAnalyticsClient) ListLabelSourceDetails(ctx context.Context, request ListLabelSourceDetailsRequest) (response ListLabelSourceDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLabelSourceDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLabelSourceDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLabelSourceDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLabelSourceDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLabelSourceDetailsResponse")
	}
	return
}

// listLabelSourceDetails implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLabelSourceDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labelSourceDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLabelSourceDetailsResponse
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

// ListLabels Returns a list of labels, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as label name, display name, description and priority.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLabels.go.html to see an example of how to use ListLabels API.
func (client LogAnalyticsClient) ListLabels(ctx context.Context, request ListLabelsRequest) (response ListLabelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLabels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLabelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLabelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLabelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLabelsResponse")
	}
	return
}

// listLabels implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLabels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLabelsResponse
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

// ListLogAnalyticsEmBridges Return a list of log analytics enterprise manager bridges.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEmBridges.go.html to see an example of how to use ListLogAnalyticsEmBridges API.
func (client LogAnalyticsClient) ListLogAnalyticsEmBridges(ctx context.Context, request ListLogAnalyticsEmBridgesRequest) (response ListLogAnalyticsEmBridgesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsEmBridges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsEmBridgesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsEmBridgesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsEmBridgesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsEmBridgesResponse")
	}
	return
}

// listLogAnalyticsEmBridges implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsEmBridges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEmBridges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsEmBridgesResponse
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

// ListLogAnalyticsEntities Return a list of log analytics entities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEntities.go.html to see an example of how to use ListLogAnalyticsEntities API.
func (client LogAnalyticsClient) ListLogAnalyticsEntities(ctx context.Context, request ListLogAnalyticsEntitiesRequest) (response ListLogAnalyticsEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsEntitiesResponse")
	}
	return
}

// listLogAnalyticsEntities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsEntitiesResponse
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

// ListLogAnalyticsEntityTopology Return a log analytics entity topology collection that contains a set of log analytics entities and a set of relationships between those, for the input source entity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEntityTopology.go.html to see an example of how to use ListLogAnalyticsEntityTopology API.
func (client LogAnalyticsClient) ListLogAnalyticsEntityTopology(ctx context.Context, request ListLogAnalyticsEntityTopologyRequest) (response ListLogAnalyticsEntityTopologyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsEntityTopology, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsEntityTopologyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsEntityTopologyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsEntityTopologyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsEntityTopologyResponse")
	}
	return
}

// listLogAnalyticsEntityTopology implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsEntityTopology(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/entityTopology", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsEntityTopologyResponse
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

// ListLogAnalyticsEntityTypes Return a list of log analytics entity types.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEntityTypes.go.html to see an example of how to use ListLogAnalyticsEntityTypes API.
func (client LogAnalyticsClient) ListLogAnalyticsEntityTypes(ctx context.Context, request ListLogAnalyticsEntityTypesRequest) (response ListLogAnalyticsEntityTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsEntityTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsEntityTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsEntityTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsEntityTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsEntityTypesResponse")
	}
	return
}

// listLogAnalyticsEntityTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsEntityTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntityTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsEntityTypesResponse
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

// ListLogAnalyticsLogGroups Returns a list of log groups in a compartment. You may limit the number of log groups, provide sorting options, and filter the results by specifying a display name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsLogGroups.go.html to see an example of how to use ListLogAnalyticsLogGroups API.
func (client LogAnalyticsClient) ListLogAnalyticsLogGroups(ctx context.Context, request ListLogAnalyticsLogGroupsRequest) (response ListLogAnalyticsLogGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsLogGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsLogGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsLogGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsLogGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsLogGroupsResponse")
	}
	return
}

// listLogAnalyticsLogGroups implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsLogGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsLogGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsLogGroupsResponse
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

// ListLogAnalyticsObjectCollectionRules Gets list of collection rules.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsObjectCollectionRules.go.html to see an example of how to use ListLogAnalyticsObjectCollectionRules API.
func (client LogAnalyticsClient) ListLogAnalyticsObjectCollectionRules(ctx context.Context, request ListLogAnalyticsObjectCollectionRulesRequest) (response ListLogAnalyticsObjectCollectionRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsObjectCollectionRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsObjectCollectionRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsObjectCollectionRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsObjectCollectionRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsObjectCollectionRulesResponse")
	}
	return
}

// listLogAnalyticsObjectCollectionRules implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsObjectCollectionRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsObjectCollectionRulesResponse
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

// ListLogSets This API returns a list of log sets.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogSets.go.html to see an example of how to use ListLogSets API.
func (client LogAnalyticsClient) ListLogSets(ctx context.Context, request ListLogSetsRequest) (response ListLogSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogSetsResponse")
	}
	return
}

// listLogSets implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage/logSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLogSetsResponse
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

// ListLookups Returns a list of lookups, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as lookup name, description and type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLookups.go.html to see an example of how to use ListLookups API.
func (client LogAnalyticsClient) ListLookups(ctx context.Context, request ListLookupsRequest) (response ListLookupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLookups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLookupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLookupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLookupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLookupsResponse")
	}
	return
}

// listLookups implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLookups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/lookups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLookupsResponse
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

// ListMetaSourceTypes Lists the types of log sources supported.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListMetaSourceTypes.go.html to see an example of how to use ListMetaSourceTypes API.
func (client LogAnalyticsClient) ListMetaSourceTypes(ctx context.Context, request ListMetaSourceTypesRequest) (response ListMetaSourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetaSourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMetaSourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMetaSourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetaSourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetaSourceTypesResponse")
	}
	return
}

// listMetaSourceTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listMetaSourceTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceMetaTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMetaSourceTypesResponse
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

// ListNamespaces Given a tenancy OCID, this API returns the namespace of the tenancy if it is valid and subscribed to the region.  The
// result also indicates if the tenancy is onboarded with Logging Analytics.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListNamespaces.go.html to see an example of how to use ListNamespaces API.
func (client LogAnalyticsClient) ListNamespaces(ctx context.Context, request ListNamespacesRequest) (response ListNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNamespacesResponse")
	}
	return
}

// listNamespaces implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNamespacesResponse
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

// ListParserFunctions Lists the parser functions defined for the specified parser.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParserFunctions.go.html to see an example of how to use ListParserFunctions API.
func (client LogAnalyticsClient) ListParserFunctions(ctx context.Context, request ListParserFunctionsRequest) (response ListParserFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParserFunctions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParserFunctionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParserFunctionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParserFunctionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParserFunctionsResponse")
	}
	return
}

// listParserFunctions implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listParserFunctions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parserFunctions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListParserFunctionsResponse
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

// ListParserMetaPlugins Lists the parser meta plugins available for defining parser functions.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParserMetaPlugins.go.html to see an example of how to use ListParserMetaPlugins API.
func (client LogAnalyticsClient) ListParserMetaPlugins(ctx context.Context, request ListParserMetaPluginsRequest) (response ListParserMetaPluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParserMetaPlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParserMetaPluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParserMetaPluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParserMetaPluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParserMetaPluginsResponse")
	}
	return
}

// listParserMetaPlugins implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listParserMetaPlugins(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parserMetaPlugins", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListParserMetaPluginsResponse
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

// ListParsers Returns a list of parsers, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as parser name, type, display name and description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParsers.go.html to see an example of how to use ListParsers API.
func (client LogAnalyticsClient) ListParsers(ctx context.Context, request ListParsersRequest) (response ListParsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParsersResponse")
	}
	return
}

// listParsers implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listParsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListParsersResponse
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

// ListQueryWorkRequests List active asynchronous queries.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListQueryWorkRequests.go.html to see an example of how to use ListQueryWorkRequests API.
func (client LogAnalyticsClient) ListQueryWorkRequests(ctx context.Context, request ListQueryWorkRequestsRequest) (response ListQueryWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQueryWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQueryWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQueryWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQueryWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQueryWorkRequestsResponse")
	}
	return
}

// listQueryWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listQueryWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/queryWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListQueryWorkRequestsResponse
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

// ListRecalledData This API returns the list of recalled data of a tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListRecalledData.go.html to see an example of how to use ListRecalledData API.
func (client LogAnalyticsClient) ListRecalledData(ctx context.Context, request ListRecalledDataRequest) (response ListRecalledDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecalledData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecalledDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecalledDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecalledDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecalledDataResponse")
	}
	return
}

// listRecalledData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listRecalledData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage/recalledData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecalledDataResponse
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

// ListResourceCategories Returns a list of resources and their category assignments.
// You may limit the number of results, provide sorting order, and filter by information such as resource type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListResourceCategories.go.html to see an example of how to use ListResourceCategories API.
func (client LogAnalyticsClient) ListResourceCategories(ctx context.Context, request ListResourceCategoriesRequest) (response ListResourceCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceCategoriesResponse")
	}
	return
}

// listResourceCategories implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listResourceCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/categories/resourceCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceCategoriesResponse
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

// ListScheduledTasks Lists scheduled tasks.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListScheduledTasks.go.html to see an example of how to use ListScheduledTasks API.
func (client LogAnalyticsClient) ListScheduledTasks(ctx context.Context, request ListScheduledTasksRequest) (response ListScheduledTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledTasksResponse")
	}
	return
}

// listScheduledTasks implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listScheduledTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/scheduledTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledTasksResponse
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

// ListSourceAssociations Returns the list of entity associations in the input compartment for the specified source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceAssociations.go.html to see an example of how to use ListSourceAssociations API.
func (client LogAnalyticsClient) ListSourceAssociations(ctx context.Context, request ListSourceAssociationsRequest) (response ListSourceAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceAssociationsResponse")
	}
	return
}

// listSourceAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourceAssociationsResponse
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

// ListSourceEventTypes Lists the event types mapped to the source with the specified name. The event type string could be the fully qualified name or a prefix that matches the event type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceEventTypes.go.html to see an example of how to use ListSourceEventTypes API.
func (client LogAnalyticsClient) ListSourceEventTypes(ctx context.Context, request ListSourceEventTypesRequest) (response ListSourceEventTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceEventTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceEventTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceEventTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceEventTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceEventTypesResponse")
	}
	return
}

// listSourceEventTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceEventTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/eventTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourceEventTypesResponse
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

// ListSourceExtendedFieldDefinitions Lists the extended field definitions for the source with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceExtendedFieldDefinitions.go.html to see an example of how to use ListSourceExtendedFieldDefinitions API.
func (client LogAnalyticsClient) ListSourceExtendedFieldDefinitions(ctx context.Context, request ListSourceExtendedFieldDefinitionsRequest) (response ListSourceExtendedFieldDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceExtendedFieldDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceExtendedFieldDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceExtendedFieldDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceExtendedFieldDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceExtendedFieldDefinitionsResponse")
	}
	return
}

// listSourceExtendedFieldDefinitions implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceExtendedFieldDefinitions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/extendedFieldDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourceExtendedFieldDefinitionsResponse
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

// ListSourceLabelOperators Lists the supported conditional operators that could be used for matching log field values to generate a label. You may use patterns to specify a condition. If a log entry matches that condition, it is tagged with the corresponding label.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceLabelOperators.go.html to see an example of how to use ListSourceLabelOperators API.
func (client LogAnalyticsClient) ListSourceLabelOperators(ctx context.Context, request ListSourceLabelOperatorsRequest) (response ListSourceLabelOperatorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceLabelOperators, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceLabelOperatorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceLabelOperatorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceLabelOperatorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceLabelOperatorsResponse")
	}
	return
}

// listSourceLabelOperators implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceLabelOperators(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceLabelOperators", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourceLabelOperatorsResponse
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

// ListSourceMetaFunctions Lists the functions that could be used to enrich log entries based on meaningful information extracted from the log fields.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceMetaFunctions.go.html to see an example of how to use ListSourceMetaFunctions API.
func (client LogAnalyticsClient) ListSourceMetaFunctions(ctx context.Context, request ListSourceMetaFunctionsRequest) (response ListSourceMetaFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceMetaFunctions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceMetaFunctionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceMetaFunctionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceMetaFunctionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceMetaFunctionsResponse")
	}
	return
}

// listSourceMetaFunctions implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceMetaFunctions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceMetaFunctions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourceMetaFunctionsResponse
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

// ListSourcePatterns Lists the source patterns for the source with the specified name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourcePatterns.go.html to see an example of how to use ListSourcePatterns API.
func (client LogAnalyticsClient) ListSourcePatterns(ctx context.Context, request ListSourcePatternsRequest) (response ListSourcePatternsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourcePatterns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourcePatternsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourcePatternsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourcePatternsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourcePatternsResponse")
	}
	return
}

// listSourcePatterns implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourcePatterns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/patterns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourcePatternsResponse
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

// ListSources Returns a list of sources, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as display name, description and entity type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSources.go.html to see an example of how to use ListSources API.
func (client LogAnalyticsClient) ListSources(ctx context.Context, request ListSourcesRequest) (response ListSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourcesResponse")
	}
	return
}

// listSources implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSourcesResponse
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

// ListStorageWorkRequestErrors This API returns the list of work request errors if any.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListStorageWorkRequestErrors.go.html to see an example of how to use ListStorageWorkRequestErrors API.
func (client LogAnalyticsClient) ListStorageWorkRequestErrors(ctx context.Context, request ListStorageWorkRequestErrorsRequest) (response ListStorageWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStorageWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStorageWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStorageWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStorageWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStorageWorkRequestErrorsResponse")
	}
	return
}

// listStorageWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listStorageWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storageWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStorageWorkRequestErrorsResponse
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

// ListStorageWorkRequests This API lists storage work requests.  Use query parameters to narrow down or sort the result list.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListStorageWorkRequests.go.html to see an example of how to use ListStorageWorkRequests API.
func (client LogAnalyticsClient) ListStorageWorkRequests(ctx context.Context, request ListStorageWorkRequestsRequest) (response ListStorageWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStorageWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStorageWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStorageWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStorageWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStorageWorkRequestsResponse")
	}
	return
}

// listStorageWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listStorageWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storageWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStorageWorkRequestsResponse
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

// ListSupportedCharEncodings Gets list of character encodings which are supported by on-demand upload.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSupportedCharEncodings.go.html to see an example of how to use ListSupportedCharEncodings API.
func (client LogAnalyticsClient) ListSupportedCharEncodings(ctx context.Context, request ListSupportedCharEncodingsRequest) (response ListSupportedCharEncodingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedCharEncodings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedCharEncodingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedCharEncodingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedCharEncodingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedCharEncodingsResponse")
	}
	return
}

// listSupportedCharEncodings implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSupportedCharEncodings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/supportedCharEncodings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSupportedCharEncodingsResponse
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

// ListSupportedTimezones Gets list of timezones which are supported by on-demand upload.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSupportedTimezones.go.html to see an example of how to use ListSupportedTimezones API.
func (client LogAnalyticsClient) ListSupportedTimezones(ctx context.Context, request ListSupportedTimezonesRequest) (response ListSupportedTimezonesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedTimezones, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedTimezonesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedTimezonesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedTimezonesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedTimezonesResponse")
	}
	return
}

// listSupportedTimezones implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSupportedTimezones(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/supportedTimezones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSupportedTimezonesResponse
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

// ListUploadFiles Gets list of files in an upload along with its processing state.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListUploadFiles.go.html to see an example of how to use ListUploadFiles API.
func (client LogAnalyticsClient) ListUploadFiles(ctx context.Context, request ListUploadFilesRequest) (response ListUploadFilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUploadFiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUploadFilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUploadFilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUploadFilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUploadFilesResponse")
	}
	return
}

// listUploadFiles implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listUploadFiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads/{uploadReference}/files", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUploadFilesResponse
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

// ListUploadWarnings Gets list of warnings in an upload caused by incorrect configuration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListUploadWarnings.go.html to see an example of how to use ListUploadWarnings API.
func (client LogAnalyticsClient) ListUploadWarnings(ctx context.Context, request ListUploadWarningsRequest) (response ListUploadWarningsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUploadWarnings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUploadWarningsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUploadWarningsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUploadWarningsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUploadWarningsResponse")
	}
	return
}

// listUploadWarnings implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listUploadWarnings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads/{uploadReference}/warnings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUploadWarningsResponse
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

// ListUploads Gets a list of all On-demand uploads.
// To use this and other API operations, you must be authorized in an IAM policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListUploads.go.html to see an example of how to use ListUploads API.
func (client LogAnalyticsClient) ListUploads(ctx context.Context, request ListUploadsRequest) (response ListUploadsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUploads, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUploadsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUploadsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUploadsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUploadsResponse")
	}
	return
}

// listUploads implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listUploads(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUploadsResponse
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

// ListWarnings Returns a list of collection warnings, containing detailed information about them. You may limit the number of results, provide sorting order, and filter by information such as start time, end time, warning type, warning state, source name, source pattern and entity name.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListWarnings.go.html to see an example of how to use ListWarnings API.
func (client LogAnalyticsClient) ListWarnings(ctx context.Context, request ListWarningsRequest) (response ListWarningsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWarnings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWarningsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWarningsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWarningsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWarningsResponse")
	}
	return
}

// listWarnings implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listWarnings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/warnings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWarningsResponse
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client LogAnalyticsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LogAnalyticsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
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

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client LogAnalyticsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LogAnalyticsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
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
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client LogAnalyticsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LogAnalyticsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests", binaryReqBody, extraHeaders)
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

// OffboardNamespace Off-boards a tenant from Logging Analytics
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/OffboardNamespace.go.html to see an example of how to use OffboardNamespace API.
func (client LogAnalyticsClient) OffboardNamespace(ctx context.Context, request OffboardNamespaceRequest) (response OffboardNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.offboardNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = OffboardNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = OffboardNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(OffboardNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into OffboardNamespaceResponse")
	}
	return
}

// offboardNamespace implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) offboardNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/offboard", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response OffboardNamespaceResponse
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

// OnboardNamespace On-boards a tenant to Logging Analytics.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/OnboardNamespace.go.html to see an example of how to use OnboardNamespace API.
func (client LogAnalyticsClient) OnboardNamespace(ctx context.Context, request OnboardNamespaceRequest) (response OnboardNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.onboardNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = OnboardNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = OnboardNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(OnboardNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into OnboardNamespaceResponse")
	}
	return
}

// onboardNamespace implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) onboardNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/onboard", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response OnboardNamespaceResponse
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

// ParseQuery Describe query
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ParseQuery.go.html to see an example of how to use ParseQuery API.
func (client LogAnalyticsClient) ParseQuery(ctx context.Context, request ParseQueryRequest) (response ParseQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.parseQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ParseQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ParseQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ParseQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ParseQueryResponse")
	}
	return
}

// parseQuery implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) parseQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/parse", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ParseQueryResponse
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

// PauseScheduledTask Pause the scheduled task specified by {scheduledTaskId}.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/PauseScheduledTask.go.html to see an example of how to use PauseScheduledTask API.
func (client LogAnalyticsClient) PauseScheduledTask(ctx context.Context, request PauseScheduledTaskRequest) (response PauseScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.pauseScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PauseScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PauseScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PauseScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PauseScheduledTaskResponse")
	}
	return
}

// pauseScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) pauseScheduledTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/pause", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PauseScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &scheduledtask{})
	return response, err
}

// PurgeStorageData This API submits a work request to purge data. Only data from log groups that the user has permission to delete
// will be purged.  To purge all data, the user must have permission to all log groups.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/PurgeStorageData.go.html to see an example of how to use PurgeStorageData API.
func (client LogAnalyticsClient) PurgeStorageData(ctx context.Context, request PurgeStorageDataRequest) (response PurgeStorageDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.purgeStorageData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PurgeStorageDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PurgeStorageDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PurgeStorageDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PurgeStorageDataResponse")
	}
	return
}

// purgeStorageData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) purgeStorageData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/purgeData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PurgeStorageDataResponse
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

// PutQueryWorkRequestBackground Put the work request specified by {workRequestId} into the background. Backgrounded queries will preserve query results on query completion for up to 7 days for recall. After 7 days the results and query expire.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/PutQueryWorkRequestBackground.go.html to see an example of how to use PutQueryWorkRequestBackground API.
func (client LogAnalyticsClient) PutQueryWorkRequestBackground(ctx context.Context, request PutQueryWorkRequestBackgroundRequest) (response PutQueryWorkRequestBackgroundResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putQueryWorkRequestBackground, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutQueryWorkRequestBackgroundResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutQueryWorkRequestBackgroundResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutQueryWorkRequestBackgroundResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutQueryWorkRequestBackgroundResponse")
	}
	return
}

// putQueryWorkRequestBackground implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) putQueryWorkRequestBackground(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/queryWorkRequests/{workRequestId}/actions/background", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutQueryWorkRequestBackgroundResponse
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

// Query Performs a log analytics search, if shouldRunAsync is false returns the query results once they become available subject to 60 second timeout. If a query is subject to exceed that time then it should be run asynchronously. Asynchronous query submissions return the queryWorkRequestId to use for execution tracking, query submission lifecycle actions and to poll for query results.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/Query.go.html to see an example of how to use Query API.
func (client LogAnalyticsClient) Query(ctx context.Context, request QueryRequest) (response QueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.query, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryResponse")
	}
	return
}

// query implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) query(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/query", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
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

// RecallArchivedData This API submits a work request to recall archived data based on time interval and data type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/RecallArchivedData.go.html to see an example of how to use RecallArchivedData API.
func (client LogAnalyticsClient) RecallArchivedData(ctx context.Context, request RecallArchivedDataRequest) (response RecallArchivedDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.recallArchivedData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecallArchivedDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecallArchivedDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecallArchivedDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecallArchivedDataResponse")
	}
	return
}

// recallArchivedData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) recallArchivedData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/recallArchivedData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RecallArchivedDataResponse
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

// RegisterLookup Creates a lookup with the specified name, type and description. The csv file containing the lookup content is passed in as binary data in the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/RegisterLookup.go.html to see an example of how to use RegisterLookup API.
func (client LogAnalyticsClient) RegisterLookup(ctx context.Context, request RegisterLookupRequest) (response RegisterLookupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.registerLookup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RegisterLookupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RegisterLookupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RegisterLookupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RegisterLookupResponse")
	}
	return
}

// registerLookup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) registerLookup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/lookups/actions/register", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RegisterLookupResponse
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

// ReleaseRecalledData This API submits a work request to release recalled data based on time interval and data type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ReleaseRecalledData.go.html to see an example of how to use ReleaseRecalledData API.
func (client LogAnalyticsClient) ReleaseRecalledData(ctx context.Context, request ReleaseRecalledDataRequest) (response ReleaseRecalledDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.releaseRecalledData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReleaseRecalledDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReleaseRecalledDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReleaseRecalledDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReleaseRecalledDataResponse")
	}
	return
}

// releaseRecalledData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) releaseRecalledData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/releaseRecalledData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReleaseRecalledDataResponse
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

// RemoveEntityAssociations Delete association between input source log analytics entity and destination entities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/RemoveEntityAssociations.go.html to see an example of how to use RemoveEntityAssociations API.
func (client LogAnalyticsClient) RemoveEntityAssociations(ctx context.Context, request RemoveEntityAssociationsRequest) (response RemoveEntityAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeEntityAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveEntityAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveEntityAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveEntityAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveEntityAssociationsResponse")
	}
	return
}

// removeEntityAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) removeEntityAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/actions/removeEntityAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveEntityAssociationsResponse
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

// RemovePreferences Removes the tenant preferences. Currently, only "DEFAULT_HOMEPAGE" is supported.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/RemovePreferences.go.html to see an example of how to use RemovePreferences API.
func (client LogAnalyticsClient) RemovePreferences(ctx context.Context, request RemovePreferencesRequest) (response RemovePreferencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removePreferences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePreferencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePreferencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePreferencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePreferencesResponse")
	}
	return
}

// removePreferences implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) removePreferences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/preferences/actions/removePreferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePreferencesResponse
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

// RemoveResourceCategories Removes the category assignments of DASHBOARD and SAVEDSEARCH resources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/RemoveResourceCategories.go.html to see an example of how to use RemoveResourceCategories API.
func (client LogAnalyticsClient) RemoveResourceCategories(ctx context.Context, request RemoveResourceCategoriesRequest) (response RemoveResourceCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeResourceCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveResourceCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveResourceCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveResourceCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveResourceCategoriesResponse")
	}
	return
}

// removeResourceCategories implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) removeResourceCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/categories/actions/removeResourceCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveResourceCategoriesResponse
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

// RemoveSourceEventTypes Remove one or more event types from a source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/RemoveSourceEventTypes.go.html to see an example of how to use RemoveSourceEventTypes API.
func (client LogAnalyticsClient) RemoveSourceEventTypes(ctx context.Context, request RemoveSourceEventTypesRequest) (response RemoveSourceEventTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeSourceEventTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveSourceEventTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveSourceEventTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveSourceEventTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveSourceEventTypesResponse")
	}
	return
}

// removeSourceEventTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) removeSourceEventTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/{sourceName}/actions/removeEventTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveSourceEventTypesResponse
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

// ResumeScheduledTask Resume the scheduled task specified by {scheduledTaskId}.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ResumeScheduledTask.go.html to see an example of how to use ResumeScheduledTask API.
func (client LogAnalyticsClient) ResumeScheduledTask(ctx context.Context, request ResumeScheduledTaskRequest) (response ResumeScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resumeScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeScheduledTaskResponse")
	}
	return
}

// resumeScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) resumeScheduledTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &scheduledtask{})
	return response, err
}

// Run Execute the saved search acceleration task in the foreground.
// The ScheduledTask taskType must be ACCELERATION.
// Optionally specify time range (timeStart and timeEnd). The default is all time.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/Run.go.html to see an example of how to use Run API.
func (client LogAnalyticsClient) Run(ctx context.Context, request RunRequest) (response RunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.run, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunResponse")
	}
	return
}

// run implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) run(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/run", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunResponse
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

// SetUnprocessedDataBucket This API configures a bucket to store unprocessed payloads.
// While processing there could be reasons a payload cannot be processed (mismatched structure, corrupted archive format, etc),
// if configured the payload would be uploaded to the bucket for verification.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/SetUnprocessedDataBucket.go.html to see an example of how to use SetUnprocessedDataBucket API.
func (client LogAnalyticsClient) SetUnprocessedDataBucket(ctx context.Context, request SetUnprocessedDataBucketRequest) (response SetUnprocessedDataBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.setUnprocessedDataBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetUnprocessedDataBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetUnprocessedDataBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetUnprocessedDataBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetUnprocessedDataBucketResponse")
	}
	return
}

// setUnprocessedDataBucket implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) setUnprocessedDataBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/setUnprocessedDataBucket", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetUnprocessedDataBucketResponse
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

// Suggest Returns a context specific list of either commands, fields, or values to append to the end of the specified query string if applicable.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/Suggest.go.html to see an example of how to use Suggest API.
func (client LogAnalyticsClient) Suggest(ctx context.Context, request SuggestRequest) (response SuggestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.suggest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuggestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuggestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuggestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuggestResponse")
	}
	return
}

// suggest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) suggest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/suggest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SuggestResponse
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

// SuppressWarning Supresses a list of warnings. Any unsuppressed warnings in the input list would be suppressed. Warnings in the input list which are already suppressed will not be modified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/SuppressWarning.go.html to see an example of how to use SuppressWarning API.
func (client LogAnalyticsClient) SuppressWarning(ctx context.Context, request SuppressWarningRequest) (response SuppressWarningResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.suppressWarning, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuppressWarningResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuppressWarningResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuppressWarningResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuppressWarningResponse")
	}
	return
}

// suppressWarning implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) suppressWarning(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/warnings/actions/suppress", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SuppressWarningResponse
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

// TestParser Tests the parser definition against the specified example content to ensure fields are successfully extracted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/TestParser.go.html to see an example of how to use TestParser API.
func (client LogAnalyticsClient) TestParser(ctx context.Context, request TestParserRequest) (response TestParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.testParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestParserResponse")
	}
	return
}

// testParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) testParser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestParserResponse
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

// UnsuppressWarning Unsupresses a list of warnings. Any suppressed warnings in the input list would be unsuppressed. Warnings in the input list which are already unsuppressed will not be modified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UnsuppressWarning.go.html to see an example of how to use UnsuppressWarning API.
func (client LogAnalyticsClient) UnsuppressWarning(ctx context.Context, request UnsuppressWarningRequest) (response UnsuppressWarningResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.unsuppressWarning, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnsuppressWarningResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnsuppressWarningResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnsuppressWarningResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnsuppressWarningResponse")
	}
	return
}

// unsuppressWarning implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) unsuppressWarning(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/warnings/actions/unsuppress", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnsuppressWarningResponse
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

// UpdateLogAnalyticsEmBridge Update log analytics enterprise manager bridge with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsEmBridge.go.html to see an example of how to use UpdateLogAnalyticsEmBridge API.
func (client LogAnalyticsClient) UpdateLogAnalyticsEmBridge(ctx context.Context, request UpdateLogAnalyticsEmBridgeRequest) (response UpdateLogAnalyticsEmBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsEmBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsEmBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsEmBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsEmBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsEmBridgeResponse")
	}
	return
}

// updateLogAnalyticsEmBridge implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsEmBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsEmBridges/{logAnalyticsEmBridgeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsEmBridgeResponse
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

// UpdateLogAnalyticsEntity Update the log analytics entity with the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsEntity.go.html to see an example of how to use UpdateLogAnalyticsEntity API.
func (client LogAnalyticsClient) UpdateLogAnalyticsEntity(ctx context.Context, request UpdateLogAnalyticsEntityRequest) (response UpdateLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsEntityResponse")
	}
	return
}

// updateLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsEntityResponse
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

// UpdateLogAnalyticsEntityType Update custom log analytics entity type. Out of box entity types cannot be udpated.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsEntityType.go.html to see an example of how to use UpdateLogAnalyticsEntityType API.
func (client LogAnalyticsClient) UpdateLogAnalyticsEntityType(ctx context.Context, request UpdateLogAnalyticsEntityTypeRequest) (response UpdateLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsEntityTypeResponse")
	}
	return
}

// updateLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsEntityTypes/{entityTypeName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsEntityTypeResponse
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

// UpdateLogAnalyticsLogGroup Updates the specified log group's display name, description, defined tags, and free-form tags.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsLogGroup.go.html to see an example of how to use UpdateLogAnalyticsLogGroup API.
func (client LogAnalyticsClient) UpdateLogAnalyticsLogGroup(ctx context.Context, request UpdateLogAnalyticsLogGroupRequest) (response UpdateLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsLogGroupResponse")
	}
	return
}

// updateLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsLogGroupResponse
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

// UpdateLogAnalyticsObjectCollectionRule Updates configuration of the object collection rule for the given id.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsObjectCollectionRule.go.html to see an example of how to use UpdateLogAnalyticsObjectCollectionRule API.
func (client LogAnalyticsClient) UpdateLogAnalyticsObjectCollectionRule(ctx context.Context, request UpdateLogAnalyticsObjectCollectionRuleRequest) (response UpdateLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// updateLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsObjectCollectionRuleResponse
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

// UpdateLookup Updates the metadata of the specified lookup, such as the lookup description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLookup.go.html to see an example of how to use UpdateLookup API.
func (client LogAnalyticsClient) UpdateLookup(ctx context.Context, request UpdateLookupRequest) (response UpdateLookupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateLookup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLookupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLookupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLookupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLookupResponse")
	}
	return
}

// updateLookup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLookup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/lookups/{lookupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLookupResponse
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

// UpdateLookupData Updates the lookup content. The csv file containing the content to be updated is passed in as binary data in the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLookupData.go.html to see an example of how to use UpdateLookupData API.
func (client LogAnalyticsClient) UpdateLookupData(ctx context.Context, request UpdateLookupDataRequest) (response UpdateLookupDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateLookupData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLookupDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLookupDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLookupDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLookupDataResponse")
	}
	return
}

// updateLookupData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLookupData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/lookups/{lookupName}/actions/updateData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLookupDataResponse
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

// UpdatePreferences Updates the tenant preferences. Currently, only "DEFAULT_HOMEPAGE" is supported.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdatePreferences.go.html to see an example of how to use UpdatePreferences API.
func (client LogAnalyticsClient) UpdatePreferences(ctx context.Context, request UpdatePreferencesRequest) (response UpdatePreferencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updatePreferences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePreferencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePreferencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePreferencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePreferencesResponse")
	}
	return
}

// updatePreferences implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updatePreferences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/preferences/actions/updatePreferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePreferencesResponse
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

// UpdateResourceCategories Updates the category assignments of DASHBOARD and SAVEDSEARCH resources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateResourceCategories.go.html to see an example of how to use UpdateResourceCategories API.
func (client LogAnalyticsClient) UpdateResourceCategories(ctx context.Context, request UpdateResourceCategoriesRequest) (response UpdateResourceCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResourceCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResourceCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResourceCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResourceCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResourceCategoriesResponse")
	}
	return
}

// updateResourceCategories implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateResourceCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/categories/actions/updateResourceCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResourceCategoriesResponse
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

// UpdateScheduledTask Update the scheduled task. Schedules may be updated only for taskType SAVED_SEARCH and PURGE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateScheduledTask.go.html to see an example of how to use UpdateScheduledTask API.
func (client LogAnalyticsClient) UpdateScheduledTask(ctx context.Context, request UpdateScheduledTaskRequest) (response UpdateScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduledTaskResponse")
	}
	return
}

// updateScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateScheduledTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &scheduledtask{})
	return response, err
}

// UpdateStorage This API updates the archiving configuration
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateStorage.go.html to see an example of how to use UpdateStorage API.
func (client LogAnalyticsClient) UpdateStorage(ctx context.Context, request UpdateStorageRequest) (response UpdateStorageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStorageResponse")
	}
	return
}

// updateStorage implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateStorage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/storage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStorageResponse
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

// UploadLogEventsFile Accepts log events for processing by Logging Analytics.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UploadLogEventsFile.go.html to see an example of how to use UploadLogEventsFile API.
func (client LogAnalyticsClient) UploadLogEventsFile(ctx context.Context, request UploadLogEventsFileRequest) (response UploadLogEventsFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.uploadLogEventsFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadLogEventsFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadLogEventsFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadLogEventsFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadLogEventsFileResponse")
	}
	return
}

// uploadLogEventsFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) uploadLogEventsFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/uploadLogEventsFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadLogEventsFileResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadLogFile Accepts log data for processing by Logging Analytics.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UploadLogFile.go.html to see an example of how to use UploadLogFile API.
func (client LogAnalyticsClient) UploadLogFile(ctx context.Context, request UploadLogFileRequest) (response UploadLogFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.uploadLogFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadLogFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadLogFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadLogFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadLogFileResponse")
	}
	return
}

// uploadLogFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) uploadLogFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/uploadLogFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadLogFileResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpsertAssociations Creates or updates associations between sources and entities. All entities should belong to the specified input compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpsertAssociations.go.html to see an example of how to use UpsertAssociations API.
func (client LogAnalyticsClient) UpsertAssociations(ctx context.Context, request UpsertAssociationsRequest) (response UpsertAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertAssociationsResponse")
	}
	return
}

// upsertAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/associations/actions/upsert", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpsertAssociationsResponse
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

// UpsertField Creates or updates a field that could be used in parser expressions to extract and assign value. To create a field, specify its display name. A name would be generated for the field. For subsequent calls to update the field, include the name attribute.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpsertField.go.html to see an example of how to use UpsertField API.
func (client LogAnalyticsClient) UpsertField(ctx context.Context, request UpsertFieldRequest) (response UpsertFieldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertField, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertFieldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertFieldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertFieldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertFieldResponse")
	}
	return
}

// upsertField implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertField(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/fields/actions/upsert", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpsertFieldResponse
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

// UpsertLabel Creates or updates a label that could be used to tag a log entry. You may optionally designate the label as a problem, and assign it a priority. You may also provide its related terms (aliases). To create a label, specify its display name. A name would be generated for the label. For subsequent calls to update the label, include the name attribute.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpsertLabel.go.html to see an example of how to use UpsertLabel API.
func (client LogAnalyticsClient) UpsertLabel(ctx context.Context, request UpsertLabelRequest) (response UpsertLabelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertLabel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertLabelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertLabelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertLabelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertLabelResponse")
	}
	return
}

// upsertLabel implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertLabel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/labels/actions/upsert", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpsertLabelResponse
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

// UpsertParser Creates or updates a parser, which defines how fields are extracted from a log entry.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpsertParser.go.html to see an example of how to use UpsertParser API.
func (client LogAnalyticsClient) UpsertParser(ctx context.Context, request UpsertParserRequest) (response UpsertParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertParserResponse")
	}
	return
}

// upsertParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertParser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/upsert", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpsertParserResponse
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

// UpsertSource Creates or updates a log source. You may also specify parsers, labels, extended fields etc., for the source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpsertSource.go.html to see an example of how to use UpsertSource API.
func (client LogAnalyticsClient) UpsertSource(ctx context.Context, request UpsertSourceRequest) (response UpsertSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertSourceResponse")
	}
	return
}

// upsertSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/actions/upsert", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpsertSourceResponse
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

// ValidateAssociationParameters Checks if the passed in entities could be associated with the specified sources. The validation is performed to ensure that the entities have the relevant property values that are used in the corresponding source patterns.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateAssociationParameters.go.html to see an example of how to use ValidateAssociationParameters API.
func (client LogAnalyticsClient) ValidateAssociationParameters(ctx context.Context, request ValidateAssociationParametersRequest) (response ValidateAssociationParametersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateAssociationParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateAssociationParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateAssociationParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateAssociationParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateAssociationParametersResponse")
	}
	return
}

// validateAssociationParameters implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateAssociationParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/associations/actions/validateParameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateAssociationParametersResponse
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

// ValidateFile Validates a log file to check whether it is eligible to be uploaded or not.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateFile.go.html to see an example of how to use ValidateFile API.
func (client LogAnalyticsClient) ValidateFile(ctx context.Context, request ValidateFileRequest) (response ValidateFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateFileResponse")
	}
	return
}

// validateFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/uploads/actions/validateFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateFileResponse
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

// ValidateSource Checks if the specified input is a valid log source definition.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateSource.go.html to see an example of how to use ValidateSource API.
func (client LogAnalyticsClient) ValidateSource(ctx context.Context, request ValidateSourceRequest) (response ValidateSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSourceResponse")
	}
	return
}

// validateSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/actions/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateSourceResponse
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

// ValidateSourceExtendedFieldDetails Checks if the specified input contains valid extended field definitions against the provided example content.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateSourceExtendedFieldDetails.go.html to see an example of how to use ValidateSourceExtendedFieldDetails API.
func (client LogAnalyticsClient) ValidateSourceExtendedFieldDetails(ctx context.Context, request ValidateSourceExtendedFieldDetailsRequest) (response ValidateSourceExtendedFieldDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateSourceExtendedFieldDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSourceExtendedFieldDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSourceExtendedFieldDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSourceExtendedFieldDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSourceExtendedFieldDetailsResponse")
	}
	return
}

// validateSourceExtendedFieldDetails implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateSourceExtendedFieldDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/actions/validateExtendedFields", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateSourceExtendedFieldDetailsResponse
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

// ValidateSourceMapping Validates the source mapping for a given file and provides match status and the parsed representation of log data.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateSourceMapping.go.html to see an example of how to use ValidateSourceMapping API.
func (client LogAnalyticsClient) ValidateSourceMapping(ctx context.Context, request ValidateSourceMappingRequest) (response ValidateSourceMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateSourceMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSourceMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSourceMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSourceMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSourceMappingResponse")
	}
	return
}

// validateSourceMapping implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateSourceMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/uploads/actions/validateSourceMapping", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateSourceMappingResponse
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

// Verify Verify the accelerated saved search task specified by {scheduledTaskId}.
// For internal use only.
// Optionally specify whether to return accelerated search results; the default is false.
// The ScheduledTask taskType must be ACCELERATION.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/Verify.go.html to see an example of how to use Verify API.
func (client LogAnalyticsClient) Verify(ctx context.Context, request VerifyRequest) (response VerifyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.verify, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = VerifyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = VerifyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(VerifyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into VerifyResponse")
	}
	return
}

// verify implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) verify(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/verify", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response VerifyResponse
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
