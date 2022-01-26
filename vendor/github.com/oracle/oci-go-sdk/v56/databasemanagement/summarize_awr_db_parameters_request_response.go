// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeAwrDbParametersRequest wrapper for the SummarizeAwrDbParameters operation
//
// See also
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

var mappingSummarizeAwrDbParametersValueChanged = map[string]SummarizeAwrDbParametersValueChangedEnum{
	"Y": SummarizeAwrDbParametersValueChangedY,
	"N": SummarizeAwrDbParametersValueChangedN,
}

// GetSummarizeAwrDbParametersValueChangedEnumValues Enumerates the set of values for SummarizeAwrDbParametersValueChangedEnum
func GetSummarizeAwrDbParametersValueChangedEnumValues() []SummarizeAwrDbParametersValueChangedEnum {
	values := make([]SummarizeAwrDbParametersValueChangedEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersValueChanged {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbParametersValueDefaultEnum Enum with underlying type: string
type SummarizeAwrDbParametersValueDefaultEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersValueDefaultEnum
const (
	SummarizeAwrDbParametersValueDefaultTrue  SummarizeAwrDbParametersValueDefaultEnum = "TRUE"
	SummarizeAwrDbParametersValueDefaultFalse SummarizeAwrDbParametersValueDefaultEnum = "FALSE"
)

var mappingSummarizeAwrDbParametersValueDefault = map[string]SummarizeAwrDbParametersValueDefaultEnum{
	"TRUE":  SummarizeAwrDbParametersValueDefaultTrue,
	"FALSE": SummarizeAwrDbParametersValueDefaultFalse,
}

// GetSummarizeAwrDbParametersValueDefaultEnumValues Enumerates the set of values for SummarizeAwrDbParametersValueDefaultEnum
func GetSummarizeAwrDbParametersValueDefaultEnumValues() []SummarizeAwrDbParametersValueDefaultEnum {
	values := make([]SummarizeAwrDbParametersValueDefaultEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersValueDefault {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbParametersValueModifiedEnum Enum with underlying type: string
type SummarizeAwrDbParametersValueModifiedEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersValueModifiedEnum
const (
	SummarizeAwrDbParametersValueModifiedModified  SummarizeAwrDbParametersValueModifiedEnum = "MODIFIED"
	SummarizeAwrDbParametersValueModifiedSystemMod SummarizeAwrDbParametersValueModifiedEnum = "SYSTEM_MOD"
	SummarizeAwrDbParametersValueModifiedFalse     SummarizeAwrDbParametersValueModifiedEnum = "FALSE"
)

var mappingSummarizeAwrDbParametersValueModified = map[string]SummarizeAwrDbParametersValueModifiedEnum{
	"MODIFIED":   SummarizeAwrDbParametersValueModifiedModified,
	"SYSTEM_MOD": SummarizeAwrDbParametersValueModifiedSystemMod,
	"FALSE":      SummarizeAwrDbParametersValueModifiedFalse,
}

// GetSummarizeAwrDbParametersValueModifiedEnumValues Enumerates the set of values for SummarizeAwrDbParametersValueModifiedEnum
func GetSummarizeAwrDbParametersValueModifiedEnumValues() []SummarizeAwrDbParametersValueModifiedEnum {
	values := make([]SummarizeAwrDbParametersValueModifiedEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersValueModified {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbParametersSortByEnum Enum with underlying type: string
type SummarizeAwrDbParametersSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersSortByEnum
const (
	SummarizeAwrDbParametersSortByIsChanged SummarizeAwrDbParametersSortByEnum = "IS_CHANGED"
	SummarizeAwrDbParametersSortByName      SummarizeAwrDbParametersSortByEnum = "NAME"
)

var mappingSummarizeAwrDbParametersSortBy = map[string]SummarizeAwrDbParametersSortByEnum{
	"IS_CHANGED": SummarizeAwrDbParametersSortByIsChanged,
	"NAME":       SummarizeAwrDbParametersSortByName,
}

// GetSummarizeAwrDbParametersSortByEnumValues Enumerates the set of values for SummarizeAwrDbParametersSortByEnum
func GetSummarizeAwrDbParametersSortByEnumValues() []SummarizeAwrDbParametersSortByEnum {
	values := make([]SummarizeAwrDbParametersSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbParametersSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbParametersSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParametersSortOrderEnum
const (
	SummarizeAwrDbParametersSortOrderAsc  SummarizeAwrDbParametersSortOrderEnum = "ASC"
	SummarizeAwrDbParametersSortOrderDesc SummarizeAwrDbParametersSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbParametersSortOrder = map[string]SummarizeAwrDbParametersSortOrderEnum{
	"ASC":  SummarizeAwrDbParametersSortOrderAsc,
	"DESC": SummarizeAwrDbParametersSortOrderDesc,
}

// GetSummarizeAwrDbParametersSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbParametersSortOrderEnum
func GetSummarizeAwrDbParametersSortOrderEnumValues() []SummarizeAwrDbParametersSortOrderEnum {
	values := make([]SummarizeAwrDbParametersSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbParametersSortOrder {
		values = append(values, v)
	}
	return values
}
