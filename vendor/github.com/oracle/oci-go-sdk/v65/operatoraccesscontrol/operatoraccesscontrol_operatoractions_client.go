// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OperatorActionsClient a client for OperatorActions
type OperatorActionsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOperatorActionsClientWithConfigurationProvider Creates a new default OperatorActions client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOperatorActionsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OperatorActionsClient, err error) {
	if enabled := common.CheckForEnabledServices("operatoraccesscontrol"); !enabled {
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
	return newOperatorActionsClientFromBaseClient(baseClient, provider)
}

// NewOperatorActionsClientWithOboToken Creates a new default OperatorActions client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOperatorActionsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OperatorActionsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOperatorActionsClientFromBaseClient(baseClient, configProvider)
}

func newOperatorActionsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OperatorActionsClient, err error) {
	// OperatorActions service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OperatorActions"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OperatorActionsClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OperatorActionsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("operatoraccesscontrol", "https://operator-access-control.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OperatorActionsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OperatorActionsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetOperatorAction Gets the operator action associated with the specified operator action ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/GetOperatorAction.go.html to see an example of how to use GetOperatorAction API.
// A default retry strategy applies to this operation GetOperatorAction()
func (client OperatorActionsClient) GetOperatorAction(ctx context.Context, request GetOperatorActionRequest) (response GetOperatorActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOperatorAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOperatorActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOperatorActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOperatorActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOperatorActionResponse")
	}
	return
}

// getOperatorAction implements the OCIOperation interface (enables retrying operations)
func (client OperatorActionsClient) getOperatorAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorActions/{operatorActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOperatorActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operatoraccesscontrol/20200630/OperatorAction/GetOperatorAction"
		err = common.PostProcessServiceError(err, "OperatorActions", "GetOperatorAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOperatorActions Lists all the OperatorActions available in the system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorActions.go.html to see an example of how to use ListOperatorActions API.
// A default retry strategy applies to this operation ListOperatorActions()
func (client OperatorActionsClient) ListOperatorActions(ctx context.Context, request ListOperatorActionsRequest) (response ListOperatorActionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOperatorActions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOperatorActionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOperatorActionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOperatorActionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOperatorActionsResponse")
	}
	return
}

// listOperatorActions implements the OCIOperation interface (enables retrying operations)
func (client OperatorActionsClient) listOperatorActions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorActions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOperatorActionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operatoraccesscontrol/20200630/OperatorAction/ListOperatorActions"
		err = common.PostProcessServiceError(err, "OperatorActions", "ListOperatorActions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
