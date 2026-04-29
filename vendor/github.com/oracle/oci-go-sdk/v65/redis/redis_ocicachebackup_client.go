// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OciCacheBackupClient a client for OciCacheBackup
type OciCacheBackupClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOciCacheBackupClientWithConfigurationProvider Creates a new default OciCacheBackup client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOciCacheBackupClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OciCacheBackupClient, err error) {
	if enabled := common.CheckForEnabledServices("redis"); !enabled {
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
	return newOciCacheBackupClientFromBaseClient(baseClient, provider)
}

// NewOciCacheBackupClientWithOboToken Creates a new default OciCacheBackup client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOciCacheBackupClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OciCacheBackupClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOciCacheBackupClientFromBaseClient(baseClient, configProvider)
}

func newOciCacheBackupClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OciCacheBackupClient, err error) {
	// OciCacheBackup service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OciCacheBackup"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OciCacheBackupClient{BaseClient: baseClient}
	client.BasePath = "20220315"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OciCacheBackupClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("redis", "https://redis.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OciCacheBackupClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OciCacheBackupClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOciCacheBackupCompartment Moves an OCI Cache Backup resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ChangeOciCacheBackupCompartment.go.html to see an example of how to use ChangeOciCacheBackupCompartment API.
// A default retry strategy applies to this operation ChangeOciCacheBackupCompartment()
func (client OciCacheBackupClient) ChangeOciCacheBackupCompartment(ctx context.Context, request ChangeOciCacheBackupCompartmentRequest) (response ChangeOciCacheBackupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOciCacheBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOciCacheBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOciCacheBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOciCacheBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOciCacheBackupCompartmentResponse")
	}
	return
}

// changeOciCacheBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) changeOciCacheBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheBackups/{ociCacheBackupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOciCacheBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "ChangeOciCacheBackupCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackup/ChangeOciCacheBackupCompartment"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "ChangeOciCacheBackupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOciCacheBackup Creates a new OCI Cache Backup.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/CreateOciCacheBackup.go.html to see an example of how to use CreateOciCacheBackup API.
// A default retry strategy applies to this operation CreateOciCacheBackup()
func (client OciCacheBackupClient) CreateOciCacheBackup(ctx context.Context, request CreateOciCacheBackupRequest) (response CreateOciCacheBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOciCacheBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOciCacheBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOciCacheBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOciCacheBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOciCacheBackupResponse")
	}
	return
}

// createOciCacheBackup implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) createOciCacheBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOciCacheBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "CreateOciCacheBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackup/CreateOciCacheBackup"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "CreateOciCacheBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOciCacheBackup Deletes an OCI Cache Backup resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/DeleteOciCacheBackup.go.html to see an example of how to use DeleteOciCacheBackup API.
// A default retry strategy applies to this operation DeleteOciCacheBackup()
func (client OciCacheBackupClient) DeleteOciCacheBackup(ctx context.Context, request DeleteOciCacheBackupRequest) (response DeleteOciCacheBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOciCacheBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOciCacheBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOciCacheBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOciCacheBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOciCacheBackupResponse")
	}
	return
}

// deleteOciCacheBackup implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) deleteOciCacheBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/ociCacheBackups/{ociCacheBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOciCacheBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "DeleteOciCacheBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackup/DeleteOciCacheBackup"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "DeleteOciCacheBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportOciCacheBackupToObjectStorage Initiates an asynchronous export of the backup’s RDB file(s) to the specified Object Storage bucket. The service generates the object names. For sharded backups, one object is written per shard under the optional prefix.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ExportOciCacheBackupToObjectStorage.go.html to see an example of how to use ExportOciCacheBackupToObjectStorage API.
// A default retry strategy applies to this operation ExportOciCacheBackupToObjectStorage()
func (client OciCacheBackupClient) ExportOciCacheBackupToObjectStorage(ctx context.Context, request ExportOciCacheBackupToObjectStorageRequest) (response ExportOciCacheBackupToObjectStorageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportOciCacheBackupToObjectStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportOciCacheBackupToObjectStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportOciCacheBackupToObjectStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportOciCacheBackupToObjectStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportOciCacheBackupToObjectStorageResponse")
	}
	return
}

// exportOciCacheBackupToObjectStorage implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) exportOciCacheBackupToObjectStorage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheBackups/{ociCacheBackupId}/actions/exportToObjectStorage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportOciCacheBackupToObjectStorageResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "ExportOciCacheBackupToObjectStorage")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackup/ExportOciCacheBackupToObjectStorage"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "ExportOciCacheBackupToObjectStorage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOciCacheBackup Gets an OCI Cache Backup by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/GetOciCacheBackup.go.html to see an example of how to use GetOciCacheBackup API.
// A default retry strategy applies to this operation GetOciCacheBackup()
func (client OciCacheBackupClient) GetOciCacheBackup(ctx context.Context, request GetOciCacheBackupRequest) (response GetOciCacheBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOciCacheBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOciCacheBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOciCacheBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOciCacheBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOciCacheBackupResponse")
	}
	return
}

// getOciCacheBackup implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) getOciCacheBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheBackups/{ociCacheBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOciCacheBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "GetOciCacheBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackup/GetOciCacheBackup"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "GetOciCacheBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOciCacheBackups Returns a list of OCI Cache Backups.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheBackups.go.html to see an example of how to use ListOciCacheBackups API.
// A default retry strategy applies to this operation ListOciCacheBackups()
func (client OciCacheBackupClient) ListOciCacheBackups(ctx context.Context, request ListOciCacheBackupsRequest) (response ListOciCacheBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOciCacheBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOciCacheBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOciCacheBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOciCacheBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOciCacheBackupsResponse")
	}
	return
}

// listOciCacheBackups implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) listOciCacheBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOciCacheBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "ListOciCacheBackups")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackupSummary/ListOciCacheBackups"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "ListOciCacheBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOciCacheBackup Updates the OCI Cache Backup
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/UpdateOciCacheBackup.go.html to see an example of how to use UpdateOciCacheBackup API.
// A default retry strategy applies to this operation UpdateOciCacheBackup()
func (client OciCacheBackupClient) UpdateOciCacheBackup(ctx context.Context, request UpdateOciCacheBackupRequest) (response UpdateOciCacheBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOciCacheBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOciCacheBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOciCacheBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOciCacheBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOciCacheBackupResponse")
	}
	return
}

// updateOciCacheBackup implements the OCIOperation interface (enables retrying operations)
func (client OciCacheBackupClient) updateOciCacheBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/ociCacheBackups/{ociCacheBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOciCacheBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ociCacheBackup", "UpdateOciCacheBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheBackup/UpdateOciCacheBackup"
		err = common.PostProcessServiceError(err, "OciCacheBackup", "UpdateOciCacheBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
