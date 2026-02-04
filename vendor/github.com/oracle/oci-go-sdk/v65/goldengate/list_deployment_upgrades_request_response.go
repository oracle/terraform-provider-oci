// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentUpgradesRequest wrapper for the ListDeploymentUpgrades operation
type ListDeploymentUpgradesRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// A filter to return only the deployment upgrades having the 'lifecycleState' given.
	LifecycleState DeploymentUpgradeLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentUpgradesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListDeploymentUpgradesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentUpgradesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentUpgradesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentUpgradesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentUpgradesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentUpgradesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentUpgradeLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDeploymentUpgradeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentUpgradesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentUpgradesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentUpgradesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentUpgradesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentUpgradesResponse wrapper for the ListDeploymentUpgrades operation
type ListDeploymentUpgradesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentUpgradeCollection instances
	DeploymentUpgradeCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentUpgradesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentUpgradesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentUpgradesSortOrderEnum Enum with underlying type: string
type ListDeploymentUpgradesSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentUpgradesSortOrderEnum
const (
	ListDeploymentUpgradesSortOrderAsc  ListDeploymentUpgradesSortOrderEnum = "ASC"
	ListDeploymentUpgradesSortOrderDesc ListDeploymentUpgradesSortOrderEnum = "DESC"
)

var mappingListDeploymentUpgradesSortOrderEnum = map[string]ListDeploymentUpgradesSortOrderEnum{
	"ASC":  ListDeploymentUpgradesSortOrderAsc,
	"DESC": ListDeploymentUpgradesSortOrderDesc,
}

var mappingListDeploymentUpgradesSortOrderEnumLowerCase = map[string]ListDeploymentUpgradesSortOrderEnum{
	"asc":  ListDeploymentUpgradesSortOrderAsc,
	"desc": ListDeploymentUpgradesSortOrderDesc,
}

// GetListDeploymentUpgradesSortOrderEnumValues Enumerates the set of values for ListDeploymentUpgradesSortOrderEnum
func GetListDeploymentUpgradesSortOrderEnumValues() []ListDeploymentUpgradesSortOrderEnum {
	values := make([]ListDeploymentUpgradesSortOrderEnum, 0)
	for _, v := range mappingListDeploymentUpgradesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentUpgradesSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentUpgradesSortOrderEnum
func GetListDeploymentUpgradesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentUpgradesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentUpgradesSortOrderEnum(val string) (ListDeploymentUpgradesSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentUpgradesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentUpgradesSortByEnum Enum with underlying type: string
type ListDeploymentUpgradesSortByEnum string

// Set of constants representing the allowable values for ListDeploymentUpgradesSortByEnum
const (
	ListDeploymentUpgradesSortByTimecreated ListDeploymentUpgradesSortByEnum = "timeCreated"
	ListDeploymentUpgradesSortByDisplayname ListDeploymentUpgradesSortByEnum = "displayName"
)

var mappingListDeploymentUpgradesSortByEnum = map[string]ListDeploymentUpgradesSortByEnum{
	"timeCreated": ListDeploymentUpgradesSortByTimecreated,
	"displayName": ListDeploymentUpgradesSortByDisplayname,
}

var mappingListDeploymentUpgradesSortByEnumLowerCase = map[string]ListDeploymentUpgradesSortByEnum{
	"timecreated": ListDeploymentUpgradesSortByTimecreated,
	"displayname": ListDeploymentUpgradesSortByDisplayname,
}

// GetListDeploymentUpgradesSortByEnumValues Enumerates the set of values for ListDeploymentUpgradesSortByEnum
func GetListDeploymentUpgradesSortByEnumValues() []ListDeploymentUpgradesSortByEnum {
	values := make([]ListDeploymentUpgradesSortByEnum, 0)
	for _, v := range mappingListDeploymentUpgradesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentUpgradesSortByEnumStringValues Enumerates the set of values in String for ListDeploymentUpgradesSortByEnum
func GetListDeploymentUpgradesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentUpgradesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentUpgradesSortByEnum(val string) (ListDeploymentUpgradesSortByEnum, bool) {
	enum, ok := mappingListDeploymentUpgradesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
