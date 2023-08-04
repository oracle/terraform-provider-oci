// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// IdentityDomainsClient a client for IdentityDomains
type IdentityDomainsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIdentityDomainsClientWithConfigurationProvider Creates a new default IdentityDomains client with the given configuration provider.
// the configuration provider will be used for the default signer
func NewIdentityDomainsClientWithConfigurationProvider(configProvider common.ConfigurationProvider, endpoint string) (client IdentityDomainsClient, err error) {
	if enabled := common.CheckForEnabledServices("identitydomains"); !enabled {
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
	return newIdentityDomainsClientFromBaseClient(baseClient, provider, endpoint)
}

// NewIdentityDomainsClientWithOboToken Creates a new default IdentityDomains client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
func NewIdentityDomainsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string, endpoint string) (client IdentityDomainsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newIdentityDomainsClientFromBaseClient(baseClient, configProvider, endpoint)
}

func newIdentityDomainsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider, endpoint string) (client IdentityDomainsClient, err error) {
	// IdentityDomains service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("IdentityDomains"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = IdentityDomainsClient{BaseClient: baseClient}
	client.BasePath = ""
	client.Host = endpoint
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IdentityDomainsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *IdentityDomainsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateApiKey Create a user's API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateApiKey.go.html to see an example of how to use CreateApiKey API.
func (client IdentityDomainsClient) CreateApiKey(ctx context.Context, request CreateApiKeyRequest) (response CreateApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApiKeyResponse")
	}
	return
}

// createApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/ApiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateApp Create an App
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateApp.go.html to see an example of how to use CreateApp API.
func (client IdentityDomainsClient) CreateApp(ctx context.Context, request CreateAppRequest) (response CreateAppResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createApp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAppResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAppResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAppResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAppResponse")
	}
	return
}

// createApp implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createApp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Apps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAppResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateApp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAppRole Create an AppRole
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateAppRole.go.html to see an example of how to use CreateAppRole API.
func (client IdentityDomainsClient) CreateAppRole(ctx context.Context, request CreateAppRoleRequest) (response CreateAppRoleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAppRole, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAppRoleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAppRoleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAppRoleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAppRoleResponse")
	}
	return
}

// createAppRole implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createAppRole(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AppRoles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAppRoleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateAppRole", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAuthToken Create a user's Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateAuthToken.go.html to see an example of how to use CreateAuthToken API.
func (client IdentityDomainsClient) CreateAuthToken(ctx context.Context, request CreateAuthTokenRequest) (response CreateAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAuthTokenResponse")
	}
	return
}

// createAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AuthTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAuthenticationFactorsRemover Remove All Authentication Factor Channels for a User
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateAuthenticationFactorsRemover.go.html to see an example of how to use CreateAuthenticationFactorsRemover API.
func (client IdentityDomainsClient) CreateAuthenticationFactorsRemover(ctx context.Context, request CreateAuthenticationFactorsRemoverRequest) (response CreateAuthenticationFactorsRemoverResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAuthenticationFactorsRemover, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAuthenticationFactorsRemoverResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAuthenticationFactorsRemoverResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAuthenticationFactorsRemoverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAuthenticationFactorsRemoverResponse")
	}
	return
}

// createAuthenticationFactorsRemover implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createAuthenticationFactorsRemover(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AuthenticationFactorsRemover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAuthenticationFactorsRemoverResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateAuthenticationFactorsRemover", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCustomerSecretKey Create a user's customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateCustomerSecretKey.go.html to see an example of how to use CreateCustomerSecretKey API.
func (client IdentityDomainsClient) CreateCustomerSecretKey(ctx context.Context, request CreateCustomerSecretKeyRequest) (response CreateCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCustomerSecretKeyResponse")
	}
	return
}

// createCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/CustomerSecretKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDynamicResourceGroup Create a Dynamic Resource Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateDynamicResourceGroup.go.html to see an example of how to use CreateDynamicResourceGroup API.
func (client IdentityDomainsClient) CreateDynamicResourceGroup(ctx context.Context, request CreateDynamicResourceGroupRequest) (response CreateDynamicResourceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDynamicResourceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDynamicResourceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDynamicResourceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDynamicResourceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDynamicResourceGroupResponse")
	}
	return
}

// createDynamicResourceGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createDynamicResourceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/DynamicResourceGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDynamicResourceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateDynamicResourceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGrant Add a Grantee to an AppRole
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateGrant.go.html to see an example of how to use CreateGrant API.
func (client IdentityDomainsClient) CreateGrant(ctx context.Context, request CreateGrantRequest) (response CreateGrantResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGrant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGrantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGrantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGrantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGrantResponse")
	}
	return
}

// createGrant implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createGrant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Grants", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGrantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateGrant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGroup Create a group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateGroup.go.html to see an example of how to use CreateGroup API.
func (client IdentityDomainsClient) CreateGroup(ctx context.Context, request CreateGroupRequest) (response CreateGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGroupResponse")
	}
	return
}

// createGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Groups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateIdentityProvider Create an Identity Provider
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateIdentityProvider.go.html to see an example of how to use CreateIdentityProvider API.
func (client IdentityDomainsClient) CreateIdentityProvider(ctx context.Context, request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIdentityProviderResponse")
	}
	return
}

// createIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/IdentityProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateIdentityProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMe Self register a user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMe.go.html to see an example of how to use CreateMe API.
func (client IdentityDomainsClient) CreateMe(ctx context.Context, request CreateMeRequest) (response CreateMeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMeResponse")
	}
	return
}

// createMe implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Me", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyApiKey Add a user's own API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyApiKey.go.html to see an example of how to use CreateMyApiKey API.
func (client IdentityDomainsClient) CreateMyApiKey(ctx context.Context, request CreateMyApiKeyRequest) (response CreateMyApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyApiKeyResponse")
	}
	return
}

// createMyApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyApiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyAuthToken Create a user's own Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyAuthToken.go.html to see an example of how to use CreateMyAuthToken API.
func (client IdentityDomainsClient) CreateMyAuthToken(ctx context.Context, request CreateMyAuthTokenRequest) (response CreateMyAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyAuthTokenResponse")
	}
	return
}

// createMyAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyAuthTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyAuthenticationFactorInitiator Initiate Self Service Enrollment using the Requested MFA Factor
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyAuthenticationFactorInitiator.go.html to see an example of how to use CreateMyAuthenticationFactorInitiator API.
func (client IdentityDomainsClient) CreateMyAuthenticationFactorInitiator(ctx context.Context, request CreateMyAuthenticationFactorInitiatorRequest) (response CreateMyAuthenticationFactorInitiatorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyAuthenticationFactorInitiator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyAuthenticationFactorInitiatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyAuthenticationFactorInitiatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyAuthenticationFactorInitiatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyAuthenticationFactorInitiatorResponse")
	}
	return
}

// createMyAuthenticationFactorInitiator implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyAuthenticationFactorInitiator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyAuthenticationFactorInitiator", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyAuthenticationFactorInitiatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyAuthenticationFactorInitiator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyAuthenticationFactorValidator Validate Self Service Enrollment using the Requested MFA Factor
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyAuthenticationFactorValidator.go.html to see an example of how to use CreateMyAuthenticationFactorValidator API.
func (client IdentityDomainsClient) CreateMyAuthenticationFactorValidator(ctx context.Context, request CreateMyAuthenticationFactorValidatorRequest) (response CreateMyAuthenticationFactorValidatorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyAuthenticationFactorValidator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyAuthenticationFactorValidatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyAuthenticationFactorValidatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyAuthenticationFactorValidatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyAuthenticationFactorValidatorResponse")
	}
	return
}

// createMyAuthenticationFactorValidator implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyAuthenticationFactorValidator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyAuthenticationFactorValidator", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyAuthenticationFactorValidatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyAuthenticationFactorValidator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyAuthenticationFactorsRemover Remove All Authentication Factor Channels for a User
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyAuthenticationFactorsRemover.go.html to see an example of how to use CreateMyAuthenticationFactorsRemover API.
func (client IdentityDomainsClient) CreateMyAuthenticationFactorsRemover(ctx context.Context, request CreateMyAuthenticationFactorsRemoverRequest) (response CreateMyAuthenticationFactorsRemoverResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyAuthenticationFactorsRemover, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyAuthenticationFactorsRemoverResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyAuthenticationFactorsRemoverResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyAuthenticationFactorsRemoverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyAuthenticationFactorsRemoverResponse")
	}
	return
}

// createMyAuthenticationFactorsRemover implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyAuthenticationFactorsRemover(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyAuthenticationFactorsRemover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyAuthenticationFactorsRemoverResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyAuthenticationFactorsRemover", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyCustomerSecretKey Add a user's own customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyCustomerSecretKey.go.html to see an example of how to use CreateMyCustomerSecretKey API.
func (client IdentityDomainsClient) CreateMyCustomerSecretKey(ctx context.Context, request CreateMyCustomerSecretKeyRequest) (response CreateMyCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyCustomerSecretKeyResponse")
	}
	return
}

// createMyCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyCustomerSecretKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyOAuth2ClientCredential Create a user's own OAuth2 client credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyOAuth2ClientCredential.go.html to see an example of how to use CreateMyOAuth2ClientCredential API.
func (client IdentityDomainsClient) CreateMyOAuth2ClientCredential(ctx context.Context, request CreateMyOAuth2ClientCredentialRequest) (response CreateMyOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyOAuth2ClientCredentialResponse")
	}
	return
}

// createMyOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyOAuth2ClientCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyRequest Create a Request
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyRequest.go.html to see an example of how to use CreateMyRequest API.
func (client IdentityDomainsClient) CreateMyRequest(ctx context.Context, request CreateMyRequestRequest) (response CreateMyRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyRequestResponse")
	}
	return
}

// createMyRequest implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMySmtpCredential Create a user's own SMTP credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMySmtpCredential.go.html to see an example of how to use CreateMySmtpCredential API.
func (client IdentityDomainsClient) CreateMySmtpCredential(ctx context.Context, request CreateMySmtpCredentialRequest) (response CreateMySmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMySmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMySmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMySmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMySmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMySmtpCredentialResponse")
	}
	return
}

// createMySmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMySmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MySmtpCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMySmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMySmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMySupportAccount Create a user's own support account.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMySupportAccount.go.html to see an example of how to use CreateMySupportAccount API.
func (client IdentityDomainsClient) CreateMySupportAccount(ctx context.Context, request CreateMySupportAccountRequest) (response CreateMySupportAccountResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMySupportAccount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMySupportAccountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMySupportAccountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMySupportAccountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMySupportAccountResponse")
	}
	return
}

// createMySupportAccount implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMySupportAccount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MySupportAccounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMySupportAccountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMySupportAccount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMyUserDbCredential Create a user's own database (DB) credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateMyUserDbCredential.go.html to see an example of how to use CreateMyUserDbCredential API.
func (client IdentityDomainsClient) CreateMyUserDbCredential(ctx context.Context, request CreateMyUserDbCredentialRequest) (response CreateMyUserDbCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMyUserDbCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMyUserDbCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMyUserDbCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMyUserDbCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMyUserDbCredentialResponse")
	}
	return
}

// createMyUserDbCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createMyUserDbCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyUserDbCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMyUserDbCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateMyUserDbCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOAuth2ClientCredential Add a user's OAuth2 client credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateOAuth2ClientCredential.go.html to see an example of how to use CreateOAuth2ClientCredential API.
func (client IdentityDomainsClient) CreateOAuth2ClientCredential(ctx context.Context, request CreateOAuth2ClientCredentialRequest) (response CreateOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOAuth2ClientCredentialResponse")
	}
	return
}

// createOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/OAuth2ClientCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePasswordPolicy Create a password policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreatePasswordPolicy.go.html to see an example of how to use CreatePasswordPolicy API.
func (client IdentityDomainsClient) CreatePasswordPolicy(ctx context.Context, request CreatePasswordPolicyRequest) (response CreatePasswordPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPasswordPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePasswordPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePasswordPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePasswordPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePasswordPolicyResponse")
	}
	return
}

// createPasswordPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createPasswordPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/PasswordPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePasswordPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreatePasswordPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityQuestion Create a security question.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateSecurityQuestion.go.html to see an example of how to use CreateSecurityQuestion API.
func (client IdentityDomainsClient) CreateSecurityQuestion(ctx context.Context, request CreateSecurityQuestionRequest) (response CreateSecurityQuestionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityQuestion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityQuestionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityQuestionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityQuestionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityQuestionResponse")
	}
	return
}

// createSecurityQuestion implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createSecurityQuestion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/SecurityQuestions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityQuestionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateSecurityQuestion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSmtpCredential Create a user's SMTP credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateSmtpCredential.go.html to see an example of how to use CreateSmtpCredential API.
func (client IdentityDomainsClient) CreateSmtpCredential(ctx context.Context, request CreateSmtpCredentialRequest) (response CreateSmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSmtpCredentialResponse")
	}
	return
}

// createSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/SmtpCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateSmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUser Create a user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateUser.go.html to see an example of how to use CreateUser API.
func (client IdentityDomainsClient) CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserResponse")
	}
	return
}

// createUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUserDbCredential Create a user's database (DB) credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/CreateUserDbCredential.go.html to see an example of how to use CreateUserDbCredential API.
func (client IdentityDomainsClient) CreateUserDbCredential(ctx context.Context, request CreateUserDbCredentialRequest) (response CreateUserDbCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUserDbCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserDbCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserDbCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserDbCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserDbCredentialResponse")
	}
	return
}

// createUserDbCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) createUserDbCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/UserDbCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserDbCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "CreateUserDbCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApiKey Delete a user's API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteApiKey.go.html to see an example of how to use DeleteApiKey API.
func (client IdentityDomainsClient) DeleteApiKey(ctx context.Context, request DeleteApiKeyRequest) (response DeleteApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApiKeyResponse")
	}
	return
}

// deleteApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/ApiKeys/{apiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApp Delete an App
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteApp.go.html to see an example of how to use DeleteApp API.
func (client IdentityDomainsClient) DeleteApp(ctx context.Context, request DeleteAppRequest) (response DeleteAppResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteApp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAppResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAppResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAppResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAppResponse")
	}
	return
}

// deleteApp implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteApp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/Apps/{appId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAppResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteApp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAppRole Delete an AppRole
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteAppRole.go.html to see an example of how to use DeleteAppRole API.
func (client IdentityDomainsClient) DeleteAppRole(ctx context.Context, request DeleteAppRoleRequest) (response DeleteAppRoleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteAppRole, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAppRoleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAppRoleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAppRoleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAppRoleResponse")
	}
	return
}

// deleteAppRole implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteAppRole(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/AppRoles/{appRoleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAppRoleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteAppRole", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAuthToken Delete a user's Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteAuthToken.go.html to see an example of how to use DeleteAuthToken API.
func (client IdentityDomainsClient) DeleteAuthToken(ctx context.Context, request DeleteAuthTokenRequest) (response DeleteAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAuthTokenResponse")
	}
	return
}

// deleteAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/AuthTokens/{authTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCustomerSecretKey Delete a user's customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteCustomerSecretKey.go.html to see an example of how to use DeleteCustomerSecretKey API.
func (client IdentityDomainsClient) DeleteCustomerSecretKey(ctx context.Context, request DeleteCustomerSecretKeyRequest) (response DeleteCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCustomerSecretKeyResponse")
	}
	return
}

// deleteCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/CustomerSecretKeys/{customerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDynamicResourceGroup Delete a Dynamic Resource Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteDynamicResourceGroup.go.html to see an example of how to use DeleteDynamicResourceGroup API.
func (client IdentityDomainsClient) DeleteDynamicResourceGroup(ctx context.Context, request DeleteDynamicResourceGroupRequest) (response DeleteDynamicResourceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteDynamicResourceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDynamicResourceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDynamicResourceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDynamicResourceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDynamicResourceGroupResponse")
	}
	return
}

// deleteDynamicResourceGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteDynamicResourceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/DynamicResourceGroups/{dynamicResourceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDynamicResourceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteDynamicResourceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGrant Remove a Grantee from an AppRole
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteGrant.go.html to see an example of how to use DeleteGrant API.
func (client IdentityDomainsClient) DeleteGrant(ctx context.Context, request DeleteGrantRequest) (response DeleteGrantResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteGrant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGrantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGrantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGrantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGrantResponse")
	}
	return
}

// deleteGrant implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteGrant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/Grants/{grantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGrantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteGrant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGroup Delete a group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteGroup.go.html to see an example of how to use DeleteGroup API.
func (client IdentityDomainsClient) DeleteGroup(ctx context.Context, request DeleteGroupRequest) (response DeleteGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGroupResponse")
	}
	return
}

// deleteGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/Groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteIdentityProvider Delete an Identity Provider
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteIdentityProvider.go.html to see an example of how to use DeleteIdentityProvider API.
func (client IdentityDomainsClient) DeleteIdentityProvider(ctx context.Context, request DeleteIdentityProviderRequest) (response DeleteIdentityProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIdentityProviderResponse")
	}
	return
}

// deleteIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/IdentityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteIdentityProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyApiKey Delete a user's own API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyApiKey.go.html to see an example of how to use DeleteMyApiKey API.
func (client IdentityDomainsClient) DeleteMyApiKey(ctx context.Context, request DeleteMyApiKeyRequest) (response DeleteMyApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyApiKeyResponse")
	}
	return
}

// deleteMyApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyApiKeys/{myApiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyAuthToken Delete a user's own Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyAuthToken.go.html to see an example of how to use DeleteMyAuthToken API.
func (client IdentityDomainsClient) DeleteMyAuthToken(ctx context.Context, request DeleteMyAuthTokenRequest) (response DeleteMyAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyAuthTokenResponse")
	}
	return
}

// deleteMyAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyAuthTokens/{myAuthTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyCustomerSecretKey Delete a user's own customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyCustomerSecretKey.go.html to see an example of how to use DeleteMyCustomerSecretKey API.
func (client IdentityDomainsClient) DeleteMyCustomerSecretKey(ctx context.Context, request DeleteMyCustomerSecretKeyRequest) (response DeleteMyCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyCustomerSecretKeyResponse")
	}
	return
}

// deleteMyCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyCustomerSecretKeys/{myCustomerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyDevice Delete a Device
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyDevice.go.html to see an example of how to use DeleteMyDevice API.
func (client IdentityDomainsClient) DeleteMyDevice(ctx context.Context, request DeleteMyDeviceRequest) (response DeleteMyDeviceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyDeviceResponse")
	}
	return
}

// deleteMyDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyDevices/{myDeviceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyOAuth2ClientCredential Delete a user's own OAuth2 client credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyOAuth2ClientCredential.go.html to see an example of how to use DeleteMyOAuth2ClientCredential API.
func (client IdentityDomainsClient) DeleteMyOAuth2ClientCredential(ctx context.Context, request DeleteMyOAuth2ClientCredentialRequest) (response DeleteMyOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyOAuth2ClientCredentialResponse")
	}
	return
}

// deleteMyOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyOAuth2ClientCredentials/{myOAuth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMySmtpCredential Delete a user's own SMTP credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMySmtpCredential.go.html to see an example of how to use DeleteMySmtpCredential API.
func (client IdentityDomainsClient) DeleteMySmtpCredential(ctx context.Context, request DeleteMySmtpCredentialRequest) (response DeleteMySmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMySmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMySmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMySmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMySmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMySmtpCredentialResponse")
	}
	return
}

// deleteMySmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMySmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MySmtpCredentials/{mySmtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMySmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMySmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMySupportAccount Delete a user's own support account.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMySupportAccount.go.html to see an example of how to use DeleteMySupportAccount API.
func (client IdentityDomainsClient) DeleteMySupportAccount(ctx context.Context, request DeleteMySupportAccountRequest) (response DeleteMySupportAccountResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMySupportAccount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMySupportAccountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMySupportAccountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMySupportAccountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMySupportAccountResponse")
	}
	return
}

// deleteMySupportAccount implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMySupportAccount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MySupportAccounts/{mySupportAccountId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMySupportAccountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMySupportAccount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyTrustedUserAgent Delete a Trusted User Agent
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyTrustedUserAgent.go.html to see an example of how to use DeleteMyTrustedUserAgent API.
func (client IdentityDomainsClient) DeleteMyTrustedUserAgent(ctx context.Context, request DeleteMyTrustedUserAgentRequest) (response DeleteMyTrustedUserAgentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyTrustedUserAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyTrustedUserAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyTrustedUserAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyTrustedUserAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyTrustedUserAgentResponse")
	}
	return
}

// deleteMyTrustedUserAgent implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyTrustedUserAgent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyTrustedUserAgents/{myTrustedUserAgentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyTrustedUserAgentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyTrustedUserAgent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMyUserDbCredential Delete a user's own database (DB) credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteMyUserDbCredential.go.html to see an example of how to use DeleteMyUserDbCredential API.
func (client IdentityDomainsClient) DeleteMyUserDbCredential(ctx context.Context, request DeleteMyUserDbCredentialRequest) (response DeleteMyUserDbCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteMyUserDbCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMyUserDbCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMyUserDbCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMyUserDbCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMyUserDbCredentialResponse")
	}
	return
}

// deleteMyUserDbCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteMyUserDbCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/MyUserDbCredentials/{myUserDbCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMyUserDbCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteMyUserDbCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOAuth2ClientCredential Delete a user's OAuth2 client credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteOAuth2ClientCredential.go.html to see an example of how to use DeleteOAuth2ClientCredential API.
func (client IdentityDomainsClient) DeleteOAuth2ClientCredential(ctx context.Context, request DeleteOAuth2ClientCredentialRequest) (response DeleteOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOAuth2ClientCredentialResponse")
	}
	return
}

// deleteOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/OAuth2ClientCredentials/{oAuth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePasswordPolicy Delete a password policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeletePasswordPolicy.go.html to see an example of how to use DeletePasswordPolicy API.
func (client IdentityDomainsClient) DeletePasswordPolicy(ctx context.Context, request DeletePasswordPolicyRequest) (response DeletePasswordPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deletePasswordPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePasswordPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePasswordPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePasswordPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePasswordPolicyResponse")
	}
	return
}

// deletePasswordPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deletePasswordPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/PasswordPolicies/{passwordPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePasswordPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeletePasswordPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityQuestion Delete a security question.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteSecurityQuestion.go.html to see an example of how to use DeleteSecurityQuestion API.
func (client IdentityDomainsClient) DeleteSecurityQuestion(ctx context.Context, request DeleteSecurityQuestionRequest) (response DeleteSecurityQuestionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityQuestion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityQuestionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityQuestionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityQuestionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityQuestionResponse")
	}
	return
}

