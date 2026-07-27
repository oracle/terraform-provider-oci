// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DbClusterClient a client for DbCluster
type DbClusterClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbClusterClientWithConfigurationProvider Creates a new default DbCluster client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbClusterClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbClusterClient, err error) {
	if enabled := common.CheckForEnabledServices("mysql"); !enabled {
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
	return newDbClusterClientFromBaseClient(baseClient, provider)
}

// NewDbClusterClientWithOboToken Creates a new default DbCluster client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDbClusterClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbClusterClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbClusterClientFromBaseClient(baseClient, configProvider)
}

func newDbClusterClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbClusterClient, err error) {
	// DbCluster service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbCluster"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DbClusterClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbClusterClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbClusterClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbClusterClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDbClusterCompartment Moves a shared-storage DB cluster into a different compartment.
// When provided, If-Match is checked against the shared-storage DB cluster's ETag.
// A default retry strategy applies to this operation ChangeDbClusterCompartment()
func (client DbClusterClient) ChangeDbClusterCompartment(ctx context.Context, request ChangeDbClusterCompartmentRequest) (response ChangeDbClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDbClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDbClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDbClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDbClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDbClusterCompartmentResponse")
	}
	return
}

// changeDbClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) changeDbClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters/{dbClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDbClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "ChangeDbClusterCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/ChangeDbClusterCompartment"
		err = common.PostProcessServiceError(err, "DbCluster", "ChangeDbClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDbCluster Creates and launches a shared-storage DB cluster.
// A default retry strategy applies to this operation CreateDbCluster()
func (client DbClusterClient) CreateDbCluster(ctx context.Context, request CreateDbClusterRequest) (response CreateDbClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbClusterResponse")
	}
	return
}

// createDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) createDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "CreateDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbCluster", "CreateDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDbNode Adds a DB node to the specified shared-storage DB cluster.
// A default retry strategy applies to this operation CreateDbNode()
func (client DbClusterClient) CreateDbNode(ctx context.Context, request CreateDbNodeRequest) (response CreateDbNodeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbNodeResponse")
	}
	return
}

// createDbNode implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) createDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters/{dbClusterId}/dbNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "CreateDbNode")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbNode/CreateDbNode"
		err = common.PostProcessServiceError(err, "DbCluster", "CreateDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDbCluster Deletes a shared-storage DB cluster, including terminating, detaching,
// removing, finalizing and otherwise deleting all related resources.
// A default retry strategy applies to this operation DeleteDbCluster()
func (client DbClusterClient) DeleteDbCluster(ctx context.Context, request DeleteDbClusterRequest) (response DeleteDbClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbClusterResponse")
	}
	return
}

// deleteDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) deleteDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbClusters/{dbClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "DeleteDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/DeleteDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "DeleteDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDbNode A DB node cannot be deleted if it is the primary, set as an HA target, or delete protected.
// A default retry strategy applies to this operation DeleteDbNode()
func (client DbClusterClient) DeleteDbNode(ctx context.Context, request DeleteDbNodeRequest) (response DeleteDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbNodeResponse")
	}
	return
}

// deleteDbNode implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) deleteDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbClusters/{dbClusterId}/dbNodes/{dbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "DeleteDbNode")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbNode/DeleteDbNode"
		err = common.PostProcessServiceError(err, "DbCluster", "DeleteDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbCluster Gets information about the specified shared-storage DB cluster.
// A default retry strategy applies to this operation GetDbCluster()
func (client DbClusterClient) GetDbCluster(ctx context.Context, request GetDbClusterRequest) (response GetDbClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbClusterResponse")
	}
	return
}

// getDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) getDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbClusters/{dbClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "GetDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/GetDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "GetDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDbNode Gets information about the specified DB node.
// A default retry strategy applies to this operation GetDbNode()
func (client DbClusterClient) GetDbNode(ctx context.Context, request GetDbNodeRequest) (response GetDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbNodeResponse")
	}
	return
}

