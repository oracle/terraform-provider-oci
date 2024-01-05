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

// SummarizeAwrDatabaseParametersRequest wrapper for the SummarizeAwrDatabaseParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseParameters.go.html to see an example of how to use SummarizeAwrDatabaseParametersRequest.
type SummarizeAwrDatabaseParametersRequest struct {

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

	// The optional contains query parameter to filter the entity name by any part of the name.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The optional query parameter to filter database parameters whose values were changed.
	ValueChanged SummarizeAwrDatabaseParametersValueChangedEnum `mandatory:"false" contributesTo:"query" name:"valueChanged" omitEmpty:"true"`

	// The optional query parameter to filter the database parameters that had the default value in the last snapshot.
	ValueDefault SummarizeAwrDatabaseParametersValueDefaultEnum `mandatory:"false" contributesTo:"query" name:"valueDefault" omitEmpty:"true"`

	// The optional query parameter to filter the database parameters that had a modified value in the last snapshot.
	ValueModified SummarizeAwrDatabaseParametersValueModifiedEnum `mandatory:"false" contributesTo:"query" name:"valueModified" omitEmpty:"true"`

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

	// The option to sort the AWR database parameter change history data.
	SortBy SummarizeAwrDatabaseParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseParametersValueChangedEnum(string(request.ValueChanged)); !ok && request.ValueChanged != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueChanged: %s. Supported values are: %s.", request.ValueChanged, strings.Join(GetSummarizeAwrDatabaseParametersValueChangedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseParametersValueDefaultEnum(string(request.ValueDefault)); !ok && request.ValueDefault != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueDefault: %s. Supported values are: %s.", request.ValueDefault, strings.Join(GetSummarizeAwrDatabaseParametersValueDefaultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseParametersValueModifiedEnum(string(request.ValueModified)); !ok && request.ValueModified != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueModified: %s. Supported values are: %s.", request.ValueModified, strings.Join(GetSummarizeAwrDatabaseParametersValueModifiedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseParametersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseParametersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseParametersResponse wrapper for the SummarizeAwrDatabaseParameters operation
type SummarizeAwrDatabaseParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseParameterCollection instances
	AwrDatabaseParameterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseParametersValueChangedEnum Enum with underlying type: string
type SummarizeAwrDatabaseParametersValueChangedEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParametersValueChangedEnum
const (
	SummarizeAwrDatabaseParametersValueChangedY SummarizeAwrDatabaseParametersValueChangedEnum = "Y"
	SummarizeAwrDatabaseParametersValueChangedN SummarizeAwrDatabaseParametersValueChangedEnum = "N"
)

var mappingSummarizeAwrDatabaseParametersValueChangedEnum = map[string]SummarizeAwrDatabaseParametersValueChangedEnum{
	"Y": SummarizeAwrDatabaseParametersValueChangedY,
	"N": SummarizeAwrDatabaseParametersValueChangedN,
}

var mappingSummarizeAwrDatabaseParametersValueChangedEnumLowerCase = map[string]SummarizeAwrDatabaseParametersValueChangedEnum{
	"y": SummarizeAwrDatabaseParametersValueChangedY,
	"n": SummarizeAwrDatabaseParametersValueChangedN,
}

// GetSummarizeAwrDatabaseParametersValueChangedEnumValues Enumerates the set of values for SummarizeAwrDatabaseParametersValueChangedEnum
func GetSummarizeAwrDatabaseParametersValueChangedEnumValues() []SummarizeAwrDatabaseParametersValueChangedEnum {
	values := make([]SummarizeAwrDatabaseParametersValueChangedEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParametersValueChangedEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParametersValueChangedEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParametersValueChangedEnum
func GetSummarizeAwrDatabaseParametersValueChangedEnumStringValues() []string {
	return []string{
		"Y",
		"N",
	}
}

// GetMappingSummarizeAwrDatabaseParametersValueChangedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParametersValueChangedEnum(val string) (SummarizeAwrDatabaseParametersValueChangedEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParametersValueChangedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseParametersValueDefaultEnum Enum with underlying type: string
type SummarizeAwrDatabaseParametersValueDefaultEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParametersValueDefaultEnum
const (
	SummarizeAwrDatabaseParametersValueDefaultTrue  SummarizeAwrDatabaseParametersValueDefaultEnum = "TRUE"
	SummarizeAwrDatabaseParametersValueDefaultFalse SummarizeAwrDatabaseParametersValueDefaultEnum = "FALSE"
)

var mappingSummarizeAwrDatabaseParametersValueDefaultEnum = map[string]SummarizeAwrDatabaseParametersValueDefaultEnum{
	"TRUE":  SummarizeAwrDatabaseParametersValueDefaultTrue,
	"FALSE": SummarizeAwrDatabaseParametersValueDefaultFalse,
}

var mappingSummarizeAwrDatabaseParametersValueDefaultEnumLowerCase = map[string]SummarizeAwrDatabaseParametersValueDefaultEnum{
	"true":  SummarizeAwrDatabaseParametersValueDefaultTrue,
	"false": SummarizeAwrDatabaseParametersValueDefaultFalse,
}

// GetSummarizeAwrDatabaseParametersValueDefaultEnumValues Enumerates the set of values for SummarizeAwrDatabaseParametersValueDefaultEnum
func GetSummarizeAwrDatabaseParametersValueDefaultEnumValues() []SummarizeAwrDatabaseParametersValueDefaultEnum {
	values := make([]SummarizeAwrDatabaseParametersValueDefaultEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParametersValueDefaultEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParametersValueDefaultEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParametersValueDefaultEnum
func GetSummarizeAwrDatabaseParametersValueDefaultEnumStringValues() []string {
	return []string{
		"TRUE",
		"FALSE",
	}
}

// GetMappingSummarizeAwrDatabaseParametersValueDefaultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParametersValueDefaultEnum(val string) (SummarizeAwrDatabaseParametersValueDefaultEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParametersValueDefaultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseParametersValueModifiedEnum Enum with underlying type: string
type SummarizeAwrDatabaseParametersValueModifiedEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParametersValueModifiedEnum
const (
	SummarizeAwrDatabaseParametersValueModifiedModified  SummarizeAwrDatabaseParametersValueModifiedEnum = "MODIFIED"
	SummarizeAwrDatabaseParametersValueModifiedSystemMod SummarizeAwrDatabaseParametersValueModifiedEnum = "SYSTEM_MOD"
	SummarizeAwrDatabaseParametersValueModifiedFalse     SummarizeAwrDatabaseParametersValueModifiedEnum = "FALSE"
)

var mappingSummarizeAwrDatabaseParametersValueModifiedEnum = map[string]SummarizeAwrDatabaseParametersValueModifiedEnum{
	"MODIFIED":   SummarizeAwrDatabaseParametersValueModifiedModified,
	"SYSTEM_MOD": SummarizeAwrDatabaseParametersValueModifiedSystemMod,
	"FALSE":      SummarizeAwrDatabaseParametersValueModifiedFalse,
}

var mappingSummarizeAwrDatabaseParametersValueModifiedEnumLowerCase = map[string]SummarizeAwrDatabaseParametersValueModifiedEnum{
	"modified":   SummarizeAwrDatabaseParametersValueModifiedModified,
	"system_mod": SummarizeAwrDatabaseParametersValueModifiedSystemMod,
	"false":      SummarizeAwrDatabaseParametersValueModifiedFalse,
}

// GetSummarizeAwrDatabaseParametersValueModifiedEnumValues Enumerates the set of values for SummarizeAwrDatabaseParametersValueModifiedEnum
func GetSummarizeAwrDatabaseParametersValueModifiedEnumValues() []SummarizeAwrDatabaseParametersValueModifiedEnum {
	values := make([]SummarizeAwrDatabaseParametersValueModifiedEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParametersValueModifiedEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParametersValueModifiedEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParametersValueModifiedEnum
func GetSummarizeAwrDatabaseParametersValueModifiedEnumStringValues() []string {
	return []string{
		"MODIFIED",
		"SYSTEM_MOD",
		"FALSE",
	}
}

// GetMappingSummarizeAwrDatabaseParametersValueModifiedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParametersValueModifiedEnum(val string) (SummarizeAwrDatabaseParametersValueModifiedEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParametersValueModifiedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseParametersSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseParametersSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParametersSortByEnum
const (
	SummarizeAwrDatabaseParametersSortByIsChanged SummarizeAwrDatabaseParametersSortByEnum = "IS_CHANGED"
	SummarizeAwrDatabaseParametersSortByName      SummarizeAwrDatabaseParametersSortByEnum = "NAME"
)

var mappingSummarizeAwrDatabaseParametersSortByEnum = map[string]SummarizeAwrDatabaseParametersSortByEnum{
	"IS_CHANGED": SummarizeAwrDatabaseParametersSortByIsChanged,
	"NAME":       SummarizeAwrDatabaseParametersSortByName,
}

var mappingSummarizeAwrDatabaseParametersSortByEnumLowerCase = map[string]SummarizeAwrDatabaseParametersSortByEnum{
	"is_changed": SummarizeAwrDatabaseParametersSortByIsChanged,
	"name":       SummarizeAwrDatabaseParametersSortByName,
}

// GetSummarizeAwrDatabaseParametersSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseParametersSortByEnum
func GetSummarizeAwrDatabaseParametersSortByEnumValues() []SummarizeAwrDatabaseParametersSortByEnum {
	values := make([]SummarizeAwrDatabaseParametersSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParametersSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParametersSortByEnum
func GetSummarizeAwrDatabaseParametersSortByEnumStringValues() []string {
	return []string{
		"IS_CHANGED",
		"NAME",
	}
}

// GetMappingSummarizeAwrDatabaseParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParametersSortByEnum(val string) (SummarizeAwrDatabaseParametersSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseParametersSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseParametersSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParametersSortOrderEnum
const (
	SummarizeAwrDatabaseParametersSortOrderAsc  SummarizeAwrDatabaseParametersSortOrderEnum = "ASC"
	SummarizeAwrDatabaseParametersSortOrderDesc SummarizeAwrDatabaseParametersSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseParametersSortOrderEnum = map[string]SummarizeAwrDatabaseParametersSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseParametersSortOrderAsc,
	"DESC": SummarizeAwrDatabaseParametersSortOrderDesc,
}

var mappingSummarizeAwrDatabaseParametersSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseParametersSortOrderEnum{
	"asc":  SummarizeAwrDatabaseParametersSortOrderAsc,
	"desc": SummarizeAwrDatabaseParametersSortOrderDesc,
}

// GetSummarizeAwrDatabaseParametersSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseParametersSortOrderEnum
func GetSummarizeAwrDatabaseParametersSortOrderEnumValues() []SummarizeAwrDatabaseParametersSortOrderEnum {
	values := make([]SummarizeAwrDatabaseParametersSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParametersSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParametersSortOrderEnum
func GetSummarizeAwrDatabaseParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParametersSortOrderEnum(val string) (SummarizeAwrDatabaseParametersSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
