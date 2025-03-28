// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbSystemsRequest wrapper for the ListDbSystems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListDbSystems.go.html to see an example of how to use ListDbSystemsRequest.
type ListDbSystemsRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If true, return only DB Systems with a HeatWave cluster attached, if false
	// return only DB Systems with no HeatWave cluster attached. If not
	// present, return all DB Systems.
	IsHeatWaveClusterAttached *bool `mandatory:"false" contributesTo:"query" name:"isHeatWaveClusterAttached"`

	// The DB System OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// DbSystem Lifecycle State
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The requested Configuration instance.
	ConfigurationId *string `mandatory:"false" contributesTo:"query" name:"configurationId"`

	// Filter instances if they are using the latest revision of the
	// Configuration they are associated with.
	IsUpToDate *bool `mandatory:"false" contributesTo:"query" name:"isUpToDate"`

	// Filter DB Systems by their Database Management configuration.
	DatabaseManagement []ListDbSystemsDatabaseManagementEnum `contributesTo:"query" name:"databaseManagement" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListDbSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListDbSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbSystemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbSystemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbSystemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbSystemLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.DatabaseManagement {
		if _, ok := GetMappingListDbSystemsDatabaseManagementEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagement: %s. Supported values are: %s.", val, strings.Join(GetListDbSystemsDatabaseManagementEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListDbSystemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbSystemsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbSystemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbSystemsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbSystemsResponse wrapper for the ListDbSystems operation
type ListDbSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbSystemSummary instances
	Items []DbSystemSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbSystemsDatabaseManagementEnum Enum with underlying type: string
type ListDbSystemsDatabaseManagementEnum string

// Set of constants representing the allowable values for ListDbSystemsDatabaseManagementEnum
const (
	ListDbSystemsDatabaseManagementEnabled  ListDbSystemsDatabaseManagementEnum = "ENABLED"
	ListDbSystemsDatabaseManagementDisabled ListDbSystemsDatabaseManagementEnum = "DISABLED"
)

var mappingListDbSystemsDatabaseManagementEnum = map[string]ListDbSystemsDatabaseManagementEnum{
	"ENABLED":  ListDbSystemsDatabaseManagementEnabled,
	"DISABLED": ListDbSystemsDatabaseManagementDisabled,
}

var mappingListDbSystemsDatabaseManagementEnumLowerCase = map[string]ListDbSystemsDatabaseManagementEnum{
	"enabled":  ListDbSystemsDatabaseManagementEnabled,
	"disabled": ListDbSystemsDatabaseManagementDisabled,
}

// GetListDbSystemsDatabaseManagementEnumValues Enumerates the set of values for ListDbSystemsDatabaseManagementEnum
func GetListDbSystemsDatabaseManagementEnumValues() []ListDbSystemsDatabaseManagementEnum {
	values := make([]ListDbSystemsDatabaseManagementEnum, 0)
	for _, v := range mappingListDbSystemsDatabaseManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemsDatabaseManagementEnumStringValues Enumerates the set of values in String for ListDbSystemsDatabaseManagementEnum
func GetListDbSystemsDatabaseManagementEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingListDbSystemsDatabaseManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemsDatabaseManagementEnum(val string) (ListDbSystemsDatabaseManagementEnum, bool) {
	enum, ok := mappingListDbSystemsDatabaseManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbSystemsSortByEnum Enum with underlying type: string
type ListDbSystemsSortByEnum string

// Set of constants representing the allowable values for ListDbSystemsSortByEnum
const (
	ListDbSystemsSortByDisplayname ListDbSystemsSortByEnum = "displayName"
	ListDbSystemsSortByTimecreated ListDbSystemsSortByEnum = "timeCreated"
)

var mappingListDbSystemsSortByEnum = map[string]ListDbSystemsSortByEnum{
	"displayName": ListDbSystemsSortByDisplayname,
	"timeCreated": ListDbSystemsSortByTimecreated,
}

var mappingListDbSystemsSortByEnumLowerCase = map[string]ListDbSystemsSortByEnum{
	"displayname": ListDbSystemsSortByDisplayname,
	"timecreated": ListDbSystemsSortByTimecreated,
}

// GetListDbSystemsSortByEnumValues Enumerates the set of values for ListDbSystemsSortByEnum
func GetListDbSystemsSortByEnumValues() []ListDbSystemsSortByEnum {
	values := make([]ListDbSystemsSortByEnum, 0)
	for _, v := range mappingListDbSystemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemsSortByEnumStringValues Enumerates the set of values in String for ListDbSystemsSortByEnum
func GetListDbSystemsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListDbSystemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemsSortByEnum(val string) (ListDbSystemsSortByEnum, bool) {
	enum, ok := mappingListDbSystemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbSystemsSortOrderEnum Enum with underlying type: string
type ListDbSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListDbSystemsSortOrderEnum
const (
	ListDbSystemsSortOrderAsc  ListDbSystemsSortOrderEnum = "ASC"
	ListDbSystemsSortOrderDesc ListDbSystemsSortOrderEnum = "DESC"
)

var mappingListDbSystemsSortOrderEnum = map[string]ListDbSystemsSortOrderEnum{
	"ASC":  ListDbSystemsSortOrderAsc,
	"DESC": ListDbSystemsSortOrderDesc,
}

var mappingListDbSystemsSortOrderEnumLowerCase = map[string]ListDbSystemsSortOrderEnum{
	"asc":  ListDbSystemsSortOrderAsc,
	"desc": ListDbSystemsSortOrderDesc,
}

// GetListDbSystemsSortOrderEnumValues Enumerates the set of values for ListDbSystemsSortOrderEnum
func GetListDbSystemsSortOrderEnumValues() []ListDbSystemsSortOrderEnum {
	values := make([]ListDbSystemsSortOrderEnum, 0)
	for _, v := range mappingListDbSystemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemsSortOrderEnumStringValues Enumerates the set of values in String for ListDbSystemsSortOrderEnum
func GetListDbSystemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbSystemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemsSortOrderEnum(val string) (ListDbSystemsSortOrderEnum, bool) {
	enum, ok := mappingListDbSystemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
