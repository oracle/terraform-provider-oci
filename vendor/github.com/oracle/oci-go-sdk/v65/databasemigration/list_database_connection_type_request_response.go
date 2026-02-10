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

// ListDatabaseConnectionTypeRequest wrapper for the ListDatabaseConnectionType operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListDatabaseConnectionType.go.html to see an example of how to use ListDatabaseConnectionTypeRequest.
type ListDatabaseConnectionTypeRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The array of technology types.
	TechnologyType []TechnologyTypeEnum `contributesTo:"query" name:"technologyType" omitEmpty:"true" collectionFormat:"multi"`

	// The array of connection types.
	ConnectionType []ConnectionTypeEnum `contributesTo:"query" name:"connectionType" omitEmpty:"true" collectionFormat:"multi"`

	// The OCID of the source connection.
	SourceConnectionId *string `mandatory:"false" contributesTo:"query" name:"sourceConnectionId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseConnectionTypeSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseConnectionTypeSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseConnectionTypeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseConnectionTypeRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseConnectionTypeRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseConnectionTypeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseConnectionTypeRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.TechnologyType {
		if _, ok := GetMappingTechnologyTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", val, strings.Join(GetTechnologyTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ConnectionType {
		if _, ok := GetMappingConnectionTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", val, strings.Join(GetConnectionTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListDatabaseConnectionTypeSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseConnectionTypeSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseConnectionTypeSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseConnectionTypeSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseConnectionTypeResponse wrapper for the ListDatabaseConnectionType operation
type ListDatabaseConnectionTypeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseConnectionTypeCollection instances
	DatabaseConnectionTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseConnectionTypeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseConnectionTypeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseConnectionTypeSortByEnum Enum with underlying type: string
type ListDatabaseConnectionTypeSortByEnum string

// Set of constants representing the allowable values for ListDatabaseConnectionTypeSortByEnum
const (
	ListDatabaseConnectionTypeSortByTimecreated ListDatabaseConnectionTypeSortByEnum = "timeCreated"
	ListDatabaseConnectionTypeSortByDisplayname ListDatabaseConnectionTypeSortByEnum = "displayName"
)

var mappingListDatabaseConnectionTypeSortByEnum = map[string]ListDatabaseConnectionTypeSortByEnum{
	"timeCreated": ListDatabaseConnectionTypeSortByTimecreated,
	"displayName": ListDatabaseConnectionTypeSortByDisplayname,
}

var mappingListDatabaseConnectionTypeSortByEnumLowerCase = map[string]ListDatabaseConnectionTypeSortByEnum{
	"timecreated": ListDatabaseConnectionTypeSortByTimecreated,
	"displayname": ListDatabaseConnectionTypeSortByDisplayname,
}

// GetListDatabaseConnectionTypeSortByEnumValues Enumerates the set of values for ListDatabaseConnectionTypeSortByEnum
func GetListDatabaseConnectionTypeSortByEnumValues() []ListDatabaseConnectionTypeSortByEnum {
	values := make([]ListDatabaseConnectionTypeSortByEnum, 0)
	for _, v := range mappingListDatabaseConnectionTypeSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseConnectionTypeSortByEnumStringValues Enumerates the set of values in String for ListDatabaseConnectionTypeSortByEnum
func GetListDatabaseConnectionTypeSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseConnectionTypeSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseConnectionTypeSortByEnum(val string) (ListDatabaseConnectionTypeSortByEnum, bool) {
	enum, ok := mappingListDatabaseConnectionTypeSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseConnectionTypeSortOrderEnum Enum with underlying type: string
type ListDatabaseConnectionTypeSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseConnectionTypeSortOrderEnum
const (
	ListDatabaseConnectionTypeSortOrderAsc  ListDatabaseConnectionTypeSortOrderEnum = "ASC"
	ListDatabaseConnectionTypeSortOrderDesc ListDatabaseConnectionTypeSortOrderEnum = "DESC"
)

var mappingListDatabaseConnectionTypeSortOrderEnum = map[string]ListDatabaseConnectionTypeSortOrderEnum{
	"ASC":  ListDatabaseConnectionTypeSortOrderAsc,
	"DESC": ListDatabaseConnectionTypeSortOrderDesc,
}

var mappingListDatabaseConnectionTypeSortOrderEnumLowerCase = map[string]ListDatabaseConnectionTypeSortOrderEnum{
	"asc":  ListDatabaseConnectionTypeSortOrderAsc,
	"desc": ListDatabaseConnectionTypeSortOrderDesc,
}

// GetListDatabaseConnectionTypeSortOrderEnumValues Enumerates the set of values for ListDatabaseConnectionTypeSortOrderEnum
func GetListDatabaseConnectionTypeSortOrderEnumValues() []ListDatabaseConnectionTypeSortOrderEnum {
	values := make([]ListDatabaseConnectionTypeSortOrderEnum, 0)
	for _, v := range mappingListDatabaseConnectionTypeSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseConnectionTypeSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseConnectionTypeSortOrderEnum
func GetListDatabaseConnectionTypeSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseConnectionTypeSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseConnectionTypeSortOrderEnum(val string) (ListDatabaseConnectionTypeSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseConnectionTypeSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
