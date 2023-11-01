// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ScheduledJobClient a client for ScheduledJob
type ScheduledJobClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewScheduledJobClientWithConfigurationProvider Creates a new default ScheduledJob client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewScheduledJobClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ScheduledJobClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagementhub"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newScheduledJobClientFromBaseClient(baseClient, provider)
}

// NewScheduledJobClientWithOboToken Creates a new default ScheduledJob client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewScheduledJobClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ScheduledJobClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newScheduledJobClientFromBaseClient(baseClient, configProvider)
}

func newScheduledJobClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ScheduledJobClient, err error) {
	// ScheduledJob service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ScheduledJob"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ScheduledJobClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ScheduledJobClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ScheduledJobClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ScheduledJobClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateScheduledJob Creates a new scheduled job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateScheduledJob.go.html to see an example of how to use CreateScheduledJob API.
// A default retry strategy applies to this operation CreateScheduledJob()
func (client ScheduledJobClient) CreateScheduledJob(ctx context.Context, request CreateScheduledJobRequest) (response CreateScheduledJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduledJobResponse")
	}
	return
}

// createScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client ScheduledJobClient) createScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ScheduledJob/CreateScheduledJob"
		err = common.PostProcessServiceError(err, "ScheduledJob", "CreateScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteScheduledJob Deletes the specified scheduled job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteScheduledJob.go.html to see an example of how to use DeleteScheduledJob API.
// A default retry strategy applies to this operation DeleteScheduledJob()
func (client ScheduledJobClient) DeleteScheduledJob(ctx context.Context, request DeleteScheduledJobRequest) (response DeleteScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduledJobResponse")
	}
	return
}

// deleteScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client ScheduledJobClient) deleteScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/scheduledJobs/{scheduledJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ScheduledJob/DeleteScheduledJob"
		err = common.PostProcessServiceError(err, "ScheduledJob", "DeleteScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledJob Gets information about the specified scheduled job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetScheduledJob.go.html to see an example of how to use GetScheduledJob API.
// A default retry strategy applies to this operation GetScheduledJob()
func (client ScheduledJobClient) GetScheduledJob(ctx context.Context, request GetScheduledJobRequest) (response GetScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledJobResponse")
	}
	return
}

// getScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client ScheduledJobClient) getScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs/{scheduledJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ScheduledJob/GetScheduledJob"
		err = common.PostProcessServiceError(err, "ScheduledJob", "GetScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledJobs Lists scheduled jobs that match the specified compartment or scheduled job OCID.
// Filter the list against a variety of criteria including but not limited to its display name,
// lifecycle state, operation type, and schedule type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListScheduledJobs.go.html to see an example of how to use ListScheduledJobs API.
// A default retry strategy applies to this operation ListScheduledJobs()
func (client ScheduledJobClient) ListScheduledJobs(ctx context.Context, request ListScheduledJobsRequest) (response ListScheduledJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledJobsResponse")
	}
	return
}

// listScheduledJobs implements the OCIOperation interface (enables retrying operations)
func (client ScheduledJobClient) listScheduledJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ScheduledJob/ListScheduledJobs"
		err = common.PostProcessServiceError(err, "ScheduledJob", "ListScheduledJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunScheduledJobNow Triggers an already created RECURRING scheduled job to run immediately instead of waiting
// for its next regularly scheduled time. This operation does not support ONETIME scheduled job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RunScheduledJobNow.go.html to see an example of how to use RunScheduledJobNow API.
// A default retry strategy applies to this operation RunScheduledJobNow()
func (client ScheduledJobClient) RunScheduledJobNow(ctx context.Context, request RunScheduledJobNowRequest) (response RunScheduledJobNowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.runScheduledJobNow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunScheduledJobNowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunScheduledJobNowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunScheduledJobNowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunScheduledJobNowResponse")
	}
	return
}

// runScheduledJobNow implements the OCIOperation interface (enables retrying operations)
func (client ScheduledJobClient) runScheduledJobNow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledJobs/{scheduledJobId}/actions/runNow", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunScheduledJobNowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ScheduledJob/RunScheduledJobNow"
		err = common.PostProcessServiceError(err, "ScheduledJob", "RunScheduledJobNow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateScheduledJob Updates the specified scheduled job's name, description, and other details, such as next execution and recurrence.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateScheduledJob.go.html to see an example of how to use UpdateScheduledJob API.
// A default retry strategy applies to this operation UpdateScheduledJob()
func (client ScheduledJobClient) UpdateScheduledJob(ctx context.Context, request UpdateScheduledJobRequest) (response UpdateScheduledJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateScheduledJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduledJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduledJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduledJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduledJobResponse")
	}
	return
}

// updateScheduledJob implements the OCIOperation interface (enables retrying operations)
func (client ScheduledJobClient) updateScheduledJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/scheduledJobs/{scheduledJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduledJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ScheduledJob/UpdateScheduledJob"
		err = common.PostProcessServiceError(err, "ScheduledJob", "UpdateScheduledJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
