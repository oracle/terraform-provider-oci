// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSoftwareSourcePackagesRequest wrapper for the ListSoftwareSourcePackages operation
type ListSoftwareSourcePackagesRequest struct {

	// The OCID of the software source.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSoftwareSourcePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListSoftwareSourcePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourcePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourcePackagesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourcePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSoftwareSourcePackagesResponse wrapper for the ListSoftwareSourcePackages operation
type ListSoftwareSourcePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SoftwarePackageSummary instances
	Items []SoftwarePackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `SoftwarePackages`s. If this header
	// appears in the response, then this is a partial list of software
	// sources. Include this value as the `page` parameter in a
	// subsequent GET request to get the next batch of software
	// sources.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareSourcePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourcePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourcePackagesSortOrderEnum Enum with underlying type: string
type ListSoftwareSourcePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourcePackagesSortOrderEnum
const (
	ListSoftwareSourcePackagesSortOrderAsc  ListSoftwareSourcePackagesSortOrderEnum = "ASC"
	ListSoftwareSourcePackagesSortOrderDesc ListSoftwareSourcePackagesSortOrderEnum = "DESC"
)

var mappingListSoftwareSourcePackagesSortOrder = map[string]ListSoftwareSourcePackagesSortOrderEnum{
	"ASC":  ListSoftwareSourcePackagesSortOrderAsc,
	"DESC": ListSoftwareSourcePackagesSortOrderDesc,
}

// GetListSoftwareSourcePackagesSortOrderEnumValues Enumerates the set of values for ListSoftwareSourcePackagesSortOrderEnum
func GetListSoftwareSourcePackagesSortOrderEnumValues() []ListSoftwareSourcePackagesSortOrderEnum {
	values := make([]ListSoftwareSourcePackagesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourcePackagesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSoftwareSourcePackagesSortByEnum Enum with underlying type: string
type ListSoftwareSourcePackagesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourcePackagesSortByEnum
const (
	ListSoftwareSourcePackagesSortByTimecreated ListSoftwareSourcePackagesSortByEnum = "TIMECREATED"
	ListSoftwareSourcePackagesSortByDisplayname ListSoftwareSourcePackagesSortByEnum = "DISPLAYNAME"
)

var mappingListSoftwareSourcePackagesSortBy = map[string]ListSoftwareSourcePackagesSortByEnum{
	"TIMECREATED": ListSoftwareSourcePackagesSortByTimecreated,
	"DISPLAYNAME": ListSoftwareSourcePackagesSortByDisplayname,
}

// GetListSoftwareSourcePackagesSortByEnumValues Enumerates the set of values for ListSoftwareSourcePackagesSortByEnum
func GetListSoftwareSourcePackagesSortByEnumValues() []ListSoftwareSourcePackagesSortByEnum {
	values := make([]ListSoftwareSourcePackagesSortByEnum, 0)
	for _, v := range mappingListSoftwareSourcePackagesSortBy {
		values = append(values, v)
	}
	return values
}
