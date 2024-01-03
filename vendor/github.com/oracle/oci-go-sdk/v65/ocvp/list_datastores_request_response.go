// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatastoresRequest wrapper for the ListDatastores operation
type ListDatastoresRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Datastore.
	DatastoreId *string `mandatory:"false" contributesTo:"query" name:"datastoreId"`

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
	SortOrder ListDatastoresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListDatastoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListDatastoresLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatastoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatastoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatastoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatastoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatastoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatastoresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatastoresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatastoresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatastoresSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatastoresLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatastoresLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatastoresResponse wrapper for the ListDatastores operation
type ListDatastoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatastoreCollection instances
	DatastoreCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatastoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatastoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatastoresSortOrderEnum Enum with underlying type: string
type ListDatastoresSortOrderEnum string

// Set of constants representing the allowable values for ListDatastoresSortOrderEnum
const (
	ListDatastoresSortOrderAsc  ListDatastoresSortOrderEnum = "ASC"
	ListDatastoresSortOrderDesc ListDatastoresSortOrderEnum = "DESC"
)

var mappingListDatastoresSortOrderEnum = map[string]ListDatastoresSortOrderEnum{
	"ASC":  ListDatastoresSortOrderAsc,
	"DESC": ListDatastoresSortOrderDesc,
}

var mappingListDatastoresSortOrderEnumLowerCase = map[string]ListDatastoresSortOrderEnum{
	"asc":  ListDatastoresSortOrderAsc,
	"desc": ListDatastoresSortOrderDesc,
}

// GetListDatastoresSortOrderEnumValues Enumerates the set of values for ListDatastoresSortOrderEnum
func GetListDatastoresSortOrderEnumValues() []ListDatastoresSortOrderEnum {
	values := make([]ListDatastoresSortOrderEnum, 0)
	for _, v := range mappingListDatastoresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatastoresSortOrderEnumStringValues Enumerates the set of values in String for ListDatastoresSortOrderEnum
func GetListDatastoresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatastoresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatastoresSortOrderEnum(val string) (ListDatastoresSortOrderEnum, bool) {
	enum, ok := mappingListDatastoresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatastoresSortByEnum Enum with underlying type: string
type ListDatastoresSortByEnum string

// Set of constants representing the allowable values for ListDatastoresSortByEnum
const (
	ListDatastoresSortByTimecreated ListDatastoresSortByEnum = "timeCreated"
	ListDatastoresSortByDisplayname ListDatastoresSortByEnum = "displayName"
)

var mappingListDatastoresSortByEnum = map[string]ListDatastoresSortByEnum{
	"timeCreated": ListDatastoresSortByTimecreated,
	"displayName": ListDatastoresSortByDisplayname,
}

var mappingListDatastoresSortByEnumLowerCase = map[string]ListDatastoresSortByEnum{
	"timecreated": ListDatastoresSortByTimecreated,
	"displayname": ListDatastoresSortByDisplayname,
}

// GetListDatastoresSortByEnumValues Enumerates the set of values for ListDatastoresSortByEnum
func GetListDatastoresSortByEnumValues() []ListDatastoresSortByEnum {
	values := make([]ListDatastoresSortByEnum, 0)
	for _, v := range mappingListDatastoresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatastoresSortByEnumStringValues Enumerates the set of values in String for ListDatastoresSortByEnum
func GetListDatastoresSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatastoresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatastoresSortByEnum(val string) (ListDatastoresSortByEnum, bool) {
	enum, ok := mappingListDatastoresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatastoresLifecycleStateEnum Enum with underlying type: string
type ListDatastoresLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatastoresLifecycleStateEnum
const (
	ListDatastoresLifecycleStateCreating ListDatastoresLifecycleStateEnum = "CREATING"
	ListDatastoresLifecycleStateUpdating ListDatastoresLifecycleStateEnum = "UPDATING"
	ListDatastoresLifecycleStateActive   ListDatastoresLifecycleStateEnum = "ACTIVE"
	ListDatastoresLifecycleStateDeleting ListDatastoresLifecycleStateEnum = "DELETING"
	ListDatastoresLifecycleStateDeleted  ListDatastoresLifecycleStateEnum = "DELETED"
	ListDatastoresLifecycleStateFailed   ListDatastoresLifecycleStateEnum = "FAILED"
)

var mappingListDatastoresLifecycleStateEnum = map[string]ListDatastoresLifecycleStateEnum{
	"CREATING": ListDatastoresLifecycleStateCreating,
	"UPDATING": ListDatastoresLifecycleStateUpdating,
	"ACTIVE":   ListDatastoresLifecycleStateActive,
	"DELETING": ListDatastoresLifecycleStateDeleting,
	"DELETED":  ListDatastoresLifecycleStateDeleted,
	"FAILED":   ListDatastoresLifecycleStateFailed,
}

var mappingListDatastoresLifecycleStateEnumLowerCase = map[string]ListDatastoresLifecycleStateEnum{
	"creating": ListDatastoresLifecycleStateCreating,
	"updating": ListDatastoresLifecycleStateUpdating,
	"active":   ListDatastoresLifecycleStateActive,
	"deleting": ListDatastoresLifecycleStateDeleting,
	"deleted":  ListDatastoresLifecycleStateDeleted,
	"failed":   ListDatastoresLifecycleStateFailed,
}

// GetListDatastoresLifecycleStateEnumValues Enumerates the set of values for ListDatastoresLifecycleStateEnum
func GetListDatastoresLifecycleStateEnumValues() []ListDatastoresLifecycleStateEnum {
	values := make([]ListDatastoresLifecycleStateEnum, 0)
	for _, v := range mappingListDatastoresLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatastoresLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatastoresLifecycleStateEnum
func GetListDatastoresLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDatastoresLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatastoresLifecycleStateEnum(val string) (ListDatastoresLifecycleStateEnum, bool) {
	enum, ok := mappingListDatastoresLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
