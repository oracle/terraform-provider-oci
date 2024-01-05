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

// ListExternalAsmInstancesRequest wrapper for the ListExternalAsmInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalAsmInstances.go.html to see an example of how to use ListExternalAsmInstancesRequest.
type ListExternalAsmInstancesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external ASM.
	ExternalAsmId *string `mandatory:"false" contributesTo:"query" name:"externalAsmId"`

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
	SortBy ListExternalAsmInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalAsmInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalAsmInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalAsmInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalAsmInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalAsmInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalAsmInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalAsmInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalAsmInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalAsmInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalAsmInstancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalAsmInstancesResponse wrapper for the ListExternalAsmInstances operation
type ListExternalAsmInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalAsmInstanceCollection instances
	ExternalAsmInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalAsmInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalAsmInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalAsmInstancesSortByEnum Enum with underlying type: string
type ListExternalAsmInstancesSortByEnum string

// Set of constants representing the allowable values for ListExternalAsmInstancesSortByEnum
const (
	ListExternalAsmInstancesSortByTimecreated ListExternalAsmInstancesSortByEnum = "TIMECREATED"
	ListExternalAsmInstancesSortByDisplayname ListExternalAsmInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListExternalAsmInstancesSortByEnum = map[string]ListExternalAsmInstancesSortByEnum{
	"TIMECREATED": ListExternalAsmInstancesSortByTimecreated,
	"DISPLAYNAME": ListExternalAsmInstancesSortByDisplayname,
}

var mappingListExternalAsmInstancesSortByEnumLowerCase = map[string]ListExternalAsmInstancesSortByEnum{
	"timecreated": ListExternalAsmInstancesSortByTimecreated,
	"displayname": ListExternalAsmInstancesSortByDisplayname,
}

// GetListExternalAsmInstancesSortByEnumValues Enumerates the set of values for ListExternalAsmInstancesSortByEnum
func GetListExternalAsmInstancesSortByEnumValues() []ListExternalAsmInstancesSortByEnum {
	values := make([]ListExternalAsmInstancesSortByEnum, 0)
	for _, v := range mappingListExternalAsmInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmInstancesSortByEnumStringValues Enumerates the set of values in String for ListExternalAsmInstancesSortByEnum
func GetListExternalAsmInstancesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalAsmInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmInstancesSortByEnum(val string) (ListExternalAsmInstancesSortByEnum, bool) {
	enum, ok := mappingListExternalAsmInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalAsmInstancesSortOrderEnum Enum with underlying type: string
type ListExternalAsmInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalAsmInstancesSortOrderEnum
const (
	ListExternalAsmInstancesSortOrderAsc  ListExternalAsmInstancesSortOrderEnum = "ASC"
	ListExternalAsmInstancesSortOrderDesc ListExternalAsmInstancesSortOrderEnum = "DESC"
)

var mappingListExternalAsmInstancesSortOrderEnum = map[string]ListExternalAsmInstancesSortOrderEnum{
	"ASC":  ListExternalAsmInstancesSortOrderAsc,
	"DESC": ListExternalAsmInstancesSortOrderDesc,
}

var mappingListExternalAsmInstancesSortOrderEnumLowerCase = map[string]ListExternalAsmInstancesSortOrderEnum{
	"asc":  ListExternalAsmInstancesSortOrderAsc,
	"desc": ListExternalAsmInstancesSortOrderDesc,
}

// GetListExternalAsmInstancesSortOrderEnumValues Enumerates the set of values for ListExternalAsmInstancesSortOrderEnum
func GetListExternalAsmInstancesSortOrderEnumValues() []ListExternalAsmInstancesSortOrderEnum {
	values := make([]ListExternalAsmInstancesSortOrderEnum, 0)
	for _, v := range mappingListExternalAsmInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalAsmInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalAsmInstancesSortOrderEnum
func GetListExternalAsmInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalAsmInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalAsmInstancesSortOrderEnum(val string) (ListExternalAsmInstancesSortOrderEnum, bool) {
	enum, ok := mappingListExternalAsmInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
