// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMonitoredResourceTypesRequest wrapper for the ListMonitoredResourceTypes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoredResourceTypes.go.html to see an example of how to use ListMonitoredResourceTypesRequest.
type ListMonitoredResourceTypesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy for which
	// monitored resource types should be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return monitored resource types that match exactly with the resource type name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that matches with lifecycleState given.
	Status ListMonitoredResourceTypesStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to exclude system resource types. If set to true, system resource types will be excluded.
	IsExcludeSystemTypes *bool `mandatory:"false" contributesTo:"query" name:"isExcludeSystemTypes"`

	// A filter to return monitored resource types that has the matching namespace.
	MetricNamespace *string `mandatory:"false" contributesTo:"query" name:"metricNamespace"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for 'timeUpdated' is descending. Default order for 'name' is ascending.
	SortBy ListMonitoredResourceTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMonitoredResourceTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Partial response refers to an optimization technique offered
	// by the RESTful web APIs, to return only the information
	// (fields) required by the client. In this mechanism, the client
	// sends the required field names as the query parameters for
	// an API to the server, and the server trims down the default
	// response content by removing the fields that are not required
	// by the client. The parameter controls which fields to
	// return and should be a query string parameter called "fields" of
	// an array type, provide the values as enums, and use collectionFormat.
	// MonitoredResourceType Id, name and compartment will be added by default.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Partial response refers to an optimization technique offered
	// by the RESTful web APIs, to return all the information except
	// the fields requested to be excluded (excludeFields) by the client.
	// In this mechanism, the client
	// sends the exclude field names as the query parameters for
	// an API to the server, and the server trims down the default
	// response content by removing the fields that are not required
	// by the client. The parameter controls which fields to
	// exlude and to return and should be a query string parameter
	// called "excludeFields" of an array type, provide the values
	// as enums, and use collectionFormat.
	ExcludeFields []string `contributesTo:"query" name:"excludeFields" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitoredResourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoredResourceTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoredResourceTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoredResourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitoredResourceTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMonitoredResourceTypesStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListMonitoredResourceTypesStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredResourceTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitoredResourceTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredResourceTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitoredResourceTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitoredResourceTypesResponse wrapper for the ListMonitoredResourceTypes operation
type ListMonitoredResourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredResourceTypesCollection instances
	MonitoredResourceTypesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListMonitoredResourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoredResourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoredResourceTypesStatusEnum Enum with underlying type: string
type ListMonitoredResourceTypesStatusEnum string

// Set of constants representing the allowable values for ListMonitoredResourceTypesStatusEnum
const (
	ListMonitoredResourceTypesStatusCreating ListMonitoredResourceTypesStatusEnum = "CREATING"
	ListMonitoredResourceTypesStatusUpdating ListMonitoredResourceTypesStatusEnum = "UPDATING"
	ListMonitoredResourceTypesStatusActive   ListMonitoredResourceTypesStatusEnum = "ACTIVE"
	ListMonitoredResourceTypesStatusDeleting ListMonitoredResourceTypesStatusEnum = "DELETING"
	ListMonitoredResourceTypesStatusDeleted  ListMonitoredResourceTypesStatusEnum = "DELETED"
	ListMonitoredResourceTypesStatusFailed   ListMonitoredResourceTypesStatusEnum = "FAILED"
)

var mappingListMonitoredResourceTypesStatusEnum = map[string]ListMonitoredResourceTypesStatusEnum{
	"CREATING": ListMonitoredResourceTypesStatusCreating,
	"UPDATING": ListMonitoredResourceTypesStatusUpdating,
	"ACTIVE":   ListMonitoredResourceTypesStatusActive,
	"DELETING": ListMonitoredResourceTypesStatusDeleting,
	"DELETED":  ListMonitoredResourceTypesStatusDeleted,
	"FAILED":   ListMonitoredResourceTypesStatusFailed,
}

var mappingListMonitoredResourceTypesStatusEnumLowerCase = map[string]ListMonitoredResourceTypesStatusEnum{
	"creating": ListMonitoredResourceTypesStatusCreating,
	"updating": ListMonitoredResourceTypesStatusUpdating,
	"active":   ListMonitoredResourceTypesStatusActive,
	"deleting": ListMonitoredResourceTypesStatusDeleting,
	"deleted":  ListMonitoredResourceTypesStatusDeleted,
	"failed":   ListMonitoredResourceTypesStatusFailed,
}

