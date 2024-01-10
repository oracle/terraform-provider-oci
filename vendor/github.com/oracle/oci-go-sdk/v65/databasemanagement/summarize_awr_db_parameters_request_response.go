// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeAwrDbParametersRequest wrapper for the SummarizeAwrDbParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbParameters.go.html to see an example of how to use SummarizeAwrDbParametersRequest.
type SummarizeAwrDbParametersRequest struct {

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

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The optional multiple value query parameter to filter the entity name.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The optional contains query parameter to filter the entity name by any part of the name.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The optional query parameter to filter database parameters whose values were changed.
	ValueChanged SummarizeAwrDbParametersValueChangedEnum `mandatory:"false" contributesTo:"query" name:"valueChanged" omitEmpty:"true"`

	// The optional query parameter to filter the database parameters that had the default value in the last snapshot.
	ValueDefault SummarizeAwrDbParametersValueDefaultEnum `mandatory:"false" contributesTo:"query" name:"valueDefault" omitEmpty:"true"`

	// The optional query parameter to filter the database parameters that had a modified value in the last snapshot.
	ValueModified SummarizeAwrDbParametersValueModifiedEnum `mandatory:"false" contributesTo:"query" name:"valueModified" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in large paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR database parameter change history data.
	SortBy SummarizeAwrDbParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request SummarizeAwrDbParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDbParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDbParametersValueChangedEnum(string(request.ValueChanged)); !ok && request.ValueChanged != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueChanged: %s. Supported values are: %s.", request.ValueChanged, strings.Join(GetSummarizeAwrDbParametersValueChangedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbParametersValueDefaultEnum(string(request.ValueDefault)); !ok && request.ValueDefault != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueDefault: %s. Supported values are: %s.", request.ValueDefault, strings.Join(GetSummarizeAwrDbParametersValueDefaultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbParametersValueModifiedEnum(string(request.ValueModified)); !ok && request.ValueModified != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueModified: %s. Supported values are: %s.", request.ValueModified, strings.Join(GetSummarizeAwrDbParametersValueModifiedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDbParametersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDbParametersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDbParametersResponse wrapper for the SummarizeAwrDbParameters operation
type SummarizeAwrDbParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbParameterCollection instances
	AwrDbParameterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbParametersValueChangedEnum Enum with underlying type: string
type SummarizeAwrDbParametersValueChangedEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersValueChangedEnum
const (
	SummarizeAwrDbParametersValueChangedY SummarizeAwrDbParametersValueChangedEnum = "Y"
	SummarizeAwrDbParametersValueChangedN SummarizeAwrDbParametersValueChangedEnum = "N"
)

var mappingSummarizeAwrDbParametersValueChangedEnum = map[string]SummarizeAwrDbParametersValueChangedEnum{
	"Y": SummarizeAwrDbParametersValueChangedY,
	"N": SummarizeAwrDbParametersValueChangedN,
}

var mappingSummarizeAwrDbParametersValueChangedEnumLowerCase = map[string]SummarizeAwrDbParametersValueChangedEnum{
	"y": SummarizeAwrDbParametersValueChangedY,
	"n": SummarizeAwrDbParametersValueChangedN,
}

// GetSummarizeAwrDbParametersValueChangedEnumValues Enumerates the set of values for SummarizeAwrDbParametersValueChangedEnum
func GetSummarizeAwrDbParametersValueChangedEnumValues() []SummarizeAwrDbParametersValueChangedEnum {
	values := make([]SummarizeAwrDbParametersValueChangedEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersValueChangedEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParametersValueChangedEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParametersValueChangedEnum
func GetSummarizeAwrDbParametersValueChangedEnumStringValues() []string {
	return []string{
		"Y",
		"N",
	}
}

// GetMappingSummarizeAwrDbParametersValueChangedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParametersValueChangedEnum(val string) (SummarizeAwrDbParametersValueChangedEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParametersValueChangedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbParametersValueDefaultEnum Enum with underlying type: string
type SummarizeAwrDbParametersValueDefaultEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersValueDefaultEnum
const (
	SummarizeAwrDbParametersValueDefaultTrue  SummarizeAwrDbParametersValueDefaultEnum = "TRUE"
	SummarizeAwrDbParametersValueDefaultFalse SummarizeAwrDbParametersValueDefaultEnum = "FALSE"
)

var mappingSummarizeAwrDbParametersValueDefaultEnum = map[string]SummarizeAwrDbParametersValueDefaultEnum{
	"TRUE":  SummarizeAwrDbParametersValueDefaultTrue,
	"FALSE": SummarizeAwrDbParametersValueDefaultFalse,
}

var mappingSummarizeAwrDbParametersValueDefaultEnumLowerCase = map[string]SummarizeAwrDbParametersValueDefaultEnum{
	"true":  SummarizeAwrDbParametersValueDefaultTrue,
	"false": SummarizeAwrDbParametersValueDefaultFalse,
}

// GetSummarizeAwrDbParametersValueDefaultEnumValues Enumerates the set of values for SummarizeAwrDbParametersValueDefaultEnum
func GetSummarizeAwrDbParametersValueDefaultEnumValues() []SummarizeAwrDbParametersValueDefaultEnum {
	values := make([]SummarizeAwrDbParametersValueDefaultEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersValueDefaultEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParametersValueDefaultEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParametersValueDefaultEnum
func GetSummarizeAwrDbParametersValueDefaultEnumStringValues() []string {
	return []string{
		"TRUE",
		"FALSE",
	}
}

// GetMappingSummarizeAwrDbParametersValueDefaultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParametersValueDefaultEnum(val string) (SummarizeAwrDbParametersValueDefaultEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParametersValueDefaultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbParametersValueModifiedEnum Enum with underlying type: string
type SummarizeAwrDbParametersValueModifiedEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersValueModifiedEnum
const (
	SummarizeAwrDbParametersValueModifiedModified  SummarizeAwrDbParametersValueModifiedEnum = "MODIFIED"
	SummarizeAwrDbParametersValueModifiedSystemMod SummarizeAwrDbParametersValueModifiedEnum = "SYSTEM_MOD"
	SummarizeAwrDbParametersValueModifiedFalse     SummarizeAwrDbParametersValueModifiedEnum = "FALSE"
)

var mappingSummarizeAwrDbParametersValueModifiedEnum = map[string]SummarizeAwrDbParametersValueModifiedEnum{
	"MODIFIED":   SummarizeAwrDbParametersValueModifiedModified,
	"SYSTEM_MOD": SummarizeAwrDbParametersValueModifiedSystemMod,
	"FALSE":      SummarizeAwrDbParametersValueModifiedFalse,
}

var mappingSummarizeAwrDbParametersValueModifiedEnumLowerCase = map[string]SummarizeAwrDbParametersValueModifiedEnum{
	"modified":   SummarizeAwrDbParametersValueModifiedModified,
	"system_mod": SummarizeAwrDbParametersValueModifiedSystemMod,
	"false":      SummarizeAwrDbParametersValueModifiedFalse,
}

// GetSummarizeAwrDbParametersValueModifiedEnumValues Enumerates the set of values for SummarizeAwrDbParametersValueModifiedEnum
func GetSummarizeAwrDbParametersValueModifiedEnumValues() []SummarizeAwrDbParametersValueModifiedEnum {
	values := make([]SummarizeAwrDbParametersValueModifiedEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersValueModifiedEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParametersValueModifiedEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParametersValueModifiedEnum
func GetSummarizeAwrDbParametersValueModifiedEnumStringValues() []string {
	return []string{
		"MODIFIED",
		"SYSTEM_MOD",
		"FALSE",
	}
}

// GetMappingSummarizeAwrDbParametersValueModifiedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParametersValueModifiedEnum(val string) (SummarizeAwrDbParametersValueModifiedEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParametersValueModifiedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbParametersSortByEnum Enum with underlying type: string
type SummarizeAwrDbParametersSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersSortByEnum
const (
	SummarizeAwrDbParametersSortByIsChanged SummarizeAwrDbParametersSortByEnum = "IS_CHANGED"
	SummarizeAwrDbParametersSortByName      SummarizeAwrDbParametersSortByEnum = "NAME"
)

var mappingSummarizeAwrDbParametersSortByEnum = map[string]SummarizeAwrDbParametersSortByEnum{
	"IS_CHANGED": SummarizeAwrDbParametersSortByIsChanged,
	"NAME":       SummarizeAwrDbParametersSortByName,
}

var mappingSummarizeAwrDbParametersSortByEnumLowerCase = map[string]SummarizeAwrDbParametersSortByEnum{
	"is_changed": SummarizeAwrDbParametersSortByIsChanged,
	"name":       SummarizeAwrDbParametersSortByName,
}

// GetSummarizeAwrDbParametersSortByEnumValues Enumerates the set of values for SummarizeAwrDbParametersSortByEnum
func GetSummarizeAwrDbParametersSortByEnumValues() []SummarizeAwrDbParametersSortByEnum {
	values := make([]SummarizeAwrDbParametersSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParametersSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParametersSortByEnum
func GetSummarizeAwrDbParametersSortByEnumStringValues() []string {
	return []string{
		"IS_CHANGED",
		"NAME",
	}
}

// GetMappingSummarizeAwrDbParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParametersSortByEnum(val string) (SummarizeAwrDbParametersSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbParametersSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbParametersSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersSortOrderEnum
const (
	SummarizeAwrDbParametersSortOrderAsc  SummarizeAwrDbParametersSortOrderEnum = "ASC"
	SummarizeAwrDbParametersSortOrderDesc SummarizeAwrDbParametersSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbParametersSortOrderEnum = map[string]SummarizeAwrDbParametersSortOrderEnum{
	"ASC":  SummarizeAwrDbParametersSortOrderAsc,
	"DESC": SummarizeAwrDbParametersSortOrderDesc,
}

var mappingSummarizeAwrDbParametersSortOrderEnumLowerCase = map[string]SummarizeAwrDbParametersSortOrderEnum{
	"asc":  SummarizeAwrDbParametersSortOrderAsc,
	"desc": SummarizeAwrDbParametersSortOrderDesc,
}

// GetSummarizeAwrDbParametersSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbParametersSortOrderEnum
func GetSummarizeAwrDbParametersSortOrderEnumValues() []SummarizeAwrDbParametersSortOrderEnum {
	values := make([]SummarizeAwrDbParametersSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParametersSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParametersSortOrderEnum
func GetSummarizeAwrDbParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDbParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParametersSortOrderEnum(val string) (SummarizeAwrDbParametersSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
