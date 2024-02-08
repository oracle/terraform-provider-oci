// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DataplaneClient a client for Dataplane
type DataplaneClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataplaneClientWithConfigurationProvider Creates a new default Dataplane client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataplaneClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataplaneClient, err error) {
	if enabled := common.CheckForEnabledServices("identitydataplane"); !enabled {
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
	return newDataplaneClientFromBaseClient(baseClient, provider)
}

// NewDataplaneClientWithOboToken Creates a new default Dataplane client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDataplaneClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataplaneClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataplaneClientFromBaseClient(baseClient, configProvider)
}

func newDataplaneClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataplaneClient, err error) {
	// Dataplane service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Dataplane"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataplaneClient{BaseClient: baseClient}
	client.BasePath = "v1"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataplaneClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("identitydataplane", "https://auth.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataplaneClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataplaneClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GenerateScopedAccessToken Based on the calling Principal and the input payload, derive the claims, and generate a scoped-access token for specific resources. For example, set scope to urn:oracle:db::id::<compartment-id> for access to a database in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydataplane/GenerateScopedAccessToken.go.html to see an example of how to use GenerateScopedAccessToken API.
func (client DataplaneClient) GenerateScopedAccessToken(ctx context.Context, request GenerateScopedAccessTokenRequest) (response GenerateScopedAccessTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateScopedAccessToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateScopedAccessTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateScopedAccessTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateScopedAccessTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateScopedAccessTokenResponse")
	}
	return
}

// generateScopedAccessToken implements the OCIOperation interface (enables retrying operations)
func (client DataplaneClient) generateScopedAccessToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/generateScopedAccessToken", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateScopedAccessTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/identity-dp/v1/SecurityToken/GenerateScopedAccessToken"
		err = common.PostProcessServiceError(err, "Dataplane", "GenerateScopedAccessToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateUserSecurityToken Exchanges a valid user token-based signature (API key and UPST) for a short-lived UPST of the authenticated
// user principal. When not specified, the user session duration is set to a default of 60 minutes in all realms. Resulting UPSTs
// are refreshable while the user session has not expired.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydataplane/GenerateUserSecurityToken.go.html to see an example of how to use GenerateUserSecurityToken API.
func (client DataplaneClient) GenerateUserSecurityToken(ctx context.Context, request GenerateUserSecurityTokenRequest) (response GenerateUserSecurityTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateUserSecurityToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateUserSecurityTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateUserSecurityTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateUserSecurityTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateUserSecurityTokenResponse")
	}
	return
}

// generateUserSecurityToken implements the OCIOperation interface (enables retrying operations)
func (client DataplaneClient) generateUserSecurityToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/token/upst/actions/GenerateUpst", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateUserSecurityTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/identity-dp/v1/SecurityToken/GenerateUserSecurityToken"
		err = common.PostProcessServiceError(err, "Dataplane", "GenerateUserSecurityToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
