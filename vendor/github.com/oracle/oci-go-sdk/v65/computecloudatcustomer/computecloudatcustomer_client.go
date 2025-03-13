// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
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

// CancelWorkRequest Cancels a work request.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client ComputeCloudAtCustomerClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCccInfrastructureCompartment Moves a Compute Cloud@Customer infrastructure resource from one compartment to another.
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
// specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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

// CreateCccFlexNetwork Creates a Compute Cloud@Customer flexNetwork. Once created must be attached to an
// infrastructure subnet.
// A default retry strategy applies to this operation CreateCccFlexNetwork()
func (client ComputeCloudAtCustomerClient) CreateCccFlexNetwork(ctx context.Context, request CreateCccFlexNetworkRequest) (response CreateCccFlexNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCccFlexNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCccFlexNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCccFlexNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCccFlexNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCccFlexNetworkResponse")
	}
	return
}

// createCccFlexNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) createCccFlexNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccFlexNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCccFlexNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetwork/CreateCccFlexNetwork"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "CreateCccFlexNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCccFlexNetworkAttachment Creates a Compute Cloud@Customer flexNetworkAttachment. Once created must be attached to an
// infrastructure subnet.
// A default retry strategy applies to this operation CreateCccFlexNetworkAttachment()
func (client ComputeCloudAtCustomerClient) CreateCccFlexNetworkAttachment(ctx context.Context, request CreateCccFlexNetworkAttachmentRequest) (response CreateCccFlexNetworkAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCccFlexNetworkAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCccFlexNetworkAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCccFlexNetworkAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCccFlexNetworkAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCccFlexNetworkAttachmentResponse")
	}
	return
}

// createCccFlexNetworkAttachment implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) createCccFlexNetworkAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccFlexNetworkAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCccFlexNetworkAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetworkAttachment/CreateCccFlexNetworkAttachment"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "CreateCccFlexNetworkAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCccInfrastructure Creates a Compute Cloud@Customer infrastructure. Once created, Oracle Services
// must connect the rack in the data center to this Oracle Cloud Infrastructure resource.
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

// CreateCccProvisionedPackage Creates a new Compute Cloud@Customer marketplace provisioned package and attempts to
// provision the selected package to the identified Compute Cloud@Customer infrastructure.
// A default retry strategy applies to this operation CreateCccProvisionedPackage()
func (client ComputeCloudAtCustomerClient) CreateCccProvisionedPackage(ctx context.Context, request CreateCccProvisionedPackageRequest) (response CreateCccProvisionedPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCccProvisionedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCccProvisionedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCccProvisionedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCccProvisionedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCccProvisionedPackageResponse")
	}
	return
}

// createCccProvisionedPackage implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) createCccProvisionedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/cccProvisionedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCccProvisionedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccProvisionedPackage/CreateCccProvisionedPackage"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "CreateCccProvisionedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCccUpgradeSchedule Creates a new Compute Cloud@Customer upgrade schedule.
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

// DeleteCccFlexNetwork Deletes a Compute Cloud@Customer flexNetwork resource specified by the resource
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
func (client ComputeCloudAtCustomerClient) DeleteCccFlexNetwork(ctx context.Context, request DeleteCccFlexNetworkRequest) (response DeleteCccFlexNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCccFlexNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCccFlexNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCccFlexNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCccFlexNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCccFlexNetworkResponse")
	}
	return
}

// deleteCccFlexNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) deleteCccFlexNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cccFlexNetworks/{cccFlexNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCccFlexNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetwork/DeleteCccFlexNetwork"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "DeleteCccFlexNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCccFlexNetworkAttachment Deletes a Compute Cloud@Customer flexNetworkAttachment resource specified by the resource
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
func (client ComputeCloudAtCustomerClient) DeleteCccFlexNetworkAttachment(ctx context.Context, request DeleteCccFlexNetworkAttachmentRequest) (response DeleteCccFlexNetworkAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCccFlexNetworkAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCccFlexNetworkAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCccFlexNetworkAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCccFlexNetworkAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCccFlexNetworkAttachmentResponse")
	}
	return
}

