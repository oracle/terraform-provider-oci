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

// ListSubscribedServicesRequest wrapper for the ListSubscribedServices operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListSubscribedServices.go.html to see an example of how to use ListSubscribedServicesRequest.
type ListSubscribedServicesRequest struct {

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Line level Subscription Id
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// Order Line identifier at subscribed service level . This identifier is originated in Order Management module. Default is null.
	OrderLineId *int64 `mandatory:"false" contributesTo:"query" name:"orderLineId"`

	// This param is used to filter subscribed services based on its status
	Status *string `mandatory:"false" contributesTo:"query" name:"status"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: '500'
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListSubscribedServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order ('sortOrder').
	SortBy ListSubscribedServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscribedServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscribedServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscribedServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscribedServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubscribedServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubscribedServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubscribedServicesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubscribedServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubscribedServicesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubscribedServicesResponse wrapper for the ListSubscribedServices operation
type ListSubscribedServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SubscribedServiceSummary instances
	Items []SubscribedServiceSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSubscribedServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscribedServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscribedServicesSortOrderEnum Enum with underlying type: string
type ListSubscribedServicesSortOrderEnum string

// Set of constants representing the allowable values for ListSubscribedServicesSortOrderEnum
const (
	ListSubscribedServicesSortOrderAsc  ListSubscribedServicesSortOrderEnum = "ASC"
	ListSubscribedServicesSortOrderDesc ListSubscribedServicesSortOrderEnum = "DESC"
)

var mappingListSubscribedServicesSortOrderEnum = map[string]ListSubscribedServicesSortOrderEnum{
	"ASC":  ListSubscribedServicesSortOrderAsc,
	"DESC": ListSubscribedServicesSortOrderDesc,
}

var mappingListSubscribedServicesSortOrderEnumLowerCase = map[string]ListSubscribedServicesSortOrderEnum{
	"asc":  ListSubscribedServicesSortOrderAsc,
	"desc": ListSubscribedServicesSortOrderDesc,
}

// GetListSubscribedServicesSortOrderEnumValues Enumerates the set of values for ListSubscribedServicesSortOrderEnum
func GetListSubscribedServicesSortOrderEnumValues() []ListSubscribedServicesSortOrderEnum {
	values := make([]ListSubscribedServicesSortOrderEnum, 0)
	for _, v := range mappingListSubscribedServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscribedServicesSortOrderEnumStringValues Enumerates the set of values in String for ListSubscribedServicesSortOrderEnum
func GetListSubscribedServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubscribedServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscribedServicesSortOrderEnum(val string) (ListSubscribedServicesSortOrderEnum, bool) {
	enum, ok := mappingListSubscribedServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubscribedServicesSortByEnum Enum with underlying type: string
type ListSubscribedServicesSortByEnum string

// Set of constants representing the allowable values for ListSubscribedServicesSortByEnum
const (
	ListSubscribedServicesSortByOrdernumber   ListSubscribedServicesSortByEnum = "ORDERNUMBER"
	ListSubscribedServicesSortByTimeinvoicing ListSubscribedServicesSortByEnum = "TIMEINVOICING"
)

var mappingListSubscribedServicesSortByEnum = map[string]ListSubscribedServicesSortByEnum{
	"ORDERNUMBER":   ListSubscribedServicesSortByOrdernumber,
	"TIMEINVOICING": ListSubscribedServicesSortByTimeinvoicing,
}

var mappingListSubscribedServicesSortByEnumLowerCase = map[string]ListSubscribedServicesSortByEnum{
	"ordernumber":   ListSubscribedServicesSortByOrdernumber,
	"timeinvoicing": ListSubscribedServicesSortByTimeinvoicing,
}

// GetListSubscribedServicesSortByEnumValues Enumerates the set of values for ListSubscribedServicesSortByEnum
func GetListSubscribedServicesSortByEnumValues() []ListSubscribedServicesSortByEnum {
	values := make([]ListSubscribedServicesSortByEnum, 0)
	for _, v := range mappingListSubscribedServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscribedServicesSortByEnumStringValues Enumerates the set of values in String for ListSubscribedServicesSortByEnum
func GetListSubscribedServicesSortByEnumStringValues() []string {
	return []string{
		"ORDERNUMBER",
		"TIMEINVOICING",
	}
}

// GetMappingListSubscribedServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscribedServicesSortByEnum(val string) (ListSubscribedServicesSortByEnum, bool) {
	enum, ok := mappingListSubscribedServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
