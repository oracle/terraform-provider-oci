// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DistributedAutonomousDbServiceClient a client for DistributedAutonomousDbService
type DistributedAutonomousDbServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDistributedAutonomousDbServiceClientWithConfigurationProvider Creates a new default DistributedAutonomousDbService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDistributedAutonomousDbServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DistributedAutonomousDbServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("distributeddatabase"); !enabled {
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
	return newDistributedAutonomousDbServiceClientFromBaseClient(baseClient, provider)
}

// NewDistributedAutonomousDbServiceClientWithOboToken Creates a new default DistributedAutonomousDbService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDistributedAutonomousDbServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DistributedAutonomousDbServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDistributedAutonomousDbServiceClientFromBaseClient(baseClient, configProvider)
}

func newDistributedAutonomousDbServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DistributedAutonomousDbServiceClient, err error) {
	// DistributedAutonomousDbService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DistributedAutonomousDbService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DistributedAutonomousDbServiceClient{BaseClient: baseClient}
	client.BasePath = "20250101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DistributedAutonomousDbServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("distributeddatabase", "https://globaldb.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DistributedAutonomousDbServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DistributedAutonomousDbServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddDistributedAutonomousDatabaseGdsControlNode Add new Global database services control(GDS CTL) node for the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/AddDistributedAutonomousDatabaseGdsControlNode.go.html to see an example of how to use AddDistributedAutonomousDatabaseGdsControlNode API.
// A default retry strategy applies to this operation AddDistributedAutonomousDatabaseGdsControlNode()
func (client DistributedAutonomousDbServiceClient) AddDistributedAutonomousDatabaseGdsControlNode(ctx context.Context, request AddDistributedAutonomousDatabaseGdsControlNodeRequest) (response AddDistributedAutonomousDatabaseGdsControlNodeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addDistributedAutonomousDatabaseGdsControlNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDistributedAutonomousDatabaseGdsControlNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDistributedAutonomousDatabaseGdsControlNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDistributedAutonomousDatabaseGdsControlNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDistributedAutonomousDatabaseGdsControlNodeResponse")
	}
	return
}

// addDistributedAutonomousDatabaseGdsControlNode implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) addDistributedAutonomousDatabaseGdsControlNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/addGdsControlNode", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDistributedAutonomousDatabaseGdsControlNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/AddDistributedAutonomousDatabaseGdsControlNode"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "AddDistributedAutonomousDatabaseGdsControlNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDistributedAutonomousDatabaseCompartment Move the Globally distributed autonomous database and its dependent resources to the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ChangeDistributedAutonomousDatabaseCompartment.go.html to see an example of how to use ChangeDistributedAutonomousDatabaseCompartment API.
// A default retry strategy applies to this operation ChangeDistributedAutonomousDatabaseCompartment()
func (client DistributedAutonomousDbServiceClient) ChangeDistributedAutonomousDatabaseCompartment(ctx context.Context, request ChangeDistributedAutonomousDatabaseCompartmentRequest) (response ChangeDistributedAutonomousDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDistributedAutonomousDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDistributedAutonomousDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDistributedAutonomousDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDistributedAutonomousDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDistributedAutonomousDatabaseCompartmentResponse")
	}
	return
}

// changeDistributedAutonomousDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) changeDistributedAutonomousDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDistributedAutonomousDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ChangeDistributedAutonomousDatabaseCompartment"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ChangeDistributedAutonomousDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDistributedAutonomousDbBackupConfig Change the DbBackupConfig for the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ChangeDistributedAutonomousDbBackupConfig.go.html to see an example of how to use ChangeDistributedAutonomousDbBackupConfig API.
// A default retry strategy applies to this operation ChangeDistributedAutonomousDbBackupConfig()
func (client DistributedAutonomousDbServiceClient) ChangeDistributedAutonomousDbBackupConfig(ctx context.Context, request ChangeDistributedAutonomousDbBackupConfigRequest) (response ChangeDistributedAutonomousDbBackupConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDistributedAutonomousDbBackupConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDistributedAutonomousDbBackupConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDistributedAutonomousDbBackupConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDistributedAutonomousDbBackupConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDistributedAutonomousDbBackupConfigResponse")
	}
	return
}

