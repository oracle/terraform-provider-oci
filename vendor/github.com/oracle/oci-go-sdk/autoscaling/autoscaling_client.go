// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//AutoScalingClient a client for AutoScaling
type AutoScalingClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAutoScalingClientWithConfigurationProvider Creates a new default AutoScaling client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAutoScalingClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AutoScalingClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newAutoScalingClientFromBaseClient(baseClient, configProvider)
}

// NewAutoScalingClientWithOboToken Creates a new default AutoScaling client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAutoScalingClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AutoScalingClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newAutoScalingClientFromBaseClient(baseClient, configProvider)
}

func newAutoScalingClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AutoScalingClient, err error) {
	client = AutoScalingClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AutoScalingClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("None", "https://autoscaling.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AutoScalingClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AutoScalingClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAutoScalingConfigurationCompartment Moves an autoscaling configuration into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When you move an autoscaling configuration to a different compartment, associated resources such as instance
// pools are not moved.
func (client AutoScalingClient) ChangeAutoScalingConfigurationCompartment(ctx context.Context, request ChangeAutoScalingConfigurationCompartmentRequest) (response ChangeAutoScalingConfigurationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeAutoScalingConfigurationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAutoScalingConfigurationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAutoScalingConfigurationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAutoScalingConfigurationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAutoScalingConfigurationCompartmentResponse")
	}
	return
}

// changeAutoScalingConfigurationCompartment implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) changeAutoScalingConfigurationCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autoScalingConfigurations/{autoScalingConfigurationId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeAutoScalingConfigurationCompartmentResponse
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

// CreateAutoScalingConfiguration Creates an autoscaling configuration.
func (client AutoScalingClient) CreateAutoScalingConfiguration(ctx context.Context, request CreateAutoScalingConfigurationRequest) (response CreateAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutoScalingConfigurationResponse")
	}
	return
}

// createAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) createAutoScalingConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autoScalingConfigurations")
	if err != nil {
		return nil, err
	}

	var response CreateAutoScalingConfigurationResponse
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

// CreateAutoScalingPolicy Creates an autoscaling policy for the specified autoscaling configuration.
func (client AutoScalingClient) CreateAutoScalingPolicy(ctx context.Context, request CreateAutoScalingPolicyRequest) (response CreateAutoScalingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAutoScalingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAutoScalingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAutoScalingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAutoScalingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAutoScalingPolicyResponse")
	}
	return
}

// createAutoScalingPolicy implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) createAutoScalingPolicy(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/autoScalingConfigurations/{autoScalingConfigurationId}/policies")
	if err != nil {
		return nil, err
	}

	var response CreateAutoScalingPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &autoscalingpolicy{})
	return response, err
}

// DeleteAutoScalingConfiguration Deletes an autoscaling configuration.
func (client AutoScalingClient) DeleteAutoScalingConfiguration(ctx context.Context, request DeleteAutoScalingConfigurationRequest) (response DeleteAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutoScalingConfigurationResponse")
	}
	return
}

// deleteAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) deleteAutoScalingConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autoScalingConfigurations/{autoScalingConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAutoScalingConfigurationResponse
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

// DeleteAutoScalingPolicy Deletes an autoscaling policy for the specified autoscaling configuration.
func (client AutoScalingClient) DeleteAutoScalingPolicy(ctx context.Context, request DeleteAutoScalingPolicyRequest) (response DeleteAutoScalingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAutoScalingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAutoScalingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAutoScalingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAutoScalingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAutoScalingPolicyResponse")
	}
	return
}

// deleteAutoScalingPolicy implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) deleteAutoScalingPolicy(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/autoScalingConfigurations/{autoScalingConfigurationId}/policies/{autoScalingPolicyId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAutoScalingPolicyResponse
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

// GetAutoScalingConfiguration Gets information about the specified autoscaling configuration.
func (client AutoScalingClient) GetAutoScalingConfiguration(ctx context.Context, request GetAutoScalingConfigurationRequest) (response GetAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutoScalingConfigurationResponse")
	}
	return
}

// getAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) getAutoScalingConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autoScalingConfigurations/{autoScalingConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response GetAutoScalingConfigurationResponse
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

// GetAutoScalingPolicy Gets information about the specified autoscaling policy in the specified autoscaling configuration.
func (client AutoScalingClient) GetAutoScalingPolicy(ctx context.Context, request GetAutoScalingPolicyRequest) (response GetAutoScalingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutoScalingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutoScalingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutoScalingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutoScalingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutoScalingPolicyResponse")
	}
	return
}

// getAutoScalingPolicy implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) getAutoScalingPolicy(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autoScalingConfigurations/{autoScalingConfigurationId}/policies/{autoScalingPolicyId}")
	if err != nil {
		return nil, err
	}

	var response GetAutoScalingPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &autoscalingpolicy{})
	return response, err
}

// ListAutoScalingConfigurations Lists autoscaling configurations in the specifed compartment.
func (client AutoScalingClient) ListAutoScalingConfigurations(ctx context.Context, request ListAutoScalingConfigurationsRequest) (response ListAutoScalingConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutoScalingConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutoScalingConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutoScalingConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutoScalingConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutoScalingConfigurationsResponse")
	}
	return
}

// listAutoScalingConfigurations implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) listAutoScalingConfigurations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autoScalingConfigurations")
	if err != nil {
		return nil, err
	}

	var response ListAutoScalingConfigurationsResponse
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

// ListAutoScalingPolicies Lists the autoscaling policies in the specified autoscaling configuration.
func (client AutoScalingClient) ListAutoScalingPolicies(ctx context.Context, request ListAutoScalingPoliciesRequest) (response ListAutoScalingPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutoScalingPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutoScalingPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutoScalingPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutoScalingPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutoScalingPoliciesResponse")
	}
	return
}

// listAutoScalingPolicies implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) listAutoScalingPolicies(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/autoScalingConfigurations/{autoScalingConfigurationId}/policies")
	if err != nil {
		return nil, err
	}

	var response ListAutoScalingPoliciesResponse
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

// UpdateAutoScalingConfiguration Updates certain fields on the specified autoscaling configuration, such as the name, the cooldown period,
// and whether the autoscaling configuration is enabled.
func (client AutoScalingClient) UpdateAutoScalingConfiguration(ctx context.Context, request UpdateAutoScalingConfigurationRequest) (response UpdateAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutoScalingConfigurationResponse")
	}
	return
}

// updateAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) updateAutoScalingConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autoScalingConfigurations/{autoScalingConfigurationId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutoScalingConfigurationResponse
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

// UpdateAutoScalingPolicy Updates an autoscaling policy in the specified autoscaling configuration.
func (client AutoScalingClient) UpdateAutoScalingPolicy(ctx context.Context, request UpdateAutoScalingPolicyRequest) (response UpdateAutoScalingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateAutoScalingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutoScalingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutoScalingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutoScalingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutoScalingPolicyResponse")
	}
	return
}

// updateAutoScalingPolicy implements the OCIOperation interface (enables retrying operations)
func (client AutoScalingClient) updateAutoScalingPolicy(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/autoScalingConfigurations/{autoScalingConfigurationId}/policies/{autoScalingPolicyId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAutoScalingPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &autoscalingpolicy{})
	return response, err
}
