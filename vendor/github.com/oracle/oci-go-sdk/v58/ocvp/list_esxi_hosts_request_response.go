// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListEsxiHostsRequest wrapper for the ListEsxiHosts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListEsxiHosts.go.html to see an example of how to use ListEsxiHostsRequest.
type ListEsxiHostsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	SddcId *string `mandatory:"false" contributesTo:"query" name:"sddcId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Compute instance.
	ComputeInstanceId *string `mandatory:"false" contributesTo:"query" name:"computeInstanceId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListEsxiHostsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListEsxiHostsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListEsxiHostsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEsxiHostsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEsxiHostsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEsxiHostsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEsxiHostsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEsxiHostsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEsxiHostsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEsxiHostsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEsxiHostsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEsxiHostsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEsxiHostsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListEsxiHostsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEsxiHostsResponse wrapper for the ListEsxiHosts operation
type ListEsxiHostsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EsxiHostCollection instances
	EsxiHostCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEsxiHostsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEsxiHostsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEsxiHostsSortOrderEnum Enum with underlying type: string
type ListEsxiHostsSortOrderEnum string

// Set of constants representing the allowable values for ListEsxiHostsSortOrderEnum
const (
	ListEsxiHostsSortOrderAsc  ListEsxiHostsSortOrderEnum = "ASC"
	ListEsxiHostsSortOrderDesc ListEsxiHostsSortOrderEnum = "DESC"
)

var mappingListEsxiHostsSortOrderEnum = map[string]ListEsxiHostsSortOrderEnum{
	"ASC":  ListEsxiHostsSortOrderAsc,
	"DESC": ListEsxiHostsSortOrderDesc,
}

// GetListEsxiHostsSortOrderEnumValues Enumerates the set of values for ListEsxiHostsSortOrderEnum
func GetListEsxiHostsSortOrderEnumValues() []ListEsxiHostsSortOrderEnum {
	values := make([]ListEsxiHostsSortOrderEnum, 0)
	for _, v := range mappingListEsxiHostsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEsxiHostsSortOrderEnumStringValues Enumerates the set of values in String for ListEsxiHostsSortOrderEnum
func GetListEsxiHostsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEsxiHostsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEsxiHostsSortOrderEnum(val string) (ListEsxiHostsSortOrderEnum, bool) {
	mappingListEsxiHostsSortOrderEnumIgnoreCase := make(map[string]ListEsxiHostsSortOrderEnum)
	for k, v := range mappingListEsxiHostsSortOrderEnum {
		mappingListEsxiHostsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEsxiHostsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEsxiHostsSortByEnum Enum with underlying type: string
type ListEsxiHostsSortByEnum string

// Set of constants representing the allowable values for ListEsxiHostsSortByEnum
const (
	ListEsxiHostsSortByTimecreated ListEsxiHostsSortByEnum = "timeCreated"
	ListEsxiHostsSortByDisplayname ListEsxiHostsSortByEnum = "displayName"
)

var mappingListEsxiHostsSortByEnum = map[string]ListEsxiHostsSortByEnum{
	"timeCreated": ListEsxiHostsSortByTimecreated,
	"displayName": ListEsxiHostsSortByDisplayname,
}

// GetListEsxiHostsSortByEnumValues Enumerates the set of values for ListEsxiHostsSortByEnum
func GetListEsxiHostsSortByEnumValues() []ListEsxiHostsSortByEnum {
	values := make([]ListEsxiHostsSortByEnum, 0)
	for _, v := range mappingListEsxiHostsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEsxiHostsSortByEnumStringValues Enumerates the set of values in String for ListEsxiHostsSortByEnum
func GetListEsxiHostsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListEsxiHostsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEsxiHostsSortByEnum(val string) (ListEsxiHostsSortByEnum, bool) {
	mappingListEsxiHostsSortByEnumIgnoreCase := make(map[string]ListEsxiHostsSortByEnum)
	for k, v := range mappingListEsxiHostsSortByEnum {
		mappingListEsxiHostsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEsxiHostsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEsxiHostsLifecycleStateEnum Enum with underlying type: string
type ListEsxiHostsLifecycleStateEnum string

// Set of constants representing the allowable values for ListEsxiHostsLifecycleStateEnum
const (
	ListEsxiHostsLifecycleStateCreating ListEsxiHostsLifecycleStateEnum = "CREATING"
	ListEsxiHostsLifecycleStateUpdating ListEsxiHostsLifecycleStateEnum = "UPDATING"
	ListEsxiHostsLifecycleStateActive   ListEsxiHostsLifecycleStateEnum = "ACTIVE"
	ListEsxiHostsLifecycleStateDeleting ListEsxiHostsLifecycleStateEnum = "DELETING"
	ListEsxiHostsLifecycleStateDeleted  ListEsxiHostsLifecycleStateEnum = "DELETED"
	ListEsxiHostsLifecycleStateFailed   ListEsxiHostsLifecycleStateEnum = "FAILED"
)

var mappingListEsxiHostsLifecycleStateEnum = map[string]ListEsxiHostsLifecycleStateEnum{
	"CREATING": ListEsxiHostsLifecycleStateCreating,
	"UPDATING": ListEsxiHostsLifecycleStateUpdating,
	"ACTIVE":   ListEsxiHostsLifecycleStateActive,
	"DELETING": ListEsxiHostsLifecycleStateDeleting,
	"DELETED":  ListEsxiHostsLifecycleStateDeleted,
	"FAILED":   ListEsxiHostsLifecycleStateFailed,
}

// GetListEsxiHostsLifecycleStateEnumValues Enumerates the set of values for ListEsxiHostsLifecycleStateEnum
func GetListEsxiHostsLifecycleStateEnumValues() []ListEsxiHostsLifecycleStateEnum {
	values := make([]ListEsxiHostsLifecycleStateEnum, 0)
	for _, v := range mappingListEsxiHostsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListEsxiHostsLifecycleStateEnumStringValues Enumerates the set of values in String for ListEsxiHostsLifecycleStateEnum
func GetListEsxiHostsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListEsxiHostsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEsxiHostsLifecycleStateEnum(val string) (ListEsxiHostsLifecycleStateEnum, bool) {
	mappingListEsxiHostsLifecycleStateEnumIgnoreCase := make(map[string]ListEsxiHostsLifecycleStateEnum)
	for k, v := range mappingListEsxiHostsLifecycleStateEnum {
		mappingListEsxiHostsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEsxiHostsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
