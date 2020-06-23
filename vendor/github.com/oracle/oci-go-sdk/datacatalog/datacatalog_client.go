// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DataCatalogClient a client for DataCatalog
type DataCatalogClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataCatalogClientWithConfigurationProvider Creates a new default DataCatalog client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataCatalogClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataCatalogClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newDataCatalogClientFromBaseClient(baseClient, configProvider)
}

// NewDataCatalogClientWithOboToken Creates a new default DataCatalog client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataCatalogClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataCatalogClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newDataCatalogClientFromBaseClient(baseClient, configProvider)
}

func newDataCatalogClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataCatalogClient, err error) {
	client = DataCatalogClient{BaseClient: baseClient}
	client.BasePath = "20190325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataCatalogClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datacatalog", "https://datacatalog.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataCatalogClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataCatalogClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachCatalogPrivateEndpoint Attaches a private reverse connection endpoint resource to a data catalog resource. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client DataCatalogClient) AttachCatalogPrivateEndpoint(ctx context.Context, request AttachCatalogPrivateEndpointRequest) (response AttachCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.attachCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachCatalogPrivateEndpointResponse")
	}
	return
}

// attachCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) attachCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/attachCatalogPrivateEndpoint")
	if err != nil {
		return nil, err
	}

	var response AttachCatalogPrivateEndpointResponse
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

// ChangeCatalogCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client DataCatalogClient) ChangeCatalogCompartment(ctx context.Context, request ChangeCatalogCompartmentRequest) (response ChangeCatalogCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeCatalogCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCatalogCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCatalogCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCatalogCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCatalogCompartmentResponse")
	}
	return
}

// changeCatalogCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) changeCatalogCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeCatalogCompartmentResponse
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

// ChangeCatalogPrivateEndpointCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client DataCatalogClient) ChangeCatalogPrivateEndpointCompartment(ctx context.Context, request ChangeCatalogPrivateEndpointCompartmentRequest) (response ChangeCatalogPrivateEndpointCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeCatalogPrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCatalogPrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCatalogPrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCatalogPrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCatalogPrivateEndpointCompartmentResponse")
	}
	return
}

// changeCatalogPrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) changeCatalogPrivateEndpointCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeCatalogPrivateEndpointCompartmentResponse
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

// CreateAttribute Creates a new entity attribute.
func (client DataCatalogClient) CreateAttribute(ctx context.Context, request CreateAttributeRequest) (response CreateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAttributeResponse")
	}
	return
}

// createAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes")
	if err != nil {
		return nil, err
	}

	var response CreateAttributeResponse
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

// CreateAttributeTag Creates a new entity attribute tag.
func (client DataCatalogClient) CreateAttributeTag(ctx context.Context, request CreateAttributeTagRequest) (response CreateAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAttributeTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAttributeTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAttributeTagResponse")
	}
	return
}

// createAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createAttributeTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateAttributeTagResponse
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

// CreateCatalog Creates a new data catalog instance that includes a console and an API URL for managing metadata operations.
// For more information, please see the documentation.
func (client DataCatalogClient) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (response CreateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogResponse")
	}
	return
}

// createCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs")
	if err != nil {
		return nil, err
	}

	var response CreateCatalogResponse
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

// CreateCatalogPrivateEndpoint Create a new private reverse connection endpoint.
func (client DataCatalogClient) CreateCatalogPrivateEndpoint(ctx context.Context, request CreateCatalogPrivateEndpointRequest) (response CreateCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogPrivateEndpointResponse")
	}
	return
}

// createCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogPrivateEndpoints")
	if err != nil {
		return nil, err
	}

	var response CreateCatalogPrivateEndpointResponse
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

// CreateConnection Creates a new connection.
func (client DataCatalogClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionResponse")
	}
	return
}

// createConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections")
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
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

// CreateDataAsset Create a new data asset.
func (client DataCatalogClient) CreateDataAsset(ctx context.Context, request CreateDataAssetRequest) (response CreateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetResponse")
	}
	return
}

// createDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets")
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetResponse
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

// CreateDataAssetTag Creates a new data asset tag.
func (client DataCatalogClient) CreateDataAssetTag(ctx context.Context, request CreateDataAssetTagRequest) (response CreateDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataAssetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataAssetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetTagResponse")
	}
	return
}

// createDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createDataAssetTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetTagResponse
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

// CreateEntity Creates a new data entity.
func (client DataCatalogClient) CreateEntity(ctx context.Context, request CreateEntityRequest) (response CreateEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityResponse")
	}
	return
}

// createEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities")
	if err != nil {
		return nil, err
	}

	var response CreateEntityResponse
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

// CreateEntityTag Creates a new entity tag.
func (client DataCatalogClient) CreateEntityTag(ctx context.Context, request CreateEntityTagRequest) (response CreateEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEntityTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEntityTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityTagResponse")
	}
	return
}

// createEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createEntityTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateEntityTagResponse
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

// CreateFolder Creates a new folder.
func (client DataCatalogClient) CreateFolder(ctx context.Context, request CreateFolderRequest) (response CreateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderResponse")
	}
	return
}

// createFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders")
	if err != nil {
		return nil, err
	}

	var response CreateFolderResponse
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

// CreateFolderTag Creates a new folder tag.
func (client DataCatalogClient) CreateFolderTag(ctx context.Context, request CreateFolderTagRequest) (response CreateFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFolderTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFolderTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderTagResponse")
	}
	return
}

// createFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createFolderTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags")
	if err != nil {
		return nil, err
	}

	var response CreateFolderTagResponse
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

// CreateGlossary Creates a new glossary.
func (client DataCatalogClient) CreateGlossary(ctx context.Context, request CreateGlossaryRequest) (response CreateGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGlossaryResponse")
	}
	return
}

// createGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries")
	if err != nil {
		return nil, err
	}

	var response CreateGlossaryResponse
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

// CreateJob Creates a new job.
func (client DataCatalogClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobResponse")
	}
	return
}

// createJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobs")
	if err != nil {
		return nil, err
	}

	var response CreateJobResponse
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

// CreateJobDefinition Creates a new job definition.
func (client DataCatalogClient) CreateJobDefinition(ctx context.Context, request CreateJobDefinitionRequest) (response CreateJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobDefinitionResponse")
	}
	return
}

// createJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobDefinitions")
	if err != nil {
		return nil, err
	}

	var response CreateJobDefinitionResponse
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

// CreateJobExecution Creates a new job execution.
func (client DataCatalogClient) CreateJobExecution(ctx context.Context, request CreateJobExecutionRequest) (response CreateJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobExecutionResponse")
	}
	return
}

// createJobExecution implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJobExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobs/{jobKey}/executions")
	if err != nil {
		return nil, err
	}

	var response CreateJobExecutionResponse
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

// CreateTerm Create a new term within a glossary.
func (client DataCatalogClient) CreateTerm(ctx context.Context, request CreateTermRequest) (response CreateTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTermResponse")
	}
	return
}

// createTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms")
	if err != nil {
		return nil, err
	}

	var response CreateTermResponse
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

// CreateTermRelationship Creates a new term relationship for this term within a glossary.
func (client DataCatalogClient) CreateTermRelationship(ctx context.Context, request CreateTermRelationshipRequest) (response CreateTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTermRelationshipResponse")
	}
	return
}

// createTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createTermRelationship(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships")
	if err != nil {
		return nil, err
	}

	var response CreateTermRelationshipResponse
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

// DeleteAttribute Deletes a specific entity attribute.
func (client DataCatalogClient) DeleteAttribute(ctx context.Context, request DeleteAttributeRequest) (response DeleteAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAttributeResponse")
	}
	return
}

// deleteAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteAttributeResponse
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

// DeleteAttributeTag Deletes a specific entity attribute tag.
func (client DataCatalogClient) DeleteAttributeTag(ctx context.Context, request DeleteAttributeTagRequest) (response DeleteAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAttributeTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAttributeTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAttributeTagResponse")
	}
	return
}

// deleteAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteAttributeTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteAttributeTagResponse
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

// DeleteCatalog Deletes a data catalog resource by identifier.
func (client DataCatalogClient) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (response DeleteCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogResponse")
	}
	return
}

// deleteCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogResponse
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

// DeleteCatalogPrivateEndpoint Deletes a private reverse connection endpoint by identifier.
func (client DataCatalogClient) DeleteCatalogPrivateEndpoint(ctx context.Context, request DeleteCatalogPrivateEndpointRequest) (response DeleteCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogPrivateEndpointResponse")
	}
	return
}

// deleteCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogPrivateEndpointResponse
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

// DeleteConnection Deletes a specific connection of a data asset.
func (client DataCatalogClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionResponse")
	}
	return
}

// deleteConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
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

// DeleteDataAsset Deletes a specific data asset identified by it's key.
func (client DataCatalogClient) DeleteDataAsset(ctx context.Context, request DeleteDataAssetRequest) (response DeleteDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetResponse")
	}
	return
}

// deleteDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetResponse
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

// DeleteDataAssetTag Deletes a specific data asset tag.
func (client DataCatalogClient) DeleteDataAssetTag(ctx context.Context, request DeleteDataAssetTagRequest) (response DeleteDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataAssetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataAssetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetTagResponse")
	}
	return
}

// deleteDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteDataAssetTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetTagResponse
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

// DeleteEntity Deletes a specific data entity.
func (client DataCatalogClient) DeleteEntity(ctx context.Context, request DeleteEntityRequest) (response DeleteEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEntityResponse")
	}
	return
}

// deleteEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteEntityResponse
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

// DeleteEntityTag Deletes a specific entity tag.
func (client DataCatalogClient) DeleteEntityTag(ctx context.Context, request DeleteEntityTagRequest) (response DeleteEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEntityTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEntityTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEntityTagResponse")
	}
	return
}

// deleteEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteEntityTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteEntityTagResponse
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

// DeleteFolder Deletes a specific folder of a data asset identified by it's key.
func (client DataCatalogClient) DeleteFolder(ctx context.Context, request DeleteFolderRequest) (response DeleteFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderResponse")
	}
	return
}

// deleteFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteFolderResponse
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

// DeleteFolderTag Deletes a specific folder tag.
func (client DataCatalogClient) DeleteFolderTag(ctx context.Context, request DeleteFolderTagRequest) (response DeleteFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFolderTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFolderTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderTagResponse")
	}
	return
}

// deleteFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteFolderTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteFolderTagResponse
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

// DeleteGlossary Deletes a specific glossary identified by it's key.
func (client DataCatalogClient) DeleteGlossary(ctx context.Context, request DeleteGlossaryRequest) (response DeleteGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGlossaryResponse")
	}
	return
}

// deleteGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteGlossaryResponse
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

// DeleteJob Deletes a specific job identified by it's key.
func (client DataCatalogClient) DeleteJob(ctx context.Context, request DeleteJobRequest) (response DeleteJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobResponse")
	}
	return
}

// deleteJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/jobs/{jobKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteJobResponse
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

// DeleteJobDefinition Deletes a specific job definition identified by it's key.
func (client DataCatalogClient) DeleteJobDefinition(ctx context.Context, request DeleteJobDefinitionRequest) (response DeleteJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobDefinitionResponse")
	}
	return
}

// deleteJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteJobDefinitionResponse
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

// DeleteTerm Deletes a specific glossary term.
func (client DataCatalogClient) DeleteTerm(ctx context.Context, request DeleteTermRequest) (response DeleteTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTermResponse")
	}
	return
}

// deleteTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteTermResponse
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

// DeleteTermRelationship Deletes a specific glossary term relationship.
func (client DataCatalogClient) DeleteTermRelationship(ctx context.Context, request DeleteTermRelationshipRequest) (response DeleteTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTermRelationshipResponse")
	}
	return
}

// deleteTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteTermRelationship(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships/{termRelationshipKey}")
	if err != nil {
		return nil, err
	}

	var response DeleteTermRelationshipResponse
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

// DetachCatalogPrivateEndpoint Detaches a private reverse connection endpoint resource to a data catalog resource. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client DataCatalogClient) DetachCatalogPrivateEndpoint(ctx context.Context, request DetachCatalogPrivateEndpointRequest) (response DetachCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detachCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachCatalogPrivateEndpointResponse")
	}
	return
}

// detachCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) detachCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/detachCatalogPrivateEndpoint")
	if err != nil {
		return nil, err
	}

	var response DetachCatalogPrivateEndpointResponse
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

// ExpandTreeForGlossary Returns the fully expanded tree hierarchy of parent and child terms in this glossary.
func (client DataCatalogClient) ExpandTreeForGlossary(ctx context.Context, request ExpandTreeForGlossaryRequest) (response ExpandTreeForGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.expandTreeForGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExpandTreeForGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExpandTreeForGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExpandTreeForGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExpandTreeForGlossaryResponse")
	}
	return
}

// expandTreeForGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) expandTreeForGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/expandTree")
	if err != nil {
		return nil, err
	}

	var response ExpandTreeForGlossaryResponse
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

// ExportGlossary Export the glossary and the terms and return the exported glossary as csv or json.
func (client DataCatalogClient) ExportGlossary(ctx context.Context, request ExportGlossaryRequest) (response ExportGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.exportGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportGlossaryResponse")
	}
	return
}

// exportGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) exportGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/export")
	if err != nil {
		return nil, err
	}

	var response ExportGlossaryResponse
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

// GetAttribute Gets a specific entity attribute by key.
func (client DataCatalogClient) GetAttribute(ctx context.Context, request GetAttributeRequest) (response GetAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAttributeResponse")
	}
	return
}

// getAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}")
	if err != nil {
		return nil, err
	}

	var response GetAttributeResponse
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

// GetAttributeTag Gets a specific entity attribute tag by key.
func (client DataCatalogClient) GetAttributeTag(ctx context.Context, request GetAttributeTagRequest) (response GetAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAttributeTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAttributeTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAttributeTagResponse")
	}
	return
}

// getAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getAttributeTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetAttributeTagResponse
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

// GetCatalog Gets a data catalog by identifier.
func (client DataCatalogClient) GetCatalog(ctx context.Context, request GetCatalogRequest) (response GetCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogResponse")
	}
	return
}

// getCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}")
	if err != nil {
		return nil, err
	}

	var response GetCatalogResponse
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

// GetCatalogPrivateEndpoint Gets a specific private reverse connection by identifier.
func (client DataCatalogClient) GetCatalogPrivateEndpoint(ctx context.Context, request GetCatalogPrivateEndpointRequest) (response GetCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogPrivateEndpointResponse")
	}
	return
}

// getCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}")
	if err != nil {
		return nil, err
	}

	var response GetCatalogPrivateEndpointResponse
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

// GetConnection Gets a specific data asset connection by key.
func (client DataCatalogClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionResponse")
	}
	return
}

// getConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}")
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
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

// GetDataAsset Gets a specific data asset for the given key within a data catalog.
func (client DataCatalogClient) GetDataAsset(ctx context.Context, request GetDataAssetRequest) (response GetDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetResponse")
	}
	return
}

// getDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}")
	if err != nil {
		return nil, err
	}

	var response GetDataAssetResponse
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

// GetDataAssetTag Gets a specific data asset tag by key.
func (client DataCatalogClient) GetDataAssetTag(ctx context.Context, request GetDataAssetTagRequest) (response GetDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataAssetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataAssetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetTagResponse")
	}
	return
}

// getDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getDataAssetTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetDataAssetTagResponse
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

// GetEntity Gets a specific data entity by key for a data asset.
func (client DataCatalogClient) GetEntity(ctx context.Context, request GetEntityRequest) (response GetEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEntityResponse")
	}
	return
}

// getEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}")
	if err != nil {
		return nil, err
	}

	var response GetEntityResponse
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

// GetEntityTag Gets a specific entity tag by key.
func (client DataCatalogClient) GetEntityTag(ctx context.Context, request GetEntityTagRequest) (response GetEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEntityTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEntityTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEntityTagResponse")
	}
	return
}

// getEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getEntityTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetEntityTagResponse
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

// GetFolder Gets a specific data asset folder by key.
func (client DataCatalogClient) GetFolder(ctx context.Context, request GetFolderRequest) (response GetFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderResponse")
	}
	return
}

// getFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}")
	if err != nil {
		return nil, err
	}

	var response GetFolderResponse
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

// GetFolderTag Gets a specific folder tag by key.
func (client DataCatalogClient) GetFolderTag(ctx context.Context, request GetFolderTagRequest) (response GetFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFolderTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFolderTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderTagResponse")
	}
	return
}

// getFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getFolderTag(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags/{tagKey}")
	if err != nil {
		return nil, err
	}

	var response GetFolderTagResponse
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

// GetGlossary Gets a specific glossary by key within a data catalog.
func (client DataCatalogClient) GetGlossary(ctx context.Context, request GetGlossaryRequest) (response GetGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGlossaryResponse")
	}
	return
}

// getGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}")
	if err != nil {
		return nil, err
	}

	var response GetGlossaryResponse
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

// GetJob Gets a specific job by key within a data catalog.
func (client DataCatalogClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobResponse")
	}
	return
}

// getJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobResponse
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

// GetJobDefinition Gets a specific job definition by key within a data catalog.
func (client DataCatalogClient) GetJobDefinition(ctx context.Context, request GetJobDefinitionRequest) (response GetJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobDefinitionResponse")
	}
	return
}

// getJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobDefinitionResponse
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

// GetJobExecution Gets a specific job execution by key.
func (client DataCatalogClient) GetJobExecution(ctx context.Context, request GetJobExecutionRequest) (response GetJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobExecutionResponse")
	}
	return
}

// getJobExecution implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobExecutionResponse
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

// GetJobLog Gets a specific job log by key.
func (client DataCatalogClient) GetJobLog(ctx context.Context, request GetJobLogRequest) (response GetJobLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogResponse")
	}
	return
}

// getJobLog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobLog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/logs/{jobLogKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobLogResponse
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

// GetJobMetrics Gets a specific job metric by key.
func (client DataCatalogClient) GetJobMetrics(ctx context.Context, request GetJobMetricsRequest) (response GetJobMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobMetricsResponse")
	}
	return
}

// getJobMetrics implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/metrics/{jobMetricsKey}")
	if err != nil {
		return nil, err
	}

	var response GetJobMetricsResponse
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

// GetTerm Gets a specific glossary term by key.
func (client DataCatalogClient) GetTerm(ctx context.Context, request GetTermRequest) (response GetTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermResponse")
	}
	return
}

// getTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}")
	if err != nil {
		return nil, err
	}

	var response GetTermResponse
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

// GetTermRelationship Gets a specific glossary term relationship by key.
func (client DataCatalogClient) GetTermRelationship(ctx context.Context, request GetTermRelationshipRequest) (response GetTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermRelationshipResponse")
	}
	return
}

// getTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getTermRelationship(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships/{termRelationshipKey}")
	if err != nil {
		return nil, err
	}

	var response GetTermRelationshipResponse
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

// GetType Gets a specific type by key within a data catalog.
func (client DataCatalogClient) GetType(ctx context.Context, request GetTypeRequest) (response GetTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTypeResponse")
	}
	return
}

// getType implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/types/{typeKey}")
	if err != nil {
		return nil, err
	}

	var response GetTypeResponse
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

// GetWorkRequest Gets the status of the work request with the given OCID.
func (client DataCatalogClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DataCatalogClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ImportConnection Import new connection for this data asset.
func (client DataCatalogClient) ImportConnection(ctx context.Context, request ImportConnectionRequest) (response ImportConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportConnectionResponse")
	}
	return
}

// importConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) importConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/importConnection")
	if err != nil {
		return nil, err
	}

	var response ImportConnectionResponse
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

// ImportGlossary Import the glossary and the terms from csv or json files and return the imported glossary resource.
func (client DataCatalogClient) ImportGlossary(ctx context.Context, request ImportGlossaryRequest) (response ImportGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportGlossaryResponse")
	}
	return
}

// importGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) importGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/import")
	if err != nil {
		return nil, err
	}

	var response ImportGlossaryResponse
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

// ListAttributeTags Returns a list of all tags for an entity attribute.
func (client DataCatalogClient) ListAttributeTags(ctx context.Context, request ListAttributeTagsRequest) (response ListAttributeTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttributeTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAttributeTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAttributeTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttributeTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttributeTagsResponse")
	}
	return
}

// listAttributeTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAttributeTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListAttributeTagsResponse
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

// ListAttributes Returns a list of all attributes of an data entity.
func (client DataCatalogClient) ListAttributes(ctx context.Context, request ListAttributesRequest) (response ListAttributesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttributesResponse")
	}
	return
}

// listAttributes implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAttributes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes")
	if err != nil {
		return nil, err
	}

	var response ListAttributesResponse
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

// ListCatalogPrivateEndpoints Returns a list of all the catalog private endpoints in the specified compartment.
func (client DataCatalogClient) ListCatalogPrivateEndpoints(ctx context.Context, request ListCatalogPrivateEndpointsRequest) (response ListCatalogPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCatalogPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCatalogPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogPrivateEndpointsResponse")
	}
	return
}

// listCatalogPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCatalogPrivateEndpoints(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogPrivateEndpoints")
	if err != nil {
		return nil, err
	}

	var response ListCatalogPrivateEndpointsResponse
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

// ListCatalogs Returns a list of all the data catalogs in the specified compartment.
func (client DataCatalogClient) ListCatalogs(ctx context.Context, request ListCatalogsRequest) (response ListCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCatalogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCatalogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogsResponse")
	}
	return
}

// listCatalogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCatalogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs")
	if err != nil {
		return nil, err
	}

	var response ListCatalogsResponse
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

// ListConnections Returns a list of all Connections for a data asset.
func (client DataCatalogClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionsResponse")
	}
	return
}

// listConnections implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listConnections(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections")
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
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

// ListDataAssetTags Returns a list of all tags for a data asset.
func (client DataCatalogClient) ListDataAssetTags(ctx context.Context, request ListDataAssetTagsRequest) (response ListDataAssetTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssetTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataAssetTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataAssetTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetTagsResponse")
	}
	return
}

// listDataAssetTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssetTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListDataAssetTagsResponse
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

// ListDataAssets Returns a list of data assets within a data catalog.
func (client DataCatalogClient) ListDataAssets(ctx context.Context, request ListDataAssetsRequest) (response ListDataAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataAssetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataAssetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetsResponse")
	}
	return
}

// listDataAssets implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssets(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets")
	if err != nil {
		return nil, err
	}

	var response ListDataAssetsResponse
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

// ListEntities Returns a list of all entities of a data asset.
func (client DataCatalogClient) ListEntities(ctx context.Context, request ListEntitiesRequest) (response ListEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntitiesResponse")
	}
	return
}

// listEntities implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listEntities(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities")
	if err != nil {
		return nil, err
	}

	var response ListEntitiesResponse
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

// ListEntityTags Returns a list of all tags for a data entity.
func (client DataCatalogClient) ListEntityTags(ctx context.Context, request ListEntityTagsRequest) (response ListEntityTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntityTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntityTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntityTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntityTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntityTagsResponse")
	}
	return
}

// listEntityTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listEntityTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListEntityTagsResponse
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

// ListFolderTags Returns a list of all tags for a folder.
func (client DataCatalogClient) ListFolderTags(ctx context.Context, request ListFolderTagsRequest) (response ListFolderTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolderTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFolderTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFolderTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFolderTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFolderTagsResponse")
	}
	return
}

// listFolderTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listFolderTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags")
	if err != nil {
		return nil, err
	}

	var response ListFolderTagsResponse
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

// ListFolders Returns a list of all folders.
func (client DataCatalogClient) ListFolders(ctx context.Context, request ListFoldersRequest) (response ListFoldersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFoldersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFoldersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFoldersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFoldersResponse")
	}
	return
}

// listFolders implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listFolders(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders")
	if err != nil {
		return nil, err
	}

	var response ListFoldersResponse
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

// ListGlossaries Returns a list of all glossaries within a data catalog.
func (client DataCatalogClient) ListGlossaries(ctx context.Context, request ListGlossariesRequest) (response ListGlossariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGlossaries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGlossariesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGlossariesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGlossariesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGlossariesResponse")
	}
	return
}

// listGlossaries implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listGlossaries(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries")
	if err != nil {
		return nil, err
	}

	var response ListGlossariesResponse
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

// ListJobDefinitions Returns a list of job definitions within a data catalog.
func (client DataCatalogClient) ListJobDefinitions(ctx context.Context, request ListJobDefinitionsRequest) (response ListJobDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobDefinitionsResponse")
	}
	return
}

// listJobDefinitions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobDefinitions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions")
	if err != nil {
		return nil, err
	}

	var response ListJobDefinitionsResponse
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

// ListJobExecutions Returns a list of job executions for a job.
func (client DataCatalogClient) ListJobExecutions(ctx context.Context, request ListJobExecutionsRequest) (response ListJobExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobExecutionsResponse")
	}
	return
}

// listJobExecutions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions")
	if err != nil {
		return nil, err
	}

	var response ListJobExecutionsResponse
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

