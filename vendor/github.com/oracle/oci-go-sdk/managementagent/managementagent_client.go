// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ManagementAgentClient a client for ManagementAgent
type ManagementAgentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagementAgentClientWithConfigurationProvider Creates a new default ManagementAgent client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagementAgentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagementAgentClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newManagementAgentClientFromBaseClient(baseClient, configProvider)
}

// NewManagementAgentClientWithOboToken Creates a new default ManagementAgent client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewManagementAgentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagementAgentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newManagementAgentClientFromBaseClient(baseClient, configProvider)
}

func newManagementAgentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagementAgentClient, err error) {
	client = ManagementAgentClient{BaseClient: baseClient}
	client.BasePath = "20200202"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagementAgentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("managementagent", "https://management-agent.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagementAgentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ManagementAgentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateManagementAgentInstallKey User creates a new install key as part of this API.
func (client ManagementAgentClient) CreateManagementAgentInstallKey(ctx context.Context, request CreateManagementAgentInstallKeyRequest) (response CreateManagementAgentInstallKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createManagementAgentInstallKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementAgentInstallKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementAgentInstallKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementAgentInstallKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementAgentInstallKeyResponse")
	}
	return
}

// createManagementAgentInstallKey implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) createManagementAgentInstallKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementAgentInstallKeys")
	if err != nil {
		return nil, err
	}

	var response CreateManagementAgentInstallKeyResponse
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

// DeleteManagementAgent Deletes a Management Agent resource by identifier
func (client ManagementAgentClient) DeleteManagementAgent(ctx context.Context, request DeleteManagementAgentRequest) (response DeleteManagementAgentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementAgentResponse")
	}
	return
}

// deleteManagementAgent implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) deleteManagementAgent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementAgents/{managementAgentId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagementAgentResponse
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

// DeleteManagementAgentInstallKey Deletes a Management Agent install Key resource by identifier
func (client ManagementAgentClient) DeleteManagementAgentInstallKey(ctx context.Context, request DeleteManagementAgentInstallKeyRequest) (response DeleteManagementAgentInstallKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementAgentInstallKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementAgentInstallKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementAgentInstallKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementAgentInstallKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementAgentInstallKeyResponse")
	}
	return
}

// deleteManagementAgentInstallKey implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) deleteManagementAgentInstallKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementAgentInstallKeys/{managementAgentInstallKeyId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagementAgentInstallKeyResponse
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

// DeleteWorkRequest Cancel the work request with the given ID.
func (client ManagementAgentClient) DeleteWorkRequest(ctx context.Context, request DeleteWorkRequestRequest) (response DeleteWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkRequestResponse")
	}
	return
}

// deleteWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) deleteWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response DeleteWorkRequestResponse
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

// DeployPlugins Deploys Plugins to a given list of agentIds.
func (client ManagementAgentClient) DeployPlugins(ctx context.Context, request DeployPluginsRequest) (response DeployPluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deployPlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeployPluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeployPluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeployPluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeployPluginsResponse")
	}
	return
}

// deployPlugins implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) deployPlugins(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementAgents/actions/deployPlugins")
	if err != nil {
		return nil, err
	}

	var response DeployPluginsResponse
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

// GetManagementAgent Gets complete details of the inventory of a given agent id
func (client ManagementAgentClient) GetManagementAgent(ctx context.Context, request GetManagementAgentRequest) (response GetManagementAgentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagementAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementAgentResponse")
	}
	return
}

// getManagementAgent implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) getManagementAgent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgents/{managementAgentId}")
	if err != nil {
		return nil, err
	}

	var response GetManagementAgentResponse
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

// GetManagementAgentInstallKey Gets complete details of the Agent install Key for a given key id
func (client ManagementAgentClient) GetManagementAgentInstallKey(ctx context.Context, request GetManagementAgentInstallKeyRequest) (response GetManagementAgentInstallKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagementAgentInstallKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementAgentInstallKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementAgentInstallKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementAgentInstallKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementAgentInstallKeyResponse")
	}
	return
}

// getManagementAgentInstallKey implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) getManagementAgentInstallKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgentInstallKeys/{managementAgentInstallKeyId}")
	if err != nil {
		return nil, err
	}

	var response GetManagementAgentInstallKeyResponse
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

