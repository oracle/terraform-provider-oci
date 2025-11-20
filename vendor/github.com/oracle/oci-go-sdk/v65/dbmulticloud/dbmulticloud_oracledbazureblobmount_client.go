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
// <br>
// <b>AWS</b>:<br>
// <b>Oracle AWS Connector Resource:</b>&nbsp;&nbsp;The Oracle AWS Connector Resource is used to install the AWS Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Google AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OracleDBAzureBlobMountClient a client for OracleDBAzureBlobMount
type OracleDBAzureBlobMountClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOracleDBAzureBlobMountClientWithConfigurationProvider Creates a new default OracleDBAzureBlobMount client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOracleDBAzureBlobMountClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OracleDBAzureBlobMountClient, err error) {
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
	return newOracleDBAzureBlobMountClientFromBaseClient(baseClient, provider)
}

// NewOracleDBAzureBlobMountClientWithOboToken Creates a new default OracleDBAzureBlobMount client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOracleDBAzureBlobMountClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OracleDBAzureBlobMountClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOracleDBAzureBlobMountClientFromBaseClient(baseClient, configProvider)
}

func newOracleDBAzureBlobMountClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OracleDBAzureBlobMountClient, err error) {
	// OracleDBAzureBlobMount service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OracleDBAzureBlobMount"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OracleDBAzureBlobMountClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OracleDBAzureBlobMountClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OracleDBAzureBlobMountClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OracleDBAzureBlobMountClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOracleDbAzureBlobMountCompartment Moves the Oracle DB Azure Blob Mount resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ChangeOracleDbAzureBlobMountCompartment.go.html to see an example of how to use ChangeOracleDbAzureBlobMountCompartment API.
// A default retry strategy applies to this operation ChangeOracleDbAzureBlobMountCompartment()
func (client OracleDBAzureBlobMountClient) ChangeOracleDbAzureBlobMountCompartment(ctx context.Context, request ChangeOracleDbAzureBlobMountCompartmentRequest) (response ChangeOracleDbAzureBlobMountCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAzureBlobMountCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAzureBlobMountCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAzureBlobMountCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAzureBlobMountCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAzureBlobMountCompartmentResponse")
	}
	return
}

// changeOracleDbAzureBlobMountCompartment implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobMountClient) changeOracleDbAzureBlobMountCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureBlobMount/{oracleDbAzureBlobMountId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAzureBlobMountCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobMount/ChangeOracleDbAzureBlobMountCompartment"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobMount", "ChangeOracleDbAzureBlobMountCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAzureBlobMount Creates Oracle DB Azure Blob Mount resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateOracleDbAzureBlobMount.go.html to see an example of how to use CreateOracleDbAzureBlobMount API.
// A default retry strategy applies to this operation CreateOracleDbAzureBlobMount()
func (client OracleDBAzureBlobMountClient) CreateOracleDbAzureBlobMount(ctx context.Context, request CreateOracleDbAzureBlobMountRequest) (response CreateOracleDbAzureBlobMountResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAzureBlobMount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAzureBlobMountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAzureBlobMountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAzureBlobMountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAzureBlobMountResponse")
	}
	return
}

// createOracleDbAzureBlobMount implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobMountClient) createOracleDbAzureBlobMount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureBlobMount", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAzureBlobMountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobMount/CreateOracleDbAzureBlobMount"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobMount", "CreateOracleDbAzureBlobMount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAzureBlobMount Unmounts Oracle DB Azure Blob Mount resource from an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D) and deletes Oracle DB Azure Blob Mount resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/DeleteOracleDbAzureBlobMount.go.html to see an example of how to use DeleteOracleDbAzureBlobMount API.
// A default retry strategy applies to this operation DeleteOracleDbAzureBlobMount()
func (client OracleDBAzureBlobMountClient) DeleteOracleDbAzureBlobMount(ctx context.Context, request DeleteOracleDbAzureBlobMountRequest) (response DeleteOracleDbAzureBlobMountResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAzureBlobMount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAzureBlobMountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAzureBlobMountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAzureBlobMountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAzureBlobMountResponse")
	}
	return
}

// deleteOracleDbAzureBlobMount implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobMountClient) deleteOracleDbAzureBlobMount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAzureBlobMount/{oracleDbAzureBlobMountId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAzureBlobMountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobMount/DeleteOracleDbAzureBlobMount"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobMount", "DeleteOracleDbAzureBlobMount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAzureBlobMount Retrieves the Oracle DB Azure Blob Mount resource for a specified resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureBlobMount.go.html to see an example of how to use GetOracleDbAzureBlobMount API.
// A default retry strategy applies to this operation GetOracleDbAzureBlobMount()
func (client OracleDBAzureBlobMountClient) GetOracleDbAzureBlobMount(ctx context.Context, request GetOracleDbAzureBlobMountRequest) (response GetOracleDbAzureBlobMountResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAzureBlobMount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAzureBlobMountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAzureBlobMountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAzureBlobMountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAzureBlobMountResponse")
	}
	return
}

// getOracleDbAzureBlobMount implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobMountClient) getOracleDbAzureBlobMount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureBlobMount/{oracleDbAzureBlobMountId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAzureBlobMountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobMount/GetOracleDbAzureBlobMount"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobMount", "GetOracleDbAzureBlobMount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAzureBlobMounts Lists all Oracle DB Azure Blob Mount resources based on the specified filters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureBlobMounts.go.html to see an example of how to use ListOracleDbAzureBlobMounts API.
// A default retry strategy applies to this operation ListOracleDbAzureBlobMounts()
func (client OracleDBAzureBlobMountClient) ListOracleDbAzureBlobMounts(ctx context.Context, request ListOracleDbAzureBlobMountsRequest) (response ListOracleDbAzureBlobMountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAzureBlobMounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAzureBlobMountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAzureBlobMountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAzureBlobMountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAzureBlobMountsResponse")
	}
	return
}

// listOracleDbAzureBlobMounts implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobMountClient) listOracleDbAzureBlobMounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureBlobMount", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAzureBlobMountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobMount/ListOracleDbAzureBlobMounts"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobMount", "ListOracleDbAzureBlobMounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAzureBlobMount Modifies the existing Oracle DB Azure Blob Mount resource for a given OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/UpdateOracleDbAzureBlobMount.go.html to see an example of how to use UpdateOracleDbAzureBlobMount API.
// A default retry strategy applies to this operation UpdateOracleDbAzureBlobMount()
func (client OracleDBAzureBlobMountClient) UpdateOracleDbAzureBlobMount(ctx context.Context, request UpdateOracleDbAzureBlobMountRequest) (response UpdateOracleDbAzureBlobMountResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAzureBlobMount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAzureBlobMountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAzureBlobMountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAzureBlobMountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAzureBlobMountResponse")
	}
	return
}

// updateOracleDbAzureBlobMount implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobMountClient) updateOracleDbAzureBlobMount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAzureBlobMount/{oracleDbAzureBlobMountId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAzureBlobMountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobMount/UpdateOracleDbAzureBlobMount"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobMount", "UpdateOracleDbAzureBlobMount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
