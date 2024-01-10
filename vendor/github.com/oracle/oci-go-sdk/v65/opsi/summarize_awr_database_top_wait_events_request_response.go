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

// SummarizeAwrDatabaseTopWaitEventsRequest wrapper for the SummarizeAwrDatabaseTopWaitEvents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseTopWaitEvents.go.html to see an example of how to use SummarizeAwrDatabaseTopWaitEventsRequest.
type SummarizeAwrDatabaseTopWaitEventsRequest struct {

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

	// The optional query parameter to filter ASH activities by FOREGROUND or BACKGROUND.
	SessionType SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum `mandatory:"false" contributesTo:"query" name:"sessionType" omitEmpty:"true"`

	// The optional query parameter to filter the number of top categories to be returned.
	TopN *int `mandatory:"false" contributesTo:"query" name:"topN"`

	// The option to sort the AWR top event summary data.
	SortBy SummarizeAwrDatabaseTopWaitEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseTopWaitEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseTopWaitEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseTopWaitEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseTopWaitEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseTopWaitEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseTopWaitEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnum(string(request.SessionType)); !ok && request.SessionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionType: %s. Supported values are: %s.", request.SessionType, strings.Join(GetSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseTopWaitEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseTopWaitEventsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseTopWaitEventsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseTopWaitEventsResponse wrapper for the SummarizeAwrDatabaseTopWaitEvents operation
type SummarizeAwrDatabaseTopWaitEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDatabaseTopWaitEventCollection instance
	AwrDatabaseTopWaitEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseTopWaitEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseTopWaitEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum Enum with underlying type: string
type SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum
const (
	SummarizeAwrDatabaseTopWaitEventsSessionTypeForeground SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum = "FOREGROUND"
	SummarizeAwrDatabaseTopWaitEventsSessionTypeBackground SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum = "BACKGROUND"
	SummarizeAwrDatabaseTopWaitEventsSessionTypeAll        SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum = "ALL"
)

var mappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnum = map[string]SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum{
	"FOREGROUND": SummarizeAwrDatabaseTopWaitEventsSessionTypeForeground,
	"BACKGROUND": SummarizeAwrDatabaseTopWaitEventsSessionTypeBackground,
	"ALL":        SummarizeAwrDatabaseTopWaitEventsSessionTypeAll,
}

var mappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumLowerCase = map[string]SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum{
	"foreground": SummarizeAwrDatabaseTopWaitEventsSessionTypeForeground,
	"background": SummarizeAwrDatabaseTopWaitEventsSessionTypeBackground,
	"all":        SummarizeAwrDatabaseTopWaitEventsSessionTypeAll,
}

// GetSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumValues Enumerates the set of values for SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum
func GetSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumValues() []SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum {
	values := make([]SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum
func GetSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumStringValues() []string {
	return []string{
		"FOREGROUND",
		"BACKGROUND",
		"ALL",
	}
}

// GetMappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnum(val string) (SummarizeAwrDatabaseTopWaitEventsSessionTypeEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseTopWaitEventsSessionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseTopWaitEventsSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseTopWaitEventsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseTopWaitEventsSortByEnum
const (
	SummarizeAwrDatabaseTopWaitEventsSortByWaitsPersec       SummarizeAwrDatabaseTopWaitEventsSortByEnum = "WAITS_PERSEC"
	SummarizeAwrDatabaseTopWaitEventsSortByAvgWaitTimePersec SummarizeAwrDatabaseTopWaitEventsSortByEnum = "AVG_WAIT_TIME_PERSEC"
)

var mappingSummarizeAwrDatabaseTopWaitEventsSortByEnum = map[string]SummarizeAwrDatabaseTopWaitEventsSortByEnum{
	"WAITS_PERSEC":         SummarizeAwrDatabaseTopWaitEventsSortByWaitsPersec,
	"AVG_WAIT_TIME_PERSEC": SummarizeAwrDatabaseTopWaitEventsSortByAvgWaitTimePersec,
}

var mappingSummarizeAwrDatabaseTopWaitEventsSortByEnumLowerCase = map[string]SummarizeAwrDatabaseTopWaitEventsSortByEnum{
	"waits_persec":         SummarizeAwrDatabaseTopWaitEventsSortByWaitsPersec,
	"avg_wait_time_persec": SummarizeAwrDatabaseTopWaitEventsSortByAvgWaitTimePersec,
}

// GetSummarizeAwrDatabaseTopWaitEventsSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseTopWaitEventsSortByEnum
func GetSummarizeAwrDatabaseTopWaitEventsSortByEnumValues() []SummarizeAwrDatabaseTopWaitEventsSortByEnum {
	values := make([]SummarizeAwrDatabaseTopWaitEventsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseTopWaitEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseTopWaitEventsSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseTopWaitEventsSortByEnum
func GetSummarizeAwrDatabaseTopWaitEventsSortByEnumStringValues() []string {
	return []string{
		"WAITS_PERSEC",
		"AVG_WAIT_TIME_PERSEC",
	}
}

// GetMappingSummarizeAwrDatabaseTopWaitEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseTopWaitEventsSortByEnum(val string) (SummarizeAwrDatabaseTopWaitEventsSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseTopWaitEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseTopWaitEventsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseTopWaitEventsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseTopWaitEventsSortOrderEnum
const (
	SummarizeAwrDatabaseTopWaitEventsSortOrderAsc  SummarizeAwrDatabaseTopWaitEventsSortOrderEnum = "ASC"
	SummarizeAwrDatabaseTopWaitEventsSortOrderDesc SummarizeAwrDatabaseTopWaitEventsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnum = map[string]SummarizeAwrDatabaseTopWaitEventsSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseTopWaitEventsSortOrderAsc,
	"DESC": SummarizeAwrDatabaseTopWaitEventsSortOrderDesc,
}

var mappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseTopWaitEventsSortOrderEnum{
	"asc":  SummarizeAwrDatabaseTopWaitEventsSortOrderAsc,
	"desc": SummarizeAwrDatabaseTopWaitEventsSortOrderDesc,
}

// GetSummarizeAwrDatabaseTopWaitEventsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseTopWaitEventsSortOrderEnum
func GetSummarizeAwrDatabaseTopWaitEventsSortOrderEnumValues() []SummarizeAwrDatabaseTopWaitEventsSortOrderEnum {
	values := make([]SummarizeAwrDatabaseTopWaitEventsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseTopWaitEventsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseTopWaitEventsSortOrderEnum
func GetSummarizeAwrDatabaseTopWaitEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnum(val string) (SummarizeAwrDatabaseTopWaitEventsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseTopWaitEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
