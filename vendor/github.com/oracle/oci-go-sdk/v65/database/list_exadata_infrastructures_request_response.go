// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExadataInfrastructuresRequest wrapper for the ListExadataInfrastructures operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExadataInfrastructures.go.html to see an example of how to use ListExadataInfrastructuresRequest.
type ListExadataInfrastructuresRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListExadataInfrastructuresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExadataInfrastructuresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ExadataInfrastructureSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExadataInfrastructuresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadataInfrastructuresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExadataInfrastructuresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadataInfrastructuresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExadataInfrastructuresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExadataInfrastructuresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExadataInfrastructuresSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExadataInfrastructuresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExadataInfrastructuresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetExadataInfrastructureSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExadataInfrastructuresResponse wrapper for the ListExadataInfrastructures operation
type ListExadataInfrastructuresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExadataInfrastructureSummary instances
	Items []ExadataInfrastructureSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExadataInfrastructuresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadataInfrastructuresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadataInfrastructuresSortByEnum Enum with underlying type: string
type ListExadataInfrastructuresSortByEnum string

// Set of constants representing the allowable values for ListExadataInfrastructuresSortByEnum
const (
	ListExadataInfrastructuresSortByTimecreated ListExadataInfrastructuresSortByEnum = "TIMECREATED"
	ListExadataInfrastructuresSortByDisplayname ListExadataInfrastructuresSortByEnum = "DISPLAYNAME"
)

var mappingListExadataInfrastructuresSortByEnum = map[string]ListExadataInfrastructuresSortByEnum{
	"TIMECREATED": ListExadataInfrastructuresSortByTimecreated,
	"DISPLAYNAME": ListExadataInfrastructuresSortByDisplayname,
}

var mappingListExadataInfrastructuresSortByEnumLowerCase = map[string]ListExadataInfrastructuresSortByEnum{
	"timecreated": ListExadataInfrastructuresSortByTimecreated,
	"displayname": ListExadataInfrastructuresSortByDisplayname,
}

// GetListExadataInfrastructuresSortByEnumValues Enumerates the set of values for ListExadataInfrastructuresSortByEnum
func GetListExadataInfrastructuresSortByEnumValues() []ListExadataInfrastructuresSortByEnum {
	values := make([]ListExadataInfrastructuresSortByEnum, 0)
	for _, v := range mappingListExadataInfrastructuresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExadataInfrastructuresSortByEnumStringValues Enumerates the set of values in String for ListExadataInfrastructuresSortByEnum
func GetListExadataInfrastructuresSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExadataInfrastructuresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExadataInfrastructuresSortByEnum(val string) (ListExadataInfrastructuresSortByEnum, bool) {
	enum, ok := mappingListExadataInfrastructuresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExadataInfrastructuresSortOrderEnum Enum with underlying type: string
type ListExadataInfrastructuresSortOrderEnum string

// Set of constants representing the allowable values for ListExadataInfrastructuresSortOrderEnum
const (
	ListExadataInfrastructuresSortOrderAsc  ListExadataInfrastructuresSortOrderEnum = "ASC"
	ListExadataInfrastructuresSortOrderDesc ListExadataInfrastructuresSortOrderEnum = "DESC"
)

var mappingListExadataInfrastructuresSortOrderEnum = map[string]ListExadataInfrastructuresSortOrderEnum{
	"ASC":  ListExadataInfrastructuresSortOrderAsc,
	"DESC": ListExadataInfrastructuresSortOrderDesc,
}

var mappingListExadataInfrastructuresSortOrderEnumLowerCase = map[string]ListExadataInfrastructuresSortOrderEnum{
	"asc":  ListExadataInfrastructuresSortOrderAsc,
	"desc": ListExadataInfrastructuresSortOrderDesc,
}

// GetListExadataInfrastructuresSortOrderEnumValues Enumerates the set of values for ListExadataInfrastructuresSortOrderEnum
func GetListExadataInfrastructuresSortOrderEnumValues() []ListExadataInfrastructuresSortOrderEnum {
	values := make([]ListExadataInfrastructuresSortOrderEnum, 0)
	for _, v := range mappingListExadataInfrastructuresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExadataInfrastructuresSortOrderEnumStringValues Enumerates the set of values in String for ListExadataInfrastructuresSortOrderEnum
func GetListExadataInfrastructuresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExadataInfrastructuresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExadataInfrastructuresSortOrderEnum(val string) (ListExadataInfrastructuresSortOrderEnum, bool) {
	enum, ok := mappingListExadataInfrastructuresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
