// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// BlockchainPlatformClient a client for BlockchainPlatform
type BlockchainPlatformClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBlockchainPlatformClientWithConfigurationProvider Creates a new default BlockchainPlatform client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBlockchainPlatformClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BlockchainPlatformClient, err error) {
	if enabled := common.CheckForEnabledServices("blockchain"); !enabled {
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
	return newBlockchainPlatformClientFromBaseClient(baseClient, provider)
}

// NewBlockchainPlatformClientWithOboToken Creates a new default BlockchainPlatform client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewBlockchainPlatformClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BlockchainPlatformClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newBlockchainPlatformClientFromBaseClient(baseClient, configProvider)
}

func newBlockchainPlatformClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BlockchainPlatformClient, err error) {
	// BlockchainPlatform service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("BlockchainPlatform"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = BlockchainPlatformClient{BaseClient: baseClient}
	client.BasePath = "20191010"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BlockchainPlatformClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("blockchain", "https://blockchain.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BlockchainPlatformClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BlockchainPlatformClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBlockchainPlatformCompartment Change Blockchain Platform Compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ChangeBlockchainPlatformCompartment.go.html to see an example of how to use ChangeBlockchainPlatformCompartment API.
func (client BlockchainPlatformClient) ChangeBlockchainPlatformCompartment(ctx context.Context, request ChangeBlockchainPlatformCompartmentRequest) (response ChangeBlockchainPlatformCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeBlockchainPlatformCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBlockchainPlatformCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBlockchainPlatformCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBlockchainPlatformCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBlockchainPlatformCompartmentResponse")
	}
	return
}

// changeBlockchainPlatformCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) changeBlockchainPlatformCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBlockchainPlatformCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/ChangeBlockchainPlatformCompartment"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ChangeBlockchainPlatformCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBlockchainPlatform Creates a new Blockchain Platform.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/CreateBlockchainPlatform.go.html to see an example of how to use CreateBlockchainPlatform API.
func (client BlockchainPlatformClient) CreateBlockchainPlatform(ctx context.Context, request CreateBlockchainPlatformRequest) (response CreateBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBlockchainPlatformResponse")
	}
	return
}

// createBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) createBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/CreateBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "CreateBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOsn Create Blockchain Platform Osn
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/CreateOsn.go.html to see an example of how to use CreateOsn API.
func (client BlockchainPlatformClient) CreateOsn(ctx context.Context, request CreateOsnRequest) (response CreateOsnResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOsn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOsnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOsnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOsnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOsnResponse")
	}
	return
}

// createOsn implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) createOsn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/osns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOsnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/CreateOsn"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "CreateOsn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePeer Create Blockchain Platform Peer
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/CreatePeer.go.html to see an example of how to use CreatePeer API.
func (client BlockchainPlatformClient) CreatePeer(ctx context.Context, request CreatePeerRequest) (response CreatePeerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePeerResponse")
	}
	return
}

// createPeer implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) createPeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/peers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/CreatePeer"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "CreatePeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBlockchainPlatform Delete a particular of a Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/DeleteBlockchainPlatform.go.html to see an example of how to use DeleteBlockchainPlatform API.
func (client BlockchainPlatformClient) DeleteBlockchainPlatform(ctx context.Context, request DeleteBlockchainPlatformRequest) (response DeleteBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBlockchainPlatformResponse")
	}
	return
}

// deleteBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) deleteBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/blockchainPlatforms/{blockchainPlatformId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/DeleteBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "DeleteBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOsn Delete a particular OSN of a Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/DeleteOsn.go.html to see an example of how to use DeleteOsn API.
func (client BlockchainPlatformClient) DeleteOsn(ctx context.Context, request DeleteOsnRequest) (response DeleteOsnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOsn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOsnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOsnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOsnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOsnResponse")
	}
	return
}

// deleteOsn implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) deleteOsn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/blockchainPlatforms/{blockchainPlatformId}/osns/{osnId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOsnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/DeleteOsn"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "DeleteOsn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePeer Delete a particular peer of a Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/DeletePeer.go.html to see an example of how to use DeletePeer API.
func (client BlockchainPlatformClient) DeletePeer(ctx context.Context, request DeletePeerRequest) (response DeletePeerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deletePeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePeerResponse")
	}
	return
}

