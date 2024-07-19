// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DataSafeClient a client for DataSafe
type DataSafeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataSafeClientWithConfigurationProvider Creates a new default DataSafe client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataSafeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataSafeClient, err error) {
	if enabled := common.CheckForEnabledServices("datasafe"); !enabled {
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
	return newDataSafeClientFromBaseClient(baseClient, provider)
}

// NewDataSafeClientWithOboToken Creates a new default DataSafe client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDataSafeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataSafeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataSafeClientFromBaseClient(baseClient, configProvider)
}

func newDataSafeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataSafeClient, err error) {
	// DataSafe service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DataSafe"))
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
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DataSafeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateTargetDatabase Reactivates a previously deactivated Data Safe target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ActivateTargetDatabase.go.html to see an example of how to use ActivateTargetDatabase API.
// A default retry strategy applies to this operation ActivateTargetDatabase()
func (client DataSafeClient) ActivateTargetDatabase(ctx context.Context, request ActivateTargetDatabaseRequest) (response ActivateTargetDatabaseResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/ActivateTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "ActivateTargetDatabase", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/AddMaskingColumnsFromSdm.go.html to see an example of how to use AddMaskingColumnsFromSdm API.
// A default retry strategy applies to this operation AddMaskingColumnsFromSdm()
func (client DataSafeClient) AddMaskingColumnsFromSdm(ctx context.Context, request AddMaskingColumnsFromSdmRequest) (response AddMaskingColumnsFromSdmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/AddMaskingColumnsFromSdm"
		err = common.PostProcessServiceError(err, "DataSafe", "AddMaskingColumnsFromSdm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AlertsUpdate Updates alerts in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/AlertsUpdate.go.html to see an example of how to use AlertsUpdate API.
// A default retry strategy applies to this operation AlertsUpdate()
func (client DataSafeClient) AlertsUpdate(ctx context.Context, request AlertsUpdateRequest) (response AlertsUpdateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.alertsUpdate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AlertsUpdateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AlertsUpdateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AlertsUpdateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AlertsUpdateResponse")
	}
	return
}

// alertsUpdate implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) alertsUpdate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alerts/actions/updateAll", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AlertsUpdateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Alert/AlertsUpdate"
		err = common.PostProcessServiceError(err, "DataSafe", "AlertsUpdate", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ApplyDiscoveryJobResults.go.html to see an example of how to use ApplyDiscoveryJobResults API.
// A default retry strategy applies to this operation ApplyDiscoveryJobResults()
func (client DataSafeClient) ApplyDiscoveryJobResults(ctx context.Context, request ApplyDiscoveryJobResultsRequest) (response ApplyDiscoveryJobResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/ApplyDiscoveryJobResults"
		err = common.PostProcessServiceError(err, "DataSafe", "ApplyDiscoveryJobResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ApplySdmMaskingPolicyDifference Applies the difference of a SDM Masking policy difference resource to the specified masking policy. Note that the plannedAction attribute
// of difference columns is used for processing. You should first use PatchSdmMaskingPolicyDifferenceColumns to set the plannedAction
// attribute of the difference columns you want to process. ApplySdmMaskingPolicyDifference automatically reads the plannedAction
// attribute and updates the masking policy to reflect the actions you planned. If the sdmMaskingPolicydifferenceId is not passed, the
// latest sdmMaskingPolicydifference is used. Note that if the masking policy associated with the SdmMaskingPolicyDifference used for this
// operation is not associated with the original SDM anymore, this operation won't be allowed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ApplySdmMaskingPolicyDifference.go.html to see an example of how to use ApplySdmMaskingPolicyDifference API.
// A default retry strategy applies to this operation ApplySdmMaskingPolicyDifference()
func (client DataSafeClient) ApplySdmMaskingPolicyDifference(ctx context.Context, request ApplySdmMaskingPolicyDifferenceRequest) (response ApplySdmMaskingPolicyDifferenceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.applySdmMaskingPolicyDifference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApplySdmMaskingPolicyDifferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApplySdmMaskingPolicyDifferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApplySdmMaskingPolicyDifferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApplySdmMaskingPolicyDifferenceResponse")
	}
	return
}

// applySdmMaskingPolicyDifference implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) applySdmMaskingPolicyDifference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/maskingColumns/actions/applyDifferenceToMaskingColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApplySdmMaskingPolicyDifferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/ApplySdmMaskingPolicyDifference"
		err = common.PostProcessServiceError(err, "DataSafe", "ApplySdmMaskingPolicyDifference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CalculateAuditVolumeAvailable Calculates the volume of audit events available on the target database to be collected. Measurable up to the defined retention period of the audit target resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CalculateAuditVolumeAvailable.go.html to see an example of how to use CalculateAuditVolumeAvailable API.
// A default retry strategy applies to this operation CalculateAuditVolumeAvailable()
func (client DataSafeClient) CalculateAuditVolumeAvailable(ctx context.Context, request CalculateAuditVolumeAvailableRequest) (response CalculateAuditVolumeAvailableResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/CalculateAuditVolumeAvailable"
		err = common.PostProcessServiceError(err, "DataSafe", "CalculateAuditVolumeAvailable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CalculateAuditVolumeCollected Calculates the volume of audit events collected by data safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CalculateAuditVolumeCollected.go.html to see an example of how to use CalculateAuditVolumeCollected API.
// A default retry strategy applies to this operation CalculateAuditVolumeCollected()
func (client DataSafeClient) CalculateAuditVolumeCollected(ctx context.Context, request CalculateAuditVolumeCollectedRequest) (response CalculateAuditVolumeCollectedResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/CalculateAuditVolumeCollected"
		err = common.PostProcessServiceError(err, "DataSafe", "CalculateAuditVolumeCollected", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancel the specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client DataSafeClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "DataSafe", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAlertCompartment Moves the specified alert into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAlertCompartment.go.html to see an example of how to use ChangeAlertCompartment API.
// A default retry strategy applies to this operation ChangeAlertCompartment()
func (client DataSafeClient) ChangeAlertCompartment(ctx context.Context, request ChangeAlertCompartmentRequest) (response ChangeAlertCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Alert/ChangeAlertCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeAlertCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAuditArchiveRetrievalCompartment Moves the archive retreival to the specified compartment. When provided, if-Match is checked against ETag value of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAuditArchiveRetrievalCompartment.go.html to see an example of how to use ChangeAuditArchiveRetrievalCompartment API.
// A default retry strategy applies to this operation ChangeAuditArchiveRetrievalCompartment()
func (client DataSafeClient) ChangeAuditArchiveRetrievalCompartment(ctx context.Context, request ChangeAuditArchiveRetrievalCompartmentRequest) (response ChangeAuditArchiveRetrievalCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditArchiveRetrieval/ChangeAuditArchiveRetrievalCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeAuditArchiveRetrievalCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAuditPolicyCompartment Moves the specified audit policy and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAuditPolicyCompartment.go.html to see an example of how to use ChangeAuditPolicyCompartment API.
// A default retry strategy applies to this operation ChangeAuditPolicyCompartment()
func (client DataSafeClient) ChangeAuditPolicyCompartment(ctx context.Context, request ChangeAuditPolicyCompartmentRequest) (response ChangeAuditPolicyCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicy/ChangeAuditPolicyCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeAuditPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAuditProfileCompartment Moves the specified audit profile and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeAuditProfileCompartment.go.html to see an example of how to use ChangeAuditProfileCompartment API.
// A default retry strategy applies to this operation ChangeAuditProfileCompartment()
func (client DataSafeClient) ChangeAuditProfileCompartment(ctx context.Context, request ChangeAuditProfileCompartmentRequest) (response ChangeAuditProfileCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/ChangeAuditProfileCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeAuditProfileCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDataSafePrivateEndpointCompartment Moves the Data Safe private endpoint and its dependent resources to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeDataSafePrivateEndpointCompartment.go.html to see an example of how to use ChangeDataSafePrivateEndpointCompartment API.
// A default retry strategy applies to this operation ChangeDataSafePrivateEndpointCompartment()
func (client DataSafeClient) ChangeDataSafePrivateEndpointCompartment(ctx context.Context, request ChangeDataSafePrivateEndpointCompartmentRequest) (response ChangeDataSafePrivateEndpointCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafePrivateEndpoint/ChangeDataSafePrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeDataSafePrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseSecurityConfigCompartment Moves the specified database security configuration and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeDatabaseSecurityConfigCompartment.go.html to see an example of how to use ChangeDatabaseSecurityConfigCompartment API.
// A default retry strategy applies to this operation ChangeDatabaseSecurityConfigCompartment()
func (client DataSafeClient) ChangeDatabaseSecurityConfigCompartment(ctx context.Context, request ChangeDatabaseSecurityConfigCompartmentRequest) (response ChangeDatabaseSecurityConfigCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseSecurityConfigCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseSecurityConfigCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseSecurityConfigCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseSecurityConfigCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseSecurityConfigCompartmentResponse")
	}
	return
}

// changeDatabaseSecurityConfigCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeDatabaseSecurityConfigCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseSecurityConfigs/{databaseSecurityConfigId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseSecurityConfigCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseSecurityConfig/ChangeDatabaseSecurityConfigCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeDatabaseSecurityConfigCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDiscoveryJobCompartment Moves the specified discovery job and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeDiscoveryJobCompartment.go.html to see an example of how to use ChangeDiscoveryJobCompartment API.
// A default retry strategy applies to this operation ChangeDiscoveryJobCompartment()
func (client DataSafeClient) ChangeDiscoveryJobCompartment(ctx context.Context, request ChangeDiscoveryJobCompartmentRequest) (response ChangeDiscoveryJobCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJob/ChangeDiscoveryJobCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeDiscoveryJobCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLibraryMaskingFormatCompartment Moves the specified library masking format into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeLibraryMaskingFormatCompartment.go.html to see an example of how to use ChangeLibraryMaskingFormatCompartment API.
// A default retry strategy applies to this operation ChangeLibraryMaskingFormatCompartment()
func (client DataSafeClient) ChangeLibraryMaskingFormatCompartment(ctx context.Context, request ChangeLibraryMaskingFormatCompartmentRequest) (response ChangeLibraryMaskingFormatCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/LibraryMaskingFormat/ChangeLibraryMaskingFormatCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeLibraryMaskingFormatCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMaskingPolicyCompartment Moves the specified masking policy and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeMaskingPolicyCompartment.go.html to see an example of how to use ChangeMaskingPolicyCompartment API.
// A default retry strategy applies to this operation ChangeMaskingPolicyCompartment()
func (client DataSafeClient) ChangeMaskingPolicyCompartment(ctx context.Context, request ChangeMaskingPolicyCompartmentRequest) (response ChangeMaskingPolicyCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/ChangeMaskingPolicyCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeMaskingPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMaskingPolicyHealthReportCompartment Moves the specified masking policy health report and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeMaskingPolicyHealthReportCompartment.go.html to see an example of how to use ChangeMaskingPolicyHealthReportCompartment API.
// A default retry strategy applies to this operation ChangeMaskingPolicyHealthReportCompartment()
func (client DataSafeClient) ChangeMaskingPolicyHealthReportCompartment(ctx context.Context, request ChangeMaskingPolicyHealthReportCompartmentRequest) (response ChangeMaskingPolicyHealthReportCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMaskingPolicyHealthReportCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMaskingPolicyHealthReportCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMaskingPolicyHealthReportCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMaskingPolicyHealthReportCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMaskingPolicyHealthReportCompartmentResponse")
	}
	return
}

// changeMaskingPolicyHealthReportCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeMaskingPolicyHealthReportCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicyHealthReports/{maskingPolicyHealthReportId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMaskingPolicyHealthReportCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicyHealthReport/ChangeMaskingPolicyHealthReportCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeMaskingPolicyHealthReportCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOnPremConnectorCompartment Moves the specified on-premises connector into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeOnPremConnectorCompartment.go.html to see an example of how to use ChangeOnPremConnectorCompartment API.
// A default retry strategy applies to this operation ChangeOnPremConnectorCompartment()
func (client DataSafeClient) ChangeOnPremConnectorCompartment(ctx context.Context, request ChangeOnPremConnectorCompartmentRequest) (response ChangeOnPremConnectorCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/ChangeOnPremConnectorCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeOnPremConnectorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeReportCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeReportCompartment.go.html to see an example of how to use ChangeReportCompartment API.
// A default retry strategy applies to this operation ChangeReportCompartment()
func (client DataSafeClient) ChangeReportCompartment(ctx context.Context, request ChangeReportCompartmentRequest) (response ChangeReportCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Report/ChangeReportCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeReportCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeReportDefinitionCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeReportDefinitionCompartment.go.html to see an example of how to use ChangeReportDefinitionCompartment API.
// A default retry strategy applies to this operation ChangeReportDefinitionCompartment()
func (client DataSafeClient) ChangeReportDefinitionCompartment(ctx context.Context, request ChangeReportDefinitionCompartmentRequest) (response ChangeReportDefinitionCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/ChangeReportDefinitionCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeReportDefinitionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeRetention Change the online and offline months .
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeRetention.go.html to see an example of how to use ChangeRetention API.
// A default retry strategy applies to this operation ChangeRetention()
func (client DataSafeClient) ChangeRetention(ctx context.Context, request ChangeRetentionRequest) (response ChangeRetentionResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/ChangeRetention"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeRetention", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSdmMaskingPolicyDifferenceCompartment Moves the specified SDM masking policy difference into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSdmMaskingPolicyDifferenceCompartment.go.html to see an example of how to use ChangeSdmMaskingPolicyDifferenceCompartment API.
// A default retry strategy applies to this operation ChangeSdmMaskingPolicyDifferenceCompartment()
func (client DataSafeClient) ChangeSdmMaskingPolicyDifferenceCompartment(ctx context.Context, request ChangeSdmMaskingPolicyDifferenceCompartmentRequest) (response ChangeSdmMaskingPolicyDifferenceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSdmMaskingPolicyDifferenceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSdmMaskingPolicyDifferenceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSdmMaskingPolicyDifferenceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSdmMaskingPolicyDifferenceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSdmMaskingPolicyDifferenceCompartmentResponse")
	}
	return
}

// changeSdmMaskingPolicyDifferenceCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSdmMaskingPolicyDifferenceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSdmMaskingPolicyDifferenceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/ChangeSdmMaskingPolicyDifferenceCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSdmMaskingPolicyDifferenceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityAssessmentCompartment Moves the specified saved security assessment or future scheduled assessments into a different compartment.
// To start, call first the operation ListSecurityAssessments with filters "type = save_schedule". This returns the scheduleAssessmentId. Then, call this changeCompartment with the scheduleAssessmentId.
// The existing saved security assessments created due to the schedule are not moved. However, all new saves will be associated with the new compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSecurityAssessmentCompartment.go.html to see an example of how to use ChangeSecurityAssessmentCompartment API.
// A default retry strategy applies to this operation ChangeSecurityAssessmentCompartment()
func (client DataSafeClient) ChangeSecurityAssessmentCompartment(ctx context.Context, request ChangeSecurityAssessmentCompartmentRequest) (response ChangeSecurityAssessmentCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/ChangeSecurityAssessmentCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSecurityAssessmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityPolicyCompartment Moves the specified security policy and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSecurityPolicyCompartment.go.html to see an example of how to use ChangeSecurityPolicyCompartment API.
// A default retry strategy applies to this operation ChangeSecurityPolicyCompartment()
func (client DataSafeClient) ChangeSecurityPolicyCompartment(ctx context.Context, request ChangeSecurityPolicyCompartmentRequest) (response ChangeSecurityPolicyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityPolicyCompartmentResponse")
	}
	return
}

// changeSecurityPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSecurityPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityPolicies/{securityPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicy/ChangeSecurityPolicyCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSecurityPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityPolicyDeploymentCompartment Moves the specified security policy deployment and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSecurityPolicyDeploymentCompartment.go.html to see an example of how to use ChangeSecurityPolicyDeploymentCompartment API.
// A default retry strategy applies to this operation ChangeSecurityPolicyDeploymentCompartment()
func (client DataSafeClient) ChangeSecurityPolicyDeploymentCompartment(ctx context.Context, request ChangeSecurityPolicyDeploymentCompartmentRequest) (response ChangeSecurityPolicyDeploymentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityPolicyDeploymentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityPolicyDeploymentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityPolicyDeploymentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityPolicyDeploymentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityPolicyDeploymentCompartmentResponse")
	}
	return
}

// changeSecurityPolicyDeploymentCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSecurityPolicyDeploymentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityPolicyDeployments/{securityPolicyDeploymentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityPolicyDeploymentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyDeployment/ChangeSecurityPolicyDeploymentCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSecurityPolicyDeploymentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSensitiveDataModelCompartment Moves the specified sensitive data model and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSensitiveDataModelCompartment.go.html to see an example of how to use ChangeSensitiveDataModelCompartment API.
// A default retry strategy applies to this operation ChangeSensitiveDataModelCompartment()
func (client DataSafeClient) ChangeSensitiveDataModelCompartment(ctx context.Context, request ChangeSensitiveDataModelCompartmentRequest) (response ChangeSensitiveDataModelCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/ChangeSensitiveDataModelCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSensitiveDataModelCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSensitiveTypeCompartment Moves the specified sensitive type into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSensitiveTypeCompartment.go.html to see an example of how to use ChangeSensitiveTypeCompartment API.
// A default retry strategy applies to this operation ChangeSensitiveTypeCompartment()
func (client DataSafeClient) ChangeSensitiveTypeCompartment(ctx context.Context, request ChangeSensitiveTypeCompartmentRequest) (response ChangeSensitiveTypeCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveType/ChangeSensitiveTypeCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSensitiveTypeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSqlCollectionCompartment Moves the specified SQL collection and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSqlCollectionCompartment.go.html to see an example of how to use ChangeSqlCollectionCompartment API.
// A default retry strategy applies to this operation ChangeSqlCollectionCompartment()
func (client DataSafeClient) ChangeSqlCollectionCompartment(ctx context.Context, request ChangeSqlCollectionCompartmentRequest) (response ChangeSqlCollectionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSqlCollectionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSqlCollectionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSqlCollectionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSqlCollectionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSqlCollectionCompartmentResponse")
	}
	return
}

// changeSqlCollectionCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSqlCollectionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections/{sqlCollectionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSqlCollectionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/ChangeSqlCollectionCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSqlCollectionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSqlFirewallPolicyCompartment Moves the specified SQL Firewall policy and its dependent resources into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeSqlFirewallPolicyCompartment.go.html to see an example of how to use ChangeSqlFirewallPolicyCompartment API.
// A default retry strategy applies to this operation ChangeSqlFirewallPolicyCompartment()
func (client DataSafeClient) ChangeSqlFirewallPolicyCompartment(ctx context.Context, request ChangeSqlFirewallPolicyCompartmentRequest) (response ChangeSqlFirewallPolicyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSqlFirewallPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSqlFirewallPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSqlFirewallPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSqlFirewallPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSqlFirewallPolicyCompartmentResponse")
	}
	return
}

// changeSqlFirewallPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) changeSqlFirewallPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlFirewallPolicies/{sqlFirewallPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSqlFirewallPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallPolicy/ChangeSqlFirewallPolicyCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeSqlFirewallPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTargetAlertPolicyAssociationCompartment Moves the specified target-alert policy Association into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeTargetAlertPolicyAssociationCompartment.go.html to see an example of how to use ChangeTargetAlertPolicyAssociationCompartment API.
// A default retry strategy applies to this operation ChangeTargetAlertPolicyAssociationCompartment()
func (client DataSafeClient) ChangeTargetAlertPolicyAssociationCompartment(ctx context.Context, request ChangeTargetAlertPolicyAssociationCompartmentRequest) (response ChangeTargetAlertPolicyAssociationCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociation/ChangeTargetAlertPolicyAssociationCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeTargetAlertPolicyAssociationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTargetDatabaseCompartment Moves the Data Safe target database to the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeTargetDatabaseCompartment.go.html to see an example of how to use ChangeTargetDatabaseCompartment API.
// A default retry strategy applies to this operation ChangeTargetDatabaseCompartment()
func (client DataSafeClient) ChangeTargetDatabaseCompartment(ctx context.Context, request ChangeTargetDatabaseCompartmentRequest) (response ChangeTargetDatabaseCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/ChangeTargetDatabaseCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeTargetDatabaseCompartment", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ChangeUserAssessmentCompartment.go.html to see an example of how to use ChangeUserAssessmentCompartment API.
// A default retry strategy applies to this operation ChangeUserAssessmentCompartment()
func (client DataSafeClient) ChangeUserAssessmentCompartment(ctx context.Context, request ChangeUserAssessmentCompartmentRequest) (response ChangeUserAssessmentCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/ChangeUserAssessmentCompartment"
		err = common.PostProcessServiceError(err, "DataSafe", "ChangeUserAssessmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CompareSecurityAssessment Compares two security assessments. For this comparison, a security assessment can be a saved assessment, a latest assessment, or a baseline assessment.
// For example, you can compare saved assessment or a latest assessment against a baseline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CompareSecurityAssessment.go.html to see an example of how to use CompareSecurityAssessment API.
// A default retry strategy applies to this operation CompareSecurityAssessment()
func (client DataSafeClient) CompareSecurityAssessment(ctx context.Context, request CompareSecurityAssessmentRequest) (response CompareSecurityAssessmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/CompareSecurityAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "CompareSecurityAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CompareUserAssessment Compares two user assessments. For this comparison, a user assessment can be a saved, a latest assessment, or a baseline.
// As an example, it can be used to compare a user assessment saved or a latest assessment with a baseline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CompareUserAssessment.go.html to see an example of how to use CompareUserAssessment API.
// A default retry strategy applies to this operation CompareUserAssessment()
func (client DataSafeClient) CompareUserAssessment(ctx context.Context, request CompareUserAssessmentRequest) (response CompareUserAssessmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/CompareUserAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "CompareUserAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAuditArchiveRetrieval Creates a work request to retrieve archived audit data. This asynchronous process will usually take over an hour to complete.
// Save the id from the response of this operation. Call GetAuditArchiveRetrieval operation after an hour, passing the id to know the status of
// this operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateAuditArchiveRetrieval.go.html to see an example of how to use CreateAuditArchiveRetrieval API.
// A default retry strategy applies to this operation CreateAuditArchiveRetrieval()
func (client DataSafeClient) CreateAuditArchiveRetrieval(ctx context.Context, request CreateAuditArchiveRetrievalRequest) (response CreateAuditArchiveRetrievalResponse, err error) {
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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataSafe", "CreateAuditArchiveRetrieval", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataSafePrivateEndpoint Creates a new Data Safe private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateDataSafePrivateEndpoint.go.html to see an example of how to use CreateDataSafePrivateEndpoint API.
// A default retry strategy applies to this operation CreateDataSafePrivateEndpoint()
func (client DataSafeClient) CreateDataSafePrivateEndpoint(ctx context.Context, request CreateDataSafePrivateEndpointRequest) (response CreateDataSafePrivateEndpointResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafePrivateEndpoint/CreateDataSafePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateDataSafePrivateEndpoint", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateDiscoveryJob.go.html to see an example of how to use CreateDiscoveryJob API.
// A default retry strategy applies to this operation CreateDiscoveryJob()
func (client DataSafeClient) CreateDiscoveryJob(ctx context.Context, request CreateDiscoveryJobRequest) (response CreateDiscoveryJobResponse, err error) {
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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataSafe", "CreateDiscoveryJob", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateLibraryMaskingFormat.go.html to see an example of how to use CreateLibraryMaskingFormat API.
// A default retry strategy applies to this operation CreateLibraryMaskingFormat()
func (client DataSafeClient) CreateLibraryMaskingFormat(ctx context.Context, request CreateLibraryMaskingFormatRequest) (response CreateLibraryMaskingFormatResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/LibraryMaskingFormat/CreateLibraryMaskingFormat"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateLibraryMaskingFormat", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateMaskingColumn.go.html to see an example of how to use CreateMaskingColumn API.
// A default retry strategy applies to this operation CreateMaskingColumn()
func (client DataSafeClient) CreateMaskingColumn(ctx context.Context, request CreateMaskingColumnRequest) (response CreateMaskingColumnResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/CreateMaskingColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateMaskingColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMaskingPolicy Creates a new masking policy and associates it with a sensitive data model or a target database.
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateMaskingPolicy.go.html to see an example of how to use CreateMaskingPolicy API.
// A default retry strategy applies to this operation CreateMaskingPolicy()
func (client DataSafeClient) CreateMaskingPolicy(ctx context.Context, request CreateMaskingPolicyRequest) (response CreateMaskingPolicyResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/CreateMaskingPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateMaskingPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOnPremConnector Creates a new on-premises connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateOnPremConnector.go.html to see an example of how to use CreateOnPremConnector API.
// A default retry strategy applies to this operation CreateOnPremConnector()
func (client DataSafeClient) CreateOnPremConnector(ctx context.Context, request CreateOnPremConnectorRequest) (response CreateOnPremConnectorResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/CreateOnPremConnector"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateOnPremConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePeerTargetDatabase Creates the peer target database under the primary target database in Data Safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreatePeerTargetDatabase.go.html to see an example of how to use CreatePeerTargetDatabase API.
// A default retry strategy applies to this operation CreatePeerTargetDatabase()
func (client DataSafeClient) CreatePeerTargetDatabase(ctx context.Context, request CreatePeerTargetDatabaseRequest) (response CreatePeerTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPeerTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePeerTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePeerTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePeerTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePeerTargetDatabaseResponse")
	}
	return
}

// createPeerTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createPeerTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/peerTargetDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePeerTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/PeerTargetDatabase/CreatePeerTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "CreatePeerTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateReportDefinition Creates a new report definition with parameters specified in the body. The report definition is stored in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateReportDefinition.go.html to see an example of how to use CreateReportDefinition API.
// A default retry strategy applies to this operation CreateReportDefinition()
func (client DataSafeClient) CreateReportDefinition(ctx context.Context, request CreateReportDefinitionRequest) (response CreateReportDefinitionResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/CreateReportDefinition"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateReportDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSdmMaskingPolicyDifference Creates SDM masking policy difference for the specified masking policy. It finds the difference between
// masking columns of the masking policy and sensitive columns of the SDM. After performing this operation,
// you can use ListDifferenceColumns to view the difference columns, PatchSdmMaskingPolicyDifferenceColumns
// to specify the action you want perform on these columns, and then ApplySdmMaskingPolicyDifference to process the
// difference columns and apply them to the masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSdmMaskingPolicyDifference.go.html to see an example of how to use CreateSdmMaskingPolicyDifference API.
// A default retry strategy applies to this operation CreateSdmMaskingPolicyDifference()
func (client DataSafeClient) CreateSdmMaskingPolicyDifference(ctx context.Context, request CreateSdmMaskingPolicyDifferenceRequest) (response CreateSdmMaskingPolicyDifferenceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSdmMaskingPolicyDifference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSdmMaskingPolicyDifferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSdmMaskingPolicyDifferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSdmMaskingPolicyDifferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSdmMaskingPolicyDifferenceResponse")
	}
	return
}

// createSdmMaskingPolicyDifference implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSdmMaskingPolicyDifference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sdmMaskingPolicyDifferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSdmMaskingPolicyDifferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataSafe", "CreateSdmMaskingPolicyDifference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityAssessment Creates a new saved security assessment for one or multiple targets in a compartment. When this operation is performed,
// it will save the latest assessments in the specified compartment. If a schedule is passed, it will persist the latest assessments,
// at the defined date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSecurityAssessment.go.html to see an example of how to use CreateSecurityAssessment API.
// A default retry strategy applies to this operation CreateSecurityAssessment()
func (client DataSafeClient) CreateSecurityAssessment(ctx context.Context, request CreateSecurityAssessmentRequest) (response CreateSecurityAssessmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/CreateSecurityAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateSecurityAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSensitiveColumn Creates a new sensitive column in the specified sensitive data model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSensitiveColumn.go.html to see an example of how to use CreateSensitiveColumn API.
// A default retry strategy applies to this operation CreateSensitiveColumn()
func (client DataSafeClient) CreateSensitiveColumn(ctx context.Context, request CreateSensitiveColumnRequest) (response CreateSensitiveColumnResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveColumn/CreateSensitiveColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateSensitiveColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSensitiveDataModel Creates a new sensitive data model. If schemas and sensitive types are provided, it automatically runs data discovery
// and adds the discovered columns to the sensitive data model. Otherwise, it creates an empty sensitive data model
// that can be updated later.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSensitiveDataModel.go.html to see an example of how to use CreateSensitiveDataModel API.
// A default retry strategy applies to this operation CreateSensitiveDataModel()
func (client DataSafeClient) CreateSensitiveDataModel(ctx context.Context, request CreateSensitiveDataModelRequest) (response CreateSensitiveDataModelResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/CreateSensitiveDataModel"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateSensitiveDataModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSensitiveType Creates a new sensitive type, which can be a basic sensitive type with regular expressions or a sensitive category.
// While sensitive types are used for data discovery, sensitive categories are used for logically grouping the related
// or similar sensitive types.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSensitiveType.go.html to see an example of how to use CreateSensitiveType API.
// A default retry strategy applies to this operation CreateSensitiveType()
func (client DataSafeClient) CreateSensitiveType(ctx context.Context, request CreateSensitiveTypeRequest) (response CreateSensitiveTypeResponse, err error) {
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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataSafe", "CreateSensitiveType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &sensitivetype{})
	return response, err
}

// CreateSqlCollection Creates a new SQL collection resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateSqlCollection.go.html to see an example of how to use CreateSqlCollection API.
// A default retry strategy applies to this operation CreateSqlCollection()
func (client DataSafeClient) CreateSqlCollection(ctx context.Context, request CreateSqlCollectionRequest) (response CreateSqlCollectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSqlCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSqlCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSqlCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSqlCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSqlCollectionResponse")
	}
	return
}

// createSqlCollection implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) createSqlCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSqlCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/CreateSqlCollection"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateSqlCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetAlertPolicyAssociation Creates a new target-alert policy association to track a alert policy applied on target.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateTargetAlertPolicyAssociation.go.html to see an example of how to use CreateTargetAlertPolicyAssociation API.
// A default retry strategy applies to this operation CreateTargetAlertPolicyAssociation()
func (client DataSafeClient) CreateTargetAlertPolicyAssociation(ctx context.Context, request CreateTargetAlertPolicyAssociationRequest) (response CreateTargetAlertPolicyAssociationResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociation/CreateTargetAlertPolicyAssociation"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateTargetAlertPolicyAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetDatabase Registers the specified database with Data Safe and creates a Data Safe target database in the Data Safe Console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateTargetDatabase.go.html to see an example of how to use CreateTargetDatabase API.
// A default retry strategy applies to this operation CreateTargetDatabase()
func (client DataSafeClient) CreateTargetDatabase(ctx context.Context, request CreateTargetDatabaseRequest) (response CreateTargetDatabaseResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/CreateTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUserAssessment Creates a new saved user assessment for one or multiple targets in a compartment. It saves the latest assessments in the
// specified compartment. If a scheduled is passed in, this operation persists the latest assessments that exist at the defined
// date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/CreateUserAssessment.go.html to see an example of how to use CreateUserAssessment API.
// A default retry strategy applies to this operation CreateUserAssessment()
func (client DataSafeClient) CreateUserAssessment(ctx context.Context, request CreateUserAssessmentRequest) (response CreateUserAssessmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/CreateUserAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "CreateUserAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeactivateTargetDatabase Deactivates a target database in Data Safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeactivateTargetDatabase.go.html to see an example of how to use DeactivateTargetDatabase API.
// A default retry strategy applies to this operation DeactivateTargetDatabase()
func (client DataSafeClient) DeactivateTargetDatabase(ctx context.Context, request DeactivateTargetDatabaseRequest) (response DeactivateTargetDatabaseResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/DeactivateTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "DeactivateTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAuditArchiveRetrieval To unload retrieved archive data, call the operation ListAuditArchiveRetrieval first.
// This will return the auditArchiveRetrievalId. Then call this operation with auditArchiveRetrievalId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteAuditArchiveRetrieval.go.html to see an example of how to use DeleteAuditArchiveRetrieval API.
// A default retry strategy applies to this operation DeleteAuditArchiveRetrieval()
func (client DataSafeClient) DeleteAuditArchiveRetrieval(ctx context.Context, request DeleteAuditArchiveRetrievalRequest) (response DeleteAuditArchiveRetrievalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditArchiveRetrieval/DeleteAuditArchiveRetrieval"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteAuditArchiveRetrieval", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAuditTrail Deletes the specified audit trail.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteAuditTrail.go.html to see an example of how to use DeleteAuditTrail API.
// A default retry strategy applies to this operation DeleteAuditTrail()
func (client DataSafeClient) DeleteAuditTrail(ctx context.Context, request DeleteAuditTrailRequest) (response DeleteAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/DeleteAuditTrail"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteAuditTrail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataSafePrivateEndpoint Deletes the specified Data Safe private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDataSafePrivateEndpoint.go.html to see an example of how to use DeleteDataSafePrivateEndpoint API.
// A default retry strategy applies to this operation DeleteDataSafePrivateEndpoint()
func (client DataSafeClient) DeleteDataSafePrivateEndpoint(ctx context.Context, request DeleteDataSafePrivateEndpointRequest) (response DeleteDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafePrivateEndpoint/DeleteDataSafePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteDataSafePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDiscoveryJob Deletes the specified discovery job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDiscoveryJob.go.html to see an example of how to use DeleteDiscoveryJob API.
// A default retry strategy applies to this operation DeleteDiscoveryJob()
func (client DataSafeClient) DeleteDiscoveryJob(ctx context.Context, request DeleteDiscoveryJobRequest) (response DeleteDiscoveryJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJob/DeleteDiscoveryJob"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteDiscoveryJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDiscoveryJobResult Deletes the specified discovery result.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteDiscoveryJobResult.go.html to see an example of how to use DeleteDiscoveryJobResult API.
// A default retry strategy applies to this operation DeleteDiscoveryJobResult()
func (client DataSafeClient) DeleteDiscoveryJobResult(ctx context.Context, request DeleteDiscoveryJobResultRequest) (response DeleteDiscoveryJobResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJobResult/DeleteDiscoveryJobResult"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteDiscoveryJobResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLibraryMaskingFormat Deletes the specified library masking format.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteLibraryMaskingFormat.go.html to see an example of how to use DeleteLibraryMaskingFormat API.
// A default retry strategy applies to this operation DeleteLibraryMaskingFormat()
func (client DataSafeClient) DeleteLibraryMaskingFormat(ctx context.Context, request DeleteLibraryMaskingFormatRequest) (response DeleteLibraryMaskingFormatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/LibraryMaskingFormat/DeleteLibraryMaskingFormat"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteLibraryMaskingFormat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMaskingColumn Deletes the specified masking column.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteMaskingColumn.go.html to see an example of how to use DeleteMaskingColumn API.
// A default retry strategy applies to this operation DeleteMaskingColumn()
func (client DataSafeClient) DeleteMaskingColumn(ctx context.Context, request DeleteMaskingColumnRequest) (response DeleteMaskingColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/DeleteMaskingColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteMaskingColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMaskingPolicy Deletes the specified masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteMaskingPolicy.go.html to see an example of how to use DeleteMaskingPolicy API.
// A default retry strategy applies to this operation DeleteMaskingPolicy()
func (client DataSafeClient) DeleteMaskingPolicy(ctx context.Context, request DeleteMaskingPolicyRequest) (response DeleteMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/DeleteMaskingPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteMaskingPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMaskingPolicyHealthReport Deletes the specified masking policy health report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteMaskingPolicyHealthReport.go.html to see an example of how to use DeleteMaskingPolicyHealthReport API.
// A default retry strategy applies to this operation DeleteMaskingPolicyHealthReport()
func (client DataSafeClient) DeleteMaskingPolicyHealthReport(ctx context.Context, request DeleteMaskingPolicyHealthReportRequest) (response DeleteMaskingPolicyHealthReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMaskingPolicyHealthReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMaskingPolicyHealthReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMaskingPolicyHealthReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMaskingPolicyHealthReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMaskingPolicyHealthReportResponse")
	}
	return
}

// deleteMaskingPolicyHealthReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteMaskingPolicyHealthReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/maskingPolicyHealthReports/{maskingPolicyHealthReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMaskingPolicyHealthReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicyHealthReport/DeleteMaskingPolicyHealthReport"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteMaskingPolicyHealthReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOnPremConnector Deletes the specified on-premises connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteOnPremConnector.go.html to see an example of how to use DeleteOnPremConnector API.
// A default retry strategy applies to this operation DeleteOnPremConnector()
func (client DataSafeClient) DeleteOnPremConnector(ctx context.Context, request DeleteOnPremConnectorRequest) (response DeleteOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/DeleteOnPremConnector"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteOnPremConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePeerTargetDatabase Removes the specified peer target database from Data Safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeletePeerTargetDatabase.go.html to see an example of how to use DeletePeerTargetDatabase API.
// A default retry strategy applies to this operation DeletePeerTargetDatabase()
func (client DataSafeClient) DeletePeerTargetDatabase(ctx context.Context, request DeletePeerTargetDatabaseRequest) (response DeletePeerTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePeerTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePeerTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePeerTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePeerTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePeerTargetDatabaseResponse")
	}
	return
}

// deletePeerTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deletePeerTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targetDatabases/{targetDatabaseId}/peerTargetDatabases/{peerTargetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePeerTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/PeerTargetDatabase/DeletePeerTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "DeletePeerTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteReportDefinition Deletes the specified report definition. Only the user created report definition can be deleted. The seeded report definitions cannot be deleted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteReportDefinition.go.html to see an example of how to use DeleteReportDefinition API.
// A default retry strategy applies to this operation DeleteReportDefinition()
func (client DataSafeClient) DeleteReportDefinition(ctx context.Context, request DeleteReportDefinitionRequest) (response DeleteReportDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/DeleteReportDefinition"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteReportDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSdmMaskingPolicyDifference Deletes the specified SDM Masking policy difference.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSdmMaskingPolicyDifference.go.html to see an example of how to use DeleteSdmMaskingPolicyDifference API.
// A default retry strategy applies to this operation DeleteSdmMaskingPolicyDifference()
func (client DataSafeClient) DeleteSdmMaskingPolicyDifference(ctx context.Context, request DeleteSdmMaskingPolicyDifferenceRequest) (response DeleteSdmMaskingPolicyDifferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSdmMaskingPolicyDifference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSdmMaskingPolicyDifferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSdmMaskingPolicyDifferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSdmMaskingPolicyDifferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSdmMaskingPolicyDifferenceResponse")
	}
	return
}

// deleteSdmMaskingPolicyDifference implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSdmMaskingPolicyDifference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSdmMaskingPolicyDifferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/DeleteSdmMaskingPolicyDifference"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSdmMaskingPolicyDifference", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSecurityAssessment.go.html to see an example of how to use DeleteSecurityAssessment API.
// A default retry strategy applies to this operation DeleteSecurityAssessment()
func (client DataSafeClient) DeleteSecurityAssessment(ctx context.Context, request DeleteSecurityAssessmentRequest) (response DeleteSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/DeleteSecurityAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSecurityAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSensitiveColumn Deletes the specified sensitive column.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSensitiveColumn.go.html to see an example of how to use DeleteSensitiveColumn API.
// A default retry strategy applies to this operation DeleteSensitiveColumn()
func (client DataSafeClient) DeleteSensitiveColumn(ctx context.Context, request DeleteSensitiveColumnRequest) (response DeleteSensitiveColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveColumn/DeleteSensitiveColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSensitiveColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSensitiveDataModel Deletes the specified sensitive data model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSensitiveDataModel.go.html to see an example of how to use DeleteSensitiveDataModel API.
// A default retry strategy applies to this operation DeleteSensitiveDataModel()
func (client DataSafeClient) DeleteSensitiveDataModel(ctx context.Context, request DeleteSensitiveDataModelRequest) (response DeleteSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/DeleteSensitiveDataModel"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSensitiveDataModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSensitiveType Deletes the specified sensitive type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSensitiveType.go.html to see an example of how to use DeleteSensitiveType API.
// A default retry strategy applies to this operation DeleteSensitiveType()
func (client DataSafeClient) DeleteSensitiveType(ctx context.Context, request DeleteSensitiveTypeRequest) (response DeleteSensitiveTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveType/DeleteSensitiveType"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSensitiveType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSqlCollection Deletes the specified SQL collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSqlCollection.go.html to see an example of how to use DeleteSqlCollection API.
// A default retry strategy applies to this operation DeleteSqlCollection()
func (client DataSafeClient) DeleteSqlCollection(ctx context.Context, request DeleteSqlCollectionRequest) (response DeleteSqlCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSqlCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSqlCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSqlCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSqlCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSqlCollectionResponse")
	}
	return
}

// deleteSqlCollection implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSqlCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sqlCollections/{sqlCollectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSqlCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/DeleteSqlCollection"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSqlCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSqlFirewallPolicy Deletes the SQL Firewall policy resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteSqlFirewallPolicy.go.html to see an example of how to use DeleteSqlFirewallPolicy API.
// A default retry strategy applies to this operation DeleteSqlFirewallPolicy()
func (client DataSafeClient) DeleteSqlFirewallPolicy(ctx context.Context, request DeleteSqlFirewallPolicyRequest) (response DeleteSqlFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSqlFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSqlFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSqlFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSqlFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSqlFirewallPolicyResponse")
	}
	return
}

// deleteSqlFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) deleteSqlFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sqlFirewallPolicies/{sqlFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSqlFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallPolicy/DeleteSqlFirewallPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteSqlFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetAlertPolicyAssociation Deletes the specified target-alert policy Association.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteTargetAlertPolicyAssociation.go.html to see an example of how to use DeleteTargetAlertPolicyAssociation API.
// A default retry strategy applies to this operation DeleteTargetAlertPolicyAssociation()
func (client DataSafeClient) DeleteTargetAlertPolicyAssociation(ctx context.Context, request DeleteTargetAlertPolicyAssociationRequest) (response DeleteTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociation/DeleteTargetAlertPolicyAssociation"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteTargetAlertPolicyAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetDatabase Deregisters the specified database from Data Safe and removes the target database from the Data Safe Console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteTargetDatabase.go.html to see an example of how to use DeleteTargetDatabase API.
// A default retry strategy applies to this operation DeleteTargetDatabase()
func (client DataSafeClient) DeleteTargetDatabase(ctx context.Context, request DeleteTargetDatabaseRequest) (response DeleteTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/DeleteTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteTargetDatabase", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DeleteUserAssessment.go.html to see an example of how to use DeleteUserAssessment API.
// A default retry strategy applies to this operation DeleteUserAssessment()
func (client DataSafeClient) DeleteUserAssessment(ctx context.Context, request DeleteUserAssessmentRequest) (response DeleteUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/DeleteUserAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "DeleteUserAssessment", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DiscoverAuditTrails.go.html to see an example of how to use DiscoverAuditTrails API.
// A default retry strategy applies to this operation DiscoverAuditTrails()
func (client DataSafeClient) DiscoverAuditTrails(ctx context.Context, request DiscoverAuditTrailsRequest) (response DiscoverAuditTrailsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/DiscoverAuditTrails"
		err = common.PostProcessServiceError(err, "DataSafe", "DiscoverAuditTrails", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadDiscoveryReport.go.html to see an example of how to use DownloadDiscoveryReport API.
// A default retry strategy applies to this operation DownloadDiscoveryReport()
func (client DataSafeClient) DownloadDiscoveryReport(ctx context.Context, request DownloadDiscoveryReportRequest) (response DownloadDiscoveryReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/DownloadDiscoveryReport"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadDiscoveryReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadMaskingLog Downloads the masking log generated by the last masking operation on a target database using the specified masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadMaskingLog.go.html to see an example of how to use DownloadMaskingLog API.
// A default retry strategy applies to this operation DownloadMaskingLog()
func (client DataSafeClient) DownloadMaskingLog(ctx context.Context, request DownloadMaskingLogRequest) (response DownloadMaskingLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/DownloadMaskingLog"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadMaskingLog", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadMaskingPolicy.go.html to see an example of how to use DownloadMaskingPolicy API.
// A default retry strategy applies to this operation DownloadMaskingPolicy()
func (client DataSafeClient) DownloadMaskingPolicy(ctx context.Context, request DownloadMaskingPolicyRequest) (response DownloadMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/DownloadMaskingPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadMaskingPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadMaskingReport Downloads an already-generated masking report. Note that the GenerateMaskingReportForDownload
// operation is a prerequisite for the DownloadMaskingReport operation. Use GenerateMaskingReportForDownload
// to generate a masking report file and then use DownloadMaskingReport to download the generated file.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadMaskingReport.go.html to see an example of how to use DownloadMaskingReport API.
// A default retry strategy applies to this operation DownloadMaskingReport()
func (client DataSafeClient) DownloadMaskingReport(ctx context.Context, request DownloadMaskingReportRequest) (response DownloadMaskingReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/DownloadMaskingReport"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadMaskingReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadPrivilegeScript Downloads the privilege script to grant/revoke required roles from the Data Safe account on the target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadPrivilegeScript.go.html to see an example of how to use DownloadPrivilegeScript API.
// A default retry strategy applies to this operation DownloadPrivilegeScript()
func (client DataSafeClient) DownloadPrivilegeScript(ctx context.Context, request DownloadPrivilegeScriptRequest) (response DownloadPrivilegeScriptResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/DownloadPrivilegeScript"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadPrivilegeScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadSecurityAssessmentReport Downloads the report of the specified security assessment. To download the security assessment report, it needs to be generated first.
// Please use GenerateSecurityAssessmentReport to generate a downloadable report in the preferred format (PDF, XLS).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadSecurityAssessmentReport.go.html to see an example of how to use DownloadSecurityAssessmentReport API.
// A default retry strategy applies to this operation DownloadSecurityAssessmentReport()
func (client DataSafeClient) DownloadSecurityAssessmentReport(ctx context.Context, request DownloadSecurityAssessmentReportRequest) (response DownloadSecurityAssessmentReportResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/DownloadSecurityAssessmentReport"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadSecurityAssessmentReport", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadSensitiveDataModel.go.html to see an example of how to use DownloadSensitiveDataModel API.
// A default retry strategy applies to this operation DownloadSensitiveDataModel()
func (client DataSafeClient) DownloadSensitiveDataModel(ctx context.Context, request DownloadSensitiveDataModelRequest) (response DownloadSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/DownloadSensitiveDataModel"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadSensitiveDataModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadUserAssessmentReport Downloads the report of the specified user assessment. To download the user assessment report, it needs to be generated first.
// Please use GenerateUserAssessmentReport to generate a downloadable report in the preferred format (PDF, XLS).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/DownloadUserAssessmentReport.go.html to see an example of how to use DownloadUserAssessmentReport API.
// A default retry strategy applies to this operation DownloadUserAssessmentReport()
func (client DataSafeClient) DownloadUserAssessmentReport(ctx context.Context, request DownloadUserAssessmentReportRequest) (response DownloadUserAssessmentReportResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/DownloadUserAssessmentReport"
		err = common.PostProcessServiceError(err, "DataSafe", "DownloadUserAssessmentReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDataSafeConfiguration Enables Data Safe in the tenancy and region.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/EnableDataSafeConfiguration.go.html to see an example of how to use EnableDataSafeConfiguration API.
// A default retry strategy applies to this operation EnableDataSafeConfiguration()
func (client DataSafeClient) EnableDataSafeConfiguration(ctx context.Context, request EnableDataSafeConfigurationRequest) (response EnableDataSafeConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafeConfiguration/EnableDataSafeConfiguration"
		err = common.PostProcessServiceError(err, "DataSafe", "EnableDataSafeConfiguration", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateDiscoveryReportForDownload.go.html to see an example of how to use GenerateDiscoveryReportForDownload API.
// A default retry strategy applies to this operation GenerateDiscoveryReportForDownload()
func (client DataSafeClient) GenerateDiscoveryReportForDownload(ctx context.Context, request GenerateDiscoveryReportForDownloadRequest) (response GenerateDiscoveryReportForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/GenerateDiscoveryReportForDownload"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateDiscoveryReportForDownload", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateHealthReport Performs health check on the masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateHealthReport.go.html to see an example of how to use GenerateHealthReport API.
// A default retry strategy applies to this operation GenerateHealthReport()
func (client DataSafeClient) GenerateHealthReport(ctx context.Context, request GenerateHealthReportRequest) (response GenerateHealthReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateHealthReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateHealthReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateHealthReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateHealthReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateHealthReportResponse")
	}
	return
}

// generateHealthReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateHealthReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maskingPolicies/{maskingPolicyId}/actions/generateHealthReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateHealthReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicyHealthReport/GenerateHealthReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateHealthReport", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateMaskingPolicyForDownload.go.html to see an example of how to use GenerateMaskingPolicyForDownload API.
// A default retry strategy applies to this operation GenerateMaskingPolicyForDownload()
func (client DataSafeClient) GenerateMaskingPolicyForDownload(ctx context.Context, request GenerateMaskingPolicyForDownloadRequest) (response GenerateMaskingPolicyForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/GenerateMaskingPolicyForDownload"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateMaskingPolicyForDownload", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateMaskingReportForDownload.go.html to see an example of how to use GenerateMaskingReportForDownload API.
// A default retry strategy applies to this operation GenerateMaskingReportForDownload()
func (client DataSafeClient) GenerateMaskingReportForDownload(ctx context.Context, request GenerateMaskingReportForDownloadRequest) (response GenerateMaskingReportForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/GenerateMaskingReportForDownload"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateMaskingReportForDownload", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateOnPremConnectorConfiguration Creates and downloads the configuration of the specified on-premises connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateOnPremConnectorConfiguration.go.html to see an example of how to use GenerateOnPremConnectorConfiguration API.
// A default retry strategy applies to this operation GenerateOnPremConnectorConfiguration()
func (client DataSafeClient) GenerateOnPremConnectorConfiguration(ctx context.Context, request GenerateOnPremConnectorConfigurationRequest) (response GenerateOnPremConnectorConfigurationResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/GenerateOnPremConnectorConfiguration"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateOnPremConnectorConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateReport Generates a .xls or .pdf report based on parameters and report definition.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateReport.go.html to see an example of how to use GenerateReport API.
// A default retry strategy applies to this operation GenerateReport()
func (client DataSafeClient) GenerateReport(ctx context.Context, request GenerateReportRequest) (response GenerateReportResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/GenerateReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateSecurityAssessmentReport Generates the report of the specified security assessment. You can get the report in PDF or XLS format.
// After generating the report, use DownloadSecurityAssessmentReport to download it in the preferred format.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateSecurityAssessmentReport.go.html to see an example of how to use GenerateSecurityAssessmentReport API.
// A default retry strategy applies to this operation GenerateSecurityAssessmentReport()
func (client DataSafeClient) GenerateSecurityAssessmentReport(ctx context.Context, request GenerateSecurityAssessmentReportRequest) (response GenerateSecurityAssessmentReportResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/GenerateSecurityAssessmentReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateSecurityAssessmentReport", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateSensitiveDataModelForDownload.go.html to see an example of how to use GenerateSensitiveDataModelForDownload API.
// A default retry strategy applies to this operation GenerateSensitiveDataModelForDownload()
func (client DataSafeClient) GenerateSensitiveDataModelForDownload(ctx context.Context, request GenerateSensitiveDataModelForDownloadRequest) (response GenerateSensitiveDataModelForDownloadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/GenerateSensitiveDataModelForDownload"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateSensitiveDataModelForDownload", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateSqlFirewallPolicy Generates or appends to the SQL Firewall policy using the specified SQL collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateSqlFirewallPolicy.go.html to see an example of how to use GenerateSqlFirewallPolicy API.
// A default retry strategy applies to this operation GenerateSqlFirewallPolicy()
func (client DataSafeClient) GenerateSqlFirewallPolicy(ctx context.Context, request GenerateSqlFirewallPolicyRequest) (response GenerateSqlFirewallPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateSqlFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateSqlFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateSqlFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateSqlFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateSqlFirewallPolicyResponse")
	}
	return
}

// generateSqlFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) generateSqlFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections/{sqlCollectionId}/actions/generateSqlFirewallPolicy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateSqlFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/GenerateSqlFirewallPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateSqlFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateUserAssessmentReport Generates the report of the specified user assessment. The report is available in PDF or XLS format.
// After generating the report, use DownloadUserAssessmentReport to download it in the preferred format.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GenerateUserAssessmentReport.go.html to see an example of how to use GenerateUserAssessmentReport API.
// A default retry strategy applies to this operation GenerateUserAssessmentReport()
func (client DataSafeClient) GenerateUserAssessmentReport(ctx context.Context, request GenerateUserAssessmentReportRequest) (response GenerateUserAssessmentReportResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/GenerateUserAssessmentReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GenerateUserAssessmentReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlert Gets the details of the specified alerts.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAlert.go.html to see an example of how to use GetAlert API.
// A default retry strategy applies to this operation GetAlert()
func (client DataSafeClient) GetAlert(ctx context.Context, request GetAlertRequest) (response GetAlertResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Alert/GetAlert"
		err = common.PostProcessServiceError(err, "DataSafe", "GetAlert", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlertPolicy Gets the details of alert policy by its ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAlertPolicy.go.html to see an example of how to use GetAlertPolicy API.
// A default retry strategy applies to this operation GetAlertPolicy()
func (client DataSafeClient) GetAlertPolicy(ctx context.Context, request GetAlertPolicyRequest) (response GetAlertPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AlertPolicy/GetAlertPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "GetAlertPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuditArchiveRetrieval Gets the details of the specified archive retreival.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditArchiveRetrieval.go.html to see an example of how to use GetAuditArchiveRetrieval API.
// A default retry strategy applies to this operation GetAuditArchiveRetrieval()
func (client DataSafeClient) GetAuditArchiveRetrieval(ctx context.Context, request GetAuditArchiveRetrievalRequest) (response GetAuditArchiveRetrievalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditArchiveRetrieval/GetAuditArchiveRetrieval"
		err = common.PostProcessServiceError(err, "DataSafe", "GetAuditArchiveRetrieval", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuditPolicy Gets a audit policy by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditPolicy.go.html to see an example of how to use GetAuditPolicy API.
// A default retry strategy applies to this operation GetAuditPolicy()
func (client DataSafeClient) GetAuditPolicy(ctx context.Context, request GetAuditPolicyRequest) (response GetAuditPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicy/GetAuditPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "GetAuditPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuditProfile Gets the details of audit profile resource and associated audit trails of the audit profile.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditProfile.go.html to see an example of how to use GetAuditProfile API.
// A default retry strategy applies to this operation GetAuditProfile()
func (client DataSafeClient) GetAuditProfile(ctx context.Context, request GetAuditProfileRequest) (response GetAuditProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/GetAuditProfile"
		err = common.PostProcessServiceError(err, "DataSafe", "GetAuditProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuditTrail Gets the details of audit trail.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetAuditTrail.go.html to see an example of how to use GetAuditTrail API.
// A default retry strategy applies to this operation GetAuditTrail()
func (client DataSafeClient) GetAuditTrail(ctx context.Context, request GetAuditTrailRequest) (response GetAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/GetAuditTrail"
		err = common.PostProcessServiceError(err, "DataSafe", "GetAuditTrail", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetCompatibleFormatsForDataTypes.go.html to see an example of how to use GetCompatibleFormatsForDataTypes API.
// A default retry strategy applies to this operation GetCompatibleFormatsForDataTypes()
func (client DataSafeClient) GetCompatibleFormatsForDataTypes(ctx context.Context, request GetCompatibleFormatsForDataTypesRequest) (response GetCompatibleFormatsForDataTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/GetCompatibleFormatsForDataTypes"
		err = common.PostProcessServiceError(err, "DataSafe", "GetCompatibleFormatsForDataTypes", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetCompatibleFormatsForSensitiveTypes.go.html to see an example of how to use GetCompatibleFormatsForSensitiveTypes API.
// A default retry strategy applies to this operation GetCompatibleFormatsForSensitiveTypes()
func (client DataSafeClient) GetCompatibleFormatsForSensitiveTypes(ctx context.Context, request GetCompatibleFormatsForSensitiveTypesRequest) (response GetCompatibleFormatsForSensitiveTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/GetCompatibleFormatsForSensitiveTypes"
		err = common.PostProcessServiceError(err, "DataSafe", "GetCompatibleFormatsForSensitiveTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataSafeConfiguration Gets the details of the Data Safe configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDataSafeConfiguration.go.html to see an example of how to use GetDataSafeConfiguration API.
// A default retry strategy applies to this operation GetDataSafeConfiguration()
func (client DataSafeClient) GetDataSafeConfiguration(ctx context.Context, request GetDataSafeConfigurationRequest) (response GetDataSafeConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafeConfiguration/GetDataSafeConfiguration"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDataSafeConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataSafePrivateEndpoint Gets the details of the specified Data Safe private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDataSafePrivateEndpoint.go.html to see an example of how to use GetDataSafePrivateEndpoint API.
// A default retry strategy applies to this operation GetDataSafePrivateEndpoint()
func (client DataSafeClient) GetDataSafePrivateEndpoint(ctx context.Context, request GetDataSafePrivateEndpointRequest) (response GetDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafePrivateEndpoint/GetDataSafePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDataSafePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseSecurityConfig Gets a database security configuration by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDatabaseSecurityConfig.go.html to see an example of how to use GetDatabaseSecurityConfig API.
// A default retry strategy applies to this operation GetDatabaseSecurityConfig()
func (client DataSafeClient) GetDatabaseSecurityConfig(ctx context.Context, request GetDatabaseSecurityConfigRequest) (response GetDatabaseSecurityConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseSecurityConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseSecurityConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseSecurityConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseSecurityConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseSecurityConfigResponse")
	}
	return
}

// getDatabaseSecurityConfig implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDatabaseSecurityConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseSecurityConfigs/{databaseSecurityConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseSecurityConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseSecurityConfig/GetDatabaseSecurityConfig"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDatabaseSecurityConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseTableAccessEntry Gets a database table access entry object by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDatabaseTableAccessEntry.go.html to see an example of how to use GetDatabaseTableAccessEntry API.
// A default retry strategy applies to this operation GetDatabaseTableAccessEntry()
func (client DataSafeClient) GetDatabaseTableAccessEntry(ctx context.Context, request GetDatabaseTableAccessEntryRequest) (response GetDatabaseTableAccessEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseTableAccessEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseTableAccessEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseTableAccessEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseTableAccessEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseTableAccessEntryResponse")
	}
	return
}

// getDatabaseTableAccessEntry implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDatabaseTableAccessEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports/{securityPolicyReportId}/databaseTableAccessEntries/{databaseTableAccessEntryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseTableAccessEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseTableAccessEntry/GetDatabaseTableAccessEntry"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDatabaseTableAccessEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseViewAccessEntry Gets a database view access object by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDatabaseViewAccessEntry.go.html to see an example of how to use GetDatabaseViewAccessEntry API.
// A default retry strategy applies to this operation GetDatabaseViewAccessEntry()
func (client DataSafeClient) GetDatabaseViewAccessEntry(ctx context.Context, request GetDatabaseViewAccessEntryRequest) (response GetDatabaseViewAccessEntryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseViewAccessEntry, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseViewAccessEntryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseViewAccessEntryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseViewAccessEntryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseViewAccessEntryResponse")
	}
	return
}

// getDatabaseViewAccessEntry implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDatabaseViewAccessEntry(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports/{securityPolicyReportId}/databaseViewAccessEntries/{databaseViewAccessEntryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseViewAccessEntryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseViewAccessEntry/GetDatabaseViewAccessEntry"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDatabaseViewAccessEntry", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDifferenceColumn Gets the details of the specified SDM Masking policy difference column.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDifferenceColumn.go.html to see an example of how to use GetDifferenceColumn API.
// A default retry strategy applies to this operation GetDifferenceColumn()
func (client DataSafeClient) GetDifferenceColumn(ctx context.Context, request GetDifferenceColumnRequest) (response GetDifferenceColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDifferenceColumn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDifferenceColumnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDifferenceColumnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDifferenceColumnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDifferenceColumnResponse")
	}
	return
}

// getDifferenceColumn implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getDifferenceColumn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}/differenceColumns/{differenceColumnKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDifferenceColumnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DifferenceColumn/GetDifferenceColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDifferenceColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDiscoveryJob Gets the details of the specified discovery job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDiscoveryJob.go.html to see an example of how to use GetDiscoveryJob API.
// A default retry strategy applies to this operation GetDiscoveryJob()
func (client DataSafeClient) GetDiscoveryJob(ctx context.Context, request GetDiscoveryJobRequest) (response GetDiscoveryJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJob/GetDiscoveryJob"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDiscoveryJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDiscoveryJobResult Gets the details of the specified discovery result.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetDiscoveryJobResult.go.html to see an example of how to use GetDiscoveryJobResult API.
// A default retry strategy applies to this operation GetDiscoveryJobResult()
func (client DataSafeClient) GetDiscoveryJobResult(ctx context.Context, request GetDiscoveryJobResultRequest) (response GetDiscoveryJobResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJobResult/GetDiscoveryJobResult"
		err = common.PostProcessServiceError(err, "DataSafe", "GetDiscoveryJobResult", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLibraryMaskingFormat Gets the details of the specified library masking format.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetLibraryMaskingFormat.go.html to see an example of how to use GetLibraryMaskingFormat API.
// A default retry strategy applies to this operation GetLibraryMaskingFormat()
func (client DataSafeClient) GetLibraryMaskingFormat(ctx context.Context, request GetLibraryMaskingFormatRequest) (response GetLibraryMaskingFormatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/LibraryMaskingFormat/GetLibraryMaskingFormat"
		err = common.PostProcessServiceError(err, "DataSafe", "GetLibraryMaskingFormat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaskingColumn Gets the details of the specified masking column.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingColumn.go.html to see an example of how to use GetMaskingColumn API.
// A default retry strategy applies to this operation GetMaskingColumn()
func (client DataSafeClient) GetMaskingColumn(ctx context.Context, request GetMaskingColumnRequest) (response GetMaskingColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/GetMaskingColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "GetMaskingColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaskingPolicy Gets the details of the specified masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingPolicy.go.html to see an example of how to use GetMaskingPolicy API.
// A default retry strategy applies to this operation GetMaskingPolicy()
func (client DataSafeClient) GetMaskingPolicy(ctx context.Context, request GetMaskingPolicyRequest) (response GetMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/GetMaskingPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "GetMaskingPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaskingPolicyHealthReport Gets the details of the specified masking policy health report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingPolicyHealthReport.go.html to see an example of how to use GetMaskingPolicyHealthReport API.
// A default retry strategy applies to this operation GetMaskingPolicyHealthReport()
func (client DataSafeClient) GetMaskingPolicyHealthReport(ctx context.Context, request GetMaskingPolicyHealthReportRequest) (response GetMaskingPolicyHealthReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaskingPolicyHealthReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaskingPolicyHealthReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaskingPolicyHealthReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaskingPolicyHealthReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaskingPolicyHealthReportResponse")
	}
	return
}

// getMaskingPolicyHealthReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getMaskingPolicyHealthReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicyHealthReports/{maskingPolicyHealthReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaskingPolicyHealthReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicyHealthReport/GetMaskingPolicyHealthReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GetMaskingPolicyHealthReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaskingReport Gets the details of the specified masking report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetMaskingReport.go.html to see an example of how to use GetMaskingReport API.
// A default retry strategy applies to this operation GetMaskingReport()
func (client DataSafeClient) GetMaskingReport(ctx context.Context, request GetMaskingReportRequest) (response GetMaskingReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingReport/GetMaskingReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GetMaskingReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOnPremConnector Gets the details of the specified on-premises connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetOnPremConnector.go.html to see an example of how to use GetOnPremConnector API.
// A default retry strategy applies to this operation GetOnPremConnector()
func (client DataSafeClient) GetOnPremConnector(ctx context.Context, request GetOnPremConnectorRequest) (response GetOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/GetOnPremConnector"
		err = common.PostProcessServiceError(err, "DataSafe", "GetOnPremConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPeerTargetDatabase Returns the details of the specified Data Safe peer target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetPeerTargetDatabase.go.html to see an example of how to use GetPeerTargetDatabase API.
// A default retry strategy applies to this operation GetPeerTargetDatabase()
func (client DataSafeClient) GetPeerTargetDatabase(ctx context.Context, request GetPeerTargetDatabaseRequest) (response GetPeerTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPeerTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPeerTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPeerTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPeerTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPeerTargetDatabaseResponse")
	}
	return
}

// getPeerTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getPeerTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}/peerTargetDatabases/{peerTargetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPeerTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/PeerTargetDatabase/GetPeerTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "GetPeerTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProfile Lists the details of given profile available on the target.
// The GetProfile operation returns only the profiles in the specified 'userAssessmentId'.
// This does not include any subcompartments of the current compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetProfile.go.html to see an example of how to use GetProfile API.
// A default retry strategy applies to this operation GetProfile()
func (client DataSafeClient) GetProfile(ctx context.Context, request GetProfileRequest) (response GetProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProfileResponse")
	}
	return
}

// getProfile implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/profiles/{profileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/GetProfile"
		err = common.PostProcessServiceError(err, "DataSafe", "GetProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReport Gets a report by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetReport.go.html to see an example of how to use GetReport API.
// A default retry strategy applies to this operation GetReport()
func (client DataSafeClient) GetReport(ctx context.Context, request GetReportRequest) (response GetReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Report/GetReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GetReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReportContent Downloads the specified report in the form of .xls or .pdf.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetReportContent.go.html to see an example of how to use GetReportContent API.
// A default retry strategy applies to this operation GetReportContent()
func (client DataSafeClient) GetReportContent(ctx context.Context, request GetReportContentRequest) (response GetReportContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Report/GetReportContent"
		err = common.PostProcessServiceError(err, "DataSafe", "GetReportContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReportDefinition Gets the details of report definition specified by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetReportDefinition.go.html to see an example of how to use GetReportDefinition API.
// A default retry strategy applies to this operation GetReportDefinition()
func (client DataSafeClient) GetReportDefinition(ctx context.Context, request GetReportDefinitionRequest) (response GetReportDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/GetReportDefinition"
		err = common.PostProcessServiceError(err, "DataSafe", "GetReportDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSdmMaskingPolicyDifference Gets the details of the specified SDM Masking policy difference.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSdmMaskingPolicyDifference.go.html to see an example of how to use GetSdmMaskingPolicyDifference API.
// A default retry strategy applies to this operation GetSdmMaskingPolicyDifference()
func (client DataSafeClient) GetSdmMaskingPolicyDifference(ctx context.Context, request GetSdmMaskingPolicyDifferenceRequest) (response GetSdmMaskingPolicyDifferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSdmMaskingPolicyDifference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSdmMaskingPolicyDifferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSdmMaskingPolicyDifferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSdmMaskingPolicyDifferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSdmMaskingPolicyDifferenceResponse")
	}
	return
}

// getSdmMaskingPolicyDifference implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSdmMaskingPolicyDifference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSdmMaskingPolicyDifferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/GetSdmMaskingPolicyDifference"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSdmMaskingPolicyDifference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAssessment Gets the details of the specified security assessment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessment.go.html to see an example of how to use GetSecurityAssessment API.
// A default retry strategy applies to this operation GetSecurityAssessment()
func (client DataSafeClient) GetSecurityAssessment(ctx context.Context, request GetSecurityAssessmentRequest) (response GetSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/GetSecurityAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSecurityAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAssessmentComparison Gets the details of the comparison report for the security assessments submitted for comparison.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessmentComparison.go.html to see an example of how to use GetSecurityAssessmentComparison API.
// A default retry strategy applies to this operation GetSecurityAssessmentComparison()
func (client DataSafeClient) GetSecurityAssessmentComparison(ctx context.Context, request GetSecurityAssessmentComparisonRequest) (response GetSecurityAssessmentComparisonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/GetSecurityAssessmentComparison"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSecurityAssessmentComparison", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityPolicy Gets a security policy by the specified OCID of the security policy resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityPolicy.go.html to see an example of how to use GetSecurityPolicy API.
// A default retry strategy applies to this operation GetSecurityPolicy()
func (client DataSafeClient) GetSecurityPolicy(ctx context.Context, request GetSecurityPolicyRequest) (response GetSecurityPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityPolicyResponse")
	}
	return
}

// getSecurityPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicies/{securityPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicy/GetSecurityPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSecurityPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityPolicyDeployment Gets a security policy deployment by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityPolicyDeployment.go.html to see an example of how to use GetSecurityPolicyDeployment API.
// A default retry strategy applies to this operation GetSecurityPolicyDeployment()
func (client DataSafeClient) GetSecurityPolicyDeployment(ctx context.Context, request GetSecurityPolicyDeploymentRequest) (response GetSecurityPolicyDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityPolicyDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityPolicyDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityPolicyDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityPolicyDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityPolicyDeploymentResponse")
	}
	return
}

// getSecurityPolicyDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityPolicyDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyDeployments/{securityPolicyDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityPolicyDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyDeployment/GetSecurityPolicyDeployment"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSecurityPolicyDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityPolicyEntryState Gets a security policy entity states by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityPolicyEntryState.go.html to see an example of how to use GetSecurityPolicyEntryState API.
// A default retry strategy applies to this operation GetSecurityPolicyEntryState()
func (client DataSafeClient) GetSecurityPolicyEntryState(ctx context.Context, request GetSecurityPolicyEntryStateRequest) (response GetSecurityPolicyEntryStateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityPolicyEntryState, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityPolicyEntryStateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityPolicyEntryStateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityPolicyEntryStateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityPolicyEntryStateResponse")
	}
	return
}

// getSecurityPolicyEntryState implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityPolicyEntryState(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyDeployments/{securityPolicyDeploymentId}/securityPolicyEntryStates/{securityPolicyEntryStateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityPolicyEntryStateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyEntryState/GetSecurityPolicyEntryState"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSecurityPolicyEntryState", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityPolicyReport Gets a security policy report by the specified OCID of the security policy report resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityPolicyReport.go.html to see an example of how to use GetSecurityPolicyReport API.
// A default retry strategy applies to this operation GetSecurityPolicyReport()
func (client DataSafeClient) GetSecurityPolicyReport(ctx context.Context, request GetSecurityPolicyReportRequest) (response GetSecurityPolicyReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityPolicyReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityPolicyReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityPolicyReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityPolicyReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityPolicyReportResponse")
	}
	return
}

// getSecurityPolicyReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSecurityPolicyReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports/{securityPolicyReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityPolicyReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyReport/GetSecurityPolicyReport"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSecurityPolicyReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSensitiveColumn Gets the details of the specified sensitive column.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSensitiveColumn.go.html to see an example of how to use GetSensitiveColumn API.
// A default retry strategy applies to this operation GetSensitiveColumn()
func (client DataSafeClient) GetSensitiveColumn(ctx context.Context, request GetSensitiveColumnRequest) (response GetSensitiveColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveColumn/GetSensitiveColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSensitiveColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSensitiveDataModel Gets the details of the specified sensitive data model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSensitiveDataModel.go.html to see an example of how to use GetSensitiveDataModel API.
// A default retry strategy applies to this operation GetSensitiveDataModel()
func (client DataSafeClient) GetSensitiveDataModel(ctx context.Context, request GetSensitiveDataModelRequest) (response GetSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/GetSensitiveDataModel"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSensitiveDataModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSensitiveType Gets the details of the specified sensitive type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSensitiveType.go.html to see an example of how to use GetSensitiveType API.
// A default retry strategy applies to this operation GetSensitiveType()
func (client DataSafeClient) GetSensitiveType(ctx context.Context, request GetSensitiveTypeRequest) (response GetSensitiveTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveType/GetSensitiveType"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSensitiveType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &sensitivetype{})
	return response, err
}

// GetSqlCollection Gets a SQL collection by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSqlCollection.go.html to see an example of how to use GetSqlCollection API.
// A default retry strategy applies to this operation GetSqlCollection()
func (client DataSafeClient) GetSqlCollection(ctx context.Context, request GetSqlCollectionRequest) (response GetSqlCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSqlCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSqlCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSqlCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSqlCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSqlCollectionResponse")
	}
	return
}

// getSqlCollection implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSqlCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlCollections/{sqlCollectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSqlCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/GetSqlCollection"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSqlCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSqlFirewallPolicy Gets a SQL Firewall policy by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSqlFirewallPolicy.go.html to see an example of how to use GetSqlFirewallPolicy API.
// A default retry strategy applies to this operation GetSqlFirewallPolicy()
func (client DataSafeClient) GetSqlFirewallPolicy(ctx context.Context, request GetSqlFirewallPolicyRequest) (response GetSqlFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSqlFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSqlFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSqlFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSqlFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSqlFirewallPolicyResponse")
	}
	return
}

// getSqlFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) getSqlFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallPolicies/{sqlFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSqlFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallPolicy/GetSqlFirewallPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "GetSqlFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetAlertPolicyAssociation Gets the details of target-alert policy association by its ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetTargetAlertPolicyAssociation.go.html to see an example of how to use GetTargetAlertPolicyAssociation API.
// A default retry strategy applies to this operation GetTargetAlertPolicyAssociation()
func (client DataSafeClient) GetTargetAlertPolicyAssociation(ctx context.Context, request GetTargetAlertPolicyAssociationRequest) (response GetTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociation/GetTargetAlertPolicyAssociation"
		err = common.PostProcessServiceError(err, "DataSafe", "GetTargetAlertPolicyAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetDatabase Returns the details of the specified Data Safe target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetTargetDatabase.go.html to see an example of how to use GetTargetDatabase API.
// A default retry strategy applies to this operation GetTargetDatabase()
func (client DataSafeClient) GetTargetDatabase(ctx context.Context, request GetTargetDatabaseRequest) (response GetTargetDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/GetTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "GetTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserAssessment Gets a user assessment by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetUserAssessment.go.html to see an example of how to use GetUserAssessment API.
// A default retry strategy applies to this operation GetUserAssessment()
func (client DataSafeClient) GetUserAssessment(ctx context.Context, request GetUserAssessmentRequest) (response GetUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/GetUserAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "GetUserAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUserAssessmentComparison Gets the details of the comparison report for the user assessments submitted for comparison.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetUserAssessmentComparison.go.html to see an example of how to use GetUserAssessmentComparison API.
// A default retry strategy applies to this operation GetUserAssessmentComparison()
func (client DataSafeClient) GetUserAssessmentComparison(ctx context.Context, request GetUserAssessmentComparisonRequest) (response GetUserAssessmentComparisonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/GetUserAssessmentComparison"
		err = common.PostProcessServiceError(err, "DataSafe", "GetUserAssessmentComparison", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of the specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DataSafeClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DataSafe", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlertAnalytics Returns the aggregation details of the alerts.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertAnalytics.go.html to see an example of how to use ListAlertAnalytics API.
// A default retry strategy applies to this operation ListAlertAnalytics()
func (client DataSafeClient) ListAlertAnalytics(ctx context.Context, request ListAlertAnalyticsRequest) (response ListAlertAnalyticsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AlertSummary/ListAlertAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAlertAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlertPolicies Gets a list of all alert policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertPolicies.go.html to see an example of how to use ListAlertPolicies API.
// A default retry strategy applies to this operation ListAlertPolicies()
func (client DataSafeClient) ListAlertPolicies(ctx context.Context, request ListAlertPoliciesRequest) (response ListAlertPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AlertPolicy/ListAlertPolicies"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAlertPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlertPolicyRules Lists the rules of the specified alert policy. The alert policy is said to be satisfied when all rules in the policy evaulate to true.
// If there are three rules: rule1,rule2 and rule3, the policy is satisfied if rule1 AND rule2 AND rule3 is True.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertPolicyRules.go.html to see an example of how to use ListAlertPolicyRules API.
// A default retry strategy applies to this operation ListAlertPolicyRules()
func (client DataSafeClient) ListAlertPolicyRules(ctx context.Context, request ListAlertPolicyRulesRequest) (response ListAlertPolicyRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AlertPolicy/ListAlertPolicyRules"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAlertPolicyRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlerts Gets a list of all alerts.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlerts.go.html to see an example of how to use ListAlerts API.
// A default retry strategy applies to this operation ListAlerts()
func (client DataSafeClient) ListAlerts(ctx context.Context, request ListAlertsRequest) (response ListAlertsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AlertSummary/ListAlerts"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAlerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuditArchiveRetrievals Returns the list of audit archive retrieval.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditArchiveRetrievals.go.html to see an example of how to use ListAuditArchiveRetrievals API.
// A default retry strategy applies to this operation ListAuditArchiveRetrievals()
func (client DataSafeClient) ListAuditArchiveRetrievals(ctx context.Context, request ListAuditArchiveRetrievalsRequest) (response ListAuditArchiveRetrievalsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditArchiveRetrieval/ListAuditArchiveRetrievals"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditArchiveRetrievals", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuditEventAnalytics By default the ListAuditEventAnalytics operation will return all of the summary columns. To filter for a specific summary column, specify
// it in the `summaryField` query parameter.
// **Example:**
// /ListAuditEventAnalytics?summaryField=targetName&summaryField=userName&summaryField=clientHostname
// &summaryField=dmls&summaryField=privilegeChanges&summaryField=ddls&summaryField=loginFailure&summaryField=loginSuccess
// &summaryField=allRecord&q=(auditEventTime ge "2021-06-13T23:49:14")
// /ListAuditEventAnalytics?timeStarted=2022-08-18T11:02:26.000Z&timeEnded=2022-08-24T11:02:26.000Z
// This will give number of events grouped by periods. Period can be 1 day, 1 week, etc.
// /ListAuditEventAnalytics?summaryField=targetName&groupBy=targetName
// This will give the number of events group by targetName. Only targetName summary column would be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditEventAnalytics.go.html to see an example of how to use ListAuditEventAnalytics API.
// A default retry strategy applies to this operation ListAuditEventAnalytics()
func (client DataSafeClient) ListAuditEventAnalytics(ctx context.Context, request ListAuditEventAnalyticsRequest) (response ListAuditEventAnalyticsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditEventSummary/ListAuditEventAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditEventAnalytics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditEvents.go.html to see an example of how to use ListAuditEvents API.
// A default retry strategy applies to this operation ListAuditEvents()
func (client DataSafeClient) ListAuditEvents(ctx context.Context, request ListAuditEventsRequest) (response ListAuditEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditEventSummary/ListAuditEvents"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditEvents", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditPolicies.go.html to see an example of how to use ListAuditPolicies API.
// A default retry strategy applies to this operation ListAuditPolicies()
func (client DataSafeClient) ListAuditPolicies(ctx context.Context, request ListAuditPoliciesRequest) (response ListAuditPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicyCollection/ListAuditPolicies"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuditPolicyAnalytics Gets a list of aggregated audit policy details on the target databases. A audit policy aggregation
// helps understand the overall state of policies provisioned on targets.
// It is especially useful to create dashboards or to support analytics.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform SummarizedAuditPolicyInfo on the specified
// `compartmentId` and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
// **Example:** ListAuditPolicyAnalytics?groupBy=auditPolicyCategory
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditPolicyAnalytics.go.html to see an example of how to use ListAuditPolicyAnalytics API.
// A default retry strategy applies to this operation ListAuditPolicyAnalytics()
func (client DataSafeClient) ListAuditPolicyAnalytics(ctx context.Context, request ListAuditPolicyAnalyticsRequest) (response ListAuditPolicyAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuditPolicyAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuditPolicyAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuditPolicyAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuditPolicyAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuditPolicyAnalyticsResponse")
	}
	return
}

// listAuditPolicyAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listAuditPolicyAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/auditPolicyAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuditPolicyAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicyAnalyticCollection/ListAuditPolicyAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditPolicyAnalytics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditProfileAnalytics.go.html to see an example of how to use ListAuditProfileAnalytics API.
// A default retry strategy applies to this operation ListAuditProfileAnalytics()
func (client DataSafeClient) ListAuditProfileAnalytics(ctx context.Context, request ListAuditProfileAnalyticsRequest) (response ListAuditProfileAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfileAnalyticCollection/ListAuditProfileAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditProfileAnalytics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditProfiles.go.html to see an example of how to use ListAuditProfiles API.
// A default retry strategy applies to this operation ListAuditProfiles()
func (client DataSafeClient) ListAuditProfiles(ctx context.Context, request ListAuditProfilesRequest) (response ListAuditProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/ListAuditProfiles"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditProfiles", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditTrailAnalytics.go.html to see an example of how to use ListAuditTrailAnalytics API.
// A default retry strategy applies to this operation ListAuditTrailAnalytics()
func (client DataSafeClient) ListAuditTrailAnalytics(ctx context.Context, request ListAuditTrailAnalyticsRequest) (response ListAuditTrailAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrailAnalyticCollection/ListAuditTrailAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditTrailAnalytics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditTrails.go.html to see an example of how to use ListAuditTrails API.
// A default retry strategy applies to this operation ListAuditTrails()
func (client DataSafeClient) ListAuditTrails(ctx context.Context, request ListAuditTrailsRequest) (response ListAuditTrailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/ListAuditTrails"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAuditTrails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAvailableAuditVolumes Retrieves a list of audit trails, and associated audit event volume for each trail up to defined start date.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAvailableAuditVolumes.go.html to see an example of how to use ListAvailableAuditVolumes API.
// A default retry strategy applies to this operation ListAvailableAuditVolumes()
func (client DataSafeClient) ListAvailableAuditVolumes(ctx context.Context, request ListAvailableAuditVolumesRequest) (response ListAvailableAuditVolumesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/ListAvailableAuditVolumes"
		err = common.PostProcessServiceError(err, "DataSafe", "ListAvailableAuditVolumes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCollectedAuditVolumes Gets a list of all collected audit volume data points.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCollectedAuditVolumes.go.html to see an example of how to use ListCollectedAuditVolumes API.
// A default retry strategy applies to this operation ListCollectedAuditVolumes()
func (client DataSafeClient) ListCollectedAuditVolumes(ctx context.Context, request ListCollectedAuditVolumesRequest) (response ListCollectedAuditVolumesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/ListCollectedAuditVolumes"
		err = common.PostProcessServiceError(err, "DataSafe", "ListCollectedAuditVolumes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListColumns Returns a list of column metadata objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListColumns.go.html to see an example of how to use ListColumns API.
// A default retry strategy applies to this operation ListColumns()
func (client DataSafeClient) ListColumns(ctx context.Context, request ListColumnsRequest) (response ListColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/ListColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "ListColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataSafePrivateEndpoints Gets a list of Data Safe private endpoints.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDataSafePrivateEndpoints.go.html to see an example of how to use ListDataSafePrivateEndpoints API.
// A default retry strategy applies to this operation ListDataSafePrivateEndpoints()
func (client DataSafeClient) ListDataSafePrivateEndpoints(ctx context.Context, request ListDataSafePrivateEndpointsRequest) (response ListDataSafePrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafePrivateEndpointSummary/ListDataSafePrivateEndpoints"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDataSafePrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseSecurityConfigs Retrieves a list of all database security configurations in Data Safe.
// The ListDatabaseSecurityConfigs operation returns only the database security configurations in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListDatabaseSecurityConfigs on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDatabaseSecurityConfigs.go.html to see an example of how to use ListDatabaseSecurityConfigs API.
// A default retry strategy applies to this operation ListDatabaseSecurityConfigs()
func (client DataSafeClient) ListDatabaseSecurityConfigs(ctx context.Context, request ListDatabaseSecurityConfigsRequest) (response ListDatabaseSecurityConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseSecurityConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseSecurityConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseSecurityConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseSecurityConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseSecurityConfigsResponse")
	}
	return
}

// listDatabaseSecurityConfigs implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDatabaseSecurityConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseSecurityConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseSecurityConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseSecurityConfigCollection/ListDatabaseSecurityConfigs"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDatabaseSecurityConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseTableAccessEntries Retrieves a list of all database table access entries in Data Safe.
// The ListDatabaseTableAccessEntries operation returns only the database table access reports for the specified security policy report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDatabaseTableAccessEntries.go.html to see an example of how to use ListDatabaseTableAccessEntries API.
// A default retry strategy applies to this operation ListDatabaseTableAccessEntries()
func (client DataSafeClient) ListDatabaseTableAccessEntries(ctx context.Context, request ListDatabaseTableAccessEntriesRequest) (response ListDatabaseTableAccessEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseTableAccessEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseTableAccessEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseTableAccessEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseTableAccessEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseTableAccessEntriesResponse")
	}
	return
}

// listDatabaseTableAccessEntries implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDatabaseTableAccessEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports/{securityPolicyReportId}/databaseTableAccessEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseTableAccessEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseTableAccessEntryCollection/ListDatabaseTableAccessEntries"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDatabaseTableAccessEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseViewAccessEntries Retrieves a list of all database view access entries in Data Safe.
// The ListDatabaseViewAccessEntries operation returns only the database view access objects for the specified security policy report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDatabaseViewAccessEntries.go.html to see an example of how to use ListDatabaseViewAccessEntries API.
// A default retry strategy applies to this operation ListDatabaseViewAccessEntries()
func (client DataSafeClient) ListDatabaseViewAccessEntries(ctx context.Context, request ListDatabaseViewAccessEntriesRequest) (response ListDatabaseViewAccessEntriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseViewAccessEntries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseViewAccessEntriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseViewAccessEntriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseViewAccessEntriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseViewAccessEntriesResponse")
	}
	return
}

// listDatabaseViewAccessEntries implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDatabaseViewAccessEntries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports/{securityPolicyReportId}/databaseViewAccessEntries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseViewAccessEntriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseViewAccessEntryCollection/ListDatabaseViewAccessEntries"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDatabaseViewAccessEntries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDifferenceColumns Gets a list of columns of a SDM masking policy difference resource based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDifferenceColumns.go.html to see an example of how to use ListDifferenceColumns API.
// A default retry strategy applies to this operation ListDifferenceColumns()
func (client DataSafeClient) ListDifferenceColumns(ctx context.Context, request ListDifferenceColumnsRequest) (response ListDifferenceColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDifferenceColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDifferenceColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDifferenceColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDifferenceColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDifferenceColumnsResponse")
	}
	return
}

// listDifferenceColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listDifferenceColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}/differenceColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDifferenceColumnsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/ListDifferenceColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDifferenceColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDiscoveryAnalytics Gets consolidated discovery analytics data based on the specified query parameters.
// If CompartmentIdInSubtreeQueryParam is specified as true, the behaviour
// is equivalent to accessLevel "ACCESSIBLE" by default.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryAnalytics.go.html to see an example of how to use ListDiscoveryAnalytics API.
// A default retry strategy applies to this operation ListDiscoveryAnalytics()
func (client DataSafeClient) ListDiscoveryAnalytics(ctx context.Context, request ListDiscoveryAnalyticsRequest) (response ListDiscoveryAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/ListDiscoveryAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDiscoveryAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDiscoveryJobResults Gets a list of discovery results based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryJobResults.go.html to see an example of how to use ListDiscoveryJobResults API.
// A default retry strategy applies to this operation ListDiscoveryJobResults()
func (client DataSafeClient) ListDiscoveryJobResults(ctx context.Context, request ListDiscoveryJobResultsRequest) (response ListDiscoveryJobResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJob/ListDiscoveryJobResults"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDiscoveryJobResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDiscoveryJobs Gets a list of incremental discovery jobs based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryJobs.go.html to see an example of how to use ListDiscoveryJobs API.
// A default retry strategy applies to this operation ListDiscoveryJobs()
func (client DataSafeClient) ListDiscoveryJobs(ctx context.Context, request ListDiscoveryJobsRequest) (response ListDiscoveryJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJob/ListDiscoveryJobs"
		err = common.PostProcessServiceError(err, "DataSafe", "ListDiscoveryJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFindingAnalytics Gets a list of findings aggregated details in the specified compartment. This provides information about the overall state
// of security assessment findings. You can use groupBy to get the count of findings under a certain risk level and with a certain findingKey,
// and as well as get the list of the targets that match the condition.
// This data is especially useful content for the statistic chart or to support analytics.
// When you perform the ListFindingAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
// parameter accessLevel is set to ACCESSIBLE, then the operation returns statistics from the compartments in which the requestor has INSPECT
// permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
// root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
// compartmentId, then "Not Authorized" is returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindingAnalytics.go.html to see an example of how to use ListFindingAnalytics API.
// A default retry strategy applies to this operation ListFindingAnalytics()
func (client DataSafeClient) ListFindingAnalytics(ctx context.Context, request ListFindingAnalyticsRequest) (response ListFindingAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFindingAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFindingAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFindingAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFindingAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFindingAnalyticsResponse")
	}
	return
}

// listFindingAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listFindingAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/findingAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFindingAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/ListFindingAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListFindingAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFindings List all the findings from all the targets in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindings.go.html to see an example of how to use ListFindings API.
// A default retry strategy applies to this operation ListFindings()
func (client DataSafeClient) ListFindings(ctx context.Context, request ListFindingsRequest) (response ListFindingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/ListFindings"
		err = common.PostProcessServiceError(err, "DataSafe", "ListFindings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFindingsChangeAuditLogs List all changes made by user to risk level of findings of the specified assessment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindingsChangeAuditLogs.go.html to see an example of how to use ListFindingsChangeAuditLogs API.
// A default retry strategy applies to this operation ListFindingsChangeAuditLogs()
func (client DataSafeClient) ListFindingsChangeAuditLogs(ctx context.Context, request ListFindingsChangeAuditLogsRequest) (response ListFindingsChangeAuditLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFindingsChangeAuditLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFindingsChangeAuditLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFindingsChangeAuditLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFindingsChangeAuditLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFindingsChangeAuditLogsResponse")
	}
	return
}

// listFindingsChangeAuditLogs implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listFindingsChangeAuditLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/{securityAssessmentId}/findingsChangeAuditLogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFindingsChangeAuditLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/ListFindingsChangeAuditLogs"
		err = common.PostProcessServiceError(err, "DataSafe", "ListFindingsChangeAuditLogs", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListGrants.go.html to see an example of how to use ListGrants API.
// A default retry strategy applies to this operation ListGrants()
func (client DataSafeClient) ListGrants(ctx context.Context, request ListGrantsRequest) (response ListGrantsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/ListGrants"
		err = common.PostProcessServiceError(err, "DataSafe", "ListGrants", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLibraryMaskingFormats Gets a list of library masking formats based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListLibraryMaskingFormats.go.html to see an example of how to use ListLibraryMaskingFormats API.
// A default retry strategy applies to this operation ListLibraryMaskingFormats()
func (client DataSafeClient) ListLibraryMaskingFormats(ctx context.Context, request ListLibraryMaskingFormatsRequest) (response ListLibraryMaskingFormatsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/LibraryMaskingFormatSummary/ListLibraryMaskingFormats"
		err = common.PostProcessServiceError(err, "DataSafe", "ListLibraryMaskingFormats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskedColumns Gets a list of masked columns present in the specified masking report and based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskedColumns.go.html to see an example of how to use ListMaskedColumns API.
// A default retry strategy applies to this operation ListMaskedColumns()
func (client DataSafeClient) ListMaskedColumns(ctx context.Context, request ListMaskedColumnsRequest) (response ListMaskedColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskedColumnSummary/ListMaskedColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskedColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingAnalytics Gets consolidated masking analytics data based on the specified query parameters.
// If CompartmentIdInSubtreeQueryParam is specified as true, the behaviour
// is equivalent to accessLevel "ACCESSIBLE" by default.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingAnalytics.go.html to see an example of how to use ListMaskingAnalytics API.
// A default retry strategy applies to this operation ListMaskingAnalytics()
func (client DataSafeClient) ListMaskingAnalytics(ctx context.Context, request ListMaskingAnalyticsRequest) (response ListMaskingAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/ListMaskingAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingColumns Gets a list of masking columns present in the specified masking policy and based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingColumns.go.html to see an example of how to use ListMaskingColumns API.
// A default retry strategy applies to this operation ListMaskingColumns()
func (client DataSafeClient) ListMaskingColumns(ctx context.Context, request ListMaskingColumnsRequest) (response ListMaskingColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/ListMaskingColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingObjects Gets a list of masking objects present in the specified masking policy and based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingObjects.go.html to see an example of how to use ListMaskingObjects API.
// A default retry strategy applies to this operation ListMaskingObjects()
func (client DataSafeClient) ListMaskingObjects(ctx context.Context, request ListMaskingObjectsRequest) (response ListMaskingObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingObjectsResponse")
	}
	return
}

// listMaskingObjects implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicies/{maskingPolicyId}/maskingObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingObjectCollection/ListMaskingObjects"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingPolicies Gets a list of masking policies based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicies.go.html to see an example of how to use ListMaskingPolicies API.
// A default retry strategy applies to this operation ListMaskingPolicies()
func (client DataSafeClient) ListMaskingPolicies(ctx context.Context, request ListMaskingPoliciesRequest) (response ListMaskingPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/ListMaskingPolicies"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingPolicyHealthReportLogs Gets a list of errors and warnings from a masking policy health check.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicyHealthReportLogs.go.html to see an example of how to use ListMaskingPolicyHealthReportLogs API.
// A default retry strategy applies to this operation ListMaskingPolicyHealthReportLogs()
func (client DataSafeClient) ListMaskingPolicyHealthReportLogs(ctx context.Context, request ListMaskingPolicyHealthReportLogsRequest) (response ListMaskingPolicyHealthReportLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingPolicyHealthReportLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingPolicyHealthReportLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingPolicyHealthReportLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingPolicyHealthReportLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingPolicyHealthReportLogsResponse")
	}
	return
}

// listMaskingPolicyHealthReportLogs implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingPolicyHealthReportLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicyHealthReports/{maskingPolicyHealthReportId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingPolicyHealthReportLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicyHealthReport/ListMaskingPolicyHealthReportLogs"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingPolicyHealthReportLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingPolicyHealthReports Gets a list of masking policy health reports based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicyHealthReports.go.html to see an example of how to use ListMaskingPolicyHealthReports API.
// A default retry strategy applies to this operation ListMaskingPolicyHealthReports()
func (client DataSafeClient) ListMaskingPolicyHealthReports(ctx context.Context, request ListMaskingPolicyHealthReportsRequest) (response ListMaskingPolicyHealthReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingPolicyHealthReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingPolicyHealthReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingPolicyHealthReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingPolicyHealthReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingPolicyHealthReportsResponse")
	}
	return
}

// listMaskingPolicyHealthReports implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingPolicyHealthReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicyHealthReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingPolicyHealthReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicyHealthReport/ListMaskingPolicyHealthReports"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingPolicyHealthReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingReports Gets a list of masking reports based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingReports.go.html to see an example of how to use ListMaskingReports API.
// A default retry strategy applies to this operation ListMaskingReports()
func (client DataSafeClient) ListMaskingReports(ctx context.Context, request ListMaskingReportsRequest) (response ListMaskingReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/ListMaskingReports"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaskingSchemas Gets a list of masking schemas present in the specified masking policy and based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingSchemas.go.html to see an example of how to use ListMaskingSchemas API.
// A default retry strategy applies to this operation ListMaskingSchemas()
func (client DataSafeClient) ListMaskingSchemas(ctx context.Context, request ListMaskingSchemasRequest) (response ListMaskingSchemasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaskingSchemas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaskingSchemasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaskingSchemasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaskingSchemasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaskingSchemasResponse")
	}
	return
}

// listMaskingSchemas implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listMaskingSchemas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maskingPolicies/{maskingPolicyId}/maskingSchemas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaskingSchemasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingSchemaCollection/ListMaskingSchemas"
		err = common.PostProcessServiceError(err, "DataSafe", "ListMaskingSchemas", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOnPremConnectors Gets a list of on-premises connectors.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListOnPremConnectors.go.html to see an example of how to use ListOnPremConnectors API.
// A default retry strategy applies to this operation ListOnPremConnectors()
func (client DataSafeClient) ListOnPremConnectors(ctx context.Context, request ListOnPremConnectorsRequest) (response ListOnPremConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnectorSummary/ListOnPremConnectors"
		err = common.PostProcessServiceError(err, "DataSafe", "ListOnPremConnectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPeerTargetDatabases Lists all the peer target databases under the primary target database identified by the OCID passed as path parameter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListPeerTargetDatabases.go.html to see an example of how to use ListPeerTargetDatabases API.
// A default retry strategy applies to this operation ListPeerTargetDatabases()
func (client DataSafeClient) ListPeerTargetDatabases(ctx context.Context, request ListPeerTargetDatabasesRequest) (response ListPeerTargetDatabasesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listPeerTargetDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPeerTargetDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPeerTargetDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPeerTargetDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPeerTargetDatabasesResponse")
	}
	return
}

// listPeerTargetDatabases implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listPeerTargetDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targetDatabases/{targetDatabaseId}/peerTargetDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPeerTargetDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/PeerTargetDatabase/ListPeerTargetDatabases"
		err = common.PostProcessServiceError(err, "DataSafe", "ListPeerTargetDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfileAnalytics Gets a list of aggregated user profile details in the specified compartment. This provides information about the
// overall profiles available. For example, the user profile details include how many users have the profile assigned
// and do how many use password verification function. This data is especially useful content for dashboards or to support analytics.
// When you perform the ListProfileAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
// parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
// permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
// root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
// compartmentId, then "Not Authorized" is returned.
// The parameter compartmentIdInSubtree applies when you perform ListProfileAnalytics on the compartmentId passed and when it is
// set to true, the entire hierarchy of compartments can be returned.
// To use ListProfileAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
// set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListProfileAnalytics.go.html to see an example of how to use ListProfileAnalytics API.
// A default retry strategy applies to this operation ListProfileAnalytics()
func (client DataSafeClient) ListProfileAnalytics(ctx context.Context, request ListProfileAnalyticsRequest) (response ListProfileAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfileAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfileAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfileAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfileAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfileAnalyticsResponse")
	}
	return
}

// listProfileAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listProfileAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/profileAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfileAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Profile/ListProfileAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListProfileAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfileSummaries Gets a list of user profiles containing the profile details along with the target id and user counts.
// The ListProfiles operation returns only the profiles belonging to a certain target. If compartment type user assessment
// id is provided, then profile information for all the targets belonging to the pertaining compartment is returned.
// The list does not include any subcompartments of the compartment under consideration.
// The parameter 'accessLevel' specifies whether to return only those compartments for which the requestor has
// INSPECT permissions on at least one resource directly or indirectly (ACCESSIBLE) (the resource can be in a
// subcompartment) or to return Not Authorized if Principal doesn't have access to even one of the child compartments.
// This is valid only when 'compartmentIdInSubtree' is set to 'true'.
// The parameter 'compartmentIdInSubtree' applies when you perform ListUserProfiles on the 'compartmentId' belonging
// to the assessmentId passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment), set the parameter
// 'compartmentIdInSubtree' to true and 'accessLevel' to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListProfileSummaries.go.html to see an example of how to use ListProfileSummaries API.
// A default retry strategy applies to this operation ListProfileSummaries()
func (client DataSafeClient) ListProfileSummaries(ctx context.Context, request ListProfileSummariesRequest) (response ListProfileSummariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfileSummaries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfileSummariesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfileSummariesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfileSummariesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfileSummariesResponse")
	}
	return
}

// listProfileSummaries implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listProfileSummaries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/profiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfileSummariesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/ListProfileSummaries"
		err = common.PostProcessServiceError(err, "DataSafe", "ListProfileSummaries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReportDefinitions Gets a list of report definitions.
// The ListReportDefinitions operation returns only the report definitions in the specified `compartmentId`.
// It also returns the seeded report definitions which are available to all the compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReportDefinitions.go.html to see an example of how to use ListReportDefinitions API.
// A default retry strategy applies to this operation ListReportDefinitions()
func (client DataSafeClient) ListReportDefinitions(ctx context.Context, request ListReportDefinitionsRequest) (response ListReportDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/ListReportDefinitions"
		err = common.PostProcessServiceError(err, "DataSafe", "ListReportDefinitions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReports Gets a list of all the reports in the compartment. It contains information such as report generation time.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReports.go.html to see an example of how to use ListReports API.
// A default retry strategy applies to this operation ListReports()
func (client DataSafeClient) ListReports(ctx context.Context, request ListReportsRequest) (response ListReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportSummary/ListReports"
		err = common.PostProcessServiceError(err, "DataSafe", "ListReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoleGrantPaths Retrieves a list of all role grant paths for a particular user.
// The ListRoleGrantPaths operation returns only the role grant paths for the specified security policy report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListRoleGrantPaths.go.html to see an example of how to use ListRoleGrantPaths API.
// A default retry strategy applies to this operation ListRoleGrantPaths()
func (client DataSafeClient) ListRoleGrantPaths(ctx context.Context, request ListRoleGrantPathsRequest) (response ListRoleGrantPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoleGrantPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoleGrantPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoleGrantPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoleGrantPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoleGrantPathsResponse")
	}
	return
}

// listRoleGrantPaths implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listRoleGrantPaths(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports/{securityPolicyReportId}/roleGrantPaths", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRoleGrantPathsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/RoleGrantPathCollection/ListRoleGrantPaths"
		err = common.PostProcessServiceError(err, "DataSafe", "ListRoleGrantPaths", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoles Returns a list of role metadata objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListRoles.go.html to see an example of how to use ListRoles API.
// A default retry strategy applies to this operation ListRoles()
func (client DataSafeClient) ListRoles(ctx context.Context, request ListRolesRequest) (response ListRolesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/ListRoles"
		err = common.PostProcessServiceError(err, "DataSafe", "ListRoles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchemas Returns list of schema.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSchemas.go.html to see an example of how to use ListSchemas API.
// A default retry strategy applies to this operation ListSchemas()
func (client DataSafeClient) ListSchemas(ctx context.Context, request ListSchemasRequest) (response ListSchemasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/ListSchemas"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSchemas", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSdmMaskingPolicyDifferences Gets a list of SDM and masking policy difference resources based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSdmMaskingPolicyDifferences.go.html to see an example of how to use ListSdmMaskingPolicyDifferences API.
// A default retry strategy applies to this operation ListSdmMaskingPolicyDifferences()
func (client DataSafeClient) ListSdmMaskingPolicyDifferences(ctx context.Context, request ListSdmMaskingPolicyDifferencesRequest) (response ListSdmMaskingPolicyDifferencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSdmMaskingPolicyDifferences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSdmMaskingPolicyDifferencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSdmMaskingPolicyDifferencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSdmMaskingPolicyDifferencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSdmMaskingPolicyDifferencesResponse")
	}
	return
}

// listSdmMaskingPolicyDifferences implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSdmMaskingPolicyDifferences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdmMaskingPolicyDifferences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSdmMaskingPolicyDifferencesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/ListSdmMaskingPolicyDifferences"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSdmMaskingPolicyDifferences", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityAssessments.go.html to see an example of how to use ListSecurityAssessments API.
// A default retry strategy applies to this operation ListSecurityAssessments()
func (client DataSafeClient) ListSecurityAssessments(ctx context.Context, request ListSecurityAssessmentsRequest) (response ListSecurityAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessmentSummary/ListSecurityAssessments"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityAssessments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityFeatureAnalytics Gets a list of Database security feature usage aggregated details in the specified compartment. This provides information about the
// overall security controls, by returning the counting number of the target databases using the security features.
// When you perform the ListSecurityFeatureAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
// parameter accessLevel is set to ACCESSIBLE, then the operation returns statistics from the compartments in which the requestor has INSPECT
// permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
// root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
// compartmentId, then "Not Authorized" is returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityFeatureAnalytics.go.html to see an example of how to use ListSecurityFeatureAnalytics API.
// A default retry strategy applies to this operation ListSecurityFeatureAnalytics()
func (client DataSafeClient) ListSecurityFeatureAnalytics(ctx context.Context, request ListSecurityFeatureAnalyticsRequest) (response ListSecurityFeatureAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityFeatureAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityFeatureAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityFeatureAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityFeatureAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityFeatureAnalyticsResponse")
	}
	return
}

// listSecurityFeatureAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityFeatureAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/securityFeatureAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityFeatureAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/ListSecurityFeatureAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityFeatureAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityFeatures Lists the usage of Database security features for a given compartment or a target level, based on the filters provided.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityFeatures.go.html to see an example of how to use ListSecurityFeatures API.
// A default retry strategy applies to this operation ListSecurityFeatures()
func (client DataSafeClient) ListSecurityFeatures(ctx context.Context, request ListSecurityFeaturesRequest) (response ListSecurityFeaturesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityFeatures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityFeaturesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityFeaturesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityFeaturesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityFeaturesResponse")
	}
	return
}

// listSecurityFeatures implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityFeatures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAssessments/securityFeatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityFeaturesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/ListSecurityFeatures"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityFeatures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityPolicies Retrieves a list of all security policies in Data Safe.
// The ListSecurityPolicies operation returns only the security policies in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSecurityPolicies on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicies.go.html to see an example of how to use ListSecurityPolicies API.
// A default retry strategy applies to this operation ListSecurityPolicies()
func (client DataSafeClient) ListSecurityPolicies(ctx context.Context, request ListSecurityPoliciesRequest) (response ListSecurityPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityPoliciesResponse")
	}
	return
}

// listSecurityPolicies implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyCollection/ListSecurityPolicies"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityPolicyDeployments Retrieves a list of all security policy deployments in Data Safe.
// The ListSecurityPolicyDeployments operation returns only the security policy deployments in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSecurityPolicyDeployments on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyDeployments.go.html to see an example of how to use ListSecurityPolicyDeployments API.
// A default retry strategy applies to this operation ListSecurityPolicyDeployments()
func (client DataSafeClient) ListSecurityPolicyDeployments(ctx context.Context, request ListSecurityPolicyDeploymentsRequest) (response ListSecurityPolicyDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityPolicyDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityPolicyDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityPolicyDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityPolicyDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityPolicyDeploymentsResponse")
	}
	return
}

// listSecurityPolicyDeployments implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityPolicyDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityPolicyDeploymentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyDeploymentCollection/ListSecurityPolicyDeployments"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityPolicyDeployments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityPolicyEntryStates Retrieves a list of all security policy entry states in Data Safe.
// The ListSecurityPolicyEntryStates operation returns only the security policy entry states for the specified security policy entry.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyEntryStates.go.html to see an example of how to use ListSecurityPolicyEntryStates API.
// A default retry strategy applies to this operation ListSecurityPolicyEntryStates()
func (client DataSafeClient) ListSecurityPolicyEntryStates(ctx context.Context, request ListSecurityPolicyEntryStatesRequest) (response ListSecurityPolicyEntryStatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityPolicyEntryStates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityPolicyEntryStatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityPolicyEntryStatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityPolicyEntryStatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityPolicyEntryStatesResponse")
	}
	return
}

// listSecurityPolicyEntryStates implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityPolicyEntryStates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyDeployments/{securityPolicyDeploymentId}/securityPolicyEntryStates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityPolicyEntryStatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyEntryStateCollection/ListSecurityPolicyEntryStates"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityPolicyEntryStates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityPolicyReports Retrieves a list of all security policy reports in Data Safe.
// The ListSecurityPolicyReports operation returns only the security policy reports in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSecurityPolicyReports on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyReports.go.html to see an example of how to use ListSecurityPolicyReports API.
// A default retry strategy applies to this operation ListSecurityPolicyReports()
func (client DataSafeClient) ListSecurityPolicyReports(ctx context.Context, request ListSecurityPolicyReportsRequest) (response ListSecurityPolicyReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityPolicyReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityPolicyReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityPolicyReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityPolicyReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityPolicyReportsResponse")
	}
	return
}

// listSecurityPolicyReports implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSecurityPolicyReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicyReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityPolicyReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyReportCollection/ListSecurityPolicyReports"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSecurityPolicyReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSensitiveColumns Gets a list of sensitive columns present in the specified sensitive data model based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveColumns.go.html to see an example of how to use ListSensitiveColumns API.
// A default retry strategy applies to this operation ListSensitiveColumns()
func (client DataSafeClient) ListSensitiveColumns(ctx context.Context, request ListSensitiveColumnsRequest) (response ListSensitiveColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveColumn/ListSensitiveColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSensitiveColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSensitiveDataModelSensitiveTypes Gets a list of sensitive type Ids present in the specified sensitive data model.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveDataModelSensitiveTypes.go.html to see an example of how to use ListSensitiveDataModelSensitiveTypes API.
// A default retry strategy applies to this operation ListSensitiveDataModelSensitiveTypes()
func (client DataSafeClient) ListSensitiveDataModelSensitiveTypes(ctx context.Context, request ListSensitiveDataModelSensitiveTypesRequest) (response ListSensitiveDataModelSensitiveTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSensitiveDataModelSensitiveTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSensitiveDataModelSensitiveTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSensitiveDataModelSensitiveTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSensitiveDataModelSensitiveTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSensitiveDataModelSensitiveTypesResponse")
	}
	return
}

// listSensitiveDataModelSensitiveTypes implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSensitiveDataModelSensitiveTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSensitiveDataModelSensitiveTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModelSensitiveTypeCollection/ListSensitiveDataModelSensitiveTypes"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSensitiveDataModelSensitiveTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSensitiveDataModels Gets a list of sensitive data models based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveDataModels.go.html to see an example of how to use ListSensitiveDataModels API.
// A default retry strategy applies to this operation ListSensitiveDataModels()
func (client DataSafeClient) ListSensitiveDataModels(ctx context.Context, request ListSensitiveDataModelsRequest) (response ListSensitiveDataModelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/ListSensitiveDataModels"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSensitiveDataModels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSensitiveObjects Gets a list of sensitive objects present in the specified sensitive data model based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveObjects.go.html to see an example of how to use ListSensitiveObjects API.
// A default retry strategy applies to this operation ListSensitiveObjects()
func (client DataSafeClient) ListSensitiveObjects(ctx context.Context, request ListSensitiveObjectsRequest) (response ListSensitiveObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSensitiveObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSensitiveObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSensitiveObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSensitiveObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSensitiveObjectsResponse")
	}
	return
}

// listSensitiveObjects implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSensitiveObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSensitiveObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveObjectCollection/ListSensitiveObjects"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSensitiveObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSensitiveSchemas Gets a list of sensitive schemas present in the specified sensitive data model based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveSchemas.go.html to see an example of how to use ListSensitiveSchemas API.
// A default retry strategy applies to this operation ListSensitiveSchemas()
func (client DataSafeClient) ListSensitiveSchemas(ctx context.Context, request ListSensitiveSchemasRequest) (response ListSensitiveSchemasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSensitiveSchemas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSensitiveSchemasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSensitiveSchemasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSensitiveSchemasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSensitiveSchemasResponse")
	}
	return
}

// listSensitiveSchemas implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSensitiveSchemas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sensitiveDataModels/{sensitiveDataModelId}/sensitiveSchemas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSensitiveSchemasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveSchemaCollection/ListSensitiveSchemas"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSensitiveSchemas", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSensitiveTypes Gets a list of sensitive types based on the specified query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveTypes.go.html to see an example of how to use ListSensitiveTypes API.
// A default retry strategy applies to this operation ListSensitiveTypes()
func (client DataSafeClient) ListSensitiveTypes(ctx context.Context, request ListSensitiveTypesRequest) (response ListSensitiveTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveType/ListSensitiveTypes"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSensitiveTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlCollectionAnalytics Retrieves a list of all SQL collection analytics in Data Safe.
// The ListSqlCollectionAnalytics operation returns only the analytics for the SQL collections in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSqlCollections on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlCollectionAnalytics.go.html to see an example of how to use ListSqlCollectionAnalytics API.
// A default retry strategy applies to this operation ListSqlCollectionAnalytics()
func (client DataSafeClient) ListSqlCollectionAnalytics(ctx context.Context, request ListSqlCollectionAnalyticsRequest) (response ListSqlCollectionAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlCollectionAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlCollectionAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlCollectionAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlCollectionAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlCollectionAnalyticsResponse")
	}
	return
}

// listSqlCollectionAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlCollectionAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlCollectionAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlCollectionAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollectionAnalyticsCollection/ListSqlCollectionAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlCollectionAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlCollectionLogInsights Retrieves a list of the SQL collection log analytics.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlCollectionLogInsights.go.html to see an example of how to use ListSqlCollectionLogInsights API.
// A default retry strategy applies to this operation ListSqlCollectionLogInsights()
func (client DataSafeClient) ListSqlCollectionLogInsights(ctx context.Context, request ListSqlCollectionLogInsightsRequest) (response ListSqlCollectionLogInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlCollectionLogInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlCollectionLogInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlCollectionLogInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlCollectionLogInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlCollectionLogInsightsResponse")
	}
	return
}

// listSqlCollectionLogInsights implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlCollectionLogInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlCollections/{sqlCollectionId}/logInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlCollectionLogInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollectionLogInsightsCollection/ListSqlCollectionLogInsights"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlCollectionLogInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlCollections Retrieves a list of all SQL collections in Data Safe.
// The ListSqlCollections operation returns only the SQL collections in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSqlCollections on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlCollections.go.html to see an example of how to use ListSqlCollections API.
// A default retry strategy applies to this operation ListSqlCollections()
func (client DataSafeClient) ListSqlCollections(ctx context.Context, request ListSqlCollectionsRequest) (response ListSqlCollectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlCollections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlCollectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlCollectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlCollectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlCollectionsResponse")
	}
	return
}

// listSqlCollections implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlCollections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlCollections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlCollectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollectionCollection/ListSqlCollections"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlCollections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlFirewallAllowedSqlAnalytics Returns the aggregation details of all SQL Firewall allowed SQL statements.
// The ListSqlFirewallAllowedSqlAnalytics operation returns the aggregates of the SQL Firewall allowed SQL statements in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSqlFirewallAllowedSqlAnalytics on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallAllowedSqlAnalytics.go.html to see an example of how to use ListSqlFirewallAllowedSqlAnalytics API.
// A default retry strategy applies to this operation ListSqlFirewallAllowedSqlAnalytics()
func (client DataSafeClient) ListSqlFirewallAllowedSqlAnalytics(ctx context.Context, request ListSqlFirewallAllowedSqlAnalyticsRequest) (response ListSqlFirewallAllowedSqlAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlFirewallAllowedSqlAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlFirewallAllowedSqlAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlFirewallAllowedSqlAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlFirewallAllowedSqlAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlFirewallAllowedSqlAnalyticsResponse")
	}
	return
}

// listSqlFirewallAllowedSqlAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlFirewallAllowedSqlAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallAllowedSqlAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlFirewallAllowedSqlAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallAllowedSqlAnalyticsCollection/ListSqlFirewallAllowedSqlAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlFirewallAllowedSqlAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlFirewallAllowedSqls Retrieves a list of all SQL Firewall allowed SQL statements.
// The ListSqlFirewallAllowedSqls operation returns only the SQL Firewall allowed SQL statements in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSqlFirewallPolicies on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallAllowedSqls.go.html to see an example of how to use ListSqlFirewallAllowedSqls API.
// A default retry strategy applies to this operation ListSqlFirewallAllowedSqls()
func (client DataSafeClient) ListSqlFirewallAllowedSqls(ctx context.Context, request ListSqlFirewallAllowedSqlsRequest) (response ListSqlFirewallAllowedSqlsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlFirewallAllowedSqls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlFirewallAllowedSqlsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlFirewallAllowedSqlsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlFirewallAllowedSqlsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlFirewallAllowedSqlsResponse")
	}
	return
}

// listSqlFirewallAllowedSqls implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlFirewallAllowedSqls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallAllowedSqls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlFirewallAllowedSqlsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallAllowedSqlCollection/ListSqlFirewallAllowedSqls"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlFirewallAllowedSqls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlFirewallPolicies Retrieves a list of all SQL Firewall policies.
// The ListSqlFirewallPolicies operation returns only the SQL Firewall policies in the specified `compartmentId`.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSqlFirewallPolicies on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallPolicies.go.html to see an example of how to use ListSqlFirewallPolicies API.
// A default retry strategy applies to this operation ListSqlFirewallPolicies()
func (client DataSafeClient) ListSqlFirewallPolicies(ctx context.Context, request ListSqlFirewallPoliciesRequest) (response ListSqlFirewallPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlFirewallPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlFirewallPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlFirewallPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlFirewallPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlFirewallPoliciesResponse")
	}
	return
}

// listSqlFirewallPolicies implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlFirewallPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlFirewallPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallPolicyCollection/ListSqlFirewallPolicies"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlFirewallPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlFirewallPolicyAnalytics Gets a list of aggregated SQL Firewall policy details.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform SummarizedSqlFirewallPolicyInfo on the specified
// `compartmentId` and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallPolicyAnalytics.go.html to see an example of how to use ListSqlFirewallPolicyAnalytics API.
// A default retry strategy applies to this operation ListSqlFirewallPolicyAnalytics()
func (client DataSafeClient) ListSqlFirewallPolicyAnalytics(ctx context.Context, request ListSqlFirewallPolicyAnalyticsRequest) (response ListSqlFirewallPolicyAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlFirewallPolicyAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlFirewallPolicyAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlFirewallPolicyAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlFirewallPolicyAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlFirewallPolicyAnalyticsResponse")
	}
	return
}

// listSqlFirewallPolicyAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlFirewallPolicyAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallPolicyAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlFirewallPolicyAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallPolicyAnalyticsCollection/ListSqlFirewallPolicyAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlFirewallPolicyAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlFirewallViolationAnalytics Returns the aggregation details of the SQL Firewall violations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallViolationAnalytics.go.html to see an example of how to use ListSqlFirewallViolationAnalytics API.
// A default retry strategy applies to this operation ListSqlFirewallViolationAnalytics()
func (client DataSafeClient) ListSqlFirewallViolationAnalytics(ctx context.Context, request ListSqlFirewallViolationAnalyticsRequest) (response ListSqlFirewallViolationAnalyticsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listSqlFirewallViolationAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlFirewallViolationAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlFirewallViolationAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlFirewallViolationAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlFirewallViolationAnalyticsResponse")
	}
	return
}

// listSqlFirewallViolationAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlFirewallViolationAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallViolationAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlFirewallViolationAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallViolationSummary/ListSqlFirewallViolationAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlFirewallViolationAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlFirewallViolations Gets a list of all the SQL Firewall violations captured by the firewall.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallViolations.go.html to see an example of how to use ListSqlFirewallViolations API.
// A default retry strategy applies to this operation ListSqlFirewallViolations()
func (client DataSafeClient) ListSqlFirewallViolations(ctx context.Context, request ListSqlFirewallViolationsRequest) (response ListSqlFirewallViolationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlFirewallViolations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlFirewallViolationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlFirewallViolationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlFirewallViolationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlFirewallViolationsResponse")
	}
	return
}

// listSqlFirewallViolations implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listSqlFirewallViolations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sqlFirewallViolations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlFirewallViolationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallViolationSummary/ListSqlFirewallViolations"
		err = common.PostProcessServiceError(err, "DataSafe", "ListSqlFirewallViolations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTables Returns a list of table metadata objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTables.go.html to see an example of how to use ListTables API.
// A default retry strategy applies to this operation ListTables()
func (client DataSafeClient) ListTables(ctx context.Context, request ListTablesRequest) (response ListTablesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/ListTables"
		err = common.PostProcessServiceError(err, "DataSafe", "ListTables", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetAlertPolicyAssociations Gets a list of all target-alert policy associations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetAlertPolicyAssociations.go.html to see an example of how to use ListTargetAlertPolicyAssociations API.
// A default retry strategy applies to this operation ListTargetAlertPolicyAssociations()
func (client DataSafeClient) ListTargetAlertPolicyAssociations(ctx context.Context, request ListTargetAlertPolicyAssociationsRequest) (response ListTargetAlertPolicyAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociationSummary/ListTargetAlertPolicyAssociations"
		err = common.PostProcessServiceError(err, "DataSafe", "ListTargetAlertPolicyAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetDatabases Returns the list of registered target databases in Data Safe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetDatabases.go.html to see an example of how to use ListTargetDatabases API.
// A default retry strategy applies to this operation ListTargetDatabases()
func (client DataSafeClient) ListTargetDatabases(ctx context.Context, request ListTargetDatabasesRequest) (response ListTargetDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabaseSummary/ListTargetDatabases"
		err = common.PostProcessServiceError(err, "DataSafe", "ListTargetDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserAccessAnalytics Gets a list of aggregated user access analytics in the specified target in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAccessAnalytics.go.html to see an example of how to use ListUserAccessAnalytics API.
// A default retry strategy applies to this operation ListUserAccessAnalytics()
func (client DataSafeClient) ListUserAccessAnalytics(ctx context.Context, request ListUserAccessAnalyticsRequest) (response ListUserAccessAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserAccessAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserAccessAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserAccessAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserAccessAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserAccessAnalyticsResponse")
	}
	return
}

// listUserAccessAnalytics implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) listUserAccessAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userAssessments/{userAssessmentId}/userAccessAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserAccessAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/ListUserAccessAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListUserAccessAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUserAnalytics Gets a list of aggregated user details from the specified user assessment. This provides information about the overall state.
// of database user security.  For example, the user details include how many users have the DBA role and how many users are in
// the critical category. This data is especially useful content for dashboards or to support analytics.
// When you perform the ListUserAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
// parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
// permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
// root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
// compartmentId, then "Not Authorized" is returned.
// The parameter compartmentIdInSubtree applies when you perform ListUserAnalytics on the compartmentId passed and when it is
// set to true, the entire hierarchy of compartments can be returned.
// To use ListUserAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
// set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAnalytics.go.html to see an example of how to use ListUserAnalytics API.
// A default retry strategy applies to this operation ListUserAnalytics()
func (client DataSafeClient) ListUserAnalytics(ctx context.Context, request ListUserAnalyticsRequest) (response ListUserAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/ListUserAnalytics"
		err = common.PostProcessServiceError(err, "DataSafe", "ListUserAnalytics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAssessments.go.html to see an example of how to use ListUserAssessments API.
// A default retry strategy applies to this operation ListUserAssessments()
func (client DataSafeClient) ListUserAssessments(ctx context.Context, request ListUserAssessmentsRequest) (response ListUserAssessmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessmentSummary/ListUserAssessments"
		err = common.PostProcessServiceError(err, "DataSafe", "ListUserAssessments", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUsers.go.html to see an example of how to use ListUsers API.
// A default retry strategy applies to this operation ListUsers()
func (client DataSafeClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/ListUsers"
		err = common.PostProcessServiceError(err, "DataSafe", "ListUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Gets a list of errors for the specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DataSafeClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DataSafe", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Gets a list of log entries for the specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DataSafeClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DataSafe", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Gets a list of work requests.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DataSafeClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DataSafe", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MaskData Masks data using the specified masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/MaskData.go.html to see an example of how to use MaskData API.
// A default retry strategy applies to this operation MaskData()
func (client DataSafeClient) MaskData(ctx context.Context, request MaskDataRequest) (response MaskDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/MaskData"
		err = common.PostProcessServiceError(err, "DataSafe", "MaskData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ModifyGlobalSettings Modifies Global Settings in Data Safe in the tenancy and region.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ModifyGlobalSettings.go.html to see an example of how to use ModifyGlobalSettings API.
// A default retry strategy applies to this operation ModifyGlobalSettings()
func (client DataSafeClient) ModifyGlobalSettings(ctx context.Context, request ModifyGlobalSettingsRequest) (response ModifyGlobalSettingsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafeConfiguration/ModifyGlobalSettings"
		err = common.PostProcessServiceError(err, "DataSafe", "ModifyGlobalSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchAlerts Updates the status of one or more alert specified by the alert IDs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchAlerts.go.html to see an example of how to use PatchAlerts API.
// A default retry strategy applies to this operation PatchAlerts()
func (client DataSafeClient) PatchAlerts(ctx context.Context, request PatchAlertsRequest) (response PatchAlertsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Alert/PatchAlerts"
		err = common.PostProcessServiceError(err, "DataSafe", "PatchAlerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchDiscoveryJobResults Patches one or more discovery results. You can use this operation to set the plannedAction attribute before using
// ApplyDiscoveryJobResults to process the results based on this attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchDiscoveryJobResults.go.html to see an example of how to use PatchDiscoveryJobResults API.
// A default retry strategy applies to this operation PatchDiscoveryJobResults()
func (client DataSafeClient) PatchDiscoveryJobResults(ctx context.Context, request PatchDiscoveryJobResultsRequest) (response PatchDiscoveryJobResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DiscoveryJob/PatchDiscoveryJobResults"
		err = common.PostProcessServiceError(err, "DataSafe", "PatchDiscoveryJobResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchMaskingColumns Patches one or more columns in the specified masking policy. Use it to create, or update
// masking columns. To create masking columns, use CreateMaskingColumnDetails as the patch
// value. And to update masking columns, use UpdateMaskingColumnDetails as the patch value.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchMaskingColumns.go.html to see an example of how to use PatchMaskingColumns API.
// A default retry strategy applies to this operation PatchMaskingColumns()
func (client DataSafeClient) PatchMaskingColumns(ctx context.Context, request PatchMaskingColumnsRequest) (response PatchMaskingColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/PatchMaskingColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "PatchMaskingColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchSdmMaskingPolicyDifferenceColumns Patches one or more SDM masking policy difference columns. You can use this operation to set the plannedAction attribute before using
// ApplySdmMaskingPolicyDifference to process the difference based on this attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchSdmMaskingPolicyDifferenceColumns.go.html to see an example of how to use PatchSdmMaskingPolicyDifferenceColumns API.
// A default retry strategy applies to this operation PatchSdmMaskingPolicyDifferenceColumns()
func (client DataSafeClient) PatchSdmMaskingPolicyDifferenceColumns(ctx context.Context, request PatchSdmMaskingPolicyDifferenceColumnsRequest) (response PatchSdmMaskingPolicyDifferenceColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchSdmMaskingPolicyDifferenceColumns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchSdmMaskingPolicyDifferenceColumnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchSdmMaskingPolicyDifferenceColumnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchSdmMaskingPolicyDifferenceColumnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchSdmMaskingPolicyDifferenceColumnsResponse")
	}
	return
}

// patchSdmMaskingPolicyDifferenceColumns implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) patchSdmMaskingPolicyDifferenceColumns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}/differenceColumns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchSdmMaskingPolicyDifferenceColumnsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/PatchSdmMaskingPolicyDifferenceColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "PatchSdmMaskingPolicyDifferenceColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchSensitiveColumns Patches one or more columns in the specified sensitive data model. Use it to create, update, or delete sensitive columns.
// To create sensitive columns, use CreateSensitiveColumnDetails as the patch value. And to update sensitive columns,
// use UpdateSensitiveColumnDetails as the patch value.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchSensitiveColumns.go.html to see an example of how to use PatchSensitiveColumns API.
// A default retry strategy applies to this operation PatchSensitiveColumns()
func (client DataSafeClient) PatchSensitiveColumns(ctx context.Context, request PatchSensitiveColumnsRequest) (response PatchSensitiveColumnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveColumn/PatchSensitiveColumns"
		err = common.PostProcessServiceError(err, "DataSafe", "PatchSensitiveColumns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchTargetAlertPolicyAssociation Creates new target-alert policy associations that will be applied on the target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchTargetAlertPolicyAssociation.go.html to see an example of how to use PatchTargetAlertPolicyAssociation API.
// A default retry strategy applies to this operation PatchTargetAlertPolicyAssociation()
func (client DataSafeClient) PatchTargetAlertPolicyAssociation(ctx context.Context, request PatchTargetAlertPolicyAssociationRequest) (response PatchTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchTargetAlertPolicyAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchTargetAlertPolicyAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchTargetAlertPolicyAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchTargetAlertPolicyAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchTargetAlertPolicyAssociationResponse")
	}
	return
}

// patchTargetAlertPolicyAssociation implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) patchTargetAlertPolicyAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/targetAlertPolicyAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchTargetAlertPolicyAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociation/PatchTargetAlertPolicyAssociation"
		err = common.PostProcessServiceError(err, "DataSafe", "PatchTargetAlertPolicyAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ProvisionAuditPolicy Provision audit policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ProvisionAuditPolicy.go.html to see an example of how to use ProvisionAuditPolicy API.
// A default retry strategy applies to this operation ProvisionAuditPolicy()
func (client DataSafeClient) ProvisionAuditPolicy(ctx context.Context, request ProvisionAuditPolicyRequest) (response ProvisionAuditPolicyResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicy/ProvisionAuditPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "ProvisionAuditPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PurgeSqlCollectionLogs Purge the SQL collection logs for the specified SqlCollection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PurgeSqlCollectionLogs.go.html to see an example of how to use PurgeSqlCollectionLogs API.
// A default retry strategy applies to this operation PurgeSqlCollectionLogs()
func (client DataSafeClient) PurgeSqlCollectionLogs(ctx context.Context, request PurgeSqlCollectionLogsRequest) (response PurgeSqlCollectionLogsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.purgeSqlCollectionLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PurgeSqlCollectionLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PurgeSqlCollectionLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PurgeSqlCollectionLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PurgeSqlCollectionLogsResponse")
	}
	return
}

// purgeSqlCollectionLogs implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) purgeSqlCollectionLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections/{sqlCollectionId}/actions/purgeLogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PurgeSqlCollectionLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/PurgeSqlCollectionLogs"
		err = common.PostProcessServiceError(err, "DataSafe", "PurgeSqlCollectionLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshDatabaseSecurityConfiguration Refreshes the specified database security configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshDatabaseSecurityConfiguration.go.html to see an example of how to use RefreshDatabaseSecurityConfiguration API.
// A default retry strategy applies to this operation RefreshDatabaseSecurityConfiguration()
func (client DataSafeClient) RefreshDatabaseSecurityConfiguration(ctx context.Context, request RefreshDatabaseSecurityConfigurationRequest) (response RefreshDatabaseSecurityConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshDatabaseSecurityConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshDatabaseSecurityConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshDatabaseSecurityConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshDatabaseSecurityConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshDatabaseSecurityConfigurationResponse")
	}
	return
}

// refreshDatabaseSecurityConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshDatabaseSecurityConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseSecurityConfigs/{databaseSecurityConfigId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshDatabaseSecurityConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseSecurityConfig/RefreshDatabaseSecurityConfiguration"
		err = common.PostProcessServiceError(err, "DataSafe", "RefreshDatabaseSecurityConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshSecurityAssessment Runs a security assessment, refreshes the latest assessment, and saves it for future reference.
// The assessment runs with a securityAssessmentId of type LATEST. Before you start, first call the ListSecurityAssessments operation with filter "type = latest" to get the security assessment id for the target's latest assessment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshSecurityAssessment.go.html to see an example of how to use RefreshSecurityAssessment API.
// A default retry strategy applies to this operation RefreshSecurityAssessment()
func (client DataSafeClient) RefreshSecurityAssessment(ctx context.Context, request RefreshSecurityAssessmentRequest) (response RefreshSecurityAssessmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/RefreshSecurityAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "RefreshSecurityAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshSqlCollectionLogInsights Refresh the specified SQL collection Log Insights.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshSqlCollectionLogInsights.go.html to see an example of how to use RefreshSqlCollectionLogInsights API.
// A default retry strategy applies to this operation RefreshSqlCollectionLogInsights()
func (client DataSafeClient) RefreshSqlCollectionLogInsights(ctx context.Context, request RefreshSqlCollectionLogInsightsRequest) (response RefreshSqlCollectionLogInsightsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshSqlCollectionLogInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshSqlCollectionLogInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshSqlCollectionLogInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshSqlCollectionLogInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshSqlCollectionLogInsightsResponse")
	}
	return
}

// refreshSqlCollectionLogInsights implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshSqlCollectionLogInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections/{sqlCollectionId}/actions/refreshLogInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshSqlCollectionLogInsightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/RefreshSqlCollectionLogInsights"
		err = common.PostProcessServiceError(err, "DataSafe", "RefreshSqlCollectionLogInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshTargetDatabase Refreshes the Data Safe target database to update it's state.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshTargetDatabase.go.html to see an example of how to use RefreshTargetDatabase API.
// A default retry strategy applies to this operation RefreshTargetDatabase()
func (client DataSafeClient) RefreshTargetDatabase(ctx context.Context, request RefreshTargetDatabaseRequest) (response RefreshTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshTargetDatabaseResponse")
	}
	return
}

// refreshTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) refreshTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targetDatabases/{targetDatabaseId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/RefreshTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "RefreshTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshUserAssessment Refreshes the latest assessment and saves it for future reference. This operation runs with a userAssessmentId of type LATEST.
// Before you start, first call the ListUserAssessments operation with filter "type = latest" to get the user assessment ID for
// the target's latest assessment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RefreshUserAssessment.go.html to see an example of how to use RefreshUserAssessment API.
// A default retry strategy applies to this operation RefreshUserAssessment()
func (client DataSafeClient) RefreshUserAssessment(ctx context.Context, request RefreshUserAssessmentRequest) (response RefreshUserAssessmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/RefreshUserAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "RefreshUserAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveScheduleReport Deletes the schedule of a .xls or .pdf report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RemoveScheduleReport.go.html to see an example of how to use RemoveScheduleReport API.
// A default retry strategy applies to this operation RemoveScheduleReport()
func (client DataSafeClient) RemoveScheduleReport(ctx context.Context, request RemoveScheduleReportRequest) (response RemoveScheduleReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeScheduleReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveScheduleReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveScheduleReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveScheduleReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveScheduleReportResponse")
	}
	return
}

// removeScheduleReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) removeScheduleReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reportDefinitions/{reportDefinitionId}/actions/removeScheduleReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveScheduleReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/RemoveScheduleReport"
		err = common.PostProcessServiceError(err, "DataSafe", "RemoveScheduleReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResumeAuditTrail Resumes the specified audit trail once it got stopped.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ResumeAuditTrail.go.html to see an example of how to use ResumeAuditTrail API.
// A default retry strategy applies to this operation ResumeAuditTrail()
func (client DataSafeClient) ResumeAuditTrail(ctx context.Context, request ResumeAuditTrailRequest) (response ResumeAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/ResumeAuditTrail"
		err = common.PostProcessServiceError(err, "DataSafe", "ResumeAuditTrail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResumeWorkRequest Resume the given work request. Issuing a resume does not guarantee of immediate resume of the work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ResumeWorkRequest.go.html to see an example of how to use ResumeWorkRequest API.
// A default retry strategy applies to this operation ResumeWorkRequest()
func (client DataSafeClient) ResumeWorkRequest(ctx context.Context, request ResumeWorkRequestRequest) (response ResumeWorkRequestResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequest/ResumeWorkRequest"
		err = common.PostProcessServiceError(err, "DataSafe", "ResumeWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveAuditPolicies Retrieves the audit policy details from the source target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/RetrieveAuditPolicies.go.html to see an example of how to use RetrieveAuditPolicies API.
// A default retry strategy applies to this operation RetrieveAuditPolicies()
func (client DataSafeClient) RetrieveAuditPolicies(ctx context.Context, request RetrieveAuditPoliciesRequest) (response RetrieveAuditPoliciesResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicy/RetrieveAuditPolicies"
		err = common.PostProcessServiceError(err, "DataSafe", "RetrieveAuditPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleReport Schedules a .xls or .pdf report based on parameters and report definition.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ScheduleReport.go.html to see an example of how to use ScheduleReport API.
// A default retry strategy applies to this operation ScheduleReport()
func (client DataSafeClient) ScheduleReport(ctx context.Context, request ScheduleReportRequest) (response ScheduleReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scheduleReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleReportResponse")
	}
	return
}

// scheduleReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) scheduleReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reportDefinitions/{reportDefinitionId}/actions/scheduleReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/ScheduleReport"
		err = common.PostProcessServiceError(err, "DataSafe", "ScheduleReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetSecurityAssessmentBaseline Sets the saved security assessment as the baseline in the compartment where the the specified assessment resides. The security assessment needs to be of type 'SAVED'.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SetSecurityAssessmentBaseline.go.html to see an example of how to use SetSecurityAssessmentBaseline API.
// A default retry strategy applies to this operation SetSecurityAssessmentBaseline()
func (client DataSafeClient) SetSecurityAssessmentBaseline(ctx context.Context, request SetSecurityAssessmentBaselineRequest) (response SetSecurityAssessmentBaselineResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/SetSecurityAssessmentBaseline"
		err = common.PostProcessServiceError(err, "DataSafe", "SetSecurityAssessmentBaseline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetUserAssessmentBaseline Sets the saved user assessment as the baseline in the compartment where the specified assessment resides. The user assessment needs to be of type 'SAVED'.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SetUserAssessmentBaseline.go.html to see an example of how to use SetUserAssessmentBaseline API.
// A default retry strategy applies to this operation SetUserAssessmentBaseline()
func (client DataSafeClient) SetUserAssessmentBaseline(ctx context.Context, request SetUserAssessmentBaselineRequest) (response SetUserAssessmentBaselineResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/SetUserAssessmentBaseline"
		err = common.PostProcessServiceError(err, "DataSafe", "SetUserAssessmentBaseline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartAuditTrail Starts collection of audit records on the specified audit trail.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/StartAuditTrail.go.html to see an example of how to use StartAuditTrail API.
// A default retry strategy applies to this operation StartAuditTrail()
func (client DataSafeClient) StartAuditTrail(ctx context.Context, request StartAuditTrailRequest) (response StartAuditTrailResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/StartAuditTrail"
		err = common.PostProcessServiceError(err, "DataSafe", "StartAuditTrail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartSqlCollection Start the specified SQL collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/StartSqlCollection.go.html to see an example of how to use StartSqlCollection API.
// A default retry strategy applies to this operation StartSqlCollection()
func (client DataSafeClient) StartSqlCollection(ctx context.Context, request StartSqlCollectionRequest) (response StartSqlCollectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startSqlCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartSqlCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartSqlCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartSqlCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartSqlCollectionResponse")
	}
	return
}

// startSqlCollection implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) startSqlCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections/{sqlCollectionId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartSqlCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/sqlCollection/StartSqlCollection"
		err = common.PostProcessServiceError(err, "DataSafe", "StartSqlCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopAuditTrail Stops the specified audit trail.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/StopAuditTrail.go.html to see an example of how to use StopAuditTrail API.
// A default retry strategy applies to this operation StopAuditTrail()
func (client DataSafeClient) StopAuditTrail(ctx context.Context, request StopAuditTrailRequest) (response StopAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/StopAuditTrail"
		err = common.PostProcessServiceError(err, "DataSafe", "StopAuditTrail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopSqlCollection Stops the specified SQL collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/StopSqlCollection.go.html to see an example of how to use StopSqlCollection API.
// A default retry strategy applies to this operation StopSqlCollection()
func (client DataSafeClient) StopSqlCollection(ctx context.Context, request StopSqlCollectionRequest) (response StopSqlCollectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopSqlCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopSqlCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopSqlCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopSqlCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopSqlCollectionResponse")
	}
	return
}

// stopSqlCollection implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) stopSqlCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sqlCollections/{sqlCollectionId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopSqlCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/StopSqlCollection"
		err = common.PostProcessServiceError(err, "DataSafe", "StopSqlCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SuspendWorkRequest Suspend the given work request. Issuing a suspend does not guarantee of a immediate suspend of the work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/SuspendWorkRequest.go.html to see an example of how to use SuspendWorkRequest API.
// A default retry strategy applies to this operation SuspendWorkRequest()
func (client DataSafeClient) SuspendWorkRequest(ctx context.Context, request SuspendWorkRequestRequest) (response SuspendWorkRequestResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/WorkRequest/SuspendWorkRequest"
		err = common.PostProcessServiceError(err, "DataSafe", "SuspendWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UnsetSecurityAssessmentBaseline Removes the baseline setting for the saved security assessment associated with the targetId passed via body.
// If no body or empty body is passed then the baseline settings of all the saved security assessments pertaining to the baseline assessment OCID provided in the path will be removed.
// Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UnsetSecurityAssessmentBaseline.go.html to see an example of how to use UnsetSecurityAssessmentBaseline API.
// A default retry strategy applies to this operation UnsetSecurityAssessmentBaseline()
func (client DataSafeClient) UnsetSecurityAssessmentBaseline(ctx context.Context, request UnsetSecurityAssessmentBaselineRequest) (response UnsetSecurityAssessmentBaselineResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/UnsetSecurityAssessmentBaseline"
		err = common.PostProcessServiceError(err, "DataSafe", "UnsetSecurityAssessmentBaseline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UnsetUserAssessmentBaseline Removes the baseline setting for the saved user assessment associated with the targetId passed via body.
// If no body or empty body is passed then the baseline settings of all the saved user assessments pertaining to the baseline assessment OCID provided in the path will be removed.
// Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UnsetUserAssessmentBaseline.go.html to see an example of how to use UnsetUserAssessmentBaseline API.
// A default retry strategy applies to this operation UnsetUserAssessmentBaseline()
func (client DataSafeClient) UnsetUserAssessmentBaseline(ctx context.Context, request UnsetUserAssessmentBaselineRequest) (response UnsetUserAssessmentBaselineResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/UnsetUserAssessmentBaseline"
		err = common.PostProcessServiceError(err, "DataSafe", "UnsetUserAssessmentBaseline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAlert Updates the status of the specified alert.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAlert.go.html to see an example of how to use UpdateAlert API.
// A default retry strategy applies to this operation UpdateAlert()
func (client DataSafeClient) UpdateAlert(ctx context.Context, request UpdateAlertRequest) (response UpdateAlertResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Alert/UpdateAlert"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateAlert", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAuditArchiveRetrieval Updates the audit archive retrieval.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditArchiveRetrieval.go.html to see an example of how to use UpdateAuditArchiveRetrieval API.
// A default retry strategy applies to this operation UpdateAuditArchiveRetrieval()
func (client DataSafeClient) UpdateAuditArchiveRetrieval(ctx context.Context, request UpdateAuditArchiveRetrievalRequest) (response UpdateAuditArchiveRetrievalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditArchiveRetrieval/UpdateAuditArchiveRetrieval"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateAuditArchiveRetrieval", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAuditPolicy Updates the audit policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditPolicy.go.html to see an example of how to use UpdateAuditPolicy API.
// A default retry strategy applies to this operation UpdateAuditPolicy()
func (client DataSafeClient) UpdateAuditPolicy(ctx context.Context, request UpdateAuditPolicyRequest) (response UpdateAuditPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditPolicy/UpdateAuditPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateAuditPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAuditProfile Updates one or more attributes of the specified audit profile.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditProfile.go.html to see an example of how to use UpdateAuditProfile API.
// A default retry strategy applies to this operation UpdateAuditProfile()
func (client DataSafeClient) UpdateAuditProfile(ctx context.Context, request UpdateAuditProfileRequest) (response UpdateAuditProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditProfile/UpdateAuditProfile"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateAuditProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAuditTrail Updates one or more attributes of the specified audit trail.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateAuditTrail.go.html to see an example of how to use UpdateAuditTrail API.
// A default retry strategy applies to this operation UpdateAuditTrail()
func (client DataSafeClient) UpdateAuditTrail(ctx context.Context, request UpdateAuditTrailRequest) (response UpdateAuditTrailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/AuditTrail/UpdateAuditTrail"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateAuditTrail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataSafePrivateEndpoint Updates one or more attributes of the specified Data Safe private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateDataSafePrivateEndpoint.go.html to see an example of how to use UpdateDataSafePrivateEndpoint API.
// A default retry strategy applies to this operation UpdateDataSafePrivateEndpoint()
func (client DataSafeClient) UpdateDataSafePrivateEndpoint(ctx context.Context, request UpdateDataSafePrivateEndpointRequest) (response UpdateDataSafePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DataSafePrivateEndpoint/UpdateDataSafePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateDataSafePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseSecurityConfig Updates the database security configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateDatabaseSecurityConfig.go.html to see an example of how to use UpdateDatabaseSecurityConfig API.
// A default retry strategy applies to this operation UpdateDatabaseSecurityConfig()
func (client DataSafeClient) UpdateDatabaseSecurityConfig(ctx context.Context, request UpdateDatabaseSecurityConfigRequest) (response UpdateDatabaseSecurityConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseSecurityConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseSecurityConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseSecurityConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseSecurityConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseSecurityConfigResponse")
	}
	return
}

// updateDatabaseSecurityConfig implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateDatabaseSecurityConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseSecurityConfigs/{databaseSecurityConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseSecurityConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/DatabaseSecurityConfig/UpdateDatabaseSecurityConfig"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateDatabaseSecurityConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFinding Updates one or more attributes of the specified finding.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateFinding.go.html to see an example of how to use UpdateFinding API.
// A default retry strategy applies to this operation UpdateFinding()
func (client DataSafeClient) UpdateFinding(ctx context.Context, request UpdateFindingRequest) (response UpdateFindingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateFinding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFindingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFindingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFindingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFindingResponse")
	}
	return
}

// updateFinding implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateFinding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityAssessments/{securityAssessmentId}/findings/{findingKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFindingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Finding/UpdateFinding"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateFinding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLibraryMaskingFormat Updates one or more attributes of the specified library masking format. Note that updating the formatEntries attribute replaces all the existing masking format entries with the specified format entries.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateLibraryMaskingFormat.go.html to see an example of how to use UpdateLibraryMaskingFormat API.
// A default retry strategy applies to this operation UpdateLibraryMaskingFormat()
func (client DataSafeClient) UpdateLibraryMaskingFormat(ctx context.Context, request UpdateLibraryMaskingFormatRequest) (response UpdateLibraryMaskingFormatResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/LibraryMaskingFormat/UpdateLibraryMaskingFormat"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateLibraryMaskingFormat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMaskingColumn Updates one or more attributes of the specified masking column. Note that updating the maskingFormats
// attribute replaces the currently assigned masking formats with the specified masking formats.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateMaskingColumn.go.html to see an example of how to use UpdateMaskingColumn API.
// A default retry strategy applies to this operation UpdateMaskingColumn()
func (client DataSafeClient) UpdateMaskingColumn(ctx context.Context, request UpdateMaskingColumnRequest) (response UpdateMaskingColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingColumn/UpdateMaskingColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateMaskingColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMaskingPolicy Updates one or more attributes of the specified masking policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateMaskingPolicy.go.html to see an example of how to use UpdateMaskingPolicy API.
// A default retry strategy applies to this operation UpdateMaskingPolicy()
func (client DataSafeClient) UpdateMaskingPolicy(ctx context.Context, request UpdateMaskingPolicyRequest) (response UpdateMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/UpdateMaskingPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateMaskingPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOnPremConnector Updates one or more attributes of the specified on-premises connector.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateOnPremConnector.go.html to see an example of how to use UpdateOnPremConnector API.
// A default retry strategy applies to this operation UpdateOnPremConnector()
func (client DataSafeClient) UpdateOnPremConnector(ctx context.Context, request UpdateOnPremConnectorRequest) (response UpdateOnPremConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/UpdateOnPremConnector"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateOnPremConnector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOnPremConnectorWallet Updates the wallet for the specified on-premises connector to a new version.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateOnPremConnectorWallet.go.html to see an example of how to use UpdateOnPremConnectorWallet API.
// A default retry strategy applies to this operation UpdateOnPremConnectorWallet()
func (client DataSafeClient) UpdateOnPremConnectorWallet(ctx context.Context, request UpdateOnPremConnectorWalletRequest) (response UpdateOnPremConnectorWalletResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/OnPremConnector/UpdateOnPremConnectorWallet"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateOnPremConnectorWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePeerTargetDatabase Updates one or more attributes of the specified Data Safe peer target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdatePeerTargetDatabase.go.html to see an example of how to use UpdatePeerTargetDatabase API.
// A default retry strategy applies to this operation UpdatePeerTargetDatabase()
func (client DataSafeClient) UpdatePeerTargetDatabase(ctx context.Context, request UpdatePeerTargetDatabaseRequest) (response UpdatePeerTargetDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePeerTargetDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePeerTargetDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePeerTargetDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePeerTargetDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePeerTargetDatabaseResponse")
	}
	return
}

// updatePeerTargetDatabase implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updatePeerTargetDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targetDatabases/{targetDatabaseId}/peerTargetDatabases/{peerTargetDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePeerTargetDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/PeerTargetDatabase/UpdatePeerTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdatePeerTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateReport Updates the specified report. Only tags can be updated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateReport.go.html to see an example of how to use UpdateReport API.
// A default retry strategy applies to this operation UpdateReport()
func (client DataSafeClient) UpdateReport(ctx context.Context, request UpdateReportRequest) (response UpdateReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateReportResponse")
	}
	return
}

// updateReport implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/reports/{reportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/Report/UpdateReport"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateReportDefinition Updates the specified report definition. Only user created report definition can be updated. Seeded report definitions need to be saved as new report definition first.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateReportDefinition.go.html to see an example of how to use UpdateReportDefinition API.
// A default retry strategy applies to this operation UpdateReportDefinition()
func (client DataSafeClient) UpdateReportDefinition(ctx context.Context, request UpdateReportDefinitionRequest) (response UpdateReportDefinitionResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/ReportDefinition/UpdateReportDefinition"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateReportDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSdmMaskingPolicyDifference Updates one or more attributes of the specified sdm masking policy difference.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSdmMaskingPolicyDifference.go.html to see an example of how to use UpdateSdmMaskingPolicyDifference API.
// A default retry strategy applies to this operation UpdateSdmMaskingPolicyDifference()
func (client DataSafeClient) UpdateSdmMaskingPolicyDifference(ctx context.Context, request UpdateSdmMaskingPolicyDifferenceRequest) (response UpdateSdmMaskingPolicyDifferenceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSdmMaskingPolicyDifference, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSdmMaskingPolicyDifferenceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSdmMaskingPolicyDifferenceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSdmMaskingPolicyDifferenceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSdmMaskingPolicyDifferenceResponse")
	}
	return
}

// updateSdmMaskingPolicyDifference implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSdmMaskingPolicyDifference(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sdmMaskingPolicyDifferences/{sdmMaskingPolicyDifferenceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSdmMaskingPolicyDifferenceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SdmMaskingPolicyDifference/UpdateSdmMaskingPolicyDifference"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSdmMaskingPolicyDifference", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityAssessment Updates one or more attributes of the specified security assessment. This operation allows to update the security assessment displayName, description, or schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSecurityAssessment.go.html to see an example of how to use UpdateSecurityAssessment API.
// A default retry strategy applies to this operation UpdateSecurityAssessment()
func (client DataSafeClient) UpdateSecurityAssessment(ctx context.Context, request UpdateSecurityAssessmentRequest) (response UpdateSecurityAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityAssessment/UpdateSecurityAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSecurityAssessment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityPolicy Updates the security policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSecurityPolicy.go.html to see an example of how to use UpdateSecurityPolicy API.
// A default retry strategy applies to this operation UpdateSecurityPolicy()
func (client DataSafeClient) UpdateSecurityPolicy(ctx context.Context, request UpdateSecurityPolicyRequest) (response UpdateSecurityPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityPolicyResponse")
	}
	return
}

// updateSecurityPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSecurityPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityPolicies/{securityPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicy/UpdateSecurityPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSecurityPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityPolicyDeployment Updates the security policy deployment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSecurityPolicyDeployment.go.html to see an example of how to use UpdateSecurityPolicyDeployment API.
// A default retry strategy applies to this operation UpdateSecurityPolicyDeployment()
func (client DataSafeClient) UpdateSecurityPolicyDeployment(ctx context.Context, request UpdateSecurityPolicyDeploymentRequest) (response UpdateSecurityPolicyDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityPolicyDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityPolicyDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityPolicyDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityPolicyDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityPolicyDeploymentResponse")
	}
	return
}

// updateSecurityPolicyDeployment implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSecurityPolicyDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityPolicyDeployments/{securityPolicyDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityPolicyDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SecurityPolicyDeployment/UpdateSecurityPolicyDeployment"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSecurityPolicyDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSensitiveColumn Updates one or more attributes of the specified sensitive column.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSensitiveColumn.go.html to see an example of how to use UpdateSensitiveColumn API.
// A default retry strategy applies to this operation UpdateSensitiveColumn()
func (client DataSafeClient) UpdateSensitiveColumn(ctx context.Context, request UpdateSensitiveColumnRequest) (response UpdateSensitiveColumnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveColumn/UpdateSensitiveColumn"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSensitiveColumn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSensitiveDataModel Updates one or more attributes of the specified sensitive data model. Note that updating any attribute of a sensitive
// data model does not perform data discovery.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSensitiveDataModel.go.html to see an example of how to use UpdateSensitiveDataModel API.
// A default retry strategy applies to this operation UpdateSensitiveDataModel()
func (client DataSafeClient) UpdateSensitiveDataModel(ctx context.Context, request UpdateSensitiveDataModelRequest) (response UpdateSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/UpdateSensitiveDataModel"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSensitiveDataModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSensitiveType Updates one or more attributes of the specified sensitive type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSensitiveType.go.html to see an example of how to use UpdateSensitiveType API.
// A default retry strategy applies to this operation UpdateSensitiveType()
func (client DataSafeClient) UpdateSensitiveType(ctx context.Context, request UpdateSensitiveTypeRequest) (response UpdateSensitiveTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveType/UpdateSensitiveType"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSensitiveType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSqlCollection Updates the SQL collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSqlCollection.go.html to see an example of how to use UpdateSqlCollection API.
// A default retry strategy applies to this operation UpdateSqlCollection()
func (client DataSafeClient) UpdateSqlCollection(ctx context.Context, request UpdateSqlCollectionRequest) (response UpdateSqlCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSqlCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSqlCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSqlCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSqlCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSqlCollectionResponse")
	}
	return
}

// updateSqlCollection implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSqlCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sqlCollections/{sqlCollectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSqlCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlCollection/UpdateSqlCollection"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSqlCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSqlFirewallPolicy Updates the SQL Firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateSqlFirewallPolicy.go.html to see an example of how to use UpdateSqlFirewallPolicy API.
// A default retry strategy applies to this operation UpdateSqlFirewallPolicy()
func (client DataSafeClient) UpdateSqlFirewallPolicy(ctx context.Context, request UpdateSqlFirewallPolicyRequest) (response UpdateSqlFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSqlFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSqlFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSqlFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSqlFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSqlFirewallPolicyResponse")
	}
	return
}

// updateSqlFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client DataSafeClient) updateSqlFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sqlFirewallPolicies/{sqlFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSqlFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SqlFirewallPolicy/UpdateSqlFirewallPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateSqlFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetAlertPolicyAssociation Updates the specified target-alert policy association.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateTargetAlertPolicyAssociation.go.html to see an example of how to use UpdateTargetAlertPolicyAssociation API.
// A default retry strategy applies to this operation UpdateTargetAlertPolicyAssociation()
func (client DataSafeClient) UpdateTargetAlertPolicyAssociation(ctx context.Context, request UpdateTargetAlertPolicyAssociationRequest) (response UpdateTargetAlertPolicyAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetAlertPolicyAssociation/UpdateTargetAlertPolicyAssociation"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateTargetAlertPolicyAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetDatabase Updates one or more attributes of the specified Data Safe target database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateTargetDatabase.go.html to see an example of how to use UpdateTargetDatabase API.
// A default retry strategy applies to this operation UpdateTargetDatabase()
func (client DataSafeClient) UpdateTargetDatabase(ctx context.Context, request UpdateTargetDatabaseRequest) (response UpdateTargetDatabaseResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/TargetDatabase/UpdateTargetDatabase"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateTargetDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUserAssessment Updates one or more attributes of the specified user assessment. This operation allows to update the user assessment displayName, description, or schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UpdateUserAssessment.go.html to see an example of how to use UpdateUserAssessment API.
// A default retry strategy applies to this operation UpdateUserAssessment()
func (client DataSafeClient) UpdateUserAssessment(ctx context.Context, request UpdateUserAssessmentRequest) (response UpdateUserAssessmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/UserAssessment/UpdateUserAssessment"
		err = common.PostProcessServiceError(err, "DataSafe", "UpdateUserAssessment", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UploadMaskingPolicy.go.html to see an example of how to use UploadMaskingPolicy API.
// A default retry strategy applies to this operation UploadMaskingPolicy()
func (client DataSafeClient) UploadMaskingPolicy(ctx context.Context, request UploadMaskingPolicyRequest) (response UploadMaskingPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/MaskingPolicy/UploadMaskingPolicy"
		err = common.PostProcessServiceError(err, "DataSafe", "UploadMaskingPolicy", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/UploadSensitiveDataModel.go.html to see an example of how to use UploadSensitiveDataModel API.
// A default retry strategy applies to this operation UploadSensitiveDataModel()
func (client DataSafeClient) UploadSensitiveDataModel(ctx context.Context, request UploadSensitiveDataModelRequest) (response UploadSensitiveDataModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-safe/20181201/SensitiveDataModel/UploadSensitiveDataModel"
		err = common.PostProcessServiceError(err, "DataSafe", "UploadSensitiveDataModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
