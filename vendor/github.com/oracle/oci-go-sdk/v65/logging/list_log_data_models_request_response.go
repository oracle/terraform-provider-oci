// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLogDataModelsRequest wrapper for the ListLogDataModels operation
type ListLogDataModelsRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Specifies whether or not nested compartments should be traversed. Defaults to false.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// Resource name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Template name query parameter of the log data model object.
	Template *string `mandatory:"false" contributesTo:"query" name:"template"`

	// Lifecycle state query parameter of the log data model object.
	LifecycleState ListLogDataModelsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogDataModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogDataModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogDataModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogDataModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogDataModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogDataModelsRequest) RetryPolicy() common.OCIRetry {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogDataModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogDataModelsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLogDataModelsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogDataModelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogDataModelsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogDataModelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogDataModelsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogDataModelsResponse wrapper for the ListLogDataModels operation
type ListLogDataModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogDataModelSummaryCollection instances
	LogDataModelSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`
}

func (response ListLogDataModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogDataModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogDataModelsLifecycleStateEnum Enum with underlying type: string
type ListLogDataModelsLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogDataModelsLifecycleStateEnum
const (
	ListLogDataModelsLifecycleStateCreating ListLogDataModelsLifecycleStateEnum = "CREATING"
	ListLogDataModelsLifecycleStateUpdating ListLogDataModelsLifecycleStateEnum = "UPDATING"
	ListLogDataModelsLifecycleStateDeleting ListLogDataModelsLifecycleStateEnum = "DELETING"
	ListLogDataModelsLifecycleStateActive   ListLogDataModelsLifecycleStateEnum = "ACTIVE"
	ListLogDataModelsLifecycleStateFailed   ListLogDataModelsLifecycleStateEnum = "FAILED"
)

var mappingListLogDataModelsLifecycleStateEnum = map[string]ListLogDataModelsLifecycleStateEnum{
	"CREATING": ListLogDataModelsLifecycleStateCreating,
	"UPDATING": ListLogDataModelsLifecycleStateUpdating,
	"DELETING": ListLogDataModelsLifecycleStateDeleting,
	"ACTIVE":   ListLogDataModelsLifecycleStateActive,
	"FAILED":   ListLogDataModelsLifecycleStateFailed,
}

var mappingListLogDataModelsLifecycleStateEnumLowerCase = map[string]ListLogDataModelsLifecycleStateEnum{
	"creating": ListLogDataModelsLifecycleStateCreating,
	"updating": ListLogDataModelsLifecycleStateUpdating,
	"deleting": ListLogDataModelsLifecycleStateDeleting,
	"active":   ListLogDataModelsLifecycleStateActive,
	"failed":   ListLogDataModelsLifecycleStateFailed,
}

// GetListLogDataModelsLifecycleStateEnumValues Enumerates the set of values for ListLogDataModelsLifecycleStateEnum
func GetListLogDataModelsLifecycleStateEnumValues() []ListLogDataModelsLifecycleStateEnum {
	values := make([]ListLogDataModelsLifecycleStateEnum, 0)
	for _, v := range mappingListLogDataModelsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogDataModelsLifecycleStateEnumStringValues Enumerates the set of values in String for ListLogDataModelsLifecycleStateEnum
func GetListLogDataModelsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}

// GetMappingListLogDataModelsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogDataModelsLifecycleStateEnum(val string) (ListLogDataModelsLifecycleStateEnum, bool) {
	enum, ok := mappingListLogDataModelsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogDataModelsSortByEnum Enum with underlying type: string
type ListLogDataModelsSortByEnum string

// Set of constants representing the allowable values for ListLogDataModelsSortByEnum
const (
	ListLogDataModelsSortByTimecreated ListLogDataModelsSortByEnum = "timeCreated"
	ListLogDataModelsSortByDisplayname ListLogDataModelsSortByEnum = "displayName"
)

var mappingListLogDataModelsSortByEnum = map[string]ListLogDataModelsSortByEnum{
	"timeCreated": ListLogDataModelsSortByTimecreated,
	"displayName": ListLogDataModelsSortByDisplayname,
}

var mappingListLogDataModelsSortByEnumLowerCase = map[string]ListLogDataModelsSortByEnum{
	"timecreated": ListLogDataModelsSortByTimecreated,
	"displayname": ListLogDataModelsSortByDisplayname,
}

// GetListLogDataModelsSortByEnumValues Enumerates the set of values for ListLogDataModelsSortByEnum
func GetListLogDataModelsSortByEnumValues() []ListLogDataModelsSortByEnum {
	values := make([]ListLogDataModelsSortByEnum, 0)
	for _, v := range mappingListLogDataModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogDataModelsSortByEnumStringValues Enumerates the set of values in String for ListLogDataModelsSortByEnum
func GetListLogDataModelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLogDataModelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogDataModelsSortByEnum(val string) (ListLogDataModelsSortByEnum, bool) {
	enum, ok := mappingListLogDataModelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogDataModelsSortOrderEnum Enum with underlying type: string
type ListLogDataModelsSortOrderEnum string

// Set of constants representing the allowable values for ListLogDataModelsSortOrderEnum
const (
	ListLogDataModelsSortOrderAsc  ListLogDataModelsSortOrderEnum = "ASC"
	ListLogDataModelsSortOrderDesc ListLogDataModelsSortOrderEnum = "DESC"
)

var mappingListLogDataModelsSortOrderEnum = map[string]ListLogDataModelsSortOrderEnum{
	"ASC":  ListLogDataModelsSortOrderAsc,
	"DESC": ListLogDataModelsSortOrderDesc,
}

var mappingListLogDataModelsSortOrderEnumLowerCase = map[string]ListLogDataModelsSortOrderEnum{
	"asc":  ListLogDataModelsSortOrderAsc,
	"desc": ListLogDataModelsSortOrderDesc,
}

// GetListLogDataModelsSortOrderEnumValues Enumerates the set of values for ListLogDataModelsSortOrderEnum
func GetListLogDataModelsSortOrderEnumValues() []ListLogDataModelsSortOrderEnum {
	values := make([]ListLogDataModelsSortOrderEnum, 0)
	for _, v := range mappingListLogDataModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogDataModelsSortOrderEnumStringValues Enumerates the set of values in String for ListLogDataModelsSortOrderEnum
func GetListLogDataModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogDataModelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogDataModelsSortOrderEnum(val string) (ListLogDataModelsSortOrderEnum, bool) {
	enum, ok := mappingListLogDataModelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
