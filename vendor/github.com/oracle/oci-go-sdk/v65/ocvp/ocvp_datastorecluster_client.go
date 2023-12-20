// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DatastoreClusterClient a client for DatastoreCluster
type DatastoreClusterClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatastoreClusterClientWithConfigurationProvider Creates a new default DatastoreCluster client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatastoreClusterClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatastoreClusterClient, err error) {
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
	return newDatastoreClusterClientFromBaseClient(baseClient, provider)
}

// NewDatastoreClusterClientWithOboToken Creates a new default DatastoreCluster client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatastoreClusterClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatastoreClusterClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatastoreClusterClientFromBaseClient(baseClient, configProvider)
}

func newDatastoreClusterClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatastoreClusterClient, err error) {
	// DatastoreCluster service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DatastoreCluster"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatastoreClusterClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatastoreClusterClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatastoreClusterClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatastoreClusterClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddDatastoreToDatastoreCluster Add the specified Datastore to the provided Datastore Cluster.
// A default retry strategy applies to this operation AddDatastoreToDatastoreCluster()
func (client DatastoreClusterClient) AddDatastoreToDatastoreCluster(ctx context.Context, request AddDatastoreToDatastoreClusterRequest) (response AddDatastoreToDatastoreClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addDatastoreToDatastoreCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDatastoreToDatastoreClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDatastoreToDatastoreClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDatastoreToDatastoreClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDatastoreToDatastoreClusterResponse")
	}
	return
}

// addDatastoreToDatastoreCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) addDatastoreToDatastoreCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/addDatastore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDatastoreToDatastoreClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/AddDatastoreToDatastoreCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "AddDatastoreToDatastoreCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachDatastoreClusterToCluster Attach the specified Datastore Cluster to the provided Vmware Cluster.
// Use the WorkRequest operations to track the
// attachment of the Datastore.
// A default retry strategy applies to this operation AttachDatastoreClusterToCluster()
func (client DatastoreClusterClient) AttachDatastoreClusterToCluster(ctx context.Context, request AttachDatastoreClusterToClusterRequest) (response AttachDatastoreClusterToClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachDatastoreClusterToCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachDatastoreClusterToClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachDatastoreClusterToClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachDatastoreClusterToClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachDatastoreClusterToClusterResponse")
	}
	return
}

// attachDatastoreClusterToCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) attachDatastoreClusterToCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/attachToCluster", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachDatastoreClusterToClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/AttachDatastoreClusterToCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "AttachDatastoreClusterToCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachDatastoreClusterToEsxiHost Attach the specified Datastore Cluster to the provided ESXi Hosts.
// Use the WorkRequest operations to track the
// attachment of the Datastore.
// A default retry strategy applies to this operation AttachDatastoreClusterToEsxiHost()
func (client DatastoreClusterClient) AttachDatastoreClusterToEsxiHost(ctx context.Context, request AttachDatastoreClusterToEsxiHostRequest) (response AttachDatastoreClusterToEsxiHostResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachDatastoreClusterToEsxiHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachDatastoreClusterToEsxiHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachDatastoreClusterToEsxiHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachDatastoreClusterToEsxiHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachDatastoreClusterToEsxiHostResponse")
	}
	return
}

// attachDatastoreClusterToEsxiHost implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) attachDatastoreClusterToEsxiHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/attachToEsxiHost", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachDatastoreClusterToEsxiHostResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/AttachDatastoreClusterToEsxiHost"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "AttachDatastoreClusterToEsxiHost", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatastoreClusterCompartment Moves an Datastore Cluster into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation ChangeDatastoreClusterCompartment()
func (client DatastoreClusterClient) ChangeDatastoreClusterCompartment(ctx context.Context, request ChangeDatastoreClusterCompartmentRequest) (response ChangeDatastoreClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatastoreClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatastoreClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatastoreClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatastoreClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatastoreClusterCompartmentResponse")
	}
	return
}

// changeDatastoreClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) changeDatastoreClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatastoreClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/ChangeDatastoreClusterCompartment"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "ChangeDatastoreClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatastoreCluster Creates a Oracle Cloud VMware Solution Datastore Cluster.
// A default retry strategy applies to this operation CreateDatastoreCluster()
func (client DatastoreClusterClient) CreateDatastoreCluster(ctx context.Context, request CreateDatastoreClusterRequest) (response CreateDatastoreClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatastoreCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatastoreClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatastoreClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatastoreClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatastoreClusterResponse")
	}
	return
}

// createDatastoreCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) createDatastoreCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatastoreClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/CreateDatastoreCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "CreateDatastoreCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatastoreCluster Deletes the specified Datastore Cluster.
// A default retry strategy applies to this operation DeleteDatastoreCluster()
func (client DatastoreClusterClient) DeleteDatastoreCluster(ctx context.Context, request DeleteDatastoreClusterRequest) (response DeleteDatastoreClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatastoreCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatastoreClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatastoreClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatastoreClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatastoreClusterResponse")
	}
	return
}

// deleteDatastoreCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) deleteDatastoreCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/datastoreClusters/{datastoreClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatastoreClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/DeleteDatastoreCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "DeleteDatastoreCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachDatastoreClusterFromCluster Detach the specified Datastore Cluster from the provided Vmware Cluster.
// Use the WorkRequest operations to track the
// detachment of the Datastore.
// A default retry strategy applies to this operation DetachDatastoreClusterFromCluster()
func (client DatastoreClusterClient) DetachDatastoreClusterFromCluster(ctx context.Context, request DetachDatastoreClusterFromClusterRequest) (response DetachDatastoreClusterFromClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachDatastoreClusterFromCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachDatastoreClusterFromClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachDatastoreClusterFromClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachDatastoreClusterFromClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachDatastoreClusterFromClusterResponse")
	}
	return
}

// detachDatastoreClusterFromCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) detachDatastoreClusterFromCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/detachFromCluster", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachDatastoreClusterFromClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/DetachDatastoreClusterFromCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "DetachDatastoreClusterFromCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachDatastoreClusterFromEsxiHost Detach the specified Datastore Cluster from the provided ESXi Hosts.
// Use the WorkRequest operations to track the
// detachment of the Datastore.
// A default retry strategy applies to this operation DetachDatastoreClusterFromEsxiHost()
func (client DatastoreClusterClient) DetachDatastoreClusterFromEsxiHost(ctx context.Context, request DetachDatastoreClusterFromEsxiHostRequest) (response DetachDatastoreClusterFromEsxiHostResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachDatastoreClusterFromEsxiHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachDatastoreClusterFromEsxiHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachDatastoreClusterFromEsxiHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachDatastoreClusterFromEsxiHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachDatastoreClusterFromEsxiHostResponse")
	}
	return
}

// detachDatastoreClusterFromEsxiHost implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) detachDatastoreClusterFromEsxiHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/detachFromEsxiHost", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachDatastoreClusterFromEsxiHostResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/DetachDatastoreClusterFromEsxiHost"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "DetachDatastoreClusterFromEsxiHost", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatastoreCluster Get the specified Datastore Cluster information.
// A default retry strategy applies to this operation GetDatastoreCluster()
func (client DatastoreClusterClient) GetDatastoreCluster(ctx context.Context, request GetDatastoreClusterRequest) (response GetDatastoreClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatastoreCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatastoreClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatastoreClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatastoreClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatastoreClusterResponse")
	}
	return
}

// getDatastoreCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) getDatastoreCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/datastoreClusters/{datastoreClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatastoreClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/GetDatastoreCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "GetDatastoreCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatastoreClusters List the Datastore Clusters in the specified compartment. The list can be filtered
// by compartment, Datastore Cluster, Display name and Lifecycle state
// A default retry strategy applies to this operation ListDatastoreClusters()
func (client DatastoreClusterClient) ListDatastoreClusters(ctx context.Context, request ListDatastoreClustersRequest) (response ListDatastoreClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatastoreClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatastoreClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatastoreClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatastoreClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatastoreClustersResponse")
	}
	return
}

// listDatastoreClusters implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) listDatastoreClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/datastoreClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatastoreClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/ListDatastoreClusters"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "ListDatastoreClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDatastoreFromDatastoreCluster Remove the specified Datastore from the provided Datastore Cluster.
// A default retry strategy applies to this operation RemoveDatastoreFromDatastoreCluster()
func (client DatastoreClusterClient) RemoveDatastoreFromDatastoreCluster(ctx context.Context, request RemoveDatastoreFromDatastoreClusterRequest) (response RemoveDatastoreFromDatastoreClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeDatastoreFromDatastoreCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDatastoreFromDatastoreClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDatastoreFromDatastoreClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDatastoreFromDatastoreClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDatastoreFromDatastoreClusterResponse")
	}
	return
}

// removeDatastoreFromDatastoreCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) removeDatastoreFromDatastoreCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/datastoreClusters/{datastoreClusterId}/actions/removeDatastore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDatastoreFromDatastoreClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/RemoveDatastoreFromDatastoreCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "RemoveDatastoreFromDatastoreCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatastoreCluster Updates the specified Datastore Cluster.
// **Important:** Updating a Datastore Cluster affects only certain attributes in the `Datastore Cluster`
// object and does not affect the VMware environment currently running.
// A default retry strategy applies to this operation UpdateDatastoreCluster()
func (client DatastoreClusterClient) UpdateDatastoreCluster(ctx context.Context, request UpdateDatastoreClusterRequest) (response UpdateDatastoreClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatastoreCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatastoreClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatastoreClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatastoreClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatastoreClusterResponse")
	}
	return
}

// updateDatastoreCluster implements the OCIOperation interface (enables retrying operations)
func (client DatastoreClusterClient) updateDatastoreCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/datastoreClusters/{datastoreClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatastoreClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/DatastoreCluster/UpdateDatastoreCluster"
		err = common.PostProcessServiceError(err, "DatastoreCluster", "UpdateDatastoreCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
