// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListParamsForActionTypeRequest wrapper for the ListParamsForActionType operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListParamsForActionType.go.html to see an example of how to use ListParamsForActionTypeRequest.
type ListParamsForActionTypeRequest struct {

	// The type of the scheduled action
	Type RecommendedScheduledActionSummaryActionTypeEnum `mandatory:"true" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The plan intent the action will be used for. Relevant to action type that can be used in multiple plans
	PlanIntent ListParamsForActionTypePlanIntentEnum `mandatory:"false" contributesTo:"query" name:"planIntent" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParamsForActionTypeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParamsForActionTypeRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListParamsForActionTypeRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParamsForActionTypeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListParamsForActionTypeRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecommendedScheduledActionSummaryActionTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetRecommendedScheduledActionSummaryActionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParamsForActionTypePlanIntentEnum(string(request.PlanIntent)); !ok && request.PlanIntent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanIntent: %s. Supported values are: %s.", request.PlanIntent, strings.Join(GetListParamsForActionTypePlanIntentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListParamsForActionTypeResponse wrapper for the ListParamsForActionType operation
type ListParamsForActionTypeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ActionParamValuesCollection instances
	ActionParamValuesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListParamsForActionTypeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParamsForActionTypeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParamsForActionTypePlanIntentEnum Enum with underlying type: string
type ListParamsForActionTypePlanIntentEnum string

// Set of constants representing the allowable values for ListParamsForActionTypePlanIntentEnum
const (
	ListParamsForActionTypePlanIntentFullSoftwareUpdate ListParamsForActionTypePlanIntentEnum = "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"
	ListParamsForActionTypePlanIntentSecurityUpdate     ListParamsForActionTypePlanIntentEnum = "EXADATA_INFRASTRUCTURE_SECURITY_UPDATE"
)

var mappingListParamsForActionTypePlanIntentEnum = map[string]ListParamsForActionTypePlanIntentEnum{
	"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE": ListParamsForActionTypePlanIntentFullSoftwareUpdate,
	"EXADATA_INFRASTRUCTURE_SECURITY_UPDATE":      ListParamsForActionTypePlanIntentSecurityUpdate,
}

var mappingListParamsForActionTypePlanIntentEnumLowerCase = map[string]ListParamsForActionTypePlanIntentEnum{
	"exadata_infrastructure_full_software_update": ListParamsForActionTypePlanIntentFullSoftwareUpdate,
	"exadata_infrastructure_security_update":      ListParamsForActionTypePlanIntentSecurityUpdate,
}

// GetListParamsForActionTypePlanIntentEnumValues Enumerates the set of values for ListParamsForActionTypePlanIntentEnum
func GetListParamsForActionTypePlanIntentEnumValues() []ListParamsForActionTypePlanIntentEnum {
	values := make([]ListParamsForActionTypePlanIntentEnum, 0)
	for _, v := range mappingListParamsForActionTypePlanIntentEnum {
		values = append(values, v)
	}
	return values
}

// GetListParamsForActionTypePlanIntentEnumStringValues Enumerates the set of values in String for ListParamsForActionTypePlanIntentEnum
func GetListParamsForActionTypePlanIntentEnumStringValues() []string {
	return []string{
		"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE",
		"EXADATA_INFRASTRUCTURE_SECURITY_UPDATE",
	}
}

// GetMappingListParamsForActionTypePlanIntentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParamsForActionTypePlanIntentEnum(val string) (ListParamsForActionTypePlanIntentEnum, bool) {
	enum, ok := mappingListParamsForActionTypePlanIntentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