// GetManagementAgentInstallKeyContent Returns a file with Management Agent install Key in it
func (client ManagementAgentClient) GetManagementAgentInstallKeyContent(ctx context.Context, request GetManagementAgentInstallKeyContentRequest) (response GetManagementAgentInstallKeyContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagementAgentInstallKeyContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementAgentInstallKeyContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementAgentInstallKeyContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementAgentInstallKeyContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementAgentInstallKeyContentResponse")
	}
	return
}

// getManagementAgentInstallKeyContent implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) getManagementAgentInstallKeyContent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgentInstallKeys/{managementAgentInstallKeyId}/content")
	if err != nil {
		return nil, err
	}

	var response GetManagementAgentInstallKeyContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
func (client ManagementAgentClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ManagementAgentClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
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

// ListManagementAgentImages Get supported agent image information
func (client ManagementAgentClient) ListManagementAgentImages(ctx context.Context, request ListManagementAgentImagesRequest) (response ListManagementAgentImagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.listManagementAgentImages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementAgentImagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementAgentImagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementAgentImagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementAgentImagesResponse")
	}
	return
}

// listManagementAgentImages implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) listManagementAgentImages(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgentImages")
	if err != nil {
		return nil, err
	}

	var response ListManagementAgentImagesResponse
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

// ListManagementAgentInstallKeys Returns a list of Management Agent installed Keys.
func (client ManagementAgentClient) ListManagementAgentInstallKeys(ctx context.Context, request ListManagementAgentInstallKeysRequest) (response ListManagementAgentInstallKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementAgentInstallKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementAgentInstallKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementAgentInstallKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementAgentInstallKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementAgentInstallKeysResponse")
	}
	return
}

// listManagementAgentInstallKeys implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) listManagementAgentInstallKeys(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgentInstallKeys")
	if err != nil {
		return nil, err
	}

	var response ListManagementAgentInstallKeysResponse
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

// ListManagementAgentPlugins Returns a list of managementAgentPlugins.
func (client ManagementAgentClient) ListManagementAgentPlugins(ctx context.Context, request ListManagementAgentPluginsRequest) (response ListManagementAgentPluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementAgentPlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementAgentPluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementAgentPluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementAgentPluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementAgentPluginsResponse")
	}
	return
}

// listManagementAgentPlugins implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) listManagementAgentPlugins(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgentPlugins")
	if err != nil {
		return nil, err
	}

	var response ListManagementAgentPluginsResponse
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

// ListManagementAgents Returns a list of Management Agent.
func (client ManagementAgentClient) ListManagementAgents(ctx context.Context, request ListManagementAgentsRequest) (response ListManagementAgentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementAgents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementAgentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementAgentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementAgentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementAgentsResponse")
	}
	return
}

// listManagementAgents implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) listManagementAgents(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAgents")
	if err != nil {
		return nil, err
	}

	var response ListManagementAgentsResponse
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
func (client ManagementAgentClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ManagementAgentClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
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

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
func (client ManagementAgentClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ManagementAgentClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
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
func (client ManagementAgentClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ManagementAgentClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
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

// UpdateManagementAgent API to update the console managed properties of the Management Agent.
func (client ManagementAgentClient) UpdateManagementAgent(ctx context.Context, request UpdateManagementAgentRequest) (response UpdateManagementAgentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateManagementAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementAgentResponse")
	}
	return
}

// updateManagementAgent implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) updateManagementAgent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementAgents/{managementAgentId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagementAgentResponse
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

// UpdateManagementAgentInstallKey API to update the modifiable properties of the Management Agent install key.
func (client ManagementAgentClient) UpdateManagementAgentInstallKey(ctx context.Context, request UpdateManagementAgentInstallKeyRequest) (response UpdateManagementAgentInstallKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateManagementAgentInstallKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementAgentInstallKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementAgentInstallKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementAgentInstallKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementAgentInstallKeyResponse")
	}
	return
}

// updateManagementAgentInstallKey implements the OCIOperation interface (enables retrying operations)
func (client ManagementAgentClient) updateManagementAgentInstallKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementAgentInstallKeys/{managementAgentInstallKeyId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagementAgentInstallKeyResponse
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
