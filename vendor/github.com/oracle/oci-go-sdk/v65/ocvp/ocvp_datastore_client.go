// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DatastoreClient a client for Datastore
type DatastoreClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatastoreClientWithConfigurationProvider Creates a new default Datastore client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatastoreClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatastoreClient, err error) {
	if enabled := common.CheckForEnabledServices("ocvp"); !enabled {
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
	return newDatastoreClientFromBaseClient(baseClient, provider)
}

// NewDatastoreClientWithOboToken Creates a new default Datastore client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatastoreClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatastoreClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatastoreClientFromBaseClient(baseClient, configProvider)
}

func newDatastoreClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatastoreClient, err error) {
	// Datastore service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Datastore"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatastoreClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatastoreClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatastoreClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatastoreClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddBlockVolumeToDatastore Add the specified Block Volume to the provided Datastore.
// Use the WorkRequest operations to track the
// addition of the block volume to the Datastore.
// A default retry strategy applies to this operation AddBlockVolumeToDatastore()
func (client DatastoreClient) AddBlockVolumeToDatastore(ctx context.Context, request AddBlockVolumeToDatastoreRequest) (response AddBlockVolumeToDatastoreResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addBlockVolumeToDatastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddBlockVolumeToDatastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddBlockVolumeToDatastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddBlockVolumeToDatastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddBlockVolumeToDatastoreResponse")
	}
	return
}

// addBlockVolumeToDatastore implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) addBlockVolumeToDatastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastores/{datastoreId}/actions/addBlockVolume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddBlockVolumeToDatastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/AddBlockVolumeToDatastore"
		err = common.PostProcessServiceError(err, "Datastore", "AddBlockVolumeToDatastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatastoreCompartment Moves an Datastore into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation ChangeDatastoreCompartment()
func (client DatastoreClient) ChangeDatastoreCompartment(ctx context.Context, request ChangeDatastoreCompartmentRequest) (response ChangeDatastoreCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatastoreCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatastoreCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatastoreCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatastoreCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatastoreCompartmentResponse")
	}
	return
}

// changeDatastoreCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) changeDatastoreCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastores/{datastoreId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatastoreCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/ChangeDatastoreCompartment"
		err = common.PostProcessServiceError(err, "Datastore", "ChangeDatastoreCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatastore Creates a Oracle Cloud VMware Solution Datastore.
// Use the WorkRequest operations to track the
// creation of the Datastore.
// A default retry strategy applies to this operation CreateDatastore()
func (client DatastoreClient) CreateDatastore(ctx context.Context, request CreateDatastoreRequest) (response CreateDatastoreResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatastoreResponse")
	}
	return
}

// createDatastore implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) createDatastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/CreateDatastore"
		err = common.PostProcessServiceError(err, "Datastore", "CreateDatastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatastore Deletes the specified Datastore.
// Use the WorkRequest operations to track the
// deletion of the Datastore.
// A default retry strategy applies to this operation DeleteDatastore()
func (client DatastoreClient) DeleteDatastore(ctx context.Context, request DeleteDatastoreRequest) (response DeleteDatastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatastoreResponse")
	}
	return
}

// deleteDatastore implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) deleteDatastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/datastores/{datastoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/DeleteDatastore"
		err = common.PostProcessServiceError(err, "Datastore", "DeleteDatastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatastore Get the specified Datastore's information.
// A default retry strategy applies to this operation GetDatastore()
func (client DatastoreClient) GetDatastore(ctx context.Context, request GetDatastoreRequest) (response GetDatastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatastoreResponse")
	}
	return
}

// getDatastore implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) getDatastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/datastores/{datastoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/GetDatastore"
		err = common.PostProcessServiceError(err, "Datastore", "GetDatastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatastores List the Datastores in the specified compartment. The list can be filtered
// by compartment, datastore id, display name and lifecycle state.
// A default retry strategy applies to this operation ListDatastores()
func (client DatastoreClient) ListDatastores(ctx context.Context, request ListDatastoresRequest) (response ListDatastoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatastores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatastoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatastoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatastoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatastoresResponse")
	}
	return
}

// listDatastores implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) listDatastores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/datastores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatastoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/ListDatastores"
		err = common.PostProcessServiceError(err, "Datastore", "ListDatastores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatastore Updates the specified Datastore.
// **Important:** Updating a Datastore affects only certain attributes in the `Datastore`
// object and does not affect the VMware environment currently running.
// A default retry strategy applies to this operation UpdateDatastore()
func (client DatastoreClient) UpdateDatastore(ctx context.Context, request UpdateDatastoreRequest) (response UpdateDatastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatastoreResponse")
	}
	return
}

// updateDatastore implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClient) updateDatastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/datastores/{datastoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Datastore/UpdateDatastore"
		err = common.PostProcessServiceError(err, "Datastore", "UpdateDatastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
