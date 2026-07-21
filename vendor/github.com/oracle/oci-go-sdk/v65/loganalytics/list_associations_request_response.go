// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssociationsRequest wrapper for the ListAssociations operation
type ListAssociationsRequest struct {

	// The Log Analytics namespace used for the request. The namespace can be obtained by running 'oci os ns get'
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The collection rule OCID.
	CollectionRuleId *string `mandatory:"false" contributesTo:"query" name:"collectionRuleId"`

	// A list of source names used for filtering.  Only fields used by the specified
	// sources will be returned.
	SourceNames *string `mandatory:"false" contributesTo:"query" name:"sourceNames"`

	// The life cycle state used for filtering.  Only associations with the specified
	// life cycle state will be returned.
	LifeCycleState ListAssociationsLifeCycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifeCycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned associations
	SortBy ListAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssociationsLifeCycleStateEnum(string(request.LifeCycleState)); !ok && request.LifeCycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifeCycleState: %s. Supported values are: %s.", request.LifeCycleState, strings.Join(GetListAssociationsLifeCycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssociationsResponse wrapper for the ListAssociations operation
type ListAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociationCollection instances
	AssociationCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociationsLifeCycleStateEnum Enum with underlying type: string
type ListAssociationsLifeCycleStateEnum string

// Set of constants representing the allowable values for ListAssociationsLifeCycleStateEnum
const (
	ListAssociationsLifeCycleStateAll        ListAssociationsLifeCycleStateEnum = "ALL"
	ListAssociationsLifeCycleStateAccepted   ListAssociationsLifeCycleStateEnum = "ACCEPTED"
	ListAssociationsLifeCycleStateInProgress ListAssociationsLifeCycleStateEnum = "IN_PROGRESS"
	ListAssociationsLifeCycleStateSucceeded  ListAssociationsLifeCycleStateEnum = "SUCCEEDED"
	ListAssociationsLifeCycleStateFailed     ListAssociationsLifeCycleStateEnum = "FAILED"
)

var mappingListAssociationsLifeCycleStateEnum = map[string]ListAssociationsLifeCycleStateEnum{
	"ALL":         ListAssociationsLifeCycleStateAll,
	"ACCEPTED":    ListAssociationsLifeCycleStateAccepted,
	"IN_PROGRESS": ListAssociationsLifeCycleStateInProgress,
	"SUCCEEDED":   ListAssociationsLifeCycleStateSucceeded,
	"FAILED":      ListAssociationsLifeCycleStateFailed,
}

var mappingListAssociationsLifeCycleStateEnumLowerCase = map[string]ListAssociationsLifeCycleStateEnum{
	"all":         ListAssociationsLifeCycleStateAll,
	"accepted":    ListAssociationsLifeCycleStateAccepted,
	"in_progress": ListAssociationsLifeCycleStateInProgress,
	"succeeded":   ListAssociationsLifeCycleStateSucceeded,
	"failed":      ListAssociationsLifeCycleStateFailed,
}

// GetListAssociationsLifeCycleStateEnumValues Enumerates the set of values for ListAssociationsLifeCycleStateEnum
func GetListAssociationsLifeCycleStateEnumValues() []ListAssociationsLifeCycleStateEnum {
	values := make([]ListAssociationsLifeCycleStateEnum, 0)
	for _, v := range mappingListAssociationsLifeCycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociationsLifeCycleStateEnumStringValues Enumerates the set of values in String for ListAssociationsLifeCycleStateEnum
func GetListAssociationsLifeCycleStateEnumStringValues() []string {
	return []string{
		"ALL",
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingListAssociationsLifeCycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociationsLifeCycleStateEnum(val string) (ListAssociationsLifeCycleStateEnum, bool) {
	enum, ok := mappingListAssociationsLifeCycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociationsSortOrderEnum Enum with underlying type: string
type ListAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListAssociationsSortOrderEnum
const (
	ListAssociationsSortOrderAsc  ListAssociationsSortOrderEnum = "ASC"
	ListAssociationsSortOrderDesc ListAssociationsSortOrderEnum = "DESC"
)

var mappingListAssociationsSortOrderEnum = map[string]ListAssociationsSortOrderEnum{
	"ASC":  ListAssociationsSortOrderAsc,
	"DESC": ListAssociationsSortOrderDesc,
}

var mappingListAssociationsSortOrderEnumLowerCase = map[string]ListAssociationsSortOrderEnum{
	"asc":  ListAssociationsSortOrderAsc,
	"desc": ListAssociationsSortOrderDesc,
}

// GetListAssociationsSortOrderEnumValues Enumerates the set of values for ListAssociationsSortOrderEnum
func GetListAssociationsSortOrderEnumValues() []ListAssociationsSortOrderEnum {
	values := make([]ListAssociationsSortOrderEnum, 0)
	for _, v := range mappingListAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListAssociationsSortOrderEnum
func GetListAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociationsSortOrderEnum(val string) (ListAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociationsSortByEnum Enum with underlying type: string
type ListAssociationsSortByEnum string

// Set of constants representing the allowable values for ListAssociationsSortByEnum
const (
	ListAssociationsSortByTimelastattempted ListAssociationsSortByEnum = "timeLastAttempted"
	ListAssociationsSortByStatus            ListAssociationsSortByEnum = "status"
)

var mappingListAssociationsSortByEnum = map[string]ListAssociationsSortByEnum{
	"timeLastAttempted": ListAssociationsSortByTimelastattempted,
	"status":            ListAssociationsSortByStatus,
}

var mappingListAssociationsSortByEnumLowerCase = map[string]ListAssociationsSortByEnum{
	"timelastattempted": ListAssociationsSortByTimelastattempted,
	"status":            ListAssociationsSortByStatus,
}

// GetListAssociationsSortByEnumValues Enumerates the set of values for ListAssociationsSortByEnum
func GetListAssociationsSortByEnumValues() []ListAssociationsSortByEnum {
	values := make([]ListAssociationsSortByEnum, 0)
	for _, v := range mappingListAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociationsSortByEnumStringValues Enumerates the set of values in String for ListAssociationsSortByEnum
func GetListAssociationsSortByEnumStringValues() []string {
	return []string{
		"timeLastAttempted",
		"status",
	}
}

// GetMappingListAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociationsSortByEnum(val string) (ListAssociationsSortByEnum, bool) {
	enum, ok := mappingListAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
