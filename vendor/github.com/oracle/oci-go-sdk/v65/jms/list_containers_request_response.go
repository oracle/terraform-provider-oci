// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListContainersRequest wrapper for the ListContainers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListContainers.go.html to see an example of how to use ListContainersRequest.
type ListContainersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The name of the application.
	ApplicationName *string `mandatory:"false" contributesTo:"query" name:"applicationName"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The security status of the Java Runtime.
	JreSecurityStatus ListContainersJreSecurityStatusEnum `mandatory:"false" contributesTo:"query" name:"jreSecurityStatus" omitEmpty:"true"`

	// If specified, only containers with a start time later than or equal to this parameter will be included in the response (formatted according to RFC3339).
	TimeStartedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStartedGreaterThanOrEqualTo"`

	// If specified, only containers with a start time earlier than or equal to this parameter will be included in the response (formatted according to RFC3339).
	TimeStartedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStartedLessThanOrEqualTo"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListContainersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the containers. Only one sort order can be provided.
	// Default order for _displayName_, _namespace_, _podName_, _applicationName_, or _jreVersion_ is **ascending**.
	// Default order for _timeStarted_ is **descending**.
	// If no value is specified _timeStarted_ is default.
	SortBy ListContainersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContainersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListContainersJreSecurityStatusEnum(string(request.JreSecurityStatus)); !ok && request.JreSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JreSecurityStatus: %s. Supported values are: %s.", request.JreSecurityStatus, strings.Join(GetListContainersJreSecurityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListContainersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListContainersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListContainersResponse wrapper for the ListContainers operation
type ListContainersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContainerCollection instances
	ContainerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListContainersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainersJreSecurityStatusEnum Enum with underlying type: string
type ListContainersJreSecurityStatusEnum string

// Set of constants representing the allowable values for ListContainersJreSecurityStatusEnum
const (
	ListContainersJreSecurityStatusEarlyAccess     ListContainersJreSecurityStatusEnum = "EARLY_ACCESS"
	ListContainersJreSecurityStatusUnknown         ListContainersJreSecurityStatusEnum = "UNKNOWN"
	ListContainersJreSecurityStatusUpToDate        ListContainersJreSecurityStatusEnum = "UP_TO_DATE"
	ListContainersJreSecurityStatusUpdateRequired  ListContainersJreSecurityStatusEnum = "UPDATE_REQUIRED"
	ListContainersJreSecurityStatusUpgradeRequired ListContainersJreSecurityStatusEnum = "UPGRADE_REQUIRED"
)

var mappingListContainersJreSecurityStatusEnum = map[string]ListContainersJreSecurityStatusEnum{
	"EARLY_ACCESS":     ListContainersJreSecurityStatusEarlyAccess,
	"UNKNOWN":          ListContainersJreSecurityStatusUnknown,
	"UP_TO_DATE":       ListContainersJreSecurityStatusUpToDate,
	"UPDATE_REQUIRED":  ListContainersJreSecurityStatusUpdateRequired,
	"UPGRADE_REQUIRED": ListContainersJreSecurityStatusUpgradeRequired,
}

var mappingListContainersJreSecurityStatusEnumLowerCase = map[string]ListContainersJreSecurityStatusEnum{
	"early_access":     ListContainersJreSecurityStatusEarlyAccess,
	"unknown":          ListContainersJreSecurityStatusUnknown,
	"up_to_date":       ListContainersJreSecurityStatusUpToDate,
	"update_required":  ListContainersJreSecurityStatusUpdateRequired,
	"upgrade_required": ListContainersJreSecurityStatusUpgradeRequired,
}

