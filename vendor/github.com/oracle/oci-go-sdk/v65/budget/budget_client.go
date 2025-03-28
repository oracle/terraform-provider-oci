// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// BudgetClient a client for Budget
type BudgetClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBudgetClientWithConfigurationProvider Creates a new default Budget client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBudgetClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BudgetClient, err error) {
	if enabled := common.CheckForEnabledServices("budget"); !enabled {
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
	return newBudgetClientFromBaseClient(baseClient, provider)
}

// NewBudgetClientWithOboToken Creates a new default Budget client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewBudgetClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BudgetClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newBudgetClientFromBaseClient(baseClient, configProvider)
}

func newBudgetClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BudgetClient, err error) {
	// Budget service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Budget"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = BudgetClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BudgetClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("budget", "https://usage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BudgetClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BudgetClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAlertRule Creates a new Alert Rule.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/CreateAlertRule.go.html to see an example of how to use CreateAlertRule API.
// A default retry strategy applies to this operation CreateAlertRule()
func (client BudgetClient) CreateAlertRule(ctx context.Context, request CreateAlertRuleRequest) (response CreateAlertRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAlertRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAlertRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAlertRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAlertRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAlertRuleResponse")
	}
	return
}

// createAlertRule implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) createAlertRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/budgets/{budgetId}/alertRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAlertRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/AlertRule/CreateAlertRule"
		err = common.PostProcessServiceError(err, "Budget", "CreateAlertRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBudget Creates a new budget.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/CreateBudget.go.html to see an example of how to use CreateBudget API.
// A default retry strategy applies to this operation CreateBudget()
func (client BudgetClient) CreateBudget(ctx context.Context, request CreateBudgetRequest) (response CreateBudgetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBudget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBudgetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBudgetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBudgetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBudgetResponse")
	}
	return
}

// createBudget implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) createBudget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/budgets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBudgetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/Budget/CreateBudget"
		err = common.PostProcessServiceError(err, "Budget", "CreateBudget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAlertRule Deletes a specified Alert Rule resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/DeleteAlertRule.go.html to see an example of how to use DeleteAlertRule API.
// A default retry strategy applies to this operation DeleteAlertRule()
func (client BudgetClient) DeleteAlertRule(ctx context.Context, request DeleteAlertRuleRequest) (response DeleteAlertRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAlertRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAlertRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAlertRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAlertRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAlertRuleResponse")
	}
	return
}

// deleteAlertRule implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) deleteAlertRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/budgets/{budgetId}/alertRules/{alertRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAlertRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/AlertRule/DeleteAlertRule"
		err = common.PostProcessServiceError(err, "Budget", "DeleteAlertRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBudget Deletes a specified budget resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/DeleteBudget.go.html to see an example of how to use DeleteBudget API.
// A default retry strategy applies to this operation DeleteBudget()
func (client BudgetClient) DeleteBudget(ctx context.Context, request DeleteBudgetRequest) (response DeleteBudgetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBudget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBudgetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBudgetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBudgetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBudgetResponse")
	}
	return
}

// deleteBudget implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) deleteBudget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/budgets/{budgetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBudgetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/Budget/DeleteBudget"
		err = common.PostProcessServiceError(err, "Budget", "DeleteBudget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlertRule Gets an Alert Rule for a specified budget.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/GetAlertRule.go.html to see an example of how to use GetAlertRule API.
// A default retry strategy applies to this operation GetAlertRule()
func (client BudgetClient) GetAlertRule(ctx context.Context, request GetAlertRuleRequest) (response GetAlertRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlertRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAlertRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAlertRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlertRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlertRuleResponse")
	}
	return
}

// getAlertRule implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) getAlertRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/budgets/{budgetId}/alertRules/{alertRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAlertRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/AlertRule/GetAlertRule"
		err = common.PostProcessServiceError(err, "Budget", "GetAlertRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBudget Gets a budget by the identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/GetBudget.go.html to see an example of how to use GetBudget API.
// A default retry strategy applies to this operation GetBudget()
func (client BudgetClient) GetBudget(ctx context.Context, request GetBudgetRequest) (response GetBudgetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBudget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBudgetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBudgetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBudgetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBudgetResponse")
	}
	return
}

// getBudget implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) getBudget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/budgets/{budgetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBudgetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/Budget/GetBudget"
		err = common.PostProcessServiceError(err, "Budget", "GetBudget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlertRules Returns a list of Alert Rules for a specified budget.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListAlertRules.go.html to see an example of how to use ListAlertRules API.
// A default retry strategy applies to this operation ListAlertRules()
func (client BudgetClient) ListAlertRules(ctx context.Context, request ListAlertRulesRequest) (response ListAlertRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlertRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlertRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlertRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlertRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlertRulesResponse")
	}
	return
}

// listAlertRules implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) listAlertRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/budgets/{budgetId}/alertRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlertRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/AlertRuleSummary/ListAlertRules"
		err = common.PostProcessServiceError(err, "Budget", "ListAlertRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBudgets Gets a list of budgets in a compartment.
// By default, ListBudgets returns budgets of the 'COMPARTMENT' target type, and the budget records with only one target compartment OCID.
// To list all budgets, set the targetType query parameter to ALL (for example: 'targetType=ALL').
// Clients should ignore new targetTypes, or upgrade to the latest version of the client SDK to handle new targetTypes.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListBudgets.go.html to see an example of how to use ListBudgets API.
// A default retry strategy applies to this operation ListBudgets()
func (client BudgetClient) ListBudgets(ctx context.Context, request ListBudgetsRequest) (response ListBudgetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBudgets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBudgetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBudgetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBudgetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBudgetsResponse")
	}
	return
}

// listBudgets implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) listBudgets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/budgets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBudgetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/BudgetSummary/ListBudgets"
		err = common.PostProcessServiceError(err, "Budget", "ListBudgets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAlertRule Update an Alert Rule for the budget identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/UpdateAlertRule.go.html to see an example of how to use UpdateAlertRule API.
// A default retry strategy applies to this operation UpdateAlertRule()
func (client BudgetClient) UpdateAlertRule(ctx context.Context, request UpdateAlertRuleRequest) (response UpdateAlertRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAlertRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAlertRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAlertRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAlertRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAlertRuleResponse")
	}
	return
}

// updateAlertRule implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) updateAlertRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/budgets/{budgetId}/alertRules/{alertRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAlertRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/AlertRule/UpdateAlertRule"
		err = common.PostProcessServiceError(err, "Budget", "UpdateAlertRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBudget Update a budget identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/UpdateBudget.go.html to see an example of how to use UpdateBudget API.
// A default retry strategy applies to this operation UpdateBudget()
func (client BudgetClient) UpdateBudget(ctx context.Context, request UpdateBudgetRequest) (response UpdateBudgetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBudget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBudgetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBudgetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBudgetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBudgetResponse")
	}
	return
}

// updateBudget implements the OCIOperation interface (enables retrying operations)
func (client BudgetClient) updateBudget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/budgets/{budgetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBudgetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/Budget/UpdateBudget"
		err = common.PostProcessServiceError(err, "Budget", "UpdateBudget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
