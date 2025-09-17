// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data Plane Integration
//
// <b>Microsoft Azure:</b> <br>
// <b>Oracle Azure Connector Resource:</b>:&nbsp;&nbsp;The Oracle Azure Connector Resource is used to install the Azure Arc Server on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
//  The supported method to install the Azure Arc Server (Azure Identity) on the Exadata VM cluster:
// <ul>
//  <li>Using a Bearer Access Token</li>
// </ul>
// <b>Oracle Azure Blob Container Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Container Resource is used to capture the details of an Azure Blob Container.
// This resource can then be reused across multiple Exadata VM clusters in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D) to mount the Azure container.
// <b>Oracle Azure Blob Mount Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Mount Resource is used to mount an Azure Blob Container on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// It relies on both the Oracle Azure Connector and the Oracle Azure Blob Container Resource to perform the mount operation.
// <b>Discover Azure Vaults and Keys Resource:</b>&nbsp;&nbsp;The Discover Oracle Azure Vaults and Azure Keys Resource is used to discover Azure Vaults and the associated encryption keys available in your Azure project.
// <b>Oracle Azure Vault:</b>&nbsp;&nbsp;The Oracle Azure Vault Resource is used to manage Azure Vaults within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Oracle Azure Key:</b>&nbsp;&nbsp;Oracle Azure Key Resource is used to register and manage a Oracle Azure Key Key within Oracle Cloud Infrastructure (OCI) under an associated Azure Vault.
// <br>
// <b>Google Cloud:</b><br>
// <b>Oracle Google Cloud Connector Resource:</b>&nbsp;&nbsp;The Oracle Google Cloud Connector Resource is used to install the Google Cloud Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Discover Google Key Rings and Keys Resource:</b>&nbsp;&nbsp;The Discover Google Key Rings and Keys Resource is used to discover Google Cloud Key Rings and the associated encryption keys available in your Google Cloud project.
// <b>Google Key Rings Resource:</b>&nbsp;&nbsp;The Google Key Rings Resource is used to register and manage Google Cloud Key Rings within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Google Key Resource:</b>&nbsp;&nbsp;The Google Key Resource is used to register and manage a Google Cloud Key within Oracle Cloud Infrastructure (OCI) under an associated Google Key Ring.
//

package dbmulticloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OracleDbAzureVaultClient a client for OracleDbAzureVault
type OracleDbAzureVaultClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOracleDbAzureVaultClientWithConfigurationProvider Creates a new default OracleDbAzureVault client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOracleDbAzureVaultClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OracleDbAzureVaultClient, err error) {
	if enabled := common.CheckForEnabledServices("dbmulticloud"); !enabled {
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
	return newOracleDbAzureVaultClientFromBaseClient(baseClient, provider)
}

// NewOracleDbAzureVaultClientWithOboToken Creates a new default OracleDbAzureVault client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOracleDbAzureVaultClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OracleDbAzureVaultClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOracleDbAzureVaultClientFromBaseClient(baseClient, configProvider)
}

func newOracleDbAzureVaultClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OracleDbAzureVaultClient, err error) {
	// OracleDbAzureVault service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OracleDbAzureVault"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OracleDbAzureVaultClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OracleDbAzureVaultClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OracleDbAzureVaultClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OracleDbAzureVaultClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOracleDbAzureVaultCompartment Moves the DB Azure Vault resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ChangeOracleDbAzureVaultCompartment.go.html to see an example of how to use ChangeOracleDbAzureVaultCompartment API.
// A default retry strategy applies to this operation ChangeOracleDbAzureVaultCompartment()
func (client OracleDbAzureVaultClient) ChangeOracleDbAzureVaultCompartment(ctx context.Context, request ChangeOracleDbAzureVaultCompartmentRequest) (response ChangeOracleDbAzureVaultCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAzureVaultCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAzureVaultCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAzureVaultCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAzureVaultCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAzureVaultCompartmentResponse")
	}
	return
}

// changeOracleDbAzureVaultCompartment implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) changeOracleDbAzureVaultCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureVault/{oracleDbAzureVaultId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAzureVaultCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/ChangeOracleDbAzureVaultCompartment"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "ChangeOracleDbAzureVaultCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAzureVault Creates DB Azure Vault resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateOracleDbAzureVault.go.html to see an example of how to use CreateOracleDbAzureVault API.
// A default retry strategy applies to this operation CreateOracleDbAzureVault()
func (client OracleDbAzureVaultClient) CreateOracleDbAzureVault(ctx context.Context, request CreateOracleDbAzureVaultRequest) (response CreateOracleDbAzureVaultResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAzureVault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAzureVaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAzureVaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAzureVaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAzureVaultResponse")
	}
	return
}

