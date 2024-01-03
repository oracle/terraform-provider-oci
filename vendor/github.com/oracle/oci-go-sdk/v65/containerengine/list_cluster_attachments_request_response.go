// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListClusterAttachmentsRequest wrapper for the ListClusterAttachments operation
type ListClusterAttachmentsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ClusterAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique ClusterAttachment identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The optional order in which to sort the results.
	SortOrder ListClusterAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListClusterAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClusterAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClusterAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClusterAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClusterAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListClusterAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterAttachmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetClusterAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListClusterAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListClusterAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListClusterAttachmentsResponse wrapper for the ListClusterAttachments operation
type ListClusterAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ClusterAttachmentCollection instances
	ClusterAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListClusterAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClusterAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClusterAttachmentsSortOrderEnum Enum with underlying type: string
type ListClusterAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListClusterAttachmentsSortOrderEnum
const (
	ListClusterAttachmentsSortOrderAsc  ListClusterAttachmentsSortOrderEnum = "ASC"
	ListClusterAttachmentsSortOrderDesc ListClusterAttachmentsSortOrderEnum = "DESC"
)

var mappingListClusterAttachmentsSortOrderEnum = map[string]ListClusterAttachmentsSortOrderEnum{
	"ASC":  ListClusterAttachmentsSortOrderAsc,
	"DESC": ListClusterAttachmentsSortOrderDesc,
}

var mappingListClusterAttachmentsSortOrderEnumLowerCase = map[string]ListClusterAttachmentsSortOrderEnum{
	"asc":  ListClusterAttachmentsSortOrderAsc,
	"desc": ListClusterAttachmentsSortOrderDesc,
}

// GetListClusterAttachmentsSortOrderEnumValues Enumerates the set of values for ListClusterAttachmentsSortOrderEnum
func GetListClusterAttachmentsSortOrderEnumValues() []ListClusterAttachmentsSortOrderEnum {
	values := make([]ListClusterAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListClusterAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListClusterAttachmentsSortOrderEnum
func GetListClusterAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListClusterAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterAttachmentsSortOrderEnum(val string) (ListClusterAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListClusterAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClusterAttachmentsSortByEnum Enum with underlying type: string
type ListClusterAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListClusterAttachmentsSortByEnum
const (
	ListClusterAttachmentsSortById          ListClusterAttachmentsSortByEnum = "ID"
	ListClusterAttachmentsSortByName        ListClusterAttachmentsSortByEnum = "NAME"
	ListClusterAttachmentsSortByTimeCreated ListClusterAttachmentsSortByEnum = "TIME_CREATED"
)

var mappingListClusterAttachmentsSortByEnum = map[string]ListClusterAttachmentsSortByEnum{
	"ID":           ListClusterAttachmentsSortById,
	"NAME":         ListClusterAttachmentsSortByName,
	"TIME_CREATED": ListClusterAttachmentsSortByTimeCreated,
}

var mappingListClusterAttachmentsSortByEnumLowerCase = map[string]ListClusterAttachmentsSortByEnum{
	"id":           ListClusterAttachmentsSortById,
	"name":         ListClusterAttachmentsSortByName,
	"time_created": ListClusterAttachmentsSortByTimeCreated,
}

// GetListClusterAttachmentsSortByEnumValues Enumerates the set of values for ListClusterAttachmentsSortByEnum
func GetListClusterAttachmentsSortByEnumValues() []ListClusterAttachmentsSortByEnum {
	values := make([]ListClusterAttachmentsSortByEnum, 0)
	for _, v := range mappingListClusterAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListClusterAttachmentsSortByEnum
func GetListClusterAttachmentsSortByEnumStringValues() []string {
	return []string{
		"ID",
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingListClusterAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterAttachmentsSortByEnum(val string) (ListClusterAttachmentsSortByEnum, bool) {
	enum, ok := mappingListClusterAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
