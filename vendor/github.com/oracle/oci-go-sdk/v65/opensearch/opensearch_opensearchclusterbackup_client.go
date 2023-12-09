// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OpensearchClusterBackupClient a client for OpensearchClusterBackup
type OpensearchClusterBackupClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOpensearchClusterBackupClientWithConfigurationProvider Creates a new default OpensearchClusterBackup client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOpensearchClusterBackupClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OpensearchClusterBackupClient, err error) {
	if enabled := common.CheckForEnabledServices("opensearch"); !enabled {
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
	return newOpensearchClusterBackupClientFromBaseClient(baseClient, provider)
}

// NewOpensearchClusterBackupClientWithOboToken Creates a new default OpensearchClusterBackup client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOpensearchClusterBackupClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OpensearchClusterBackupClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOpensearchClusterBackupClientFromBaseClient(baseClient, configProvider)
}

func newOpensearchClusterBackupClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OpensearchClusterBackupClient, err error) {
	// OpensearchClusterBackup service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OpensearchClusterBackup"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OpensearchClusterBackupClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OpensearchClusterBackupClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("opensearch", "https://search-indexing.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OpensearchClusterBackupClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OpensearchClusterBackupClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DeleteOpensearchClusterBackup Deletes a OpensearchClusterBackup resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/DeleteOpensearchClusterBackup.go.html to see an example of how to use DeleteOpensearchClusterBackup API.
func (client OpensearchClusterBackupClient) DeleteOpensearchClusterBackup(ctx context.Context, request DeleteOpensearchClusterBackupRequest) (response DeleteOpensearchClusterBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOpensearchClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOpensearchClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOpensearchClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOpensearchClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOpensearchClusterBackupResponse")
	}
	return
}

// deleteOpensearchClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterBackupClient) deleteOpensearchClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/opensearchClusterBackups/{opensearchClusterBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOpensearchClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterBackup/DeleteOpensearchClusterBackup"
		err = common.PostProcessServiceError(err, "OpensearchClusterBackup", "DeleteOpensearchClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOpensearchClusterBackup Gets a OpensearchClusterBackup by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/GetOpensearchClusterBackup.go.html to see an example of how to use GetOpensearchClusterBackup API.
func (client OpensearchClusterBackupClient) GetOpensearchClusterBackup(ctx context.Context, request GetOpensearchClusterBackupRequest) (response GetOpensearchClusterBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOpensearchClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOpensearchClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOpensearchClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOpensearchClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOpensearchClusterBackupResponse")
	}
	return
}

// getOpensearchClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterBackupClient) getOpensearchClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opensearchClusterBackups/{opensearchClusterBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOpensearchClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterBackup/GetOpensearchClusterBackup"
		err = common.PostProcessServiceError(err, "OpensearchClusterBackup", "GetOpensearchClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOpensearchClusterBackups Returns a list of OpensearchClusterBackups.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/ListOpensearchClusterBackups.go.html to see an example of how to use ListOpensearchClusterBackups API.
func (client OpensearchClusterBackupClient) ListOpensearchClusterBackups(ctx context.Context, request ListOpensearchClusterBackupsRequest) (response ListOpensearchClusterBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOpensearchClusterBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOpensearchClusterBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOpensearchClusterBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOpensearchClusterBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOpensearchClusterBackupsResponse")
	}
	return
}

// listOpensearchClusterBackups implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterBackupClient) listOpensearchClusterBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opensearchClusterBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOpensearchClusterBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterBackupCollection/ListOpensearchClusterBackups"
		err = common.PostProcessServiceError(err, "OpensearchClusterBackup", "ListOpensearchClusterBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOpensearchClusterBackup Updates the OpensearchClusterBackup
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/UpdateOpensearchClusterBackup.go.html to see an example of how to use UpdateOpensearchClusterBackup API.
func (client OpensearchClusterBackupClient) UpdateOpensearchClusterBackup(ctx context.Context, request UpdateOpensearchClusterBackupRequest) (response UpdateOpensearchClusterBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOpensearchClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOpensearchClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOpensearchClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOpensearchClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOpensearchClusterBackupResponse")
	}
	return
}

// updateOpensearchClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterBackupClient) updateOpensearchClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/opensearchClusterBackups/{opensearchClusterBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOpensearchClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterBackup/UpdateOpensearchClusterBackup"
		err = common.PostProcessServiceError(err, "OpensearchClusterBackup", "UpdateOpensearchClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
