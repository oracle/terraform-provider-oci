// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//TraceClient a client for Trace
type TraceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTraceClientWithConfigurationProvider Creates a new default Trace client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTraceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TraceClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newTraceClientFromBaseClient(baseClient, provider)
}

// NewTraceClientWithOboToken Creates a new default Trace client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewTraceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TraceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTraceClientFromBaseClient(baseClient, configProvider)
}

func newTraceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TraceClient, err error) {
	// Trace service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Trace"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = TraceClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TraceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmtraces", "https://apm-trace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TraceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TraceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkActivateAttribute Activates a set of attributes for the given APM Domain.  The API is case in-sensitive.  Any duplicates present in the bulk activation
// request are de-duplicated and only unique attributes are activated.  A maximum number of 700 string attributes and 100 numeric attributes
// can be activated in an APM Domain subject to the available string and numeric slots.  Once an attribute has been activated, it may take sometime
// for it to be appear in searches as the span processor might not have picked up the changes or any associated caches might not have refreshed.  The
// bulk activation operation is atomic, and the operation succeeds only if all the attributes in the request have been processed successfully and they
// get a success status back.  If the processing of any attribute results in a processing or validation error, then none of the attributes in the bulk
// request are activated.  Attributes that are activated are un-pinned by default if they are pinned.
func (client TraceClient) BulkActivateAttribute(ctx context.Context, request BulkActivateAttributeRequest) (response BulkActivateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkActivateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkActivateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkActivateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkActivateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkActivateAttributeResponse")
	}
	return
}

// bulkActivateAttribute implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) bulkActivateAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/activateAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkActivateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/BulkActivateAttribute"
		err = common.PostProcessServiceError(err, "Trace", "BulkActivateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkDeActivateAttribute De-activates a set of attributes for the given APM Domain.  The API is case in-sensitive.  Any duplicates present in the bulk de-activation
// request are de-duplicated and only unique attributes are de-activated.  A maximum number of 700 string attributes and 100 numeric attributes
// can be activated in an APM Domain subject to the available string and numeric slots.  Out of box attributes (Trace and Span) cannot be
// de-activated, and will result in a processing error.  Once an attribute has been de-activated, it may take sometime for it to dissappear in
// searches as the span processor might not have picked up the changes or any associated caches might not have refreshed.  The bulk de-activation
// operation is atomic, and the operation succeeds only if all the attributes in the request have been processed successfully and they get a success
// status back.  If the processing of any attribute results in a processing or validation error, then none of the attributes in the bulk request
// are de-activated.
func (client TraceClient) BulkDeActivateAttribute(ctx context.Context, request BulkDeActivateAttributeRequest) (response BulkDeActivateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkDeActivateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkDeActivateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkDeActivateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkDeActivateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkDeActivateAttributeResponse")
	}
	return
}

// bulkDeActivateAttribute implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) bulkDeActivateAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/deActivateAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkDeActivateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/BulkDeActivateAttribute"
		err = common.PostProcessServiceError(err, "Trace", "BulkDeActivateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkPinAttribute Pin a set of attributes in the APM Domain.  Attributes the are marked pinned are not auto-promoted by the span processor.
// Attributes that are de-activated are pinned by default.
func (client TraceClient) BulkPinAttribute(ctx context.Context, request BulkPinAttributeRequest) (response BulkPinAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkPinAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkPinAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkPinAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkPinAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkPinAttributeResponse")
	}
	return
}

// bulkPinAttribute implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) bulkPinAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/pinAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkPinAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/BulkPinAttribute"
		err = common.PostProcessServiceError(err, "Trace", "BulkPinAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUnpinAttribute Un-pin a set of attributes in the APM Domain.
func (client TraceClient) BulkUnpinAttribute(ctx context.Context, request BulkUnpinAttributeRequest) (response BulkUnpinAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkUnpinAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUnpinAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUnpinAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUnpinAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUnpinAttributeResponse")
	}
	return
}

