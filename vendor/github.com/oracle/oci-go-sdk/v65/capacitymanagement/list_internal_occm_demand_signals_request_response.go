// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInternalOccmDemandSignalsRequest wrapper for the ListInternalOccmDemandSignals operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignals.go.html to see an example of how to use ListInternalOccmDemandSignalsRequest.
type ListInternalOccmDemandSignalsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// A query parameter to filter the list of demand signals based on it's OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A query parameter to filter the list of demand signals based on its state.
	LifecycleDetails *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetails"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccmDemandSignalsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand Signal API. Only one sort order may be provided. The default order for timeCreated is reverse chronological order (latest date at the top). The default order for displayName is case sensitive alphabetical order.
	SortBy ListInternalOccmDemandSignalsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccmDemandSignalsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccmDemandSignalsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccmDemandSignalsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccmDemandSignalsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccmDemandSignalsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccmDemandSignalsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccmDemandSignalsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccmDemandSignalsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccmDemandSignalsResponse wrapper for the ListInternalOccmDemandSignals operation
type ListInternalOccmDemandSignalsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InternalOccmDemandSignalCollection instances
	InternalOccmDemandSignalCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccmDemandSignalsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccmDemandSignalsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccmDemandSignalsSortOrderEnum Enum with underlying type: string
type ListInternalOccmDemandSignalsSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalsSortOrderEnum
const (
	ListInternalOccmDemandSignalsSortOrderAsc  ListInternalOccmDemandSignalsSortOrderEnum = "ASC"
	ListInternalOccmDemandSignalsSortOrderDesc ListInternalOccmDemandSignalsSortOrderEnum = "DESC"
)

var mappingListInternalOccmDemandSignalsSortOrderEnum = map[string]ListInternalOccmDemandSignalsSortOrderEnum{
	"ASC":  ListInternalOccmDemandSignalsSortOrderAsc,
	"DESC": ListInternalOccmDemandSignalsSortOrderDesc,
}

var mappingListInternalOccmDemandSignalsSortOrderEnumLowerCase = map[string]ListInternalOccmDemandSignalsSortOrderEnum{
	"asc":  ListInternalOccmDemandSignalsSortOrderAsc,
	"desc": ListInternalOccmDemandSignalsSortOrderDesc,
}

// GetListInternalOccmDemandSignalsSortOrderEnumValues Enumerates the set of values for ListInternalOccmDemandSignalsSortOrderEnum
func GetListInternalOccmDemandSignalsSortOrderEnumValues() []ListInternalOccmDemandSignalsSortOrderEnum {
	values := make([]ListInternalOccmDemandSignalsSortOrderEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalsSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalsSortOrderEnum
func GetListInternalOccmDemandSignalsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccmDemandSignalsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalsSortOrderEnum(val string) (ListInternalOccmDemandSignalsSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalsSortByEnum Enum with underlying type: string
type ListInternalOccmDemandSignalsSortByEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalsSortByEnum
const (
	ListInternalOccmDemandSignalsSortByTimecreated ListInternalOccmDemandSignalsSortByEnum = "timeCreated"
	ListInternalOccmDemandSignalsSortByDisplayname ListInternalOccmDemandSignalsSortByEnum = "displayName"
)

var mappingListInternalOccmDemandSignalsSortByEnum = map[string]ListInternalOccmDemandSignalsSortByEnum{
	"timeCreated": ListInternalOccmDemandSignalsSortByTimecreated,
	"displayName": ListInternalOccmDemandSignalsSortByDisplayname,
}

var mappingListInternalOccmDemandSignalsSortByEnumLowerCase = map[string]ListInternalOccmDemandSignalsSortByEnum{
	"timecreated": ListInternalOccmDemandSignalsSortByTimecreated,
	"displayname": ListInternalOccmDemandSignalsSortByDisplayname,
}

// GetListInternalOccmDemandSignalsSortByEnumValues Enumerates the set of values for ListInternalOccmDemandSignalsSortByEnum
func GetListInternalOccmDemandSignalsSortByEnumValues() []ListInternalOccmDemandSignalsSortByEnum {
	values := make([]ListInternalOccmDemandSignalsSortByEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalsSortByEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalsSortByEnum
func GetListInternalOccmDemandSignalsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListInternalOccmDemandSignalsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalsSortByEnum(val string) (ListInternalOccmDemandSignalsSortByEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
