// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//DataSafeClient a client for DataSafe
type DataSafeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataSafeClientWithConfigurationProvider Creates a new default DataSafe client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataSafeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataSafeClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDataSafeClientFromBaseClient(baseClient, provider)
}

// NewDataSafeClientWithOboToken Creates a new default DataSafe client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataSafeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataSafeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataSafeClientFromBaseClient(baseClient, configProvider)
}

func newDataSafeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataSafeClient, err error) {
	// DataSafe service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataSafeClient{BaseClient: baseClient}
	client.BasePath = "20181201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataSafeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datasafe", "https://datasafe.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataSafeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataSafeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateTargetDatabase Reactivates a previously deactivated Data Safe target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ActivateTargetDatabase.go.html to see an example of how to use ActivateTargetDatabase API.
func (client DataSafeClient) ActivateTargetDatabase(ctx context.Context, request ActivateTargetDatabaseRequest) (response ActivateTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateTargetDatabaseResponse")
	}
	return
}

// activateTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) activateTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateTargetDatabaseResponse
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

// AddMaskingColumnsFromSdm Adds columns to the specified masking policy from the associated sensitive data model. It
// automatically pulls all the sensitive columns and their relationships from the sensitive
// data model and uses this information to create columns in the masking policy. It also assigns
// default masking formats to these columns based on the associated sensitive types.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/AddMaskingColumnsFromSdm.go.html to see an example of how to use AddMaskingColumnsFromSdm API.
func (client DataSafeClient) AddMaskingColumnsFromSdm(ctx context.Context, request AddMaskingColumnsFromSdmRequest) (response AddMaskingColumnsFromSdmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addMaskingColumnsFromSdm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMaskingColumnsFromSdmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMaskingColumnsFromSdmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMaskingColumnsFromSdmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMaskingColumnsFromSdmResponse")
	}
	return
}

// addMaskingColumnsFromSdm implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) addMaskingColumnsFromSdm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/addMaskingColumnsFromSdm", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMaskingColumnsFromSdmResponse
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

// ApplyDiscoveryJobResults Applies the results of a discovery job to the specified sensitive data model. Note that the plannedAction attribute
// of discovery results is used for processing them. You should first use PatchDiscoveryJobResults to set the plannedAction
// attribute of the discovery results you want to process. ApplyDiscoveryJobResults automatically reads the plannedAction
// attribute and updates the sensitive data model to reflect the actions you planned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ApplyDiscoveryJobResults.go.html to see an example of how to use ApplyDiscoveryJobResults API.
func (client DataSafeClient) ApplyDiscoveryJobResults(ctx context.Context, request ApplyDiscoveryJobResultsRequest) (response ApplyDiscoveryJobResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.applyDiscoveryJobResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApplyDiscoveryJobResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApplyDiscoveryJobResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApplyDiscoveryJobResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApplyDiscoveryJobResultsResponse")
	}
	return
}

// applyDiscoveryJobResults implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) applyDiscoveryJobResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns/actions/applyDiscoveryJobResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApplyDiscoveryJobResultsResponse
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

// CalculateAuditVolumeAvailable Calculates the volume of audit events available on the target database to be collected. Measurable up to the defined retention period of the audit target resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CalculateAuditVolumeAvailable.go.html to see an example of how to use CalculateAuditVolumeAvailable API.
func (client DataSafeClient) CalculateAuditVolumeAvailable(ctx context.Context, request CalculateAuditVolumeAvailableRequest) (response CalculateAuditVolumeAvailableResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.calculateAuditVolumeAvailable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CalculateAuditVolumeAvailableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CalculateAuditVolumeAvailableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CalculateAuditVolumeAvailableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CalculateAuditVolumeAvailableResponse")
	}
	return
}

// calculateAuditVolumeAvailable implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) calculateAuditVolumeAvailable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditProfiles/{auditProfileId}/actions/calculateAuditVolumeAvailable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CalculateAuditVolumeAvailableResponse
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

// CalculateAuditVolumeCollected Calculates the volume of audit events collected by data safe.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CalculateAuditVolumeCollected.go.html to see an example of how to use CalculateAuditVolumeCollected API.
func (client DataSafeClient) CalculateAuditVolumeCollected(ctx context.Context, request CalculateAuditVolumeCollectedRequest) (response CalculateAuditVolumeCollectedResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.calculateAuditVolumeCollected, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CalculateAuditVolumeCollectedResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CalculateAuditVolumeCollectedResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CalculateAuditVolumeCollectedResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CalculateAuditVolumeCollectedResponse")
	}
	return
}

// calculateAuditVolumeCollected implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) calculateAuditVolumeCollected(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditProfiles/{auditProfileId}/actions/calculateAuditVolumeCollected", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CalculateAuditVolumeCollectedResponse
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

// CancelWorkRequest Cancel the given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
func (client DataSafeClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
func (client DataSafeClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAlertCompartment Moves the specified alert into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAlertCompartment.go.html to see an example of how to use ChangeAlertCompartment API.
func (client DataSafeClient) ChangeAlertCompartment(ctx context.Context, request ChangeAlertCompartmentRequest) (response ChangeAlertCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAlertCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAlertCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAlertCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAlertCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAlertCompartmentResponse")
	}
	return
}

// changeAlertCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeAlertCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alerts/{alertId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAlertCompartmentResponse
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

// ChangeAuditArchiveRetrievalCompartment Moves the archive retreival to the specified compartment. When provided, if-Match is checked against ETag value of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAuditArchiveRetrievalCompartment.go.html to see an example of how to use ChangeAuditArchiveRetrievalCompartment API.
func (client DataSafeClient) ChangeAuditArchiveRetrievalCompartment(ctx context.Context, request ChangeAuditArchiveRetrievalCompartmentRequest) (response ChangeAuditArchiveRetrievalCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAuditArchiveRetrievalCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAuditArchiveRetrievalCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAuditArchiveRetrievalCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAuditArchiveRetrievalCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAuditArchiveRetrievalCompartmentResponse")
	}
	return
}

// changeAuditArchiveRetrievalCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeAuditArchiveRetrievalCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditArchiveRetrievals/{auditArchiveRetrievalId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAuditArchiveRetrievalCompartmentResponse
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

// ChangeAuditPolicyCompartment Moves the specified audit policy and its dependent resources into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAuditPolicyCompartment.go.html to see an example of how to use ChangeAuditPolicyCompartment API.
func (client DataSafeClient) ChangeAuditPolicyCompartment(ctx context.Context, request ChangeAuditPolicyCompartmentRequest) (response ChangeAuditPolicyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAuditPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAuditPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAuditPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAuditPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAuditPolicyCompartmentResponse")
	}
	return
}

// changeAuditPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeAuditPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditPolicies/{auditPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAuditPolicyCompartmentResponse
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

// ChangeAuditProfileCompartment Moves the specified audit profile and its dependent resources into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAuditProfileCompartment.go.html to see an example of how to use ChangeAuditProfileCompartment API.
func (client DataSafeClient) ChangeAuditProfileCompartment(ctx context.Context, request ChangeAuditProfileCompartmentRequest) (response ChangeAuditProfileCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAuditProfileCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAuditProfileCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAuditProfileCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAuditProfileCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAuditProfileCompartmentResponse")
	}
	return
}

// changeAuditProfileCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeAuditProfileCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditProfiles/{auditProfileId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAuditProfileCompartmentResponse
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

// ChangeDataSafePrivateEndpointCompartment Moves the Data Safe private endpoint and its dependent resources to the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeDataSafePrivateEndpointCompartment.go.html to see an example of how to use ChangeDataSafePrivateEndpointCompartment API.
func (client DataSafeClient) ChangeDataSafePrivateEndpointCompartment(ctx context.Context, request ChangeDataSafePrivateEndpointCompartmentRequest) (response ChangeDataSafePrivateEndpointCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDataSafePrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDataSafePrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDataSafePrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDataSafePrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDataSafePrivateEndpointCompartmentResponse")
	}
	return
}

// changeDataSafePrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeDataSafePrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDataSafePrivateEndpointCompartmentResponse
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

// ChangeDiscoveryJobCompartment Moves the specified discovery job and its dependent resources into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeDiscoveryJobCompartment.go.html to see an example of how to use ChangeDiscoveryJobCompartment API.
func (client DataSafeClient) ChangeDiscoveryJobCompartment(ctx context.Context, request ChangeDiscoveryJobCompartmentRequest) (response ChangeDiscoveryJobCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDiscoveryJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDiscoveryJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDiscoveryJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDiscoveryJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDiscoveryJobCompartmentResponse")
	}
	return
}

// changeDiscoveryJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeDiscoveryJobCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/discoveryJobs/{discoveryJobId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDiscoveryJobCompartmentResponse
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

// ChangeLibraryMaskingFormatCompartment Moves the specified library masking format into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeLibraryMaskingFormatCompartment.go.html to see an example of how to use ChangeLibraryMaskingFormatCompartment API.
func (client DataSafeClient) ChangeLibraryMaskingFormatCompartment(ctx context.Context, request ChangeLibraryMaskingFormatCompartmentRequest) (response ChangeLibraryMaskingFormatCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeLibraryMaskingFormatCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLibraryMaskingFormatCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLibraryMaskingFormatCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLibraryMaskingFormatCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLibraryMaskingFormatCompartmentResponse")
	}
	return
}

// changeLibraryMaskingFormatCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeLibraryMaskingFormatCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/libraryMaskingFormats/{libraryMaskingFormatId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeLibraryMaskingFormatCompartmentResponse
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

// ChangeMaskingPolicyCompartment Moves the specified masking policy and its dependent resources into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeMaskingPolicyCompartment.go.html to see an example of how to use ChangeMaskingPolicyCompartment API.
func (client DataSafeClient) ChangeMaskingPolicyCompartment(ctx context.Context, request ChangeMaskingPolicyCompartmentRequest) (response ChangeMaskingPolicyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMaskingPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMaskingPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMaskingPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMaskingPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMaskingPolicyCompartmentResponse")
	}
	return
}

// changeMaskingPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeMaskingPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMaskingPolicyCompartmentResponse
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

// ChangeOnPremConnectorCompartment Moves the specified on-premises connector into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeOnPremConnectorCompartment.go.html to see an example of how to use ChangeOnPremConnectorCompartment API.
func (client DataSafeClient) ChangeOnPremConnectorCompartment(ctx context.Context, request ChangeOnPremConnectorCompartmentRequest) (response ChangeOnPremConnectorCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOnPremConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOnPremConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOnPremConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOnPremConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOnPremConnectorCompartmentResponse")
	}
	return
}

// changeOnPremConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeOnPremConnectorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/onPremConnectors/{onPremConnectorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOnPremConnectorCompartmentResponse
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

// ChangeReportCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeReportCompartment.go.html to see an example of how to use ChangeReportCompartment API.
func (client DataSafeClient) ChangeReportCompartment(ctx context.Context, request ChangeReportCompartmentRequest) (response ChangeReportCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeReportCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeReportCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeReportCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeReportCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeReportCompartmentResponse")
	}
	return
}

// changeReportCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeReportCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reports/{reportId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeReportCompartmentResponse
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

// ChangeReportDefinitionCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeReportDefinitionCompartment.go.html to see an example of how to use ChangeReportDefinitionCompartment API.
func (client DataSafeClient) ChangeReportDefinitionCompartment(ctx context.Context, request ChangeReportDefinitionCompartmentRequest) (response ChangeReportDefinitionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeReportDefinitionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeReportDefinitionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeReportDefinitionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeReportDefinitionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeReportDefinitionCompartmentResponse")
	}
	return
}

// changeReportDefinitionCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeReportDefinitionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reportDefinitions/{reportDefinitionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeReportDefinitionCompartmentResponse
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

// ChangeRetention Change the online and offline months .
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeRetention.go.html to see an example of how to use ChangeRetention API.
func (client DataSafeClient) ChangeRetention(ctx context.Context, request ChangeRetentionRequest) (response ChangeRetentionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRetention, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRetentionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRetentionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRetentionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRetentionResponse")
	}
	return
}

// changeRetention implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeRetention(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditProfiles/{auditProfileId}/actions/changeRetention", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRetentionResponse
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

// ChangeSecurityAssessmentCompartment Moves the specified saved security assessment or future scheduled assessments into a different compartment.
// To start, call first the operation ListSecurityAssessments with filters "type = save_schedule". This returns the scheduleAssessmentId. Then, call this changeCompartment with the scheduleAssessmentId.
// The existing saved security assessments created due to the schedule are not moved. However, all new saves will be associated with the new compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSecurityAssessmentCompartment.go.html to see an example of how to use ChangeSecurityAssessmentCompartment API.
func (client DataSafeClient) ChangeSecurityAssessmentCompartment(ctx context.Context, request ChangeSecurityAssessmentCompartmentRequest) (response ChangeSecurityAssessmentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityAssessmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityAssessmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityAssessmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityAssessmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityAssessmentCompartmentResponse")
	}
	return
}

// changeSecurityAssessmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSecurityAssessmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityAssessmentCompartmentResponse
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

// ChangeSensitiveDataModelCompartment Moves the specified sensitive data model and its dependent resources into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSensitiveDataModelCompartment.go.html to see an example of how to use ChangeSensitiveDataModelCompartment API.
func (client DataSafeClient) ChangeSensitiveDataModelCompartment(ctx context.Context, request ChangeSensitiveDataModelCompartmentRequest) (response ChangeSensitiveDataModelCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSensitiveDataModelCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSensitiveDataModelCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSensitiveDataModelCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSensitiveDataModelCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSensitiveDataModelCompartmentResponse")
	}
	return
}

// changeSensitiveDataModelCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSensitiveDataModelCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSensitiveDataModelCompartmentResponse
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

// ChangeSensitiveTypeCompartment Moves the specified sensitive type into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSensitiveTypeCompartment.go.html to see an example of how to use ChangeSensitiveTypeCompartment API.
func (client DataSafeClient) ChangeSensitiveTypeCompartment(ctx context.Context, request ChangeSensitiveTypeCompartmentRequest) (response ChangeSensitiveTypeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSensitiveTypeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSensitiveTypeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSensitiveTypeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSensitiveTypeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSensitiveTypeCompartmentResponse")
	}
	return
}

// changeSensitiveTypeCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSensitiveTypeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveTypes/{sensitiveTypeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSensitiveTypeCompartmentResponse
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

