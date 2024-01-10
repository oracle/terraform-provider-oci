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

// ListWebAppAccelerationPoliciesRequest wrapper for the ListWebAppAccelerationPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waa/ListWebAppAccelerationPolicies.go.html to see an example of how to use ListWebAppAccelerationPoliciesRequest.
type ListWebAppAccelerationPoliciesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycleState.
	LifecycleState []WebAppAccelerationPolicyLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the WebAppAccelerationPolicy with the given OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListWebAppAccelerationPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	// If no value is specified timeCreated is default.
	SortBy ListWebAppAccelerationPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWebAppAccelerationPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWebAppAccelerationPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWebAppAccelerationPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWebAppAccelerationPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWebAppAccelerationPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingWebAppAccelerationPolicyLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetWebAppAccelerationPolicyLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListWebAppAccelerationPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWebAppAccelerationPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWebAppAccelerationPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWebAppAccelerationPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWebAppAccelerationPoliciesResponse wrapper for the ListWebAppAccelerationPolicies operation
type ListWebAppAccelerationPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WebAppAccelerationPolicyCollection instances
	WebAppAccelerationPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWebAppAccelerationPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWebAppAccelerationPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWebAppAccelerationPoliciesSortOrderEnum Enum with underlying type: string
type ListWebAppAccelerationPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListWebAppAccelerationPoliciesSortOrderEnum
const (
	ListWebAppAccelerationPoliciesSortOrderAsc  ListWebAppAccelerationPoliciesSortOrderEnum = "ASC"
	ListWebAppAccelerationPoliciesSortOrderDesc ListWebAppAccelerationPoliciesSortOrderEnum = "DESC"
)

var mappingListWebAppAccelerationPoliciesSortOrderEnum = map[string]ListWebAppAccelerationPoliciesSortOrderEnum{
	"ASC":  ListWebAppAccelerationPoliciesSortOrderAsc,
	"DESC": ListWebAppAccelerationPoliciesSortOrderDesc,
}

var mappingListWebAppAccelerationPoliciesSortOrderEnumLowerCase = map[string]ListWebAppAccelerationPoliciesSortOrderEnum{
	"asc":  ListWebAppAccelerationPoliciesSortOrderAsc,
	"desc": ListWebAppAccelerationPoliciesSortOrderDesc,
}

// GetListWebAppAccelerationPoliciesSortOrderEnumValues Enumerates the set of values for ListWebAppAccelerationPoliciesSortOrderEnum
func GetListWebAppAccelerationPoliciesSortOrderEnumValues() []ListWebAppAccelerationPoliciesSortOrderEnum {
	values := make([]ListWebAppAccelerationPoliciesSortOrderEnum, 0)
	for _, v := range mappingListWebAppAccelerationPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppAccelerationPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListWebAppAccelerationPoliciesSortOrderEnum
func GetListWebAppAccelerationPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWebAppAccelerationPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppAccelerationPoliciesSortOrderEnum(val string) (ListWebAppAccelerationPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListWebAppAccelerationPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWebAppAccelerationPoliciesSortByEnum Enum with underlying type: string
type ListWebAppAccelerationPoliciesSortByEnum string

// Set of constants representing the allowable values for ListWebAppAccelerationPoliciesSortByEnum
const (
	ListWebAppAccelerationPoliciesSortByTimecreated ListWebAppAccelerationPoliciesSortByEnum = "timeCreated"
	ListWebAppAccelerationPoliciesSortByDisplayname ListWebAppAccelerationPoliciesSortByEnum = "displayName"
)

var mappingListWebAppAccelerationPoliciesSortByEnum = map[string]ListWebAppAccelerationPoliciesSortByEnum{
	"timeCreated": ListWebAppAccelerationPoliciesSortByTimecreated,
	"displayName": ListWebAppAccelerationPoliciesSortByDisplayname,
}

var mappingListWebAppAccelerationPoliciesSortByEnumLowerCase = map[string]ListWebAppAccelerationPoliciesSortByEnum{
	"timecreated": ListWebAppAccelerationPoliciesSortByTimecreated,
	"displayname": ListWebAppAccelerationPoliciesSortByDisplayname,
}

// GetListWebAppAccelerationPoliciesSortByEnumValues Enumerates the set of values for ListWebAppAccelerationPoliciesSortByEnum
func GetListWebAppAccelerationPoliciesSortByEnumValues() []ListWebAppAccelerationPoliciesSortByEnum {
	values := make([]ListWebAppAccelerationPoliciesSortByEnum, 0)
	for _, v := range mappingListWebAppAccelerationPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppAccelerationPoliciesSortByEnumStringValues Enumerates the set of values in String for ListWebAppAccelerationPoliciesSortByEnum
func GetListWebAppAccelerationPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWebAppAccelerationPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppAccelerationPoliciesSortByEnum(val string) (ListWebAppAccelerationPoliciesSortByEnum, bool) {
	enum, ok := mappingListWebAppAccelerationPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
