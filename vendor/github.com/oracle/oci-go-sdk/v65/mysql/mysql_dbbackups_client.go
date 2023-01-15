// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//DbBackupsClient a client for DbBackups
type DbBackupsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbBackupsClientWithConfigurationProvider Creates a new default DbBackups client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbBackupsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbBackupsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDbBackupsClientFromBaseClient(baseClient, provider)
}

// NewDbBackupsClientWithOboToken Creates a new default DbBackups client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDbBackupsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbBackupsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbBackupsClientFromBaseClient(baseClient, configProvider)
}

func newDbBackupsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbBackupsClient, err error) {
	// DbBackups service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbBackups"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DbBackupsClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbBackupsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbBackupsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbBackupsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBackupCompartment Moves a DB System Backup into a different compartment.
// When provided, If-Match is checked against ETag values of the Backup.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ChangeBackupCompartment.go.html to see an example of how to use ChangeBackupCompartment API.
func (client DbBackupsClient) ChangeBackupCompartment(ctx context.Context, request ChangeBackupCompartmentRequest) (response ChangeBackupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBackupCompartmentResponse")
	}
	return
}

// changeBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbBackupsClient) changeBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups/{backupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/Backup/ChangeBackupCompartment"
		err = common.PostProcessServiceError(err, "DbBackups", "ChangeBackupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBackup Create a backup of a DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/CreateBackup.go.html to see an example of how to use CreateBackup API.
// A default retry strategy applies to this operation CreateBackup()
func (client DbBackupsClient) CreateBackup(ctx context.Context, request CreateBackupRequest) (response CreateBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBackupResponse")
	}
	return
}

// createBackup implements the OCIOperation interface (enables retrying operations)
func (client DbBackupsClient) createBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/backups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/Backup/CreateBackup"
		err = common.PostProcessServiceError(err, "DbBackups", "CreateBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBackup Delete a Backup.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/DeleteBackup.go.html to see an example of how to use DeleteBackup API.
// A default retry strategy applies to this operation DeleteBackup()
func (client DbBackupsClient) DeleteBackup(ctx context.Context, request DeleteBackupRequest) (response DeleteBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBackupResponse")
	}
	return
}

// deleteBackup implements the OCIOperation interface (enables retrying operations)
func (client DbBackupsClient) deleteBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/backups/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/Backup/DeleteBackup"
		err = common.PostProcessServiceError(err, "DbBackups", "DeleteBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackup Get information about the specified Backup
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetBackup.go.html to see an example of how to use GetBackup API.
// A default retry strategy applies to this operation GetBackup()
func (client DbBackupsClient) GetBackup(ctx context.Context, request GetBackupRequest) (response GetBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackupResponse")
	}
	return
}

// getBackup implements the OCIOperation interface (enables retrying operations)
func (client DbBackupsClient) getBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/Backup/GetBackup"
		err = common.PostProcessServiceError(err, "DbBackups", "GetBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBackups Get a list of DB System backups.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListBackups.go.html to see an example of how to use ListBackups API.
// A default retry strategy applies to this operation ListBackups()
func (client DbBackupsClient) ListBackups(ctx context.Context, request ListBackupsRequest) (response ListBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBackupsResponse")
	}
	return
}

// listBackups implements the OCIOperation interface (enables retrying operations)
func (client DbBackupsClient) listBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/backups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BackupSummary/ListBackups"
		err = common.PostProcessServiceError(err, "DbBackups", "ListBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBackup Update the metadata of a Backup. Metadata such as the displayName or description
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/UpdateBackup.go.html to see an example of how to use UpdateBackup API.
// A default retry strategy applies to this operation UpdateBackup()
func (client DbBackupsClient) UpdateBackup(ctx context.Context, request UpdateBackupRequest) (response UpdateBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBackupResponse")
	}
	return
}

// updateBackup implements the OCIOperation interface (enables retrying operations)
func (client DbBackupsClient) updateBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/backups/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/Backup/UpdateBackup"
		err = common.PostProcessServiceError(err, "DbBackups", "UpdateBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
