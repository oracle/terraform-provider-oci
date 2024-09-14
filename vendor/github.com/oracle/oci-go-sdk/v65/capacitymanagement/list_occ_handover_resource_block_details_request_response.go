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

// ListOccHandoverResourceBlockDetailsRequest wrapper for the ListOccHandoverResourceBlockDetails operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccHandoverResourceBlockDetails.go.html to see an example of how to use ListOccHandoverResourceBlockDetailsRequest.
type ListOccHandoverResourceBlockDetailsRequest struct {

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
	SortOrder ListOccHandoverResourceBlockDetailsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// The default order for handoverDate is chronological order(latest date item at the end).
	SortBy ListOccHandoverResourceBlockDetailsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccHandoverResourceBlockDetailsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccHandoverResourceBlockDetailsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccHandoverResourceBlockDetailsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccHandoverResourceBlockDetailsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccHandoverResourceBlockDetailsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccHandoverResourceBlockDetailsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccHandoverResourceBlockDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccHandoverResourceBlockDetailsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccHandoverResourceBlockDetailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccHandoverResourceBlockDetailsResponse wrapper for the ListOccHandoverResourceBlockDetails operation
type ListOccHandoverResourceBlockDetailsResponse struct {

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

func (response ListOccHandoverResourceBlockDetailsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccHandoverResourceBlockDetailsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccHandoverResourceBlockDetailsSortOrderEnum Enum with underlying type: string
type ListOccHandoverResourceBlockDetailsSortOrderEnum string

// Set of constants representing the allowable values for ListOccHandoverResourceBlockDetailsSortOrderEnum
const (
	ListOccHandoverResourceBlockDetailsSortOrderAsc  ListOccHandoverResourceBlockDetailsSortOrderEnum = "ASC"
	ListOccHandoverResourceBlockDetailsSortOrderDesc ListOccHandoverResourceBlockDetailsSortOrderEnum = "DESC"
)

var mappingListOccHandoverResourceBlockDetailsSortOrderEnum = map[string]ListOccHandoverResourceBlockDetailsSortOrderEnum{
	"ASC":  ListOccHandoverResourceBlockDetailsSortOrderAsc,
	"DESC": ListOccHandoverResourceBlockDetailsSortOrderDesc,
}

var mappingListOccHandoverResourceBlockDetailsSortOrderEnumLowerCase = map[string]ListOccHandoverResourceBlockDetailsSortOrderEnum{
	"asc":  ListOccHandoverResourceBlockDetailsSortOrderAsc,
	"desc": ListOccHandoverResourceBlockDetailsSortOrderDesc,
}

// GetListOccHandoverResourceBlockDetailsSortOrderEnumValues Enumerates the set of values for ListOccHandoverResourceBlockDetailsSortOrderEnum
func GetListOccHandoverResourceBlockDetailsSortOrderEnumValues() []ListOccHandoverResourceBlockDetailsSortOrderEnum {
	values := make([]ListOccHandoverResourceBlockDetailsSortOrderEnum, 0)
	for _, v := range mappingListOccHandoverResourceBlockDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccHandoverResourceBlockDetailsSortOrderEnumStringValues Enumerates the set of values in String for ListOccHandoverResourceBlockDetailsSortOrderEnum
func GetListOccHandoverResourceBlockDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccHandoverResourceBlockDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccHandoverResourceBlockDetailsSortOrderEnum(val string) (ListOccHandoverResourceBlockDetailsSortOrderEnum, bool) {
	enum, ok := mappingListOccHandoverResourceBlockDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccHandoverResourceBlockDetailsSortByEnum Enum with underlying type: string
type ListOccHandoverResourceBlockDetailsSortByEnum string

// Set of constants representing the allowable values for ListOccHandoverResourceBlockDetailsSortByEnum
const (
	ListOccHandoverResourceBlockDetailsSortByHandoverdate ListOccHandoverResourceBlockDetailsSortByEnum = "handoverDate"
)

var mappingListOccHandoverResourceBlockDetailsSortByEnum = map[string]ListOccHandoverResourceBlockDetailsSortByEnum{
	"handoverDate": ListOccHandoverResourceBlockDetailsSortByHandoverdate,
}

var mappingListOccHandoverResourceBlockDetailsSortByEnumLowerCase = map[string]ListOccHandoverResourceBlockDetailsSortByEnum{
	"handoverdate": ListOccHandoverResourceBlockDetailsSortByHandoverdate,
}

// GetListOccHandoverResourceBlockDetailsSortByEnumValues Enumerates the set of values for ListOccHandoverResourceBlockDetailsSortByEnum
func GetListOccHandoverResourceBlockDetailsSortByEnumValues() []ListOccHandoverResourceBlockDetailsSortByEnum {
	values := make([]ListOccHandoverResourceBlockDetailsSortByEnum, 0)
	for _, v := range mappingListOccHandoverResourceBlockDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccHandoverResourceBlockDetailsSortByEnumStringValues Enumerates the set of values in String for ListOccHandoverResourceBlockDetailsSortByEnum
func GetListOccHandoverResourceBlockDetailsSortByEnumStringValues() []string {
	return []string{
		"handoverDate",
	}
}

// GetMappingListOccHandoverResourceBlockDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccHandoverResourceBlockDetailsSortByEnum(val string) (ListOccHandoverResourceBlockDetailsSortByEnum, bool) {
	enum, ok := mappingListOccHandoverResourceBlockDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
