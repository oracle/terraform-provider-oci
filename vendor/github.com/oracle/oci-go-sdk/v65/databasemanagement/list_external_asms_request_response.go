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

// ListExternalAsmsRequest wrapper for the ListExternalAsms operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsms.go.html to see an example of how to use ListExternalAsmsRequest.
type ListExternalAsmsRequest struct {

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
	SortBy ListExternalAsmsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalAsmsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalAsmsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalAsmsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalAsmsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalAsmsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalAsmsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalAsmsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalAsmsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalAsmsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalAsmsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalAsmsResponse wrapper for the ListExternalAsms operation
type ListExternalAsmsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalAsmCollection instances
	ExternalAsmCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalAsmsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalAsmsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalAsmsSortByEnum Enum with underlying type: string
type ListExternalAsmsSortByEnum string

// Set of constants representing the allowable values for ListExternalAsmsSortByEnum
const (
	ListExternalAsmsSortByTimecreated ListExternalAsmsSortByEnum = "TIMECREATED"
	ListExternalAsmsSortByDisplayname ListExternalAsmsSortByEnum = "DISPLAYNAME"
)

var mappingListExternalAsmsSortByEnum = map[string]ListExternalAsmsSortByEnum{
	"TIMECREATED": ListExternalAsmsSortByTimecreated,
	"DISPLAYNAME": ListExternalAsmsSortByDisplayname,
}

var mappingListExternalAsmsSortByEnumLowerCase = map[string]ListExternalAsmsSortByEnum{
	"timecreated": ListExternalAsmsSortByTimecreated,
	"displayname": ListExternalAsmsSortByDisplayname,
}

// GetListExternalAsmsSortByEnumValues Enumerates the set of values for ListExternalAsmsSortByEnum
func GetListExternalAsmsSortByEnumValues() []ListExternalAsmsSortByEnum {
	values := make([]ListExternalAsmsSortByEnum, 0)
	for _, v := range mappingListExternalAsmsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmsSortByEnumStringValues Enumerates the set of values in String for ListExternalAsmsSortByEnum
func GetListExternalAsmsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalAsmsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmsSortByEnum(val string) (ListExternalAsmsSortByEnum, bool) {
	enum, ok := mappingListExternalAsmsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalAsmsSortOrderEnum Enum with underlying type: string
type ListExternalAsmsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalAsmsSortOrderEnum
const (
	ListExternalAsmsSortOrderAsc  ListExternalAsmsSortOrderEnum = "ASC"
	ListExternalAsmsSortOrderDesc ListExternalAsmsSortOrderEnum = "DESC"
)

var mappingListExternalAsmsSortOrderEnum = map[string]ListExternalAsmsSortOrderEnum{
	"ASC":  ListExternalAsmsSortOrderAsc,
	"DESC": ListExternalAsmsSortOrderDesc,
}

var mappingListExternalAsmsSortOrderEnumLowerCase = map[string]ListExternalAsmsSortOrderEnum{
	"asc":  ListExternalAsmsSortOrderAsc,
	"desc": ListExternalAsmsSortOrderDesc,
}

// GetListExternalAsmsSortOrderEnumValues Enumerates the set of values for ListExternalAsmsSortOrderEnum
func GetListExternalAsmsSortOrderEnumValues() []ListExternalAsmsSortOrderEnum {
	values := make([]ListExternalAsmsSortOrderEnum, 0)
	for _, v := range mappingListExternalAsmsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmsSortOrderEnumStringValues Enumerates the set of values in String for ListExternalAsmsSortOrderEnum
func GetListExternalAsmsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalAsmsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmsSortOrderEnum(val string) (ListExternalAsmsSortOrderEnum, bool) {
	enum, ok := mappingListExternalAsmsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