// deleteCccFlexNetworkAttachment implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) deleteCccFlexNetworkAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cccFlexNetworkAttachments/{cccFlexNetworkAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCccFlexNetworkAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetworkAttachment/DeleteCccFlexNetworkAttachment"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "DeleteCccFlexNetworkAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCccInfrastructure Deletes a Compute Cloud@Customer infrastructure resource specified by the resource
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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

// DeleteCccProvisionedPackage Deletes a Compute Cloud@Customer marketplace provisioned package identified by the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). The package will also be deprovisioned
// from the associated Compute Cloud@Customer infrastructure.
func (client ComputeCloudAtCustomerClient) DeleteCccProvisionedPackage(ctx context.Context, request DeleteCccProvisionedPackageRequest) (response DeleteCccProvisionedPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCccProvisionedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCccProvisionedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCccProvisionedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCccProvisionedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCccProvisionedPackageResponse")
	}
	return
}

// deleteCccProvisionedPackage implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) deleteCccProvisionedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/cccProvisionedPackages/{cccProvisionedPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCccProvisionedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccProvisionedPackage/DeleteCccProvisionedPackage"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "DeleteCccProvisionedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCccUpgradeSchedule Deletes a Compute Cloud@Customer upgrade schedule by the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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

// GetCccFlexNetwork Gets a Compute Cloud@Customer flexNetwork using the flexNetwork
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation GetCccFlexNetwork()
func (client ComputeCloudAtCustomerClient) GetCccFlexNetwork(ctx context.Context, request GetCccFlexNetworkRequest) (response GetCccFlexNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCccFlexNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCccFlexNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCccFlexNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCccFlexNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCccFlexNetworkResponse")
	}
	return
}

// getCccFlexNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) getCccFlexNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccFlexNetworks/{cccFlexNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCccFlexNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetwork/GetCccFlexNetwork"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetCccFlexNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCccFlexNetworkAttachment Gets a Compute Cloud@Customer flexNetworkAttachment using the flexNetworkAttachment
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation GetCccFlexNetworkAttachment()
func (client ComputeCloudAtCustomerClient) GetCccFlexNetworkAttachment(ctx context.Context, request GetCccFlexNetworkAttachmentRequest) (response GetCccFlexNetworkAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCccFlexNetworkAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCccFlexNetworkAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCccFlexNetworkAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCccFlexNetworkAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCccFlexNetworkAttachmentResponse")
	}
	return
}

// getCccFlexNetworkAttachment implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) getCccFlexNetworkAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccFlexNetworkAttachments/{cccFlexNetworkAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCccFlexNetworkAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetworkAttachment/GetCccFlexNetworkAttachment"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetCccFlexNetworkAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCccInfrastructure Gets a Compute Cloud@Customer infrastructure using the infrastructure
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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

// GetCccListing Gets detailed information about a listing, including the listing's name, version, description, and
// resources.
// The compartmentId query parameter is required but is used only for authorization to the operation, the
// results of the request will not change based on the compartmentId parameter.
// If you plan to launch an instance from an image listing on a Compute Cloud@Customer infrastructure,
// you must create a provisioned package from the listing and a specific package version. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Creating a provisioned package requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetCccPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/)
// To get the image ID to launch an instance, issue a CreateCccProvisionedPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// Once the asynchronous provisioning operation is complete, call the GetCccProvisionedPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// Lastly, to launch the instance on the Compute Cloud@Customer infrastructure directly,
// use the image ID of the listing resource version to issue the LaunchInstance (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// A default retry strategy applies to this operation GetCccListing()
func (client ComputeCloudAtCustomerClient) GetCccListing(ctx context.Context, request GetCccListingRequest) (response GetCccListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCccListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCccListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCccListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCccListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCccListingResponse")
	}
	return
}

// getCccListing implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) getCccListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccListings/{cccListingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCccListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccListing/GetCccListing"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetCccListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCccProvisionedPackage Gets a Compute Cloud@Customer marketplace provisioned package by the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// A default retry strategy applies to this operation GetCccProvisionedPackage()
func (client ComputeCloudAtCustomerClient) GetCccProvisionedPackage(ctx context.Context, request GetCccProvisionedPackageRequest) (response GetCccProvisionedPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCccProvisionedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCccProvisionedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCccProvisionedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCccProvisionedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCccProvisionedPackageResponse")
	}
	return
}

