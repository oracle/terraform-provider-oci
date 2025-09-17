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

// OracleDbAzureVaultAssociationClient a client for OracleDbAzureVaultAssociation
type OracleDbAzureVaultAssociationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOracleDbAzureVaultAssociationClientWithConfigurationProvider Creates a new default OracleDbAzureVaultAssociation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOracleDbAzureVaultAssociationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OracleDbAzureVaultAssociationClient, err error) {
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
	return newOracleDbAzureVaultAssociationClientFromBaseClient(baseClient, provider)
}

// NewOracleDbAzureVaultAssociationClientWithOboToken Creates a new default OracleDbAzureVaultAssociation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOracleDbAzureVaultAssociationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OracleDbAzureVaultAssociationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOracleDbAzureVaultAssociationClientFromBaseClient(baseClient, configProvider)
}

func newOracleDbAzureVaultAssociationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OracleDbAzureVaultAssociationClient, err error) {
	// OracleDbAzureVaultAssociation service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OracleDbAzureVaultAssociation"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OracleDbAzureVaultAssociationClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OracleDbAzureVaultAssociationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OracleDbAzureVaultAssociationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OracleDbAzureVaultAssociationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CascadingDeleteOracleDbAzureVaultAssociation Deletes the Oracle DB Azure Vault Association resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CascadingDeleteOracleDbAzureVaultAssociation.go.html to see an example of how to use CascadingDeleteOracleDbAzureVaultAssociation API.
// A default retry strategy applies to this operation CascadingDeleteOracleDbAzureVaultAssociation()
func (client OracleDbAzureVaultAssociationClient) CascadingDeleteOracleDbAzureVaultAssociation(ctx context.Context, request CascadingDeleteOracleDbAzureVaultAssociationRequest) (response CascadingDeleteOracleDbAzureVaultAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cascadingDeleteOracleDbAzureVaultAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CascadingDeleteOracleDbAzureVaultAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CascadingDeleteOracleDbAzureVaultAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CascadingDeleteOracleDbAzureVaultAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CascadingDeleteOracleDbAzureVaultAssociationResponse")
	}
	return
}

// cascadingDeleteOracleDbAzureVaultAssociation implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) cascadingDeleteOracleDbAzureVaultAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureVaultAssociation/{oracleDbAzureVaultAssociationId}/actions/cascadingDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CascadingDeleteOracleDbAzureVaultAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/CascadingDeleteOracleDbAzureVaultAssociation"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "CascadingDeleteOracleDbAzureVaultAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOracleDbAzureVaultAssociationCompartment Moves the Oracle DB Azure Vault Association resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ChangeOracleDbAzureVaultAssociationCompartment.go.html to see an example of how to use ChangeOracleDbAzureVaultAssociationCompartment API.
// A default retry strategy applies to this operation ChangeOracleDbAzureVaultAssociationCompartment()
func (client OracleDbAzureVaultAssociationClient) ChangeOracleDbAzureVaultAssociationCompartment(ctx context.Context, request ChangeOracleDbAzureVaultAssociationCompartmentRequest) (response ChangeOracleDbAzureVaultAssociationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAzureVaultAssociationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAzureVaultAssociationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAzureVaultAssociationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAzureVaultAssociationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAzureVaultAssociationCompartmentResponse")
	}
	return
}

// changeOracleDbAzureVaultAssociationCompartment implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) changeOracleDbAzureVaultAssociationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureVaultAssociation/{oracleDbAzureVaultAssociationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAzureVaultAssociationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/ChangeOracleDbAzureVaultAssociationCompartment"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "ChangeOracleDbAzureVaultAssociationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAzureVaultAssociation Creates Oracle DB Azure Vault Association resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateOracleDbAzureVaultAssociation.go.html to see an example of how to use CreateOracleDbAzureVaultAssociation API.
// A default retry strategy applies to this operation CreateOracleDbAzureVaultAssociation()
func (client OracleDbAzureVaultAssociationClient) CreateOracleDbAzureVaultAssociation(ctx context.Context, request CreateOracleDbAzureVaultAssociationRequest) (response CreateOracleDbAzureVaultAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAzureVaultAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAzureVaultAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAzureVaultAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAzureVaultAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAzureVaultAssociationResponse")
	}
	return
}

