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

// ListCollectedAuditVolumesRequest wrapper for the ListCollectedAuditVolumes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCollectedAuditVolumes.go.html to see an example of how to use ListCollectedAuditVolumesRequest.
type ListCollectedAuditVolumesRequest struct {

	// The OCID of the audit.
	AuditProfileId *string `mandatory:"true" contributesTo:"path" name:"auditProfileId"`

	// The OCID of the work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"query" name:"workRequestId"`

	// Specifying `monthInConsiderationGreaterThan` parameter
	// will retrieve all items for which the event month is
	// greater than the date and time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T00:00:00.000Z
	MonthInConsiderationGreaterThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"monthInConsiderationGreaterThan"`

	// Specifying `monthInConsiderationLessThan` parameter
	// will retrieve all items for which the event month is
	// less than the date and time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T00:00:00.000Z
	MonthInConsiderationLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"monthInConsiderationLessThan"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCollectedAuditVolumesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order(sortOrder). The default order for all fields is ascending.
	SortBy ListCollectedAuditVolumesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCollectedAuditVolumesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCollectedAuditVolumesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCollectedAuditVolumesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCollectedAuditVolumesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCollectedAuditVolumesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCollectedAuditVolumesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCollectedAuditVolumesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCollectedAuditVolumesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCollectedAuditVolumesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCollectedAuditVolumesResponse wrapper for the ListCollectedAuditVolumes operation
type ListCollectedAuditVolumesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CollectedAuditVolumeCollection instances
	CollectedAuditVolumeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCollectedAuditVolumesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCollectedAuditVolumesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCollectedAuditVolumesSortOrderEnum Enum with underlying type: string
type ListCollectedAuditVolumesSortOrderEnum string

// Set of constants representing the allowable values for ListCollectedAuditVolumesSortOrderEnum
const (
	ListCollectedAuditVolumesSortOrderAsc  ListCollectedAuditVolumesSortOrderEnum = "ASC"
	ListCollectedAuditVolumesSortOrderDesc ListCollectedAuditVolumesSortOrderEnum = "DESC"
)

var mappingListCollectedAuditVolumesSortOrderEnum = map[string]ListCollectedAuditVolumesSortOrderEnum{
	"ASC":  ListCollectedAuditVolumesSortOrderAsc,
	"DESC": ListCollectedAuditVolumesSortOrderDesc,
}

var mappingListCollectedAuditVolumesSortOrderEnumLowerCase = map[string]ListCollectedAuditVolumesSortOrderEnum{
	"asc":  ListCollectedAuditVolumesSortOrderAsc,
	"desc": ListCollectedAuditVolumesSortOrderDesc,
}

// GetListCollectedAuditVolumesSortOrderEnumValues Enumerates the set of values for ListCollectedAuditVolumesSortOrderEnum
func GetListCollectedAuditVolumesSortOrderEnumValues() []ListCollectedAuditVolumesSortOrderEnum {
	values := make([]ListCollectedAuditVolumesSortOrderEnum, 0)
	for _, v := range mappingListCollectedAuditVolumesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectedAuditVolumesSortOrderEnumStringValues Enumerates the set of values in String for ListCollectedAuditVolumesSortOrderEnum
func GetListCollectedAuditVolumesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCollectedAuditVolumesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectedAuditVolumesSortOrderEnum(val string) (ListCollectedAuditVolumesSortOrderEnum, bool) {
	enum, ok := mappingListCollectedAuditVolumesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCollectedAuditVolumesSortByEnum Enum with underlying type: string
type ListCollectedAuditVolumesSortByEnum string

// Set of constants representing the allowable values for ListCollectedAuditVolumesSortByEnum
const (
	ListCollectedAuditVolumesSortByMonthinconsideration ListCollectedAuditVolumesSortByEnum = "monthInConsideration"
	ListCollectedAuditVolumesSortByOnlinevolume         ListCollectedAuditVolumesSortByEnum = "onlineVolume"
	ListCollectedAuditVolumesSortByArchivedvolume       ListCollectedAuditVolumesSortByEnum = "archivedVolume"
)

var mappingListCollectedAuditVolumesSortByEnum = map[string]ListCollectedAuditVolumesSortByEnum{
	"monthInConsideration": ListCollectedAuditVolumesSortByMonthinconsideration,
	"onlineVolume":         ListCollectedAuditVolumesSortByOnlinevolume,
	"archivedVolume":       ListCollectedAuditVolumesSortByArchivedvolume,
}

var mappingListCollectedAuditVolumesSortByEnumLowerCase = map[string]ListCollectedAuditVolumesSortByEnum{
	"monthinconsideration": ListCollectedAuditVolumesSortByMonthinconsideration,
	"onlinevolume":         ListCollectedAuditVolumesSortByOnlinevolume,
	"archivedvolume":       ListCollectedAuditVolumesSortByArchivedvolume,
}

// GetListCollectedAuditVolumesSortByEnumValues Enumerates the set of values for ListCollectedAuditVolumesSortByEnum
func GetListCollectedAuditVolumesSortByEnumValues() []ListCollectedAuditVolumesSortByEnum {
	values := make([]ListCollectedAuditVolumesSortByEnum, 0)
	for _, v := range mappingListCollectedAuditVolumesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCollectedAuditVolumesSortByEnumStringValues Enumerates the set of values in String for ListCollectedAuditVolumesSortByEnum
func GetListCollectedAuditVolumesSortByEnumStringValues() []string {
	return []string{
		"monthInConsideration",
		"onlineVolume",
		"archivedVolume",
	}
}

// GetMappingListCollectedAuditVolumesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCollectedAuditVolumesSortByEnum(val string) (ListCollectedAuditVolumesSortByEnum, bool) {
	enum, ok := mappingListCollectedAuditVolumesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