// GetListMonitoredResourceTypesStatusEnumValues Enumerates the set of values for ListMonitoredResourceTypesStatusEnum
func GetListMonitoredResourceTypesStatusEnumValues() []ListMonitoredResourceTypesStatusEnum {
	values := make([]ListMonitoredResourceTypesStatusEnum, 0)
	for _, v := range mappingListMonitoredResourceTypesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourceTypesStatusEnumStringValues Enumerates the set of values in String for ListMonitoredResourceTypesStatusEnum
func GetListMonitoredResourceTypesStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListMonitoredResourceTypesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourceTypesStatusEnum(val string) (ListMonitoredResourceTypesStatusEnum, bool) {
	enum, ok := mappingListMonitoredResourceTypesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoredResourceTypesSortByEnum Enum with underlying type: string
type ListMonitoredResourceTypesSortByEnum string

// Set of constants representing the allowable values for ListMonitoredResourceTypesSortByEnum
const (
	ListMonitoredResourceTypesSortByTimeupdated ListMonitoredResourceTypesSortByEnum = "timeUpdated"
	ListMonitoredResourceTypesSortByName        ListMonitoredResourceTypesSortByEnum = "name"
)

var mappingListMonitoredResourceTypesSortByEnum = map[string]ListMonitoredResourceTypesSortByEnum{
	"timeUpdated": ListMonitoredResourceTypesSortByTimeupdated,
	"name":        ListMonitoredResourceTypesSortByName,
}

var mappingListMonitoredResourceTypesSortByEnumLowerCase = map[string]ListMonitoredResourceTypesSortByEnum{
	"timeupdated": ListMonitoredResourceTypesSortByTimeupdated,
	"name":        ListMonitoredResourceTypesSortByName,
}

// GetListMonitoredResourceTypesSortByEnumValues Enumerates the set of values for ListMonitoredResourceTypesSortByEnum
func GetListMonitoredResourceTypesSortByEnumValues() []ListMonitoredResourceTypesSortByEnum {
	values := make([]ListMonitoredResourceTypesSortByEnum, 0)
	for _, v := range mappingListMonitoredResourceTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourceTypesSortByEnumStringValues Enumerates the set of values in String for ListMonitoredResourceTypesSortByEnum
func GetListMonitoredResourceTypesSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
		"name",
	}
}

// GetMappingListMonitoredResourceTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourceTypesSortByEnum(val string) (ListMonitoredResourceTypesSortByEnum, bool) {
	enum, ok := mappingListMonitoredResourceTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoredResourceTypesSortOrderEnum Enum with underlying type: string
type ListMonitoredResourceTypesSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoredResourceTypesSortOrderEnum
const (
	ListMonitoredResourceTypesSortOrderAsc  ListMonitoredResourceTypesSortOrderEnum = "ASC"
	ListMonitoredResourceTypesSortOrderDesc ListMonitoredResourceTypesSortOrderEnum = "DESC"
)

var mappingListMonitoredResourceTypesSortOrderEnum = map[string]ListMonitoredResourceTypesSortOrderEnum{
	"ASC":  ListMonitoredResourceTypesSortOrderAsc,
	"DESC": ListMonitoredResourceTypesSortOrderDesc,
}

var mappingListMonitoredResourceTypesSortOrderEnumLowerCase = map[string]ListMonitoredResourceTypesSortOrderEnum{
	"asc":  ListMonitoredResourceTypesSortOrderAsc,
	"desc": ListMonitoredResourceTypesSortOrderDesc,
}

// GetListMonitoredResourceTypesSortOrderEnumValues Enumerates the set of values for ListMonitoredResourceTypesSortOrderEnum
func GetListMonitoredResourceTypesSortOrderEnumValues() []ListMonitoredResourceTypesSortOrderEnum {
	values := make([]ListMonitoredResourceTypesSortOrderEnum, 0)
	for _, v := range mappingListMonitoredResourceTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourceTypesSortOrderEnumStringValues Enumerates the set of values in String for ListMonitoredResourceTypesSortOrderEnum
func GetListMonitoredResourceTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitoredResourceTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourceTypesSortOrderEnum(val string) (ListMonitoredResourceTypesSortOrderEnum, bool) {
	enum, ok := mappingListMonitoredResourceTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
