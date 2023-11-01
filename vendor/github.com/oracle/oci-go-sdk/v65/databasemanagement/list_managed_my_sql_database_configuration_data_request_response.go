// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedMySqlDatabaseConfigurationDataRequest wrapper for the ListManagedMySqlDatabaseConfigurationData operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedMySqlDatabaseConfigurationData.go.html to see an example of how to use ListManagedMySqlDatabaseConfigurationDataRequest.
type ListManagedMySqlDatabaseConfigurationDataRequest struct {

	// The OCID of the Managed MySQL Database.
	ManagedMySqlDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedMySqlDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder ListManagedMySqlDatabaseConfigurationDataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListManagedMySqlDatabaseConfigurationDataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedMySqlDatabaseConfigurationDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedMySqlDatabaseConfigurationDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedMySqlDatabaseConfigurationDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedMySqlDatabaseConfigurationDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedMySqlDatabaseConfigurationDataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedMySqlDatabaseConfigurationDataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedMySqlDatabaseConfigurationDataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedMySqlDatabaseConfigurationDataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedMySqlDatabaseConfigurationDataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedMySqlDatabaseConfigurationDataResponse wrapper for the ListManagedMySqlDatabaseConfigurationData operation
type ListManagedMySqlDatabaseConfigurationDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MySqlConfigurationDataCollection instances
	MySqlConfigurationDataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedMySqlDatabaseConfigurationDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedMySqlDatabaseConfigurationDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedMySqlDatabaseConfigurationDataSortOrderEnum Enum with underlying type: string
type ListManagedMySqlDatabaseConfigurationDataSortOrderEnum string

// Set of constants representing the allowable values for ListManagedMySqlDatabaseConfigurationDataSortOrderEnum
const (
	ListManagedMySqlDatabaseConfigurationDataSortOrderAsc  ListManagedMySqlDatabaseConfigurationDataSortOrderEnum = "ASC"
	ListManagedMySqlDatabaseConfigurationDataSortOrderDesc ListManagedMySqlDatabaseConfigurationDataSortOrderEnum = "DESC"
)

var mappingListManagedMySqlDatabaseConfigurationDataSortOrderEnum = map[string]ListManagedMySqlDatabaseConfigurationDataSortOrderEnum{
	"ASC":  ListManagedMySqlDatabaseConfigurationDataSortOrderAsc,
	"DESC": ListManagedMySqlDatabaseConfigurationDataSortOrderDesc,
}

var mappingListManagedMySqlDatabaseConfigurationDataSortOrderEnumLowerCase = map[string]ListManagedMySqlDatabaseConfigurationDataSortOrderEnum{
	"asc":  ListManagedMySqlDatabaseConfigurationDataSortOrderAsc,
	"desc": ListManagedMySqlDatabaseConfigurationDataSortOrderDesc,
}

// GetListManagedMySqlDatabaseConfigurationDataSortOrderEnumValues Enumerates the set of values for ListManagedMySqlDatabaseConfigurationDataSortOrderEnum
func GetListManagedMySqlDatabaseConfigurationDataSortOrderEnumValues() []ListManagedMySqlDatabaseConfigurationDataSortOrderEnum {
	values := make([]ListManagedMySqlDatabaseConfigurationDataSortOrderEnum, 0)
	for _, v := range mappingListManagedMySqlDatabaseConfigurationDataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedMySqlDatabaseConfigurationDataSortOrderEnumStringValues Enumerates the set of values in String for ListManagedMySqlDatabaseConfigurationDataSortOrderEnum
func GetListManagedMySqlDatabaseConfigurationDataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedMySqlDatabaseConfigurationDataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedMySqlDatabaseConfigurationDataSortOrderEnum(val string) (ListManagedMySqlDatabaseConfigurationDataSortOrderEnum, bool) {
	enum, ok := mappingListManagedMySqlDatabaseConfigurationDataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedMySqlDatabaseConfigurationDataSortByEnum Enum with underlying type: string
type ListManagedMySqlDatabaseConfigurationDataSortByEnum string

// Set of constants representing the allowable values for ListManagedMySqlDatabaseConfigurationDataSortByEnum
const (
	ListManagedMySqlDatabaseConfigurationDataSortByTimecreated ListManagedMySqlDatabaseConfigurationDataSortByEnum = "TIMECREATED"
	ListManagedMySqlDatabaseConfigurationDataSortByName        ListManagedMySqlDatabaseConfigurationDataSortByEnum = "NAME"
)

var mappingListManagedMySqlDatabaseConfigurationDataSortByEnum = map[string]ListManagedMySqlDatabaseConfigurationDataSortByEnum{
	"TIMECREATED": ListManagedMySqlDatabaseConfigurationDataSortByTimecreated,
	"NAME":        ListManagedMySqlDatabaseConfigurationDataSortByName,
}

var mappingListManagedMySqlDatabaseConfigurationDataSortByEnumLowerCase = map[string]ListManagedMySqlDatabaseConfigurationDataSortByEnum{
	"timecreated": ListManagedMySqlDatabaseConfigurationDataSortByTimecreated,
	"name":        ListManagedMySqlDatabaseConfigurationDataSortByName,
}

// GetListManagedMySqlDatabaseConfigurationDataSortByEnumValues Enumerates the set of values for ListManagedMySqlDatabaseConfigurationDataSortByEnum
func GetListManagedMySqlDatabaseConfigurationDataSortByEnumValues() []ListManagedMySqlDatabaseConfigurationDataSortByEnum {
	values := make([]ListManagedMySqlDatabaseConfigurationDataSortByEnum, 0)
	for _, v := range mappingListManagedMySqlDatabaseConfigurationDataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedMySqlDatabaseConfigurationDataSortByEnumStringValues Enumerates the set of values in String for ListManagedMySqlDatabaseConfigurationDataSortByEnum
func GetListManagedMySqlDatabaseConfigurationDataSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListManagedMySqlDatabaseConfigurationDataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedMySqlDatabaseConfigurationDataSortByEnum(val string) (ListManagedMySqlDatabaseConfigurationDataSortByEnum, bool) {
	enum, ok := mappingListManagedMySqlDatabaseConfigurationDataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
