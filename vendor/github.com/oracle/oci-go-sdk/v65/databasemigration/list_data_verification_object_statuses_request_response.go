// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataVerificationObjectStatusesRequest wrapper for the ListDataVerificationObjectStatuses operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListDataVerificationObjectStatuses.go.html to see an example of how to use ListDataVerificationObjectStatusesRequest.
type ListDataVerificationObjectStatusesRequest struct {

	// The OCID of the migration
	MigrationId *string `mandatory:"true" contributesTo:"path" name:"migrationId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only results for a specific owner.
	Owner *string `mandatory:"false" contributesTo:"query" name:"owner"`

	// A filter to return only results for a specific object type.
	// The allowed values depend on the migration's `databaseCombination`:
	// - Oracle migrations: `OracleDatabaseObjectTypes`
	// - MySQL migrations: `MySqlDatabaseObjectTypes`
	ObjectType *string `mandatory:"false" contributesTo:"query" name:"objectType"`

	// Free-text filter applied by the service to relevant fields for the report.
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// Match filter for object status results.
	// When `true`, returns only rows where source and target match.
	// When `false`, returns only rows with mismatches.
	// When omitted, returns all rows, ordered by mismatch first.
	IsMatch *bool `mandatory:"false" contributesTo:"query" name:"isMatch"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// By default, non-matching object statuses are returned first, followed by matching statuses.
	// Rows with the same match result are then sorted by owner, object type, and object name.
	SortBy ListDataVerificationObjectStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataVerificationObjectStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataVerificationObjectStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataVerificationObjectStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataVerificationObjectStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataVerificationObjectStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataVerificationObjectStatusesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataVerificationObjectStatusesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataVerificationObjectStatusesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataVerificationObjectStatusesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataVerificationObjectStatusesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataVerificationObjectStatusesResponse wrapper for the ListDataVerificationObjectStatuses operation
type ListDataVerificationObjectStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataVerificationObjectStatusCollection instances
	DataVerificationObjectStatusCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataVerificationObjectStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataVerificationObjectStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataVerificationObjectStatusesSortByEnum Enum with underlying type: string
type ListDataVerificationObjectStatusesSortByEnum string

// Set of constants representing the allowable values for ListDataVerificationObjectStatusesSortByEnum
const (
	ListDataVerificationObjectStatusesSortByIsmatch    ListDataVerificationObjectStatusesSortByEnum = "isMatch"
	ListDataVerificationObjectStatusesSortByOwner      ListDataVerificationObjectStatusesSortByEnum = "owner"
	ListDataVerificationObjectStatusesSortByObjecttype ListDataVerificationObjectStatusesSortByEnum = "objectType"
	ListDataVerificationObjectStatusesSortByObjectname ListDataVerificationObjectStatusesSortByEnum = "objectName"
)

var mappingListDataVerificationObjectStatusesSortByEnum = map[string]ListDataVerificationObjectStatusesSortByEnum{
	"isMatch":    ListDataVerificationObjectStatusesSortByIsmatch,
	"owner":      ListDataVerificationObjectStatusesSortByOwner,
	"objectType": ListDataVerificationObjectStatusesSortByObjecttype,
	"objectName": ListDataVerificationObjectStatusesSortByObjectname,
}

var mappingListDataVerificationObjectStatusesSortByEnumLowerCase = map[string]ListDataVerificationObjectStatusesSortByEnum{
	"ismatch":    ListDataVerificationObjectStatusesSortByIsmatch,
	"owner":      ListDataVerificationObjectStatusesSortByOwner,
	"objecttype": ListDataVerificationObjectStatusesSortByObjecttype,
	"objectname": ListDataVerificationObjectStatusesSortByObjectname,
}

// GetListDataVerificationObjectStatusesSortByEnumValues Enumerates the set of values for ListDataVerificationObjectStatusesSortByEnum
func GetListDataVerificationObjectStatusesSortByEnumValues() []ListDataVerificationObjectStatusesSortByEnum {
	values := make([]ListDataVerificationObjectStatusesSortByEnum, 0)
	for _, v := range mappingListDataVerificationObjectStatusesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataVerificationObjectStatusesSortByEnumStringValues Enumerates the set of values in String for ListDataVerificationObjectStatusesSortByEnum
func GetListDataVerificationObjectStatusesSortByEnumStringValues() []string {
	return []string{
		"isMatch",
		"owner",
		"objectType",
		"objectName",
	}
}

// GetMappingListDataVerificationObjectStatusesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataVerificationObjectStatusesSortByEnum(val string) (ListDataVerificationObjectStatusesSortByEnum, bool) {
	enum, ok := mappingListDataVerificationObjectStatusesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataVerificationObjectStatusesSortOrderEnum Enum with underlying type: string
type ListDataVerificationObjectStatusesSortOrderEnum string

// Set of constants representing the allowable values for ListDataVerificationObjectStatusesSortOrderEnum
const (
	ListDataVerificationObjectStatusesSortOrderAsc  ListDataVerificationObjectStatusesSortOrderEnum = "ASC"
	ListDataVerificationObjectStatusesSortOrderDesc ListDataVerificationObjectStatusesSortOrderEnum = "DESC"
)

var mappingListDataVerificationObjectStatusesSortOrderEnum = map[string]ListDataVerificationObjectStatusesSortOrderEnum{
	"ASC":  ListDataVerificationObjectStatusesSortOrderAsc,
	"DESC": ListDataVerificationObjectStatusesSortOrderDesc,
}

var mappingListDataVerificationObjectStatusesSortOrderEnumLowerCase = map[string]ListDataVerificationObjectStatusesSortOrderEnum{
	"asc":  ListDataVerificationObjectStatusesSortOrderAsc,
	"desc": ListDataVerificationObjectStatusesSortOrderDesc,
}

// GetListDataVerificationObjectStatusesSortOrderEnumValues Enumerates the set of values for ListDataVerificationObjectStatusesSortOrderEnum
func GetListDataVerificationObjectStatusesSortOrderEnumValues() []ListDataVerificationObjectStatusesSortOrderEnum {
	values := make([]ListDataVerificationObjectStatusesSortOrderEnum, 0)
	for _, v := range mappingListDataVerificationObjectStatusesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataVerificationObjectStatusesSortOrderEnumStringValues Enumerates the set of values in String for ListDataVerificationObjectStatusesSortOrderEnum
func GetListDataVerificationObjectStatusesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataVerificationObjectStatusesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataVerificationObjectStatusesSortOrderEnum(val string) (ListDataVerificationObjectStatusesSortOrderEnum, bool) {
	enum, ok := mappingListDataVerificationObjectStatusesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