// ChangeTargetAlertPolicyAssociationCompartment Moves the specified target-alert policy Association into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeTargetAlertPolicyAssociationCompartment.go.html to see an example of how to use ChangeTargetAlertPolicyAssociationCompartment API.
func (client DataSafeClient) ChangeTargetAlertPolicyAssociationCompartment(ctx context.Context, request ChangeTargetAlertPolicyAssociationCompartmentRequest) (response ChangeTargetAlertPolicyAssociationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTargetAlertPolicyAssociationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTargetAlertPolicyAssociationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTargetAlertPolicyAssociationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTargetAlertPolicyAssociationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTargetAlertPolicyAssociationCompartmentResponse")
	}
	return
}

// changeTargetAlertPolicyAssociationCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeTargetAlertPolicyAssociationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetAlertPolicyAssociations/{targetAlertPolicyAssociationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTargetAlertPolicyAssociationCompartmentResponse
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

// ChangeTargetDatabaseCompartment Moves the Data Safe target database to the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeTargetDatabaseCompartment.go.html to see an example of how to use ChangeTargetDatabaseCompartment API.
func (client DataSafeClient) ChangeTargetDatabaseCompartment(ctx context.Context, request ChangeTargetDatabaseCompartmentRequest) (response ChangeTargetDatabaseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTargetDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTargetDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTargetDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTargetDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTargetDatabaseCompartmentResponse")
	}
	return
}

// changeTargetDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeTargetDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTargetDatabaseCompartmentResponse
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

// ChangeUserAssessmentCompartment Moves the specified saved user assessment or future scheduled assessments into a different compartment.
// To start storing scheduled user assessments on a different compartment, first call the operation ListUserAssessments with
// the filters "type = save_schedule". That call returns the scheduleAssessmentId. Then call
// ChangeUserAssessmentCompartment with the scheduleAssessmentId. The existing saved user assessments created per the schedule
// are not be moved. However, all new saves will be associated with the new compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeUserAssessmentCompartment.go.html to see an example of how to use ChangeUserAssessmentCompartment API.
func (client DataSafeClient) ChangeUserAssessmentCompartment(ctx context.Context, request ChangeUserAssessmentCompartmentRequest) (response ChangeUserAssessmentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeUserAssessmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeUserAssessmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeUserAssessmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeUserAssessmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeUserAssessmentCompartmentResponse")
	}
	return
}

// changeUserAssessmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeUserAssessmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeUserAssessmentCompartmentResponse
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

// CompareSecurityAssessment Compares two security assessments. For this comparison, a security assessment can be a saved assessment, a latest assessment, or a baseline assessment.
// For example, you can compare saved assessment or a latest assessment against a baseline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CompareSecurityAssessment.go.html to see an example of how to use CompareSecurityAssessment API.
func (client DataSafeClient) CompareSecurityAssessment(ctx context.Context, request CompareSecurityAssessmentRequest) (response CompareSecurityAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.compareSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompareSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompareSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompareSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompareSecurityAssessmentResponse")
	}
	return
}

// compareSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) compareSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/compare", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompareSecurityAssessmentResponse
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

// CompareUserAssessment Compares two user assessments. For this comparison, a user assessment can be a saved, a latest assessment, or a baseline.
// As an example, it can be used to compare a user assessment saved or a latest assessment with a baseline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CompareUserAssessment.go.html to see an example of how to use CompareUserAssessment API.
func (client DataSafeClient) CompareUserAssessment(ctx context.Context, request CompareUserAssessmentRequest) (response CompareUserAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.compareUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompareUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompareUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompareUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompareUserAssessmentResponse")
	}
	return
}

// compareUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) compareUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/compare", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompareUserAssessmentResponse
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

// CreateAuditArchiveRetrieval Creates a work request to retrieve archived audit data. This asynchronous process will usually take over an hour to complete.
// Save the id from the response of this operation. Call GetAuditArchiveRetrieval operation after an hour, passing the id to know the status of
// this operation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateAuditArchiveRetrieval.go.html to see an example of how to use CreateAuditArchiveRetrieval API.
func (client DataSafeClient) CreateAuditArchiveRetrieval(ctx context.Context, request CreateAuditArchiveRetrievalRequest) (response CreateAuditArchiveRetrievalResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAuditArchiveRetrieval, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAuditArchiveRetrievalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAuditArchiveRetrievalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAuditArchiveRetrievalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAuditArchiveRetrievalResponse")
	}
	return
}

// createAuditArchiveRetrieval implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createAuditArchiveRetrieval(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditArchiveRetrievals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAuditArchiveRetrievalResponse
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

// CreateDataSafePrivateEndpoint Creates a new Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateDataSafePrivateEndpoint.go.html to see an example of how to use CreateDataSafePrivateEndpoint API.
func (client DataSafeClient) CreateDataSafePrivateEndpoint(ctx context.Context, request CreateDataSafePrivateEndpointRequest) (response CreateDataSafePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataSafePrivateEndpointResponse")
	}
	return
}

// createDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSafePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataSafePrivateEndpointResponse
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

// CreateDiscoveryJob Performs incremental data discovery for the specified sensitive data model. It uses the target database associated
// with the sensitive data model.
// After performing data discovery, you can use ListDiscoveryJobResults to view the discovery results, PatchDiscoveryJobResults
// to specify the action you want perform on these results, and then ApplyDiscoveryJobResults to process the results
// and apply them to the sensitive data model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateDiscoveryJob.go.html to see an example of how to use CreateDiscoveryJob API.
func (client DataSafeClient) CreateDiscoveryJob(ctx context.Context, request CreateDiscoveryJobRequest) (response CreateDiscoveryJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDiscoveryJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDiscoveryJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDiscoveryJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDiscoveryJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDiscoveryJobResponse")
	}
	return
}

// createDiscoveryJob implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createDiscoveryJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/discoveryJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDiscoveryJobResponse
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

// CreateLibraryMaskingFormat Creates a new library masking format. A masking format can have one or more
// format entries. The combined output of all the format entries is used for masking.
// It provides the flexibility to define a masking format that can generate different
// parts of a data value separately and then combine them to get the final data value
// for masking. Note that you cannot define masking condition in a library masking format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateLibraryMaskingFormat.go.html to see an example of how to use CreateLibraryMaskingFormat API.
func (client DataSafeClient) CreateLibraryMaskingFormat(ctx context.Context, request CreateLibraryMaskingFormatRequest) (response CreateLibraryMaskingFormatResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createLibraryMaskingFormat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLibraryMaskingFormatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLibraryMaskingFormatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLibraryMaskingFormatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLibraryMaskingFormatResponse")
	}
	return
}

// createLibraryMaskingFormat implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createLibraryMaskingFormat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/libraryMaskingFormats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLibraryMaskingFormatResponse
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

// CreateMaskingColumn Creates a new masking column in the specified masking policy. Use this operation
// to add parent columns only. It automatically adds the child columns from the
// associated sensitive data model or target database. If you provide the
// sensitiveTypeId attribute but not the maskingFormats attribute, it automatically
// assigns the default masking format associated with the specified sensitive type.
// Alternatively, if you provide the maskingFormats attribute, the specified masking
// formats are assigned to the column.
// Using the maskingFormats attribute, you can assign one or more masking formats
// to a column. You need to specify a condition as part of each masking format. It
// enables you to do <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>
// so that you can mask the column data values differently using different
// masking conditions. A masking format can have one or more format entries. The
// combined output of all the format entries is used for masking. It provides the
// flexibility to define a masking format that can generate different parts of a data
// value separately and then combine them to get the final data value for masking.
// You can use the maskingColumnGroup attribute to group the columns that you would
// like to mask together. It enables you to do <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037">group or compound masking</a> that ensures that the
// masked data across the columns in a group continue to retain the same logical relationship.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateMaskingColumn.go.html to see an example of how to use CreateMaskingColumn API.
func (client DataSafeClient) CreateMaskingColumn(ctx context.Context, request CreateMaskingColumnRequest) (response CreateMaskingColumnResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMaskingColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMaskingColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMaskingColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMaskingColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMaskingColumnResponse")
	}
	return
}

// createMaskingColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createMaskingColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/maskingColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMaskingColumnResponse
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

// CreateMaskingPolicy Creates a new masking policy and associates it with a sensitive data model or a reference target database.
// To use a sensitive data model as the source of masking columns, set the columnSource attribute to
// SENSITIVE_DATA_MODEL and provide the sensitiveDataModelId attribute. After creating a masking policy,
// you can use the AddMaskingColumnsFromSdm operation to automatically add all the columns from
// the associated sensitive data model. In this case, the target database associated with the
// sensitive data model is used for column and masking format validations.
// You can also create a masking policy without using a sensitive data model. In this case,
// you need to associate your masking policy with a target database by setting the columnSource
// attribute to TARGET and providing the targetId attribute. The specified target database
// is used for column and masking format validations.
// After creating a masking policy, you can use the CreateMaskingColumn or PatchMaskingColumns
// operation to manually add columns to the policy. You need to add the parent columns only,
// and it automatically adds the child columns (in referential relationship with the parent columns)
// from the associated sensitive data model or target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateMaskingPolicy.go.html to see an example of how to use CreateMaskingPolicy API.
func (client DataSafeClient) CreateMaskingPolicy(ctx context.Context, request CreateMaskingPolicyRequest) (response CreateMaskingPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMaskingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMaskingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMaskingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMaskingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMaskingPolicyResponse")
	}
	return
}

// createMaskingPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createMaskingPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMaskingPolicyResponse
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

// CreateOnPremConnector Creates a new on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateOnPremConnector.go.html to see an example of how to use CreateOnPremConnector API.
func (client DataSafeClient) CreateOnPremConnector(ctx context.Context, request CreateOnPremConnectorRequest) (response CreateOnPremConnectorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOnPremConnectorResponse")
	}
	return
}

// createOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/onPremConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOnPremConnectorResponse
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

// CreateReportDefinition Creates a new report definition with parameters specified in the body. The report definition is stored in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateReportDefinition.go.html to see an example of how to use CreateReportDefinition API.
func (client DataSafeClient) CreateReportDefinition(ctx context.Context, request CreateReportDefinitionRequest) (response CreateReportDefinitionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createReportDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateReportDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateReportDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateReportDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateReportDefinitionResponse")
	}
	return
}

// createReportDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createReportDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reportDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateReportDefinitionResponse
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

// CreateSecurityAssessment Creates a new saved security assessment for one or multiple targets in a compartment. When this operation is performed,
// it will save the latest assessments in the specified compartment. If a schedule is passed, it will persist the latest assessments,
// at the defined date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSecurityAssessment.go.html to see an example of how to use CreateSecurityAssessment API.
func (client DataSafeClient) CreateSecurityAssessment(ctx context.Context, request CreateSecurityAssessmentRequest) (response CreateSecurityAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityAssessmentResponse")
	}
	return
}

// createSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityAssessmentResponse
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

// CreateSensitiveColumn Creates a new sensitive column in the specified sensitive data model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSensitiveColumn.go.html to see an example of how to use CreateSensitiveColumn API.
func (client DataSafeClient) CreateSensitiveColumn(ctx context.Context, request CreateSensitiveColumnRequest) (response CreateSensitiveColumnResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSensitiveColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSensitiveColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSensitiveColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSensitiveColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSensitiveColumnResponse")
	}
	return
}

// createSensitiveColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSensitiveColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSensitiveColumnResponse
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

// CreateSensitiveDataModel Creates a new sensitive data model. If schemas and sensitive types are provided, it automatically runs data discovery
// and adds the discovered columns to the sensitive data model. Otherwise, it creates an empty sensitive data model
// that can be updated later.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSensitiveDataModel.go.html to see an example of how to use CreateSensitiveDataModel API.
func (client DataSafeClient) CreateSensitiveDataModel(ctx context.Context, request CreateSensitiveDataModelRequest) (response CreateSensitiveDataModelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSensitiveDataModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSensitiveDataModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSensitiveDataModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSensitiveDataModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSensitiveDataModelResponse")
	}
	return
}

// createSensitiveDataModel implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSensitiveDataModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSensitiveDataModelResponse
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

// CreateSensitiveType Creates a new sensitive type, which can be a basic sensitive type with regular expressions or a sensitive category.
// While sensitive types are used for data discovery, sensitive categories are used for logically grouping the related
// or similar sensitive types.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSensitiveType.go.html to see an example of how to use CreateSensitiveType API.
func (client DataSafeClient) CreateSensitiveType(ctx context.Context, request CreateSensitiveTypeRequest) (response CreateSensitiveTypeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSensitiveType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSensitiveTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSensitiveTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSensitiveTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSensitiveTypeResponse")
	}
	return
}

// createSensitiveType implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSensitiveType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSensitiveTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &sensitivetype{})
	return response, err
}

// CreateTargetAlertPolicyAssociation Creates a new target-alert policy association to track a alert policy applied on target.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateTargetAlertPolicyAssociation.go.html to see an example of how to use CreateTargetAlertPolicyAssociation API.
func (client DataSafeClient) CreateTargetAlertPolicyAssociation(ctx context.Context, request CreateTargetAlertPolicyAssociationRequest) (response CreateTargetAlertPolicyAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTargetAlertPolicyAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetAlertPolicyAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetAlertPolicyAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetAlertPolicyAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetAlertPolicyAssociationResponse")
	}
	return
}

// createTargetAlertPolicyAssociation implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createTargetAlertPolicyAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetAlertPolicyAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTargetAlertPolicyAssociationResponse
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

// CreateTargetDatabase Registers the specified database with Data Safe and creates a Data Safe target database in the Data Safe Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateTargetDatabase.go.html to see an example of how to use CreateTargetDatabase API.
func (client DataSafeClient) CreateTargetDatabase(ctx context.Context, request CreateTargetDatabaseRequest) (response CreateTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetDatabaseResponse")
	}
	return
}

// createTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTargetDatabaseResponse
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

// CreateUserAssessment Creates a new saved user assessment for one or multiple targets in a compartment. It saves the latest assessments in the
// specified compartment. If a scheduled is passed in, this operation persists the latest assessments that exist at the defined
// date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateUserAssessment.go.html to see an example of how to use CreateUserAssessment API.
func (client DataSafeClient) CreateUserAssessment(ctx context.Context, request CreateUserAssessmentRequest) (response CreateUserAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserAssessmentResponse")
	}
	return
}

// createUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserAssessmentResponse
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

