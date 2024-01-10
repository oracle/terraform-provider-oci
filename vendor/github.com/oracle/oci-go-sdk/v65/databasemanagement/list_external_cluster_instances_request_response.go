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

// ListExternalClusterInstancesRequest wrapper for the ListExternalClusterInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalClusterInstances.go.html to see an example of how to use ListExternalClusterInstancesRequest.
type ListExternalClusterInstancesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external cluster.
	ExternalClusterId *string `mandatory:"false" contributesTo:"query" name:"externalClusterId"`

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
	SortBy ListExternalClusterInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalClusterInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalClusterInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalClusterInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalClusterInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalClusterInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalClusterInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalClusterInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalClusterInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalClusterInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalClusterInstancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalClusterInstancesResponse wrapper for the ListExternalClusterInstances operation
type ListExternalClusterInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalClusterInstanceCollection instances
	ExternalClusterInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalClusterInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalClusterInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalClusterInstancesSortByEnum Enum with underlying type: string
type ListExternalClusterInstancesSortByEnum string

// Set of constants representing the allowable values for ListExternalClusterInstancesSortByEnum
const (
	ListExternalClusterInstancesSortByTimecreated ListExternalClusterInstancesSortByEnum = "TIMECREATED"
	ListExternalClusterInstancesSortByDisplayname ListExternalClusterInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListExternalClusterInstancesSortByEnum = map[string]ListExternalClusterInstancesSortByEnum{
	"TIMECREATED": ListExternalClusterInstancesSortByTimecreated,
	"DISPLAYNAME": ListExternalClusterInstancesSortByDisplayname,
}

var mappingListExternalClusterInstancesSortByEnumLowerCase = map[string]ListExternalClusterInstancesSortByEnum{
	"timecreated": ListExternalClusterInstancesSortByTimecreated,
	"displayname": ListExternalClusterInstancesSortByDisplayname,
}

// GetListExternalClusterInstancesSortByEnumValues Enumerates the set of values for ListExternalClusterInstancesSortByEnum
func GetListExternalClusterInstancesSortByEnumValues() []ListExternalClusterInstancesSortByEnum {
	values := make([]ListExternalClusterInstancesSortByEnum, 0)
	for _, v := range mappingListExternalClusterInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalClusterInstancesSortByEnumStringValues Enumerates the set of values in String for ListExternalClusterInstancesSortByEnum
func GetListExternalClusterInstancesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalClusterInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalClusterInstancesSortByEnum(val string) (ListExternalClusterInstancesSortByEnum, bool) {
	enum, ok := mappingListExternalClusterInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalClusterInstancesSortOrderEnum Enum with underlying type: string
type ListExternalClusterInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalClusterInstancesSortOrderEnum
const (
	ListExternalClusterInstancesSortOrderAsc  ListExternalClusterInstancesSortOrderEnum = "ASC"
	ListExternalClusterInstancesSortOrderDesc ListExternalClusterInstancesSortOrderEnum = "DESC"
)

var mappingListExternalClusterInstancesSortOrderEnum = map[string]ListExternalClusterInstancesSortOrderEnum{
	"ASC":  ListExternalClusterInstancesSortOrderAsc,
	"DESC": ListExternalClusterInstancesSortOrderDesc,
}

var mappingListExternalClusterInstancesSortOrderEnumLowerCase = map[string]ListExternalClusterInstancesSortOrderEnum{
	"asc":  ListExternalClusterInstancesSortOrderAsc,
	"desc": ListExternalClusterInstancesSortOrderDesc,
}

// GetListExternalClusterInstancesSortOrderEnumValues Enumerates the set of values for ListExternalClusterInstancesSortOrderEnum
func GetListExternalClusterInstancesSortOrderEnumValues() []ListExternalClusterInstancesSortOrderEnum {
	values := make([]ListExternalClusterInstancesSortOrderEnum, 0)
	for _, v := range mappingListExternalClusterInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalClusterInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalClusterInstancesSortOrderEnum
func GetListExternalClusterInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalClusterInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalClusterInstancesSortOrderEnum(val string) (ListExternalClusterInstancesSortOrderEnum, bool) {
	enum, ok := mappingListExternalClusterInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
