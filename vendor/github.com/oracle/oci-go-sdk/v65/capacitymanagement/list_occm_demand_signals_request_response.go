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

// ListOccmDemandSignalsRequest wrapper for the ListOccmDemandSignals operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignals.go.html to see an example of how to use ListOccmDemandSignalsRequest.
type ListOccmDemandSignalsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortOrder ListOccmDemandSignalsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand Signal API. Only one sort order may be provided. The default order for timeCreated is reverse chronological order (latest date at the top). The default order for displayName is case sensitive alphabetical order.
	SortBy ListOccmDemandSignalsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccmDemandSignalsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccmDemandSignalsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccmDemandSignalsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccmDemandSignalsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccmDemandSignalsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccmDemandSignalsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccmDemandSignalsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccmDemandSignalsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccmDemandSignalsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccmDemandSignalsResponse wrapper for the ListOccmDemandSignals operation
type ListOccmDemandSignalsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccmDemandSignalCollection instances
	OccmDemandSignalCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccmDemandSignalsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccmDemandSignalsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccmDemandSignalsSortOrderEnum Enum with underlying type: string
type ListOccmDemandSignalsSortOrderEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalsSortOrderEnum
const (
	ListOccmDemandSignalsSortOrderAsc  ListOccmDemandSignalsSortOrderEnum = "ASC"
	ListOccmDemandSignalsSortOrderDesc ListOccmDemandSignalsSortOrderEnum = "DESC"
)

var mappingListOccmDemandSignalsSortOrderEnum = map[string]ListOccmDemandSignalsSortOrderEnum{
	"ASC":  ListOccmDemandSignalsSortOrderAsc,
	"DESC": ListOccmDemandSignalsSortOrderDesc,
}

var mappingListOccmDemandSignalsSortOrderEnumLowerCase = map[string]ListOccmDemandSignalsSortOrderEnum{
	"asc":  ListOccmDemandSignalsSortOrderAsc,
	"desc": ListOccmDemandSignalsSortOrderDesc,
}

// GetListOccmDemandSignalsSortOrderEnumValues Enumerates the set of values for ListOccmDemandSignalsSortOrderEnum
func GetListOccmDemandSignalsSortOrderEnumValues() []ListOccmDemandSignalsSortOrderEnum {
	values := make([]ListOccmDemandSignalsSortOrderEnum, 0)
	for _, v := range mappingListOccmDemandSignalsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalsSortOrderEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalsSortOrderEnum
func GetListOccmDemandSignalsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccmDemandSignalsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalsSortOrderEnum(val string) (ListOccmDemandSignalsSortOrderEnum, bool) {
	enum, ok := mappingListOccmDemandSignalsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccmDemandSignalsSortByEnum Enum with underlying type: string
type ListOccmDemandSignalsSortByEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalsSortByEnum
const (
	ListOccmDemandSignalsSortByTimecreated ListOccmDemandSignalsSortByEnum = "timeCreated"
	ListOccmDemandSignalsSortByDisplayname ListOccmDemandSignalsSortByEnum = "displayName"
)

var mappingListOccmDemandSignalsSortByEnum = map[string]ListOccmDemandSignalsSortByEnum{
	"timeCreated": ListOccmDemandSignalsSortByTimecreated,
	"displayName": ListOccmDemandSignalsSortByDisplayname,
}

var mappingListOccmDemandSignalsSortByEnumLowerCase = map[string]ListOccmDemandSignalsSortByEnum{
	"timecreated": ListOccmDemandSignalsSortByTimecreated,
	"displayname": ListOccmDemandSignalsSortByDisplayname,
}

// GetListOccmDemandSignalsSortByEnumValues Enumerates the set of values for ListOccmDemandSignalsSortByEnum
func GetListOccmDemandSignalsSortByEnumValues() []ListOccmDemandSignalsSortByEnum {
	values := make([]ListOccmDemandSignalsSortByEnum, 0)
	for _, v := range mappingListOccmDemandSignalsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalsSortByEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalsSortByEnum
func GetListOccmDemandSignalsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOccmDemandSignalsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalsSortByEnum(val string) (ListOccmDemandSignalsSortByEnum, bool) {
	enum, ok := mappingListOccmDemandSignalsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