// deleteSecurityQuestion implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteSecurityQuestion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/SecurityQuestions/{securityQuestionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityQuestionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteSecurityQuestion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSmtpCredential Delete a user's SMTP credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteSmtpCredential.go.html to see an example of how to use DeleteSmtpCredential API.
func (client IdentityDomainsClient) DeleteSmtpCredential(ctx context.Context, request DeleteSmtpCredentialRequest) (response DeleteSmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSmtpCredentialResponse")
	}
	return
}

// deleteSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/SmtpCredentials/{smtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteSmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUser Delete a user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteUser.go.html to see an example of how to use DeleteUser API.
func (client IdentityDomainsClient) DeleteUser(ctx context.Context, request DeleteUserRequest) (response DeleteUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserResponse")
	}
	return
}

// deleteUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/Users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUserDbCredential Delete a user's database (DB) credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/DeleteUserDbCredential.go.html to see an example of how to use DeleteUserDbCredential API.
func (client IdentityDomainsClient) DeleteUserDbCredential(ctx context.Context, request DeleteUserDbCredentialRequest) (response DeleteUserDbCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteUserDbCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserDbCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserDbCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserDbCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserDbCredentialResponse")
	}
	return
}

// deleteUserDbCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) deleteUserDbCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/admin/v1/UserDbCredentials/{userDbCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserDbCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "DeleteUserDbCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAccountMgmtInfo Get Account Mgmt Info
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetAccountMgmtInfo.go.html to see an example of how to use GetAccountMgmtInfo API.
func (client IdentityDomainsClient) GetAccountMgmtInfo(ctx context.Context, request GetAccountMgmtInfoRequest) (response GetAccountMgmtInfoResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAccountMgmtInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAccountMgmtInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAccountMgmtInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAccountMgmtInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAccountMgmtInfoResponse")
	}
	return
}

// getAccountMgmtInfo implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getAccountMgmtInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AccountMgmtInfos/{accountMgmtInfoId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAccountMgmtInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetAccountMgmtInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAccountRecoverySetting Get an account recovery setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetAccountRecoverySetting.go.html to see an example of how to use GetAccountRecoverySetting API.
func (client IdentityDomainsClient) GetAccountRecoverySetting(ctx context.Context, request GetAccountRecoverySettingRequest) (response GetAccountRecoverySettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAccountRecoverySetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAccountRecoverySettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAccountRecoverySettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAccountRecoverySettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAccountRecoverySettingResponse")
	}
	return
}

// getAccountRecoverySetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getAccountRecoverySetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AccountRecoverySettings/{accountRecoverySettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAccountRecoverySettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetAccountRecoverySetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApiKey Get a user's API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetApiKey.go.html to see an example of how to use GetApiKey API.
func (client IdentityDomainsClient) GetApiKey(ctx context.Context, request GetApiKeyRequest) (response GetApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApiKeyResponse")
	}
	return
}

// getApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/ApiKeys/{apiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApp Get an App
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetApp.go.html to see an example of how to use GetApp API.
func (client IdentityDomainsClient) GetApp(ctx context.Context, request GetAppRequest) (response GetAppResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getApp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAppResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAppResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAppResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAppResponse")
	}
	return
}

// getApp implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getApp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Apps/{appId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAppResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetApp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAppRole Get an AppRole
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetAppRole.go.html to see an example of how to use GetAppRole API.
func (client IdentityDomainsClient) GetAppRole(ctx context.Context, request GetAppRoleRequest) (response GetAppRoleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAppRole, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAppRoleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAppRoleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAppRoleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAppRoleResponse")
	}
	return
}

// getAppRole implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getAppRole(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AppRoles/{appRoleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAppRoleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetAppRole", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuthToken Get a user's Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetAuthToken.go.html to see an example of how to use GetAuthToken API.
func (client IdentityDomainsClient) GetAuthToken(ctx context.Context, request GetAuthTokenRequest) (response GetAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuthTokenResponse")
	}
	return
}

// getAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AuthTokens/{authTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuthenticationFactorSetting Get Authentication Factor Settings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetAuthenticationFactorSetting.go.html to see an example of how to use GetAuthenticationFactorSetting API.
func (client IdentityDomainsClient) GetAuthenticationFactorSetting(ctx context.Context, request GetAuthenticationFactorSettingRequest) (response GetAuthenticationFactorSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getAuthenticationFactorSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuthenticationFactorSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuthenticationFactorSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuthenticationFactorSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuthenticationFactorSettingResponse")
	}
	return
}

// getAuthenticationFactorSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getAuthenticationFactorSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AuthenticationFactorSettings/{authenticationFactorSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuthenticationFactorSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetAuthenticationFactorSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCustomerSecretKey Get a user's customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetCustomerSecretKey.go.html to see an example of how to use GetCustomerSecretKey API.
func (client IdentityDomainsClient) GetCustomerSecretKey(ctx context.Context, request GetCustomerSecretKeyRequest) (response GetCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCustomerSecretKeyResponse")
	}
	return
}

// getCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/CustomerSecretKeys/{customerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDynamicResourceGroup Get a Dynamic Resource Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetDynamicResourceGroup.go.html to see an example of how to use GetDynamicResourceGroup API.
func (client IdentityDomainsClient) GetDynamicResourceGroup(ctx context.Context, request GetDynamicResourceGroupRequest) (response GetDynamicResourceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getDynamicResourceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDynamicResourceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDynamicResourceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDynamicResourceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDynamicResourceGroupResponse")
	}
	return
}

// getDynamicResourceGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getDynamicResourceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/DynamicResourceGroups/{dynamicResourceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDynamicResourceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetDynamicResourceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGrant Get a Grant
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetGrant.go.html to see an example of how to use GetGrant API.
func (client IdentityDomainsClient) GetGrant(ctx context.Context, request GetGrantRequest) (response GetGrantResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getGrant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGrantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGrantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGrantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGrantResponse")
	}
	return
}

// getGrant implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getGrant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Grants/{grantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGrantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetGrant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGroup Get a group. <b>Important:</b> The Group SEARCH and GET operations on users and members will throw an exception if the response has more than 10,000 members. To avoid the exception, use the pagination filter to GET or SEARCH group members.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetGroup.go.html to see an example of how to use GetGroup API.
func (client IdentityDomainsClient) GetGroup(ctx context.Context, request GetGroupRequest) (response GetGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGroupResponse")
	}
	return
}

// getGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIdentityProvider Get an Identity Provider
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetIdentityProvider.go.html to see an example of how to use GetIdentityProvider API.
func (client IdentityDomainsClient) GetIdentityProvider(ctx context.Context, request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIdentityProviderResponse")
	}
	return
}

// getIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/IdentityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetIdentityProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIdentitySetting Get an Identity setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetIdentitySetting.go.html to see an example of how to use GetIdentitySetting API.
func (client IdentityDomainsClient) GetIdentitySetting(ctx context.Context, request GetIdentitySettingRequest) (response GetIdentitySettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getIdentitySetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIdentitySettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIdentitySettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIdentitySettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIdentitySettingResponse")
	}
	return
}

// getIdentitySetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getIdentitySetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/IdentitySettings/{identitySettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIdentitySettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetIdentitySetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetKmsiSetting Get KmsiSettings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetKmsiSetting.go.html to see an example of how to use GetKmsiSetting API.
func (client IdentityDomainsClient) GetKmsiSetting(ctx context.Context, request GetKmsiSettingRequest) (response GetKmsiSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getKmsiSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetKmsiSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetKmsiSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetKmsiSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetKmsiSettingResponse")
	}
	return
}

// getKmsiSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getKmsiSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/KmsiSettings/{kmsiSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetKmsiSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetKmsiSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMe Get a user's own information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMe.go.html to see an example of how to use GetMe API.
func (client IdentityDomainsClient) GetMe(ctx context.Context, request GetMeRequest) (response GetMeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMeResponse")
	}
	return
}

// getMe implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Me", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyApiKey Get a user's own API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyApiKey.go.html to see an example of how to use GetMyApiKey API.
func (client IdentityDomainsClient) GetMyApiKey(ctx context.Context, request GetMyApiKeyRequest) (response GetMyApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyApiKeyResponse")
	}
	return
}

// getMyApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyApiKeys/{myApiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyAuthToken Get a user's own Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyAuthToken.go.html to see an example of how to use GetMyAuthToken API.
func (client IdentityDomainsClient) GetMyAuthToken(ctx context.Context, request GetMyAuthTokenRequest) (response GetMyAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyAuthTokenResponse")
	}
	return
}

// getMyAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyAuthTokens/{myAuthTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyCustomerSecretKey Get a user's own customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyCustomerSecretKey.go.html to see an example of how to use GetMyCustomerSecretKey API.
func (client IdentityDomainsClient) GetMyCustomerSecretKey(ctx context.Context, request GetMyCustomerSecretKeyRequest) (response GetMyCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyCustomerSecretKeyResponse")
	}
	return
}

// getMyCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyCustomerSecretKeys/{myCustomerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyDevice Get a Device
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyDevice.go.html to see an example of how to use GetMyDevice API.
func (client IdentityDomainsClient) GetMyDevice(ctx context.Context, request GetMyDeviceRequest) (response GetMyDeviceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyDeviceResponse")
	}
	return
}

// getMyDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyDevices/{myDeviceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyOAuth2ClientCredential Get a user's own OAuth2 client credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyOAuth2ClientCredential.go.html to see an example of how to use GetMyOAuth2ClientCredential API.
func (client IdentityDomainsClient) GetMyOAuth2ClientCredential(ctx context.Context, request GetMyOAuth2ClientCredentialRequest) (response GetMyOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyOAuth2ClientCredentialResponse")
	}
	return
}

// getMyOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyOAuth2ClientCredentials/{myOAuth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMySmtpCredential Get a user's own SMTP credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMySmtpCredential.go.html to see an example of how to use GetMySmtpCredential API.
func (client IdentityDomainsClient) GetMySmtpCredential(ctx context.Context, request GetMySmtpCredentialRequest) (response GetMySmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMySmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMySmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMySmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMySmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMySmtpCredentialResponse")
	}
	return
}

// getMySmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMySmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MySmtpCredentials/{mySmtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMySmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMySmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMySupportAccount Get a user's own support account.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMySupportAccount.go.html to see an example of how to use GetMySupportAccount API.
func (client IdentityDomainsClient) GetMySupportAccount(ctx context.Context, request GetMySupportAccountRequest) (response GetMySupportAccountResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMySupportAccount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMySupportAccountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMySupportAccountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMySupportAccountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMySupportAccountResponse")
	}
	return
}

// getMySupportAccount implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMySupportAccount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MySupportAccounts/{mySupportAccountId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMySupportAccountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMySupportAccount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyTrustedUserAgent Get a Trusted User Agent
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyTrustedUserAgent.go.html to see an example of how to use GetMyTrustedUserAgent API.
func (client IdentityDomainsClient) GetMyTrustedUserAgent(ctx context.Context, request GetMyTrustedUserAgentRequest) (response GetMyTrustedUserAgentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyTrustedUserAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyTrustedUserAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyTrustedUserAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyTrustedUserAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyTrustedUserAgentResponse")
	}
	return
}

