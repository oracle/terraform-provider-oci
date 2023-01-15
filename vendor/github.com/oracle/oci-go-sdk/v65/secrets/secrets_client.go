// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Retrieval API
//
// Use the Secret Retrieval API to retrieve secrets and secret versions from vaults. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package secrets

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//SecretsClient a client for Secrets
type SecretsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSecretsClientWithConfigurationProvider Creates a new default Secrets client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSecretsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SecretsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newSecretsClientFromBaseClient(baseClient, provider)
}

// NewSecretsClientWithOboToken Creates a new default Secrets client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewSecretsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SecretsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSecretsClientFromBaseClient(baseClient, configProvider)
}

func newSecretsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SecretsClient, err error) {
	// Secrets service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Secrets"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SecretsClient{BaseClient: baseClient}
	client.BasePath = "20190301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SecretsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("secrets", "https://secrets.vaults.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SecretsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *SecretsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetSecretBundle Gets a secret bundle that matches either the specified `stage`, `secretVersionName`, or `versionNumber` parameter.
// If none of these parameters are provided, the bundle for the secret version marked as `CURRENT` will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/secrets/GetSecretBundle.go.html to see an example of how to use GetSecretBundle API.
// A default retry strategy applies to this operation GetSecretBundle()
func (client SecretsClient) GetSecretBundle(ctx context.Context, request GetSecretBundleRequest) (response GetSecretBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecretBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecretBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecretBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecretBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecretBundleResponse")
	}
	return
}

// getSecretBundle implements the OCIOperation interface (enables retrying operations)
func (client SecretsClient) getSecretBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secretbundles/{secretId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecretBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretretrieval/20190301/SecretBundle/GetSecretBundle"
		err = common.PostProcessServiceError(err, "Secrets", "GetSecretBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecretBundleByName Gets a secret bundle by secret name and vault ID, and secret version that matches either the specified `stage`, `secretVersionName`, or `versionNumber` parameter.
// If none of these parameters are provided, the bundle for the secret version marked as `CURRENT` is returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/secrets/GetSecretBundleByName.go.html to see an example of how to use GetSecretBundleByName API.
// A default retry strategy applies to this operation GetSecretBundleByName()
func (client SecretsClient) GetSecretBundleByName(ctx context.Context, request GetSecretBundleByNameRequest) (response GetSecretBundleByNameResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecretBundleByName, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecretBundleByNameResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecretBundleByNameResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecretBundleByNameResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecretBundleByNameResponse")
	}
	return
}

// getSecretBundleByName implements the OCIOperation interface (enables retrying operations)
func (client SecretsClient) getSecretBundleByName(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secretbundles/actions/getByName", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecretBundleByNameResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretretrieval/20190301/SecretBundle/GetSecretBundleByName"
		err = common.PostProcessServiceError(err, "Secrets", "GetSecretBundleByName", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecretBundleVersions Lists all secret bundle versions for the specified secret.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/secrets/ListSecretBundleVersions.go.html to see an example of how to use ListSecretBundleVersions API.
// A default retry strategy applies to this operation ListSecretBundleVersions()
func (client SecretsClient) ListSecretBundleVersions(ctx context.Context, request ListSecretBundleVersionsRequest) (response ListSecretBundleVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecretBundleVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecretBundleVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecretBundleVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecretBundleVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecretBundleVersionsResponse")
	}
	return
}

// listSecretBundleVersions implements the OCIOperation interface (enables retrying operations)
func (client SecretsClient) listSecretBundleVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secretbundles/{secretId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecretBundleVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretretrieval/20190301/SecretBundleVersionSummary/ListSecretBundleVersions"
		err = common.PostProcessServiceError(err, "Secrets", "ListSecretBundleVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
