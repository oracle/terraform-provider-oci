// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListServiceAttachmentsRequest wrapper for the ListServiceAttachments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListServiceAttachments.go.html to see an example of how to use ListServiceAttachmentsRequest.
type ListServiceAttachmentsRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter that returns all resources that match the specified lifecycle state.
	LifecycleState ServiceAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns all resources that match the specified lifecycle state.
	ServiceInstanceType ServiceAttachmentServiceInstanceTypeEnum `mandatory:"false" contributesTo:"query" name:"serviceInstanceType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListServiceAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListServiceAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceAttachmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetServiceAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingServiceAttachmentServiceInstanceTypeEnum(string(request.ServiceInstanceType)); !ok && request.ServiceInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceInstanceType: %s. Supported values are: %s.", request.ServiceInstanceType, strings.Join(GetServiceAttachmentServiceInstanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceAttachmentsResponse wrapper for the ListServiceAttachments operation
type ListServiceAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceAttachmentCollection instances
	ServiceAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceAttachmentsSortOrderEnum Enum with underlying type: string
type ListServiceAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceAttachmentsSortOrderEnum
const (
	ListServiceAttachmentsSortOrderAsc  ListServiceAttachmentsSortOrderEnum = "ASC"
	ListServiceAttachmentsSortOrderDesc ListServiceAttachmentsSortOrderEnum = "DESC"
)

var mappingListServiceAttachmentsSortOrderEnum = map[string]ListServiceAttachmentsSortOrderEnum{
	"ASC":  ListServiceAttachmentsSortOrderAsc,
	"DESC": ListServiceAttachmentsSortOrderDesc,
}

var mappingListServiceAttachmentsSortOrderEnumLowerCase = map[string]ListServiceAttachmentsSortOrderEnum{
	"asc":  ListServiceAttachmentsSortOrderAsc,
	"desc": ListServiceAttachmentsSortOrderDesc,
}

// GetListServiceAttachmentsSortOrderEnumValues Enumerates the set of values for ListServiceAttachmentsSortOrderEnum
func GetListServiceAttachmentsSortOrderEnumValues() []ListServiceAttachmentsSortOrderEnum {
	values := make([]ListServiceAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListServiceAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListServiceAttachmentsSortOrderEnum
func GetListServiceAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceAttachmentsSortOrderEnum(val string) (ListServiceAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListServiceAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceAttachmentsSortByEnum Enum with underlying type: string
type ListServiceAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListServiceAttachmentsSortByEnum
const (
	ListServiceAttachmentsSortByTimeCreated ListServiceAttachmentsSortByEnum = "TIME_CREATED"
	ListServiceAttachmentsSortByDisplayName ListServiceAttachmentsSortByEnum = "DISPLAY_NAME"
)

var mappingListServiceAttachmentsSortByEnum = map[string]ListServiceAttachmentsSortByEnum{
	"TIME_CREATED": ListServiceAttachmentsSortByTimeCreated,
	"DISPLAY_NAME": ListServiceAttachmentsSortByDisplayName,
}

var mappingListServiceAttachmentsSortByEnumLowerCase = map[string]ListServiceAttachmentsSortByEnum{
	"time_created": ListServiceAttachmentsSortByTimeCreated,
	"display_name": ListServiceAttachmentsSortByDisplayName,
}

// GetListServiceAttachmentsSortByEnumValues Enumerates the set of values for ListServiceAttachmentsSortByEnum
func GetListServiceAttachmentsSortByEnumValues() []ListServiceAttachmentsSortByEnum {
	values := make([]ListServiceAttachmentsSortByEnum, 0)
	for _, v := range mappingListServiceAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListServiceAttachmentsSortByEnum
func GetListServiceAttachmentsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListServiceAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceAttachmentsSortByEnum(val string) (ListServiceAttachmentsSortByEnum, bool) {
	enum, ok := mappingListServiceAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
