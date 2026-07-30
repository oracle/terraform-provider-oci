// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
// <b>Oracle AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DbMulticloudAwsProviderClient a client for DbMulticloudAwsProvider
type DbMulticloudAwsProviderClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbMulticloudAwsProviderClientWithConfigurationProvider Creates a new default DbMulticloudAwsProvider client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbMulticloudAwsProviderClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbMulticloudAwsProviderClient, err error) {
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
	return newDbMulticloudAwsProviderClientFromBaseClient(baseClient, provider)
}

// NewDbMulticloudAwsProviderClientWithOboToken Creates a new default DbMulticloudAwsProvider client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDbMulticloudAwsProviderClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbMulticloudAwsProviderClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbMulticloudAwsProviderClientFromBaseClient(baseClient, configProvider)
}

func newDbMulticloudAwsProviderClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbMulticloudAwsProviderClient, err error) {
	// DbMulticloudAwsProvider service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbMulticloudAwsProvider"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DbMulticloudAwsProviderClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbMulticloudAwsProviderClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbMulticloudAwsProviderClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbMulticloudAwsProviderClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOracleDbAwsIdentityConnectorCompartment Moves the Oracle DB AWS Identity Connector resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
// A default retry strategy applies to this operation ChangeOracleDbAwsIdentityConnectorCompartment()
func (client DbMulticloudAwsProviderClient) ChangeOracleDbAwsIdentityConnectorCompartment(ctx context.Context, request ChangeOracleDbAwsIdentityConnectorCompartmentRequest) (response ChangeOracleDbAwsIdentityConnectorCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAwsIdentityConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAwsIdentityConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAwsIdentityConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAwsIdentityConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAwsIdentityConnectorCompartmentResponse")
	}
	return
}

// changeOracleDbAwsIdentityConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) changeOracleDbAwsIdentityConnectorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsIdentityConnector/{oracleDbAwsIdentityConnectorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAwsIdentityConnectorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ChangeOracleDbAwsIdentityConnectorCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/ChangeOracleDbAwsIdentityConnectorCompartment"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ChangeOracleDbAwsIdentityConnectorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOracleDbAwsKeyCompartment Moves the AWS Key resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
// A default retry strategy applies to this operation ChangeOracleDbAwsKeyCompartment()
func (client DbMulticloudAwsProviderClient) ChangeOracleDbAwsKeyCompartment(ctx context.Context, request ChangeOracleDbAwsKeyCompartmentRequest) (response ChangeOracleDbAwsKeyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAwsKeyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAwsKeyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAwsKeyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAwsKeyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAwsKeyCompartmentResponse")
	}
	return
}

// changeOracleDbAwsKeyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) changeOracleDbAwsKeyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKey/{oracleDbAwsKeyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAwsKeyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ChangeOracleDbAwsKeyCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/ChangeOracleDbAwsKeyCompartment"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ChangeOracleDbAwsKeyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOracleDbAwsKeyPoolCompartment Moves the AWS Key Pool resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
// A default retry strategy applies to this operation ChangeOracleDbAwsKeyPoolCompartment()
func (client DbMulticloudAwsProviderClient) ChangeOracleDbAwsKeyPoolCompartment(ctx context.Context, request ChangeOracleDbAwsKeyPoolCompartmentRequest) (response ChangeOracleDbAwsKeyPoolCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAwsKeyPoolCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAwsKeyPoolCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAwsKeyPoolCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAwsKeyPoolCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAwsKeyPoolCompartmentResponse")
	}
	return
}

// changeOracleDbAwsKeyPoolCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) changeOracleDbAwsKeyPoolCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKeyPool/{oracleDbAwsKeyPoolId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAwsKeyPoolCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ChangeOracleDbAwsKeyPoolCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/ChangeOracleDbAwsKeyPoolCompartment"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ChangeOracleDbAwsKeyPoolCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAwsIdentityConnector Creates Oracle DB AWS Identity Connector resource.
// A default retry strategy applies to this operation CreateOracleDbAwsIdentityConnector()
func (client DbMulticloudAwsProviderClient) CreateOracleDbAwsIdentityConnector(ctx context.Context, request CreateOracleDbAwsIdentityConnectorRequest) (response CreateOracleDbAwsIdentityConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAwsIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAwsIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAwsIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAwsIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAwsIdentityConnectorResponse")
	}
	return
}

// createOracleDbAwsIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) createOracleDbAwsIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsIdentityConnector", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAwsIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "CreateOracleDbAwsIdentityConnector")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/CreateOracleDbAwsIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "CreateOracleDbAwsIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAwsKey Create DB AWS Key resource.
// A default retry strategy applies to this operation CreateOracleDbAwsKey()
func (client DbMulticloudAwsProviderClient) CreateOracleDbAwsKey(ctx context.Context, request CreateOracleDbAwsKeyRequest) (response CreateOracleDbAwsKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAwsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAwsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAwsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAwsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAwsKeyResponse")
	}
	return
}

// createOracleDbAwsKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) createOracleDbAwsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAwsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "CreateOracleDbAwsKey")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/CreateOracleDbAwsKey"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "CreateOracleDbAwsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAwsKeyPool Create DB AWS Key Pool resource.
// A default retry strategy applies to this operation CreateOracleDbAwsKeyPool()
func (client DbMulticloudAwsProviderClient) CreateOracleDbAwsKeyPool(ctx context.Context, request CreateOracleDbAwsKeyPoolRequest) (response CreateOracleDbAwsKeyPoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAwsKeyPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAwsKeyPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAwsKeyPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAwsKeyPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAwsKeyPoolResponse")
	}
	return
}

// createOracleDbAwsKeyPool implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) createOracleDbAwsKeyPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKeyPool", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAwsKeyPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "CreateOracleDbAwsKeyPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/CreateOracleDbAwsKeyPool"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "CreateOracleDbAwsKeyPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAwsIdentityConnector Deletes a Oracle DB AWS Identity Connector resource.
// A default retry strategy applies to this operation DeleteOracleDbAwsIdentityConnector()
func (client DbMulticloudAwsProviderClient) DeleteOracleDbAwsIdentityConnector(ctx context.Context, request DeleteOracleDbAwsIdentityConnectorRequest) (response DeleteOracleDbAwsIdentityConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAwsIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAwsIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAwsIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAwsIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAwsIdentityConnectorResponse")
	}
	return
}

// deleteOracleDbAwsIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) deleteOracleDbAwsIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAwsIdentityConnector/{oracleDbAwsIdentityConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAwsIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "DeleteOracleDbAwsIdentityConnector")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/DeleteOracleDbAwsIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "DeleteOracleDbAwsIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAwsKey Delete AWS Key resource.
// A default retry strategy applies to this operation DeleteOracleDbAwsKey()
func (client DbMulticloudAwsProviderClient) DeleteOracleDbAwsKey(ctx context.Context, request DeleteOracleDbAwsKeyRequest) (response DeleteOracleDbAwsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAwsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAwsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAwsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAwsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAwsKeyResponse")
	}
	return
}

// deleteOracleDbAwsKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) deleteOracleDbAwsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAwsKey/{oracleDbAwsKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAwsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "DeleteOracleDbAwsKey")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/DeleteOracleDbAwsKey"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "DeleteOracleDbAwsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAwsKeyPool Delete AWS Key Pool resource.
// A default retry strategy applies to this operation DeleteOracleDbAwsKeyPool()
func (client DbMulticloudAwsProviderClient) DeleteOracleDbAwsKeyPool(ctx context.Context, request DeleteOracleDbAwsKeyPoolRequest) (response DeleteOracleDbAwsKeyPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAwsKeyPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAwsKeyPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAwsKeyPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAwsKeyPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAwsKeyPoolResponse")
	}
	return
}

// deleteOracleDbAwsKeyPool implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) deleteOracleDbAwsKeyPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAwsKeyPool/{oracleDbAwsKeyPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAwsKeyPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "DeleteOracleDbAwsKeyPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/DeleteOracleDbAwsKeyPool"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "DeleteOracleDbAwsKeyPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAwsIdentityConnector Retrieves detailed information about a Oracle DB AWS Identity Connector resource by specifying its unique resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation GetOracleDbAwsIdentityConnector()
func (client DbMulticloudAwsProviderClient) GetOracleDbAwsIdentityConnector(ctx context.Context, request GetOracleDbAwsIdentityConnectorRequest) (response GetOracleDbAwsIdentityConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAwsIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAwsIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAwsIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAwsIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAwsIdentityConnectorResponse")
	}
	return
}

// getOracleDbAwsIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) getOracleDbAwsIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAwsIdentityConnector/{oracleDbAwsIdentityConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAwsIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "GetOracleDbAwsIdentityConnector")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/GetOracleDbAwsIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "GetOracleDbAwsIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAwsKey Retrieves detailed information about a Oracle AWS Key resource by specifying its unique resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation GetOracleDbAwsKey()
func (client DbMulticloudAwsProviderClient) GetOracleDbAwsKey(ctx context.Context, request GetOracleDbAwsKeyRequest) (response GetOracleDbAwsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAwsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAwsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAwsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAwsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAwsKeyResponse")
	}
	return
}

// getOracleDbAwsKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) getOracleDbAwsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAwsKey/{oracleDbAwsKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAwsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "GetOracleDbAwsKey")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/GetOracleDbAwsKey"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "GetOracleDbAwsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAwsKeyPool Retrieves detailed information about a Oracle AWS Key Pool resource by specifying its unique resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation GetOracleDbAwsKeyPool()
func (client DbMulticloudAwsProviderClient) GetOracleDbAwsKeyPool(ctx context.Context, request GetOracleDbAwsKeyPoolRequest) (response GetOracleDbAwsKeyPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAwsKeyPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAwsKeyPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAwsKeyPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAwsKeyPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAwsKeyPoolResponse")
	}
	return
}

// getOracleDbAwsKeyPool implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) getOracleDbAwsKeyPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAwsKeyPool/{oracleDbAwsKeyPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAwsKeyPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "GetOracleDbAwsKeyPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/GetOracleDbAwsKeyPool"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "GetOracleDbAwsKeyPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAwsIdentityConnectors Lists all Oracle DB AWS Identity Connectors based on the specified filters.
// A default retry strategy applies to this operation ListOracleDbAwsIdentityConnectors()
func (client DbMulticloudAwsProviderClient) ListOracleDbAwsIdentityConnectors(ctx context.Context, request ListOracleDbAwsIdentityConnectorsRequest) (response ListOracleDbAwsIdentityConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAwsIdentityConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAwsIdentityConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAwsIdentityConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAwsIdentityConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAwsIdentityConnectorsResponse")
	}
	return
}

// listOracleDbAwsIdentityConnectors implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) listOracleDbAwsIdentityConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAwsIdentityConnector", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAwsIdentityConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ListOracleDbAwsIdentityConnectors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/ListOracleDbAwsIdentityConnectors"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ListOracleDbAwsIdentityConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAwsKeyPools Lists all DB AWS Key Pools based on the specified filters.
// A default retry strategy applies to this operation ListOracleDbAwsKeyPools()
func (client DbMulticloudAwsProviderClient) ListOracleDbAwsKeyPools(ctx context.Context, request ListOracleDbAwsKeyPoolsRequest) (response ListOracleDbAwsKeyPoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAwsKeyPools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAwsKeyPoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAwsKeyPoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAwsKeyPoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAwsKeyPoolsResponse")
	}
	return
}

// listOracleDbAwsKeyPools implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) listOracleDbAwsKeyPools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAwsKeyPool", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAwsKeyPoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ListOracleDbAwsKeyPools")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/ListOracleDbAwsKeyPools"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ListOracleDbAwsKeyPools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAwsKeys Lists all DB AWS Keys based on the specified filters.
// A default retry strategy applies to this operation ListOracleDbAwsKeys()
func (client DbMulticloudAwsProviderClient) ListOracleDbAwsKeys(ctx context.Context, request ListOracleDbAwsKeysRequest) (response ListOracleDbAwsKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAwsKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAwsKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAwsKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAwsKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAwsKeysResponse")
	}
	return
}

// listOracleDbAwsKeys implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) listOracleDbAwsKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAwsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAwsKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ListOracleDbAwsKeys")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/ListOracleDbAwsKeys"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ListOracleDbAwsKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshOracleDbAwsIdentityConnector Refreshes the Oracle DB AWS Connector resource.
// A default retry strategy applies to this operation RefreshOracleDbAwsIdentityConnector()
func (client DbMulticloudAwsProviderClient) RefreshOracleDbAwsIdentityConnector(ctx context.Context, request RefreshOracleDbAwsIdentityConnectorRequest) (response RefreshOracleDbAwsIdentityConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshOracleDbAwsIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshOracleDbAwsIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshOracleDbAwsIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshOracleDbAwsIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshOracleDbAwsIdentityConnectorResponse")
	}
	return
}

// refreshOracleDbAwsIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) refreshOracleDbAwsIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsIdentityConnector/{oracleDbAwsIdentityConnectorId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshOracleDbAwsIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "RefreshOracleDbAwsIdentityConnector")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/RefreshOracleDbAwsIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "RefreshOracleDbAwsIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshOracleDbAwsKey Refreshes Oracle AWS Key resource.
// A default retry strategy applies to this operation RefreshOracleDbAwsKey()
func (client DbMulticloudAwsProviderClient) RefreshOracleDbAwsKey(ctx context.Context, request RefreshOracleDbAwsKeyRequest) (response RefreshOracleDbAwsKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshOracleDbAwsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshOracleDbAwsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshOracleDbAwsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshOracleDbAwsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshOracleDbAwsKeyResponse")
	}
	return
}

// refreshOracleDbAwsKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) refreshOracleDbAwsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKey/{oracleDbAwsKeyId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshOracleDbAwsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "RefreshOracleDbAwsKey")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/RefreshOracleDbAwsKey"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "RefreshOracleDbAwsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReplicateOracleDbAwsKey Replicate Oracle AWS Key resource to target region.
// A default retry strategy applies to this operation ReplicateOracleDbAwsKey()
func (client DbMulticloudAwsProviderClient) ReplicateOracleDbAwsKey(ctx context.Context, request ReplicateOracleDbAwsKeyRequest) (response ReplicateOracleDbAwsKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.replicateOracleDbAwsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReplicateOracleDbAwsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReplicateOracleDbAwsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReplicateOracleDbAwsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReplicateOracleDbAwsKeyResponse")
	}
	return
}

// replicateOracleDbAwsKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) replicateOracleDbAwsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKey/{oracleDbAwsKeyId}/actions/replicate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReplicateOracleDbAwsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ReplicateOracleDbAwsKey")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/ReplicateOracleDbAwsKey"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ReplicateOracleDbAwsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReplicateOracleDbAwsKeyPool Replicate Oracle AWS Key Pool resource to target region.
// A default retry strategy applies to this operation ReplicateOracleDbAwsKeyPool()
func (client DbMulticloudAwsProviderClient) ReplicateOracleDbAwsKeyPool(ctx context.Context, request ReplicateOracleDbAwsKeyPoolRequest) (response ReplicateOracleDbAwsKeyPoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.replicateOracleDbAwsKeyPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReplicateOracleDbAwsKeyPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReplicateOracleDbAwsKeyPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReplicateOracleDbAwsKeyPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReplicateOracleDbAwsKeyPoolResponse")
	}
	return
}

// replicateOracleDbAwsKeyPool implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) replicateOracleDbAwsKeyPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAwsKeyPool/{oracleDbAwsKeyPoolId}/actions/replicate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReplicateOracleDbAwsKeyPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "ReplicateOracleDbAwsKeyPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/ReplicateOracleDbAwsKeyPool"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "ReplicateOracleDbAwsKeyPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAwsIdentityConnector Modifies the existing Oracle DB AWS Identity Connector resource for a given OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation UpdateOracleDbAwsIdentityConnector()
func (client DbMulticloudAwsProviderClient) UpdateOracleDbAwsIdentityConnector(ctx context.Context, request UpdateOracleDbAwsIdentityConnectorRequest) (response UpdateOracleDbAwsIdentityConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAwsIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAwsIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAwsIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAwsIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAwsIdentityConnectorResponse")
	}
	return
}

// updateOracleDbAwsIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) updateOracleDbAwsIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAwsIdentityConnector/{oracleDbAwsIdentityConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAwsIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "UpdateOracleDbAwsIdentityConnector")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsIdentityConnector/UpdateOracleDbAwsIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "UpdateOracleDbAwsIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAwsKey Modifies the existing Oracle AWS Key Details for a given OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation UpdateOracleDbAwsKey()
func (client DbMulticloudAwsProviderClient) UpdateOracleDbAwsKey(ctx context.Context, request UpdateOracleDbAwsKeyRequest) (response UpdateOracleDbAwsKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAwsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAwsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAwsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAwsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAwsKeyResponse")
	}
	return
}

// updateOracleDbAwsKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) updateOracleDbAwsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAwsKey/{oracleDbAwsKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAwsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "UpdateOracleDbAwsKey")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKey/UpdateOracleDbAwsKey"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "UpdateOracleDbAwsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAwsKeyPool Modifies the existing Oracle AWS Key Pool Details for a given OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation UpdateOracleDbAwsKeyPool()
func (client DbMulticloudAwsProviderClient) UpdateOracleDbAwsKeyPool(ctx context.Context, request UpdateOracleDbAwsKeyPoolRequest) (response UpdateOracleDbAwsKeyPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAwsKeyPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAwsKeyPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAwsKeyPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAwsKeyPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAwsKeyPoolResponse")
	}
	return
}

// updateOracleDbAwsKeyPool implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudAwsProviderClient) updateOracleDbAwsKeyPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAwsKeyPool/{oracleDbAwsKeyPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAwsKeyPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbMulticloudAwsProvider", "UpdateOracleDbAwsKeyPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAwsKeyPool/UpdateOracleDbAwsKeyPool"
		err = common.PostProcessServiceError(err, "DbMulticloudAwsProvider", "UpdateOracleDbAwsKeyPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
