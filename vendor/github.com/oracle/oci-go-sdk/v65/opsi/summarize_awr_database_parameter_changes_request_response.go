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

// SummarizeAwrDatabaseParameterChangesRequest wrapper for the SummarizeAwrDatabaseParameterChanges operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseParameterChanges.go.html to see an example of how to use SummarizeAwrDatabaseParameterChangesRequest.
type SummarizeAwrDatabaseParameterChangesRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The required single value query parameter to filter the entity name.
	Name *string `mandatory:"true" contributesTo:"query" name:"name"`

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
	SortBy SummarizeAwrDatabaseParameterChangesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseParameterChangesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseParameterChangesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseParameterChangesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseParameterChangesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseParameterChangesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseParameterChangesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseParameterChangesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseParameterChangesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseParameterChangesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseParameterChangesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseParameterChangesResponse wrapper for the SummarizeAwrDatabaseParameterChanges operation
type SummarizeAwrDatabaseParameterChangesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseParameterChangeCollection instances
	AwrDatabaseParameterChangeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseParameterChangesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseParameterChangesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseParameterChangesSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseParameterChangesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParameterChangesSortByEnum
const (
	SummarizeAwrDatabaseParameterChangesSortByIsChanged SummarizeAwrDatabaseParameterChangesSortByEnum = "IS_CHANGED"
	SummarizeAwrDatabaseParameterChangesSortByName      SummarizeAwrDatabaseParameterChangesSortByEnum = "NAME"
)

var mappingSummarizeAwrDatabaseParameterChangesSortByEnum = map[string]SummarizeAwrDatabaseParameterChangesSortByEnum{
	"IS_CHANGED": SummarizeAwrDatabaseParameterChangesSortByIsChanged,
	"NAME":       SummarizeAwrDatabaseParameterChangesSortByName,
}

var mappingSummarizeAwrDatabaseParameterChangesSortByEnumLowerCase = map[string]SummarizeAwrDatabaseParameterChangesSortByEnum{
	"is_changed": SummarizeAwrDatabaseParameterChangesSortByIsChanged,
	"name":       SummarizeAwrDatabaseParameterChangesSortByName,
}

// GetSummarizeAwrDatabaseParameterChangesSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseParameterChangesSortByEnum
func GetSummarizeAwrDatabaseParameterChangesSortByEnumValues() []SummarizeAwrDatabaseParameterChangesSortByEnum {
	values := make([]SummarizeAwrDatabaseParameterChangesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParameterChangesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParameterChangesSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParameterChangesSortByEnum
func GetSummarizeAwrDatabaseParameterChangesSortByEnumStringValues() []string {
	return []string{
		"IS_CHANGED",
		"NAME",
	}
}

// GetMappingSummarizeAwrDatabaseParameterChangesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParameterChangesSortByEnum(val string) (SummarizeAwrDatabaseParameterChangesSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParameterChangesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseParameterChangesSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseParameterChangesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseParameterChangesSortOrderEnum
const (
	SummarizeAwrDatabaseParameterChangesSortOrderAsc  SummarizeAwrDatabaseParameterChangesSortOrderEnum = "ASC"
	SummarizeAwrDatabaseParameterChangesSortOrderDesc SummarizeAwrDatabaseParameterChangesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseParameterChangesSortOrderEnum = map[string]SummarizeAwrDatabaseParameterChangesSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseParameterChangesSortOrderAsc,
	"DESC": SummarizeAwrDatabaseParameterChangesSortOrderDesc,
}

var mappingSummarizeAwrDatabaseParameterChangesSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseParameterChangesSortOrderEnum{
	"asc":  SummarizeAwrDatabaseParameterChangesSortOrderAsc,
	"desc": SummarizeAwrDatabaseParameterChangesSortOrderDesc,
}

// GetSummarizeAwrDatabaseParameterChangesSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseParameterChangesSortOrderEnum
func GetSummarizeAwrDatabaseParameterChangesSortOrderEnumValues() []SummarizeAwrDatabaseParameterChangesSortOrderEnum {
	values := make([]SummarizeAwrDatabaseParameterChangesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseParameterChangesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseParameterChangesSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseParameterChangesSortOrderEnum
func GetSummarizeAwrDatabaseParameterChangesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseParameterChangesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseParameterChangesSortOrderEnum(val string) (SummarizeAwrDatabaseParameterChangesSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseParameterChangesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
