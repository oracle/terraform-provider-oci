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

// ListExternalDbSystemConnectorsRequest wrapper for the ListExternalDbSystemConnectors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalDbSystemConnectors.go.html to see an example of how to use ListExternalDbSystemConnectorsRequest.
type ListExternalDbSystemConnectorsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system.
	ExternalDbSystemId *string `mandatory:"false" contributesTo:"query" name:"externalDbSystemId"`

	// A filter to only return the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for `TIMECREATED` is descending and the default sort order for `DISPLAYNAME` is ascending.
	// The `DISPLAYNAME` sort order is case-sensitive.
	SortBy ListExternalDbSystemConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalDbSystemConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalDbSystemConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalDbSystemConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalDbSystemConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalDbSystemConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalDbSystemConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalDbSystemConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalDbSystemConnectorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalDbSystemConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalDbSystemConnectorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalDbSystemConnectorsResponse wrapper for the ListExternalDbSystemConnectors operation
type ListExternalDbSystemConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalDbSystemConnectorCollection instances
	ExternalDbSystemConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalDbSystemConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalDbSystemConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalDbSystemConnectorsSortByEnum Enum with underlying type: string
type ListExternalDbSystemConnectorsSortByEnum string

// Set of constants representing the allowable values for ListExternalDbSystemConnectorsSortByEnum
const (
	ListExternalDbSystemConnectorsSortByTimecreated ListExternalDbSystemConnectorsSortByEnum = "TIMECREATED"
	ListExternalDbSystemConnectorsSortByDisplayname ListExternalDbSystemConnectorsSortByEnum = "DISPLAYNAME"
)

var mappingListExternalDbSystemConnectorsSortByEnum = map[string]ListExternalDbSystemConnectorsSortByEnum{
	"TIMECREATED": ListExternalDbSystemConnectorsSortByTimecreated,
	"DISPLAYNAME": ListExternalDbSystemConnectorsSortByDisplayname,
}

var mappingListExternalDbSystemConnectorsSortByEnumLowerCase = map[string]ListExternalDbSystemConnectorsSortByEnum{
	"timecreated": ListExternalDbSystemConnectorsSortByTimecreated,
	"displayname": ListExternalDbSystemConnectorsSortByDisplayname,
}

// GetListExternalDbSystemConnectorsSortByEnumValues Enumerates the set of values for ListExternalDbSystemConnectorsSortByEnum
func GetListExternalDbSystemConnectorsSortByEnumValues() []ListExternalDbSystemConnectorsSortByEnum {
	values := make([]ListExternalDbSystemConnectorsSortByEnum, 0)
	for _, v := range mappingListExternalDbSystemConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDbSystemConnectorsSortByEnumStringValues Enumerates the set of values in String for ListExternalDbSystemConnectorsSortByEnum
func GetListExternalDbSystemConnectorsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalDbSystemConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDbSystemConnectorsSortByEnum(val string) (ListExternalDbSystemConnectorsSortByEnum, bool) {
	enum, ok := mappingListExternalDbSystemConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalDbSystemConnectorsSortOrderEnum Enum with underlying type: string
type ListExternalDbSystemConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalDbSystemConnectorsSortOrderEnum
const (
	ListExternalDbSystemConnectorsSortOrderAsc  ListExternalDbSystemConnectorsSortOrderEnum = "ASC"
	ListExternalDbSystemConnectorsSortOrderDesc ListExternalDbSystemConnectorsSortOrderEnum = "DESC"
)

var mappingListExternalDbSystemConnectorsSortOrderEnum = map[string]ListExternalDbSystemConnectorsSortOrderEnum{
	"ASC":  ListExternalDbSystemConnectorsSortOrderAsc,
	"DESC": ListExternalDbSystemConnectorsSortOrderDesc,
}

var mappingListExternalDbSystemConnectorsSortOrderEnumLowerCase = map[string]ListExternalDbSystemConnectorsSortOrderEnum{
	"asc":  ListExternalDbSystemConnectorsSortOrderAsc,
	"desc": ListExternalDbSystemConnectorsSortOrderDesc,
}

// GetListExternalDbSystemConnectorsSortOrderEnumValues Enumerates the set of values for ListExternalDbSystemConnectorsSortOrderEnum
func GetListExternalDbSystemConnectorsSortOrderEnumValues() []ListExternalDbSystemConnectorsSortOrderEnum {
	values := make([]ListExternalDbSystemConnectorsSortOrderEnum, 0)
	for _, v := range mappingListExternalDbSystemConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalDbSystemConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListExternalDbSystemConnectorsSortOrderEnum
func GetListExternalDbSystemConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalDbSystemConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalDbSystemConnectorsSortOrderEnum(val string) (ListExternalDbSystemConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListExternalDbSystemConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
