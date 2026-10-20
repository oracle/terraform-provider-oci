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

// ListDataVerificationObjectTypeCountsRequest wrapper for the ListDataVerificationObjectTypeCounts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListDataVerificationObjectTypeCounts.go.html to see an example of how to use ListDataVerificationObjectTypeCountsRequest.
type ListDataVerificationObjectTypeCountsRequest struct {

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

	// A filter to return only results for a specific schema/owner.
	SchemaName *string `mandatory:"false" contributesTo:"query" name:"schemaName"`

	// A filter to return only results for a specific object type.
	// The allowed values depend on the migration's `databaseCombination`:
	// - Oracle migrations: `OracleDatabaseObjectTypes`
	// - MySQL migrations: `MySqlDatabaseObjectTypes`
	ObjectType *string `mandatory:"false" contributesTo:"query" name:"objectType"`

	// Free-text filter applied by the service to relevant fields for the report.
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// Minimum absolute deltaPercent threshold (magnitude) to return.
	// The service filters results where `abs(deltaPercent) >= minAbsDeltaPercent`.
	// Must be non-negative.
	// Note: This filter applies to the object type counts report, which uses `deltaPercent`.
	// Table row count reports use `variancePercent` instead.
	MinAbsDeltaPercent *float64 `mandatory:"false" contributesTo:"query" name:"minAbsDeltaPercent"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for deltaPercent is descending.
	SortBy ListDataVerificationObjectTypeCountsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataVerificationObjectTypeCountsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataVerificationObjectTypeCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataVerificationObjectTypeCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataVerificationObjectTypeCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataVerificationObjectTypeCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataVerificationObjectTypeCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataVerificationObjectTypeCountsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataVerificationObjectTypeCountsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataVerificationObjectTypeCountsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataVerificationObjectTypeCountsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataVerificationObjectTypeCountsResponse wrapper for the ListDataVerificationObjectTypeCounts operation
type ListDataVerificationObjectTypeCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataVerificationObjectTypeCountCollection instances
	DataVerificationObjectTypeCountCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataVerificationObjectTypeCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataVerificationObjectTypeCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataVerificationObjectTypeCountsSortByEnum Enum with underlying type: string
type ListDataVerificationObjectTypeCountsSortByEnum string

// Set of constants representing the allowable values for ListDataVerificationObjectTypeCountsSortByEnum
const (
	ListDataVerificationObjectTypeCountsSortByDeltapercent       ListDataVerificationObjectTypeCountsSortByEnum = "deltaPercent"
	ListDataVerificationObjectTypeCountsSortBySchemaname         ListDataVerificationObjectTypeCountsSortByEnum = "schemaName"
	ListDataVerificationObjectTypeCountsSortByObjecttype         ListDataVerificationObjectTypeCountsSortByEnum = "objectType"
	ListDataVerificationObjectTypeCountsSortBySourceobjectcount  ListDataVerificationObjectTypeCountsSortByEnum = "sourceObjectCount"
	ListDataVerificationObjectTypeCountsSortByTargetobjectcount  ListDataVerificationObjectTypeCountsSortByEnum = "targetObjectCount"
	ListDataVerificationObjectTypeCountsSortBySourceinvalidcount ListDataVerificationObjectTypeCountsSortByEnum = "sourceInvalidCount"
	ListDataVerificationObjectTypeCountsSortByTargetinvalidcount ListDataVerificationObjectTypeCountsSortByEnum = "targetInvalidCount"
)

var mappingListDataVerificationObjectTypeCountsSortByEnum = map[string]ListDataVerificationObjectTypeCountsSortByEnum{
	"deltaPercent":       ListDataVerificationObjectTypeCountsSortByDeltapercent,
	"schemaName":         ListDataVerificationObjectTypeCountsSortBySchemaname,
	"objectType":         ListDataVerificationObjectTypeCountsSortByObjecttype,
	"sourceObjectCount":  ListDataVerificationObjectTypeCountsSortBySourceobjectcount,
	"targetObjectCount":  ListDataVerificationObjectTypeCountsSortByTargetobjectcount,
	"sourceInvalidCount": ListDataVerificationObjectTypeCountsSortBySourceinvalidcount,
	"targetInvalidCount": ListDataVerificationObjectTypeCountsSortByTargetinvalidcount,
}

var mappingListDataVerificationObjectTypeCountsSortByEnumLowerCase = map[string]ListDataVerificationObjectTypeCountsSortByEnum{
	"deltapercent":       ListDataVerificationObjectTypeCountsSortByDeltapercent,
	"schemaname":         ListDataVerificationObjectTypeCountsSortBySchemaname,
	"objecttype":         ListDataVerificationObjectTypeCountsSortByObjecttype,
	"sourceobjectcount":  ListDataVerificationObjectTypeCountsSortBySourceobjectcount,
	"targetobjectcount":  ListDataVerificationObjectTypeCountsSortByTargetobjectcount,
	"sourceinvalidcount": ListDataVerificationObjectTypeCountsSortBySourceinvalidcount,
	"targetinvalidcount": ListDataVerificationObjectTypeCountsSortByTargetinvalidcount,
}

// GetListDataVerificationObjectTypeCountsSortByEnumValues Enumerates the set of values for ListDataVerificationObjectTypeCountsSortByEnum
func GetListDataVerificationObjectTypeCountsSortByEnumValues() []ListDataVerificationObjectTypeCountsSortByEnum {
	values := make([]ListDataVerificationObjectTypeCountsSortByEnum, 0)
	for _, v := range mappingListDataVerificationObjectTypeCountsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataVerificationObjectTypeCountsSortByEnumStringValues Enumerates the set of values in String for ListDataVerificationObjectTypeCountsSortByEnum
func GetListDataVerificationObjectTypeCountsSortByEnumStringValues() []string {
	return []string{
		"deltaPercent",
		"schemaName",
		"objectType",
		"sourceObjectCount",
		"targetObjectCount",
		"sourceInvalidCount",
		"targetInvalidCount",
	}
}

// GetMappingListDataVerificationObjectTypeCountsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataVerificationObjectTypeCountsSortByEnum(val string) (ListDataVerificationObjectTypeCountsSortByEnum, bool) {
	enum, ok := mappingListDataVerificationObjectTypeCountsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataVerificationObjectTypeCountsSortOrderEnum Enum with underlying type: string
type ListDataVerificationObjectTypeCountsSortOrderEnum string

// Set of constants representing the allowable values for ListDataVerificationObjectTypeCountsSortOrderEnum
const (
	ListDataVerificationObjectTypeCountsSortOrderAsc  ListDataVerificationObjectTypeCountsSortOrderEnum = "ASC"
	ListDataVerificationObjectTypeCountsSortOrderDesc ListDataVerificationObjectTypeCountsSortOrderEnum = "DESC"
)

var mappingListDataVerificationObjectTypeCountsSortOrderEnum = map[string]ListDataVerificationObjectTypeCountsSortOrderEnum{
	"ASC":  ListDataVerificationObjectTypeCountsSortOrderAsc,
	"DESC": ListDataVerificationObjectTypeCountsSortOrderDesc,
}

var mappingListDataVerificationObjectTypeCountsSortOrderEnumLowerCase = map[string]ListDataVerificationObjectTypeCountsSortOrderEnum{
	"asc":  ListDataVerificationObjectTypeCountsSortOrderAsc,
	"desc": ListDataVerificationObjectTypeCountsSortOrderDesc,
}

// GetListDataVerificationObjectTypeCountsSortOrderEnumValues Enumerates the set of values for ListDataVerificationObjectTypeCountsSortOrderEnum
func GetListDataVerificationObjectTypeCountsSortOrderEnumValues() []ListDataVerificationObjectTypeCountsSortOrderEnum {
	values := make([]ListDataVerificationObjectTypeCountsSortOrderEnum, 0)
	for _, v := range mappingListDataVerificationObjectTypeCountsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataVerificationObjectTypeCountsSortOrderEnumStringValues Enumerates the set of values in String for ListDataVerificationObjectTypeCountsSortOrderEnum
func GetListDataVerificationObjectTypeCountsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataVerificationObjectTypeCountsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataVerificationObjectTypeCountsSortOrderEnum(val string) (ListDataVerificationObjectTypeCountsSortOrderEnum, bool) {
	enum, ok := mappingListDataVerificationObjectTypeCountsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
