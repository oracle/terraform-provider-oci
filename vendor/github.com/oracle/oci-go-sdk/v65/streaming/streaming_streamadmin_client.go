// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//StreamAdminClient a client for StreamAdmin
type StreamAdminClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewStreamAdminClientWithConfigurationProvider Creates a new default StreamAdmin client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewStreamAdminClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client StreamAdminClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newStreamAdminClientFromBaseClient(baseClient, provider)
}

// NewStreamAdminClientWithOboToken Creates a new default StreamAdmin client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewStreamAdminClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client StreamAdminClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newStreamAdminClientFromBaseClient(baseClient, configProvider)
}

func newStreamAdminClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client StreamAdminClient, err error) {
	// StreamAdmin service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("StreamAdmin"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = StreamAdminClient{BaseClient: baseClient}
	client.BasePath = "20180418"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *StreamAdminClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("streams", "https://streaming.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *StreamAdminClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *StreamAdminClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeConnectHarnessCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ChangeConnectHarnessCompartment.go.html to see an example of how to use ChangeConnectHarnessCompartment API.
// A default retry strategy applies to this operation ChangeConnectHarnessCompartment()
func (client StreamAdminClient) ChangeConnectHarnessCompartment(ctx context.Context, request ChangeConnectHarnessCompartmentRequest) (response ChangeConnectHarnessCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeConnectHarnessCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeConnectHarnessCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeConnectHarnessCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeConnectHarnessCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeConnectHarnessCompartmentResponse")
	}
	return
}

// changeConnectHarnessCompartment implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) changeConnectHarnessCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connectharnesses/{connectHarnessId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeConnectHarnessCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarness/ChangeConnectHarnessCompartment"
		err = common.PostProcessServiceError(err, "StreamAdmin", "ChangeConnectHarnessCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeStreamCompartment Moves a resource into a different compartment.
// When provided, If-Match is checked against ETag values of the resource.
// The stream will also be moved into the default stream pool in the destination compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ChangeStreamCompartment.go.html to see an example of how to use ChangeStreamCompartment API.
// A default retry strategy applies to this operation ChangeStreamCompartment()
func (client StreamAdminClient) ChangeStreamCompartment(ctx context.Context, request ChangeStreamCompartmentRequest) (response ChangeStreamCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeStreamCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeStreamCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeStreamCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeStreamCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeStreamCompartmentResponse")
	}
	return
}

// changeStreamCompartment implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) changeStreamCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeStreamCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Stream/ChangeStreamCompartment"
		err = common.PostProcessServiceError(err, "StreamAdmin", "ChangeStreamCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeStreamPoolCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ChangeStreamPoolCompartment.go.html to see an example of how to use ChangeStreamPoolCompartment API.
// A default retry strategy applies to this operation ChangeStreamPoolCompartment()
func (client StreamAdminClient) ChangeStreamPoolCompartment(ctx context.Context, request ChangeStreamPoolCompartmentRequest) (response ChangeStreamPoolCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeStreamPoolCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeStreamPoolCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeStreamPoolCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeStreamPoolCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeStreamPoolCompartmentResponse")
	}
	return
}

// changeStreamPoolCompartment implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) changeStreamPoolCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streampools/{streamPoolId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeStreamPoolCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/ChangeStreamPoolCompartment"
		err = common.PostProcessServiceError(err, "StreamAdmin", "ChangeStreamPoolCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConnectHarness Starts the provisioning of a new connect harness.
// To track the progress of the provisioning, you can periodically call ConnectHarness object tells you its current state.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/CreateConnectHarness.go.html to see an example of how to use CreateConnectHarness API.
// A default retry strategy applies to this operation CreateConnectHarness()
func (client StreamAdminClient) CreateConnectHarness(ctx context.Context, request CreateConnectHarnessRequest) (response CreateConnectHarnessResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConnectHarness, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectHarnessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectHarnessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectHarnessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectHarnessResponse")
	}
	return
}

// createConnectHarness implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) createConnectHarness(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connectharnesses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectHarnessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarness/CreateConnectHarness"
		err = common.PostProcessServiceError(err, "StreamAdmin", "CreateConnectHarness", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateStream Starts the provisioning of a new stream.
// The stream will be created in the given compartment id or stream pool id, depending on which parameter is specified.
// Compartment id and stream pool id cannot be specified at the same time.
// To track the progress of the provisioning, you can periodically call GetStream.
// In the response, the `lifecycleState` parameter of the Stream object tells you its current state.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/CreateStream.go.html to see an example of how to use CreateStream API.
// A default retry strategy applies to this operation CreateStream()
func (client StreamAdminClient) CreateStream(ctx context.Context, request CreateStreamRequest) (response CreateStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStreamResponse")
	}
	return
}

// createStream implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) createStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Stream/CreateStream"
		err = common.PostProcessServiceError(err, "StreamAdmin", "CreateStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateStreamPool Starts the provisioning of a new stream pool.
// To track the progress of the provisioning, you can periodically call GetStreamPool.
// In the response, the `lifecycleState` parameter of the object tells you its current state.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/CreateStreamPool.go.html to see an example of how to use CreateStreamPool API.
// A default retry strategy applies to this operation CreateStreamPool()
func (client StreamAdminClient) CreateStreamPool(ctx context.Context, request CreateStreamPoolRequest) (response CreateStreamPoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createStreamPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStreamPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStreamPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStreamPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStreamPoolResponse")
	}
	return
}

// createStreamPool implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) createStreamPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streampools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStreamPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/CreateStreamPool"
		err = common.PostProcessServiceError(err, "StreamAdmin", "CreateStreamPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnectHarness Deletes a connect harness and its content. Connect harness contents are deleted immediately. The service retains records of the connect harness itself for 90 days after deletion.
// The `lifecycleState` parameter of the `ConnectHarness` object changes to `DELETING` and the connect harness becomes inaccessible for read or write operations.
// To verify that a connect harness has been deleted, make a GetConnectHarness request. If the call returns the connect harness's
// lifecycle state as `DELETED`, then the connect harness has been deleted. If the call returns a "404 Not Found" error, that means all records of the
// connect harness have been deleted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/DeleteConnectHarness.go.html to see an example of how to use DeleteConnectHarness API.
// A default retry strategy applies to this operation DeleteConnectHarness()
func (client StreamAdminClient) DeleteConnectHarness(ctx context.Context, request DeleteConnectHarnessRequest) (response DeleteConnectHarnessResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnectHarness, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectHarnessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectHarnessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectHarnessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectHarnessResponse")
	}
	return
}

// deleteConnectHarness implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) deleteConnectHarness(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/connectharnesses/{connectHarnessId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectHarnessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarness/DeleteConnectHarness"
		err = common.PostProcessServiceError(err, "StreamAdmin", "DeleteConnectHarness", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteStream Deletes a stream and its content. Stream contents are deleted immediately. The service retains records of the stream itself for 90 days after deletion.
// The `lifecycleState` parameter of the `Stream` object changes to `DELETING` and the stream becomes inaccessible for read or write operations.
// To verify that a stream has been deleted, make a GetStream request. If the call returns the stream's
// lifecycle state as `DELETED`, then the stream has been deleted. If the call returns a "404 Not Found" error, that means all records of the
// stream have been deleted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/DeleteStream.go.html to see an example of how to use DeleteStream API.
// A default retry strategy applies to this operation DeleteStream()
func (client StreamAdminClient) DeleteStream(ctx context.Context, request DeleteStreamRequest) (response DeleteStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStreamResponse")
	}
	return
}

// deleteStream implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) deleteStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/streams/{streamId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Stream/DeleteStream"
		err = common.PostProcessServiceError(err, "StreamAdmin", "DeleteStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteStreamPool Deletes a stream pool. All containing streams will also be deleted.
// The default stream pool of a compartment cannot be deleted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/DeleteStreamPool.go.html to see an example of how to use DeleteStreamPool API.
// A default retry strategy applies to this operation DeleteStreamPool()
func (client StreamAdminClient) DeleteStreamPool(ctx context.Context, request DeleteStreamPoolRequest) (response DeleteStreamPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStreamPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteStreamPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteStreamPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStreamPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStreamPoolResponse")
	}
	return
}

// deleteStreamPool implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) deleteStreamPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/streampools/{streamPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteStreamPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/DeleteStreamPool"
		err = common.PostProcessServiceError(err, "StreamAdmin", "DeleteStreamPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConnectHarness Gets detailed information about a connect harness.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/GetConnectHarness.go.html to see an example of how to use GetConnectHarness API.
// A default retry strategy applies to this operation GetConnectHarness()
func (client StreamAdminClient) GetConnectHarness(ctx context.Context, request GetConnectHarnessRequest) (response GetConnectHarnessResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnectHarness, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectHarnessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectHarnessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectHarnessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectHarnessResponse")
	}
	return
}

// getConnectHarness implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) getConnectHarness(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connectharnesses/{connectHarnessId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectHarnessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarness/GetConnectHarness"
		err = common.PostProcessServiceError(err, "StreamAdmin", "GetConnectHarness", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStream Gets detailed information about a stream, including the number of partitions.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/GetStream.go.html to see an example of how to use GetStream API.
// A default retry strategy applies to this operation GetStream()
func (client StreamAdminClient) GetStream(ctx context.Context, request GetStreamRequest) (response GetStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStreamResponse")
	}
	return
}

// getStream implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) getStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams/{streamId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Stream/GetStream"
		err = common.PostProcessServiceError(err, "StreamAdmin", "GetStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStreamPool Gets detailed information about the stream pool, such as Kafka settings.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/GetStreamPool.go.html to see an example of how to use GetStreamPool API.
// A default retry strategy applies to this operation GetStreamPool()
func (client StreamAdminClient) GetStreamPool(ctx context.Context, request GetStreamPoolRequest) (response GetStreamPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStreamPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStreamPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStreamPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStreamPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStreamPoolResponse")
	}
	return
}

// getStreamPool implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) getStreamPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streampools/{streamPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStreamPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/GetStreamPool"
		err = common.PostProcessServiceError(err, "StreamAdmin", "GetStreamPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnectHarnesses Lists the connectharness.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ListConnectHarnesses.go.html to see an example of how to use ListConnectHarnesses API.
// A default retry strategy applies to this operation ListConnectHarnesses()
func (client StreamAdminClient) ListConnectHarnesses(ctx context.Context, request ListConnectHarnessesRequest) (response ListConnectHarnessesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnectHarnesses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectHarnessesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectHarnessesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectHarnessesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectHarnessesResponse")
	}
	return
}

// listConnectHarnesses implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) listConnectHarnesses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connectharnesses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectHarnessesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarnessSummary/ListConnectHarnesses"
		err = common.PostProcessServiceError(err, "StreamAdmin", "ListConnectHarnesses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStreamPools List the stream pools for a given compartment ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ListStreamPools.go.html to see an example of how to use ListStreamPools API.
// A default retry strategy applies to this operation ListStreamPools()
func (client StreamAdminClient) ListStreamPools(ctx context.Context, request ListStreamPoolsRequest) (response ListStreamPoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStreamPools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStreamPoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStreamPoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStreamPoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStreamPoolsResponse")
	}
	return
}

// listStreamPools implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) listStreamPools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streampools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStreamPoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamPoolSummary/ListStreamPools"
		err = common.PostProcessServiceError(err, "StreamAdmin", "ListStreamPools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStreams Lists the streams in the given compartment id.
// If the compartment id is specified, it will list streams in the compartment, regardless of their stream pool.
// If the stream pool id is specified, the action will be scoped to that stream pool.
// The compartment id and stream pool id cannot be specified at the same time.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ListStreams.go.html to see an example of how to use ListStreams API.
// A default retry strategy applies to this operation ListStreams()
func (client StreamAdminClient) ListStreams(ctx context.Context, request ListStreamsRequest) (response ListStreamsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStreams, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStreamsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStreamsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStreamsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStreamsResponse")
	}
	return
}

// listStreams implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) listStreams(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStreamsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamSummary/ListStreams"
		err = common.PostProcessServiceError(err, "StreamAdmin", "ListStreams", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConnectHarness Updates the tags applied to the connect harness.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/UpdateConnectHarness.go.html to see an example of how to use UpdateConnectHarness API.
// A default retry strategy applies to this operation UpdateConnectHarness()
func (client StreamAdminClient) UpdateConnectHarness(ctx context.Context, request UpdateConnectHarnessRequest) (response UpdateConnectHarnessResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnectHarness, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConnectHarnessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConnectHarnessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectHarnessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectHarnessResponse")
	}
	return
}

// updateConnectHarness implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) updateConnectHarness(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/connectharnesses/{connectHarnessId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConnectHarnessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarness/UpdateConnectHarness"
		err = common.PostProcessServiceError(err, "StreamAdmin", "UpdateConnectHarness", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateStream Updates the stream. Only specified values will be updated.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/UpdateStream.go.html to see an example of how to use UpdateStream API.
// A default retry strategy applies to this operation UpdateStream()
func (client StreamAdminClient) UpdateStream(ctx context.Context, request UpdateStreamRequest) (response UpdateStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStreamResponse")
	}
	return
}

// updateStream implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) updateStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streams/{streamId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/Stream/UpdateStream"
		err = common.PostProcessServiceError(err, "StreamAdmin", "UpdateStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateStreamPool Updates the specified stream pool.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/UpdateStreamPool.go.html to see an example of how to use UpdateStreamPool API.
// A default retry strategy applies to this operation UpdateStreamPool()
func (client StreamAdminClient) UpdateStreamPool(ctx context.Context, request UpdateStreamPoolRequest) (response UpdateStreamPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStreamPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStreamPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStreamPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStreamPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStreamPoolResponse")
	}
	return
}

// updateStreamPool implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) updateStreamPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streampools/{streamPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStreamPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/UpdateStreamPool"
		err = common.PostProcessServiceError(err, "StreamAdmin", "UpdateStreamPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
