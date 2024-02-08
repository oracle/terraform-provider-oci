// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWebAppAccelerationsRequest wrapper for the ListWebAppAccelerations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/ListWebAppAccelerations.go.html to see an example of how to use ListWebAppAccelerationsRequest.
type ListWebAppAccelerationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the WebAppAcceleration with the given OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only the WebAppAcceleration with the given
	// OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of related WebAppAccelerationPolicy.
	WebAppAccelerationPolicyId *string `mandatory:"false" contributesTo:"query" name:"webAppAccelerationPolicyId"`

	// A filter to return only resources that match the given lifecycleState.
	LifecycleState []WebAppAccelerationLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListWebAppAccelerationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	// If no value is specified timeCreated is default.
	SortBy ListWebAppAccelerationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWebAppAccelerationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWebAppAccelerationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWebAppAccelerationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWebAppAccelerationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWebAppAccelerationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingWebAppAccelerationLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetWebAppAccelerationLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListWebAppAccelerationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWebAppAccelerationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWebAppAccelerationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWebAppAccelerationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWebAppAccelerationsResponse wrapper for the ListWebAppAccelerations operation
type ListWebAppAccelerationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WebAppAccelerationCollection instances
	WebAppAccelerationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWebAppAccelerationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWebAppAccelerationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWebAppAccelerationsSortOrderEnum Enum with underlying type: string
type ListWebAppAccelerationsSortOrderEnum string

// Set of constants representing the allowable values for ListWebAppAccelerationsSortOrderEnum
const (
	ListWebAppAccelerationsSortOrderAsc  ListWebAppAccelerationsSortOrderEnum = "ASC"
	ListWebAppAccelerationsSortOrderDesc ListWebAppAccelerationsSortOrderEnum = "DESC"
)

var mappingListWebAppAccelerationsSortOrderEnum = map[string]ListWebAppAccelerationsSortOrderEnum{
	"ASC":  ListWebAppAccelerationsSortOrderAsc,
	"DESC": ListWebAppAccelerationsSortOrderDesc,
}

var mappingListWebAppAccelerationsSortOrderEnumLowerCase = map[string]ListWebAppAccelerationsSortOrderEnum{
	"asc":  ListWebAppAccelerationsSortOrderAsc,
	"desc": ListWebAppAccelerationsSortOrderDesc,
}

// GetListWebAppAccelerationsSortOrderEnumValues Enumerates the set of values for ListWebAppAccelerationsSortOrderEnum
func GetListWebAppAccelerationsSortOrderEnumValues() []ListWebAppAccelerationsSortOrderEnum {
	values := make([]ListWebAppAccelerationsSortOrderEnum, 0)
	for _, v := range mappingListWebAppAccelerationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppAccelerationsSortOrderEnumStringValues Enumerates the set of values in String for ListWebAppAccelerationsSortOrderEnum
func GetListWebAppAccelerationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWebAppAccelerationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppAccelerationsSortOrderEnum(val string) (ListWebAppAccelerationsSortOrderEnum, bool) {
	enum, ok := mappingListWebAppAccelerationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWebAppAccelerationsSortByEnum Enum with underlying type: string
type ListWebAppAccelerationsSortByEnum string

// Set of constants representing the allowable values for ListWebAppAccelerationsSortByEnum
const (
	ListWebAppAccelerationsSortByTimecreated ListWebAppAccelerationsSortByEnum = "timeCreated"
	ListWebAppAccelerationsSortByDisplayname ListWebAppAccelerationsSortByEnum = "displayName"
)

var mappingListWebAppAccelerationsSortByEnum = map[string]ListWebAppAccelerationsSortByEnum{
	"timeCreated": ListWebAppAccelerationsSortByTimecreated,
	"displayName": ListWebAppAccelerationsSortByDisplayname,
}

var mappingListWebAppAccelerationsSortByEnumLowerCase = map[string]ListWebAppAccelerationsSortByEnum{
	"timecreated": ListWebAppAccelerationsSortByTimecreated,
	"displayname": ListWebAppAccelerationsSortByDisplayname,
}

// GetListWebAppAccelerationsSortByEnumValues Enumerates the set of values for ListWebAppAccelerationsSortByEnum
func GetListWebAppAccelerationsSortByEnumValues() []ListWebAppAccelerationsSortByEnum {
	values := make([]ListWebAppAccelerationsSortByEnum, 0)
	for _, v := range mappingListWebAppAccelerationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppAccelerationsSortByEnumStringValues Enumerates the set of values in String for ListWebAppAccelerationsSortByEnum
func GetListWebAppAccelerationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWebAppAccelerationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppAccelerationsSortByEnum(val string) (ListWebAppAccelerationsSortByEnum, bool) {
	enum, ok := mappingListWebAppAccelerationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
