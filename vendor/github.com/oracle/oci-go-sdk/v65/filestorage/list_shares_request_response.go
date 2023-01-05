// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSharesRequest wrapper for the ListShares operation
type ListSharesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A user-friendly share name. It does not have to be unique, and it is changeable.
	// Example: `My share`
	ShareName *string `mandatory:"false" contributesTo:"query" name:"shareName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the share set.
	ShareSetId *string `mandatory:"false" contributesTo:"query" name:"shareSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"false" contributesTo:"query" name:"fileSystemId"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type shares.
	LifecycleState ShareLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by share name, results are
	// shown in ascending alphanumeric order.
	SortBy ListSharesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListSharesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSharesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSharesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSharesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSharesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSharesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShareLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetShareLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSharesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSharesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSharesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSharesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSharesResponse wrapper for the ListShares operation
type ListSharesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ShareSummary instances
	Items []ShareSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSharesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSharesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSharesSortByEnum Enum with underlying type: string
type ListSharesSortByEnum string

// Set of constants representing the allowable values for ListSharesSortByEnum
const (
	ListSharesSortByTimecreated ListSharesSortByEnum = "TIMECREATED"
	ListSharesSortBySharename   ListSharesSortByEnum = "SHARENAME"
)

var mappingListSharesSortByEnum = map[string]ListSharesSortByEnum{
	"TIMECREATED": ListSharesSortByTimecreated,
	"SHARENAME":   ListSharesSortBySharename,
}

var mappingListSharesSortByEnumLowerCase = map[string]ListSharesSortByEnum{
	"timecreated": ListSharesSortByTimecreated,
	"sharename":   ListSharesSortBySharename,
}

// GetListSharesSortByEnumValues Enumerates the set of values for ListSharesSortByEnum
func GetListSharesSortByEnumValues() []ListSharesSortByEnum {
	values := make([]ListSharesSortByEnum, 0)
	for _, v := range mappingListSharesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSharesSortByEnumStringValues Enumerates the set of values in String for ListSharesSortByEnum
func GetListSharesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"SHARENAME",
	}
}

// GetMappingListSharesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSharesSortByEnum(val string) (ListSharesSortByEnum, bool) {
	enum, ok := mappingListSharesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSharesSortOrderEnum Enum with underlying type: string
type ListSharesSortOrderEnum string

// Set of constants representing the allowable values for ListSharesSortOrderEnum
const (
	ListSharesSortOrderAsc  ListSharesSortOrderEnum = "ASC"
	ListSharesSortOrderDesc ListSharesSortOrderEnum = "DESC"
)

var mappingListSharesSortOrderEnum = map[string]ListSharesSortOrderEnum{
	"ASC":  ListSharesSortOrderAsc,
	"DESC": ListSharesSortOrderDesc,
}

var mappingListSharesSortOrderEnumLowerCase = map[string]ListSharesSortOrderEnum{
	"asc":  ListSharesSortOrderAsc,
	"desc": ListSharesSortOrderDesc,
}

// GetListSharesSortOrderEnumValues Enumerates the set of values for ListSharesSortOrderEnum
func GetListSharesSortOrderEnumValues() []ListSharesSortOrderEnum {
	values := make([]ListSharesSortOrderEnum, 0)
	for _, v := range mappingListSharesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSharesSortOrderEnumStringValues Enumerates the set of values in String for ListSharesSortOrderEnum
func GetListSharesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSharesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSharesSortOrderEnum(val string) (ListSharesSortOrderEnum, bool) {
	enum, ok := mappingListSharesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
