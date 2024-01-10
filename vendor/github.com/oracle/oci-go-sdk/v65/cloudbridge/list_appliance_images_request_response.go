// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApplianceImagesRequest wrapper for the ListApplianceImages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListApplianceImages.go.html to see an example of how to use ListApplianceImagesRequest.
type ListApplianceImagesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListApplianceImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListApplianceImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplianceImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplianceImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplianceImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplianceImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplianceImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplianceImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplianceImagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplianceImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplianceImagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplianceImagesResponse wrapper for the ListApplianceImages operation
type ListApplianceImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplianceImageCollection instances
	ApplianceImageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApplianceImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplianceImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplianceImagesSortOrderEnum Enum with underlying type: string
type ListApplianceImagesSortOrderEnum string

// Set of constants representing the allowable values for ListApplianceImagesSortOrderEnum
const (
	ListApplianceImagesSortOrderAsc  ListApplianceImagesSortOrderEnum = "ASC"
	ListApplianceImagesSortOrderDesc ListApplianceImagesSortOrderEnum = "DESC"
)

var mappingListApplianceImagesSortOrderEnum = map[string]ListApplianceImagesSortOrderEnum{
	"ASC":  ListApplianceImagesSortOrderAsc,
	"DESC": ListApplianceImagesSortOrderDesc,
}

var mappingListApplianceImagesSortOrderEnumLowerCase = map[string]ListApplianceImagesSortOrderEnum{
	"asc":  ListApplianceImagesSortOrderAsc,
	"desc": ListApplianceImagesSortOrderDesc,
}

// GetListApplianceImagesSortOrderEnumValues Enumerates the set of values for ListApplianceImagesSortOrderEnum
func GetListApplianceImagesSortOrderEnumValues() []ListApplianceImagesSortOrderEnum {
	values := make([]ListApplianceImagesSortOrderEnum, 0)
	for _, v := range mappingListApplianceImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplianceImagesSortOrderEnumStringValues Enumerates the set of values in String for ListApplianceImagesSortOrderEnum
func GetListApplianceImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplianceImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplianceImagesSortOrderEnum(val string) (ListApplianceImagesSortOrderEnum, bool) {
	enum, ok := mappingListApplianceImagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplianceImagesSortByEnum Enum with underlying type: string
type ListApplianceImagesSortByEnum string

// Set of constants representing the allowable values for ListApplianceImagesSortByEnum
const (
	ListApplianceImagesSortByTimecreated ListApplianceImagesSortByEnum = "timeCreated"
	ListApplianceImagesSortByTimeupdated ListApplianceImagesSortByEnum = "timeUpdated"
	ListApplianceImagesSortByDisplayname ListApplianceImagesSortByEnum = "displayName"
)

var mappingListApplianceImagesSortByEnum = map[string]ListApplianceImagesSortByEnum{
	"timeCreated": ListApplianceImagesSortByTimecreated,
	"timeUpdated": ListApplianceImagesSortByTimeupdated,
	"displayName": ListApplianceImagesSortByDisplayname,
}

var mappingListApplianceImagesSortByEnumLowerCase = map[string]ListApplianceImagesSortByEnum{
	"timecreated": ListApplianceImagesSortByTimecreated,
	"timeupdated": ListApplianceImagesSortByTimeupdated,
	"displayname": ListApplianceImagesSortByDisplayname,
}

// GetListApplianceImagesSortByEnumValues Enumerates the set of values for ListApplianceImagesSortByEnum
func GetListApplianceImagesSortByEnumValues() []ListApplianceImagesSortByEnum {
	values := make([]ListApplianceImagesSortByEnum, 0)
	for _, v := range mappingListApplianceImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplianceImagesSortByEnumStringValues Enumerates the set of values in String for ListApplianceImagesSortByEnum
func GetListApplianceImagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListApplianceImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplianceImagesSortByEnum(val string) (ListApplianceImagesSortByEnum, bool) {
	enum, ok := mappingListApplianceImagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
