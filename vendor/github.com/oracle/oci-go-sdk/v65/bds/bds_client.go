// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// BdsClient a client for Bds
type BdsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBdsClientWithConfigurationProvider Creates a new default Bds client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBdsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BdsClient, err error) {
	if enabled := common.CheckForEnabledServices("bds"); !enabled {
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
	return newBdsClientFromBaseClient(baseClient, provider)
}

// NewBdsClientWithOboToken Creates a new default Bds client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewBdsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BdsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newBdsClientFromBaseClient(baseClient, configProvider)
}

func newBdsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BdsClient, err error) {
	// Bds service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Bds"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = BdsClient{BaseClient: baseClient}
	client.BasePath = "20190531"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BdsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("bds", "https://bigdataservice.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BdsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BdsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateBdsMetastoreConfiguration Activate specified metastore configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ActivateBdsMetastoreConfiguration.go.html to see an example of how to use ActivateBdsMetastoreConfiguration API.
func (client BdsClient) ActivateBdsMetastoreConfiguration(ctx context.Context, request ActivateBdsMetastoreConfigurationRequest) (response ActivateBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.activateBdsMetastoreConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateBdsMetastoreConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateBdsMetastoreConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateBdsMetastoreConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateBdsMetastoreConfigurationResponse")
	}
	return
}

// activateBdsMetastoreConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) activateBdsMetastoreConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/metastoreConfigs/{metastoreConfigId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateBdsMetastoreConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/ActivateBdsMetastoreConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "ActivateBdsMetastoreConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddAutoScalingConfiguration Add an autoscale configuration to the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddAutoScalingConfiguration.go.html to see an example of how to use AddAutoScalingConfiguration API.
func (client BdsClient) AddAutoScalingConfiguration(ctx context.Context, request AddAutoScalingConfigurationRequest) (response AddAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddAutoScalingConfigurationResponse")
	}
	return
}

// addAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddAutoScalingConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "AddAutoScalingConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddBlockStorage Adds block storage to existing worker/compute only worker nodes. The same amount of  storage will be added to all worker/compute only worker nodes. No change will be made to storage that is already attached. Block storage cannot be removed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddBlockStorage.go.html to see an example of how to use AddBlockStorage API.
func (client BdsClient) AddBlockStorage(ctx context.Context, request AddBlockStorageRequest) (response AddBlockStorageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addBlockStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddBlockStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddBlockStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddBlockStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddBlockStorageResponse")
	}
	return
}

// addBlockStorage implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addBlockStorage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addBlockStorage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddBlockStorageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddBlockStorage"
		err = common.PostProcessServiceError(err, "Bds", "AddBlockStorage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddCloudSql Adds Cloud SQL to your cluster. You can use Cloud SQL to query against non-relational data stored in multiple big data sources, including Apache Hive, HDFS, Oracle NoSQL Database, and Apache HBase. Adding Cloud SQL adds a query server node to the cluster and creates cell servers on all the worker nodes in the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddCloudSql.go.html to see an example of how to use AddCloudSql API.
func (client BdsClient) AddCloudSql(ctx context.Context, request AddCloudSqlRequest) (response AddCloudSqlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addCloudSql, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddCloudSqlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddCloudSqlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddCloudSqlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddCloudSqlResponse")
	}
	return
}

// addCloudSql implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addCloudSql(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addCloudSql", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddCloudSqlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddCloudSql"
		err = common.PostProcessServiceError(err, "Bds", "AddCloudSql", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddKafka Adds Kafka to a cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddKafka.go.html to see an example of how to use AddKafka API.
func (client BdsClient) AddKafka(ctx context.Context, request AddKafkaRequest) (response AddKafkaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addKafka, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddKafkaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddKafkaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddKafkaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddKafkaResponse")
	}
	return
}

// addKafka implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addKafka(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addKafka", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddKafkaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddKafka"
		err = common.PostProcessServiceError(err, "Bds", "AddKafka", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddMasterNodes Increases the size (scales out) of a cluster by adding master nodes. The added master nodes will have the same shape and will have the same amount of attached block storage as other master nodes in the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddMasterNodes.go.html to see an example of how to use AddMasterNodes API.
func (client BdsClient) AddMasterNodes(ctx context.Context, request AddMasterNodesRequest) (response AddMasterNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addMasterNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMasterNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMasterNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMasterNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMasterNodesResponse")
	}
	return
}

