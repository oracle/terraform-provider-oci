// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListPublicationPackagesRequest wrapper for the ListPublicationPackages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListPublicationPackages.go.html to see an example of how to use ListPublicationPackagesRequest.
type ListPublicationPackagesRequest struct {

	// The unique identifier for the publication.
	PublicationId *string `mandatory:"true" contributesTo:"path" name:"publicationId"`

	// The version of the package. Package versions are unique within a listing.
	PackageVersion *string `mandatory:"false" contributesTo:"query" name:"packageVersion"`

	// A filter to return only packages that match the given package type exactly.
	PackageType *string `mandatory:"false" contributesTo:"query" name:"packageType"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `TIMERELEASED` displays results in descending order by default.
	// You can change your preference by specifying a different sort order.
	SortBy ListPublicationPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either `ASC` or `DESC`.
	SortOrder ListPublicationPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublicationPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublicationPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublicationPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublicationPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPublicationPackagesResponse wrapper for the ListPublicationPackages operation
type ListPublicationPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PublicationPackageSummary instances
	Items []PublicationPackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPublicationPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublicationPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublicationPackagesSortByEnum Enum with underlying type: string
type ListPublicationPackagesSortByEnum string

// Set of constants representing the allowable values for ListPublicationPackagesSortByEnum
const (
	ListPublicationPackagesSortByTimereleased ListPublicationPackagesSortByEnum = "TIMERELEASED"
)

var mappingListPublicationPackagesSortBy = map[string]ListPublicationPackagesSortByEnum{
	"TIMERELEASED": ListPublicationPackagesSortByTimereleased,
}

// GetListPublicationPackagesSortByEnumValues Enumerates the set of values for ListPublicationPackagesSortByEnum
func GetListPublicationPackagesSortByEnumValues() []ListPublicationPackagesSortByEnum {
	values := make([]ListPublicationPackagesSortByEnum, 0)
	for _, v := range mappingListPublicationPackagesSortBy {
		values = append(values, v)
	}
	return values
}

// ListPublicationPackagesSortOrderEnum Enum with underlying type: string
type ListPublicationPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListPublicationPackagesSortOrderEnum
const (
	ListPublicationPackagesSortOrderAsc  ListPublicationPackagesSortOrderEnum = "ASC"
	ListPublicationPackagesSortOrderDesc ListPublicationPackagesSortOrderEnum = "DESC"
)

var mappingListPublicationPackagesSortOrder = map[string]ListPublicationPackagesSortOrderEnum{
	"ASC":  ListPublicationPackagesSortOrderAsc,
	"DESC": ListPublicationPackagesSortOrderDesc,
}

// GetListPublicationPackagesSortOrderEnumValues Enumerates the set of values for ListPublicationPackagesSortOrderEnum
func GetListPublicationPackagesSortOrderEnumValues() []ListPublicationPackagesSortOrderEnum {
	values := make([]ListPublicationPackagesSortOrderEnum, 0)
	for _, v := range mappingListPublicationPackagesSortOrder {
		values = append(values, v)
	}
	return values
}
