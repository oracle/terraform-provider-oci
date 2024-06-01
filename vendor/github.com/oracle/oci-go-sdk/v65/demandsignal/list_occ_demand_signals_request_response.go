// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package demandsignal

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOccDemandSignalsRequest wrapper for the ListOccDemandSignals operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/ListOccDemandSignals.go.html to see an example of how to use ListOccDemandSignalsRequest.
type ListOccDemandSignalsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState OccDemandSignalLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OccDemandSignal.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOccDemandSignalsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListOccDemandSignalsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccDemandSignalsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccDemandSignalsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccDemandSignalsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccDemandSignalsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccDemandSignalsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccDemandSignalLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOccDemandSignalLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccDemandSignalsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccDemandSignalsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccDemandSignalsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccDemandSignalsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccDemandSignalsResponse wrapper for the ListOccDemandSignals operation
type ListOccDemandSignalsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccDemandSignalCollection instances
	OccDemandSignalCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccDemandSignalsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccDemandSignalsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccDemandSignalsSortOrderEnum Enum with underlying type: string
type ListOccDemandSignalsSortOrderEnum string

// Set of constants representing the allowable values for ListOccDemandSignalsSortOrderEnum
const (
	ListOccDemandSignalsSortOrderAsc  ListOccDemandSignalsSortOrderEnum = "ASC"
	ListOccDemandSignalsSortOrderDesc ListOccDemandSignalsSortOrderEnum = "DESC"
)

var mappingListOccDemandSignalsSortOrderEnum = map[string]ListOccDemandSignalsSortOrderEnum{
	"ASC":  ListOccDemandSignalsSortOrderAsc,
	"DESC": ListOccDemandSignalsSortOrderDesc,
}

var mappingListOccDemandSignalsSortOrderEnumLowerCase = map[string]ListOccDemandSignalsSortOrderEnum{
	"asc":  ListOccDemandSignalsSortOrderAsc,
	"desc": ListOccDemandSignalsSortOrderDesc,
}

// GetListOccDemandSignalsSortOrderEnumValues Enumerates the set of values for ListOccDemandSignalsSortOrderEnum
func GetListOccDemandSignalsSortOrderEnumValues() []ListOccDemandSignalsSortOrderEnum {
	values := make([]ListOccDemandSignalsSortOrderEnum, 0)
	for _, v := range mappingListOccDemandSignalsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccDemandSignalsSortOrderEnumStringValues Enumerates the set of values in String for ListOccDemandSignalsSortOrderEnum
func GetListOccDemandSignalsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccDemandSignalsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccDemandSignalsSortOrderEnum(val string) (ListOccDemandSignalsSortOrderEnum, bool) {
	enum, ok := mappingListOccDemandSignalsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccDemandSignalsSortByEnum Enum with underlying type: string
type ListOccDemandSignalsSortByEnum string

// Set of constants representing the allowable values for ListOccDemandSignalsSortByEnum
const (
	ListOccDemandSignalsSortByTimecreated ListOccDemandSignalsSortByEnum = "timeCreated"
	ListOccDemandSignalsSortByDisplayname ListOccDemandSignalsSortByEnum = "displayName"
)

var mappingListOccDemandSignalsSortByEnum = map[string]ListOccDemandSignalsSortByEnum{
	"timeCreated": ListOccDemandSignalsSortByTimecreated,
	"displayName": ListOccDemandSignalsSortByDisplayname,
}

var mappingListOccDemandSignalsSortByEnumLowerCase = map[string]ListOccDemandSignalsSortByEnum{
	"timecreated": ListOccDemandSignalsSortByTimecreated,
	"displayname": ListOccDemandSignalsSortByDisplayname,
}

// GetListOccDemandSignalsSortByEnumValues Enumerates the set of values for ListOccDemandSignalsSortByEnum
func GetListOccDemandSignalsSortByEnumValues() []ListOccDemandSignalsSortByEnum {
	values := make([]ListOccDemandSignalsSortByEnum, 0)
	for _, v := range mappingListOccDemandSignalsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccDemandSignalsSortByEnumStringValues Enumerates the set of values in String for ListOccDemandSignalsSortByEnum
func GetListOccDemandSignalsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOccDemandSignalsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccDemandSignalsSortByEnum(val string) (ListOccDemandSignalsSortByEnum, bool) {
	enum, ok := mappingListOccDemandSignalsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
