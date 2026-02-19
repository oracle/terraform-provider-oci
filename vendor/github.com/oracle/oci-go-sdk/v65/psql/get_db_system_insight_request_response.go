// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetDbSystemInsightRequest wrapper for the GetDbSystemInsight operation
type GetDbSystemInsightRequest struct {

	// A unique identifier for the database system.
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// High-level category of insight data.
	// Known values:
	// - QUERY_INSIGHT: Query performance and SQL-related insights
	InsightType GetDbSystemInsightInsightTypeEnum `mandatory:"true" contributesTo:"path" name:"insightType"`

	// Specific subtype of insight data within the selected insightType.
	// Supported values depend on the insightType:
	// When insightType = QUERY_INSIGHT:
	// - AAS_TIME_SERIES:      Average Active Sessions time-series
	// - TOP_QUERIES:          Top SQL queries ranked by performance
	InsightDataType GetDbSystemInsightInsightDataTypeEnum `mandatory:"true" contributesTo:"query" name:"insightDataType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The start date for getting backups. An RFC 3339 (https://tools.ietf.org/rfc/rfc3339) formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// The end date for getting backups. An RFC 3339 (https://tools.ietf.org/rfc/rfc3339) formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The type of wait event used to filter insight data.
	// This parameter is applicable only for specific insight combinations:
	//     - insightType: QUERY_INSIGHT
	//       insightDataType:
	//           - TOP_QUERIES
	// `Activity`: The server process is idle or waiting for background tasks.
	// `BufferPin`: Waiting for exclusive access to a data buffer in memory.
	// `Client`: Waiting for data from the client (e.g., query submission).
	// `CPU`: Active processing time (thread is running, not waiting).
	// `Extension`: Waiting for an extension-specific event or condition.
	// `IO`: Waiting for a read/write operation on a data file or WAL.
	// `IPC`: Waiting for inter-process communication (e.g., shared memory).
	// `Lock`: Waiting for a heavyweight SQL-level lock (e.g., table lock).
	// `LWLock`: Waiting for a lightweight lock on internal data structures.
	// `Timeout`: Waiting for a specialized timeout to expire.
	WaitEventType GetDbSystemInsightWaitEventTypeEnum `mandatory:"false" contributesTo:"query" name:"waitEventType" omitEmpty:"true"`

	// Name of the database used to filter insight data.
	// This parameter is applicable only for specific insight combinations:
	//     - insightType: QUERY_INSIGHT
	//       insightDataType:
	//           - TOP_QUERIES
	DatabaseName *string `mandatory:"false" contributesTo:"query" name:"databaseName"`

	// SQL query text or partial query text used to filter insight data.
	// This parameter is applicable only for specific insight combinations:
	//   - insightType: QUERY_INSIGHT
	//     insightDataType:
	//         - TOP_QUERIES
	QueryText *string `mandatory:"false" contributesTo:"query" name:"queryText"`

	// Role of the database instance used to filter insight data.
	// This parameter is applicable only for specific insight combinations:
	//  - insightType: QUERY_INSIGHT
	//    insightDataType:
	//        - TOP_QUERIES
	DbInstanceRole GetDbSystemInsightDbInstanceRoleEnum `mandatory:"false" contributesTo:"query" name:"dbInstanceRole" omitEmpty:"true"`

	// Unique identifier of the database instance node for which insights are requested.
	// This parameter is applicable only for specific insight combinations:
	//     - insightType: QUERY_INSIGHT
	//       insightDataType:
	//          - TOP_QUERIES
	DbInstanceId *string `mandatory:"false" contributesTo:"query" name:"dbInstanceId"`

	// The field to sort the results by following known values
	// - AAS: sorted in descending order by default
	// - AVG_EXECUTION_TIME: sorted in descending order by default
	// - QUERY_COUNT: sorted in descending order by default
	// This parameter is applicable only for specific insight combinations:
	//   - insightType: QUERY_INSIGHT
	//     insightDataType:
	//         - TOP_QUERIES
	SortInsightDataBy GetDbSystemInsightSortInsightDataByEnum `mandatory:"false" contributesTo:"query" name:"sortInsightDataBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder GetDbSystemInsightSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDbSystemInsightRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDbSystemInsightRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDbSystemInsightRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDbSystemInsightRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDbSystemInsightRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetDbSystemInsightInsightTypeEnum(string(request.InsightType)); !ok && request.InsightType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightType: %s. Supported values are: %s.", request.InsightType, strings.Join(GetGetDbSystemInsightInsightTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDbSystemInsightInsightDataTypeEnum(string(request.InsightDataType)); !ok && request.InsightDataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightDataType: %s. Supported values are: %s.", request.InsightDataType, strings.Join(GetGetDbSystemInsightInsightDataTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDbSystemInsightWaitEventTypeEnum(string(request.WaitEventType)); !ok && request.WaitEventType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WaitEventType: %s. Supported values are: %s.", request.WaitEventType, strings.Join(GetGetDbSystemInsightWaitEventTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDbSystemInsightDbInstanceRoleEnum(string(request.DbInstanceRole)); !ok && request.DbInstanceRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbInstanceRole: %s. Supported values are: %s.", request.DbInstanceRole, strings.Join(GetGetDbSystemInsightDbInstanceRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDbSystemInsightSortInsightDataByEnum(string(request.SortInsightDataBy)); !ok && request.SortInsightDataBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortInsightDataBy: %s. Supported values are: %s.", request.SortInsightDataBy, strings.Join(GetGetDbSystemInsightSortInsightDataByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDbSystemInsightSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetDbSystemInsightSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDbSystemInsightResponse wrapper for the GetDbSystemInsight operation
type GetDbSystemInsightResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbSystemInsight instances
	DbSystemInsight `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetDbSystemInsightResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDbSystemInsightResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDbSystemInsightInsightTypeEnum Enum with underlying type: string
type GetDbSystemInsightInsightTypeEnum string

// Set of constants representing the allowable values for GetDbSystemInsightInsightTypeEnum
const (
	GetDbSystemInsightInsightTypeQueryInsight GetDbSystemInsightInsightTypeEnum = "QUERY_INSIGHT"
)

var mappingGetDbSystemInsightInsightTypeEnum = map[string]GetDbSystemInsightInsightTypeEnum{
	"QUERY_INSIGHT": GetDbSystemInsightInsightTypeQueryInsight,
}

var mappingGetDbSystemInsightInsightTypeEnumLowerCase = map[string]GetDbSystemInsightInsightTypeEnum{
	"query_insight": GetDbSystemInsightInsightTypeQueryInsight,
}

// GetGetDbSystemInsightInsightTypeEnumValues Enumerates the set of values for GetDbSystemInsightInsightTypeEnum
func GetGetDbSystemInsightInsightTypeEnumValues() []GetDbSystemInsightInsightTypeEnum {
	values := make([]GetDbSystemInsightInsightTypeEnum, 0)
	for _, v := range mappingGetDbSystemInsightInsightTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemInsightInsightTypeEnumStringValues Enumerates the set of values in String for GetDbSystemInsightInsightTypeEnum
func GetGetDbSystemInsightInsightTypeEnumStringValues() []string {
	return []string{
		"QUERY_INSIGHT",
	}
}

// GetMappingGetDbSystemInsightInsightTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemInsightInsightTypeEnum(val string) (GetDbSystemInsightInsightTypeEnum, bool) {
	enum, ok := mappingGetDbSystemInsightInsightTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDbSystemInsightInsightDataTypeEnum Enum with underlying type: string
type GetDbSystemInsightInsightDataTypeEnum string

// Set of constants representing the allowable values for GetDbSystemInsightInsightDataTypeEnum
const (
	GetDbSystemInsightInsightDataTypeAasTimeSeries GetDbSystemInsightInsightDataTypeEnum = "AAS_TIME_SERIES"
	GetDbSystemInsightInsightDataTypeTopQueries    GetDbSystemInsightInsightDataTypeEnum = "TOP_QUERIES"
)

var mappingGetDbSystemInsightInsightDataTypeEnum = map[string]GetDbSystemInsightInsightDataTypeEnum{
	"AAS_TIME_SERIES": GetDbSystemInsightInsightDataTypeAasTimeSeries,
	"TOP_QUERIES":     GetDbSystemInsightInsightDataTypeTopQueries,
}

var mappingGetDbSystemInsightInsightDataTypeEnumLowerCase = map[string]GetDbSystemInsightInsightDataTypeEnum{
	"aas_time_series": GetDbSystemInsightInsightDataTypeAasTimeSeries,
	"top_queries":     GetDbSystemInsightInsightDataTypeTopQueries,
}

// GetGetDbSystemInsightInsightDataTypeEnumValues Enumerates the set of values for GetDbSystemInsightInsightDataTypeEnum
func GetGetDbSystemInsightInsightDataTypeEnumValues() []GetDbSystemInsightInsightDataTypeEnum {
	values := make([]GetDbSystemInsightInsightDataTypeEnum, 0)
	for _, v := range mappingGetDbSystemInsightInsightDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemInsightInsightDataTypeEnumStringValues Enumerates the set of values in String for GetDbSystemInsightInsightDataTypeEnum
func GetGetDbSystemInsightInsightDataTypeEnumStringValues() []string {
	return []string{
		"AAS_TIME_SERIES",
		"TOP_QUERIES",
	}
}

// GetMappingGetDbSystemInsightInsightDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemInsightInsightDataTypeEnum(val string) (GetDbSystemInsightInsightDataTypeEnum, bool) {
	enum, ok := mappingGetDbSystemInsightInsightDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDbSystemInsightWaitEventTypeEnum Enum with underlying type: string
type GetDbSystemInsightWaitEventTypeEnum string

// Set of constants representing the allowable values for GetDbSystemInsightWaitEventTypeEnum
const (
	GetDbSystemInsightWaitEventTypeActivity  GetDbSystemInsightWaitEventTypeEnum = "Activity"
	GetDbSystemInsightWaitEventTypeBufferpin GetDbSystemInsightWaitEventTypeEnum = "BufferPin"
	GetDbSystemInsightWaitEventTypeClient    GetDbSystemInsightWaitEventTypeEnum = "Client"
	GetDbSystemInsightWaitEventTypeCpu       GetDbSystemInsightWaitEventTypeEnum = "CPU"
	GetDbSystemInsightWaitEventTypeExtension GetDbSystemInsightWaitEventTypeEnum = "Extension"
	GetDbSystemInsightWaitEventTypeIo        GetDbSystemInsightWaitEventTypeEnum = "IO"
	GetDbSystemInsightWaitEventTypeIpc       GetDbSystemInsightWaitEventTypeEnum = "IPC"
	GetDbSystemInsightWaitEventTypeLock      GetDbSystemInsightWaitEventTypeEnum = "Lock"
	GetDbSystemInsightWaitEventTypeLwlock    GetDbSystemInsightWaitEventTypeEnum = "LWLock"
	GetDbSystemInsightWaitEventTypeTimeout   GetDbSystemInsightWaitEventTypeEnum = "Timeout"
)

var mappingGetDbSystemInsightWaitEventTypeEnum = map[string]GetDbSystemInsightWaitEventTypeEnum{
	"Activity":  GetDbSystemInsightWaitEventTypeActivity,
	"BufferPin": GetDbSystemInsightWaitEventTypeBufferpin,
	"Client":    GetDbSystemInsightWaitEventTypeClient,
	"CPU":       GetDbSystemInsightWaitEventTypeCpu,
	"Extension": GetDbSystemInsightWaitEventTypeExtension,
	"IO":        GetDbSystemInsightWaitEventTypeIo,
	"IPC":       GetDbSystemInsightWaitEventTypeIpc,
	"Lock":      GetDbSystemInsightWaitEventTypeLock,
	"LWLock":    GetDbSystemInsightWaitEventTypeLwlock,
	"Timeout":   GetDbSystemInsightWaitEventTypeTimeout,
}

var mappingGetDbSystemInsightWaitEventTypeEnumLowerCase = map[string]GetDbSystemInsightWaitEventTypeEnum{
	"activity":  GetDbSystemInsightWaitEventTypeActivity,
	"bufferpin": GetDbSystemInsightWaitEventTypeBufferpin,
	"client":    GetDbSystemInsightWaitEventTypeClient,
	"cpu":       GetDbSystemInsightWaitEventTypeCpu,
	"extension": GetDbSystemInsightWaitEventTypeExtension,
	"io":        GetDbSystemInsightWaitEventTypeIo,
	"ipc":       GetDbSystemInsightWaitEventTypeIpc,
	"lock":      GetDbSystemInsightWaitEventTypeLock,
	"lwlock":    GetDbSystemInsightWaitEventTypeLwlock,
	"timeout":   GetDbSystemInsightWaitEventTypeTimeout,
}

// GetGetDbSystemInsightWaitEventTypeEnumValues Enumerates the set of values for GetDbSystemInsightWaitEventTypeEnum
func GetGetDbSystemInsightWaitEventTypeEnumValues() []GetDbSystemInsightWaitEventTypeEnum {
	values := make([]GetDbSystemInsightWaitEventTypeEnum, 0)
	for _, v := range mappingGetDbSystemInsightWaitEventTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemInsightWaitEventTypeEnumStringValues Enumerates the set of values in String for GetDbSystemInsightWaitEventTypeEnum
func GetGetDbSystemInsightWaitEventTypeEnumStringValues() []string {
	return []string{
		"Activity",
		"BufferPin",
		"Client",
		"CPU",
		"Extension",
		"IO",
		"IPC",
		"Lock",
		"LWLock",
		"Timeout",
	}
}

// GetMappingGetDbSystemInsightWaitEventTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemInsightWaitEventTypeEnum(val string) (GetDbSystemInsightWaitEventTypeEnum, bool) {
	enum, ok := mappingGetDbSystemInsightWaitEventTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDbSystemInsightDbInstanceRoleEnum Enum with underlying type: string
type GetDbSystemInsightDbInstanceRoleEnum string

// Set of constants representing the allowable values for GetDbSystemInsightDbInstanceRoleEnum
const (
	GetDbSystemInsightDbInstanceRolePrimary GetDbSystemInsightDbInstanceRoleEnum = "PRIMARY"
	GetDbSystemInsightDbInstanceRoleReplica GetDbSystemInsightDbInstanceRoleEnum = "REPLICA"
)

var mappingGetDbSystemInsightDbInstanceRoleEnum = map[string]GetDbSystemInsightDbInstanceRoleEnum{
	"PRIMARY": GetDbSystemInsightDbInstanceRolePrimary,
	"REPLICA": GetDbSystemInsightDbInstanceRoleReplica,
}

var mappingGetDbSystemInsightDbInstanceRoleEnumLowerCase = map[string]GetDbSystemInsightDbInstanceRoleEnum{
	"primary": GetDbSystemInsightDbInstanceRolePrimary,
	"replica": GetDbSystemInsightDbInstanceRoleReplica,
}

// GetGetDbSystemInsightDbInstanceRoleEnumValues Enumerates the set of values for GetDbSystemInsightDbInstanceRoleEnum
func GetGetDbSystemInsightDbInstanceRoleEnumValues() []GetDbSystemInsightDbInstanceRoleEnum {
	values := make([]GetDbSystemInsightDbInstanceRoleEnum, 0)
	for _, v := range mappingGetDbSystemInsightDbInstanceRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemInsightDbInstanceRoleEnumStringValues Enumerates the set of values in String for GetDbSystemInsightDbInstanceRoleEnum
func GetGetDbSystemInsightDbInstanceRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"REPLICA",
	}
}

// GetMappingGetDbSystemInsightDbInstanceRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemInsightDbInstanceRoleEnum(val string) (GetDbSystemInsightDbInstanceRoleEnum, bool) {
	enum, ok := mappingGetDbSystemInsightDbInstanceRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDbSystemInsightSortInsightDataByEnum Enum with underlying type: string
type GetDbSystemInsightSortInsightDataByEnum string

// Set of constants representing the allowable values for GetDbSystemInsightSortInsightDataByEnum
const (
	GetDbSystemInsightSortInsightDataByAas              GetDbSystemInsightSortInsightDataByEnum = "AAS"
	GetDbSystemInsightSortInsightDataByAvgExecutionTime GetDbSystemInsightSortInsightDataByEnum = "AVG_EXECUTION_TIME"
	GetDbSystemInsightSortInsightDataByQueryCount       GetDbSystemInsightSortInsightDataByEnum = "QUERY_COUNT"
)

var mappingGetDbSystemInsightSortInsightDataByEnum = map[string]GetDbSystemInsightSortInsightDataByEnum{
	"AAS":                GetDbSystemInsightSortInsightDataByAas,
	"AVG_EXECUTION_TIME": GetDbSystemInsightSortInsightDataByAvgExecutionTime,
	"QUERY_COUNT":        GetDbSystemInsightSortInsightDataByQueryCount,
}

var mappingGetDbSystemInsightSortInsightDataByEnumLowerCase = map[string]GetDbSystemInsightSortInsightDataByEnum{
	"aas":                GetDbSystemInsightSortInsightDataByAas,
	"avg_execution_time": GetDbSystemInsightSortInsightDataByAvgExecutionTime,
	"query_count":        GetDbSystemInsightSortInsightDataByQueryCount,
}

// GetGetDbSystemInsightSortInsightDataByEnumValues Enumerates the set of values for GetDbSystemInsightSortInsightDataByEnum
func GetGetDbSystemInsightSortInsightDataByEnumValues() []GetDbSystemInsightSortInsightDataByEnum {
	values := make([]GetDbSystemInsightSortInsightDataByEnum, 0)
	for _, v := range mappingGetDbSystemInsightSortInsightDataByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemInsightSortInsightDataByEnumStringValues Enumerates the set of values in String for GetDbSystemInsightSortInsightDataByEnum
func GetGetDbSystemInsightSortInsightDataByEnumStringValues() []string {
	return []string{
		"AAS",
		"AVG_EXECUTION_TIME",
		"QUERY_COUNT",
	}
}

// GetMappingGetDbSystemInsightSortInsightDataByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemInsightSortInsightDataByEnum(val string) (GetDbSystemInsightSortInsightDataByEnum, bool) {
	enum, ok := mappingGetDbSystemInsightSortInsightDataByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDbSystemInsightSortOrderEnum Enum with underlying type: string
type GetDbSystemInsightSortOrderEnum string

// Set of constants representing the allowable values for GetDbSystemInsightSortOrderEnum
const (
	GetDbSystemInsightSortOrderAsc  GetDbSystemInsightSortOrderEnum = "ASC"
	GetDbSystemInsightSortOrderDesc GetDbSystemInsightSortOrderEnum = "DESC"
)

var mappingGetDbSystemInsightSortOrderEnum = map[string]GetDbSystemInsightSortOrderEnum{
	"ASC":  GetDbSystemInsightSortOrderAsc,
	"DESC": GetDbSystemInsightSortOrderDesc,
}

var mappingGetDbSystemInsightSortOrderEnumLowerCase = map[string]GetDbSystemInsightSortOrderEnum{
	"asc":  GetDbSystemInsightSortOrderAsc,
	"desc": GetDbSystemInsightSortOrderDesc,
}

// GetGetDbSystemInsightSortOrderEnumValues Enumerates the set of values for GetDbSystemInsightSortOrderEnum
func GetGetDbSystemInsightSortOrderEnumValues() []GetDbSystemInsightSortOrderEnum {
	values := make([]GetDbSystemInsightSortOrderEnum, 0)
	for _, v := range mappingGetDbSystemInsightSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemInsightSortOrderEnumStringValues Enumerates the set of values in String for GetDbSystemInsightSortOrderEnum
func GetGetDbSystemInsightSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetDbSystemInsightSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemInsightSortOrderEnum(val string) (GetDbSystemInsightSortOrderEnum, bool) {
	enum, ok := mappingGetDbSystemInsightSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
