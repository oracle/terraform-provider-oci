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

// ListMediaAssetDistributionChannelAttachmentsRequest wrapper for the ListMediaAssetDistributionChannelAttachments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaAssetDistributionChannelAttachments.go.html to see an example of how to use ListMediaAssetDistributionChannelAttachmentsRequest.
type ListMediaAssetDistributionChannelAttachmentsRequest struct {

	// Unique MediaAsset identifier
	MediaAssetId *string `mandatory:"true" contributesTo:"path" name:"mediaAssetId"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaAssetDistributionChannelAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListMediaAssetDistributionChannelAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique DistributionChannel identifier.
	DistributionChannelId *string `mandatory:"false" contributesTo:"query" name:"distributionChannelId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaAssetDistributionChannelAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaAssetDistributionChannelAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaAssetDistributionChannelAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaAssetDistributionChannelAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaAssetDistributionChannelAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMediaAssetDistributionChannelAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaAssetDistributionChannelAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaAssetDistributionChannelAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaAssetDistributionChannelAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaAssetDistributionChannelAttachmentsResponse wrapper for the ListMediaAssetDistributionChannelAttachments operation
type ListMediaAssetDistributionChannelAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaAssetDistributionChannelAttachmentCollection instances
	MediaAssetDistributionChannelAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaAssetDistributionChannelAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaAssetDistributionChannelAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaAssetDistributionChannelAttachmentsSortOrderEnum Enum with underlying type: string
type ListMediaAssetDistributionChannelAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaAssetDistributionChannelAttachmentsSortOrderEnum
const (
	ListMediaAssetDistributionChannelAttachmentsSortOrderAsc  ListMediaAssetDistributionChannelAttachmentsSortOrderEnum = "ASC"
	ListMediaAssetDistributionChannelAttachmentsSortOrderDesc ListMediaAssetDistributionChannelAttachmentsSortOrderEnum = "DESC"
)

var mappingListMediaAssetDistributionChannelAttachmentsSortOrderEnum = map[string]ListMediaAssetDistributionChannelAttachmentsSortOrderEnum{
	"ASC":  ListMediaAssetDistributionChannelAttachmentsSortOrderAsc,
	"DESC": ListMediaAssetDistributionChannelAttachmentsSortOrderDesc,
}

var mappingListMediaAssetDistributionChannelAttachmentsSortOrderEnumLowerCase = map[string]ListMediaAssetDistributionChannelAttachmentsSortOrderEnum{
	"asc":  ListMediaAssetDistributionChannelAttachmentsSortOrderAsc,
	"desc": ListMediaAssetDistributionChannelAttachmentsSortOrderDesc,
}

// GetListMediaAssetDistributionChannelAttachmentsSortOrderEnumValues Enumerates the set of values for ListMediaAssetDistributionChannelAttachmentsSortOrderEnum
func GetListMediaAssetDistributionChannelAttachmentsSortOrderEnumValues() []ListMediaAssetDistributionChannelAttachmentsSortOrderEnum {
	values := make([]ListMediaAssetDistributionChannelAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListMediaAssetDistributionChannelAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaAssetDistributionChannelAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaAssetDistributionChannelAttachmentsSortOrderEnum
func GetListMediaAssetDistributionChannelAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaAssetDistributionChannelAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaAssetDistributionChannelAttachmentsSortOrderEnum(val string) (ListMediaAssetDistributionChannelAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListMediaAssetDistributionChannelAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaAssetDistributionChannelAttachmentsSortByEnum Enum with underlying type: string
type ListMediaAssetDistributionChannelAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListMediaAssetDistributionChannelAttachmentsSortByEnum
const (
	ListMediaAssetDistributionChannelAttachmentsSortByMediaassetid          ListMediaAssetDistributionChannelAttachmentsSortByEnum = "mediaAssetId"
	ListMediaAssetDistributionChannelAttachmentsSortByDistributionchannelid ListMediaAssetDistributionChannelAttachmentsSortByEnum = "distributionChannelId"
	ListMediaAssetDistributionChannelAttachmentsSortByDisplayname           ListMediaAssetDistributionChannelAttachmentsSortByEnum = "displayName"
	ListMediaAssetDistributionChannelAttachmentsSortByVersion               ListMediaAssetDistributionChannelAttachmentsSortByEnum = "version"
)

var mappingListMediaAssetDistributionChannelAttachmentsSortByEnum = map[string]ListMediaAssetDistributionChannelAttachmentsSortByEnum{
	"mediaAssetId":          ListMediaAssetDistributionChannelAttachmentsSortByMediaassetid,
	"distributionChannelId": ListMediaAssetDistributionChannelAttachmentsSortByDistributionchannelid,
	"displayName":           ListMediaAssetDistributionChannelAttachmentsSortByDisplayname,
	"version":               ListMediaAssetDistributionChannelAttachmentsSortByVersion,
}

var mappingListMediaAssetDistributionChannelAttachmentsSortByEnumLowerCase = map[string]ListMediaAssetDistributionChannelAttachmentsSortByEnum{
	"mediaassetid":          ListMediaAssetDistributionChannelAttachmentsSortByMediaassetid,
	"distributionchannelid": ListMediaAssetDistributionChannelAttachmentsSortByDistributionchannelid,
	"displayname":           ListMediaAssetDistributionChannelAttachmentsSortByDisplayname,
	"version":               ListMediaAssetDistributionChannelAttachmentsSortByVersion,
}

// GetListMediaAssetDistributionChannelAttachmentsSortByEnumValues Enumerates the set of values for ListMediaAssetDistributionChannelAttachmentsSortByEnum
func GetListMediaAssetDistributionChannelAttachmentsSortByEnumValues() []ListMediaAssetDistributionChannelAttachmentsSortByEnum {
	values := make([]ListMediaAssetDistributionChannelAttachmentsSortByEnum, 0)
	for _, v := range mappingListMediaAssetDistributionChannelAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaAssetDistributionChannelAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListMediaAssetDistributionChannelAttachmentsSortByEnum
func GetListMediaAssetDistributionChannelAttachmentsSortByEnumStringValues() []string {
	return []string{
		"mediaAssetId",
		"distributionChannelId",
		"displayName",
		"version",
	}
}

// GetMappingListMediaAssetDistributionChannelAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaAssetDistributionChannelAttachmentsSortByEnum(val string) (ListMediaAssetDistributionChannelAttachmentsSortByEnum, bool) {
	enum, ok := mappingListMediaAssetDistributionChannelAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
