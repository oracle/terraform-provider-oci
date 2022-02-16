// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListManagementAgentImagesRequest wrapper for the ListManagementAgentImages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListManagementAgentImages.go.html to see an example of how to use ListManagementAgentImagesRequest.
type ListManagementAgentImagesRequest struct {

	// The OCID of the compartment to which a request will be scoped.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for platformType is descending. Default order for version is descending. If no value is specified platformType is default.
	SortBy ListManagementAgentImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire platform name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentImagesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.
	InstallType ListManagementAgentImagesInstallTypeEnum `mandatory:"false" contributesTo:"query" name:"installType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementAgentImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementAgentImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementAgentImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementAgentImagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementAgentImagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentImagesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagementAgentImagesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentImagesInstallTypeEnum(string(request.InstallType)); !ok && request.InstallType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstallType: %s. Supported values are: %s.", request.InstallType, strings.Join(GetListManagementAgentImagesInstallTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementAgentImagesResponse wrapper for the ListManagementAgentImages operation
type ListManagementAgentImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentImageSummary instances
	Items []ManagementAgentImageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAgentImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentImagesSortOrderEnum Enum with underlying type: string
type ListManagementAgentImagesSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesSortOrderEnum
const (
	ListManagementAgentImagesSortOrderAsc  ListManagementAgentImagesSortOrderEnum = "ASC"
	ListManagementAgentImagesSortOrderDesc ListManagementAgentImagesSortOrderEnum = "DESC"
)

var mappingListManagementAgentImagesSortOrderEnum = map[string]ListManagementAgentImagesSortOrderEnum{
	"ASC":  ListManagementAgentImagesSortOrderAsc,
	"DESC": ListManagementAgentImagesSortOrderDesc,
}

// GetListManagementAgentImagesSortOrderEnumValues Enumerates the set of values for ListManagementAgentImagesSortOrderEnum
func GetListManagementAgentImagesSortOrderEnumValues() []ListManagementAgentImagesSortOrderEnum {
	values := make([]ListManagementAgentImagesSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentImagesSortOrderEnumStringValues Enumerates the set of values in String for ListManagementAgentImagesSortOrderEnum
func GetListManagementAgentImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementAgentImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentImagesSortOrderEnum(val string) (ListManagementAgentImagesSortOrderEnum, bool) {
	mappingListManagementAgentImagesSortOrderEnumIgnoreCase := make(map[string]ListManagementAgentImagesSortOrderEnum)
	for k, v := range mappingListManagementAgentImagesSortOrderEnum {
		mappingListManagementAgentImagesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListManagementAgentImagesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentImagesSortByEnum Enum with underlying type: string
type ListManagementAgentImagesSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesSortByEnum
const (
	ListManagementAgentImagesSortByPlatformtype ListManagementAgentImagesSortByEnum = "platformType"
	ListManagementAgentImagesSortByVersion      ListManagementAgentImagesSortByEnum = "version"
)

var mappingListManagementAgentImagesSortByEnum = map[string]ListManagementAgentImagesSortByEnum{
	"platformType": ListManagementAgentImagesSortByPlatformtype,
	"version":      ListManagementAgentImagesSortByVersion,
}

// GetListManagementAgentImagesSortByEnumValues Enumerates the set of values for ListManagementAgentImagesSortByEnum
func GetListManagementAgentImagesSortByEnumValues() []ListManagementAgentImagesSortByEnum {
	values := make([]ListManagementAgentImagesSortByEnum, 0)
	for _, v := range mappingListManagementAgentImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentImagesSortByEnumStringValues Enumerates the set of values in String for ListManagementAgentImagesSortByEnum
func GetListManagementAgentImagesSortByEnumStringValues() []string {
	return []string{
		"platformType",
		"version",
	}
}

// GetMappingListManagementAgentImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentImagesSortByEnum(val string) (ListManagementAgentImagesSortByEnum, bool) {
	mappingListManagementAgentImagesSortByEnumIgnoreCase := make(map[string]ListManagementAgentImagesSortByEnum)
	for k, v := range mappingListManagementAgentImagesSortByEnum {
		mappingListManagementAgentImagesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListManagementAgentImagesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentImagesLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentImagesLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesLifecycleStateEnum
const (
	ListManagementAgentImagesLifecycleStateCreating   ListManagementAgentImagesLifecycleStateEnum = "CREATING"
	ListManagementAgentImagesLifecycleStateUpdating   ListManagementAgentImagesLifecycleStateEnum = "UPDATING"
	ListManagementAgentImagesLifecycleStateActive     ListManagementAgentImagesLifecycleStateEnum = "ACTIVE"
	ListManagementAgentImagesLifecycleStateInactive   ListManagementAgentImagesLifecycleStateEnum = "INACTIVE"
	ListManagementAgentImagesLifecycleStateTerminated ListManagementAgentImagesLifecycleStateEnum = "TERMINATED"
	ListManagementAgentImagesLifecycleStateDeleting   ListManagementAgentImagesLifecycleStateEnum = "DELETING"
	ListManagementAgentImagesLifecycleStateDeleted    ListManagementAgentImagesLifecycleStateEnum = "DELETED"
	ListManagementAgentImagesLifecycleStateFailed     ListManagementAgentImagesLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentImagesLifecycleStateEnum = map[string]ListManagementAgentImagesLifecycleStateEnum{
	"CREATING":   ListManagementAgentImagesLifecycleStateCreating,
	"UPDATING":   ListManagementAgentImagesLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentImagesLifecycleStateActive,
	"INACTIVE":   ListManagementAgentImagesLifecycleStateInactive,
	"TERMINATED": ListManagementAgentImagesLifecycleStateTerminated,
	"DELETING":   ListManagementAgentImagesLifecycleStateDeleting,
	"DELETED":    ListManagementAgentImagesLifecycleStateDeleted,
	"FAILED":     ListManagementAgentImagesLifecycleStateFailed,
}

// GetListManagementAgentImagesLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentImagesLifecycleStateEnum
func GetListManagementAgentImagesLifecycleStateEnumValues() []ListManagementAgentImagesLifecycleStateEnum {
	values := make([]ListManagementAgentImagesLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentImagesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentImagesLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagementAgentImagesLifecycleStateEnum
func GetListManagementAgentImagesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"TERMINATED",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListManagementAgentImagesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentImagesLifecycleStateEnum(val string) (ListManagementAgentImagesLifecycleStateEnum, bool) {
	mappingListManagementAgentImagesLifecycleStateEnumIgnoreCase := make(map[string]ListManagementAgentImagesLifecycleStateEnum)
	for k, v := range mappingListManagementAgentImagesLifecycleStateEnum {
		mappingListManagementAgentImagesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListManagementAgentImagesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentImagesInstallTypeEnum Enum with underlying type: string
type ListManagementAgentImagesInstallTypeEnum string

// Set of constants representing the allowable values for ListManagementAgentImagesInstallTypeEnum
const (
	ListManagementAgentImagesInstallTypeAgent   ListManagementAgentImagesInstallTypeEnum = "AGENT"
	ListManagementAgentImagesInstallTypeGateway ListManagementAgentImagesInstallTypeEnum = "GATEWAY"
)

var mappingListManagementAgentImagesInstallTypeEnum = map[string]ListManagementAgentImagesInstallTypeEnum{
	"AGENT":   ListManagementAgentImagesInstallTypeAgent,
	"GATEWAY": ListManagementAgentImagesInstallTypeGateway,
}

// GetListManagementAgentImagesInstallTypeEnumValues Enumerates the set of values for ListManagementAgentImagesInstallTypeEnum
func GetListManagementAgentImagesInstallTypeEnumValues() []ListManagementAgentImagesInstallTypeEnum {
	values := make([]ListManagementAgentImagesInstallTypeEnum, 0)
	for _, v := range mappingListManagementAgentImagesInstallTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentImagesInstallTypeEnumStringValues Enumerates the set of values in String for ListManagementAgentImagesInstallTypeEnum
func GetListManagementAgentImagesInstallTypeEnumStringValues() []string {
	return []string{
		"AGENT",
		"GATEWAY",
	}
}

// GetMappingListManagementAgentImagesInstallTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentImagesInstallTypeEnum(val string) (ListManagementAgentImagesInstallTypeEnum, bool) {
	mappingListManagementAgentImagesInstallTypeEnumIgnoreCase := make(map[string]ListManagementAgentImagesInstallTypeEnum)
	for k, v := range mappingListManagementAgentImagesInstallTypeEnum {
		mappingListManagementAgentImagesInstallTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListManagementAgentImagesInstallTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
