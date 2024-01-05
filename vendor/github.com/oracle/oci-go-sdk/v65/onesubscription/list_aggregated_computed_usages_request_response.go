// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAggregatedComputedUsagesRequest wrapper for the ListAggregatedComputedUsages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListAggregatedComputedUsages.go.html to see an example of how to use ListAggregatedComputedUsagesRequest.
type ListAggregatedComputedUsagesRequest struct {

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

	// Grouping criteria to use for aggregate the computed Usage, either hourly (`HOURLY`), daily (`DAILY`), monthly(`MONTHLY`) or none (`NONE`) to not follow a grouping criteria by date.
	Grouping ListAggregatedComputedUsagesGroupingEnum `mandatory:"false" contributesTo:"query" name:"grouping" omitEmpty:"true"`

	// The maximum number aggregatedComputedUsages of items to return within the Subscription "List" call, this
	// counts the overall count across all items
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAggregatedComputedUsagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAggregatedComputedUsagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAggregatedComputedUsagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAggregatedComputedUsagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAggregatedComputedUsagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAggregatedComputedUsagesGroupingEnum(string(request.Grouping)); !ok && request.Grouping != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Grouping: %s. Supported values are: %s.", request.Grouping, strings.Join(GetListAggregatedComputedUsagesGroupingEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAggregatedComputedUsagesResponse wrapper for the ListAggregatedComputedUsages operation
type ListAggregatedComputedUsagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AggregatedComputedUsageSummary instances
	Items []AggregatedComputedUsageSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAggregatedComputedUsagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAggregatedComputedUsagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAggregatedComputedUsagesGroupingEnum Enum with underlying type: string
type ListAggregatedComputedUsagesGroupingEnum string

// Set of constants representing the allowable values for ListAggregatedComputedUsagesGroupingEnum
const (
	ListAggregatedComputedUsagesGroupingHourly  ListAggregatedComputedUsagesGroupingEnum = "HOURLY"
	ListAggregatedComputedUsagesGroupingDaily   ListAggregatedComputedUsagesGroupingEnum = "DAILY"
	ListAggregatedComputedUsagesGroupingMonthly ListAggregatedComputedUsagesGroupingEnum = "MONTHLY"
	ListAggregatedComputedUsagesGroupingNone    ListAggregatedComputedUsagesGroupingEnum = "NONE"
)

var mappingListAggregatedComputedUsagesGroupingEnum = map[string]ListAggregatedComputedUsagesGroupingEnum{
	"HOURLY":  ListAggregatedComputedUsagesGroupingHourly,
	"DAILY":   ListAggregatedComputedUsagesGroupingDaily,
	"MONTHLY": ListAggregatedComputedUsagesGroupingMonthly,
	"NONE":    ListAggregatedComputedUsagesGroupingNone,
}

var mappingListAggregatedComputedUsagesGroupingEnumLowerCase = map[string]ListAggregatedComputedUsagesGroupingEnum{
	"hourly":  ListAggregatedComputedUsagesGroupingHourly,
	"daily":   ListAggregatedComputedUsagesGroupingDaily,
	"monthly": ListAggregatedComputedUsagesGroupingMonthly,
	"none":    ListAggregatedComputedUsagesGroupingNone,
}

// GetListAggregatedComputedUsagesGroupingEnumValues Enumerates the set of values for ListAggregatedComputedUsagesGroupingEnum
func GetListAggregatedComputedUsagesGroupingEnumValues() []ListAggregatedComputedUsagesGroupingEnum {
	values := make([]ListAggregatedComputedUsagesGroupingEnum, 0)
	for _, v := range mappingListAggregatedComputedUsagesGroupingEnum {
		values = append(values, v)
	}
	return values
}

// GetListAggregatedComputedUsagesGroupingEnumStringValues Enumerates the set of values in String for ListAggregatedComputedUsagesGroupingEnum
func GetListAggregatedComputedUsagesGroupingEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
		"MONTHLY",
		"NONE",
	}
}

// GetMappingListAggregatedComputedUsagesGroupingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAggregatedComputedUsagesGroupingEnum(val string) (ListAggregatedComputedUsagesGroupingEnum, bool) {
	enum, ok := mappingListAggregatedComputedUsagesGroupingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