// getCccProvisionedPackage implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) getCccProvisionedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccProvisionedPackages/{cccProvisionedPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCccProvisionedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccProvisionedPackage/GetCccProvisionedPackage"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetCccProvisionedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCccUpgradeSchedule Gets a Compute Cloud@Customer upgrade schedule by the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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

// GetWorkRequest Gets the details of a work request.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ComputeCloudAtCustomerClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client ComputeCloudAtCustomerClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccAgreements Returns the terms of use agreements that must be accepted before you can deploy the specified version of a package.
// The compartmentId query parameter is required but is used only for authorization to the operation, the
// results of the request will not change based on the compartmentId parameter.
// A default retry strategy applies to this operation ListCccAgreements()
func (client ComputeCloudAtCustomerClient) ListCccAgreements(ctx context.Context, request ListCccAgreementsRequest) (response ListCccAgreementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccAgreements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccAgreementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccAgreementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccAgreementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccAgreementsResponse")
	}
	return
}

// listCccAgreements implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccAgreements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccAgreements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccAgreementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccAgreementCollection/ListCccAgreements"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccAgreements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccFlexNetworkAttachments Returns a list of Compute Cloud@Customer flexNetworkAttachments associated with a flexNetwork.
// A default retry strategy applies to this operation ListCccFlexNetworkAttachments()
func (client ComputeCloudAtCustomerClient) ListCccFlexNetworkAttachments(ctx context.Context, request ListCccFlexNetworkAttachmentsRequest) (response ListCccFlexNetworkAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccFlexNetworkAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccFlexNetworkAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccFlexNetworkAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccFlexNetworkAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccFlexNetworkAttachmentsResponse")
	}
	return
}

// listCccFlexNetworkAttachments implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccFlexNetworkAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccFlexNetworkAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccFlexNetworkAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetworkAttachmentCollection/ListCccFlexNetworkAttachments"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccFlexNetworkAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccFlexNetworks Returns a list of Compute Cloud@Customer flexNetworks.
// A default retry strategy applies to this operation ListCccFlexNetworks()
func (client ComputeCloudAtCustomerClient) ListCccFlexNetworks(ctx context.Context, request ListCccFlexNetworksRequest) (response ListCccFlexNetworksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccFlexNetworks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccFlexNetworksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccFlexNetworksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccFlexNetworksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccFlexNetworksResponse")
	}
	return
}

// listCccFlexNetworks implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccFlexNetworks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccFlexNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccFlexNetworksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetworkCollection/ListCccFlexNetworks"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccFlexNetworks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccInfrastructures Returns a list of Compute Cloud@Customer infrastructures.
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

// ListCccListings Gets a list of listings from Compute Cloud@Customer Marketplace by searching keywords and
// filtering according to listing attributes.
// The compartmentId query parameter is required but is used only for authorization to the operation, the
// results of the request will not change based on the compartmentId parameter.
// If you plan to launch an instance from an image listing on a Compute Cloud@Customer infrastructure,
// you must create a provisioned package from the listing and a specific package version. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Creating a provisioned package requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetCccPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/)
// To get the image ID to launch an instance, issue a CreateCccProvisionedPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// Once the asynchronous provisioning operation is complete, call the GetCccProvisionedPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// Lastly, to launch the instance on the Compute Cloud@Customer infrastructure directly,
// use the image ID of the listing resource version to issue the LaunchInstance (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// A default retry strategy applies to this operation ListCccListings()
func (client ComputeCloudAtCustomerClient) ListCccListings(ctx context.Context, request ListCccListingsRequest) (response ListCccListingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccListings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccListingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccListingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccListingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccListingsResponse")
	}
	return
}

