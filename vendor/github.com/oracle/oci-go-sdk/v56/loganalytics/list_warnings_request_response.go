// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListWarningsRequest wrapper for the ListWarnings operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListWarnings.go.html to see an example of how to use ListWarningsRequest.
type ListWarningsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The warning state used for filtering.  A value of SUPPRESSED will return only
	// suppressed warnings, a value of UNSUPPRESSED will return only unsuppressed
	// warnings, and a value of ALL will return all warnings regardless of their
	// suppression state.  Default is UNSUPPRESSED.
	WarningState ListWarningsWarningStateEnum `mandatory:"false" contributesTo:"query" name:"warningState" omitEmpty:"true"`

	// The source name.
	SourceName *string `mandatory:"false" contributesTo:"query" name:"sourceName"`

	// The source pattern used for filtering.  Only warnings associated with a source with the
	// specified pattern will be returned.
	SourcePattern *string `mandatory:"false" contributesTo:"query" name:"sourcePattern"`

	// warning message query parameter
	WarningMessage *string `mandatory:"false" contributesTo:"query" name:"warningMessage"`

	// The entity name used for filtering.  Only warnings associated with an entity with the
	// specified name will be returned.
	EntityName *string `mandatory:"false" contributesTo:"query" name:"entityName"`

	// The entity type used for filtering.  Only associations on an entity with the
	// specified type will be returned.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// The warning type query parameter.
	WarningType *string `mandatory:"false" contributesTo:"query" name:"warningType"`

	// A flag indicating whether to filter warnings based on source display name or on warning level.
	// A value of true will filter based on warning level (rule, source, or pattern), while a
	// value of false will filter based on source display name.
	IsNoSource *bool `mandatory:"false" contributesTo:"query" name:"isNoSource"`

	// The warning start date query parameter.
	StartTime *string `mandatory:"false" contributesTo:"query" name:"startTime"`

	// The warning end date query parameter.
	EndTime *string `mandatory:"false" contributesTo:"query" name:"endTime"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListWarningsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned warnings
	SortBy ListWarningsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWarningsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWarningsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWarningsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWarningsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWarningsResponse wrapper for the ListWarnings operation
type ListWarningsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsWarningCollection instances
	LogAnalyticsWarningCollection `presentIn:"body"`

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

func (response ListWarningsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWarningsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWarningsWarningStateEnum Enum with underlying type: string
type ListWarningsWarningStateEnum string

// Set of constants representing the allowable values for ListWarningsWarningStateEnum
const (
	ListWarningsWarningStateAll          ListWarningsWarningStateEnum = "ALL"
	ListWarningsWarningStateSuppressed   ListWarningsWarningStateEnum = "SUPPRESSED"
	ListWarningsWarningStateUnsuppressed ListWarningsWarningStateEnum = "UNSUPPRESSED"
)

var mappingListWarningsWarningState = map[string]ListWarningsWarningStateEnum{
	"ALL":          ListWarningsWarningStateAll,
	"SUPPRESSED":   ListWarningsWarningStateSuppressed,
	"UNSUPPRESSED": ListWarningsWarningStateUnsuppressed,
}

// GetListWarningsWarningStateEnumValues Enumerates the set of values for ListWarningsWarningStateEnum
func GetListWarningsWarningStateEnumValues() []ListWarningsWarningStateEnum {
	values := make([]ListWarningsWarningStateEnum, 0)
	for _, v := range mappingListWarningsWarningState {
		values = append(values, v)
	}
	return values
}

// ListWarningsSortOrderEnum Enum with underlying type: string
type ListWarningsSortOrderEnum string

// Set of constants representing the allowable values for ListWarningsSortOrderEnum
const (
	ListWarningsSortOrderAsc  ListWarningsSortOrderEnum = "ASC"
	ListWarningsSortOrderDesc ListWarningsSortOrderEnum = "DESC"
)

var mappingListWarningsSortOrder = map[string]ListWarningsSortOrderEnum{
	"ASC":  ListWarningsSortOrderAsc,
	"DESC": ListWarningsSortOrderDesc,
}

// GetListWarningsSortOrderEnumValues Enumerates the set of values for ListWarningsSortOrderEnum
func GetListWarningsSortOrderEnumValues() []ListWarningsSortOrderEnum {
	values := make([]ListWarningsSortOrderEnum, 0)
	for _, v := range mappingListWarningsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWarningsSortByEnum Enum with underlying type: string
type ListWarningsSortByEnum string

// Set of constants representing the allowable values for ListWarningsSortByEnum
const (
	ListWarningsSortByEntitytype         ListWarningsSortByEnum = "EntityType"
	ListWarningsSortBySourcename         ListWarningsSortByEnum = "SourceName"
	ListWarningsSortByPatterntext        ListWarningsSortByEnum = "PatternText"
	ListWarningsSortByFirstreported      ListWarningsSortByEnum = "FirstReported"
	ListWarningsSortByWarningmessage     ListWarningsSortByEnum = "WarningMessage"
	ListWarningsSortByHost               ListWarningsSortByEnum = "Host"
	ListWarningsSortByEntityname         ListWarningsSortByEnum = "EntityName"
	ListWarningsSortByInitialwarningdate ListWarningsSortByEnum = "InitialWarningDate"
)

var mappingListWarningsSortBy = map[string]ListWarningsSortByEnum{
	"EntityType":         ListWarningsSortByEntitytype,
	"SourceName":         ListWarningsSortBySourcename,
	"PatternText":        ListWarningsSortByPatterntext,
	"FirstReported":      ListWarningsSortByFirstreported,
	"WarningMessage":     ListWarningsSortByWarningmessage,
	"Host":               ListWarningsSortByHost,
	"EntityName":         ListWarningsSortByEntityname,
	"InitialWarningDate": ListWarningsSortByInitialwarningdate,
}

// GetListWarningsSortByEnumValues Enumerates the set of values for ListWarningsSortByEnum
func GetListWarningsSortByEnumValues() []ListWarningsSortByEnum {
	values := make([]ListWarningsSortByEnum, 0)
	for _, v := range mappingListWarningsSortBy {
		values = append(values, v)
	}
	return values
}
