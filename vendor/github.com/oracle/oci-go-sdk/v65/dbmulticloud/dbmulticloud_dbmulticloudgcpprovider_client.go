// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// <b>Microsoft Azure</b>:<br>
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
// <b>Google Cloud</b>:<br>
// 1. Oracle Google Cloud Connector Resource: This is for installing Google Identity in ExaCS VM Cluster.<br>
// 2. Discover Google Key-Rings and Keys Resource: This is for to discover Google Key-Rings.<br>
// 3. Google Key-Rings Resource: This is for to maintain Google Key-Rings in Oracle Cloud.<br>
// 4. Google Key Resource: This is for to maintain Google Key in Oracle Cloud for a Google Key-Ring.<br>
//

package dbmulticloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DbMulticloudGCPProviderClient a client for DbMulticloudGCPProvider
type DbMulticloudGCPProviderClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbMulticloudGCPProviderClientWithConfigurationProvider Creates a new default DbMulticloudGCPProvider client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbMulticloudGCPProviderClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbMulticloudGCPProviderClient, err error) {
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
	return newDbMulticloudGCPProviderClientFromBaseClient(baseClient, provider)
}

// NewDbMulticloudGCPProviderClientWithOboToken Creates a new default DbMulticloudGCPProvider client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDbMulticloudGCPProviderClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbMulticloudGCPProviderClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbMulticloudGCPProviderClientFromBaseClient(baseClient, configProvider)
}

func newDbMulticloudGCPProviderClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbMulticloudGCPProviderClient, err error) {
	// DbMulticloudGCPProvider service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbMulticloudGCPProvider"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DbMulticloudGCPProviderClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbMulticloudGCPProviderClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbMulticloudGCPProviderClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbMulticloudGCPProviderClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOracleDbGcpIdentityConnectorCompartment Moves the Oracle DB GCP Identity Connector resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
// A default retry strategy applies to this operation ChangeOracleDbGcpIdentityConnectorCompartment()
func (client DbMulticloudGCPProviderClient) ChangeOracleDbGcpIdentityConnectorCompartment(ctx context.Context, request ChangeOracleDbGcpIdentityConnectorCompartmentRequest) (response ChangeOracleDbGcpIdentityConnectorCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbGcpIdentityConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbGcpIdentityConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbGcpIdentityConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbGcpIdentityConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbGcpIdentityConnectorCompartmentResponse")
	}
	return
}

// changeOracleDbGcpIdentityConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) changeOracleDbGcpIdentityConnectorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbGcpIdentityConnector/{oracleDbGcpIdentityConnectorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbGcpIdentityConnectorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpIdentityConnector/ChangeOracleDbGcpIdentityConnectorCompartment"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "ChangeOracleDbGcpIdentityConnectorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOracleDbGcpKeyRingCompartment Moves the GCP Key Ring resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
// A default retry strategy applies to this operation ChangeOracleDbGcpKeyRingCompartment()
func (client DbMulticloudGCPProviderClient) ChangeOracleDbGcpKeyRingCompartment(ctx context.Context, request ChangeOracleDbGcpKeyRingCompartmentRequest) (response ChangeOracleDbGcpKeyRingCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOracleDbGcpKeyRingCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOracleDbGcpKeyRingCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOracleDbGcpKeyRingCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOracleDbGcpKeyRingCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOracleDbGcpKeyRingCompartmentResponse")
	}
	return
}

// changeOracleDbGcpKeyRingCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) changeOracleDbGcpKeyRingCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbGcpKeyRing/{oracleDbGcpKeyRingId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOracleDbGcpKeyRingCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/ChangeOracleDbGcpKeyRingCompartment"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "ChangeOracleDbGcpKeyRingCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbGcpIdentityConnector Creates Oracle DB GCP Identity Connector Resource.
// A default retry strategy applies to this operation CreateOracleDbGcpIdentityConnector()
func (client DbMulticloudGCPProviderClient) CreateOracleDbGcpIdentityConnector(ctx context.Context, request CreateOracleDbGcpIdentityConnectorRequest) (response CreateOracleDbGcpIdentityConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbGcpIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbGcpIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbGcpIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbGcpIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbGcpIdentityConnectorResponse")
	}
	return
}

// createOracleDbGcpIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) createOracleDbGcpIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbGcpIdentityConnector", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbGcpIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpIdentityConnector/CreateOracleDbGcpIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "CreateOracleDbGcpIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOracleDbGcpKeyRing Create DB GCP Key Rings based on the provided information, this will fetch Keys related to GCP Key Rings.
// A default retry strategy applies to this operation CreateOracleDbGcpKeyRing()
func (client DbMulticloudGCPProviderClient) CreateOracleDbGcpKeyRing(ctx context.Context, request CreateOracleDbGcpKeyRingRequest) (response CreateOracleDbGcpKeyRingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOracleDbGcpKeyRing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOracleDbGcpKeyRingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOracleDbGcpKeyRingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOracleDbGcpKeyRingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOracleDbGcpKeyRingResponse")
	}
	return
}