// createOracleDbAzureVaultAssociation implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) createOracleDbAzureVaultAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureVaultAssociation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAzureVaultAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/CreateOracleDbAzureVaultAssociation"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "CreateOracleDbAzureVaultAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAzureVaultAssociation Deletes the Oracle DB Azure Vault Association resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/DeleteOracleDbAzureVaultAssociation.go.html to see an example of how to use DeleteOracleDbAzureVaultAssociation API.
// A default retry strategy applies to this operation DeleteOracleDbAzureVaultAssociation()
func (client OracleDbAzureVaultAssociationClient) DeleteOracleDbAzureVaultAssociation(ctx context.Context, request DeleteOracleDbAzureVaultAssociationRequest) (response DeleteOracleDbAzureVaultAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAzureVaultAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAzureVaultAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAzureVaultAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAzureVaultAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAzureVaultAssociationResponse")
	}
	return
}

// deleteOracleDbAzureVaultAssociation implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) deleteOracleDbAzureVaultAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAzureVaultAssociation/{oracleDbAzureVaultAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAzureVaultAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/DeleteOracleDbAzureVaultAssociation"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "DeleteOracleDbAzureVaultAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAzureVaultAssociation Retrieves the details of an Oracle DB Azure Vault Association using the specified association resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureVaultAssociation.go.html to see an example of how to use GetOracleDbAzureVaultAssociation API.
// A default retry strategy applies to this operation GetOracleDbAzureVaultAssociation()
func (client OracleDbAzureVaultAssociationClient) GetOracleDbAzureVaultAssociation(ctx context.Context, request GetOracleDbAzureVaultAssociationRequest) (response GetOracleDbAzureVaultAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAzureVaultAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAzureVaultAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAzureVaultAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAzureVaultAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAzureVaultAssociationResponse")
	}
	return
}

// getOracleDbAzureVaultAssociation implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) getOracleDbAzureVaultAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureVaultAssociation/{oracleDbAzureVaultAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAzureVaultAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/GetOracleDbAzureVaultAssociation"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "GetOracleDbAzureVaultAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAzureVaultAssociations Lists all Oracle DB Azure Associations based on the specified filters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureVaultAssociations.go.html to see an example of how to use ListOracleDbAzureVaultAssociations API.
// A default retry strategy applies to this operation ListOracleDbAzureVaultAssociations()
func (client OracleDbAzureVaultAssociationClient) ListOracleDbAzureVaultAssociations(ctx context.Context, request ListOracleDbAzureVaultAssociationsRequest) (response ListOracleDbAzureVaultAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAzureVaultAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAzureVaultAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAzureVaultAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAzureVaultAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAzureVaultAssociationsResponse")
	}
	return
}

// listOracleDbAzureVaultAssociations implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) listOracleDbAzureVaultAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureVaultAssociation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAzureVaultAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/ListOracleDbAzureVaultAssociations"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "ListOracleDbAzureVaultAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAzureVaultAssociation Modifies the properties of an Oracle DB Azure Vault Association resource identified by the specified resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/UpdateOracleDbAzureVaultAssociation.go.html to see an example of how to use UpdateOracleDbAzureVaultAssociation API.
// A default retry strategy applies to this operation UpdateOracleDbAzureVaultAssociation()
func (client OracleDbAzureVaultAssociationClient) UpdateOracleDbAzureVaultAssociation(ctx context.Context, request UpdateOracleDbAzureVaultAssociationRequest) (response UpdateOracleDbAzureVaultAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAzureVaultAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAzureVaultAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAzureVaultAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAzureVaultAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAzureVaultAssociationResponse")
	}
	return
}

// updateOracleDbAzureVaultAssociation implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureVaultAssociationClient) updateOracleDbAzureVaultAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAzureVaultAssociation/{oracleDbAzureVaultAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAzureVaultAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureVaultAssociation/UpdateOracleDbAzureVaultAssociation"
		err = common.PostProcessServiceError(err, "OracleDbAzureVaultAssociation", "UpdateOracleDbAzureVaultAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
