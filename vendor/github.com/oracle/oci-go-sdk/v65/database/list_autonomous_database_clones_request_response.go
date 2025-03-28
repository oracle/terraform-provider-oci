// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDatabaseClonesRequest wrapper for the ListAutonomousDatabaseClones operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseClones.go.html to see an example of how to use ListAutonomousDatabaseClonesRequest.
type ListAutonomousDatabaseClonesRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDatabaseClonesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousDatabaseSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousDatabaseClonesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given clone type exactly.
	CloneType ListAutonomousDatabaseClonesCloneTypeEnum `mandatory:"false" contributesTo:"query" name:"cloneType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDatabaseClonesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDatabaseClonesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDatabaseClonesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDatabaseClonesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDatabaseClonesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutonomousDatabaseClonesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousDatabaseClonesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAutonomousDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDatabaseClonesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutonomousDatabaseClonesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDatabaseClonesCloneTypeEnum(string(request.CloneType)); !ok && request.CloneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneType: %s. Supported values are: %s.", request.CloneType, strings.Join(GetListAutonomousDatabaseClonesCloneTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDatabaseClonesResponse wrapper for the ListAutonomousDatabaseClones operation
type ListAutonomousDatabaseClonesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousDatabaseSummary instances
	Items []AutonomousDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDatabaseClonesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDatabaseClonesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDatabaseClonesSortOrderEnum Enum with underlying type: string
type ListAutonomousDatabaseClonesSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseClonesSortOrderEnum
const (
	ListAutonomousDatabaseClonesSortOrderAsc  ListAutonomousDatabaseClonesSortOrderEnum = "ASC"
	ListAutonomousDatabaseClonesSortOrderDesc ListAutonomousDatabaseClonesSortOrderEnum = "DESC"
)

var mappingListAutonomousDatabaseClonesSortOrderEnum = map[string]ListAutonomousDatabaseClonesSortOrderEnum{
	"ASC":  ListAutonomousDatabaseClonesSortOrderAsc,
	"DESC": ListAutonomousDatabaseClonesSortOrderDesc,
}

var mappingListAutonomousDatabaseClonesSortOrderEnumLowerCase = map[string]ListAutonomousDatabaseClonesSortOrderEnum{
	"asc":  ListAutonomousDatabaseClonesSortOrderAsc,
	"desc": ListAutonomousDatabaseClonesSortOrderDesc,
}

// GetListAutonomousDatabaseClonesSortOrderEnumValues Enumerates the set of values for ListAutonomousDatabaseClonesSortOrderEnum
func GetListAutonomousDatabaseClonesSortOrderEnumValues() []ListAutonomousDatabaseClonesSortOrderEnum {
	values := make([]ListAutonomousDatabaseClonesSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDatabaseClonesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseClonesSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseClonesSortOrderEnum
func GetListAutonomousDatabaseClonesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousDatabaseClonesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseClonesSortOrderEnum(val string) (ListAutonomousDatabaseClonesSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseClonesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousDatabaseClonesSortByEnum Enum with underlying type: string
type ListAutonomousDatabaseClonesSortByEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseClonesSortByEnum
const (
	ListAutonomousDatabaseClonesSortByNone        ListAutonomousDatabaseClonesSortByEnum = "NONE"
	ListAutonomousDatabaseClonesSortByTimecreated ListAutonomousDatabaseClonesSortByEnum = "TIMECREATED"
	ListAutonomousDatabaseClonesSortByDisplayname ListAutonomousDatabaseClonesSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousDatabaseClonesSortByEnum = map[string]ListAutonomousDatabaseClonesSortByEnum{
	"NONE":        ListAutonomousDatabaseClonesSortByNone,
	"TIMECREATED": ListAutonomousDatabaseClonesSortByTimecreated,
	"DISPLAYNAME": ListAutonomousDatabaseClonesSortByDisplayname,
}

var mappingListAutonomousDatabaseClonesSortByEnumLowerCase = map[string]ListAutonomousDatabaseClonesSortByEnum{
	"none":        ListAutonomousDatabaseClonesSortByNone,
	"timecreated": ListAutonomousDatabaseClonesSortByTimecreated,
	"displayname": ListAutonomousDatabaseClonesSortByDisplayname,
}

// GetListAutonomousDatabaseClonesSortByEnumValues Enumerates the set of values for ListAutonomousDatabaseClonesSortByEnum
func GetListAutonomousDatabaseClonesSortByEnumValues() []ListAutonomousDatabaseClonesSortByEnum {
	values := make([]ListAutonomousDatabaseClonesSortByEnum, 0)
	for _, v := range mappingListAutonomousDatabaseClonesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseClonesSortByEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseClonesSortByEnum
func GetListAutonomousDatabaseClonesSortByEnumStringValues() []string {
	return []string{
		"NONE",
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutonomousDatabaseClonesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseClonesSortByEnum(val string) (ListAutonomousDatabaseClonesSortByEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseClonesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousDatabaseClonesCloneTypeEnum Enum with underlying type: string
type ListAutonomousDatabaseClonesCloneTypeEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseClonesCloneTypeEnum
const (
	ListAutonomousDatabaseClonesCloneTypeRefreshableClone ListAutonomousDatabaseClonesCloneTypeEnum = "REFRESHABLE_CLONE"
)

var mappingListAutonomousDatabaseClonesCloneTypeEnum = map[string]ListAutonomousDatabaseClonesCloneTypeEnum{
	"REFRESHABLE_CLONE": ListAutonomousDatabaseClonesCloneTypeRefreshableClone,
}

var mappingListAutonomousDatabaseClonesCloneTypeEnumLowerCase = map[string]ListAutonomousDatabaseClonesCloneTypeEnum{
	"refreshable_clone": ListAutonomousDatabaseClonesCloneTypeRefreshableClone,
}

// GetListAutonomousDatabaseClonesCloneTypeEnumValues Enumerates the set of values for ListAutonomousDatabaseClonesCloneTypeEnum
func GetListAutonomousDatabaseClonesCloneTypeEnumValues() []ListAutonomousDatabaseClonesCloneTypeEnum {
	values := make([]ListAutonomousDatabaseClonesCloneTypeEnum, 0)
	for _, v := range mappingListAutonomousDatabaseClonesCloneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseClonesCloneTypeEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseClonesCloneTypeEnum
func GetListAutonomousDatabaseClonesCloneTypeEnumStringValues() []string {
	return []string{
		"REFRESHABLE_CLONE",
	}
}

// GetMappingListAutonomousDatabaseClonesCloneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseClonesCloneTypeEnum(val string) (ListAutonomousDatabaseClonesCloneTypeEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseClonesCloneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
