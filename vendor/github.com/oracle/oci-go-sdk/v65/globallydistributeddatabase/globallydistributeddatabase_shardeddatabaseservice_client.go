// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ShardedDatabaseServiceClient a client for ShardedDatabaseService
type ShardedDatabaseServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewShardedDatabaseServiceClientWithConfigurationProvider Creates a new default ShardedDatabaseService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewShardedDatabaseServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ShardedDatabaseServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("globallydistributeddatabase"); !enabled {
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
	return newShardedDatabaseServiceClientFromBaseClient(baseClient, provider)
}

// NewShardedDatabaseServiceClientWithOboToken Creates a new default ShardedDatabaseService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewShardedDatabaseServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ShardedDatabaseServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newShardedDatabaseServiceClientFromBaseClient(baseClient, configProvider)
}

func newShardedDatabaseServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ShardedDatabaseServiceClient, err error) {
	// ShardedDatabaseService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ShardedDatabaseService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ShardedDatabaseServiceClient{BaseClient: baseClient}
	client.BasePath = "20230301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ShardedDatabaseServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("globallydistributeddatabase", "https://globaldb.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ShardedDatabaseServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ShardedDatabaseServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangePrivateEndpointCompartment Move the private endpoint to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ChangePrivateEndpointCompartment.go.html to see an example of how to use ChangePrivateEndpointCompartment API.
// A default retry strategy applies to this operation ChangePrivateEndpointCompartment()
func (client ShardedDatabaseServiceClient) ChangePrivateEndpointCompartment(ctx context.Context, request ChangePrivateEndpointCompartmentRequest) (response ChangePrivateEndpointCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePrivateEndpointCompartmentResponse")
	}
	return
}

// changePrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) changePrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateEndpoints/{privateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/PrivateEndpoint/ChangePrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ChangePrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeShardedDatabaseCompartment Move the sharded database database and its dependent resources to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ChangeShardedDatabaseCompartment.go.html to see an example of how to use ChangeShardedDatabaseCompartment API.
// A default retry strategy applies to this operation ChangeShardedDatabaseCompartment()
func (client ShardedDatabaseServiceClient) ChangeShardedDatabaseCompartment(ctx context.Context, request ChangeShardedDatabaseCompartmentRequest) (response ChangeShardedDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeShardedDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeShardedDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeShardedDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeShardedDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeShardedDatabaseCompartmentResponse")
	}
	return
}

// changeShardedDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) changeShardedDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeShardedDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/ChangeShardedDatabaseCompartment"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ChangeShardedDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureShardedDatabaseGsms Configure new Global Service Manager(GSM aka shard manager) instances for the sharded database. Specify the names
// of old GSM instances that need to be replaced via parameter oldGsmNames in the request payload. Also specify
// whether rotated GSM instances shall be provisioned with latest image of GSM software or the image used by
// existing GSM instances shall be used.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ConfigureShardedDatabaseGsms.go.html to see an example of how to use ConfigureShardedDatabaseGsms API.
// A default retry strategy applies to this operation ConfigureShardedDatabaseGsms()
func (client ShardedDatabaseServiceClient) ConfigureShardedDatabaseGsms(ctx context.Context, request ConfigureShardedDatabaseGsmsRequest) (response ConfigureShardedDatabaseGsmsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureShardedDatabaseGsms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureShardedDatabaseGsmsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureShardedDatabaseGsmsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureShardedDatabaseGsmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureShardedDatabaseGsmsResponse")
	}
	return
}

