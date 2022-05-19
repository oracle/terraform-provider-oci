// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVirtualServicesRequest wrapper for the ListVirtualServices operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListVirtualServices.go.html to see an example of how to use ListVirtualServicesRequest.
type ListVirtualServicesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVirtualServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for 'timeCreated' is descending. Default order for 'name' is ascending.
	SortBy ListVirtualServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique Mesh identifier.
	MeshId *string `mandatory:"false" contributesTo:"query" name:"meshId"`

	// Unique VirtualService identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVirtualServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVirtualServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVirtualServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVirtualServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVirtualServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVirtualServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVirtualServicesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVirtualServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVirtualServicesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVirtualServicesResponse wrapper for the ListVirtualServices operation
type ListVirtualServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VirtualServiceCollection instances
	VirtualServiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVirtualServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVirtualServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVirtualServicesSortOrderEnum Enum with underlying type: string
type ListVirtualServicesSortOrderEnum string

// Set of constants representing the allowable values for ListVirtualServicesSortOrderEnum
const (
	ListVirtualServicesSortOrderAsc  ListVirtualServicesSortOrderEnum = "ASC"
	ListVirtualServicesSortOrderDesc ListVirtualServicesSortOrderEnum = "DESC"
)

var mappingListVirtualServicesSortOrderEnum = map[string]ListVirtualServicesSortOrderEnum{
	"ASC":  ListVirtualServicesSortOrderAsc,
	"DESC": ListVirtualServicesSortOrderDesc,
}

var mappingListVirtualServicesSortOrderEnumLowerCase = map[string]ListVirtualServicesSortOrderEnum{
	"asc":  ListVirtualServicesSortOrderAsc,
	"desc": ListVirtualServicesSortOrderDesc,
}

// GetListVirtualServicesSortOrderEnumValues Enumerates the set of values for ListVirtualServicesSortOrderEnum
func GetListVirtualServicesSortOrderEnumValues() []ListVirtualServicesSortOrderEnum {
	values := make([]ListVirtualServicesSortOrderEnum, 0)
	for _, v := range mappingListVirtualServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVirtualServicesSortOrderEnumStringValues Enumerates the set of values in String for ListVirtualServicesSortOrderEnum
func GetListVirtualServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVirtualServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVirtualServicesSortOrderEnum(val string) (ListVirtualServicesSortOrderEnum, bool) {
	enum, ok := mappingListVirtualServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVirtualServicesSortByEnum Enum with underlying type: string
type ListVirtualServicesSortByEnum string

// Set of constants representing the allowable values for ListVirtualServicesSortByEnum
const (
	ListVirtualServicesSortById          ListVirtualServicesSortByEnum = "id"
	ListVirtualServicesSortByTimecreated ListVirtualServicesSortByEnum = "timeCreated"
	ListVirtualServicesSortByName        ListVirtualServicesSortByEnum = "name"
)

var mappingListVirtualServicesSortByEnum = map[string]ListVirtualServicesSortByEnum{
	"id":          ListVirtualServicesSortById,
	"timeCreated": ListVirtualServicesSortByTimecreated,
	"name":        ListVirtualServicesSortByName,
}

var mappingListVirtualServicesSortByEnumLowerCase = map[string]ListVirtualServicesSortByEnum{
	"id":          ListVirtualServicesSortById,
	"timecreated": ListVirtualServicesSortByTimecreated,
	"name":        ListVirtualServicesSortByName,
}

// GetListVirtualServicesSortByEnumValues Enumerates the set of values for ListVirtualServicesSortByEnum
func GetListVirtualServicesSortByEnumValues() []ListVirtualServicesSortByEnum {
	values := make([]ListVirtualServicesSortByEnum, 0)
	for _, v := range mappingListVirtualServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVirtualServicesSortByEnumStringValues Enumerates the set of values in String for ListVirtualServicesSortByEnum
func GetListVirtualServicesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"name",
	}
}

// GetMappingListVirtualServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVirtualServicesSortByEnum(val string) (ListVirtualServicesSortByEnum, bool) {
	enum, ok := mappingListVirtualServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