// DeactivateTargetDatabase Deactivates a target database in Data Safe.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeactivateTargetDatabase.go.html to see an example of how to use DeactivateTargetDatabase API.
func (client DataSafeClient) DeactivateTargetDatabase(ctx context.Context, request DeactivateTargetDatabaseRequest) (response DeactivateTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deactivateTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateTargetDatabaseResponse")
	}
	return
}

// deactivateTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deactivateTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateTargetDatabaseResponse
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

// DeleteAuditArchiveRetrieval To unload retrieved archive data, call the operation ListAuditArchiveRetrieval first.
// This will return the auditArchiveRetrievalId. Then call this operation with auditArchiveRetrievalId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteAuditArchiveRetrieval.go.html to see an example of how to use DeleteAuditArchiveRetrieval API.
func (client DataSafeClient) DeleteAuditArchiveRetrieval(ctx context.Context, request DeleteAuditArchiveRetrievalRequest) (response DeleteAuditArchiveRetrievalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAuditArchiveRetrieval, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAuditArchiveRetrievalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAuditArchiveRetrievalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAuditArchiveRetrievalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAuditArchiveRetrievalResponse")
	}
	return
}

// deleteAuditArchiveRetrieval implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteAuditArchiveRetrieval(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/auditArchiveRetrievals/{auditArchiveRetrievalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAuditArchiveRetrievalResponse
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

// DeleteAuditTrail Deletes the specified audit trail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteAuditTrail.go.html to see an example of how to use DeleteAuditTrail API.
func (client DataSafeClient) DeleteAuditTrail(ctx context.Context, request DeleteAuditTrailRequest) (response DeleteAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAuditTrail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAuditTrailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAuditTrailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAuditTrailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAuditTrailResponse")
	}
	return
}

// deleteAuditTrail implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteAuditTrail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/auditTrails/{auditTrailId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAuditTrailResponse
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

// DeleteDataSafePrivateEndpoint Deletes the specified Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDataSafePrivateEndpoint.go.html to see an example of how to use DeleteDataSafePrivateEndpoint API.
func (client DataSafeClient) DeleteDataSafePrivateEndpoint(ctx context.Context, request DeleteDataSafePrivateEndpointRequest) (response DeleteDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataSafePrivateEndpointResponse")
	}
	return
}

// deleteDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataSafePrivateEndpointResponse
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

// DeleteDiscoveryJob Deletes the specified discovery job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDiscoveryJob.go.html to see an example of how to use DeleteDiscoveryJob API.
func (client DataSafeClient) DeleteDiscoveryJob(ctx context.Context, request DeleteDiscoveryJobRequest) (response DeleteDiscoveryJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDiscoveryJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDiscoveryJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDiscoveryJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDiscoveryJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDiscoveryJobResponse")
	}
	return
}

// deleteDiscoveryJob implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteDiscoveryJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/discoveryJobs/{discoveryJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDiscoveryJobResponse
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

// DeleteDiscoveryJobResult Deletes the specified discovery result.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDiscoveryJobResult.go.html to see an example of how to use DeleteDiscoveryJobResult API.
func (client DataSafeClient) DeleteDiscoveryJobResult(ctx context.Context, request DeleteDiscoveryJobResultRequest) (response DeleteDiscoveryJobResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDiscoveryJobResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDiscoveryJobResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDiscoveryJobResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDiscoveryJobResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDiscoveryJobResultResponse")
	}
	return
}

// deleteDiscoveryJobResult implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteDiscoveryJobResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/discoveryJobs/{discoveryJobId}/results/{resultKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDiscoveryJobResultResponse
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

// DeleteLibraryMaskingFormat Deletes the specified library masking format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteLibraryMaskingFormat.go.html to see an example of how to use DeleteLibraryMaskingFormat API.
func (client DataSafeClient) DeleteLibraryMaskingFormat(ctx context.Context, request DeleteLibraryMaskingFormatRequest) (response DeleteLibraryMaskingFormatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLibraryMaskingFormat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLibraryMaskingFormatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLibraryMaskingFormatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLibraryMaskingFormatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLibraryMaskingFormatResponse")
	}
	return
}

// deleteLibraryMaskingFormat implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteLibraryMaskingFormat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/libraryMaskingFormats/{libraryMaskingFormatId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLibraryMaskingFormatResponse
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

// DeleteMaskingColumn Deletes the specified masking column.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteMaskingColumn.go.html to see an example of how to use DeleteMaskingColumn API.
func (client DataSafeClient) DeleteMaskingColumn(ctx context.Context, request DeleteMaskingColumnRequest) (response DeleteMaskingColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMaskingColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMaskingColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMaskingColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMaskingColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMaskingColumnResponse")
	}
	return
}

// deleteMaskingColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteMaskingColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/maskingPolicies/{maskingPolicyId}/maskingColumns/{maskingColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMaskingColumnResponse
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

// DeleteMaskingPolicy Deletes the specified masking policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteMaskingPolicy.go.html to see an example of how to use DeleteMaskingPolicy API.
func (client DataSafeClient) DeleteMaskingPolicy(ctx context.Context, request DeleteMaskingPolicyRequest) (response DeleteMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMaskingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMaskingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMaskingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMaskingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMaskingPolicyResponse")
	}
	return
}

// deleteMaskingPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteMaskingPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/maskingPolicies/{maskingPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMaskingPolicyResponse
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

// DeleteOnPremConnector Deletes the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteOnPremConnector.go.html to see an example of how to use DeleteOnPremConnector API.
func (client DataSafeClient) DeleteOnPremConnector(ctx context.Context, request DeleteOnPremConnectorRequest) (response DeleteOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOnPremConnectorResponse")
	}
	return
}

// deleteOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/onPremConnectors/{onPremConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOnPremConnectorResponse
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

// DeleteReportDefinition Deletes the specified report definition. Only the user created report definition can be deleted. The seeded report definitions cannot be deleted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteReportDefinition.go.html to see an example of how to use DeleteReportDefinition API.
func (client DataSafeClient) DeleteReportDefinition(ctx context.Context, request DeleteReportDefinitionRequest) (response DeleteReportDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteReportDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteReportDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteReportDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteReportDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteReportDefinitionResponse")
	}
	return
}

// deleteReportDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteReportDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/reportDefinitions/{reportDefinitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteReportDefinitionResponse
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

// DeleteSecurityAssessment Deletes the specified saved security assessment or schedule. To delete a security assessment schedule,
// first call the operation ListSecurityAssessments with filters "type = save_schedule".
// That operation returns the scheduleAssessmentId. Then, call DeleteSecurityAssessment with the scheduleAssessmentId.
// If the assessment being deleted is the baseline for that compartment, then it will impact all baselines in the compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSecurityAssessment.go.html to see an example of how to use DeleteSecurityAssessment API.
func (client DataSafeClient) DeleteSecurityAssessment(ctx context.Context, request DeleteSecurityAssessmentRequest) (response DeleteSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityAssessmentResponse")
	}
	return
}

// deleteSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/securityAssessments/{securityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityAssessmentResponse
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

// DeleteSensitiveColumn Deletes the specified sensitive column.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSensitiveColumn.go.html to see an example of how to use DeleteSensitiveColumn API.
func (client DataSafeClient) DeleteSensitiveColumn(ctx context.Context, request DeleteSensitiveColumnRequest) (response DeleteSensitiveColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSensitiveColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSensitiveColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSensitiveColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSensitiveColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSensitiveColumnResponse")
	}
	return
}

// deleteSensitiveColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSensitiveColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns/{sensitiveColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSensitiveColumnResponse
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

// DeleteSensitiveDataModel Deletes the specified sensitive data model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSensitiveDataModel.go.html to see an example of how to use DeleteSensitiveDataModel API.
func (client DataSafeClient) DeleteSensitiveDataModel(ctx context.Context, request DeleteSensitiveDataModelRequest) (response DeleteSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSensitiveDataModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSensitiveDataModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSensitiveDataModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSensitiveDataModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSensitiveDataModelResponse")
	}
	return
}

// deleteSensitiveDataModel implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSensitiveDataModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sensitiveDataModels/{sensitiveDataModelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSensitiveDataModelResponse
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

// DeleteSensitiveType Deletes the specified sensitive type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSensitiveType.go.html to see an example of how to use DeleteSensitiveType API.
func (client DataSafeClient) DeleteSensitiveType(ctx context.Context, request DeleteSensitiveTypeRequest) (response DeleteSensitiveTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSensitiveType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSensitiveTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSensitiveTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSensitiveTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSensitiveTypeResponse")
	}
	return
}

// deleteSensitiveType implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSensitiveType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sensitiveTypes/{sensitiveTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSensitiveTypeResponse
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

// DeleteTargetAlertPolicyAssociation Deletes the specified target-alert policy Association.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteTargetAlertPolicyAssociation.go.html to see an example of how to use DeleteTargetAlertPolicyAssociation API.
func (client DataSafeClient) DeleteTargetAlertPolicyAssociation(ctx context.Context, request DeleteTargetAlertPolicyAssociationRequest) (response DeleteTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetAlertPolicyAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetAlertPolicyAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetAlertPolicyAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetAlertPolicyAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetAlertPolicyAssociationResponse")
	}
	return
}

// deleteTargetAlertPolicyAssociation implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteTargetAlertPolicyAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targetAlertPolicyAssociations/{targetAlertPolicyAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTargetAlertPolicyAssociationResponse
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

// DeleteTargetDatabase Deregisters the specified database from Data Safe and removes the target database from the Data Safe Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteTargetDatabase.go.html to see an example of how to use DeleteTargetDatabase API.
func (client DataSafeClient) DeleteTargetDatabase(ctx context.Context, request DeleteTargetDatabaseRequest) (response DeleteTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetDatabaseResponse")
	}
	return
}

// deleteTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targetDatabases/{targetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTargetDatabaseResponse
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

// DeleteUserAssessment Deletes the specified saved user assessment or schedule. To delete a user assessment schedule, first call the operation
// ListUserAssessments with filters "type = save_schedule".
// That call returns the scheduleAssessmentId. Then call DeleteUserAssessment with the scheduleAssessmentId.
// If the assessment being deleted is the baseline for that compartment, then it will impact all baselines in the compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteUserAssessment.go.html to see an example of how to use DeleteUserAssessment API.
func (client DataSafeClient) DeleteUserAssessment(ctx context.Context, request DeleteUserAssessmentRequest) (response DeleteUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserAssessmentResponse")
	}
	return
}

// deleteUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/userAssessments/{userAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserAssessmentResponse
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

// DiscoverAuditTrails Updates the list of audit trails created under audit profile.The
// operation can be used to create new audit trails for target database
// when they become available for audit collection because of change of database version
// or change of database unified mode or change of data base  edition or being deleted previously etc.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DiscoverAuditTrails.go.html to see an example of how to use DiscoverAuditTrails API.
func (client DataSafeClient) DiscoverAuditTrails(ctx context.Context, request DiscoverAuditTrailsRequest) (response DiscoverAuditTrailsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.discoverAuditTrails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DiscoverAuditTrailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DiscoverAuditTrailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DiscoverAuditTrailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DiscoverAuditTrailsResponse")
	}
	return
}

// discoverAuditTrails implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) discoverAuditTrails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditProfiles/{auditProfileId}/actions/discoverAuditTrails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DiscoverAuditTrailsResponse
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

// DownloadDiscoveryReport Downloads an already-generated discovery report. Note that the GenerateDiscoveryReportForDownload operation is a
// prerequisite for the DownloadDiscoveryReport operation. Use GenerateDiscoveryReportForDownload to generate a discovery
// report file and then use DownloadDiscoveryReport to download the generated file. By default, it downloads report for
// all the columns in a sensitive data model. Use the discoveryJobId attribute to download report for a specific discovery job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadDiscoveryReport.go.html to see an example of how to use DownloadDiscoveryReport API.
func (client DataSafeClient) DownloadDiscoveryReport(ctx context.Context, request DownloadDiscoveryReportRequest) (response DownloadDiscoveryReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadDiscoveryReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadDiscoveryReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadDiscoveryReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadDiscoveryReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadDiscoveryReportResponse")
	}
	return
}

// downloadDiscoveryReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadDiscoveryReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/actions/downloadReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadDiscoveryReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadMaskingLog Downloads the masking log generated by the last masking operation on a target database using the specified masking policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadMaskingLog.go.html to see an example of how to use DownloadMaskingLog API.
func (client DataSafeClient) DownloadMaskingLog(ctx context.Context, request DownloadMaskingLogRequest) (response DownloadMaskingLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadMaskingLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadMaskingLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadMaskingLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadMaskingLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadMaskingLogResponse")
	}
	return
}

// downloadMaskingLog implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadMaskingLog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/downloadLog", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadMaskingLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadMaskingPolicy Downloads an already-generated file corresponding to the specified masking policy.
// Note that the GenerateMaskingPolicyForDownload operation is a prerequisite for the
// DownloadMaskingPolicy operation. Use GenerateMaskingPolicyForDownload to generate
// a masking policy file and then use DownloadMaskingPolicy to download the generated file.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadMaskingPolicy.go.html to see an example of how to use DownloadMaskingPolicy API.
func (client DataSafeClient) DownloadMaskingPolicy(ctx context.Context, request DownloadMaskingPolicyRequest) (response DownloadMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadMaskingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadMaskingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadMaskingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadMaskingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadMaskingPolicyResponse")
	}
	return
}

// downloadMaskingPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadMaskingPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/download", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadMaskingPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadMaskingReport Downloads an already-generated masking report. Note that the GenerateMaskingReportForDownload
// operation is a prerequisite for the DownloadMaskingReport operation. Use GenerateMaskingReportForDownload
// to generate a masking report file and then use DownloadMaskingReport to download the generated file.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadMaskingReport.go.html to see an example of how to use DownloadMaskingReport API.
func (client DataSafeClient) DownloadMaskingReport(ctx context.Context, request DownloadMaskingReportRequest) (response DownloadMaskingReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadMaskingReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadMaskingReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadMaskingReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadMaskingReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadMaskingReportResponse")
	}
	return
}

// downloadMaskingReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadMaskingReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/downloadReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadMaskingReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadPrivilegeScript Downloads the privilege script to grant/revoke required roles from the Data Safe account on the target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadPrivilegeScript.go.html to see an example of how to use DownloadPrivilegeScript API.
func (client DataSafeClient) DownloadPrivilegeScript(ctx context.Context, request DownloadPrivilegeScriptRequest) (response DownloadPrivilegeScriptResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadPrivilegeScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadPrivilegeScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadPrivilegeScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadPrivilegeScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadPrivilegeScriptResponse")
	}
	return
}

// downloadPrivilegeScript implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadPrivilegeScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/downloadPrivilegeScript", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadPrivilegeScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadSecurityAssessmentReport Downloads the report of the specified security assessment. To download the security assessment report, it needs to be generated first.
// Please use GenerateSecurityAssessmentReport to generate a downloadable report in the preferred format (PDF, XLS).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadSecurityAssessmentReport.go.html to see an example of how to use DownloadSecurityAssessmentReport API.
func (client DataSafeClient) DownloadSecurityAssessmentReport(ctx context.Context, request DownloadSecurityAssessmentReportRequest) (response DownloadSecurityAssessmentReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadSecurityAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadSecurityAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadSecurityAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadSecurityAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadSecurityAssessmentReportResponse")
	}
	return
}

// downloadSecurityAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadSecurityAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/downloadReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadSecurityAssessmentReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadSensitiveDataModel Downloads an already-generated file corresponding to the specified sensitive data model. Note that the
// GenerateSensitiveDataModelForDownload operation is a prerequisite for the DownloadSensitiveDataModel operation.
// Use GenerateSensitiveDataModelForDownload to generate a data model file and then use DownloadSensitiveDataModel
// to download the generated file.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadSensitiveDataModel.go.html to see an example of how to use DownloadSensitiveDataModel API.
func (client DataSafeClient) DownloadSensitiveDataModel(ctx context.Context, request DownloadSensitiveDataModelRequest) (response DownloadSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadSensitiveDataModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadSensitiveDataModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadSensitiveDataModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadSensitiveDataModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadSensitiveDataModelResponse")
	}
	return
}

// downloadSensitiveDataModel implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadSensitiveDataModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/actions/download", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadSensitiveDataModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadUserAssessmentReport Downloads the report of the specified user assessment. To download the user assessment report, it needs to be generated first.
// Please use GenerateUserAssessmentReport to generate a downloadable report in the preferred format (PDF, XLS).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadUserAssessmentReport.go.html to see an example of how to use DownloadUserAssessmentReport API.
func (client DataSafeClient) DownloadUserAssessmentReport(ctx context.Context, request DownloadUserAssessmentReportRequest) (response DownloadUserAssessmentReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadUserAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadUserAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadUserAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadUserAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadUserAssessmentReportResponse")
	}
	return
}

// downloadUserAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) downloadUserAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/downloadReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadUserAssessmentReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDataSafeConfiguration Enables Data Safe in the tenancy and region.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/EnableDataSafeConfiguration.go.html to see an example of how to use EnableDataSafeConfiguration API.
func (client DataSafeClient) EnableDataSafeConfiguration(ctx context.Context, request EnableDataSafeConfigurationRequest) (response EnableDataSafeConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableDataSafeConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableDataSafeConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableDataSafeConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableDataSafeConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableDataSafeConfigurationResponse")
	}
	return
}

// enableDataSafeConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) enableDataSafeConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableDataSafeConfigurationResponse
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

// GenerateDiscoveryReportForDownload Generates a downloadable discovery report. It's a prerequisite for the DownloadDiscoveryReport operation. Use this
// endpoint to generate a discovery report file and then use DownloadDiscoveryReport to download the generated file.
// By default, it generates report for all the columns in a sensitive data model. Use the discoveryJobId attribute
// to generate report for a specific discovery job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateDiscoveryReportForDownload.go.html to see an example of how to use GenerateDiscoveryReportForDownload API.
func (client DataSafeClient) GenerateDiscoveryReportForDownload(ctx context.Context, request GenerateDiscoveryReportForDownloadRequest) (response GenerateDiscoveryReportForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateDiscoveryReportForDownload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateDiscoveryReportForDownloadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateDiscoveryReportForDownloadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateDiscoveryReportForDownloadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateDiscoveryReportForDownloadResponse")
	}
	return
}

// generateDiscoveryReportForDownload implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateDiscoveryReportForDownload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/actions/generateReportForDownload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateDiscoveryReportForDownloadResponse
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

// GenerateMaskingPolicyForDownload Generates a downloadable file corresponding to the specified masking policy. It's
// a prerequisite for the DownloadMaskingPolicy operation. Use this endpoint to generate
// a masking policy file and then use DownloadMaskingPolicy to download the generated file.
// Note that file generation and download are serial operations. The download operation
// can't be invoked while the generate operation is in progress.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateMaskingPolicyForDownload.go.html to see an example of how to use GenerateMaskingPolicyForDownload API.
func (client DataSafeClient) GenerateMaskingPolicyForDownload(ctx context.Context, request GenerateMaskingPolicyForDownloadRequest) (response GenerateMaskingPolicyForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateMaskingPolicyForDownload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateMaskingPolicyForDownloadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateMaskingPolicyForDownloadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateMaskingPolicyForDownloadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateMaskingPolicyForDownloadResponse")
	}
	return
}

// generateMaskingPolicyForDownload implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateMaskingPolicyForDownload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/generatePolicyForDownload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateMaskingPolicyForDownloadResponse
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

// GenerateMaskingReportForDownload Generates a downloadable masking report. It's a prerequisite for the
// DownloadMaskingReport operation. Use this endpoint to generate a
// masking report file and then use DownloadMaskingReport to download
// the generated file.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateMaskingReportForDownload.go.html to see an example of how to use GenerateMaskingReportForDownload API.
func (client DataSafeClient) GenerateMaskingReportForDownload(ctx context.Context, request GenerateMaskingReportForDownloadRequest) (response GenerateMaskingReportForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateMaskingReportForDownload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateMaskingReportForDownloadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateMaskingReportForDownloadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateMaskingReportForDownloadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateMaskingReportForDownloadResponse")
	}
	return
}

// generateMaskingReportForDownload implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateMaskingReportForDownload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/generateReportForDownload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateMaskingReportForDownloadResponse
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

// GenerateOnPremConnectorConfiguration Creates and downloads the configuration of the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateOnPremConnectorConfiguration.go.html to see an example of how to use GenerateOnPremConnectorConfiguration API.
func (client DataSafeClient) GenerateOnPremConnectorConfiguration(ctx context.Context, request GenerateOnPremConnectorConfigurationRequest) (response GenerateOnPremConnectorConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateOnPremConnectorConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateOnPremConnectorConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateOnPremConnectorConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateOnPremConnectorConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateOnPremConnectorConfigurationResponse")
	}
	return
}

// generateOnPremConnectorConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateOnPremConnectorConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/onPremConnectors/{onPremConnectorId}/actions/generateConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateOnPremConnectorConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateReport Generates a PDF or XLS report based on parameters and report definition.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateReport.go.html to see an example of how to use GenerateReport API.
func (client DataSafeClient) GenerateReport(ctx context.Context, request GenerateReportRequest) (response GenerateReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateReportResponse")
	}
	return
}

// generateReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reportDefinitions/{reportDefinitionId}/actions/generateReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateReportResponse
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

// GenerateSecurityAssessmentReport Generates the report of the specified security assessment. You can get the report in PDF or XLS format.
// After generating the report, use DownloadSecurityAssessmentReport to download it in the preferred format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateSecurityAssessmentReport.go.html to see an example of how to use GenerateSecurityAssessmentReport API.
func (client DataSafeClient) GenerateSecurityAssessmentReport(ctx context.Context, request GenerateSecurityAssessmentReportRequest) (response GenerateSecurityAssessmentReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateSecurityAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateSecurityAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateSecurityAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateSecurityAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateSecurityAssessmentReportResponse")
	}
	return
}

// generateSecurityAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateSecurityAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/generateReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateSecurityAssessmentReportResponse
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

// GenerateSensitiveDataModelForDownload Generates a downloadable file corresponding to the specified sensitive data model. It's a prerequisite for the
// DownloadSensitiveDataModel operation. Use this endpoint to generate a data model file and then use DownloadSensitiveDataModel
// to download the generated file. Note that file generation and download are serial operations. The download operation
// can't be invoked while the generate operation is in progress.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateSensitiveDataModelForDownload.go.html to see an example of how to use GenerateSensitiveDataModelForDownload API.
func (client DataSafeClient) GenerateSensitiveDataModelForDownload(ctx context.Context, request GenerateSensitiveDataModelForDownloadRequest) (response GenerateSensitiveDataModelForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateSensitiveDataModelForDownload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateSensitiveDataModelForDownloadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateSensitiveDataModelForDownloadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateSensitiveDataModelForDownloadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateSensitiveDataModelForDownloadResponse")
	}
	return
}

// generateSensitiveDataModelForDownload implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateSensitiveDataModelForDownload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/actions/generateDataModelForDownload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateSensitiveDataModelForDownloadResponse
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

// GenerateUserAssessmentReport Generates the report of the specified user assessment. The report is available in PDF or XLS format.
// After generating the report, use DownloadUserAssessmentReport to download it in the preferred format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateUserAssessmentReport.go.html to see an example of how to use GenerateUserAssessmentReport API.
func (client DataSafeClient) GenerateUserAssessmentReport(ctx context.Context, request GenerateUserAssessmentReportRequest) (response GenerateUserAssessmentReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateUserAssessmentReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateUserAssessmentReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateUserAssessmentReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateUserAssessmentReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateUserAssessmentReportResponse")
	}
	return
}

// generateUserAssessmentReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateUserAssessmentReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/generateReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateUserAssessmentReportResponse
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

// GetAlert Gets the details of alert by its ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAlert.go.html to see an example of how to use GetAlert API.
func (client DataSafeClient) GetAlert(ctx context.Context, request GetAlertRequest) (response GetAlertResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlert, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAlertResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAlertResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlertResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlertResponse")
	}
	return
}

// getAlert implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getAlert(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alerts/{alertId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAlertResponse
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

// GetAlertPolicy Gets the details of alert policy by its ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAlertPolicy.go.html to see an example of how to use GetAlertPolicy API.
func (client DataSafeClient) GetAlertPolicy(ctx context.Context, request GetAlertPolicyRequest) (response GetAlertPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlertPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAlertPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAlertPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlertPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlertPolicyResponse")
	}
	return
}

// getAlertPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getAlertPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alertPolicies/{alertPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAlertPolicyResponse
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

// GetAuditArchiveRetrieval Gets the details of the specified archive retreival.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditArchiveRetrieval.go.html to see an example of how to use GetAuditArchiveRetrieval API.
func (client DataSafeClient) GetAuditArchiveRetrieval(ctx context.Context, request GetAuditArchiveRetrievalRequest) (response GetAuditArchiveRetrievalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAuditArchiveRetrieval, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuditArchiveRetrievalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuditArchiveRetrievalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuditArchiveRetrievalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuditArchiveRetrievalResponse")
	}
	return
}

// getAuditArchiveRetrieval implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getAuditArchiveRetrieval(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditArchiveRetrievals/{auditArchiveRetrievalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuditArchiveRetrievalResponse
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

// GetAuditPolicy Gets a audit policy by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditPolicy.go.html to see an example of how to use GetAuditPolicy API.
func (client DataSafeClient) GetAuditPolicy(ctx context.Context, request GetAuditPolicyRequest) (response GetAuditPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAuditPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuditPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuditPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuditPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuditPolicyResponse")
	}
	return
}

// getAuditPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getAuditPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditPolicies/{auditPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuditPolicyResponse
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

// GetAuditProfile Gets the details of audit profile resource and associated audit trails of the audit profile.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditProfile.go.html to see an example of how to use GetAuditProfile API.
func (client DataSafeClient) GetAuditProfile(ctx context.Context, request GetAuditProfileRequest) (response GetAuditProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAuditProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuditProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuditProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuditProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuditProfileResponse")
	}
	return
}

// getAuditProfile implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getAuditProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditProfiles/{auditProfileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuditProfileResponse
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

// GetAuditTrail Gets the details of audit trail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditTrail.go.html to see an example of how to use GetAuditTrail API.
func (client DataSafeClient) GetAuditTrail(ctx context.Context, request GetAuditTrailRequest) (response GetAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAuditTrail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuditTrailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuditTrailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuditTrailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuditTrailResponse")
	}
	return
}

// getAuditTrail implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getAuditTrail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditTrails/{auditTrailId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuditTrailResponse
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

// GetCompatibleFormatsForDataTypes Gets a list of basic masking formats compatible with the supported data types.
// The data types are grouped into the following categories -
// Character - Includes CHAR, NCHAR, VARCHAR2, and NVARCHAR2
// Numeric - Includes NUMBER, FLOAT, RAW, BINARY_FLOAT, and BINARY_DOUBLE
// Date - Includes DATE and TIMESTAMP
// LOB - Includes BLOB, CLOB, and NCLOB
// All - Includes all the supported data types
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetCompatibleFormatsForDataTypes.go.html to see an example of how to use GetCompatibleFormatsForDataTypes API.
func (client DataSafeClient) GetCompatibleFormatsForDataTypes(ctx context.Context, request GetCompatibleFormatsForDataTypesRequest) (response GetCompatibleFormatsForDataTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCompatibleFormatsForDataTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCompatibleFormatsForDataTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCompatibleFormatsForDataTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCompatibleFormatsForDataTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCompatibleFormatsForDataTypesResponse")
	}
	return
}

// getCompatibleFormatsForDataTypes implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getCompatibleFormatsForDataTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compatibleFormatsForDataTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCompatibleFormatsForDataTypesResponse
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

// GetCompatibleFormatsForSensitiveTypes Gets a list of library masking formats compatible with the existing sensitive types.
// For each sensitive type, it returns the assigned default masking format as well as
// the other library masking formats that have the sensitiveTypeIds attribute containing
// the OCID of the sensitive type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetCompatibleFormatsForSensitiveTypes.go.html to see an example of how to use GetCompatibleFormatsForSensitiveTypes API.
func (client DataSafeClient) GetCompatibleFormatsForSensitiveTypes(ctx context.Context, request GetCompatibleFormatsForSensitiveTypesRequest) (response GetCompatibleFormatsForSensitiveTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCompatibleFormatsForSensitiveTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCompatibleFormatsForSensitiveTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCompatibleFormatsForSensitiveTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCompatibleFormatsForSensitiveTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCompatibleFormatsForSensitiveTypesResponse")
	}
	return
}