// changeDistributedAutonomousDbBackupConfig implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) changeDistributedAutonomousDbBackupConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/changeDbBackupConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDistributedAutonomousDbBackupConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ChangeDistributedAutonomousDbBackupConfig"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ChangeDistributedAutonomousDbBackupConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureDistributedAutonomousDatabaseGsmWallet Configure wallets on Global Service Manager(GSM) instances for a Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ConfigureDistributedAutonomousDatabaseGsmWallet.go.html to see an example of how to use ConfigureDistributedAutonomousDatabaseGsmWallet API.
// A default retry strategy applies to this operation ConfigureDistributedAutonomousDatabaseGsmWallet()
func (client DistributedAutonomousDbServiceClient) ConfigureDistributedAutonomousDatabaseGsmWallet(ctx context.Context, request ConfigureDistributedAutonomousDatabaseGsmWalletRequest) (response ConfigureDistributedAutonomousDatabaseGsmWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureDistributedAutonomousDatabaseGsmWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureDistributedAutonomousDatabaseGsmWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureDistributedAutonomousDatabaseGsmWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureDistributedAutonomousDatabaseGsmWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureDistributedAutonomousDatabaseGsmWalletResponse")
	}
	return
}

// configureDistributedAutonomousDatabaseGsmWallet implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) configureDistributedAutonomousDatabaseGsmWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/configureGsmWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureDistributedAutonomousDatabaseGsmWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ConfigureDistributedAutonomousDatabaseGsmWallet"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ConfigureDistributedAutonomousDatabaseGsmWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureDistributedAutonomousDatabaseGsms Configure new Global Service Manager(GSM aka shard manager) instances for the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ConfigureDistributedAutonomousDatabaseGsms.go.html to see an example of how to use ConfigureDistributedAutonomousDatabaseGsms API.
// A default retry strategy applies to this operation ConfigureDistributedAutonomousDatabaseGsms()
func (client DistributedAutonomousDbServiceClient) ConfigureDistributedAutonomousDatabaseGsms(ctx context.Context, request ConfigureDistributedAutonomousDatabaseGsmsRequest) (response ConfigureDistributedAutonomousDatabaseGsmsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureDistributedAutonomousDatabaseGsms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureDistributedAutonomousDatabaseGsmsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureDistributedAutonomousDatabaseGsmsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureDistributedAutonomousDatabaseGsmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureDistributedAutonomousDatabaseGsmsResponse")
	}
	return
}

// configureDistributedAutonomousDatabaseGsms implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) configureDistributedAutonomousDatabaseGsms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/configureGsms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureDistributedAutonomousDatabaseGsmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ConfigureDistributedAutonomousDatabaseGsms"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ConfigureDistributedAutonomousDatabaseGsms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureDistributedAutonomousDatabaseSharding Once all components of Globally distributed autonomous database are provisioned, this
// api shall be invoked to configure sharding on the Globally distributed autonomous database. Note that this 'ConfigureSharding' API also needs to be
// invoked after successfully adding a new shard to the Globally distributed autonomous database using PATCH api. If this API is not
// invoked after successfully adding a new shard, then that new shard will not be a participant in sharding topology of
// the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ConfigureDistributedAutonomousDatabaseSharding.go.html to see an example of how to use ConfigureDistributedAutonomousDatabaseSharding API.
// A default retry strategy applies to this operation ConfigureDistributedAutonomousDatabaseSharding()
func (client DistributedAutonomousDbServiceClient) ConfigureDistributedAutonomousDatabaseSharding(ctx context.Context, request ConfigureDistributedAutonomousDatabaseShardingRequest) (response ConfigureDistributedAutonomousDatabaseShardingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureDistributedAutonomousDatabaseSharding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureDistributedAutonomousDatabaseShardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureDistributedAutonomousDatabaseShardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureDistributedAutonomousDatabaseShardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureDistributedAutonomousDatabaseShardingResponse")
	}
	return
}

