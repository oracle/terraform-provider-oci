// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v43/common"
	"github.com/oracle/oci-go-sdk/v43/common/auth"
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *BdsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddAutoScalingConfiguration Add an autoscale configuration to the cluster.
//
// See also
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
func (client BdsClient) addAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AddAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddBlockStorage Adds block storage to existing worker nodes. The same amount of  storage will be added to all worker nodes. No change will be made  to storage that is already attached. Block storage cannot be removed.
//
// See also
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
func (client BdsClient) addBlockStorage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addBlockStorage", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AddBlockStorageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddCloudSql Adds Cloud SQL to your cluster. You can use Cloud SQL to query against non-relational data stored in multiple big data sources, including Apache Hive, HDFS, Oracle NoSQL Database, and Apache HBase. Adding Cloud SQL adds a query server node to the cluster and creates cell servers on all the worker nodes in the cluster.
//
// See also
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
func (client BdsClient) addCloudSql(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addCloudSql", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AddCloudSqlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddWorkerNodes Increases the size (scales out) a cluster by adding worker nodes. The added worker nodes will have the same shape and will have the same amount of attached block storage as other worker nodes in the cluster.
//
// See also
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
func (client BdsClient) addWorkerNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/addWorkerNodes", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AddWorkerNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeBdsInstanceCompartment Moves a Big Data Service cluster into a different compartment.
//
// See also
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
func (client BdsClient) changeBdsInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeBdsInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeShape Changes the size of a cluster by scaling up or scaling down the nodes. Nodes are scaled up or down by changing the shapes of all the nodes of the same type to the next larger or smaller shape. The node types are master, utility, worker, and Cloud SQL. Only nodes with VM-STANDARD shapes can be scaled.
//
// See also
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
func (client BdsClient) changeShape(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/changeShape", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeShapeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBdsInstance Creates a Big Data Service cluster.
//
// See also
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
func (client BdsClient) createBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBdsInstance Deletes the cluster identified by the given ID.
//
// See also
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
func (client BdsClient) deleteBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/bdsInstances/{bdsInstanceId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAutoScalingConfiguration Returns details of the autoscale configuration identified by the given ID.
//
// See also
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
func (client BdsClient) getAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBdsInstance Returns information about the Big Data Service cluster identified by the given ID.
//
// See also
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
func (client BdsClient) getBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Returns the status of the work request identified by the given ID.
//
// See also
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
func (client BdsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAutoScalingConfigurations Returns information about the autoscaling configurations for a cluster.
//
// See also
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
func (client BdsClient) listAutoScalingConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListAutoScalingConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBdsInstances Returns a list of all Big Data Service clusters in a compartment.
//
// See also
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
func (client BdsClient) listBdsInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/bdsInstances", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListBdsInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a paginated list of errors for a work request identified by the given ID.
//
// See also
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
func (client BdsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for a given work request.
//
// See also
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
func (client BdsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// See also
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
func (client BdsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAutoScalingConfiguration Deletes an autoscale configuration.
//
// See also
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
func (client BdsClient) removeAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}/actions/remove", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RemoveAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveCloudSql Removes Cloud SQL from the cluster.
//
// See also
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
func (client BdsClient) removeCloudSql(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/removeCloudSql", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RemoveCloudSqlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartNode Restarts a single node of a Big Data Service cluster
//
// See also
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
func (client BdsClient) restartNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/bdsInstances/{bdsInstanceId}/actions/restartNode", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestartNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAutoScalingConfiguration Updates fields on an autoscale configuration, including the name, the threshold value, and whether the autoscale configuration is enabled.
//
// See also
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
func (client BdsClient) updateAutoScalingConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bdsInstances/{bdsInstanceId}/autoScalingConfiguration/{autoScalingConfigurationId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAutoScalingConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBdsInstance Updates the Big Data Service cluster identified by the given ID.
//
// See also
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
func (client BdsClient) updateBdsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/bdsInstances/{bdsInstanceId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateBdsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