// getCompatibleFormatsForSensitiveTypes implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getCompatibleFormatsForSensitiveTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compatibleFormatsForSensitiveTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCompatibleFormatsForSensitiveTypesResponse
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

// GetDataSafeConfiguration Gets the details of the Data Safe configuration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDataSafeConfiguration.go.html to see an example of how to use GetDataSafeConfiguration API.
func (client DataSafeClient) GetDataSafeConfiguration(ctx context.Context, request GetDataSafeConfigurationRequest) (response GetDataSafeConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataSafeConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataSafeConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataSafeConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataSafeConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataSafeConfigurationResponse")
	}
	return
}

// getDataSafeConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDataSafeConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataSafeConfigurationResponse
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

// GetDataSafePrivateEndpoint Gets the details of the specified Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDataSafePrivateEndpoint.go.html to see an example of how to use GetDataSafePrivateEndpoint API.
func (client DataSafeClient) GetDataSafePrivateEndpoint(ctx context.Context, request GetDataSafePrivateEndpointRequest) (response GetDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataSafePrivateEndpointResponse")
	}
	return
}

// getDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataSafePrivateEndpointResponse
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

// GetDiscoveryJob Gets the details of the specified discovery job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDiscoveryJob.go.html to see an example of how to use GetDiscoveryJob API.
func (client DataSafeClient) GetDiscoveryJob(ctx context.Context, request GetDiscoveryJobRequest) (response GetDiscoveryJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDiscoveryJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDiscoveryJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDiscoveryJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDiscoveryJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDiscoveryJobResponse")
	}
	return
}

// getDiscoveryJob implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDiscoveryJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs/{discoveryJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDiscoveryJobResponse
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

// GetDiscoveryJobResult Gets the details of the specified discovery result.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDiscoveryJobResult.go.html to see an example of how to use GetDiscoveryJobResult API.
func (client DataSafeClient) GetDiscoveryJobResult(ctx context.Context, request GetDiscoveryJobResultRequest) (response GetDiscoveryJobResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDiscoveryJobResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDiscoveryJobResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDiscoveryJobResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDiscoveryJobResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDiscoveryJobResultResponse")
	}
	return
}

// getDiscoveryJobResult implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDiscoveryJobResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs/{discoveryJobId}/results/{resultKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDiscoveryJobResultResponse
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

// GetLibraryMaskingFormat Gets the details of the specified library masking format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetLibraryMaskingFormat.go.html to see an example of how to use GetLibraryMaskingFormat API.
func (client DataSafeClient) GetLibraryMaskingFormat(ctx context.Context, request GetLibraryMaskingFormatRequest) (response GetLibraryMaskingFormatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLibraryMaskingFormat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLibraryMaskingFormatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLibraryMaskingFormatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLibraryMaskingFormatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLibraryMaskingFormatResponse")
	}
	return
}

// getLibraryMaskingFormat implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getLibraryMaskingFormat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/libraryMaskingFormats/{libraryMaskingFormatId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLibraryMaskingFormatResponse
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

// GetMaskingColumn Gets the details of the specified masking column.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingColumn.go.html to see an example of how to use GetMaskingColumn API.
func (client DataSafeClient) GetMaskingColumn(ctx context.Context, request GetMaskingColumnRequest) (response GetMaskingColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaskingColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaskingColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaskingColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaskingColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaskingColumnResponse")
	}
	return
}

// getMaskingColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getMaskingColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicies/{maskingPolicyId}/maskingColumns/{maskingColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaskingColumnResponse
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

// GetMaskingPolicy Gets the details of the specified masking policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingPolicy.go.html to see an example of how to use GetMaskingPolicy API.
func (client DataSafeClient) GetMaskingPolicy(ctx context.Context, request GetMaskingPolicyRequest) (response GetMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaskingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaskingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaskingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaskingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaskingPolicyResponse")
	}
	return
}

// getMaskingPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getMaskingPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicies/{maskingPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaskingPolicyResponse
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

// GetMaskingReport Gets the details of the specified masking report.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingReport.go.html to see an example of how to use GetMaskingReport API.
func (client DataSafeClient) GetMaskingReport(ctx context.Context, request GetMaskingReportRequest) (response GetMaskingReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaskingReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaskingReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaskingReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaskingReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaskingReportResponse")
	}
	return
}

// getMaskingReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getMaskingReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingReports/{maskingReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaskingReportResponse
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

// GetOnPremConnector Gets the details of the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetOnPremConnector.go.html to see an example of how to use GetOnPremConnector API.
func (client DataSafeClient) GetOnPremConnector(ctx context.Context, request GetOnPremConnectorRequest) (response GetOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOnPremConnectorResponse")
	}
	return
}

// getOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/onPremConnectors/{onPremConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOnPremConnectorResponse
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

// GetReport Gets a report by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetReport.go.html to see an example of how to use GetReport API.
func (client DataSafeClient) GetReport(ctx context.Context, request GetReportRequest) (response GetReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportResponse")
	}
	return
}

// getReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportResponse
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

// GetReportContent Downloads the specified report in the form of PDF or XLXS.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetReportContent.go.html to see an example of how to use GetReportContent API.
func (client DataSafeClient) GetReportContent(ctx context.Context, request GetReportContentRequest) (response GetReportContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReportContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportContentResponse")
	}
	return
}

// getReportContent implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getReportContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReportDefinition Gets the details of report definition specified by the identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetReportDefinition.go.html to see an example of how to use GetReportDefinition API.
func (client DataSafeClient) GetReportDefinition(ctx context.Context, request GetReportDefinitionRequest) (response GetReportDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReportDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportDefinitionResponse")
	}
	return
}

// getReportDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getReportDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reportDefinitions/{reportDefinitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportDefinitionResponse
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

// GetSecurityAssessment Gets the details of the specified security assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessment.go.html to see an example of how to use GetSecurityAssessment API.
func (client DataSafeClient) GetSecurityAssessment(ctx context.Context, request GetSecurityAssessmentRequest) (response GetSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAssessmentResponse")
	}
	return
}

// getSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAssessmentResponse
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

// GetSecurityAssessmentComparison Gets the details of the comparison report on the security assessments submitted for comparison.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessmentComparison.go.html to see an example of how to use GetSecurityAssessmentComparison API.
func (client DataSafeClient) GetSecurityAssessmentComparison(ctx context.Context, request GetSecurityAssessmentComparisonRequest) (response GetSecurityAssessmentComparisonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAssessmentComparison, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAssessmentComparisonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAssessmentComparisonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAssessmentComparisonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAssessmentComparisonResponse")
	}
	return
}

// getSecurityAssessmentComparison implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityAssessmentComparison(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}/comparison/{comparisonSecurityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAssessmentComparisonResponse
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

// GetSensitiveColumn Gets the details of the specified sensitive column.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSensitiveColumn.go.html to see an example of how to use GetSensitiveColumn API.
func (client DataSafeClient) GetSensitiveColumn(ctx context.Context, request GetSensitiveColumnRequest) (response GetSensitiveColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSensitiveColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSensitiveColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSensitiveColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSensitiveColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSensitiveColumnResponse")
	}
	return
}

// getSensitiveColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSensitiveColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns/{sensitiveColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSensitiveColumnResponse
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

// GetSensitiveDataModel Gets the details of the specified sensitive data model.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSensitiveDataModel.go.html to see an example of how to use GetSensitiveDataModel API.
func (client DataSafeClient) GetSensitiveDataModel(ctx context.Context, request GetSensitiveDataModelRequest) (response GetSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSensitiveDataModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSensitiveDataModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSensitiveDataModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSensitiveDataModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSensitiveDataModelResponse")
	}
	return
}

// getSensitiveDataModel implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSensitiveDataModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels/{sensitiveDataModelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSensitiveDataModelResponse
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

// GetSensitiveType Gets the details of the specified sensitive type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSensitiveType.go.html to see an example of how to use GetSensitiveType API.
func (client DataSafeClient) GetSensitiveType(ctx context.Context, request GetSensitiveTypeRequest) (response GetSensitiveTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSensitiveType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSensitiveTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSensitiveTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSensitiveTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSensitiveTypeResponse")
	}
	return
}

// getSensitiveType implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSensitiveType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveTypes/{sensitiveTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSensitiveTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &sensitivetype{})
	return response, err
}

// GetTargetAlertPolicyAssociation Gets the details of target-alert policy association by its ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetTargetAlertPolicyAssociation.go.html to see an example of how to use GetTargetAlertPolicyAssociation API.
func (client DataSafeClient) GetTargetAlertPolicyAssociation(ctx context.Context, request GetTargetAlertPolicyAssociationRequest) (response GetTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetAlertPolicyAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetAlertPolicyAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetAlertPolicyAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetAlertPolicyAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetAlertPolicyAssociationResponse")
	}
	return
}

// getTargetAlertPolicyAssociation implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getTargetAlertPolicyAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetAlertPolicyAssociations/{targetAlertPolicyAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetAlertPolicyAssociationResponse
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

// GetTargetDatabase Returns the details of the specified Data Safe target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetTargetDatabase.go.html to see an example of how to use GetTargetDatabase API.
func (client DataSafeClient) GetTargetDatabase(ctx context.Context, request GetTargetDatabaseRequest) (response GetTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetDatabaseResponse")
	}
	return
}

// getTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetDatabaseResponse
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

// GetUserAssessment Gets a user assessment by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetUserAssessment.go.html to see an example of how to use GetUserAssessment API.
func (client DataSafeClient) GetUserAssessment(ctx context.Context, request GetUserAssessmentRequest) (response GetUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserAssessmentResponse")
	}
	return
}

// getUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserAssessmentResponse
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

// GetUserAssessmentComparison Gets the details of the comparison report for the user assessments provided.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetUserAssessmentComparison.go.html to see an example of how to use GetUserAssessmentComparison API.
func (client DataSafeClient) GetUserAssessmentComparison(ctx context.Context, request GetUserAssessmentComparisonRequest) (response GetUserAssessmentComparisonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserAssessmentComparison, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserAssessmentComparisonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserAssessmentComparisonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserAssessmentComparisonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserAssessmentComparisonResponse")
	}
	return
}

// getUserAssessmentComparison implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getUserAssessmentComparison(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/comparison/{comparisonUserAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserAssessmentComparisonResponse
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

// GetWorkRequest Gets the details of the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client DataSafeClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DataSafeClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlertAnalytics Returns aggregation details of alerts.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertAnalytics.go.html to see an example of how to use ListAlertAnalytics API.
func (client DataSafeClient) ListAlertAnalytics(ctx context.Context, request ListAlertAnalyticsRequest) (response ListAlertAnalyticsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAlertAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlertAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlertAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlertAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlertAnalyticsResponse")
	}
	return
}

// listAlertAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAlertAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alertAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlertAnalyticsResponse
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

// ListAlertPolicies Gets a list of all alert policies.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertPolicies.go.html to see an example of how to use ListAlertPolicies API.
func (client DataSafeClient) ListAlertPolicies(ctx context.Context, request ListAlertPoliciesRequest) (response ListAlertPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlertPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlertPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlertPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlertPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlertPoliciesResponse")
	}
	return
}

// listAlertPolicies implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAlertPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alertPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlertPoliciesResponse
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

// ListAlertPolicyRules Lists the rules of the specified alert policy. The alert policy is said to be satisfied when all rules in the policy evaulate to true.
// If there are three rules: rule1,rule2 and rule3, the policy is satisfied if rule1 AND rule2 AND rule3 is True.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertPolicyRules.go.html to see an example of how to use ListAlertPolicyRules API.
func (client DataSafeClient) ListAlertPolicyRules(ctx context.Context, request ListAlertPolicyRulesRequest) (response ListAlertPolicyRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlertPolicyRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlertPolicyRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlertPolicyRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlertPolicyRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlertPolicyRulesResponse")
	}
	return
}

// listAlertPolicyRules implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAlertPolicyRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alertPolicies/{alertPolicyId}/rules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlertPolicyRulesResponse
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

// ListAlerts Gets a list of all alerts.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlerts.go.html to see an example of how to use ListAlerts API.
func (client DataSafeClient) ListAlerts(ctx context.Context, request ListAlertsRequest) (response ListAlertsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlertsResponse")
	}
	return
}

// listAlerts implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAlerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlertsResponse
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

// ListAuditArchiveRetrievals Returns the list of audit archive retrieval.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditArchiveRetrievals.go.html to see an example of how to use ListAuditArchiveRetrievals API.
func (client DataSafeClient) ListAuditArchiveRetrievals(ctx context.Context, request ListAuditArchiveRetrievalsRequest) (response ListAuditArchiveRetrievalsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditArchiveRetrievals, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditArchiveRetrievalsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditArchiveRetrievalsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditArchiveRetrievalsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditArchiveRetrievalsResponse")
	}
	return
}

// listAuditArchiveRetrievals implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditArchiveRetrievals(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditArchiveRetrievals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditArchiveRetrievalsResponse
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

// ListAuditEventAnalytics By default ListAuditEventAnalytics operation will return all of the summary columns. To filter desired summary columns, specify
// it in the `summaryOf` query parameter.
// **Example:** /ListAuditEventAnalytics?summaryField=targetName&summaryField=userName&summaryField=clientHostName&summaryField
//              &summaryField=dmls&summaryField=privilege_changes&summaryField=ddls&summaryField=login_failure&summaryField=login_success
//              &summaryField=eventcount&q=(operationTime ge '2021-06-13T23:49:14')&groupBy=targetName
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditEventAnalytics.go.html to see an example of how to use ListAuditEventAnalytics API.
func (client DataSafeClient) ListAuditEventAnalytics(ctx context.Context, request ListAuditEventAnalyticsRequest) (response ListAuditEventAnalyticsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listAuditEventAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditEventAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditEventAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditEventAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditEventAnalyticsResponse")
	}
	return
}

// listAuditEventAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditEventAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditEventAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditEventAnalyticsResponse
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

