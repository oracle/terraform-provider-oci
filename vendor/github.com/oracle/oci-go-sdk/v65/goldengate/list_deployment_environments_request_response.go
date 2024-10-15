// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentEnvironmentsRequest wrapper for the ListDeploymentEnvironments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentEnvironments.go.html to see an example of how to use ListDeploymentEnvironmentsRequest.
type ListDeploymentEnvironmentsRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListDeploymentEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentEnvironmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentEnvironmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentEnvironmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentEnvironmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentEnvironmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentEnvironmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentEnvironmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentEnvironmentsResponse wrapper for the ListDeploymentEnvironments operation
type ListDeploymentEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentEnvironmentCollection instances
	DeploymentEnvironmentCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentEnvironmentsSortOrderEnum Enum with underlying type: string
type ListDeploymentEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentEnvironmentsSortOrderEnum
const (
	ListDeploymentEnvironmentsSortOrderAsc  ListDeploymentEnvironmentsSortOrderEnum = "ASC"
	ListDeploymentEnvironmentsSortOrderDesc ListDeploymentEnvironmentsSortOrderEnum = "DESC"
)

var mappingListDeploymentEnvironmentsSortOrderEnum = map[string]ListDeploymentEnvironmentsSortOrderEnum{
	"ASC":  ListDeploymentEnvironmentsSortOrderAsc,
	"DESC": ListDeploymentEnvironmentsSortOrderDesc,
}

var mappingListDeploymentEnvironmentsSortOrderEnumLowerCase = map[string]ListDeploymentEnvironmentsSortOrderEnum{
	"asc":  ListDeploymentEnvironmentsSortOrderAsc,
	"desc": ListDeploymentEnvironmentsSortOrderDesc,
}

// GetListDeploymentEnvironmentsSortOrderEnumValues Enumerates the set of values for ListDeploymentEnvironmentsSortOrderEnum
func GetListDeploymentEnvironmentsSortOrderEnumValues() []ListDeploymentEnvironmentsSortOrderEnum {
	values := make([]ListDeploymentEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentEnvironmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentEnvironmentsSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentEnvironmentsSortOrderEnum
func GetListDeploymentEnvironmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentEnvironmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentEnvironmentsSortOrderEnum(val string) (ListDeploymentEnvironmentsSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentEnvironmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentEnvironmentsSortByEnum Enum with underlying type: string
type ListDeploymentEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentEnvironmentsSortByEnum
const (
	ListDeploymentEnvironmentsSortByTimecreated ListDeploymentEnvironmentsSortByEnum = "timeCreated"
	ListDeploymentEnvironmentsSortByDisplayname ListDeploymentEnvironmentsSortByEnum = "displayName"
)

var mappingListDeploymentEnvironmentsSortByEnum = map[string]ListDeploymentEnvironmentsSortByEnum{
	"timeCreated": ListDeploymentEnvironmentsSortByTimecreated,
	"displayName": ListDeploymentEnvironmentsSortByDisplayname,
}

var mappingListDeploymentEnvironmentsSortByEnumLowerCase = map[string]ListDeploymentEnvironmentsSortByEnum{
	"timecreated": ListDeploymentEnvironmentsSortByTimecreated,
	"displayname": ListDeploymentEnvironmentsSortByDisplayname,
}

// GetListDeploymentEnvironmentsSortByEnumValues Enumerates the set of values for ListDeploymentEnvironmentsSortByEnum
func GetListDeploymentEnvironmentsSortByEnumValues() []ListDeploymentEnvironmentsSortByEnum {
	values := make([]ListDeploymentEnvironmentsSortByEnum, 0)
	for _, v := range mappingListDeploymentEnvironmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentEnvironmentsSortByEnumStringValues Enumerates the set of values in String for ListDeploymentEnvironmentsSortByEnum
func GetListDeploymentEnvironmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentEnvironmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentEnvironmentsSortByEnum(val string) (ListDeploymentEnvironmentsSortByEnum, bool) {
	enum, ok := mappingListDeploymentEnvironmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