// GetListContainersJreSecurityStatusEnumValues Enumerates the set of values for ListContainersJreSecurityStatusEnum
func GetListContainersJreSecurityStatusEnumValues() []ListContainersJreSecurityStatusEnum {
	values := make([]ListContainersJreSecurityStatusEnum, 0)
	for _, v := range mappingListContainersJreSecurityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersJreSecurityStatusEnumStringValues Enumerates the set of values in String for ListContainersJreSecurityStatusEnum
func GetListContainersJreSecurityStatusEnumStringValues() []string {
	return []string{
		"EARLY_ACCESS",
		"UNKNOWN",
		"UP_TO_DATE",
		"UPDATE_REQUIRED",
		"UPGRADE_REQUIRED",
	}
}

// GetMappingListContainersJreSecurityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersJreSecurityStatusEnum(val string) (ListContainersJreSecurityStatusEnum, bool) {
	enum, ok := mappingListContainersJreSecurityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainersSortOrderEnum Enum with underlying type: string
type ListContainersSortOrderEnum string

// Set of constants representing the allowable values for ListContainersSortOrderEnum
const (
	ListContainersSortOrderAsc  ListContainersSortOrderEnum = "ASC"
	ListContainersSortOrderDesc ListContainersSortOrderEnum = "DESC"
)

var mappingListContainersSortOrderEnum = map[string]ListContainersSortOrderEnum{
	"ASC":  ListContainersSortOrderAsc,
	"DESC": ListContainersSortOrderDesc,
}

var mappingListContainersSortOrderEnumLowerCase = map[string]ListContainersSortOrderEnum{
	"asc":  ListContainersSortOrderAsc,
	"desc": ListContainersSortOrderDesc,
}

// GetListContainersSortOrderEnumValues Enumerates the set of values for ListContainersSortOrderEnum
func GetListContainersSortOrderEnumValues() []ListContainersSortOrderEnum {
	values := make([]ListContainersSortOrderEnum, 0)
	for _, v := range mappingListContainersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersSortOrderEnumStringValues Enumerates the set of values in String for ListContainersSortOrderEnum
func GetListContainersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListContainersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersSortOrderEnum(val string) (ListContainersSortOrderEnum, bool) {
	enum, ok := mappingListContainersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainersSortByEnum Enum with underlying type: string
type ListContainersSortByEnum string

// Set of constants representing the allowable values for ListContainersSortByEnum
const (
	ListContainersSortByDisplayname       ListContainersSortByEnum = "displayName"
	ListContainersSortByNamespace         ListContainersSortByEnum = "namespace"
	ListContainersSortByPodname           ListContainersSortByEnum = "podName"
	ListContainersSortByApplicationname   ListContainersSortByEnum = "applicationName"
	ListContainersSortByJreversion        ListContainersSortByEnum = "jreVersion"
	ListContainersSortByJresecuritystatus ListContainersSortByEnum = "jreSecurityStatus"
	ListContainersSortByTimestarted       ListContainersSortByEnum = "timeStarted"
)

var mappingListContainersSortByEnum = map[string]ListContainersSortByEnum{
	"displayName":       ListContainersSortByDisplayname,
	"namespace":         ListContainersSortByNamespace,
	"podName":           ListContainersSortByPodname,
	"applicationName":   ListContainersSortByApplicationname,
	"jreVersion":        ListContainersSortByJreversion,
	"jreSecurityStatus": ListContainersSortByJresecuritystatus,
	"timeStarted":       ListContainersSortByTimestarted,
}

var mappingListContainersSortByEnumLowerCase = map[string]ListContainersSortByEnum{
	"displayname":       ListContainersSortByDisplayname,
	"namespace":         ListContainersSortByNamespace,
	"podname":           ListContainersSortByPodname,
	"applicationname":   ListContainersSortByApplicationname,
	"jreversion":        ListContainersSortByJreversion,
	"jresecuritystatus": ListContainersSortByJresecuritystatus,
	"timestarted":       ListContainersSortByTimestarted,
}

// GetListContainersSortByEnumValues Enumerates the set of values for ListContainersSortByEnum
func GetListContainersSortByEnumValues() []ListContainersSortByEnum {
	values := make([]ListContainersSortByEnum, 0)
	for _, v := range mappingListContainersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersSortByEnumStringValues Enumerates the set of values in String for ListContainersSortByEnum
func GetListContainersSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"namespace",
		"podName",
		"applicationName",
		"jreVersion",
		"jreSecurityStatus",
		"timeStarted",
	}
}

// GetMappingListContainersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersSortByEnum(val string) (ListContainersSortByEnum, bool) {
	enum, ok := mappingListContainersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
