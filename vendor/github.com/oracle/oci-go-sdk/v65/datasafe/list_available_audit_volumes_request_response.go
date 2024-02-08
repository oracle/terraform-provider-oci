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

// ListAvailableAuditVolumesRequest wrapper for the ListAvailableAuditVolumes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAvailableAuditVolumes.go.html to see an example of how to use ListAvailableAuditVolumesRequest.
type ListAvailableAuditVolumesRequest struct {

	// The OCID of the audit.
	AuditProfileId *string `mandatory:"true" contributesTo:"path" name:"auditProfileId"`

	// The OCID of the work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"query" name:"workRequestId"`

	// The audit trail location.
	TrailLocation *string `mandatory:"false" contributesTo:"query" name:"trailLocation"`

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
	SortOrder ListAvailableAuditVolumesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order(sortOrder). The default order for all fields is ascending.
	SortBy ListAvailableAuditVolumesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableAuditVolumesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableAuditVolumesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableAuditVolumesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableAuditVolumesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableAuditVolumesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableAuditVolumesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableAuditVolumesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableAuditVolumesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableAuditVolumesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableAuditVolumesResponse wrapper for the ListAvailableAuditVolumes operation
type ListAvailableAuditVolumesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableAuditVolumeCollection instances
	AvailableAuditVolumeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAvailableAuditVolumesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableAuditVolumesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableAuditVolumesSortOrderEnum Enum with underlying type: string
type ListAvailableAuditVolumesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableAuditVolumesSortOrderEnum
const (
	ListAvailableAuditVolumesSortOrderAsc  ListAvailableAuditVolumesSortOrderEnum = "ASC"
	ListAvailableAuditVolumesSortOrderDesc ListAvailableAuditVolumesSortOrderEnum = "DESC"
)

var mappingListAvailableAuditVolumesSortOrderEnum = map[string]ListAvailableAuditVolumesSortOrderEnum{
	"ASC":  ListAvailableAuditVolumesSortOrderAsc,
	"DESC": ListAvailableAuditVolumesSortOrderDesc,
}

var mappingListAvailableAuditVolumesSortOrderEnumLowerCase = map[string]ListAvailableAuditVolumesSortOrderEnum{
	"asc":  ListAvailableAuditVolumesSortOrderAsc,
	"desc": ListAvailableAuditVolumesSortOrderDesc,
}

// GetListAvailableAuditVolumesSortOrderEnumValues Enumerates the set of values for ListAvailableAuditVolumesSortOrderEnum
func GetListAvailableAuditVolumesSortOrderEnumValues() []ListAvailableAuditVolumesSortOrderEnum {
	values := make([]ListAvailableAuditVolumesSortOrderEnum, 0)
	for _, v := range mappingListAvailableAuditVolumesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableAuditVolumesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableAuditVolumesSortOrderEnum
func GetListAvailableAuditVolumesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableAuditVolumesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableAuditVolumesSortOrderEnum(val string) (ListAvailableAuditVolumesSortOrderEnum, bool) {
	enum, ok := mappingListAvailableAuditVolumesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableAuditVolumesSortByEnum Enum with underlying type: string
type ListAvailableAuditVolumesSortByEnum string

// Set of constants representing the allowable values for ListAvailableAuditVolumesSortByEnum
const (
	ListAvailableAuditVolumesSortByMonthinconsideration ListAvailableAuditVolumesSortByEnum = "monthInConsideration"
	ListAvailableAuditVolumesSortByVolume               ListAvailableAuditVolumesSortByEnum = "volume"
	ListAvailableAuditVolumesSortByTraillocation        ListAvailableAuditVolumesSortByEnum = "trailLocation"
)

var mappingListAvailableAuditVolumesSortByEnum = map[string]ListAvailableAuditVolumesSortByEnum{
	"monthInConsideration": ListAvailableAuditVolumesSortByMonthinconsideration,
	"volume":               ListAvailableAuditVolumesSortByVolume,
	"trailLocation":        ListAvailableAuditVolumesSortByTraillocation,
}

var mappingListAvailableAuditVolumesSortByEnumLowerCase = map[string]ListAvailableAuditVolumesSortByEnum{
	"monthinconsideration": ListAvailableAuditVolumesSortByMonthinconsideration,
	"volume":               ListAvailableAuditVolumesSortByVolume,
	"traillocation":        ListAvailableAuditVolumesSortByTraillocation,
}

// GetListAvailableAuditVolumesSortByEnumValues Enumerates the set of values for ListAvailableAuditVolumesSortByEnum
func GetListAvailableAuditVolumesSortByEnumValues() []ListAvailableAuditVolumesSortByEnum {
	values := make([]ListAvailableAuditVolumesSortByEnum, 0)
	for _, v := range mappingListAvailableAuditVolumesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableAuditVolumesSortByEnumStringValues Enumerates the set of values in String for ListAvailableAuditVolumesSortByEnum
func GetListAvailableAuditVolumesSortByEnumStringValues() []string {
	return []string{
		"monthInConsideration",
		"volume",
		"trailLocation",
	}
}

// GetMappingListAvailableAuditVolumesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableAuditVolumesSortByEnum(val string) (ListAvailableAuditVolumesSortByEnum, bool) {
	enum, ok := mappingListAvailableAuditVolumesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
