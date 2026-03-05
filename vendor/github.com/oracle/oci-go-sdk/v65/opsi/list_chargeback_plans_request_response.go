// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListChargebackPlansRequest wrapper for the ListChargebackPlans operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListChargebackPlans.go.html to see an example of how to use ListChargebackPlansRequest.
type ListChargebackPlansRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Ops Insights chargeback plan.
	ChargebackplanId *string `mandatory:"false" contributesTo:"query" name:"chargebackplanId"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListChargebackPlansSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort chargeback plans.
	SortBy ListChargebackPlansSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChargebackPlansRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChargebackPlansRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChargebackPlansRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChargebackPlansRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListChargebackPlansRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListChargebackPlansSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListChargebackPlansSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChargebackPlansSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListChargebackPlansSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListChargebackPlansResponse wrapper for the ListChargebackPlans operation
type ListChargebackPlansResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ChargebackPlanCollection instances
	ChargebackPlanCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListChargebackPlansResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChargebackPlansResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChargebackPlansSortOrderEnum Enum with underlying type: string
type ListChargebackPlansSortOrderEnum string

// Set of constants representing the allowable values for ListChargebackPlansSortOrderEnum
const (
	ListChargebackPlansSortOrderAsc  ListChargebackPlansSortOrderEnum = "ASC"
	ListChargebackPlansSortOrderDesc ListChargebackPlansSortOrderEnum = "DESC"
)

var mappingListChargebackPlansSortOrderEnum = map[string]ListChargebackPlansSortOrderEnum{
	"ASC":  ListChargebackPlansSortOrderAsc,
	"DESC": ListChargebackPlansSortOrderDesc,
}

var mappingListChargebackPlansSortOrderEnumLowerCase = map[string]ListChargebackPlansSortOrderEnum{
	"asc":  ListChargebackPlansSortOrderAsc,
	"desc": ListChargebackPlansSortOrderDesc,
}

// GetListChargebackPlansSortOrderEnumValues Enumerates the set of values for ListChargebackPlansSortOrderEnum
func GetListChargebackPlansSortOrderEnumValues() []ListChargebackPlansSortOrderEnum {
	values := make([]ListChargebackPlansSortOrderEnum, 0)
	for _, v := range mappingListChargebackPlansSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListChargebackPlansSortOrderEnumStringValues Enumerates the set of values in String for ListChargebackPlansSortOrderEnum
func GetListChargebackPlansSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListChargebackPlansSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChargebackPlansSortOrderEnum(val string) (ListChargebackPlansSortOrderEnum, bool) {
	enum, ok := mappingListChargebackPlansSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChargebackPlansSortByEnum Enum with underlying type: string
type ListChargebackPlansSortByEnum string

// Set of constants representing the allowable values for ListChargebackPlansSortByEnum
const (
	ListChargebackPlansSortByTimecreated    ListChargebackPlansSortByEnum = "timeCreated"
	ListChargebackPlansSortById             ListChargebackPlansSortByEnum = "id"
	ListChargebackPlansSortByLifecyclestate ListChargebackPlansSortByEnum = "lifecycleState"
)

var mappingListChargebackPlansSortByEnum = map[string]ListChargebackPlansSortByEnum{
	"timeCreated":    ListChargebackPlansSortByTimecreated,
	"id":             ListChargebackPlansSortById,
	"lifecycleState": ListChargebackPlansSortByLifecyclestate,
}

var mappingListChargebackPlansSortByEnumLowerCase = map[string]ListChargebackPlansSortByEnum{
	"timecreated":    ListChargebackPlansSortByTimecreated,
	"id":             ListChargebackPlansSortById,
	"lifecyclestate": ListChargebackPlansSortByLifecyclestate,
}

// GetListChargebackPlansSortByEnumValues Enumerates the set of values for ListChargebackPlansSortByEnum
func GetListChargebackPlansSortByEnumValues() []ListChargebackPlansSortByEnum {
	values := make([]ListChargebackPlansSortByEnum, 0)
	for _, v := range mappingListChargebackPlansSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListChargebackPlansSortByEnumStringValues Enumerates the set of values in String for ListChargebackPlansSortByEnum
func GetListChargebackPlansSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"id",
		"lifecycleState",
	}
}

// GetMappingListChargebackPlansSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChargebackPlansSortByEnum(val string) (ListChargebackPlansSortByEnum, bool) {
	enum, ok := mappingListChargebackPlansSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
