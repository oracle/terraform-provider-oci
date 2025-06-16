// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwareImagesRequest wrapper for the ListSoftwareImages operation
type ListSoftwareImagesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique SoftwareImage identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSoftwareImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSoftwareImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Patch Recommendation Status identifier
	PatchRecommendationStatus ListSoftwareImagesPatchRecommendationStatusEnum `mandatory:"false" contributesTo:"query" name:"patchRecommendationStatus" omitEmpty:"true"`

	// A filter to return only database that match the given release version.
	DatabaseRelease *string `mandatory:"false" contributesTo:"query" name:"databaseRelease"`

	// The version Key of the image version
	ImageVersionKey *int64 `mandatory:"false" contributesTo:"query" name:"imageVersionKey"`

	// Deployment type to use in lists.
	DeploymentType ListSoftwareImagesDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"deploymentType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSoftwareImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareImagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareImagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareImagesPatchRecommendationStatusEnum(string(request.PatchRecommendationStatus)); !ok && request.PatchRecommendationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchRecommendationStatus: %s. Supported values are: %s.", request.PatchRecommendationStatus, strings.Join(GetListSoftwareImagesPatchRecommendationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareImagesDeploymentTypeEnum(string(request.DeploymentType)); !ok && request.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", request.DeploymentType, strings.Join(GetListSoftwareImagesDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareImagesResponse wrapper for the ListSoftwareImages operation
type ListSoftwareImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwareImagesCollection instances
	SoftwareImagesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareImagesSortOrderEnum Enum with underlying type: string
type ListSoftwareImagesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareImagesSortOrderEnum
const (
	ListSoftwareImagesSortOrderAsc  ListSoftwareImagesSortOrderEnum = "ASC"
	ListSoftwareImagesSortOrderDesc ListSoftwareImagesSortOrderEnum = "DESC"
)

var mappingListSoftwareImagesSortOrderEnum = map[string]ListSoftwareImagesSortOrderEnum{
	"ASC":  ListSoftwareImagesSortOrderAsc,
	"DESC": ListSoftwareImagesSortOrderDesc,
}

var mappingListSoftwareImagesSortOrderEnumLowerCase = map[string]ListSoftwareImagesSortOrderEnum{
	"asc":  ListSoftwareImagesSortOrderAsc,
	"desc": ListSoftwareImagesSortOrderDesc,
}

// GetListSoftwareImagesSortOrderEnumValues Enumerates the set of values for ListSoftwareImagesSortOrderEnum
func GetListSoftwareImagesSortOrderEnumValues() []ListSoftwareImagesSortOrderEnum {
	values := make([]ListSoftwareImagesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareImagesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareImagesSortOrderEnum
func GetListSoftwareImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareImagesSortOrderEnum(val string) (ListSoftwareImagesSortOrderEnum, bool) {
	enum, ok := mappingListSoftwareImagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareImagesSortByEnum Enum with underlying type: string
type ListSoftwareImagesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareImagesSortByEnum
const (
	ListSoftwareImagesSortByTimecreated ListSoftwareImagesSortByEnum = "timeCreated"
	ListSoftwareImagesSortByDisplayname ListSoftwareImagesSortByEnum = "displayName"
)

var mappingListSoftwareImagesSortByEnum = map[string]ListSoftwareImagesSortByEnum{
	"timeCreated": ListSoftwareImagesSortByTimecreated,
	"displayName": ListSoftwareImagesSortByDisplayname,
}

var mappingListSoftwareImagesSortByEnumLowerCase = map[string]ListSoftwareImagesSortByEnum{
	"timecreated": ListSoftwareImagesSortByTimecreated,
	"displayname": ListSoftwareImagesSortByDisplayname,
}

// GetListSoftwareImagesSortByEnumValues Enumerates the set of values for ListSoftwareImagesSortByEnum
func GetListSoftwareImagesSortByEnumValues() []ListSoftwareImagesSortByEnum {
	values := make([]ListSoftwareImagesSortByEnum, 0)
	for _, v := range mappingListSoftwareImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareImagesSortByEnumStringValues Enumerates the set of values in String for ListSoftwareImagesSortByEnum
func GetListSoftwareImagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSoftwareImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareImagesSortByEnum(val string) (ListSoftwareImagesSortByEnum, bool) {
	enum, ok := mappingListSoftwareImagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareImagesPatchRecommendationStatusEnum Enum with underlying type: string
type ListSoftwareImagesPatchRecommendationStatusEnum string

// Set of constants representing the allowable values for ListSoftwareImagesPatchRecommendationStatusEnum
const (
	ListSoftwareImagesPatchRecommendationStatusAll            ListSoftwareImagesPatchRecommendationStatusEnum = "ALL"
	ListSoftwareImagesPatchRecommendationStatusUpToDate       ListSoftwareImagesPatchRecommendationStatusEnum = "UP_TO_DATE"
	ListSoftwareImagesPatchRecommendationStatusPatchAvailable ListSoftwareImagesPatchRecommendationStatusEnum = "PATCH_AVAILABLE"
)

var mappingListSoftwareImagesPatchRecommendationStatusEnum = map[string]ListSoftwareImagesPatchRecommendationStatusEnum{
	"ALL":             ListSoftwareImagesPatchRecommendationStatusAll,
	"UP_TO_DATE":      ListSoftwareImagesPatchRecommendationStatusUpToDate,
	"PATCH_AVAILABLE": ListSoftwareImagesPatchRecommendationStatusPatchAvailable,
}

var mappingListSoftwareImagesPatchRecommendationStatusEnumLowerCase = map[string]ListSoftwareImagesPatchRecommendationStatusEnum{
	"all":             ListSoftwareImagesPatchRecommendationStatusAll,
	"up_to_date":      ListSoftwareImagesPatchRecommendationStatusUpToDate,
	"patch_available": ListSoftwareImagesPatchRecommendationStatusPatchAvailable,
}

// GetListSoftwareImagesPatchRecommendationStatusEnumValues Enumerates the set of values for ListSoftwareImagesPatchRecommendationStatusEnum
func GetListSoftwareImagesPatchRecommendationStatusEnumValues() []ListSoftwareImagesPatchRecommendationStatusEnum {
	values := make([]ListSoftwareImagesPatchRecommendationStatusEnum, 0)
	for _, v := range mappingListSoftwareImagesPatchRecommendationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareImagesPatchRecommendationStatusEnumStringValues Enumerates the set of values in String for ListSoftwareImagesPatchRecommendationStatusEnum
func GetListSoftwareImagesPatchRecommendationStatusEnumStringValues() []string {
	return []string{
		"ALL",
		"UP_TO_DATE",
		"PATCH_AVAILABLE",
	}
}

// GetMappingListSoftwareImagesPatchRecommendationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareImagesPatchRecommendationStatusEnum(val string) (ListSoftwareImagesPatchRecommendationStatusEnum, bool) {
	enum, ok := mappingListSoftwareImagesPatchRecommendationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareImagesDeploymentTypeEnum Enum with underlying type: string
type ListSoftwareImagesDeploymentTypeEnum string

// Set of constants representing the allowable values for ListSoftwareImagesDeploymentTypeEnum
const (
	ListSoftwareImagesDeploymentTypeExternal ListSoftwareImagesDeploymentTypeEnum = "EXTERNAL"
	ListSoftwareImagesDeploymentTypeVm       ListSoftwareImagesDeploymentTypeEnum = "VM"
)

var mappingListSoftwareImagesDeploymentTypeEnum = map[string]ListSoftwareImagesDeploymentTypeEnum{
	"EXTERNAL": ListSoftwareImagesDeploymentTypeExternal,
	"VM":       ListSoftwareImagesDeploymentTypeVm,
}

var mappingListSoftwareImagesDeploymentTypeEnumLowerCase = map[string]ListSoftwareImagesDeploymentTypeEnum{
	"external": ListSoftwareImagesDeploymentTypeExternal,
	"vm":       ListSoftwareImagesDeploymentTypeVm,
}

// GetListSoftwareImagesDeploymentTypeEnumValues Enumerates the set of values for ListSoftwareImagesDeploymentTypeEnum
func GetListSoftwareImagesDeploymentTypeEnumValues() []ListSoftwareImagesDeploymentTypeEnum {
	values := make([]ListSoftwareImagesDeploymentTypeEnum, 0)
	for _, v := range mappingListSoftwareImagesDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareImagesDeploymentTypeEnumStringValues Enumerates the set of values in String for ListSoftwareImagesDeploymentTypeEnum
func GetListSoftwareImagesDeploymentTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"VM",
	}
}

// GetMappingListSoftwareImagesDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareImagesDeploymentTypeEnum(val string) (ListSoftwareImagesDeploymentTypeEnum, bool) {
	enum, ok := mappingListSoftwareImagesDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