// createOracleDbGcpKeyRing implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) createOracleDbGcpKeyRing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbGcpKeyRing", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOracleDbGcpKeyRingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/CreateOracleDbGcpKeyRing"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "CreateOracleDbGcpKeyRing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbGcpIdentityConnector Delete Oracle DB GCP Identity Connector Resource and delete connector too from Database Resource.
// A default retry strategy applies to this operation DeleteOracleDbGcpIdentityConnector()
func (client DbMulticloudGCPProviderClient) DeleteOracleDbGcpIdentityConnector(ctx context.Context, request DeleteOracleDbGcpIdentityConnectorRequest) (response DeleteOracleDbGcpIdentityConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbGcpIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbGcpIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbGcpIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbGcpIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbGcpIdentityConnectorResponse")
	}
	return
}

// deleteOracleDbGcpIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) deleteOracleDbGcpIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbGcpIdentityConnector/{oracleDbGcpIdentityConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbGcpIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpIdentityConnector/DeleteOracleDbGcpIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "DeleteOracleDbGcpIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOracleDbGcpKeyRing Delete  GCP Key Ring details.
// A default retry strategy applies to this operation DeleteOracleDbGcpKeyRing()
func (client DbMulticloudGCPProviderClient) DeleteOracleDbGcpKeyRing(ctx context.Context, request DeleteOracleDbGcpKeyRingRequest) (response DeleteOracleDbGcpKeyRingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOracleDbGcpKeyRing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOracleDbGcpKeyRingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOracleDbGcpKeyRingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOracleDbGcpKeyRingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOracleDbGcpKeyRingResponse")
	}
	return
}

// deleteOracleDbGcpKeyRing implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) deleteOracleDbGcpKeyRing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oracleDbGcpKeyRing/{oracleDbGcpKeyRingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOracleDbGcpKeyRingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/DeleteOracleDbGcpKeyRing"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "DeleteOracleDbGcpKeyRing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbGcpIdentityConnector Get Oracle DB Oracle DB GCP Identity Connector form a particular Resource ID.
// A default retry strategy applies to this operation GetOracleDbGcpIdentityConnector()
func (client DbMulticloudGCPProviderClient) GetOracleDbGcpIdentityConnector(ctx context.Context, request GetOracleDbGcpIdentityConnectorRequest) (response GetOracleDbGcpIdentityConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbGcpIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbGcpIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbGcpIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbGcpIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbGcpIdentityConnectorResponse")
	}
	return
}

// getOracleDbGcpIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) getOracleDbGcpIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbGcpIdentityConnector/{oracleDbGcpIdentityConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbGcpIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpIdentityConnector/GetOracleDbGcpIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "GetOracleDbGcpIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbGcpKey Get Oracle DB Google Cloud Key Details form a particular resource ID.
// A default retry strategy applies to this operation GetOracleDbGcpKey()
func (client DbMulticloudGCPProviderClient) GetOracleDbGcpKey(ctx context.Context, request GetOracleDbGcpKeyRequest) (response GetOracleDbGcpKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbGcpKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbGcpKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbGcpKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbGcpKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbGcpKeyResponse")
	}
	return
}

// getOracleDbGcpKey implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) getOracleDbGcpKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbGcpKey/{oracleDbGcpKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbGcpKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKey/GetOracleDbGcpKey"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "GetOracleDbGcpKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOracleDbGcpKeyRing Get Oracle GCP Key Ring Details form a particular Container Resource ID.
// A default retry strategy applies to this operation GetOracleDbGcpKeyRing()
func (client DbMulticloudGCPProviderClient) GetOracleDbGcpKeyRing(ctx context.Context, request GetOracleDbGcpKeyRingRequest) (response GetOracleDbGcpKeyRingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbGcpKeyRing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbGcpKeyRingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbGcpKeyRingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbGcpKeyRingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbGcpKeyRingResponse")
	}
	return
}

// getOracleDbGcpKeyRing implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) getOracleDbGcpKeyRing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbGcpKeyRing/{oracleDbGcpKeyRingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbGcpKeyRingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/GetOracleDbGcpKeyRing"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "GetOracleDbGcpKeyRing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbGcpIdentityConnectors Lists the all Oracle DB GCP Identity Connectors based on filters.
// A default retry strategy applies to this operation ListOracleDbGcpIdentityConnectors()
func (client DbMulticloudGCPProviderClient) ListOracleDbGcpIdentityConnectors(ctx context.Context, request ListOracleDbGcpIdentityConnectorsRequest) (response ListOracleDbGcpIdentityConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbGcpIdentityConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbGcpIdentityConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbGcpIdentityConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbGcpIdentityConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbGcpIdentityConnectorsResponse")
	}
	return
}

