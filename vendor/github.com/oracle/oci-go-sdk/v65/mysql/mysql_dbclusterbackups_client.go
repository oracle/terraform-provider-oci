// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DbClusterBackupsClient a client for DbClusterBackups
type DbClusterBackupsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbClusterBackupsClientWithConfigurationProvider Creates a new default DbClusterBackups client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbClusterBackupsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbClusterBackupsClient, err error) {
	if enabled := common.CheckForEnabledServices("mysql"); !enabled {
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
	return newDbClusterBackupsClientFromBaseClient(baseClient, provider)
}

// NewDbClusterBackupsClientWithOboToken Creates a new default DbClusterBackups client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDbClusterBackupsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbClusterBackupsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbClusterBackupsClientFromBaseClient(baseClient, configProvider)
}

func newDbClusterBackupsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbClusterBackupsClient, err error) {
	// DbClusterBackups service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbClusterBackups"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DbClusterBackupsClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbClusterBackupsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbClusterBackupsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbClusterBackupsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDbClusterBackupCompartment Moves a shared-storage DB cluster backup into a different compartment.
// When provided, If-Match is checked against ETag values of the backup.
// A default retry strategy applies to this operation ChangeDbClusterBackupCompartment()
func (client DbClusterBackupsClient) ChangeDbClusterBackupCompartment(ctx context.Context, request ChangeDbClusterBackupCompartmentRequest) (response ChangeDbClusterBackupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDbClusterBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDbClusterBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDbClusterBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDbClusterBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDbClusterBackupCompartmentResponse")
	}
	return
}

// changeDbClusterBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbClusterBackupsClient) changeDbClusterBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusterBackups/{dbClusterBackupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDbClusterBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbClusterBackups", "ChangeDbClusterBackupCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterBackup/ChangeDbClusterBackupCompartment"
		err = common.PostProcessServiceError(err, "DbClusterBackups", "ChangeDbClusterBackupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDbClusterBackup Creates a backup of a shared-storage DB cluster.
// A default retry strategy applies to this operation CreateDbClusterBackup()
func (client DbClusterBackupsClient) CreateDbClusterBackup(ctx context.Context, request CreateDbClusterBackupRequest) (response CreateDbClusterBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDbClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDbClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDbClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbClusterBackupResponse")
	}
	return
}

// createDbClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client DbClusterBackupsClient) createDbClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusterBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDbClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbClusterBackups", "CreateDbClusterBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterBackup/CreateDbClusterBackup"
		err = common.PostProcessServiceError(err, "DbClusterBackups", "CreateDbClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDbClusterBackup Deletes a shared-storage DB cluster backup.
// A default retry strategy applies to this operation DeleteDbClusterBackup()
func (client DbClusterBackupsClient) DeleteDbClusterBackup(ctx context.Context, request DeleteDbClusterBackupRequest) (response DeleteDbClusterBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDbClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDbClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbClusterBackupResponse")
	}
	return
}

// deleteDbClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client DbClusterBackupsClient) deleteDbClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbClusterBackups/{dbClusterBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDbClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbClusterBackups", "DeleteDbClusterBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterBackup/DeleteDbClusterBackup"
		err = common.PostProcessServiceError(err, "DbClusterBackups", "DeleteDbClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbClusterBackup Gets information about the specified shared-storage DB cluster backup.
// A default retry strategy applies to this operation GetDbClusterBackup()
func (client DbClusterBackupsClient) GetDbClusterBackup(ctx context.Context, request GetDbClusterBackupRequest) (response GetDbClusterBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbClusterBackupResponse")
	}
	return
}

// getDbClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client DbClusterBackupsClient) getDbClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbClusterBackups/{dbClusterBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbClusterBackups", "GetDbClusterBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterBackup/GetDbClusterBackup"
		err = common.PostProcessServiceError(err, "DbClusterBackups", "GetDbClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbClusterBackups Gets a list of shared-storage DB cluster backups in the specified compartment.
// A default retry strategy applies to this operation ListDbClusterBackups()
func (client DbClusterBackupsClient) ListDbClusterBackups(ctx context.Context, request ListDbClusterBackupsRequest) (response ListDbClusterBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbClusterBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbClusterBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbClusterBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbClusterBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbClusterBackupsResponse")
	}
	return
}

// listDbClusterBackups implements the OCIOperation interface (enables retrying operations)
func (client DbClusterBackupsClient) listDbClusterBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbClusterBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbClusterBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbClusterBackups", "ListDbClusterBackups")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterBackupCollection/ListDbClusterBackups"
		err = common.PostProcessServiceError(err, "DbClusterBackups", "ListDbClusterBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbClusterBackup Updates the details of a shared-storage DB cluster backup.
// A default retry strategy applies to this operation UpdateDbClusterBackup()
func (client DbClusterBackupsClient) UpdateDbClusterBackup(ctx context.Context, request UpdateDbClusterBackupRequest) (response UpdateDbClusterBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbClusterBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbClusterBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbClusterBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbClusterBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbClusterBackupResponse")
	}
	return
}

// updateDbClusterBackup implements the OCIOperation interface (enables retrying operations)
func (client DbClusterBackupsClient) updateDbClusterBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbClusterBackups/{dbClusterBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbClusterBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbClusterBackups", "UpdateDbClusterBackup")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterBackup/UpdateDbClusterBackup"
		err = common.PostProcessServiceError(err, "DbClusterBackups", "UpdateDbClusterBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
