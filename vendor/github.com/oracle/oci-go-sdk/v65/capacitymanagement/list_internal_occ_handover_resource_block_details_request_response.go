// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInternalOccHandoverResourceBlockDetailsRequest wrapper for the ListInternalOccHandoverResourceBlockDetails operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccHandoverResourceBlockDetails.go.html to see an example of how to use ListInternalOccHandoverResourceBlockDetailsRequest.
type ListInternalOccHandoverResourceBlockDetailsRequest struct {

	// The OCID of the OccHandoverResource which is a required query parameter for listing OccHandoverResourceDetails.
	OccHandoverResourceBlockId *string `mandatory:"true" contributesTo:"query" name:"occHandoverResourceBlockId"`

	// This fiter is applicable only for COMPUTE namespace. It helps in fetching of all resource block details for which the hostId is equal to the one provided in this query param.
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccHandoverResourceBlockDetailsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// The default order for handoverDate is chronological order(latest date item at the end).
	SortBy ListInternalOccHandoverResourceBlockDetailsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccHandoverResourceBlockDetailsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccHandoverResourceBlockDetailsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccHandoverResourceBlockDetailsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccHandoverResourceBlockDetailsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccHandoverResourceBlockDetailsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccHandoverResourceBlockDetailsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccHandoverResourceBlockDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccHandoverResourceBlockDetailsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccHandoverResourceBlockDetailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccHandoverResourceBlockDetailsResponse wrapper for the ListInternalOccHandoverResourceBlockDetails operation
type ListInternalOccHandoverResourceBlockDetailsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccHandoverResourceBlockDetailCollection instances
	OccHandoverResourceBlockDetailCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccHandoverResourceBlockDetailsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccHandoverResourceBlockDetailsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccHandoverResourceBlockDetailsSortOrderEnum Enum with underlying type: string
type ListInternalOccHandoverResourceBlockDetailsSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccHandoverResourceBlockDetailsSortOrderEnum
const (
	ListInternalOccHandoverResourceBlockDetailsSortOrderAsc  ListInternalOccHandoverResourceBlockDetailsSortOrderEnum = "ASC"
	ListInternalOccHandoverResourceBlockDetailsSortOrderDesc ListInternalOccHandoverResourceBlockDetailsSortOrderEnum = "DESC"
)

var mappingListInternalOccHandoverResourceBlockDetailsSortOrderEnum = map[string]ListInternalOccHandoverResourceBlockDetailsSortOrderEnum{
	"ASC":  ListInternalOccHandoverResourceBlockDetailsSortOrderAsc,
	"DESC": ListInternalOccHandoverResourceBlockDetailsSortOrderDesc,
}

var mappingListInternalOccHandoverResourceBlockDetailsSortOrderEnumLowerCase = map[string]ListInternalOccHandoverResourceBlockDetailsSortOrderEnum{
	"asc":  ListInternalOccHandoverResourceBlockDetailsSortOrderAsc,
	"desc": ListInternalOccHandoverResourceBlockDetailsSortOrderDesc,
}

// GetListInternalOccHandoverResourceBlockDetailsSortOrderEnumValues Enumerates the set of values for ListInternalOccHandoverResourceBlockDetailsSortOrderEnum
func GetListInternalOccHandoverResourceBlockDetailsSortOrderEnumValues() []ListInternalOccHandoverResourceBlockDetailsSortOrderEnum {
	values := make([]ListInternalOccHandoverResourceBlockDetailsSortOrderEnum, 0)
	for _, v := range mappingListInternalOccHandoverResourceBlockDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccHandoverResourceBlockDetailsSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccHandoverResourceBlockDetailsSortOrderEnum
func GetListInternalOccHandoverResourceBlockDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccHandoverResourceBlockDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccHandoverResourceBlockDetailsSortOrderEnum(val string) (ListInternalOccHandoverResourceBlockDetailsSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccHandoverResourceBlockDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccHandoverResourceBlockDetailsSortByEnum Enum with underlying type: string
type ListInternalOccHandoverResourceBlockDetailsSortByEnum string

// Set of constants representing the allowable values for ListInternalOccHandoverResourceBlockDetailsSortByEnum
const (
	ListInternalOccHandoverResourceBlockDetailsSortByHandoverdate ListInternalOccHandoverResourceBlockDetailsSortByEnum = "handoverDate"
)

var mappingListInternalOccHandoverResourceBlockDetailsSortByEnum = map[string]ListInternalOccHandoverResourceBlockDetailsSortByEnum{
	"handoverDate": ListInternalOccHandoverResourceBlockDetailsSortByHandoverdate,
}

var mappingListInternalOccHandoverResourceBlockDetailsSortByEnumLowerCase = map[string]ListInternalOccHandoverResourceBlockDetailsSortByEnum{
	"handoverdate": ListInternalOccHandoverResourceBlockDetailsSortByHandoverdate,
}

// GetListInternalOccHandoverResourceBlockDetailsSortByEnumValues Enumerates the set of values for ListInternalOccHandoverResourceBlockDetailsSortByEnum
func GetListInternalOccHandoverResourceBlockDetailsSortByEnumValues() []ListInternalOccHandoverResourceBlockDetailsSortByEnum {
	values := make([]ListInternalOccHandoverResourceBlockDetailsSortByEnum, 0)
	for _, v := range mappingListInternalOccHandoverResourceBlockDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccHandoverResourceBlockDetailsSortByEnumStringValues Enumerates the set of values in String for ListInternalOccHandoverResourceBlockDetailsSortByEnum
func GetListInternalOccHandoverResourceBlockDetailsSortByEnumStringValues() []string {
	return []string{
		"handoverDate",
	}
}

// GetMappingListInternalOccHandoverResourceBlockDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccHandoverResourceBlockDetailsSortByEnum(val string) (ListInternalOccHandoverResourceBlockDetailsSortByEnum, bool) {
	enum, ok := mappingListInternalOccHandoverResourceBlockDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
