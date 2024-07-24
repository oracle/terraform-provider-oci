// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetDatabaseGroupAuditProfileDeviationsRequest wrapper for the ListTargetDatabaseGroupAuditProfileDeviations operation
type ListTargetDatabaseGroupAuditProfileDeviationsRequest struct {

	// The OCID of the audit.
	TargetDatabaseGroupAuditProfileId *string `mandatory:"true" contributesTo:"path" name:"targetDatabaseGroupAuditProfileId"`

	// The OCID of the work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"query" name:"workRequestId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetDatabaseGroupAuditProfileDeviationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetDatabaseGroupAuditProfileDeviationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetDatabaseGroupAuditProfileDeviationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetDatabaseGroupAuditProfileDeviationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetDatabaseGroupAuditProfileDeviationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetDatabaseGroupAuditProfileDeviationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetDatabaseGroupAuditProfileDeviationsResponse wrapper for the ListTargetDatabaseGroupAuditProfileDeviations operation
type ListTargetDatabaseGroupAuditProfileDeviationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetDeviationCollection instances
	TargetDeviationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetDatabaseGroupAuditProfileDeviationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetDatabaseGroupAuditProfileDeviationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum Enum with underlying type: string
type ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum
const (
	ListTargetDatabaseGroupAuditProfileDeviationsSortByTimecreated ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum = "TIMECREATED"
	ListTargetDatabaseGroupAuditProfileDeviationsSortByDisplayname ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum = "DISPLAYNAME"
)

var mappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnum = map[string]ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum{
	"TIMECREATED": ListTargetDatabaseGroupAuditProfileDeviationsSortByTimecreated,
	"DISPLAYNAME": ListTargetDatabaseGroupAuditProfileDeviationsSortByDisplayname,
}

var mappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnumLowerCase = map[string]ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum{
	"timecreated": ListTargetDatabaseGroupAuditProfileDeviationsSortByTimecreated,
	"displayname": ListTargetDatabaseGroupAuditProfileDeviationsSortByDisplayname,
}

// GetListTargetDatabaseGroupAuditProfileDeviationsSortByEnumValues Enumerates the set of values for ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum
func GetListTargetDatabaseGroupAuditProfileDeviationsSortByEnumValues() []ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum {
	values := make([]ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupAuditProfileDeviationsSortByEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum
func GetListTargetDatabaseGroupAuditProfileDeviationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnum(val string) (ListTargetDatabaseGroupAuditProfileDeviationsSortByEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupAuditProfileDeviationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum Enum with underlying type: string
type ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum
const (
	ListTargetDatabaseGroupAuditProfileDeviationsSortOrderAsc  ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum = "ASC"
	ListTargetDatabaseGroupAuditProfileDeviationsSortOrderDesc ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum = "DESC"
)

var mappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum = map[string]ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum{
	"ASC":  ListTargetDatabaseGroupAuditProfileDeviationsSortOrderAsc,
	"DESC": ListTargetDatabaseGroupAuditProfileDeviationsSortOrderDesc,
}

var mappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumLowerCase = map[string]ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum{
	"asc":  ListTargetDatabaseGroupAuditProfileDeviationsSortOrderAsc,
	"desc": ListTargetDatabaseGroupAuditProfileDeviationsSortOrderDesc,
}

// GetListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumValues Enumerates the set of values for ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum
func GetListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumValues() []ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum {
	values := make([]ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum
func GetListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum(val string) (ListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupAuditProfileDeviationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
