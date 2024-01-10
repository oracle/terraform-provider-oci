// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMediaAssetsRequest wrapper for the ListMediaAssets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaAssets.go.html to see an example of how to use ListMediaAssetsRequest.
type ListMediaAssetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources with lifecycleState matching the given lifecycleState.
	LifecycleState ListMediaAssetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListMediaAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique DistributionChannel identifier.
	DistributionChannelId *string `mandatory:"false" contributesTo:"query" name:"distributionChannelId"`

	// Unique MediaAsset identifier of the asset from which this asset is derived.
	ParentMediaAssetId *string `mandatory:"false" contributesTo:"query" name:"parentMediaAssetId"`

	// Unique MediaAsset identifier of the first asset upload.
	MasterMediaAssetId *string `mandatory:"false" contributesTo:"query" name:"masterMediaAssetId"`

	// Filter MediaAsset by the asset type.
	Type ListMediaAssetsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Filter MediaAsset by the bucket where the object is stored.
	BucketName *string `mandatory:"false" contributesTo:"query" name:"bucketName"`

	// Filter MediaAsset by the name of the object in object storage.
	ObjectName *string `mandatory:"false" contributesTo:"query" name:"objectName"`

	// The ID of the MediaWorkflowJob used to produce this asset, if this parameter is supplied then the workflow ID must also be supplied.
	MediaWorkflowJobId *string `mandatory:"false" contributesTo:"query" name:"mediaWorkflowJobId"`

	// The ID of the MediaWorkflow used to produce this asset.
	SourceMediaWorkflowId *string `mandatory:"false" contributesTo:"query" name:"sourceMediaWorkflowId"`

	// The version of the MediaWorkflow used to produce this asset.
	SourceMediaWorkflowVersion *int64 `mandatory:"false" contributesTo:"query" name:"sourceMediaWorkflowVersion"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMediaAssetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMediaAssetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaAssetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaAssetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaAssetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaAssetsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListMediaAssetsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaAssetsResponse wrapper for the ListMediaAssets operation
type ListMediaAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaAssetCollection instances
	MediaAssetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaAssetsLifecycleStateEnum Enum with underlying type: string
type ListMediaAssetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMediaAssetsLifecycleStateEnum
const (
	ListMediaAssetsLifecycleStateCreating ListMediaAssetsLifecycleStateEnum = "CREATING"
	ListMediaAssetsLifecycleStateUpdating ListMediaAssetsLifecycleStateEnum = "UPDATING"
	ListMediaAssetsLifecycleStateActive   ListMediaAssetsLifecycleStateEnum = "ACTIVE"
	ListMediaAssetsLifecycleStateDeleting ListMediaAssetsLifecycleStateEnum = "DELETING"
	ListMediaAssetsLifecycleStateDeleted  ListMediaAssetsLifecycleStateEnum = "DELETED"
	ListMediaAssetsLifecycleStateFailed   ListMediaAssetsLifecycleStateEnum = "FAILED"
)

var mappingListMediaAssetsLifecycleStateEnum = map[string]ListMediaAssetsLifecycleStateEnum{
	"CREATING": ListMediaAssetsLifecycleStateCreating,
	"UPDATING": ListMediaAssetsLifecycleStateUpdating,
	"ACTIVE":   ListMediaAssetsLifecycleStateActive,
	"DELETING": ListMediaAssetsLifecycleStateDeleting,
	"DELETED":  ListMediaAssetsLifecycleStateDeleted,
	"FAILED":   ListMediaAssetsLifecycleStateFailed,
}

var mappingListMediaAssetsLifecycleStateEnumLowerCase = map[string]ListMediaAssetsLifecycleStateEnum{
	"creating": ListMediaAssetsLifecycleStateCreating,
	"updating": ListMediaAssetsLifecycleStateUpdating,
	"active":   ListMediaAssetsLifecycleStateActive,
	"deleting": ListMediaAssetsLifecycleStateDeleting,
	"deleted":  ListMediaAssetsLifecycleStateDeleted,
	"failed":   ListMediaAssetsLifecycleStateFailed,
}

// GetListMediaAssetsLifecycleStateEnumValues Enumerates the set of values for ListMediaAssetsLifecycleStateEnum
func GetListMediaAssetsLifecycleStateEnumValues() []ListMediaAssetsLifecycleStateEnum {
	values := make([]ListMediaAssetsLifecycleStateEnum, 0)
	for _, v := range mappingListMediaAssetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaAssetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMediaAssetsLifecycleStateEnum
func GetListMediaAssetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListMediaAssetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaAssetsLifecycleStateEnum(val string) (ListMediaAssetsLifecycleStateEnum, bool) {
	enum, ok := mappingListMediaAssetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaAssetsSortOrderEnum Enum with underlying type: string
type ListMediaAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaAssetsSortOrderEnum
const (
	ListMediaAssetsSortOrderAsc  ListMediaAssetsSortOrderEnum = "ASC"
	ListMediaAssetsSortOrderDesc ListMediaAssetsSortOrderEnum = "DESC"
)

var mappingListMediaAssetsSortOrderEnum = map[string]ListMediaAssetsSortOrderEnum{
	"ASC":  ListMediaAssetsSortOrderAsc,
	"DESC": ListMediaAssetsSortOrderDesc,
}

var mappingListMediaAssetsSortOrderEnumLowerCase = map[string]ListMediaAssetsSortOrderEnum{
	"asc":  ListMediaAssetsSortOrderAsc,
	"desc": ListMediaAssetsSortOrderDesc,
}

// GetListMediaAssetsSortOrderEnumValues Enumerates the set of values for ListMediaAssetsSortOrderEnum
func GetListMediaAssetsSortOrderEnumValues() []ListMediaAssetsSortOrderEnum {
	values := make([]ListMediaAssetsSortOrderEnum, 0)
	for _, v := range mappingListMediaAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaAssetsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaAssetsSortOrderEnum
func GetListMediaAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaAssetsSortOrderEnum(val string) (ListMediaAssetsSortOrderEnum, bool) {
	enum, ok := mappingListMediaAssetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaAssetsSortByEnum Enum with underlying type: string
type ListMediaAssetsSortByEnum string

// Set of constants representing the allowable values for ListMediaAssetsSortByEnum
const (
	ListMediaAssetsSortByCompartmentid      ListMediaAssetsSortByEnum = "compartmentId"
	ListMediaAssetsSortByType               ListMediaAssetsSortByEnum = "type"
	ListMediaAssetsSortByLifecyclestate     ListMediaAssetsSortByEnum = "lifecycleState"
	ListMediaAssetsSortByParentmediaassetid ListMediaAssetsSortByEnum = "parentMediaAssetId"
	ListMediaAssetsSortByMastermediaassetid ListMediaAssetsSortByEnum = "masterMediaAssetId"
	ListMediaAssetsSortByDisplayname        ListMediaAssetsSortByEnum = "displayName"
	ListMediaAssetsSortByTimecreated        ListMediaAssetsSortByEnum = "timeCreated"
	ListMediaAssetsSortByTimeupdated        ListMediaAssetsSortByEnum = "timeUpdated"
)

var mappingListMediaAssetsSortByEnum = map[string]ListMediaAssetsSortByEnum{
	"compartmentId":      ListMediaAssetsSortByCompartmentid,
	"type":               ListMediaAssetsSortByType,
	"lifecycleState":     ListMediaAssetsSortByLifecyclestate,
	"parentMediaAssetId": ListMediaAssetsSortByParentmediaassetid,
	"masterMediaAssetId": ListMediaAssetsSortByMastermediaassetid,
	"displayName":        ListMediaAssetsSortByDisplayname,
	"timeCreated":        ListMediaAssetsSortByTimecreated,
	"timeUpdated":        ListMediaAssetsSortByTimeupdated,
}

var mappingListMediaAssetsSortByEnumLowerCase = map[string]ListMediaAssetsSortByEnum{
	"compartmentid":      ListMediaAssetsSortByCompartmentid,
	"type":               ListMediaAssetsSortByType,
	"lifecyclestate":     ListMediaAssetsSortByLifecyclestate,
	"parentmediaassetid": ListMediaAssetsSortByParentmediaassetid,
	"mastermediaassetid": ListMediaAssetsSortByMastermediaassetid,
	"displayname":        ListMediaAssetsSortByDisplayname,
	"timecreated":        ListMediaAssetsSortByTimecreated,
	"timeupdated":        ListMediaAssetsSortByTimeupdated,
}

// GetListMediaAssetsSortByEnumValues Enumerates the set of values for ListMediaAssetsSortByEnum
func GetListMediaAssetsSortByEnumValues() []ListMediaAssetsSortByEnum {
	values := make([]ListMediaAssetsSortByEnum, 0)
	for _, v := range mappingListMediaAssetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaAssetsSortByEnumStringValues Enumerates the set of values in String for ListMediaAssetsSortByEnum
func GetListMediaAssetsSortByEnumStringValues() []string {
	return []string{
		"compartmentId",
		"type",
		"lifecycleState",
		"parentMediaAssetId",
		"masterMediaAssetId",
		"displayName",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListMediaAssetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaAssetsSortByEnum(val string) (ListMediaAssetsSortByEnum, bool) {
	enum, ok := mappingListMediaAssetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaAssetsTypeEnum Enum with underlying type: string
type ListMediaAssetsTypeEnum string

// Set of constants representing the allowable values for ListMediaAssetsTypeEnum
const (
	ListMediaAssetsTypeAudio       ListMediaAssetsTypeEnum = "AUDIO"
	ListMediaAssetsTypeVideo       ListMediaAssetsTypeEnum = "VIDEO"
	ListMediaAssetsTypePlaylist    ListMediaAssetsTypeEnum = "PLAYLIST"
	ListMediaAssetsTypeImage       ListMediaAssetsTypeEnum = "IMAGE"
	ListMediaAssetsTypeCaptionFile ListMediaAssetsTypeEnum = "CAPTION_FILE"
	ListMediaAssetsTypeUnknown     ListMediaAssetsTypeEnum = "UNKNOWN"
)

var mappingListMediaAssetsTypeEnum = map[string]ListMediaAssetsTypeEnum{
	"AUDIO":        ListMediaAssetsTypeAudio,
	"VIDEO":        ListMediaAssetsTypeVideo,
	"PLAYLIST":     ListMediaAssetsTypePlaylist,
	"IMAGE":        ListMediaAssetsTypeImage,
	"CAPTION_FILE": ListMediaAssetsTypeCaptionFile,
	"UNKNOWN":      ListMediaAssetsTypeUnknown,
}

var mappingListMediaAssetsTypeEnumLowerCase = map[string]ListMediaAssetsTypeEnum{
	"audio":        ListMediaAssetsTypeAudio,
	"video":        ListMediaAssetsTypeVideo,
	"playlist":     ListMediaAssetsTypePlaylist,
	"image":        ListMediaAssetsTypeImage,
	"caption_file": ListMediaAssetsTypeCaptionFile,
	"unknown":      ListMediaAssetsTypeUnknown,
}

// GetListMediaAssetsTypeEnumValues Enumerates the set of values for ListMediaAssetsTypeEnum
func GetListMediaAssetsTypeEnumValues() []ListMediaAssetsTypeEnum {
	values := make([]ListMediaAssetsTypeEnum, 0)
	for _, v := range mappingListMediaAssetsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaAssetsTypeEnumStringValues Enumerates the set of values in String for ListMediaAssetsTypeEnum
func GetListMediaAssetsTypeEnumStringValues() []string {
	return []string{
		"AUDIO",
		"VIDEO",
		"PLAYLIST",
		"IMAGE",
		"CAPTION_FILE",
		"UNKNOWN",
	}
}

// GetMappingListMediaAssetsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaAssetsTypeEnum(val string) (ListMediaAssetsTypeEnum, bool) {
	enum, ok := mappingListMediaAssetsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