// getMyTrustedUserAgent implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyTrustedUserAgent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyTrustedUserAgents/{myTrustedUserAgentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyTrustedUserAgentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyTrustedUserAgent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMyUserDbCredential Get a user's own database (DB) credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetMyUserDbCredential.go.html to see an example of how to use GetMyUserDbCredential API.
func (client IdentityDomainsClient) GetMyUserDbCredential(ctx context.Context, request GetMyUserDbCredentialRequest) (response GetMyUserDbCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getMyUserDbCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMyUserDbCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMyUserDbCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMyUserDbCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMyUserDbCredentialResponse")
	}
	return
}

// getMyUserDbCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getMyUserDbCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyUserDbCredentials/{myUserDbCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMyUserDbCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetMyUserDbCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOAuth2ClientCredential Get a user's OAuth2 client credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetOAuth2ClientCredential.go.html to see an example of how to use GetOAuth2ClientCredential API.
func (client IdentityDomainsClient) GetOAuth2ClientCredential(ctx context.Context, request GetOAuth2ClientCredentialRequest) (response GetOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOAuth2ClientCredentialResponse")
	}
	return
}

// getOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/OAuth2ClientCredentials/{oAuth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPasswordPolicy Get a password policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetPasswordPolicy.go.html to see an example of how to use GetPasswordPolicy API.
func (client IdentityDomainsClient) GetPasswordPolicy(ctx context.Context, request GetPasswordPolicyRequest) (response GetPasswordPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getPasswordPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPasswordPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPasswordPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPasswordPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPasswordPolicyResponse")
	}
	return
}

// getPasswordPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getPasswordPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/PasswordPolicies/{passwordPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPasswordPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetPasswordPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityQuestion Get a security question.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetSecurityQuestion.go.html to see an example of how to use GetSecurityQuestion API.
func (client IdentityDomainsClient) GetSecurityQuestion(ctx context.Context, request GetSecurityQuestionRequest) (response GetSecurityQuestionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getSecurityQuestion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityQuestionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityQuestionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityQuestionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityQuestionResponse")
	}
	return
}

// getSecurityQuestion implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getSecurityQuestion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/SecurityQuestions/{securityQuestionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityQuestionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetSecurityQuestion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityQuestionSetting Get a security question setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetSecurityQuestionSetting.go.html to see an example of how to use GetSecurityQuestionSetting API.
func (client IdentityDomainsClient) GetSecurityQuestionSetting(ctx context.Context, request GetSecurityQuestionSettingRequest) (response GetSecurityQuestionSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getSecurityQuestionSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityQuestionSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityQuestionSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityQuestionSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityQuestionSettingResponse")
	}
	return
}

// getSecurityQuestionSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getSecurityQuestionSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/SecurityQuestionSettings/{securityQuestionSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityQuestionSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetSecurityQuestionSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSmtpCredential Get a user's SMTP credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetSmtpCredential.go.html to see an example of how to use GetSmtpCredential API.
func (client IdentityDomainsClient) GetSmtpCredential(ctx context.Context, request GetSmtpCredentialRequest) (response GetSmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSmtpCredentialResponse")
	}
	return
}

// getSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/SmtpCredentials/{smtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetSmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUser Get a user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetUser.go.html to see an example of how to use GetUser API.
func (client IdentityDomainsClient) GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserResponse")
	}
	return
}

// getUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserAttributesSetting Get User Schema Attribute Settings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetUserAttributesSetting.go.html to see an example of how to use GetUserAttributesSetting API.
func (client IdentityDomainsClient) GetUserAttributesSetting(ctx context.Context, request GetUserAttributesSettingRequest) (response GetUserAttributesSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getUserAttributesSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserAttributesSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserAttributesSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserAttributesSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserAttributesSettingResponse")
	}
	return
}

// getUserAttributesSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getUserAttributesSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/UserAttributesSettings/{userAttributesSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserAttributesSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetUserAttributesSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserDbCredential Get a user's database (DB) credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/GetUserDbCredential.go.html to see an example of how to use GetUserDbCredential API.
func (client IdentityDomainsClient) GetUserDbCredential(ctx context.Context, request GetUserDbCredentialRequest) (response GetUserDbCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getUserDbCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserDbCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserDbCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserDbCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserDbCredentialResponse")
	}
	return
}

// getUserDbCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) getUserDbCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/UserDbCredentials/{userDbCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserDbCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "GetUserDbCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAccountMgmtInfos Search Account Mgmt Info
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListAccountMgmtInfos.go.html to see an example of how to use ListAccountMgmtInfos API.
func (client IdentityDomainsClient) ListAccountMgmtInfos(ctx context.Context, request ListAccountMgmtInfosRequest) (response ListAccountMgmtInfosResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAccountMgmtInfos, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAccountMgmtInfosResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAccountMgmtInfosResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAccountMgmtInfosResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAccountMgmtInfosResponse")
	}
	return
}

// listAccountMgmtInfos implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listAccountMgmtInfos(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AccountMgmtInfos", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAccountMgmtInfosResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListAccountMgmtInfos", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAccountRecoverySettings Search for account recovery settings.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListAccountRecoverySettings.go.html to see an example of how to use ListAccountRecoverySettings API.
func (client IdentityDomainsClient) ListAccountRecoverySettings(ctx context.Context, request ListAccountRecoverySettingsRequest) (response ListAccountRecoverySettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAccountRecoverySettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAccountRecoverySettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAccountRecoverySettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAccountRecoverySettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAccountRecoverySettingsResponse")
	}
	return
}

// listAccountRecoverySettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listAccountRecoverySettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AccountRecoverySettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAccountRecoverySettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListAccountRecoverySettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApiKeys Search API keys.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListApiKeys.go.html to see an example of how to use ListApiKeys API.
func (client IdentityDomainsClient) ListApiKeys(ctx context.Context, request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listApiKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApiKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApiKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApiKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApiKeysResponse")
	}
	return
}

// listApiKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listApiKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/ApiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApiKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListApiKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAppRoles Search AppRoles
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListAppRoles.go.html to see an example of how to use ListAppRoles API.
func (client IdentityDomainsClient) ListAppRoles(ctx context.Context, request ListAppRolesRequest) (response ListAppRolesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAppRoles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAppRolesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAppRolesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAppRolesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAppRolesResponse")
	}
	return
}

// listAppRoles implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listAppRoles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AppRoles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAppRolesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListAppRoles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApps Search Apps
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListApps.go.html to see an example of how to use ListApps API.
func (client IdentityDomainsClient) ListApps(ctx context.Context, request ListAppsRequest) (response ListAppsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listApps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAppsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAppsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAppsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAppsResponse")
	}
	return
}

// listApps implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listApps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Apps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAppsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListApps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuthTokens Search for Auth tokens.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListAuthTokens.go.html to see an example of how to use ListAuthTokens API.
func (client IdentityDomainsClient) ListAuthTokens(ctx context.Context, request ListAuthTokensRequest) (response ListAuthTokensResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAuthTokens, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuthTokensResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuthTokensResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuthTokensResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuthTokensResponse")
	}
	return
}

// listAuthTokens implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listAuthTokens(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AuthTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuthTokensResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListAuthTokens", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuthenticationFactorSettings Search Authentication Factor Settings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListAuthenticationFactorSettings.go.html to see an example of how to use ListAuthenticationFactorSettings API.
func (client IdentityDomainsClient) ListAuthenticationFactorSettings(ctx context.Context, request ListAuthenticationFactorSettingsRequest) (response ListAuthenticationFactorSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAuthenticationFactorSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuthenticationFactorSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuthenticationFactorSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuthenticationFactorSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuthenticationFactorSettingsResponse")
	}
	return
}

// listAuthenticationFactorSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listAuthenticationFactorSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/AuthenticationFactorSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuthenticationFactorSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListAuthenticationFactorSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCustomerSecretKeys Search for a user's customer secret keys.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListCustomerSecretKeys.go.html to see an example of how to use ListCustomerSecretKeys API.
func (client IdentityDomainsClient) ListCustomerSecretKeys(ctx context.Context, request ListCustomerSecretKeysRequest) (response ListCustomerSecretKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listCustomerSecretKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCustomerSecretKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCustomerSecretKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCustomerSecretKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCustomerSecretKeysResponse")
	}
	return
}

// listCustomerSecretKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listCustomerSecretKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/CustomerSecretKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCustomerSecretKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListCustomerSecretKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDynamicResourceGroups Search for Dynamic Resource Groups.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListDynamicResourceGroups.go.html to see an example of how to use ListDynamicResourceGroups API.
func (client IdentityDomainsClient) ListDynamicResourceGroups(ctx context.Context, request ListDynamicResourceGroupsRequest) (response ListDynamicResourceGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listDynamicResourceGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDynamicResourceGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDynamicResourceGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDynamicResourceGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDynamicResourceGroupsResponse")
	}
	return
}

// listDynamicResourceGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listDynamicResourceGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/DynamicResourceGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDynamicResourceGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListDynamicResourceGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGrants Search Grants
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListGrants.go.html to see an example of how to use ListGrants API.
func (client IdentityDomainsClient) ListGrants(ctx context.Context, request ListGrantsRequest) (response ListGrantsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listGrants, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGrantsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGrantsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGrantsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGrantsResponse")
	}
	return
}

// listGrants implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listGrants(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Grants", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGrantsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListGrants", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGroups Search for groups. <b>Important:</b> The Group SEARCH and GET operations on users and members will throw an exception if the response has more than 10,000 members. To avoid the exception, use the pagination filter to GET or SEARCH group members.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListGroups.go.html to see an example of how to use ListGroups API.
func (client IdentityDomainsClient) ListGroups(ctx context.Context, request ListGroupsRequest) (response ListGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGroupsResponse")
	}
	return
}

// listGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Groups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIdentityProviders Search Identity Providers
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListIdentityProviders.go.html to see an example of how to use ListIdentityProviders API.
func (client IdentityDomainsClient) ListIdentityProviders(ctx context.Context, request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listIdentityProviders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIdentityProvidersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIdentityProvidersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIdentityProvidersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIdentityProvidersResponse")
	}
	return
}

// listIdentityProviders implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listIdentityProviders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/IdentityProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIdentityProvidersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListIdentityProviders", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIdentitySettings Search for Identity settings.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListIdentitySettings.go.html to see an example of how to use ListIdentitySettings API.
func (client IdentityDomainsClient) ListIdentitySettings(ctx context.Context, request ListIdentitySettingsRequest) (response ListIdentitySettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listIdentitySettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIdentitySettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIdentitySettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIdentitySettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIdentitySettingsResponse")
	}
	return
}

// listIdentitySettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listIdentitySettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/IdentitySettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIdentitySettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListIdentitySettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListKmsiSettings Search KmsiSettings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListKmsiSettings.go.html to see an example of how to use ListKmsiSettings API.
func (client IdentityDomainsClient) ListKmsiSettings(ctx context.Context, request ListKmsiSettingsRequest) (response ListKmsiSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listKmsiSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListKmsiSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListKmsiSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListKmsiSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKmsiSettingsResponse")
	}
	return
}

// listKmsiSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listKmsiSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/KmsiSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListKmsiSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListKmsiSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyApiKeys Search for a user's own API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyApiKeys.go.html to see an example of how to use ListMyApiKeys API.
func (client IdentityDomainsClient) ListMyApiKeys(ctx context.Context, request ListMyApiKeysRequest) (response ListMyApiKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyApiKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyApiKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyApiKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyApiKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyApiKeysResponse")
	}
	return
}

// listMyApiKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyApiKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyApiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyApiKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyApiKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyApps Search My Apps
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyApps.go.html to see an example of how to use ListMyApps API.
func (client IdentityDomainsClient) ListMyApps(ctx context.Context, request ListMyAppsRequest) (response ListMyAppsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyApps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyAppsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyAppsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyAppsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyAppsResponse")
	}
	return
}

// listMyApps implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyApps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyApps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyAppsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyApps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyAuthTokens Search for a user's own Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyAuthTokens.go.html to see an example of how to use ListMyAuthTokens API.
func (client IdentityDomainsClient) ListMyAuthTokens(ctx context.Context, request ListMyAuthTokensRequest) (response ListMyAuthTokensResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyAuthTokens, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyAuthTokensResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyAuthTokensResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyAuthTokensResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyAuthTokensResponse")
	}
	return
}

// listMyAuthTokens implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyAuthTokens(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyAuthTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyAuthTokensResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyAuthTokens", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyCustomerSecretKeys Search for a user's own customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyCustomerSecretKeys.go.html to see an example of how to use ListMyCustomerSecretKeys API.
func (client IdentityDomainsClient) ListMyCustomerSecretKeys(ctx context.Context, request ListMyCustomerSecretKeysRequest) (response ListMyCustomerSecretKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyCustomerSecretKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyCustomerSecretKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyCustomerSecretKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyCustomerSecretKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyCustomerSecretKeysResponse")
	}
	return
}

// listMyCustomerSecretKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyCustomerSecretKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyCustomerSecretKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyCustomerSecretKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyCustomerSecretKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyDevices Search Devices
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyDevices.go.html to see an example of how to use ListMyDevices API.
func (client IdentityDomainsClient) ListMyDevices(ctx context.Context, request ListMyDevicesRequest) (response ListMyDevicesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyDevices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyDevicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyDevicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyDevicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyDevicesResponse")
	}
	return
}

// listMyDevices implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyDevices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyDevicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyDevices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyGroups Search for 'My Groups'.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyGroups.go.html to see an example of how to use ListMyGroups API.
func (client IdentityDomainsClient) ListMyGroups(ctx context.Context, request ListMyGroupsRequest) (response ListMyGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyGroupsResponse")
	}
	return
}

// listMyGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyOAuth2ClientCredentials Search for a user's own OAuth2 client credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyOAuth2ClientCredentials.go.html to see an example of how to use ListMyOAuth2ClientCredentials API.
func (client IdentityDomainsClient) ListMyOAuth2ClientCredentials(ctx context.Context, request ListMyOAuth2ClientCredentialsRequest) (response ListMyOAuth2ClientCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyOAuth2ClientCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyOAuth2ClientCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyOAuth2ClientCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyOAuth2ClientCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyOAuth2ClientCredentialsResponse")
	}
	return
}

// listMyOAuth2ClientCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyOAuth2ClientCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyOAuth2ClientCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyOAuth2ClientCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyOAuth2ClientCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyRequestableGroups Search My Requestable Groups
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyRequestableGroups.go.html to see an example of how to use ListMyRequestableGroups API.
func (client IdentityDomainsClient) ListMyRequestableGroups(ctx context.Context, request ListMyRequestableGroupsRequest) (response ListMyRequestableGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyRequestableGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyRequestableGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyRequestableGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyRequestableGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyRequestableGroupsResponse")
	}
	return
}

// listMyRequestableGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyRequestableGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyRequestableGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyRequestableGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyRequestableGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyRequests Search My Requests
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyRequests.go.html to see an example of how to use ListMyRequests API.
func (client IdentityDomainsClient) ListMyRequests(ctx context.Context, request ListMyRequestsRequest) (response ListMyRequestsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyRequestsResponse")
	}
	return
}

// listMyRequests implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMySmtpCredentials Search for a user's own SMTP credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMySmtpCredentials.go.html to see an example of how to use ListMySmtpCredentials API.
func (client IdentityDomainsClient) ListMySmtpCredentials(ctx context.Context, request ListMySmtpCredentialsRequest) (response ListMySmtpCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMySmtpCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMySmtpCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMySmtpCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMySmtpCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMySmtpCredentialsResponse")
	}
	return
}

// listMySmtpCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMySmtpCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MySmtpCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMySmtpCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMySmtpCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMySupportAccounts Search for a user's own support account.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMySupportAccounts.go.html to see an example of how to use ListMySupportAccounts API.
func (client IdentityDomainsClient) ListMySupportAccounts(ctx context.Context, request ListMySupportAccountsRequest) (response ListMySupportAccountsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMySupportAccounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMySupportAccountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMySupportAccountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMySupportAccountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMySupportAccountsResponse")
	}
	return
}

// listMySupportAccounts implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMySupportAccounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MySupportAccounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMySupportAccountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMySupportAccounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyTrustedUserAgents Search Trusted User Agents
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyTrustedUserAgents.go.html to see an example of how to use ListMyTrustedUserAgents API.
func (client IdentityDomainsClient) ListMyTrustedUserAgents(ctx context.Context, request ListMyTrustedUserAgentsRequest) (response ListMyTrustedUserAgentsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyTrustedUserAgents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyTrustedUserAgentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyTrustedUserAgentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyTrustedUserAgentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyTrustedUserAgentsResponse")
	}
	return
}

// listMyTrustedUserAgents implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyTrustedUserAgents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyTrustedUserAgents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyTrustedUserAgentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyTrustedUserAgents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMyUserDbCredentials Search for a user's own database (DB) credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMyUserDbCredentials.go.html to see an example of how to use ListMyUserDbCredentials API.
func (client IdentityDomainsClient) ListMyUserDbCredentials(ctx context.Context, request ListMyUserDbCredentialsRequest) (response ListMyUserDbCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listMyUserDbCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMyUserDbCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMyUserDbCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMyUserDbCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMyUserDbCredentialsResponse")
	}
	return
}

// listMyUserDbCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listMyUserDbCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/MyUserDbCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMyUserDbCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListMyUserDbCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOAuth2ClientCredentials Search for a user's OAuth2 client credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListOAuth2ClientCredentials.go.html to see an example of how to use ListOAuth2ClientCredentials API.
func (client IdentityDomainsClient) ListOAuth2ClientCredentials(ctx context.Context, request ListOAuth2ClientCredentialsRequest) (response ListOAuth2ClientCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listOAuth2ClientCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOAuth2ClientCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOAuth2ClientCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOAuth2ClientCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOAuth2ClientCredentialsResponse")
	}
	return
}

// listOAuth2ClientCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listOAuth2ClientCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/OAuth2ClientCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOAuth2ClientCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListOAuth2ClientCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPasswordPolicies Search for password policies.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListPasswordPolicies.go.html to see an example of how to use ListPasswordPolicies API.
func (client IdentityDomainsClient) ListPasswordPolicies(ctx context.Context, request ListPasswordPoliciesRequest) (response ListPasswordPoliciesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listPasswordPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPasswordPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPasswordPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPasswordPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPasswordPoliciesResponse")
	}
	return
}

// listPasswordPolicies implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listPasswordPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/PasswordPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPasswordPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListPasswordPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceTypeSchemaAttributes Search Resource Type Schema Attributes
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListResourceTypeSchemaAttributes.go.html to see an example of how to use ListResourceTypeSchemaAttributes API.
func (client IdentityDomainsClient) ListResourceTypeSchemaAttributes(ctx context.Context, request ListResourceTypeSchemaAttributesRequest) (response ListResourceTypeSchemaAttributesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listResourceTypeSchemaAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceTypeSchemaAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceTypeSchemaAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceTypeSchemaAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceTypeSchemaAttributesResponse")
	}
	return
}

// listResourceTypeSchemaAttributes implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listResourceTypeSchemaAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/ResourceTypeSchemaAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceTypeSchemaAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListResourceTypeSchemaAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityQuestionSettings Search for security question settings.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListSecurityQuestionSettings.go.html to see an example of how to use ListSecurityQuestionSettings API.
func (client IdentityDomainsClient) ListSecurityQuestionSettings(ctx context.Context, request ListSecurityQuestionSettingsRequest) (response ListSecurityQuestionSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listSecurityQuestionSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityQuestionSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityQuestionSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityQuestionSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityQuestionSettingsResponse")
	}
	return
}

// listSecurityQuestionSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listSecurityQuestionSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/SecurityQuestionSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityQuestionSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListSecurityQuestionSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityQuestions Search for security questions.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListSecurityQuestions.go.html to see an example of how to use ListSecurityQuestions API.
func (client IdentityDomainsClient) ListSecurityQuestions(ctx context.Context, request ListSecurityQuestionsRequest) (response ListSecurityQuestionsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listSecurityQuestions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityQuestionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityQuestionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityQuestionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityQuestionsResponse")
	}
	return
}

// listSecurityQuestions implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listSecurityQuestions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/SecurityQuestions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityQuestionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListSecurityQuestions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSmtpCredentials Search for SMTP credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListSmtpCredentials.go.html to see an example of how to use ListSmtpCredentials API.
func (client IdentityDomainsClient) ListSmtpCredentials(ctx context.Context, request ListSmtpCredentialsRequest) (response ListSmtpCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listSmtpCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSmtpCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSmtpCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSmtpCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSmtpCredentialsResponse")
	}
	return
}

// listSmtpCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listSmtpCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/SmtpCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSmtpCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListSmtpCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserAttributesSettings Search User Schema Attribute Settings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListUserAttributesSettings.go.html to see an example of how to use ListUserAttributesSettings API.
func (client IdentityDomainsClient) ListUserAttributesSettings(ctx context.Context, request ListUserAttributesSettingsRequest) (response ListUserAttributesSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listUserAttributesSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserAttributesSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserAttributesSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserAttributesSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserAttributesSettingsResponse")
	}
	return
}

// listUserAttributesSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listUserAttributesSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/UserAttributesSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserAttributesSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListUserAttributesSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserDbCredentials Search for a user's database (DB) credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListUserDbCredentials.go.html to see an example of how to use ListUserDbCredentials API.
func (client IdentityDomainsClient) ListUserDbCredentials(ctx context.Context, request ListUserDbCredentialsRequest) (response ListUserDbCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listUserDbCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserDbCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserDbCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserDbCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserDbCredentialsResponse")
	}
	return
}

// listUserDbCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listUserDbCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/UserDbCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserDbCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListUserDbCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUsers Search for users.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListUsers.go.html to see an example of how to use ListUsers API.
func (client IdentityDomainsClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsersResponse")
	}
	return
}

// listUsers implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) listUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/admin/v1/Users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "ListUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchAccountRecoverySetting Update an account recovery setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchAccountRecoverySetting.go.html to see an example of how to use PatchAccountRecoverySetting API.
func (client IdentityDomainsClient) PatchAccountRecoverySetting(ctx context.Context, request PatchAccountRecoverySettingRequest) (response PatchAccountRecoverySettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchAccountRecoverySetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchAccountRecoverySettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchAccountRecoverySettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchAccountRecoverySettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchAccountRecoverySettingResponse")
	}
	return
}

// patchAccountRecoverySetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchAccountRecoverySetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/AccountRecoverySettings/{accountRecoverySettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchAccountRecoverySettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchAccountRecoverySetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchApiKey Update a user's API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchApiKey.go.html to see an example of how to use PatchApiKey API.
func (client IdentityDomainsClient) PatchApiKey(ctx context.Context, request PatchApiKeyRequest) (response PatchApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchApiKeyResponse")
	}
	return
}

// patchApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/ApiKeys/{apiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchApp Update an App
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchApp.go.html to see an example of how to use PatchApp API.
func (client IdentityDomainsClient) PatchApp(ctx context.Context, request PatchAppRequest) (response PatchAppResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchApp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchAppResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchAppResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchAppResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchAppResponse")
	}
	return
}

// patchApp implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchApp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/Apps/{appId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchAppResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchApp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchAppRole Update an AppRole
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchAppRole.go.html to see an example of how to use PatchAppRole API.
func (client IdentityDomainsClient) PatchAppRole(ctx context.Context, request PatchAppRoleRequest) (response PatchAppRoleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchAppRole, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchAppRoleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchAppRoleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchAppRoleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchAppRoleResponse")
	}
	return
}

// patchAppRole implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchAppRole(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/AppRoles/{appRoleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchAppRoleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchAppRole", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchAuthToken Update a user's Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchAuthToken.go.html to see an example of how to use PatchAuthToken API.
func (client IdentityDomainsClient) PatchAuthToken(ctx context.Context, request PatchAuthTokenRequest) (response PatchAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchAuthTokenResponse")
	}
	return
}

// patchAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/AuthTokens/{authTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchCustomerSecretKey Update a user's customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchCustomerSecretKey.go.html to see an example of how to use PatchCustomerSecretKey API.
func (client IdentityDomainsClient) PatchCustomerSecretKey(ctx context.Context, request PatchCustomerSecretKeyRequest) (response PatchCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchCustomerSecretKeyResponse")
	}
	return
}

// patchCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/CustomerSecretKeys/{customerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchDynamicResourceGroup Update a Dynamic Resource Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchDynamicResourceGroup.go.html to see an example of how to use PatchDynamicResourceGroup API.
func (client IdentityDomainsClient) PatchDynamicResourceGroup(ctx context.Context, request PatchDynamicResourceGroupRequest) (response PatchDynamicResourceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchDynamicResourceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchDynamicResourceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchDynamicResourceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchDynamicResourceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchDynamicResourceGroupResponse")
	}
	return
}

// patchDynamicResourceGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchDynamicResourceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/DynamicResourceGroups/{dynamicResourceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchDynamicResourceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchDynamicResourceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchGrant Update a Grant
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchGrant.go.html to see an example of how to use PatchGrant API.
func (client IdentityDomainsClient) PatchGrant(ctx context.Context, request PatchGrantRequest) (response PatchGrantResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchGrant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchGrantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchGrantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchGrantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchGrantResponse")
	}
	return
}

// patchGrant implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchGrant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/Grants/{grantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchGrantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchGrant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchGroup Update a group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchGroup.go.html to see an example of how to use PatchGroup API.
func (client IdentityDomainsClient) PatchGroup(ctx context.Context, request PatchGroupRequest) (response PatchGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchGroupResponse")
	}
	return
}

// patchGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/Groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchIdentityProvider Update an Identity Provider
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchIdentityProvider.go.html to see an example of how to use PatchIdentityProvider API.
func (client IdentityDomainsClient) PatchIdentityProvider(ctx context.Context, request PatchIdentityProviderRequest) (response PatchIdentityProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchIdentityProviderResponse")
	}
	return
}

// patchIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/IdentityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchIdentityProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchIdentitySetting Update an Identity setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchIdentitySetting.go.html to see an example of how to use PatchIdentitySetting API.
func (client IdentityDomainsClient) PatchIdentitySetting(ctx context.Context, request PatchIdentitySettingRequest) (response PatchIdentitySettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchIdentitySetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchIdentitySettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchIdentitySettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchIdentitySettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchIdentitySettingResponse")
	}
	return
}

// patchIdentitySetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchIdentitySetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/IdentitySettings/{identitySettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchIdentitySettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchIdentitySetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchKmsiSetting Update a Setting
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchKmsiSetting.go.html to see an example of how to use PatchKmsiSetting API.
func (client IdentityDomainsClient) PatchKmsiSetting(ctx context.Context, request PatchKmsiSettingRequest) (response PatchKmsiSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchKmsiSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchKmsiSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchKmsiSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchKmsiSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchKmsiSettingResponse")
	}
	return
}

// patchKmsiSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchKmsiSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/KmsiSettings/{kmsiSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchKmsiSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchKmsiSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMe Update a user's own information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMe.go.html to see an example of how to use PatchMe API.
func (client IdentityDomainsClient) PatchMe(ctx context.Context, request PatchMeRequest) (response PatchMeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMeResponse")
	}
	return
}

// patchMe implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/Me", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMyApiKey Update a user's own API key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMyApiKey.go.html to see an example of how to use PatchMyApiKey API.
func (client IdentityDomainsClient) PatchMyApiKey(ctx context.Context, request PatchMyApiKeyRequest) (response PatchMyApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMyApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMyApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMyApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMyApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMyApiKeyResponse")
	}
	return
}

// patchMyApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMyApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/MyApiKeys/{myApiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMyApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMyApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMyAuthToken Update a user's own Auth token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMyAuthToken.go.html to see an example of how to use PatchMyAuthToken API.
func (client IdentityDomainsClient) PatchMyAuthToken(ctx context.Context, request PatchMyAuthTokenRequest) (response PatchMyAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMyAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMyAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMyAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMyAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMyAuthTokenResponse")
	}
	return
}

// patchMyAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMyAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/MyAuthTokens/{myAuthTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMyAuthTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMyAuthToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMyCustomerSecretKey Update a user's own customer secret key.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMyCustomerSecretKey.go.html to see an example of how to use PatchMyCustomerSecretKey API.
func (client IdentityDomainsClient) PatchMyCustomerSecretKey(ctx context.Context, request PatchMyCustomerSecretKeyRequest) (response PatchMyCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMyCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMyCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMyCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMyCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMyCustomerSecretKeyResponse")
	}
	return
}

// patchMyCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMyCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/MyCustomerSecretKeys/{myCustomerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMyCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMyCustomerSecretKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMyDevice Update a Device
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMyDevice.go.html to see an example of how to use PatchMyDevice API.
func (client IdentityDomainsClient) PatchMyDevice(ctx context.Context, request PatchMyDeviceRequest) (response PatchMyDeviceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMyDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMyDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMyDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMyDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMyDeviceResponse")
	}
	return
}

// patchMyDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMyDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/MyDevices/{myDeviceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMyDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMyDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMyOAuth2ClientCredential Update a user's own OAuth2 client credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMyOAuth2ClientCredential.go.html to see an example of how to use PatchMyOAuth2ClientCredential API.
func (client IdentityDomainsClient) PatchMyOAuth2ClientCredential(ctx context.Context, request PatchMyOAuth2ClientCredentialRequest) (response PatchMyOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMyOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMyOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMyOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMyOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMyOAuth2ClientCredentialResponse")
	}
	return
}

// patchMyOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMyOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/MyOAuth2ClientCredentials/{myOAuth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMyOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMyOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMySmtpCredential Update a user's own SMTP credential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchMySmtpCredential.go.html to see an example of how to use PatchMySmtpCredential API.
func (client IdentityDomainsClient) PatchMySmtpCredential(ctx context.Context, request PatchMySmtpCredentialRequest) (response PatchMySmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchMySmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMySmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMySmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMySmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMySmtpCredentialResponse")
	}
	return
}

// patchMySmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchMySmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/MySmtpCredentials/{mySmtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMySmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchMySmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchOAuth2ClientCredential Update a user's OAuth2 client credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchOAuth2ClientCredential.go.html to see an example of how to use PatchOAuth2ClientCredential API.
func (client IdentityDomainsClient) PatchOAuth2ClientCredential(ctx context.Context, request PatchOAuth2ClientCredentialRequest) (response PatchOAuth2ClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchOAuth2ClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchOAuth2ClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchOAuth2ClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchOAuth2ClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchOAuth2ClientCredentialResponse")
	}
	return
}

// patchOAuth2ClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchOAuth2ClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/OAuth2ClientCredentials/{oAuth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchOAuth2ClientCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchOAuth2ClientCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchPasswordPolicy Update a password policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchPasswordPolicy.go.html to see an example of how to use PatchPasswordPolicy API.
func (client IdentityDomainsClient) PatchPasswordPolicy(ctx context.Context, request PatchPasswordPolicyRequest) (response PatchPasswordPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchPasswordPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchPasswordPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchPasswordPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchPasswordPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchPasswordPolicyResponse")
	}
	return
}

// patchPasswordPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchPasswordPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/PasswordPolicies/{passwordPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchPasswordPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchPasswordPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchSecurityQuestion Update a security question.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchSecurityQuestion.go.html to see an example of how to use PatchSecurityQuestion API.
func (client IdentityDomainsClient) PatchSecurityQuestion(ctx context.Context, request PatchSecurityQuestionRequest) (response PatchSecurityQuestionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchSecurityQuestion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchSecurityQuestionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchSecurityQuestionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchSecurityQuestionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchSecurityQuestionResponse")
	}
	return
}

// patchSecurityQuestion implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchSecurityQuestion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/SecurityQuestions/{securityQuestionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchSecurityQuestionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchSecurityQuestion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchSecurityQuestionSetting Update a security question setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchSecurityQuestionSetting.go.html to see an example of how to use PatchSecurityQuestionSetting API.
func (client IdentityDomainsClient) PatchSecurityQuestionSetting(ctx context.Context, request PatchSecurityQuestionSettingRequest) (response PatchSecurityQuestionSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchSecurityQuestionSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchSecurityQuestionSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchSecurityQuestionSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchSecurityQuestionSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchSecurityQuestionSettingResponse")
	}
	return
}

// patchSecurityQuestionSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchSecurityQuestionSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/SecurityQuestionSettings/{securityQuestionSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchSecurityQuestionSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchSecurityQuestionSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchSmtpCredential Update a user's SMTP credentials.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchSmtpCredential.go.html to see an example of how to use PatchSmtpCredential API.
func (client IdentityDomainsClient) PatchSmtpCredential(ctx context.Context, request PatchSmtpCredentialRequest) (response PatchSmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchSmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchSmtpCredentialResponse")
	}
	return
}

// patchSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/SmtpCredentials/{smtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchSmtpCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchSmtpCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchUser Update a user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchUser.go.html to see an example of how to use PatchUser API.
func (client IdentityDomainsClient) PatchUser(ctx context.Context, request PatchUserRequest) (response PatchUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchUserResponse")
	}
	return
}

// patchUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/Users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchUserAttributesSetting Update User Schema Attribute Settings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PatchUserAttributesSetting.go.html to see an example of how to use PatchUserAttributesSetting API.
func (client IdentityDomainsClient) PatchUserAttributesSetting(ctx context.Context, request PatchUserAttributesSettingRequest) (response PatchUserAttributesSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchUserAttributesSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchUserAttributesSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchUserAttributesSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchUserAttributesSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchUserAttributesSettingResponse")
	}
	return
}

// patchUserAttributesSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) patchUserAttributesSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/admin/v1/UserAttributesSettings/{userAttributesSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchUserAttributesSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PatchUserAttributesSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutAccountRecoverySetting Replace an account recovery setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutAccountRecoverySetting.go.html to see an example of how to use PutAccountRecoverySetting API.
func (client IdentityDomainsClient) PutAccountRecoverySetting(ctx context.Context, request PutAccountRecoverySettingRequest) (response PutAccountRecoverySettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putAccountRecoverySetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutAccountRecoverySettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutAccountRecoverySettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutAccountRecoverySettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutAccountRecoverySettingResponse")
	}
	return
}

// putAccountRecoverySetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putAccountRecoverySetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/AccountRecoverySettings/{accountRecoverySettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutAccountRecoverySettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutAccountRecoverySetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutApp Replace an App
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutApp.go.html to see an example of how to use PutApp API.
func (client IdentityDomainsClient) PutApp(ctx context.Context, request PutAppRequest) (response PutAppResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putApp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutAppResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutAppResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutAppResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutAppResponse")
	}
	return
}

// putApp implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putApp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/Apps/{appId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutAppResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutApp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutAppStatusChanger Activate/Deactivate an App
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutAppStatusChanger.go.html to see an example of how to use PutAppStatusChanger API.
func (client IdentityDomainsClient) PutAppStatusChanger(ctx context.Context, request PutAppStatusChangerRequest) (response PutAppStatusChangerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putAppStatusChanger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutAppStatusChangerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutAppStatusChangerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutAppStatusChangerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutAppStatusChangerResponse")
	}
	return
}

// putAppStatusChanger implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putAppStatusChanger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/AppStatusChanger/{appStatusChangerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutAppStatusChangerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutAppStatusChanger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutAuthenticationFactorSetting Replace Authentication Factor Settings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutAuthenticationFactorSetting.go.html to see an example of how to use PutAuthenticationFactorSetting API.
func (client IdentityDomainsClient) PutAuthenticationFactorSetting(ctx context.Context, request PutAuthenticationFactorSettingRequest) (response PutAuthenticationFactorSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putAuthenticationFactorSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutAuthenticationFactorSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutAuthenticationFactorSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutAuthenticationFactorSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutAuthenticationFactorSettingResponse")
	}
	return
}

// putAuthenticationFactorSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putAuthenticationFactorSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/AuthenticationFactorSettings/{authenticationFactorSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutAuthenticationFactorSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutAuthenticationFactorSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutDynamicResourceGroup Replace a Dynamic Resource Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutDynamicResourceGroup.go.html to see an example of how to use PutDynamicResourceGroup API.
func (client IdentityDomainsClient) PutDynamicResourceGroup(ctx context.Context, request PutDynamicResourceGroupRequest) (response PutDynamicResourceGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putDynamicResourceGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutDynamicResourceGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutDynamicResourceGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutDynamicResourceGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutDynamicResourceGroupResponse")
	}
	return
}

// putDynamicResourceGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putDynamicResourceGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/DynamicResourceGroups/{dynamicResourceGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutDynamicResourceGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutDynamicResourceGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutGroup Replace a group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutGroup.go.html to see an example of how to use PutGroup API.
func (client IdentityDomainsClient) PutGroup(ctx context.Context, request PutGroupRequest) (response PutGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutGroupResponse")
	}
	return
}

// putGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/Groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutIdentityProvider Replace an Identity Provider
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutIdentityProvider.go.html to see an example of how to use PutIdentityProvider API.
func (client IdentityDomainsClient) PutIdentityProvider(ctx context.Context, request PutIdentityProviderRequest) (response PutIdentityProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutIdentityProviderResponse")
	}
	return
}

// putIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/IdentityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutIdentityProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutIdentitySetting Replace an Identity setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutIdentitySetting.go.html to see an example of how to use PutIdentitySetting API.
func (client IdentityDomainsClient) PutIdentitySetting(ctx context.Context, request PutIdentitySettingRequest) (response PutIdentitySettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putIdentitySetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutIdentitySettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutIdentitySettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutIdentitySettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutIdentitySettingResponse")
	}
	return
}

// putIdentitySetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putIdentitySetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/IdentitySettings/{identitySettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutIdentitySettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutIdentitySetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutKmsiSetting Replace KmsiSettings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutKmsiSetting.go.html to see an example of how to use PutKmsiSetting API.
func (client IdentityDomainsClient) PutKmsiSetting(ctx context.Context, request PutKmsiSettingRequest) (response PutKmsiSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putKmsiSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutKmsiSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutKmsiSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutKmsiSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutKmsiSettingResponse")
	}
	return
}

// putKmsiSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putKmsiSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/KmsiSettings/{kmsiSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutKmsiSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutKmsiSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutMe Replace a user's own information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutMe.go.html to see an example of how to use PutMe API.
func (client IdentityDomainsClient) PutMe(ctx context.Context, request PutMeRequest) (response PutMeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putMe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutMeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutMeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutMeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutMeResponse")
	}
	return
}

// putMe implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putMe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/Me", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutMeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutMe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutMePasswordChanger Update a user's own password.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutMePasswordChanger.go.html to see an example of how to use PutMePasswordChanger API.
func (client IdentityDomainsClient) PutMePasswordChanger(ctx context.Context, request PutMePasswordChangerRequest) (response PutMePasswordChangerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putMePasswordChanger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutMePasswordChangerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutMePasswordChangerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutMePasswordChangerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutMePasswordChangerResponse")
	}
	return
}

// putMePasswordChanger implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putMePasswordChanger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/MePasswordChanger", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutMePasswordChangerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutMePasswordChanger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutPasswordPolicy Replace a password policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutPasswordPolicy.go.html to see an example of how to use PutPasswordPolicy API.
func (client IdentityDomainsClient) PutPasswordPolicy(ctx context.Context, request PutPasswordPolicyRequest) (response PutPasswordPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putPasswordPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutPasswordPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutPasswordPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutPasswordPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutPasswordPolicyResponse")
	}
	return
}

// putPasswordPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putPasswordPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/PasswordPolicies/{passwordPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutPasswordPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutPasswordPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutSecurityQuestionSetting Replace a security question setting.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutSecurityQuestionSetting.go.html to see an example of how to use PutSecurityQuestionSetting API.
func (client IdentityDomainsClient) PutSecurityQuestionSetting(ctx context.Context, request PutSecurityQuestionSettingRequest) (response PutSecurityQuestionSettingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putSecurityQuestionSetting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutSecurityQuestionSettingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutSecurityQuestionSettingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutSecurityQuestionSettingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutSecurityQuestionSettingResponse")
	}
	return
}

// putSecurityQuestionSetting implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putSecurityQuestionSetting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/SecurityQuestionSettings/{securityQuestionSettingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutSecurityQuestionSettingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutSecurityQuestionSetting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutUser Replace a user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutUser.go.html to see an example of how to use PutUser API.
func (client IdentityDomainsClient) PutUser(ctx context.Context, request PutUserRequest) (response PutUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutUserResponse")
	}
	return
}

// putUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/Users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutUserCapabilitiesChanger Change a user's capabilities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutUserCapabilitiesChanger.go.html to see an example of how to use PutUserCapabilitiesChanger API.
func (client IdentityDomainsClient) PutUserCapabilitiesChanger(ctx context.Context, request PutUserCapabilitiesChangerRequest) (response PutUserCapabilitiesChangerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putUserCapabilitiesChanger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutUserCapabilitiesChangerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutUserCapabilitiesChangerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutUserCapabilitiesChangerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutUserCapabilitiesChangerResponse")
	}
	return
}

// putUserCapabilitiesChanger implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putUserCapabilitiesChanger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/UserCapabilitiesChanger/{userCapabilitiesChangerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutUserCapabilitiesChangerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutUserCapabilitiesChanger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutUserPasswordChanger Change a user's password to a known value.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutUserPasswordChanger.go.html to see an example of how to use PutUserPasswordChanger API.
func (client IdentityDomainsClient) PutUserPasswordChanger(ctx context.Context, request PutUserPasswordChangerRequest) (response PutUserPasswordChangerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putUserPasswordChanger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutUserPasswordChangerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutUserPasswordChangerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutUserPasswordChangerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutUserPasswordChangerResponse")
	}
	return
}

// putUserPasswordChanger implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putUserPasswordChanger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/UserPasswordChanger/{userPasswordChangerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutUserPasswordChangerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutUserPasswordChanger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutUserPasswordResetter Reset a user's password to a randomly-generated value.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutUserPasswordResetter.go.html to see an example of how to use PutUserPasswordResetter API.
func (client IdentityDomainsClient) PutUserPasswordResetter(ctx context.Context, request PutUserPasswordResetterRequest) (response PutUserPasswordResetterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putUserPasswordResetter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutUserPasswordResetterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutUserPasswordResetterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutUserPasswordResetterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutUserPasswordResetterResponse")
	}
	return
}

// putUserPasswordResetter implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putUserPasswordResetter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/UserPasswordResetter/{userPasswordResetterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutUserPasswordResetterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutUserPasswordResetter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutUserStatusChanger Change a user's status.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutUserStatusChanger.go.html to see an example of how to use PutUserStatusChanger API.
func (client IdentityDomainsClient) PutUserStatusChanger(ctx context.Context, request PutUserStatusChangerRequest) (response PutUserStatusChangerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putUserStatusChanger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutUserStatusChangerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutUserStatusChangerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutUserStatusChangerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutUserStatusChangerResponse")
	}
	return
}

// putUserStatusChanger implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) putUserStatusChanger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/admin/v1/UserStatusChanger/{userStatusChangerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutUserStatusChangerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "PutUserStatusChanger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchAccountMgmtInfos Search Account Mgmt Info Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchAccountMgmtInfos.go.html to see an example of how to use SearchAccountMgmtInfos API.
func (client IdentityDomainsClient) SearchAccountMgmtInfos(ctx context.Context, request SearchAccountMgmtInfosRequest) (response SearchAccountMgmtInfosResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchAccountMgmtInfos, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchAccountMgmtInfosResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchAccountMgmtInfosResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchAccountMgmtInfosResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchAccountMgmtInfosResponse")
	}
	return
}

// searchAccountMgmtInfos implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchAccountMgmtInfos(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AccountMgmtInfos/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchAccountMgmtInfosResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchAccountMgmtInfos", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchApiKeys Search ApiKeys Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchApiKeys.go.html to see an example of how to use SearchApiKeys API.
func (client IdentityDomainsClient) SearchApiKeys(ctx context.Context, request SearchApiKeysRequest) (response SearchApiKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchApiKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchApiKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchApiKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchApiKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchApiKeysResponse")
	}
	return
}

// searchApiKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchApiKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/ApiKeys/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchApiKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchApiKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchAppRoles Search AppRoles Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchAppRoles.go.html to see an example of how to use SearchAppRoles API.
func (client IdentityDomainsClient) SearchAppRoles(ctx context.Context, request SearchAppRolesRequest) (response SearchAppRolesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchAppRoles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchAppRolesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchAppRolesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchAppRolesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchAppRolesResponse")
	}
	return
}

// searchAppRoles implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchAppRoles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AppRoles/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchAppRolesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchAppRoles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchApps Search Apps Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchApps.go.html to see an example of how to use SearchApps API.
func (client IdentityDomainsClient) SearchApps(ctx context.Context, request SearchAppsRequest) (response SearchAppsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchApps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchAppsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchAppsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchAppsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchAppsResponse")
	}
	return
}

// searchApps implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchApps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Apps/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchAppsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchApps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchAuthTokens Search for Auth tokens using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchAuthTokens.go.html to see an example of how to use SearchAuthTokens API.
func (client IdentityDomainsClient) SearchAuthTokens(ctx context.Context, request SearchAuthTokensRequest) (response SearchAuthTokensResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchAuthTokens, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchAuthTokensResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchAuthTokensResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchAuthTokensResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchAuthTokensResponse")
	}
	return
}