// bulkUnpinAttribute implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) bulkUnpinAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/unPinAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUnpinAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/BulkUnpinAttribute"
		err = common.PostProcessServiceError(err, "Trace", "BulkUnpinAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUpdateAttributeNotes Add or edit notes to a set of attributes in the APM Domain.  Notes can be added to either an active or an inactive attribute.  The
// notes will be preserved even if the attribute changes state (when an active attribute is de-activated or when an inactive attribute
// is activated).
func (client TraceClient) BulkUpdateAttributeNotes(ctx context.Context, request BulkUpdateAttributeNotesRequest) (response BulkUpdateAttributeNotesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkUpdateAttributeNotes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUpdateAttributeNotesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUpdateAttributeNotesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUpdateAttributeNotesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUpdateAttributeNotesResponse")
	}
	return
}

// bulkUpdateAttributeNotes implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) bulkUpdateAttributeNotes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/updateNotes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUpdateAttributeNotesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/BulkUpdateAttributeNotes"
		err = common.PostProcessServiceError(err, "Trace", "BulkUpdateAttributeNotes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAggregatedSnapshot Gets the aggregated snapshot identified by trace ID.
func (client TraceClient) GetAggregatedSnapshot(ctx context.Context, request GetAggregatedSnapshotRequest) (response GetAggregatedSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAggregatedSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAggregatedSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAggregatedSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAggregatedSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAggregatedSnapshotResponse")
	}
	return
}

// getAggregatedSnapshot implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getAggregatedSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/traces/{traceKey}/aggregatedSnapshotData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAggregatedSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/AggregatedSnapshot/GetAggregatedSnapshot"
		err = common.PostProcessServiceError(err, "Trace", "GetAggregatedSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSpan Gets the span details identified by spanId.
func (client TraceClient) GetSpan(ctx context.Context, request GetSpanRequest) (response GetSpanResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSpan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSpanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSpanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSpanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSpanResponse")
	}
	return
}

// getSpan implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getSpan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/spans/{traceKey}/{spanKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSpanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Span/GetSpan"
		err = common.PostProcessServiceError(err, "Trace", "GetSpan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStatusAutoActivate Get auto activation status for a private data key or public data key in the APM Domain.
func (client TraceClient) GetStatusAutoActivate(ctx context.Context, request GetStatusAutoActivateRequest) (response GetStatusAutoActivateResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStatusAutoActivate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStatusAutoActivateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStatusAutoActivateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStatusAutoActivateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStatusAutoActivateResponse")
	}
	return
}

// getStatusAutoActivate implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getStatusAutoActivate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/attributes/autoActivateStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStatusAutoActivateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/GetStatusAutoActivate"
		err = common.PostProcessServiceError(err, "Trace", "GetStatusAutoActivate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTrace Gets the trace details identified by traceId.
func (client TraceClient) GetTrace(ctx context.Context, request GetTraceRequest) (response GetTraceResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTrace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTraceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTraceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTraceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTraceResponse")
	}
	return
}

// getTrace implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getTrace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/traces/{traceKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTraceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/GetTrace"
		err = common.PostProcessServiceError(err, "Trace", "GetTrace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTraceSnapshot Gets the trace snapshots data identified by trace ID.
func (client TraceClient) GetTraceSnapshot(ctx context.Context, request GetTraceSnapshotRequest) (response GetTraceSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTraceSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTraceSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTraceSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTraceSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTraceSnapshotResponse")
	}
	return
}

// getTraceSnapshot implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) getTraceSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/traces/{traceKey}/snapshotData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTraceSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/TraceSnapshot/GetTraceSnapshot"
		err = common.PostProcessServiceError(err, "Trace", "GetTraceSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutToggleAutoActivate Turn on or off auto activate for private data key or public data key traffic a given APM Domain.
func (client TraceClient) PutToggleAutoActivate(ctx context.Context, request PutToggleAutoActivateRequest) (response PutToggleAutoActivateResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putToggleAutoActivate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutToggleAutoActivateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutToggleAutoActivateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutToggleAutoActivateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutToggleAutoActivateResponse")
	}
	return
}

// putToggleAutoActivate implements the OCIOperation interface (enables retrying operations)
func (client TraceClient) putToggleAutoActivate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/autoActivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutToggleAutoActivateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/Trace/PutToggleAutoActivate"
		err = common.PostProcessServiceError(err, "Trace", "PutToggleAutoActivate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
