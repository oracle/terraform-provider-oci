// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osubusage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListComputedUsagesRequest wrapper for the ListComputedUsages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubusage/ListComputedUsages.go.html to see an example of how to use ListComputedUsagesRequest.
type ListComputedUsagesRequest struct {

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// Initial date to filter Computed Usage data in SPM. In the case of non aggregated data the time period between of fromDate and toDate , expressed in RFC 3339 timestamp format.
	TimeFrom *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeFrom"`

	// Final date to filter Computed Usage data in SPM, expressed in RFC 3339 timestamp format.
	TimeTo *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeTo"`

	// Product part number for subscribed service line, called parent product.
	ParentProduct *string `mandatory:"false" contributesTo:"query" name:"parentProduct"`

	// Product part number for Computed Usage .
	ComputedProduct *string `mandatory:"false" contributesTo:"query" name:"computedProduct"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListComputedUsagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	SortBy ListComputedUsagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComputedUsagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComputedUsagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListComputedUsagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComputedUsagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListComputedUsagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListComputedUsagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListComputedUsagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListComputedUsagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListComputedUsagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListComputedUsagesResponse wrapper for the ListComputedUsages operation
type ListComputedUsagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ComputedUsageSummary instances
	Items []ComputedUsageSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListComputedUsagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComputedUsagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComputedUsagesSortOrderEnum Enum with underlying type: string
type ListComputedUsagesSortOrderEnum string

// Set of constants representing the allowable values for ListComputedUsagesSortOrderEnum
const (
	ListComputedUsagesSortOrderAsc  ListComputedUsagesSortOrderEnum = "ASC"
	ListComputedUsagesSortOrderDesc ListComputedUsagesSortOrderEnum = "DESC"
)

var mappingListComputedUsagesSortOrderEnum = map[string]ListComputedUsagesSortOrderEnum{
	"ASC":  ListComputedUsagesSortOrderAsc,
	"DESC": ListComputedUsagesSortOrderDesc,
}

var mappingListComputedUsagesSortOrderEnumLowerCase = map[string]ListComputedUsagesSortOrderEnum{
	"asc":  ListComputedUsagesSortOrderAsc,
	"desc": ListComputedUsagesSortOrderDesc,
}

// GetListComputedUsagesSortOrderEnumValues Enumerates the set of values for ListComputedUsagesSortOrderEnum
func GetListComputedUsagesSortOrderEnumValues() []ListComputedUsagesSortOrderEnum {
	values := make([]ListComputedUsagesSortOrderEnum, 0)
	for _, v := range mappingListComputedUsagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputedUsagesSortOrderEnumStringValues Enumerates the set of values in String for ListComputedUsagesSortOrderEnum
func GetListComputedUsagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListComputedUsagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputedUsagesSortOrderEnum(val string) (ListComputedUsagesSortOrderEnum, bool) {
	enum, ok := mappingListComputedUsagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListComputedUsagesSortByEnum Enum with underlying type: string
type ListComputedUsagesSortByEnum string

// Set of constants representing the allowable values for ListComputedUsagesSortByEnum
const (
	ListComputedUsagesSortByTimecreated   ListComputedUsagesSortByEnum = "timeCreated"
	ListComputedUsagesSortByTimeofarrival ListComputedUsagesSortByEnum = "timeOfArrival"
	ListComputedUsagesSortByTimemeteredon ListComputedUsagesSortByEnum = "timeMeteredOn"
)

var mappingListComputedUsagesSortByEnum = map[string]ListComputedUsagesSortByEnum{
	"timeCreated":   ListComputedUsagesSortByTimecreated,
	"timeOfArrival": ListComputedUsagesSortByTimeofarrival,
	"timeMeteredOn": ListComputedUsagesSortByTimemeteredon,
}

var mappingListComputedUsagesSortByEnumLowerCase = map[string]ListComputedUsagesSortByEnum{
	"timecreated":   ListComputedUsagesSortByTimecreated,
	"timeofarrival": ListComputedUsagesSortByTimeofarrival,
	"timemeteredon": ListComputedUsagesSortByTimemeteredon,
}

// GetListComputedUsagesSortByEnumValues Enumerates the set of values for ListComputedUsagesSortByEnum
func GetListComputedUsagesSortByEnumValues() []ListComputedUsagesSortByEnum {
	values := make([]ListComputedUsagesSortByEnum, 0)
	for _, v := range mappingListComputedUsagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputedUsagesSortByEnumStringValues Enumerates the set of values in String for ListComputedUsagesSortByEnum
func GetListComputedUsagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeOfArrival",
		"timeMeteredOn",
	}
}

// GetMappingListComputedUsagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputedUsagesSortByEnum(val string) (ListComputedUsagesSortByEnum, bool) {
	enum, ok := mappingListComputedUsagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
