// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListHostedApplicationsIamRequest wrapper for the ListHostedApplicationsIam operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListHostedApplicationsIam.go.html to see an example of how to use ListHostedApplicationsIamRequest.
type ListHostedApplicationsIamRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the hosted applications that their lifecycle state matches the given lifecycle state.
	LifecycleState HostedApplicationBaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListHostedApplicationsIamSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListHostedApplicationsIamSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHostedApplicationsIamRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHostedApplicationsIamRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHostedApplicationsIamRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHostedApplicationsIamRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHostedApplicationsIamRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostedApplicationBaseLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetHostedApplicationBaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostedApplicationsIamSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHostedApplicationsIamSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostedApplicationsIamSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHostedApplicationsIamSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHostedApplicationsIamResponse wrapper for the ListHostedApplicationsIam operation
type ListHostedApplicationsIamResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HostedApplicationCollection instances
	HostedApplicationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHostedApplicationsIamResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHostedApplicationsIamResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHostedApplicationsIamSortOrderEnum Enum with underlying type: string
type ListHostedApplicationsIamSortOrderEnum string

// Set of constants representing the allowable values for ListHostedApplicationsIamSortOrderEnum
const (
	ListHostedApplicationsIamSortOrderAsc  ListHostedApplicationsIamSortOrderEnum = "ASC"
	ListHostedApplicationsIamSortOrderDesc ListHostedApplicationsIamSortOrderEnum = "DESC"
)

var mappingListHostedApplicationsIamSortOrderEnum = map[string]ListHostedApplicationsIamSortOrderEnum{
	"ASC":  ListHostedApplicationsIamSortOrderAsc,
	"DESC": ListHostedApplicationsIamSortOrderDesc,
}

var mappingListHostedApplicationsIamSortOrderEnumLowerCase = map[string]ListHostedApplicationsIamSortOrderEnum{
	"asc":  ListHostedApplicationsIamSortOrderAsc,
	"desc": ListHostedApplicationsIamSortOrderDesc,
}

// GetListHostedApplicationsIamSortOrderEnumValues Enumerates the set of values for ListHostedApplicationsIamSortOrderEnum
func GetListHostedApplicationsIamSortOrderEnumValues() []ListHostedApplicationsIamSortOrderEnum {
	values := make([]ListHostedApplicationsIamSortOrderEnum, 0)
	for _, v := range mappingListHostedApplicationsIamSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostedApplicationsIamSortOrderEnumStringValues Enumerates the set of values in String for ListHostedApplicationsIamSortOrderEnum
func GetListHostedApplicationsIamSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHostedApplicationsIamSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostedApplicationsIamSortOrderEnum(val string) (ListHostedApplicationsIamSortOrderEnum, bool) {
	enum, ok := mappingListHostedApplicationsIamSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostedApplicationsIamSortByEnum Enum with underlying type: string
type ListHostedApplicationsIamSortByEnum string

// Set of constants representing the allowable values for ListHostedApplicationsIamSortByEnum
const (
	ListHostedApplicationsIamSortByTimecreated    ListHostedApplicationsIamSortByEnum = "timeCreated"
	ListHostedApplicationsIamSortByDisplayname    ListHostedApplicationsIamSortByEnum = "displayName"
	ListHostedApplicationsIamSortByLifecyclestate ListHostedApplicationsIamSortByEnum = "lifecycleState"
)

var mappingListHostedApplicationsIamSortByEnum = map[string]ListHostedApplicationsIamSortByEnum{
	"timeCreated":    ListHostedApplicationsIamSortByTimecreated,
	"displayName":    ListHostedApplicationsIamSortByDisplayname,
	"lifecycleState": ListHostedApplicationsIamSortByLifecyclestate,
}

var mappingListHostedApplicationsIamSortByEnumLowerCase = map[string]ListHostedApplicationsIamSortByEnum{
	"timecreated":    ListHostedApplicationsIamSortByTimecreated,
	"displayname":    ListHostedApplicationsIamSortByDisplayname,
	"lifecyclestate": ListHostedApplicationsIamSortByLifecyclestate,
}

// GetListHostedApplicationsIamSortByEnumValues Enumerates the set of values for ListHostedApplicationsIamSortByEnum
func GetListHostedApplicationsIamSortByEnumValues() []ListHostedApplicationsIamSortByEnum {
	values := make([]ListHostedApplicationsIamSortByEnum, 0)
	for _, v := range mappingListHostedApplicationsIamSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostedApplicationsIamSortByEnumStringValues Enumerates the set of values in String for ListHostedApplicationsIamSortByEnum
func GetListHostedApplicationsIamSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListHostedApplicationsIamSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostedApplicationsIamSortByEnum(val string) (ListHostedApplicationsIamSortByEnum, bool) {
	enum, ok := mappingListHostedApplicationsIamSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
