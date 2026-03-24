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

// ListHostedDeploymentsRequest wrapper for the ListHostedDeployments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListHostedDeployments.go.html to see an example of how to use ListHostedDeploymentsRequest.
type ListHostedDeploymentsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// A filter to return only the hosted deployments that their lifecycle state matches the given lifecycle state.
	LifecycleState HostedDeploymentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted deployment.
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
	SortOrder ListHostedDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListHostedDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHostedDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHostedDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHostedDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHostedDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHostedDeploymentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostedDeploymentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetHostedDeploymentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostedDeploymentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHostedDeploymentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostedDeploymentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHostedDeploymentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHostedDeploymentsResponse wrapper for the ListHostedDeployments operation
type ListHostedDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HostedDeploymentCollection instances
	HostedDeploymentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHostedDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHostedDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHostedDeploymentsSortOrderEnum Enum with underlying type: string
type ListHostedDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListHostedDeploymentsSortOrderEnum
const (
	ListHostedDeploymentsSortOrderAsc  ListHostedDeploymentsSortOrderEnum = "ASC"
	ListHostedDeploymentsSortOrderDesc ListHostedDeploymentsSortOrderEnum = "DESC"
)

var mappingListHostedDeploymentsSortOrderEnum = map[string]ListHostedDeploymentsSortOrderEnum{
	"ASC":  ListHostedDeploymentsSortOrderAsc,
	"DESC": ListHostedDeploymentsSortOrderDesc,
}

var mappingListHostedDeploymentsSortOrderEnumLowerCase = map[string]ListHostedDeploymentsSortOrderEnum{
	"asc":  ListHostedDeploymentsSortOrderAsc,
	"desc": ListHostedDeploymentsSortOrderDesc,
}

// GetListHostedDeploymentsSortOrderEnumValues Enumerates the set of values for ListHostedDeploymentsSortOrderEnum
func GetListHostedDeploymentsSortOrderEnumValues() []ListHostedDeploymentsSortOrderEnum {
	values := make([]ListHostedDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListHostedDeploymentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostedDeploymentsSortOrderEnumStringValues Enumerates the set of values in String for ListHostedDeploymentsSortOrderEnum
func GetListHostedDeploymentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHostedDeploymentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostedDeploymentsSortOrderEnum(val string) (ListHostedDeploymentsSortOrderEnum, bool) {
	enum, ok := mappingListHostedDeploymentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostedDeploymentsSortByEnum Enum with underlying type: string
type ListHostedDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListHostedDeploymentsSortByEnum
const (
	ListHostedDeploymentsSortByTimecreated    ListHostedDeploymentsSortByEnum = "timeCreated"
	ListHostedDeploymentsSortByDisplayname    ListHostedDeploymentsSortByEnum = "displayName"
	ListHostedDeploymentsSortByLifecyclestate ListHostedDeploymentsSortByEnum = "lifecycleState"
)

var mappingListHostedDeploymentsSortByEnum = map[string]ListHostedDeploymentsSortByEnum{
	"timeCreated":    ListHostedDeploymentsSortByTimecreated,
	"displayName":    ListHostedDeploymentsSortByDisplayname,
	"lifecycleState": ListHostedDeploymentsSortByLifecyclestate,
}

var mappingListHostedDeploymentsSortByEnumLowerCase = map[string]ListHostedDeploymentsSortByEnum{
	"timecreated":    ListHostedDeploymentsSortByTimecreated,
	"displayname":    ListHostedDeploymentsSortByDisplayname,
	"lifecyclestate": ListHostedDeploymentsSortByLifecyclestate,
}

// GetListHostedDeploymentsSortByEnumValues Enumerates the set of values for ListHostedDeploymentsSortByEnum
func GetListHostedDeploymentsSortByEnumValues() []ListHostedDeploymentsSortByEnum {
	values := make([]ListHostedDeploymentsSortByEnum, 0)
	for _, v := range mappingListHostedDeploymentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostedDeploymentsSortByEnumStringValues Enumerates the set of values in String for ListHostedDeploymentsSortByEnum
func GetListHostedDeploymentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListHostedDeploymentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostedDeploymentsSortByEnum(val string) (ListHostedDeploymentsSortByEnum, bool) {
	enum, ok := mappingListHostedDeploymentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
