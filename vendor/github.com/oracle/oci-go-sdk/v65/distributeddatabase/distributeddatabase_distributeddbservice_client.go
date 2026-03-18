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

// DistributedDbServiceClient a client for DistributedDbService
type DistributedDbServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDistributedDbServiceClientWithConfigurationProvider Creates a new default DistributedDbService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDistributedDbServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DistributedDbServiceClient, err error) {
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
	return newDistributedDbServiceClientFromBaseClient(baseClient, provider)
}

// NewDistributedDbServiceClientWithOboToken Creates a new default DistributedDbService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDistributedDbServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DistributedDbServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDistributedDbServiceClientFromBaseClient(baseClient, configProvider)
}

func newDistributedDbServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DistributedDbServiceClient, err error) {
	// DistributedDbService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DistributedDbService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DistributedDbServiceClient{BaseClient: baseClient}
	client.BasePath = "20250101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DistributedDbServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("distributeddatabase", "https://globaldb.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DistributedDbServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DistributedDbServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddDistributedDatabaseGdsControlNode Add new Global database services control(GDS CTL) node for the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/AddDistributedDatabaseGdsControlNode.go.html to see an example of how to use AddDistributedDatabaseGdsControlNode API.
// A default retry strategy applies to this operation AddDistributedDatabaseGdsControlNode()
func (client DistributedDbServiceClient) AddDistributedDatabaseGdsControlNode(ctx context.Context, request AddDistributedDatabaseGdsControlNodeRequest) (response AddDistributedDatabaseGdsControlNodeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addDistributedDatabaseGdsControlNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDistributedDatabaseGdsControlNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDistributedDatabaseGdsControlNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDistributedDatabaseGdsControlNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDistributedDatabaseGdsControlNodeResponse")
	}
	return
}

// addDistributedDatabaseGdsControlNode implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) addDistributedDatabaseGdsControlNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/addGdsControlNode", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDistributedDatabaseGdsControlNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/AddDistributedDatabaseGdsControlNode"
		err = common.PostProcessServiceError(err, "DistributedDbService", "AddDistributedDatabaseGdsControlNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDistributedDatabaseCompartment Move the Globally distributed database and its dependent resources to the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ChangeDistributedDatabaseCompartment.go.html to see an example of how to use ChangeDistributedDatabaseCompartment API.
// A default retry strategy applies to this operation ChangeDistributedDatabaseCompartment()
func (client DistributedDbServiceClient) ChangeDistributedDatabaseCompartment(ctx context.Context, request ChangeDistributedDatabaseCompartmentRequest) (response ChangeDistributedDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDistributedDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDistributedDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDistributedDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDistributedDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDistributedDatabaseCompartmentResponse")
	}
	return
}

// changeDistributedDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) changeDistributedDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDistributedDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/ChangeDistributedDatabaseCompartment"
		err = common.PostProcessServiceError(err, "DistributedDbService", "ChangeDistributedDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDistributedDbBackupConfig Change the DbBackupConfig for the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ChangeDistributedDbBackupConfig.go.html to see an example of how to use ChangeDistributedDbBackupConfig API.
// A default retry strategy applies to this operation ChangeDistributedDbBackupConfig()
func (client DistributedDbServiceClient) ChangeDistributedDbBackupConfig(ctx context.Context, request ChangeDistributedDbBackupConfigRequest) (response ChangeDistributedDbBackupConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDistributedDbBackupConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDistributedDbBackupConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDistributedDbBackupConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDistributedDbBackupConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDistributedDbBackupConfigResponse")
	}
	return
}

// changeDistributedDbBackupConfig implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) changeDistributedDbBackupConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/changeDbBackupConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDistributedDbBackupConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/ChangeDistributedDbBackupConfig"
		err = common.PostProcessServiceError(err, "DistributedDbService", "ChangeDistributedDbBackupConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureDistributedDatabaseGsms Configure new Global Service Manager(GSM aka shard manager) instances for the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ConfigureDistributedDatabaseGsms.go.html to see an example of how to use ConfigureDistributedDatabaseGsms API.
// A default retry strategy applies to this operation ConfigureDistributedDatabaseGsms()
func (client DistributedDbServiceClient) ConfigureDistributedDatabaseGsms(ctx context.Context, request ConfigureDistributedDatabaseGsmsRequest) (response ConfigureDistributedDatabaseGsmsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureDistributedDatabaseGsms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureDistributedDatabaseGsmsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureDistributedDatabaseGsmsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureDistributedDatabaseGsmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureDistributedDatabaseGsmsResponse")
	}
	return
}

