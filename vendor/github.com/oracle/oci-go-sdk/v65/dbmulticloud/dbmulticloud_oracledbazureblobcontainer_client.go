// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OracleDBAzureBlobContainerClient a client for OracleDBAzureBlobContainer
type OracleDBAzureBlobContainerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOracleDBAzureBlobContainerClientWithConfigurationProvider Creates a new default OracleDBAzureBlobContainer client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOracleDBAzureBlobContainerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OracleDBAzureBlobContainerClient, err error) {
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
	return newOracleDBAzureBlobContainerClientFromBaseClient(baseClient, provider)
}

// NewOracleDBAzureBlobContainerClientWithOboToken Creates a new default OracleDBAzureBlobContainer client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOracleDBAzureBlobContainerClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OracleDBAzureBlobContainerClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOracleDBAzureBlobContainerClientFromBaseClient(baseClient, configProvider)
}

func newOracleDBAzureBlobContainerClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OracleDBAzureBlobContainerClient, err error) {
	// OracleDBAzureBlobContainer service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OracleDBAzureBlobContainer"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OracleDBAzureBlobContainerClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OracleDBAzureBlobContainerClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OracleDBAzureBlobContainerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OracleDBAzureBlobContainerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOracleDbAzureBlobContainerCompartment Moves the Oracle DB Azure Blob Container resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ChangeOracleDbAzureBlobContainerCompartment.go.html to see an example of how to use ChangeOracleDbAzureBlobContainerCompartment API.
// A default retry strategy applies to this operation ChangeOracleDbAzureBlobContainerCompartment()
func (client OracleDBAzureBlobContainerClient) ChangeOracleDbAzureBlobContainerCompartment(ctx context.Context, request ChangeOracleDbAzureBlobContainerCompartmentRequest) (response ChangeOracleDbAzureBlobContainerCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbAzureBlobContainerCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbAzureBlobContainerCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbAzureBlobContainerCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbAzureBlobContainerCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbAzureBlobContainerCompartmentResponse")
	}
	return
}

// changeOracleDbAzureBlobContainerCompartment implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobContainerClient) changeOracleDbAzureBlobContainerCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureBlobContainer/{oracleDbAzureBlobContainerId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbAzureBlobContainerCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobContainer/ChangeOracleDbAzureBlobContainerCompartment"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobContainer", "ChangeOracleDbAzureBlobContainerCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbAzureBlobContainer Capture Azure Container details for mounting Azure Container on multiple OCI Database Resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateOracleDbAzureBlobContainer.go.html to see an example of how to use CreateOracleDbAzureBlobContainer API.
// A default retry strategy applies to this operation CreateOracleDbAzureBlobContainer()
func (client OracleDBAzureBlobContainerClient) CreateOracleDbAzureBlobContainer(ctx context.Context, request CreateOracleDbAzureBlobContainerRequest) (response CreateOracleDbAzureBlobContainerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbAzureBlobContainer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbAzureBlobContainerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbAzureBlobContainerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbAzureBlobContainerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbAzureBlobContainerResponse")
	}
	return
}

// createOracleDbAzureBlobContainer implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobContainerClient) createOracleDbAzureBlobContainer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbAzureBlobContainer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbAzureBlobContainerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobContainer/CreateOracleDbAzureBlobContainer"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobContainer", "CreateOracleDbAzureBlobContainer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbAzureBlobContainer Delete Oracle DB Azure Blob Container details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/DeleteOracleDbAzureBlobContainer.go.html to see an example of how to use DeleteOracleDbAzureBlobContainer API.
// A default retry strategy applies to this operation DeleteOracleDbAzureBlobContainer()
func (client OracleDBAzureBlobContainerClient) DeleteOracleDbAzureBlobContainer(ctx context.Context, request DeleteOracleDbAzureBlobContainerRequest) (response DeleteOracleDbAzureBlobContainerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbAzureBlobContainer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbAzureBlobContainerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbAzureBlobContainerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbAzureBlobContainerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbAzureBlobContainerResponse")
	}
	return
}

// deleteOracleDbAzureBlobContainer implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobContainerClient) deleteOracleDbAzureBlobContainer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbAzureBlobContainer/{oracleDbAzureBlobContainerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbAzureBlobContainerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobContainer/DeleteOracleDbAzureBlobContainer"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobContainer", "DeleteOracleDbAzureBlobContainer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbAzureBlobContainer Get Oracle DB Azure Blob Container Details form a particular Container Resource ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureBlobContainer.go.html to see an example of how to use GetOracleDbAzureBlobContainer API.
// A default retry strategy applies to this operation GetOracleDbAzureBlobContainer()
func (client OracleDBAzureBlobContainerClient) GetOracleDbAzureBlobContainer(ctx context.Context, request GetOracleDbAzureBlobContainerRequest) (response GetOracleDbAzureBlobContainerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAzureBlobContainer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAzureBlobContainerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAzureBlobContainerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAzureBlobContainerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAzureBlobContainerResponse")
	}
	return
}

// getOracleDbAzureBlobContainer implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobContainerClient) getOracleDbAzureBlobContainer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureBlobContainer/{oracleDbAzureBlobContainerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAzureBlobContainerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobContainer/GetOracleDbAzureBlobContainer"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobContainer", "GetOracleDbAzureBlobContainer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAzureBlobContainers Lists the all Oracle DB Azure Blob Container based on filter.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureBlobContainers.go.html to see an example of how to use ListOracleDbAzureBlobContainers API.
// A default retry strategy applies to this operation ListOracleDbAzureBlobContainers()
func (client OracleDBAzureBlobContainerClient) ListOracleDbAzureBlobContainers(ctx context.Context, request ListOracleDbAzureBlobContainersRequest) (response ListOracleDbAzureBlobContainersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAzureBlobContainers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAzureBlobContainersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAzureBlobContainersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAzureBlobContainersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAzureBlobContainersResponse")
	}
	return
}

// listOracleDbAzureBlobContainers implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobContainerClient) listOracleDbAzureBlobContainers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureBlobContainer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAzureBlobContainersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobContainer/ListOracleDbAzureBlobContainers"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobContainer", "ListOracleDbAzureBlobContainers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbAzureBlobContainer Modifies the existing Oracle DB Azure Blob Container for a given ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/UpdateOracleDbAzureBlobContainer.go.html to see an example of how to use UpdateOracleDbAzureBlobContainer API.
// A default retry strategy applies to this operation UpdateOracleDbAzureBlobContainer()
func (client OracleDBAzureBlobContainerClient) UpdateOracleDbAzureBlobContainer(ctx context.Context, request UpdateOracleDbAzureBlobContainerRequest) (response UpdateOracleDbAzureBlobContainerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbAzureBlobContainer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbAzureBlobContainerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbAzureBlobContainerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbAzureBlobContainerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbAzureBlobContainerResponse")
	}
	return
}

// updateOracleDbAzureBlobContainer implements the OCIOperation interface (enables retrying operations)
func (client OracleDBAzureBlobContainerClient) updateOracleDbAzureBlobContainer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbAzureBlobContainer/{oracleDbAzureBlobContainerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbAzureBlobContainerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbAzureBlobContainer/UpdateOracleDbAzureBlobContainer"
		err = common.PostProcessServiceError(err, "OracleDBAzureBlobContainer", "UpdateOracleDbAzureBlobContainer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