// deletePeer implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) deletePeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/blockchainPlatforms/{blockchainPlatformId}/peers/{peerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/DeletePeer"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "DeletePeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkRequest Attempts to cancel the work request with the given ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/DeleteWorkRequest.go.html to see an example of how to use DeleteWorkRequest API.
func (client BlockchainPlatformClient) DeleteWorkRequest(ctx context.Context, request DeleteWorkRequestRequest) (response DeleteWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client BlockchainPlatformClient) deleteWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/WorkRequest/DeleteWorkRequest"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "DeleteWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBlockchainPlatform Gets information about a Blockchain Platform identified by the specific id
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/GetBlockchainPlatform.go.html to see an example of how to use GetBlockchainPlatform API.
func (client BlockchainPlatformClient) GetBlockchainPlatform(ctx context.Context, request GetBlockchainPlatformRequest) (response GetBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBlockchainPlatformResponse")
	}
	return
}

// getBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) getBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/GetBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "GetBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOsn Gets information about an OSN identified by the specific id
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/GetOsn.go.html to see an example of how to use GetOsn API.
func (client BlockchainPlatformClient) GetOsn(ctx context.Context, request GetOsnRequest) (response GetOsnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOsn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOsnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOsnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOsnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOsnResponse")
	}
	return
}

// getOsn implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) getOsn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}/osns/{osnId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOsnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/Osn/GetOsn"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "GetOsn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPeer Gets information about a peer identified by the specific id
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/GetPeer.go.html to see an example of how to use GetPeer API.
func (client BlockchainPlatformClient) GetPeer(ctx context.Context, request GetPeerRequest) (response GetPeerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPeerResponse")
	}
	return
}

// getPeer implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) getPeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}/peers/{peerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/Peer/GetPeer"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "GetPeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client BlockchainPlatformClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client BlockchainPlatformClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBlockchainPlatformPatches List Blockchain Platform Patches
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListBlockchainPlatformPatches.go.html to see an example of how to use ListBlockchainPlatformPatches API.
func (client BlockchainPlatformClient) ListBlockchainPlatformPatches(ctx context.Context, request ListBlockchainPlatformPatchesRequest) (response ListBlockchainPlatformPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlockchainPlatformPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlockchainPlatformPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlockchainPlatformPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlockchainPlatformPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlockchainPlatformPatchesResponse")
	}
	return
}

// listBlockchainPlatformPatches implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listBlockchainPlatformPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBlockchainPlatformPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/ListBlockchainPlatformPatches"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListBlockchainPlatformPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBlockchainPlatforms Returns a list Blockchain Platform Instances in a compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListBlockchainPlatforms.go.html to see an example of how to use ListBlockchainPlatforms API.
func (client BlockchainPlatformClient) ListBlockchainPlatforms(ctx context.Context, request ListBlockchainPlatformsRequest) (response ListBlockchainPlatformsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlockchainPlatforms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlockchainPlatformsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlockchainPlatformsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlockchainPlatformsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlockchainPlatformsResponse")
	}
	return
}

// listBlockchainPlatforms implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listBlockchainPlatforms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBlockchainPlatformsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/ListBlockchainPlatforms"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListBlockchainPlatforms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOsns List Blockchain Platform OSNs
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListOsns.go.html to see an example of how to use ListOsns API.
func (client BlockchainPlatformClient) ListOsns(ctx context.Context, request ListOsnsRequest) (response ListOsnsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listOsns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOsnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOsnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOsnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOsnsResponse")
	}
	return
}

// listOsns implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listOsns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}/osns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOsnsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/Osn/ListOsns"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListOsns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPeers List Blockchain Platform Peers
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListPeers.go.html to see an example of how to use ListPeers API.
func (client BlockchainPlatformClient) ListPeers(ctx context.Context, request ListPeersRequest) (response ListPeersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listPeers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPeersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPeersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPeersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPeersResponse")
	}
	return
}

// listPeers implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) listPeers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blockchainPlatforms/{blockchainPlatformId}/peers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPeersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/Peer/ListPeers"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListPeers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client BlockchainPlatformClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client BlockchainPlatformClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client BlockchainPlatformClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client BlockchainPlatformClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client BlockchainPlatformClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client BlockchainPlatformClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PreviewScaleBlockchainPlatform Preview Scale Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/PreviewScaleBlockchainPlatform.go.html to see an example of how to use PreviewScaleBlockchainPlatform API.
func (client BlockchainPlatformClient) PreviewScaleBlockchainPlatform(ctx context.Context, request PreviewScaleBlockchainPlatformRequest) (response PreviewScaleBlockchainPlatformResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.previewScaleBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PreviewScaleBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PreviewScaleBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PreviewScaleBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PreviewScaleBlockchainPlatformResponse")
	}
	return
}

// previewScaleBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) previewScaleBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/scale/preview", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PreviewScaleBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/PreviewScaleBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "PreviewScaleBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScaleBlockchainPlatform Scale Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/ScaleBlockchainPlatform.go.html to see an example of how to use ScaleBlockchainPlatform API.
func (client BlockchainPlatformClient) ScaleBlockchainPlatform(ctx context.Context, request ScaleBlockchainPlatformRequest) (response ScaleBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scaleBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScaleBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScaleBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScaleBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScaleBlockchainPlatformResponse")
	}
	return
}

// scaleBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) scaleBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/scale", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScaleBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/ScaleBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "ScaleBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartBlockchainPlatform Start a Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/StartBlockchainPlatform.go.html to see an example of how to use StartBlockchainPlatform API.
func (client BlockchainPlatformClient) StartBlockchainPlatform(ctx context.Context, request StartBlockchainPlatformRequest) (response StartBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartBlockchainPlatformResponse")
	}
	return
}

// startBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) startBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/StartBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "StartBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopBlockchainPlatform Stop a Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/StopBlockchainPlatform.go.html to see an example of how to use StopBlockchainPlatform API.
func (client BlockchainPlatformClient) StopBlockchainPlatform(ctx context.Context, request StopBlockchainPlatformRequest) (response StopBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopBlockchainPlatformResponse")
	}
	return
}

// stopBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) stopBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/StopBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "StopBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBlockchainPlatform Update a particular of a Blockchain Platform
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/UpdateBlockchainPlatform.go.html to see an example of how to use UpdateBlockchainPlatform API.
func (client BlockchainPlatformClient) UpdateBlockchainPlatform(ctx context.Context, request UpdateBlockchainPlatformRequest) (response UpdateBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBlockchainPlatformResponse")
	}
	return
}

// updateBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) updateBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/blockchainPlatforms/{blockchainPlatformId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/UpdateBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "UpdateBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOsn Update Blockchain Platform OSN
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/UpdateOsn.go.html to see an example of how to use UpdateOsn API.
func (client BlockchainPlatformClient) UpdateOsn(ctx context.Context, request UpdateOsnRequest) (response UpdateOsnResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateOsn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOsnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOsnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOsnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOsnResponse")
	}
	return
}

// updateOsn implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) updateOsn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/blockchainPlatforms/{blockchainPlatformId}/osns/{osnId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOsnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/UpdateOsn"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "UpdateOsn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePeer Update Blockchain Platform Peer
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/UpdatePeer.go.html to see an example of how to use UpdatePeer API.
func (client BlockchainPlatformClient) UpdatePeer(ctx context.Context, request UpdatePeerRequest) (response UpdatePeerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePeerResponse")
	}
	return
}

// updatePeer implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) updatePeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/blockchainPlatforms/{blockchainPlatformId}/peers/{peerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/UpdatePeer"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "UpdatePeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpgradeBlockchainPlatform Upgrade a Blockchain Platform version
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/blockchain/UpgradeBlockchainPlatform.go.html to see an example of how to use UpgradeBlockchainPlatform API.
func (client BlockchainPlatformClient) UpgradeBlockchainPlatform(ctx context.Context, request UpgradeBlockchainPlatformRequest) (response UpgradeBlockchainPlatformResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.upgradeBlockchainPlatform, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpgradeBlockchainPlatformResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpgradeBlockchainPlatformResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpgradeBlockchainPlatformResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpgradeBlockchainPlatformResponse")
	}
	return
}

// upgradeBlockchainPlatform implements the OCIOperation interface (enables retrying operations)
func (client BlockchainPlatformClient) upgradeBlockchainPlatform(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blockchainPlatforms/{blockchainPlatformId}/actions/upgrade", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpgradeBlockchainPlatformResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/blockchain/20191010/BlockchainPlatform/UpgradeBlockchainPlatform"
		err = common.PostProcessServiceError(err, "BlockchainPlatform", "UpgradeBlockchainPlatform", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
