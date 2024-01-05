// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ospgateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubscriptionsRequest wrapper for the ListSubscriptions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/ListSubscriptions.go.html to see an example of how to use ListSubscriptionsRequest.
type ListSubscriptionsRequest struct {

	// The home region's public name of the logged in user.
	OspHomeRegion *string `mandatory:"true" contributesTo:"query" name:"ospHomeRegion"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one field can be selected for sorting.
	SortBy ListSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ascending or descending).
	SortOrder ListSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubscriptionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubscriptionsResponse wrapper for the ListSubscriptions operation
type ListSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubscriptionCollection instances
	SubscriptionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. If this header appears in the response, then this
	// is a partial list of invoices. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of invoices.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscriptionsSortByEnum Enum with underlying type: string
type ListSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortByEnum
const (
	ListSubscriptionsSortByInvoiceNo   ListSubscriptionsSortByEnum = "INVOICE_NO"
	ListSubscriptionsSortByRefNo       ListSubscriptionsSortByEnum = "REF_NO"
	ListSubscriptionsSortByStatus      ListSubscriptionsSortByEnum = "STATUS"
	ListSubscriptionsSortByType        ListSubscriptionsSortByEnum = "TYPE"
	ListSubscriptionsSortByInvoiceDate ListSubscriptionsSortByEnum = "INVOICE_DATE"
	ListSubscriptionsSortByDueDate     ListSubscriptionsSortByEnum = "DUE_DATE"
	ListSubscriptionsSortByPaymRef     ListSubscriptionsSortByEnum = "PAYM_REF"
	ListSubscriptionsSortByTotalAmount ListSubscriptionsSortByEnum = "TOTAL_AMOUNT"
	ListSubscriptionsSortByBalanceDue  ListSubscriptionsSortByEnum = "BALANCE_DUE"
)

var mappingListSubscriptionsSortByEnum = map[string]ListSubscriptionsSortByEnum{
	"INVOICE_NO":   ListSubscriptionsSortByInvoiceNo,
	"REF_NO":       ListSubscriptionsSortByRefNo,
	"STATUS":       ListSubscriptionsSortByStatus,
	"TYPE":         ListSubscriptionsSortByType,
	"INVOICE_DATE": ListSubscriptionsSortByInvoiceDate,
	"DUE_DATE":     ListSubscriptionsSortByDueDate,
	"PAYM_REF":     ListSubscriptionsSortByPaymRef,
	"TOTAL_AMOUNT": ListSubscriptionsSortByTotalAmount,
	"BALANCE_DUE":  ListSubscriptionsSortByBalanceDue,
}

var mappingListSubscriptionsSortByEnumLowerCase = map[string]ListSubscriptionsSortByEnum{
	"invoice_no":   ListSubscriptionsSortByInvoiceNo,
	"ref_no":       ListSubscriptionsSortByRefNo,
	"status":       ListSubscriptionsSortByStatus,
	"type":         ListSubscriptionsSortByType,
	"invoice_date": ListSubscriptionsSortByInvoiceDate,
	"due_date":     ListSubscriptionsSortByDueDate,
	"paym_ref":     ListSubscriptionsSortByPaymRef,
	"total_amount": ListSubscriptionsSortByTotalAmount,
	"balance_due":  ListSubscriptionsSortByBalanceDue,
}

// GetListSubscriptionsSortByEnumValues Enumerates the set of values for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumValues() []ListSubscriptionsSortByEnum {
	values := make([]ListSubscriptionsSortByEnum, 0)
	for _, v := range mappingListSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"INVOICE_NO",
		"REF_NO",
		"STATUS",
		"TYPE",
		"INVOICE_DATE",
		"DUE_DATE",
		"PAYM_REF",
		"TOTAL_AMOUNT",
		"BALANCE_DUE",
	}
}

// GetMappingListSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsSortByEnum(val string) (ListSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubscriptionsSortOrderEnum Enum with underlying type: string
type ListSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortOrderEnum
const (
	ListSubscriptionsSortOrderAsc  ListSubscriptionsSortOrderEnum = "ASC"
	ListSubscriptionsSortOrderDesc ListSubscriptionsSortOrderEnum = "DESC"
)

var mappingListSubscriptionsSortOrderEnum = map[string]ListSubscriptionsSortOrderEnum{
	"ASC":  ListSubscriptionsSortOrderAsc,
	"DESC": ListSubscriptionsSortOrderDesc,
}

var mappingListSubscriptionsSortOrderEnumLowerCase = map[string]ListSubscriptionsSortOrderEnum{
	"asc":  ListSubscriptionsSortOrderAsc,
	"desc": ListSubscriptionsSortOrderDesc,
}

// GetListSubscriptionsSortOrderEnumValues Enumerates the set of values for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumValues() []ListSubscriptionsSortOrderEnum {
	values := make([]ListSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsSortOrderEnum(val string) (ListSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
