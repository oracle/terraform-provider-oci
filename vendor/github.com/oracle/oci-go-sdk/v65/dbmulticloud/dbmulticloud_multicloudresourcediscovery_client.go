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

// MultiCloudResourceDiscoveryClient a client for MultiCloudResourceDiscovery
type MultiCloudResourceDiscoveryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMultiCloudResourceDiscoveryClientWithConfigurationProvider Creates a new default MultiCloudResourceDiscovery client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMultiCloudResourceDiscoveryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MultiCloudResourceDiscoveryClient, err error) {
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
	return newMultiCloudResourceDiscoveryClientFromBaseClient(baseClient, provider)
}

// NewMultiCloudResourceDiscoveryClientWithOboToken Creates a new default MultiCloudResourceDiscovery client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMultiCloudResourceDiscoveryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MultiCloudResourceDiscoveryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMultiCloudResourceDiscoveryClientFromBaseClient(baseClient, configProvider)
}

func newMultiCloudResourceDiscoveryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MultiCloudResourceDiscoveryClient, err error) {
	// MultiCloudResourceDiscovery service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MultiCloudResourceDiscovery"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MultiCloudResourceDiscoveryClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MultiCloudResourceDiscoveryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MultiCloudResourceDiscoveryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MultiCloudResourceDiscoveryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeMultiCloudResourceDiscoveryCompartment Moves the Multicloud Resource Discovery resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ChangeMultiCloudResourceDiscoveryCompartment.go.html to see an example of how to use ChangeMultiCloudResourceDiscoveryCompartment API.
// A default retry strategy applies to this operation ChangeMultiCloudResourceDiscoveryCompartment()
func (client MultiCloudResourceDiscoveryClient) ChangeMultiCloudResourceDiscoveryCompartment(ctx context.Context, request ChangeMultiCloudResourceDiscoveryCompartmentRequest) (response ChangeMultiCloudResourceDiscoveryCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMultiCloudResourceDiscoveryCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMultiCloudResourceDiscoveryCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMultiCloudResourceDiscoveryCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMultiCloudResourceDiscoveryCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMultiCloudResourceDiscoveryCompartmentResponse")
	}
	return
}

// changeMultiCloudResourceDiscoveryCompartment implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudResourceDiscoveryClient) changeMultiCloudResourceDiscoveryCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/multiCloudResourceDiscovery/{multiCloudResourceDiscoveryId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMultiCloudResourceDiscoveryCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/MultiCloudResourceDiscovery/ChangeMultiCloudResourceDiscoveryCompartment"
		err = common.PostProcessServiceError(err, "MultiCloudResourceDiscovery", "ChangeMultiCloudResourceDiscoveryCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMultiCloudResourceDiscovery Discovers Multicloud Resource and their associated resources based on the information provided.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateMultiCloudResourceDiscovery.go.html to see an example of how to use CreateMultiCloudResourceDiscovery API.
// A default retry strategy applies to this operation CreateMultiCloudResourceDiscovery()
func (client MultiCloudResourceDiscoveryClient) CreateMultiCloudResourceDiscovery(ctx context.Context, request CreateMultiCloudResourceDiscoveryRequest) (response CreateMultiCloudResourceDiscoveryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMultiCloudResourceDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMultiCloudResourceDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMultiCloudResourceDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMultiCloudResourceDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMultiCloudResourceDiscoveryResponse")
	}
	return
}

// createMultiCloudResourceDiscovery implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudResourceDiscoveryClient) createMultiCloudResourceDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/multiCloudResourceDiscovery", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMultiCloudResourceDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/MultiCloudResourceDiscovery/CreateMultiCloudResourceDiscovery"
		err = common.PostProcessServiceError(err, "MultiCloudResourceDiscovery", "CreateMultiCloudResourceDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMultiCloudResourceDiscovery Deletes the Multicloud Resource Discovery resource and removes its associated metadata from Oracle Cloud Infrastructure.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/DeleteMultiCloudResourceDiscovery.go.html to see an example of how to use DeleteMultiCloudResourceDiscovery API.
// A default retry strategy applies to this operation DeleteMultiCloudResourceDiscovery()
func (client MultiCloudResourceDiscoveryClient) DeleteMultiCloudResourceDiscovery(ctx context.Context, request DeleteMultiCloudResourceDiscoveryRequest) (response DeleteMultiCloudResourceDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMultiCloudResourceDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMultiCloudResourceDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMultiCloudResourceDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMultiCloudResourceDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMultiCloudResourceDiscoveryResponse")
	}
	return
}

