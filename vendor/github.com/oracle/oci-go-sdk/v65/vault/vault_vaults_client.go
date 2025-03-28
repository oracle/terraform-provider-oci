// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// VaultsClient a client for Vaults
type VaultsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewVaultsClientWithConfigurationProvider Creates a new default Vaults client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewVaultsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client VaultsClient, err error) {
	if enabled := common.CheckForEnabledServices("vault"); !enabled {
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
	return newVaultsClientFromBaseClient(baseClient, provider)
}

// NewVaultsClientWithOboToken Creates a new default Vaults client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewVaultsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client VaultsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newVaultsClientFromBaseClient(baseClient, configProvider)
}

func newVaultsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client VaultsClient, err error) {
	// Vaults service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Vaults"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = VaultsClient{BaseClient: baseClient}
	client.BasePath = "20180608"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *VaultsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("vault", "https://vaults.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *VaultsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *VaultsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelSecretDeletion Cancels the pending deletion of the specified secret. Canceling
// a scheduled deletion restores the secret's lifecycle state to what
// it was before you scheduled the secret for deletion.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/CancelSecretDeletion.go.html to see an example of how to use CancelSecretDeletion API.
func (client VaultsClient) CancelSecretDeletion(ctx context.Context, request CancelSecretDeletionRequest) (response CancelSecretDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelSecretDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelSecretDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelSecretDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelSecretDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelSecretDeletionResponse")
	}
	return
}

// cancelSecretDeletion implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) cancelSecretDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelSecretDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/CancelSecretDeletion"
		err = common.PostProcessServiceError(err, "Vaults", "CancelSecretDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelSecretRotation Cancels the ongoing secret rotation. The cancellation is contingent on how
// far the rotation process has progressed. Upon cancelling a rotation, all
// future rotations are also disabled.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/CancelSecretRotation.go.html to see an example of how to use CancelSecretRotation API.
// A default retry strategy applies to this operation CancelSecretRotation()
func (client VaultsClient) CancelSecretRotation(ctx context.Context, request CancelSecretRotationRequest) (response CancelSecretRotationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelSecretRotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelSecretRotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelSecretRotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelSecretRotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelSecretRotationResponse")
	}
	return
}

// cancelSecretRotation implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) cancelSecretRotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/cancelRotation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelSecretRotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/CancelSecretRotation"
		err = common.PostProcessServiceError(err, "Vaults", "CancelSecretRotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelSecretVersionDeletion Cancels the scheduled deletion of a secret version.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/CancelSecretVersionDeletion.go.html to see an example of how to use CancelSecretVersionDeletion API.
func (client VaultsClient) CancelSecretVersionDeletion(ctx context.Context, request CancelSecretVersionDeletionRequest) (response CancelSecretVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelSecretVersionDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelSecretVersionDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelSecretVersionDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelSecretVersionDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelSecretVersionDeletionResponse")
	}
	return
}

// cancelSecretVersionDeletion implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) cancelSecretVersionDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/version/{secretVersionNumber}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelSecretVersionDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/SecretVersion/CancelSecretVersionDeletion"
		err = common.PostProcessServiceError(err, "Vaults", "CancelSecretVersionDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecretCompartment Moves a secret into a different compartment within the same tenancy. For information about
// moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, if-match is checked against the ETag values of the secret.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/ChangeSecretCompartment.go.html to see an example of how to use ChangeSecretCompartment API.
func (client VaultsClient) ChangeSecretCompartment(ctx context.Context, request ChangeSecretCompartmentRequest) (response ChangeSecretCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecretCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecretCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecretCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecretCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecretCompartmentResponse")
	}
	return
}

// changeSecretCompartment implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) changeSecretCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecretCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/ChangeSecretCompartment"
		err = common.PostProcessServiceError(err, "Vaults", "ChangeSecretCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecret Creates a new secret according to the details of the request.
// This operation is not supported by the Oracle Cloud Infrastructure Terraform Provider.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/CreateSecret.go.html to see an example of how to use CreateSecret API.
// A default retry strategy applies to this operation CreateSecret()
func (client VaultsClient) CreateSecret(ctx context.Context, request CreateSecretRequest) (response CreateSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecretResponse")
	}
	return
}

// createSecret implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) createSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/CreateSecret"
		err = common.PostProcessServiceError(err, "Vaults", "CreateSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecret Gets information about the specified secret.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/GetSecret.go.html to see an example of how to use GetSecret API.
// A default retry strategy applies to this operation GetSecret()
func (client VaultsClient) GetSecret(ctx context.Context, request GetSecretRequest) (response GetSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecretResponse")
	}
	return
}

// getSecret implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) getSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets/{secretId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/GetSecret"
		err = common.PostProcessServiceError(err, "Vaults", "GetSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecretVersion Gets information about the specified version of a secret.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/GetSecretVersion.go.html to see an example of how to use GetSecretVersion API.
// A default retry strategy applies to this operation GetSecretVersion()
func (client VaultsClient) GetSecretVersion(ctx context.Context, request GetSecretVersionRequest) (response GetSecretVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecretVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecretVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecretVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecretVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecretVersionResponse")
	}
	return
}

// getSecretVersion implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) getSecretVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets/{secretId}/version/{secretVersionNumber}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecretVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/SecretVersion/GetSecretVersion"
		err = common.PostProcessServiceError(err, "Vaults", "GetSecretVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecretVersions Lists all secret versions for the specified secret.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/ListSecretVersions.go.html to see an example of how to use ListSecretVersions API.
// A default retry strategy applies to this operation ListSecretVersions()
func (client VaultsClient) ListSecretVersions(ctx context.Context, request ListSecretVersionsRequest) (response ListSecretVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecretVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecretVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecretVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecretVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecretVersionsResponse")
	}
	return
}

// listSecretVersions implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) listSecretVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets/{secretId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecretVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/SecretVersionSummary/ListSecretVersions"
		err = common.PostProcessServiceError(err, "Vaults", "ListSecretVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecrets Lists all secrets in the specified vault and compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/ListSecrets.go.html to see an example of how to use ListSecrets API.
// A default retry strategy applies to this operation ListSecrets()
func (client VaultsClient) ListSecrets(ctx context.Context, request ListSecretsRequest) (response ListSecretsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecrets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecretsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecretsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecretsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecretsResponse")
	}
	return
}

// listSecrets implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) listSecrets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecretsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/SecretSummary/ListSecrets"
		err = common.PostProcessServiceError(err, "Vaults", "ListSecrets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateSecret API to force rotation of an existing secret in Vault and the specified target system; expects secret to have a valid Target System Details object
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/RotateSecret.go.html to see an example of how to use RotateSecret API.
// A default retry strategy applies to this operation RotateSecret()
func (client VaultsClient) RotateSecret(ctx context.Context, request RotateSecretRequest) (response RotateSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.rotateSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateSecretResponse")
	}
	return
}

// rotateSecret implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) rotateSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/rotate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/RotateSecret"
		err = common.PostProcessServiceError(err, "Vaults", "RotateSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleSecretDeletion Schedules the deletion of the specified secret. This sets the lifecycle state of the secret
// to `PENDING_DELETION` and then deletes it after the specified retention period ends.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/ScheduleSecretDeletion.go.html to see an example of how to use ScheduleSecretDeletion API.
func (client VaultsClient) ScheduleSecretDeletion(ctx context.Context, request ScheduleSecretDeletionRequest) (response ScheduleSecretDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.scheduleSecretDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleSecretDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleSecretDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleSecretDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleSecretDeletionResponse")
	}
	return
}

// scheduleSecretDeletion implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) scheduleSecretDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleSecretDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/ScheduleSecretDeletion"
		err = common.PostProcessServiceError(err, "Vaults", "ScheduleSecretDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleSecretVersionDeletion Schedules the deletion of the specified secret version. This deletes it after the specified retention period ends. You can only
// delete a secret version if the secret version rotation state is marked as `DEPRECATED`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/ScheduleSecretVersionDeletion.go.html to see an example of how to use ScheduleSecretVersionDeletion API.
func (client VaultsClient) ScheduleSecretVersionDeletion(ctx context.Context, request ScheduleSecretVersionDeletionRequest) (response ScheduleSecretVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.scheduleSecretVersionDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleSecretVersionDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleSecretVersionDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleSecretVersionDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleSecretVersionDeletionResponse")
	}
	return
}

// scheduleSecretVersionDeletion implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) scheduleSecretVersionDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/version/{secretVersionNumber}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleSecretVersionDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/SecretVersion/ScheduleSecretVersionDeletion"
		err = common.PostProcessServiceError(err, "Vaults", "ScheduleSecretVersionDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecret Updates the properties of a secret. Specifically, you can update the version number of the secret to make
// that version number the current version. You can also update a secret's description, its free-form or defined tags, rules
// and the secret contents. Updating the secret content automatically creates a new secret version. You cannot, however, update the current secret version number, secret contents, and secret rules at the
// same time. Furthermore, the secret must in an `ACTIVE` lifecycle state to be updated.
// This operation is not supported by the Oracle Cloud Infrastructure Terraform Provider.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/UpdateSecret.go.html to see an example of how to use UpdateSecret API.
func (client VaultsClient) UpdateSecret(ctx context.Context, request UpdateSecretRequest) (response UpdateSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecretResponse")
	}
	return
}

// updateSecret implements the OCIOperation interface (enables retrying operations)
func (client VaultsClient) updateSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/secrets/{secretId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/secretmgmt/20180608/Secret/UpdateSecret"
		err = common.PostProcessServiceError(err, "Vaults", "UpdateSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
