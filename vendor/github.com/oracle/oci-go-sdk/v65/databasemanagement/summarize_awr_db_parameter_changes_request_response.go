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

// SummarizeAwrDbParameterChangesRequest wrapper for the SummarizeAwrDbParameterChanges operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbParameterChanges.go.html to see an example of how to use SummarizeAwrDbParameterChangesRequest.
type SummarizeAwrDbParameterChangesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The required single value query parameter to filter the entity name.
	Name *string `mandatory:"true" contributesTo:"query" name:"name"`

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

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in large paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR database parameter change history data.
	SortBy SummarizeAwrDbParameterChangesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbParameterChangesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request SummarizeAwrDbParameterChangesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbParameterChangesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbParameterChangesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbParameterChangesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDbParameterChangesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDbParameterChangesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDbParameterChangesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbParameterChangesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDbParameterChangesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDbParameterChangesResponse wrapper for the SummarizeAwrDbParameterChanges operation
type SummarizeAwrDbParameterChangesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbParameterChangeCollection instances
	AwrDbParameterChangeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbParameterChangesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbParameterChangesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbParameterChangesSortByEnum Enum with underlying type: string
type SummarizeAwrDbParameterChangesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParameterChangesSortByEnum
const (
	SummarizeAwrDbParameterChangesSortByIsChanged SummarizeAwrDbParameterChangesSortByEnum = "IS_CHANGED"
	SummarizeAwrDbParameterChangesSortByName      SummarizeAwrDbParameterChangesSortByEnum = "NAME"
)

var mappingSummarizeAwrDbParameterChangesSortByEnum = map[string]SummarizeAwrDbParameterChangesSortByEnum{
	"IS_CHANGED": SummarizeAwrDbParameterChangesSortByIsChanged,
	"NAME":       SummarizeAwrDbParameterChangesSortByName,
}

var mappingSummarizeAwrDbParameterChangesSortByEnumLowerCase = map[string]SummarizeAwrDbParameterChangesSortByEnum{
	"is_changed": SummarizeAwrDbParameterChangesSortByIsChanged,
	"name":       SummarizeAwrDbParameterChangesSortByName,
}

// GetSummarizeAwrDbParameterChangesSortByEnumValues Enumerates the set of values for SummarizeAwrDbParameterChangesSortByEnum
func GetSummarizeAwrDbParameterChangesSortByEnumValues() []SummarizeAwrDbParameterChangesSortByEnum {
	values := make([]SummarizeAwrDbParameterChangesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbParameterChangesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParameterChangesSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParameterChangesSortByEnum
func GetSummarizeAwrDbParameterChangesSortByEnumStringValues() []string {
	return []string{
		"IS_CHANGED",
		"NAME",
	}
}

// GetMappingSummarizeAwrDbParameterChangesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParameterChangesSortByEnum(val string) (SummarizeAwrDbParameterChangesSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParameterChangesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbParameterChangesSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbParameterChangesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbParameterChangesSortOrderEnum
const (
	SummarizeAwrDbParameterChangesSortOrderAsc  SummarizeAwrDbParameterChangesSortOrderEnum = "ASC"
	SummarizeAwrDbParameterChangesSortOrderDesc SummarizeAwrDbParameterChangesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbParameterChangesSortOrderEnum = map[string]SummarizeAwrDbParameterChangesSortOrderEnum{
	"ASC":  SummarizeAwrDbParameterChangesSortOrderAsc,
	"DESC": SummarizeAwrDbParameterChangesSortOrderDesc,
}

var mappingSummarizeAwrDbParameterChangesSortOrderEnumLowerCase = map[string]SummarizeAwrDbParameterChangesSortOrderEnum{
	"asc":  SummarizeAwrDbParameterChangesSortOrderAsc,
	"desc": SummarizeAwrDbParameterChangesSortOrderDesc,
}

// GetSummarizeAwrDbParameterChangesSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbParameterChangesSortOrderEnum
func GetSummarizeAwrDbParameterChangesSortOrderEnumValues() []SummarizeAwrDbParameterChangesSortOrderEnum {
	values := make([]SummarizeAwrDbParameterChangesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbParameterChangesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbParameterChangesSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDbParameterChangesSortOrderEnum
func GetSummarizeAwrDbParameterChangesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDbParameterChangesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbParameterChangesSortOrderEnum(val string) (SummarizeAwrDbParameterChangesSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDbParameterChangesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