// ListAuditEvents The ListAuditEvents operation returns specified `compartmentId` audit Events only.
// The list does not include any audit Events associated with the `subcompartments` of the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListAuditEvents on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditEvents.go.html to see an example of how to use ListAuditEvents API.
func (client DataSafeClient) ListAuditEvents(ctx context.Context, request ListAuditEventsRequest) (response ListAuditEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditEventsResponse")
	}
	return
}

// listAuditEvents implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditEventsResponse
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

// ListAuditPolicies Retrieves a list of all audited targets with their corresponding provisioned audit policies, and their provisioning conditions.
// The ListAuditPolicies operation returns only the audit policies in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListAuditPolicies on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditPolicies.go.html to see an example of how to use ListAuditPolicies API.
func (client DataSafeClient) ListAuditPolicies(ctx context.Context, request ListAuditPoliciesRequest) (response ListAuditPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditPoliciesResponse")
	}
	return
}

// listAuditPolicies implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditPoliciesResponse
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

// ListAuditProfileAnalytics Gets a list of audit profile aggregated details . A audit profile  aggregation helps understand the overall  state of audit profile profiles.
// As an example, it helps understand how many audit profiles have paid usage. It is especially useful to create dashboards or to support analytics.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform AuditProfileAnalytics on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditProfileAnalytics.go.html to see an example of how to use ListAuditProfileAnalytics API.
func (client DataSafeClient) ListAuditProfileAnalytics(ctx context.Context, request ListAuditProfileAnalyticsRequest) (response ListAuditProfileAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditProfileAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditProfileAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditProfileAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditProfileAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditProfileAnalyticsResponse")
	}
	return
}

// listAuditProfileAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditProfileAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditProfileAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditProfileAnalyticsResponse
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

// ListAuditProfiles Gets a list of all audit profiles.
// The ListAuditProfiles operation returns only the audit profiles in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListAuditProfiles on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditProfiles.go.html to see an example of how to use ListAuditProfiles API.
func (client DataSafeClient) ListAuditProfiles(ctx context.Context, request ListAuditProfilesRequest) (response ListAuditProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditProfilesResponse")
	}
	return
}

// listAuditProfiles implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditProfilesResponse
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

// ListAuditTrailAnalytics Gets a list of audit trail aggregated details . A audit trail aggregation helps understand the overall  state of trails.
// As an example, it helps understand how many trails are running or stopped. It is especially useful to create dashboards or to support analytics.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform AuditTrailAnalytics on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditTrailAnalytics.go.html to see an example of how to use ListAuditTrailAnalytics API.
func (client DataSafeClient) ListAuditTrailAnalytics(ctx context.Context, request ListAuditTrailAnalyticsRequest) (response ListAuditTrailAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditTrailAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditTrailAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditTrailAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditTrailAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditTrailAnalyticsResponse")
	}
	return
}

// listAuditTrailAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditTrailAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditTrailAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditTrailAnalyticsResponse
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

// ListAuditTrails Gets a list of all audit trails.
// The ListAuditTrails operation returns only the audit trails in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListAuditTrails on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditTrails.go.html to see an example of how to use ListAuditTrails API.
func (client DataSafeClient) ListAuditTrails(ctx context.Context, request ListAuditTrailsRequest) (response ListAuditTrailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditTrails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditTrailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditTrailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditTrailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditTrailsResponse")
	}
	return
}

// listAuditTrails implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditTrails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditTrails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditTrailsResponse
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

// ListAvailableAuditVolumes Retrieves a list of audit trails, and associated audit event volume for each trail up to defined start date.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAvailableAuditVolumes.go.html to see an example of how to use ListAvailableAuditVolumes API.
func (client DataSafeClient) ListAvailableAuditVolumes(ctx context.Context, request ListAvailableAuditVolumesRequest) (response ListAvailableAuditVolumesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailableAuditVolumes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailableAuditVolumesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailableAuditVolumesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailableAuditVolumesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailableAuditVolumesResponse")
	}
	return
}

// listAvailableAuditVolumes implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAvailableAuditVolumes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditProfiles/{auditProfileId}/availableAuditVolumes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailableAuditVolumesResponse
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

// ListCollectedAuditVolumes Gets a list of all collected audit volume data points.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCollectedAuditVolumes.go.html to see an example of how to use ListCollectedAuditVolumes API.
func (client DataSafeClient) ListCollectedAuditVolumes(ctx context.Context, request ListCollectedAuditVolumesRequest) (response ListCollectedAuditVolumesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCollectedAuditVolumes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCollectedAuditVolumesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCollectedAuditVolumesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCollectedAuditVolumesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCollectedAuditVolumesResponse")
	}
	return
}

// listCollectedAuditVolumes implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listCollectedAuditVolumes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditProfiles/{auditProfileId}/collectedAuditVolumes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCollectedAuditVolumesResponse
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

// ListColumns Returns a list of column metadata objects.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListColumns.go.html to see an example of how to use ListColumns API.
func (client DataSafeClient) ListColumns(ctx context.Context, request ListColumnsRequest) (response ListColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListColumnsResponse")
	}
	return
}

// listColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}/columns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListColumnsResponse
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

// ListDataSafePrivateEndpoints Gets a list of Data Safe private endpoints.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDataSafePrivateEndpoints.go.html to see an example of how to use ListDataSafePrivateEndpoints API.
func (client DataSafeClient) ListDataSafePrivateEndpoints(ctx context.Context, request ListDataSafePrivateEndpointsRequest) (response ListDataSafePrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataSafePrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataSafePrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataSafePrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataSafePrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataSafePrivateEndpointsResponse")
	}
	return
}

// listDataSafePrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDataSafePrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSafePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataSafePrivateEndpointsResponse
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

// ListDiscoveryAnalytics Gets consolidated discovery analytics data based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryAnalytics.go.html to see an example of how to use ListDiscoveryAnalytics API.
func (client DataSafeClient) ListDiscoveryAnalytics(ctx context.Context, request ListDiscoveryAnalyticsRequest) (response ListDiscoveryAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDiscoveryAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDiscoveryAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDiscoveryAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDiscoveryAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDiscoveryAnalyticsResponse")
	}
	return
}

// listDiscoveryAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDiscoveryAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDiscoveryAnalyticsResponse
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

// ListDiscoveryJobResults Gets a list of discovery results based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryJobResults.go.html to see an example of how to use ListDiscoveryJobResults API.
func (client DataSafeClient) ListDiscoveryJobResults(ctx context.Context, request ListDiscoveryJobResultsRequest) (response ListDiscoveryJobResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDiscoveryJobResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDiscoveryJobResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDiscoveryJobResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDiscoveryJobResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDiscoveryJobResultsResponse")
	}
	return
}

// listDiscoveryJobResults implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDiscoveryJobResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs/{discoveryJobId}/results", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDiscoveryJobResultsResponse
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

// ListDiscoveryJobs Gets a list of incremental discovery jobs based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryJobs.go.html to see an example of how to use ListDiscoveryJobs API.
func (client DataSafeClient) ListDiscoveryJobs(ctx context.Context, request ListDiscoveryJobsRequest) (response ListDiscoveryJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDiscoveryJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDiscoveryJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDiscoveryJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDiscoveryJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDiscoveryJobsResponse")
	}
	return
}

// listDiscoveryJobs implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDiscoveryJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDiscoveryJobsResponse
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

// ListFindings List all the findings from all the targets in the specified assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindings.go.html to see an example of how to use ListFindings API.
func (client DataSafeClient) ListFindings(ctx context.Context, request ListFindingsRequest) (response ListFindingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFindings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFindingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFindingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFindingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFindingsResponse")
	}
	return
}

// listFindings implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listFindings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}/findings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFindingsResponse
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

// ListGrants Gets a list of grants for a particular user in the specified user assessment. A user grant contains details such as the
// privilege name, type, category, and depth level. The depth level indicates how deep in the hierarchy of roles granted to
// roles a privilege grant is. The userKey in this operation is a system-generated identifier. Perform the operation ListUsers
// to get the userKey for a particular user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListGrants.go.html to see an example of how to use ListGrants API.
func (client DataSafeClient) ListGrants(ctx context.Context, request ListGrantsRequest) (response ListGrantsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGrants, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGrantsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGrantsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGrantsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGrantsResponse")
	}
	return
}

// listGrants implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listGrants(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/users/{userKey}/grants", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGrantsResponse
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

// ListLibraryMaskingFormats Gets a list of library masking formats based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListLibraryMaskingFormats.go.html to see an example of how to use ListLibraryMaskingFormats API.
func (client DataSafeClient) ListLibraryMaskingFormats(ctx context.Context, request ListLibraryMaskingFormatsRequest) (response ListLibraryMaskingFormatsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLibraryMaskingFormats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLibraryMaskingFormatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLibraryMaskingFormatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLibraryMaskingFormatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLibraryMaskingFormatsResponse")
	}
	return
}

// listLibraryMaskingFormats implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listLibraryMaskingFormats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/libraryMaskingFormats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLibraryMaskingFormatsResponse
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

// ListMaskedColumns Gets a list of masked columns present in the specified masking report and based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskedColumns.go.html to see an example of how to use ListMaskedColumns API.
func (client DataSafeClient) ListMaskedColumns(ctx context.Context, request ListMaskedColumnsRequest) (response ListMaskedColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskedColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskedColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskedColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskedColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskedColumnsResponse")
	}
	return
}

// listMaskedColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskedColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingReports/{maskingReportId}/maskedColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskedColumnsResponse
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

// ListMaskingAnalytics Gets consolidated masking analytics data based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingAnalytics.go.html to see an example of how to use ListMaskingAnalytics API.
func (client DataSafeClient) ListMaskingAnalytics(ctx context.Context, request ListMaskingAnalyticsRequest) (response ListMaskingAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingAnalyticsResponse")
	}
	return
}

// listMaskingAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingAnalyticsResponse
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

// ListMaskingColumns Gets a list of masking columns present in the specified masking policy and based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingColumns.go.html to see an example of how to use ListMaskingColumns API.
func (client DataSafeClient) ListMaskingColumns(ctx context.Context, request ListMaskingColumnsRequest) (response ListMaskingColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingColumnsResponse")
	}
	return
}

// listMaskingColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicies/{maskingPolicyId}/maskingColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingColumnsResponse
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

// ListMaskingPolicies Gets a list of masking policies based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicies.go.html to see an example of how to use ListMaskingPolicies API.
func (client DataSafeClient) ListMaskingPolicies(ctx context.Context, request ListMaskingPoliciesRequest) (response ListMaskingPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingPoliciesResponse")
	}
	return
}

// listMaskingPolicies implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingPoliciesResponse
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

// ListMaskingReports Gets a list of masking reports based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingReports.go.html to see an example of how to use ListMaskingReports API.
func (client DataSafeClient) ListMaskingReports(ctx context.Context, request ListMaskingReportsRequest) (response ListMaskingReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingReportsResponse")
	}
	return
}

// listMaskingReports implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingReportsResponse
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

// ListOnPremConnectors Gets a list of on-premises connectors.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListOnPremConnectors.go.html to see an example of how to use ListOnPremConnectors API.
func (client DataSafeClient) ListOnPremConnectors(ctx context.Context, request ListOnPremConnectorsRequest) (response ListOnPremConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOnPremConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOnPremConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOnPremConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOnPremConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOnPremConnectorsResponse")
	}
	return
}

// listOnPremConnectors implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listOnPremConnectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/onPremConnectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOnPremConnectorsResponse
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

// ListReportDefinitions Gets a list of report definitions.
// The ListReportDefinitions operation returns only the report definitions in the specified `compartmentId`.
// It also returns the seeded report definitions which are available to all the compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReportDefinitions.go.html to see an example of how to use ListReportDefinitions API.
func (client DataSafeClient) ListReportDefinitions(ctx context.Context, request ListReportDefinitionsRequest) (response ListReportDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReportDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReportDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReportDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReportDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReportDefinitionsResponse")
	}
	return
}

// listReportDefinitions implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listReportDefinitions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reportDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReportDefinitionsResponse
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

// ListReports Gets a list of all the reports in the compartment. It contains information such as report generation time.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReports.go.html to see an example of how to use ListReports API.
func (client DataSafeClient) ListReports(ctx context.Context, request ListReportsRequest) (response ListReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReportsResponse")
	}
	return
}

// listReports implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReportsResponse
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

// ListRoles Returns a list of role metadata objects.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListRoles.go.html to see an example of how to use ListRoles API.
func (client DataSafeClient) ListRoles(ctx context.Context, request ListRolesRequest) (response ListRolesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRolesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRolesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRolesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRolesResponse")
	}
	return
}

// listRoles implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listRoles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}/roles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRolesResponse
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

// ListSchemas Returns list of schema.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSchemas.go.html to see an example of how to use ListSchemas API.
func (client DataSafeClient) ListSchemas(ctx context.Context, request ListSchemasRequest) (response ListSchemasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSchemas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSchemasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSchemasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSchemasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSchemasResponse")
	}
	return
}

// listSchemas implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSchemas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}/schemas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchemasResponse
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

// ListSecurityAssessments Gets a list of security assessments.
// The ListSecurityAssessments operation returns only the assessments in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSecurityAssessments on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityAssessments.go.html to see an example of how to use ListSecurityAssessments API.
func (client DataSafeClient) ListSecurityAssessments(ctx context.Context, request ListSecurityAssessmentsRequest) (response ListSecurityAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAssessments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAssessmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAssessmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAssessmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAssessmentsResponse")
	}
	return
}

// listSecurityAssessments implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityAssessments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAssessmentsResponse
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

// ListSensitiveColumns Gets a list of sensitive columns present in the specified sensitive data model based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveColumns.go.html to see an example of how to use ListSensitiveColumns API.
func (client DataSafeClient) ListSensitiveColumns(ctx context.Context, request ListSensitiveColumnsRequest) (response ListSensitiveColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSensitiveColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSensitiveColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSensitiveColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSensitiveColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSensitiveColumnsResponse")
	}
	return
}

// listSensitiveColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSensitiveColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSensitiveColumnsResponse
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

