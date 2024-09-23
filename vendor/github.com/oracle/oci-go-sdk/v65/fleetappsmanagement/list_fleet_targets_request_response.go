// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetTargetsRequest wrapper for the ListFleetTargets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListFleetTargets.go.html to see an example of how to use ListFleetTargetsRequest.
type ListFleetTargetsRequest struct {

	// unique Fleet identifier
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Product Name
	Product *string `mandatory:"false" contributesTo:"query" name:"product"`

	// Resource Identifier
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Resource Display Name
	ResourceDisplayName *string `mandatory:"false" contributesTo:"query" name:"resourceDisplayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFleetTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName,product and resourceDisplayName is ascending.
	SortBy ListFleetTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFleetTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetTargetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetTargetsResponse wrapper for the ListFleetTargets operation
type ListFleetTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetTargetCollection instances
	FleetTargetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetTargetsSortOrderEnum Enum with underlying type: string
type ListFleetTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetTargetsSortOrderEnum
const (
	ListFleetTargetsSortOrderAsc  ListFleetTargetsSortOrderEnum = "ASC"
	ListFleetTargetsSortOrderDesc ListFleetTargetsSortOrderEnum = "DESC"
)

var mappingListFleetTargetsSortOrderEnum = map[string]ListFleetTargetsSortOrderEnum{
	"ASC":  ListFleetTargetsSortOrderAsc,
	"DESC": ListFleetTargetsSortOrderDesc,
}

var mappingListFleetTargetsSortOrderEnumLowerCase = map[string]ListFleetTargetsSortOrderEnum{
	"asc":  ListFleetTargetsSortOrderAsc,
	"desc": ListFleetTargetsSortOrderDesc,
}

// GetListFleetTargetsSortOrderEnumValues Enumerates the set of values for ListFleetTargetsSortOrderEnum
func GetListFleetTargetsSortOrderEnumValues() []ListFleetTargetsSortOrderEnum {
	values := make([]ListFleetTargetsSortOrderEnum, 0)
	for _, v := range mappingListFleetTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListFleetTargetsSortOrderEnum
func GetListFleetTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetTargetsSortOrderEnum(val string) (ListFleetTargetsSortOrderEnum, bool) {
	enum, ok := mappingListFleetTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetTargetsSortByEnum Enum with underlying type: string
type ListFleetTargetsSortByEnum string

// Set of constants representing the allowable values for ListFleetTargetsSortByEnum
const (
	ListFleetTargetsSortByDisplayname         ListFleetTargetsSortByEnum = "displayName"
	ListFleetTargetsSortByProduct             ListFleetTargetsSortByEnum = "product"
	ListFleetTargetsSortByResourcedisplayname ListFleetTargetsSortByEnum = "resourceDisplayName"
)

var mappingListFleetTargetsSortByEnum = map[string]ListFleetTargetsSortByEnum{
	"displayName":         ListFleetTargetsSortByDisplayname,
	"product":             ListFleetTargetsSortByProduct,
	"resourceDisplayName": ListFleetTargetsSortByResourcedisplayname,
}

var mappingListFleetTargetsSortByEnumLowerCase = map[string]ListFleetTargetsSortByEnum{
	"displayname":         ListFleetTargetsSortByDisplayname,
	"product":             ListFleetTargetsSortByProduct,
	"resourcedisplayname": ListFleetTargetsSortByResourcedisplayname,
}

// GetListFleetTargetsSortByEnumValues Enumerates the set of values for ListFleetTargetsSortByEnum
func GetListFleetTargetsSortByEnumValues() []ListFleetTargetsSortByEnum {
	values := make([]ListFleetTargetsSortByEnum, 0)
	for _, v := range mappingListFleetTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetTargetsSortByEnumStringValues Enumerates the set of values in String for ListFleetTargetsSortByEnum
func GetListFleetTargetsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"product",
		"resourceDisplayName",
	}
}

// GetMappingListFleetTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetTargetsSortByEnum(val string) (ListFleetTargetsSortByEnum, bool) {
	enum, ok := mappingListFleetTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
