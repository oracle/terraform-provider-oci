// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// SummarizeAwrDbCpuUsagesRequest wrapper for the SummarizeAwrDbCpuUsages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbCpuUsages.go.html to see an example of how to use SummarizeAwrDbCpuUsagesRequest.
type SummarizeAwrDbCpuUsagesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
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

	// The optional query parameter to filter ASH activities by FOREGROUND or BACKGROUND.
	SessionType SummarizeAwrDbCpuUsagesSessionTypeEnum `mandatory:"false" contributesTo:"query" name:"sessionType" omitEmpty:"true"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in large paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR CPU usage summary data.
	SortBy SummarizeAwrDbCpuUsagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbCpuUsagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDbCpuUsagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbCpuUsagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbCpuUsagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbCpuUsagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDbCpuUsagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDbCpuUsagesSessionTypeEnum(string(request.SessionType)); !ok && request.SessionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionType: %s. Supported values are: %s.", request.SessionType, strings.Join(GetSummarizeAwrDbCpuUsagesSessionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbCpuUsagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDbCpuUsagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbCpuUsagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDbCpuUsagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDbCpuUsagesResponse wrapper for the SummarizeAwrDbCpuUsages operation
type SummarizeAwrDbCpuUsagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbCpuUsageCollection instances
	AwrDbCpuUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbCpuUsagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbCpuUsagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbCpuUsagesSessionTypeEnum Enum with underlying type: string
type SummarizeAwrDbCpuUsagesSessionTypeEnum string

// Set of constants representing the allowable values for SummarizeAwrDbCpuUsagesSessionTypeEnum
const (
	SummarizeAwrDbCpuUsagesSessionTypeForeground SummarizeAwrDbCpuUsagesSessionTypeEnum = "FOREGROUND"
	SummarizeAwrDbCpuUsagesSessionTypeBackground SummarizeAwrDbCpuUsagesSessionTypeEnum = "BACKGROUND"
	SummarizeAwrDbCpuUsagesSessionTypeAll        SummarizeAwrDbCpuUsagesSessionTypeEnum = "ALL"
)

var mappingSummarizeAwrDbCpuUsagesSessionTypeEnum = map[string]SummarizeAwrDbCpuUsagesSessionTypeEnum{
	"FOREGROUND": SummarizeAwrDbCpuUsagesSessionTypeForeground,
	"BACKGROUND": SummarizeAwrDbCpuUsagesSessionTypeBackground,
	"ALL":        SummarizeAwrDbCpuUsagesSessionTypeAll,
}

// GetSummarizeAwrDbCpuUsagesSessionTypeEnumValues Enumerates the set of values for SummarizeAwrDbCpuUsagesSessionTypeEnum
func GetSummarizeAwrDbCpuUsagesSessionTypeEnumValues() []SummarizeAwrDbCpuUsagesSessionTypeEnum {
	values := make([]SummarizeAwrDbCpuUsagesSessionTypeEnum, 0)
	for _, v := range mappingSummarizeAwrDbCpuUsagesSessionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbCpuUsagesSessionTypeEnumStringValues Enumerates the set of values in String for SummarizeAwrDbCpuUsagesSessionTypeEnum
func GetSummarizeAwrDbCpuUsagesSessionTypeEnumStringValues() []string {
	return []string{
		"FOREGROUND",
		"BACKGROUND",
		"ALL",
	}
}

// GetMappingSummarizeAwrDbCpuUsagesSessionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbCpuUsagesSessionTypeEnum(val string) (SummarizeAwrDbCpuUsagesSessionTypeEnum, bool) {
	mappingSummarizeAwrDbCpuUsagesSessionTypeEnumIgnoreCase := make(map[string]SummarizeAwrDbCpuUsagesSessionTypeEnum)
	for k, v := range mappingSummarizeAwrDbCpuUsagesSessionTypeEnum {
		mappingSummarizeAwrDbCpuUsagesSessionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeAwrDbCpuUsagesSessionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbCpuUsagesSortByEnum Enum with underlying type: string
type SummarizeAwrDbCpuUsagesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbCpuUsagesSortByEnum
const (
	SummarizeAwrDbCpuUsagesSortByTimeSampled SummarizeAwrDbCpuUsagesSortByEnum = "TIME_SAMPLED"
	SummarizeAwrDbCpuUsagesSortByAvgValue    SummarizeAwrDbCpuUsagesSortByEnum = "AVG_VALUE"
)

var mappingSummarizeAwrDbCpuUsagesSortByEnum = map[string]SummarizeAwrDbCpuUsagesSortByEnum{
	"TIME_SAMPLED": SummarizeAwrDbCpuUsagesSortByTimeSampled,
	"AVG_VALUE":    SummarizeAwrDbCpuUsagesSortByAvgValue,
}

// GetSummarizeAwrDbCpuUsagesSortByEnumValues Enumerates the set of values for SummarizeAwrDbCpuUsagesSortByEnum
func GetSummarizeAwrDbCpuUsagesSortByEnumValues() []SummarizeAwrDbCpuUsagesSortByEnum {
	values := make([]SummarizeAwrDbCpuUsagesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbCpuUsagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbCpuUsagesSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDbCpuUsagesSortByEnum
func GetSummarizeAwrDbCpuUsagesSortByEnumStringValues() []string {
	return []string{
		"TIME_SAMPLED",
		"AVG_VALUE",
	}
}

// GetMappingSummarizeAwrDbCpuUsagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbCpuUsagesSortByEnum(val string) (SummarizeAwrDbCpuUsagesSortByEnum, bool) {
	mappingSummarizeAwrDbCpuUsagesSortByEnumIgnoreCase := make(map[string]SummarizeAwrDbCpuUsagesSortByEnum)
	for k, v := range mappingSummarizeAwrDbCpuUsagesSortByEnum {
		mappingSummarizeAwrDbCpuUsagesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeAwrDbCpuUsagesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbCpuUsagesSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbCpuUsagesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbCpuUsagesSortOrderEnum
const (
	SummarizeAwrDbCpuUsagesSortOrderAsc  SummarizeAwrDbCpuUsagesSortOrderEnum = "ASC"
	SummarizeAwrDbCpuUsagesSortOrderDesc SummarizeAwrDbCpuUsagesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbCpuUsagesSortOrderEnum = map[string]SummarizeAwrDbCpuUsagesSortOrderEnum{
	"ASC":  SummarizeAwrDbCpuUsagesSortOrderAsc,
	"DESC": SummarizeAwrDbCpuUsagesSortOrderDesc,
}

// GetSummarizeAwrDbCpuUsagesSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbCpuUsagesSortOrderEnum
func GetSummarizeAwrDbCpuUsagesSortOrderEnumValues() []SummarizeAwrDbCpuUsagesSortOrderEnum {
	values := make([]SummarizeAwrDbCpuUsagesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbCpuUsagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbCpuUsagesSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDbCpuUsagesSortOrderEnum
func GetSummarizeAwrDbCpuUsagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDbCpuUsagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbCpuUsagesSortOrderEnum(val string) (SummarizeAwrDbCpuUsagesSortOrderEnum, bool) {
	mappingSummarizeAwrDbCpuUsagesSortOrderEnumIgnoreCase := make(map[string]SummarizeAwrDbCpuUsagesSortOrderEnum)
	for k, v := range mappingSummarizeAwrDbCpuUsagesSortOrderEnum {
		mappingSummarizeAwrDbCpuUsagesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeAwrDbCpuUsagesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