// listCccListings implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccListings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccListings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccListingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccListingCollection/ListCccListings"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccListings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccPackages Gets the list of packages for a listing.
// The compartmentId query parameter is required but is used only for authorization to the operation, the
// results of the request will not change based on the compartmentId parameter.
// If you plan to launch an instance from an image listing on a Compute Cloud@Customer infrastructure,
// you must create a provisioned package from the listing and a specific package version. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Creating a provisioned package requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetCccPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/)
// To get the image ID to launch an instance, issue a CreateCccProvisionedPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// Once the asynchronous provisioning operation is complete, call the GetCccProvisionedPackage (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// Lastly, to launch the instance on the Compute Cloud@Customer infrastructure directly,
// use the image ID of the listing resource version to issue the LaunchInstance (https://docs.oracle.com/en-us/iaas/api/#/en/compute-cloud-at-customer/20221208/CccInfrastructure/) API call.
// A default retry strategy applies to this operation ListCccPackages()
func (client ComputeCloudAtCustomerClient) ListCccPackages(ctx context.Context, request ListCccPackagesRequest) (response ListCccPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccPackagesResponse")
	}
	return
}

// listCccPackages implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccPackageCollection/ListCccPackages"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccProvisionedPackages Returns a list of Compute Cloud@Customer marketplace provisioned packages.
// A default retry strategy applies to this operation ListCccProvisionedPackages()
func (client ComputeCloudAtCustomerClient) ListCccProvisionedPackages(ctx context.Context, request ListCccProvisionedPackagesRequest) (response ListCccProvisionedPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCccProvisionedPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCccProvisionedPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCccProvisionedPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCccProvisionedPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCccProvisionedPackagesResponse")
	}
	return
}

// listCccProvisionedPackages implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) listCccProvisionedPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cccProvisionedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCccProvisionedPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccProvisionedPackageCollection/ListCccProvisionedPackages"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListCccProvisionedPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCccUpgradeSchedules Returns a list of Compute Cloud@Customer upgrade schedules.
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

// ListWorkRequestErrors Lists the errors for a work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ComputeCloudAtCustomerClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client ComputeCloudAtCustomerClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs for a work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ComputeCloudAtCustomerClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client ComputeCloudAtCustomerClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ComputeCloudAtCustomerClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client ComputeCloudAtCustomerClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCccFlexNetwork Updates Compute Cloud@Customer flexNetwork resource.
func (client ComputeCloudAtCustomerClient) UpdateCccFlexNetwork(ctx context.Context, request UpdateCccFlexNetworkRequest) (response UpdateCccFlexNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCccFlexNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCccFlexNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCccFlexNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCccFlexNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCccFlexNetworkResponse")
	}
	return
}

// updateCccFlexNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) updateCccFlexNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cccFlexNetworks/{cccFlexNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCccFlexNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetwork/UpdateCccFlexNetwork"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "UpdateCccFlexNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCccFlexNetworkAttachment Updates Compute Cloud@Customer flexNetworkAttachment resource.
func (client ComputeCloudAtCustomerClient) UpdateCccFlexNetworkAttachment(ctx context.Context, request UpdateCccFlexNetworkAttachmentRequest) (response UpdateCccFlexNetworkAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCccFlexNetworkAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCccFlexNetworkAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCccFlexNetworkAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCccFlexNetworkAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCccFlexNetworkAttachmentResponse")
	}
	return
}

// updateCccFlexNetworkAttachment implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) updateCccFlexNetworkAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cccFlexNetworkAttachments/{cccFlexNetworkAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCccFlexNetworkAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccFlexNetworkAttachment/UpdateCccFlexNetworkAttachment"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "UpdateCccFlexNetworkAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCccInfrastructure Updates Compute Cloud@Customer infrastructure resource.
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

// UpdateCccProvisionedPackage Updates the Compute Cloud@Customer marketplace provisioned package.
func (client ComputeCloudAtCustomerClient) UpdateCccProvisionedPackage(ctx context.Context, request UpdateCccProvisionedPackageRequest) (response UpdateCccProvisionedPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCccProvisionedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCccProvisionedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCccProvisionedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCccProvisionedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCccProvisionedPackageResponse")
	}
	return
}

// updateCccProvisionedPackage implements the OCIOperation interface (enables retrying operations)
func (client ComputeCloudAtCustomerClient) updateCccProvisionedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/cccProvisionedPackages/{cccProvisionedPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCccProvisionedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/compute-cloud-at-customer/20221208/CccProvisionedPackage/UpdateCccProvisionedPackage"
		err = common.PostProcessServiceError(err, "ComputeCloudAtCustomer", "UpdateCccProvisionedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCccUpgradeSchedule Updates the Compute Cloud@Customer upgrade schedule.
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
