// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOccAvailabilitiesRequest wrapper for the ListOccAvailabilities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccAvailabilities.go.html to see an example of how to use ListOccAvailabilitiesRequest.
type ListOccAvailabilitiesRequest struct {

	// The OCID of the availability catalog.
	OccAvailabilityCatalogId *string `mandatory:"true" contributesTo:"path" name:"occAvailabilityCatalogId"`

	// The capacity handover date of the capacity constraint to filter the list of capacity constraints.
	DateExpectedCapacityHandover *string `mandatory:"false" contributesTo:"query" name:"dateExpectedCapacityHandover"`

	// The name of the resource to filter the list of capacity constraints.
	ResourceName *string `mandatory:"false" contributesTo:"query" name:"resourceName"`

	// Resource type using which the capacity constraints of an availability catalog can be filtered.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// Workload type using the resources in an availability catalog can be filtered.
	WorkloadType *string `mandatory:"false" contributesTo:"query" name:"workloadType"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccAvailabilitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.  Only one sort order may be provided. The default order for resource name is ascending. The default order for date of capacity handover is descending.
	SortBy ListOccAvailabilitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccAvailabilitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccAvailabilitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccAvailabilitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccAvailabilitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccAvailabilitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccAvailabilitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccAvailabilitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccAvailabilitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccAvailabilitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccAvailabilitiesResponse wrapper for the ListOccAvailabilities operation
type ListOccAvailabilitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccAvailabilityCollection instances
	OccAvailabilityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccAvailabilitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccAvailabilitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccAvailabilitiesSortOrderEnum Enum with underlying type: string
type ListOccAvailabilitiesSortOrderEnum string

// Set of constants representing the allowable values for ListOccAvailabilitiesSortOrderEnum
const (
	ListOccAvailabilitiesSortOrderAsc  ListOccAvailabilitiesSortOrderEnum = "ASC"
	ListOccAvailabilitiesSortOrderDesc ListOccAvailabilitiesSortOrderEnum = "DESC"
)

var mappingListOccAvailabilitiesSortOrderEnum = map[string]ListOccAvailabilitiesSortOrderEnum{
	"ASC":  ListOccAvailabilitiesSortOrderAsc,
	"DESC": ListOccAvailabilitiesSortOrderDesc,
}

var mappingListOccAvailabilitiesSortOrderEnumLowerCase = map[string]ListOccAvailabilitiesSortOrderEnum{
	"asc":  ListOccAvailabilitiesSortOrderAsc,
	"desc": ListOccAvailabilitiesSortOrderDesc,
}

// GetListOccAvailabilitiesSortOrderEnumValues Enumerates the set of values for ListOccAvailabilitiesSortOrderEnum
func GetListOccAvailabilitiesSortOrderEnumValues() []ListOccAvailabilitiesSortOrderEnum {
	values := make([]ListOccAvailabilitiesSortOrderEnum, 0)
	for _, v := range mappingListOccAvailabilitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilitiesSortOrderEnumStringValues Enumerates the set of values in String for ListOccAvailabilitiesSortOrderEnum
func GetListOccAvailabilitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccAvailabilitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilitiesSortOrderEnum(val string) (ListOccAvailabilitiesSortOrderEnum, bool) {
	enum, ok := mappingListOccAvailabilitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccAvailabilitiesSortByEnum Enum with underlying type: string
type ListOccAvailabilitiesSortByEnum string

// Set of constants representing the allowable values for ListOccAvailabilitiesSortByEnum
const (
	ListOccAvailabilitiesSortByDateexpectedcapacityhandover ListOccAvailabilitiesSortByEnum = "dateExpectedCapacityHandover"
	ListOccAvailabilitiesSortByResourcename                 ListOccAvailabilitiesSortByEnum = "resourceName"
)

var mappingListOccAvailabilitiesSortByEnum = map[string]ListOccAvailabilitiesSortByEnum{
	"dateExpectedCapacityHandover": ListOccAvailabilitiesSortByDateexpectedcapacityhandover,
	"resourceName":                 ListOccAvailabilitiesSortByResourcename,
}

var mappingListOccAvailabilitiesSortByEnumLowerCase = map[string]ListOccAvailabilitiesSortByEnum{
	"dateexpectedcapacityhandover": ListOccAvailabilitiesSortByDateexpectedcapacityhandover,
	"resourcename":                 ListOccAvailabilitiesSortByResourcename,
}

// GetListOccAvailabilitiesSortByEnumValues Enumerates the set of values for ListOccAvailabilitiesSortByEnum
func GetListOccAvailabilitiesSortByEnumValues() []ListOccAvailabilitiesSortByEnum {
	values := make([]ListOccAvailabilitiesSortByEnum, 0)
	for _, v := range mappingListOccAvailabilitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilitiesSortByEnumStringValues Enumerates the set of values in String for ListOccAvailabilitiesSortByEnum
func GetListOccAvailabilitiesSortByEnumStringValues() []string {
	return []string{
		"dateExpectedCapacityHandover",
		"resourceName",
	}
}

// GetMappingListOccAvailabilitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilitiesSortByEnum(val string) (ListOccAvailabilitiesSortByEnum, bool) {
	enum, ok := mappingListOccAvailabilitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
