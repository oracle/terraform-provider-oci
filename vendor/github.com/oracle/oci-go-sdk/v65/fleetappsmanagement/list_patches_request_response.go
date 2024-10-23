// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchesRequest wrapper for the ListPatches operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListPatches.go.html to see an example of how to use ListPatchesRequest.
type ListPatchesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Product platformConfigurationId associated with the Patch.
	ProductId *string `mandatory:"false" contributesTo:"query" name:"productId"`

	// Product version
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// DefinedBy type.
	Type PatchTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Patch Type platformConfigurationId associated with the Patch.
	PatchTypeId *string `mandatory:"false" contributesTo:"query" name:"patchTypeId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// unique Patch identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Patch Released Date
	TimeReleasedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeReleasedGreaterThanOrEqualTo"`

	// Patch Released Date
	TimeReleasedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeReleasedLessThan"`

	// Filter patch based on compliance policy rules for the Product
	ShouldCompliancePolicyRulesBeApplied *bool `mandatory:"false" contributesTo:"query" name:"shouldCompliancePolicyRulesBeApplied"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the Patch.
	LifecycleState PatchLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetPatchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPatchLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchesResponse wrapper for the ListPatches operation
type ListPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchCollection instances
	PatchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchesSortByEnum Enum with underlying type: string
type ListPatchesSortByEnum string

// Set of constants representing the allowable values for ListPatchesSortByEnum
const (
	ListPatchesSortByTimecreated ListPatchesSortByEnum = "timeCreated"
	ListPatchesSortByDisplayname ListPatchesSortByEnum = "displayName"
)

var mappingListPatchesSortByEnum = map[string]ListPatchesSortByEnum{
	"timeCreated": ListPatchesSortByTimecreated,
	"displayName": ListPatchesSortByDisplayname,
}

var mappingListPatchesSortByEnumLowerCase = map[string]ListPatchesSortByEnum{
	"timecreated": ListPatchesSortByTimecreated,
	"displayname": ListPatchesSortByDisplayname,
}

// GetListPatchesSortByEnumValues Enumerates the set of values for ListPatchesSortByEnum
func GetListPatchesSortByEnumValues() []ListPatchesSortByEnum {
	values := make([]ListPatchesSortByEnum, 0)
	for _, v := range mappingListPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchesSortByEnumStringValues Enumerates the set of values in String for ListPatchesSortByEnum
func GetListPatchesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchesSortByEnum(val string) (ListPatchesSortByEnum, bool) {
	enum, ok := mappingListPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchesSortOrderEnum Enum with underlying type: string
type ListPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListPatchesSortOrderEnum
const (
	ListPatchesSortOrderAsc  ListPatchesSortOrderEnum = "ASC"
	ListPatchesSortOrderDesc ListPatchesSortOrderEnum = "DESC"
)

var mappingListPatchesSortOrderEnum = map[string]ListPatchesSortOrderEnum{
	"ASC":  ListPatchesSortOrderAsc,
	"DESC": ListPatchesSortOrderDesc,
}

var mappingListPatchesSortOrderEnumLowerCase = map[string]ListPatchesSortOrderEnum{
	"asc":  ListPatchesSortOrderAsc,
	"desc": ListPatchesSortOrderDesc,
}

// GetListPatchesSortOrderEnumValues Enumerates the set of values for ListPatchesSortOrderEnum
func GetListPatchesSortOrderEnumValues() []ListPatchesSortOrderEnum {
	values := make([]ListPatchesSortOrderEnum, 0)
	for _, v := range mappingListPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListPatchesSortOrderEnum
func GetListPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchesSortOrderEnum(val string) (ListPatchesSortOrderEnum, bool) {
	enum, ok := mappingListPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