// addMasterNodes implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addMasterNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addMasterNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMasterNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddMasterNodes"
		err = common.PostProcessServiceError(err, "Bds", "AddMasterNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddUtilityNodes Increases the size (scales out) of a cluster by adding utility nodes. The added utility nodes will have the same shape and will have the same amount of attached block storage as other utility nodes in the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddUtilityNodes.go.html to see an example of how to use AddUtilityNodes API.
func (client BdsClient) AddUtilityNodes(ctx context.Context, request AddUtilityNodesRequest) (response AddUtilityNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addUtilityNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddUtilityNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddUtilityNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddUtilityNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddUtilityNodesResponse")
	}
	return
}

// addUtilityNodes implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addUtilityNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addUtilityNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddUtilityNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddUtilityNodes"
		err = common.PostProcessServiceError(err, "Bds", "AddUtilityNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddWorkerNodes Increases the size (scales out) a cluster by adding worker nodes(data/compute). The added worker nodes will have the same shape and will have the same amount of attached block storage as other worker nodes in the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/AddWorkerNodes.go.html to see an example of how to use AddWorkerNodes API.
func (client BdsClient) AddWorkerNodes(ctx context.Context, request AddWorkerNodesRequest) (response AddWorkerNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addWorkerNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddWorkerNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddWorkerNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddWorkerNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddWorkerNodesResponse")
	}
	return
}

// addWorkerNodes implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) addWorkerNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addWorkerNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddWorkerNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/AddWorkerNodes"
		err = common.PostProcessServiceError(err, "Bds", "AddWorkerNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CertificateServiceInfo A list of services and their certificate details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/CertificateServiceInfo.go.html to see an example of how to use CertificateServiceInfo API.
func (client BdsClient) CertificateServiceInfo(ctx context.Context, request CertificateServiceInfoRequest) (response CertificateServiceInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.certificateServiceInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CertificateServiceInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CertificateServiceInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CertificateServiceInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CertificateServiceInfoResponse")
	}
	return
}

// certificateServiceInfo implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) certificateServiceInfo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/fetchOdhServiceCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CertificateServiceInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/CertificateServiceInfo"
		err = common.PostProcessServiceError(err, "Bds", "CertificateServiceInfo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeBdsInstanceCompartment Moves a Big Data Service cluster into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ChangeBdsInstanceCompartment.go.html to see an example of how to use ChangeBdsInstanceCompartment API.
func (client BdsClient) ChangeBdsInstanceCompartment(ctx context.Context, request ChangeBdsInstanceCompartmentRequest) (response ChangeBdsInstanceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeBdsInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBdsInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBdsInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBdsInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBdsInstanceCompartmentResponse")
	}
	return
}

// changeBdsInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) changeBdsInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBdsInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ChangeBdsInstanceCompartment"
		err = common.PostProcessServiceError(err, "Bds", "ChangeBdsInstanceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeShape Changes the size of a cluster by scaling up or scaling down the nodes. Nodes are scaled up or down by changing the shapes of all the nodes of the same type to the next larger or smaller shape. The node types are master, utility, worker, and Cloud SQL. Only nodes with VM-STANDARD shapes can be scaled.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ChangeShape.go.html to see an example of how to use ChangeShape API.
func (client BdsClient) ChangeShape(ctx context.Context, request ChangeShapeRequest) (response ChangeShapeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeShape, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeShapeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeShapeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeShapeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeShapeResponse")
	}
	return
}

// changeShape implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) changeShape(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/changeShape", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeShapeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ChangeShape"
		err = common.PostProcessServiceError(err, "Bds", "ChangeShape", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBdsApiKey Create an API key on behalf of the specified user.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/CreateBdsApiKey.go.html to see an example of how to use CreateBdsApiKey API.
func (client BdsClient) CreateBdsApiKey(ctx context.Context, request CreateBdsApiKeyRequest) (response CreateBdsApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBdsApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBdsApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBdsApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBdsApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBdsApiKeyResponse")
	}
	return
}

