// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//RewardsClient a client for Rewards
type RewardsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRewardsClientWithConfigurationProvider Creates a new default Rewards client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRewardsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RewardsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRewardsClientFromBaseClient(baseClient, provider)
}

// NewRewardsClientWithOboToken Creates a new default Rewards client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewRewardsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RewardsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRewardsClientFromBaseClient(baseClient, configProvider)
}

func newRewardsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RewardsClient, err error) {
	// Rewards service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RewardsClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RewardsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("usage", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RewardsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RewardsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateRedeemableUser Add list of redeemable user email ids for a subscription Id
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/CreateRedeemableUser.go.html to see an example of how to use CreateRedeemableUser API.
func (client RewardsClient) CreateRedeemableUser(ctx context.Context, request CreateRedeemableUserRequest) (response CreateRedeemableUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRedeemableUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRedeemableUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRedeemableUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRedeemableUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRedeemableUserResponse")
	}
	return
}

// createRedeemableUser implements the OCIOperation interface (enables retrying operations)
func (client RewardsClient) createRedeemableUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/subscriptions/{subscriptionId}/redeemableUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRedeemableUserResponse
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

// DeleteRedeemableUser Delete list of redeemable user email ids for a subscription Id
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/DeleteRedeemableUser.go.html to see an example of how to use DeleteRedeemableUser API.
func (client RewardsClient) DeleteRedeemableUser(ctx context.Context, request DeleteRedeemableUserRequest) (response DeleteRedeemableUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRedeemableUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRedeemableUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRedeemableUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRedeemableUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRedeemableUserResponse")
	}
	return
}

// deleteRedeemableUser implements the OCIOperation interface (enables retrying operations)
func (client RewardsClient) deleteRedeemableUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/subscriptions/{subscriptionId}/redeemableUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRedeemableUserResponse
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

// ListProducts This API provides usage period specific product and its usage details.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListProducts.go.html to see an example of how to use ListProducts API.
func (client RewardsClient) ListProducts(ctx context.Context, request ListProductsRequest) (response ListProductsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProducts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProductsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProductsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProductsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProductsResponse")
	}
	return
}

// listProducts implements the OCIOperation interface (enables retrying operations)
func (client RewardsClient) listProducts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}/products", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProductsResponse
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

// ListRedeemableUsers Provides emailids of redeemable users for the given subscriptionId
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListRedeemableUsers.go.html to see an example of how to use ListRedeemableUsers API.
func (client RewardsClient) ListRedeemableUsers(ctx context.Context, request ListRedeemableUsersRequest) (response ListRedeemableUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRedeemableUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRedeemableUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRedeemableUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRedeemableUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRedeemableUsersResponse")
	}
	return
}

// listRedeemableUsers implements the OCIOperation interface (enables retrying operations)
func (client RewardsClient) listRedeemableUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}/redeemableUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRedeemableUsersResponse
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

// ListRewards This API returns list of rewards for a subscription Id
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListRewards.go.html to see an example of how to use ListRewards API.
func (client RewardsClient) ListRewards(ctx context.Context, request ListRewardsRequest) (response ListRewardsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRewards, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRewardsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRewardsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRewardsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRewardsResponse")
	}
	return
}

// listRewards implements the OCIOperation interface (enables retrying operations)
func (client RewardsClient) listRewards(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/subscriptions/{subscriptionId}/rewards", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRewardsResponse
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