// ListJobLogs Returns a list of job logs.
func (client DataCatalogClient) ListJobLogs(ctx context.Context, request ListJobLogsRequest) (response ListJobLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobLogsResponse")
	}
	return
}

// listJobLogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/logs")
	if err != nil {
		return nil, err
	}

	var response ListJobLogsResponse
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

// ListJobMetrics Returns a list of job metrics.
func (client DataCatalogClient) ListJobMetrics(ctx context.Context, request ListJobMetricsRequest) (response ListJobMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobMetricsResponse")
	}
	return
}

// listJobMetrics implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/metrics")
	if err != nil {
		return nil, err
	}

	var response ListJobMetricsResponse
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

// ListJobs Returns a list of jobs within a data catalog.
func (client DataCatalogClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobsResponse")
	}
	return
}

// listJobs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs")
	if err != nil {
		return nil, err
	}

	var response ListJobsResponse
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

// ListTags Returns a list of all user created tags in the system.
func (client DataCatalogClient) ListTags(ctx context.Context, request ListTagsRequest) (response ListTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTagsResponse")
	}
	return
}

// listTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTags(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/tags")
	if err != nil {
		return nil, err
	}

	var response ListTagsResponse
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

// ListTermRelationships Returns a list of all term relationships within a glossary.
func (client DataCatalogClient) ListTermRelationships(ctx context.Context, request ListTermRelationshipsRequest) (response ListTermRelationshipsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTermRelationships, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTermRelationshipsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTermRelationshipsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTermRelationshipsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTermRelationshipsResponse")
	}
	return
}

// listTermRelationships implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTermRelationships(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships")
	if err != nil {
		return nil, err
	}

	var response ListTermRelationshipsResponse
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

// ListTerms Returns a list of all terms within a glossary.
func (client DataCatalogClient) ListTerms(ctx context.Context, request ListTermsRequest) (response ListTermsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTerms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTermsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTermsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTermsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTermsResponse")
	}
	return
}

// listTerms implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTerms(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms")
	if err != nil {
		return nil, err
	}

	var response ListTermsResponse
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

// ListTypes Returns a list of all types within a data catalog.
func (client DataCatalogClient) ListTypes(ctx context.Context, request ListTypesRequest) (response ListTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTypesResponse")
	}
	return
}

// listTypes implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/types")
	if err != nil {
		return nil, err
	}

	var response ListTypesResponse
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

// ListWorkRequestErrors Returns a (paginated) list of errors for a given work request.
func (client DataCatalogClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DataCatalogClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequestLogs Returns a (paginated) list of logs for a given work request.
func (client DataCatalogClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DataCatalogClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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
func (client DataCatalogClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DataCatalogClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ObjectStats Returns stats on objects by type in the repository.
func (client DataCatalogClient) ObjectStats(ctx context.Context, request ObjectStatsRequest) (response ObjectStatsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.objectStats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ObjectStatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ObjectStatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ObjectStatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ObjectStatsResponse")
	}
	return
}

// objectStats implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) objectStats(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/objectStats")
	if err != nil {
		return nil, err
	}

	var response ObjectStatsResponse
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

// ParseConnection Parse data asset references through connections from this data asset.
func (client DataCatalogClient) ParseConnection(ctx context.Context, request ParseConnectionRequest) (response ParseConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.parseConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ParseConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ParseConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ParseConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ParseConnectionResponse")
	}
	return
}

// parseConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) parseConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/parseConnection")
	if err != nil {
		return nil, err
	}

	var response ParseConnectionResponse
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

// SearchCriteria Returns a list of search results within a data catalog.
func (client DataCatalogClient) SearchCriteria(ctx context.Context, request SearchCriteriaRequest) (response SearchCriteriaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchCriteria, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchCriteriaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchCriteriaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchCriteriaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchCriteriaResponse")
	}
	return
}

// searchCriteria implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) searchCriteria(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/search")
	if err != nil {
		return nil, err
	}

	var response SearchCriteriaResponse
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

// TestConnection Test the connection by connecting to the data asset using credentials in the metadata.
func (client DataCatalogClient) TestConnection(ctx context.Context, request TestConnectionRequest) (response TestConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.testConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestConnectionResponse")
	}
	return
}

// testConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) testConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}/actions/test")
	if err != nil {
		return nil, err
	}

	var response TestConnectionResponse
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