// createBdsApiKey implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) createBdsApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/apiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBdsApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsApiKey/CreateBdsApiKey"
		err = common.PostProcessServiceError(err, "Bds", "CreateBdsApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBdsInstance Creates a Big Data Service cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/CreateBdsInstance.go.html to see an example of how to use CreateBdsInstance API.
func (client BdsClient) CreateBdsInstance(ctx context.Context, request CreateBdsInstanceRequest) (response CreateBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBdsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBdsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBdsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBdsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBdsInstanceResponse")
	}
	return
}

// createBdsInstance implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) createBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/CreateBdsInstance"
		err = common.PostProcessServiceError(err, "Bds", "CreateBdsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBdsMetastoreConfiguration Create and activate external metastore configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/CreateBdsMetastoreConfiguration.go.html to see an example of how to use CreateBdsMetastoreConfiguration API.
func (client BdsClient) CreateBdsMetastoreConfiguration(ctx context.Context, request CreateBdsMetastoreConfigurationRequest) (response CreateBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBdsMetastoreConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBdsMetastoreConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBdsMetastoreConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBdsMetastoreConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBdsMetastoreConfigurationResponse")
	}
	return
}

// createBdsMetastoreConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) createBdsMetastoreConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/metastoreConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBdsMetastoreConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/CreateBdsMetastoreConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "CreateBdsMetastoreConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBdsApiKey Deletes the user's API key represented by the provided ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/DeleteBdsApiKey.go.html to see an example of how to use DeleteBdsApiKey API.
func (client BdsClient) DeleteBdsApiKey(ctx context.Context, request DeleteBdsApiKeyRequest) (response DeleteBdsApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBdsApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBdsApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBdsApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBdsApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBdsApiKeyResponse")
	}
	return
}

// deleteBdsApiKey implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) deleteBdsApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bdsInstances/{bdsInstanceId}/apiKeys/{apiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBdsApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsApiKey/DeleteBdsApiKey"
		err = common.PostProcessServiceError(err, "Bds", "DeleteBdsApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBdsInstance Deletes the cluster identified by the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/DeleteBdsInstance.go.html to see an example of how to use DeleteBdsInstance API.
func (client BdsClient) DeleteBdsInstance(ctx context.Context, request DeleteBdsInstanceRequest) (response DeleteBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBdsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBdsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBdsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBdsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBdsInstanceResponse")
	}
	return
}

// deleteBdsInstance implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) deleteBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bdsInstances/{bdsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/DeleteBdsInstance"
		err = common.PostProcessServiceError(err, "Bds", "DeleteBdsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBdsMetastoreConfiguration Delete the BDS metastore configuration represented by the provided ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/DeleteBdsMetastoreConfiguration.go.html to see an example of how to use DeleteBdsMetastoreConfiguration API.
func (client BdsClient) DeleteBdsMetastoreConfiguration(ctx context.Context, request DeleteBdsMetastoreConfigurationRequest) (response DeleteBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBdsMetastoreConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBdsMetastoreConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBdsMetastoreConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBdsMetastoreConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBdsMetastoreConfigurationResponse")
	}
	return
}

// deleteBdsMetastoreConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) deleteBdsMetastoreConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bdsInstances/{bdsInstanceId}/metastoreConfigs/{metastoreConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBdsMetastoreConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/DeleteBdsMetastoreConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "DeleteBdsMetastoreConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableCertificate Disabling TLS/SSL for various ODH services running on the BDS cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/DisableCertificate.go.html to see an example of how to use DisableCertificate API.
func (client BdsClient) DisableCertificate(ctx context.Context, request DisableCertificateRequest) (response DisableCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.disableCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableCertificateResponse")
	}
	return
}

// disableCertificate implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) disableCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/disableOdhServiceCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/DisableCertificate"
		err = common.PostProcessServiceError(err, "Bds", "DisableCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableCertificate Configuring TLS/SSL for various ODH services running on the BDS cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/EnableCertificate.go.html to see an example of how to use EnableCertificate API.
func (client BdsClient) EnableCertificate(ctx context.Context, request EnableCertificateRequest) (response EnableCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.enableCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableCertificateResponse")
	}
	return
}

