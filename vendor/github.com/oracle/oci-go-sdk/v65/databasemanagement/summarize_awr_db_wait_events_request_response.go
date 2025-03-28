// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeAwrDbWaitEventsRequest wrapper for the SummarizeAwrDbWaitEvents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbWaitEvents.go.html to see an example of how to use SummarizeAwrDbWaitEventsRequest.
type SummarizeAwrDbWaitEventsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The optional single value query parameter to filter the database instance number.
	InstNum *string `mandatory:"false" contributesTo:"query" name:"instNum"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnIdGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot ID.
	EndSnIdLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnIdLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional multiple value query parameter to filter the entity name.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The optional query parameter to filter ASH activities by FOREGROUND or BACKGROUND.
	SessionType SummarizeAwrDbWaitEventsSessionTypeEnum `mandatory:"false" contributesTo:"query" name:"sessionType" omitEmpty:"true"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in large paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the data within a time period.
	SortBy SummarizeAwrDbWaitEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbWaitEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDbWaitEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbWaitEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbWaitEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbWaitEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDbWaitEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDbWaitEventsSessionTypeEnum(string(request.SessionType)); !ok && request.SessionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionType: %s. Supported values are: %s.", request.SessionType, strings.Join(GetSummarizeAwrDbWaitEventsSessionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbWaitEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDbWaitEventsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbWaitEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDbWaitEventsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDbWaitEventsResponse wrapper for the SummarizeAwrDbWaitEvents operation
type SummarizeAwrDbWaitEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbWaitEventCollection instances
	AwrDbWaitEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbWaitEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbWaitEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbWaitEventsSessionTypeEnum Enum with underlying type: string
type SummarizeAwrDbWaitEventsSessionTypeEnum string

// Set of constants representing the allowable values for SummarizeAwrDbWaitEventsSessionTypeEnum
const (
	SummarizeAwrDbWaitEventsSessionTypeForeground SummarizeAwrDbWaitEventsSessionTypeEnum = "FOREGROUND"
	SummarizeAwrDbWaitEventsSessionTypeBackground SummarizeAwrDbWaitEventsSessionTypeEnum = "BACKGROUND"
	SummarizeAwrDbWaitEventsSessionTypeAll        SummarizeAwrDbWaitEventsSessionTypeEnum = "ALL"
)

var mappingSummarizeAwrDbWaitEventsSessionTypeEnum = map[string]SummarizeAwrDbWaitEventsSessionTypeEnum{
	"FOREGROUND": SummarizeAwrDbWaitEventsSessionTypeForeground,
	"BACKGROUND": SummarizeAwrDbWaitEventsSessionTypeBackground,
	"ALL":        SummarizeAwrDbWaitEventsSessionTypeAll,
}

var mappingSummarizeAwrDbWaitEventsSessionTypeEnumLowerCase = map[string]SummarizeAwrDbWaitEventsSessionTypeEnum{
	"foreground": SummarizeAwrDbWaitEventsSessionTypeForeground,
	"background": SummarizeAwrDbWaitEventsSessionTypeBackground,
	"all":        SummarizeAwrDbWaitEventsSessionTypeAll,
}

// GetSummarizeAwrDbWaitEventsSessionTypeEnumValues Enumerates the set of values for SummarizeAwrDbWaitEventsSessionTypeEnum
func GetSummarizeAwrDbWaitEventsSessionTypeEnumValues() []SummarizeAwrDbWaitEventsSessionTypeEnum {
	values := make([]SummarizeAwrDbWaitEventsSessionTypeEnum, 0)
	for _, v := range mappingSummarizeAwrDbWaitEventsSessionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbWaitEventsSessionTypeEnumStringValues Enumerates the set of values in String for SummarizeAwrDbWaitEventsSessionTypeEnum
func GetSummarizeAwrDbWaitEventsSessionTypeEnumStringValues() []string {
	return []string{
		"FOREGROUND",
		"BACKGROUND",
		"ALL",
	}
}

// GetMappingSummarizeAwrDbWaitEventsSessionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbWaitEventsSessionTypeEnum(val string) (SummarizeAwrDbWaitEventsSessionTypeEnum, bool) {
	enum, ok := mappingSummarizeAwrDbWaitEventsSessionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbWaitEventsSortByEnum Enum with underlying type: string
type SummarizeAwrDbWaitEventsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbWaitEventsSortByEnum
const (
	SummarizeAwrDbWaitEventsSortByTimeBegin SummarizeAwrDbWaitEventsSortByEnum = "TIME_BEGIN"
	SummarizeAwrDbWaitEventsSortByName      SummarizeAwrDbWaitEventsSortByEnum = "NAME"
)

var mappingSummarizeAwrDbWaitEventsSortByEnum = map[string]SummarizeAwrDbWaitEventsSortByEnum{
	"TIME_BEGIN": SummarizeAwrDbWaitEventsSortByTimeBegin,
	"NAME":       SummarizeAwrDbWaitEventsSortByName,
}

var mappingSummarizeAwrDbWaitEventsSortByEnumLowerCase = map[string]SummarizeAwrDbWaitEventsSortByEnum{
	"time_begin": SummarizeAwrDbWaitEventsSortByTimeBegin,
	"name":       SummarizeAwrDbWaitEventsSortByName,
}

// GetSummarizeAwrDbWaitEventsSortByEnumValues Enumerates the set of values for SummarizeAwrDbWaitEventsSortByEnum
func GetSummarizeAwrDbWaitEventsSortByEnumValues() []SummarizeAwrDbWaitEventsSortByEnum {
	values := make([]SummarizeAwrDbWaitEventsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbWaitEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbWaitEventsSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDbWaitEventsSortByEnum
func GetSummarizeAwrDbWaitEventsSortByEnumStringValues() []string {
	return []string{
		"TIME_BEGIN",
		"NAME",
	}
}

// GetMappingSummarizeAwrDbWaitEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbWaitEventsSortByEnum(val string) (SummarizeAwrDbWaitEventsSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDbWaitEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbWaitEventsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbWaitEventsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbWaitEventsSortOrderEnum
const (
	SummarizeAwrDbWaitEventsSortOrderAsc  SummarizeAwrDbWaitEventsSortOrderEnum = "ASC"
	SummarizeAwrDbWaitEventsSortOrderDesc SummarizeAwrDbWaitEventsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbWaitEventsSortOrderEnum = map[string]SummarizeAwrDbWaitEventsSortOrderEnum{
	"ASC":  SummarizeAwrDbWaitEventsSortOrderAsc,
	"DESC": SummarizeAwrDbWaitEventsSortOrderDesc,
}

var mappingSummarizeAwrDbWaitEventsSortOrderEnumLowerCase = map[string]SummarizeAwrDbWaitEventsSortOrderEnum{
	"asc":  SummarizeAwrDbWaitEventsSortOrderAsc,
	"desc": SummarizeAwrDbWaitEventsSortOrderDesc,
}

// GetSummarizeAwrDbWaitEventsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbWaitEventsSortOrderEnum
func GetSummarizeAwrDbWaitEventsSortOrderEnumValues() []SummarizeAwrDbWaitEventsSortOrderEnum {
	values := make([]SummarizeAwrDbWaitEventsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbWaitEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbWaitEventsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDbWaitEventsSortOrderEnum
func GetSummarizeAwrDbWaitEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDbWaitEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbWaitEventsSortOrderEnum(val string) (SummarizeAwrDbWaitEventsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDbWaitEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
