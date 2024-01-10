// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListKnowledgeBasesRequest wrapper for the ListKnowledgeBases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListKnowledgeBases.go.html to see an example of how to use ListKnowledgeBasesRequest.
type ListKnowledgeBasesRequest struct {

	// A filter to return only resources that match the specified identifier.
	// Required only if the compartmentId query parameter is not specified.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field used to sort Knowledge Bases. Only one sort order is allowed.
	// Default order for _displayName_ is **ascending alphabetical order**.
	// Default order for _lifecyleState_ is the following sequence: **CREATING, ACTIVE, UPDATING, FAILED, DELETING, and DELETED**.Default order for _timeCreated_ is **descending**.
	// Default order for _timeUpdated_ is **descending**.
	SortBy ListKnowledgeBasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only Knowledge Bases that match the specified lifecycleState.
	LifecycleState KnowledgeBaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListKnowledgeBasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that belong to the specified compartment identifier.
	// Required only if the id query param is not specified.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKnowledgeBasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKnowledgeBasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKnowledgeBasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKnowledgeBasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKnowledgeBasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListKnowledgeBasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListKnowledgeBasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKnowledgeBaseLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetKnowledgeBaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKnowledgeBasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListKnowledgeBasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKnowledgeBasesResponse wrapper for the ListKnowledgeBases operation
type ListKnowledgeBasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of KnowledgeBaseCollection instances
	KnowledgeBaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListKnowledgeBasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKnowledgeBasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKnowledgeBasesSortByEnum Enum with underlying type: string
type ListKnowledgeBasesSortByEnum string

// Set of constants representing the allowable values for ListKnowledgeBasesSortByEnum
const (
	ListKnowledgeBasesSortByDisplayName    ListKnowledgeBasesSortByEnum = "DISPLAY_NAME"
	ListKnowledgeBasesSortByLifecycleState ListKnowledgeBasesSortByEnum = "LIFECYCLE_STATE"
	ListKnowledgeBasesSortByTimeCreated    ListKnowledgeBasesSortByEnum = "TIME_CREATED"
	ListKnowledgeBasesSortByTimeUpdated    ListKnowledgeBasesSortByEnum = "TIME_UPDATED"
)

var mappingListKnowledgeBasesSortByEnum = map[string]ListKnowledgeBasesSortByEnum{
	"DISPLAY_NAME":    ListKnowledgeBasesSortByDisplayName,
	"LIFECYCLE_STATE": ListKnowledgeBasesSortByLifecycleState,
	"TIME_CREATED":    ListKnowledgeBasesSortByTimeCreated,
	"TIME_UPDATED":    ListKnowledgeBasesSortByTimeUpdated,
}

var mappingListKnowledgeBasesSortByEnumLowerCase = map[string]ListKnowledgeBasesSortByEnum{
	"display_name":    ListKnowledgeBasesSortByDisplayName,
	"lifecycle_state": ListKnowledgeBasesSortByLifecycleState,
	"time_created":    ListKnowledgeBasesSortByTimeCreated,
	"time_updated":    ListKnowledgeBasesSortByTimeUpdated,
}

// GetListKnowledgeBasesSortByEnumValues Enumerates the set of values for ListKnowledgeBasesSortByEnum
func GetListKnowledgeBasesSortByEnumValues() []ListKnowledgeBasesSortByEnum {
	values := make([]ListKnowledgeBasesSortByEnum, 0)
	for _, v := range mappingListKnowledgeBasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListKnowledgeBasesSortByEnumStringValues Enumerates the set of values in String for ListKnowledgeBasesSortByEnum
func GetListKnowledgeBasesSortByEnumStringValues() []string {
	return []string{
		"DISPLAY_NAME",
		"LIFECYCLE_STATE",
		"TIME_CREATED",
		"TIME_UPDATED",
	}
}

// GetMappingListKnowledgeBasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKnowledgeBasesSortByEnum(val string) (ListKnowledgeBasesSortByEnum, bool) {
	enum, ok := mappingListKnowledgeBasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListKnowledgeBasesSortOrderEnum Enum with underlying type: string
type ListKnowledgeBasesSortOrderEnum string

// Set of constants representing the allowable values for ListKnowledgeBasesSortOrderEnum
const (
	ListKnowledgeBasesSortOrderAsc  ListKnowledgeBasesSortOrderEnum = "ASC"
	ListKnowledgeBasesSortOrderDesc ListKnowledgeBasesSortOrderEnum = "DESC"
)

var mappingListKnowledgeBasesSortOrderEnum = map[string]ListKnowledgeBasesSortOrderEnum{
	"ASC":  ListKnowledgeBasesSortOrderAsc,
	"DESC": ListKnowledgeBasesSortOrderDesc,
}

var mappingListKnowledgeBasesSortOrderEnumLowerCase = map[string]ListKnowledgeBasesSortOrderEnum{
	"asc":  ListKnowledgeBasesSortOrderAsc,
	"desc": ListKnowledgeBasesSortOrderDesc,
}

// GetListKnowledgeBasesSortOrderEnumValues Enumerates the set of values for ListKnowledgeBasesSortOrderEnum
func GetListKnowledgeBasesSortOrderEnumValues() []ListKnowledgeBasesSortOrderEnum {
	values := make([]ListKnowledgeBasesSortOrderEnum, 0)
	for _, v := range mappingListKnowledgeBasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListKnowledgeBasesSortOrderEnumStringValues Enumerates the set of values in String for ListKnowledgeBasesSortOrderEnum
func GetListKnowledgeBasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListKnowledgeBasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKnowledgeBasesSortOrderEnum(val string) (ListKnowledgeBasesSortOrderEnum, bool) {
	enum, ok := mappingListKnowledgeBasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
