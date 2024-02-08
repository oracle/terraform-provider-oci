// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStreamsRequest wrapper for the ListStreams operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ListStreams.go.html to see an example of how to use ListStreamsRequest.
type ListStreamsRequest struct {

	// The OCID of the compartment. Is exclusive with the `streamPoolId` parameter. One of them is required.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the stream pool. Is exclusive with the `compartmentId` parameter. One of them is required.
	StreamPoolId *string `mandatory:"false" contributesTo:"query" name:"streamPoolId"`

	// A filter to return only resources that match the given ID exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return. The value must be between 1 and 50. The default is 10.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide no more than one sort order. By default, `TIMECREATED` sorts results in descending order and `NAME` sorts results in ascending order.
	SortBy ListStreamsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListStreamsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState StreamLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStreamsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStreamLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStreamLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamsResponse wrapper for the ListStreams operation
type ListStreamsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []StreamSummary instances
	Items []StreamSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListStreamsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamsSortByEnum Enum with underlying type: string
type ListStreamsSortByEnum string

// Set of constants representing the allowable values for ListStreamsSortByEnum
const (
	ListStreamsSortByName        ListStreamsSortByEnum = "NAME"
	ListStreamsSortByTimecreated ListStreamsSortByEnum = "TIMECREATED"
)

var mappingListStreamsSortByEnum = map[string]ListStreamsSortByEnum{
	"NAME":        ListStreamsSortByName,
	"TIMECREATED": ListStreamsSortByTimecreated,
}

var mappingListStreamsSortByEnumLowerCase = map[string]ListStreamsSortByEnum{
	"name":        ListStreamsSortByName,
	"timecreated": ListStreamsSortByTimecreated,
}

// GetListStreamsSortByEnumValues Enumerates the set of values for ListStreamsSortByEnum
func GetListStreamsSortByEnumValues() []ListStreamsSortByEnum {
	values := make([]ListStreamsSortByEnum, 0)
	for _, v := range mappingListStreamsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamsSortByEnumStringValues Enumerates the set of values in String for ListStreamsSortByEnum
func GetListStreamsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListStreamsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamsSortByEnum(val string) (ListStreamsSortByEnum, bool) {
	enum, ok := mappingListStreamsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamsSortOrderEnum Enum with underlying type: string
type ListStreamsSortOrderEnum string

// Set of constants representing the allowable values for ListStreamsSortOrderEnum
const (
	ListStreamsSortOrderAsc  ListStreamsSortOrderEnum = "ASC"
	ListStreamsSortOrderDesc ListStreamsSortOrderEnum = "DESC"
)

var mappingListStreamsSortOrderEnum = map[string]ListStreamsSortOrderEnum{
	"ASC":  ListStreamsSortOrderAsc,
	"DESC": ListStreamsSortOrderDesc,
}

var mappingListStreamsSortOrderEnumLowerCase = map[string]ListStreamsSortOrderEnum{
	"asc":  ListStreamsSortOrderAsc,
	"desc": ListStreamsSortOrderDesc,
}

// GetListStreamsSortOrderEnumValues Enumerates the set of values for ListStreamsSortOrderEnum
func GetListStreamsSortOrderEnumValues() []ListStreamsSortOrderEnum {
	values := make([]ListStreamsSortOrderEnum, 0)
	for _, v := range mappingListStreamsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamsSortOrderEnumStringValues Enumerates the set of values in String for ListStreamsSortOrderEnum
func GetListStreamsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamsSortOrderEnum(val string) (ListStreamsSortOrderEnum, bool) {
	enum, ok := mappingListStreamsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