// ListSensitiveDataModels Gets a list of sensitive data models based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveDataModels.go.html to see an example of how to use ListSensitiveDataModels API.
func (client DataSafeClient) ListSensitiveDataModels(ctx context.Context, request ListSensitiveDataModelsRequest) (response ListSensitiveDataModelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSensitiveDataModels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSensitiveDataModelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSensitiveDataModelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSensitiveDataModelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSensitiveDataModelsResponse")
	}
	return
}

// listSensitiveDataModels implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSensitiveDataModels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSensitiveDataModelsResponse
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

// ListSensitiveTypes Gets a list of sensitive types based on the specified query parameters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveTypes.go.html to see an example of how to use ListSensitiveTypes API.
func (client DataSafeClient) ListSensitiveTypes(ctx context.Context, request ListSensitiveTypesRequest) (response ListSensitiveTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSensitiveTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSensitiveTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSensitiveTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSensitiveTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSensitiveTypesResponse")
	}
	return
}

// listSensitiveTypes implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSensitiveTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSensitiveTypesResponse
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

// ListTables Returns a list of table metadata objects.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTables.go.html to see an example of how to use ListTables API.
func (client DataSafeClient) ListTables(ctx context.Context, request ListTablesRequest) (response ListTablesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTables, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTablesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTablesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTablesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTablesResponse")
	}
	return
}

// listTables implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listTables(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}/tables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTablesResponse
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

// ListTargetAlertPolicyAssociations Gets a list of all target-alert policy associations.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetAlertPolicyAssociations.go.html to see an example of how to use ListTargetAlertPolicyAssociations API.
func (client DataSafeClient) ListTargetAlertPolicyAssociations(ctx context.Context, request ListTargetAlertPolicyAssociationsRequest) (response ListTargetAlertPolicyAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetAlertPolicyAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetAlertPolicyAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetAlertPolicyAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetAlertPolicyAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetAlertPolicyAssociationsResponse")
	}
	return
}

// listTargetAlertPolicyAssociations implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listTargetAlertPolicyAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetAlertPolicyAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetAlertPolicyAssociationsResponse
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

// ListTargetDatabases Returns the list of registered target databases in Data Safe.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetDatabases.go.html to see an example of how to use ListTargetDatabases API.
func (client DataSafeClient) ListTargetDatabases(ctx context.Context, request ListTargetDatabasesRequest) (response ListTargetDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetDatabasesResponse")
	}
	return
}

// listTargetDatabases implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listTargetDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetDatabasesResponse
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

// ListUserAnalytics Gets a list of aggregated user details from the specified user assessment. This provides information about the overall state
// of database user security.  For example, the user details include how many users have the DBA role and how many users are in
// the critical category. This data is especially useful content for dashboards or to support analytics.
// When you perform the ListUserAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
// parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
// permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
// root compartment. If the requestor does not have access to at least one subcompartment of the compartment specified by
// compartmentId, then "Not Authorized" is returned.
// The parameter compartmentIdInSubtree applies when you perform ListUserAnalytics on the compartmentId passed and when it is
// set to true, the entire hierarchy of compartments can be returned.
// To use ListUserAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
// set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAnalytics.go.html to see an example of how to use ListUserAnalytics API.
func (client DataSafeClient) ListUserAnalytics(ctx context.Context, request ListUserAnalyticsRequest) (response ListUserAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserAnalyticsResponse")
	}
	return
}

// listUserAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listUserAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/userAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserAnalyticsResponse
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

// ListUserAssessments Gets a list of user assessments.
// The ListUserAssessments operation returns only the assessments in the specified `compartmentId`.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListUserAssessments on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAssessments.go.html to see an example of how to use ListUserAssessments API.
func (client DataSafeClient) ListUserAssessments(ctx context.Context, request ListUserAssessmentsRequest) (response ListUserAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserAssessments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserAssessmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserAssessmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserAssessmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserAssessmentsResponse")
	}
	return
}

// listUserAssessments implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listUserAssessments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserAssessmentsResponse
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

// ListUsers Gets a list of users of the specified user assessment. The result contains the database user details for each user, such
// as user type, account status, last login time, user creation time, authentication type, user profile, and the date and time
// of the latest password change. It also contains the user category derived from these user details as well as privileges
// granted to each user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUsers.go.html to see an example of how to use ListUsers API.
func (client DataSafeClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsersResponse")
	}
	return
}

// listUsers implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsersResponse
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

// ListWorkRequestErrors Gets a list of errors for the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client DataSafeClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DataSafeClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Gets a list of log entries for the specified work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client DataSafeClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DataSafeClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Gets a list of work requests.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client DataSafeClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DataSafeClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MaskData Masks data using the specified masking policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/MaskData.go.html to see an example of how to use MaskData API.
func (client DataSafeClient) MaskData(ctx context.Context, request MaskDataRequest) (response MaskDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.maskData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MaskDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MaskDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MaskDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MaskDataResponse")
	}
	return
}

// maskData implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) maskData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/mask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MaskDataResponse
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

// ModifyGlobalSettings Modifies Global Settings in Data Safe in the tenancy and region.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ModifyGlobalSettings.go.html to see an example of how to use ModifyGlobalSettings API.
func (client DataSafeClient) ModifyGlobalSettings(ctx context.Context, request ModifyGlobalSettingsRequest) (response ModifyGlobalSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.modifyGlobalSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ModifyGlobalSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ModifyGlobalSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ModifyGlobalSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ModifyGlobalSettingsResponse")
	}
	return
}

// modifyGlobalSettings implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) modifyGlobalSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configuration/{compartmentId}/actions/modifyGlobalSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ModifyGlobalSettingsResponse
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

// PatchAlerts Patch alerts. Updates one or more alerts by specifying alert Ids.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchAlerts.go.html to see an example of how to use PatchAlerts API.
func (client DataSafeClient) PatchAlerts(ctx context.Context, request PatchAlertsRequest) (response PatchAlertsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchAlerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchAlertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchAlertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchAlertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchAlertsResponse")
	}
	return
}

// patchAlerts implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) patchAlerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/alerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchAlertsResponse
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

// PatchDiscoveryJobResults Patches one or more discovery results. You can use this operation to set the plannedAction attribute before using
// ApplyDiscoveryJobResults to process the results based on this attribute.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchDiscoveryJobResults.go.html to see an example of how to use PatchDiscoveryJobResults API.
func (client DataSafeClient) PatchDiscoveryJobResults(ctx context.Context, request PatchDiscoveryJobResultsRequest) (response PatchDiscoveryJobResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchDiscoveryJobResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchDiscoveryJobResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchDiscoveryJobResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchDiscoveryJobResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchDiscoveryJobResultsResponse")
	}
	return
}

// patchDiscoveryJobResults implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) patchDiscoveryJobResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/discoveryJobs/{discoveryJobId}/results", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchDiscoveryJobResultsResponse
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

// PatchMaskingColumns Patches one or more columns in the specified masking policy. Use it to create, or update
// masking columns. To create masking columns, use CreateMaskingColumnDetails as the patch
// value. And to update masking columns, use UpdateMaskingColumnDetails as the patch value.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchMaskingColumns.go.html to see an example of how to use PatchMaskingColumns API.
func (client DataSafeClient) PatchMaskingColumns(ctx context.Context, request PatchMaskingColumnsRequest) (response PatchMaskingColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchMaskingColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchMaskingColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchMaskingColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchMaskingColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchMaskingColumnsResponse")
	}
	return
}

// patchMaskingColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) patchMaskingColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/maskingPolicies/{maskingPolicyId}/maskingColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchMaskingColumnsResponse
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

// PatchSensitiveColumns Patches one or more columns in the specified sensitive data model. Use it to create, update, or delete sensitive columns.
// To create sensitive columns, use CreateSensitiveColumnDetails as the patch value. And to update sensitive columns,
// use UpdateSensitiveColumnDetails as the patch value.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchSensitiveColumns.go.html to see an example of how to use PatchSensitiveColumns API.
func (client DataSafeClient) PatchSensitiveColumns(ctx context.Context, request PatchSensitiveColumnsRequest) (response PatchSensitiveColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchSensitiveColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchSensitiveColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchSensitiveColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchSensitiveColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchSensitiveColumnsResponse")
	}
	return
}

// patchSensitiveColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) patchSensitiveColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchSensitiveColumnsResponse
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

// ProvisionAuditPolicy Provision audit policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ProvisionAuditPolicy.go.html to see an example of how to use ProvisionAuditPolicy API.
func (client DataSafeClient) ProvisionAuditPolicy(ctx context.Context, request ProvisionAuditPolicyRequest) (response ProvisionAuditPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.provisionAuditPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ProvisionAuditPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ProvisionAuditPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ProvisionAuditPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ProvisionAuditPolicyResponse")
	}
	return
}

// provisionAuditPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) provisionAuditPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditPolicies/{auditPolicyId}/actions/provision", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ProvisionAuditPolicyResponse
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

// RefreshSecurityAssessment Runs a security assessment, refreshes the latest assessment, and saves it for future reference.
// The assessment runs with a securityAssessmentId of type LATEST. Before you start, first call the ListSecurityAssessments operation with filter "type = latest" to get the security assessment id for the target's latest assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshSecurityAssessment.go.html to see an example of how to use RefreshSecurityAssessment API.
func (client DataSafeClient) RefreshSecurityAssessment(ctx context.Context, request RefreshSecurityAssessmentRequest) (response RefreshSecurityAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshSecurityAssessmentResponse")
	}
	return
}

// refreshSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshSecurityAssessmentResponse
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

// RefreshUserAssessment Refreshes the latest assessment and saves it for future reference. This operation runs with a userAssessmentId of type LATEST.
// Before you start, first call the ListUserAssessments operation with filter "type = latest" to get the user assessment ID for
// the target's latest assessment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshUserAssessment.go.html to see an example of how to use RefreshUserAssessment API.
func (client DataSafeClient) RefreshUserAssessment(ctx context.Context, request RefreshUserAssessmentRequest) (response RefreshUserAssessmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshUserAssessmentResponse")
	}
	return
}

// refreshUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshUserAssessmentResponse
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

// ResumeAuditTrail Resumes the specified audit trail once it got stopped.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ResumeAuditTrail.go.html to see an example of how to use ResumeAuditTrail API.
func (client DataSafeClient) ResumeAuditTrail(ctx context.Context, request ResumeAuditTrailRequest) (response ResumeAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resumeAuditTrail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeAuditTrailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeAuditTrailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeAuditTrailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeAuditTrailResponse")
	}
	return
}

// resumeAuditTrail implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) resumeAuditTrail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditTrails/{auditTrailId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeAuditTrailResponse
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

// ResumeWorkRequest Resume the given work request. Issuing a resume does not guarantee of immediate resume of the work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ResumeWorkRequest.go.html to see an example of how to use ResumeWorkRequest API.
func (client DataSafeClient) ResumeWorkRequest(ctx context.Context, request ResumeWorkRequestRequest) (response ResumeWorkRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resumeWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeWorkRequestResponse")
	}
	return
}

// resumeWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) resumeWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workRequests/{workRequestId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeWorkRequestResponse
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

// RetrieveAuditPolicies Retrieves the audit policy details from the source target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RetrieveAuditPolicies.go.html to see an example of how to use RetrieveAuditPolicies API.
func (client DataSafeClient) RetrieveAuditPolicies(ctx context.Context, request RetrieveAuditPoliciesRequest) (response RetrieveAuditPoliciesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retrieveAuditPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveAuditPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveAuditPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveAuditPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveAuditPoliciesResponse")
	}
	return
}

// retrieveAuditPolicies implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) retrieveAuditPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditPolicies/{auditPolicyId}/actions/retrieveFromTarget", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveAuditPoliciesResponse
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

// SetSecurityAssessmentBaseline Sets the saved security assessment as the baseline in the compartment where the the specified assessment resides. The security assessment needs to be of type 'SAVED'.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SetSecurityAssessmentBaseline.go.html to see an example of how to use SetSecurityAssessmentBaseline API.
func (client DataSafeClient) SetSecurityAssessmentBaseline(ctx context.Context, request SetSecurityAssessmentBaselineRequest) (response SetSecurityAssessmentBaselineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.setSecurityAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetSecurityAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetSecurityAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetSecurityAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetSecurityAssessmentBaselineResponse")
	}
	return
}

// setSecurityAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) setSecurityAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/setBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetSecurityAssessmentBaselineResponse
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

// SetUserAssessmentBaseline Sets the saved user assessment as the baseline in the compartment where the specified assessment resides. The user assessment needs to be of type 'SAVED'.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SetUserAssessmentBaseline.go.html to see an example of how to use SetUserAssessmentBaseline API.
func (client DataSafeClient) SetUserAssessmentBaseline(ctx context.Context, request SetUserAssessmentBaselineRequest) (response SetUserAssessmentBaselineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.setUserAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetUserAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetUserAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetUserAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetUserAssessmentBaselineResponse")
	}
	return
}

// setUserAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) setUserAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/setBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetUserAssessmentBaselineResponse
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

// StartAuditTrail Starts collection of audit records on the specified audit trail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/StartAuditTrail.go.html to see an example of how to use StartAuditTrail API.
func (client DataSafeClient) StartAuditTrail(ctx context.Context, request StartAuditTrailRequest) (response StartAuditTrailResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startAuditTrail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartAuditTrailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartAuditTrailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAuditTrailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAuditTrailResponse")
	}
	return
}

// startAuditTrail implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) startAuditTrail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditTrails/{auditTrailId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartAuditTrailResponse
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

// StopAuditTrail Stops the specified audit trail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/StopAuditTrail.go.html to see an example of how to use StopAuditTrail API.
func (client DataSafeClient) StopAuditTrail(ctx context.Context, request StopAuditTrailRequest) (response StopAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopAuditTrail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopAuditTrailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopAuditTrailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAuditTrailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAuditTrailResponse")
	}
	return
}

// stopAuditTrail implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) stopAuditTrail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/auditTrails/{auditTrailId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopAuditTrailResponse
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

// SuspendWorkRequest Suspend the given work request. Issuing a suspend does not guarantee of a immediate suspend of the work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SuspendWorkRequest.go.html to see an example of how to use SuspendWorkRequest API.
func (client DataSafeClient) SuspendWorkRequest(ctx context.Context, request SuspendWorkRequestRequest) (response SuspendWorkRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.suspendWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuspendWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuspendWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuspendWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuspendWorkRequestResponse")
	}
	return
}

// suspendWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) suspendWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/workRequests/{workRequestId}/actions/suspend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SuspendWorkRequestResponse
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

