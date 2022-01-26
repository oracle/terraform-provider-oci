// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SearchSoftwarePackagesRequest wrapper for the SearchSoftwarePackages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/SearchSoftwarePackages.go.html to see an example of how to use SearchSoftwarePackagesRequest.
type SearchSoftwarePackagesRequest struct {

	// the identifier for the software package (not an OCID)
	SoftwarePackageName *string `mandatory:"false" contributesTo:"query" name:"softwarePackageName"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The name of the CVE as published.
	// Example: `CVE-2006-4535`
	CveName *string `mandatory:"false" contributesTo:"query" name:"cveName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SearchSoftwarePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy SearchSoftwarePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SearchSoftwarePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SearchSoftwarePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SearchSoftwarePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SearchSoftwarePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SearchSoftwarePackagesResponse wrapper for the SearchSoftwarePackages operation
type SearchSoftwarePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SoftwarePackageSearchSummary instances
	Items []SoftwarePackageSearchSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SearchSoftwarePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SearchSoftwarePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SearchSoftwarePackagesSortOrderEnum Enum with underlying type: string
type SearchSoftwarePackagesSortOrderEnum string

// Set of constants representing the allowable values for SearchSoftwarePackagesSortOrderEnum
const (
	SearchSoftwarePackagesSortOrderAsc  SearchSoftwarePackagesSortOrderEnum = "ASC"
	SearchSoftwarePackagesSortOrderDesc SearchSoftwarePackagesSortOrderEnum = "DESC"
)

var mappingSearchSoftwarePackagesSortOrder = map[string]SearchSoftwarePackagesSortOrderEnum{
	"ASC":  SearchSoftwarePackagesSortOrderAsc,
	"DESC": SearchSoftwarePackagesSortOrderDesc,
}

// GetSearchSoftwarePackagesSortOrderEnumValues Enumerates the set of values for SearchSoftwarePackagesSortOrderEnum
func GetSearchSoftwarePackagesSortOrderEnumValues() []SearchSoftwarePackagesSortOrderEnum {
	values := make([]SearchSoftwarePackagesSortOrderEnum, 0)
	for _, v := range mappingSearchSoftwarePackagesSortOrder {
		values = append(values, v)
	}
	return values
}

// SearchSoftwarePackagesSortByEnum Enum with underlying type: string
type SearchSoftwarePackagesSortByEnum string

// Set of constants representing the allowable values for SearchSoftwarePackagesSortByEnum
const (
	SearchSoftwarePackagesSortByTimecreated SearchSoftwarePackagesSortByEnum = "TIMECREATED"
	SearchSoftwarePackagesSortByDisplayname SearchSoftwarePackagesSortByEnum = "DISPLAYNAME"
)

var mappingSearchSoftwarePackagesSortBy = map[string]SearchSoftwarePackagesSortByEnum{
	"TIMECREATED": SearchSoftwarePackagesSortByTimecreated,
	"DISPLAYNAME": SearchSoftwarePackagesSortByDisplayname,
}

// GetSearchSoftwarePackagesSortByEnumValues Enumerates the set of values for SearchSoftwarePackagesSortByEnum
func GetSearchSoftwarePackagesSortByEnumValues() []SearchSoftwarePackagesSortByEnum {
	values := make([]SearchSoftwarePackagesSortByEnum, 0)
	for _, v := range mappingSearchSoftwarePackagesSortBy {
		values = append(values, v)
	}
	return values
}
