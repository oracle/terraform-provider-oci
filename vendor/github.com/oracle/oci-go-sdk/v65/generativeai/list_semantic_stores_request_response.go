// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSemanticStoresRequest wrapper for the ListSemanticStores operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListSemanticStores.go.html to see an example of how to use ListSemanticStoresRequest.
type ListSemanticStoresRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycle state matches the given array.
	LifecycleState []ListSemanticStoresLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SemanticStore.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSemanticStoresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListSemanticStoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources whose queryingConnectionId matches with this id.
	DataSourceQueryingConnectionId *string `mandatory:"false" contributesTo:"query" name:"dataSourceQueryingConnectionId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSemanticStoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSemanticStoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSemanticStoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSemanticStoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSemanticStoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingListSemanticStoresLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetListSemanticStoresLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSemanticStoresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSemanticStoresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSemanticStoresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSemanticStoresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSemanticStoresResponse wrapper for the ListSemanticStores operation
type ListSemanticStoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SemanticStoreCollection instances
	SemanticStoreCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSemanticStoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSemanticStoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSemanticStoresLifecycleStateEnum Enum with underlying type: string
type ListSemanticStoresLifecycleStateEnum string

// Set of constants representing the allowable values for ListSemanticStoresLifecycleStateEnum
const (
	ListSemanticStoresLifecycleStateActive   ListSemanticStoresLifecycleStateEnum = "ACTIVE"
	ListSemanticStoresLifecycleStateCreating ListSemanticStoresLifecycleStateEnum = "CREATING"
	ListSemanticStoresLifecycleStateUpdating ListSemanticStoresLifecycleStateEnum = "UPDATING"
	ListSemanticStoresLifecycleStateDeleting ListSemanticStoresLifecycleStateEnum = "DELETING"
	ListSemanticStoresLifecycleStateDeleted  ListSemanticStoresLifecycleStateEnum = "DELETED"
	ListSemanticStoresLifecycleStateFailed   ListSemanticStoresLifecycleStateEnum = "FAILED"
)

var mappingListSemanticStoresLifecycleStateEnum = map[string]ListSemanticStoresLifecycleStateEnum{
	"ACTIVE":   ListSemanticStoresLifecycleStateActive,
	"CREATING": ListSemanticStoresLifecycleStateCreating,
	"UPDATING": ListSemanticStoresLifecycleStateUpdating,
	"DELETING": ListSemanticStoresLifecycleStateDeleting,
	"DELETED":  ListSemanticStoresLifecycleStateDeleted,
	"FAILED":   ListSemanticStoresLifecycleStateFailed,
}

var mappingListSemanticStoresLifecycleStateEnumLowerCase = map[string]ListSemanticStoresLifecycleStateEnum{
	"active":   ListSemanticStoresLifecycleStateActive,
	"creating": ListSemanticStoresLifecycleStateCreating,
	"updating": ListSemanticStoresLifecycleStateUpdating,
	"deleting": ListSemanticStoresLifecycleStateDeleting,
	"deleted":  ListSemanticStoresLifecycleStateDeleted,
	"failed":   ListSemanticStoresLifecycleStateFailed,
}

// GetListSemanticStoresLifecycleStateEnumValues Enumerates the set of values for ListSemanticStoresLifecycleStateEnum
func GetListSemanticStoresLifecycleStateEnumValues() []ListSemanticStoresLifecycleStateEnum {
	values := make([]ListSemanticStoresLifecycleStateEnum, 0)
	for _, v := range mappingListSemanticStoresLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSemanticStoresLifecycleStateEnumStringValues Enumerates the set of values in String for ListSemanticStoresLifecycleStateEnum
func GetListSemanticStoresLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSemanticStoresLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSemanticStoresLifecycleStateEnum(val string) (ListSemanticStoresLifecycleStateEnum, bool) {
	enum, ok := mappingListSemanticStoresLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSemanticStoresSortOrderEnum Enum with underlying type: string
type ListSemanticStoresSortOrderEnum string

// Set of constants representing the allowable values for ListSemanticStoresSortOrderEnum
const (
	ListSemanticStoresSortOrderAsc  ListSemanticStoresSortOrderEnum = "ASC"
	ListSemanticStoresSortOrderDesc ListSemanticStoresSortOrderEnum = "DESC"
)

var mappingListSemanticStoresSortOrderEnum = map[string]ListSemanticStoresSortOrderEnum{
	"ASC":  ListSemanticStoresSortOrderAsc,
	"DESC": ListSemanticStoresSortOrderDesc,
}

var mappingListSemanticStoresSortOrderEnumLowerCase = map[string]ListSemanticStoresSortOrderEnum{
	"asc":  ListSemanticStoresSortOrderAsc,
	"desc": ListSemanticStoresSortOrderDesc,
}

// GetListSemanticStoresSortOrderEnumValues Enumerates the set of values for ListSemanticStoresSortOrderEnum
func GetListSemanticStoresSortOrderEnumValues() []ListSemanticStoresSortOrderEnum {
	values := make([]ListSemanticStoresSortOrderEnum, 0)
	for _, v := range mappingListSemanticStoresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSemanticStoresSortOrderEnumStringValues Enumerates the set of values in String for ListSemanticStoresSortOrderEnum
func GetListSemanticStoresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSemanticStoresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSemanticStoresSortOrderEnum(val string) (ListSemanticStoresSortOrderEnum, bool) {
	enum, ok := mappingListSemanticStoresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSemanticStoresSortByEnum Enum with underlying type: string
type ListSemanticStoresSortByEnum string

// Set of constants representing the allowable values for ListSemanticStoresSortByEnum
const (
	ListSemanticStoresSortByDisplayname ListSemanticStoresSortByEnum = "displayName"
	ListSemanticStoresSortByTimecreated ListSemanticStoresSortByEnum = "timeCreated"
)

var mappingListSemanticStoresSortByEnum = map[string]ListSemanticStoresSortByEnum{
	"displayName": ListSemanticStoresSortByDisplayname,
	"timeCreated": ListSemanticStoresSortByTimecreated,
}

var mappingListSemanticStoresSortByEnumLowerCase = map[string]ListSemanticStoresSortByEnum{
	"displayname": ListSemanticStoresSortByDisplayname,
	"timecreated": ListSemanticStoresSortByTimecreated,
}

// GetListSemanticStoresSortByEnumValues Enumerates the set of values for ListSemanticStoresSortByEnum
func GetListSemanticStoresSortByEnumValues() []ListSemanticStoresSortByEnum {
	values := make([]ListSemanticStoresSortByEnum, 0)
	for _, v := range mappingListSemanticStoresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSemanticStoresSortByEnumStringValues Enumerates the set of values in String for ListSemanticStoresSortByEnum
func GetListSemanticStoresSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListSemanticStoresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSemanticStoresSortByEnum(val string) (ListSemanticStoresSortByEnum, bool) {
	enum, ok := mappingListSemanticStoresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
