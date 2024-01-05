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

// ListInvoicesRequest wrapper for the ListInvoices operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListInvoices.go.html to see an example of how to use ListInvoicesRequest.
type ListInvoicesRequest struct {

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// AR Unique identifier for an invoice .
	ArCustomerTransactionId *string `mandatory:"true" contributesTo:"query" name:"arCustomerTransactionId"`

	// Initial date to filter Invoice data in SPM.
	TimeFrom *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFrom"`

	// Final date to filter Invoice data in SPM.
	TimeTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeTo"`

	// The sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListInvoicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order ('sortOrder').
	SortBy ListInvoicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: '500'
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Partial response refers to an optimization technique offered
	// by the RESTful web APIs to return only the information
	// (fields) required by the client. This parameter is used to control what fields to
	// return.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInvoicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInvoicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInvoicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInvoicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInvoicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInvoicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInvoicesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInvoicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInvoicesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInvoicesResponse wrapper for the ListInvoices operation
type ListInvoicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InvoiceSummary instances
	Items []InvoiceSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInvoicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInvoicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInvoicesSortOrderEnum Enum with underlying type: string
type ListInvoicesSortOrderEnum string

// Set of constants representing the allowable values for ListInvoicesSortOrderEnum
const (
	ListInvoicesSortOrderAsc  ListInvoicesSortOrderEnum = "ASC"
	ListInvoicesSortOrderDesc ListInvoicesSortOrderEnum = "DESC"
)

var mappingListInvoicesSortOrderEnum = map[string]ListInvoicesSortOrderEnum{
	"ASC":  ListInvoicesSortOrderAsc,
	"DESC": ListInvoicesSortOrderDesc,
}

var mappingListInvoicesSortOrderEnumLowerCase = map[string]ListInvoicesSortOrderEnum{
	"asc":  ListInvoicesSortOrderAsc,
	"desc": ListInvoicesSortOrderDesc,
}

// GetListInvoicesSortOrderEnumValues Enumerates the set of values for ListInvoicesSortOrderEnum
func GetListInvoicesSortOrderEnumValues() []ListInvoicesSortOrderEnum {
	values := make([]ListInvoicesSortOrderEnum, 0)
	for _, v := range mappingListInvoicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInvoicesSortOrderEnumStringValues Enumerates the set of values in String for ListInvoicesSortOrderEnum
func GetListInvoicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInvoicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInvoicesSortOrderEnum(val string) (ListInvoicesSortOrderEnum, bool) {
	enum, ok := mappingListInvoicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInvoicesSortByEnum Enum with underlying type: string
type ListInvoicesSortByEnum string

// Set of constants representing the allowable values for ListInvoicesSortByEnum
const (
	ListInvoicesSortByOrdernumber   ListInvoicesSortByEnum = "ORDERNUMBER"
	ListInvoicesSortByTimeinvoicing ListInvoicesSortByEnum = "TIMEINVOICING"
)

var mappingListInvoicesSortByEnum = map[string]ListInvoicesSortByEnum{
	"ORDERNUMBER":   ListInvoicesSortByOrdernumber,
	"TIMEINVOICING": ListInvoicesSortByTimeinvoicing,
}

var mappingListInvoicesSortByEnumLowerCase = map[string]ListInvoicesSortByEnum{
	"ordernumber":   ListInvoicesSortByOrdernumber,
	"timeinvoicing": ListInvoicesSortByTimeinvoicing,
}

// GetListInvoicesSortByEnumValues Enumerates the set of values for ListInvoicesSortByEnum
func GetListInvoicesSortByEnumValues() []ListInvoicesSortByEnum {
	values := make([]ListInvoicesSortByEnum, 0)
	for _, v := range mappingListInvoicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInvoicesSortByEnumStringValues Enumerates the set of values in String for ListInvoicesSortByEnum
func GetListInvoicesSortByEnumStringValues() []string {
	return []string{
		"ORDERNUMBER",
		"TIMEINVOICING",
	}
}

// GetMappingListInvoicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInvoicesSortByEnum(val string) (ListInvoicesSortByEnum, bool) {
	enum, ok := mappingListInvoicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