// deleteMultiCloudResourceDiscovery implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudResourceDiscoveryClient) deleteMultiCloudResourceDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/multiCloudResourceDiscovery/{multiCloudResourceDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMultiCloudResourceDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/MultiCloudResourceDiscovery/DeleteMultiCloudResourceDiscovery"
		err = common.PostProcessServiceError(err, "MultiCloudResourceDiscovery", "DeleteMultiCloudResourceDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMultiCloudResourceDiscovery Retrieves detailed information about a Multicloud discovered resource by specifying its unique resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetMultiCloudResourceDiscovery.go.html to see an example of how to use GetMultiCloudResourceDiscovery API.
// A default retry strategy applies to this operation GetMultiCloudResourceDiscovery()
func (client MultiCloudResourceDiscoveryClient) GetMultiCloudResourceDiscovery(ctx context.Context, request GetMultiCloudResourceDiscoveryRequest) (response GetMultiCloudResourceDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMultiCloudResourceDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMultiCloudResourceDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMultiCloudResourceDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMultiCloudResourceDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMultiCloudResourceDiscoveryResponse")
	}
	return
}

// getMultiCloudResourceDiscovery implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudResourceDiscoveryClient) getMultiCloudResourceDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/multiCloudResourceDiscovery/{multiCloudResourceDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMultiCloudResourceDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/MultiCloudResourceDiscovery/GetMultiCloudResourceDiscovery"
		err = common.PostProcessServiceError(err, "MultiCloudResourceDiscovery", "GetMultiCloudResourceDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMultiCloudResourceDiscoveries Lists all Multicloud Resource Discovery resources based on the specified filters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListMultiCloudResourceDiscoveries.go.html to see an example of how to use ListMultiCloudResourceDiscoveries API.
// A default retry strategy applies to this operation ListMultiCloudResourceDiscoveries()
func (client MultiCloudResourceDiscoveryClient) ListMultiCloudResourceDiscoveries(ctx context.Context, request ListMultiCloudResourceDiscoveriesRequest) (response ListMultiCloudResourceDiscoveriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMultiCloudResourceDiscoveries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMultiCloudResourceDiscoveriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMultiCloudResourceDiscoveriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMultiCloudResourceDiscoveriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMultiCloudResourceDiscoveriesResponse")
	}
	return
}

// listMultiCloudResourceDiscoveries implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudResourceDiscoveryClient) listMultiCloudResourceDiscoveries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/multiCloudResourceDiscovery", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMultiCloudResourceDiscoveriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/MultiCloudResourceDiscovery/ListMultiCloudResourceDiscoveries"
		err = common.PostProcessServiceError(err, "MultiCloudResourceDiscovery", "ListMultiCloudResourceDiscoveries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMultiCloudResourceDiscovery Modifies the properties of an Azure discovered resource identified by the specified resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/UpdateMultiCloudResourceDiscovery.go.html to see an example of how to use UpdateMultiCloudResourceDiscovery API.
// A default retry strategy applies to this operation UpdateMultiCloudResourceDiscovery()
func (client MultiCloudResourceDiscoveryClient) UpdateMultiCloudResourceDiscovery(ctx context.Context, request UpdateMultiCloudResourceDiscoveryRequest) (response UpdateMultiCloudResourceDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMultiCloudResourceDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMultiCloudResourceDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMultiCloudResourceDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMultiCloudResourceDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMultiCloudResourceDiscoveryResponse")
	}
	return
}

// updateMultiCloudResourceDiscovery implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudResourceDiscoveryClient) updateMultiCloudResourceDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/multiCloudResourceDiscovery/{multiCloudResourceDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMultiCloudResourceDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/MultiCloudResourceDiscovery/UpdateMultiCloudResourceDiscovery"
		err = common.PostProcessServiceError(err, "MultiCloudResourceDiscovery", "UpdateMultiCloudResourceDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