// configureDistributedDatabaseGsms implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) configureDistributedDatabaseGsms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/configureGsms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureDistributedDatabaseGsmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/ConfigureDistributedDatabaseGsms"
		err = common.PostProcessServiceError(err, "DistributedDbService", "ConfigureDistributedDatabaseGsms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureDistributedDatabaseSharding Once all components of Globally distributed database are provisioned, and signed GSM certificates are successfully uploaded, this
// api shall be invoked to configure sharding on the Globally distributed database. Note that this 'ConfigureSharding' API also needs to be
// invoked after successfully adding a new shard to the Globally distributed database using PATCH api. If this API is not
// invoked after successfully adding a new shard, then that new shard will not be a participant in sharding topology of
// the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ConfigureDistributedDatabaseSharding.go.html to see an example of how to use ConfigureDistributedDatabaseSharding API.
// A default retry strategy applies to this operation ConfigureDistributedDatabaseSharding()
func (client DistributedDbServiceClient) ConfigureDistributedDatabaseSharding(ctx context.Context, request ConfigureDistributedDatabaseShardingRequest) (response ConfigureDistributedDatabaseShardingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureDistributedDatabaseSharding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureDistributedDatabaseShardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureDistributedDatabaseShardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureDistributedDatabaseShardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureDistributedDatabaseShardingResponse")
	}
	return
}

// configureDistributedDatabaseSharding implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) configureDistributedDatabaseSharding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/configureSharding", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureDistributedDatabaseShardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/ConfigureDistributedDatabaseSharding"
		err = common.PostProcessServiceError(err, "DistributedDbService", "ConfigureDistributedDatabaseSharding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDistributedDatabase Creates a Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/CreateDistributedDatabase.go.html to see an example of how to use CreateDistributedDatabase API.
// A default retry strategy applies to this operation CreateDistributedDatabase()
func (client DistributedDbServiceClient) CreateDistributedDatabase(ctx context.Context, request CreateDistributedDatabaseRequest) (response CreateDistributedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDistributedDatabaseResponse")
	}
	return
}

// createDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) createDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DistributedDbService", "CreateDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDistributedDatabase Terminate the given Globally distributed databases.
// For an EXADB_XS based distributed database, if the parameter mustDeleteInfra is set to true,
// then the VmCluster and DbStorageVault associated with each shard and catalog will also be deleted.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/DeleteDistributedDatabase.go.html to see an example of how to use DeleteDistributedDatabase API.
// A default retry strategy applies to this operation DeleteDistributedDatabase()
func (client DistributedDbServiceClient) DeleteDistributedDatabase(ctx context.Context, request DeleteDistributedDatabaseRequest) (response DeleteDistributedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDistributedDatabaseResponse")
	}
	return
}

// deleteDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) deleteDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/distributedDatabases/{distributedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/DeleteDistributedDatabase"
		err = common.PostProcessServiceError(err, "DistributedDbService", "DeleteDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadDistributedDatabaseGsmCertificateSigningRequest Generate the common certificate signing request for GSMs. Download the <globaldb-prefix>.csr file from
// API response. Users can use this .csr file to generate the CA signed certificate, and as a next step
// use 'uploadSignedCertificateAndGenerateWallet' API to upload the CA signed certificate to GSM, and
// generate wallets for the GSM instances of the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/DownloadDistributedDatabaseGsmCertificateSigningRequest.go.html to see an example of how to use DownloadDistributedDatabaseGsmCertificateSigningRequest API.
// A default retry strategy applies to this operation DownloadDistributedDatabaseGsmCertificateSigningRequest()
func (client DistributedDbServiceClient) DownloadDistributedDatabaseGsmCertificateSigningRequest(ctx context.Context, request DownloadDistributedDatabaseGsmCertificateSigningRequestRequest) (response DownloadDistributedDatabaseGsmCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadDistributedDatabaseGsmCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadDistributedDatabaseGsmCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadDistributedDatabaseGsmCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadDistributedDatabaseGsmCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadDistributedDatabaseGsmCertificateSigningRequestResponse")
	}
	return
}

// downloadDistributedDatabaseGsmCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) downloadDistributedDatabaseGsmCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/downloadGsmCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadDistributedDatabaseGsmCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/DownloadDistributedDatabaseGsmCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "DistributedDbService", "DownloadDistributedDatabaseGsmCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateDistributedDatabaseGsmCertificateSigningRequest Generate the certificate signing request for GSM instances of the Globally distributed database. Once certificate signing
// request is generated, then customers can download the certificate signing request using
// 'downloadGsmCertificateSigningRequest' api call.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GenerateDistributedDatabaseGsmCertificateSigningRequest.go.html to see an example of how to use GenerateDistributedDatabaseGsmCertificateSigningRequest API.
// A default retry strategy applies to this operation GenerateDistributedDatabaseGsmCertificateSigningRequest()
func (client DistributedDbServiceClient) GenerateDistributedDatabaseGsmCertificateSigningRequest(ctx context.Context, request GenerateDistributedDatabaseGsmCertificateSigningRequestRequest) (response GenerateDistributedDatabaseGsmCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateDistributedDatabaseGsmCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateDistributedDatabaseGsmCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateDistributedDatabaseGsmCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateDistributedDatabaseGsmCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateDistributedDatabaseGsmCertificateSigningRequestResponse")
	}
	return
}

// generateDistributedDatabaseGsmCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) generateDistributedDatabaseGsmCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/generateGsmCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateDistributedDatabaseGsmCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/GenerateDistributedDatabaseGsmCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "DistributedDbService", "GenerateDistributedDatabaseGsmCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateDistributedDatabaseWallet Generate the wallet associated with Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GenerateDistributedDatabaseWallet.go.html to see an example of how to use GenerateDistributedDatabaseWallet API.
// A default retry strategy applies to this operation GenerateDistributedDatabaseWallet()
func (client DistributedDbServiceClient) GenerateDistributedDatabaseWallet(ctx context.Context, request GenerateDistributedDatabaseWalletRequest) (response GenerateDistributedDatabaseWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateDistributedDatabaseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateDistributedDatabaseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateDistributedDatabaseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateDistributedDatabaseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateDistributedDatabaseWalletResponse")
	}
	return
}

// generateDistributedDatabaseWallet implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) generateDistributedDatabaseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/generateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateDistributedDatabaseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/GenerateDistributedDatabaseWallet"
		err = common.PostProcessServiceError(err, "DistributedDbService", "GenerateDistributedDatabaseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDistributedDatabase Gets the details of the Globally distributed database identified by given id.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GetDistributedDatabase.go.html to see an example of how to use GetDistributedDatabase API.
// A default retry strategy applies to this operation GetDistributedDatabase()
func (client DistributedDbServiceClient) GetDistributedDatabase(ctx context.Context, request GetDistributedDatabaseRequest) (response GetDistributedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDistributedDatabaseResponse")
	}
	return
}

// getDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) getDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedDatabases/{distributedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/GetDistributedDatabase"
		err = common.PostProcessServiceError(err, "DistributedDbService", "GetDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDistributedDatabaseRaftMetric Operation to retrieve RAFT metrics for the Globally distributed database. If the Globally distributed database is not
// RAFT based then empty response is returned from the API.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GetDistributedDatabaseRaftMetric.go.html to see an example of how to use GetDistributedDatabaseRaftMetric API.
// A default retry strategy applies to this operation GetDistributedDatabaseRaftMetric()
func (client DistributedDbServiceClient) GetDistributedDatabaseRaftMetric(ctx context.Context, request GetDistributedDatabaseRaftMetricRequest) (response GetDistributedDatabaseRaftMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDistributedDatabaseRaftMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDistributedDatabaseRaftMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDistributedDatabaseRaftMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDistributedDatabaseRaftMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDistributedDatabaseRaftMetricResponse")
	}
	return
}

// getDistributedDatabaseRaftMetric implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) getDistributedDatabaseRaftMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedDatabases/{distributedDatabaseId}/raftMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDistributedDatabaseRaftMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/GetDistributedDatabaseRaftMetric"
		err = common.PostProcessServiceError(err, "DistributedDbService", "GetDistributedDatabaseRaftMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDistributedDatabases List of Globally distributed databases.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ListDistributedDatabases.go.html to see an example of how to use ListDistributedDatabases API.
// A default retry strategy applies to this operation ListDistributedDatabases()
func (client DistributedDbServiceClient) ListDistributedDatabases(ctx context.Context, request ListDistributedDatabasesRequest) (response ListDistributedDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDistributedDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDistributedDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDistributedDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDistributedDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDistributedDatabasesResponse")
	}
	return
}

