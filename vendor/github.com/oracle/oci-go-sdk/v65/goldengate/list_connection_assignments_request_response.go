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

// ListConnectionAssignmentsRequest wrapper for the ListConnectionAssignments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListConnectionAssignments.go.html to see an example of how to use ListConnectionAssignmentsRequest.
type ListConnectionAssignmentsRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection.
	ConnectionId *string `mandatory:"false" contributesTo:"query" name:"connectionId"`

	// The name of the connection in the assignment (aliasName).
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only connection assignments having the 'lifecycleState' given.
	LifecycleState ConnectionAssignmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectionAssignmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListConnectionAssignmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectionAssignmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionAssignmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectionAssignmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionAssignmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectionAssignmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionAssignmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConnectionAssignmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionAssignmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectionAssignmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionAssignmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectionAssignmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectionAssignmentsResponse wrapper for the ListConnectionAssignments operation
type ListConnectionAssignmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionAssignmentCollection instances
	ConnectionAssignmentCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConnectionAssignmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionAssignmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionAssignmentsSortOrderEnum Enum with underlying type: string
type ListConnectionAssignmentsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionAssignmentsSortOrderEnum
const (
	ListConnectionAssignmentsSortOrderAsc  ListConnectionAssignmentsSortOrderEnum = "ASC"
	ListConnectionAssignmentsSortOrderDesc ListConnectionAssignmentsSortOrderEnum = "DESC"
)

var mappingListConnectionAssignmentsSortOrderEnum = map[string]ListConnectionAssignmentsSortOrderEnum{
	"ASC":  ListConnectionAssignmentsSortOrderAsc,
	"DESC": ListConnectionAssignmentsSortOrderDesc,
}

var mappingListConnectionAssignmentsSortOrderEnumLowerCase = map[string]ListConnectionAssignmentsSortOrderEnum{
	"asc":  ListConnectionAssignmentsSortOrderAsc,
	"desc": ListConnectionAssignmentsSortOrderDesc,
}

// GetListConnectionAssignmentsSortOrderEnumValues Enumerates the set of values for ListConnectionAssignmentsSortOrderEnum
func GetListConnectionAssignmentsSortOrderEnumValues() []ListConnectionAssignmentsSortOrderEnum {
	values := make([]ListConnectionAssignmentsSortOrderEnum, 0)
	for _, v := range mappingListConnectionAssignmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionAssignmentsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectionAssignmentsSortOrderEnum
func GetListConnectionAssignmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectionAssignmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionAssignmentsSortOrderEnum(val string) (ListConnectionAssignmentsSortOrderEnum, bool) {
	enum, ok := mappingListConnectionAssignmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionAssignmentsSortByEnum Enum with underlying type: string
type ListConnectionAssignmentsSortByEnum string

// Set of constants representing the allowable values for ListConnectionAssignmentsSortByEnum
const (
	ListConnectionAssignmentsSortByTimecreated ListConnectionAssignmentsSortByEnum = "timeCreated"
	ListConnectionAssignmentsSortByDisplayname ListConnectionAssignmentsSortByEnum = "displayName"
)

var mappingListConnectionAssignmentsSortByEnum = map[string]ListConnectionAssignmentsSortByEnum{
	"timeCreated": ListConnectionAssignmentsSortByTimecreated,
	"displayName": ListConnectionAssignmentsSortByDisplayname,
}

var mappingListConnectionAssignmentsSortByEnumLowerCase = map[string]ListConnectionAssignmentsSortByEnum{
	"timecreated": ListConnectionAssignmentsSortByTimecreated,
	"displayname": ListConnectionAssignmentsSortByDisplayname,
}

// GetListConnectionAssignmentsSortByEnumValues Enumerates the set of values for ListConnectionAssignmentsSortByEnum
func GetListConnectionAssignmentsSortByEnumValues() []ListConnectionAssignmentsSortByEnum {
	values := make([]ListConnectionAssignmentsSortByEnum, 0)
	for _, v := range mappingListConnectionAssignmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionAssignmentsSortByEnumStringValues Enumerates the set of values in String for ListConnectionAssignmentsSortByEnum
func GetListConnectionAssignmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConnectionAssignmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionAssignmentsSortByEnum(val string) (ListConnectionAssignmentsSortByEnum, bool) {
	enum, ok := mappingListConnectionAssignmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