// UnsetSecurityAssessmentBaseline Removes the baseline setting for the saved security assessment. The saved security assessment is no longer considered a baseline.
// Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UnsetSecurityAssessmentBaseline.go.html to see an example of how to use UnsetSecurityAssessmentBaseline API.
func (client DataSafeClient) UnsetSecurityAssessmentBaseline(ctx context.Context, request UnsetSecurityAssessmentBaselineRequest) (response UnsetSecurityAssessmentBaselineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.unsetSecurityAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnsetSecurityAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnsetSecurityAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnsetSecurityAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnsetSecurityAssessmentBaselineResponse")
	}
	return
}

// unsetSecurityAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) unsetSecurityAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAssessments/{securityAssessmentId}/actions/unsetBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnsetSecurityAssessmentBaselineResponse
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

// UnsetUserAssessmentBaseline Removes the baseline setting for the saved user assessment. The saved user assessment is no longer considered a baseline.
// Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UnsetUserAssessmentBaseline.go.html to see an example of how to use UnsetUserAssessmentBaseline API.
func (client DataSafeClient) UnsetUserAssessmentBaseline(ctx context.Context, request UnsetUserAssessmentBaselineRequest) (response UnsetUserAssessmentBaselineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.unsetUserAssessmentBaseline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnsetUserAssessmentBaselineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnsetUserAssessmentBaselineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnsetUserAssessmentBaselineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnsetUserAssessmentBaselineResponse")
	}
	return
}

// unsetUserAssessmentBaseline implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) unsetUserAssessmentBaseline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userAssessments/{userAssessmentId}/actions/unsetBaseline", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnsetUserAssessmentBaselineResponse
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

// UpdateAlert Updates alert status of the specified alert.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAlert.go.html to see an example of how to use UpdateAlert API.
func (client DataSafeClient) UpdateAlert(ctx context.Context, request UpdateAlertRequest) (response UpdateAlertResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAlert, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAlertResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAlertResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAlertResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAlertResponse")
	}
	return
}

// updateAlert implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateAlert(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/alerts/{alertId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAlertResponse
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

// UpdateAuditArchiveRetrieval Updates the audit archive retrieval.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditArchiveRetrieval.go.html to see an example of how to use UpdateAuditArchiveRetrieval API.
func (client DataSafeClient) UpdateAuditArchiveRetrieval(ctx context.Context, request UpdateAuditArchiveRetrievalRequest) (response UpdateAuditArchiveRetrievalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuditArchiveRetrieval, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuditArchiveRetrievalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuditArchiveRetrievalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuditArchiveRetrievalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuditArchiveRetrievalResponse")
	}
	return
}

// updateAuditArchiveRetrieval implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateAuditArchiveRetrieval(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/auditArchiveRetrievals/{auditArchiveRetrievalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuditArchiveRetrievalResponse
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

// UpdateAuditPolicy Updates the audit policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditPolicy.go.html to see an example of how to use UpdateAuditPolicy API.
func (client DataSafeClient) UpdateAuditPolicy(ctx context.Context, request UpdateAuditPolicyRequest) (response UpdateAuditPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuditPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuditPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuditPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuditPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuditPolicyResponse")
	}
	return
}

// updateAuditPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateAuditPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/auditPolicies/{auditPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuditPolicyResponse
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

// UpdateAuditProfile Updates one or more attributes of the specified audit profile.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditProfile.go.html to see an example of how to use UpdateAuditProfile API.
func (client DataSafeClient) UpdateAuditProfile(ctx context.Context, request UpdateAuditProfileRequest) (response UpdateAuditProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuditProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuditProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuditProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuditProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuditProfileResponse")
	}
	return
}

// updateAuditProfile implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateAuditProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/auditProfiles/{auditProfileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuditProfileResponse
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

// UpdateAuditTrail Updates one or more attributes of the specified audit trail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditTrail.go.html to see an example of how to use UpdateAuditTrail API.
func (client DataSafeClient) UpdateAuditTrail(ctx context.Context, request UpdateAuditTrailRequest) (response UpdateAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuditTrail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuditTrailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuditTrailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuditTrailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuditTrailResponse")
	}
	return
}

// updateAuditTrail implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateAuditTrail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/auditTrails/{auditTrailId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuditTrailResponse
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

// UpdateDataSafePrivateEndpoint Updates one or more attributes of the specified Data Safe private endpoint.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateDataSafePrivateEndpoint.go.html to see an example of how to use UpdateDataSafePrivateEndpoint API.
func (client DataSafeClient) UpdateDataSafePrivateEndpoint(ctx context.Context, request UpdateDataSafePrivateEndpointRequest) (response UpdateDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataSafePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataSafePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataSafePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataSafePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataSafePrivateEndpointResponse")
	}
	return
}

// updateDataSafePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateDataSafePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dataSafePrivateEndpoints/{dataSafePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataSafePrivateEndpointResponse
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

// UpdateLibraryMaskingFormat Updates one or more attributes of the specified library masking format. Note that updating the formatEntries attribute replaces all the existing masking format entries with the specified format entries.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateLibraryMaskingFormat.go.html to see an example of how to use UpdateLibraryMaskingFormat API.
func (client DataSafeClient) UpdateLibraryMaskingFormat(ctx context.Context, request UpdateLibraryMaskingFormatRequest) (response UpdateLibraryMaskingFormatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLibraryMaskingFormat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLibraryMaskingFormatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLibraryMaskingFormatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLibraryMaskingFormatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLibraryMaskingFormatResponse")
	}
	return
}

// updateLibraryMaskingFormat implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateLibraryMaskingFormat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/libraryMaskingFormats/{libraryMaskingFormatId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLibraryMaskingFormatResponse
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

// UpdateMaskingColumn Updates one or more attributes of the specified masking column. Note that updating the maskingFormats
// attribute replaces the currently assigned masking formats with the specified masking formats.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateMaskingColumn.go.html to see an example of how to use UpdateMaskingColumn API.
func (client DataSafeClient) UpdateMaskingColumn(ctx context.Context, request UpdateMaskingColumnRequest) (response UpdateMaskingColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaskingColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMaskingColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMaskingColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaskingColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaskingColumnResponse")
	}
	return
}

// updateMaskingColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateMaskingColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maskingPolicies/{maskingPolicyId}/maskingColumns/{maskingColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMaskingColumnResponse
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

// UpdateMaskingPolicy Updates one or more attributes of the specified masking policy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateMaskingPolicy.go.html to see an example of how to use UpdateMaskingPolicy API.
func (client DataSafeClient) UpdateMaskingPolicy(ctx context.Context, request UpdateMaskingPolicyRequest) (response UpdateMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaskingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMaskingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMaskingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaskingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaskingPolicyResponse")
	}
	return
}

// updateMaskingPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateMaskingPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maskingPolicies/{maskingPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMaskingPolicyResponse
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

// UpdateOnPremConnector Updates one or more attributes of the specified on-premises connector.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateOnPremConnector.go.html to see an example of how to use UpdateOnPremConnector API.
func (client DataSafeClient) UpdateOnPremConnector(ctx context.Context, request UpdateOnPremConnectorRequest) (response UpdateOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOnPremConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOnPremConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOnPremConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOnPremConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOnPremConnectorResponse")
	}
	return
}

// updateOnPremConnector implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateOnPremConnector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/onPremConnectors/{onPremConnectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOnPremConnectorResponse
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

// UpdateOnPremConnectorWallet Updates the wallet for the specified on-premises connector to a new version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateOnPremConnectorWallet.go.html to see an example of how to use UpdateOnPremConnectorWallet API.
func (client DataSafeClient) UpdateOnPremConnectorWallet(ctx context.Context, request UpdateOnPremConnectorWalletRequest) (response UpdateOnPremConnectorWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateOnPremConnectorWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOnPremConnectorWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOnPremConnectorWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOnPremConnectorWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOnPremConnectorWalletResponse")
	}
	return
}

// updateOnPremConnectorWallet implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateOnPremConnectorWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/onPremConnectors/{onPremConnectorId}/wallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOnPremConnectorWalletResponse
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

// UpdateReportDefinition Updates the specified report definition. Only user created report definition can be updated. Seeded report definitions need to be saved as new report definition first.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateReportDefinition.go.html to see an example of how to use UpdateReportDefinition API.
func (client DataSafeClient) UpdateReportDefinition(ctx context.Context, request UpdateReportDefinitionRequest) (response UpdateReportDefinitionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateReportDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateReportDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateReportDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateReportDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateReportDefinitionResponse")
	}
	return
}

// updateReportDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateReportDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/reportDefinitions/{reportDefinitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateReportDefinitionResponse
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

// UpdateSecurityAssessment Updates one or more attributes of the specified security assessment. This operation allows to update the security assessment displayName, description, or schedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSecurityAssessment.go.html to see an example of how to use UpdateSecurityAssessment API.
func (client DataSafeClient) UpdateSecurityAssessment(ctx context.Context, request UpdateSecurityAssessmentRequest) (response UpdateSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityAssessmentResponse")
	}
	return
}

// updateSecurityAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSecurityAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityAssessments/{securityAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityAssessmentResponse
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

// UpdateSensitiveColumn Updates one or more attributes of the specified sensitive column.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSensitiveColumn.go.html to see an example of how to use UpdateSensitiveColumn API.
func (client DataSafeClient) UpdateSensitiveColumn(ctx context.Context, request UpdateSensitiveColumnRequest) (response UpdateSensitiveColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSensitiveColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSensitiveColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSensitiveColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSensitiveColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSensitiveColumnResponse")
	}
	return
}

// updateSensitiveColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSensitiveColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveColumns/{sensitiveColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSensitiveColumnResponse
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

// UpdateSensitiveDataModel Updates one or more attributes of the specified sensitive data model. Note that updating any attribute of a sensitive
// data model does not perform data discovery.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSensitiveDataModel.go.html to see an example of how to use UpdateSensitiveDataModel API.
func (client DataSafeClient) UpdateSensitiveDataModel(ctx context.Context, request UpdateSensitiveDataModelRequest) (response UpdateSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSensitiveDataModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSensitiveDataModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSensitiveDataModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSensitiveDataModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSensitiveDataModelResponse")
	}
	return
}

// updateSensitiveDataModel implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSensitiveDataModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sensitiveDataModels/{sensitiveDataModelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSensitiveDataModelResponse
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

// UpdateSensitiveType Updates one or more attributes of the specified sensitive type.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSensitiveType.go.html to see an example of how to use UpdateSensitiveType API.
func (client DataSafeClient) UpdateSensitiveType(ctx context.Context, request UpdateSensitiveTypeRequest) (response UpdateSensitiveTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSensitiveType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSensitiveTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSensitiveTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSensitiveTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSensitiveTypeResponse")
	}
	return
}

// updateSensitiveType implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSensitiveType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sensitiveTypes/{sensitiveTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSensitiveTypeResponse
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

// UpdateTargetAlertPolicyAssociation Updates the specified target-alert policy association.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateTargetAlertPolicyAssociation.go.html to see an example of how to use UpdateTargetAlertPolicyAssociation API.
func (client DataSafeClient) UpdateTargetAlertPolicyAssociation(ctx context.Context, request UpdateTargetAlertPolicyAssociationRequest) (response UpdateTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetAlertPolicyAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetAlertPolicyAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetAlertPolicyAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetAlertPolicyAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetAlertPolicyAssociationResponse")
	}
	return
}

// updateTargetAlertPolicyAssociation implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateTargetAlertPolicyAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targetAlertPolicyAssociations/{targetAlertPolicyAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetAlertPolicyAssociationResponse
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

// UpdateTargetDatabase Updates one or more attributes of the specified Data Safe target database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateTargetDatabase.go.html to see an example of how to use UpdateTargetDatabase API.
func (client DataSafeClient) UpdateTargetDatabase(ctx context.Context, request UpdateTargetDatabaseRequest) (response UpdateTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetDatabaseResponse")
	}
	return
}

// updateTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targetDatabases/{targetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetDatabaseResponse
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

// UpdateUserAssessment Updates one or more attributes of the specified user assessment. This operation allows to update the user assessment displayName, description, or schedule.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateUserAssessment.go.html to see an example of how to use UpdateUserAssessment API.
func (client DataSafeClient) UpdateUserAssessment(ctx context.Context, request UpdateUserAssessmentRequest) (response UpdateUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUserAssessment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUserAssessmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUserAssessmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUserAssessmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUserAssessmentResponse")
	}
	return
}

// updateUserAssessment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateUserAssessment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/userAssessments/{userAssessmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUserAssessmentResponse
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

// UploadMaskingPolicy Uploads a masking policy file (also called template) to update the specified masking policy.
// To create a new masking policy using a file, first use the CreateMaskingPolicy operation
// to create an empty masking policy and then use this endpoint to upload the masking policy file.
// Note that the upload operation replaces the content of the specified masking policy,
// including all the existing columns and masking formats, with the content of the file.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UploadMaskingPolicy.go.html to see an example of how to use UploadMaskingPolicy API.
func (client DataSafeClient) UploadMaskingPolicy(ctx context.Context, request UploadMaskingPolicyRequest) (response UploadMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.uploadMaskingPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadMaskingPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadMaskingPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadMaskingPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadMaskingPolicyResponse")
	}
	return
}

// uploadMaskingPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) uploadMaskingPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/upload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadMaskingPolicyResponse
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

// UploadSensitiveDataModel Uploads a sensitive data model file (also called template) to update the specified sensitive data model. To create
// a new sensitive data model using a file, first use the CreateSensitiveDataModel operation to create an empty data model
// and then use this endpoint to upload the data model file. Note that the upload operation replaces the content of the
// specified sensitive data model, including all the existing columns and their relationships, with the content of the file.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UploadSensitiveDataModel.go.html to see an example of how to use UploadSensitiveDataModel API.
func (client DataSafeClient) UploadSensitiveDataModel(ctx context.Context, request UploadSensitiveDataModelRequest) (response UploadSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.uploadSensitiveDataModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadSensitiveDataModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadSensitiveDataModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadSensitiveDataModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadSensitiveDataModelResponse")
	}
	return
}

// uploadSensitiveDataModel implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) uploadSensitiveDataModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sensitiveDataModels/{sensitiveDataModelId}/actions/upload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadSensitiveDataModelResponse
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
