// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalExadataStorageServersRequest wrapper for the ListExternalExadataStorageServers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalExadataStorageServers.go.html to see an example of how to use ListExternalExadataStorageServersRequest.
type ListExternalExadataStorageServersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExternalExadataInfrastructureId *string `mandatory:"true" contributesTo:"query" name:"externalExadataInfrastructureId"`

	// The optional single value query filter parameter on the entity display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListExternalExadataStorageServersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalExadataStorageServersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalExadataStorageServersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalExadataStorageServersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalExadataStorageServersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalExadataStorageServersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalExadataStorageServersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalExadataStorageServersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalExadataStorageServersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalExadataStorageServersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalExadataStorageServersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalExadataStorageServersResponse wrapper for the ListExternalExadataStorageServers operation
type ListExternalExadataStorageServersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalExadataStorageServerCollection instances
	ExternalExadataStorageServerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalExadataStorageServersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalExadataStorageServersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalExadataStorageServersSortByEnum Enum with underlying type: string
type ListExternalExadataStorageServersSortByEnum string

// Set of constants representing the allowable values for ListExternalExadataStorageServersSortByEnum
const (
	ListExternalExadataStorageServersSortByTimecreated ListExternalExadataStorageServersSortByEnum = "TIMECREATED"
	ListExternalExadataStorageServersSortByName        ListExternalExadataStorageServersSortByEnum = "NAME"
)

var mappingListExternalExadataStorageServersSortByEnum = map[string]ListExternalExadataStorageServersSortByEnum{
	"TIMECREATED": ListExternalExadataStorageServersSortByTimecreated,
	"NAME":        ListExternalExadataStorageServersSortByName,
}

var mappingListExternalExadataStorageServersSortByEnumLowerCase = map[string]ListExternalExadataStorageServersSortByEnum{
	"timecreated": ListExternalExadataStorageServersSortByTimecreated,
	"name":        ListExternalExadataStorageServersSortByName,
}

// GetListExternalExadataStorageServersSortByEnumValues Enumerates the set of values for ListExternalExadataStorageServersSortByEnum
func GetListExternalExadataStorageServersSortByEnumValues() []ListExternalExadataStorageServersSortByEnum {
	values := make([]ListExternalExadataStorageServersSortByEnum, 0)
	for _, v := range mappingListExternalExadataStorageServersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalExadataStorageServersSortByEnumStringValues Enumerates the set of values in String for ListExternalExadataStorageServersSortByEnum
func GetListExternalExadataStorageServersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListExternalExadataStorageServersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalExadataStorageServersSortByEnum(val string) (ListExternalExadataStorageServersSortByEnum, bool) {
	enum, ok := mappingListExternalExadataStorageServersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalExadataStorageServersSortOrderEnum Enum with underlying type: string
type ListExternalExadataStorageServersSortOrderEnum string

// Set of constants representing the allowable values for ListExternalExadataStorageServersSortOrderEnum
const (
	ListExternalExadataStorageServersSortOrderAsc  ListExternalExadataStorageServersSortOrderEnum = "ASC"
	ListExternalExadataStorageServersSortOrderDesc ListExternalExadataStorageServersSortOrderEnum = "DESC"
)

var mappingListExternalExadataStorageServersSortOrderEnum = map[string]ListExternalExadataStorageServersSortOrderEnum{
	"ASC":  ListExternalExadataStorageServersSortOrderAsc,
	"DESC": ListExternalExadataStorageServersSortOrderDesc,
}

var mappingListExternalExadataStorageServersSortOrderEnumLowerCase = map[string]ListExternalExadataStorageServersSortOrderEnum{
	"asc":  ListExternalExadataStorageServersSortOrderAsc,
	"desc": ListExternalExadataStorageServersSortOrderDesc,
}

// GetListExternalExadataStorageServersSortOrderEnumValues Enumerates the set of values for ListExternalExadataStorageServersSortOrderEnum
func GetListExternalExadataStorageServersSortOrderEnumValues() []ListExternalExadataStorageServersSortOrderEnum {
	values := make([]ListExternalExadataStorageServersSortOrderEnum, 0)
	for _, v := range mappingListExternalExadataStorageServersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalExadataStorageServersSortOrderEnumStringValues Enumerates the set of values in String for ListExternalExadataStorageServersSortOrderEnum
func GetListExternalExadataStorageServersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalExadataStorageServersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalExadataStorageServersSortOrderEnum(val string) (ListExternalExadataStorageServersSortOrderEnum, bool) {
	enum, ok := mappingListExternalExadataStorageServersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
