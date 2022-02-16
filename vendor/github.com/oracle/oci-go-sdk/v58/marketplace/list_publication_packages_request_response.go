// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPublicationPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPublicationPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPublicationPackagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPublicationPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPublicationPackagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListPublicationPackagesSortByEnum = map[string]ListPublicationPackagesSortByEnum{
	"TIMERELEASED": ListPublicationPackagesSortByTimereleased,
}

// GetListPublicationPackagesSortByEnumValues Enumerates the set of values for ListPublicationPackagesSortByEnum
func GetListPublicationPackagesSortByEnumValues() []ListPublicationPackagesSortByEnum {
	values := make([]ListPublicationPackagesSortByEnum, 0)
	for _, v := range mappingListPublicationPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicationPackagesSortByEnumStringValues Enumerates the set of values in String for ListPublicationPackagesSortByEnum
func GetListPublicationPackagesSortByEnumStringValues() []string {
	return []string{
		"TIMERELEASED",
	}
}

// GetMappingListPublicationPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicationPackagesSortByEnum(val string) (ListPublicationPackagesSortByEnum, bool) {
	mappingListPublicationPackagesSortByEnumIgnoreCase := make(map[string]ListPublicationPackagesSortByEnum)
	for k, v := range mappingListPublicationPackagesSortByEnum {
		mappingListPublicationPackagesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublicationPackagesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListPublicationPackagesSortOrderEnum Enum with underlying type: string
type ListPublicationPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListPublicationPackagesSortOrderEnum
const (
	ListPublicationPackagesSortOrderAsc  ListPublicationPackagesSortOrderEnum = "ASC"
	ListPublicationPackagesSortOrderDesc ListPublicationPackagesSortOrderEnum = "DESC"
)

var mappingListPublicationPackagesSortOrderEnum = map[string]ListPublicationPackagesSortOrderEnum{
	"ASC":  ListPublicationPackagesSortOrderAsc,
	"DESC": ListPublicationPackagesSortOrderDesc,
}

// GetListPublicationPackagesSortOrderEnumValues Enumerates the set of values for ListPublicationPackagesSortOrderEnum
func GetListPublicationPackagesSortOrderEnumValues() []ListPublicationPackagesSortOrderEnum {
	values := make([]ListPublicationPackagesSortOrderEnum, 0)
	for _, v := range mappingListPublicationPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicationPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListPublicationPackagesSortOrderEnum
func GetListPublicationPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPublicationPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicationPackagesSortOrderEnum(val string) (ListPublicationPackagesSortOrderEnum, bool) {
	mappingListPublicationPackagesSortOrderEnumIgnoreCase := make(map[string]ListPublicationPackagesSortOrderEnum)
	for k, v := range mappingListPublicationPackagesSortOrderEnum {
		mappingListPublicationPackagesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublicationPackagesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