// createOracleDbAzureVault implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) createOracleDbAzureVault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureVault", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAzureVaultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/CreateOracleDbAzureVault"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "CreateOracleDbAzureVault", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAzureVault Deletes the DB Azure Vault resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/DeleteOracleDbAzureVault.go.html to see an example of how to use DeleteOracleDbAzureVault API.
// A default retry strategy applies to this operation DeleteOracleDbAzureVault()
func (client OracleDbAzureVaultClient) DeleteOracleDbAzureVault(ctx context.Context, request DeleteOracleDbAzureVaultRequest) (response DeleteOracleDbAzureVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAzureVault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAzureVaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAzureVaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAzureVaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAzureVaultResponse")
	}
	return
}

// deleteOracleDbAzureVault implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) deleteOracleDbAzureVault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAzureVault/{oracleDbAzureVaultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAzureVaultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/DeleteOracleDbAzureVault"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "DeleteOracleDbAzureVault", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAzureVault Retrieves detailed information about an Oracle Database Azure Vault resource using its unique resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). This operation returns metadata and configuration details associated with the specified vault resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureVault.go.html to see an example of how to use GetOracleDbAzureVault API.
// A default retry strategy applies to this operation GetOracleDbAzureVault()
func (client OracleDbAzureVaultClient) GetOracleDbAzureVault(ctx context.Context, request GetOracleDbAzureVaultRequest) (response GetOracleDbAzureVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAzureVault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAzureVaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAzureVaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAzureVaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAzureVaultResponse")
	}
	return
}

// getOracleDbAzureVault implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) getOracleDbAzureVault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureVault/{oracleDbAzureVaultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAzureVaultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/GetOracleDbAzureVault"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "GetOracleDbAzureVault", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAzureVaults Lists all DB Azure Vault resources based on the specified filters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureVaults.go.html to see an example of how to use ListOracleDbAzureVaults API.
// A default retry strategy applies to this operation ListOracleDbAzureVaults()
func (client OracleDbAzureVaultClient) ListOracleDbAzureVaults(ctx context.Context, request ListOracleDbAzureVaultsRequest) (response ListOracleDbAzureVaultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAzureVaults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAzureVaultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAzureVaultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAzureVaultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAzureVaultsResponse")
	}
	return
}

// listOracleDbAzureVaults implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) listOracleDbAzureVaults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureVault", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAzureVaultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/ListOracleDbAzureVaults"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "ListOracleDbAzureVaults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshOracleDbAzureVault Refreshes the Oracle DB Azure Vault resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/RefreshOracleDbAzureVault.go.html to see an example of how to use RefreshOracleDbAzureVault API.
// A default retry strategy applies to this operation RefreshOracleDbAzureVault()
func (client OracleDbAzureVaultClient) RefreshOracleDbAzureVault(ctx context.Context, request RefreshOracleDbAzureVaultRequest) (response RefreshOracleDbAzureVaultResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshOracleDbAzureVault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshOracleDbAzureVaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshOracleDbAzureVaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshOracleDbAzureVaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshOracleDbAzureVaultResponse")
	}
	return
}

// refreshOracleDbAzureVault implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) refreshOracleDbAzureVault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureVault/{oracleDbAzureVaultId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshOracleDbAzureVaultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/RefreshOracleDbAzureVault"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "RefreshOracleDbAzureVault", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAzureVault Modifies the configuration details of an existing Oracle Database Azure Vault resource identified by its unique OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). This operation updates only the specified fields in the request body.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/UpdateOracleDbAzureVault.go.html to see an example of how to use UpdateOracleDbAzureVault API.
// A default retry strategy applies to this operation UpdateOracleDbAzureVault()
func (client OracleDbAzureVaultClient) UpdateOracleDbAzureVault(ctx context.Context, request UpdateOracleDbAzureVaultRequest) (response UpdateOracleDbAzureVaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAzureVault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAzureVaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAzureVaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAzureVaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAzureVaultResponse")
	}
	return
}

// updateOracleDbAzureVault implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultClient) updateOracleDbAzureVault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAzureVault/{oracleDbAzureVaultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAzureVaultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVault/UpdateOracleDbAzureVault"
		err = common.PostProcessServiceError(err, "OracleDbAzureVault", "UpdateOracleDbAzureVault", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
