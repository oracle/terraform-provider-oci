// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v41/common"
	"github.com/oracle/oci-go-sdk/v41/common/auth"
	"net/http"
)

//DbSystemClient a client for DbSystem
type DbSystemClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbSystemClientWithConfigurationProvider Creates a new default DbSystem client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbSystemClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbSystemClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDbSystemClientFromBaseClient(baseClient, provider)
}

// NewDbSystemClientWithOboToken Creates a new default DbSystem client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDbSystemClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbSystemClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbSystemClientFromBaseClient(baseClient, configProvider)
}

func newDbSystemClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbSystemClient, err error) {
	client = DbSystemClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbSystemClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbSystemClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbSystemClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Adds an Analytics Cluster to the DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/AddAnalyticsCluster.go.html to see an example of how to use AddAnalyticsCluster API.
func (client DbSystemClient) AddAnalyticsCluster(ctx context.Context, request AddAnalyticsClusterRequest) (response AddAnalyticsClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddAnalyticsClusterResponse")
	}
	return
}

// addAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) addAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/analyticsCluster/actions/add", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AddAnalyticsClusterResponse
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

// AddHeatWaveCluster Adds a HeatWave cluster to the DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/AddHeatWaveCluster.go.html to see an example of how to use AddHeatWaveCluster API.
func (client DbSystemClient) AddHeatWaveCluster(ctx context.Context, request AddHeatWaveClusterRequest) (response AddHeatWaveClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddHeatWaveClusterResponse")
	}
	return
}

// addHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) addHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/heatWaveCluster/actions/add", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response AddHeatWaveClusterResponse
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

// CreateDbSystem Creates and launches a DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/CreateDbSystem.go.html to see an example of how to use CreateDbSystem API.
func (client DbSystemClient) CreateDbSystem(ctx context.Context, request CreateDbSystemRequest) (response CreateDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDbSystemResponse")
	}
	return
}

// createDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) createDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDbSystemResponse
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

// DeleteAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Deletes the Analytics Cluster including terminating, detaching, removing, finalizing and
// otherwise deleting all related resources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/DeleteAnalyticsCluster.go.html to see an example of how to use DeleteAnalyticsCluster API.
func (client DbSystemClient) DeleteAnalyticsCluster(ctx context.Context, request DeleteAnalyticsClusterRequest) (response DeleteAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAnalyticsClusterResponse")
	}
	return
}

// deleteAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) deleteAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbSystems/{dbSystemId}/analyticsCluster", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteAnalyticsClusterResponse
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

// DeleteDbSystem Delete a DB System, including terminating, detaching,
// removing, finalizing and otherwise deleting all related resources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/DeleteDbSystem.go.html to see an example of how to use DeleteDbSystem API.
func (client DbSystemClient) DeleteDbSystem(ctx context.Context, request DeleteDbSystemRequest) (response DeleteDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDbSystemResponse")
	}
	return
}

// deleteDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) deleteDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbSystems/{dbSystemId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDbSystemResponse
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

// DeleteHeatWaveCluster Deletes the HeatWave cluster including terminating, detaching, removing, finalizing and
// otherwise deleting all related resources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/DeleteHeatWaveCluster.go.html to see an example of how to use DeleteHeatWaveCluster API.
func (client DbSystemClient) DeleteHeatWaveCluster(ctx context.Context, request DeleteHeatWaveClusterRequest) (response DeleteHeatWaveClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteHeatWaveClusterResponse")
	}
	return
}

// deleteHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) deleteHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dbSystems/{dbSystemId}/heatWaveCluster", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteHeatWaveClusterResponse
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

// GenerateAnalyticsClusterMemoryEstimate DEPRECATED -- please use HeatWave API instead.
// Sends a request to estimate the memory footprints of user tables when loaded to Analytics Cluster memory.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GenerateAnalyticsClusterMemoryEstimate.go.html to see an example of how to use GenerateAnalyticsClusterMemoryEstimate API.
func (client DbSystemClient) GenerateAnalyticsClusterMemoryEstimate(ctx context.Context, request GenerateAnalyticsClusterMemoryEstimateRequest) (response GenerateAnalyticsClusterMemoryEstimateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateAnalyticsClusterMemoryEstimate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateAnalyticsClusterMemoryEstimateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateAnalyticsClusterMemoryEstimateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateAnalyticsClusterMemoryEstimateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateAnalyticsClusterMemoryEstimateResponse")
	}
	return
}

// generateAnalyticsClusterMemoryEstimate implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) generateAnalyticsClusterMemoryEstimate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/analyticsClusterMemoryEstimate/actions/generate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GenerateAnalyticsClusterMemoryEstimateResponse
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

// GenerateHeatWaveClusterMemoryEstimate Sends a request to estimate the memory footprints of user tables when loaded to HeatWave cluster memory.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GenerateHeatWaveClusterMemoryEstimate.go.html to see an example of how to use GenerateHeatWaveClusterMemoryEstimate API.
func (client DbSystemClient) GenerateHeatWaveClusterMemoryEstimate(ctx context.Context, request GenerateHeatWaveClusterMemoryEstimateRequest) (response GenerateHeatWaveClusterMemoryEstimateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateHeatWaveClusterMemoryEstimate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateHeatWaveClusterMemoryEstimateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateHeatWaveClusterMemoryEstimateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateHeatWaveClusterMemoryEstimateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateHeatWaveClusterMemoryEstimateResponse")
	}
	return
}

