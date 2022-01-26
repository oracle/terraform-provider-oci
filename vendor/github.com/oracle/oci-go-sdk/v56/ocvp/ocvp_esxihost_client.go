// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//EsxiHostClient a client for EsxiHost
type EsxiHostClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEsxiHostClientWithConfigurationProvider Creates a new default EsxiHost client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEsxiHostClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EsxiHostClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newEsxiHostClientFromBaseClient(baseClient, provider)
}

// NewEsxiHostClientWithOboToken Creates a new default EsxiHost client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewEsxiHostClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EsxiHostClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEsxiHostClientFromBaseClient(baseClient, configProvider)
}

func newEsxiHostClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EsxiHostClient, err error) {
	// EsxiHost service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EsxiHostClient{BaseClient: baseClient}
	client.BasePath = "20200501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EsxiHostClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EsxiHostClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EsxiHostClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateEsxiHost Adds another ESXi host to an existing SDDC. The attributes of the specified
// `Sddc` determine the VMware software and other configuration settings used
// by the ESXi host.
// Use the WorkRequest operations to track the
// creation of the ESXi host.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/CreateEsxiHost.go.html to see an example of how to use CreateEsxiHost API.
// A default retry strategy applies to this operation CreateEsxiHost()
func (client EsxiHostClient) CreateEsxiHost(ctx context.Context, request CreateEsxiHostRequest) (response CreateEsxiHostResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEsxiHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEsxiHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEsxiHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEsxiHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEsxiHostResponse")
	}
	return
}

// createEsxiHost implements the OCIOperation interface (enables retrying operations)
func (client EsxiHostClient) createEsxiHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/esxiHosts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEsxiHostResponse
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

// DeleteEsxiHost Deletes the specified ESXi host. Before deleting the host, back up or
// migrate any VMware workloads running on it.
// When you delete an ESXi host, Oracle does not remove the node
// configuration within the VMware environment itself. That is
// your responsibility.
// **Note:** If you delete EXSi hosts from the SDDC to total less than 3,
// you are still billed for the 3 minimum recommended EXSi hosts. Also,
// you cannot add more VMware workloads to the SDDC until it again has at
// least 3 ESXi hosts.
// Use the WorkRequest operations to track the
// deletion of the ESXi host.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/DeleteEsxiHost.go.html to see an example of how to use DeleteEsxiHost API.
// A default retry strategy applies to this operation DeleteEsxiHost()
func (client EsxiHostClient) DeleteEsxiHost(ctx context.Context, request DeleteEsxiHostRequest) (response DeleteEsxiHostResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEsxiHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEsxiHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEsxiHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEsxiHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEsxiHostResponse")
	}
	return
}

// deleteEsxiHost implements the OCIOperation interface (enables retrying operations)
func (client EsxiHostClient) deleteEsxiHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/esxiHosts/{esxiHostId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEsxiHostResponse
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

// GetEsxiHost Gets the specified ESXi host's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/GetEsxiHost.go.html to see an example of how to use GetEsxiHost API.
// A default retry strategy applies to this operation GetEsxiHost()
func (client EsxiHostClient) GetEsxiHost(ctx context.Context, request GetEsxiHostRequest) (response GetEsxiHostResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEsxiHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEsxiHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEsxiHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEsxiHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEsxiHostResponse")
	}
	return
}

// getEsxiHost implements the OCIOperation interface (enables retrying operations)
func (client EsxiHostClient) getEsxiHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/esxiHosts/{esxiHostId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEsxiHostResponse
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

// ListEsxiHosts Lists the ESXi hosts in the specified SDDC. The list can be filtered
// by Compute instance OCID or ESXi display name.
// Remember that in terms of implementation, an ESXi host is a Compute instance that
// is configured with the chosen bundle of VMware software. Each `EsxiHost`
// object has its own OCID (`id`), and a separate attribute for the OCID of
// the Compute instance (`computeInstanceId`). When filtering the list of
// ESXi hosts, you can specify the OCID of the Compute instance, not the
// ESXi host OCID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListEsxiHosts.go.html to see an example of how to use ListEsxiHosts API.
// A default retry strategy applies to this operation ListEsxiHosts()
func (client EsxiHostClient) ListEsxiHosts(ctx context.Context, request ListEsxiHostsRequest) (response ListEsxiHostsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEsxiHosts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEsxiHostsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEsxiHostsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEsxiHostsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEsxiHostsResponse")
	}
	return
}

// listEsxiHosts implements the OCIOperation interface (enables retrying operations)
func (client EsxiHostClient) listEsxiHosts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/esxiHosts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEsxiHostsResponse
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

// UpdateEsxiHost Updates the specified ESXi host.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/UpdateEsxiHost.go.html to see an example of how to use UpdateEsxiHost API.
// A default retry strategy applies to this operation UpdateEsxiHost()
func (client EsxiHostClient) UpdateEsxiHost(ctx context.Context, request UpdateEsxiHostRequest) (response UpdateEsxiHostResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEsxiHost, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEsxiHostResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEsxiHostResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEsxiHostResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEsxiHostResponse")
	}
	return
}

// updateEsxiHost implements the OCIOperation interface (enables retrying operations)
func (client EsxiHostClient) updateEsxiHost(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/esxiHosts/{esxiHostId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEsxiHostResponse
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