// listOracleDbGcpIdentityConnectors implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) listOracleDbGcpIdentityConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbGcpIdentityConnector", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbGcpIdentityConnectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpIdentityConnector/ListOracleDbGcpIdentityConnectors"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "ListOracleDbGcpIdentityConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbGcpKeyRings Lists the all DB GCP Key Rings based on filters.
// A default retry strategy applies to this operation ListOracleDbGcpKeyRings()
func (client DbMulticloudGCPProviderClient) ListOracleDbGcpKeyRings(ctx context.Context, request ListOracleDbGcpKeyRingsRequest) (response ListOracleDbGcpKeyRingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbGcpKeyRings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbGcpKeyRingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbGcpKeyRingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbGcpKeyRingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbGcpKeyRingsResponse")
	}
	return
}

// listOracleDbGcpKeyRings implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) listOracleDbGcpKeyRings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbGcpKeyRing", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbGcpKeyRingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/ListOracleDbGcpKeyRings"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "ListOracleDbGcpKeyRings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbGcpKeys Lists the all Oracle DB Google Cloud Keys based on filters.
// A default retry strategy applies to this operation ListOracleDbGcpKeys()
func (client DbMulticloudGCPProviderClient) ListOracleDbGcpKeys(ctx context.Context, request ListOracleDbGcpKeysRequest) (response ListOracleDbGcpKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbGcpKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbGcpKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbGcpKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbGcpKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbGcpKeysResponse")
	}
	return
}

// listOracleDbGcpKeys implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) listOracleDbGcpKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbGcpKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbGcpKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKey/ListOracleDbGcpKeys"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "ListOracleDbGcpKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshOracleDbGcpKeyRing Refresh Oracle GCP Key Ring details from backend.
// A default retry strategy applies to this operation RefreshOracleDbGcpKeyRing()
func (client DbMulticloudGCPProviderClient) RefreshOracleDbGcpKeyRing(ctx context.Context, request RefreshOracleDbGcpKeyRingRequest) (response RefreshOracleDbGcpKeyRingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshOracleDbGcpKeyRing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshOracleDbGcpKeyRingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshOracleDbGcpKeyRingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshOracleDbGcpKeyRingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshOracleDbGcpKeyRingResponse")
	}
	return
}

// refreshOracleDbGcpKeyRing implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) refreshOracleDbGcpKeyRing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oracleDbGcpKeyRing/{oracleDbGcpKeyRingId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshOracleDbGcpKeyRingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/RefreshOracleDbGcpKeyRing"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "RefreshOracleDbGcpKeyRing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbGcpIdentityConnector Modifies the existing Oracle DB GCP Identity Connector resource for a given ID.
// A default retry strategy applies to this operation UpdateOracleDbGcpIdentityConnector()
func (client DbMulticloudGCPProviderClient) UpdateOracleDbGcpIdentityConnector(ctx context.Context, request UpdateOracleDbGcpIdentityConnectorRequest) (response UpdateOracleDbGcpIdentityConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbGcpIdentityConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbGcpIdentityConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbGcpIdentityConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbGcpIdentityConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbGcpIdentityConnectorResponse")
	}
	return
}

// updateOracleDbGcpIdentityConnector implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) updateOracleDbGcpIdentityConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbGcpIdentityConnector/{oracleDbGcpIdentityConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbGcpIdentityConnectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpIdentityConnector/UpdateOracleDbGcpIdentityConnector"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "UpdateOracleDbGcpIdentityConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOracleDbGcpKeyRing Modifies the existing Oracle GCP Key Ring Details for a given ID.
// A default retry strategy applies to this operation UpdateOracleDbGcpKeyRing()
func (client DbMulticloudGCPProviderClient) UpdateOracleDbGcpKeyRing(ctx context.Context, request UpdateOracleDbGcpKeyRingRequest) (response UpdateOracleDbGcpKeyRingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOracleDbGcpKeyRing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOracleDbGcpKeyRingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOracleDbGcpKeyRingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOracleDbGcpKeyRingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOracleDbGcpKeyRingResponse")
	}
	return
}

// updateOracleDbGcpKeyRing implements the OCIOperation interface (enables retrying operations)
func (client DbMulticloudGCPProviderClient) updateOracleDbGcpKeyRing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oracleDbGcpKeyRing/{oracleDbGcpKeyRingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOracleDbGcpKeyRingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-multicloud-integrations/20240501/OracleDbGcpKeyRing/UpdateOracleDbGcpKeyRing"
		err = common.PostProcessServiceError(err, "DbMulticloudGCPProvider", "UpdateOracleDbGcpKeyRing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