// listDistributedDatabases implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) listDistributedDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDistributedDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabaseCollection/ListDistributedDatabases"
		err = common.PostProcessServiceError(err, "DistributedDbService", "ListDistributedDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MoveDistributedDatabaseReplicationUnit Move the replication units for RAFT based globally distributed database from source shard to destination shard.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/MoveDistributedDatabaseReplicationUnit.go.html to see an example of how to use MoveDistributedDatabaseReplicationUnit API.
// A default retry strategy applies to this operation MoveDistributedDatabaseReplicationUnit()
func (client DistributedDbServiceClient) MoveDistributedDatabaseReplicationUnit(ctx context.Context, request MoveDistributedDatabaseReplicationUnitRequest) (response MoveDistributedDatabaseReplicationUnitResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.moveDistributedDatabaseReplicationUnit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MoveDistributedDatabaseReplicationUnitResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MoveDistributedDatabaseReplicationUnitResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MoveDistributedDatabaseReplicationUnitResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MoveDistributedDatabaseReplicationUnitResponse")
	}
	return
}

// moveDistributedDatabaseReplicationUnit implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) moveDistributedDatabaseReplicationUnit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/moveReplicationUnit", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MoveDistributedDatabaseReplicationUnitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/MoveDistributedDatabaseReplicationUnit"
		err = common.PostProcessServiceError(err, "DistributedDbService", "MoveDistributedDatabaseReplicationUnit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchDistributedDatabase Patch operation to add, remove or update shards to the Globally distributed database topology. In single patch
// operation, multiple shards can be either added, or removed or updated. Combination of inserts, update
// and remove in single operation is not allowed.
// For an EXADB_XS based distributed database, removing a shard with the parameter mustDeleteInfra set to true
// will also delete the associated VmCluster and DbStorageVault.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/PatchDistributedDatabase.go.html to see an example of how to use PatchDistributedDatabase API.
// A default retry strategy applies to this operation PatchDistributedDatabase()
func (client DistributedDbServiceClient) PatchDistributedDatabase(ctx context.Context, request PatchDistributedDatabaseRequest) (response PatchDistributedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchDistributedDatabaseResponse")
	}
	return
}

// patchDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) patchDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/distributedDatabases/{distributedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/PatchDistributedDatabase"
		err = common.PostProcessServiceError(err, "DistributedDbService", "PatchDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RecreateFailedDistributedDatabaseResource Recreate the failed resource for the Globally Distributed Database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/RecreateFailedDistributedDatabaseResource.go.html to see an example of how to use RecreateFailedDistributedDatabaseResource API.
// A default retry strategy applies to this operation RecreateFailedDistributedDatabaseResource()
func (client DistributedDbServiceClient) RecreateFailedDistributedDatabaseResource(ctx context.Context, request RecreateFailedDistributedDatabaseResourceRequest) (response RecreateFailedDistributedDatabaseResourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.recreateFailedDistributedDatabaseResource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecreateFailedDistributedDatabaseResourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecreateFailedDistributedDatabaseResourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecreateFailedDistributedDatabaseResourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecreateFailedDistributedDatabaseResourceResponse")
	}
	return
}

// recreateFailedDistributedDatabaseResource implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) recreateFailedDistributedDatabaseResource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/recreateFailedResource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RecreateFailedDistributedDatabaseResourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/RecreateFailedDistributedDatabaseResource"
		err = common.PostProcessServiceError(err, "DistributedDbService", "RecreateFailedDistributedDatabaseResource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateDistributedDatabasePasswords Rotate the gsmuser and gsmcatuser passwords for shards and catalog of the Globally distributed database.  This operation will also remove GdsCtlNodes if present.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/RotateDistributedDatabasePasswords.go.html to see an example of how to use RotateDistributedDatabasePasswords API.
// A default retry strategy applies to this operation RotateDistributedDatabasePasswords()
func (client DistributedDbServiceClient) RotateDistributedDatabasePasswords(ctx context.Context, request RotateDistributedDatabasePasswordsRequest) (response RotateDistributedDatabasePasswordsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateDistributedDatabasePasswords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateDistributedDatabasePasswordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateDistributedDatabasePasswordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateDistributedDatabasePasswordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateDistributedDatabasePasswordsResponse")
	}
	return
}

// rotateDistributedDatabasePasswords implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) rotateDistributedDatabasePasswords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/rotateDbPasswords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateDistributedDatabasePasswordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/RotateDistributedDatabasePasswords"
		err = common.PostProcessServiceError(err, "DistributedDbService", "RotateDistributedDatabasePasswords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartDistributedDatabase Start the shards, catalog and GSMs of Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/StartDistributedDatabase.go.html to see an example of how to use StartDistributedDatabase API.
// A default retry strategy applies to this operation StartDistributedDatabase()
func (client DistributedDbServiceClient) StartDistributedDatabase(ctx context.Context, request StartDistributedDatabaseRequest) (response StartDistributedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartDistributedDatabaseResponse")
	}
	return
}

// startDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) startDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/startDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/StartDistributedDatabase"
		err = common.PostProcessServiceError(err, "DistributedDbService", "StartDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopDistributedDatabase Stop the shards, catalog and GSM instances for the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/StopDistributedDatabase.go.html to see an example of how to use StopDistributedDatabase API.
// A default retry strategy applies to this operation StopDistributedDatabase()
func (client DistributedDbServiceClient) StopDistributedDatabase(ctx context.Context, request StopDistributedDatabaseRequest) (response StopDistributedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopDistributedDatabaseResponse")
	}
	return
}

// stopDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) stopDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/stopDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/StopDistributedDatabase"
		err = common.PostProcessServiceError(err, "DistributedDbService", "StopDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDistributedDatabase Updates the configuration of the Globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/UpdateDistributedDatabase.go.html to see an example of how to use UpdateDistributedDatabase API.
// A default retry strategy applies to this operation UpdateDistributedDatabase()
func (client DistributedDbServiceClient) UpdateDistributedDatabase(ctx context.Context, request UpdateDistributedDatabaseRequest) (response UpdateDistributedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDistributedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDistributedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDistributedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDistributedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDistributedDatabaseResponse")
	}
	return
}

// updateDistributedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) updateDistributedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/distributedDatabases/{distributedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDistributedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/UpdateDistributedDatabase"
		err = common.PostProcessServiceError(err, "DistributedDbService", "UpdateDistributedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadDistributedDatabaseSignedCertificateAndGenerateWallet Upload the CA signed certificate to the GSM instances and generate wallets for GSM instances of the
// Globally distributed database. Customer shall provide the CA signed certificate key details by adding the certificate
// in request body.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/UploadDistributedDatabaseSignedCertificateAndGenerateWallet.go.html to see an example of how to use UploadDistributedDatabaseSignedCertificateAndGenerateWallet API.
// A default retry strategy applies to this operation UploadDistributedDatabaseSignedCertificateAndGenerateWallet()
func (client DistributedDbServiceClient) UploadDistributedDatabaseSignedCertificateAndGenerateWallet(ctx context.Context, request UploadDistributedDatabaseSignedCertificateAndGenerateWalletRequest) (response UploadDistributedDatabaseSignedCertificateAndGenerateWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.uploadDistributedDatabaseSignedCertificateAndGenerateWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadDistributedDatabaseSignedCertificateAndGenerateWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadDistributedDatabaseSignedCertificateAndGenerateWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadDistributedDatabaseSignedCertificateAndGenerateWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadDistributedDatabaseSignedCertificateAndGenerateWalletResponse")
	}
	return
}

// uploadDistributedDatabaseSignedCertificateAndGenerateWallet implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) uploadDistributedDatabaseSignedCertificateAndGenerateWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/uploadSignedCertificateAndGenerateWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadDistributedDatabaseSignedCertificateAndGenerateWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/UploadDistributedDatabaseSignedCertificateAndGenerateWallet"
		err = common.PostProcessServiceError(err, "DistributedDbService", "UploadDistributedDatabaseSignedCertificateAndGenerateWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateDistributedDatabaseNetwork Validate the network connectivity between components of the globally distributed database.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ValidateDistributedDatabaseNetwork.go.html to see an example of how to use ValidateDistributedDatabaseNetwork API.
// A default retry strategy applies to this operation ValidateDistributedDatabaseNetwork()
func (client DistributedDbServiceClient) ValidateDistributedDatabaseNetwork(ctx context.Context, request ValidateDistributedDatabaseNetworkRequest) (response ValidateDistributedDatabaseNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateDistributedDatabaseNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateDistributedDatabaseNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateDistributedDatabaseNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateDistributedDatabaseNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateDistributedDatabaseNetworkResponse")
	}
	return
}

// validateDistributedDatabaseNetwork implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbServiceClient) validateDistributedDatabaseNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabases/{distributedDatabaseId}/actions/validateNetwork", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateDistributedDatabaseNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabase/ValidateDistributedDatabaseNetwork"
		err = common.PostProcessServiceError(err, "DistributedDbService", "ValidateDistributedDatabaseNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