// getDbNode implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) getDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbClusters/{dbClusterId}/dbNodes/{dbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "GetDbNode")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbNode/GetDbNode"
		err = common.PostProcessServiceError(err, "DbCluster", "GetDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbClusters Gets a list of shared-storage DB clusters in the specified compartment.
// The default sort order is by timeUpdated, descending.
// A default retry strategy applies to this operation ListDbClusters()
func (client DbClusterClient) ListDbClusters(ctx context.Context, request ListDbClustersRequest) (response ListDbClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbClustersResponse")
	}
	return
}

// listDbClusters implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) listDbClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "ListDbClusters")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbClusterCollection/ListDbClusters"
		err = common.PostProcessServiceError(err, "DbCluster", "ListDbClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDbNodes Lists DB nodes in the specified shared-storage DB cluster.
// Default sort order is by timeUpdated, descending.
// A default retry strategy applies to this operation ListDbNodes()
func (client DbClusterClient) ListDbNodes(ctx context.Context, request ListDbNodesRequest) (response ListDbNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbNodesResponse")
	}
	return
}

// listDbNodes implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) listDbNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbClusters/{dbClusterId}/dbNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDbNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "ListDbNodes")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbNodeCollection/ListDbNodes"
		err = common.PostProcessServiceError(err, "DbCluster", "ListDbNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartDbCluster Restarts the specified shared-storage DB cluster.
// A default retry strategy applies to this operation RestartDbCluster()
func (client DbClusterClient) RestartDbCluster(ctx context.Context, request RestartDbClusterRequest) (response RestartDbClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restartDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartDbClusterResponse")
	}
	return
}

// restartDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) restartDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters/{dbClusterId}/actions/restart", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestartDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "RestartDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/RestartDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "RestartDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartDbCluster Starts the specified shared-storage DB cluster.
// A default retry strategy applies to this operation StartDbCluster()
func (client DbClusterClient) StartDbCluster(ctx context.Context, request StartDbClusterRequest) (response StartDbClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartDbClusterResponse")
	}
	return
}

// startDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) startDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters/{dbClusterId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "StartDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/StartDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "StartDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopDbCluster Stops the specified shared-storage DB cluster.
// A default retry strategy applies to this operation StopDbCluster()
func (client DbClusterClient) StopDbCluster(ctx context.Context, request StopDbClusterRequest) (response StopDbClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopDbClusterResponse")
	}
	return
}

// stopDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) stopDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters/{dbClusterId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "StopDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/StopDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "StopDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchoverDbCluster Promotes a secondary DB node to become the new primary DB node in the shared-storage DB cluster.
// A default retry strategy applies to this operation SwitchoverDbCluster()
func (client DbClusterClient) SwitchoverDbCluster(ctx context.Context, request SwitchoverDbClusterRequest) (response SwitchoverDbClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.switchoverDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchoverDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchoverDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverDbClusterResponse")
	}
	return
}

// switchoverDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) switchoverDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbClusters/{dbClusterId}/actions/switchover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchoverDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "SwitchoverDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/SwitchoverDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "SwitchoverDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbCluster Updates the shared-storage DB cluster settings.
// A default retry strategy applies to this operation UpdateDbCluster()
func (client DbClusterClient) UpdateDbCluster(ctx context.Context, request UpdateDbClusterRequest) (response UpdateDbClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbClusterResponse")
	}
	return
}

// updateDbCluster implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) updateDbCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbClusters/{dbClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "UpdateDbCluster")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbCluster/UpdateDbCluster"
		err = common.PostProcessServiceError(err, "DbCluster", "UpdateDbCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDbNode Updates the specified DB node. Only secondaries are updated.
// To update the role of a DB node, either update the shared-storage DB cluster or perform a primary switchover.
// A default retry strategy applies to this operation UpdateDbNode()
func (client DbClusterClient) UpdateDbNode(ctx context.Context, request UpdateDbNodeRequest) (response UpdateDbNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbNodeResponse")
	}
	return
}

// updateDbNode implements the OCIOperation interface (enables retrying operations)
func (client DbClusterClient) updateDbNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbClusters/{dbClusterId}/dbNodes/{dbNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDbNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dbCluster", "UpdateDbNode")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/DbNode/UpdateDbNode"
		err = common.PostProcessServiceError(err, "DbCluster", "UpdateDbNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
