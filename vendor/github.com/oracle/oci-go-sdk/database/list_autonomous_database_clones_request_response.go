// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAutonomousDatabaseClonesRequest wrapper for the ListAutonomousDatabaseClones operation
type ListAutonomousDatabaseClonesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
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
func (request ListAutonomousDatabaseClonesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDatabaseClonesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
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
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
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

var mappingListAutonomousDatabaseClonesSortOrder = map[string]ListAutonomousDatabaseClonesSortOrderEnum{
	"ASC":  ListAutonomousDatabaseClonesSortOrderAsc,
	"DESC": ListAutonomousDatabaseClonesSortOrderDesc,
}

// GetListAutonomousDatabaseClonesSortOrderEnumValues Enumerates the set of values for ListAutonomousDatabaseClonesSortOrderEnum
func GetListAutonomousDatabaseClonesSortOrderEnumValues() []ListAutonomousDatabaseClonesSortOrderEnum {
	values := make([]ListAutonomousDatabaseClonesSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDatabaseClonesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAutonomousDatabaseClonesSortByEnum Enum with underlying type: string
type ListAutonomousDatabaseClonesSortByEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseClonesSortByEnum
const (
	ListAutonomousDatabaseClonesSortByNone        ListAutonomousDatabaseClonesSortByEnum = "NONE"
	ListAutonomousDatabaseClonesSortByTimecreated ListAutonomousDatabaseClonesSortByEnum = "TIMECREATED"
	ListAutonomousDatabaseClonesSortByDisplayname ListAutonomousDatabaseClonesSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousDatabaseClonesSortBy = map[string]ListAutonomousDatabaseClonesSortByEnum{
	"NONE":        ListAutonomousDatabaseClonesSortByNone,
	"TIMECREATED": ListAutonomousDatabaseClonesSortByTimecreated,
	"DISPLAYNAME": ListAutonomousDatabaseClonesSortByDisplayname,
}

// GetListAutonomousDatabaseClonesSortByEnumValues Enumerates the set of values for ListAutonomousDatabaseClonesSortByEnum
func GetListAutonomousDatabaseClonesSortByEnumValues() []ListAutonomousDatabaseClonesSortByEnum {
	values := make([]ListAutonomousDatabaseClonesSortByEnum, 0)
	for _, v := range mappingListAutonomousDatabaseClonesSortBy {
		values = append(values, v)
	}
	return values
}

// ListAutonomousDatabaseClonesCloneTypeEnum Enum with underlying type: string
type ListAutonomousDatabaseClonesCloneTypeEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseClonesCloneTypeEnum
const (
	ListAutonomousDatabaseClonesCloneTypeRefreshableClone ListAutonomousDatabaseClonesCloneTypeEnum = "REFRESHABLE_CLONE"
)

var mappingListAutonomousDatabaseClonesCloneType = map[string]ListAutonomousDatabaseClonesCloneTypeEnum{
	"REFRESHABLE_CLONE": ListAutonomousDatabaseClonesCloneTypeRefreshableClone,
}

// GetListAutonomousDatabaseClonesCloneTypeEnumValues Enumerates the set of values for ListAutonomousDatabaseClonesCloneTypeEnum
func GetListAutonomousDatabaseClonesCloneTypeEnumValues() []ListAutonomousDatabaseClonesCloneTypeEnum {
	values := make([]ListAutonomousDatabaseClonesCloneTypeEnum, 0)
	for _, v := range mappingListAutonomousDatabaseClonesCloneType {
		values = append(values, v)
	}
	return values
}