// enableCertificate implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) enableCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/enableOdhServiceCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/EnableCertificate"
		err = common.PostProcessServiceError(err, "Bds", "EnableCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExecuteBootstrapScript Execute bootstrap script.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ExecuteBootstrapScript.go.html to see an example of how to use ExecuteBootstrapScript API.
func (client BdsClient) ExecuteBootstrapScript(ctx context.Context, request ExecuteBootstrapScriptRequest) (response ExecuteBootstrapScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.executeBootstrapScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExecuteBootstrapScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExecuteBootstrapScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExecuteBootstrapScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExecuteBootstrapScriptResponse")
	}
	return
}

// executeBootstrapScript implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) executeBootstrapScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/executeBootstrapScript", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExecuteBootstrapScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ExecuteBootstrapScript"
		err = common.PostProcessServiceError(err, "Bds", "ExecuteBootstrapScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutoScalingConfiguration Returns details of the autoscale configuration identified by the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetAutoScalingConfiguration.go.html to see an example of how to use GetAutoScalingConfiguration API.
func (client BdsClient) GetAutoScalingConfiguration(ctx context.Context, request GetAutoScalingConfigurationRequest) (response GetAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAutoScalingConfigurationResponse")
	}
	return
}

// getAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) getAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/GetAutoScalingConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "GetAutoScalingConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBdsApiKey Returns the user's API key information for the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetBdsApiKey.go.html to see an example of how to use GetBdsApiKey API.
func (client BdsClient) GetBdsApiKey(ctx context.Context, request GetBdsApiKeyRequest) (response GetBdsApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBdsApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBdsApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBdsApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBdsApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBdsApiKeyResponse")
	}
	return
}

// getBdsApiKey implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) getBdsApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/apiKeys/{apiKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBdsApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsApiKey/GetBdsApiKey"
		err = common.PostProcessServiceError(err, "Bds", "GetBdsApiKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBdsInstance Returns information about the Big Data Service cluster identified by the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetBdsInstance.go.html to see an example of how to use GetBdsInstance API.
func (client BdsClient) GetBdsInstance(ctx context.Context, request GetBdsInstanceRequest) (response GetBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBdsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBdsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBdsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBdsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBdsInstanceResponse")
	}
	return
}

// getBdsInstance implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) getBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/GetBdsInstance"
		err = common.PostProcessServiceError(err, "Bds", "GetBdsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBdsMetastoreConfiguration Returns the BDS Metastore configuration information for the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetBdsMetastoreConfiguration.go.html to see an example of how to use GetBdsMetastoreConfiguration API.
func (client BdsClient) GetBdsMetastoreConfiguration(ctx context.Context, request GetBdsMetastoreConfigurationRequest) (response GetBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBdsMetastoreConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBdsMetastoreConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBdsMetastoreConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBdsMetastoreConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBdsMetastoreConfigurationResponse")
	}
	return
}

// getBdsMetastoreConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) getBdsMetastoreConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/metastoreConfigs/{metastoreConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBdsMetastoreConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/GetBdsMetastoreConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "GetBdsMetastoreConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOsPatchDetails Get the details of an os patch
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetOsPatchDetails.go.html to see an example of how to use GetOsPatchDetails API.
func (client BdsClient) GetOsPatchDetails(ctx context.Context, request GetOsPatchDetailsRequest) (response GetOsPatchDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getOsPatchDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOsPatchDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOsPatchDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOsPatchDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOsPatchDetailsResponse")
	}
	return
}

// getOsPatchDetails implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) getOsPatchDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/getOsPatch", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOsPatchDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/GetOsPatchDetails"
		err = common.PostProcessServiceError(err, "Bds", "GetOsPatchDetails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Returns the status of the work request identified by the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client BdsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client BdsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Bds", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallOsPatch Install an os patch on a cluster
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/InstallOsPatch.go.html to see an example of how to use InstallOsPatch API.
func (client BdsClient) InstallOsPatch(ctx context.Context, request InstallOsPatchRequest) (response InstallOsPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.installOsPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallOsPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallOsPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallOsPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallOsPatchResponse")
	}
	return
}