// configureDistributedAutonomousDatabaseSharding implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) configureDistributedAutonomousDatabaseSharding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/configureSharding", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureDistributedAutonomousDatabaseShardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ConfigureDistributedAutonomousDatabaseSharding"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ConfigureDistributedAutonomousDatabaseSharding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDistributedAutonomousDatabase Creates a Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/CreateDistributedAutonomousDatabase.go.html to see an example of how to use CreateDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation CreateDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) CreateDistributedAutonomousDatabase(ctx context.Context, request CreateDistributedAutonomousDatabaseRequest) (response CreateDistributedAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDistributedAutonomousDatabaseResponse")
	}
	return
}

// createDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) createDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "CreateDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDistributedAutonomousDatabase Terminate the given Globally distributed autonomous databases.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/DeleteDistributedAutonomousDatabase.go.html to see an example of how to use DeleteDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation DeleteDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) DeleteDistributedAutonomousDatabase(ctx context.Context, request DeleteDistributedAutonomousDatabaseRequest) (response DeleteDistributedAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDistributedAutonomousDatabaseResponse")
	}
	return
}

// deleteDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) deleteDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/DeleteDistributedAutonomousDatabase"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "DeleteDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest Generate the common certificate signing request for GSMs. Download the <globalautonomousdb-prefix>.csr file from
// API response. Users can use this .csr file to generate the CA signed certificate, and as a next step
// use 'uploadSignedCertificateAndGenerateWallet' API to upload the CA signed certificate to GSM, and
// generate wallets for the GSM instances of the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest.go.html to see an example of how to use DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest API.
// A default retry strategy applies to this operation DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest()
func (client DistributedAutonomousDbServiceClient) DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx context.Context, request DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestRequest) (response DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadDistributedAutonomousDatabaseGsmCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse")
	}
	return
}

// downloadDistributedAutonomousDatabaseGsmCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) downloadDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/downloadGsmCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "DownloadDistributedAutonomousDatabaseGsmCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest Generate the certificate signing request for GSM instances of the Globally distributed autonomous database. Once certificate signing
// request is generated, then customers can download the certificate signing request using
// 'downloadGsmCertificateSigningRequest' api call.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest.go.html to see an example of how to use GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest API.
// A default retry strategy applies to this operation GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest()
func (client DistributedAutonomousDbServiceClient) GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx context.Context, request GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestRequest) (response GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateDistributedAutonomousDatabaseGsmCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse")
	}
	return
}

// generateDistributedAutonomousDatabaseGsmCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) generateDistributedAutonomousDatabaseGsmCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/generateGsmCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "GenerateDistributedAutonomousDatabaseGsmCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateDistributedAutonomousDatabaseWallet Generate the wallet associated with Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GenerateDistributedAutonomousDatabaseWallet.go.html to see an example of how to use GenerateDistributedAutonomousDatabaseWallet API.
// A default retry strategy applies to this operation GenerateDistributedAutonomousDatabaseWallet()
func (client DistributedAutonomousDbServiceClient) GenerateDistributedAutonomousDatabaseWallet(ctx context.Context, request GenerateDistributedAutonomousDatabaseWalletRequest) (response GenerateDistributedAutonomousDatabaseWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateDistributedAutonomousDatabaseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateDistributedAutonomousDatabaseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateDistributedAutonomousDatabaseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateDistributedAutonomousDatabaseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateDistributedAutonomousDatabaseWalletResponse")
	}
	return
}

// generateDistributedAutonomousDatabaseWallet implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) generateDistributedAutonomousDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/generateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateDistributedAutonomousDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/GenerateDistributedAutonomousDatabaseWallet"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "GenerateDistributedAutonomousDatabaseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDistributedAutonomousDatabase Gets the details of the Globally distributed autonomous database identified by given id.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GetDistributedAutonomousDatabase.go.html to see an example of how to use GetDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation GetDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) GetDistributedAutonomousDatabase(ctx context.Context, request GetDistributedAutonomousDatabaseRequest) (response GetDistributedAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDistributedAutonomousDatabaseResponse")
	}
	return
}

// getDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) getDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/GetDistributedAutonomousDatabase"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "GetDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDistributedAutonomousDatabaseRaftMetric Operation to retrieve RAFT metrics for the Globally distributed autonomous database. If the Globally distributed
// autonomous database is not RAFT based then empty response is returned from the API.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GetDistributedAutonomousDatabaseRaftMetric.go.html to see an example of how to use GetDistributedAutonomousDatabaseRaftMetric API.
// A default retry strategy applies to this operation GetDistributedAutonomousDatabaseRaftMetric()
func (client DistributedAutonomousDbServiceClient) GetDistributedAutonomousDatabaseRaftMetric(ctx context.Context, request GetDistributedAutonomousDatabaseRaftMetricRequest) (response GetDistributedAutonomousDatabaseRaftMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDistributedAutonomousDatabaseRaftMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDistributedAutonomousDatabaseRaftMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDistributedAutonomousDatabaseRaftMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDistributedAutonomousDatabaseRaftMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDistributedAutonomousDatabaseRaftMetricResponse")
	}
	return
}

// getDistributedAutonomousDatabaseRaftMetric implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) getDistributedAutonomousDatabaseRaftMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/raftMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDistributedAutonomousDatabaseRaftMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/GetDistributedAutonomousDatabaseRaftMetric"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "GetDistributedAutonomousDatabaseRaftMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDistributedAutonomousDatabases List of Globally distributed autonomous databases.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ListDistributedAutonomousDatabases.go.html to see an example of how to use ListDistributedAutonomousDatabases API.
// A default retry strategy applies to this operation ListDistributedAutonomousDatabases()
func (client DistributedAutonomousDbServiceClient) ListDistributedAutonomousDatabases(ctx context.Context, request ListDistributedAutonomousDatabasesRequest) (response ListDistributedAutonomousDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDistributedAutonomousDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDistributedAutonomousDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDistributedAutonomousDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDistributedAutonomousDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDistributedAutonomousDatabasesResponse")
	}
	return
}

// listDistributedAutonomousDatabases implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) listDistributedAutonomousDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedAutonomousDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDistributedAutonomousDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabaseCollection/ListDistributedAutonomousDatabases"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ListDistributedAutonomousDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MoveDistributedAutonomousDatabaseReplicationUnit Move the replication units for RAFT based globally distributed autonomous database from source shard to destination shard.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/MoveDistributedAutonomousDatabaseReplicationUnit.go.html to see an example of how to use MoveDistributedAutonomousDatabaseReplicationUnit API.
// A default retry strategy applies to this operation MoveDistributedAutonomousDatabaseReplicationUnit()
func (client DistributedAutonomousDbServiceClient) MoveDistributedAutonomousDatabaseReplicationUnit(ctx context.Context, request MoveDistributedAutonomousDatabaseReplicationUnitRequest) (response MoveDistributedAutonomousDatabaseReplicationUnitResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.moveDistributedAutonomousDatabaseReplicationUnit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MoveDistributedAutonomousDatabaseReplicationUnitResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MoveDistributedAutonomousDatabaseReplicationUnitResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MoveDistributedAutonomousDatabaseReplicationUnitResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MoveDistributedAutonomousDatabaseReplicationUnitResponse")
	}
	return
}

// moveDistributedAutonomousDatabaseReplicationUnit implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) moveDistributedAutonomousDatabaseReplicationUnit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/moveReplicationUnit", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MoveDistributedAutonomousDatabaseReplicationUnitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/MoveDistributedAutonomousDatabaseReplicationUnit"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "MoveDistributedAutonomousDatabaseReplicationUnit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchDistributedAutonomousDatabase Patch operation to add, remove or update shards to the Globally distributed autonomous database topology. In single patch
// operation, multiple shards can be either added, or removed or updated. Combination of inserts, update
// and remove in single operation is not allowed.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/PatchDistributedAutonomousDatabase.go.html to see an example of how to use PatchDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation PatchDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) PatchDistributedAutonomousDatabase(ctx context.Context, request PatchDistributedAutonomousDatabaseRequest) (response PatchDistributedAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchDistributedAutonomousDatabaseResponse")
	}
	return
}

// patchDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) patchDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/PatchDistributedAutonomousDatabase"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "PatchDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RecreateFailedDistributedAutonomousDatabaseResource Recreate the failed resource for the Globally Distributed Autonomous Database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/RecreateFailedDistributedAutonomousDatabaseResource.go.html to see an example of how to use RecreateFailedDistributedAutonomousDatabaseResource API.
// A default retry strategy applies to this operation RecreateFailedDistributedAutonomousDatabaseResource()
func (client DistributedAutonomousDbServiceClient) RecreateFailedDistributedAutonomousDatabaseResource(ctx context.Context, request RecreateFailedDistributedAutonomousDatabaseResourceRequest) (response RecreateFailedDistributedAutonomousDatabaseResourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.recreateFailedDistributedAutonomousDatabaseResource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecreateFailedDistributedAutonomousDatabaseResourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecreateFailedDistributedAutonomousDatabaseResourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecreateFailedDistributedAutonomousDatabaseResourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecreateFailedDistributedAutonomousDatabaseResourceResponse")
	}
	return
}

// recreateFailedDistributedAutonomousDatabaseResource implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) recreateFailedDistributedAutonomousDatabaseResource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/recreateFailedResource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RecreateFailedDistributedAutonomousDatabaseResourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/RecreateFailedDistributedAutonomousDatabaseResource"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "RecreateFailedDistributedAutonomousDatabaseResource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateDistributedAutonomousDatabasePasswords Rotate the gsmuser and gsmcatuser passwords for shards and catalog of the Globally distributed autonomous database. This operation will also remove GdsCtlNodes if present.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/RotateDistributedAutonomousDatabasePasswords.go.html to see an example of how to use RotateDistributedAutonomousDatabasePasswords API.
// A default retry strategy applies to this operation RotateDistributedAutonomousDatabasePasswords()
func (client DistributedAutonomousDbServiceClient) RotateDistributedAutonomousDatabasePasswords(ctx context.Context, request RotateDistributedAutonomousDatabasePasswordsRequest) (response RotateDistributedAutonomousDatabasePasswordsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateDistributedAutonomousDatabasePasswords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateDistributedAutonomousDatabasePasswordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateDistributedAutonomousDatabasePasswordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateDistributedAutonomousDatabasePasswordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateDistributedAutonomousDatabasePasswordsResponse")
	}
	return
}

// rotateDistributedAutonomousDatabasePasswords implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) rotateDistributedAutonomousDatabasePasswords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/rotateDbPasswords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateDistributedAutonomousDatabasePasswordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/RotateDistributedAutonomousDatabasePasswords"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "RotateDistributedAutonomousDatabasePasswords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartDistributedAutonomousDatabase Start the shards, catalog and GSMs of Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/StartDistributedAutonomousDatabase.go.html to see an example of how to use StartDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation StartDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) StartDistributedAutonomousDatabase(ctx context.Context, request StartDistributedAutonomousDatabaseRequest) (response StartDistributedAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartDistributedAutonomousDatabaseResponse")
	}
	return
}

// startDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) startDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/startDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/StartDistributedAutonomousDatabase"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "StartDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopDistributedAutonomousDatabase Stop the shards, catalog and GSM instances for the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/StopDistributedAutonomousDatabase.go.html to see an example of how to use StopDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation StopDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) StopDistributedAutonomousDatabase(ctx context.Context, request StopDistributedAutonomousDatabaseRequest) (response StopDistributedAutonomousDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopDistributedAutonomousDatabaseResponse")
	}
	return
}

// stopDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) stopDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/stopDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/StopDistributedAutonomousDatabase"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "StopDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDistributedAutonomousDatabase Updates the configuration of the Globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/UpdateDistributedAutonomousDatabase.go.html to see an example of how to use UpdateDistributedAutonomousDatabase API.
// A default retry strategy applies to this operation UpdateDistributedAutonomousDatabase()
func (client DistributedAutonomousDbServiceClient) UpdateDistributedAutonomousDatabase(ctx context.Context, request UpdateDistributedAutonomousDatabaseRequest) (response UpdateDistributedAutonomousDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDistributedAutonomousDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDistributedAutonomousDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDistributedAutonomousDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDistributedAutonomousDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDistributedAutonomousDatabaseResponse")
	}
	return
}

// updateDistributedAutonomousDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) updateDistributedAutonomousDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDistributedAutonomousDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/UpdateDistributedAutonomousDatabase"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "UpdateDistributedAutonomousDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet Upload the CA signed certificate to the GSM instances and generate wallets for GSM instances of the
// Globally distributed autonomous database. Customer shall provide the CA signed certificate key details by adding the certificate
// in request body.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet.go.html to see an example of how to use UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet API.
// A default retry strategy applies to this operation UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet()
func (client DistributedAutonomousDbServiceClient) UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx context.Context, request UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletRequest) (response UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.uploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletResponse")
	}
	return
}

// uploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) uploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/uploadSignedCertificateAndGenerateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "UploadDistributedAutonomousDatabaseSignedCertificateAndGenerateWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateDistributedAutonomousDatabaseCaBundle Validate the CA Bundles consistency of the globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ValidateDistributedAutonomousDatabaseCaBundle.go.html to see an example of how to use ValidateDistributedAutonomousDatabaseCaBundle API.
// A default retry strategy applies to this operation ValidateDistributedAutonomousDatabaseCaBundle()
func (client DistributedAutonomousDbServiceClient) ValidateDistributedAutonomousDatabaseCaBundle(ctx context.Context, request ValidateDistributedAutonomousDatabaseCaBundleRequest) (response ValidateDistributedAutonomousDatabaseCaBundleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateDistributedAutonomousDatabaseCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateDistributedAutonomousDatabaseCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateDistributedAutonomousDatabaseCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateDistributedAutonomousDatabaseCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateDistributedAutonomousDatabaseCaBundleResponse")
	}
	return
}

// validateDistributedAutonomousDatabaseCaBundle implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) validateDistributedAutonomousDatabaseCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/validateCaBundle", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDistributedAutonomousDatabaseCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ValidateDistributedAutonomousDatabaseCaBundle"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ValidateDistributedAutonomousDatabaseCaBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateDistributedAutonomousDatabaseNetwork Validate the network connectivity between components of the globally distributed autonomous database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ValidateDistributedAutonomousDatabaseNetwork.go.html to see an example of how to use ValidateDistributedAutonomousDatabaseNetwork API.
// A default retry strategy applies to this operation ValidateDistributedAutonomousDatabaseNetwork()
func (client DistributedAutonomousDbServiceClient) ValidateDistributedAutonomousDatabaseNetwork(ctx context.Context, request ValidateDistributedAutonomousDatabaseNetworkRequest) (response ValidateDistributedAutonomousDatabaseNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateDistributedAutonomousDatabaseNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateDistributedAutonomousDatabaseNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateDistributedAutonomousDatabaseNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateDistributedAutonomousDatabaseNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateDistributedAutonomousDatabaseNetworkResponse")
	}
	return
}

// validateDistributedAutonomousDatabaseNetwork implements the OCIOperation interface (enables retrying operations)
func (client DistributedAutonomousDbServiceClient) validateDistributedAutonomousDatabaseNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedAutonomousDatabases/{distributedAutonomousDatabaseId}/actions/validateNetwork", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDistributedAutonomousDatabaseNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedAutonomousDatabase/ValidateDistributedAutonomousDatabaseNetwork"
		err = common.PostProcessServiceError(err, "DistributedAutonomousDbService", "ValidateDistributedAutonomousDatabaseNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
