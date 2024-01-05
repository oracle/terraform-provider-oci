// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ComputeCloudAtCustomerClient a client for ComputeCloudAtCustomer
type ComputeCloudAtCustomerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewComputeCloudAtCustomerClientWithConfigurationProvider Creates a new default ComputeCloudAtCustomer client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComputeCloudAtCustomerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComputeCloudAtCustomerClient, err error) {
	if enabled := common.CheckForEnabledServices("computecloudatcustomer"); !enabled {
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
	return newComputeCloudAtCustomerClientFromBaseClient(baseClient, provider)
}

// NewComputeCloudAtCustomerClientWithOboToken Creates a new default ComputeCloudAtCustomer client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewComputeCloudAtCustomerClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ComputeCloudAtCustomerClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newComputeCloudAtCustomerClientFromBaseClient(baseClient, configProvider)
}

func newComputeCloudAtCustomerClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ComputeCloudAtCustomerClient, err error) {
	// ComputeCloudAtCustomer service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ComputeCloudAtCustomer"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ComputeCloudAtCustomerClient{BaseClient: baseClient}
	client.BasePath = "20221208"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ComputeCloudAtCustomerClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("computecloudatcustomer", "https://ccc.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ComputeCloudAtCustomerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ComputeCloudAtCustomerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeCccInfrastructureCompartment Moves a Compute Cloud@Customer infrastructure resource from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/ChangeCccInfrastructureCompartment.go.html to see an example of how to use ChangeCccInfrastructureCompartment API.
// A default retry strategy applies to this operation ChangeCccInfrastructureCompartment()
func (client ComputeCloudAtCustomerClient) ChangeCccInfrastructureCompartment(ctx context.Context, request ChangeCccInfrastructureCompartmentRequest) (response ChangeCccInfrastructureCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCccInfrastructureCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCccInfrastructureCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCccInfrastructureCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCccInfrastructureCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCccInfrastructureCompartmentResponse")
	}
	return
}

// changeCccInfrastructureCompartment implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) changeCccInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccInfrastructures/{cccInfrastructureId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCccInfrastructureCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/ChangeCccInfrastructureCompartment"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ChangeCccInfrastructureCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCccUpgradeScheduleCompartment Moves a Compute Cloud@Customer upgrade schedule from one compartment to another using the
// specified OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/ChangeCccUpgradeScheduleCompartment.go.html to see an example of how to use ChangeCccUpgradeScheduleCompartment API.
// A default retry strategy applies to this operation ChangeCccUpgradeScheduleCompartment()
func (client ComputeCloudAtCustomerClient) ChangeCccUpgradeScheduleCompartment(ctx context.Context, request ChangeCccUpgradeScheduleCompartmentRequest) (response ChangeCccUpgradeScheduleCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCccUpgradeScheduleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCccUpgradeScheduleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCccUpgradeScheduleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCccUpgradeScheduleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCccUpgradeScheduleCompartmentResponse")
	}
	return
}

// changeCccUpgradeScheduleCompartment implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) changeCccUpgradeScheduleCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccUpgradeSchedules/{cccUpgradeScheduleId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCccUpgradeScheduleCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccUpgradeSchedule/ChangeCccUpgradeScheduleCompartment"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ChangeCccUpgradeScheduleCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCccInfrastructure Creates a Compute Cloud@Customer infrastructure. Once created, Oracle Services
// must connect the rack in the data center to this Oracle Cloud Infrastructure resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/CreateCccInfrastructure.go.html to see an example of how to use CreateCccInfrastructure API.
// A default retry strategy applies to this operation CreateCccInfrastructure()
func (client ComputeCloudAtCustomerClient) CreateCccInfrastructure(ctx context.Context, request CreateCccInfrastructureRequest) (response CreateCccInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCccInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCccInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCccInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCccInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCccInfrastructureResponse")
	}
	return
}

// createCccInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) createCccInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCccInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/CreateCccInfrastructure"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "CreateCccInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCccUpgradeSchedule Creates a new Compute Cloud@Customer upgrade schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/CreateCccUpgradeSchedule.go.html to see an example of how to use CreateCccUpgradeSchedule API.
// A default retry strategy applies to this operation CreateCccUpgradeSchedule()
func (client ComputeCloudAtCustomerClient) CreateCccUpgradeSchedule(ctx context.Context, request CreateCccUpgradeScheduleRequest) (response CreateCccUpgradeScheduleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCccUpgradeSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCccUpgradeScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCccUpgradeScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCccUpgradeScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCccUpgradeScheduleResponse")
	}
	return
}

// createCccUpgradeSchedule implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) createCccUpgradeSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccUpgradeSchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCccUpgradeScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccUpgradeSchedule/CreateCccUpgradeSchedule"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "CreateCccUpgradeSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCccInfrastructure Deletes a Compute Cloud@Customer infrastructure resource specified by the resource
// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/DeleteCccInfrastructure.go.html to see an example of how to use DeleteCccInfrastructure API.
func (client ComputeCloudAtCustomerClient) DeleteCccInfrastructure(ctx context.Context, request DeleteCccInfrastructureRequest) (response DeleteCccInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCccInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCccInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCccInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCccInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCccInfrastructureResponse")
	}
	return
}

// deleteCccInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) deleteCccInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cccInfrastructures/{cccInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCccInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/DeleteCccInfrastructure"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "DeleteCccInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCccUpgradeSchedule Deletes a Compute Cloud@Customer upgrade schedule by the specified
// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/DeleteCccUpgradeSchedule.go.html to see an example of how to use DeleteCccUpgradeSchedule API.
func (client ComputeCloudAtCustomerClient) DeleteCccUpgradeSchedule(ctx context.Context, request DeleteCccUpgradeScheduleRequest) (response DeleteCccUpgradeScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCccUpgradeSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCccUpgradeScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCccUpgradeScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCccUpgradeScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCccUpgradeScheduleResponse")
	}
	return
}

// deleteCccUpgradeSchedule implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) deleteCccUpgradeSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cccUpgradeSchedules/{cccUpgradeScheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCccUpgradeScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccUpgradeSchedule/DeleteCccUpgradeSchedule"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "DeleteCccUpgradeSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCccInfrastructure Gets a Compute Cloud@Customer infrastructure using the infrastructure
// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/GetCccInfrastructure.go.html to see an example of how to use GetCccInfrastructure API.
// A default retry strategy applies to this operation GetCccInfrastructure()
func (client ComputeCloudAtCustomerClient) GetCccInfrastructure(ctx context.Context, request GetCccInfrastructureRequest) (response GetCccInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCccInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCccInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCccInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCccInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCccInfrastructureResponse")
	}
	return
}

// getCccInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) getCccInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccInfrastructures/{cccInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCccInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/GetCccInfrastructure"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetCccInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCccUpgradeSchedule Gets a Compute Cloud@Customer upgrade schedule by the specified
// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/GetCccUpgradeSchedule.go.html to see an example of how to use GetCccUpgradeSchedule API.
// A default retry strategy applies to this operation GetCccUpgradeSchedule()
func (client ComputeCloudAtCustomerClient) GetCccUpgradeSchedule(ctx context.Context, request GetCccUpgradeScheduleRequest) (response GetCccUpgradeScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCccUpgradeSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCccUpgradeScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCccUpgradeScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCccUpgradeScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCccUpgradeScheduleResponse")
	}
	return
}

// getCccUpgradeSchedule implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) getCccUpgradeSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccUpgradeSchedules/{cccUpgradeScheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCccUpgradeScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccUpgradeSchedule/GetCccUpgradeSchedule"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetCccUpgradeSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccInfrastructures Returns a list of Compute Cloud@Customer infrastructures.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/ListCccInfrastructures.go.html to see an example of how to use ListCccInfrastructures API.
// A default retry strategy applies to this operation ListCccInfrastructures()
func (client ComputeCloudAtCustomerClient) ListCccInfrastructures(ctx context.Context, request ListCccInfrastructuresRequest) (response ListCccInfrastructuresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccInfrastructures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccInfrastructuresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccInfrastructuresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccInfrastructuresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccInfrastructuresResponse")
	}
	return
}

// listCccInfrastructures implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccInfrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccInfrastructuresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructureCollection/ListCccInfrastructures"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccInfrastructures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccUpgradeSchedules Returns a list of Compute Cloud@Customer upgrade schedules.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/ListCccUpgradeSchedules.go.html to see an example of how to use ListCccUpgradeSchedules API.
// A default retry strategy applies to this operation ListCccUpgradeSchedules()
func (client ComputeCloudAtCustomerClient) ListCccUpgradeSchedules(ctx context.Context, request ListCccUpgradeSchedulesRequest) (response ListCccUpgradeSchedulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccUpgradeSchedules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccUpgradeSchedulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccUpgradeSchedulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccUpgradeSchedulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccUpgradeSchedulesResponse")
	}
	return
}

// listCccUpgradeSchedules implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccUpgradeSchedules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccUpgradeSchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccUpgradeSchedulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccUpgradeScheduleCollection/ListCccUpgradeSchedules"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccUpgradeSchedules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCccInfrastructure Updates Compute Cloud@Customer infrastructure resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/UpdateCccInfrastructure.go.html to see an example of how to use UpdateCccInfrastructure API.
func (client ComputeCloudAtCustomerClient) UpdateCccInfrastructure(ctx context.Context, request UpdateCccInfrastructureRequest) (response UpdateCccInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCccInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCccInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCccInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCccInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCccInfrastructureResponse")
	}
	return
}

// updateCccInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) updateCccInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cccInfrastructures/{cccInfrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCccInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/UpdateCccInfrastructure"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "UpdateCccInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCccUpgradeSchedule Updates the Compute Cloud@Customer upgrade schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/UpdateCccUpgradeSchedule.go.html to see an example of how to use UpdateCccUpgradeSchedule API.
func (client ComputeCloudAtCustomerClient) UpdateCccUpgradeSchedule(ctx context.Context, request UpdateCccUpgradeScheduleRequest) (response UpdateCccUpgradeScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCccUpgradeSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCccUpgradeScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCccUpgradeScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCccUpgradeScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCccUpgradeScheduleResponse")
	}
	return
}

// updateCccUpgradeSchedule implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) updateCccUpgradeSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cccUpgradeSchedules/{cccUpgradeScheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCccUpgradeScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccUpgradeSchedule/UpdateCccUpgradeSchedule"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "UpdateCccUpgradeSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
