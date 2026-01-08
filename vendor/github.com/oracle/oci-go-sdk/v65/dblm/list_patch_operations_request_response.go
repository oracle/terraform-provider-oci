// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchOperationsRequest wrapper for the ListPatchOperations operation
type ListPatchOperationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState FppServerLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only those patch operations whose status matches the given status.
	Status PatchOperationStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique PatchOperation identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources whose timeStarted is greater than or equal to the given date-time.
	TimeStartedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStartedGreaterThanOrEqualTo"`

	// A filter to return only resources whose timeStarted is less than the given date-time.
	TimeStartedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStartedLessThan"`

	// A filter to return only resources whose isPrerequisitesOnly property value matches the given value.
	IsPrerequisitesOnly *bool `mandatory:"false" contributesTo:"query" name:"isPrerequisitesOnly"`

	// A filter to return only resources whose isQuickPrerequisitesCheck property value matches the given value.
	IsQuickPrerequisitesCheck *bool `mandatory:"false" contributesTo:"query" name:"isQuickPrerequisitesCheck"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPatchOperationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.
	SortBy ListPatchOperationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Deployment type to use in lists.
	DeploymentType ListPatchOperationsDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"deploymentType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchOperationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchOperationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchOperationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchOperationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchOperationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFppServerLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFppServerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchOperationStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetPatchOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchOperationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchOperationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchOperationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchOperationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchOperationsDeploymentTypeEnum(string(request.DeploymentType)); !ok && request.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", request.DeploymentType, strings.Join(GetListPatchOperationsDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchOperationsResponse wrapper for the ListPatchOperations operation
type ListPatchOperationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchOperationCollection instances
	PatchOperationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatchOperationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchOperationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchOperationsSortOrderEnum Enum with underlying type: string
type ListPatchOperationsSortOrderEnum string

// Set of constants representing the allowable values for ListPatchOperationsSortOrderEnum
const (
	ListPatchOperationsSortOrderAsc  ListPatchOperationsSortOrderEnum = "ASC"
	ListPatchOperationsSortOrderDesc ListPatchOperationsSortOrderEnum = "DESC"
)

var mappingListPatchOperationsSortOrderEnum = map[string]ListPatchOperationsSortOrderEnum{
	"ASC":  ListPatchOperationsSortOrderAsc,
	"DESC": ListPatchOperationsSortOrderDesc,
}

var mappingListPatchOperationsSortOrderEnumLowerCase = map[string]ListPatchOperationsSortOrderEnum{
	"asc":  ListPatchOperationsSortOrderAsc,
	"desc": ListPatchOperationsSortOrderDesc,
}

// GetListPatchOperationsSortOrderEnumValues Enumerates the set of values for ListPatchOperationsSortOrderEnum
func GetListPatchOperationsSortOrderEnumValues() []ListPatchOperationsSortOrderEnum {
	values := make([]ListPatchOperationsSortOrderEnum, 0)
	for _, v := range mappingListPatchOperationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchOperationsSortOrderEnumStringValues Enumerates the set of values in String for ListPatchOperationsSortOrderEnum
func GetListPatchOperationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchOperationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchOperationsSortOrderEnum(val string) (ListPatchOperationsSortOrderEnum, bool) {
	enum, ok := mappingListPatchOperationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchOperationsSortByEnum Enum with underlying type: string
type ListPatchOperationsSortByEnum string

// Set of constants representing the allowable values for ListPatchOperationsSortByEnum
const (
	ListPatchOperationsSortByTimecreated          ListPatchOperationsSortByEnum = "timeCreated"
	ListPatchOperationsSortByDisplayname          ListPatchOperationsSortByEnum = "displayName"
	ListPatchOperationsSortByTimestarted          ListPatchOperationsSortByEnum = "timeStarted"
	ListPatchOperationsSortByTimecompleted        ListPatchOperationsSortByEnum = "timeCompleted"
	ListPatchOperationsSortByTimeelapsedinseconds ListPatchOperationsSortByEnum = "timeElapsedInSeconds"
	ListPatchOperationsSortByStatus               ListPatchOperationsSortByEnum = "status"
	ListPatchOperationsSortByResources            ListPatchOperationsSortByEnum = "resources"
)

var mappingListPatchOperationsSortByEnum = map[string]ListPatchOperationsSortByEnum{
	"timeCreated":          ListPatchOperationsSortByTimecreated,
	"displayName":          ListPatchOperationsSortByDisplayname,
	"timeStarted":          ListPatchOperationsSortByTimestarted,
	"timeCompleted":        ListPatchOperationsSortByTimecompleted,
	"timeElapsedInSeconds": ListPatchOperationsSortByTimeelapsedinseconds,
	"status":               ListPatchOperationsSortByStatus,
	"resources":            ListPatchOperationsSortByResources,
}

var mappingListPatchOperationsSortByEnumLowerCase = map[string]ListPatchOperationsSortByEnum{
	"timecreated":          ListPatchOperationsSortByTimecreated,
	"displayname":          ListPatchOperationsSortByDisplayname,
	"timestarted":          ListPatchOperationsSortByTimestarted,
	"timecompleted":        ListPatchOperationsSortByTimecompleted,
	"timeelapsedinseconds": ListPatchOperationsSortByTimeelapsedinseconds,
	"status":               ListPatchOperationsSortByStatus,
	"resources":            ListPatchOperationsSortByResources,
}

// GetListPatchOperationsSortByEnumValues Enumerates the set of values for ListPatchOperationsSortByEnum
func GetListPatchOperationsSortByEnumValues() []ListPatchOperationsSortByEnum {
	values := make([]ListPatchOperationsSortByEnum, 0)
	for _, v := range mappingListPatchOperationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchOperationsSortByEnumStringValues Enumerates the set of values in String for ListPatchOperationsSortByEnum
func GetListPatchOperationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"timeStarted",
		"timeCompleted",
		"timeElapsedInSeconds",
		"status",
		"resources",
	}
}

// GetMappingListPatchOperationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchOperationsSortByEnum(val string) (ListPatchOperationsSortByEnum, bool) {
	enum, ok := mappingListPatchOperationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchOperationsDeploymentTypeEnum Enum with underlying type: string
type ListPatchOperationsDeploymentTypeEnum string

// Set of constants representing the allowable values for ListPatchOperationsDeploymentTypeEnum
const (
	ListPatchOperationsDeploymentTypeExternal ListPatchOperationsDeploymentTypeEnum = "EXTERNAL"
	ListPatchOperationsDeploymentTypeVm       ListPatchOperationsDeploymentTypeEnum = "VM"
)

var mappingListPatchOperationsDeploymentTypeEnum = map[string]ListPatchOperationsDeploymentTypeEnum{
	"EXTERNAL": ListPatchOperationsDeploymentTypeExternal,
	"VM":       ListPatchOperationsDeploymentTypeVm,
}

var mappingListPatchOperationsDeploymentTypeEnumLowerCase = map[string]ListPatchOperationsDeploymentTypeEnum{
	"external": ListPatchOperationsDeploymentTypeExternal,
	"vm":       ListPatchOperationsDeploymentTypeVm,
}

// GetListPatchOperationsDeploymentTypeEnumValues Enumerates the set of values for ListPatchOperationsDeploymentTypeEnum
func GetListPatchOperationsDeploymentTypeEnumValues() []ListPatchOperationsDeploymentTypeEnum {
	values := make([]ListPatchOperationsDeploymentTypeEnum, 0)
	for _, v := range mappingListPatchOperationsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchOperationsDeploymentTypeEnumStringValues Enumerates the set of values in String for ListPatchOperationsDeploymentTypeEnum
func GetListPatchOperationsDeploymentTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"VM",
	}
}

// GetMappingListPatchOperationsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchOperationsDeploymentTypeEnum(val string) (ListPatchOperationsDeploymentTypeEnum, bool) {
	enum, ok := mappingListPatchOperationsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
