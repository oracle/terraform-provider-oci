// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//VaultsClient a client for Vaults
type VaultsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewVaultsClientWithConfigurationProvider Creates a new default Vaults client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewVaultsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client VaultsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newVaultsClientFromBaseClient(baseClient, configProvider)
}

// NewVaultsClientWithOboToken Creates a new default Vaults client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewVaultsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client VaultsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newVaultsClientFromBaseClient(baseClient, configProvider)
}

func newVaultsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client VaultsClient, err error) {
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
func (client VaultsClient) CancelSecretDeletion(ctx context.Context, request CancelSecretDeletionRequest) (response CancelSecretDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) cancelSecretDeletion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/cancelDeletion")
	if err != nil {
		return nil, err
	}

	var response CancelSecretDeletionResponse
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

// CancelSecretVersionDeletion Cancels the scheduled deletion of a secret version.
func (client VaultsClient) CancelSecretVersionDeletion(ctx context.Context, request CancelSecretVersionDeletionRequest) (response CancelSecretVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) cancelSecretVersionDeletion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/version/{secretVersionNumber}/actions/cancelDeletion")
	if err != nil {
		return nil, err
	}

	var response CancelSecretVersionDeletionResponse
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

// ChangeSecretCompartment Moves a secret into a different compartment within the same tenancy. For information about
// moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, if-match is checked against the ETag values of the secret.
func (client VaultsClient) ChangeSecretCompartment(ctx context.Context, request ChangeSecretCompartmentRequest) (response ChangeSecretCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) changeSecretCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeSecretCompartmentResponse
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

// CreateSecret Creates a new secret according to the details of the request.
// This operation is not supported by the Oracle Cloud Infrastructure Terraform Provider.
func (client VaultsClient) CreateSecret(ctx context.Context, request CreateSecretRequest) (response CreateSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) createSecret(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets")
	if err != nil {
		return nil, err
	}

	var response CreateSecretResponse
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

// GetSecret Gets information about the specified secret.
func (client VaultsClient) GetSecret(ctx context.Context, request GetSecretRequest) (response GetSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) getSecret(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets/{secretId}")
	if err != nil {
		return nil, err
	}

	var response GetSecretResponse
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

// GetSecretVersion Gets information about the specified version of a secret.
func (client VaultsClient) GetSecretVersion(ctx context.Context, request GetSecretVersionRequest) (response GetSecretVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) getSecretVersion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets/{secretId}/version/{secretVersionNumber}")
	if err != nil {
		return nil, err
	}

	var response GetSecretVersionResponse
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

// ListSecretVersions Lists all secret versions for the specified secret.
func (client VaultsClient) ListSecretVersions(ctx context.Context, request ListSecretVersionsRequest) (response ListSecretVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) listSecretVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets/{secretId}/versions")
	if err != nil {
		return nil, err
	}

	var response ListSecretVersionsResponse
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

// ListSecrets Lists all secrets in the specified vault and compartment.
func (client VaultsClient) ListSecrets(ctx context.Context, request ListSecretsRequest) (response ListSecretsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) listSecrets(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/secrets")
	if err != nil {
		return nil, err
	}

	var response ListSecretsResponse
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

// ScheduleSecretDeletion Schedules the deletion of the specified secret. This sets the lifecycle state of the secret
// to `PENDING_DELETION` and then deletes it after the specified retention period ends.
func (client VaultsClient) ScheduleSecretDeletion(ctx context.Context, request ScheduleSecretDeletionRequest) (response ScheduleSecretDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) scheduleSecretDeletion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/actions/scheduleDeletion")
	if err != nil {
		return nil, err
	}

	var response ScheduleSecretDeletionResponse
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

// ScheduleSecretVersionDeletion Schedules the deletion of the specified secret version. This deletes it after the specified retention period ends. You can only
// delete a secret version if the secret version rotation state is marked as `DEPRECATED`.
func (client VaultsClient) ScheduleSecretVersionDeletion(ctx context.Context, request ScheduleSecretVersionDeletionRequest) (response ScheduleSecretVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) scheduleSecretVersionDeletion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/secrets/{secretId}/version/{secretVersionNumber}/actions/scheduleDeletion")
	if err != nil {
		return nil, err
	}

	var response ScheduleSecretVersionDeletionResponse
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

// UpdateSecret Updates the properties of a secret. Specifically, you can update the version number of the secret to make
// that version number the current version. You can also update a secret's description, its free-form or defined tags, rules
// and the secret contents. Updating the secret content automatically creates a new secret version. You cannot, however, update the current secret version number and the secret contents and the rules at the
// same time. Furthermore, the secret must in an `ACTIVE` lifecycle state to be updated.
// This operation is not supported by the Oracle Cloud Infrastructure Terraform Provider.
func (client VaultsClient) UpdateSecret(ctx context.Context, request UpdateSecretRequest) (response UpdateSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client VaultsClient) updateSecret(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/secrets/{secretId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSecretResponse
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
