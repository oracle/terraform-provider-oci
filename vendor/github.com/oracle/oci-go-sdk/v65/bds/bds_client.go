// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

//BdsClient a client for Bds
type BdsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBdsClientWithConfigurationProvider Creates a new default Bds client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBdsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BdsClient, err error) {
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
//  as well as reading the region
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
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *BdsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateBdsLakehouseConfiguration Activate the specified lakehouse configuration.
func (client BdsClient) ActivateBdsLakehouseConfiguration(ctx context.Context, request ActivateBdsLakehouseConfigurationRequest) (response ActivateBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.activateBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateBdsLakehouseConfigurationResponse")
	}
	return
}

// activateBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) activateBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs/{lakehouseConfigId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/ActivateBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "ActivateBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ActivateBdsMetastoreConfiguration Activate specified metastore configuration.
func (client BdsClient) ActivateBdsMetastoreConfiguration(ctx context.Context, request ActivateBdsMetastoreConfigurationRequest) (response ActivateBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) AddAutoScalingConfiguration(ctx context.Context, request AddAutoScalingConfigurationRequest) (response AddAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) AddBlockStorage(ctx context.Context, request AddBlockStorageRequest) (response AddBlockStorageResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) AddCloudSql(ctx context.Context, request AddCloudSqlRequest) (response AddCloudSqlResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// AddWorkerNodes Increases the size (scales out) a cluster by adding worker nodes(data/compute). The added worker nodes will have the same shape and will have the same amount of attached block storage as other worker nodes in the cluster.
func (client BdsClient) AddWorkerNodes(ctx context.Context, request AddWorkerNodesRequest) (response AddWorkerNodesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// ChangeBdsInstanceCompartment Moves a Big Data Service cluster into a different compartment.
func (client BdsClient) ChangeBdsInstanceCompartment(ctx context.Context, request ChangeBdsInstanceCompartmentRequest) (response ChangeBdsInstanceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ChangeShape(ctx context.Context, request ChangeShapeRequest) (response ChangeShapeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) CreateBdsApiKey(ctx context.Context, request CreateBdsApiKeyRequest) (response CreateBdsApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) CreateBdsInstance(ctx context.Context, request CreateBdsInstanceRequest) (response CreateBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// CreateBdsLakehouseConfiguration Create and activate the lakehouse configuration.
func (client BdsClient) CreateBdsLakehouseConfiguration(ctx context.Context, request CreateBdsLakehouseConfigurationRequest) (response CreateBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBdsLakehouseConfigurationResponse")
	}
	return
}

// createBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) createBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/CreateBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "CreateBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBdsMetastoreConfiguration Create and activate external metastore configuration.
func (client BdsClient) CreateBdsMetastoreConfiguration(ctx context.Context, request CreateBdsMetastoreConfigurationRequest) (response CreateBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// DeactivateBdsLakehouseConfiguration Deactivate the specified lakehouse configuration.
func (client BdsClient) DeactivateBdsLakehouseConfiguration(ctx context.Context, request DeactivateBdsLakehouseConfigurationRequest) (response DeactivateBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deactivateBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateBdsLakehouseConfigurationResponse")
	}
	return
}

// deactivateBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) deactivateBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs/{lakehouseConfigId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/DeactivateBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "DeactivateBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBdsApiKey Deletes the user's API key represented by the provided ID.
func (client BdsClient) DeleteBdsApiKey(ctx context.Context, request DeleteBdsApiKeyRequest) (response DeleteBdsApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) DeleteBdsInstance(ctx context.Context, request DeleteBdsInstanceRequest) (response DeleteBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// DeleteBdsLakehouseConfiguration Delete the BDS lakehouse configuration for the given ID.
func (client BdsClient) DeleteBdsLakehouseConfiguration(ctx context.Context, request DeleteBdsLakehouseConfigurationRequest) (response DeleteBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBdsLakehouseConfigurationResponse")
	}
	return
}

// deleteBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) deleteBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs/{lakehouseConfigId}/actions/delete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/DeleteBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "DeleteBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBdsMetastoreConfiguration Delete the BDS metastore configuration represented by the provided ID.
func (client BdsClient) DeleteBdsMetastoreConfiguration(ctx context.Context, request DeleteBdsMetastoreConfigurationRequest) (response DeleteBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// GetAutoScalingConfiguration Returns details of the autoscale configuration identified by the given ID.
func (client BdsClient) GetAutoScalingConfiguration(ctx context.Context, request GetAutoScalingConfigurationRequest) (response GetAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) GetBdsApiKey(ctx context.Context, request GetBdsApiKeyRequest) (response GetBdsApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) GetBdsInstance(ctx context.Context, request GetBdsInstanceRequest) (response GetBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// GetBdsLakehouseConfiguration Returns the BDS lakehouse configuration information for the given ID.
func (client BdsClient) GetBdsLakehouseConfiguration(ctx context.Context, request GetBdsLakehouseConfigurationRequest) (response GetBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBdsLakehouseConfigurationResponse")
	}
	return
}

// getBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) getBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs/{lakehouseConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/GetBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "GetBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBdsMetastoreConfiguration Returns the BDS Metastore configuration information for the given ID.
func (client BdsClient) GetBdsMetastoreConfiguration(ctx context.Context, request GetBdsMetastoreConfigurationRequest) (response GetBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// GetWorkRequest Returns the status of the work request identified by the given ID.
func (client BdsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// InstallPatch Install the specified patch to this cluster.
func (client BdsClient) InstallPatch(ctx context.Context, request InstallPatchRequest) (response InstallPatchResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListAutoScalingConfigurations(ctx context.Context, request ListAutoScalingConfigurationsRequest) (response ListAutoScalingConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListBdsApiKeys(ctx context.Context, request ListBdsApiKeysRequest) (response ListBdsApiKeysResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListBdsInstances(ctx context.Context, request ListBdsInstancesRequest) (response ListBdsInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// ListBdsLakehouseConfigurations Returns a list of lakehouse configurations associated with this Big Data Service cluster.
func (client BdsClient) ListBdsLakehouseConfigurations(ctx context.Context, request ListBdsLakehouseConfigurationsRequest) (response ListBdsLakehouseConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBdsLakehouseConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBdsLakehouseConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBdsLakehouseConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBdsLakehouseConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBdsLakehouseConfigurationsResponse")
	}
	return
}

// listBdsLakehouseConfigurations implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) listBdsLakehouseConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBdsLakehouseConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/ListBdsLakehouseConfigurations"
		err = common.PostProcessServiceError(err, "Bds", "ListBdsLakehouseConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBdsMetastoreConfigurations Returns a list of metastore configurations ssociated with this Big Data Service cluster.
func (client BdsClient) ListBdsMetastoreConfigurations(ctx context.Context, request ListBdsMetastoreConfigurationsRequest) (response ListBdsMetastoreConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// ListPatchHistories List the patch history of this cluster.
func (client BdsClient) ListPatchHistories(ctx context.Context, request ListPatchHistoriesRequest) (response ListPatchHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListPatches(ctx context.Context, request ListPatchesRequest) (response ListPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// RefreshDpCertificate Refreshes the specified cluster's DP certificate by replacing a new DP certificate for the cluster.
func (client BdsClient) RefreshDpCertificate(ctx context.Context, request RefreshDpCertificateRequest) (response RefreshDpCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.refreshDpCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshDpCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshDpCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshDpCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshDpCertificateResponse")
	}
	return
}

// refreshDpCertificate implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) refreshDpCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/refreshDpCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshDpCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsInstance/RefreshDpCertificate"
		err = common.PostProcessServiceError(err, "Bds", "RefreshDpCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAutoScalingConfiguration Deletes an autoscale configuration.
func (client BdsClient) RemoveAutoScalingConfiguration(ctx context.Context, request RemoveAutoScalingConfigurationRequest) (response RemoveAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) RemoveCloudSql(ctx context.Context, request RemoveCloudSqlRequest) (response RemoveCloudSqlResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// RemoveNode Remove a single node of a Big Data Service cluster
func (client BdsClient) RemoveNode(ctx context.Context, request RemoveNodeRequest) (response RemoveNodeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// RestartNode Restarts a single node of a Big Data Service cluster
func (client BdsClient) RestartNode(ctx context.Context, request RestartNodeRequest) (response RestartNodeResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// TestBdsLakehouseConfiguration Test the specified lakehouse configuration.
func (client BdsClient) TestBdsLakehouseConfiguration(ctx context.Context, request TestBdsLakehouseConfigurationRequest) (response TestBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.testBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestBdsLakehouseConfigurationResponse")
	}
	return
}

// testBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) testBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs/{lakehouseConfigId}/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/TestBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "TestBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestBdsMetastoreConfiguration Test specified metastore configuration.
func (client BdsClient) TestBdsMetastoreConfiguration(ctx context.Context, request TestBdsMetastoreConfigurationRequest) (response TestBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) TestBdsObjectStorageConnection(ctx context.Context, request TestBdsObjectStorageConnectionRequest) (response TestBdsObjectStorageConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) UpdateAutoScalingConfiguration(ctx context.Context, request UpdateAutoScalingConfigurationRequest) (response UpdateAutoScalingConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
func (client BdsClient) UpdateBdsInstance(ctx context.Context, request UpdateBdsInstanceRequest) (response UpdateBdsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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

// UpdateBdsLakehouseConfiguration Update the BDS lakehouse configuration.
func (client BdsClient) UpdateBdsLakehouseConfiguration(ctx context.Context, request UpdateBdsLakehouseConfigurationRequest) (response UpdateBdsLakehouseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBdsLakehouseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBdsLakehouseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBdsLakehouseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBdsLakehouseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBdsLakehouseConfigurationResponse")
	}
	return
}

// updateBdsLakehouseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client BdsClient) updateBdsLakehouseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bdsInstances/{bdsInstanceId}/lakehouseConfigs/{lakehouseConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBdsLakehouseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/bigdata/20190531/BdsLakehouseConfiguration/UpdateBdsLakehouseConfiguration"
		err = common.PostProcessServiceError(err, "Bds", "UpdateBdsLakehouseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBdsMetastoreConfiguration Update the BDS metastore configuration represented by the provided ID.
func (client BdsClient) UpdateBdsMetastoreConfiguration(ctx context.Context, request UpdateBdsMetastoreConfigurationRequest) (response UpdateBdsMetastoreConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
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
