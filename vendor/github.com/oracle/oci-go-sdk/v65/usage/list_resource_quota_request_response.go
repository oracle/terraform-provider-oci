// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourceQuotaRequest wrapper for the ListResourceQuota operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListResourceQuota.go.html to see an example of how to use ListResourceQuotaRequest.
type ListResourceQuotaRequest struct {

	// Service Name.
	ServiceName *string `mandatory:"true" contributesTo:"query" name:"serviceName"`

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Service entitlement Id.
	ServiceEntitlement *string `mandatory:"false" contributesTo:"query" name:"serviceEntitlement"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, which can be ascending (ASC) or descending (DESC).
	SortOrder ListResourceQuotaSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Supports one sort order.
	SortBy ListResourceQuotaSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceQuotaRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceQuotaRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceQuotaRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceQuotaRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceQuotaRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceQuotaSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceQuotaSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceQuotaSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceQuotaSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceQuotaResponse wrapper for the ListResourceQuota operation
type ListResourceQuotaResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceQuotumCollection instances
	ResourceQuotumCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListResourceQuotaResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceQuotaResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceQuotaSortOrderEnum Enum with underlying type: string
type ListResourceQuotaSortOrderEnum string

// Set of constants representing the allowable values for ListResourceQuotaSortOrderEnum
const (
	ListResourceQuotaSortOrderAsc  ListResourceQuotaSortOrderEnum = "ASC"
	ListResourceQuotaSortOrderDesc ListResourceQuotaSortOrderEnum = "DESC"
)

var mappingListResourceQuotaSortOrderEnum = map[string]ListResourceQuotaSortOrderEnum{
	"ASC":  ListResourceQuotaSortOrderAsc,
	"DESC": ListResourceQuotaSortOrderDesc,
}

var mappingListResourceQuotaSortOrderEnumLowerCase = map[string]ListResourceQuotaSortOrderEnum{
	"asc":  ListResourceQuotaSortOrderAsc,
	"desc": ListResourceQuotaSortOrderDesc,
}

// GetListResourceQuotaSortOrderEnumValues Enumerates the set of values for ListResourceQuotaSortOrderEnum
func GetListResourceQuotaSortOrderEnumValues() []ListResourceQuotaSortOrderEnum {
	values := make([]ListResourceQuotaSortOrderEnum, 0)
	for _, v := range mappingListResourceQuotaSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceQuotaSortOrderEnumStringValues Enumerates the set of values in String for ListResourceQuotaSortOrderEnum
func GetListResourceQuotaSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceQuotaSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceQuotaSortOrderEnum(val string) (ListResourceQuotaSortOrderEnum, bool) {
	enum, ok := mappingListResourceQuotaSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceQuotaSortByEnum Enum with underlying type: string
type ListResourceQuotaSortByEnum string

// Set of constants representing the allowable values for ListResourceQuotaSortByEnum
const (
	ListResourceQuotaSortByTimecreated ListResourceQuotaSortByEnum = "TIMECREATED"
	ListResourceQuotaSortByTimestart   ListResourceQuotaSortByEnum = "TIMESTART"
)

var mappingListResourceQuotaSortByEnum = map[string]ListResourceQuotaSortByEnum{
	"TIMECREATED": ListResourceQuotaSortByTimecreated,
	"TIMESTART":   ListResourceQuotaSortByTimestart,
}

var mappingListResourceQuotaSortByEnumLowerCase = map[string]ListResourceQuotaSortByEnum{
	"timecreated": ListResourceQuotaSortByTimecreated,
	"timestart":   ListResourceQuotaSortByTimestart,
}

// GetListResourceQuotaSortByEnumValues Enumerates the set of values for ListResourceQuotaSortByEnum
func GetListResourceQuotaSortByEnumValues() []ListResourceQuotaSortByEnum {
	values := make([]ListResourceQuotaSortByEnum, 0)
	for _, v := range mappingListResourceQuotaSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceQuotaSortByEnumStringValues Enumerates the set of values in String for ListResourceQuotaSortByEnum
func GetListResourceQuotaSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMESTART",
	}
}

// GetMappingListResourceQuotaSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceQuotaSortByEnum(val string) (ListResourceQuotaSortByEnum, bool) {
	enum, ok := mappingListResourceQuotaSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
