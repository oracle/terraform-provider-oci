// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeAwrDatabaseWaitEventsRequest wrapper for the SummarizeAwrDatabaseWaitEvents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseWaitEvents.go.html to see an example of how to use SummarizeAwrDatabaseWaitEventsRequest.
type SummarizeAwrDatabaseWaitEventsRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The optional single value query parameter to filter by database instance number.
	InstanceNumber *string `mandatory:"false" contributesTo:"query" name:"instanceNumber"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnapshotIdentifierGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnapshotIdentifierGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot Identifier.
	EndSnapshotIdentifierLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnapshotIdentifierLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional multiple value query parameter to filter the entity name.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The optional query parameter to filter ASH activities by FOREGROUND or BACKGROUND.
	SessionType SummarizeAwrDatabaseWaitEventsSessionTypeEnum `mandatory:"false" contributesTo:"query" name:"sessionType" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the data within a time period.
	SortBy SummarizeAwrDatabaseWaitEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseWaitEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseWaitEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseWaitEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseWaitEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseWaitEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseWaitEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseWaitEventsSessionTypeEnum(string(request.SessionType)); !ok && request.SessionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionType: %s. Supported values are: %s.", request.SessionType, strings.Join(GetSummarizeAwrDatabaseWaitEventsSessionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseWaitEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseWaitEventsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseWaitEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseWaitEventsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseWaitEventsResponse wrapper for the SummarizeAwrDatabaseWaitEvents operation
type SummarizeAwrDatabaseWaitEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseWaitEventCollection instances
	AwrDatabaseWaitEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseWaitEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseWaitEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseWaitEventsSessionTypeEnum Enum with underlying type: string
type SummarizeAwrDatabaseWaitEventsSessionTypeEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseWaitEventsSessionTypeEnum
const (
	SummarizeAwrDatabaseWaitEventsSessionTypeForeground SummarizeAwrDatabaseWaitEventsSessionTypeEnum = "FOREGROUND"
	SummarizeAwrDatabaseWaitEventsSessionTypeBackground SummarizeAwrDatabaseWaitEventsSessionTypeEnum = "BACKGROUND"
	SummarizeAwrDatabaseWaitEventsSessionTypeAll        SummarizeAwrDatabaseWaitEventsSessionTypeEnum = "ALL"
)

var mappingSummarizeAwrDatabaseWaitEventsSessionTypeEnum = map[string]SummarizeAwrDatabaseWaitEventsSessionTypeEnum{
	"FOREGROUND": SummarizeAwrDatabaseWaitEventsSessionTypeForeground,
	"BACKGROUND": SummarizeAwrDatabaseWaitEventsSessionTypeBackground,
	"ALL":        SummarizeAwrDatabaseWaitEventsSessionTypeAll,
}

var mappingSummarizeAwrDatabaseWaitEventsSessionTypeEnumLowerCase = map[string]SummarizeAwrDatabaseWaitEventsSessionTypeEnum{
	"foreground": SummarizeAwrDatabaseWaitEventsSessionTypeForeground,
	"background": SummarizeAwrDatabaseWaitEventsSessionTypeBackground,
	"all":        SummarizeAwrDatabaseWaitEventsSessionTypeAll,
}

// GetSummarizeAwrDatabaseWaitEventsSessionTypeEnumValues Enumerates the set of values for SummarizeAwrDatabaseWaitEventsSessionTypeEnum
func GetSummarizeAwrDatabaseWaitEventsSessionTypeEnumValues() []SummarizeAwrDatabaseWaitEventsSessionTypeEnum {
	values := make([]SummarizeAwrDatabaseWaitEventsSessionTypeEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseWaitEventsSessionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseWaitEventsSessionTypeEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseWaitEventsSessionTypeEnum
func GetSummarizeAwrDatabaseWaitEventsSessionTypeEnumStringValues() []string {
	return []string{
		"FOREGROUND",
		"BACKGROUND",
		"ALL",
	}
}

// GetMappingSummarizeAwrDatabaseWaitEventsSessionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseWaitEventsSessionTypeEnum(val string) (SummarizeAwrDatabaseWaitEventsSessionTypeEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseWaitEventsSessionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseWaitEventsSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseWaitEventsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseWaitEventsSortByEnum
const (
	SummarizeAwrDatabaseWaitEventsSortByTimeBegin SummarizeAwrDatabaseWaitEventsSortByEnum = "TIME_BEGIN"
	SummarizeAwrDatabaseWaitEventsSortByName      SummarizeAwrDatabaseWaitEventsSortByEnum = "NAME"
)

var mappingSummarizeAwrDatabaseWaitEventsSortByEnum = map[string]SummarizeAwrDatabaseWaitEventsSortByEnum{
	"TIME_BEGIN": SummarizeAwrDatabaseWaitEventsSortByTimeBegin,
	"NAME":       SummarizeAwrDatabaseWaitEventsSortByName,
}

var mappingSummarizeAwrDatabaseWaitEventsSortByEnumLowerCase = map[string]SummarizeAwrDatabaseWaitEventsSortByEnum{
	"time_begin": SummarizeAwrDatabaseWaitEventsSortByTimeBegin,
	"name":       SummarizeAwrDatabaseWaitEventsSortByName,
}

// GetSummarizeAwrDatabaseWaitEventsSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseWaitEventsSortByEnum
func GetSummarizeAwrDatabaseWaitEventsSortByEnumValues() []SummarizeAwrDatabaseWaitEventsSortByEnum {
	values := make([]SummarizeAwrDatabaseWaitEventsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseWaitEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseWaitEventsSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseWaitEventsSortByEnum
func GetSummarizeAwrDatabaseWaitEventsSortByEnumStringValues() []string {
	return []string{
		"TIME_BEGIN",
		"NAME",
	}
}

// GetMappingSummarizeAwrDatabaseWaitEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseWaitEventsSortByEnum(val string) (SummarizeAwrDatabaseWaitEventsSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseWaitEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseWaitEventsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseWaitEventsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseWaitEventsSortOrderEnum
const (
	SummarizeAwrDatabaseWaitEventsSortOrderAsc  SummarizeAwrDatabaseWaitEventsSortOrderEnum = "ASC"
	SummarizeAwrDatabaseWaitEventsSortOrderDesc SummarizeAwrDatabaseWaitEventsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseWaitEventsSortOrderEnum = map[string]SummarizeAwrDatabaseWaitEventsSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseWaitEventsSortOrderAsc,
	"DESC": SummarizeAwrDatabaseWaitEventsSortOrderDesc,
}

var mappingSummarizeAwrDatabaseWaitEventsSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseWaitEventsSortOrderEnum{
	"asc":  SummarizeAwrDatabaseWaitEventsSortOrderAsc,
	"desc": SummarizeAwrDatabaseWaitEventsSortOrderDesc,
}

// GetSummarizeAwrDatabaseWaitEventsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseWaitEventsSortOrderEnum
func GetSummarizeAwrDatabaseWaitEventsSortOrderEnumValues() []SummarizeAwrDatabaseWaitEventsSortOrderEnum {
	values := make([]SummarizeAwrDatabaseWaitEventsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseWaitEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseWaitEventsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseWaitEventsSortOrderEnum
func GetSummarizeAwrDatabaseWaitEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseWaitEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseWaitEventsSortOrderEnum(val string) (SummarizeAwrDatabaseWaitEventsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseWaitEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
