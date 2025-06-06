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

// ListInboundReplicationsRequest wrapper for the ListInboundReplications operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListInboundReplications.go.html to see an example of how to use ListInboundReplicationsRequest.
type ListInboundReplicationsRequest struct {

	// The OCID of the Managed MySQL Database.
	ManagedMySqlDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedMySqlDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListInboundReplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used.
	// The default sort by field is ‘sourceHost’.
	// The default sort order for ‘sourceHost’ is ascending.
	// The ‘sourceHost’ sort order is case-sensitive.
	SortBy ListInboundReplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInboundReplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInboundReplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInboundReplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInboundReplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInboundReplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInboundReplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInboundReplicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInboundReplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInboundReplicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInboundReplicationsResponse wrapper for the ListInboundReplications operation
type ListInboundReplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedMySqlDatabaseInboundReplicationCollection instances
	ManagedMySqlDatabaseInboundReplicationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInboundReplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInboundReplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInboundReplicationsSortOrderEnum Enum with underlying type: string
type ListInboundReplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListInboundReplicationsSortOrderEnum
const (
	ListInboundReplicationsSortOrderAsc  ListInboundReplicationsSortOrderEnum = "ASC"
	ListInboundReplicationsSortOrderDesc ListInboundReplicationsSortOrderEnum = "DESC"
)

var mappingListInboundReplicationsSortOrderEnum = map[string]ListInboundReplicationsSortOrderEnum{
	"ASC":  ListInboundReplicationsSortOrderAsc,
	"DESC": ListInboundReplicationsSortOrderDesc,
}

var mappingListInboundReplicationsSortOrderEnumLowerCase = map[string]ListInboundReplicationsSortOrderEnum{
	"asc":  ListInboundReplicationsSortOrderAsc,
	"desc": ListInboundReplicationsSortOrderDesc,
}

// GetListInboundReplicationsSortOrderEnumValues Enumerates the set of values for ListInboundReplicationsSortOrderEnum
func GetListInboundReplicationsSortOrderEnumValues() []ListInboundReplicationsSortOrderEnum {
	values := make([]ListInboundReplicationsSortOrderEnum, 0)
	for _, v := range mappingListInboundReplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInboundReplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListInboundReplicationsSortOrderEnum
func GetListInboundReplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInboundReplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInboundReplicationsSortOrderEnum(val string) (ListInboundReplicationsSortOrderEnum, bool) {
	enum, ok := mappingListInboundReplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInboundReplicationsSortByEnum Enum with underlying type: string
type ListInboundReplicationsSortByEnum string

// Set of constants representing the allowable values for ListInboundReplicationsSortByEnum
const (
	ListInboundReplicationsSortBySourcehost ListInboundReplicationsSortByEnum = "sourceHost"
)

var mappingListInboundReplicationsSortByEnum = map[string]ListInboundReplicationsSortByEnum{
	"sourceHost": ListInboundReplicationsSortBySourcehost,
}

var mappingListInboundReplicationsSortByEnumLowerCase = map[string]ListInboundReplicationsSortByEnum{
	"sourcehost": ListInboundReplicationsSortBySourcehost,
}

// GetListInboundReplicationsSortByEnumValues Enumerates the set of values for ListInboundReplicationsSortByEnum
func GetListInboundReplicationsSortByEnumValues() []ListInboundReplicationsSortByEnum {
	values := make([]ListInboundReplicationsSortByEnum, 0)
	for _, v := range mappingListInboundReplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInboundReplicationsSortByEnumStringValues Enumerates the set of values in String for ListInboundReplicationsSortByEnum
func GetListInboundReplicationsSortByEnumStringValues() []string {
	return []string{
		"sourceHost",
	}
}

// GetMappingListInboundReplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInboundReplicationsSortByEnum(val string) (ListInboundReplicationsSortByEnum, bool) {
	enum, ok := mappingListInboundReplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
