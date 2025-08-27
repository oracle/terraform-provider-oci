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

// ListOutboundReplicationsRequest wrapper for the ListOutboundReplications operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOutboundReplications.go.html to see an example of how to use ListOutboundReplicationsRequest.
type ListOutboundReplicationsRequest struct {

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
	SortOrder ListOutboundReplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used.
	// The default sort by field is ‘replicaHost’.
	// The default sort order for ‘replicaHost’ is ascending.
	// The ‘replicaHost’ sort order is case-sensitive.
	SortBy ListOutboundReplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOutboundReplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOutboundReplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOutboundReplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOutboundReplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOutboundReplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOutboundReplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOutboundReplicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOutboundReplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOutboundReplicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOutboundReplicationsResponse wrapper for the ListOutboundReplications operation
type ListOutboundReplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedMySqlDatabaseOutboundReplicationCollection instances
	ManagedMySqlDatabaseOutboundReplicationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOutboundReplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOutboundReplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOutboundReplicationsSortOrderEnum Enum with underlying type: string
type ListOutboundReplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListOutboundReplicationsSortOrderEnum
const (
	ListOutboundReplicationsSortOrderAsc  ListOutboundReplicationsSortOrderEnum = "ASC"
	ListOutboundReplicationsSortOrderDesc ListOutboundReplicationsSortOrderEnum = "DESC"
)

var mappingListOutboundReplicationsSortOrderEnum = map[string]ListOutboundReplicationsSortOrderEnum{
	"ASC":  ListOutboundReplicationsSortOrderAsc,
	"DESC": ListOutboundReplicationsSortOrderDesc,
}

var mappingListOutboundReplicationsSortOrderEnumLowerCase = map[string]ListOutboundReplicationsSortOrderEnum{
	"asc":  ListOutboundReplicationsSortOrderAsc,
	"desc": ListOutboundReplicationsSortOrderDesc,
}

// GetListOutboundReplicationsSortOrderEnumValues Enumerates the set of values for ListOutboundReplicationsSortOrderEnum
func GetListOutboundReplicationsSortOrderEnumValues() []ListOutboundReplicationsSortOrderEnum {
	values := make([]ListOutboundReplicationsSortOrderEnum, 0)
	for _, v := range mappingListOutboundReplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOutboundReplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListOutboundReplicationsSortOrderEnum
func GetListOutboundReplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOutboundReplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOutboundReplicationsSortOrderEnum(val string) (ListOutboundReplicationsSortOrderEnum, bool) {
	enum, ok := mappingListOutboundReplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOutboundReplicationsSortByEnum Enum with underlying type: string
type ListOutboundReplicationsSortByEnum string

// Set of constants representing the allowable values for ListOutboundReplicationsSortByEnum
const (
	ListOutboundReplicationsSortByReplicahost ListOutboundReplicationsSortByEnum = "replicaHost"
)

var mappingListOutboundReplicationsSortByEnum = map[string]ListOutboundReplicationsSortByEnum{
	"replicaHost": ListOutboundReplicationsSortByReplicahost,
}

var mappingListOutboundReplicationsSortByEnumLowerCase = map[string]ListOutboundReplicationsSortByEnum{
	"replicahost": ListOutboundReplicationsSortByReplicahost,
}

// GetListOutboundReplicationsSortByEnumValues Enumerates the set of values for ListOutboundReplicationsSortByEnum
func GetListOutboundReplicationsSortByEnumValues() []ListOutboundReplicationsSortByEnum {
	values := make([]ListOutboundReplicationsSortByEnum, 0)
	for _, v := range mappingListOutboundReplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOutboundReplicationsSortByEnumStringValues Enumerates the set of values in String for ListOutboundReplicationsSortByEnum
func GetListOutboundReplicationsSortByEnumStringValues() []string {
	return []string{
		"replicaHost",
	}
}

// GetMappingListOutboundReplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOutboundReplicationsSortByEnum(val string) (ListOutboundReplicationsSortByEnum, bool) {
	enum, ok := mappingListOutboundReplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
