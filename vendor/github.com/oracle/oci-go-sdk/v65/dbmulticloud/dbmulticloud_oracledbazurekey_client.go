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

// OracleDbAzureKeyClient a client for OracleDbAzureKey
type OracleDbAzureKeyClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOracleDbAzureKeyClientWithConfigurationProvider Creates a new default OracleDbAzureKey client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOracleDbAzureKeyClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OracleDbAzureKeyClient, err error) {
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
	return newOracleDbAzureKeyClientFromBaseClient(baseClient, provider)
}

// NewOracleDbAzureKeyClientWithOboToken Creates a new default OracleDbAzureKey client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOracleDbAzureKeyClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OracleDbAzureKeyClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOracleDbAzureKeyClientFromBaseClient(baseClient, configProvider)
}

func newOracleDbAzureKeyClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OracleDbAzureKeyClient, err error) {
	// OracleDbAzureKey service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OracleDbAzureKey"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OracleDbAzureKeyClient{BaseClient: baseClient}
	client.BasePath = "20240501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OracleDbAzureKeyClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dbmulticloud", "https://dbmulticloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OracleDbAzureKeyClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OracleDbAzureKeyClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetOracleDbAzureKey Get Oracle DB Azure Key Details form a particular Container Resource ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureKey.go.html to see an example of how to use GetOracleDbAzureKey API.
// A default retry strategy applies to this operation GetOracleDbAzureKey()
func (client OracleDbAzureKeyClient) GetOracleDbAzureKey(ctx context.Context, request GetOracleDbAzureKeyRequest) (response GetOracleDbAzureKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOracleDbAzureKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOracleDbAzureKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOracleDbAzureKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOracleDbAzureKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOracleDbAzureKeyResponse")
	}
	return
}

// getOracleDbAzureKey implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureKeyClient) getOracleDbAzureKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureKey/{oracleDbAzureKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOracleDbAzureKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OracleDbAzureKey", "GetOracleDbAzureKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOracleDbAzureKeys Lists the all Oracle DB Azure Keys based on filters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureKeys.go.html to see an example of how to use ListOracleDbAzureKeys API.
// A default retry strategy applies to this operation ListOracleDbAzureKeys()
func (client OracleDbAzureKeyClient) ListOracleDbAzureKeys(ctx context.Context, request ListOracleDbAzureKeysRequest) (response ListOracleDbAzureKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOracleDbAzureKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOracleDbAzureKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOracleDbAzureKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOracleDbAzureKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOracleDbAzureKeysResponse")
	}
	return
}

// listOracleDbAzureKeys implements the OCIOperation interface (enables retrying operations)
func (client OracleDbAzureKeyClient) listOracleDbAzureKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oracleDbAzureKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOracleDbAzureKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OracleDbAzureKey", "ListOracleDbAzureKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
