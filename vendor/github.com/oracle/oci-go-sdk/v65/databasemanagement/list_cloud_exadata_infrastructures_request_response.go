// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudExadataInfrastructuresRequest wrapper for the ListCloudExadataInfrastructures operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudExadataInfrastructures.go.html to see an example of how to use ListCloudExadataInfrastructuresRequest.
type ListCloudExadataInfrastructuresRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortBy ListCloudExadataInfrastructuresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudExadataInfrastructuresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudExadataInfrastructuresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudExadataInfrastructuresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudExadataInfrastructuresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudExadataInfrastructuresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudExadataInfrastructuresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudExadataInfrastructuresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudExadataInfrastructuresSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudExadataInfrastructuresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudExadataInfrastructuresSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudExadataInfrastructuresResponse wrapper for the ListCloudExadataInfrastructures operation
type ListCloudExadataInfrastructuresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudExadataInfrastructureCollection instances
	CloudExadataInfrastructureCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudExadataInfrastructuresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudExadataInfrastructuresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudExadataInfrastructuresSortByEnum Enum with underlying type: string
type ListCloudExadataInfrastructuresSortByEnum string

// Set of constants representing the allowable values for ListCloudExadataInfrastructuresSortByEnum
const (
	ListCloudExadataInfrastructuresSortByTimecreated ListCloudExadataInfrastructuresSortByEnum = "TIMECREATED"
	ListCloudExadataInfrastructuresSortByName        ListCloudExadataInfrastructuresSortByEnum = "NAME"
)

var mappingListCloudExadataInfrastructuresSortByEnum = map[string]ListCloudExadataInfrastructuresSortByEnum{
	"TIMECREATED": ListCloudExadataInfrastructuresSortByTimecreated,
	"NAME":        ListCloudExadataInfrastructuresSortByName,
}

var mappingListCloudExadataInfrastructuresSortByEnumLowerCase = map[string]ListCloudExadataInfrastructuresSortByEnum{
	"timecreated": ListCloudExadataInfrastructuresSortByTimecreated,
	"name":        ListCloudExadataInfrastructuresSortByName,
}

// GetListCloudExadataInfrastructuresSortByEnumValues Enumerates the set of values for ListCloudExadataInfrastructuresSortByEnum
func GetListCloudExadataInfrastructuresSortByEnumValues() []ListCloudExadataInfrastructuresSortByEnum {
	values := make([]ListCloudExadataInfrastructuresSortByEnum, 0)
	for _, v := range mappingListCloudExadataInfrastructuresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudExadataInfrastructuresSortByEnumStringValues Enumerates the set of values in String for ListCloudExadataInfrastructuresSortByEnum
func GetListCloudExadataInfrastructuresSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListCloudExadataInfrastructuresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudExadataInfrastructuresSortByEnum(val string) (ListCloudExadataInfrastructuresSortByEnum, bool) {
	enum, ok := mappingListCloudExadataInfrastructuresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudExadataInfrastructuresSortOrderEnum Enum with underlying type: string
type ListCloudExadataInfrastructuresSortOrderEnum string

// Set of constants representing the allowable values for ListCloudExadataInfrastructuresSortOrderEnum
const (
	ListCloudExadataInfrastructuresSortOrderAsc  ListCloudExadataInfrastructuresSortOrderEnum = "ASC"
	ListCloudExadataInfrastructuresSortOrderDesc ListCloudExadataInfrastructuresSortOrderEnum = "DESC"
)

var mappingListCloudExadataInfrastructuresSortOrderEnum = map[string]ListCloudExadataInfrastructuresSortOrderEnum{
	"ASC":  ListCloudExadataInfrastructuresSortOrderAsc,
	"DESC": ListCloudExadataInfrastructuresSortOrderDesc,
}

var mappingListCloudExadataInfrastructuresSortOrderEnumLowerCase = map[string]ListCloudExadataInfrastructuresSortOrderEnum{
	"asc":  ListCloudExadataInfrastructuresSortOrderAsc,
	"desc": ListCloudExadataInfrastructuresSortOrderDesc,
}

// GetListCloudExadataInfrastructuresSortOrderEnumValues Enumerates the set of values for ListCloudExadataInfrastructuresSortOrderEnum
func GetListCloudExadataInfrastructuresSortOrderEnumValues() []ListCloudExadataInfrastructuresSortOrderEnum {
	values := make([]ListCloudExadataInfrastructuresSortOrderEnum, 0)
	for _, v := range mappingListCloudExadataInfrastructuresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudExadataInfrastructuresSortOrderEnumStringValues Enumerates the set of values in String for ListCloudExadataInfrastructuresSortOrderEnum
func GetListCloudExadataInfrastructuresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudExadataInfrastructuresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudExadataInfrastructuresSortOrderEnum(val string) (ListCloudExadataInfrastructuresSortOrderEnum, bool) {
	enum, ok := mappingListCloudExadataInfrastructuresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