// generateHeatWaveClusterMemoryEstimate implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) generateHeatWaveClusterMemoryEstimate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/heatWaveClusterMemoryEstimate/actions/generate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GenerateHeatWaveClusterMemoryEstimateResponse
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

// GetAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Gets information about the Analytics Cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetAnalyticsCluster.go.html to see an example of how to use GetAnalyticsCluster API.
func (client DbSystemClient) GetAnalyticsCluster(ctx context.Context, request GetAnalyticsClusterRequest) (response GetAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnalyticsClusterResponse")
	}
	return
}

// getAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) getAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/analyticsCluster", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAnalyticsClusterResponse
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

// GetAnalyticsClusterMemoryEstimate DEPRECATED -- please use HeatWave API instead.
// Gets the most recent Analytics Cluster memory estimate that can be used to determine a suitable
// Analytics Cluster size.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetAnalyticsClusterMemoryEstimate.go.html to see an example of how to use GetAnalyticsClusterMemoryEstimate API.
func (client DbSystemClient) GetAnalyticsClusterMemoryEstimate(ctx context.Context, request GetAnalyticsClusterMemoryEstimateRequest) (response GetAnalyticsClusterMemoryEstimateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnalyticsClusterMemoryEstimate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnalyticsClusterMemoryEstimateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnalyticsClusterMemoryEstimateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnalyticsClusterMemoryEstimateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnalyticsClusterMemoryEstimateResponse")
	}
	return
}

// getAnalyticsClusterMemoryEstimate implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) getAnalyticsClusterMemoryEstimate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/analyticsClusterMemoryEstimate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetAnalyticsClusterMemoryEstimateResponse
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

// GetDbSystem Get information about the specified DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetDbSystem.go.html to see an example of how to use GetDbSystem API.
func (client DbSystemClient) GetDbSystem(ctx context.Context, request GetDbSystemRequest) (response GetDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDbSystemResponse")
	}
	return
}

// getDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) getDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDbSystemResponse
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

// GetHeatWaveCluster Gets information about the HeatWave cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetHeatWaveCluster.go.html to see an example of how to use GetHeatWaveCluster API.
func (client DbSystemClient) GetHeatWaveCluster(ctx context.Context, request GetHeatWaveClusterRequest) (response GetHeatWaveClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHeatWaveClusterResponse")
	}
	return
}

// getHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) getHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/heatWaveCluster", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetHeatWaveClusterResponse
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

// GetHeatWaveClusterMemoryEstimate Gets the most recent HeatWave cluster memory estimate that can be used to determine a suitable
// HeatWave cluster size.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetHeatWaveClusterMemoryEstimate.go.html to see an example of how to use GetHeatWaveClusterMemoryEstimate API.
func (client DbSystemClient) GetHeatWaveClusterMemoryEstimate(ctx context.Context, request GetHeatWaveClusterMemoryEstimateRequest) (response GetHeatWaveClusterMemoryEstimateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHeatWaveClusterMemoryEstimate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHeatWaveClusterMemoryEstimateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHeatWaveClusterMemoryEstimateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHeatWaveClusterMemoryEstimateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHeatWaveClusterMemoryEstimateResponse")
	}
	return
}

// getHeatWaveClusterMemoryEstimate implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) getHeatWaveClusterMemoryEstimate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems/{dbSystemId}/heatWaveClusterMemoryEstimate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetHeatWaveClusterMemoryEstimateResponse
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

// ListDbSystems Get a list of DB Systems in the specified compartment.
// The default sort order is by timeUpdated, descending.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListDbSystems.go.html to see an example of how to use ListDbSystems API.
func (client DbSystemClient) ListDbSystems(ctx context.Context, request ListDbSystemsRequest) (response ListDbSystemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDbSystems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDbSystemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDbSystemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDbSystemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDbSystemsResponse")
	}
	return
}

// listDbSystems implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) listDbSystems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dbSystems", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDbSystemsResponse
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

// RestartAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Restarts the Analytics Cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/RestartAnalyticsCluster.go.html to see an example of how to use RestartAnalyticsCluster API.
func (client DbSystemClient) RestartAnalyticsCluster(ctx context.Context, request RestartAnalyticsClusterRequest) (response RestartAnalyticsClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restartAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartAnalyticsClusterResponse")
	}
	return
}

// restartAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) restartAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/analyticsCluster/actions/restart", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestartAnalyticsClusterResponse
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

// RestartDbSystem Restarts the specified DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/RestartDbSystem.go.html to see an example of how to use RestartDbSystem API.
func (client DbSystemClient) RestartDbSystem(ctx context.Context, request RestartDbSystemRequest) (response RestartDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restartDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartDbSystemResponse")
	}
	return
}

// restartDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) restartDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/restart", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestartDbSystemResponse
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

