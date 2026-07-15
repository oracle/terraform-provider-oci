// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInfrastructuresRequest wrapper for the ListInfrastructures operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListInfrastructures.go.html to see an example of how to use ListInfrastructuresRequest.
type ListInfrastructuresRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	// For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided,
	// it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the entire display name given. The match is case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that match the specified lifecycle state.
	LifecycleState []InfrastructureLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// The maximum number of items to return in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which you want to start retrieving results. This token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order that you want to use, which is either `ASC` or `DESC`.
	SortOrder ListInfrastructuresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which you want to sort. You can provide only one type of sort order. The default order for `timeCreated` is descending. The default order for `displayName` is ascending. If no value is specified, then `timeCreated` is the default.
	// When listing software images within the same `version`, using `sortBy=buildIdentifier` is recommended. `buildIdentifier` is a monotonically increasing, time-ordered string marker (yyyy-mm-dd-hh:mm:ss) stored with the image.
	SortBy ListInfrastructuresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInfrastructuresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInfrastructuresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInfrastructuresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInfrastructuresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInfrastructuresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingInfrastructureLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetInfrastructureLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListInfrastructuresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInfrastructuresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInfrastructuresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInfrastructuresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInfrastructuresResponse wrapper for the ListInfrastructures operation
type ListInfrastructuresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InfrastructureCollection instances
	InfrastructureCollection `presentIn:"body"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then it can mean that a partial list was returned. To obtain the next batch of items, include this value as the `page` parameter for your next GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInfrastructuresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInfrastructuresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInfrastructuresSortOrderEnum Enum with underlying type: string
type ListInfrastructuresSortOrderEnum string

// Set of constants representing the allowable values for ListInfrastructuresSortOrderEnum
const (
	ListInfrastructuresSortOrderAsc  ListInfrastructuresSortOrderEnum = "ASC"
	ListInfrastructuresSortOrderDesc ListInfrastructuresSortOrderEnum = "DESC"
)

var mappingListInfrastructuresSortOrderEnum = map[string]ListInfrastructuresSortOrderEnum{
	"ASC":  ListInfrastructuresSortOrderAsc,
	"DESC": ListInfrastructuresSortOrderDesc,
}

var mappingListInfrastructuresSortOrderEnumLowerCase = map[string]ListInfrastructuresSortOrderEnum{
	"asc":  ListInfrastructuresSortOrderAsc,
	"desc": ListInfrastructuresSortOrderDesc,
}

// GetListInfrastructuresSortOrderEnumValues Enumerates the set of values for ListInfrastructuresSortOrderEnum
func GetListInfrastructuresSortOrderEnumValues() []ListInfrastructuresSortOrderEnum {
	values := make([]ListInfrastructuresSortOrderEnum, 0)
	for _, v := range mappingListInfrastructuresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInfrastructuresSortOrderEnumStringValues Enumerates the set of values in String for ListInfrastructuresSortOrderEnum
func GetListInfrastructuresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInfrastructuresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInfrastructuresSortOrderEnum(val string) (ListInfrastructuresSortOrderEnum, bool) {
	enum, ok := mappingListInfrastructuresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInfrastructuresSortByEnum Enum with underlying type: string
type ListInfrastructuresSortByEnum string

// Set of constants representing the allowable values for ListInfrastructuresSortByEnum
const (
	ListInfrastructuresSortByTimecreated     ListInfrastructuresSortByEnum = "timeCreated"
	ListInfrastructuresSortByDisplayname     ListInfrastructuresSortByEnum = "displayName"
	ListInfrastructuresSortByBuildidentifier ListInfrastructuresSortByEnum = "buildIdentifier"
)

var mappingListInfrastructuresSortByEnum = map[string]ListInfrastructuresSortByEnum{
	"timeCreated":     ListInfrastructuresSortByTimecreated,
	"displayName":     ListInfrastructuresSortByDisplayname,
	"buildIdentifier": ListInfrastructuresSortByBuildidentifier,
}

var mappingListInfrastructuresSortByEnumLowerCase = map[string]ListInfrastructuresSortByEnum{
	"timecreated":     ListInfrastructuresSortByTimecreated,
	"displayname":     ListInfrastructuresSortByDisplayname,
	"buildidentifier": ListInfrastructuresSortByBuildidentifier,
}

// GetListInfrastructuresSortByEnumValues Enumerates the set of values for ListInfrastructuresSortByEnum
func GetListInfrastructuresSortByEnumValues() []ListInfrastructuresSortByEnum {
	values := make([]ListInfrastructuresSortByEnum, 0)
	for _, v := range mappingListInfrastructuresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInfrastructuresSortByEnumStringValues Enumerates the set of values in String for ListInfrastructuresSortByEnum
func GetListInfrastructuresSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"buildIdentifier",
	}
}

// GetMappingListInfrastructuresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInfrastructuresSortByEnum(val string) (ListInfrastructuresSortByEnum, bool) {
	enum, ok := mappingListInfrastructuresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