// UpdateAttribute Updates a specific data asset attribute.
func (client DataCatalogClient) UpdateAttribute(ctx context.Context, request UpdateAttributeRequest) (response UpdateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAttributeResponse")
	}
	return
}

// updateAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateAttribute(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateAttributeResponse
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

// UpdateCatalog Updates the data catalog.
func (client DataCatalogClient) UpdateCatalog(ctx context.Context, request UpdateCatalogRequest) (response UpdateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogResponse")
	}
	return
}

// updateCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateCatalog(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogResponse
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

// UpdateCatalogPrivateEndpoint Updates the private reverse connection endpoint.
func (client DataCatalogClient) UpdateCatalogPrivateEndpoint(ctx context.Context, request UpdateCatalogPrivateEndpointRequest) (response UpdateCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogPrivateEndpointResponse")
	}
	return
}

// updateCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogPrivateEndpointResponse
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

// UpdateConnection Updates a specific connection of a data asset.
func (client DataCatalogClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectionResponse")
	}
	return
}

// updateConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
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

// UpdateDataAsset Updates a specific data asset identified by the given key.
func (client DataCatalogClient) UpdateDataAsset(ctx context.Context, request UpdateDataAssetRequest) (response UpdateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataAssetResponse")
	}
	return
}

// updateDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateDataAsset(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateDataAssetResponse
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

// UpdateEntity Updates a specific data entity.
func (client DataCatalogClient) UpdateEntity(ctx context.Context, request UpdateEntityRequest) (response UpdateEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEntityResponse")
	}
	return
}

// updateEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateEntityResponse
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

// UpdateFolder Updates a specific folder of a data asset.
func (client DataCatalogClient) UpdateFolder(ctx context.Context, request UpdateFolderRequest) (response UpdateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFolderResponse")
	}
	return
}

// updateFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateFolder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateFolderResponse
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

// UpdateGlossary Updates a specific glossary identified by the given key.
func (client DataCatalogClient) UpdateGlossary(ctx context.Context, request UpdateGlossaryRequest) (response UpdateGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGlossaryResponse")
	}
	return
}

// updateGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateGlossary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateGlossaryResponse
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

// UpdateJob Updates a specific job identified by the given key.
func (client DataCatalogClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobResponse")
	}
	return
}

// updateJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/jobs/{jobKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateJobResponse
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

// UpdateJobDefinition Update a specific job definition identified by the given key.
func (client DataCatalogClient) UpdateJobDefinition(ctx context.Context, request UpdateJobDefinitionRequest) (response UpdateJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobDefinitionResponse")
	}
	return
}

// updateJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateJobDefinition(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateJobDefinitionResponse
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

// UpdateTerm Updates a specific glossary term.
func (client DataCatalogClient) UpdateTerm(ctx context.Context, request UpdateTermRequest) (response UpdateTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermResponse")
	}
	return
}

// updateTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateTerm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateTermResponse
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

// UpdateTermRelationship Updates a specific glossary term relationship.
func (client DataCatalogClient) UpdateTermRelationship(ctx context.Context, request UpdateTermRelationshipRequest) (response UpdateTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermRelationshipResponse")
	}
	return
}

// updateTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateTermRelationship(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships/{termRelationshipKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateTermRelationshipResponse
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

// UploadCredentials Upload connection credentails and metadata for this connection.
func (client DataCatalogClient) UploadCredentials(ctx context.Context, request UploadCredentialsRequest) (response UploadCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.uploadCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadCredentialsResponse")
	}
	return
}

// uploadCredentials implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) uploadCredentials(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}/actions/uploadCredentials")
	if err != nil {
		return nil, err
	}

	var response UploadCredentialsResponse
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

// Users Returns active users in the system.
func (client DataCatalogClient) Users(ctx context.Context, request UsersRequest) (response UsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.users, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UsersResponse")
	}
	return
}

// users implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) users(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/users")
	if err != nil {
		return nil, err
	}

	var response UsersResponse
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

// ValidateConnection Validate connection by connecting to the data asset using credentials in metadata.
func (client DataCatalogClient) ValidateConnection(ctx context.Context, request ValidateConnectionRequest) (response ValidateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateConnectionResponse")
	}
	return
}

// validateConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) validateConnection(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/validateConnection")
	if err != nil {
		return nil, err
	}

	var response ValidateConnectionResponse
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
