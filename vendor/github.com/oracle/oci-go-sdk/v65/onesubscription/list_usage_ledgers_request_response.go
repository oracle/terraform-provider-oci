// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUsageLedgersRequest wrapper for the ListUsageLedgers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListUsageLedgers.go.html to see an example of how to use ListUsageLedgersRequest.
type ListUsageLedgersRequest struct {

	// Subscription identifier.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: '500'
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The usage-ledger sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListUsageLedgersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The usage-ledger field to sort by. You can provide one sort order ('sortOrder').
	SortBy ListUsageLedgersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The lower bound of the consumption-date range, expressed in RFC 3339 timestamp format. The value is inclusive (>=).
	TimeConsumptionGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeConsumptionGreaterThanOrEqualTo"`

	// The upper bound of the consumption-date range, expressed in RFC 3339 timestamp format. The value is exclusive (<).
	TimeConsumptionLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeConsumptionLessThan"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsageLedgersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsageLedgersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsageLedgersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsageLedgersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUsageLedgersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUsageLedgersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUsageLedgersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsageLedgersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUsageLedgersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUsageLedgersResponse wrapper for the ListUsageLedgers operation
type ListUsageLedgersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UsageLedgerSummary instances
	Items []UsageLedgerSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListUsageLedgersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsageLedgersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsageLedgersSortOrderEnum Enum with underlying type: string
type ListUsageLedgersSortOrderEnum string

// Set of constants representing the allowable values for ListUsageLedgersSortOrderEnum
const (
	ListUsageLedgersSortOrderAsc  ListUsageLedgersSortOrderEnum = "ASC"
	ListUsageLedgersSortOrderDesc ListUsageLedgersSortOrderEnum = "DESC"
)

var mappingListUsageLedgersSortOrderEnum = map[string]ListUsageLedgersSortOrderEnum{
	"ASC":  ListUsageLedgersSortOrderAsc,
	"DESC": ListUsageLedgersSortOrderDesc,
}

var mappingListUsageLedgersSortOrderEnumLowerCase = map[string]ListUsageLedgersSortOrderEnum{
	"asc":  ListUsageLedgersSortOrderAsc,
	"desc": ListUsageLedgersSortOrderDesc,
}

// GetListUsageLedgersSortOrderEnumValues Enumerates the set of values for ListUsageLedgersSortOrderEnum
func GetListUsageLedgersSortOrderEnumValues() []ListUsageLedgersSortOrderEnum {
	values := make([]ListUsageLedgersSortOrderEnum, 0)
	for _, v := range mappingListUsageLedgersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsageLedgersSortOrderEnumStringValues Enumerates the set of values in String for ListUsageLedgersSortOrderEnum
func GetListUsageLedgersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUsageLedgersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsageLedgersSortOrderEnum(val string) (ListUsageLedgersSortOrderEnum, bool) {
	enum, ok := mappingListUsageLedgersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUsageLedgersSortByEnum Enum with underlying type: string
type ListUsageLedgersSortByEnum string

// Set of constants representing the allowable values for ListUsageLedgersSortByEnum
const (
	ListUsageLedgersSortByTimeconsumption   ListUsageLedgersSortByEnum = "timeConsumption"
	ListUsageLedgersSortByConsumptionamount ListUsageLedgersSortByEnum = "consumptionAmount"
)

var mappingListUsageLedgersSortByEnum = map[string]ListUsageLedgersSortByEnum{
	"timeConsumption":   ListUsageLedgersSortByTimeconsumption,
	"consumptionAmount": ListUsageLedgersSortByConsumptionamount,
}

var mappingListUsageLedgersSortByEnumLowerCase = map[string]ListUsageLedgersSortByEnum{
	"timeconsumption":   ListUsageLedgersSortByTimeconsumption,
	"consumptionamount": ListUsageLedgersSortByConsumptionamount,
}

// GetListUsageLedgersSortByEnumValues Enumerates the set of values for ListUsageLedgersSortByEnum
func GetListUsageLedgersSortByEnumValues() []ListUsageLedgersSortByEnum {
	values := make([]ListUsageLedgersSortByEnum, 0)
	for _, v := range mappingListUsageLedgersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsageLedgersSortByEnumStringValues Enumerates the set of values in String for ListUsageLedgersSortByEnum
func GetListUsageLedgersSortByEnumStringValues() []string {
	return []string{
		"timeConsumption",
		"consumptionAmount",
	}
}

// GetMappingListUsageLedgersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsageLedgersSortByEnum(val string) (ListUsageLedgersSortByEnum, bool) {
	enum, ok := mappingListUsageLedgersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