// configureShardedDatabaseGsms implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) configureShardedDatabaseGsms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/configureGsms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureShardedDatabaseGsmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/ConfigureShardedDatabaseGsms"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ConfigureShardedDatabaseGsms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureSharding Once all components of sharded database are provisioned, and signed GSM certificates are successfully uploaded, this
// api shall be invoked to configure sharding on the sharded database. Note that this 'ConfigureSharding' API also needs to be
// invoked after successfully adding a new shard to the sharded database using PATCH api. If this API is not
// invoked after successfully adding a new shard, then that new shard will not be a participant in sharding topology of
// the sharded database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ConfigureSharding.go.html to see an example of how to use ConfigureSharding API.
// A default retry strategy applies to this operation ConfigureSharding()
func (client ShardedDatabaseServiceClient) ConfigureSharding(ctx context.Context, request ConfigureShardingRequest) (response ConfigureShardingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureSharding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureShardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureShardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureShardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureShardingResponse")
	}
	return
}

// configureSharding implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) configureSharding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/configureSharding", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureShardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/ConfigureSharding"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ConfigureSharding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivateEndpoint Creates a PrivateEndpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/CreatePrivateEndpoint.go.html to see an example of how to use CreatePrivateEndpoint API.
// A default retry strategy applies to this operation CreatePrivateEndpoint()
func (client ShardedDatabaseServiceClient) CreatePrivateEndpoint(ctx context.Context, request CreatePrivateEndpointRequest) (response CreatePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivateEndpointResponse")
	}
	return
}

// createPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) createPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "CreatePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateShardedDatabase Creates a Sharded Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/CreateShardedDatabase.go.html to see an example of how to use CreateShardedDatabase API.
// A default retry strategy applies to this operation CreateShardedDatabase()
func (client ShardedDatabaseServiceClient) CreateShardedDatabase(ctx context.Context, request CreateShardedDatabaseRequest) (response CreateShardedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateShardedDatabaseResponse")
	}
	return
}

// createShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) createShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "CreateShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &shardeddatabase{})
	return response, err
}

// DeletePrivateEndpoint Delete the given private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/DeletePrivateEndpoint.go.html to see an example of how to use DeletePrivateEndpoint API.
// A default retry strategy applies to this operation DeletePrivateEndpoint()
func (client ShardedDatabaseServiceClient) DeletePrivateEndpoint(ctx context.Context, request DeletePrivateEndpointRequest) (response DeletePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deletePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePrivateEndpointResponse")
	}
	return
}

// deletePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) deletePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/privateEndpoints/{privateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/PrivateEndpoint/DeletePrivateEndpoint"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "DeletePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteShardedDatabase Terminate the given sharded databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/DeleteShardedDatabase.go.html to see an example of how to use DeleteShardedDatabase API.
// A default retry strategy applies to this operation DeleteShardedDatabase()
func (client ShardedDatabaseServiceClient) DeleteShardedDatabase(ctx context.Context, request DeleteShardedDatabaseRequest) (response DeleteShardedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteShardedDatabaseResponse")
	}
	return
}

// deleteShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) deleteShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/shardedDatabases/{shardedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/DeleteShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "DeleteShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadGsmCertificateSigningRequest Generate the common certificate signing request for GSMs. Download the <sdb-prefix>.csr file from
// API response. Users can use this .csr file to generate the CA signed certificate, and as a next step
// use 'uploadSignedCertificateAndGenerateWallet' API to upload the CA signed certificate to GSM, and
// generate wallets for the GSM instances of the sharded database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/DownloadGsmCertificateSigningRequest.go.html to see an example of how to use DownloadGsmCertificateSigningRequest API.
// A default retry strategy applies to this operation DownloadGsmCertificateSigningRequest()
func (client ShardedDatabaseServiceClient) DownloadGsmCertificateSigningRequest(ctx context.Context, request DownloadGsmCertificateSigningRequestRequest) (response DownloadGsmCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadGsmCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadGsmCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadGsmCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadGsmCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadGsmCertificateSigningRequestResponse")
	}
	return
}

// downloadGsmCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) downloadGsmCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/downloadGsmCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadGsmCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/DownloadGsmCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "DownloadGsmCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FetchConnectionString Gets the Sharded Database Connection Strings.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/FetchConnectionString.go.html to see an example of how to use FetchConnectionString API.
// A default retry strategy applies to this operation FetchConnectionString()
func (client ShardedDatabaseServiceClient) FetchConnectionString(ctx context.Context, request FetchConnectionStringRequest) (response FetchConnectionStringResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.fetchConnectionString, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FetchConnectionStringResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FetchConnectionStringResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FetchConnectionStringResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FetchConnectionStringResponse")
	}
	return
}

// fetchConnectionString implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) fetchConnectionString(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/getConnectionString", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FetchConnectionStringResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/FetchConnectionString"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "FetchConnectionString", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FetchShardableCloudAutonomousVmClusters List of cloudAutonomousVMClusters for the given tenancy, that can be sharded.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/FetchShardableCloudAutonomousVmClusters.go.html to see an example of how to use FetchShardableCloudAutonomousVmClusters API.
// A default retry strategy applies to this operation FetchShardableCloudAutonomousVmClusters()
func (client ShardedDatabaseServiceClient) FetchShardableCloudAutonomousVmClusters(ctx context.Context, request FetchShardableCloudAutonomousVmClustersRequest) (response FetchShardableCloudAutonomousVmClustersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.fetchShardableCloudAutonomousVmClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FetchShardableCloudAutonomousVmClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FetchShardableCloudAutonomousVmClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FetchShardableCloudAutonomousVmClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FetchShardableCloudAutonomousVmClustersResponse")
	}
	return
}

// fetchShardableCloudAutonomousVmClusters implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) fetchShardableCloudAutonomousVmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/actions/listShardableCloudAutonomousVmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FetchShardableCloudAutonomousVmClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/FetchShardableCloudAutonomousVmClusters"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "FetchShardableCloudAutonomousVmClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateGsmCertificateSigningRequest Generate the certificate signing request for GSM instances of the sharded database. Once certificate signing
// request is generated, then customers can download the certificate signing request using
// 'downloadGsmCertificateSigningRequest' api call.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/GenerateGsmCertificateSigningRequest.go.html to see an example of how to use GenerateGsmCertificateSigningRequest API.
// A default retry strategy applies to this operation GenerateGsmCertificateSigningRequest()
func (client ShardedDatabaseServiceClient) GenerateGsmCertificateSigningRequest(ctx context.Context, request GenerateGsmCertificateSigningRequestRequest) (response GenerateGsmCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateGsmCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateGsmCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateGsmCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateGsmCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateGsmCertificateSigningRequestResponse")
	}
	return
}

// generateGsmCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) generateGsmCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/generateGsmCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateGsmCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/GenerateGsmCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "GenerateGsmCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateWallet Generate the wallet associated with sharded database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/GenerateWallet.go.html to see an example of how to use GenerateWallet API.
// A default retry strategy applies to this operation GenerateWallet()
func (client ShardedDatabaseServiceClient) GenerateWallet(ctx context.Context, request GenerateWalletRequest) (response GenerateWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateWalletResponse")
	}
	return
}

// generateWallet implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) generateWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/generateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/GenerateWallet"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "GenerateWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateEndpoint Get the PrivateEndpoint resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/GetPrivateEndpoint.go.html to see an example of how to use GetPrivateEndpoint API.
// A default retry strategy applies to this operation GetPrivateEndpoint()
func (client ShardedDatabaseServiceClient) GetPrivateEndpoint(ctx context.Context, request GetPrivateEndpointRequest) (response GetPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateEndpointResponse")
	}
	return
}

// getPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) getPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateEndpoints/{privateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/PrivateEndpoint/GetPrivateEndpoint"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "GetPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetShardedDatabase Gets the details of the Sharded database identified by given id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/GetShardedDatabase.go.html to see an example of how to use GetShardedDatabase API.
// A default retry strategy applies to this operation GetShardedDatabase()
func (client ShardedDatabaseServiceClient) GetShardedDatabase(ctx context.Context, request GetShardedDatabaseRequest) (response GetShardedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetShardedDatabaseResponse")
	}
	return
}

// getShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) getShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/shardedDatabases/{shardedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/GetShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "GetShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &shardeddatabase{})
	return response, err
}

// GetWorkRequest Gets details of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ShardedDatabaseServiceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivateEndpoints List of PrivateEndpoints.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListPrivateEndpoints.go.html to see an example of how to use ListPrivateEndpoints API.
// A default retry strategy applies to this operation ListPrivateEndpoints()
func (client ShardedDatabaseServiceClient) ListPrivateEndpoints(ctx context.Context, request ListPrivateEndpointsRequest) (response ListPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivateEndpointsResponse")
	}
	return
}

// listPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) listPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/PrivateEndpointCollection/ListPrivateEndpoints"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ListPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListShardedDatabases List of Sharded databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListShardedDatabases.go.html to see an example of how to use ListShardedDatabases API.
// A default retry strategy applies to this operation ListShardedDatabases()
func (client ShardedDatabaseServiceClient) ListShardedDatabases(ctx context.Context, request ListShardedDatabasesRequest) (response ListShardedDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listShardedDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListShardedDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListShardedDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListShardedDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListShardedDatabasesResponse")
	}
	return
}

// listShardedDatabases implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) listShardedDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/shardedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListShardedDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabaseCollection/ListShardedDatabases"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ListShardedDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a (paginated) list of errors for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ShardedDatabaseServiceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ShardedDatabaseServiceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ShardedDatabaseServiceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchShardedDatabase Patch operation to add, remove or update shards to the sharded database topology. In single patch
// operation, multiple shards can be either added, or removed or updated. Combination of inserts, update
// and remove in single operation is not allowed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/PatchShardedDatabase.go.html to see an example of how to use PatchShardedDatabase API.
// A default retry strategy applies to this operation PatchShardedDatabase()
func (client ShardedDatabaseServiceClient) PatchShardedDatabase(ctx context.Context, request PatchShardedDatabaseRequest) (response PatchShardedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchShardedDatabaseResponse")
	}
	return
}

// patchShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) patchShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/shardedDatabases/{shardedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/PatchShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "PatchShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PrevalidateShardedDatabase Sharded database pre-validation request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/PrevalidateShardedDatabase.go.html to see an example of how to use PrevalidateShardedDatabase API.
// A default retry strategy applies to this operation PrevalidateShardedDatabase()
func (client ShardedDatabaseServiceClient) PrevalidateShardedDatabase(ctx context.Context, request PrevalidateShardedDatabaseRequest) (response PrevalidateShardedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.prevalidateShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PrevalidateShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PrevalidateShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PrevalidateShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PrevalidateShardedDatabaseResponse")
	}
	return
}

// prevalidateShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) prevalidateShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/actions/prevalidate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PrevalidateShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/PrevalidateShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "PrevalidateShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReinstateProxyInstance API to reinstate the proxy instances associated with the private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ReinstateProxyInstance.go.html to see an example of how to use ReinstateProxyInstance API.
// A default retry strategy applies to this operation ReinstateProxyInstance()
func (client ShardedDatabaseServiceClient) ReinstateProxyInstance(ctx context.Context, request ReinstateProxyInstanceRequest) (response ReinstateProxyInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.reinstateProxyInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReinstateProxyInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReinstateProxyInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReinstateProxyInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReinstateProxyInstanceResponse")
	}
	return
}

// reinstateProxyInstance implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) reinstateProxyInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateEndpoints/{privateEndpointId}/actions/reinstateProxyInstance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReinstateProxyInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/PrivateEndpoint/ReinstateProxyInstance"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ReinstateProxyInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartShardedDatabase Start the shards, catalog and GSMs of Sharded Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/StartShardedDatabase.go.html to see an example of how to use StartShardedDatabase API.
// A default retry strategy applies to this operation StartShardedDatabase()
func (client ShardedDatabaseServiceClient) StartShardedDatabase(ctx context.Context, request StartShardedDatabaseRequest) (response StartShardedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartShardedDatabaseResponse")
	}
	return
}

// startShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) startShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/startDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/StartShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "StartShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopShardedDatabase Stop the shards, catalog and GSM instances for the sharded database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/StopShardedDatabase.go.html to see an example of how to use StopShardedDatabase API.
// A default retry strategy applies to this operation StopShardedDatabase()
func (client ShardedDatabaseServiceClient) StopShardedDatabase(ctx context.Context, request StopShardedDatabaseRequest) (response StopShardedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopShardedDatabaseResponse")
	}
	return
}

// stopShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) stopShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/stopDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/StopShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "StopShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePrivateEndpoint Updates the configuration of privateendpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/UpdatePrivateEndpoint.go.html to see an example of how to use UpdatePrivateEndpoint API.
// A default retry strategy applies to this operation UpdatePrivateEndpoint()
func (client ShardedDatabaseServiceClient) UpdatePrivateEndpoint(ctx context.Context, request UpdatePrivateEndpointRequest) (response UpdatePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePrivateEndpointResponse")
	}
	return
}

// updatePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) updatePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/privateEndpoints/{privateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/PrivateEndpoint/UpdatePrivateEndpoint"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "UpdatePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateShardedDatabase Updates the configuration of sharded database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/UpdateShardedDatabase.go.html to see an example of how to use UpdateShardedDatabase API.
// A default retry strategy applies to this operation UpdateShardedDatabase()
func (client ShardedDatabaseServiceClient) UpdateShardedDatabase(ctx context.Context, request UpdateShardedDatabaseRequest) (response UpdateShardedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateShardedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateShardedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateShardedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateShardedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateShardedDatabaseResponse")
	}
	return
}

// updateShardedDatabase implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) updateShardedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/shardedDatabases/{shardedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateShardedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/UpdateShardedDatabase"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "UpdateShardedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &shardeddatabase{})
	return response, err
}

// UploadSignedCertificateAndGenerateWallet Upload the CA signed certificate to the GSM instances and generate wallets for GSM instances of the
// sharded database. Customer shall provide the CA signed certificate key details by adding the certificate
// in request body.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/UploadSignedCertificateAndGenerateWallet.go.html to see an example of how to use UploadSignedCertificateAndGenerateWallet API.
// A default retry strategy applies to this operation UploadSignedCertificateAndGenerateWallet()
func (client ShardedDatabaseServiceClient) UploadSignedCertificateAndGenerateWallet(ctx context.Context, request UploadSignedCertificateAndGenerateWalletRequest) (response UploadSignedCertificateAndGenerateWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.uploadSignedCertificateAndGenerateWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadSignedCertificateAndGenerateWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadSignedCertificateAndGenerateWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadSignedCertificateAndGenerateWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadSignedCertificateAndGenerateWalletResponse")
	}
	return
}

// uploadSignedCertificateAndGenerateWallet implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) uploadSignedCertificateAndGenerateWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/uploadSignedCertificateAndGenerateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadSignedCertificateAndGenerateWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/UploadSignedCertificateAndGenerateWallet"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "UploadSignedCertificateAndGenerateWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateNetwork Validate the network connectivity between components of sharded database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/ValidateNetwork.go.html to see an example of how to use ValidateNetwork API.
// A default retry strategy applies to this operation ValidateNetwork()
func (client ShardedDatabaseServiceClient) ValidateNetwork(ctx context.Context, request ValidateNetworkRequest) (response ValidateNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateNetworkResponse")
	}
	return
}

// validateNetwork implements the OCIOperation interface (enables retrying operations)
func (client ShardedDatabaseServiceClient) validateNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/shardedDatabases/{shardedDatabaseId}/actions/validateNetwork", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-autonomous-database/20230301/ShardedDatabase/ValidateNetwork"
		err = common.PostProcessServiceError(err, "ShardedDatabaseService", "ValidateNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
