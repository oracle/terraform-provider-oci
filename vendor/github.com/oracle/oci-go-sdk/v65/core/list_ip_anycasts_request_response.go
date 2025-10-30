// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIpAnycastsRequest wrapper for the ListIpAnycasts operation
type ListIpAnycastsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListIpAnycastsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListIpAnycastsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIpAnycastsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIpAnycastsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIpAnycastsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListIpAnycastsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["compartmentId"] != nil {
		templateParam := mandatoryParamMap["compartmentId"]
		for _, template := range templateParam {
			replacementParam := *request.CompartmentId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIpAnycastsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIpAnycastsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIpAnycastsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIpAnycastsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIpAnycastsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIpAnycastsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIpAnycastsResponse wrapper for the ListIpAnycasts operation
type ListIpAnycastsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IpAnycastCollection instances
	IpAnycastCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListIpAnycastsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIpAnycastsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIpAnycastsSortByEnum Enum with underlying type: string
type ListIpAnycastsSortByEnum string

// Set of constants representing the allowable values for ListIpAnycastsSortByEnum
const (
	ListIpAnycastsSortByTimecreated ListIpAnycastsSortByEnum = "TIMECREATED"
	ListIpAnycastsSortByDisplayname ListIpAnycastsSortByEnum = "DISPLAYNAME"
)

var mappingListIpAnycastsSortByEnum = map[string]ListIpAnycastsSortByEnum{
	"TIMECREATED": ListIpAnycastsSortByTimecreated,
	"DISPLAYNAME": ListIpAnycastsSortByDisplayname,
}

var mappingListIpAnycastsSortByEnumLowerCase = map[string]ListIpAnycastsSortByEnum{
	"timecreated": ListIpAnycastsSortByTimecreated,
	"displayname": ListIpAnycastsSortByDisplayname,
}

// GetListIpAnycastsSortByEnumValues Enumerates the set of values for ListIpAnycastsSortByEnum
func GetListIpAnycastsSortByEnumValues() []ListIpAnycastsSortByEnum {
	values := make([]ListIpAnycastsSortByEnum, 0)
	for _, v := range mappingListIpAnycastsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIpAnycastsSortByEnumStringValues Enumerates the set of values in String for ListIpAnycastsSortByEnum
func GetListIpAnycastsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListIpAnycastsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIpAnycastsSortByEnum(val string) (ListIpAnycastsSortByEnum, bool) {
	enum, ok := mappingListIpAnycastsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIpAnycastsSortOrderEnum Enum with underlying type: string
type ListIpAnycastsSortOrderEnum string

// Set of constants representing the allowable values for ListIpAnycastsSortOrderEnum
const (
	ListIpAnycastsSortOrderAsc  ListIpAnycastsSortOrderEnum = "ASC"
	ListIpAnycastsSortOrderDesc ListIpAnycastsSortOrderEnum = "DESC"
)

var mappingListIpAnycastsSortOrderEnum = map[string]ListIpAnycastsSortOrderEnum{
	"ASC":  ListIpAnycastsSortOrderAsc,
	"DESC": ListIpAnycastsSortOrderDesc,
}

var mappingListIpAnycastsSortOrderEnumLowerCase = map[string]ListIpAnycastsSortOrderEnum{
	"asc":  ListIpAnycastsSortOrderAsc,
	"desc": ListIpAnycastsSortOrderDesc,
}

// GetListIpAnycastsSortOrderEnumValues Enumerates the set of values for ListIpAnycastsSortOrderEnum
func GetListIpAnycastsSortOrderEnumValues() []ListIpAnycastsSortOrderEnum {
	values := make([]ListIpAnycastsSortOrderEnum, 0)
	for _, v := range mappingListIpAnycastsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIpAnycastsSortOrderEnumStringValues Enumerates the set of values in String for ListIpAnycastsSortOrderEnum
func GetListIpAnycastsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIpAnycastsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIpAnycastsSortOrderEnum(val string) (ListIpAnycastsSortOrderEnum, bool) {
	enum, ok := mappingListIpAnycastsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
