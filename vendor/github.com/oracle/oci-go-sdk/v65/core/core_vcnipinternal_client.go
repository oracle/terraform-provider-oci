// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// VcnipInternalClient a client for VcnipInternal
type VcnipInternalClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewVcnipInternalClientWithConfigurationProvider Creates a new default VcnipInternal client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewVcnipInternalClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client VcnipInternalClient, err error) {
	if enabled := common.CheckForEnabledServices("core"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newVcnipInternalClientFromBaseClient(baseClient, provider)
}

// NewVcnipInternalClientWithOboToken Creates a new default VcnipInternal client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewVcnipInternalClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client VcnipInternalClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newVcnipInternalClientFromBaseClient(baseClient, configProvider)
}

func newVcnipInternalClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client VcnipInternalClient, err error) {
	// VcnipInternal service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("VcnipInternal"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = VcnipInternalClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *VcnipInternalClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("iaas", "https://iaas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *VcnipInternalClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *VcnipInternalClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ByoipRangeLock Lock ByoipRange
func (client VcnipInternalClient) ByoipRangeLock(ctx context.Context, request ByoipRangeLockRequest) (response ByoipRangeLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.byoipRangeLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ByoipRangeLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ByoipRangeLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ByoipRangeLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ByoipRangeLockResponse")
	}
	return
}

// byoipRangeLock implements the OCIOperation interface (enables retrying operations)
func (client VcnipInternalClient) byoipRangeLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byoipRanges/{byoipRangeId}/actions/lock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ByoipRangeLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ByoipRangeResponse/ByoipRangeLock"
		err = common.PostProcessServiceError(err, "VcnipInternal", "ByoipRangeLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ByoipRangeUnlock Lock ByoipRange
func (client VcnipInternalClient) ByoipRangeUnlock(ctx context.Context, request ByoipRangeUnlockRequest) (response ByoipRangeUnlockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.byoipRangeUnlock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ByoipRangeUnlockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ByoipRangeUnlockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ByoipRangeUnlockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ByoipRangeUnlockResponse")
	}
	return
}

// byoipRangeUnlock implements the OCIOperation interface (enables retrying operations)
func (client VcnipInternalClient) byoipRangeUnlock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byoipRanges/{byoipRangeId}/actions/unlock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ByoipRangeUnlockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ByoipRangeResponse/ByoipRangeUnlock"
		err = common.PostProcessServiceError(err, "VcnipInternal", "ByoipRangeUnlock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