// searchAuthTokens implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchAuthTokens(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AuthTokens/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchAuthTokensResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchAuthTokens", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchAuthenticationFactorSettings Search Authentication Factor Settings Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchAuthenticationFactorSettings.go.html to see an example of how to use SearchAuthenticationFactorSettings API.
func (client IdentityDomainsClient) SearchAuthenticationFactorSettings(ctx context.Context, request SearchAuthenticationFactorSettingsRequest) (response SearchAuthenticationFactorSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchAuthenticationFactorSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchAuthenticationFactorSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchAuthenticationFactorSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchAuthenticationFactorSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchAuthenticationFactorSettingsResponse")
	}
	return
}

// searchAuthenticationFactorSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchAuthenticationFactorSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/AuthenticationFactorSettings/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchAuthenticationFactorSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchAuthenticationFactorSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchCustomerSecretKeys Search for customer secret keys using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchCustomerSecretKeys.go.html to see an example of how to use SearchCustomerSecretKeys API.
func (client IdentityDomainsClient) SearchCustomerSecretKeys(ctx context.Context, request SearchCustomerSecretKeysRequest) (response SearchCustomerSecretKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchCustomerSecretKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchCustomerSecretKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchCustomerSecretKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchCustomerSecretKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchCustomerSecretKeysResponse")
	}
	return
}

// searchCustomerSecretKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchCustomerSecretKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/CustomerSecretKeys/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchCustomerSecretKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchCustomerSecretKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchDynamicResourceGroups Search for Dynamic Resource Groups using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchDynamicResourceGroups.go.html to see an example of how to use SearchDynamicResourceGroups API.
func (client IdentityDomainsClient) SearchDynamicResourceGroups(ctx context.Context, request SearchDynamicResourceGroupsRequest) (response SearchDynamicResourceGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchDynamicResourceGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchDynamicResourceGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchDynamicResourceGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchDynamicResourceGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchDynamicResourceGroupsResponse")
	}
	return
}

// searchDynamicResourceGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchDynamicResourceGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/DynamicResourceGroups/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchDynamicResourceGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchDynamicResourceGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchGrants Search Grants Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchGrants.go.html to see an example of how to use SearchGrants API.
func (client IdentityDomainsClient) SearchGrants(ctx context.Context, request SearchGrantsRequest) (response SearchGrantsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchGrants, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchGrantsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchGrantsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchGrantsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchGrantsResponse")
	}
	return
}

// searchGrants implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchGrants(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Grants/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchGrantsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchGrants", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchGroups Search for groups using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchGroups.go.html to see an example of how to use SearchGroups API.
func (client IdentityDomainsClient) SearchGroups(ctx context.Context, request SearchGroupsRequest) (response SearchGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchGroupsResponse")
	}
	return
}

// searchGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Groups/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchIdentityProviders Search Identity Providers Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchIdentityProviders.go.html to see an example of how to use SearchIdentityProviders API.
func (client IdentityDomainsClient) SearchIdentityProviders(ctx context.Context, request SearchIdentityProvidersRequest) (response SearchIdentityProvidersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchIdentityProviders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchIdentityProvidersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchIdentityProvidersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchIdentityProvidersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchIdentityProvidersResponse")
	}
	return
}

// searchIdentityProviders implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchIdentityProviders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/IdentityProviders/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchIdentityProvidersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchIdentityProviders", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchIdentitySettings Search for Identity settings using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchIdentitySettings.go.html to see an example of how to use SearchIdentitySettings API.
func (client IdentityDomainsClient) SearchIdentitySettings(ctx context.Context, request SearchIdentitySettingsRequest) (response SearchIdentitySettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchIdentitySettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchIdentitySettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchIdentitySettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchIdentitySettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchIdentitySettingsResponse")
	}
	return
}

// searchIdentitySettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchIdentitySettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/IdentitySettings/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchIdentitySettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchIdentitySettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchKmsiSettings Search KmsiSettings Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchKmsiSettings.go.html to see an example of how to use SearchKmsiSettings API.
func (client IdentityDomainsClient) SearchKmsiSettings(ctx context.Context, request SearchKmsiSettingsRequest) (response SearchKmsiSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchKmsiSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchKmsiSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchKmsiSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchKmsiSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchKmsiSettingsResponse")
	}
	return
}

// searchKmsiSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchKmsiSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/KmsiSettings/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchKmsiSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchKmsiSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMyApps Search My Apps Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchMyApps.go.html to see an example of how to use SearchMyApps API.
func (client IdentityDomainsClient) SearchMyApps(ctx context.Context, request SearchMyAppsRequest) (response SearchMyAppsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMyApps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMyAppsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMyAppsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMyAppsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMyAppsResponse")
	}
	return
}

// searchMyApps implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchMyApps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyApps/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMyAppsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchMyApps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMyGroups Search for 'My Groups' using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchMyGroups.go.html to see an example of how to use SearchMyGroups API.
func (client IdentityDomainsClient) SearchMyGroups(ctx context.Context, request SearchMyGroupsRequest) (response SearchMyGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMyGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMyGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMyGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMyGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMyGroupsResponse")
	}
	return
}

// searchMyGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchMyGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyGroups/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMyGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchMyGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMyRequestableGroups Search My Requestable Groups Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchMyRequestableGroups.go.html to see an example of how to use SearchMyRequestableGroups API.
func (client IdentityDomainsClient) SearchMyRequestableGroups(ctx context.Context, request SearchMyRequestableGroupsRequest) (response SearchMyRequestableGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMyRequestableGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMyRequestableGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMyRequestableGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMyRequestableGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMyRequestableGroupsResponse")
	}
	return
}

// searchMyRequestableGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchMyRequestableGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyRequestableGroups/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMyRequestableGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchMyRequestableGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMyRequests Search My Requests Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchMyRequests.go.html to see an example of how to use SearchMyRequests API.
func (client IdentityDomainsClient) SearchMyRequests(ctx context.Context, request SearchMyRequestsRequest) (response SearchMyRequestsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMyRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMyRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMyRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMyRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMyRequestsResponse")
	}
	return
}

// searchMyRequests implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchMyRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/MyRequests/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMyRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchMyRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchOAuth2ClientCredentials Search for OAuth2 client credentials using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchOAuth2ClientCredentials.go.html to see an example of how to use SearchOAuth2ClientCredentials API.
func (client IdentityDomainsClient) SearchOAuth2ClientCredentials(ctx context.Context, request SearchOAuth2ClientCredentialsRequest) (response SearchOAuth2ClientCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchOAuth2ClientCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchOAuth2ClientCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchOAuth2ClientCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchOAuth2ClientCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchOAuth2ClientCredentialsResponse")
	}
	return
}

// searchOAuth2ClientCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchOAuth2ClientCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/OAuth2ClientCredentials/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchOAuth2ClientCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchOAuth2ClientCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchPasswordPolicies Search for password policies using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchPasswordPolicies.go.html to see an example of how to use SearchPasswordPolicies API.
func (client IdentityDomainsClient) SearchPasswordPolicies(ctx context.Context, request SearchPasswordPoliciesRequest) (response SearchPasswordPoliciesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchPasswordPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchPasswordPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchPasswordPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchPasswordPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchPasswordPoliciesResponse")
	}
	return
}

// searchPasswordPolicies implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchPasswordPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/PasswordPolicies/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchPasswordPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchPasswordPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchResourceTypeSchemaAttributes Search Resource Type Schema Attributes Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchResourceTypeSchemaAttributes.go.html to see an example of how to use SearchResourceTypeSchemaAttributes API.
func (client IdentityDomainsClient) SearchResourceTypeSchemaAttributes(ctx context.Context, request SearchResourceTypeSchemaAttributesRequest) (response SearchResourceTypeSchemaAttributesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchResourceTypeSchemaAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchResourceTypeSchemaAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchResourceTypeSchemaAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchResourceTypeSchemaAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchResourceTypeSchemaAttributesResponse")
	}
	return
}

// searchResourceTypeSchemaAttributes implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchResourceTypeSchemaAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/ResourceTypeSchemaAttributes/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchResourceTypeSchemaAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchResourceTypeSchemaAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSecurityQuestionSettings Search for security question settings using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchSecurityQuestionSettings.go.html to see an example of how to use SearchSecurityQuestionSettings API.
func (client IdentityDomainsClient) SearchSecurityQuestionSettings(ctx context.Context, request SearchSecurityQuestionSettingsRequest) (response SearchSecurityQuestionSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchSecurityQuestionSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSecurityQuestionSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSecurityQuestionSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSecurityQuestionSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSecurityQuestionSettingsResponse")
	}
	return
}

// searchSecurityQuestionSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchSecurityQuestionSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/SecurityQuestionSettings/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSecurityQuestionSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchSecurityQuestionSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSecurityQuestions Search for security questions using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchSecurityQuestions.go.html to see an example of how to use SearchSecurityQuestions API.
func (client IdentityDomainsClient) SearchSecurityQuestions(ctx context.Context, request SearchSecurityQuestionsRequest) (response SearchSecurityQuestionsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchSecurityQuestions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSecurityQuestionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSecurityQuestionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSecurityQuestionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSecurityQuestionsResponse")
	}
	return
}

// searchSecurityQuestions implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchSecurityQuestions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/SecurityQuestions/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSecurityQuestionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchSecurityQuestions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchSmtpCredentials Search for SMTP credentials using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchSmtpCredentials.go.html to see an example of how to use SearchSmtpCredentials API.
func (client IdentityDomainsClient) SearchSmtpCredentials(ctx context.Context, request SearchSmtpCredentialsRequest) (response SearchSmtpCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchSmtpCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchSmtpCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchSmtpCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchSmtpCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchSmtpCredentialsResponse")
	}
	return
}

// searchSmtpCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchSmtpCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/SmtpCredentials/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchSmtpCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchSmtpCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchUserAttributesSettings Search User Schema Attribute Settings Using POST
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchUserAttributesSettings.go.html to see an example of how to use SearchUserAttributesSettings API.
func (client IdentityDomainsClient) SearchUserAttributesSettings(ctx context.Context, request SearchUserAttributesSettingsRequest) (response SearchUserAttributesSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchUserAttributesSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchUserAttributesSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchUserAttributesSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchUserAttributesSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchUserAttributesSettingsResponse")
	}
	return
}

// searchUserAttributesSettings implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchUserAttributesSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/UserAttributesSettings/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchUserAttributesSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchUserAttributesSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchUserDbCredentials Search for a user's database (DB) credentials using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchUserDbCredentials.go.html to see an example of how to use SearchUserDbCredentials API.
func (client IdentityDomainsClient) SearchUserDbCredentials(ctx context.Context, request SearchUserDbCredentialsRequest) (response SearchUserDbCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchUserDbCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchUserDbCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchUserDbCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchUserDbCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchUserDbCredentialsResponse")
	}
	return
}

// searchUserDbCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchUserDbCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/UserDbCredentials/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchUserDbCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchUserDbCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchUsers Search for users using POST.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/SearchUsers.go.html to see an example of how to use SearchUsers API.
func (client IdentityDomainsClient) SearchUsers(ctx context.Context, request SearchUsersRequest) (response SearchUsersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchUsersResponse")
	}
	return
}

// searchUsers implements the OCIOperation interface (enables retrying operations)
func (client IdentityDomainsClient) searchUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/admin/v1/Users/.search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchUsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "IdentityDomains", "SearchUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
