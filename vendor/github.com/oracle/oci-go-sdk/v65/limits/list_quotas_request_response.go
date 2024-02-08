// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListQuotasRequest wrapper for the ListQuotas operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limits/ListQuotas.go.html to see an example of how to use ListQuotasRequest.
type ListQuotasRequest struct {

	// The OCID of the parent compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filters returned quotas based on the given state.
	LifecycleState ListQuotasLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'. By default, it is ascending.
	SortOrder ListQuotasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListQuotasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQuotasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQuotasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQuotasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQuotasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListQuotasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListQuotasLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListQuotasLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQuotasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListQuotasSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQuotasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListQuotasSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListQuotasResponse wrapper for the ListQuotas operation
type ListQuotasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []QuotaSummary instances
	Items []QuotaSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListQuotasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQuotasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQuotasLifecycleStateEnum Enum with underlying type: string
type ListQuotasLifecycleStateEnum string

// Set of constants representing the allowable values for ListQuotasLifecycleStateEnum
const (
	ListQuotasLifecycleStateActive ListQuotasLifecycleStateEnum = "ACTIVE"
)

var mappingListQuotasLifecycleStateEnum = map[string]ListQuotasLifecycleStateEnum{
	"ACTIVE": ListQuotasLifecycleStateActive,
}

var mappingListQuotasLifecycleStateEnumLowerCase = map[string]ListQuotasLifecycleStateEnum{
	"active": ListQuotasLifecycleStateActive,
}

// GetListQuotasLifecycleStateEnumValues Enumerates the set of values for ListQuotasLifecycleStateEnum
func GetListQuotasLifecycleStateEnumValues() []ListQuotasLifecycleStateEnum {
	values := make([]ListQuotasLifecycleStateEnum, 0)
	for _, v := range mappingListQuotasLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuotasLifecycleStateEnumStringValues Enumerates the set of values in String for ListQuotasLifecycleStateEnum
func GetListQuotasLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingListQuotasLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuotasLifecycleStateEnum(val string) (ListQuotasLifecycleStateEnum, bool) {
	enum, ok := mappingListQuotasLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListQuotasSortOrderEnum Enum with underlying type: string
type ListQuotasSortOrderEnum string

// Set of constants representing the allowable values for ListQuotasSortOrderEnum
const (
	ListQuotasSortOrderAsc  ListQuotasSortOrderEnum = "ASC"
	ListQuotasSortOrderDesc ListQuotasSortOrderEnum = "DESC"
)

var mappingListQuotasSortOrderEnum = map[string]ListQuotasSortOrderEnum{
	"ASC":  ListQuotasSortOrderAsc,
	"DESC": ListQuotasSortOrderDesc,
}

var mappingListQuotasSortOrderEnumLowerCase = map[string]ListQuotasSortOrderEnum{
	"asc":  ListQuotasSortOrderAsc,
	"desc": ListQuotasSortOrderDesc,
}

// GetListQuotasSortOrderEnumValues Enumerates the set of values for ListQuotasSortOrderEnum
func GetListQuotasSortOrderEnumValues() []ListQuotasSortOrderEnum {
	values := make([]ListQuotasSortOrderEnum, 0)
	for _, v := range mappingListQuotasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuotasSortOrderEnumStringValues Enumerates the set of values in String for ListQuotasSortOrderEnum
func GetListQuotasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListQuotasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuotasSortOrderEnum(val string) (ListQuotasSortOrderEnum, bool) {
	enum, ok := mappingListQuotasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListQuotasSortByEnum Enum with underlying type: string
type ListQuotasSortByEnum string

// Set of constants representing the allowable values for ListQuotasSortByEnum
const (
	ListQuotasSortByName        ListQuotasSortByEnum = "NAME"
	ListQuotasSortByTimecreated ListQuotasSortByEnum = "TIMECREATED"
)

var mappingListQuotasSortByEnum = map[string]ListQuotasSortByEnum{
	"NAME":        ListQuotasSortByName,
	"TIMECREATED": ListQuotasSortByTimecreated,
}

var mappingListQuotasSortByEnumLowerCase = map[string]ListQuotasSortByEnum{
	"name":        ListQuotasSortByName,
	"timecreated": ListQuotasSortByTimecreated,
}

// GetListQuotasSortByEnumValues Enumerates the set of values for ListQuotasSortByEnum
func GetListQuotasSortByEnumValues() []ListQuotasSortByEnum {
	values := make([]ListQuotasSortByEnum, 0)
	for _, v := range mappingListQuotasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuotasSortByEnumStringValues Enumerates the set of values in String for ListQuotasSortByEnum
func GetListQuotasSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListQuotasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuotasSortByEnum(val string) (ListQuotasSortByEnum, bool) {
	enum, ok := mappingListQuotasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
