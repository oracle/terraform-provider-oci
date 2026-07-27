// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBlueGreenDeploymentsRequest wrapper for the ListBlueGreenDeployments operation
type ListBlueGreenDeploymentsRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Source DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) filter.
	SourceDbSystemId *string `mandatory:"false" contributesTo:"query" name:"sourceDbSystemId"`

	// Target DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) filter.
	TargetDbSystemId *string `mandatory:"false" contributesTo:"query" name:"targetDbSystemId"`

	// Filters deployments by display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filters deployments by lifecycle state.
	LifecycleState BlueGreenDeploymentSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListBlueGreenDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListBlueGreenDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBlueGreenDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBlueGreenDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBlueGreenDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBlueGreenDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBlueGreenDeploymentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBlueGreenDeploymentSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBlueGreenDeploymentSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBlueGreenDeploymentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBlueGreenDeploymentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBlueGreenDeploymentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBlueGreenDeploymentsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBlueGreenDeploymentsResponse wrapper for the ListBlueGreenDeployments operation
type ListBlueGreenDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BlueGreenDeploymentCollection instances
	BlueGreenDeploymentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBlueGreenDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBlueGreenDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBlueGreenDeploymentsSortByEnum Enum with underlying type: string
type ListBlueGreenDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListBlueGreenDeploymentsSortByEnum
const (
	ListBlueGreenDeploymentsSortByDisplayname ListBlueGreenDeploymentsSortByEnum = "displayName"
	ListBlueGreenDeploymentsSortByTimecreated ListBlueGreenDeploymentsSortByEnum = "timeCreated"
)

var mappingListBlueGreenDeploymentsSortByEnum = map[string]ListBlueGreenDeploymentsSortByEnum{
	"displayName": ListBlueGreenDeploymentsSortByDisplayname,
	"timeCreated": ListBlueGreenDeploymentsSortByTimecreated,
}

var mappingListBlueGreenDeploymentsSortByEnumLowerCase = map[string]ListBlueGreenDeploymentsSortByEnum{
	"displayname": ListBlueGreenDeploymentsSortByDisplayname,
	"timecreated": ListBlueGreenDeploymentsSortByTimecreated,
}

// GetListBlueGreenDeploymentsSortByEnumValues Enumerates the set of values for ListBlueGreenDeploymentsSortByEnum
func GetListBlueGreenDeploymentsSortByEnumValues() []ListBlueGreenDeploymentsSortByEnum {
	values := make([]ListBlueGreenDeploymentsSortByEnum, 0)
	for _, v := range mappingListBlueGreenDeploymentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBlueGreenDeploymentsSortByEnumStringValues Enumerates the set of values in String for ListBlueGreenDeploymentsSortByEnum
func GetListBlueGreenDeploymentsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListBlueGreenDeploymentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBlueGreenDeploymentsSortByEnum(val string) (ListBlueGreenDeploymentsSortByEnum, bool) {
	enum, ok := mappingListBlueGreenDeploymentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBlueGreenDeploymentsSortOrderEnum Enum with underlying type: string
type ListBlueGreenDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListBlueGreenDeploymentsSortOrderEnum
const (
	ListBlueGreenDeploymentsSortOrderAsc  ListBlueGreenDeploymentsSortOrderEnum = "ASC"
	ListBlueGreenDeploymentsSortOrderDesc ListBlueGreenDeploymentsSortOrderEnum = "DESC"
)

var mappingListBlueGreenDeploymentsSortOrderEnum = map[string]ListBlueGreenDeploymentsSortOrderEnum{
	"ASC":  ListBlueGreenDeploymentsSortOrderAsc,
	"DESC": ListBlueGreenDeploymentsSortOrderDesc,
}

var mappingListBlueGreenDeploymentsSortOrderEnumLowerCase = map[string]ListBlueGreenDeploymentsSortOrderEnum{
	"asc":  ListBlueGreenDeploymentsSortOrderAsc,
	"desc": ListBlueGreenDeploymentsSortOrderDesc,
}

// GetListBlueGreenDeploymentsSortOrderEnumValues Enumerates the set of values for ListBlueGreenDeploymentsSortOrderEnum
func GetListBlueGreenDeploymentsSortOrderEnumValues() []ListBlueGreenDeploymentsSortOrderEnum {
	values := make([]ListBlueGreenDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListBlueGreenDeploymentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBlueGreenDeploymentsSortOrderEnumStringValues Enumerates the set of values in String for ListBlueGreenDeploymentsSortOrderEnum
func GetListBlueGreenDeploymentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBlueGreenDeploymentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBlueGreenDeploymentsSortOrderEnum(val string) (ListBlueGreenDeploymentsSortOrderEnum, bool) {
	enum, ok := mappingListBlueGreenDeploymentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
