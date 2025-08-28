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

// ListCloudClusterInstancesRequest wrapper for the ListCloudClusterInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudClusterInstances.go.html to see an example of how to use ListCloudClusterInstancesRequest.
type ListCloudClusterInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster.
	CloudClusterId *string `mandatory:"false" contributesTo:"query" name:"cloudClusterId"`

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
	SortBy ListCloudClusterInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudClusterInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudClusterInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudClusterInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudClusterInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudClusterInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudClusterInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudClusterInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudClusterInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudClusterInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudClusterInstancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudClusterInstancesResponse wrapper for the ListCloudClusterInstances operation
type ListCloudClusterInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudClusterInstanceCollection instances
	CloudClusterInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudClusterInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudClusterInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudClusterInstancesSortByEnum Enum with underlying type: string
type ListCloudClusterInstancesSortByEnum string

// Set of constants representing the allowable values for ListCloudClusterInstancesSortByEnum
const (
	ListCloudClusterInstancesSortByTimecreated ListCloudClusterInstancesSortByEnum = "TIMECREATED"
	ListCloudClusterInstancesSortByDisplayname ListCloudClusterInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListCloudClusterInstancesSortByEnum = map[string]ListCloudClusterInstancesSortByEnum{
	"TIMECREATED": ListCloudClusterInstancesSortByTimecreated,
	"DISPLAYNAME": ListCloudClusterInstancesSortByDisplayname,
}

var mappingListCloudClusterInstancesSortByEnumLowerCase = map[string]ListCloudClusterInstancesSortByEnum{
	"timecreated": ListCloudClusterInstancesSortByTimecreated,
	"displayname": ListCloudClusterInstancesSortByDisplayname,
}

// GetListCloudClusterInstancesSortByEnumValues Enumerates the set of values for ListCloudClusterInstancesSortByEnum
func GetListCloudClusterInstancesSortByEnumValues() []ListCloudClusterInstancesSortByEnum {
	values := make([]ListCloudClusterInstancesSortByEnum, 0)
	for _, v := range mappingListCloudClusterInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudClusterInstancesSortByEnumStringValues Enumerates the set of values in String for ListCloudClusterInstancesSortByEnum
func GetListCloudClusterInstancesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudClusterInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudClusterInstancesSortByEnum(val string) (ListCloudClusterInstancesSortByEnum, bool) {
	enum, ok := mappingListCloudClusterInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudClusterInstancesSortOrderEnum Enum with underlying type: string
type ListCloudClusterInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudClusterInstancesSortOrderEnum
const (
	ListCloudClusterInstancesSortOrderAsc  ListCloudClusterInstancesSortOrderEnum = "ASC"
	ListCloudClusterInstancesSortOrderDesc ListCloudClusterInstancesSortOrderEnum = "DESC"
)

var mappingListCloudClusterInstancesSortOrderEnum = map[string]ListCloudClusterInstancesSortOrderEnum{
	"ASC":  ListCloudClusterInstancesSortOrderAsc,
	"DESC": ListCloudClusterInstancesSortOrderDesc,
}

var mappingListCloudClusterInstancesSortOrderEnumLowerCase = map[string]ListCloudClusterInstancesSortOrderEnum{
	"asc":  ListCloudClusterInstancesSortOrderAsc,
	"desc": ListCloudClusterInstancesSortOrderDesc,
}

// GetListCloudClusterInstancesSortOrderEnumValues Enumerates the set of values for ListCloudClusterInstancesSortOrderEnum
func GetListCloudClusterInstancesSortOrderEnumValues() []ListCloudClusterInstancesSortOrderEnum {
	values := make([]ListCloudClusterInstancesSortOrderEnum, 0)
	for _, v := range mappingListCloudClusterInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudClusterInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudClusterInstancesSortOrderEnum
func GetListCloudClusterInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudClusterInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudClusterInstancesSortOrderEnum(val string) (ListCloudClusterInstancesSortOrderEnum, bool) {
	enum, ok := mappingListCloudClusterInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
