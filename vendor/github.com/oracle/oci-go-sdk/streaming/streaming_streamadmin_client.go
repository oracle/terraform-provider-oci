// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
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
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *StreamAdminClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeStreamCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
func (client StreamAdminClient) ChangeStreamCompartment(ctx context.Context, request ChangeStreamCompartmentRequest) (response ChangeStreamCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeStreamCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			response = ChangeStreamCompartmentResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client StreamAdminClient) changeStreamCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeStreamCompartmentResponse
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

// CreateArchiver Starts the provisioning of a new stream archiver.
// To track the progress of the provisioning, you can periodically call GetArchiver.
// In the response, the `lifecycleState` parameter of the Archiver object tells you its current state.
func (client StreamAdminClient) CreateArchiver(ctx context.Context, request CreateArchiverRequest) (response CreateArchiverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createArchiver, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateArchiverResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateArchiverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateArchiverResponse")
	}
	return
}

// createArchiver implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) createArchiver(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/archiver")
	if err != nil {
		return nil, err
	}

	var response CreateArchiverResponse
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

// CreateStream Starts the provisioning of a new stream.
// To track the progress of the provisioning, you can periodically call GetStream.
// In the response, the `lifecycleState` parameter of the Stream object tells you its current state.
func (client StreamAdminClient) CreateStream(ctx context.Context, request CreateStreamRequest) (response CreateStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createStream, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateStreamResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client StreamAdminClient) createStream(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams")
	if err != nil {
		return nil, err
	}

	var response CreateStreamResponse
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

// DeleteStream Deletes a stream and its content. Stream contents are deleted immediately. The service retains records of the stream itself for 90 days after deletion.
// The `lifecycleState` parameter of the `Stream` object changes to `DELETING` and the stream becomes inaccessible for read or write operations.
// To verify that a stream has been deleted, make a GetStream request. If the call returns the stream's
// lifecycle state as `DELETED`, then the stream has been deleted. If the call returns a "404 Not Found" error, that means all records of the
// stream have been deleted.
func (client StreamAdminClient) DeleteStream(ctx context.Context, request DeleteStreamRequest) (response DeleteStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStream, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteStreamResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client StreamAdminClient) deleteStream(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/streams/{streamId}")
	if err != nil {
		return nil, err
	}

	var response DeleteStreamResponse
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

// GetArchiver Returns the current state of the stream archiver.
func (client StreamAdminClient) GetArchiver(ctx context.Context, request GetArchiverRequest) (response GetArchiverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getArchiver, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetArchiverResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetArchiverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetArchiverResponse")
	}
	return
}

// getArchiver implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) getArchiver(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams/{streamId}/archiver")
	if err != nil {
		return nil, err
	}

	var response GetArchiverResponse
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

// GetStream Gets detailed information about a stream, including the number of partitions.
func (client StreamAdminClient) GetStream(ctx context.Context, request GetStreamRequest) (response GetStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStream, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetStreamResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client StreamAdminClient) getStream(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams/{streamId}")
	if err != nil {
		return nil, err
	}

	var response GetStreamResponse
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

// ListStreams Lists the streams.
func (client StreamAdminClient) ListStreams(ctx context.Context, request ListStreamsRequest) (response ListStreamsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStreams, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListStreamsResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client StreamAdminClient) listStreams(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streams")
	if err != nil {
		return nil, err
	}

	var response ListStreamsResponse
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

// StartArchiver Start the archiver for the specified stream.
func (client StreamAdminClient) StartArchiver(ctx context.Context, request StartArchiverRequest) (response StartArchiverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startArchiver, policy)
	if err != nil {
		if ociResponse != nil {
			response = StartArchiverResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartArchiverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartArchiverResponse")
	}
	return
}

// startArchiver implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) startArchiver(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/archiver/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartArchiverResponse
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

// StopArchiver Stop the archiver for the specified stream.
func (client StreamAdminClient) StopArchiver(ctx context.Context, request StopArchiverRequest) (response StopArchiverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopArchiver, policy)
	if err != nil {
		if ociResponse != nil {
			response = StopArchiverResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopArchiverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopArchiverResponse")
	}
	return
}

// stopArchiver implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) stopArchiver(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streams/{streamId}/archiver/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopArchiverResponse
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

// UpdateArchiver Update the stream archiver parameters.
func (client StreamAdminClient) UpdateArchiver(ctx context.Context, request UpdateArchiverRequest) (response UpdateArchiverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateArchiver, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateArchiverResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateArchiverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateArchiverResponse")
	}
	return
}

// updateArchiver implements the OCIOperation interface (enables retrying operations)
func (client StreamAdminClient) updateArchiver(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streams/{streamId}/archiver")
	if err != nil {
		return nil, err
	}

	var response UpdateArchiverResponse
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

// UpdateStream Updates the tags applied to the stream.
func (client StreamAdminClient) UpdateStream(ctx context.Context, request UpdateStreamRequest) (response UpdateStreamResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStream, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateStreamResponse{RawResponse: ociResponse.HTTPResponse()}
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
func (client StreamAdminClient) updateStream(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streams/{streamId}")
	if err != nil {
		return nil, err
	}

	var response UpdateStreamResponse
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