// RestartHeatWaveCluster Restarts the HeatWave cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/RestartHeatWaveCluster.go.html to see an example of how to use RestartHeatWaveCluster API.
func (client DbSystemClient) RestartHeatWaveCluster(ctx context.Context, request RestartHeatWaveClusterRequest) (response RestartHeatWaveClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restartHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartHeatWaveClusterResponse")
	}
	return
}

// restartHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) restartHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/heatWaveCluster/actions/restart", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response RestartHeatWaveClusterResponse
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

// StartAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Starts the Analytics Cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/StartAnalyticsCluster.go.html to see an example of how to use StartAnalyticsCluster API.
func (client DbSystemClient) StartAnalyticsCluster(ctx context.Context, request StartAnalyticsClusterRequest) (response StartAnalyticsClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAnalyticsClusterResponse")
	}
	return
}

// startAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) startAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/analyticsCluster/actions/start", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StartAnalyticsClusterResponse
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

// StartDbSystem Start the specified DB System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/StartDbSystem.go.html to see an example of how to use StartDbSystem API.
func (client DbSystemClient) StartDbSystem(ctx context.Context, request StartDbSystemRequest) (response StartDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartDbSystemResponse")
	}
	return
}

// startDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) startDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/start", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StartDbSystemResponse
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

// StartHeatWaveCluster Starts the HeatWave cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/StartHeatWaveCluster.go.html to see an example of how to use StartHeatWaveCluster API.
func (client DbSystemClient) StartHeatWaveCluster(ctx context.Context, request StartHeatWaveClusterRequest) (response StartHeatWaveClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartHeatWaveClusterResponse")
	}
	return
}

// startHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) startHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/heatWaveCluster/actions/start", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StartHeatWaveClusterResponse
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

// StopAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Stops the Analytics Cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/StopAnalyticsCluster.go.html to see an example of how to use StopAnalyticsCluster API.
func (client DbSystemClient) StopAnalyticsCluster(ctx context.Context, request StopAnalyticsClusterRequest) (response StopAnalyticsClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAnalyticsClusterResponse")
	}
	return
}

// stopAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) stopAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/analyticsCluster/actions/stop", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StopAnalyticsClusterResponse
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

// StopDbSystem Stops the specified DB System.
// A stopped DB System is not billed.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/StopDbSystem.go.html to see an example of how to use StopDbSystem API.
func (client DbSystemClient) StopDbSystem(ctx context.Context, request StopDbSystemRequest) (response StopDbSystemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopDbSystemResponse")
	}
	return
}

// stopDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) stopDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/actions/stop", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StopDbSystemResponse
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

// StopHeatWaveCluster Stops the HeatWave cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/StopHeatWaveCluster.go.html to see an example of how to use StopHeatWaveCluster API.
func (client DbSystemClient) StopHeatWaveCluster(ctx context.Context, request StopHeatWaveClusterRequest) (response StopHeatWaveClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopHeatWaveClusterResponse")
	}
	return
}

// stopHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) stopHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dbSystems/{dbSystemId}/heatWaveCluster/actions/stop", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response StopHeatWaveClusterResponse
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

// UpdateAnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// Updates the Analytics Cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/UpdateAnalyticsCluster.go.html to see an example of how to use UpdateAnalyticsCluster API.
func (client DbSystemClient) UpdateAnalyticsCluster(ctx context.Context, request UpdateAnalyticsClusterRequest) (response UpdateAnalyticsClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnalyticsCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnalyticsClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnalyticsClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnalyticsClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnalyticsClusterResponse")
	}
	return
}

// updateAnalyticsCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) updateAnalyticsCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}/analyticsCluster", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateAnalyticsClusterResponse
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

// UpdateDbSystem Update the configuration of a DB System.
// Updating different fields in the DB System will have different results
// on the uptime of the DB System. For example, changing the displayName of
// a DB System will take effect immediately, but changing the shape of a
// DB System is an asynchronous operation that involves provisioning new
// Compute resources, pausing the DB System and migrating storage
// before making the DB System available again.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/UpdateDbSystem.go.html to see an example of how to use UpdateDbSystem API.
func (client DbSystemClient) UpdateDbSystem(ctx context.Context, request UpdateDbSystemRequest) (response UpdateDbSystemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDbSystem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDbSystemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDbSystemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDbSystemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDbSystemResponse")
	}
	return
}

// updateDbSystem implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) updateDbSystem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDbSystemResponse
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

// UpdateHeatWaveCluster Updates the HeatWave cluster.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/UpdateHeatWaveCluster.go.html to see an example of how to use UpdateHeatWaveCluster API.
func (client DbSystemClient) UpdateHeatWaveCluster(ctx context.Context, request UpdateHeatWaveClusterRequest) (response UpdateHeatWaveClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateHeatWaveCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHeatWaveClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHeatWaveClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHeatWaveClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHeatWaveClusterResponse")
	}
	return
}

// updateHeatWaveCluster implements the OCIOperation interface (enables retrying operations)
func (client DbSystemClient) updateHeatWaveCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dbSystems/{dbSystemId}/heatWaveCluster", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateHeatWaveClusterResponse
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
