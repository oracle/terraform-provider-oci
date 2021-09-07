// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/common/auth"
	"net/http"
)

//FunctionsInvokeClient a client for FunctionsInvoke
type FunctionsInvokeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFunctionsInvokeClientWithConfigurationProvider Creates a new default FunctionsInvoke client with the given configuration provider.
// the configuration provider will be used for the default signer
func NewFunctionsInvokeClientWithConfigurationProvider(configProvider common.ConfigurationProvider, endpoint string) (client FunctionsInvokeClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newFunctionsInvokeClientFromBaseClient(baseClient, provider, endpoint)
}

// NewFunctionsInvokeClientWithOboToken Creates a new default FunctionsInvoke client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
func NewFunctionsInvokeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string, endpoint string) (client FunctionsInvokeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFunctionsInvokeClientFromBaseClient(baseClient, configProvider, endpoint)
}

func newFunctionsInvokeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider, endpoint string) (client FunctionsInvokeClient, err error) {
	client = FunctionsInvokeClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	client.Host = endpoint
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FunctionsInvokeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *FunctionsInvokeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// InvokeFunction Invokes a function
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/InvokeFunction.go.html to see an example of how to use InvokeFunction API.
func (client FunctionsInvokeClient) InvokeFunction(ctx context.Context, request InvokeFunctionRequest) (response InvokeFunctionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.invokeFunction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InvokeFunctionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InvokeFunctionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InvokeFunctionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InvokeFunctionResponse")
	}
	return
}

// invokeFunction implements the OCIOperation interface (enables retrying operations)
func (client FunctionsInvokeClient) invokeFunction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/functions/{functionId}/actions/invoke", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InvokeFunctionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