// installOsPatch implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) installOsPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/installOsPatch", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallOsPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/InstallOsPatch"
		err = common.PostProcessServiceError(err, "Bds", "InstallOsPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPatch Install the specified patch to this cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/InstallPatch.go.html to see an example of how to use InstallPatch API.
func (client BdsClient) InstallPatch(ctx context.Context, request InstallPatchRequest) (response InstallPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.installPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallPatchResponse")
	}
	return
}

// installPatch implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) installPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/installPatch", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/InstallPatch"
		err = common.PostProcessServiceError(err, "Bds", "InstallPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutoScalingConfigurations Returns information about the autoscaling configurations for a cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListAutoScalingConfigurations.go.html to see an example of how to use ListAutoScalingConfigurations API.
func (client BdsClient) ListAutoScalingConfigurations(ctx context.Context, request ListAutoScalingConfigurationsRequest) (response ListAutoScalingConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAutoScalingConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAutoScalingConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAutoScalingConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAutoScalingConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAutoScalingConfigurationsResponse")
	}
	return
}

// listAutoScalingConfigurations implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listAutoScalingConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAutoScalingConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ListAutoScalingConfigurations"
		err = common.PostProcessServiceError(err, "Bds", "ListAutoScalingConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBdsApiKeys Returns a list of all API keys associated with this Big Data Service cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListBdsApiKeys.go.html to see an example of how to use ListBdsApiKeys API.
func (client BdsClient) ListBdsApiKeys(ctx context.Context, request ListBdsApiKeysRequest) (response ListBdsApiKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBdsApiKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBdsApiKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBdsApiKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBdsApiKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBdsApiKeysResponse")
	}
	return
}

// listBdsApiKeys implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listBdsApiKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/apiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBdsApiKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsApiKey/ListBdsApiKeys"
		err = common.PostProcessServiceError(err, "Bds", "ListBdsApiKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBdsInstances Returns a list of all Big Data Service clusters in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListBdsInstances.go.html to see an example of how to use ListBdsInstances API.
func (client BdsClient) ListBdsInstances(ctx context.Context, request ListBdsInstancesRequest) (response ListBdsInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBdsInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBdsInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBdsInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBdsInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBdsInstancesResponse")
	}
	return
}

// listBdsInstances implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listBdsInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBdsInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstanceSummary/ListBdsInstances"
		err = common.PostProcessServiceError(err, "Bds", "ListBdsInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBdsMetastoreConfigurations Returns a list of metastore configurations ssociated with this Big Data Service cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListBdsMetastoreConfigurations.go.html to see an example of how to use ListBdsMetastoreConfigurations API.
func (client BdsClient) ListBdsMetastoreConfigurations(ctx context.Context, request ListBdsMetastoreConfigurationsRequest) (response ListBdsMetastoreConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBdsMetastoreConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBdsMetastoreConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBdsMetastoreConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBdsMetastoreConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBdsMetastoreConfigurationsResponse")
	}
	return
}

// listBdsMetastoreConfigurations implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listBdsMetastoreConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/metastoreConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBdsMetastoreConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/ListBdsMetastoreConfigurations"
		err = common.PostProcessServiceError(err, "Bds", "ListBdsMetastoreConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOsPatches List all available os patches for a given cluster
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListOsPatches.go.html to see an example of how to use ListOsPatches API.
func (client BdsClient) ListOsPatches(ctx context.Context, request ListOsPatchesRequest) (response ListOsPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.listOsPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOsPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOsPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOsPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOsPatchesResponse")
	}
	return
}

// listOsPatches implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listOsPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/listOsPatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOsPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ListOsPatches"
		err = common.PostProcessServiceError(err, "Bds", "ListOsPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatchHistories List the patch history of this cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListPatchHistories.go.html to see an example of how to use ListPatchHistories API.
func (client BdsClient) ListPatchHistories(ctx context.Context, request ListPatchHistoriesRequest) (response ListPatchHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatchHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchHistoriesResponse")
	}
	return
}

