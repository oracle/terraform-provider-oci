// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailTrackConfigsRequest wrapper for the ListEmailTrackConfigs operation
type ListEmailTrackConfigsRequest struct {

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource on which email tracking is configured.
	TrackConfigScopeId *string `mandatory:"false" contributesTo:"query" name:"trackConfigScopeId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailTrackConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListEmailTrackConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter returned list by specified lifecyclestate. This parameter is case-insensitive.
	LifecycleState EmailTrackConfigLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailTrackConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailTrackConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailTrackConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailTrackConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailTrackConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailTrackConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailTrackConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailTrackConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailTrackConfigsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailTrackConfigLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailTrackConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailTrackConfigsResponse wrapper for the ListEmailTrackConfigs operation
type ListEmailTrackConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailTrackConfigCollection instances
	EmailTrackConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEmailTrackConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailTrackConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailTrackConfigsSortOrderEnum Enum with underlying type: string
type ListEmailTrackConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailTrackConfigsSortOrderEnum
const (
	ListEmailTrackConfigsSortOrderAsc  ListEmailTrackConfigsSortOrderEnum = "ASC"
	ListEmailTrackConfigsSortOrderDesc ListEmailTrackConfigsSortOrderEnum = "DESC"
)

var mappingListEmailTrackConfigsSortOrderEnum = map[string]ListEmailTrackConfigsSortOrderEnum{
	"ASC":  ListEmailTrackConfigsSortOrderAsc,
	"DESC": ListEmailTrackConfigsSortOrderDesc,
}

var mappingListEmailTrackConfigsSortOrderEnumLowerCase = map[string]ListEmailTrackConfigsSortOrderEnum{
	"asc":  ListEmailTrackConfigsSortOrderAsc,
	"desc": ListEmailTrackConfigsSortOrderDesc,
}

// GetListEmailTrackConfigsSortOrderEnumValues Enumerates the set of values for ListEmailTrackConfigsSortOrderEnum
func GetListEmailTrackConfigsSortOrderEnumValues() []ListEmailTrackConfigsSortOrderEnum {
	values := make([]ListEmailTrackConfigsSortOrderEnum, 0)
	for _, v := range mappingListEmailTrackConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailTrackConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailTrackConfigsSortOrderEnum
func GetListEmailTrackConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailTrackConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailTrackConfigsSortOrderEnum(val string) (ListEmailTrackConfigsSortOrderEnum, bool) {
	enum, ok := mappingListEmailTrackConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailTrackConfigsSortByEnum Enum with underlying type: string
type ListEmailTrackConfigsSortByEnum string

// Set of constants representing the allowable values for ListEmailTrackConfigsSortByEnum
const (
	ListEmailTrackConfigsSortByTimecreated ListEmailTrackConfigsSortByEnum = "TIMECREATED"
	ListEmailTrackConfigsSortByDisplayname ListEmailTrackConfigsSortByEnum = "DISPLAYNAME"
)

var mappingListEmailTrackConfigsSortByEnum = map[string]ListEmailTrackConfigsSortByEnum{
	"TIMECREATED": ListEmailTrackConfigsSortByTimecreated,
	"DISPLAYNAME": ListEmailTrackConfigsSortByDisplayname,
}

var mappingListEmailTrackConfigsSortByEnumLowerCase = map[string]ListEmailTrackConfigsSortByEnum{
	"timecreated": ListEmailTrackConfigsSortByTimecreated,
	"displayname": ListEmailTrackConfigsSortByDisplayname,
}

// GetListEmailTrackConfigsSortByEnumValues Enumerates the set of values for ListEmailTrackConfigsSortByEnum
func GetListEmailTrackConfigsSortByEnumValues() []ListEmailTrackConfigsSortByEnum {
	values := make([]ListEmailTrackConfigsSortByEnum, 0)
	for _, v := range mappingListEmailTrackConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailTrackConfigsSortByEnumStringValues Enumerates the set of values in String for ListEmailTrackConfigsSortByEnum
func GetListEmailTrackConfigsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListEmailTrackConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailTrackConfigsSortByEnum(val string) (ListEmailTrackConfigsSortByEnum, bool) {
	enum, ok := mappingListEmailTrackConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
