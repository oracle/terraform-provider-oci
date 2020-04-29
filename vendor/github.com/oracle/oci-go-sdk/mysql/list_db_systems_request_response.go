// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDbSystemsRequest wrapper for the ListDbSystems operation
type ListDbSystemsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// DbSystem Lifecycle State
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The requested Configuration instance.
	ConfigurationId *string `mandatory:"false" contributesTo:"query" name:"configurationId"`

	// Filter instances if they are using the latest revision of the
	// Configuration they are associated with.
	IsUpToDate *bool `mandatory:"false" contributesTo:"query" name:"isUpToDate"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListDbSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListDbSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbSystemsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDbSystemsResponse wrapper for the ListDbSystems operation
type ListDbSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbSystemSummary instances
	Items []DbSystemSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbSystemsSortByEnum Enum with underlying type: string
type ListDbSystemsSortByEnum string

// Set of constants representing the allowable values for ListDbSystemsSortByEnum
const (
	ListDbSystemsSortByDisplayname ListDbSystemsSortByEnum = "displayName"
	ListDbSystemsSortByTimecreated ListDbSystemsSortByEnum = "timeCreated"
)

var mappingListDbSystemsSortBy = map[string]ListDbSystemsSortByEnum{
	"displayName": ListDbSystemsSortByDisplayname,
	"timeCreated": ListDbSystemsSortByTimecreated,
}

// GetListDbSystemsSortByEnumValues Enumerates the set of values for ListDbSystemsSortByEnum
func GetListDbSystemsSortByEnumValues() []ListDbSystemsSortByEnum {
	values := make([]ListDbSystemsSortByEnum, 0)
	for _, v := range mappingListDbSystemsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDbSystemsSortOrderEnum Enum with underlying type: string
type ListDbSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListDbSystemsSortOrderEnum
const (
	ListDbSystemsSortOrderAsc  ListDbSystemsSortOrderEnum = "ASC"
	ListDbSystemsSortOrderDesc ListDbSystemsSortOrderEnum = "DESC"
)

var mappingListDbSystemsSortOrder = map[string]ListDbSystemsSortOrderEnum{
	"ASC":  ListDbSystemsSortOrderAsc,
	"DESC": ListDbSystemsSortOrderDesc,
}

// GetListDbSystemsSortOrderEnumValues Enumerates the set of values for ListDbSystemsSortOrderEnum
func GetListDbSystemsSortOrderEnumValues() []ListDbSystemsSortOrderEnum {
	values := make([]ListDbSystemsSortOrderEnum, 0)
	for _, v := range mappingListDbSystemsSortOrder {
		values = append(values, v)
	}
	return values
}