// listPatchHistories implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listPatchHistories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/patchHistory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ListPatchHistories"
		err = common.PostProcessServiceError(err, "Bds", "ListPatchHistories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatches List all the available patches for this cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListPatches.go.html to see an example of how to use ListPatches API.
func (client BdsClient) ListPatches(ctx context.Context, request ListPatchesRequest) (response ListPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchesResponse")
	}
	return
}

// listPatches implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/ListPatches"
		err = common.PostProcessServiceError(err, "Bds", "ListPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a paginated list of errors for a work request identified by the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client BdsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client BdsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Bds", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client BdsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client BdsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Bds", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client BdsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client BdsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Bds", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAutoScalingConfiguration Deletes an autoscale configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/RemoveAutoScalingConfiguration.go.html to see an example of how to use RemoveAutoScalingConfiguration API.
func (client BdsClient) RemoveAutoScalingConfiguration(ctx context.Context, request RemoveAutoScalingConfigurationRequest) (response RemoveAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveAutoScalingConfigurationResponse")
	}
	return
}

// removeAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) removeAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}/actions/remove", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RemoveAutoScalingConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "RemoveAutoScalingConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveCloudSql Removes Cloud SQL from the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/RemoveCloudSql.go.html to see an example of how to use RemoveCloudSql API.
func (client BdsClient) RemoveCloudSql(ctx context.Context, request RemoveCloudSqlRequest) (response RemoveCloudSqlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeCloudSql, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveCloudSqlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveCloudSqlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveCloudSqlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveCloudSqlResponse")
	}
	return
}

// removeCloudSql implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) removeCloudSql(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/removeCloudSql", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveCloudSqlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RemoveCloudSql"
		err = common.PostProcessServiceError(err, "Bds", "RemoveCloudSql", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveKafka Remove Kafka from the cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/RemoveKafka.go.html to see an example of how to use RemoveKafka API.
func (client BdsClient) RemoveKafka(ctx context.Context, request RemoveKafkaRequest) (response RemoveKafkaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeKafka, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveKafkaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveKafkaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveKafkaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveKafkaResponse")
	}
	return
}

// removeKafka implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) removeKafka(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/removeKafka", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveKafkaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RemoveKafka"
		err = common.PostProcessServiceError(err, "Bds", "RemoveKafka", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveNode Remove a single node of a Big Data Service cluster
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/RemoveNode.go.html to see an example of how to use RemoveNode API.
func (client BdsClient) RemoveNode(ctx context.Context, request RemoveNodeRequest) (response RemoveNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveNodeResponse")
	}
	return
}

// removeNode implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) removeNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/removeNode", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RemoveNode"
		err = common.PostProcessServiceError(err, "Bds", "RemoveNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RenewCertificate Renewing TLS/SSL for various ODH services running on the BDS cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/RenewCertificate.go.html to see an example of how to use RenewCertificate API.
func (client BdsClient) RenewCertificate(ctx context.Context, request RenewCertificateRequest) (response RenewCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.renewCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RenewCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RenewCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RenewCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RenewCertificateResponse")
	}
	return
}

// renewCertificate implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) renewCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/renewOdhServiceCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RenewCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RenewCertificate"
		err = common.PostProcessServiceError(err, "Bds", "RenewCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartNode Restarts a single node of a Big Data Service cluster
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/RestartNode.go.html to see an example of how to use RestartNode API.
func (client BdsClient) RestartNode(ctx context.Context, request RestartNodeRequest) (response RestartNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.restartNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartNodeResponse")
	}
	return
}

// restartNode implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) restartNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/restartNode", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestartNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RestartNode"
		err = common.PostProcessServiceError(err, "Bds", "RestartNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartBdsInstance Starts the BDS cluster that was stopped earlier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/StartBdsInstance.go.html to see an example of how to use StartBdsInstance API.
func (client BdsClient) StartBdsInstance(ctx context.Context, request StartBdsInstanceRequest) (response StartBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startBdsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartBdsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartBdsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartBdsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartBdsInstanceResponse")
	}
	return
}

