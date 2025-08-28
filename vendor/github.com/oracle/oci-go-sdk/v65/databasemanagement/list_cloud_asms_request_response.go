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

// ListCloudAsmsRequest wrapper for the ListCloudAsms operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudAsms.go.html to see an example of how to use ListCloudAsmsRequest.
type ListCloudAsmsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"false" contributesTo:"query" name:"cloudDbSystemId"`

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
	SortBy ListCloudAsmsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudAsmsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudAsmsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudAsmsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudAsmsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudAsmsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudAsmsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudAsmsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudAsmsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudAsmsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudAsmsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudAsmsResponse wrapper for the ListCloudAsms operation
type ListCloudAsmsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudAsmCollection instances
	CloudAsmCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudAsmsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudAsmsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudAsmsSortByEnum Enum with underlying type: string
type ListCloudAsmsSortByEnum string

// Set of constants representing the allowable values for ListCloudAsmsSortByEnum
const (
	ListCloudAsmsSortByTimecreated ListCloudAsmsSortByEnum = "TIMECREATED"
	ListCloudAsmsSortByDisplayname ListCloudAsmsSortByEnum = "DISPLAYNAME"
)

var mappingListCloudAsmsSortByEnum = map[string]ListCloudAsmsSortByEnum{
	"TIMECREATED": ListCloudAsmsSortByTimecreated,
	"DISPLAYNAME": ListCloudAsmsSortByDisplayname,
}

var mappingListCloudAsmsSortByEnumLowerCase = map[string]ListCloudAsmsSortByEnum{
	"timecreated": ListCloudAsmsSortByTimecreated,
	"displayname": ListCloudAsmsSortByDisplayname,
}

// GetListCloudAsmsSortByEnumValues Enumerates the set of values for ListCloudAsmsSortByEnum
func GetListCloudAsmsSortByEnumValues() []ListCloudAsmsSortByEnum {
	values := make([]ListCloudAsmsSortByEnum, 0)
	for _, v := range mappingListCloudAsmsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmsSortByEnumStringValues Enumerates the set of values in String for ListCloudAsmsSortByEnum
func GetListCloudAsmsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudAsmsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmsSortByEnum(val string) (ListCloudAsmsSortByEnum, bool) {
	enum, ok := mappingListCloudAsmsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudAsmsSortOrderEnum Enum with underlying type: string
type ListCloudAsmsSortOrderEnum string

// Set of constants representing the allowable values for ListCloudAsmsSortOrderEnum
const (
	ListCloudAsmsSortOrderAsc  ListCloudAsmsSortOrderEnum = "ASC"
	ListCloudAsmsSortOrderDesc ListCloudAsmsSortOrderEnum = "DESC"
)

var mappingListCloudAsmsSortOrderEnum = map[string]ListCloudAsmsSortOrderEnum{
	"ASC":  ListCloudAsmsSortOrderAsc,
	"DESC": ListCloudAsmsSortOrderDesc,
}

var mappingListCloudAsmsSortOrderEnumLowerCase = map[string]ListCloudAsmsSortOrderEnum{
	"asc":  ListCloudAsmsSortOrderAsc,
	"desc": ListCloudAsmsSortOrderDesc,
}

// GetListCloudAsmsSortOrderEnumValues Enumerates the set of values for ListCloudAsmsSortOrderEnum
func GetListCloudAsmsSortOrderEnumValues() []ListCloudAsmsSortOrderEnum {
	values := make([]ListCloudAsmsSortOrderEnum, 0)
	for _, v := range mappingListCloudAsmsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmsSortOrderEnumStringValues Enumerates the set of values in String for ListCloudAsmsSortOrderEnum
func GetListCloudAsmsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudAsmsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmsSortOrderEnum(val string) (ListCloudAsmsSortOrderEnum, bool) {
	enum, ok := mappingListCloudAsmsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
