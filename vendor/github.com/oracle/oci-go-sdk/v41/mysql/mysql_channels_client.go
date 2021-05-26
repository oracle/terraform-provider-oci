// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v41/common"
	"github.com/oracle/oci-go-sdk/v41/common/auth"
	"net/http"
)

//ChannelsClient a client for Channels
type ChannelsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewChannelsClientWithConfigurationProvider Creates a new default Channels client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewChannelsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ChannelsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newChannelsClientFromBaseClient(baseClient, provider)
}

// NewChannelsClientWithOboToken Creates a new default Channels client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewChannelsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ChannelsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newChannelsClientFromBaseClient(baseClient, configProvider)
}

func newChannelsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ChannelsClient, err error) {
	client = ChannelsClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ChannelsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ChannelsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ChannelsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateChannel Creates a Channel to establish replication from a source to a target.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/CreateChannel.go.html to see an example of how to use CreateChannel API.
func (client ChannelsClient) CreateChannel(ctx context.Context, request CreateChannelRequest) (response CreateChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateChannelResponse")
	}
	return
}

// createChannel implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) createChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/channels", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateChannelResponse
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

// DeleteChannel Deletes the specified Channel.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/DeleteChannel.go.html to see an example of how to use DeleteChannel API.
func (client ChannelsClient) DeleteChannel(ctx context.Context, request DeleteChannelRequest) (response DeleteChannelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteChannelResponse")
	}
	return
}

// deleteChannel implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) deleteChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/channels/{channelId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteChannelResponse
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

// GetChannel Gets the full details of the specified Channel, including the user-specified
// configuration parameters (passwords are omitted), as well as information about
// the state of the Channel, its sources and targets.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetChannel.go.html to see an example of how to use GetChannel API.
func (client ChannelsClient) GetChannel(ctx context.Context, request GetChannelRequest) (response GetChannelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetChannelResponse")
	}
	return
}

// getChannel implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) getChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/channels/{channelId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetChannelResponse
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

// ListChannels Lists all the Channels that match the specified filters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListChannels.go.html to see an example of how to use ListChannels API.
func (client ChannelsClient) ListChannels(ctx context.Context, request ListChannelsRequest) (response ListChannelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listChannels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListChannelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListChannelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListChannelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListChannelsResponse")
	}
	return
}

// listChannels implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) listChannels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/channels", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListChannelsResponse
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

// ResetChannel Resets the specified Channel by purging its cached information, leaving the Channel
// as if it had just been created. This operation is only accepted in Inactive Channels.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ResetChannel.go.html to see an example of how to use ResetChannel API.
func (client ChannelsClient) ResetChannel(ctx context.Context, request ResetChannelRequest) (response ResetChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resetChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResetChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResetChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResetChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResetChannelResponse")
	}
	return
}

// resetChannel implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) resetChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/channels/{channelId}/actions/reset", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ResetChannelResponse
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

// ResumeChannel Resumes an enabled Channel that has become Inactive due to an error. The resume operation
// requires that the error that cause the Channel to become Inactive has already been fixed,
// otherwise the operation may fail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ResumeChannel.go.html to see an example of how to use ResumeChannel API.
func (client ChannelsClient) ResumeChannel(ctx context.Context, request ResumeChannelRequest) (response ResumeChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resumeChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeChannelResponse")
	}
	return
}

// resumeChannel implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) resumeChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/channels/{channelId}/actions/resume", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ResumeChannelResponse
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

// UpdateChannel Updates the properties of the specified Channel.
// If the Channel is Active the Update operation will asynchronously apply the new configuration
// parameters to the Channel and the Channel may become temporarily unavailable. Otherwise, the
// new configuration will be applied the next time the Channel becomes Active.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/UpdateChannel.go.html to see an example of how to use UpdateChannel API.
func (client ChannelsClient) UpdateChannel(ctx context.Context, request UpdateChannelRequest) (response UpdateChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateChannelResponse")
	}
	return
}

// updateChannel implements the OCIOperation interface (enables retrying operations)
func (client ChannelsClient) updateChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/channels/{channelId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateChannelResponse
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
