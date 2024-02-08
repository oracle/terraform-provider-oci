// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModelVersionSetsRequest wrapper for the ListModelVersionSets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelVersionSets.go.html to see an example of how to use ListModelVersionSetsRequest.
type ListModelVersionSetsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListModelVersionSetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListModelVersionSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown in descending order.
	SortBy ListModelVersionSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelVersionSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelVersionSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelVersionSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelVersionSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelVersionSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModelVersionSetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListModelVersionSetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelVersionSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelVersionSetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelVersionSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelVersionSetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelVersionSetsResponse wrapper for the ListModelVersionSets operation
type ListModelVersionSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelVersionSetSummary instances
	Items []ModelVersionSetSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelVersionSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelVersionSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelVersionSetsLifecycleStateEnum Enum with underlying type: string
type ListModelVersionSetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelVersionSetsLifecycleStateEnum
const (
	ListModelVersionSetsLifecycleStateActive   ListModelVersionSetsLifecycleStateEnum = "ACTIVE"
	ListModelVersionSetsLifecycleStateDeleting ListModelVersionSetsLifecycleStateEnum = "DELETING"
	ListModelVersionSetsLifecycleStateDeleted  ListModelVersionSetsLifecycleStateEnum = "DELETED"
	ListModelVersionSetsLifecycleStateFailed   ListModelVersionSetsLifecycleStateEnum = "FAILED"
)

var mappingListModelVersionSetsLifecycleStateEnum = map[string]ListModelVersionSetsLifecycleStateEnum{
	"ACTIVE":   ListModelVersionSetsLifecycleStateActive,
	"DELETING": ListModelVersionSetsLifecycleStateDeleting,
	"DELETED":  ListModelVersionSetsLifecycleStateDeleted,
	"FAILED":   ListModelVersionSetsLifecycleStateFailed,
}

var mappingListModelVersionSetsLifecycleStateEnumLowerCase = map[string]ListModelVersionSetsLifecycleStateEnum{
	"active":   ListModelVersionSetsLifecycleStateActive,
	"deleting": ListModelVersionSetsLifecycleStateDeleting,
	"deleted":  ListModelVersionSetsLifecycleStateDeleted,
	"failed":   ListModelVersionSetsLifecycleStateFailed,
}

// GetListModelVersionSetsLifecycleStateEnumValues Enumerates the set of values for ListModelVersionSetsLifecycleStateEnum
func GetListModelVersionSetsLifecycleStateEnumValues() []ListModelVersionSetsLifecycleStateEnum {
	values := make([]ListModelVersionSetsLifecycleStateEnum, 0)
	for _, v := range mappingListModelVersionSetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelVersionSetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListModelVersionSetsLifecycleStateEnum
func GetListModelVersionSetsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListModelVersionSetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelVersionSetsLifecycleStateEnum(val string) (ListModelVersionSetsLifecycleStateEnum, bool) {
	enum, ok := mappingListModelVersionSetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelVersionSetsSortOrderEnum Enum with underlying type: string
type ListModelVersionSetsSortOrderEnum string

// Set of constants representing the allowable values for ListModelVersionSetsSortOrderEnum
const (
	ListModelVersionSetsSortOrderAsc  ListModelVersionSetsSortOrderEnum = "ASC"
	ListModelVersionSetsSortOrderDesc ListModelVersionSetsSortOrderEnum = "DESC"
)

var mappingListModelVersionSetsSortOrderEnum = map[string]ListModelVersionSetsSortOrderEnum{
	"ASC":  ListModelVersionSetsSortOrderAsc,
	"DESC": ListModelVersionSetsSortOrderDesc,
}

var mappingListModelVersionSetsSortOrderEnumLowerCase = map[string]ListModelVersionSetsSortOrderEnum{
	"asc":  ListModelVersionSetsSortOrderAsc,
	"desc": ListModelVersionSetsSortOrderDesc,
}

// GetListModelVersionSetsSortOrderEnumValues Enumerates the set of values for ListModelVersionSetsSortOrderEnum
func GetListModelVersionSetsSortOrderEnumValues() []ListModelVersionSetsSortOrderEnum {
	values := make([]ListModelVersionSetsSortOrderEnum, 0)
	for _, v := range mappingListModelVersionSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelVersionSetsSortOrderEnumStringValues Enumerates the set of values in String for ListModelVersionSetsSortOrderEnum
func GetListModelVersionSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelVersionSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelVersionSetsSortOrderEnum(val string) (ListModelVersionSetsSortOrderEnum, bool) {
	enum, ok := mappingListModelVersionSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelVersionSetsSortByEnum Enum with underlying type: string
type ListModelVersionSetsSortByEnum string

// Set of constants representing the allowable values for ListModelVersionSetsSortByEnum
const (
	ListModelVersionSetsSortByTimecreated    ListModelVersionSetsSortByEnum = "timeCreated"
	ListModelVersionSetsSortByName           ListModelVersionSetsSortByEnum = "name"
	ListModelVersionSetsSortByLifecyclestate ListModelVersionSetsSortByEnum = "lifecycleState"
)

var mappingListModelVersionSetsSortByEnum = map[string]ListModelVersionSetsSortByEnum{
	"timeCreated":    ListModelVersionSetsSortByTimecreated,
	"name":           ListModelVersionSetsSortByName,
	"lifecycleState": ListModelVersionSetsSortByLifecyclestate,
}

var mappingListModelVersionSetsSortByEnumLowerCase = map[string]ListModelVersionSetsSortByEnum{
	"timecreated":    ListModelVersionSetsSortByTimecreated,
	"name":           ListModelVersionSetsSortByName,
	"lifecyclestate": ListModelVersionSetsSortByLifecyclestate,
}

// GetListModelVersionSetsSortByEnumValues Enumerates the set of values for ListModelVersionSetsSortByEnum
func GetListModelVersionSetsSortByEnumValues() []ListModelVersionSetsSortByEnum {
	values := make([]ListModelVersionSetsSortByEnum, 0)
	for _, v := range mappingListModelVersionSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelVersionSetsSortByEnumStringValues Enumerates the set of values in String for ListModelVersionSetsSortByEnum
func GetListModelVersionSetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"lifecycleState",
	}
}

// GetMappingListModelVersionSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelVersionSetsSortByEnum(val string) (ListModelVersionSetsSortByEnum, bool) {
	enum, ok := mappingListModelVersionSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
