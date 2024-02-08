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

// SummarizeAwrDatabaseCpuUsagesRequest wrapper for the SummarizeAwrDatabaseCpuUsages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseCpuUsages.go.html to see an example of how to use SummarizeAwrDatabaseCpuUsagesRequest.
type SummarizeAwrDatabaseCpuUsagesRequest struct {

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
	SessionType SummarizeAwrDatabaseCpuUsagesSessionTypeEnum `mandatory:"false" contributesTo:"query" name:"sessionType" omitEmpty:"true"`

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

	// The option to sort the AWR CPU usage summary data.
	SortBy SummarizeAwrDatabaseCpuUsagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseCpuUsagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseCpuUsagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseCpuUsagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseCpuUsagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseCpuUsagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseCpuUsagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnum(string(request.SessionType)); !ok && request.SessionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionType: %s. Supported values are: %s.", request.SessionType, strings.Join(GetSummarizeAwrDatabaseCpuUsagesSessionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseCpuUsagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseCpuUsagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseCpuUsagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseCpuUsagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseCpuUsagesResponse wrapper for the SummarizeAwrDatabaseCpuUsages operation
type SummarizeAwrDatabaseCpuUsagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseCpuUsageCollection instances
	AwrDatabaseCpuUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseCpuUsagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseCpuUsagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseCpuUsagesSessionTypeEnum Enum with underlying type: string
type SummarizeAwrDatabaseCpuUsagesSessionTypeEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseCpuUsagesSessionTypeEnum
const (
	SummarizeAwrDatabaseCpuUsagesSessionTypeForeground SummarizeAwrDatabaseCpuUsagesSessionTypeEnum = "FOREGROUND"
	SummarizeAwrDatabaseCpuUsagesSessionTypeBackground SummarizeAwrDatabaseCpuUsagesSessionTypeEnum = "BACKGROUND"
	SummarizeAwrDatabaseCpuUsagesSessionTypeAll        SummarizeAwrDatabaseCpuUsagesSessionTypeEnum = "ALL"
)

var mappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnum = map[string]SummarizeAwrDatabaseCpuUsagesSessionTypeEnum{
	"FOREGROUND": SummarizeAwrDatabaseCpuUsagesSessionTypeForeground,
	"BACKGROUND": SummarizeAwrDatabaseCpuUsagesSessionTypeBackground,
	"ALL":        SummarizeAwrDatabaseCpuUsagesSessionTypeAll,
}

var mappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnumLowerCase = map[string]SummarizeAwrDatabaseCpuUsagesSessionTypeEnum{
	"foreground": SummarizeAwrDatabaseCpuUsagesSessionTypeForeground,
	"background": SummarizeAwrDatabaseCpuUsagesSessionTypeBackground,
	"all":        SummarizeAwrDatabaseCpuUsagesSessionTypeAll,
}

// GetSummarizeAwrDatabaseCpuUsagesSessionTypeEnumValues Enumerates the set of values for SummarizeAwrDatabaseCpuUsagesSessionTypeEnum
func GetSummarizeAwrDatabaseCpuUsagesSessionTypeEnumValues() []SummarizeAwrDatabaseCpuUsagesSessionTypeEnum {
	values := make([]SummarizeAwrDatabaseCpuUsagesSessionTypeEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseCpuUsagesSessionTypeEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseCpuUsagesSessionTypeEnum
func GetSummarizeAwrDatabaseCpuUsagesSessionTypeEnumStringValues() []string {
	return []string{
		"FOREGROUND",
		"BACKGROUND",
		"ALL",
	}
}

// GetMappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnum(val string) (SummarizeAwrDatabaseCpuUsagesSessionTypeEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseCpuUsagesSessionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseCpuUsagesSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseCpuUsagesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseCpuUsagesSortByEnum
const (
	SummarizeAwrDatabaseCpuUsagesSortByTimeSampled SummarizeAwrDatabaseCpuUsagesSortByEnum = "TIME_SAMPLED"
	SummarizeAwrDatabaseCpuUsagesSortByAvgValue    SummarizeAwrDatabaseCpuUsagesSortByEnum = "AVG_VALUE"
)

var mappingSummarizeAwrDatabaseCpuUsagesSortByEnum = map[string]SummarizeAwrDatabaseCpuUsagesSortByEnum{
	"TIME_SAMPLED": SummarizeAwrDatabaseCpuUsagesSortByTimeSampled,
	"AVG_VALUE":    SummarizeAwrDatabaseCpuUsagesSortByAvgValue,
}

var mappingSummarizeAwrDatabaseCpuUsagesSortByEnumLowerCase = map[string]SummarizeAwrDatabaseCpuUsagesSortByEnum{
	"time_sampled": SummarizeAwrDatabaseCpuUsagesSortByTimeSampled,
	"avg_value":    SummarizeAwrDatabaseCpuUsagesSortByAvgValue,
}

// GetSummarizeAwrDatabaseCpuUsagesSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseCpuUsagesSortByEnum
func GetSummarizeAwrDatabaseCpuUsagesSortByEnumValues() []SummarizeAwrDatabaseCpuUsagesSortByEnum {
	values := make([]SummarizeAwrDatabaseCpuUsagesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseCpuUsagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseCpuUsagesSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseCpuUsagesSortByEnum
func GetSummarizeAwrDatabaseCpuUsagesSortByEnumStringValues() []string {
	return []string{
		"TIME_SAMPLED",
		"AVG_VALUE",
	}
}

// GetMappingSummarizeAwrDatabaseCpuUsagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseCpuUsagesSortByEnum(val string) (SummarizeAwrDatabaseCpuUsagesSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseCpuUsagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseCpuUsagesSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseCpuUsagesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseCpuUsagesSortOrderEnum
const (
	SummarizeAwrDatabaseCpuUsagesSortOrderAsc  SummarizeAwrDatabaseCpuUsagesSortOrderEnum = "ASC"
	SummarizeAwrDatabaseCpuUsagesSortOrderDesc SummarizeAwrDatabaseCpuUsagesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseCpuUsagesSortOrderEnum = map[string]SummarizeAwrDatabaseCpuUsagesSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseCpuUsagesSortOrderAsc,
	"DESC": SummarizeAwrDatabaseCpuUsagesSortOrderDesc,
}

var mappingSummarizeAwrDatabaseCpuUsagesSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseCpuUsagesSortOrderEnum{
	"asc":  SummarizeAwrDatabaseCpuUsagesSortOrderAsc,
	"desc": SummarizeAwrDatabaseCpuUsagesSortOrderDesc,
}

// GetSummarizeAwrDatabaseCpuUsagesSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseCpuUsagesSortOrderEnum
func GetSummarizeAwrDatabaseCpuUsagesSortOrderEnumValues() []SummarizeAwrDatabaseCpuUsagesSortOrderEnum {
	values := make([]SummarizeAwrDatabaseCpuUsagesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseCpuUsagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseCpuUsagesSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseCpuUsagesSortOrderEnum
func GetSummarizeAwrDatabaseCpuUsagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseCpuUsagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseCpuUsagesSortOrderEnum(val string) (SummarizeAwrDatabaseCpuUsagesSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseCpuUsagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