// startBdsInstance implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) startBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/StartBdsInstance"
		err = common.PostProcessServiceError(err, "Bds", "StartBdsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopBdsInstance Stops the BDS cluster that can be started at later point of time.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/StopBdsInstance.go.html to see an example of how to use StopBdsInstance API.
func (client BdsClient) StopBdsInstance(ctx context.Context, request StopBdsInstanceRequest) (response StopBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopBdsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopBdsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopBdsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopBdsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopBdsInstanceResponse")
	}
	return
}

// stopBdsInstance implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) stopBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/StopBdsInstance"
		err = common.PostProcessServiceError(err, "Bds", "StopBdsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestBdsMetastoreConfiguration Test specified metastore configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/TestBdsMetastoreConfiguration.go.html to see an example of how to use TestBdsMetastoreConfiguration API.
func (client BdsClient) TestBdsMetastoreConfiguration(ctx context.Context, request TestBdsMetastoreConfigurationRequest) (response TestBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.testBdsMetastoreConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestBdsMetastoreConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestBdsMetastoreConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestBdsMetastoreConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestBdsMetastoreConfigurationResponse")
	}
	return
}

// testBdsMetastoreConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) testBdsMetastoreConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/metastoreConfigs/{metastoreConfigId}/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestBdsMetastoreConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/TestBdsMetastoreConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "TestBdsMetastoreConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestBdsObjectStorageConnection Test access to specified Object Storage bucket using the API key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/TestBdsObjectStorageConnection.go.html to see an example of how to use TestBdsObjectStorageConnection API.
func (client BdsClient) TestBdsObjectStorageConnection(ctx context.Context, request TestBdsObjectStorageConnectionRequest) (response TestBdsObjectStorageConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.testBdsObjectStorageConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestBdsObjectStorageConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestBdsObjectStorageConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestBdsObjectStorageConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestBdsObjectStorageConnectionResponse")
	}
	return
}

// testBdsObjectStorageConnection implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) testBdsObjectStorageConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/apiKeys/{apiKeyId}/actions/testObjectStorageConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestBdsObjectStorageConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsApiKey/TestBdsObjectStorageConnection"
		err = common.PostProcessServiceError(err, "Bds", "TestBdsObjectStorageConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutoScalingConfiguration Updates fields on an autoscale configuration, including the name, the threshold value, and whether the autoscale configuration is enabled.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/UpdateAutoScalingConfiguration.go.html to see an example of how to use UpdateAutoScalingConfiguration API.
func (client BdsClient) UpdateAutoScalingConfiguration(ctx context.Context, request UpdateAutoScalingConfigurationRequest) (response UpdateAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateAutoScalingConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAutoScalingConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAutoScalingConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAutoScalingConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAutoScalingConfigurationResponse")
	}
	return
}

// updateAutoScalingConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) updateAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/UpdateAutoScalingConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "UpdateAutoScalingConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBdsInstance Updates the Big Data Service cluster identified by the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/UpdateBdsInstance.go.html to see an example of how to use UpdateBdsInstance API.
func (client BdsClient) UpdateBdsInstance(ctx context.Context, request UpdateBdsInstanceRequest) (response UpdateBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBdsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBdsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBdsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBdsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBdsInstanceResponse")
	}
	return
}

// updateBdsInstance implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) updateBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bdsInstances/{bdsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/UpdateBdsInstance"
		err = common.PostProcessServiceError(err, "Bds", "UpdateBdsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBdsMetastoreConfiguration Update the BDS metastore configuration represented by the provided ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/UpdateBdsMetastoreConfiguration.go.html to see an example of how to use UpdateBdsMetastoreConfiguration API.
func (client BdsClient) UpdateBdsMetastoreConfiguration(ctx context.Context, request UpdateBdsMetastoreConfigurationRequest) (response UpdateBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBdsMetastoreConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBdsMetastoreConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBdsMetastoreConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBdsMetastoreConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBdsMetastoreConfigurationResponse")
	}
	return
}

// updateBdsMetastoreConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) updateBdsMetastoreConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bdsInstances/{bdsInstanceId}/metastoreConfigs/{metastoreConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBdsMetastoreConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsMetastoreConfiguration/UpdateBdsMetastoreConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "UpdateBdsMetastoreConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
