// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Subscription, Commitment and and Rate Card Details
//
// Set of APIs that return the Subscription Details, Commitment and Effective Rate Card Details
//

package osubsubscription

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CommitmentClient a client for Commitment
type CommitmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCommitmentClientWithConfigurationProvider Creates a new default Commitment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCommitmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CommitmentClient, err error) {
	if enabled := common.CheckForEnabledServices("osubsubscription"); !enabled {
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
	return newCommitmentClientFromBaseClient(baseClient, provider)
}

// NewCommitmentClientWithOboToken Creates a new default Commitment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCommitmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CommitmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCommitmentClientFromBaseClient(baseClient, configProvider)
}

func newCommitmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CommitmentClient, err error) {
	// Commitment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Commitment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CommitmentClient{BaseClient: baseClient}
	client.BasePath = "oalapp/service/onesubs/proxy/20210501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CommitmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osubsubscription", "https://csaap-e.oracle.com")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CommitmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CommitmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetCommitment This API returns the commitment details corresponding to the id provided
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubsubscription/GetCommitment.go.html to see an example of how to use GetCommitment API.
func (client CommitmentClient) GetCommitment(ctx context.Context, request GetCommitmentRequest) (response GetCommitmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCommitment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCommitmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCommitmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCommitmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCommitmentResponse")
	}
	return
}

// getCommitment implements the OCIOperation interface (enables retrying operations)
func (client CommitmentClient) getCommitment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/commitments/{commitmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCommitmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Commitment", "GetCommitment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCommitments This list API returns all commitments for a particular Subscribed Service
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubsubscription/ListCommitments.go.html to see an example of how to use ListCommitments API.
func (client CommitmentClient) ListCommitments(ctx context.Context, request ListCommitmentsRequest) (response ListCommitmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCommitments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCommitmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCommitmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCommitmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCommitmentsResponse")
	}
	return
}

// listCommitments implements the OCIOperation interface (enables retrying operations)
func (client CommitmentClient) listCommitments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/commitments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCommitmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Commitment", "ListCommitments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
