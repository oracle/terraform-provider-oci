// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailDeliveryConfigsRequest wrapper for the ListEmailDeliveryConfigs operation
type ListEmailDeliveryConfigsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailDeliveryConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailDeliveryConfigLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the attribute with which to sort the return paths.
	// Default: `timeCreated`
	// * **timeCreated:** Sorts by timeCreated.
	// * **name:** Sorts by name.
	SortBy ListEmailDeliveryConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailDeliveryConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailDeliveryConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailDeliveryConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailDeliveryConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailDeliveryConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailDeliveryConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailDeliveryConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailDeliveryConfigLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailDeliveryConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailDeliveryConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailDeliveryConfigsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailDeliveryConfigsResponse wrapper for the ListEmailDeliveryConfigs operation
type ListEmailDeliveryConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailDeliveryConfigCollection instances
	EmailDeliveryConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListEmailDeliveryConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailDeliveryConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailDeliveryConfigsSortOrderEnum Enum with underlying type: string
type ListEmailDeliveryConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailDeliveryConfigsSortOrderEnum
const (
	ListEmailDeliveryConfigsSortOrderAsc  ListEmailDeliveryConfigsSortOrderEnum = "ASC"
	ListEmailDeliveryConfigsSortOrderDesc ListEmailDeliveryConfigsSortOrderEnum = "DESC"
)

var mappingListEmailDeliveryConfigsSortOrderEnum = map[string]ListEmailDeliveryConfigsSortOrderEnum{
	"ASC":  ListEmailDeliveryConfigsSortOrderAsc,
	"DESC": ListEmailDeliveryConfigsSortOrderDesc,
}

var mappingListEmailDeliveryConfigsSortOrderEnumLowerCase = map[string]ListEmailDeliveryConfigsSortOrderEnum{
	"asc":  ListEmailDeliveryConfigsSortOrderAsc,
	"desc": ListEmailDeliveryConfigsSortOrderDesc,
}

// GetListEmailDeliveryConfigsSortOrderEnumValues Enumerates the set of values for ListEmailDeliveryConfigsSortOrderEnum
func GetListEmailDeliveryConfigsSortOrderEnumValues() []ListEmailDeliveryConfigsSortOrderEnum {
	values := make([]ListEmailDeliveryConfigsSortOrderEnum, 0)
	for _, v := range mappingListEmailDeliveryConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailDeliveryConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailDeliveryConfigsSortOrderEnum
func GetListEmailDeliveryConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailDeliveryConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailDeliveryConfigsSortOrderEnum(val string) (ListEmailDeliveryConfigsSortOrderEnum, bool) {
	enum, ok := mappingListEmailDeliveryConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailDeliveryConfigsSortByEnum Enum with underlying type: string
type ListEmailDeliveryConfigsSortByEnum string

// Set of constants representing the allowable values for ListEmailDeliveryConfigsSortByEnum
const (
	ListEmailDeliveryConfigsSortByTimecreated ListEmailDeliveryConfigsSortByEnum = "timeCreated"
	ListEmailDeliveryConfigsSortByName        ListEmailDeliveryConfigsSortByEnum = "name"
)

var mappingListEmailDeliveryConfigsSortByEnum = map[string]ListEmailDeliveryConfigsSortByEnum{
	"timeCreated": ListEmailDeliveryConfigsSortByTimecreated,
	"name":        ListEmailDeliveryConfigsSortByName,
}

var mappingListEmailDeliveryConfigsSortByEnumLowerCase = map[string]ListEmailDeliveryConfigsSortByEnum{
	"timecreated": ListEmailDeliveryConfigsSortByTimecreated,
	"name":        ListEmailDeliveryConfigsSortByName,
}

// GetListEmailDeliveryConfigsSortByEnumValues Enumerates the set of values for ListEmailDeliveryConfigsSortByEnum
func GetListEmailDeliveryConfigsSortByEnumValues() []ListEmailDeliveryConfigsSortByEnum {
	values := make([]ListEmailDeliveryConfigsSortByEnum, 0)
	for _, v := range mappingListEmailDeliveryConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailDeliveryConfigsSortByEnumStringValues Enumerates the set of values in String for ListEmailDeliveryConfigsSortByEnum
func GetListEmailDeliveryConfigsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListEmailDeliveryConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailDeliveryConfigsSortByEnum(val string) (ListEmailDeliveryConfigsSortByEnum, bool) {
	enum, ok := mappingListEmailDeliveryConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
